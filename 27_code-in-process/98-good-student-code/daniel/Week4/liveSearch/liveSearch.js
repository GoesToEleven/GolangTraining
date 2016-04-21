var searchBox = document.forms.search.searchBox;

function onKeyUp() {
    var regex = new RegExp(searchBox.value, 'i'),
        names = document.querySelectorAll('.name');
    for (var i = 0; i < names.length; i++) {
        var obj = names[i];
        if (obj.innerHTML.search(regex) === -1) {
            obj.parentNode.style.display = 'none';
        }
        else {
            obj.parentNode.style.display = 'block';
        }
    }
}

function positionPeople() {
    var people = document.querySelectorAll('.person'),
        currentPosition = 0;
    for (var i = 0; i < people.length; i++) {
        var person = people[i];
        if (person.style.display !== 'none') {
            console.log(person);
            person.style.top = currentPosition + 'px';
            currentPosition += person.clientHeight;
        }
    }
}

function getNames() {
    var request = new XMLHttpRequest();
    request.open('GET', 'data.json');
    request.addEventListener('readystatechange', function () {
        if (request.status === 200 && request.readyState === 4) {
            var data = JSON.parse(request.responseText),
                template = document.querySelector('#template').innerHTML;
            document.querySelector('#data-display').innerHTML = Mustache.to_html(template, data);
        }
    });
    request.send();
}

function onLoad() {
    getNames();
    searchBox.addEventListener('keyup', onKeyUp);
    setInterval(positionPeople, 100);
}

window.addEventListener('load', onLoad);