package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/CaoYouXin/goserve/migrate"
	"github.com/CaoYouXin/goserve/orm"
	"github.com/CaoYouXin/goserve/routes"
)

func main() {
	orm.Init()
	defer orm.Close()

	if err := migrate.Init(); err != nil {
		fmt.Printf("%v", err.Error())
		return
	}

	e := echo.New()

	go func() {
		e.Use(middleware.Logger())
		e.Use(middleware.Recover())

		e.Static("/", "public")

		routes.Init(e)

		// e.Logger.Fatal(e.Start(":1323"))
		if err := e.Start(":1323"); err != nil {
			e.Logger.Info("shutting down the server")
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
