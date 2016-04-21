package main

import (
	"fmt"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/mail"
	"net/http"
)

func followedEmail(w http.ResponseWriter, r *http.Request, recip string) {
	ctx := appengine.NewContext(r)
	msg := &mail.Message{
		Sender:  "TwitClone Support <support@example.com>",
		To:      []string{recip},
		Subject: "You are being followed",
		Body:    fmt.Sprintf(confirmMessage),
	}
	if err := mail.Send(ctx, msg); err != nil {
		log.Errorf(ctx, "Couldn't send email: %v", err)
	}
}

const confirmMessage = `
Someone is now following you. Don't say you didn't ask for it.
`
