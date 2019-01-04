package style

import "github.com/fatih/color"

// Subtle use magenta and faint to format the text
func Subtle(s string) string {
	return color.New(color.FgMagenta, color.Faint).Sprint(s)
}

// Seperator use magenta and faint to format the text
func Seperator(s string) string {
	return color.New(color.FgBlack, color.Faint, color.Concealed).Sprint(s)
}

// Title make the text bold
func Title(s string) string {
	return color.New(color.Bold).Sprint(s)
}

// Loading make the text red and blink
func Loading(s string) string {
	return color.New(color.BlinkSlow, color.FgRed).Sprint(s)
}

// Completed make the text green
func Completed(s string) string {
	return color.New(color.FgGreen).Sprint(s)
}

// Header make the background blue and text white
func Header(s string) string {
	return color.New(color.BgBlue, color.FgWhite).Sprint(s)
}
