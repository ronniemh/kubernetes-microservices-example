const express = require("express");
const app = express();

app.get('/getPoint', function (req, res) {

  res.send({pointX:Math.random(), pointY: Math.random()});
});
app.listen(3000, () => {
 console.log("getPoint run at port 3000");
});