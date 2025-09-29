package service_test

import (
	"bytes"
	"encoding/json"
	"monitoring-demo/internal/service"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Мокккк
type MockStorage struct {
	mock.Mock
}

func (m *MockStorage) Save(code, url string) {
	m.Called(code, url)
}

func (m *MockStorage) Load(code string) (string, bool) {
	args := m.Called(code)
	return args.String(0), args.Bool(1)
}

func TestUserService_Save(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockStorage := new(MockStorage)
	userService := service.NewUserService(mockStorage)

	payload := map[string]string{"url": "https://google.com"}
	body, _ := json.Marshal(payload)

	req := httptest.NewRequest(http.MethodPost, "/shorten", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	mockStorage.On("Save", mock.AnythingOfType("string"), "https://google.com").Return().Once()

	userService.Save(c)

	assert.Equal(t, http.StatusOK, w.Code)
	mockStorage.AssertExpectations(t)
}

func TestUserService_Load(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockStorage := new(MockStorage)
	userService := service.NewUserService(mockStorage)

	req := httptest.NewRequest(http.MethodGet, "/testCode", nil)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = []gin.Param{{Key: "code", Value: "testCode"}}
	c.Request = req

	mockStorage.On("Load", "testCode").Return("https://google.com", true).Once()

	userService.Load(c)

	assert.Equal(t, http.StatusFound, w.Code)
	assert.Equal(t, "https://google.com", w.Header().Get("Location"))
	mockStorage.AssertExpectations(t)
}

func TestUserService_Load_NotFound(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockStorage := new(MockStorage)
	userService := service.NewUserService(mockStorage)

	req := httptest.NewRequest(http.MethodGet, "/invalid", nil)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = []gin.Param{{Key: "code", Value: "invalid"}}
	c.Request = req

	mockStorage.On("Load", "invalid").Return("", false).Once()

	userService.Load(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	mockStorage.AssertExpectations(t)
}
