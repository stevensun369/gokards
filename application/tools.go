package application

import (
	"github.com/gofiber/fiber/v2"

	// internal 
	"github.com/stevensun369/kards/models"

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

func authUsersKard(c *fiber.Ctx, kard models.Kard) {
	userEmail := c.Cookies("user_email")

	if kard.From != userEmail && kard.To != userEmail {
		c.Redirect("/auth/login")
	}
}