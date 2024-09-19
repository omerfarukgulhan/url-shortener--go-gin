package domain

type Url struct {
	ID       int64  `gorm:"primaryKey"`
	ShortUrl string `gorm:"uniqueIndex"`
	LongUrl  string `gorm:"uniqueIndex"`
}
