package models_test

import (
	"monitoring-demo/internal/storage"
	"monitoring-demo/models"
	"testing"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/stretchr/testify/assert"
)

func TestSendLogs(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	assert.NoError(t, err)

	db.AutoMigrate(&storage.Logs{})

	go models.SendLogs(db)

	models.LogChan <- storage.Logs{
		Duration: time.Second,
		Status:   200,
		IP:       "127.0.0.1",
		Code:     "abc123",
		URL:      "https://google.com",
		Path:     "/abc123",
	}

	time.Sleep(4 * time.Second)

	var logs []storage.Logs
	err = db.Find(&logs).Error
	assert.NoError(t, err)
	assert.Len(t, logs, 1)
}
