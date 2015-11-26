package chat

import (
	"encoding/json"
	"net/http"

	"fmt"
	"google.golang.org/appengine"
	"google.golang.org/appengine/channel"
	"google.golang.org/appengine/user"
)

// API handles api calls
type API struct {
	root string
}

// NewAPI creates a new API, root should be set to the root url for the API
func NewAPI(root string) *API {
	api := &API{
		root: root,
	}
	return api
}

func (api *API) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	endpoint := req.URL.Path[len(api.root):]
	method := req.Method

	var err error
	switch endpoint {
	case "channels":
		switch method {
		case "POST":
			err = api.handlePostChannel(res, req)
		default:
			http.Error(res, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
	case "messages":
		switch method {
		case "POST":
			err = api.handlePostMessage(res, req)
		default:
			http.Error(res, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
	default:
		http.NotFound(res, req)
		return
	}

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(err.Error())
	}
}

// create the channel
func (api *API) handlePostChannel(res http.ResponseWriter, req *http.Request) error {
	ctx := appengine.NewContext(req)
	u := user.Current(ctx)
	tok, err := channel.Create(ctx, u.Email)
	if err != nil {
		return err
	}
	room, err := getRoom(ctx)
	if err != nil {
		return err
	}
	room.Emails = append(room.Emails, u.Email)
	err = putRoom(ctx, room)
	if err != nil {
		return err
	}
	json.NewEncoder(res).Encode(tok)
	return nil
}

func (api *API) handlePostMessage(res http.ResponseWriter, req *http.Request) error {
	ctx := appengine.NewContext(req)
	type Message struct {
		Text string
	}
	var message Message
	err := json.NewDecoder(req.Body).Decode(&message)
	if err != nil {
		return err
	}
	fmt.Println("***********IS THE MESSAGE HERE: ", message)
	// u.Email == "test@example.com"
	room, err := getRoom(ctx)
	if err != nil {
		return err
	}
	for _, email := range room.Emails {
		err = channel.SendJSON(ctx, email, message)
		if err != nil {
			return err
		}
	}
	json.NewEncoder(res).Encode(true)
	return nil
}
