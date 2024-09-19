package persistence

import (
	"gorm.io/gorm"
	"url-shortener--go-gin/domain"
)

type IUrlRepository interface {
	GetAllUrls() ([]domain.Url, error)
	GetUrlById(urlId int64) (domain.Url, error)
	GetUrlByShort(shortUrl string) (domain.Url, error)
	AddUrl(url domain.Url) (domain.Url, error)
	UpdateUrl(urlId int64, url domain.Url) (domain.Url, error)
	DeleteUrl(urlId int64) error
}

type UrlRepository struct {
	DB *gorm.DB
}

func NewUrlRepository(db *gorm.DB) IUrlRepository {
	return &UrlRepository{DB: db}
}

func (urlRepository *UrlRepository) GetAllUrls() ([]domain.Url, error) {
	var urls []domain.Url
	result := urlRepository.DB.Find(&urls)
	if result.Error != nil {
		return nil, result.Error
	}

	return urls, nil
}

func (urlRepository *UrlRepository) GetUrlById(urlId int64) (domain.Url, error) {
	var url domain.Url
	result := urlRepository.DB.First(&url, urlId)
	if result.Error != nil {
		return url, result.Error
	}

	return url, nil
}

func (urlRepository *UrlRepository) GetUrlByShort(shortUrl string) (domain.Url, error) {
	var url domain.Url
	result := urlRepository.DB.Where("short_url = ?", shortUrl).First(&url)
	if result.Error != nil {
		return url, result.Error
	}

	return url, nil
}

func (urlRepository *UrlRepository) AddUrl(url domain.Url) (domain.Url, error) {
	result := urlRepository.DB.Create(&url)
	if result.Error != nil {
		return url, result.Error
	}

	return url, nil
}

func (urlRepository *UrlRepository) UpdateUrl(urlId int64, url domain.Url) (domain.Url, error) {
	var existingUrl domain.Url
	result := urlRepository.DB.First(&existingUrl, urlId)
	if result.Error != nil {
		return existingUrl, result.Error
	}

	urlRepository.DB.Model(&existingUrl).Updates(url)

	return existingUrl, nil
}

func (urlRepository *UrlRepository) DeleteUrl(urlId int64) error {
	result := urlRepository.DB.Delete(&domain.Url{}, urlId)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
