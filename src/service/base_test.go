package service

import (
	"fmt"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	globalSetup()
	code := m.Run()
	globalTeardown()
	os.Exit(code)
}

func globalSetup() {
	fmt.Println("global setup service")
}

func globalTeardown() {
	fmt.Println("global teardown service")
}
