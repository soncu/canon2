package main

import (
	"os"
	"testing"
)

// Test functions for different scenarios
func TestShred(t *testing.T) {
	// 1. Test shredding a valid file
	t.Run("ValidFile", func(t *testing.T) {
		fileName := "testfile.txt"
		err := os.WriteFile(fileName, []byte("This is a test file"), 0644)
		if err != nil {
			t.Fatalf("Failed to create file: %v", err)
		}

		err = shredPath(fileName)
		if err != nil {
			t.Errorf("Shred operation failed: %v", err)
		}

		// Check if the file has been deleted
		if _, err := os.Stat(fileName); !os.IsNotExist(err) {
			t.Errorf("File was not deleted!")
		}
	})

	// 2. Test shredding a non-existent file
	t.Run("NonExistentFile", func(t *testing.T) {
		fileName := "nonexistent.txt"
		err := shredPath(fileName)
		if err == nil {
			t.Errorf("Expected error for non-existent file, but got none.")
		}
	})

	// 3. Test shredding an empty file
	t.Run("EmptyFile", func(t *testing.T) {
		fileName := "emptyfile.txt"
		err := os.WriteFile(fileName, []byte(""), 0644)
		if err != nil {
			t.Fatalf("Failed to create empty file: %v", err)
		}

		err = shredPath(fileName)
		if err != nil {
			t.Errorf("Shred operation failed for empty file: %v", err)
		}

		// Check if the file has been deleted
		if _, err := os.Stat(fileName); !os.IsNotExist(err) {
			t.Errorf("Empty file was not deleted!")
		}
	})

	// 4. Test shredding a read-only file
	t.Run("ReadOnlyFile", func(t *testing.T) {
		fileName := "readonlyfile.txt"
		err := os.WriteFile(fileName, []byte("Read-only file"), 0444) // Read-only permissions
		if err != nil {
			t.Fatalf("Failed to create read-only file: %v", err)
		}

		err = shredPath(fileName)
		if err == nil {
			t.Errorf("Expected error for read-only file, but got none.")
		}

		// Remove the file after the test
		err = os.Remove(fileName)
		if err != nil {
			t.Fatalf("Failed to delete read-only file: %v", err)
		}
	})

	// 5. Test shredding a large file ( 100 MB)
	// You can increase the size but be careful for system storage space and time
	t.Run("LargeFile", func(t *testing.T) {
		fileName := "largefile.txt"
		largeData := make([]byte, 100 * 1024 * 1024) // 100 MB
		err := os.WriteFile(fileName, largeData, 0644)
		if err != nil {
			t.Fatalf("Failed to create large file: %v", err)
		}

		err = shredPath(fileName)
		if err != nil {
			t.Errorf("Shred operation failed for large file: %v", err)
		}

		// Check if the file has been deleted
		if _, err := os.Stat(fileName); !os.IsNotExist(err) {
			t.Errorf("Large file was not deleted!")
		}
	})
}
