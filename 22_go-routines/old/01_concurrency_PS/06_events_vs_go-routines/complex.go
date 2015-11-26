package main

import (
	"fmt"
)

func main() {
	btn := makeButton()

	handlerOne := make(chan string, 2)
	handlerTwo := make(chan string, 2)

	btn.AddEventListener("click", handlerOne)
	btn.AddEventListener("click", handlerTwo)

	go func() {
		for {
			msg := <-handlerOne
			fmt.Println("Handler One: " + msg)
		}
	}()

	go func() {
		for {
			msg := <-handlerTwo
			fmt.Println("Handler Two: " + msg)
		}
	}()

	btn.TriggerEvent("click", "button clicked!")

	btn.RemoveEventListener("click", handlerOne)

	btn.TriggerEvent("click", "button clicked again!")

	fmt.Scanln()
}

type button struct {
	eventListeners map[string][]chan string
}

func (btn *button) AddEventListener(event string, responseChannel chan string) {
	if _, present := btn.eventListeners[event]; present {
		btn.eventListeners[event] =
			append(btn.eventListeners[event], responseChannel)
	} else {
		btn.eventListeners[event] = []chan string{responseChannel}
	}
}

func (btn *button) RemoveEventListener(event string, listenerChannel chan string) {
	if _, present := btn.eventListeners[event]; present {
		for idx := range btn.eventListeners[event] {
			if btn.eventListeners[event][idx] == listenerChannel {
				btn.eventListeners[event] = append(btn.eventListeners[event][:idx],
					btn.eventListeners[event][idx+1:]...)
				break
			}
		}
	}
}

func (btn *button) TriggerEvent(event string, response string) {
	if _, present := btn.eventListeners[event]; present {
		for _, handler := range btn.eventListeners[event] {
			go func(handler chan string) {
				handler <- response
			}(handler)
		}
	}
}

func makeButton() *button {
	result := new(button)
	result.eventListeners = make(map[string][]chan string)
	return result
}
