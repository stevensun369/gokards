package application

import (
	"github.com/gofiber/fiber/v2"

	"math/rand"
	"strconv"
)

func authUser(c *fiber.Ctx) {
	userEmail := c.Cookies("user_email")

	if userEmail == "" {
		c.Redirect("/auth/login")
	}
}

func createKardID() string {
	var kardID string
	for i := 0; i < 10; i++ {
		kardID += strconv.Itoa(rand.Intn(9))
	}
	return kardID
}