const path = require('path')
const ForkTsCheckerWebpackPlugin = require('fork-ts-checker-webpack-plugin')

const VuetifyLoaderPlugin = require('vuetify-loader/lib/plugin')

const clientPath = process.env.CLIENT_PATH || './../client'
const rootPath = path.resolve(__dirname, './../', clientPath)

module.exports = {
  module: {
    rules: [
      {
        test: /\.ts$/,
        exclude: /node_modules/,
        loaders: [
          {
            loader: 'ts-loader',
            options: {
              appendTsSuffixTo: [/\.vue$/],
              transpileOnly: true
            }
          }
        ]
      },
      {
        test: /\.(c|sc|sa)ss$/,
        use: [
          'vue-style-loader',
          'style-loader',
          'css-loader',
          {
            loader: 'sass-loader',
            // Requires sass-loader@^7.0.0
            options: {
              implementation: require('sass'),
              fiber: require('fibers'),
              indentedSyntax: true // optional
            },
            // Requires sass-loader@^8.0.0
            options: {
              implementation: require('sass'),
              sassOptions: {
                fiber: require('fibers'),
                indentedSyntax: true // optional
              },
            },
          },
        ],
      },
    ]
  },

  plugins: [
    new ForkTsCheckerWebpackPlugin(),
    new VuetifyLoaderPlugin()
  ],

  resolve: {
    extensions: ['.ts', '.js', '.vue', '.json'],
    alias: {
      'static': path.resolve(rootPath, './app/static'),
      'assets': path.resolve(rootPath, './app/assets'),
      '~': path.resolve(rootPath, './app'),
      '@': path.resolve(rootPath, './app'),
      '~~': path.resolve(rootPath, './'),
      '@@': path.resolve(rootPath, './')
    }
  }
}
