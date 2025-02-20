package util

import (
	"fmt"
	"os"
	"strings"
)

var activeDebugValues = map[string]bool{
	"1":  true,
	"on": true,
	"yes": true,
}

func Debug(a ...any) {
	value, isSet := os.LookupEnv("DEBUG")
	if isSet && activeDebugValues[strings.ToLower(value)] {
		fmt.Println(a...)
	}
}
