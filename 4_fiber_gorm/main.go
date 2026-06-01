package main

import(
	"github.com/gofiber/fiber/v3"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"github.com/golang-jwt/jwt/v4"
	"log"
	"os"
	"time"
)

const (
	host 	= "localhost"
	port	= 5432
	user 	= "myuser"
	password	= "mypassword"
	dbname 	= "mydatabase"
)

var jwtSecretKey = []byte("your_super_secret_key")

func authRequired(c fiber.Ctx) error{
	cookie := c.Cookies("jwt")
	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error){
		return jwtSecretKey, nil
	})

	if err != nil || !token.Valid{
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	return c.Next()
}

func main(){
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
	host, port, user, password, dbname)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:	time.Second,
			LogLevel:				logger.Info,
			Colorful:				true,
		},
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil{
		panic("failed to connect to database")
	}

	db.AutoMigrate(&Book{})
	fmt.Println("Database migration completed!")

	app := fiber.New()

	app.Use("/books", authRequired)

	app.Get("/books", func(c fiber.Ctx) error{
		return getBooks(db, c)
	})
	app.Get("/books/:id", func(c fiber.Ctx) error{
		return getBook(db, c)
	})
	app.Post("/books", func(c fiber.Ctx) error{
		return createBook(db, c)
	})
	app.Put("/books/:id", func(c fiber.Ctx) error{
		return updateBook(db, c)
	})
	app.Delete("/books/:id", func(c fiber.Ctx) error{
		return deleteBook(db, c)
	})

	app.Post("/register", func(c fiber.Ctx) error{
		return createUser(db, c)
	})
	app.Post("/login", func(c fiber.Ctx) error{
		return loginUser(db, c)
	})

	log.Fatal(app.Listen(":8000"))
}