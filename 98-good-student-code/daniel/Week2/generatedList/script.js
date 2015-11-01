/*jslint plusplus: true, devel: true, indent: 2*/

var myLinks = [
	"http://twitter.com/",
	"http://facebook.com/",
	"http://www.tumblr.com/",
	"https://linkedin.com/",
	"https://github.com/",
	"reddit.com"
];

var re = /^(?:https?:\/\/)?(?:www\.)?([a-z0-9_\-]+)\./;
var output = '<style>\n\tli {\n\t\tdisplay: inline-block;\n\t\tmargin-left: 10px;\n\t}\n</style>\n<ul>';
var i;
for (i = 0; i < myLinks.length; i++) {
	var reArray = re.exec(myLinks[i]);
	var icon = reArray[1];
	output += '\n\t<li><a href="' + myLinks[i] + '" target="_blank"><i class="fa fa-2x fa-' + icon + '"></i></a></li>';
}
output += '\n</ul>';

console.log(output);

document.querySelector(".links").innerHTML = output;