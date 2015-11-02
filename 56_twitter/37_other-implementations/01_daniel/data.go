package twitter

import (
	"fmt"
	"strings"

	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/mail"
)

func getProfileByEmail(ctx context.Context, email string) (*profile, error) {
	key := datastore.NewKey(ctx, "profile", email, 0, nil)
	var p profile
	err := datastore.Get(ctx, key, &p)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func getProfileByUsername(ctx context.Context, usr string) (*profile, error) {
	query := datastore.NewQuery("profile")
	p := []profile{}
	_, err := query.Filter("Username =", usr).GetAll(ctx, &p)
	if err != nil {
		return nil, err
	}
	if len(p) == 0 {
		return nil, datastore.ErrNoSuchEntity
	} else if len(p) > 1 {
		return nil, datastore.ErrInvalidKey
	}
	return &p[0], nil
}

func itemIn(item string, list []string) bool {
	for _, v := range list {
		if item == v {
			return true
		}
	}
	return false
}

func getMultiTweets(ctx context.Context, emails []string) ([]tweet, error) {
	query := datastore.NewQuery("Tweets")
	t := []tweet{}
	query = query.Order("-SubmitTime")
	it := query.Run(ctx)
	for {
		var i tweet
		key, err := it.Next(&i)
		if err == datastore.Done {
			break
		} else if err != nil {
			return nil, err
		}
		p, err := getProfileByEmail(ctx, key.Parent().StringID())
		if err != nil {
			return nil, err
		}
		if !itemIn(p.Username, emails) {
			continue
		}
		i.Username = p.Username
		t = append(t, i)
	}
	return t, nil
}

func getTweets(ctx context.Context, email string) ([]tweet, error) {
	query := datastore.NewQuery("Tweets")
	t := []tweet{}
	query = query.Order("-SubmitTime")
	if email != "" {
		key := datastore.NewKey(ctx, "profile", email, 0, nil)
		query = query.Ancestor(key)
	}
	keys, err := query.GetAll(ctx, &t)
	if err != nil {
		return nil, err
	}
	for i := range t {
		p, err := getProfileByEmail(ctx, keys[i].Parent().StringID())
		if err != nil {
			return nil, err
		}
		t[i].Username = p.Username
	}
	return t, nil
}

func createProfile(ctx context.Context, username, email string) error {
	key := datastore.NewKey(ctx, "profile", email, 0, nil)
	p := profile{
		Username: username,
		Email:    email,
	}
	_, err := datastore.Put(ctx, key, &p)
	return err
}

func emailUser(ctx context.Context, username string, t *tweet) error {
	to, err := getProfileByUsername(ctx, username)
	if err == datastore.ErrNoSuchEntity {
		return nil
	} else if err != nil {
		return err
	}
	msg := &mail.Message{
		Sender:  "noreply@twitter-1012.appspotmail.com",
		To:      []string{fmt.Sprintf("%s <%s>", to.Username, to.Email)},
		Subject: fmt.Sprintf("Mention from %s", t.Username),
		Body:    fmt.Sprintf("%s has mentioned you in a tweet: %s", t.Username, t.Message),
	}
	err = mail.Send(ctx, msg)
	return err
}

func postTweet(ctx context.Context, t *tweet, email string) error {
	words := strings.Fields(t.Message)
	for _, w := range words {
		if w[0] == '@' {
			if err := emailUser(ctx, w[1:], t); err != nil {
				return err
			}
		}
	}

	profileKey := datastore.NewKey(ctx, "profile", email, 0, nil)
	key := datastore.NewIncompleteKey(ctx, "Tweets", profileKey)
	_, err := datastore.Put(ctx, key, t)
	return err
}

func addFollower(ctx context.Context, currentUser *profile, newFollowed string) error {
	if itemIn(newFollowed, currentUser.Following) || newFollowed == currentUser.Username {
		return nil
	}
	currentUser.Following = append(currentUser.Following, newFollowed)
	key := datastore.NewKey(ctx, "profile", currentUser.Email, 0, nil)
	_, err := datastore.Put(ctx, key, currentUser)
	return err
}

func removeFollower(ctx context.Context, currentUser *profile, oldFollowed string) error {
	needsStore := false
	for i, f := range currentUser.Following {
		if f == oldFollowed {
			currentUser.Following = append(currentUser.Following[:i], currentUser.Following[i+1:]...)
			needsStore = true
			break
		}
	}
	if !needsStore {
		return nil
	}
	key := datastore.NewKey(ctx, "profile", currentUser.Email, 0, nil)
	_, err := datastore.Put(ctx, key, currentUser)
	return err
}
