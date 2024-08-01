package main

type Snippet struct {
	Html   string
	Title  string
	Author string
}

var snippets []Snippet

func AddSnippet(snippet Snippet) {
	snippets = append(snippets, snippet)
}

func GetLatestSnippets(count int) []Snippet {
	return createReversed(snippets[max(len(snippets)-count-1, 0):])
}

func SeedSnippets() {
	snippet1 := Snippet{
		Html:   "<h1>hello</h1><p>meow</p>",
		Title:  "hello snippet",
		Author: "Bob",
	}

	snippet2 := Snippet{
		Html:   "<h1>hello</h1><p>meow2</p>",
		Title:  "hello snippet2",
		Author: "Bob",
	}

	snippet3 := Snippet{
		Html:   "<h1>hello</h1><p>meow3</p>",
		Title:  "hello snippet3",
		Author: "Bob",
	}

	snippet4 := Snippet{
		Html:   "<h1>hello</h1><p>meow4</p>",
		Title:  "hello snippet4",
		Author: "Bob",
	}
	snippets = []Snippet{snippet1, snippet2, snippet3, snippet4}
}
