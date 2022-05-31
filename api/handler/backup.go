package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
	"github.com/ncroxas/questionaire_api/api/model"
	"github.com/ncroxas/questionaire_api/api/utils"
	"gorm.io/gorm/clause"
)

func ExportDB(c *gin.Context) {
	db := utils.DB

	var packs []model.Pack
	db.Preload("Categories.Subcategories.Questions.Answers").Find(&packs)

	var users []model.User
	db.Unscoped().Preload(clause.Associations).Find(&users)

	data := gin.H{"packs": packs, "users": users}
	file := fmt.Sprintf("db-%v.json", time.Now().Format("2006-01-02"))
	backup, _ := json.MarshalIndent(data, "", "\t")
	_ = ioutil.WriteFile(file, backup, 0644)

	c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%v", file))
	c.Writer.Header().Add("Content-Type", "application/octet-stream")
	c.File(file)

	os.Remove(file)
}

func ImportDB(c *gin.Context) {
	db := utils.DB

	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No file received"})
		return
	}

	if err := c.SaveUploadedFile(file, file.Filename); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to save file"})
		return
	}

	content, err := ioutil.ReadFile(file.Filename)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error when opening file"})
		return
	}

	var payload map[string]interface{}
	if err := json.Unmarshal(content, &payload); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error when opening file"})
		return
	}

	var packs []model.Pack
	mapstructure.Decode(payload["packs"], &packs)
	var users []model.User
	mapstructure.Decode(payload["users"], &users)

	db.Clauses(clause.OnConflict{UpdateAll: true}).Create(&packs)
	db.Unscoped().Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.Assignments(map[string]interface{}{"deleted_at": nil}),
	}).Create(&users)

	os.Remove(file.Filename)
	c.JSON(http.StatusOK, gin.H{"message": "File uploaded"})
}
