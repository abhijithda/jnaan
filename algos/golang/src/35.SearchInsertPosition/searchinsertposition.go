package searchinsertposition

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

func searchInsert(nums []int, target int) int {
	log.Printf("Given nums: %d; target: %d", nums, target)
	startI := 0
	if len(nums) == 0 {
		return 0
	}
	endI := len(nums) - 1

	for {
		chkI := (startI + endI) / 2
		log.Printf("Checking index:%d of start:%d & end:%d", chkI, startI, endI)
		if nums[chkI] == target {
			return chkI
		}

		if startI == endI {
			// For case where nums[chkI] is lesser than target, return next index.
			// 	If last element, then also, send next index as the place to insert.
			if nums[chkI] < target {
				return chkI + 1
			}
			// For case where nums[chkI] is greater than target, return current index.
			return chkI
		}

		if nums[chkI] < target {
			startI = chkI + 1
			log.Println("Start index set to", startI)
		} else if nums[chkI] > target {
			endI = chkI - 1
			if endI < startI {
				endI = startI
			}
			log.Println("End index set to", endI)
		}

	}

}
