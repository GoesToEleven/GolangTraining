var timeLeft = null;
var timerId = null;

function countdown() {
    timeLeft--;
    setTimer();
    randomizeHitButton();
}

function randomizeHitButton() {
    var button = document.querySelector('#hit');
    button.style.left = Math.floor(Math.random() * (window.innerWidth - button.clientWidth)) + 'px';
    button.style.top = Math.floor(Math.random() * (window.innerHeight - button.clientHeight)) + 'px';
}

function lerp(start, end, interval) {
    return (end - start) * interval + start;
}

function setTimer() {
    var extraS;
    if (timeLeft === 1) {
        extraS = '';
    }
    else {
        extraS = 's';
    }
    document.querySelector('#timer').innerHTML = timeLeft + ' second' + extraS;
    var interval = (108 - timeLeft) / 108;
    var red = Math.floor(lerp(51, 255, interval));
    var green = Math.floor(lerp(51, 0, interval));
    var blue = Math.floor(lerp(51, 0, interval));
    document.querySelector('html').style.backgroundColor = 'rgb(' + red + ',' + green + ',' + blue + ')';


    if (timeLeft <= 5) {
        document.querySelector('#message').innerHTML = 'You are all going to die!';
    }
    else if (timeLeft <= 10) {
        document.querySelector('#message').innerHTML = 'Bad things will happen when time runs out!';
    }
    else if (timeLeft <= 15) {
        document.querySelector('#message').innerHTML = 'Time is running low!';
    }
    else if (timeLeft <= 20) {
        document.querySelector('#message').innerHTML = 'This is going to be dangerous!';
    }
    else if (timeLeft <= 25) {
        document.querySelector('#message').innerHTML = 'The button needs to be pressed!';
    }
    else {
        document.querySelector('#message').innerHTML = '';
    }


    if (timeLeft === 0) {
        document.querySelector('#failure').style.visibility = 'visible';
        abort();
    }
}

function hitButton() {
    timeLeft = 108;
    if (timerId === null) {
        timerId = setInterval(countdown, 1000);
    }
    setTimer();
}

function abort() {
    if (timerId !== null) {
        clearInterval(timerId);
        timerId = null;
    }
    timeLeft = null;
    document.querySelector('html').style.backgroundColor = 'rgb(51,51,51)';
    document.querySelector('#timer').innerHTML = '';
    document.querySelector('#message').innerHTML = '';
}

document.querySelector('#hit').addEventListener('click', hitButton);
document.querySelector('#abort').addEventListener('click', abort);
