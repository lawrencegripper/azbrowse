package style

import "github.com/fatih/color"

func Subtle(s string) string {
	return color.New(color.FgMagenta, color.Faint).Sprint(s)
}

func Title(s string) string {
	return color.New(color.Bold).Sprint(s)
}

func Loading(s string) string {
	return color.New(color.BlinkSlow, color.FgRed).Sprint(s)
}

func Completed(s string) string {
	return color.New(color.FgGreen).Sprint(s)
}

func Header(s string) string {
	return color.New(color.BgBlue, color.FgWhite).Sprint(s)
}
