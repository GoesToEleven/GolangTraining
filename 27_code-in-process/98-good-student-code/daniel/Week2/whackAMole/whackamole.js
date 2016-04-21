var level = 1,
    interval = 10,
    onHit = 1,
    onMiss = -1,
    onBoss = 40;

var points = 0,
    molePopup = null,
    hits = (level - 1) * interval;

function findMoleSize(size) {
    var width = window.innerWidth,
        height = Math.floor(window.innerHeight * 0.9),
        moleWidth,
        moleHeight;
    if(width < height) {
        moleWidth = width / size;
        moleHeight = moleWidth * 195 / 259;
    }
    else {
        moleHeight = height / size;
        moleWidth = moleHeight * 259 / 195;
    }
    return {width: Math.floor(moleWidth), height: Math.floor(moleHeight)};
}

function sizeMoles() {
    var size = level + 1,
        game = document.querySelector('#game'),
        moles = document.querySelectorAll('.mole'),
        bosses = document.querySelectorAll('.boss'),
        moleSize = findMoleSize(size);
    game.style.width = moleSize.width * size + 'px';
    game.style.height = moleSize.height * size + 'px';
    for (var i = 0; i < moles.length; i++) {
        var mole = moles[i];
        mole.style.width = moleSize.width - 1 + 'px';
        mole.style.height = moleSize.height - 3 + 'px';
    }
    for (var j = 0; j < bosses.length; j++) {
        var boss = bosses[j];
        boss.style.width = moleSize.width * 1.5 + 'px';
        boss.style.height = moleSize.height * 1.5 + 'px';
    }
}

function createMoles(size) {
    var game = document.querySelector('#game');
    for (var i = 0; i < size; i++) {
        var newRow = document.createElement('div');
        newRow.className = 'game-row';
        for (var j = 0; j < size; j++) {
            var newMole = document.createElement('div');
            newMole.className = 'mole';
            newMole.src = 'whackamole.png';
            newMole.dataset.timerId = 'null';
            newMole.addEventListener('click', onClickMole);
            newRow.appendChild(newMole);
        }
        game.appendChild(newRow);
    }
    sizeMoles();
    molePopup = setTimeout(randomMole, 2000);
}

function clearMoles() {
    clearTimeout(molePopup);
    molePopup = null;
    var moles = document.querySelectorAll('.mole');
    for (var i = 0; i < moles.length; i++) {
        moles[i].parentNode.removeChild(moles[i]);
    }
    var rows = document.querySelectorAll('.game-row');
    for (i = 0; i < rows.length; i++) {
        rows[i].parentNode.removeChild(rows[i]);
    }
}

function moveBossRandom() {
    var boss = document.querySelector('.boss'),
        game = document.querySelector('#game');
    boss.style.left = Math.floor(Math.random() * (game.clientWidth - boss.clientWidth)) + 'px';
    boss.style.top = Math.floor(Math.random() * (game.clientHeight - boss.clientHeight)) + 'px';
}

function getRelativePosition(element, relativeTo) {
    var xPos = 0,
        yPos = 0;
    while (element != relativeTo) {
        xPos += (element.offsetLeft - element.scrollLeft + element.clientLeft);
        yPos += (element.offsetTop - element.scrollTop + element.clientTop);
        element = element.offsetParent;
    }
    return {x: xPos, y: yPos};
}

function moveBossFarthest() {
    var boss = document.querySelector('.boss'),
        game = document.querySelector('#game'),
        relativePos = getRelativePosition(boss, game);
    if (relativePos.x < (game.clientWidth - boss.clientWidth) / 2) {
        boss.style.left = (game.clientWidth - boss.clientWidth) + 'px';
    }
    else {
        boss.style.left = '0';
    }
    if (relativePos.y < (game.clientHeight - boss.clientHeight) / 2) {
        boss.style.top = (game.clientHeight - boss.clientHeight) + 'px';
    }
    else {
        boss.style.top = '0';
    }
}

function hitBoss() {
    var boss = document.querySelector('.boss');
    boss.dataset.health = boss.dataset.health - 1;
    boss.removeEventListener('click', hitBoss);
    boss.classList.add('hit');
    setTimeout(function () {
        boss.classList.remove('hit');
        boss.addEventListener('click', hitBoss);
    }, 250);
    playHit();
    if (boss.dataset.health <= 0) {
        clearInterval(boss.dataset.intervalTimer);
        boss.parentNode.removeChild(boss);
        points += onBoss;
        level++;
        updateScores();
        createMoles(level + 1);
    }
}

function startBoss() {
    clearMoles();
    var boss = document.createElement('div');
    boss.className = 'boss';
    boss.dataset.health = 3;
    boss.dataset.intervalTimer = setInterval(moveBossRandom, 1000);
    boss.addEventListener('mouseenter', moveBossFarthest);
    boss.addEventListener('click', hitBoss);
    document.querySelector('#game').appendChild(boss);
    sizeMoles();
}

function updateScores() {
    document.querySelector('#points').innerHTML = 'Points: ' + points;
    document.querySelector('#level').innerHTML = 'Level: ' + level;
}

function playHit() {
    var newSound = document.createElement('audio');
    newSound.src = 'Socapex%20-%20big%20punch.mp3';
    newSound.play();
    document.body.appendChild(newSound);
    setTimeout(function () {
        document.body.removeChild(newSound);
    }, 800);
}

function onClickMole(e) {
    if (e.target.dataset.timerId !== 'null') {
        hits++;
        points += onHit;
        updateScores();
        clearTimeout(e.target.dataset.timerId);
        e.target.dataset.timerId = null;
        e.target.classList.add('hit');
        playHit();
        setTimeout(function () {
            e.target.classList.remove('visible');
            e.target.classList.remove('hit');
        }, 250);
        if (hits >= level * interval) {
            startBoss();
        }
    }
}

function randomMole() {
    var moles = document.querySelectorAll('.mole');
    var unavailableMoles = document.querySelectorAll('.visible, .hit');
    var isDownMole = unavailableMoles.length !== moles.length;
    var whichMole;
    if (isDownMole) {
        do {
            whichMole = moles[Math.floor(Math.random() * moles.length)];
        } while (whichMole.classList.contains('visible') || whichMole.classList.contains('hit'));
        whichMole.classList.add('visible');
        var delay = Math.floor(Math.random() * 4000 + 1000);
        whichMole.dataset.timerId = setTimeout(function () {
            whichMole.classList.remove('visible');
            points += onMiss;
            updateScores();
            whichMole.dataset.timerId = 'null';
        }, delay);
    }
    molePopup = setTimeout(randomMole, Math.floor(Math.random() * 2000 + 500));
}

function onLoad() {
    createMoles(level + 1);
    updateScores();
    window.addEventListener('resize', sizeMoles);
}

window.addEventListener('load', onLoad);