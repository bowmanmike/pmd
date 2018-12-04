package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gen2brain/beeep"
)

const (
	cyclesBeforeBreak = 4
)

type cycleInfo struct {
	Label    string
	Duration int
}

var (
	work       = cycleInfo{"Work", 25}
	shortBreak = cycleInfo{"Short Break", 5}
	longBreak  = cycleInfo{"Long Break", 20}
)

func main() {
	totalCycleCount := 0

	for {
		for i := 0; i < cyclesBeforeBreak; i++ {
			runCycle(work)
			if i == cyclesBeforeBreak-1 {
				break
			}
			runCycle(shortBreak)
		}
		runCycle(longBreak)

		totalCycleCount++
		log.Printf("%v full cycles completed!\n", totalCycleCount)
	}
}

func runCycle(options cycleInfo) {
	alertText := fmt.Sprintf("Starting %s cycle for %v minutes", options.Label, options.Duration)
	beeep.Alert("PMD", alertText, "")
	log.Printf(alertText)

	timer := time.NewTimer(time.Duration(options.Duration) * time.Minute)

	<-timer.C
	fmt.Println("     Done!")
}
