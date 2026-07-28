package utils

import (
	"fmt"
)

// ANSI Escape Codes for Terminal Colors
const (
	Reset  = "\033[0m"
	Blue   = "\033[34m"
	Yellow = "\033[33m"
	Green  = "\033[32m"
	Grey	= "\033[37m"
)

// Helper function to map a color string to its ANSI code
func getColorCode(color string) string {
	switch color {
	case "blue":
		return Blue
	case "yellow":
		return Yellow
	case "green":
		return Green
	case "grey":
		return Grey
	default:
		return Reset
	}
}

func PrintInColor(message string, colorName string) {
	colorCode := getColorCode(colorName)
	fmt.Printf("%s %s %s\n", colorCode, message, Reset)
}

