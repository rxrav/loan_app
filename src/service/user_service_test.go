package service

import (
	"strconv"
	"testing"
)

func Test_generateUsername(t *testing.T) {
	firstName := "alan"
	lastName := "spark"
	fullNameLen := len(firstName+lastName)

	got := ExportedGenerateUsername(firstName, lastName)
	name := got[:fullNameLen]
	numberAsStr := got[fullNameLen:]

	if name != (firstName+lastName) {
		t.Errorf("got %s, want %s", name, (firstName+lastName))
	}

	_, err := strconv.Atoi(numberAsStr)
	if err != nil {
		t.Errorf("got: %v", numberAsStr)
	}
}
