package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func (c *Clawer) Question4_top(args ...string) {
	doc := Query(c.url + "topsites/")

	p := args[0]
	topNumber, err := strconv.ParseInt(p, 10, 16)
	if err != nil {
		panic(err)
	}
	i := int64(1)
	doc.Find("div.tr.site-listing").Each(func(index int, ele *goquery.Selection) {
		ele.Find("div.td.DescriptionCell").Each(func(j int, d *goquery.Selection) {
			if i <= topNumber {
				fmt.Printf("%d ", i)
				fmt.Println(strings.TrimSpace(d.First().Text()))
			}
			i++
		})
	})
}
