package handler

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gosimple/slug"
	"github.com/ncroxas/questionaire_api/api/model"
	"github.com/ncroxas/questionaire_api/api/utils"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func cleanQ(question *model.Question) *model.Question {
	new_slug := fmt.Sprintf("%v-%v", question.Title, question.SubSlug)
	cleanup := &model.Question{
		Title:       strings.TrimSpace(question.Title),
		Slug:        slug.MakeLang(new_slug, "de"),
		Description: question.Description,
		Answers:     question.Answers,
		SubSlug:     question.SubSlug,
	}
	return cleanup
}

func GetQuestions(c *gin.Context) {
	db := utils.DB

	var questions []model.Question
	result := db.Preload("Answers").Where("sub_slug = ?", c.Query("subcategory")).Find(&questions)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Question not found!"})
		return
	}

	var subcategory model.Subcategory
	db.Where("slug = ?", c.Query("subcategory")).Select("title").First(&subcategory)

	c.JSON(http.StatusOK, gin.H{"questions": questions, "title": subcategory.Title})
}

func GetQuestionList(c *gin.Context) {
	db := utils.DB

	limit, _ := strconv.ParseInt(c.Query("limit"), 10, 0)
	page, _ := strconv.ParseInt(c.Query("page"), 10, 0)

	if limit == 0 {
		limit = 5
	}

	var pagination utils.Pagination
	pagination.Limit = int(limit)
	pagination.Page = int(page)
	pagination.Sort = c.Query("sort")

	var questions []model.Question
	db.Scopes(utils.Paginate(questions, &pagination, db)).Preload("Answers").Find(&questions)

	if len(questions) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No question found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"questions": questions, "pages": pagination.TotalPages})
}

func SearchQuestions(c *gin.Context) {
	db := utils.DB

	limit, _ := strconv.ParseInt(c.Query("limit"), 10, 0)
	page, _ := strconv.ParseInt(c.Query("page"), 10, 0)

	if limit == 0 {
		limit = 5
	}

	var pagination utils.Pagination
	pagination.Limit = int(limit)
	pagination.Page = int(page)
	pagination.Sort = c.Query("sort")

	var questions []model.Question
	db.Scopes(utils.Paginate(questions, &pagination, db)).Preload("Answers").Where("title LIKE ?", "%"+c.Query("search")+"%").Find(&questions)

	c.JSON(http.StatusOK, gin.H{"questions": questions, "pages": pagination.TotalPages})
}

func NewQuestion(c *gin.Context) {
	db := utils.DB

	var question model.Question
	if err := c.ShouldBind(&question); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if question.Title == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Title cannot be empty!"})
		return
	}

	result := db.Where("slug = ?", question.SubSlug).First(&model.Subcategory{})
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Subcategory " + question.SubSlug + " not found!"})
		return
	}

	newQuestion := cleanQ(&question)
	db.Create(&newQuestion)

	c.JSON(http.StatusOK, gin.H{"question": newQuestion})
}

func UpdateQuestion(c *gin.Context) {
	db := utils.DB

	var question model.Question
	if err := c.ShouldBind(&question); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if question.Title == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Title cannot be empty!"})
		return
	}

	resultP := db.Where("slug = ?", question.SubSlug).First(&model.Subcategory{})
	if errors.Is(resultP.Error, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Subcategory " + question.SubSlug + " not found!"})
		return
	}

	var oldQuestion model.Question
	result := db.Preload("Answers").Where("slug = ?", question.Slug).First(&oldQuestion)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Question " + question.Slug + " not found!"})
		return
	}

	for i := range question.Answers {
		var answer model.Answer
		result := db.Where("id = ?", question.Answers[i].ID).Find(&answer)
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Answer " + answer.QSlug + " not found!"})
			return
		}

		db.Model(&answer).Select("image", "answer", "correct").Updates(&question.Answers[i])
	}

	newQuestion := cleanQ(&question)
	db.Model(&oldQuestion).Updates(&newQuestion)

	c.JSON(http.StatusOK, gin.H{"question": newQuestion})
}

func DeleteQuestion(c *gin.Context) {
	db := utils.DB

	var question model.Question
	if err := c.ShouldBind(&question); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := db.Where("slug = ?", question.Slug).First(&question)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Question " + question.Slug + " not found!"})
		return
	}

	db.Delete(&question)
	c.JSON(http.StatusOK, gin.H{"question": question})
}

// Answer Helpers
func AddAnswer(c *gin.Context) {
	db := utils.DB

	var answer model.Answer
	if err := c.ShouldBind(&answer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var question model.Question
	result := db.Where("slug = ?", answer.QSlug).First(&question)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, "Question not found!")
		return
	}

	db.Create(&model.Answer{
		Image:   answer.Image,
		Answer:  answer.Answer,
		Correct: answer.Correct,
		Marked:  answer.Marked,
		QSlug:   question.Slug,
	})

	c.JSON(http.StatusOK, "Answer added")
}

func DeleteAnswer(c *gin.Context) {
	db := utils.DB

	var helper model.Answer
	if err := c.ShouldBind(&helper); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var answer model.Answer
	db.Where("id = ?", helper.ID).Delete(&answer)

	c.JSON(http.StatusOK, "Answer deleted")
}

func Evaluate(c *gin.Context) {
	db := utils.DB

	var user model.User
	db.Preload(clause.Associations).Where("username = ?", c.Query("username")).Find(&user)

	var subcategories []model.Subcategory
	var questions []model.Question
	for i := range user.Packs {
		var packs []model.Pack
		db.Preload("Categories.Subcategories.Questions.Answers").Where("id = ?", user.Packs[i].PackID).Find(&packs)

		var categories []model.Category
		for i := range packs {
			categories = append(categories, packs[i].Categories...)
		}

		for i := range categories {
			subcategories = append(subcategories, categories[i].Subcategories...)
		}

		for i := range subcategories {
			questions = append(questions, subcategories[i].Questions...)
		}
	}

	// Evaluate user answers
	result := make(map[int]interface{})
	if len(user.Answers) > 0 {
		for i := range questions {
			var subcategory model.Subcategory
			db.Where("slug = ?", questions[i].SubSlug).Find(&subcategory)

			var answers = questions[i].Answers
			check := false

			for j := range answers {
				var user_answer model.UserAnswers
				db.Where("answer_id = ? AND user_id = ?", answers[j].ID, user.ID).Find(&user_answer)

				if answers[j].Correct == user_answer.Marked {
					check = true
				} else {
					check = false
					break
				}
			}

			result[i] = map[string]interface{}{
				"question_title":    questions[i].Title,
				"subcategory_title": subcategory.Title,
				"subcategory_slug":  questions[i].SubSlug,
				"correct":           check,
			}
		}
	}

	c.JSON(http.StatusOK, gin.H{"result": result, "total": len(questions), "subcategories": subcategories})
}
