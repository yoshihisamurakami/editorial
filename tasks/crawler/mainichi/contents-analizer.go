package mainichi

import (
	"editorial/model"
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

const IsContentDownloadMode = true
const ContentsPageFile = "../mainichi/html/contents.html"
const MediaIdMainichi = 1

func (c ContentsPage) GetArticleContents(url string) model.CommonEditorial {
	if IsContentDownloadMode == true {
		c.Doc = c.getContentsDocByDownload("https:" + url)
	} else {
		c.Doc = c.getContentsDocFromFile()
	}
	e := model.CommonEditorial{}
	e.Url = url
	e.Date = c.getDate()
	e.No = 0
	e.MediaId = MediaIdMainichi
	e.Title = c.getTitle()
	e.Body = c.getContents()
	return e
}

func (c ContentsPage) getContentsDocByDownload(url string) *goquery.Document {
	time.Sleep(1 * time.Second)
	doc, err := goquery.NewDocument(url)
	if err != nil {
		panic(err)
	}
	return doc
}

func (c ContentsPage) getContentsDocFromFile() *goquery.Document {
	fileInfos, _ := ioutil.ReadFile(ContentsPageFile)
	stringReader := strings.NewReader(string(fileInfos))
	doc, err := goquery.NewDocumentFromReader(stringReader)
	if err != nil {
		panic(err)
	}
	return doc
}

func (c ContentsPage) getTitle() string {
	title := c.Doc.Find(".header-box > h1")
	return strings.TrimSpace(title.Text())
}

func (c ContentsPage) getDate() string {
	dateText := c.Doc.Find(".post > time").Text()
	rep := regexp.MustCompile("(\\d+)年(\\d+)月(\\d+)日")
	result := rep.FindStringSubmatch(dateText)
	month, _ := strconv.Atoi(result[2])
	day, _ := strconv.Atoi(result[3])
	return fmt.Sprintf("%s-%02d-%02d", result[1], month, day)
}

func (c ContentsPage) getContents() string {
	var body []string
	paragraphes := c.Doc.Find(".main-text > p.txt")
	paragraphes.Each(func(index int, paragraph *goquery.Selection) {
		line := strings.TrimSpace(paragraph.Text())
		if len(line) > 0 {
			body = append(body, line)
		}
	})
	return strings.Join(body, "\n")
}
