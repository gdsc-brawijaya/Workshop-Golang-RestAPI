package rest

import (
	"go-article/src/business/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Rest interface {
	Run()
}

type rest struct {
	uc  *usecase.Usecase
	gin *gin.Engine
}

func Init(uc *usecase.Usecase) Rest {
	r := &rest{
		uc:  uc,
		gin: gin.Default(),
	}
	r.Register()

	return r
}

func (r *rest) Run() {
	r.gin.Run()
}

func (r *rest) Register() {
	api := r.gin.Group("/api")
	api.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "Hello from api",
		})
	})

	articles := api.Group("/articles")
	articles.GET("/", r.GetAllArticle)
	articles.POST("/", r.CreateArticle)
	articles.PUT("/:id", r.UpdateArticle)
	articles.DELETE("/:id", r.DeleteArticle)
}
