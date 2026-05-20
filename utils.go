package main

import (
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func getTerminalWidth() int {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	if err != nil {
		return 80 // fallback
	}
	// output is "rows cols\n"
	parts := strings.Fields(strings.TrimSpace(string(out)))
	if len(parts) < 2 {
		return 80
	}
	width, err := strconv.Atoi(parts[1])
	if err != nil {
		return 80
	}
	return width
}
