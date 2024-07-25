package asciifunc

import (
	"bufio"
	"os"
)

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
