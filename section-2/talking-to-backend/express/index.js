const express = require("express");
const cors = require("cors");

const app = express();

app.use(cors());
app.use(express.json());
app.use(express.urlencoded({extended : true}));

app.get("/", (req,res)=>{
    res.status(200).send("HI go lang")
})

app.get("/get", (req,res)=>{
    res.status(200).json({ message: "JSON data go" })
})

app.post("/post", (req,res)=>{
    res.status(200).send(req.body)
})

app.post("/postform", (req,res)=>{
    res.status(200).send(JSON.stringify(req.body))
})

app.listen(3000, () => {
  console.log(`server started at http://localhost:3000`);
});
