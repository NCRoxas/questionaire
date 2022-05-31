package handler

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gosimple/slug"
	"github.com/ncroxas/questionaire_api/api/model"
	"github.com/ncroxas/questionaire_api/api/utils"
	"gorm.io/gorm"
)

func cleanP(pack *model.Pack) *model.Pack {
	cleanup := &model.Pack{
		Title: strings.TrimSpace(pack.Title),
		Slug:  slug.MakeLang(pack.Title, "de"),
	}
	return cleanup
}

func GetPackList(c *gin.Context) {
	db := utils.DB

	var packs []model.Pack
	db.Preload("Categories.Subcategories.Questions.Answers").Find(&packs)

	if len(packs) == 0 {
		c.JSON(http.StatusNotFound, "No packs found")
		return
	}

	var categories []model.Category
	for i := range packs {
		categories = append(categories, packs[i].Categories...)
	}

	var subcategories []model.Subcategory
	for i := range categories {
		subcategories = append(subcategories, categories[i].Subcategories...)
	}

	data := map[string]interface{}{
		"packs":         packs,
		"categories":    categories,
		"subcategories": subcategories,
	}
	c.JSON(http.StatusOK, data)
}

func GetPack(c *gin.Context) {
	db := utils.DB

	var pack model.Pack
	if err := c.ShouldBind(&pack); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := db.Where("slug = ?", pack.Slug).First(&pack)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Pack " + pack.Slug + " not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"pack": pack})
}

func NewPack(c *gin.Context) {
	db := utils.DB

	var pack model.Pack
	if err := c.ShouldBind(&pack); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if pack.Title == "" {
		c.JSON(http.StatusNotFound, gin.H{"error": "Title cannot be empty!"})
		return
	}

	newPack := cleanP(&pack)
	db.Create(&newPack)

	c.JSON(http.StatusOK, gin.H{"pack": newPack})
}

func UpdatePack(c *gin.Context) {
	db := utils.DB

	var pack model.Pack
	if err := c.ShouldBind(&pack); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if pack.Title == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Title cannot be empty!"})
		return
	}

	var oldPack model.Pack
	result := db.Where("slug = ?", pack.Slug).First(&oldPack)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Pack " + pack.Slug + " not found!"})
		return
	}

	newPack := cleanP(&pack)
	db.Model(&oldPack).Updates(&newPack)

	c.JSON(http.StatusOK, gin.H{"pack": newPack})
}

func DeletePack(c *gin.Context) {
	db := utils.DB

	var pack model.Pack
	if err := c.ShouldBind(&pack); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := db.Where("slug = ?", pack.Slug).First(&pack)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Pack " + pack.Slug + " not found!"})
		return
	}

	db.Delete(&pack)
	c.JSON(http.StatusOK, gin.H{"pack": pack})
}
