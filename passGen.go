package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

// Trying out git
func main() {

	rand.Seed(time.Now().UnixNano())

	// passGen -size
	passSize, err := strconv.Atoi(os.Args[1])

	fmt.Println(passSize)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	allowedChars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()_+-="

	newPassword := ""

	for i := 0; i < passSize; i++ {
		// Generate a random ASCII character
		randomChar := allowedChars[rand.Intn(62)]

		// Append the random character to the new password string
		newPassword += string(randomChar)
	}

	fmt.Println(newPassword)
}
