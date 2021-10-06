const gulp = require('gulp');
const babel = require('gulp-babel');
const webpack = require('webpack-stream');
const jsdoc = require("gulp-jsdoc3");

gulp.task('default', () =>
  gulp.src('target/generated-sources/swagger/src/**/*.js')
    .pipe(webpack( require('./target/generated-sources/swagger/webpack.config.js') ))
    .pipe(gulp.dest('target/dist/'))
    .pipe(babel({
      presets: ['@babel/env']
    }))
    .pipe(gulp.dest('target/transpile-output/'))
);
gulp.task('js-doc', function (cb) {
  gulp.src('target/dist/*.js')
    .pipe(jsdoc( require('./jsdoc-conf.json'), cb))
});
gulp.task('build-prod', () =>
  gulp.src('target/generated-sources/swagger/src/**/*.js')
    .pipe(webpack( require('./target/generated-sources/swagger/webpack.config.prod.js') ))
    .pipe(gulp.dest('target/dist/'))
    .pipe(babel({
      presets: ['@babel/env']
    }))
    .pipe(gulp.dest('target/transpile-output/'))
);