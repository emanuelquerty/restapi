package sqlite

import (
	"context"
	"database/sql"
	"fmt"
	"restapi/domain"
)

type UserStore struct {
	db *sql.DB
}

func NewUserStore(db *sql.DB) *UserStore {
	return &UserStore{
		db: db,
	}
}

func (u *UserStore) Create(ctx context.Context, user *domain.User) error {
	tx, err := u.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("error beggining transaction: %w", err)
	}
	defer tx.Rollback()

	query := "INSERT INTO users (first_name, last_name, email, password) VALUES (?, ?, ?, ?) RETURNING id"
	stmt, err := tx.PrepareContext(ctx, query)
	if err != nil {
		return fmt.Errorf("error preparing sql query: %w", err)
	}

	row := stmt.QueryRowContext(ctx, user.FirstName, user.LastName, user.Email, user.PasswordHash)
	err = row.Scan(&user.ID)
	if err != nil {
		return fmt.Errorf("error scanning sql row: %w", err)
	}

	return tx.Commit()
}

func (u *UserStore) Update(ctx context.Context, id int, updates map[string]any) (domain.User, error) {
	var user domain.User

	tx, err := u.db.BeginTx(ctx, nil)
	if err != nil {
		return user, fmt.Errorf("error beggining transaction: %w", err)
	}
	defer tx.Rollback()

	columnNames, columnValues := BuildUpdateColumns(updates)
	columnValues = append(columnValues, id)

	query := fmt.Sprintf(`UPDATE users SET %s WHERE id=? 
	RETURNING first_name, last_name, email`, columnNames)

	stmt, err := tx.PrepareContext(ctx, query)
	if err != nil {
		return user, fmt.Errorf("error preparing sql query: %w", err)
	}

	row := stmt.QueryRowContext(ctx, columnValues...)
	err = row.Scan(&user.FirstName, &user.LastName, &user.Email)
	if err != nil {
		return user, fmt.Errorf("error scanning sql row: %w", err)
	}
	user.ID = id

	return user, tx.Commit()
}

func (u *UserStore) FindByID(ctx context.Context, id int) (domain.User, error) {
	var user domain.User

	tx, err := u.db.BeginTx(ctx, nil)
	if err != nil {
		return user, fmt.Errorf("error beggining transaction: %w", err)
	}
	defer tx.Rollback()

	query := "SELECT first_name, last_name, email FROM users WHERE id=?"
	stmt, err := tx.PrepareContext(ctx, query)
	if err != nil {
		return user, fmt.Errorf("error preparing sql query: %w", err)
	}

	row := stmt.QueryRowContext(ctx, id)
	err = row.Scan(&user.FirstName, &user.LastName, &user.Email)
	if err != nil {
		return user, fmt.Errorf("error scanning sql row: %w", err)
	}

	user.ID = id
	return user, tx.Commit()
}

func (u *UserStore) Delete(ctx context.Context, id int) error {
	tx, err := u.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("error beggining transaction: %w", err)
	}
	defer tx.Rollback()

	query := "DELETE FROM users WHERE id=?"
	stmt, err := tx.PrepareContext(ctx, query)
	if err != nil {
		return fmt.Errorf("error preparing sql query: %w", err)
	}

	_, err = stmt.ExecContext(ctx, id)
	if err != nil {
		return fmt.Errorf("error executing sql statement: %w", err)
	}

	return tx.Commit()
}

func (u *UserStore) FindAll(ctx context.Context) ([]domain.User, error) {
	var users []domain.User

	tx, err := u.db.BeginTx(ctx, nil)
	if err != nil {
		return users, fmt.Errorf("error beggining transaction: %w", err)
	}
	defer tx.Rollback()

	query := "SELECT id, first_name, last_name, email FROM users"
	rows, err := tx.QueryContext(ctx, query)
	if err != nil {
		return users, fmt.Errorf("error executing sql query: %w", err)
	}

	for rows.Next() {
		var user domain.User

		err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email)
		if err != nil {
			return users, fmt.Errorf("error scanning sql row: %w", err)
		}
		users = append(users, user)
	}

	return users, tx.Commit()
}
