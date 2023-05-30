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

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	allowedChars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()_+-="

	allowedCapitalChars := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	allowedSmallChars := "abcdefghijklmnopqrstuvwxyz"
	allowedDigits := "0123456789"
	allowedSymbols := "!@#$%^&*()_+-="

	newPassword := ""

	randomChar := allowedCapitalChars[rand.Intn(25)]
	newPassword += string(randomChar)
	// fmt.Println(newPassword)

	randomChar = allowedSmallChars[rand.Intn(25)]
	newPassword += string(randomChar)
	// fmt.Println(newPassword)

	randomChar = allowedDigits[rand.Intn(10)]
	newPassword += string(randomChar)
	// fmt.Println(newPassword)

	randomChar = allowedSymbols[rand.Intn(13)]
	newPassword += string(randomChar)
	// fmt.Println(newPassword)

	if passSize < 8 {
		passSize = 8
		println("WARNING: The password must be a minimum of 8 characters in length.")
	}

	for i := 0; i < passSize-4; i++ {
		// Generate a random ASCII character
		randomChar = allowedChars[rand.Intn(62)]

		// Append the random character to the new password string
		newPassword += string(randomChar)
	}

	fmt.Println(newPassword)
}
