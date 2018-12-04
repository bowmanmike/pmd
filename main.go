package main

import (
	"fmt"
	"time"

	"github.com/gen2brain/beeep"
)

const (
	workCycle         = 25
	shortBreak        = 5
	longBreak         = 20
	cyclesBeforeBreak = 4
)

type cycleInfo struct {
	Label    string
	Duration int
}

func main() {
	totalCycleCount := 0

	for {
		for i := 0; i < cyclesBeforeBreak; i++ {
			runCycle(cycleInfo{"Work", 25})
			if i == cyclesBeforeBreak-1 {
				break
			}
			runCycle(cycleInfo{"Short Break", 5})
		}
		runCycle(cycleInfo{"Long Break", 20})

		totalCycleCount++
		fmt.Printf("%v full cycles completed!\n", totalCycleCount)
	}
}

func runCycle(options cycleInfo) {
	alertText := fmt.Sprintf("Starting %s cycle for %v minutes", options.Label, options.Duration)
	beeep.Alert("PMD", alertText, "")
	fmt.Printf(alertText)

	timer := time.NewTimer(time.Duration(options.Duration) * time.Minute)

	<-timer.C
	fmt.Println("  -  Done!")
}
