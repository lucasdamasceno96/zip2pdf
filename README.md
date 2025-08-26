# Zip to PDF Converter

A simple command-line utility written in Go that converts the contents of a `.zip` archive into a single, well-formatted PDF document. This tool is perfect for creating a readable snapshot of a project's source code.

## Features

-   Extracts the content of a `.zip` file.
-   Walks through the extracted directory and file structure.
-   Generates a multi-page PDF where each file and directory is clearly marked.
-   Uses a monospaced font for code readability.

## Architecture Overview

This project maintains a logical layered architecture to separate concerns, implemented within a simplified file structure:

-   **Handler Layer (`cmd/main.go`)**: This is the entry point of the application. It is responsible for parsing command-line arguments and handling all user interaction.

-   **Internal Logic (`internal/`)**: All core logic resides within the `internal` package, with responsibilities separated by file:
    -   **`service.go`**: Acts as the **Service Layer**. It orchestrates the conversion process, calling the utility handlers in the correct order (first extract, then generate).
    -   **`zip_handler.go`**: A utility component responsible solely for extracting `.zip` files.
    -   **`pdf_handler.go`**: A utility component responsible solely for generating the PDF document from a directory structure.

## Prerequisites

-   [Go](https://golang.org/dl/) (version 1.18 or higher).

## Building the Binary

1.  Clone the repository:
    ```sh
    git clone [https://github.com/your-username/zip2pdf.git](https://github.com/your-username/zip2pdf.git)
    cd zip2pdf
    ```

2.  Install the dependencies:
    ```sh
    go mod tidy
    ```

3.  Build the executable binary:
    ```sh
    go build -o zip2pdf ./cmd/
    ```
    This command compiles the `main.go` file inside the `cmd` directory and creates a binary named `zip2pdf` in the project's root folder.

## How to Use the Binary

Once you have built the binary, you can use it directly from your terminal.

### Syntax

```sh
./zip2pdf <path-to-your-file.zip>
```

### Example

1.  Make sure you have a zip file ready (e.g., `my-project.zip`).
2.  Run the command from your project's root directory:
    ```sh
    ./zip2pdf my-project.zip
    ```

3.  The tool will create an `output/` directory and generate the PDF inside it.

    **Example Output:**
    ```
    Starting conversion for: my-project.zip
    Extracting zip file...
    Generating PDF...

    ✅ Success! PDF generated at: output/project_content.pdf
    ```
    You can now open `output/project_content.pdf` to view the contents of your zip file.