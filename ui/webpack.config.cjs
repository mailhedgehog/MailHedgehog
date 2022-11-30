// webstorm temporal help
const path = require('path');

module.exports = {
  resolve: {
    extensions: ['.ts', '.tsx', '.js', '.json', '.vue'],
    alias: {
      '@': path.resolve('src'),
    },
  },
};
