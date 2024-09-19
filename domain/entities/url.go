package entities

type Url struct {
	ID       uint   `gorm:"primaryKey"`  // Use uint and primary key
	ShortUrl string `gorm:"uniqueIndex"` // Unique index
	LongUrl  string `gorm:"uniqueIndex"` // Unique index
}
