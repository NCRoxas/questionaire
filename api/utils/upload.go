package utils

import (
	"bytes"
	"fmt"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func ImageUpload(c *gin.Context) {
	upload, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No file received"})
		return
	}

	uploadDir := "./dist/images/vendor/"
	if ok, _ := strconv.ParseBool(os.Getenv("DEV_MODE")); ok {
		uploadDir = "./web/public/images/vendor/"
	}

	if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
		os.MkdirAll(uploadDir, 0755)
	}

	// Decode uploaded file to image
	file, _ := upload.Open()
	fileByte, _ := ioutil.ReadAll(file)
	contentType := http.DetectContentType(fileByte)

	switch contentType {
	case "image/png":
		if err := c.SaveUploadedFile(upload, uploadDir+upload.Filename); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to save file"})
			return
		}
	case "image/jpeg":
		img, err := jpeg.Decode(bytes.NewReader(fileByte))
		if err != nil {
			log.Println(err)
		}

		path := fmt.Sprintf(uploadDir+"%s.png", strings.Split(upload.Filename, ".")[0])
		out, _ := os.Create(path)
		defer out.Close()
		if err := png.Encode(out, img); err != nil {
			log.Println(err)
		}
	}
	c.JSON(http.StatusOK, "Finished")
}
