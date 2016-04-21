OAUTH
https://en.wikipedia.org/wiki/OAuth

GOLANG DOCs
https://godoc.org/golang.org/x/oauth2

TWITTER EXAMPLE
signin with twitter

(1)
learn about twitter oauth
https://dev.twitter.com/web/sign-in
https://dev.twitter.com/web/sign-in/desktop-browser

(2)
this is what we want to build - oauth like this


(3)
twitter oauth docs


EXAMPLE:
(customized for my registered application, which we'll do next)
--local development---
--deployment---

(4)
register your application

(5)
Redirect users to request GitHub access

(6)
create a callback to receive authorization code

(7)
exchange authorization code for an access token
***AUTHORIZATION CODE***
- the authorization code is the AUTHORIZATION OF THE USER
-- requires our github oauth api Client ID
***ACCESS TOKEN***
- the access token is the AUTHORIZATION OF THE APPLICATION
-- requires our github oauth api Client Secret
-- we wouldn't want to pass our Client Secret in the URL


(8)
Get user information
EMAIL EXAMPLE in code

(9)
Application flow:
- now that you have someone's VERIFIED email (github verifies emails)
-- you can associate that email with a user's account on your site