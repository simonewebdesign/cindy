# 🌼 Cindy

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

### Can I preview the email before sending it?

Sure: simply pass an extra shell argument (any will do). For example:

    cindy --preview

This will print the full email that will be sent, complete with headers. You may want to save it to a local file, in order to preview it in your web browser:

    cindy --preview > email-preview.html

When you're happy with it, just rerun the tool without the argument.


## Build from source and run

    go build main.go
    ./main

    # or

    go run main.go


## Caveats

### Don't forget to inline the CSS

It's important to also inline the CSS before sending the email:

    node_modules/.bin/juice template.html template.html

### License and credits

The template was originally forked from [leemunroe/responsive-html-email-template](https://github.com/leemunroe/responsive-html-email-template).

    The MIT License (MIT)

    Copyright (c) 2020 simonewebdesign

    Permission is hereby granted, free of charge, to any person obtaining a copy
    of this software and associated documentation files (the "Software"), to deal
    in the Software without restriction, including without limitation the rights
    to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
    copies of the Software, and to permit persons to whom the Software is
    furnished to do so, subject to the following conditions:

    The above copyright notice and this permission notice shall be included in all
    copies or substantial portions of the Software.

    THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
    IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
    FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
    AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
    LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
    OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
    SOFTWARE.
