package model

import "fmt"

type Crawler struct {
	TargetSiteName string
}

func (c *Crawler) SetTargetSiteName(name string) {
	c.TargetSiteName = name
}

func (c Crawler) Crawl(topAnalizer TopAnalizerInterface, contentsPage ContentsPageInterface) {
	latestLinks := topAnalizer.GetEditorialLinks()
	for i := 0; i < len(latestLinks); i++ {
		article := contentsPage.GetArticleContents(latestLinks[i])
		if count := article.Count(); count == 0 {
			article.Insert()
			fmt.Println(article.Url + "  Insert!")
		} else {
			article.Update()
			fmt.Println(article.Url + "  Update!")
		}
	}
}
