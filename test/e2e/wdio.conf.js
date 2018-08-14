const path = require('path');

exports.config = {
  maxInstances: 1,
  baseUrl: 'http://localhost:3000',
  capabilities: [{
    maxInstances: 1,
    browserName: 'chrome',
    chromeOptions: {
      args: ['--headless', '--disable-gpu', '--window-size=1024,768'],
    },
  }],
  coloredLogs: true,
  connectionRetryCount: 3,
  connectionRetryTimeout: 90000,
  deprecationWarnings: false,
  framework: 'mocha',
  host: '127.0.01',
  port: 4444,
  mochaOpts: {
    timeout: 100000,
  },
  logLevel: 'silent',
  plugins: {
    'wdio-webcomponents': {},
  },
  reporters: ['dot', 'junit'],
  reporterOptions: {
    junit: {
      outputDir: './',
      outputFileFormat: () => 'junit_E2eTestOutput.xml',
    },
  },
  specs: [
    './e2e.spec.js',
  ],
  sync: true,
  waitforTimeout: 10000,
}
