// prettyhtml reads a valid utf-8 encoded HTML doc and pretty prints it to stdout.

package main

import (
	"log"
	"os"

	"golang.org/x/net/html"
)

// forEachNode is a function that is designed to traverse an HTML tree,
// visiting each node and calling functions before and after visiting the node.
func forEachNode(n *html.Node, pre, post func(n *html.Node)) {

}

// prettyPrint prints out an HTML document with proper indentation.
func prettyPrint(data *os.File) {
	htmldoc, err := html.Parse(data)
	if err != nil {
		log.Printf("Error while parsing HTML document: %s", err)
		os.Exit(1)
	}

}

func main() {

	filename := os.Args[1]

	data, err := os.Open(filename)
	if err != nil {
		log.Printf("Error while opening HTML file: %s", err)
		os.Exit(1)
	}

	prettyPrint(data)
}
