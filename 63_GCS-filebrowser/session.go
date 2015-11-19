package filebrowser

import (
	"encoding/json"
	"net/http"

	"google.golang.org/appengine/memcache"

	"github.com/nu7hatch/gouuid"

	"golang.org/x/net/context"
)

type Session struct {
	ID                  string
	Bucket, Credentials string
}

func getSession(ctx context.Context, req *http.Request) Session {
	cookie, err := req.Cookie("sessionid")
	if err != nil || cookie.Value == "" {
		sessionID, _ := uuid.NewV4()
		cookie = &http.Cookie{
			Name:  "sessionid",
			Value: sessionID.String(),
		}
	}

	item, err := memcache.Get(ctx, cookie.Value)
	if err != nil {
		item = &memcache.Item{
			Key:   cookie.Value,
			Value: []byte(""),
		}
	}

	var session Session
	json.Unmarshal(item.Value, &session)
	session.ID = cookie.Value
	return session
}

func putSession(ctx context.Context, res http.ResponseWriter, session Session) {
	bs, err := json.Marshal(session)
	if err != nil {
		return
	}

	memcache.Set(ctx, &memcache.Item{
		Key:   session.ID,
		Value: bs,
	})

	http.SetCookie(res, &http.Cookie{
		Name:  "sessionid",
		Value: session.ID,
	})
}
