package usecase

import (
	"go-article/src/business/domain"
	"go-article/src/business/usecase/article"
)

type Usecase struct {
	Article article.Interface
}

func Init(d *domain.Domains) *Usecase {
	uc := &Usecase{
		Article: article.Init(d.Article),
	}

	return uc
}
