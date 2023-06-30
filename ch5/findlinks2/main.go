package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	for _, url := range os.Args[1:] {
		links, err := findLinks(url)
		if err != nil {
			fmt.Printf("findlinks2: %v\n", err)
			continue
		}
		for _, link := range links {
			fmt.Println(link)
		}
	}
}

// findLinks performs and HTTP GET request for url, parses the
// response as HTML, and extracts and return the links.

func findLinks(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s: %v", url, resp.Status)
	}
	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("parsing %s: %v", url, err)
	}
	return visit(nil, doc), nil

}

func visit(links []string, n *html.Node) []string {
	if n == nil {
		return nil
	}
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}

	links = visit(links, n.FirstChild)
	return visit(links, n.NextSibling)
}

// CountWordsAndImages does an HTTP GET request for the HTML
// document url and returns the number of words and images in it.
// func countWordsAndImages(url string) (words, images int, err error) {
// 	resp, err := http.Get(url)
// 	if err != nil {
// 		return
// 	}
// 	doc, err := html.Parse(resp.Body)
// 	resp.Body.Close()
// 	if err != nil {
// 		err = fmt.Errorf("html.Parse error: %v", err)
// 		return
// 	}
// 	words, images, _ = countWordsAndImages(doc)
// 	return
// }
