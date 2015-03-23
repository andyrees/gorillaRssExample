package main

import (
	"fmt"
	"time"

	"github.com/gorilla/feeds"
)

func main() {
	now := time.Now()
	feed := &feeds.Feed{
		Title:       "andyrees.github.io blog",
		Link:        &feeds.Link{Href: "http://andyrees.github.io/"},
		Description: "Discussion about programming in go and transcoding",
		Author:      &feeds.Author{"Andy Rees", "https://twitter.com/andygolang"},
		Created:     now,
	}

	feed.Items = []*feeds.Item{
		{
			Title:       "Francesc Campoy Flores blog",
			Link:        &feeds.Link{Href: "http://www.campoy.cat/"},
			Description: "Developer Advocate at Google",
			Author:      &feeds.Author{"Francesc Campoy Flores", "user@name.net"},
			Created:     now,
		},
		{
			Title:       "Official Go blog",
			Link:        &feeds.Link{Href: "http://blog.golang.org/"},
			Description: "The Go Blog",
			Created:     now,
		},
	}

	atom, _ := feed.ToAtom()
	rss, _ := feed.ToRss()

	fmt.Println(atom, "\n", rss)
}
