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
	// Create Box
	box := tm.NewBox(30|tm.PCT, 7, 0)

	// Print content to screen
	fmt.Fprint(box, fmt.Sprintf("* %s\n* %s\n* %s\n* %s\n* %s\n* %s\n",
	i.OSInfo,
	i.UserInfo,
	i.MemInfo,
	i.CPUInfo,
	i.NetInfo,
	i.MacInfo))

	// Move box to position
	tm.Print(tm.MoveTo(box.String(), 0|tm.PCT, 10|tm.PCT))

	//Print new line after box
	tm.Printf("\n")

	tm.Flush()
}