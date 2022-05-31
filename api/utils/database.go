package utils

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/ncroxas/questionaire_api/api/model"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

//go:generate mock database -f -d questionaire -u admin -p 5432 -a postgres.com -w pass -r 10

var DB *gorm.DB

func init() {
	var db *gorm.DB
	var err error

	if _, ok := os.LookupEnv("PGHOSTADDR"); ok {
		dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=Europe/Vienna", os.Getenv("PGHOSTADDR"), os.Getenv("PGUSER"), os.Getenv("PGPASSWORD"), os.Getenv("PGDATABASE"), os.Getenv("PGPORT"))

		if db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{}); err != nil {
			panic("failed to connect database")
		}
	} else {
		if db, err = gorm.Open(sqlite.Open("sqlite.db"), &gorm.Config{SkipDefaultTransaction: true}); err != nil {
			panic("failed to connect database")
		}
	}

	// Migrating models
	db.AutoMigrate(&model.Pack{}, &model.Category{}, &model.Subcategory{}, &model.Question{}, &model.Answer{})
	db.AutoMigrate(&model.User{}, &model.UserPacks{}, &model.UserAnswers{})

	// Save Database
	DB = db

	defaultUser()
}

// Create default user if missing
func defaultUser() {
	admin := model.User{
		Username:   "admin",
		Firstname:  "Super",
		Surname:    "Admin",
		Email:      "info@questionaire.com",
		Gender:     "Herr",
		Street:     "TestStreet",
		Zip:        "1111",
		City:       "Vienna",
		State:      "Austria",
		Birthday:   time.Date(2000, 1, 1, 1, 0, 0, 0, time.UTC),
		Role:       "admin",
		Newsletter: false,
		Active:     true,
	}

	var user model.User
	result := DB.Where("username = ?", admin.Username).First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {

		hash, err := bcrypt.GenerateFromPassword([]byte(os.Getenv("ADMIN_PASSWORD")), bcrypt.DefaultCost)
		if err != nil {
			log.Println(err)
			return
		}
		admin.Password = string(hash)
		DB.Create(&admin)
	}
}
