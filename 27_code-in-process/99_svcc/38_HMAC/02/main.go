package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"github.com/nu7hatch/gouuid"
	"io"
)

func main() {
	sessionID, _ := uuid.NewV4()
	fmt.Println(sessionID.String())
	code := getCode(sessionID.String())
	fmt.Println(code)
	// we get the same code here b/c the string is unchanged
	code = getCode(sessionID.String())
	fmt.Println(code)
	// we get a different code here b/c the string is unchanged
	code = getCode(sessionID.String() + "some other stuff")
	fmt.Println(code)

}

func getCode(data string) string {
	h := hmac.New(sha256.New, []byte("this can be anything we want it to be it is our private key"))
	io.WriteString(h, data)
	return fmt.Sprintf("%x", h.Sum(nil))
}
