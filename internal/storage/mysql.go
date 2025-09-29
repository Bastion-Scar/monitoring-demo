package storage

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
) //

type Logs struct {
	ID       uint `gorm:"primary_key;auto_increment"`
	Duration time.Duration
	Status   int
	IP       string
	Code     string
	URL      string
	Path     string
}

type URLMapping struct {
	ID   uint `gorm:"primary_key;auto_increment"`
	Code string
	URL  string
}

func InitDB() *gorm.DB {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Ошибка загрузки .env файла")
	}

	dsn := os.Getenv("DSN")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Не удалось подключиться к БД:", err)
	}
	err = db.AutoMigrate(&Logs{}, &URLMapping{})
	if err != nil {
		log.Fatal("Не удалось мигрировать БД:", err)
	}
	return db
}

type MySQLStorage struct {
	DB *gorm.DB
}

func NewMySQLStorage() *MySQLStorage {
	return &MySQLStorage{DB: InitDB()}
}

func (m MySQLStorage) Save(code string, url string) {
	entry := URLMapping{Code: code, URL: url}
	err := m.DB.Create(&entry).Error
	if err != nil {
		log.Println("Ошибка при сохранении URL: ", err)
	}
}

func (m MySQLStorage) Load(code string) (string, bool) {
	var entry URLMapping
	err := m.DB.Where("code = ?", code).First(&entry).Error
	if err != nil {
		return "", false
	}
	return entry.URL, true
}
