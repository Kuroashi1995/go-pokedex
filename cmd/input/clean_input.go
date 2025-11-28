package input

import (
	"strings"
)

func CleanInput(text string) []string {
	text_lower := strings.ToLower(text)

	trimmed := strings.Split(strings.TrimSpace(text_lower), " ")

	for i := len(trimmed) - 1; i >= 0; i-- {
		if trimmed[i] == "" {
			trimmed = append(trimmed[:i], trimmed[i+1:]...)
		}
	}
	return trimmed
}
