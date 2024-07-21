package asciifunc

func Asciigenerate(input []string) string {
	inputs := input[0]

	filenames := input[1]

	return AsciiConverter(inputs, filenames)
}
