function deleteButton(keyVal) {
  var xhr = new XMLHttpRequest();
  xhr.open('POST', '/todo.json');
  xhr.setRequestHeader('Content-Type', 'application/x-www-form-urlencoded');
  xhr.addEventListener('readystatechange', function() {
    if (xhr.readyState === 4) {
      getData();
    }
  });
  xhr.send('keyVal=' + keyVal);
}

function getData() {
  var xhr = new XMLHttpRequest();
  xhr.open('GET', '/todo.json');
  xhr.addEventListener('readystatechange', function() {
    if (xhr.readyState === 4) {
      var data = JSON.parse(xhr.responseText);
      newHTML = '';
      for (var i = 0; i < data.length; i++) {
        newHTML += '<p>' + data[i].Value + '</p>' + '<button onclick="deleteButton(\'' + data[i].KeyVal + '\');">Delete</button>';
      }
      document.querySelector('#todo-list').innerHTML = newHTML;
    }
  });
  xhr.send();
}

function sendData(dataObj, callback) {
  var xhr = new XMLHttpRequest();
  xhr.open('POST', '/todo.json');
  var data = JSON.stringify(dataObj);
  xhr.setRequestHeader('Content-Type', 'application/json');
  if (callback) {
    xhr.addEventListener('readystatechange', function() {
      if (xhr.readyState === 4) {
        callback();
      }
    });
  }
  xhr.send(data);
}

document.forms.newitem.addEventListener('submit', function(e) {
  e.preventDefault();
  var data = {Value: e.target.itemtext.value};
  sendData(data, function() {
    e.target.itemtext.value = '';
    getData();
  });
});

getData();
