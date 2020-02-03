package mainichi

import (
	"io/ioutil"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const IsTopDownloadMode = true
const EditorialTopUrl = "https://mainichi.jp/editorial/"
const TopPageFile = "../mainichi/html/top.html"
const LinksCountToGet = 4

func (t TopAnalizer) GetEditorialLinks() []string {
	t.Doc = t.getDoc()
	links := t.getLinks()
	return links[:LinksCountToGet]
}

func (t TopAnalizer) getDoc() *goquery.Document {
	if IsTopDownloadMode == true {
		return t.getDocByDownload()
	} else {
		return t.getDocFromFile()
	}
}

func (t TopAnalizer) getDocByDownload() *goquery.Document {
	doc, err := goquery.NewDocument(EditorialTopUrl)
	if err != nil {
		panic(err)
	}
	return doc
}

func (t TopAnalizer) getDocFromFile() *goquery.Document {
	fileInfos, _ := ioutil.ReadFile(TopPageFile)
	stringReader := strings.NewReader(string(fileInfos))
	doc, err := goquery.NewDocumentFromReader(stringReader)
	if err != nil {
		panic(err)
	}
	return doc
}

func (t TopAnalizer) getLinks() []string {
	var ret []string
	links := t.Doc.Find("article > a")
	links.Each(func(index int, link *goquery.Selection) {
		href, exists := link.Attr("href")
		if exists {
			ret = append(ret, href)
		}
	})
	return ret
}
