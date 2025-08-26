package internal

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
)

// ZipHandler is responsible for handling zip file operations.
type ZipHandler struct{}

// NewZipHandler creates a new ZipHandler.
func NewZipHandler() *ZipHandler {
	return &ZipHandler{}
}

// Extract unpacks a zip archive to a destination directory.
func (h *ZipHandler) Extract(source, destination string) error {
	reader, err := zip.OpenReader(source)
	if err != nil {
		return err
	}
	defer reader.Close()

	for _, file := range reader.File {
		filePath := filepath.Join(destination, file.Name)

		if file.FileInfo().IsDir() {
			os.MkdirAll(filePath, os.ModePerm)
			continue
		}

		if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
			return err
		}

		outFile, err := os.Create(filePath)
		if err != nil {
			return err
		}

		zippedFile, err := file.Open()
		if err != nil {
			outFile.Close()
			return err
		}

		_, err = io.Copy(outFile, zippedFile)

		outFile.Close()
		zippedFile.Close()

		if err != nil {
			return err
		}
	}
	return nil
}
