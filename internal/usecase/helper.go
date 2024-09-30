package usecase

import (
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

func moveFile(dstLoc string, file *multipart.FileHeader) (string, error) {
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	// Destination
	if err := os.MkdirAll(filepath.Dir(dstLoc), 0770); err != nil {
		return "", err
	}
	dst, err := os.Create(dstLoc)
	if err != nil {
		return "", err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return "", err
	}

	return dst.Name(), nil
}
