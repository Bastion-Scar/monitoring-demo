package service

import (
	"math/rand"
	"monitoring-demo/internal/storage"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type UserService struct {
	repo storage.Storage
}

func NewUserService(repo storage.Storage) *UserService {
	return &UserService{repo}
}

func (s *UserService) Save(c *gin.Context) {
	var req struct {
		URL string `json:"url"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "binding failed",
		})

		return
	}

	shortenURL, err := getRandomURL()

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "cannot create shorten url",
		})
		return
	}

	s.repo.Save(shortenURL, req.URL)

	c.JSON(200, gin.H{
		"shortURL": shortenURL,
	})

	c.Set("code", shortenURL)
	c.Set("url", req.URL)

}

func (s *UserService) Load(c *gin.Context) {
	code := c.Param("code")
	originalURL, ok := s.repo.Load(code)
	if !ok {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "shorten code not found",
		})
		return
	}

	c.Redirect(302, originalURL)
}

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func init() {
	rand.Seed(time.Now().UnixNano())
}

func getRandomURL() (string, error) {
	code := make([]byte, 6)
	for i := range code {
		code[i] = charset[rand.Intn(len(charset))]
	}
	return string(code), nil
}
