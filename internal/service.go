package internal

import (
	"fmt"
	"os"
	"path/filepath"
)

// ConversionService orchestrates the zip-to-pdf conversion process.
type ConversionService struct {
	archiver  *ZipHandler
	generator *PdfHandler
}

// NewConversionService creates a new service.
func NewConversionService(a *ZipHandler, g *PdfHandler) *ConversionService {
	return &ConversionService{
		archiver:  a,
		generator: g,
	}
}

// ConvertZipToPdf handles the entire workflow.
func (s *ConversionService) ConvertZipToPdf(zipPath string, outputDir string) (string, error) {
	// 1. Create temporary directory for extraction
	tempDir := filepath.Join(outputDir, "temp_extracted")
	if err := os.MkdirAll(tempDir, os.ModePerm); err != nil {
		return "", fmt.Errorf("could not create temp directory: %w", err)
	}
	defer os.RemoveAll(tempDir) // Clean up afterward

	// 2. Extract the zip file using the archiver
	fmt.Println("Extracting zip file...")
	if err := s.archiver.Extract(zipPath, tempDir); err != nil {
		return "", fmt.Errorf("failed to extract zip: %w", err)
	}

	// 3. Generate the PDF from the extracted files using the generator
	fmt.Println("Generating PDF...")
	pdfPath := filepath.Join(outputDir, "project_content.pdf")
	if err := s.generator.CreateFromDirectory(tempDir, pdfPath); err != nil {
		return "", fmt.Errorf("failed to generate PDF: %w", err)
	}

	return pdfPath, nil
}
