package service

import (
	"editorial/db"
	"editorial/model"
	"fmt"
	"strings"
)

func GetEditorials() []model.ArticleView {
	db.Init()
	db := db.GetDb()
	defer db.Close()

	var article []model.Article
	var articleView []model.ArticleView
	db.Order("date DESC, created_at DESC").Find(&article)

	for _, v := range article {
		av := model.ArticleView{}
		av.Init(v)
		av.FormattedDate = v.Date[:10]
		av.MediaName = av.GetMediaName()
		articleView = append(articleView, av)
	}
	return articleView
}

func GetOneEditorial(id string) model.ArticleView {
	db.Init()
	db := db.GetDb()
	defer db.Close()

	var article model.Article
	var articleView model.ArticleView
	db.Find(&article, id)
	articleView.Init(article)
	articleView.FormattedBody = getArticleBody(article.Body)
	articleView.FormattedDate = article.Date[:10]
	articleView.MediaName = articleView.GetMediaName()
	return articleView
}

func GetPrevNextEditorial(id string) (prevArticle model.Article, nextArticle model.Article) {
	db.Init()
	db := db.GetDb()
	defer db.Close()

	var article []model.Article
	var count int
	db.Order("date DESC, created_at DESC").Find(&article).Count(&count)
	max := count

	for i, _ := range article {
		if fmt.Sprint(article[i].ID) == id {
			if i == 0 {
				nextArticle = article[i+1]
			} else if i+1 == max {
				prevArticle = article[i-1]
			} else {
				prevArticle = article[i-1]
				nextArticle = article[i+1]
			}
			return
		}
	}
	return
}

func getArticleBody(body string) []string {
	return strings.Split(body, "\n")
}
