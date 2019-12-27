const path = require('path')
const ForkTsCheckerWebpackPlugin = require('fork-ts-checker-webpack-plugin')

module.exports = {
  module: {
    rules: [{
      test: /\.ts$/,
      exclude: /node_modules/,
      use: [{
        loader: 'ts-loader',
        options: {
          appendTsSuffixTo: [/\.vue$/],
          transpileOnly: true
        },
      }]
    }]
  },

  plugins: [
    new ForkTsCheckerWebpackPlugin()
  ],

  resolve: {
    extensions: ['.ts', '.js', '.vue', '.json'],
    alias: {
      'static': path.resolve(__dirname, './../../app/static'),
      'assets': path.resolve(__dirname, './../../app/assets'),
      '~': path.resolve(__dirname, './../../client/app'),
      '@': path.resolve(__dirname, './../../client/app'),
      '~~': path.resolve(__dirname, './../../client'),
      '@@': path.resolve(__dirname, './../../client')
    }
  }
}
