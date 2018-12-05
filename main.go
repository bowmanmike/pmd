package main

import (
	"log"

	"github.com/bowmanmike/pmd/timer"
)

const (
	cyclesBeforeBreak = 4
)

var (
	work       = timer.CycleInfo{Label: "Work", Duration: 25}
	shortBreak = timer.CycleInfo{Label: "Short Break", Duration: 5}
	longBreak  = timer.CycleInfo{Label: "Long Break", Duration: 20}
)

func main() {
	totalCycleCount := 0

	for {
		for i := 0; i < cyclesBeforeBreak; i++ {
			timer.RunCycle(work)
			if i == cyclesBeforeBreak-1 {
				break
			}
			timer.RunCycle(shortBreak)
		}
		timer.RunCycle(longBreak)

		totalCycleCount++
		log.Printf("%v full cycles completed!\n", totalCycleCount)
	}
}
