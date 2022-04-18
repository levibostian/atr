package ui

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/levibostian/bins/store"
)

func ShouldNotHappen(err error) {
	color.Red("[BUG] Something happened that should not have. That means there is probably a bug inside of Purslane.")
	color.Red("Report an issue here: https://github.com/levibostian/Purslane/issues/new, and give this message:")
	fmt.Print(err)
	panic("Exiting...")
}

// HandleError pass in error and we will handle it.
func HandleError(err error) {
	if err != nil {
		Error("\nError encountered!")
		fmt.Println(err)
		os.Exit(1)
	}
}

// Debug - Allows you to put anything you want inside. String, struct, etc. We will print that to the console.
func Debug(format string, args ...interface{}) {
	if store.CliConfig.Debug {
		msg := fmt.Sprintf(format, args...)
		color.Cyan("[DEBUG] " + msg)
	}
}

// Debug - Allows you to put anything you want inside. String, struct, etc. We will print that to the console.
func DebugError(err error) {
	Debug("%v", err)
}

func Abort(format string, args ...interface{}) {
	Error(format, args...)
	os.Exit(1)
}

// Error show a message in red
func Error(format string, args ...interface{}) {
	color.Red(fmt.Sprintf(format, args...))
}

// Message Show a neutral message in white
func Message(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	color.White(msg)
}

// Message Show a success message in green
func Success(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	color.Green(msg)
}
