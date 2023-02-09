package article

import (
	articleDom "go-article/src/business/domain/article"
	"go-article/src/business/entity"
)

type Interface interface {
	GetList() ([]entity.Article, error)
	Create(param entity.ArticleInputParam) error
	Update(param entity.ArticleInputParam, uriParam entity.ArticleParam) error
	Delete(param entity.ArticleParam) error
}

type article struct {
	article articleDom.Interface
}

func Init(ad articleDom.Interface) Interface {
	a := &article{
		article: ad,
	}

	return a
}

func (a *article) GetList() ([]entity.Article, error) {
	articles, err := a.article.GetList()
	if err != nil {
		return articles, err
	}

	return articles, nil
}

func (a *article) Create(param entity.ArticleInputParam) error {
	// mapping data dari struct param menuju struct Model Article
	article := entity.Article{
		Title: param.Tilte,
		Body:  param.Body,
	}

	// panggil method create dari domain layer untuk menyimpan data kedalam database
	err := a.article.Create(article)
	if err != nil {
		return err
	}

	return nil
}

func (a *article) Update(param entity.ArticleInputParam, uriParam entity.ArticleParam) error {
	article, err := a.article.Get(uriParam.Id)
	if err != nil {
		return err
	}

	article.Title = param.Tilte
	article.Body = param.Body

	err = a.article.Update(article)
	if err != nil {
		return err
	}

	return nil
}

func (a *article) Delete(param entity.ArticleParam) error {
	err := a.article.Delete(param.Id)
	if err != nil {
		return err
	}

	return nil
}
