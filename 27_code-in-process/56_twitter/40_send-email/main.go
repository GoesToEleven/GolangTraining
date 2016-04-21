package main

import (
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/mail"
	"google.golang.org/appengine/user"
)

func init() {
	http.HandleFunc("/sendmail", handleIndex)
}

func handleIndex(res http.ResponseWriter, req *http.Request) {
	ctx := appengine.NewContext(req)
	u := user.Current(ctx)

	msg := &mail.Message{
		Sender:  u.Email,
		To:      []string{"Todd McLeod <todd.mcleod@fresnocitycollege.edu>"},
		Subject: "See you tonight",
		Body:    "Don't forget our plans. Hark, 'til later.",
	}
	if err := mail.Send(ctx, msg); err != nil {
		log.Errorf(ctx, "Alas, my user, the email failed to sendeth: %v", err)
	}
}

/*

goapp deploy

appcfg.py --oauth2 update .

http://twitmock-1012.appspot.com/

*/
