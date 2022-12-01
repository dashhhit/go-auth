package repo

import (
	"auth-golang/internal/entity"
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Repo struct {
	*pgxpool.Pool
}

func New(pg *pgxpool.Pool) *Repo {
	return &Repo{pg}
}

func (r *Repo) AddUser(ctx context.Context, user entity.User) error {
	sql := `INSERT INTO users (first_name, last_name, gender, email, password) VALUES ($1, $2, $3, $4, $5);`

	_, err := r.Exec(ctx, sql, user.FirstName, user.LastName, user.Gender, user.Email, user.Password)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repo) GetUser(ctx context.Context, id string) (*entity.User, error) {
	user := &entity.User{}

	sql := `SELECT id, first_name, last_name, gender, email FROM users WHERE id=$1;`

	row := r.QueryRow(ctx, sql, id)
	err := row.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Gender, &user.Email)
	if err != nil {
		return nil, err
	}

	return user, nil

}

func (r *Repo) IsExistsEmail(ctx context.Context, email string) bool {
	var (
		result int
	)
	sql := `SELECT DISTINCT 1 FROM users WHERE email=$1`
	row := r.QueryRow(ctx, sql, email)
	err := row.Scan(&result)
	if err != nil {
		return false
	}
	if result == 1 {
		return true
	}
	return false
}
