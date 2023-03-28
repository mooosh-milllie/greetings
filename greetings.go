package greetings

import (
	"errors"
	"fmt"
)

func Welcome(name string) (string, error) {
	if name == "" {
		return "", errors.New("name not specified")
	}
	message := fmt.Sprintf("Welcome, %v, good evening!", name)

	return message, nil
}
