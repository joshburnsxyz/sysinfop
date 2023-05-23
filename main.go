package main

import (
	"fmt"
	"sysinfop/system"
	tm "github.com/buger/goterm"
)

func main() {
	tm.Clear()

	// get info
	i := system.BuildSysInfo()

	// Create Box with 30% width of current screen, and height of 20 lines
	box := tm.NewBox(30|tm.PCT, 20, 0)

	// Print content to screen
	fmt.Fprint("* %s\n* %s\n* %s\n", i.OSInfo, i.MemInfo, i.CPUInfo)

	// Move box to centre of screen
	tm.Print(tm.MoveTo(box.String(), 40|tm.PCT, 40|tm.PCT))

	tm.Flush()
}