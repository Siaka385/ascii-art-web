package asciifunc

import (
	"bufio"
	"os"
)

/*
The Filenamevalidate function checks if a given filename exists among predefined valid filenames and corrects common misspellings or variations of these filenames.
*/

func Filenamevalidate(m string) string {
	if !filenameExist(m) {
		return "FILE DOES NOT EXSIT OR CHECK THE SPELLING OF YOUR FILES (e.g standard.txt)"
	}

	if m == "shadow" || m == "shadow.text" {
		return "shadow.txt"
	} else if m == "thinkertoy" || m == "thinkertoy.txt" {
		return "thinkertoy.txt"
	} else {
		return "standard.txt"
	}
}

func filenameExist(m string) bool {
	return m == "shadow.txt" || m == "thinkertoy.txt" || m == "standard.txt"
}

/*
The CheckFileEmpty function determines if a specified file is empty or contains only an empty string.
*/
func CheckFileEmpty(m string) bool {
	file, err := os.Open(m)
	if err != nil {
		return false
	}

	defer file.Close()

	content, err := os.ReadFile(m)
	if err != nil {
		return false
	}

	t := string(content)

	if len(t) == 0 || (len(t) == 1 && t == "") {
		return false
	}
	return true
}

/*
CheckNumberOfLinesInTheFile
checks if the file file has been modified
it checks the number of lines in the files to make sure the line has 855 lines
*/
func CheckNumberOfLinesInTheFile(m string) bool {
	file, err := os.Open(m)
	if err != nil {
		return false
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	countLines := 0

	for scanner.Scan() {
		countLines++
		if countLines > 855 {
			break
		}

	}

	return countLines == 855
}
