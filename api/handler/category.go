package handler

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gosimple/slug"
	"github.com/ncroxas/questionaire_api/api/model"
	"github.com/ncroxas/questionaire_api/api/utils"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func cleanC(category *model.Category) *model.Category {
	new_slug := fmt.Sprintf("%v-%v", category.Title, category.PackSlug)
	cleanup := &model.Category{
		Title:    strings.TrimSpace(category.Title),
		Slug:     slug.MakeLang(new_slug, "de"),
		PackSlug: category.PackSlug,
	}
	return cleanup
}

func GetCategory(c *gin.Context) {
	db := utils.DB

	var category model.Category
	if err := c.ShouldBind(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := db.Where("slug = ?", category.Slug).First(&category)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Category " + category.Slug + " not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"category": category})
}

func NewCategory(c *gin.Context) {
	db := utils.DB

	var category model.Category
	if err := c.ShouldBind(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if category.Title == "" {
		c.JSON(http.StatusNotFound, gin.H{"error": "Title cannot be empty!"})
		return
	}

	result := db.Where("slug = ?", category.PackSlug).First(&model.Pack{})
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Pack " + category.PackSlug + " not found!"})
		return
	}
	
	newCategory := cleanC(&category)
	db.Create(&newCategory)

	c.JSON(http.StatusOK, gin.H{"category": newCategory})
}

func NewCategoryBatch(c *gin.Context) {
	db := utils.DB

	var helper struct {
		Pack  string   `form:"pack"`
		Title []string `form:"title"`
	}
	if err := c.ShouldBind(&helper); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var pack model.Pack
	result := db.Preload("Categories").Where("slug LIKE ?", helper.Pack).First(&pack)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, "Pack not found!")
		return
	}

	if helper.Title != nil {
		categories := make([]model.Category, len(helper.Title))
		for key, val := range helper.Title {
			uslug := fmt.Sprintf("%v-%v", val, pack.Slug)
			categories[key].Title = strings.TrimSpace(val)
			categories[key].Slug = slug.MakeLang(uslug, "de")
			categories[key].PackSlug = pack.Slug
		}
		db.Clauses(clause.OnConflict{DoNothing: true}).Create(&categories)
	}

	c.JSON(http.StatusOK, "Categories added")
}

func UpdateCategory(c *gin.Context) {
	db := utils.DB

	var category model.Category
	if err := c.ShouldBind(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if category.Title == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Title cannot be empty!"})
		return
	}

	resultP := db.Where("slug = ?", category.PackSlug).First(&model.Pack{})
	if errors.Is(resultP.Error, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Pack " + category.PackSlug + " not found!"})
		return
	}

	var oldCategory model.Category
	result := db.Where("slug = ?", category.Slug).First(&oldCategory)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Category " + category.Slug + " not found!"})
		return
	}

	newCategory := cleanC(&category)
	db.Model(&oldCategory).Updates(&newCategory)

	c.JSON(http.StatusOK, gin.H{"category": newCategory})
}

func DeleteCategory(c *gin.Context) {
	db := utils.DB

	var category model.Category
	if err := c.ShouldBind(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := db.Where("slug = ?", category.Slug).First(&category)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Category " + category.Slug + " not found!"})
		return
	}

	db.Delete(&category)
	c.JSON(http.StatusOK, gin.H{"category": category})
}

func DeleteCategoryBatch(c *gin.Context) {
	db := utils.DB

	var helper struct {
		Pack  string   `form:"pack"`
		Title []string `form:"title"`
	}
	if err := c.ShouldBind(&helper); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var category model.Category
	for _, val := range helper.Title {
		db.Select(clause.Associations).Where("title = ?", val).Delete(&category)
	}

	c.JSON(http.StatusOK, "Categories deleted")
}
