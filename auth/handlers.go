package auth

import(
	"github.com/gofiber/fiber/v2"

	"github.com/stevensun369/kards/models"
	"github.com/stevensun369/kards/database"

	"golang.org/x/crypto/bcrypt"

	"strconv"
	"time"
	"fmt"

)

// Register handlers
func getRegister(c *fiber.Ctx) error {

	password := c.Query("password")
	email := c.Query("email")

	return c.Render("auth/register", fiber.Map{
		"Title": "Înregistrare",
		"password": password,
		"email": email,
	}, "layouts/main")
}

func postRegister(c *fiber.Ctx) error {

	//  --- email verification
	email := c.FormValue("email")

	// db connection
	db := database.DBConn
	var checkUser models.User

	// --- email
	emailError := db.Where("email = ?", email).First(&checkUser).Error
	if emailError == nil {
		return c.Redirect("/auth/register?email=no")
	}

	// --- password and password confirm verification
	password := c.FormValue("password")
	passwordConfirm := c.FormValue("password_confirm")

	if password != passwordConfirm {
		return c.Redirect("/auth/register?password=no")
	}

	// --- password hashing
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return c.Redirect("/elev/register")
	}

	nume := c.FormValue("nume")
	prenume := c.FormValue("prenume")

	var user models.User

	user.Nume = nume;
	user.Prenume = prenume;
	user.Email = email;
	user.Password = string(hashedPassword);
	db.Save(&user)

	return c.Redirect("/auth/login")
}

// Login handlers
func getLogin(c *fiber.Ctx) error {

	user := c.Query("user")
	password := c.Query("password")

	return c.Render("auth/login", fiber.Map{
		"Title": "Conectare",
		"user": user,
		"password": password,
	}, "layouts/main")
}

func postLogin(c *fiber.Ctx) error {
	email := c.FormValue("email")
	password := c.FormValue("password")

	// db connection
	db := database.DBConn
	var user models.User

	query :=  db.Where("email = ?", email).First(&user)
	if exists := query.Error; exists != nil {
		return c.Redirect("/auth/login?user=no")
	}

	// if bcrypt is not null, then the password is not correct, and we redirect to login page
	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) != nil {
		fmt.Printf("no")
		return c.Redirect("/auth/login?password=no")
	}

	c.Cookie(&fiber.Cookie{
		Name: "user_id",
		Value: strconv.Itoa(user.ID),
		Expires: time.Now().Add(720 * time.Hour),
	})

	c.Cookie(&fiber.Cookie{
		Name: "user_email",
		Value: user.Email,
		Expires: time.Now().Add(720 * time.Hour),
	})

	return c.Redirect("/")
}

// Logout handler
func getLogout(c *fiber.Ctx) error {

	c.Cookie(&fiber.Cookie{
		Name: "user_id",
		Value: "",
		Expires: time.Now().Add(720 * time.Hour),
	})

	c.Cookie(&fiber.Cookie{
		Name: "user_email",
		Value: "",
		Expires: time.Now().Add(720 * time.Hour),
	})

	return c.Redirect("/auth/login")
}