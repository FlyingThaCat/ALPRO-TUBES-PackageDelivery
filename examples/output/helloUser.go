package output

import (
	"fmt"
)

// This is private function
func askName() string {
	var name string
	fmt.Print("Enter your name: ")
	fmt.Scanln(&name)
	return name
}

// This is public function
func HelloUser() {
	name := askName()
	fmt.Printf("Hello, %s!\n", name)
}