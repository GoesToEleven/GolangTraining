var gulp = require('gulp'),
	webserver = require('gulp-webserver');

gulp.task('webserver', function() {
	gulp.src('.')
		.pipe(webserver({
			port: 8080,
			livereload: true,
			open: true,
			fallback: 'mustache.html'
		}));
});

gulp.task('default', ['webserver']);