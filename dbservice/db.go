package dbservice

import (
	"database/sql"
	"log"
	"strconv"

	// we need to import mysql
	"github.com/qjouda/dignity-platform/backend/datatype"
	"github.com/qjouda/dignity-platform/backend/env"
	_ "github.com/go-sql-driver/mysql"
	migrate "github.com/rubenv/sql-migrate"
	"github.com/sirupsen/logrus"
)

// DB database connection type
type DB struct {
	*sql.DB
}

//MustConnect connects to database or panics
func MustConnect(dataSourceName string) *DB {
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		logrus.WithFields(
			logrus.Fields{"e": err.Error()},
		).Error("db:MustConnect:cant open sql")
		log.Fatal(err.Error())
	}

	// Open doesn't open a connection. Validate DSN data:
	err = db.Ping()
	if err != nil {
		logrus.WithFields(
			logrus.Fields{"e": err.Error()},
		).Error("db:MustConnect:db unreachable")
		log.Fatal(err.Error())
	}
	if err := migrateDB(db); err != nil {
		logrus.WithFields(
			logrus.Fields{"e": err.Error()},
		).Error("db:MustConnect:unable to migrate")
		log.Fatal("migrateDB: ", err)
	}
	return &DB{db}
}

// StringToID converts string to ID
func StringToID(str string) (id datatype.ID, isOK bool) {
	idInt, err := strconv.Atoi(str)
	if err != nil {
		return datatype.ID(0), false
	}
	return datatype.ID(idInt), true
}

// helper function to setup the databsae by performing automated database migration steps.
func migrateDB(db *sql.DB) error {
	var migrations = &migrate.FileMigrationSource{
		Dir: env.MustGetenv("DB_MIGRATIONS_PATH"),
	}
	_, err := migrate.Exec(db, "mysql", migrations, migrate.Up)
	return err
}
