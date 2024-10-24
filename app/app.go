package app

import (
	"database/sql"
	"log/slog"
	"net/http"
	"restapi/storage/sqlite"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type middleware func(*slog.Logger, http.Handler) http.Handler

type App struct {
	Logger      *slog.Logger
	Mux         *http.ServeMux
	middlewares []middleware

	userStore *sqlite.UserStore
}

func New(logger *slog.Logger, db *sql.DB) *App {
	mux := http.NewServeMux()

	userStore := sqlite.NewUserStore(db)

	app := &App{
		Logger:    logger,
		Mux:       mux,
		userStore: userStore,
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

	a.Handle("GET /api/health", http.HandlerFunc(a.checkHealth))
	a.Handle("/metrics", promhttp.Handler())

	a.Handle("GET /api/users/{id}", http.HandlerFunc(a.findUserByID))
	a.Handle("GET /api/users", http.HandlerFunc(a.findAllUsers))

	a.Handle("POST /api/users", http.HandlerFunc(a.createUser))

	a.Handle("PUT /api/users/{id}", http.HandlerFunc(a.updateUser))

	a.Handle("DELETE /api/users/{id}", http.HandlerFunc(a.deleteUser))

}

func (a *App) ListenAndServe(port string) error {
	server := &http.Server{
		Addr:    ":" + port,
		Handler: a.Mux,
	}

	a.Logger.Info("server is running", slog.String("port", port))
	return server.ListenAndServe()
}
