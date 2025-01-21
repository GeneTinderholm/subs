package main

import (
	"fmt"
	"os"
	"subs/shift"
)

func main() {
	if len(os.Args) < 2 {
		printUsageAndDie()
	}
	args := os.Args[1:]
	var err error
	switch args[0] {
	case shift.Command:
		err = shift.Do(args[1:]...)
	default:
		printUsageAndDie()
	}
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(-1)
	}
}

func printUsageAndDie() {
	_, _ = fmt.Fprintf(os.Stderr, `
Usage: subs {command} args...

Current commands:
%s
`, shift.Help)
	os.Exit(-1)
}
