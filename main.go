package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"github.com/CaoYouXin/goserve/orm"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	orm.Init()
	defer orm.Close()
	orm.Insert()

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Static("/", "public")

	e.GET("/s", func(c echo.Context) error {
		a := fmt.Sprintf("The square number of 13 is: %d", orm.Select(13))

		b := fmt.Sprintf("The square number of 1 is: %d", orm.Select(1))

		return c.String(200, fmt.Sprintf("%s\n%s", a, b))
	})

	e.Logger.Fatal(e.Start(":1323"))
}
