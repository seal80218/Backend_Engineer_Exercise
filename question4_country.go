package main

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func (c *Clawer) Question4_country(args ...string) {
	doc := Query(c.url + "topsites/countries/" + args[0])

	i := 1
	doc.Find("div.tr.site-listing").Each(func(index int, ele *goquery.Selection) {
		ele.Find("div.td.DescriptionCell").Each(func(j int, d *goquery.Selection) {
			if i <= 20 {
				fmt.Printf("%d ", i)
				fmt.Println(strings.TrimSpace(d.First().Text()))
			}
			i++
		})
	})
}
