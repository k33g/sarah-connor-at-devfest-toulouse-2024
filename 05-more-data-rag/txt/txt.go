package txt

import "strings"

func SplitTextWithDelimiter(text string, delimiter string) []string {
	return strings.Split(text, delimiter)
}
