package gokit

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// DbConnect Returns a connection to db or error
// it takes in the db credentials and other configurations e.g max connections.
// if this function returns an error, it is advisable to stop execution of the application.
func DbConnect(dbDriver, dbUser, dbPass, dbHost, dbPort, dbName string,
	maxOpenCon, maxIdleCon, maxOpenConLifetime, maxIdleConLifetime int) (*sql.DB, error) {
	dbUri := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=latin1&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbPort, dbName)

	if dbDriver != "mysql" {
		return nil, fmt.Errorf("%s only mysql driver is configured", LogPrefix)
	}

	Log("%s Connecting to DB...", LogPrefix)
	db, err := sql.Open(dbDriver, dbUri)
	if err != nil {
		return nil, err
	}

	Log("%s Testing connection...", LogPrefix)
	if err = db.Ping(); err != nil {
		return nil, err
	}
	Log("%s Connected to DB successfully...", LogPrefix)

	Log("%s Setting MAX_OPEN_CONNECTIONS...", LogPrefix)
	db.SetMaxOpenConns(maxOpenCon)

	Log("%s Setting MAX_IDLE_CONNECTIONS...", LogPrefix)
	db.SetMaxIdleConns(maxIdleCon)

	Log("%s Setting max open connection lifetime to %d", LogPrefix, maxOpenConLifetime)
	db.SetConnMaxLifetime(time.Duration(maxOpenConLifetime) * time.Second)

	Log("%s Setting max idle connection lifetime to %d", LogPrefix, maxIdleConLifetime)
	db.SetConnMaxIdleTime(time.Duration(maxIdleConLifetime) * time.Second)
	return db, nil
}
