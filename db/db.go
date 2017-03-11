package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

const (
	//DBDriver name of database driver
	DBDriver = "mysql"
	//DBName name of database
	DBName = "blog"
	// DBUser is user of database
	DBUser = "didiyudha"
	// DBPassword is password database
	DBPassword = "ytrewq"
	// DBURL is url of database
	DBURL = DBUser + ":" + DBPassword + "@/" + DBName
)

// GetDB return DB
func GetDB() (*sql.DB, error) {
	db, err := sql.Open(DBDriver, DBURL)
	if err != nil {
		return db, err
	}
	err = db.Ping()
	if err != nil {
		return db, err
	}
	return db, nil
}
