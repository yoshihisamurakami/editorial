package tokyo

import (
	"editorial/model"
)

type TopAnalizer struct {
	model.TopAnalizer
}

type ContentsPage struct {
	model.ContentsPage
}

type Editorial struct {
	model.CommonEditorial
}

func Crawl() {
	crawler := model.Crawler{}
	crawler.SetTargetSiteName("東京新聞")
	topAnalizer := TopAnalizer{}
	contentsPage := ContentsPage{}
	crawler.Crawl(topAnalizer, contentsPage)
}
