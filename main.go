package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	// Root level middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Group level middleware
	g := e.Group("/admin")
	g.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		if username == "joe" && password == "secret" {
			return true, nil
		}
		return false, nil
	}))
	g.GET("", func(c echo.Context) error {
		return c.String(http.StatusOK, "welcome")
	})

	// Route level middleware
	track := func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			println("request to /users")
			return next(c)
		}
	}
	e.GET("/users", func(c echo.Context) error {
		return c.String(http.StatusOK, "/users")
	}, track)

	// cookie test
	e.GET("/cookie", func(c echo.Context) error {
		cookies := c.Cookies()
		for _, cookie := range cookies {
			fmt.Println(cookie.Name)
			fmt.Println(cookie.Value)
		}
		cookie := new(http.Cookie)
		cookie.Name = fmt.Sprintf("%s%d", "cookie_", len(cookies))
		cookie.Value = fmt.Sprintf("%s%d", "value_", len(cookies))
		cookie.Expires = time.Now().Add(24 * time.Hour)
		c.SetCookie(cookie)
		return c.String(http.StatusOK, "open console to check cookies\n")
	})

	e.Logger.Fatal(e.Start(":1323"))
}
