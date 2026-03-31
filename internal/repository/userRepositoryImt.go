package repository

import (
	"backend/internal/model"
	"context"
	"database/sql"
	"errors"
	"fmt"
)

type userPostgres struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userPostgres{db: db}
}

func (r *userPostgres) GetByID(ctx context.Context, id int) (*model.User, error) {
	query := `SELECT user_id, login FROM users WHERE user_id = $1`

	var user model.User
	err := r.db.QueryRowContext(ctx, query, id).Scan(&user.ID, &user.Login)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // пользователь не найден
		}
		return nil, fmt.Errorf("failed to get user by id %d: %w", id, err)
	}

	return &user, nil
}

func (r *userPostgres) GetAll(ctx context.Context) ([]*model.User, error) {
	query := `SELECT user_id, login FROM users ORDER BY user_id`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to get all users: %w", err)
	}
	defer rows.Close()

	var users []*model.User
	for rows.Next() {
		var user model.User
		if err := rows.Scan(&user.ID, &user.Login); err != nil {
			return nil, fmt.Errorf("failed to scan user: %w", err)
		}
		users = append(users, &user)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}

	return users, nil
}
