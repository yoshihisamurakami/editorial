package model

import "fmt"

type Crawler struct {
	TargetSiteName string
}

func (c *Crawler) SetTargetSiteName(name string) {
	c.TargetSiteName = name
}

func (c Crawler) Crawl(topAnalizer TopAnalizerInterface, contentsPage ContentsPageInterface) {
	fmt.Println("crawler start...  " + c.TargetSiteName)
	latestLinks := topAnalizer.GetEditorialLinks()
	for i := 0; i < len(latestLinks); i++ {
		article := contentsPage.GetArticleContents(latestLinks[i])
		fmt.Println(article.Url)
		if count := article.Count(); count == 0 {
			article.Insert()
			fmt.Println("Insert!")
		} else {
			article.Update()
			fmt.Println("Update!")
		}
	}
}
