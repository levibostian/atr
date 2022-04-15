package ui

import (
	"fmt"
	"os"
	"time"

	"github.com/briandowns/spinner"
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

func Abort(message string) {
	Error(message)
	os.Exit(1)
}

// Error show a message in red
func Error(message string) {
	color.Red(message)
}

type Progress interface {
	Done()
}

type SpinnerProgress struct {
	Spinner *spinner.Spinner
}

func (spinner SpinnerProgress) Done() {
	spinner.Spinner.Stop()
}

func MessageProgress(format string, args ...interface{}) Progress {
	spinner := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
	spinner.Suffix = " " + fmt.Sprintf(format, args...)
	spinner.Color("magenta")

	spinner.Start()

	return SpinnerProgress{
		Spinner: spinner,
	}
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
