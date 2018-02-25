'use strict';

var ExtractTextPlugin = require('extract-text-webpack-plugin');
var path = require('path');

module.exports = {
  entry: './goblog.js',
  output: {
  	path: path.resolve(__dirname, 'public/')
    filename: 'js/app.js'
  },
  module: {
  	loaders: [
  		{
  			test: /\.js/,
  			loader: 'babel-loader',
  		},
  		{
  			test: /\.css/,
  			loader: ExtractTextPlugin.extract('css-loader'),
  		}
  	]
  },
  plugins: [
  	new ExtractTextPlugin('css/app.css')
  ]
};