package handler

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/nvlhnn/url-shortener/internal/db/mysql"
	"gorm.io/gorm"
)


type RedirectHandler struct {
	DB mysql.URLMysql
}

func (h *RedirectHandler) RedirectUrl(c *fiber.Ctx) error {

	shortUrl := c.Params("shortUrl")

	// check url exist in database
	url, err := h.DB.Load(shortUrl, false)
	if err == gorm.ErrRecordNotFound {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "URL not found"})
	}else if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Internal Server Error"})
	}

	// check if url is expired
	if  url.ExpiredAt.Before(time.Now()) {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "URL expired"})
	}

	// redirect to original url 
    return c.Redirect(url.OriginalURL, fiber.StatusMovedPermanently)
}
