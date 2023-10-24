package main

import (
	"os"
	"regexp"
	"testing"
)

const (
	expr = "\\w{2}\\d{11}\\s+\\d{4}\\s+\\d{4}-\\d{2}-\\d{2}\\s+\\d{2}:\\d{2}:\\d{2}\\s+\\d{2}:\\d{2}:\\d{2}\\s+\\d\\s+\\w\\d{2}\\w\\d\\w{2}\\d\\w{2}\\d{2}\\s+\\w+\\s+\\w+\\s+\\w+\\s+\\w"
)

func TestRegex(t *testing.T) {
	re := regexp.MustCompile(expr)
	file, err := os.ReadFile("test/log/example/quectel.txt")
	if err != nil {
		t.Error(err)
	}
	match := re.FindStringSubmatch(string(file))
	if len(match) == 0 {
		t.Error("match is empty")
	}
}
