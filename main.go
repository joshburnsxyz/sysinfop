package main

import (
	"fmt"
	"log"
	"sysinfop/system"
	tm "github.com/buger/goterm"
)

func main() {
	tm.Clear()

	// get info
	i,err := system.BuildSysInfo()
	if err != nil {
		log.Fatal(err)
	}
	// Create Box with 30% width of current screen, and height of 20 lines
	box := tm.NewBox(30|tm.PCT, 5, 0)

	// Print content to screen
	fmt.Fprint(box, fmt.Sprintf("* %s\n* %s\n* %s\n", i.OSInfo, i.MemInfo, i.CPUInfo))

	// Move box to centre of screen
	tm.Print(tm.MoveTo(box.String(), 10|tm.PCT, 10|tm.PCT))

	tm.Flush()
}