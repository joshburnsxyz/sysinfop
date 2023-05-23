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
	NetInfo string
	MacInfo string
	UserInfo string
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
	// Get system info
	var si sysinfo.SysInfo
	si.GetSysInfo()

	// get user data
	u, err := getCurrentUser()
	if err != nil {
		return &ParsedData{},err
	}

	// Get first (primary) network device
	netDevice := si.Network[0]

	// Create strings
	osistr := fmt.Sprintf("OS: %s - %s", si.OS.Name, si.OS.Release)
	useristr := fmt.Sprintf("USER: %s - %s", u.Username, u.HomeDir)
	memistr := fmt.Sprintf("RAM: %d MB @ %d MT/s", si.Memory.Size, si.Memory.Speed)
	cpuistr := fmt.Sprintf("CPU: %d Cores / %d Threads", si.CPU.Cores, si.CPU.Threads)
	netistr := fmt.Sprintf("NET: %s @ %d MB/s", netDevice.Name, netDevice.Speed)
	macaddress := fmt.Sprintf("MAC ADDR: %s", netDevice.MACAddress)

	// Build data object
	p := ParsedData{
		OSInfo: osistr,
		MemInfo: memistr,
		CPUInfo: cpuistr,
		NetInfo: netistr,
		MacInfo: macaddress,
		UserInfo: useristr,
	}

	return &p,nil
}