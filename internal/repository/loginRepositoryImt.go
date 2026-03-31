package repository

import (
	"backend/internal/model"
	"context"
	"database/sql"
	_ "fmt"
)

type loginPostgres struct {
	db *sql.DB
}

func NewLoginRepository(db *sql.DB) LoginRepository {
	return &loginPostgres{db: db}
}

func (r *loginPostgres) LogIn(ctx context.Context, req model.LoginRequest) (model.LoginResponse, error) {

	responce := model.LoginResponse{}

	rows, err := r.db.Query(
		`SELECT user_id, role, name 
		FROM users
		WHERE login = $1 and 
		password = $2
		LIMIT 1`,
		req.Login,
		req.Password)

	if err != nil {
		return responce, err
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(
			&responce.Id,
			&responce.Role,
			&responce.Name,
		)
		if err != nil {
			return responce, err
		}
	}

	if err := rows.Err(); err != nil {
		return responce, err
	}

	return responce, nil
}
