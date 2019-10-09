const path = require('path')

module.exports = {
  entry: {
    main: './client/index.js'
  },
  output: {
    filename: '[name].js',
    path: path.resolve('./client/dist')
  },
  module: {
    rules: [
      {
        test: /\.(js|jsx)$/,
        exclude: /node_modules/,
        use: {
          loader: "babel-loader"
        }
      }
    ]
  }
}
