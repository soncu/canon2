package main

import (
	"crypto/rand"
	"fmt"
	"os"
)

// shredPath function: shreds the file in the path by overwriting 3 times with random bytes
func shredPath(path string) error {
	file, err := os.OpenFile(path, os.O_WRONLY, 0)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	// Get the file size
	fileInfo, err := file.Stat()
	if err != nil {
		return fmt.Errorf("failed to get file size: %w", err)
	}
	fileSize := fileInfo.Size()

	// Overwrite the file with random data 3 times
	randomData := make([]byte, fileSize)
	for i := 0; i < 3; i++ {
		_, err = rand.Read(randomData)
		if err != nil {
			return fmt.Errorf("failed to generate random data: %w", err)
		}
		_, err = file.WriteAt(randomData, 0)
		if err != nil {
			return fmt.Errorf("failed to write data to file: %w", err)
		}
	}

	// Close and delete the file
	file.Close()
	err = os.Remove(path)
	if err != nil {
		return fmt.Errorf("failed to delete file: %w", err)
	}

	return nil
}

func main() {
	// Example to directly test the shredPath function
	fileName := "randomfile"
	err := os.WriteFile(fileName, []byte("This is a secret file needs to be destroyed!"), 0644)
	if err != nil {
		fmt.Printf("Failed to create file: %v\n", err)
		return
	}

	err = shredPath(fileName)
	if err != nil {
		fmt.Printf("Shred operation failed: %v\n", err)
		return
	}

	fmt.Printf("%s successfully shredded and deleted.\n", fileName)
}

