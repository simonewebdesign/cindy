package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Feed is the entire RSS feed
type Feed struct {
	Feed []Entry `xml:"entry"`
}

// Entry is a blog post
type Entry struct {
	Title   string `xml:"title"`
	Link    Link   `xml:"link"`
	Updated string `xml:"updated"`
	ID      string `xml:"id"`
	Content string `xml:"content"`
}

// Link is a URL
type Link struct {
	Href string `xml:"href,attr"`
}

func main() {
	resp, err := http.Get("https://www.simonewebdesign.it/atom.xml")
	if err != nil {
		log.Printf("Error reading body: %v", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	var feed Feed
	xml.Unmarshal(body, &feed)

	for i := 0; i < len(feed.Feed); i++ {
		fmt.Println("Entry Title: " + feed.Feed[i].Title)
		fmt.Println("Entry Link: " + feed.Feed[i].Link.Href)
		fmt.Println("---")
	}
}
