package utils

import (
	"fmt"
	"time"
)

func ShowDelayedMessage(message string, delay int, clearScreen bool) {
	if clearScreen {
		ClearScreen()
	}
	
	for i := 0; i < delay; i++ {
		ClearScreen()
		fmt.Println(message)
		time.Sleep(1 * time.Second)
	}
	
	if clearScreen {
		ClearScreen()
	}
}