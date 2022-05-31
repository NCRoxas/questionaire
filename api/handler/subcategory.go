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

func cleanS(subcategory *model.Subcategory) *model.Subcategory {
	new_slug := fmt.Sprintf("%v-%v", subcategory.Title, subcategory.CatSlug)
	cleanup := &model.Subcategory{
		Title:   strings.TrimSpace(subcategory.Title),
		Slug:    slug.MakeLang(new_slug, "de"),
		CatSlug: subcategory.CatSlug,
	}
	return cleanup
}

func GetSubcategory(c *gin.Context) {
	db := utils.DB

	var subcategory model.Subcategory
	if err := c.ShouldBind(&subcategory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := db.Where("slug = ?", subcategory.Slug).First(&subcategory)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Subcategory " + subcategory.Slug + " not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"subcategory": subcategory})
}

func NewSubcategory(c *gin.Context) {
	db := utils.DB

	var subcategory model.Subcategory
	if err := c.ShouldBind(&subcategory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if subcategory.Title == "" {
		c.JSON(http.StatusNotFound, gin.H{"error": "Title cannot be empty!"})
		return
	}

	result := db.Where("slug = ?", subcategory.CatSlug).First(&model.Category{})
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Category " + subcategory.CatSlug + " not found!"})
		return
	}

	newSubcategory := cleanS(&subcategory)
	db.Create(&newSubcategory)

	c.JSON(http.StatusOK, gin.H{"subcategory": newSubcategory})
}

func NewSubcategoryBatch(c *gin.Context) {
	db := utils.DB

	var helper struct {
		Category string   `form:"category"`
		Title    []string `form:"title"`
	}
	if err := c.ShouldBind(&helper); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var category model.Category
	result := db.Preload("Subcategories").Where("slug = ?", helper.Category).First(&category)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, "Category not found!")
		return
	}

	if helper.Title != nil {
		subcategories := make([]model.Subcategory, len(helper.Title))
		for key, val := range helper.Title {
			uslug := fmt.Sprintf("%v-%v", val, category.Slug)
			subcategories[key].Title = strings.TrimSpace(val)
			subcategories[key].Slug = slug.MakeLang(uslug, "de")
			subcategories[key].CatSlug = category.Slug
		}
		db.Clauses(clause.OnConflict{DoNothing: true}).Create(&subcategories)
	}

	c.JSON(http.StatusOK, "Subcategories added")
}

func UpdateSubcategory(c *gin.Context) {
	db := utils.DB

	var subcategory model.Subcategory
	if err := c.ShouldBind(&subcategory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if subcategory.Title == "" {
		c.JSON(http.StatusNotFound, gin.H{"error": "Title cannot be empty!"})
		return
	}

	resultC := db.Where("slug = ?", subcategory.CatSlug).First(&model.Category{})
	if errors.Is(resultC.Error, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Category " + subcategory.CatSlug + " not found!"})
		return
	}

	var oldSubcategory model.Subcategory
	result := db.Where("slug = ?", subcategory.Slug).First(&oldSubcategory)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Subcategory " + subcategory.Slug + " not found!"})
		return
	}

	newSubcategory := cleanS(&subcategory)
	db.Model(&oldSubcategory).Updates(&newSubcategory)

	c.JSON(http.StatusOK, gin.H{"subcategory": newSubcategory})
}

func DeleteSubcategory(c *gin.Context) {
	db := utils.DB

	var subcategory model.Subcategory
	if err := c.ShouldBind(&subcategory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := db.Where("slug = ?", subcategory.Slug).First(&subcategory)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Subcategory " + subcategory.Slug + " not found!"})
		return
	}

	db.Delete(&subcategory)
	c.JSON(http.StatusOK, gin.H{"subcategory": subcategory})
}

func DeleteSubcategoryBatch(c *gin.Context) {
	db := utils.DB

	var helper struct {
		Category string   `form:"category"`
		Title    []string `form:"title"`
	}
	if err := c.ShouldBind(&helper); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var subcategory model.Subcategory
	for _, val := range helper.Title {
		db.Where("title = ?", val).Delete(&subcategory)
	}

	c.JSON(http.StatusOK, "Subcategories deleted")
}
