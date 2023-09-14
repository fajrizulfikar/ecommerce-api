package controllers

import (
	"os"
	"os/exec"
	"testing"
)

func TestMain(m *testing.M) {
	// ensure table exists
	exec.Command("dbmate", "up").Run()

	code := m.Run()

	exec.Command("dbmate", "down").Run()

	os.Exit(code)
}
