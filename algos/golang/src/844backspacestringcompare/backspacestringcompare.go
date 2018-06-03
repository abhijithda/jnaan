package backspacestringcompare

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

func getStringAfterBackspace(st string) string {
	log.Println("Normalizing string: ", st)

	if len(st) <= 0 {
		return ""
	}

	var res string
	for i := 0; i < len(st); i++ {
		log.Println("Processing character: ", string(st[i]))
		if st[i] == '#' {
			log.Printf("Length(%v) of res string(%v)\n", len(res), res)
			if len(res) == 1 {
				res = ""
			} else if len(res) > 1 {
				log.Printf("Processing backspace # for %v(%v): %v",
					res, len(res), res[0:len(res)-1])
				res = res[0 : len(res)-1]
				log.Println("Resulting string: ", res)
			}
			continue
		}
		res = res[:] + string(st[i])
	}
	return res
}

func backspaceCompare(S string, T string) bool {
	stS := getStringAfterBackspace(S)
	stT := getStringAfterBackspace(T)

	if stS == stT {
		return true
	}
	return false
}
