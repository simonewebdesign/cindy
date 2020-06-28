package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/smtp"
	"net/url"
	"os"
	"regexp"
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
	var mailBody = makeMessageBody(latestPost.Title, latestPost.Content, latestPost.Link.Href)

	if len(os.Args) > 1 {
		fmt.Println("Preview")
		fmt.Println(makeMessage("Test Sender <test-sender@example.com>", "test-recipient@example.com", latestPost.Title, mailBody))
		os.Exit(0)
	}

	fmt.Println("Authenticating...")
	auth := smtp.PlainAuth("", os.Getenv("CINDY_AUTH_USERNAME"), os.Getenv("CINDY_AUTH_PASSWORD"), os.Getenv("CINDY_SMTP_SERVER"))

	addresses, err := ioutil.ReadFile(os.Getenv("CINDY_ADDRESSES_PATH"))
	if err != nil {
		panic(err)
	}

	var rxEmail = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

	for idx, to := range strings.Split(string(addresses), "\n") {
		if len(to) > 254 || !rxEmail.MatchString(to) {
			fmt.Printf("\033[31m✗INVALID EMAIL: %v\033[0m\n", to)
			continue
		}

		msg := makeMessage(os.Getenv("CINDY_FROM"), to, latestPost.Title, mailBody)
		msg = strings.Replace(msg, "{{UNSUB_URL}}", os.Getenv("CINDY_UNSUB_URL")+url.QueryEscape(to), -1)

		log.Printf("[%d] Sending mail to `%s'...", idx, to)

		smtpErr := smtp.SendMail(os.Getenv("CINDY_SMTP_SERVER")+":"+os.Getenv("CINDY_SMTP_PORT"), auth, os.Getenv("CINDY_SENDER_EMAIL"), []string{to}, []byte(msg))
		if smtpErr != nil {
			log.Printf("\033[31m✗FAIL: %v\033[0m\n", smtpErr)
		} else {
			log.Printf("\033[32mOK\033[0m\n")
		}
	}
}

func makeMessage(sender string, recipient string, title string, body string) string {
	return "From: " + sender + "\r\n" +
		"To: " + recipient + "\r\n" +
		"Subject: New Post: " + title + "\r\n" +
		"Content-Type: text/html; charset=utf-8\r\n" +
		"\r\n" +
		body
}

func makeMessageBody(title string, content string, link string) string {
	dat, err := ioutil.ReadFile(os.Getenv("CINDY_TEMPLATE_PATH"))
	if err != nil {
		panic(err)
	}
	var s = string(dat)
	s = strings.ReplaceAll(s, "{{POST_TITLE}}", title)
	s = strings.ReplaceAll(s, "{{POST_CONTENT}}", content)
	s = strings.ReplaceAll(s, "{{POST_URL}}", link)
	return s
}
