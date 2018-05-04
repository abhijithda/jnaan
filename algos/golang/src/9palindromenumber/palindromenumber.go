package palindromenumber

import (
	"fmt"
	"log"
	"os"
)

const logFile = "log.txt"

func init() {
	f, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Printf("Error opening file: %v", err)
		return
	}
	// defer f.Close()

	log.SetOutput(f)
}

func isPalindrome(x int) bool {
	tX := x
	log.Printf("Checking whether %v is palindrome or not...", x)
	if x < 0 {
		return false
	}

	y := tX % 10

	for int(tX/10) > 0 {
		y *= 10
		tX = int(tX / 10)
		y += tX % 10
	}

	log.Printf("%v is palindrome: %v", x, x == y)
	return x == y
}
