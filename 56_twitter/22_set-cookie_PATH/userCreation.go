package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"
	"io/ioutil"
	"net/http"
)

func checkUserName(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	ctx := appengine.NewContext(req)
	bs, err := ioutil.ReadAll(req.Body)
	sbs := string(bs)
	log.Infof(ctx, "REQUEST BODY: %v", sbs)
	var user User
	key := datastore.NewKey(ctx, "Users", sbs, 0, nil)
	err = datastore.Get(ctx, key, &user)
	// if there is an err, there is NO user
	log.Infof(ctx, "ERR: %v", err)
	if err != nil {
		// there is an err, there is a NO user
		fmt.Fprint(res, "false")
		return
	} else {
		fmt.Fprint(res, "true")
	}
}

func createUser(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	ctx := appengine.NewContext(req)
	NewUser := User{
		Email:    req.FormValue("email"),
		UserName: req.FormValue("userName"),
		Password: req.FormValue("password"),
	}
	key := datastore.NewKey(ctx, "Users", NewUser.UserName, 0, nil)
	key, err := datastore.Put(ctx, key, &NewUser)
	// this is the only error checking I added; any others on this page needed?
	if err != nil {
		log.Errorf(ctx, "error adding todo: %v", err)
		http.Error(res, err.Error(), 500)
		return
	}

	// SET COOKIE
	http.SetCookie(res, &http.Cookie{
		Name:  "my-cookies",
		Value: "some value user creation WITH PATH / ",
		Path:  "/",
	})

	// redirect
	http.Redirect(res, req, "/", 302)
}

/*
1. Introduction

Securing cookies is an important subject.
Think about an authentication cookie. When the attacker is able to grab this cookie,
he can impersonate the user. This article describes HttpOnly and secure flags that can
enhance security of cookies.

2. HTTP, HTTPS and secure Flag

When HTTP protocol is used, the traffic is sent in plaintext.
It allows the attacker to see/modify the traffic (man-in-the-middle attack).
HTTPS is a secure version of HTTP – it uses SSL/TLS to protect the data
of the application layer. When HTTPS is used, the following properties are achieved:
authentication, data integrity, confidentiality. How are HTTP and HTTPS related to
a secure flag of the cookie?

Let’s consider the case of an authentication cookie. As was previously said,
stealing this cookie is equivalent to impersonating the user. When HTTP is used,
the cookie is sent in plaintext. This is fine for the attacker eavesdropping on the
communication channel between the browser and the server – he can grab the cookie and
impersonate the user.

Now let’s assume that HTTPS is used instead of HTTP.
HTTPS provides confidentiality. That’s why the attacker can’t see the cookie.
The conclusion is to send the authentication cookie over a secure channel so that
it can’t be eavesdropped. The question that might appear in this moment is:
why do we need a secure flag if we can use HTTPS?

Let’s consider the following scenario to answer this question.
The site is available over HTTP and HTTPS. Moreover, let’s assume that there is
an attacker in the middle of the communication channel between the browser and the server.
The cookie sent over HTTPS can’t be eavesdropped. However, the attacker can take advantage of
the fact that the site is also available over HTTP. The attacker can send the link to
the HTTP version of the site to the user. The user clicks the link and the HTTP request
is generated. Since HTTP traffic is sent in plaintext, the attacker eavesdrops on the
communication channel and reads the authentication cookie of the user. Can we allow
this cookie to be sent only over HTTPS? If this was possible, we would prevent the attacker
from reading the authentication cookie in our story. It turns out that it is possible,
and a secure flag is used exactly for this purpose – the cookie with a secure flag will
only be sent over an HTTPS connection.

3. HttpOnly Flag

In the previous section, it was presented how to protect the cookie from an attacker
eavesdropping on the communication channel between the browser and the server.
However, eavesdropping is not the only attack vector to grab the cookie.

Let’s continue the story with the authentication cookie and assume that XSS
(cross-site scripting) vulnerability is present in the application. Then the
attacker can take advantage of the XSS vulnerability to steal the authentication cookie.
Can we somehow prevent this from happening? It turns out that an HttpOnly flag can be used
to solve this problem. When an HttpOnly flag is used, JavaScript will not be able to read
this authentication cookie in case of XSS exploitation. It seems like we have achieved
the goal, but the problem might still be present when cross-site tracing (XST)
vulnerability exists (this vulnerability will be explained in the next section of the article)
 – the attacker might take advantage of XSS and enabled TRACE method to read
 the authentication cookie even if HttpOnly flag is used. Let’s see how XST works.

SOURCE:
http://resources.infosecinstitute.com/securing-cookies-httponly-secure-flags/
*/
