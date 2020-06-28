# 🌼 Cindy

Cindy is a command-line syndication tool. It checks an RSS feed for new entries and sends a bunch of emails on behalf of an email address of your choice.

The recipients are validated before an email is sent, according to a regular expression you can specify. Failures are noted at the end of the process, so you can decide what to do next (i.e. maybe retry sending, or delete the bogus address from your mailing list).

The template is simple HTML and can be fully customized in any way you like.

If you're a blogger, you can see Cindy as a replacement for a service like Mailchimp, but without all the complexity. At its core, Cindy takes the most recent RSS item and sends it out as a newsletter to your subscribers.


## Installation

### Via Homebrew (macOS)

    brew install simonewebdesign/tap/cindy

### Compile from source

You'll need [Go](https://golang.org/) to compile Cindy.

    go build cindy.go && ./cindy

    # or

    go run cindy.go


## Configuration

Cindy must be configured via environment variables. There are no default values, therefore these variables are all mandatory.

    CINDY_RSS_URL        # The URL to an RSS feed, e.g.: https://www.simonewebdesign.it/atom.xml
    CINDY_FROM           # The "From" header,  e.g.: "Weekly Newsletter" <news@example.com>
    CINDY_SENDER_EMAIL   # The sender address, e.g.: news@example.com

    CINDY_AUTH_USERNAME  # For authenticating yourself on your SMTP server
    CINDY_AUTH_PASSWORD  # Self explanatory

    CINDY_SMTP_SERVER    # The address of your SMTP server, e.g.: smtp.example.com
    CINDY_SMTP_PORT      # The port to use. It's usually either 587 or 465.

    CINDY_UNSUB_URL      # The URL to unsubscribe, e.g.: https://example.com/unsubscribe?email=
                         # The email will be appended to the string at runtime.

    CINDY_TEMPLATE_PATH  # Path to the HTML email template to be sent by Cindy
    CINDY_ADDRESSES_PATH # Path to the TXT file containing the list of emails separated
                           by a newline. There should be no newline at the end of this file.


## Frequently Asked Questions

### How do I provide my own email template?

You have total freedom over that. A good starting point could be [leemunroe/responsive-html-email-template](https://github.com/leemunroe/responsive-html-email-template). Once you have a template, don't forget to set the `CINDY_TEMPLATE_PATH` environment variable.

### How can I inject data into the email template?

Simply put the following placeholders into your template and they will be replaced with the actual values at runtime.

    {{POST_URL}}         # URL to the blog post entry
    {{POST_TITLE}}       # Title of the blog post
    {{POST_CONTENT}}     # HTML content
    {{UNSUB_URL}}        # URL to unsubscribe (i.e. the CINDY_UNSUB_URL environment variable)

### Can I preview the email before sending it?

Sure: just pass an extra shell argument (any will do). For example:

    cindy preview

This will print the full email that will be sent, complete with headers. You may want to save it to a local file, in order to preview it in your web browser:

    cindy preview > email-preview.html

When you're happy with it, just rerun Cindy without any arguments.
