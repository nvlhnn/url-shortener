package handler

import (
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/nvlhnn/url-shortener/internal/db/mysql"
	"github.com/nvlhnn/url-shortener/internal/model"
)


type request struct {
	URL 		string `json:"url"`
	CustomURL	string `json:"custom_url,omitempty"`
}

type Response struct {
	URL             string        `json:"url"`
	ShortURL  		string        `json:"short_url"`
	ExpiredAt       string	  	`json:"expired_at"`
}

type ShortenHandler struct {
	DB mysql.URLMysql
}



func (h *ShortenHandler) ShortenURL(c *fiber.Ctx) error {

	var req request

	err := c.BodyParser(&req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "cannot parse JSON body"})
	}

	// validate url is valid
	if !govalidator.IsURL(req.URL) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid URL"})
	}

	// validate custom url is valid
	if req.CustomURL != "" {
		if !govalidator.IsAlphanumeric(req.CustomURL) {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid custom URL"})
		}

	}else{
		req.CustomURL = uuid.New().String()[:6]
	}

	// validate custom url max length is 16 characters
	if len(req.CustomURL) > 16 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Custom URL too long"})
	}

	// validate custom url is not in use
	if _, err := h.DB.Load(req.CustomURL, false); err == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Short URL already in use"})
	}
	
	// set expired date to 5 years from now and convert to UTC+7 datetime string
	expiredAt := time.Now().AddDate(5, 0, 0)

	// create a new URL model
	url := model.URL{
		OriginalURL: req.URL,
		ShortenedURL: req.CustomURL,
		ExpiredAt: expiredAt,
	}

	// save the url to the database
	url, err = h.DB.Save(url)
	if err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Internal Server Error"})
	}

	return c.Status(fiber.StatusCreated).JSON(Response{
		URL:            url.OriginalURL,
		ShortURL:       c.BaseURL() + "/" + url.ShortenedURL,
		ExpiredAt:      expiredAt.Format("2006-01-02 15:04:05"),
	})
}
