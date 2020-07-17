package main

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)

// ClassName is div's class
const ClassName string = ".post-content"

// URL is url for scrapping
const URL string = "https://blog.fintoc.com/mensaje-del-futuro/"

func main() {

	c := colly.NewCollector()

	c.OnHTML(ClassName, func(div *colly.HTMLElement) {
		content := strings.ToLower(div.Text)
		countF := strings.Count(content, "f")
		fmt.Println(countF)
	})

	c.Visit(URL)
}
