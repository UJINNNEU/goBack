package login

import (
	"backend/internal/model"
	"context"
)

func (l *LoginRepository) SignIn(ctx context.Context, loginRequest model.LoginRequest) (model.LoginResponse, error) {
	loginResponse := model.LoginResponse{}

	query :=
		`SELECT	
		u.id,
		u.role,
		u.name
	FROM users u
	WHERE u.login = $1 AND u.password = $2
	LIMIT 1`

	rows, err := l.db.Query(query,
		loginRequest.Login,
		loginRequest.Password)

	if err != nil {
		defer rows.Close()
	}

	for rows.Next() {
		if err := rows.Scan(
			&loginResponse.Id,
			&loginResponse.Role,
			&loginResponse.Name); err != nil {
			return loginResponse, err
		}
	}

	return loginResponse, nil
}
