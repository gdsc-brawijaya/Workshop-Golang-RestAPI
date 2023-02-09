package domain

import (
	"go-article/src/business/domain/article"

	"gorm.io/gorm"
)

type Domains struct {
	Article article.Interface
}

func Init(db *gorm.DB) *Domains {
	d := &Domains{
		Article: article.Init(db),
	}

	return d
}
