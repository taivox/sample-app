package tests

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"testing"
	"time"

	"backend/packages/api"
	"backend/packages/config"
	"backend/packages/db"

	"github.com/apex/log"
	"github.com/stretchr/testify/suite"
)

type apiTestSuite struct {
	suite.Suite
	port   string
	db     *sql.DB
	client *http.Client
}

func TestAPITest(t *testing.T) {
	err := os.Setenv("ENV", "test")
	if err != nil {
		log.WithField("reason", err.Error()).Fatal("Failed to set ENV environment variable")
	}

	config.InitConfig()
	go api.StartServer()
	time.Sleep(1 * time.Second)

	conn, err := db.ConnectDB()
	if err != nil {
		log.WithField("reason", err.Error()).Fatal("Db connection error occurred")
	}

	serverPort := config.Config[config.SERVER_PORT]

	suite.Run(t, &apiTestSuite{
		port:   serverPort,
		db:     conn,
		client: &http.Client{},
	})
}

func (s *apiTestSuite) TearDownSuite() {
	_, _ = s.db.Query(db.DeleteUser, "api_test@test.com")
	_, _ = s.db.Query(db.DeleteUser, "duplicate@test.com")
	_ = s.db.Close()
	api.StopServer()
}

func (s *apiTestSuite) baseURL() string {
	return fmt.Sprintf("http://localhost%s", s.port)
}

func (s *apiTestSuite) post(endpoint string, body map[string]string) (*http.Response, error) {
	jsonData, _ := json.Marshal(body)
	return http.Post(s.baseURL()+endpoint, "application/json", bytes.NewBuffer(jsonData))
}

func (s *apiTestSuite) readBody(resp *http.Response) string {
	body, _ := io.ReadAll(resp.Body)
	return string(body)
}

// Test registration with invalid email
func (s *apiTestSuite) TestRegister_InvalidEmail() {
	resp, err := s.post("/api/register", map[string]string{
		"name":     "Test User",
		"email":    "invalid-email",
		"password": "password123",
	})
	s.NoError(err)
	defer func() { _ = resp.Body.Close() }()

	s.Equal(http.StatusUnprocessableEntity, resp.StatusCode)
	body := s.readBody(resp)
	s.Contains(body, "Invalid email")
}

// Test registration with short password
func (s *apiTestSuite) TestRegister_ShortPassword() {
	resp, err := s.post("/api/register", map[string]string{
		"name":     "Test User",
		"email":    "test@example.com",
		"password": "abc",
	})
	s.NoError(err)
	defer func() { _ = resp.Body.Close() }()

	s.Equal(http.StatusUnprocessableEntity, resp.StatusCode)
	body := s.readBody(resp)
	s.Contains(body, "Invalid password")
}

// Test registration with empty name
func (s *apiTestSuite) TestRegister_EmptyName() {
	resp, err := s.post("/api/register", map[string]string{
		"name":     "",
		"email":    "test@example.com",
		"password": "password123",
	})
	s.NoError(err)
	defer func() { _ = resp.Body.Close() }()

	s.Equal(http.StatusUnprocessableEntity, resp.StatusCode)
	body := s.readBody(resp)
	s.Contains(body, "Invalid name")
}

// Test duplicate email registration
func (s *apiTestSuite) TestRegister_DuplicateEmail() {
	// First registration
	resp1, err := s.post("/api/register", map[string]string{
		"name":     "First User",
		"email":    "duplicate@test.com",
		"password": "password123",
	})
	s.NoError(err)
	defer func() { _ = resp1.Body.Close() }()
	s.Equal(http.StatusOK, resp1.StatusCode)

	// Second registration with same email
	resp2, err := s.post("/api/register", map[string]string{
		"name":     "Second User",
		"email":    "duplicate@test.com",
		"password": "password456",
	})
	s.NoError(err)
	defer func() { _ = resp2.Body.Close() }()

	s.Equal(http.StatusBadRequest, resp2.StatusCode)
	body := s.readBody(resp2)
	s.Contains(body, "email already exists")
}

