package orm

import (
	"database/sql"
	"fmt"
)

var pool *sql.DB

// Init db
func Init() {
	var err error

	pool, err = sql.Open("mysql", "root:123457@tcp(db:3306)/ft?charset=utf8&collation=utf8_bin")
	if err != nil {
		panic(err.Error())
	}

	err = pool.Ping()
	if err != nil {
		panic(err.Error())
	}
}

// Close db
func Close() {
	pool.Close()
	fmt.Println("\ndb pool... closed.")
}

// Exec sql
func Exec(sql string, args ...interface{}) (sql.Result, error) {
	return pool.Exec(sql, args...)
}

// Query sql
func Query(sql string, args ...interface{}) (*sql.Rows, error) {
	return pool.Query(sql, args...)
}
