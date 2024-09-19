package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"url-shortener--go-gin/common/util/result"
	"url-shortener--go-gin/domain/requests"
	"url-shortener--go-gin/service"
)

type UrlController struct {
	urlService service.IUrlService
}

func NewUrlController(urlService service.IUrlService) *UrlController {
	return &UrlController{urlService: urlService}
}

func (urlController *UrlController) RegisterUrlRoutes(router *gin.Engine) {
	router.GET("/urls/:shortUrl", urlController.GetLongUrl)
	router.POST("/urls", urlController.CreateShortUrl)
}

func (urlController *UrlController) GetLongUrl(ctx *gin.Context) {
	shortUrl := ctx.Param("shortUrl")
	if shortUrl == "" {
		ctx.JSON(http.StatusBadRequest, result.NewResult(false, "Enter a valid short URL"))
		return
	}

	longUrl, err := urlController.urlService.GetLongUrl(shortUrl)
	if err != nil {
		ctx.JSON(http.StatusNotFound, result.NewResult(false, "Short URL not found"))
		return
	}

	ctx.Redirect(http.StatusPermanentRedirect, longUrl)
}

func (urlController *UrlController) CreateShortUrl(ctx *gin.Context) {
	var urlRequest requests.UrlCreateRequest
	if err := ctx.ShouldBindJSON(&urlRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, result.NewResult(false, "Enter url in valid format"))
		return
	}

	shortUrl, err := urlController.urlService.CreateShortUrl(urlRequest.LongUrl)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, result.NewResult(false, "Error creating short URL"))
		return
	}

	ctx.JSON(http.StatusCreated, result.NewDataResult(true, "Url created", shortUrl))
}
