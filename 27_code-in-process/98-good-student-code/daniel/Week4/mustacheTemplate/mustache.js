var request = new XMLHttpRequest();
request.open('GET', 'data.json');
request.addEventListener('readystatechange', function() {
	if (request.status === 200 && request.readyState === 4) {
		var data = JSON.parse(request.responseText),
			template = document.querySelector('#template').innerHTML,
			html = Mustache.to_html(template, data);
		document.querySelector('#data-display').innerHTML = html;
	}
});
request.send();