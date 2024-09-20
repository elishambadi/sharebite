package utils

import (
	"os"
	"path/filepath"

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

	fileURL := "https://example.com/uploads/" + file.Filename

	return fileURL, nil
}
