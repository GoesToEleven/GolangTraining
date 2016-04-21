package main

import (
	"io"
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/channel"
)

func init() {
	// build routing table
	http.HandleFunc("/create_channel", handleCreateChannel)
	http.HandleFunc("/listen", handleListen)
	http.HandleFunc("/send", handleSend)
}

func handleSend(res http.ResponseWriter, req *http.Request) {
	ctx := appengine.NewContext(req)
	message := req.FormValue("message")
	err := channel.SendJSON(ctx, "example", message)
	if err != nil {
		http.Error(res, err.Error(), 500)
	}
	io.WriteString(res, "Sent!")
}

func handleCreateChannel(res http.ResponseWriter, req *http.Request) {
	ctx := appengine.NewContext(req)
	tok, err := channel.Create(ctx, "example")
	if err != nil {
		http.Error(res, err.Error(), 500)
		return
	}
	io.WriteString(res, tok)
}

func handleListen(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, `<!DOCTYPE html>
<html>
  <head>
  <script type="text/javascript" src="/_ah/channel/jsapi"></script>
  </head>
  <body>
<script>
// I want to listen on the channel
var token = prompt("Enter Your Token");
// listen
var channel = new goog.appengine.Channel(token);

// this:
 var socket = channel.open();
 socket.onopen = function(evt) {
	console.log("OPEN", arguments);
 };
 socket.onmessage = function(msg) {
	var data = JSON.parse(msg.data);
    alert("Message Received: " + data);
 };
 socket.onerror = function(evt) {
	console.log("ERROR", arguments);
 };
 socket.onclose = function(evt) {
	console.log("CLOSE", arguments);
 };

// or this:
var socket = channel.open({
  onopen: function() {
    console.log("OPEN", arguments);
  },
  onmessage: function(msg) {
    var data = JSON.parse(msg.data);
    alert("Message Received: " + data);
  },
  onerror: function() {
    console.log("ERROR", arguments);
  },
  onclose: function() {
    console.log("CLOSE", arguments);
  }
});
</script>
  </body>
</html>`)
}
