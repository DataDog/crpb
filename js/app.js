const express = require('express');
const app = express();

app.get("/", async (_, res) => {
  console.log("[test] sample log for / route test");
  res.status(200).json({"msg":"Hello world!"});
});

const port = process.env.PORT || 8080;
app.listen(port, () => console.log(`Application is now ready, listening to: ${port} randomId: ${Math.random()}`));