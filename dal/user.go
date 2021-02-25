package dal

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/oxodao/app-template/models"
)

type User struct {
	db *sqlx.DB
}

func (u *User) FindByUsername(username string) (*models.ExampleUser, error) {
	var user models.ExampleUser

	// Splitted to be SURE I don't accidentally create a SQL injection
	// Maybe I need to find a better way to do this
	rq := fmt.Sprintf(`
		SELECT USER_ID, USER_NAME 
		FROM %v
		WHERE USER_NAME = $1
	`, user.GetTableName())

	row := u.db.QueryRowx(rq, username)
	if row.Err() != nil {
		return nil, row.Err()
	}

	err := row.StructScan(&user)

	return &user, err
}
