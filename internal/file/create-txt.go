package file

import (
	"fmt"
	"io"
	"os"
)

func CreateTxt(name, destination, content string) error {
	// Open the file for writing. Create it or truncate it if it already exists.

	filePath := fmt.Sprintf("%s/%s.txt", destination, name)

	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

	// Write the content to the file
	_, err = io.WriteString(file, content)
	if err != nil {
		return fmt.Errorf("error writing to file: %w", err)
	}

	return nil
}
