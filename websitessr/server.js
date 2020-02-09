const express = require('express');
const bodyParser = require('body-parser')
const next = require('next');
const path = require('path')
const dev = process.env.NODE_ENV !== 'production';
const app = next({ dev });
const handle = app.getRequestHandler();
const jwt = require('jsonwebtoken');
const jwtSecret = "mysuperdupersecret";

app
  .prepare()
  .then(() => {
    const server = express();
    server.use(bodyParser.json())
    server.get("/robots.txt", (req, res) => {
      res.header("Content-Type", "text/plain")
      res.sendFile(path.join(__dirname, "./public/static", "robots.txt"))
    })
    server.post("/api/login", (req, res) => {
      // generate a constant token, no need to be fancy here
      const token = jwt.sign({ email: req.body.email }, jwtSecret, { expiresIn: 60 }) // 1 min token
      // return it back
      res.json({ "token": token })
    });
    server.get('*', (req, res) => {
      return handle(req, res);
    });
    server.listen(3000, err => {
      if (err) throw err;
      console.log('> Ready on http://localhost:3000');
    });
  })
  .catch(ex => {
    console.error(ex.stack);
    process.exit(1);
  });