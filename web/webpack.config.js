const webpack = require('webpack')
const path = require('path')
const HtmlWebpackPlugin = require('html-webpack-plugin')
const WebpackCleanupPlugin = require('webpack-cleanup-plugin')
const ExtractTextPlugin = require('extract-text-webpack-plugin')

const PORT = process.env.PORT || "8888" // for Dev Server

let config = {
  entry: [
    './src/index.jsx'
  ],
  output: {
    publicPath: './',
    path: path.join(__dirname, 'public'),
    filename: '[hash].js'
  },
  resolve: {
    extensions: ['.js', '.jsx']
  },
  module: {
    loaders: [
      {test: /\.jsx?$/, exclude: /(node_modules|bower_components|public\/)/,  loader: "babel-loader"},
      {
        test: /\.css$/,
        loader: ExtractTextPlugin.extract({fallback: 'style-loader', use : 'css-loader?sourceMap&localIdentName=[local]__[hash:base64:5]'}),
        exclude: ['node_modules']
      },
      {test: /\.eot(\?v=\d+\.\d+\.\d+)?$/, exclude: /(node_modules|bower_components)/, loader: "file-loader"},
      {test: /\.(woff|woff2)$/, exclude: /(node_modules|bower_components)/, loader: "file-loader?prefix=font/&limit=5000"},
      {test: /\.ttf(\?v=\d+\.\d+\.\d+)?$/, exclude: /(node_modules|bower_components)/, loader: "file-loader?limit=10000&mimetype=application/octet-stream"},
      {test: /\.svg(\?v=\d+\.\d+\.\d+)?$/, exclude: /(node_modules|bower_components)/, loader: "file-loader?limit=10000&mimetype=image/svg+xml"},
      {test: /\.gif/, exclude: /(node_modules|bower_components)/, loader: "file-loader?limit=10000&mimetype=image/gif"},
      {test: /\.jpg/, exclude: /(node_modules|bower_components)/, loader: "file-loader?limit=10000&mimetype=image/jpg"},
      {test: /\.png/, exclude: /(node_modules|bower_components)/, loader: "file-loader?limit=10000&mimetype=image/png"}
    ]
  },
  plugins: [
    new WebpackCleanupPlugin(),
    new webpack.optimize.OccurrenceOrderPlugin(),
    new webpack.optimize.UglifyJsPlugin({
      compress: {
        warnings: false,
        screw_ie8: true,
        drop_console: true,
        drop_debugger: true
      }
    }),
    new ExtractTextPlugin({
      filename: 'styles.css',
      allChunks: true
    }),
    new HtmlWebpackPlugin({
      template: './src/template.html',
      files: {
        css: ['styles.css'],
        js: ['bundle.js'],
      }
    })
  ]
}

if(process.env.NODE_ENV === "production") {
  // Add Production plugins
  config.plugins.push(
    new webpack.DefinePlugin({
      'process.env': {
        NODE_ENV: '"production"'
      }
    })
  )
} else {
  // Entry with hot loader
  config.entry = ['react-hot-loader/patch', './src/index.jsx']
  config.devtool = 'eval-source-map'
  config.devServer = {
    contentBase: "./public",
      noInfo: true, // do not print bundle build stats
      hot: true,    // Hot reloading
      inline: true, // embed the webpack-dev-server runtime into the bundle
      historyApiFallback: true, // serve index.html in place of 404 responses to allow HTML5 history
      port: PORT,
      host: '127.0.0.1'
  }

  config.plugins.push(
    new webpack.NoEmitOnErrorsPlugin(),
    new webpack.NamedModulesPlugin(),
    new webpack.HotModuleReplacementPlugin()
  )
}

module.exports = config
