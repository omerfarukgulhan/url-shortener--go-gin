package service

import (
	"url-shortener--go-gin/common/util/id"
	"url-shortener--go-gin/domain"
	"url-shortener--go-gin/persistence"
)

const base62Chars = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

type IUrlService interface {
	GetLongUrl(shortUrl string) (string, error)
	CreateShortUrl(longUrl string) (domain.Url, error)
}

type UrlService struct {
	urlRepository persistence.IUrlRepository
}

func NewUrlService(urlRepository persistence.IUrlRepository) IUrlService {
	return &UrlService{urlRepository: urlRepository}
}

func (urlService *UrlService) GetLongUrl(shortUrl string) (string, error) {
	url, err := urlService.urlRepository.GetUrlByShort(shortUrl)
	if err != nil {
		return "", err
	}
	return url.LongUrl, nil
}

func (urlService *UrlService) CreateShortUrl(longUrl string) (domain.Url, error) {
	uniqueId, err := id.GetUniqueId()
	if err != nil {
		return domain.Url{}, err
	}

	url := domain.Url{
		ID:       uniqueId,
		ShortUrl: generateShortUrl(uniqueId),
		LongUrl:  longUrl,
	}

	return urlService.urlRepository.AddUrl(url)
}

func generateShortUrl(urlId int64) string {
	isRemaining := true
	shortUrl := ""
	for isRemaining {
		shortUrl = string(base62Chars[urlId%62]) + shortUrl
		urlId = urlId / 62
		if urlId == 0 {
			isRemaining = false
		}
	}

	return shortUrl
}
