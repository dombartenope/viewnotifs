package utils

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

// Helper functions to handle nil values and conversion to string
func StringOrNil(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

func IntOrNilToString(i *int) string {
	if i == nil {
		return ""
	}
	return strconv.Itoa(*i)
}

func SecondsToMinSec(seconds int64) string {
	duration := time.Duration(seconds) * time.Second
	minutes := duration / time.Minute
	seconds = int64(duration % time.Minute / time.Second)
	return fmt.Sprintf("%d:%02d", minutes, seconds)
}

func SaveAuthKeyToFile(key, value string) {

	file, err := os.OpenFile(".env", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	defer file.Close()

	_, err = file.WriteString(fmt.Sprintf("%s=%s\n", key, value))
	if err != nil {
		log.Fatalf("error: %s", err)
	}

	fmt.Println("Saved to .env file successfully")

}

func ClearEnvFile() {
	file, err := os.OpenFile(".env", os.O_TRUNC|os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalf("error opening file: %s", err)
	}
	file.Close() // Close immediately after truncating
	fmt.Println(".env file cleared")
}
