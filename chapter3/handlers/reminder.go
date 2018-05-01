package handlers

import "github.com/labstack/echo"

func CreateReminder(ctx echo.Context) error {
	return nil
}

func GetReminder(ctx echo.Context) error {
	ctx.Logger().Info("Reminder id is: ", c.Param("id"))
	return nil
}
