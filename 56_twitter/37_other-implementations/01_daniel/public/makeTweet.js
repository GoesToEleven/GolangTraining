document.querySelector('#create').addEventListener('click', function() {
  var form = document.createElement('form');
  form.id = 'createTweet';
  form.method = 'POST';

  charactersLeft = document.createElement('p');
  charactersLeft.textContent = '140 characters left';
  charactersLeft.id = 'characters';
  form.appendChild(charactersLeft);

  textInput = document.createElement('input');
  textInput.type = 'text';
  textInput.name = 'message';
  textInput.setAttribute('maxlength', 140);
  textInput.placeholder = 'Message';
  textInput.setAttribute('required', 'true');
  textInput.addEventListener('input', function() {
    document.querySelector('#characters').textContent = (140 - textInput.value.length) + ' characters left';
  });
  form.appendChild(textInput);

  submitInput = document.createElement('input');
  submitInput.type = 'submit';
  submitInput.value = 'Tweet';
  form.appendChild(submitInput);

  form.addEventListener('submit', function(e) {
    e.preventDefault();
    xhr = new XMLHttpRequest();
    xhr.open('POST', '/tweet.json');
    xhr.send(textInput.value);
    form.parentNode.removeChild(form);
  });
  document.body.appendChild(form);
});
