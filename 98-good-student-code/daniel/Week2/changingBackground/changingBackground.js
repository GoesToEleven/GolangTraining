var backgrounds = [
	'wallpapers/729965-1440x900-[DesktopNexus.com].jpg',
	'wallpapers/813495-1440x900-[DesktopNexus.com].jpg',
	'wallpapers/821425-1440x900-[DesktopNexus.com].jpg',
	'wallpapers/1036809-1440x900-[DesktopNexus.com].jpg',
	'wallpapers/1062969-1440x900-[DesktopNexus.com].jpg'
];


function changeBackground(delta) {
	'use strict';
	window.requestAnimationFrame(changeBackground);
	var node = document.querySelector('html'),
			dateObject = new Date(),
			now = dateObject.getSeconds(),
			backgroundIndex = Math.floor(now / (60 / backgrounds.length));

	node.style.backgroundImage = 'url(' + backgrounds[backgroundIndex] + ')';
}

changeBackground(0);

// preloading hack
var images = new Array();
function preload() {
	for (i = 0; i < backgrounds.length; i++) {
		images[i] = new Image();
		images[i].src = backgrounds[i];
	}
}
preload();