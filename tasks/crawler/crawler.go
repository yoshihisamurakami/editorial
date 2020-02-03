package main

import (
	"editorial/tasks/crawler/mainichi"
	"editorial/tasks/crawler/tokyo"
)

func main() {
	mainichi.Crawl()
	tokyo.Crawl()
}
