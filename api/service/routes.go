package service

import (
	"os"
	"strconv"
	"time"

	cache "github.com/chenyahui/gin-cache"
	"github.com/chenyahui/gin-cache/persist"
	"github.com/gin-gonic/gin"
	"github.com/ncroxas/questionaire_api/api/handler"
	"github.com/ncroxas/questionaire_api/api/utils"
)

func SetupRoutes(r *gin.Engine) {
	memoryStore := persist.NewMemoryStore(1 * time.Minute)

	v1 := r.Group("/api/v1")
	{
		user := v1.Group("/user")
		{
			user.GET("", handler.GetUser)
			user.GET("/batch", handler.GetUserList)
			user.GET("/search", handler.SearchUsers)
			user.POST("", handler.NewUser)
			user.POST("/auth", handler.AuthUser)
			user.POST("/reset", handler.ResetUserRequest)
			user.POST("/verify", handler.ResetUser)
			user.PATCH("", handler.UpdateUser)
			user.DELETE("", handler.DeleteUser)
			user.DELETE("/rels", handler.DeleteUserRels)
		}
		pack := v1.Group("/pack")
		{
			pack.GET("", handler.GetPack)
			pack.GET("/batch", handler.GetPackList)
			pack.POST("", handler.NewPack)
			pack.PATCH("", handler.UpdatePack)
			pack.DELETE("", handler.DeletePack)
		}
		category := v1.Group("/category")
		{
			category.GET("", cache.CacheByRequestURI(memoryStore, 2*time.Second), handler.GetCategory)
			category.POST("", handler.NewCategory)
			category.POST("/batch", handler.NewCategoryBatch)
			category.PATCH("", handler.UpdateCategory)
			category.DELETE("", handler.DeleteCategory)
			category.DELETE("/batch", handler.DeleteCategoryBatch)
		}
		subcategory := v1.Group("/subcategory")
		{
			subcategory.GET("", cache.CacheByRequestURI(memoryStore, 2*time.Second), handler.GetSubcategory)
			subcategory.POST("", handler.NewSubcategory)
			subcategory.POST("/batch", handler.NewSubcategoryBatch)
			subcategory.PATCH("", handler.UpdateSubcategory)
			subcategory.DELETE("", handler.DeleteSubcategory)
			subcategory.DELETE("/batch", handler.DeleteSubcategoryBatch)
		}
		question := v1.Group("/question")
		{
			question.GET("", handler.GetQuestions)
			question.GET("/batch", handler.GetQuestionList)
			question.GET("/search", handler.SearchQuestions)
			question.POST("", handler.NewQuestion)
			question.PATCH("", handler.UpdateQuestion)
			question.DELETE("", handler.DeleteQuestion)
			question.POST("/answer", handler.AddAnswer)
			question.DELETE("/answer", handler.DeleteAnswer)
			question.POST("/eval", handler.Evaluate)
		}
		backup := v1.Group("/backup")
		{
			backup.GET("", handler.ExportDB)
			backup.POST("", handler.ImportDB)
		}
		v1.POST("/upload", utils.ImageUpload)
	}

	// Serve static files in production
	if ok, _ := strconv.ParseBool(os.Getenv("DEV_MODE")); !ok {
		r.Static("/static", "./dist")
		r.NoRoute(func(c *gin.Context) { // Make sure reload works
			c.File("./dist/index.html")
		})
	}

}
