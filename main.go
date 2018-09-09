package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"github.com/labstack/echo"
)

func main() {
	db, err := sql.Open("mysql", "root:123457@db/ft?charset=utf8&collation=utf8_bin")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	e := echo.New()

	e.Static("/", "public")

	e.GET("/s", func(c echo.Context) error {
		// Prepare statement for inserting data
		stmtIns, err := db.Prepare("INSERT INTO squareNum(number, squareNumber) VALUES( ?, ? )") // ? = placeholder
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		defer stmtIns.Close() // Close the statement when we leave main() / the program terminates

		// Prepare statement for reading data
		stmtOut, err := db.Prepare("SELECT squareNumber FROM squareNum WHERE number = ?")
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		defer stmtOut.Close()

		// Insert square numbers for 0-24 in the database
		for i := 0; i < 25; i++ {
			_, err = stmtIns.Exec(i, (i * i)) // Insert tuples (i, i^2)
			if err != nil {
				panic(err.Error()) // proper error handling instead of panic in your app
			}
		}

		var squareNum int // we "scan" the result in here

		// Query the square-number of 13
		err = stmtOut.QueryRow(13).Scan(&squareNum) // WHERE number = 13
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		a := fmt.Sprintf("The square number of 13 is: %d", squareNum)

		// Query another number.. 1 maybe?
		err = stmtOut.QueryRow(1).Scan(&squareNum) // WHERE number = 1
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		b := fmt.Sprintf("The square number of 1 is: %d", squareNum)

		return c.String(200, fmt.Sprintf("%s\n%s", a, b))
	})

	e.Logger.Fatal(e.Start(":1323"))
}
