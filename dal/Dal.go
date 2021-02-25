package dal

import "github.com/jmoiron/sqlx"

type Dal struct {
	User *User
}

func NewDal(db *sqlx.DB) *Dal {
	return &Dal {
		User: &User{ db },
	}
}
