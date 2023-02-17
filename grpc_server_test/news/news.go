package news

import (
	"context"
)

// Server is used to implement news.ChatServer
type Server struct {
	UnimplementedNewsServiceServer
}

// =================================================================================================

var news = *new(NewsList)

// add some news
func Init() {
	news.News = append(news.News, &News{Id: "1", Title: "News 1", Body: "Content 1", PostImage: "/someUrl1"})
	news.News = append(news.News, &News{Id: "2", Title: "News 2", Body: "Content 2", PostImage: "/someUrl2"})
	news.News = append(news.News, &News{Id: "3", Title: "News 3", Body: "Content 3", PostImage: "/someUrl3"})
	news.News = append(news.News, &News{Id: "4", Title: "News 4", Body: "Content 4", PostImage: "/someUrl4"})
	news.News = append(news.News, &News{Id: "5", Title: "News 5", Body: "Content 5", PostImage: "/someUrl5"})
}

// =================================================================================================

// GetNews implementation
func (s *Server) GetAllNews(ctx context.Context, in *Empty) (*NewsList, error) {
	return &news, nil
}

// GetNewsById implementation
func (s *Server) GetNewsById(ctx context.Context, in *News) (*News, error) {
	id := in.GetId()
	for _, n := range news.News {
		if n.GetId() == id {
			return n, nil
		}
	}
	return nil, nil
}
