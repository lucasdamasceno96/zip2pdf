package main

import (
	"fmt"
	"os"

	"github.com/lucasdamasceno96/zip2pdf/internal"
)

func main() {
	// 1. Handle CLI arguments
	if len(os.Args) < 2 {
		fmt.Println("Usage: zip2pdf <path-to-file.zip>")
		os.Exit(1)
	}
	zipFilePath := os.Args[1]

	// 2. Setup dependencies (Dependency Injection)
	zipHandler := internal.NewZipHandler()
	pdfHandler := internal.NewPdfHandler()
	conversionService := internal.NewConversionService(zipHandler, pdfHandler)

	// 3. Create output directory and run the service
	outputDir := "output"
	os.MkdirAll(outputDir, os.ModePerm)

	fmt.Printf("Starting conversion for: %s\n", zipFilePath)
	pdfPath, err := conversionService.ConvertZipToPdf(zipFilePath, outputDir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("\n✅ Success! PDF generated at: %s\n", pdfPath)
}
