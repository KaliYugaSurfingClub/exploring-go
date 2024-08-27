package main

import (
	"github.com/PuerkitoBio/goquery"
	"net/http"
)

func findAllHrefs(host string) []string {
	result := make([]string, 0)

	client := &http.Client{}
	response, err := client.Get(host)

	if err != nil {
		return []string{}
	}

	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		return []string{}
	}

	result = append(result, host)

	document.Find("a").Each(func(index int, element *goquery.Selection) {
		if href, ok := element.Attr("href"); ok {
			result = append(result, findAllHrefs(host+href)...)
		}
	})

	return result
}

func Crawl(host string) []string {
	result := make([]string, 0)
	hrefs := findAllHrefs(host)

	result = append(result, hrefs...)

	return result
}

func main() {
	//Crawl("http://localhost8080/page1.html")
}
