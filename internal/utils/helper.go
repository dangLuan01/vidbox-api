package utils

import (
	"crypto/rand"
	"encoding/base64"
	"html"
	"math/big"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func GetEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != ""{
		return value
	}
	return defaultValue
}

func GetIntEnv(key string, defaultValue int) int {
	value := os.Getenv(key);
	if value == ""{
		
		return defaultValue
	}
	intValue, err := strconv.Atoi(value)
	if err != nil {

		return defaultValue
	}
	
	return intValue
}

func GenerateRandomString(lenght int) (string, error) {
	bytes := make([]byte, lenght)

	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}

	return base64.URLEncoding.EncodeToString(bytes), nil
}

func GenerateRandomInt(lenght int) (string, error) {
	digits := "0123456789"
	number := make([]byte, lenght)

	for i := 0; i < lenght; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(digits))))
		if err != nil {
			return "", err
		}
		number[i] = digits[num.Int64()]
	}

	return string(number), nil
}

func RegexOriginalName(name string) string {
	decodedName := html.UnescapeString(name)

	re := regexp.MustCompile(`\s*\(.*\)|\d+`)
	output := re.ReplaceAllString(decodedName, "")
	return strings.Join(strings.Fields(output), " ")
}

func ExtractNumber(input string) *int {
	defaultSeason := 1
	re 	:= regexp.MustCompile(`\d+`)
	match := re.FindString(input)

	if match != "" {
		number, err := strconv.Atoi(match)
		if err == nil {
			if number > 30 {
				return &defaultSeason			
			}
			return &number
		}
	}

	return &defaultSeason
}