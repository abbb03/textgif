package inputprocessor

import (
	"bufio"
	"errors"
	"os"
	"regexp"
	"strings"
)

var ErrInvalidInput = errors.New("inputprocessor: invalid input")

func GetInput() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	inp, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	ok, err := regexp.MatchString(`^([еёЕЁа-яА-Я \n]+)$`, inp)
	if err != nil {
		return "", err
	}
	if !ok {
		return "", ErrInvalidInput
	}
	return strings.ToLower(inp), nil
}

func GetText(text string) (string, error) {
	inp := text
	ok, err := regexp.MatchString(`^([еёЕЁа-яА-Я \n]+)$`, inp)
	if err != nil {
		return "", err
	}
	if !ok {
		return "", ErrInvalidInput
	}
	return strings.ToLower(inp), nil
}
