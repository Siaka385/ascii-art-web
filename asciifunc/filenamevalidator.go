package asciifunc

import (
	"crypto/sha256"
	"fmt"
	"io"
	"os"
)

/*
The CheckFileEmpty function determines if a specified file is empty or contains only an empty string.
*/
func IsBannerAccepted(m string) bool {
	file, err := os.Open(m)
	if err != nil {
		return false
	}
	file.Close()

	return true
}

/*
CheckNumberOfLinesInTheFile
checks if the file file has been modified
it checks the number of lines in the files to make sure the line has 855 lines
*/

func CheckFileAuthencity(m string) bool {
	file, err := os.Open(m)
	if err != nil {
		return false
	}
	defer file.Close()

	// Create a SHA-256 hash instance
	hash := sha256.New()

	// Copy the file's contents to the hash function
	if _, err := io.Copy(hash, file); err != nil {
		return false
	}

	// Get the final hash value
	hashBytes := hash.Sum(nil)
	computedHash := fmt.Sprintf("%x", hashBytes) // Convert hash to hex string

	// Known hash value to compare against
	ShadowHash := "26b94d0b134b77e9fd23e0360bfd81740f80fb7f6541d1d8c5d85e73ee550f73"
	standardHash := "e194f1033442617ab8a78e1ca63a2061f5cc07a3f05ac226ed32eb9dfd22a6bf"
	thinkertoy := "64285e4960d199f4819323c4dc6319ba34f1f0dd9da14d07111345f5d76c3fa3"

	if m == "shadow.txt" {
		return computedHash == ShadowHash
	} else if m == "thinkertoy.txt" {
		return computedHash == thinkertoy
	} else if m == "standard.txt" {
		return computedHash == standardHash
	} else {
		return m == "shadow.txt" || m == "thinkertoy.txt" || m == "standard.txt"
	}
}
