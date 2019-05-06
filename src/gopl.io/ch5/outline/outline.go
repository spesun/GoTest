package main

import (
	"os"
	"fmt"
	"golang.org/x/net/html"
	"net/http"
)

func main() {

	resp, err := http.Get("http://gopl.io")
	doc, err := html.Parse(resp.Body)
	//doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "outline: %v\n", err)
		os.Exit(1)
	}
	outline(nil, doc)
}
func outline(stack []string, n *html.Node) {
	if n.Type == html.ElementNode {
		stack = append(stack, n.Data) // push tag
		fmt.Println(stack)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		outline(stack, c)
	}
}

