package main

import (
	"fmt"
)

func main() {

	msgCh := make(chan message, 1)
	errCh := make(chan failedMessage, 1)

	select {
	case receivedMsg := <-msgCh:
		fmt.Println(receivedMsg)
	case receivedError := <-errCh:
		fmt.Println(receivedError)
	default:
		fmt.Println("No messages received")
	}

}

type message struct {
	To      []string
	From    string
	Content string
}

type failedMessage struct {
	ErrorMessage    string
	OriginalMessage message
}

/* 3.5.1 - demo setup
package main

import (
	"fmt"
)

func main() {

	msgCh := make(chan message, 1)
	errCh := make(chan failedMessage, 1)

	msg := message{
		To: []string{"bilbo@underhill.me"},
		From: "gandalf@whitecouncil.org",
		Content: "Keep it secret, keep it safe",
	}

	failedMessage := failedMessage{
		ErrorMessage: "message intercepted by black rider",
		OriginalMessage: message{},
	}

	msgCh <- msg
	errCh <- failedMessage

	fmt.Println(<-msgCh)
	fmt.Println(<-errCh)

}

type message struct {
	To []string
	From string
	Content string
}

type failedMessage struct {
	ErrorMessage string
	OriginalMessage message
}
*/

/* 3.5.2 - add select, use first
package main

import (
	"fmt"
)

func main() {

	msgCh := make(chan message, 1)
	errCh := make(chan failedMessage, 1)

	msg := message{
		To: []string{"bilbo@underhill.me"},
		From: "gandalf@whitecouncil.org",
		Content: "Keep it secret, keep it safe",
	}

	msgCh <- msg

	select {
		case receivedMsg := <- msgCh:
			fmt.Println(receivedMsg)
		case receivedError := <- errCh:
			fmt.Println(receivedError)
	}

}

type message struct {
	To []string
	From string
	Content string
}

type failedMessage struct {
	ErrorMessage string
	OriginalMessage message
}
*/

/* 3.5.3 - use second select case
package main

import (
	"fmt"
)

func main() {

	msgCh := make(chan message, 1)
	errCh := make(chan failedMessage, 1)

	msg := message{
		To: []string{"bilbo@underhill.me"},
		From: "gandalf@whitecouncil.org",
		Content: "Keep it secret, keep it safe",
	}

	failedMessage := failedMessage{
		ErrorMessage: "message intercepted by black rider",
		OriginalMessage: message{},
	}

	errCh <- failedMessage

	select {
		case receivedMsg := <- msgCh:
			fmt.Println(receivedMsg)
		case receivedError := <- errCh:
			fmt.Println(receivedError)
	}

}

type message struct {
	To []string
	From string
	Content string
}

type failedMessage struct {
	ErrorMessage string
	OriginalMessage message
}
*/

/* 3.5.3 - add select, send both messages
package main

import (
	"fmt"
)

func main() {

	msgCh := make(chan message, 1)
	errCh := make(chan failedMessage, 1)

	msg := message{
		To: []string{"bilbo@underhill.me"},
		From: "gandalf@whitecouncil.org",
		Content: "Keep it secret, keep it safe",
	}

	failedMessage := failedMessage{
		ErrorMessage: "message intercepted by black rider",
		OriginalMessage: message{},
	}

	msgCh <- msg
	errCh <- failedMessage

	select {
		case receivedMsg := <- msgCh:
			fmt.Println(receivedMsg)
		case receivedError := <- errCh:
			fmt.Println(receivedError)
	}

}

type message struct {
	To []string
	From string
	Content string
}

type failedMessage struct {
	ErrorMessage string
	OriginalMessage message
}
*/

/* 3.5.4 - send no messages
package main

import (
	"fmt"
)

func main() {

	msgCh := make(chan message, 1)
	errCh := make(chan failedMessage, 1)

	select {
		case receivedMsg := <- msgCh:
			fmt.Println(receivedMsg)
		case receivedError := <- errCh:
			fmt.Println(receivedError)
	}

}

type message struct {
	To []string
	From string
	Content string
}

type failedMessage struct {
	ErrorMessage string
	OriginalMessage message
}
*/

/* 3.5.5 - non-blocking select
package main

import (
	"fmt"
)

func main() {

	msgCh := make(chan message, 1)
	errCh := make(chan failedMessage, 1)

	select {
		case receivedMsg := <- msgCh:
			fmt.Println(receivedMsg)
		case receivedError := <- errCh:
			fmt.Println(receivedError)
		default:
			fmt.Println("No messages received")
	}

}

type message struct {
	To []string
	From string
	Content string
}

type failedMessage struct {
	ErrorMessage string
	OriginalMessage message
}
*/
