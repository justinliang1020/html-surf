package main

import (
	"html/template"
	"strings"
)

type Snippet struct {
	Html   string
	Title  string
	Author string
}

var snippets []Snippet

func AddSnippet(html string, title string, author string) {
	// Use template.HTMLEscapeString to escape the HTML
	escapedHtml := template.HTMLEscapeString(html)

	// Replace " with &quot; as HTMLEscapeString doesn't do this
	escapedHtml = strings.ReplaceAll(escapedHtml, `"`, "&quot;")

	snippets = append(snippets, Snippet{Html: escapedHtml, Title: title, Author: author})
}

func GetLatestSnippets(count int) []Snippet {
	return createReversed(snippets[max(len(snippets)-count-1, 0):])
}

func SeedSnippets() {
	snippet1 := Snippet{
		Html:   "<h1>hello world!</h1>",
		Title:  "hello world snippet",
		Author: "Justin",
	}

	snippet2 := Snippet{
		Html: `
    <style>
        .cat {
            font-family: monospace;
            font-size: 20px;
            white-space: pre;
            text-align: center;
            animation: dance 1s infinite alternate;
        }

        @keyframes dance {
            0% { transform: translateY(0) rotate(0deg); }
            25% { transform: translateY(-10px) rotate(-5deg); }
            50% { transform: translateY(0) rotate(0deg); }
            75% { transform: translateY(-10px) rotate(5deg); }
            100% { transform: translateY(0) rotate(0deg); }
        }
    </style>
    <div class="cat">
  /\___/\
  ( o   o )
  (  =^=  )
  (        )
  (         )
  (          )))))))))))
    </div>
    `,
		Title:  "dancing cat",
		Author: "claude",
	}

	snippet3 := Snippet{
		Html: `
<style>
    body {
        display: flex;
        justify-content: center;
        align-items: center;
        height: 100vh;
        background-color: #f0f0f0;
        font-family: Arial, sans-serif;
    }
    .taco {
        font-size: 100px;
        animation: dance 0.5s infinite alternate;
    }
    @keyframes dance {
        from { transform: rotate(-20deg); }
        to { transform: rotate(20deg); }
    }
    .message {
        position: absolute;
        bottom: 20px;
        font-size: 24px;
        color: #333;
    }
</style>
<div class="taco">🌮</div>
<div class="message">Let's taco 'bout how awesome you are!</div>
    `,
		Title:  "dancing taco",
		Author: "claude",
	}

	snippets = []Snippet{}
	AddSnippet(snippet1.Html, snippet1.Title, snippet1.Author)
	AddSnippet(snippet2.Html, snippet2.Title, snippet2.Author)
	AddSnippet(snippet3.Html, snippet3.Title, snippet3.Author)
}
