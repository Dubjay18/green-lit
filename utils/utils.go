package utils

import (
	"encoding/json"
	"github.com/Dubjay18/green-lit/errs"
	"golang.org/x/crypto/bcrypt"
	"math/rand/v2"
	"net/http"
	"regexp"
)

var (
	boldPattern    = regexp.MustCompile(`\[b\](.*?)\[\/b\]`)
	italicPattern  = regexp.MustCompile(`\[i\](.*?)\[\/i\]`)
	newlinePattern = regexp.MustCompile(`\\n`)
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

	content = newlinePattern.ReplaceAllString(content, "<br>")
	content = boldPattern.ReplaceAllString(content, "<strong>$1</strong>")
	content = italicPattern.ReplaceAllString(content, "<em>$1</em>")

	return content
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GetUUID() (int, error) {
	minValue := 1000
	maxValue := 99999
	//if min > max {
	//	return 0, fmt.Errorf("min (%d) cannot be greater than max (%d)", min, max)
	//}

	// Create a big.Int representing the difference between max and min
	maxBigInt := int64(maxValue - minValue)

	// Generate a random number less than the difference
	randomInt := rand.IntN(int(maxBigInt))

	// Convert the random big.Int to an integer and add the minimum value to get the final result within the range
	return randomInt + minValue, nil

}
