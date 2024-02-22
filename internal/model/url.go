package model

import "time"

type URL struct {
    ID           uint       `gorm:"primaryKey"`
    OriginalURL  string     `gorm:"not null"`
    ShortenedURL string     `gorm:"not null"`
    ExpiredAt    time.Time  `gorm:"not null"`
}