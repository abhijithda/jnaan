package CampusBikes

import (
	"fmt"
	"log"
	"math"
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

func assignBikes(workers [][]int, bikes [][]int) []int {
	log.Printf("\nGiven workers: %+v, bikes: %+v\n", workers, bikes)
	distance := make([][]int, len(workers))
	for d := range distance {
		distance[d] = make([]int, len(bikes))
	}

	wIdxs := make([]int, len(workers))
	bIdxs := make([]int, len(bikes))
	sdist := int(^uint32(0))
	log.Println("Small distance init: ", sdist)
	swi := -1
	sbi := -1
	for wi := range workers {
		for bi := range bikes {
			distance[wi][bi] = int(math.Abs(float64(workers[wi][0]-bikes[bi][0]))) +
				int(math.Abs(float64(workers[wi][1]-bikes[bi][1])))
			log.Printf("Distance[%d][%d]: %d", wi, bi, distance[wi][bi])
			if distance[wi][bi] < sdist {
				sdist = distance[wi][bi]
				swi = wi
				sbi = bi
				log.Println("sdist:", sdist)
			}
		}
	}

	assignment := make([]int, len(workers))
	assignment[swi] = sbi
	wIdxs[swi] = 1
	bIdxs[sbi] = 1

	sdist = int(^uint32(0))
	// wCnt is number of workers needing bike.
	wCnt := len(wIdxs)
	for wCnt != 0 {
		for wi := range wIdxs {
			if wIdxs[wi] == 1 {
				continue
			}
			for bi := range bIdxs {
				if bIdxs[bi] == 1 {
					continue
				}
				log.Printf("Distance[%d][%d]: %d", wi, bi, distance[wi][bi])
				if distance[wi][bi] < sdist {
					sdist = distance[wi][bi]
					swi = wi
					sbi = bi
					log.Println("sdist:", sdist)
				}
			}
		}
		assignment[swi] = sbi
		wIdxs[swi] = 1
		bIdxs[sbi] = 1
		wCnt--
		log.Printf("Assignments: %v", assignment)
		log.Printf("Assign Workers %v to Bikes %v", wIdxs, bIdxs)
		sdist = int(^uint32(0))
	}
	log.Println("sdist:", sdist)

	return assignment
}

func assignBikesOld(workers [][]int, bikes [][]int) []int {
	log.Printf("\nGiven workers: %+v, bikes: %+v\n", workers, bikes)
	distance := make([][]int, len(workers))
	for d := range distance {
		distance[d] = make([]int, len(bikes))
	}

	wIdxs := map[int]bool{}
	bIdxs := map[int]bool{}
	sdist := int(^uint32(0))
	log.Println("Small distance init: ", sdist)
	for bi := range bikes {
		bIdxs[bi] = true
	}
	swi := -1
	sbi := -1
	for wi := range workers {
		wIdxs[wi] = true
		for bi := range bikes {
			distance[wi][bi] = int(math.Abs(float64(workers[wi][0]-bikes[bi][0]))) +
				int(math.Abs(float64(workers[wi][1]-bikes[bi][1])))
			log.Printf("Distance[%d][%d]: %d", wi, bi, distance[wi][bi])
			if distance[wi][bi] < sdist {
				sdist = distance[wi][bi]
				swi = wi
				sbi = bi
				log.Println("sdist:", sdist)
			}
		}
	}

	assignment := make([]int, len(workers))
	assignment[swi] = sbi
	delete(wIdxs, swi)
	delete(bIdxs, sbi)

	sdist = int(^uint32(0))
	for len(wIdxs) != 0 {
		for wi := range wIdxs {
			for bi := range bIdxs {
				log.Printf("Distance[%d][%d]: %d", wi, bi, distance[wi][bi])
				if distance[wi][bi] < sdist {
					sdist = distance[wi][bi]
					swi = wi
					sbi = bi
					log.Println("sdist:", sdist)
				}
			}
		}
		assignment[swi] = sbi
		delete(wIdxs, swi)
		delete(bIdxs, sbi)
		log.Printf("Assignments: %v", assignment)
		log.Printf("Assign Workers %v to Bikes %v", wIdxs, bIdxs)
		sdist = int(^uint32(0))
	}
	log.Println("sdist:", sdist)

	return assignment
}
