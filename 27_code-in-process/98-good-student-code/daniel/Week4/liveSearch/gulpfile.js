var gulp = require('gulp'),
    webserver = require('gulp-webserver');

gulp.task('webserver', function() {
    gulp.src('.')
        .pipe(webserver({
            port: 8080,
            livereload: true,
            open: true,
            fallback: 'liveSearch.html'
        }));
});

gulp.task('default', ['webserver']);