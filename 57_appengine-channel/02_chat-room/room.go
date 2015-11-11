package chat

import (
	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
)

type Room struct {
	Emails []string
}

func getRoom(ctx context.Context) (*Room, error) {
	key := datastore.NewKey(ctx, "Room", "GLOBAL", 0, nil)
	var room Room
	datastore.Get(ctx, key, &room)
	return &room, nil
}

func putRoom(ctx context.Context, room *Room) error {
	key := datastore.NewKey(ctx, "Room", "GLOBAL", 0, nil)
	_, err := datastore.Put(ctx, key, room)
	return err
}
