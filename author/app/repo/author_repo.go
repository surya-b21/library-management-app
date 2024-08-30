package repo

import (
	"context"

	"github.com/surya-b21/library-management-app/author/app/model"
	"github.com/surya-b21/library-management-app/author/app/service"
)

type AuthorRepository struct{}

func (b *AuthorRepository) Create(ctx context.Context, author model.AuthorAPI) model.Author {
	db := service.DB

	newAuthor := model.Author{
		AuthorAPI: author,
	}
	db.Create(&newAuthor)

	return newAuthor
}

func (b *AuthorRepository) Get(ctx context.Context, id string) model.Author {
	db := service.DB

	author := model.Author{}
	db.Where("id = ?", id).First(&author)

	return author
}

func (b *AuthorRepository) GetAll(ctx context.Context) []model.Author {
	db := service.DB

	authorList := []model.Author{}
	db.Select("id", "name").Find(&authorList)

	return authorList
}

func (b *AuthorRepository) Update(ctx context.Context, req model.AuthorAPI, id string) model.Author {
	db := service.DB

	author := model.Author{}
	db.Where("id = ?", id).First(&author)

	author.Name = req.Name
	db.Save(&author)

	return author
}

func (b *AuthorRepository) Delete(ctx context.Context, id string) {
	db := service.DB

	db.Where("id = ?", id).Delete(&model.Author{})
}
