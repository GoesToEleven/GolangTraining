package main

import (
	"fmt"
	"strings"
	"time"

	"golang.org/x/net/context"

	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/mail"
)

type Scream struct {
	ID       int64
	Username string
	Email    string
	Message  string
	Time     time.Time
}

func createScream(ctx context.Context, scream *Scream) error {
	key := datastore.NewIncompleteKey(ctx, "Scream", nil)
	_, err := datastore.Put(ctx, key, scream)
	if err != nil {
		return err
	}
	if strings.Contains(scream.Message, "@") {
		getUser := scanString(scream.Message)
		profileMention, err := getProfileByUsername(ctx, getUser)
		if err != nil {
			return err
		}
		return sendMail(ctx, profileMention)
	}
	return nil
}

func getScreams(ctx context.Context, username string) ([]*Scream, error) {
	screams := make([]*Scream, 0)
	q := datastore.NewQuery("Scream")
	if username != "" {
		q = q.Filter("Username=", username)
	}
	iterator := q.Run(ctx)
	for {
		var scream Scream
		key, err := iterator.Next(&scream)
		if err == datastore.Done {
			break
		} else if err != nil {
			return nil, err
		}
		scream.ID = key.IntID()
		screams = append(screams, &scream)
	}
	return screams, nil
}

type Profile struct {
	Username  string
	Email     string
	Following []string
}

func getProfile(ctx context.Context, email string) (*Profile, error) {
	key := datastore.NewKey(ctx, "Profile", email, 0, nil)
	var profile Profile
	return &profile, datastore.Get(ctx, key, &profile)
}

func getProfileByUsername(ctx context.Context, username string) (*Profile, error) {
	q := datastore.NewQuery("Profile").Filter("Username =", username).Limit(1)
	var profiles []Profile
	_, err := q.GetAll(ctx, &profiles)
	if err != nil {
		return nil, err
	}
	if len(profiles) == 0 {
		return nil, fmt.Errorf("profile not found")
	}
	return &profiles[0], nil
}

func createProfile(ctx context.Context, profile *Profile) error {
	key := datastore.NewKey(ctx, "Profile", profile.Email, 0, nil)
	_, err := datastore.Put(ctx, key, profile)
	return err
}
func follow(ctx context.Context, email string, followee string) error {
	profile, err := getProfile(ctx, email)
	if err != nil {
		return err
	}
	for _, f := range profile.Following {
		if f == followee {
			return nil
		}
		if followee == "" {
			return nil
		}
	}
	profile.Following = append(profile.Following, followee)
	return createProfile(ctx, profile)
}

func scanString(message string) string {
	user := ""
	words := strings.Split(message, " ")
	for _, value := range words {
		if strings.HasPrefix(value, "@") {
			user = value[1:len(value)]
		}
	}
	return user
}
func sendMail(ctx context.Context, mentionProfile *Profile) error {
	msg := &mail.Message{
		Sender:  "<admin@secret-spark-101320.appspotmail.com>",
		To:      []string{mentionProfile.Email},
		Subject: "Someone mentioned you!",
		Body:    fmt.Sprintf("Someone mentioned you!"),
	}
	return mail.Send(ctx, msg)
}
