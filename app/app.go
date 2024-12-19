package app

import (
	"database/sql"
	"encoding/json"
	"log"
	"log/slog"
	"net/http"
	"restapi/domain"
	"restapi/storage/sqlite"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type appHandler func(w http.ResponseWriter, r *http.Request) *appError

func (fn appHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if e := fn(w, r); e != nil { // e is *appError, not os.Error
		log.Printf("%v", e.Error)
		err := json.NewEncoder(w).Encode(e)
		if err != nil {
			http.Error(w, e.Message, e.Code)
		}
	}
}

type middleware func(*slog.Logger, http.Handler) http.Handler

type App struct {
	Logger      *slog.Logger
	Mux         *http.ServeMux
	middlewares []middleware

	UserService domain.UserService
}

func New(logger *slog.Logger, db *sql.DB) *App {
	mux := http.NewServeMux()

	userStore := sqlite.NewUserStore(db)

	app := &App{
		Logger:      logger,
		Mux:         mux,
		UserService: userStore,
	}
	return app
}

func (a *App) Use(m middleware) {
	a.middlewares = append(a.middlewares, m)
}

func (a *App) Handle(pattern string, handler http.Handler) {
	finalHandler := handler
	for _, middleware := range a.middlewares {
		finalHandler = middleware(a.Logger, finalHandler)
	}

	a.Mux.Handle(pattern, finalHandler)
}

func (a *App) RegisterRoutes() {
	a.Mux.Handle("/metrics", promhttp.Handler())

	a.Handle("GET /api/health", appHandler(a.checkHealth))

	a.Handle("GET /api/users/{id}", appHandler(a.findUserByID))
	a.Handle("GET /api/users", appHandler(a.findAllUsers))

	a.Handle("POST /api/users", appHandler(a.createUser))

	a.Handle("PUT /api/users/{id}", appHandler(a.updateUser))

	a.Handle("DELETE /api/users/{id}", appHandler(a.deleteUser))

}

func (a *App) checkHealth(w http.ResponseWriter, r *http.Request) *appError {
	res := struct {
		Message string
	}{"Server is running"}

	err := json.NewEncoder(w).Encode(res)
	if err != nil {
		return &appError{err, "Invalid json body", http.StatusInternalServerError}
	}
	return nil
}

func (a *App) ListenAndServe(port string) error {
	a.RegisterRoutes()

	server := &http.Server{
		Addr:    ":" + port,
		Handler: a.Mux,
	}

	a.Logger.Info("server is running", slog.String("port", port))
	return server.ListenAndServe()
}
