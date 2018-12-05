package timer

import (
	"fmt"
	"log"
	"time"

	"github.com/gen2brain/beeep"
)

type CycleInfo struct {
	Label    string
	Duration int
}

func RunCycle(options CycleInfo) {
	alertText := fmt.Sprintf("Starting %s cycle for %v minutes", options.Label, options.Duration)
	beeep.Alert("PMD", alertText, "")
	log.Printf(alertText)

	timer := time.NewTimer(time.Duration(options.Duration) * time.Minute)

	<-timer.C
	fmt.Println("     Done!")
}