// Test login with wrong password
func (s *apiTestSuite) TestLogin_WrongPassword() {
	// Register user first
	resp1, _ := s.post("/api/register", map[string]string{
		"name":     "API Test User",
		"email":    "api_test@test.com",
		"password": "correctpassword",
	})
	defer func() { _ = resp1.Body.Close() }()

	// Try login with wrong password
	resp, err := s.post("/api/login", map[string]string{
		"email":    "api_test@test.com",
		"password": "wrongpassword",
	})
	s.NoError(err)
	defer func() { _ = resp.Body.Close() }()

	s.Equal(http.StatusUnauthorized, resp.StatusCode)
	body := s.readBody(resp)
	s.Contains(body, "Incorrect credentials")
}

// Test login with non-existent email
func (s *apiTestSuite) TestLogin_NonExistentEmail() {
	resp, err := s.post("/api/login", map[string]string{
		"email":    "nonexistent@test.com",
		"password": "password123",
	})
	s.NoError(err)
	defer func() { _ = resp.Body.Close() }()

	s.Equal(http.StatusUnauthorized, resp.StatusCode)
}

// Test successful login returns token
func (s *apiTestSuite) TestLogin_Success() {
	// Ensure user exists
	_, _ = s.db.Query(db.DeleteUser, "api_test@test.com")
	resp1, _ := s.post("/api/register", map[string]string{
		"name":     "API Test User",
		"email":    "api_test@test.com",
		"password": "password123",
	})
	defer func() { _ = resp1.Body.Close() }()

	// Login
	resp, err := s.post("/api/login", map[string]string{
		"email":    "api_test@test.com",
		"password": "password123",
	})
	s.NoError(err)
	defer func() { _ = resp.Body.Close() }()

	s.Equal(http.StatusOK, resp.StatusCode)
	body := s.readBody(resp)
	s.Contains(body, "success\":true")
	s.Contains(body, "token")
}

// Test logout
func (s *apiTestSuite) TestLogout() {
	resp, err := http.Get(s.baseURL() + "/api/logout")
	s.NoError(err)
	defer func() { _ = resp.Body.Close() }()

	s.Equal(http.StatusOK, resp.StatusCode)
}

// Test session without token
func (s *apiTestSuite) TestSession_Unauthorized() {
	resp, err := http.Get(s.baseURL() + "/api/session")
	s.NoError(err)
	defer func() { _ = resp.Body.Close() }()

	s.Equal(http.StatusUnauthorized, resp.StatusCode)
}

// Test session with valid token
func (s *apiTestSuite) TestSession_WithToken() {
	// Ensure user exists and login
	_, _ = s.db.Query(db.DeleteUser, "api_test@test.com")
	resp1, _ := s.post("/api/register", map[string]string{
		"name":     "API Test User",
		"email":    "api_test@test.com",
		"password": "password123",
	})
	defer func() { _ = resp1.Body.Close() }()

	loginResp, _ := s.post("/api/login", map[string]string{
		"email":    "api_test@test.com",
		"password": "password123",
	})
	defer func() { _ = loginResp.Body.Close() }()

	// Extract token from response
	var loginResult map[string]interface{}
	body, _ := io.ReadAll(loginResp.Body)
	_ = json.Unmarshal(body, &loginResult)
	token := loginResult["token"].(string)

	// Call session with token
	req, _ := http.NewRequest("GET", s.baseURL()+"/api/session", nil)
	req.Header.Set("Authorization", token)
	resp, err := s.client.Do(req)
	s.NoError(err)
	defer func() { _ = resp.Body.Close() }()

	s.Equal(http.StatusOK, resp.StatusCode)
	sessionBody := s.readBody(resp)
	s.Contains(sessionBody, "api_test@test.com")
}
