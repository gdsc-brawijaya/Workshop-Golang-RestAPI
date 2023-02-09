package rest

import (
	"go-article/src/business/entity"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *rest) GetAllArticle(ctx *gin.Context) {
	articles, err := r.uc.Article.GetList()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"data":  nil,
			"msg":   "failed to get articles",
			"error": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":  articles,
		"msg":   "successfully get all articles",
		"error": nil,
	})
}

func (r *rest) CreateArticle(ctx *gin.Context) {
	var param entity.ArticleInputParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"data":  nil,
			"msg":   "failed to create article",
			"error": err,
		})
		return
	}

	err := r.uc.Article.Create(param)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"data":  nil,
			"msg":   "failed to create article",
			"error": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":  nil,
		"msg":   "successfully create article",
		"error": nil,
	})
}

func (r *rest) UpdateArticle(ctx *gin.Context) {
	var param entity.ArticleInputParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"data":  nil,
			"msg":   "failed to update article",
			"error": err,
		})
		return
	}

	var uriParam entity.ArticleParam
	if err := ctx.ShouldBindUri(&uriParam); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"data":  nil,
			"msg":   "failed to update article",
			"error": err,
		})
		return
	}

	err := r.uc.Article.Update(param, uriParam)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"data":  nil,
			"msg":   "failed to update article",
			"error": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":  nil,
		"msg":   "successfully update article",
		"error": nil,
	})
}

func (r *rest) DeleteArticle(ctx *gin.Context) {
	var param entity.ArticleParam
	if err := ctx.ShouldBindUri(&param); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"data":  nil,
			"msg":   "failed to delete article",
			"error": err,
		})
		return
	}

	err := r.uc.Article.Delete(param)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"data":  nil,
			"msg":   "failed to delete article",
			"error": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":  nil,
		"msg":   "successfully delete article",
		"error": nil,
	})
}
