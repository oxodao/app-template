package config

import "fmt"

const DbProvider = "postgres"

type DB struct {
	Username   string
	Password   string
	Hostname   string
	Port       int
	Database   string
	CustomArgs string
}

func (db DB) GetDSN() string {
	dsn := fmt.Sprintf("%v://%v:%v@%v:%v/%v", DbProvider, db.Username, db.Password, db.Hostname, db.Port, db.Database)
	if len(db.CustomArgs) > 0 {
		dsn += "?" + db.CustomArgs
	}

	return dsn
}
