package main

import (
	"fmt"
	"strconv"

	_ "github.com/go-sql-driver/mysql"

	"github.com/CaoYouXin/goserve/orm"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	orm.Init()
	defer orm.Close()
	// orm.Insert()

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Static("/", "public")

	e.GET("/square/of/:number", func(c echo.Context) error {
		number, err := strconv.Atoi(c.Param("number"))
		if err != nil {
			return c.String(400, "We need a number.\n")
		}

		square, err := orm.Select(number)
		if err != nil {
			return c.String(404, fmt.Sprintf("We can't find a square of %d\n", number))
		}

		return c.String(200, fmt.Sprintf("The square number of %d is: %d\n", number, square))
	})

	e.Logger.Fatal(e.Start(":1323"))
}
