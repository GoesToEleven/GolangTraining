package main

import (
	"fmt"
)

func main() {

	msgCh := make(chan Message, 1)
	errCh := make(chan FailedMessage, 1)

	select {
	case receivedMsg := <-msgCh:
		fmt.Println(receivedMsg)
	case receivedError := <-errCh:
		fmt.Println(receivedError)
	default:
		fmt.Println("No messages received")
	}

}

type Message struct {
	To      []string
	From    string
	Content string
}

type FailedMessage struct {
	ErrorMessage    string
	OriginalMessage Message
}

/* 3.5.1 - demo setup
package main

import (
	"fmt"
)

func main() {

	msgCh := make(chan Message, 1)
	errCh := make(chan FailedMessage, 1)

	msg := Message{
		To: []string{"bilbo@underhill.me"},
		From: "gandalf@whitecouncil.org",
		Content: "Keep it secret, keep it safe",
	}

	failedMessage := FailedMessage{
		ErrorMessage: "Message intercepted by black rider",
		OriginalMessage: Message{},
	}

	msgCh <- msg
	errCh <- failedMessage

	fmt.Println(<-msgCh)
	fmt.Println(<-errCh)

}

type Message struct {
	To []string
	From string
	Content string
}

type FailedMessage struct {
	ErrorMessage string
	OriginalMessage Message
}
*/

/* 3.5.2 - add select, use first
package main

import (
	"fmt"
)

func main() {

	msgCh := make(chan Message, 1)
	errCh := make(chan FailedMessage, 1)

	msg := Message{
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

type Message struct {
	To []string
	From string
	Content string
}

type FailedMessage struct {
	ErrorMessage string
	OriginalMessage Message
}
*/

/* 3.5.3 - use second select case
package main

import (
	"fmt"
)

func main() {

	msgCh := make(chan Message, 1)
	errCh := make(chan FailedMessage, 1)

	msg := Message{
		To: []string{"bilbo@underhill.me"},
		From: "gandalf@whitecouncil.org",
		Content: "Keep it secret, keep it safe",
	}

	failedMessage := FailedMessage{
		ErrorMessage: "Message intercepted by black rider",
		OriginalMessage: Message{},
	}

	errCh <- failedMessage

	select {
		case receivedMsg := <- msgCh:
			fmt.Println(receivedMsg)
		case receivedError := <- errCh:
			fmt.Println(receivedError)
	}

}

type Message struct {
	To []string
	From string
	Content string
}

type FailedMessage struct {
	ErrorMessage string
	OriginalMessage Message
}
*/

/* 3.5.3 - add select, send both messages
package main

import (
	"fmt"
)

func main() {

	msgCh := make(chan Message, 1)
	errCh := make(chan FailedMessage, 1)

	msg := Message{
		To: []string{"bilbo@underhill.me"},
		From: "gandalf@whitecouncil.org",
		Content: "Keep it secret, keep it safe",
	}

	failedMessage := FailedMessage{
		ErrorMessage: "Message intercepted by black rider",
		OriginalMessage: Message{},
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

type Message struct {
	To []string
	From string
	Content string
}

type FailedMessage struct {
	ErrorMessage string
	OriginalMessage Message
}
*/

/* 3.5.4 - send no messages
package main

import (
	"fmt"
)

func main() {

	msgCh := make(chan Message, 1)
	errCh := make(chan FailedMessage, 1)

	select {
		case receivedMsg := <- msgCh:
			fmt.Println(receivedMsg)
		case receivedError := <- errCh:
			fmt.Println(receivedError)
	}

}

type Message struct {
	To []string
	From string
	Content string
}

type FailedMessage struct {
	ErrorMessage string
	OriginalMessage Message
}
*/

/* 3.5.5 - non-blocking select
package main

import (
	"fmt"
)

func main() {

	msgCh := make(chan Message, 1)
	errCh := make(chan FailedMessage, 1)

	select {
		case receivedMsg := <- msgCh:
			fmt.Println(receivedMsg)
		case receivedError := <- errCh:
			fmt.Println(receivedError)
		default:
			fmt.Println("No messages received")
	}

}

type Message struct {
	To []string
	From string
	Content string
}

type FailedMessage struct {
	ErrorMessage string
	OriginalMessage Message
}
*/
