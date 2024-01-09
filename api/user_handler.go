package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wanyansu/holidayinn/types"
)

func HandleGetUsers(c *fiber.Ctx) error {
	u := types.User{
		FirstName: "John",
		LastName:  "Doe",
	}
	return c.JSON(u)
}
