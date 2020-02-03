package model

import (
	"editorial/db"

	"github.com/PuerkitoBio/goquery"
)

type TopAnalizer struct {
	Doc *goquery.Document
}

type TopAnalizerInterface interface {
	GetEditorialLinks() []string
}

type ContentsPage struct {
	Doc *goquery.Document
}

type ContentsPageInterface interface {
	GetArticleContents(string) CommonEditorial
}

type CommonEditorial struct {
	Url     string
	Date    string
	No      int
	MediaId int
	Title   string
	Body    string
}

func (e CommonEditorial) Count() (count int) {
	db.Init()
	db := db.GetDb()
	defer db.Close()

	db.Table("articles").Where("url = ?", e.Url).Count(&count)
	return
}

func (e CommonEditorial) Insert() {
	db.Init()
	db := db.GetDb()
	defer db.Close()

	articleEx := Article{}
	articleEx.Date = e.Date
	articleEx.MediaId = e.MediaId
	articleEx.Title = e.Title
	articleEx.Body = e.Body
	articleEx.Url = e.Url
	db.Create(&articleEx)
}

func (e CommonEditorial) Update() {
	db.Init()
	db := db.GetDb()
	defer db.Close()

	articleExBefore := Article{}

	db.Table("articles").Where("url = ?", e.Url).First(&articleExBefore)
	articleExAfter := articleExBefore
	db.First(&articleExAfter)
	articleExAfter.Date = e.Date
	articleExAfter.MediaId = e.MediaId
	articleExAfter.Title = e.Title
	articleExAfter.Body = e.Body
	db.Model(&articleExBefore).Update(&articleExAfter)
}
