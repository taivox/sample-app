package api

import (
	"io"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestPong(t *testing.T) {
	app := fiber.New()
	app.Get("/ping", Pong)

	req := httptest.NewRequest("GET", "/ping", nil)
	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	body, _ := io.ReadAll(resp.Body)
	assert.Equal(t, "pong", string(body))
}

func TestLogout(t *testing.T) {
	app := fiber.New()
	app.Get("/logout", Logout)

	req := httptest.NewRequest("GET", "/logout", nil)
	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestAuthorizeSession_NoToken(t *testing.T) {
	app := fiber.New()
	app.Get("/protected", AuthorizeSession, func(c *fiber.Ctx) error {
		return c.SendString("success")
	})

	req := httptest.NewRequest("GET", "/protected", nil)
	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, 401, resp.StatusCode)
}

func TestAuthorizeSession_InvalidToken(t *testing.T) {
	app := fiber.New()
	app.Get("/protected", AuthorizeSession, func(c *fiber.Ctx) error {
		return c.SendString("success")
	})

	req := httptest.NewRequest("GET", "/protected", nil)
	req.Header.Set("Authorization", "invalid-token")
	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, 401, resp.StatusCode)
}

func TestCreateUser_InvalidJSON(t *testing.T) {
	app := fiber.New()
	app.Post("/register", func(c *fiber.Ctx) error {
		return CreateUser(c, nil)
	})

	req := httptest.NewRequest("POST", "/register", strings.NewReader("invalid json"))
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)

	assert.NoError(t, err)
	// Fiber returns 500 for JSON parse errors
	assert.Equal(t, 500, resp.StatusCode)
}

func TestCreateUser_ValidationError(t *testing.T) {
	app := fiber.New()
	app.Post("/register", func(c *fiber.Ctx) error {
		return CreateUser(c, nil)
	})

	// Invalid email and short password
	body := `{"name":"","email":"invalid","password":"abc"}`
	req := httptest.NewRequest("POST", "/register", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, 422, resp.StatusCode)

	respBody, _ := io.ReadAll(resp.Body)
	assert.Contains(t, string(respBody), "Invalid email")
}

func TestLogin_InvalidJSON(t *testing.T) {
	app := fiber.New()
	app.Post("/login", func(c *fiber.Ctx) error {
		return Login(c, nil)
	})

	req := httptest.NewRequest("POST", "/login", strings.NewReader("invalid json"))
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)

	assert.NoError(t, err)
	// Fiber returns 500 for JSON parse errors
	assert.Equal(t, 500, resp.StatusCode)
}

func TestClaims_Structure(t *testing.T) {
	claims := &Claims{}
	assert.NotNil(t, claims)
	assert.Empty(t, claims.User.Email)
	assert.Empty(t, claims.User.Name)
}
