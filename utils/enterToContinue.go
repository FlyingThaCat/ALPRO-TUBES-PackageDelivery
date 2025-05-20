package utils

import "fmt"

func EnterToContinue() {
	// press enter to continue
	var input string
	fmt.Println("Tekan Enter untuk melanjutkan...")
	fmt.Scanln(&input)
	fmt.Println("Melanjutkan...")
	fmt.Print("\n\n")
}