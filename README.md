# Cindy

Cindy is a command-line syndication tool. It checks an RSS feed for new entries and sends a bunch of emails on behalf of an email address of your choice.

The recipients are validated before an email is sent, according to a regular expression you can specify. Failures are noted at the end of the process, so you can decide what to do next (i.e. maybe retry sending, or delete the bogus address from your mailing list).

The template is simple HTML and can be fully customized in any way you like.

If you're a blogger, you can see Cindy as a replacement for a service like Mailchimp, but without all the complexity. At its core, Cindy takes the most recent RSS item and sends it out as a newsletter to your subscribers.


## Installation

Simply clone or download this repo. You'll need [Go](https://golang.org/) to compile Cindy.


## Configuration

Cindy must be configured via environment variables. There are no default values, therefore these variables are all mandatory.

    CINDY_RSS_URL        # The URL to an RSS feed, e.g.: https://www.simonewebdesign.it/atom.xml
    CINDY_SENDER_EMAIL   # The "From" address, e.g.: no-reply@example.com

    CINDY_AUTH_USERNAME  # For authenticating yourself on your SMTP server
    CINDY_AUTH_PASSWORD  # Self explanatory

    CINDY_SMTP_SERVER    # The address of your SMTP server, e.g.: smtp.example.com
    CINDY_SMTP_PORT      # The port to use. It's usually either 587 or 465.

### How do I provide a list of email addresses?

Cindy expects an `addresses.txt` file containing the list of emails separated by a newline. There should be no newline at the end of the file.


## Build from source and run

    go build main.go
    ./main

    # or

    go run main.go
