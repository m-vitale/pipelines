const express = require('express');
const mockApiMiddleware = require('./mock-api-middleware');

const app = express();
const port = process.argv[2] || 3001;

mockApiMiddleware(app);

app.listen(port, () => {
  console.log('Server listening at http://localhost:' + port);
});
