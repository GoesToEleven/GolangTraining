function changeBackground(e) {
  var html = document.querySelector('html');
  html.style.backgroundImage = "url('" + e.target.src + "')";
}

var images = document.querySelectorAll('img');
for (var i = 0; i < images.length; i++) {
  images[i].addEventListener('click', changeBackground);
}