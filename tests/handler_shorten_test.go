package tests

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http/httptest"
	"net/url"
	"path"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/nvlhnn/url-shortener/internal/db/mysql"
	"github.com/nvlhnn/url-shortener/internal/handler"
	"github.com/stretchr/testify/assert"
)



func TestShortenWithoutCustomURL(t *testing.T) {

	app := fiber.New()

	shortenHandler := &handler.ShortenHandler{
		DB: mysql.NewURLMysql(DB),
	}
	
	app.Post("/api/v1", shortenHandler.ShortenURL)

	url_string := "https://github.com/nvlhnn";
	body := fmt.Sprintf(`{"url": "%s"}`, url_string)
	req := httptest.NewRequest("POST", "/api/v1", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("Error while testing shorten route: %v", err)
	}

	if resp.StatusCode != 201 {
		t.Fatalf("Expected status code 200, got %v", resp.StatusCode)
	}
	
	var responseBody handler.Response
	err = json.NewDecoder(resp.Body).Decode(&responseBody)
	if err != nil {
		t.Fatalf("Error while decoding response body: %v", err)
	}

	assert.Equal(t, resp.StatusCode, 201, "Status code should be 200")
	if resp.StatusCode != 201 {
		return
	}

	parsedURL, err := url.Parse(responseBody.ShortURL)
    if err != nil {
		t.Fatalf("Error parsing URL: %v", err)
        return
    }

    segment := path.Base(parsedURL.Path)

	assert.Equal(t, responseBody.URL, "https://github.com/nvlhnn", "URL should be https://github.com/nvlhnn")
	assert.Equal(t, resp.Header.Get("Content-Type"), "application/json", "Content-Type should be application/json")
	assert.Len(t, segment, 6, "Short URL should be 6 characters long")
	assert.NotEmpty(t, responseBody.ExpiredAt, "ExpiredAt should not be empty")

}

func TestShortenWithCustomURL(t *testing.T) {

	app := fiber.New()

	shortenHandler := &handler.ShortenHandler{
		DB: mysql.NewURLMysql(DB),
	}
	
	app.Post("/api/v1", shortenHandler.ShortenURL)

	body := `{"url":"https://github.com/nvlhnn", "custom_url": "nvlhnn"}`
	req := httptest.NewRequest("POST", "/api/v1", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("Error while testing shorten route: %v", err)
	}

	if resp.StatusCode != 201 {
		log.Panicln(resp)
		t.Fatalf("Expected status code 201, got %v", resp.StatusCode)
	}
	
	var responseBody handler.Response
	err = json.NewDecoder(resp.Body).Decode(&responseBody)
	if err != nil {
		t.Fatalf("Error while decoding response body: %v", err)
	}

	assert.Equal(t, resp.StatusCode, 201, "Status code should be 200")
	if resp.StatusCode != 201 {
		return
	}

	parsedURL, err := url.Parse(responseBody.ShortURL)
    if err != nil {
		t.Fatalf("Error parsing URL: %v", err)
        return
    }

    segment := path.Base(parsedURL.Path)

	assert.Equal(t, responseBody.URL, "https://github.com/nvlhnn", "URL should be https://github.com/nvlhnn")
	assert.Equal(t, resp.Header.Get("Content-Type"), "application/json", "Content-Type should be application/json")
	assert.True(t, len(segment) <= 16, "custom_url should be less than or equal to 16")
	assert.NotEmpty(t, responseBody.ExpiredAt, "ExpiredAt should not be empty")
	assert.Equal(t, segment, "nvlhnn", "Short URL should be nvlhnn")

}