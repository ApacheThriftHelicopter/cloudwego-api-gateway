package db

import (
	"context"

	"github.com/ApacheThriftHelicopter/cloudwego-api-gateway/bookdemo/pkg/constants"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	UserID  int64  `json:"user_id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func (n *Book) TableName() string {
	return constants.BookTableName
}

// CreateBook create book info
func CreateBook(ctx context.Context, books []*Book) error {
	if err := DB.WithContext(ctx).Create(books).Error; err != nil {
		return err
	}
	return nil
}

// MGetBooks multiple get list of book info
func MGetBooks(ctx context.Context, bookIDs []int64) ([]*Book, error) {
	var res []*Book
	if len(bookIDs) == 0 {
		return res, nil
	}

	if err := DB.WithContext(ctx).Where("id in ?", bookIDs).Find(&res).Error; err != nil {
		return res, err
	}
	return res, nil
}

// UpdateBook update book info
func UpdateBook(ctx context.Context, bookID, userID int64, title, content *string) error {
	params := map[string]interface{}{}
	if title != nil {
		params["title"] = *title
	}
	if content != nil {
		params["content"] = *content
	}
	return DB.WithContext(ctx).Model(&Book{}).Where("id = ? and user_id = ?", bookID, userID).
		Updates(params).Error
}

// DeleteBook delete book info
func DeleteBook(ctx context.Context, bookID, userID int64) error {
	return DB.WithContext(ctx).Where("id = ? and user_id = ? ", bookID, userID).Delete(&Book{}).Error
}

// QueryBook query list of book info
func QueryBook(ctx context.Context, userID int64, searchKey *string, limit, offset int) ([]*Book, int64, error) {
	var total int64
	var res []*Book
	conn := DB.WithContext(ctx).Model(&Book{}).Where("user_id = ?", userID)

	if searchKey != nil {
		conn = conn.Where("title like ?", "%"+*searchKey+"%")
	}

	if err := conn.Count(&total).Error; err != nil {
		return res, total, err
	}

	if err := conn.Limit(limit).Offset(offset).Find(&res).Error; err != nil {
		return res, total, err
	}

	return res, total, nil
}