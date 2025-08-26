package internal

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/jung-kurt/gofpdf"
)

// PdfHandler is responsible for generating PDF files.
type PdfHandler struct{}

// NewPdfHandler creates a new PdfHandler.
func NewPdfHandler() *PdfHandler {
	return &PdfHandler{}
}

// CreateFromDirectory walks a directory and builds a PDF from its contents.
func (h *PdfHandler) CreateFromDirectory(rootPath, pdfOutputPath string) error {
	pdf := gofpdf.New("P", "mm", "A4", "")

	err := filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		relativePath, _ := filepath.Rel(rootPath, path)
		// Skip the root directory itself in the PDF output
		if relativePath == "." {
			return nil
		}

		pdf.AddPage()

		if info.IsDir() {
			pdf.SetFont("Arial", "B", 14)
			pdf.Cell(0, 10, fmt.Sprintf("DIRECTORY: %s", relativePath))
			return nil
		}

		// File name title
		pdf.SetFont("Arial", "B", 12)
		pdf.Cell(0, 10, fmt.Sprintf("FILE: %s", relativePath))
		pdf.Ln(12)

		// File content
		content, err := os.ReadFile(path)
		if err != nil {
			// If we can't read the file, just note it and continue
			pdf.SetFont("Courier", "I", 9)
			pdf.Cell(0, 5, fmt.Sprintf("[Error reading file: %v]", err))
			return nil
		}

		pdf.SetFont("Courier", "", 9)
		for _, line := range strings.Split(string(content), "\n") {
			// Use MultiCell for automatic line wrapping and handling of long lines
			pdf.MultiCell(0, 5, line, "", "L", false)
		}

		return nil
	})

	if err != nil {
		return err
	}

	return pdf.OutputFileAndClose(pdfOutputPath)
}
