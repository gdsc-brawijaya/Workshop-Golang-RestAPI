package article

import (
	"go-article/src/business/entity"

	"gorm.io/gorm"
)

type Interface interface {
	GetList() ([]entity.Article, error)
	Create(article entity.Article) error
	Update(article entity.Article) error
	Get(id uint) (entity.Article, error)
	Delete(id uint) error
}

type article struct {
	db *gorm.DB
}

func Init(db *gorm.DB) Interface {
	a := &article{
		db: db,
	}

	return a
}

func (a *article) GetList() ([]entity.Article, error) {
	articles := []entity.Article{}

	err := a.db.Find(&articles).Error
	if err != nil {
		return articles, err
	}

	return articles, nil
}

func (a *article) Create(article entity.Article) error {
	err := a.db.Create(&article).Error
	if err != nil {
		return err
	}

	return nil
}

func (a *article) Update(article entity.Article) error {
	err := a.db.Save(&article).Error
	if err != nil {
		return err
	}

	return err
}

func (a *article) Get(id uint) (entity.Article, error) {
	article := entity.Article{}

	err := a.db.Where("id = ?", id).First(&article).Error
	if err != nil {
		return article, err
	}

	return article, nil
}

func (a *article) Delete(id uint) error {
	err := a.db.Delete(&entity.Article{}, id).Error
	if err != nil {
		return err
	}

	return nil
}
