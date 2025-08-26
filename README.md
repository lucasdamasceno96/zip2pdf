# Zip to PDF Converter

A simple command-line utility written in Go that converts the contents of a `.zip` archive into a single, well-formatted PDF document. This tool is perfect for creating a readable snapshot of a project's source code.

## Features

-   Extracts the content of a `.zip` file.
-   Walks through the extracted directory and file structure.
-   Generates a multi-page PDF where each file and directory is clearly marked.
-   Uses a monospaced font for code readability.

## Architecture Overview

This project follows a layered architecture to separate concerns, making it easier to maintain and test:

-   **Handler Layer (`cmd/zip2pdf`)**: Responsible for parsing command-line arguments and handling user interaction. It's the entry point of the application.
-   **Service Layer (`internal/service`)**: Contains the core business logic. It orchestrates the other layers to perform the conversion workflow (extract then generate).
-   **Data/Utility Layers (`internal/archiver`, `internal/generator`)**: These layers handle specific tasks. `archiver` deals with zip file extraction, and `generator` deals with PDF creation.

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
    go build -o zip2pdf ./cmd/zip2pdf/
    ```
    This command will create a binary named `zip2pdf` in the current directory.

## How to Use the Binary

Once you have built the binary, you can use it directly from your terminal.

### Syntax

```sh
./zip2pdf <path-to-your-file.zip>
```

### Example

1.  Make sure you have a zip file ready (e.g., `my-project.zip`).
2.  Run the command:
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