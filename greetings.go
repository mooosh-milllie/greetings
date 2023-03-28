package greetings

import "fmt"

func Welcome(name string) string {
	message := fmt.Sprintf("Welcome, %v, good day", name)

	return message
}
