package system

import (
	"fmt"
	"os/user"
	"encoding/json"
	"github.com/zcalusic/sysinfo"
)

type ParsedData struct {
	OSInfo string
	MemInfo string
	CPUInfo string
}

// Return information about the current user.
func GetCurrentUser() (&user.User,err) {
	u, err := user.Current()
	if err != nil {
		return &user.User{},err
	}
	return u,nil
}

// Return sysinfo object
func GetSysInfo() (&ParsedData,err) {
	si := sysinfo.GetSysInfo()
	
	// Create strings
	osistr := fmt.Sprintf("%s %s - v%s", si.OS.Vendor, si.OS.Name, si.OS.Version)
	memistr := fmt.Sprintf("%d", si.Memory.Size)
	cpuistr := fmt.Sprintf("%d Cores /%d Threads", si.CPU.Cores, si.CPU.Threads)

	// Build data object
	p := ParsedData{
		OSInfo: osistr,
		MemInfo: memistr,
		CPUInfo: cpuistr
	}

	return &p
}