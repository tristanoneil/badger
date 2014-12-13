var gulp = require('gulp'),
    sass = require('gulp-sass'),
    mainBowerFiles = require('main-bower-files'),
    concat = require('gulp-concat');

//
// Compile all application SCSS files
// to /dist/css
//
gulp.task('sass', function () {
  gulp.src('assets/scss/*.scss')
    .pipe(sass())
    .pipe(gulp.dest('dist/css'));
});

//
// Concat all JS.
//
gulp.task('js', function() {
  var scripts = mainBowerFiles();
  scripts.push('assets/js/**/*.js');

  gulp.src(scripts)
    .pipe(concat('all.js'))
    .pipe(gulp.dest('./dist/js'));
});

//
// Compile SASS and JS to a single file.
//
gulp.task('default', ['sass', 'js']);

//
// Compile SASS and JS to a single file and watch for changes.
//
gulp.task('watch', ['sass', 'js'], function() {
  gulp.watch('assets/scss/*.scss', ['sass']);
  gulp.watch('assets/js/*.js', ['js']);
});
