package persistence

import (
	"gorm.io/gorm"
	"url-shortener--go-gin/domain/entities"
)

type IUrlRepository interface {
	GetAllUrls() ([]entities.Url, error)
	GetUrlById(urlId int64) (entities.Url, error)
	GetUrlByShort(shortUrl string) (entities.Url, error)
	AddUrl(url entities.Url) (entities.Url, error)
	UpdateUrl(urlId int64, url entities.Url) (entities.Url, error)
	DeleteUrl(urlId int64) error
}

type UrlRepository struct {
	DB *gorm.DB
}

func NewUrlRepository(db *gorm.DB) IUrlRepository {
	return &UrlRepository{DB: db}
}

func (urlRepository *UrlRepository) GetAllUrls() ([]entities.Url, error) {
	var urls []entities.Url
	result := urlRepository.DB.Find(&urls)
	if result.Error != nil {
		return nil, result.Error
	}

	return urls, nil
}

func (urlRepository *UrlRepository) GetUrlById(urlId int64) (entities.Url, error) {
	var url entities.Url
	result := urlRepository.DB.First(&url, urlId)
	if result.Error != nil {
		return url, result.Error
	}

	return url, nil
}

func (urlRepository *UrlRepository) GetUrlByShort(shortUrl string) (entities.Url, error) {
	var url entities.Url
	result := urlRepository.DB.Where("short_url = ?", shortUrl).First(&url)
	if result.Error != nil {
		return url, result.Error
	}

	return url, nil
}

func (urlRepository *UrlRepository) AddUrl(url entities.Url) (entities.Url, error) {
	result := urlRepository.DB.Create(&url)
	if result.Error != nil {
		return url, result.Error
	}

	return url, nil
}

func (urlRepository *UrlRepository) UpdateUrl(urlId int64, url entities.Url) (entities.Url, error) {
	var existingUrl entities.Url
	result := urlRepository.DB.First(&existingUrl, urlId)
	if result.Error != nil {
		return existingUrl, result.Error
	}

	urlRepository.DB.Model(&existingUrl).Updates(url)

	return existingUrl, nil
}

func (urlRepository *UrlRepository) DeleteUrl(urlId int64) error {
	result := urlRepository.DB.Delete(&entities.Url{}, urlId)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
