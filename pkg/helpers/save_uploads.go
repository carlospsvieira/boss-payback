package helpers

import (
	"boss-payback/pkg/utils"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

func SaveUploadedFile(file *multipart.FileHeader) (string, error) {
	filename := utils.GenerateUUID() + filepath.Ext(file.Filename)
	uploadDir := "./uploads"

	if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
		return "", err
	}

	filePath := filepath.Join(uploadDir, filename)

	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	dst, err := os.Create(filePath)
	if err != nil {
		return "", err
	}
	defer dst.Close()

	if _, err := io.Copy(dst, src); err != nil {
		return "", err
	}

	return filePath, nil
}
