package asciifunc

import (
	"bufio"
	"os"
	"strings"
)

/*
AsciiConverter
Concatenate all converted lines and return the resul
*/
func AsciiConverter(inputs, filename string) string {
	return CreateAscii(inputs, filename)
}

/*
Calculate line numbers for each character in the input string n using getLineNumbers.
Initialize variables to store ASCII lines and the final ASCII art.
Loop through 9 lines for each character to construct the ASCII representation.
For each character, append the corresponding line from the file to asciiLines.
After processing all characters for a line, append asciiLines to asciiArt and reset asciiLines.
Convert the 2D slice asciiArt into a single string using Change2DtoString.
*/

func CreateAscii(n, filename string) string {
	numbers := getLineNumbers(n)

	var asciiLines []string
	var asciiArt [][]string
	counterLInes := 0
	moveToNextLine := 0
	for counterLInes <= 8 {
		for i := 0; i < len(numbers); i++ {
			asciiLines = append(asciiLines, PrintLine(numbers[i]+moveToNextLine, filename))
		}
		asciiArt = append(asciiArt, asciiLines)
		asciiLines = []string{}
		moveToNextLine++
		counterLInes++
	}
	return Change2DtoString(asciiArt)
}

/*
IsItAnAsciiCharacter func
checks to ensures every character it part of the ascii value range
*/
func IsItAnAsciiCharacter(m string) bool {
	m = strings.ReplaceAll(m, "\n", "k")
	for _, word := range m {
		if word < 32 || word > 126 {
			return false
		}
	}
	return true
}

/*
START OF PRINTLINES func
The function attempts to open the specified file using os.Open.
If the file opens successfully, it initializes a scanner to read through the file line by line.
It iterates through each line, incrementing a counter until it matches the desired line number (num).
Once the line number matches num, it captures that line's content.
The function returns the captured line and handles any errors that might occur during file reading.
*/

func PrintLine(num int, filename string) string {
	file, err := os.Open(filename)
	if err != nil {
		panic("Error opening the file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	linetoPrint := 0
	var line string
	for scanner.Scan() {
		linetoPrint++
		if linetoPrint == num {
			line = scanner.Text()
		}
	}

	if err := scanner.Err(); err != nil {
		panic("we have an error")
	}
	return line
}

/*
GETLINENUMBERS FUNCTION
e getLineNumbers function calculates the starting line number for ASCII art representation of each character in a given file based on its ASCII value.
*/
func getLineNumbers(word string) []int {
	var lineNumbers []int

	for _, char := range word {
		b := int(char-32)*9 + 2
		lineNumbers = append(lineNumbers, b)
	}

	return lineNumbers
}

func Change2DtoString(m [][]string) string {
	b := ""

	for i := 0; i < len(m); i++ {
		b += strings.Join(m[i], "") + "\n"
	}

	return b[:len(b)-1]
}
