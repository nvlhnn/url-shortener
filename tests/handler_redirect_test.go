package tests

import (
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/nvlhnn/url-shortener/internal/db/mysql"
	"github.com/nvlhnn/url-shortener/internal/handler"
	"github.com/stretchr/testify/assert"
)



func TestRedirectExist(t *testing.T) {

	CreateTestData()

	app := fiber.New()

	redirectHandler := &handler.RedirectHandler{
		DB: mysql.NewURLMysql(DB),
	}
	
	
	app.Get("/:shortUrl", redirectHandler.RedirectUrl)

	req := httptest.NewRequest("GET", "/test", nil)
	resp, err := app.Test(req)

	if err != nil {
		t.Error("Error while testing resolve route")
	}

	assert.Equal(t, resp.StatusCode, 301, "Status code should be 301")

}

type Error struct {
	Message string `json:"message"`
}