package userInput

import (
	"bufio"
	"fmt"
	"github.com/mailhedgehog/logger"
	"golang.org/x/crypto/ssh/terminal"
	"os"
	"strings"
)

var configuredLogger *logger.Logger

func logManager() *logger.Logger {
	if configuredLogger == nil {
		configuredLogger = logger.CreateLogger("userInput")
	}
	return configuredLogger
}

func Get(proposal string) (string, error) {
	fmt.Print(strings.TrimSpace(proposal) + " ")
	reader := bufio.NewReader(os.Stdin)
	inputValue, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	inputValue = strings.TrimSuffix(inputValue, "\n")

	logManager().Debug(fmt.Sprintf("User input: '%s'", inputValue))

	return inputValue, nil
}

func GetSilent(proposal string) (string, error) {
	fmt.Print(strings.TrimSpace(proposal) + " ")
	inputValue, err := terminal.ReadPassword(0)
	if err != nil {
		return "", err
	}
	fmt.Println(" ")
	return string(inputValue), nil
}
