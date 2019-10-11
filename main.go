package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)

func main() {
	incomputer()
	gigacomputer()
}

func incomputer() {
	// Instantiate default collector
	c := colly.NewCollector(
		// Allow requests only to store.xkcd.com
		colly.AllowedDomains("www.incomputer.cz"),
	)
	// Extract product details
	c.OnHTML(".product", func(e *colly.HTMLElement) {
		rawPrice := e.ChildText(".view-price b")
		// \xc2\xa0 is non-breakable space
		priceStr := strings.ReplaceAll(strings.ReplaceAll(rawPrice, "Kč", ""), "\xc2\xa0", "")
		price, _ := strconv.Atoi(priceStr)

		if price < 6000 || price > 10000 {
			return
		}
		name := e.ChildText("h3 a")
		desc := e.ChildText(".description")

		fmt.Println(name, "|", price, "|", desc)
	})
	// Find and visit next page links
	c.OnHTML(`a[href].ico-next`, func(e *colly.HTMLElement) {
		e.Request.Visit(e.Attr("href"))
	})
	c.Visit("https://www.incomputer.cz/inshop/scripts/search.aspx?q=thinkpad")
}

func gigacomputer() {
	// Instantiate default collector
	c := colly.NewCollector(
		// Allow requests only to store.xkcd.com
		colly.AllowedDomains("www.gigacomputer.cz"),
	)
	// Extract product details
	c.OnHTML(".product", func(e *colly.HTMLElement) {
		rawPrice := e.ChildText(".price")
		// \xc2\xa0 is non-breakable space
		priceStr := strings.ReplaceAll(strings.Split(rawPrice, ",")[0], "\xc2\xa0", "")
		price, _ := strconv.Atoi(priceStr)

		if price < 6000 || price > 10000 {
			return
		}
		name := e.ChildText("h3 a")
		desc := e.ChildText("p")

		fmt.Println(name, "|", price, "|", desc)
	})
	// Find and visit next page links
	c.OnHTML(`.page.next a[href]`, func(e *colly.HTMLElement) {
		e.Request.Visit(e.Attr("href"))
	})
	c.Visit("https://www.gigacomputer.cz/hledani/?q=thinkpad")
}
