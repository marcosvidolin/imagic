package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	// Example usage
	filePath := "images/sample3.jpeg"
	format, err := detectImageFormat(filePath)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("Detected image format: %s\n", format)
}

func detectImageFormat(filePath string) (ImageFormat, error) {
	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		return UnknownFormat, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	// Read the first 16 bytes (enough for most formats)
	header := make([]byte, 16)
	_, err = io.ReadFull(file, header)
	if err != nil && err != io.EOF {
		return UnknownFormat, fmt.Errorf("failed to read file header: %w", err)
	}

	// Match magic numbers
	for format, magic := range MagicNumbers {
		if bytes.HasPrefix(header, magic) {
			// Special case for WEBP to check "WEBP" after "RIFF"
			if format == WEBP && !bytes.Contains(header, []byte("WEBP")) {
				continue
			}

			return format, nil
		}
	}

	return UnknownFormat, nil
}
