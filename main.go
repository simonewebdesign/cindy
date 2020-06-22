package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/smtp"
	"os"
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
	resp, httpErr := http.Get("https://www.simonewebdesign.it/atom.xml")
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

	for i := 0; i < len(feed.Feed); i++ {
		fmt.Println("Entry Title: " + feed.Feed[i].Title)
		fmt.Println("Entry Link: " + feed.Feed[i].Link.Href)
		fmt.Println("---")
	}

	var latestPost = feed.Feed[0]

	fmt.Println("Authenticating...")
	auth := smtp.PlainAuth("", os.Getenv("CINDY_AUTH_EMAIL"), os.Getenv("CINDY_AUTH_PASSWORD"), os.Getenv("CINDY_SMTP_SERVER"))

	to := "hello@simonewebdesign.it"
	msg := []byte("To: " + to + "\r\n" +
		"Subject: The latest post is here\r\n" +
		"\r\n" +
		latestPost.Content +
		"Here’s the latest post. Hope you like it.\r\n")

	fmt.Println("Sending mail to " + to + "...")
	smtpErr := smtp.SendMail(os.Getenv("CINDY_SMTP_SERVER")+":"+os.Getenv("CINDY_SMTP_PORT"), auth, os.Getenv("CINDY_SENDER_EMAIL"), []string{to}, msg)
	if smtpErr != nil {
		log.Printf("Failed sending mail to `%s'; Error: %v", to, smtpErr)
	}
}
