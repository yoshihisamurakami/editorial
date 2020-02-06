package main

import (
	"editorial/tasks/crawler/mainichi"
	"editorial/tasks/crawler/tokyo"
	"time"
)

func main() {
	go mainichi.Crawl()
	go tokyo.Crawl()
	time.Sleep(10 * time.Second)
}
