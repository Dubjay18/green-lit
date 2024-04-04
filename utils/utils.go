package utils

import (
	"encoding/json"
	"github.com/Dubjay18/green-lit/errs"
	"net/http"
	"regexp"
)

var (
	boldPattern   = regexp.MustCompile(`\[b\](.*?)\[\/b\]`)
	italicPattern = regexp.MustCompile(`\[i\](.*?)\[\/i\]`)
	// Add more patterns for other formatting...
)

func WriteJson(w http.ResponseWriter, code int, i interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	err := json.NewEncoder(w).Encode(i)
	if err != nil {
		return
	}
}
func DecodeJson(r *http.Request, i interface{}) *errs.AppError {
	err := json.NewDecoder(r.Body).Decode(i)
	if err != nil {
		return errs.NewBadRequestError("Invalid request payload")
	}
	return nil
}

func ParseRichText(content string) string {
	// Replace bold formatting
	content = boldPattern.ReplaceAllString(content, "<strong>$1</strong>")

	// Replace italic formatting
	content = italicPattern.ReplaceAllString(content, "<em>$1</em>")

	// ... add replacements for other formatting

	return content
}
