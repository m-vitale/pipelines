import * as MinifyPlugin from 'babel-minify-webpack-plugin';
import * as webpack from 'webpack';
import * as merge from 'webpack-merge';
import common from './webpack.common';

export default merge(common, {
  plugins: [
    new MinifyPlugin(),
    new webpack.DefinePlugin({
      'process.env.NODE_ENV': JSON.stringify('prod')
    }),
    new webpack.optimize.AggressiveMergingPlugin(),
  ],
});
