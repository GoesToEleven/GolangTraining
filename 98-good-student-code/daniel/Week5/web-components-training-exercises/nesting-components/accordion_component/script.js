var accordions = Array.prototype.slice.call(document.querySelectorAll('.accordion-component'));
accordions.forEach(function (accordion) {
    var sections = Array.prototype.slice.call(accordion.querySelectorAll('.section'));
    sections.forEach(function (section) {
        var content = section.querySelector('.content');
        section.querySelector('h2').addEventListener('click', function () {
            content.classList.toggle('closed');
        });
    });
});