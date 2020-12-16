package application

import (
	"github.com/gofiber/fiber/v2"

	"fmt"
	"strings"

	// img related
	_ "image/png"
	_ "image/jpeg"
	"image"
	"os"

	// other appss
	"github.com/stevensun369/kards/database"
	"github.com/stevensun369/kards/models"
)

func getHome(c *fiber.Ctx) error {

	authUser(c)

	db := database.DBConn
	
	var kards []models.Kard
	db.Order("created_at desc").Where("to", c.Cookies("user_email")).Find(&kards)



	return c.Render("application/home", fiber.Map{
		"Title": "acasă",
		"kards": kards,
	}, "layouts/main")
}

func getSent(c *fiber.Ctx) error {

	authUser(c)

	db := database.DBConn
	
	var kards []models.Kard
	db.Order("created_at desc").Where("from", c.Cookies("user_email")).Find(&kards)

	return c.Render("application/sent", fiber.Map{
		"Title": "kard-urile trimise",
		"kards": kards,
	}, "layouts/main")
}

func getAdd(c *fiber.Ctx) error {

	authUser(c)

	return c.Render("application/add", fiber.Map{
		"Title": "create a new kard",
	}, "layouts/main")
}

func postAdd(c *fiber.Ctx) error {

	to := c.FormValue("to")
	toSlice := strings.Split(to, " ")
	to = ""
	for _, s := range toSlice {
		to += s
	}

	background := c.FormValue("background")
	// orientation := c.FormValue("orientation")

	message := c.FormValue("message")
	fontColor := c.FormValue("font-color")
	font := c.FormValue("font")

	db := database.DBConn

	// checking that a kard id doesn't exist
	var kardID string
	kardCheck := true
	for kardCheck {
		kardID = createKardID()
		var kardTest models.Kard
		queryErr := db.Where("kard_id", kardID).First(&kardTest).Error
		
		// if there is an error, it means that the kard id doesn't exist, and I will exit the loop
		if queryErr != nil {
			kardCheck = false
		}
	}

	// uploading the image
	imageUpload, _ := c.FormFile("image")
	fileExtensionSlice := strings.Split(imageUpload.Filename, ".")
	fileExtension := fileExtensionSlice[len(fileExtensionSlice) - 1]
	filePath := fmt.Sprintf("./media/%s", kardID) + "." + fileExtension
	fileLink := fmt.Sprintf("/media/%s", kardID) + "." + fileExtension

	c.SaveFile(imageUpload, filePath)
	
	// getting the image width and height
	var orientation string
	openFile, _ := os.Open(filePath)
	file, _, _ := image.DecodeConfig(openFile)
	if file.Width > file.Height {
		orientation = "l"
	} else if file.Height > file.Width {
		orientation = "p"
	} else if file.Height == file.Width {
		orientation = "l";
	}

	// From name
	var userFrom models.User
	db.Select("nume", "prenume").Where("id", c.Cookies("user_id")).First(&userFrom)

	// to name
	var userTo models.User
	db.Select("nume", "prenume").Where("email", to).First(&userTo)

	// create the kard
	kard := models.Kard{
		KardID: kardID,
		FromName: userFrom.Nume + " " + userFrom.Prenume,
		From: c.Cookies("user_email"),
		To: to,
		ToName: userTo.Nume + " " + userTo.Prenume,
		Background: background,
		Orientation: orientation,
		Image: fileLink,
		Message: message,
		Font: font,
		Color: fontColor,
	}

	db.Create(&kard)

	return c.Redirect("/")
}

func getKard(c *fiber.Ctx) error {
	kardID := c.Params("kardID")

	prev := c.Query("prev")

	db := database.DBConn

	// kard model
	var kard models.Kard
	kardQuery := db.Where("kard_id", kardID).First(&kard)

	if kardQuery.Error != nil {
		return c.Render("application/kard-problem", fiber.Map{}, "layouts/main")
	}

	authUsersKard(c, kard)

	return c.Render("application/kard", fiber.Map{
		"kard": kard,
		"prev": prev,
	}, "layouts/main")
}

func getHelp(c *fiber.Ctx) error {

	authUser(c)

	db := database.DBConn

	var user models.User
	db.First(&user, c.Cookies("user_id"))

	user.DidHelp = true

	db.Save(&user)

	return c.Render("application/help", fiber.Map{}, "layouts/main")
}