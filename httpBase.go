package main

import (
	"bytes"
	"io/ioutil"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func Query(url string) *goquery.Document {
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		panic(err)
	}

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	r := bytes.NewReader(b)

	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		panic(err)
	}

	return doc
}
