package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

func main() {
	pages := readPages("http://bakusai.com/thr_res/acode=3/ctgid=103/bid=454/tid=6995911/")
	pages = append(pages, "/thr_res/acode=3/ctgid=103/bid=454/tid=6995911/")
	for _, page := range pages {
		getMessages("http://bakusai.com" + page)
	}
}

func readPages(url string) []string {
	//Load the URL
	res, err := http.Get(url)
	if err != nil {
		log.Fatalf("%v", err)
		return nil
	}
	defer res.Body.Close()

	// Convert the designated charset HTML to utf-8 encoded HTML.
	utfBody := transform.NewReader(bufio.NewReader(res.Body), japanese.ShiftJIS.NewDecoder())

	// use utfBody using goquery
	doc, err := goquery.NewDocumentFromReader(utfBody)
	if err != nil {
		log.Fatalf("%v", err)
		return nil
	}
	pages := []string{}
	// use doc...
	doc.Find(".paging_numberlink").Find("a").Each(func(_ int, s *goquery.Selection) {
		page, exists := s.Attr("href")
		if exists {
			pages = append(pages, page)
		}
	})
	return pages
}

func getMessages(url string) {
	//Load the URL
	res, err := http.Get(url)
	if err != nil {
		log.Fatalf("%v", err)
	}
	defer res.Body.Close()

	// Convert the designated charset HTML to utf-8 encoded HTML.
	utfBody := transform.NewReader(bufio.NewReader(res.Body), japanese.ShiftJIS.NewDecoder())

	// use utfBody using goquery
	doc, err := goquery.NewDocumentFromReader(utfBody)
	if err != nil {
		log.Fatalf("%v", err)
	}
	// use doc...
	doc.Find(".resbody").Each(func(_ int, s *goquery.Selection) {
		text := s.Text()
		fmt.Println(text)
		fmt.Println("////////////////////")
	})
}
