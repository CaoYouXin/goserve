package routes

import (
	"fmt"
	"strconv"

	sntest "github.com/CaoYouXin/goserve/orm/sn"
	"github.com/labstack/echo"
)

// Init routes
func Init(e *echo.Echo) {
	e.GET("/square/of/:number", func(c echo.Context) error {
		number, err := strconv.Atoi(c.Param("number"))
		if err != nil {
			return c.String(400, "We need a number.\n")
		}

		square, err := sntest.GetSquareNumber(number)
		if err != nil {
			return c.String(404, fmt.Sprintf("We can't find a square of %d\n", number))
		}

		return c.String(200, fmt.Sprintf("The square number of %d is: %d\n", number, square.SquareNumber))
	})
}
