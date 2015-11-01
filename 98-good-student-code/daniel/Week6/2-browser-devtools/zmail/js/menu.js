var menuButton = document.querySelector('.nav-menu-button'),
	nav        = document.querySelector('#nav');

menuButton.addEventListener('click', function (e) {
	debugger;
    nav.classList.toggle('active');
	e.preventDefault();
});