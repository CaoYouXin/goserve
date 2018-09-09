package orm

import (
	"database/sql"
)

var pool *sql.DB

// Init test
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

// Close test
func Close() {
	pool.Close()
}
