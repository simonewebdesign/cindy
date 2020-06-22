package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/smtp"
	"os"
	"strings"
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
	resp, httpErr := http.Get(os.Getenv("CINDY_RSS_URL"))
	if httpErr != nil {
		log.Printf("HTTP error: %v", httpErr)
		return
	}
	defer resp.Body.Close()

	body, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		log.Printf("Error reading body: %v", readErr)
		return
	}

	var feed Feed
	xml.Unmarshal(body, &feed)

	var latestPost = feed.Feed[0]

	fmt.Println("Authenticating...")
	auth := smtp.PlainAuth("", os.Getenv("CINDY_AUTH_USERNAME"), os.Getenv("CINDY_AUTH_PASSWORD"), os.Getenv("CINDY_SMTP_SERVER"))

	addresses, err := ioutil.ReadFile("addresses.txt")
	if err != nil {
		panic(err)
	}

	for idx, to := range strings.Split(string(addresses), "\n") {
		msg := []byte("To: " + to + "\r\n" +
			"Subject: New Post: " + latestPost.Title + "\r\n" +
			"Content-Type: text/html; charset=utf-8\r\n" +
			"\r\n" +
			wrapInTemplate(latestPost.Content, latestPost.Link.Href))

		log.Printf("[%d] Sending mail to `%s'...", idx, to)

		smtpErr := smtp.SendMail(os.Getenv("CINDY_SMTP_SERVER")+":"+os.Getenv("CINDY_SMTP_PORT"), auth, os.Getenv("CINDY_SENDER_EMAIL"), []string{to}, msg)
		if smtpErr != nil {
			log.Printf("\033[31m✗FAIL: %v\033[0m\n", smtpErr)
		} else {
			log.Printf("\033[32mOK\033[0m\n")
		}
	}
}

func wrapInTemplate(content string, link string) string {
	dat, err := ioutil.ReadFile("template.html")
	if err != nil {
		panic(err)
	}
	var s = string(dat)
	s = strings.ReplaceAll(s, "{{POST_CONTENT}}", content)
	s = strings.ReplaceAll(s, "{{POST_LINK}}", link)
	return s
}