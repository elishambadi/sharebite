package utils

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/elishambadi/sharebite/config"
	"github.com/gin-gonic/gin"
)

func UploadFile(c *gin.Context, uploadDir string) (string, error) {
	file, err := c.FormFile("file")
	if err != nil {
		return "", err
	}

	err = os.MkdirAll(uploadDir, os.ModePerm)
	if err != nil {
		return "", err
	}

	filePath := filepath.Join(uploadDir, file.Filename)
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		return "", err
	}

	APP_URL := config.AppConfig.AppURL

	fileURL := fmt.Sprintf("%s/uploads/", APP_URL) + file.Filename

	return fileURL, nil
}
