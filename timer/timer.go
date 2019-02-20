package timer

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

type CycleInfo struct {
	Label    string
	Duration int
}

func Run() {
	for {
	}
}

func RunCycle(options CycleInfo) {
	alertText := fmt.Sprintf("Starting %s cycle for %v minutes", options.Label, options.Duration)
	// beeep.Alert("PMD", alertText, "")
	log.Printf(alertText)

	input := make(chan string)
	// output := make(chan string)

	timer := time.NewTimer(time.Duration(options.Duration) * time.Minute)

	go readInput(input)

	for elem := range input {
		fmt.Println(elem)
		if elem == "pause" {
			timer.Stop()
			break
		}
	}

	// for {
	// 	text := <-input
	// 	if text == "p" {
	// 		fmt.Println(text)
	// 		timer.Stop()
	// 		break
	// 	}
	// 	if !timer.Stop() {
	// 		<-timer.C
	// 	}
	// }

	fmt.Println("     Done!")
}

func readInput(data chan string) {
	buf := bufio.NewReader(os.Stdin)

	for {
		text, err := buf.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		data <- text
	}
}
