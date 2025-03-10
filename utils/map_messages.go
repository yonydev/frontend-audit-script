package utils

import "strings"

func MapMessagePrinter(messages []string) string {
	return strings.Join(messages, "\n")
}
