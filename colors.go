package main

import "github.com/fatih/color"

func Blue(s string) string {
	return color.New(color.FgMagenta).Sprint(s)
}
