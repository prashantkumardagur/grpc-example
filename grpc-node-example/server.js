const grpc = require('@grpc/grpc-js');
const protoLoader = require('@grpc/proto-loader');


// ==============================================================


const server = new grpc.Server();

var packageDefinition = protoLoader.loadSync("news.proto");
var News = grpc.loadPackageDefinition(packageDefinition).news;


// ==============================================================


const news = [
  { id: 1, title: 'News 1', body: 'Content 1', titleImage: "/someimagepath/1" },
  { id: 2, title: 'News 2', body: 'Content 2', titleImage: "/someimagepath/2" },
  { id: 3, title: 'News 3', body: 'Content 3', titleImage: "/someimagepath/3" },
  { id: 4, title: 'News 4', body: 'Content 4', titleImage: "/someimagepath/4" },
];


// ==============================================================


server.addService(News.NewsService.service, {

  // function to get all news
  getAllNews: (_, callback) => {
    callback(null, { news });
  },

  // function to get news by id
  getNewsById: (call, callback) => {
    const { id } = call.request;
    const newsItem = news.find((item) => item.id == id);
    callback(null, newsItem);
  },
});



// ==============================================================


server.bindAsync(
  "localhost:50051",
  grpc.ServerCredentials.createInsecure(),
  (err, port) => {
    server.start();
    console.log(`Server running on port ${port}`);
  }
);