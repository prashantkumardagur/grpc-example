const grpc = require('@grpc/grpc-js');
const protoLoader = require('@grpc/proto-loader');


// ==============================================================


var packageDefinition = protoLoader.loadSync("news.proto");
const NewsService = grpc.loadPackageDefinition(packageDefinition).news.NewsService;


// ==============================================================


const client = new NewsService('localhost:8080', grpc.credentials.createInsecure());


// ==============================================================


// client.getAllNews({}, (err, response) => {
//   console.log(response);
// });

client.getNewsById({ id: 1 }, (err, response) => {
  if(err){ console.log(err); return ;}
  console.log(response);
});