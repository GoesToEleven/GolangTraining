OAUTH
https://en.wikipedia.org/wiki/OAuth

GOLANG DOCs
https://godoc.org/golang.org/x/oauth2

GITHUB EXAMPLE

(1)
learn about github oauth
https://help.github.com/articles/connecting-with-third-party-applications/

(2)
this is what we want to build - oauth like this
https://forum.golangbridge.org/
http://codepen.io/
https://jsbin.com/

(3)
github oauth docs
https://developer.github.com/v3/oauth/
https://developer.github.com/v3/oauth/#web-application-flow

EXAMPLE:
(customized for my registered application, which we'll do next)
--local development---
https://github.com/login/oauth/authorize?client_id=fbbaa8ce5c394b7c3198&redirect_uri=http://localhost:8080/oauth2callback&state=233453423232
--deployment---
https://github.com/login/oauth/authorize?client_id=fbbaa8ce5c394b7c3198&redirect_uri=http://learning-1130.appspot.com/oauth2callback&state=233453423232

(4)
register your application
https://github.com/settings/applications/new
see pic_01.png, pic_02.png

(5)
Redirect users to request GitHub access
https://developer.github.com/v3/oauth/#web-application-flow

(6)
create a callback to receive authorization code

(7)
exchange authorization code for an access token
***AUTHORIZATION CODE***
- the authorization code is the AUTHORIZATION FROM THE USER
-- requires our github oauth api Client ID
***ACCESS TOKEN***
- the access token is the AUTHORIZATION OF THE APPLICATION
-- requires our github oauth api Client Secret
-- we wouldn't want to pass our Client Secret in the URL

https://youtu.be/oxogqJiFVYI?t=2547

(8)
Get user information
EMAIL EXAMPLE in code

(9)
Application flow:
- now that you have someone's VERIFIED email (github verifies emails)
-- you can associate that email with a user's account on your site

----

XSFR
https://en.wikipedia.org/wiki/Cross-site_request_forgery

Bank Example:
Someone is already logged into their bank
Bank uses GET (url) to send parameters
Hacker creates a url that transfers money to his account
Hacker puts that url in an img tag on web pages
when people view that webpage, that img tag makes a GET request
- not for an img, but just going to that url
- an img tag automatically fires the GET request
-- give me this img
- but instead of an img, they've made the GET request to transfer funds
bada-bing, bada-bang!
DEFEAT:
unique code
-github "state"
-you add a code to all URLs
-any requests must have this unique code
-it's impossible for others to know this beforehand
POST
-now hacker has to submit to your server using the POST method
-requires having the user fill out a form

https://youtu.be/oxogqJiFVYI?t=1001

----


