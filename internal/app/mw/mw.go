package mw

import (
	"github.com/labstack/echo/v4"
	"log"
)

const roleAdmin = "admin"

func RoleCheck(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {

		val := ctx.Request().Header.Get("User-Role")
		if val == roleAdmin {
			log.Println("Red button user detected")
		}
		err := next(ctx)
		if err != nil {
			return err
		}

		return nil
	}
}
