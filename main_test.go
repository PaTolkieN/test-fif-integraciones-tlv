package main_test

import (
	"os"
	"testing"
)

const (
	value       string = "11AB123456789A050"
	valBadRegex string = "11AB12345678-A050"
)

func TestMain(m *testing.M) {

	os.Exit(m.Run())
}
