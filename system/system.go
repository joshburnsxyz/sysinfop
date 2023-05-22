package system

import (
	"os/user"
	"github.com/zcalusic/sysinfo"
)

var si sysinfo.Sysinfo

// Return information about the current user.
func GetCurrentUser() (&user.User,err) {
	u, err := user.Current()
	if err != nil {
		return &user.User{},err
	}
	return u,nil
}

// Return sysinfo object
func GetSysInfo() (sysinfo.Sysinfo) {
	si = sysinfo.GetSysInfo()
	return si
}