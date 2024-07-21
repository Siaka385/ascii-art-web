package asciifunc

// import (
// 	"testing"
// )

// func TestAsciiConverter(t *testing.T) {
// 	result1 := "\n" + AsciiConverter("hello", "standard.txt")
// 	expected1 := "\n" + " _              _   _          \n" +
// 		"| |            | | | |         \n" +
// 		"| |__     ___  | | | |   ___   \n" +
// 		"|  _ \\   / _ \\ | | | |  / _ \\  \n" +
// 		"| | | | |  __/ | | | | | (_) | \n" +
// 		"|_| |_|  \\___| |_| |_|  \\___/  \n" +
// 		"                               \n" +
// 		"                               \n"

// 	result2 := "\n" + AsciiConverter("hello\\nThere", "standard.txt")
// 	expected2 := "\n" + " _              _   _          \n" +
// 		"| |            | | | |         \n" +
// 		"| |__     ___  | | | |   ___   \n" +
// 		"|  _ \\   / _ \\ | | | |  / _ \\  \n" +
// 		"| | | | |  __/ | | | | | (_) | \n" +
// 		"|_| |_|  \\___| |_| |_|  \\___/  \n" +
// 		"                               \n" +
// 		"                               \n" +
// 		" _______   _                           \n" +
// 		"|__   __| | |                          \n" +
// 		"   | |    | |__     ___   _ __    ___  \n" +
// 		"   | |    |  _ \\   / _ \\ | '__|  / _ \\ \n" +
// 		"   | |    | | | | |  __/ | |    |  __/ \n" +
// 		"   |_|    |_| |_|  \\___| |_|     \\___| \n" +
// 		"                                       \n" +
// 		"                                       \n"

// 	if result1 != expected1 {
// 		t.Errorf("expected this  %v  \nbut got this  %v , the length of result is \n %v \n while length of expected is \n %v", expected1, result1, len(result1), len(expected1))
// 	} else if result2 != expected2 {
// 		t.Errorf("expected this  %v  \nbut got this  %v , the length of result is \n %v \n while length of expected is \n %v", expected2, result2, len(result2), len(expected2))
// 	}
// }

// // Input string contains only ASCII characters within the printable range (32-126)
// func TestOnlyPrintableAsciiCharacters(t *testing.T) {
// 	input := "Hello, World!"
// 	result := IsItAnAsciiCharacter(input)
// 	if !result {
// 		t.Errorf("Expected true, got false for input: %s", input)
// 	}
// }
