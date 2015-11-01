var sounds = ['AmazingLee',
			  'EqueKenox',
			  'NightKitty',
			  'Phoebex',
			  'Shiloah'];

function onClick(e) {
	var player = document.querySelector('audio');
	if (e.target.className === 'playing') {
		player.pause();
		e.target.className = 'paused';
	}
	else if (e.target.className === 'paused') {
		player.play();
		e.target.className = 'playing';
	}
	else {
		var lastSong = document.querySelector('.playing, .paused');
		if (lastSong !== null) {
			lastSong.className = '';
		}
		var song = e.target.innerHTML;
		player.pause();
		player.src = 'audio/' + song + '.mp3';
		player.play();
		e.target.className = 'playing';
	}
}

function onSongEnd() {
	var song = document.querySelector('.playing, .paused');
	if (song !== null) {
		song.className = '';
	}
}

function onLoad() {
	var playerList = document.querySelector('#player-list');
	for(var i = 0; i < sounds.length; i++) {
		var newSong = document.createElement('li');
		newSong.innerHTML = sounds[i];
		newSong.addEventListener('click', onClick);
		playerList.appendChild(newSong);
	}
	var player = document.querySelector('audio');
	player.addEventListener('ended', onSongEnd);
}

window.addEventListener('load', onLoad);