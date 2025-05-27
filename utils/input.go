package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func GetInput(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func GetInt(prompt string) (int, error) {
	input := GetInput(prompt)
	return strconv.Atoi(input)
}

func GetFloat(prompt string) (float64, error) {
	input := GetInput(prompt)
	return strconv.ParseFloat(input, 64)
}

func GetConfirmation(prompt string) (bool, error) {
	input := strings.ToLower(GetInput(prompt + " (y/n): "))
	switch input {
		case "y", "yes":
			return true, nil
		case "n", "no":
			return false, nil
		default:
			return false, fmt.Errorf("silakan masukkan 'y' atau 'n'")
		}
}