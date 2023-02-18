package main

import (
	"fmt"
	"io"
	"log"
	"strings"

	"net/http"

	"github.com/takunnithan/html_link_parser"
)

var visitedURLs map[string]bool
var newURLs []string
var baseURL string

func main() {
	fmt.Println("Site map!")
	visitedURLs = make(map[string]bool)
	baseURL = "https://crawler-test.com/"
	newURLs = append(newURLs, baseURL)
	buildSitemap()
	fmt.Println(visitedURLs)
}

func getHtmlSourceReader(url string) (io.ReadCloser, error) {
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return resp.Body, nil
}

func buildSitemap() {
	i := 0
	for {
		fmt.Println("length: ", len(newURLs))
		if i >= 100 {
			break
		}
		url := newURLs[i]
		htmlSourceReader, err := getHtmlSourceReader(url)
		if err != nil {
			visitedURLs[url] = true
			i = i + 1
			continue
		}
		links := html_link_parser.GetLinks(htmlSourceReader)
		for _, link := range links {
			newURL := link.Href
			if !strings.Contains(newURL, "https://") {
				newURL = baseURL + newURL
			}
			// fmt.Println("new URL: ", newURL)
			_, ok := visitedURLs[newURL]
			if !ok && strings.Contains(newURL, baseURL) {
				// if !ok {
				fmt.Println("URL Added....", newURL)
				newURLs = append(newURLs, newURL)
			}
		}
		visitedURLs[url] = true
		i = i + 1
		fmt.Println("Index: ", i)
		htmlSourceReader.Close()
	}
}

// Hard coded site URL
// Download the html source into a variable as byte slices
// Update the parser to use byte slices
// call the get links function with the html source
// Global / function level map to keep track of visited URLs
// A list for new URLs. An infinite loop which loops over the list and breaks when it is empty
// Construct the XML using the lib
// return
