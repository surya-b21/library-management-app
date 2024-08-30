package repo

import (
	"context"

	"github.com/surya-b21/library-management-app/category/app/model"
	"github.com/surya-b21/library-management-app/category/app/service"
)

type CategoryRepository struct{}

func (b *CategoryRepository) Create(ctx context.Context, category model.CategoryAPI) model.Category {
	db := service.DB

	newCategory := model.Category{
		CategoryAPI: category,
	}
	db.Create(&newCategory)

	return newCategory
}

func (b *CategoryRepository) Get(ctx context.Context, id string) model.Category {
	db := service.DB

	category := model.Category{}
	db.Where("id = ?", id).First(&category)

	return category
}

func (b *CategoryRepository) GetAll(ctx context.Context) []model.Category {
	db := service.DB

	categoryList := []model.Category{}
	db.Select("id", "name").Find(&categoryList)

	return categoryList
}

func (b *CategoryRepository) Update(ctx context.Context, req model.CategoryAPI, id string) model.Category {
	db := service.DB

	category := model.Category{}
	db.Where("id = ?", id).First(&category)

	category.Name = req.Name
	db.Save(&category)

	return category
}

func (b *CategoryRepository) Delete(ctx context.Context, id string) {
	db := service.DB

	db.Where("id = ?", id).Delete(&model.Category{})
}
