package arvanvod

import "strings"

func RemoveSymbols(input string) string {
	input = strings.Replace(input, "(", "", -1)
	input = strings.Replace(input, ")", "", -1)
	input = strings.Replace(input, "@", "", -1)
	input = strings.Replace(input, "$", "", -1)
	input = strings.Replace(input, "*", "", -1)
	input = strings.Replace(input, "^", "", -1)
	input = strings.Replace(input, "%", "", -1)
	input = strings.Replace(input, "#", "", -1)
	input = strings.Replace(input, "!", "", -1)
	input = strings.Replace(input, "&", "", -1)
	input = strings.Replace(input, "~", "", -1)
	input = strings.Replace(input, ",", "", -1)
	input = strings.Replace(input, "-", "_", -1)
	input = strings.Replace(input, "+", "", -1)
	input = strings.Replace(input, "=", "", -1)
	return input
}
