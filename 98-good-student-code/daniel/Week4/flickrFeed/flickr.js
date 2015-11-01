function formatImage() {
	return function(item, render) {
		var parsed = render(item);
		return parsed.slice(0, -10) + 's' + parsed.slice(-9);
	};
}

function jsonFlickrFeed(data) {
	data.items.forEach(function(item) {
		item.formatImage = formatImage;
	});
	var template = document.querySelector('#template').innerHTML,
		html = Mustache.to_html(template, data.items);
	document.querySelector('#photos').innerHTML = html;
}

function onClick(target) {
	var overlay = document.createElement('div'),
		image = document.createElement('img');
	overlay.className = 'overlay';
	overlay.addEventListener('click', onCloseClick);
	image.src = target.src.slice(0, -5) + 'b' + target.src.slice(-4);
	overlay.appendChild(image);
	document.body.appendChild(overlay);
}

function onCloseClick(target) {
	var overlay = document.querySelector('.overlay');
	overlay.parentNode.removeChild(overlay);
}