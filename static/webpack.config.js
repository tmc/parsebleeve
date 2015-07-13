module.exports = {
  entry: './src/index.js',
  output: { filename: 'app.js' },
  module: { loaders: [
    { test: /\.jsx?$/, 
      loader: 'babel-loader?stage=0',
      exclude: /node_modules/
    }
  ] },
  externals: { 'react': 'React' }
};
