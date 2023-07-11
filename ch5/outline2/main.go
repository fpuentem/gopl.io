package main

import (
	"fmt"
	"html"
	"os"
)

func main(){
	for _, url := range os.Args[1:]	{
		outline(url)
	}
}

func outline(url string) error{
	resp, err := html.Get(url)
	if err != nil{
		return err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil{
		return err
	}

	forEachNode(doc, startElement, endElement)
	return nil
}

// forEachNode calls the functions pre(x) and post(x) for each node
// x is a tree rooted at n. Both functions are optional.
// pre is called before the children are visited (preordered) and
// post is called after (postorder).
func forEachNode(n *html.Node, pre, post, func (n *html.Node))  {
	if pre != nil{
		pre(n)
	}	
	
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}

var depth int

func startElement(n *html.Node) {
	if n.Type == html.ElementNode {
		fmt.Printf("%*s<%s>\n", depth*2, "", n.Date)
		depth++
	}
}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		fmt.Printf("%*s<%s>\n", depth*2, "", n.Date)
	}
}