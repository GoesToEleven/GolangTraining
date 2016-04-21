package dropbox

import (
	"encoding/json"
	"net/http"

	"github.com/nu7hatch/gouuid"
	"golang.org/x/net/context"
	"google.golang.org/appengine/memcache"
)

const cookieName = "sessionid"

type session struct {
	ID    string `json:"-"`
	State string
}

func getSession(ctx context.Context, req *http.Request) session {
	cookie, err := req.Cookie(cookieName)
	if err != nil || cookie.Value == "" {
		uid, _ := uuid.NewV4()
		cookie = &http.Cookie{
			Name:  cookieName,
			Value: uid.String(),
		}
	}

	item, err := memcache.Get(ctx, cookie.Value)
	if err != nil {
		item = &memcache.Item{
			Key:   cookie.Value,
			Value: []byte(""),
		}
	}

	var s session
	json.Unmarshal(item.Value, &s)
	s.ID = cookie.Value
	return s
}

func putSession(ctx context.Context, res http.ResponseWriter, s session) error {
	bs, err := json.Marshal(s)
	if err != nil {
		return err
	}

	err = memcache.Set(ctx, &memcache.Item{
		Key:   s.ID,
		Value: bs,
	})
	if err != nil {
		return err
	}

	http.SetCookie(res, &http.Cookie{
		Name:  cookieName,
		Value: s.ID,
	})

	return nil
}
