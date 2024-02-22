package mysql

import (
	"time"

	"github.com/nvlhnn/url-shortener/internal/model"
	"gorm.io/gorm"
)

type URLMysql interface {
	Save(originalURL model.URL) (model.URL, error)
	Load(shortenedURL string, filterExpired bool) (model.URL, error)
}

type urlMysql struct {
	db *gorm.DB
}

func NewURLMysql(db *gorm.DB) URLMysql {
	return &urlMysql{
		db : db,
	}
}


func (u *urlMysql) Save(url model.URL) (model.URL, error) {
	err := u.db.Save(&url).Error
	return url, err
}

func (u *urlMysql) Load(shortenedURL string, filterExpired bool) (model.URL, error) {
	var url model.URL

	query := u.db.Where("shortened_url = ?", shortenedURL)
	if filterExpired {
		query = query.Where("expired_at > ?", time.Now())
	}
	err := query.First(&url).Error
	return url, err
}
