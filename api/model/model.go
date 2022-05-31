package model

import (
	"log"
	"time"

	"gorm.io/gorm"
)

type Basemodel struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type User struct {
	Basemodel  `mapstructure:",squash"`
	Username   string         `json:"username" gorm:"unique"`
	Firstname  string         `json:"firstname"`
	Surname    string         `json:"surname"`
	Email      string         `json:"email" gorm:"unique"`
	Password   string         `json:"password"`
	Gender     string         `json:"gender"`
	Street     string         `json:"street"`
	Zip        string         `json:"zip"`
	City       string         `json:"city"`
	State      string         `json:"state"`
	Birthday   time.Time      `json:"birthday"`
	Role       string         `json:"role" gorm:"default:user"`
	Token      string         `json:"token"`
	Newsletter bool           `json:"newsletter" gorm:"default:false"`
	Active     bool           `json:"active" gorm:"default:true"`
	Packs      []UserPacks    `json:"userpacks"`
	Answers    []UserAnswers  `json:"useranswers"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	LastLogin  time.Time      `json:"lastlogin"`
}

type Pack struct {
	Basemodel  `mapstructure:",squash"`
	Title      string     `json:"title"`
	Slug       string     `json:"slug" gorm:"unique"`
	Categories []Category `json:"categories" gorm:"foreignKey:PackSlug;references:Slug"`
}

type Category struct {
	Basemodel     `mapstructure:",squash"`
	Title         string        `json:"title"`
	Slug          string        `json:"slug" gorm:"unique"`
	Subcategories []Subcategory `gorm:"foreignKey:CatSlug;references:Slug;"`
	PackSlug      string        `json:"pack_slug" binding:"required"`
}

type Subcategory struct {
	Basemodel `mapstructure:",squash"`
	Title     string     `json:"title"`
	Slug      string     `json:"slug" gorm:"unique"`
	Questions []Question `gorm:"foreignKey:SubSlug;references:Slug"`
	CatSlug   string     `json:"category_slug" binding:"required"`
}

type Question struct {
	Basemodel   `mapstructure:",squash"`
	Title       string   `json:"title"`
	Slug        string   `json:"slug" gorm:"unique"`
	Description string   `json:"description"`
	Answers     []Answer `gorm:"foreignKey:QSlug;references:Slug"`
	SubSlug     string   `json:"subcategory_slug" binding:"required"`
}

type Answer struct {
	Basemodel `mapstructure:",squash"`
	Image     string `json:"image"`
	Answer    string `json:"answer"`
	Correct   bool   `json:"correct"`
	Marked    bool   `json:"marked" gorm:"default:false"`
	QSlug     string `json:"question_slug" binding:"required"`
}

type UserPacks struct {
	Basemodel  `mapstructure:",squash"`
	PackID     uint      `json:"pack_id"`
	PackTitle  string    `json:"pack_title"`
	BillNumber string    `json:"billnumber"`
	Info       string    `json:"info"`
	Amount     string    `json:"amount"`
	Expiry     time.Time `json:"expiry"`
	UserID     uint      `json:"user_id"`
}

type UserAnswers struct {
	Basemodel `mapstructure:",squash"`
	AnswerID  uint `json:"answer_id"`
	Marked    bool `json:"marked" gorm:"default:false"`
	UserID    uint `json:"user_id"`
}

// Custom cascade
func (pack *Pack) AfterDelete(tx *gorm.DB) error {
	// Find relevant relations
	var categories, subcategories, questions []string
	tx.Model(&Category{}).Where("pack_slug = ?", pack.Slug).Select("slug").Find(&categories)
	tx.Model(&Subcategory{}).Where("cat_slug IN ?", categories).Select("slug").Find(&subcategories)
	tx.Model(&Question{}).Where("sub_slug IN ?", subcategories).Select("slug").Find(&questions)

	// Delete objects
	if err := tx.Model(&Answer{}).Where("q_slug IN ?", questions).Delete(&Answer{}).Error; err != nil {
		log.Println(err)
		return err
	}
	if err := tx.Model(&Question{}).Where("sub_slug IN ?", subcategories).Delete(&Question{}).Error; err != nil {
		log.Println(err)
		return err
	}
	if err := tx.Model(&Subcategory{}).Where("cat_slug IN ?", categories).Delete(&Subcategory{}).Error; err != nil {
		log.Println(err)
		return err
	}

	return tx.Model(&Category{}).Where("pack_slug = ?", pack.Slug).Delete(&Category{}).Error
}

func (category *Category) AfterDelete(tx *gorm.DB) error {
	// Find relevant relations
	var subcategories, questions []string
	tx.Model(&Subcategory{}).Where("cat_slug = ?", category.Slug).Select("slug").Find(&subcategories)
	tx.Model(&Question{}).Where("sub_slug IN ?", subcategories).Select("slug").Find(&questions)

	// Delete objects
	if err := tx.Model(&Answer{}).Where("q_slug IN ?", questions).Delete(&Answer{}).Error; err != nil {
		log.Println(err)
		return err
	}
	if err := tx.Model(&Question{}).Where("sub_slug IN ?", subcategories).Delete(&Question{}).Error; err != nil {
		log.Println(err)
		return err
	}

	return tx.Model(&Subcategory{}).Where("cat_slug = ?", category.Slug).Delete(&Subcategory{}).Error
}

func (subcategory *Subcategory) AfterDelete(tx *gorm.DB) error {
	// Find relevant relations
	var questions []string
	tx.Model(&Question{}).Where("sub_slug = ?", subcategory.Slug).Select("slug").Find(&questions)

	// Delete objects
	if err := tx.Model(&Answer{}).Where("q_slug IN ?", questions).Delete(&Answer{}).Error; err != nil {
		log.Println(err)
		return err
	}

	return tx.Model(&Question{}).Where("sub_slug = ?", subcategory.Slug).Delete(&Question{}).Error
}

func (question *Question) AfterDelete(tx *gorm.DB) error {
	return tx.Model(&Answer{}).Where("q_slug = ?", question.Slug).Delete(&Answer{}).Error
}
