package db

type Snippet struct {
	Html   string
	Title  string
	Author string
}

var Snippets []Snippet
