package requests

type UrlCreateRequest struct {
	LongUrl string `json:"longUrl" binding:"required,url"`
}
