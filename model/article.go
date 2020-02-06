package model

import (
	"editorial/db"

	"github.com/jinzhu/gorm"
)

type Article struct {
	gorm.Model
	Date    string
	No      int
	MediaId int
	Title   string
	Body    string
	Url     string
}

type ArticleView struct {
	Article
	FormattedDate string
	FormattedBody []string
	MediaName     string
}

type Editorial struct {
	url     string
	date    string
	no      int
	mediaId int
	title   string
	body    string
}

func (av *ArticleView) Init(a Article) {
	av.ID = a.ID
	av.Date = a.Date
	av.No = a.No
	av.MediaId = a.MediaId
	av.Title = a.Title
	av.Body = a.Body
	av.Url = a.Url
	av.ensureForView()
}

func (a *ArticleView) ensureForView() {
	a.FormattedDate = a.Date[:10]
	a.MediaName = a.getMediaName()
}

func (a ArticleView) getMediaName() string {
	mediaNames := []string{"", "毎日新聞", "東京新聞"}
	return mediaNames[a.MediaId]
}

func Update(e Editorial) {
	db.Init()
	db := db.GetDb()
	defer db.Close()

	articleExBefore := Article{}

	db.Table("articles").Where("url = ?", e.url).First(&articleExBefore)
	articleExAfter := articleExBefore
	db.First(&articleExAfter)
	articleExAfter.Date = e.date
	articleExAfter.MediaId = e.mediaId
	articleExAfter.Title = e.title
	articleExAfter.Body = e.body
	db.Model(&articleExBefore).Update(&articleExAfter)
}
