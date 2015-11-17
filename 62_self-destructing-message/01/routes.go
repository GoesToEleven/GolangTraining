package selfdestruct

import (
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/nu7hatch/gouuid"
	"google.golang.org/appengine"
	"google.golang.org/appengine/memcache"
)

func init() {
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/msg/", handleMessage)
}

// create a message
func handleIndex(res http.ResponseWriter, req *http.Request) {
	ctx := appengine.NewContext(req)

	// form submit
	if req.Method == "POST" {
		msg := req.FormValue("message")
		key, _ := uuid.NewV4()
		// store the message in memcache
		item := &memcache.Item{
			Key:   key.String(),
			Value: []byte(msg),
		}
		err := memcache.Add(ctx, item)
		if err != nil {
			http.Error(res, err.Error(), 500)
			return
		}
		io.WriteString(res, `<!DOCTYPE html>
<html>
  <head>

  </head>
  <body>
    Here is your self-destructing secret message ID:
    <a href="/msg/`+key.String()+`">`+key.String()+`</a>
  </body>
</html>`)
	} else {

		// render the form
		io.WriteString(res, `<!DOCTYPE html>
  <html>
    <head>

    </head>
    <body>
      <form method="POST">
        <label>Message:
          <textarea name="message"></textarea>
        </label><br>
        <input type="submit">
      </form>
    </body>
  </html>`)
	}

}

// return a message based on its id
func handleMessage(res http.ResponseWriter, req *http.Request) {
	ctx := appengine.NewContext(req)
	// get key from URL
	key := strings.SplitN(req.URL.Path, "/", 3)[2]
	// get item from memcache
	item, err := memcache.Get(ctx, key)
	if err != nil {
		http.NotFound(res, req)
		return
	}

	// delete msg after it is viewed
	// this way:
	// memcache.Delete(ctx, key)
	// or this way:
	if item.Flags == 0 {
		item.Expiration = 10 * time.Second
		item.Flags = 1
		memcache.Set(ctx, item)
	}

	res.Write(item.Value)
}
