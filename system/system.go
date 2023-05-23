package system

import (
	"fmt"
	"os/user"
	"github.com/zcalusic/sysinfo"
)

type ParsedData struct {
	OSInfo string
	MemInfo string
	CPUInfo string
	UserInfo *user.User
}

// Return information about the current user.
func getCurrentUser() (*user.User,error) {
	u, err := user.Current()
	if err != nil {
		return &user.User{},err
	}
	return u,nil
}

// Return sysinfo object
func BuildSysInfo() (*ParsedData,error) {
	si := sysinfo.GetSysInfo()
	u, err := getCurrentUser()
	if err != nil {
		return &ParsedData{},err
	}

	// Create strings
	osistr := fmt.Sprintf("%s %s - v%s", si.OS.Vendor, si.OS.Name, si.OS.Version)
	memistr := fmt.Sprintf("%d", si.Memory.Size)
	cpuistr := fmt.Sprintf("%d Cores /%d Threads", si.CPU.Cores, si.CPU.Threads)

	// Build data object
	p := ParsedData{
		OSInfo: osistr,
		MemInfo: memistr,
		CPUInfo: cpuistr,
		UserInfo: u,
	}

	return &p,nil
}