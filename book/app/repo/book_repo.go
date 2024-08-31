package repo

import (
	"context"

	"github.com/surya-b21/library-management-app/book/app/model"
	"github.com/surya-b21/library-management-app/book/app/service"
)

type BookRepository struct{}

func (b *BookRepository) Create(ctx context.Context, book model.BookAPI) model.Book {
	db := service.DB

	newBook := model.Book{
		BookAPI: book,
	}
	db.Create(&newBook)

	return newBook
}

func (b *BookRepository) Get(ctx context.Context, id string) model.Book {
	db := service.DB

	book := model.Book{}
	db.Where("id = ?", id).First(&book)

	return book
}

func (b *BookRepository) GetAll(ctx context.Context) []model.Book {
	db := service.DB

	bookList := []model.Book{}
	db.Select("id", "title", "pages", "author_id", "category_id", "stock").Find(&bookList)

	return bookList
}

func (b *BookRepository) Update(ctx context.Context, req model.BookAPI, id string) model.Book {
	db := service.DB

	book := model.Book{}
	db.Where("id = ?", id).First(&book)

	book.Title = req.Title
	book.Pages = req.Pages
	book.AuthorID = req.AuthorID
	book.CategoryID = req.CategoryID
	book.Stock = req.Stock
	db.Save(&book)

	return book
}

func (b *BookRepository) Delete(ctx context.Context, id string) {
	db := service.DB

	db.Where("id = ?", id).Delete(&model.Book{})
}

func (b *BookRepository) Borrow(ctx context.Context, id string) *string {
	db := service.DB

	book := model.Book{}
	db.Where("id = ?", id).First(&book)

	if *book.Stock == 0 {
		message := "Cannot borrow book because stock is zero"
		return &message
	}

	newStock := *book.Stock - 1
	book.Stock = &newStock
	db.Save(&book)
	return nil
}

func (b *BookRepository) Return(ctx context.Context, id string) {
	db := service.DB

	book := model.Book{}
	db.Where("id = ?", id).First(&book)

	newStock := *book.Stock + 1
	book.Stock = &newStock
	db.Save(&book)
}
