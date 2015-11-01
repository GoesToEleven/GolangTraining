var images = ['Barot_Bellingham',
			  'Constance_Smith',
			  'Hassum_Harrod',
			  'Hillary_Goldwynn',
			  'Hillary_Goldwynn_01',
			  'Hillary_Goldwynn_02',
			  'Hillary_Goldwynn_03',
			  'Hillary_Goldwynn_04',
			  'Hillary_Goldwynn_05',
			  'Hillary_Goldwynn_06',
			  'Hillary_Goldwynn_07',
			  'Jennifer_Jerome',
			  'Jonathan_Ferrar',
			  'LaVonne_LaRue',
			  'Lorenzo_Garcia_01',
			  'Lorenzo_Garcia_02',
			  'Lorenzo_Garcia_03',
			  'Lorenzo_Garcia_04',
			  'Riley_Rewington',
			  'Riley_Rewington_01',
			  'Riley_Rewington_02',
			  'Riley_Rewington_03',
			  'Riley_Rewington_04',
			  'Riley_Rewington_05',
			  'Riley_Rewington_06',
			  'Xhou_Ta'];

function onImageLoad() {
	var spinner = document.querySelector('i');
	spinner.parentNode.removeChild(spinner);
}

function onClick(e) {
	var imageIndex = e.target.dataset.itemIndex,
		overlay = document.querySelector('.overlay');
		bigImage = document.createElement('img');
		spinner = document.createElement('i');
	bigImage.src = 'images/' + images[imageIndex] + '.jpg';
	bigImage.addEventListener('load', onImageLoad);
	overlay.appendChild(bigImage);
	spinner.className = 'fa fa-spinner fa-pulse fa-4x';
	overlay.appendChild(spinner);
	overlay.style.display = '';
}

function onPageLoad() {
	var newImage = document.createElement('img');
	newImage.dataset.itemIndex = Math.floor(Math.random() * images.length);
	newImage.src = 'images/' + images[newImage.dataset.itemIndex] + '_tn.jpg';
	newImage.addEventListener('click', onClick);
	document.querySelector('.center').appendChild(newImage);
	document.querySelector('.overlay').style.display = 'none';
}

window.addEventListener('load', onPageLoad);