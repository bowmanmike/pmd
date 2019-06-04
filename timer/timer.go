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
	finishTime := time.Now().Add(time.Duration(options.Duration) * time.Minute)
	alertText := fmt.Sprintf("Starting %s cycle for %v minutes. Cycle will finish at %v", options.Label, options.Duration, finishTime.Format("3:04PM"))
	beeep.Alert("PMD", alertText, "")
	log.Printf(alertText)

	timer := time.NewTimer(time.Duration(options.Duration) * time.Minute)

	<-timer.C
	fmt.Println("     Done!")
}
