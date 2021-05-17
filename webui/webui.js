var express = require("express");
var app = express();
var redis = require("redis");

var client = redis.createClient(6379, "redis");
client.on("error", function (err) {
  console.error("Redis error", err);
});

app.get("/", function (req, res) {
  res.redirect("/index.html");
});

app.get("/json", function (req, res) {
  client.hlen("points", function (err, num_coordinates) {
    // points
    client.get("encrypt_coordinates", function (err, coordinates) {
      // encrypt_coordinates
      var now = Date.now() / 1000;
      res.json({
        num_coordinates: num_coordinates,
        coordinates: coordinates,
        now: now,
      });
    });
  });
});

app.use(express.static("files"));

var server = app.listen(80, function () {
  console.log("WEBUI running on port 80");
});
