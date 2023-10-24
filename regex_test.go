package main

import (
	"os"
	"regexp"
	"testing"
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
