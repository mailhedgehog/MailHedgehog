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

func GetString(proposal string) (string, error) {

	// Print proposal.
	fmt.Print(strings.TrimSpace(proposal) + " ")

	// Read input until new line.
	reader := bufio.NewReader(os.Stdin)
	inputStringValue, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	// Remove new line symbol from resulting string.
	inputStringValue = strings.TrimSuffix(inputStringValue, "\n")

	logManager().Debug(fmt.Sprintf("User input: '%s'", inputStringValue))

	return inputStringValue, nil
}

func GetSecretString(proposal string) (string, error) {

	// Print proposal.
	fmt.Print(strings.TrimSpace(proposal) + " ")

	// Read input wrote by user.
	inputBytesValue, err := terminal.ReadPassword(0)
	if err != nil {
		return "", err
	}

	// Print empty line to fix bug with hidden text.
	fmt.Println(" ")

	return string(inputBytesValue), nil
}
