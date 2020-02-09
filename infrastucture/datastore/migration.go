package datastore

import (
	config "awawe/configuration"
	"database/sql"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"log"
)

func MigrateMySQLDatabase(db *sql.DB) {
	if db == nil {
		db = NewMySQLDB()
	}

	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	dbMigration, err := migrate.NewWithDatabaseInstance(
		fmt.Sprintf("file:%s", config.GetAppConfig().MigrationPath),
		"mysql",
		driver,
	)
	if err != nil {
		fmt.Errorf(err.Error())
	}

	if steps := config.GetAppConfig().MigrationSteps; steps == 0 {
		if err := dbMigration.Up(); err != nil {
			fmt.Errorf(err.Error())
		}
		fmt.Println("migrating database up")
	} else {
		if err := dbMigration.Steps(steps); err != nil {
			fmt.Errorf(err.Error())
		}
		fmt.Printf("migrating database for %d step(s)", steps)
	}
}
