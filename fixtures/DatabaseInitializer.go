package fixtures

import (
	"fmt"
	"github.com/oxodao/app-template/log"
	"github.com/oxodao/app-template/models"
	"github.com/oxodao/app-template/services"
)

/**
 * This is specifically made for Postgresql
 * Requires to be rewritten if used with another DBMS
 */

func InitializeDB(prv *services.Provider, force bool) {
	// Install required plugins
	plugins := []string{
		"pgcrypto",
	}

	log.InfoAlways("Creating the database...")

	if !force {
		log.InfoAlways("Executing the creation script will remove ALL TABLES AND ALL DATA.")
		log.InfoAlways("Are you sure you want to proceed ? (y / N)")

		// @TODO read user input
		// if != Y & != y => os.exit(1)
	}

	for _, p := range plugins {
		_, err := prv.DB.Exec(fmt.Sprintf("CREATE EXTENSION IF NOT EXISTS %v;", p))
		if err != nil {
			panic(err)
		}
	}

	tables := []models.DatabaseModel{
		models.ExampleUser{},
	}

	// Dropping old tables
	for i := len(tables) - 1; i >= 0; i-- {
		_, e := prv.DB.Exec(`DROP TABLE IF EXISTS ` + tables[i].GetTableName())
		if e != nil {
			panic(e)
		}
	}

	// Re-creating tables
	for _, t := range tables {
		// No need to continue to run if there is an error
		_, e := prv.DB.Exec(t.GetCreationScript())
		if e != nil {
			panic(e)
		}

		log.InfoAlways(fmt.Sprintf("\t- Table %v created.", t.GetTableName()))
	}
}
