var channel;

function apiRequest(method, endpoint, data, callback) {
  var xhr = new XMLHttpRequest();
  xhr.open(method, '/api/' + endpoint);
  xhr.addEventListener('readystatechange', function() {
    if (xhr.readyState === 4) {
      if (Math.floor(xhr.status / 100) === 2) {
        if (callback) {
          callback(JSON.parse(xhr.responseText), null);
        }
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
  });
  xhr.send(JSON.stringify(data));
}

function sendMessage(text, callback) {
  apiRequest('POST', 'messages', {
    Text: text
  }, callback);
}

function getChannel() {
  apiRequest('POST', 'channels', {}, function(res, err) {
    if (err) {
      alert('Unable to get channel: ' + err);
      return;
    }
    channel = new goog.appengine.Channel(res);
    sock = channel.open();
    sock.onmessage = function(msg) {
      var data = JSON.parse(msg.data);
      onMessage(data);
    };
    sock.onerror = function() {
      alert('An error occured with the connection.');
    };
  });
}

function onMessage(message) {
  var el = document.querySelector('#messages');
  var p = document.createElement('p');
  p.textContent = message.Text;
  el.appendChild(p);
}

(function() {
  getChannel();
  var controls = document.querySelector('#controls');
  controls.addEventListener('submit', function(e) {
    e.preventDefault();
    var textInput = document.querySelector('#text-input');
    var text = textInput.value;
    sendMessage(text, function(res, err) {
      if (err) {
        alert(err);
      }
    });
    textInput.value = '';
  });
})();
