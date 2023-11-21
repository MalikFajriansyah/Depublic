package controller

import (
	"Depublic-App-Service/config"
	"Depublic-App-Service/model"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

/* Func untuk user*/
func LoginUser(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"message": "Login successfull"})
}

func RegisterUser(c echo.Context) error {
	db, err := config.DatabaseInit()
	if err != nil {
		log.Fatal(err)
	}
	err = db.AutoMigrate(&model.User{})
	if err != nil {
		log.Fatal(err)
	}

	var newUser model.User
	if err := c.Bind(&newUser); err != nil {
		return err
	}

	//validasi data
	if newUser.Username == "" || newUser.Password == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "username dan password harus diisi")
	}

	//cek jika data user sudah terdaftar
	var existingUser model.User
	if db.Where("username = ?", newUser.Username).First(&existingUser).Error == nil {
		return echo.NewHTTPError(http.StatusBadRequest, "username sudah digunakan")
	}

	//setel role user registrasi
	newUser.Role = "user"

	//encrypt password
	hashedPassword, err := hashPassword(newUser.Password)
	if err != nil {
		return err
	}
	newUser.Password = hashedPassword

	//create data baru ke database
	if err := db.Create(&newUser).Error; err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, map[string]string{"message": "Akun berhasil dibuat"})
}

func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

/* Func untuk user*/

/* Func untuk event*/
func GetAllEvent(c echo.Context) error {
	db, err := config.DatabaseInit()
	if err != nil {
		log.Fatal(err)
	}
	err = db.AutoMigrate(&model.Event{})
	if err != nil {
		log.Fatal(err)
	}
	var events []model.Event
	if err := db.Find(&events).Error; err != nil {
		c.JSON(http.StatusNotFound, map[string]string{"message": "Event belum tersedia"})
	}

	return c.JSON(http.StatusOK, events)
}

func CreateEvent(c echo.Context) error {
	db, err := config.DatabaseInit()
	if err != nil {
		log.Fatal(err)
	}
	err = db.AutoMigrate(&model.Event{})
	if err != nil {
		log.Fatal(err)
	}
	events := new(model.Event)
	if err := c.Bind(events); err != nil {
		return err
	}

	db.Create(&events)

	return c.JSON(http.StatusCreated, map[string]string{"message": "Berhasil menambahkan event"})
}

/* Func untuk event*/
