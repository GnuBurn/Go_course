package main

import (
	"gorm.io/gorm"
	"github.com/gofiber/fiber/v3"
)

type Book struct{
	gorm.Model
	Name				string `json:"name"`
	Author 			string `json:"author"`
	Description	string `json:"description"`
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

// func CreateBook(db *gorm.DB, book *Book){
// 	result := db.Create(book)
// 	if result.Error != nil{
// 		log.Fatalf("Error creating book: %v", result.Error)
// 	}
// 	fmt.Println("Book created successfully")
// }

// func GetBook(db *gorm.DB, id uint) *Book{
// 	var book Book
// 	result := db.First(&book, id)
// 	if result.Error != nil{
// 		log.Fatalf("Error finding book: %v", result.Error)
// 	}
// 	return &book
// }

// func UpdateBook(db *gorm.DB, book *Book){
// 	result := db.Save(book)
// 	if result.Error != nil{
// 		log.Fatalf("Error updating book: %v", result.Error)
// 	}
// 	fmt.Println("Book updated successfully")
// }

// func DeleteBook(db *gorm.DB, id uint){
// 	var book Book
// 	result := db.Delete(&book, id)
// 	if result.Error != nil{
// 		log.Fatalf("Error deleting book: %v", result.Error)
// 	}
// 	fmt.Println("Book deleted successfully")
// }

// func HardDeleteBook(db *gorm.DB, id uint){
// 	var book Book
// 	result := db.Unscoped().Delete(&book, id)
// 	if result.Error != nil{
// 		return result.Error
// 	}
// 	fmt.Println("Book hard deleted successfully")
// 	return nil
// }

// func getBooksSortedByCreatedAt(db *gorm.DB) ([]Book, error){
// 	var books []Book
// 	result := db.Order("created_at desc").Find(&books)
// 	if result.Error != nil{
// 		return nil, result.Error
// 	}
// 	return books, nil
// }

// func getBooksByAuthorName(db *gorm.DB, authorName string) ([]Book, error){
// 	var books []Book
// 	result := db.Where("author = ?", authorName).Find(&books)
// 	if result.Error != nil{
// 		return nil, result.Error
// 	}
// 	return books, nil