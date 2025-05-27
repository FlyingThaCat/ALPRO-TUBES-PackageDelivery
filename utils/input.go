package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func GetInput(prompt string, allowEmpty bool, messageEmpty string) string {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print(prompt)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if !allowEmpty && input == "" {
			if messageEmpty != "" {
				fmt.Println(messageEmpty)
			} else {
				fmt.Println("Input tidak boleh kosong.")
			}
			continue
		}
		return input
	}
}

func GetInt(prompt string, message string) (int, error) {
	if message == "" {
		message = "Silakan masukkan angka yang valid."
	}

	for {
		input := GetInput(prompt, false, message)
		value, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println(message)
			continue
		}
		return value, nil
	}
}

func GetFloat(prompt string, message string) (float64, error) {
	if message == "" {
		message = "Silakan masukkan angka desimal yang valid."
	}

	for {
		input := GetInput(prompt, false, message)
		value, err := strconv.ParseFloat(input, 64)
		if err != nil {
			fmt.Println(message)
			continue
		}
		return value, nil
	}
}

func GetConfirmation(prompt string, message string) (bool, error) {
	if message == "" {
		message = "Silakan masukkan 'y' atau 'n'."
	}

	for {
		input := strings.ToLower(GetInput(prompt+" (y/n): ", false, message))
		switch input {
		case "y", "yes":
			return true, nil
		case "n", "no":
			return false, nil
		default:
			fmt.Println("Input tidak valid. " + message)
		}
	}
}

func GetString(prompt string, message string) string {
	if message == "" {
		message = "Silakan masukkan string yang valid."
	}

	for {
		input := GetInput(prompt, false, message)
		if input == "" {
			fmt.Println(message)
			continue
		}
		return input
	}
}