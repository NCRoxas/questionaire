package handler

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ncroxas/questionaire_api/api/model"
	"github.com/ncroxas/questionaire_api/api/utils"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func GetUser(c *gin.Context) {
	db := utils.DB

	var user model.User
	result := db.Preload(clause.Associations).Where("username = ?", c.Query("username")).First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, "User not found!")
		return
	}

	pack_ids := []uint{}
	for i := range user.Packs {
		pack_ids = append(pack_ids, user.Packs[i].PackID)
	}

	// Filter subscriped packages
	var packs []model.Pack
	db.Preload("Categories.Subcategories").Where("ID IN ?", pack_ids).Find(&packs)

	var categories []model.Category
	for i := range packs {
		categories = append(categories, packs[i].Categories...)
	}

	var subcategories []model.Subcategory
	for i := range categories {
		subcategories = append(subcategories, categories[i].Subcategories...)
	}

	data := gin.H{
		"username":      user.Username,
		"firstname":     user.Firstname,
		"surname":       user.Surname,
		"email":         user.Email,
		"gender":        user.Gender,
		"street":        user.Street,
		"zip":           user.Zip,
		"city":          user.City,
		"state":         user.State,
		"birthday":      user.Birthday.Format("2006-01-02"),
		"role":          user.Role,
		"newsletter":    user.Newsletter,
		"active":        user.Active,
		"user_packs":    user.Packs,
		"user_answers":  user.Answers,
		"packs":         packs,
		"categories":    categories,
		"subcategories": subcategories,
		"lastlogin":     user.LastLogin,
	}
	c.SecureJSON(http.StatusOK, data)
}

func GetUserList(c *gin.Context) {
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

	var users []model.User
	db.Scopes(utils.Paginate(users, &pagination, db)).Preload(clause.Associations).Find(&users)

	if len(users) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No user found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"users": users, "pages": pagination.TotalPages})
}

func SearchUsers(c *gin.Context) {
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

	var users []model.User
	db.Scopes(utils.Paginate(users, &pagination, db)).Preload(clause.Associations).Where("username LIKE ?", "%"+c.Query("search")+"%").Find(&users)

	c.JSON(http.StatusOK, gin.H{"users": users, "pages": pagination.TotalPages})
}

func NewUser(c *gin.Context) {
	db := utils.DB

	var user model.User
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if user.Username == "" || user.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username/Email cannot be empty!"})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	user.Password = string(hash)
	if user.Username == "" || user.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Please provide a username and email for the new user"})
		return
	}

	db.Create(&user)
	c.JSON(http.StatusOK, gin.H{"user": user})
}

func UpdateUser(c *gin.Context) {
	db := utils.DB

	var user model.User
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var oldUser model.User
	result := db.Where("username = ?", user.Username).First(&oldUser)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{"error": "User " + user.Username + " not found!"})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	user.Password = string(hash)

	// Add User Packs
	for i := range user.Packs {
		var user_pack model.UserPacks
		db.Model(&oldUser).Where("pack_id = ?", user.Packs[i].PackID).Association("Packs").Find(&user_pack)
		if user_pack.PackID == 0 {
			db.Model(&oldUser).Association("Packs").Append(&user.Packs[i])
		}
	}

	// Add/Update User Answers
	for i := range user.Answers {
		var user_answer model.UserAnswers

		if user.Answers[i].AnswerID > 0 {
			db.Model(&user_answer).Where("answer_id = ? AND user_id = ?", user.Answers[i].AnswerID, user.ID).Find(&user_answer)
			if user_answer.AnswerID == 0 && user_answer.UserID == 0 {
				db.Model(&oldUser).Association("Answers").Append(&user.Answers[i])
			} else {
				db.Model(&user_answer).Select("Marked").Updates(&user.Answers[i])
			}
		}
	}

	db.Model(&oldUser).Updates(&user)
	c.JSON(http.StatusOK, gin.H{"user": user})
}

func AuthUser(c *gin.Context) {
	db := utils.DB

	var helper struct {
		Login    string `form:"login"`
		Password string `form:"password"`
		Remember bool   `form:"remember"`
	}
	if err := c.ShouldBind(&helper); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user model.User
	result := db.Where("email = ?", helper.Login).Or("username = ?", helper.Login).First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusUnauthorized, "Invalid Credentials")
		return
	}

	if !user.Active {
		c.JSON(http.StatusUnauthorized, "User is not active")
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(helper.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, "Invalid Credentials")
		return
	}

	token, err := utils.CreateToken(user.Username, helper.Remember)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Something went terribly wrong"})
		return
	}

	user.LastLogin = time.Now()
	user.Token = token
	db.Model(&user).Updates(user)

	c.SecureJSON(http.StatusOK, gin.H{"token": user.Token, "role": user.Role})
}

func ResetUserRequest(c *gin.Context) {
	db := utils.DB

	var helper struct {
		Login string `form:"login"`
	}
	if err := c.ShouldBind(&helper); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user model.User
	result := db.Where("email = ?", helper.Login).Or("username = ?", helper.Login).First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, "Invalid Credentials")
		return
	}

	token, err := utils.CreateResetToken(user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Something went terribly wrong"})
		return
	}

	link := fmt.Sprintf("%v/reset/%v", os.Getenv("FRONTEND"), token)

	// Sending mail to user
	body := fmt.Sprintf("Sie erhalten diese Nachricht, da Sie Ihr Passwort zurückgesetzen wollen.\n"+
		"Bitte Nutzen Sie den generierten Link um ihr Passwort zu ändern.\n%v\n"+
		"Der Link ist nur 10 Minuten gültig!\n"+
		"\n - Questionaire GmbH\n", link)
	utils.SendMail(user.Email, body)

	user.Token = token
	db.Model(&user).Updates(user)
	c.SecureJSON(http.StatusOK, user.Username)
}

func ResetUser(c *gin.Context) {
	db := utils.DB

	var user model.User
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	username, err := utils.CheckToken(user.Token)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var oldUser model.User
	result := db.Where("username = ?", username).First(&oldUser)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, "Invalid Credentials")
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Internal server error")
		return
	}

	user.Password = string(hash)
	db.Model(&oldUser).Updates(user)
	c.JSON(http.StatusOK, gin.H{"ok": "Password changed!"})

}

func DeleteUser(c *gin.Context) {
	db := utils.DB

	var user model.User
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := db.Where("username = ?", user.Username).First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{"error": "User " + user.Username + " not found!"})
		return
	}

	db.Delete(&user)
	c.JSON(http.StatusOK, gin.H{"user": user})
}

func DeleteUserRels(c *gin.Context) {
	db := utils.DB

	var helper model.User
	if err := c.ShouldBind(&helper); err != nil {
		log.Print(err)
	}

	var user model.User
	result := db.Where("username = ?", helper.Username).First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, "User not found")
		return
	}

	if len(helper.Packs) > 0 {
		for i := range helper.Packs {
			var user_pack model.UserPacks
			db.Model(&user_pack).Where("pack_id = ?", helper.Packs[i].PackID).Find(&user_pack)
			db.Delete(&user_pack)
		}

	}

	if len(helper.Answers) > 0 {
		for i := range helper.Answers {
			var user_answer model.UserAnswers
			db.Model(&user_answer).Where("answer_id = ?", helper.Answers[i].AnswerID).Find(&user_answer)
			db.Delete(&user_answer)
		}
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}
