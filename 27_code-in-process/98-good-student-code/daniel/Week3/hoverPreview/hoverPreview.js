var container = document.createElement('div');
container.style.display = 'flex';
container.style.flexFlow = 'row wrap';

function onEnter(event) {
  var big = document.querySelector('#bigger');
  big.style.backgroundImage = 'url("' + event.target.src + '")';
  big.style.transform = 'scale(1)';
}

function onMove(event) {
  var big = document.querySelector('#bigger');
  big.style.left = event.clientX - 350 + 'px';
  big.style.top = event.clientY + 10 + 'px';
}

function onLeave(event) {
  var big = document.querySelector('#bigger');
  big.style.transform = '';
}

for (var i = 1; i < 10; i++) {
  var image = document.createElement('img');
  image.src = 'http://lorempixel.com/700/700/abstract/' + i;
  image.alt = 'Sports Image ' + i;
  image.addEventListener('mouseenter', onEnter);
  image.addEventListener('mouseleave', onLeave);
  image.addEventListener('mousemove', onMove);
  container.appendChild(image);
}

var bodyNode = document.querySelector('body');
var big = document.querySelector('#bigger');
bodyNode.insertBefore(container, big);