package repo

import (
	"github.com/jackc/pgx/v4/pgxpool"
)

type Repo struct {
	*pgxpool.Conn
}

func New(pg *pgxpool.Conn) *Repo {
	return &Repo{pg}
}

func (r *Repo) addUser() {

}
