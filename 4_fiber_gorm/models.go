package main

import (
	"gorm.io/gorm"
	"github.com/gofiber/fiber/v3"
	"golang.org/x/crypto/bcrypt"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type Book struct{
	gorm.Model
	Name				string `json:"name"`
	Author 			string `json:"author"`
	Description	string `json:"description"`
}

type User struct{
	gorm.Model
	Email				string	`gorm:"unique"`
	Password		string
}

func getBooks(db *gorm.DB, c fiber.Ctx) error{
	var books []Book
	db.Find(&books)
	return c.JSON(books)
}

func getBook(db *gorm.DB, c fiber.Ctx) error{
	id := c.Params("id")
	var book Book
	db.Find(&book, id)
	return c.JSON(book)
}

func createBook(db *gorm.DB, c fiber.Ctx) error{
	book := new(Book)
	if err := c.Bind().JSON(&book); err != nil {
		return err
	}
	db.Create(&book)
	return c.JSON(book)
}

func updateBook(db *gorm.DB, c fiber.Ctx) error{
	id := c.Params("id")
	book := new(Book)
	db.First(&book, id)
	if err := c.Bind().JSON(&book); err != nil{
		return err
	}
	db.Save(&book)
	return c.JSON(book)
}

func deleteBook(db *gorm.DB, c fiber.Ctx) error{
	id := c.Params("id")
	db.Delete(&Book{}, id)
	return c.SendString("Book successfully deleted")
}

func createUser(db *gorm.DB, c fiber.Ctx) error{
	user := new(User)
	if err := c.Bind().JSON(&user); err != nil {
		return err
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil{
		return err
	}
	user.Password = string(hashedPassword)

	db.Create(user)
	return c.JSON(user)
}

func loginUser(db *gorm.DB, c fiber.Ctx) error{
	var input User
	var user User

	if err := c.Bind().JSON(&input); err != nil{
		return err
	}

	db.Where("email = ?", input.Email).First(&user)
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil{
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = user.ID
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString(jwtSecretKey)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	c.Cookie(&fiber.Cookie{
		Name: "jwt",
		Value: t,
		Expires: time.Now().Add(time.Hour * 72),
		HTTPOnly: true,
	})

	return c.JSON(fiber.Map{"message": "success"})
}