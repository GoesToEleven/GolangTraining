// make an api request
function apiRequest(method, endpoint, data, callback) {
  var xhr = new XMLHttpRequest();
  xhr.open(method, "/api/" + endpoint);
  xhr.send(JSON.stringify(data));
  xhr.onreadystatechange = function(evt) {
    if (xhr.readyState === 4) {
      if (Math.floor(xhr.status/100) === 2) {
        if (callback) {
          callback(JSON.parse(xhr.responseText), null);
        }
          // any status other than 2XX, else ...
      } else {
        var msg;
        try {
          msg = JSON.parse(xhr.responseText);
        } catch(e) {
          msg = xhr.responseText;
        }
        if (callback) {
          callback(null, msg);
        }
      }
    }
  };
}

// send a message
function sendMessage(text, callback) {
  apiRequest("POST", "messages", {
    "Text": text
  }, callback);
}

// when a message is received
function onMessage(message) {
    console.log("DID I GET THE MESSAGE " + message);
  var el = document.getElementById("messages");
  var p = document.createElement("p");
  p.textContent = message;
  el.appendChild(p);
}

(function() {

// hook up text input
var controls = document.getElementById("controls");
controls.addEventListener("submit", function(evt) {
  evt.preventDefault();
  var textInput = document.getElementById("text-input");
  var text = textInput.value;
  sendMessage(text, function(res, err) {
    if (err) {
      alert(err);
    }
  });
  textInput.value = "";
}, false);


apiRequest("POST", "channels", null, function(res, err) {
  if (err) {
    alert(err);
    return;
  }
  var token = res;
  var channel = new goog.appengine.Channel(token);
  var socket = channel.open({
    onopen: function() {
      console.log("OPEN", arguments);
    },
    onmessage: function(msg) {
      var data = JSON.parse(msg.data);
      onMessage(data.Text);
        console.log(data);
    },
    onerror: function() {
      console.log("ERROR", arguments);
    },
    onclose: function() {
      console.log("CLOSE", arguments);
    }
  });
});
})();
