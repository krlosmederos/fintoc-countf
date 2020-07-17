package main

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)

const (
	url       = "https://blog.fintoc.com/mensaje-del-futuro/"
	className = ".post-content"
)

func main() {

	c := colly.NewCollector()

	c.OnHTML(className, func(div *colly.HTMLElement) {
		content := strings.ToLower(div.Text)
		countF := strings.Count(content, "f")
		fmt.Println(countF)
	})

	c.Visit(url)
}
