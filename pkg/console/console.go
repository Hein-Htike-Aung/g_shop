package console

import (
	"fmt"
	"os"

	"github.com/mgutz/ansi"
)

// Success print success in green
func Success(msg string) {
	colorOut(msg, "green")
}

// Error print error in red
func Error(msg string) {
	colorOut(msg, "red")
}

// Warning print warning in yellow
func Warning(msg string) {
	colorOut(msg, "yellow")
}

// Exit print error and os.Exit(1)
func Exit(msg string) {
	Error(msg)
	os.Exit(1)
}

// ExitIf helper: exit if err != nil
func ExitIf(err error) {
	if err != nil {
		Exit(err.Error())
	}
}

// colorOut internal colored output
func colorOut(message, color string) {
	fmt.Fprintln(os.Stdout, ansi.Color(message, color))
}
