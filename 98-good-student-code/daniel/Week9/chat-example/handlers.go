package chat

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"io"
	"net/http"

	"golang.org/x/net/context"

	"google.golang.org/appengine"
	"google.golang.org/appengine/channel"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/user"
)

// API is the API
type API struct {
	root string
}

func newAPI(root string) *API {
	return &API{
		root: root,
	}
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

func getClientID(email string) string {
	hash := md5.New()
	io.WriteString(hash, email)
	sum := hash.Sum(nil)
	return hex.EncodeToString(sum)
}

func saveNewConnection(ctx context.Context, clientID string) error {
	key := datastore.NewKey(ctx, "connection", clientID, 0, nil)
	val := struct{ Value string }{clientID}
	_, err := datastore.Put(ctx, key, &val)
	return err
}

func (api *API) handlePostChannel(res http.ResponseWriter, req *http.Request) error {
	ctx := appengine.NewContext(req)
	u := user.Current(ctx)
	clientID := getClientID(u.Email)
	token, err := channel.Create(ctx, clientID)
	if err != nil {
		return err
	}
	err = saveNewConnection(ctx, clientID)
	if err != nil {
		return err
	}
	return json.NewEncoder(res).Encode(token)
}

func (api *API) handlePostMessage(res http.ResponseWriter, req *http.Request) error {
	ctx := appengine.NewContext(req)

	var message struct{ Text string }
	err := json.NewDecoder(req.Body).Decode(&message)
	if err != nil {
		return err
	}

	query := datastore.NewQuery("connection")
	it := query.Run(ctx)
	for {
		var conn struct{ Value string }
		_, err := it.Next(&conn)
		if err == datastore.Done {
			break
		} else if err != nil {
			return err
		}
		err = channel.SendJSON(ctx, conn.Value, message)
		if err != nil {
			return err
		}
	}
	return nil
}
