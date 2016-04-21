ENCRYPT / HASH passwords
- make your hash slow
--- prevents brute force
- salt your hash
--- store your salt with your hash so you can use it
--- each salt is unique to each password
- Corey used bcrypt to do this:
golang.org/x/crypto/bcrypt
- golang.org/x/ are maybe experimental projects made by google and not yet in standard library
- you could also us a second salt, called a pepper sometimes
-- this is a code unique to the whole website
-- it's not stored with the passwords
https://github.com/Saxleader/fall2015/tree/master/code-review_improvement

/////////////////////////////

https://github.com/FelixVicis/f15_advWeb_amills/tree/master/twitterclone
found error on line 63 of api.go, SessionData was misspelled
Started making global messages, Toots

type Toot struct {
	UserName string
	Message  string
}
- also add in
-- time of tweet
-- make sure username unique

Ran into issue with serveTemplate, may need to extend functionality
umm. server 500 error. welp. something went wonky.


/////////////////////////////

type Post struct {
	Username string
	Post     string
}


posting user tweets
using jquery
https://github.com/herrschwartz/AdvWeb


//////////////////////////////

prevent brute force
- check IP address
-- store in memcache
- limited login attempts per account for certain time period

https://godoc.org/net/http#Request
// RemoteAddr allows HTTP servers and other software to record
    // the network address that sent the request, usually for
    // logging. This field is not filled in by ReadRequest and
    // has no defined format. The HTTP server in this package
    // sets RemoteAddr to an "IP:port" address before invoking a
    // handler.
    // This field is ignored by the HTTP client.
    RemoteAddr string

- package subtle

/////////////

authenticate user when they create account
- send email and have confirmation link to click

//////////////

xss


