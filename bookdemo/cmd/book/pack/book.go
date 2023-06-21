package pack

import (
	"github.com/ApacheThriftHelicopter/cloudwego-api-gateway/bookdemo/cmd/book/dal/db"
	"github.com/ApacheThriftHelicopter/cloudwego-api-gateway/bookdemo/kitex_gen/book"
)

// Book pack book info
func Book(m *db.Book) *book.Book {
	if m == nil {
		return nil
	}

	return &book.Book{
		BookId:     int64(m.ID),
		UserId:     m.UserID,
		Title:      m.Title,
		Content:    m.Content,
		CreateTime: m.CreatedAt.Unix(),
	}
}

// Books pack list of book info
func Books(ms []*db.Book) []*book.Book {
	books := make([]*book.Book, 0)
	for _, m := range ms {
		if n := Book(m); n != nil {
			books = append(books, n)
		}
	}
	return books
}

func UserIds(ms []*db.Book) []int64 {
	uIds := make([]int64, 0)
	if len(ms) == 0 {
		return uIds
	}
	uIdMap := make(map[int64]struct{})
	for _, m := range ms {
		if m != nil {
			uIdMap[m.UserID] = struct{}{}
		}
	}
	for uId := range uIdMap {
		uIds = append(uIds, uId)
	}
	return uIds
}