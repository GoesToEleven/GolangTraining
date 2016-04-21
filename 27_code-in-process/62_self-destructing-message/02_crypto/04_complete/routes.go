package selfdestruct

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"golang.org/x/crypto/nacl/secretbox"

	"github.com/nu7hatch/gouuid"
	"google.golang.org/appengine"
	"google.golang.org/appengine/memcache"
)

func init() {
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/msg/", handleMessage)
}

func encrypt(decrypted string, password [32]byte) string {
	var nonce [24]byte
	io.ReadAtLeast(rand.Reader, nonce[:], 24)
	encrypted := secretbox.Seal(nil, []byte(decrypted), &nonce, &password)
	return fmt.Sprintf("%x:%x", nonce[:], encrypted)
}

func decrypt(encrypted string, password [32]byte) (string, error) {
	var nonce [24]byte
	parts := strings.SplitN(encrypted, ":", 2)
	if len(parts) < 2 {
		return "", fmt.Errorf("expected nonce")
	}
	bs, err := hex.DecodeString(parts[0])
	if err != nil || len(bs) != 24 {
		return "", fmt.Errorf("invalid nonce")
	}
	copy(nonce[:], bs)
	bs, err = hex.DecodeString(parts[1])
	if err != nil {
		return "", fmt.Errorf("invalid message")
	}
	decrypted, ok := secretbox.Open(nil, bs, &nonce, &password)
	if !ok {
		return "", fmt.Errorf("invalid message")
	}
	return string(decrypted), nil
}

func generatePassword() [32]byte {
	var password [32]byte
	io.ReadAtLeast(rand.Reader, password[:], 32)
	return password
}

// create a message
func handleIndex(res http.ResponseWriter, req *http.Request) {
	ctx := appengine.NewContext(req)

	// form submit
	if req.Method == "POST" {
		msg := req.FormValue("message")
		secretKey := generatePassword()
		encryptedMessage := encrypt(msg, secretKey)

		messageKey, _ := uuid.NewV4()
		// store the message in memcache
		item := &memcache.Item{
			Key:   messageKey.String(),
			Value: []byte(encryptedMessage),
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
    <a href="/msg/`+messageKey.String()+`?secret=`+fmt.Sprintf("%x", secretKey)+`">`+messageKey.String()+`</a>
    Send it to Peter Graves.
    <p>The encrypted message:</p>
    <p>`+encryptedMessage+`</p>
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

	var secretKey [32]byte
	bs, err := hex.DecodeString(req.FormValue("secret"))
	if err != nil || len(bs) != 32 {
		http.NotFound(res, req)
		return
	}
	copy(secretKey[:], bs)

	msg, err := decrypt(string(item.Value), secretKey)
	if err != nil {
		http.NotFound(res, req)
		return
	}

	// if this is the first time the message is viewed
	if item.Flags == 0 {
		item.Expiration = 30 * time.Second
		item.Flags = 1
		memcache.Set(ctx, item)
	}

	res.Write([]byte(msg))
}
