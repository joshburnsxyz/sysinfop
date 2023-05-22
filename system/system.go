package system

import (
	"encoding/json"
	"fmt"
	"os/user"

	"github.com/zcalusic/sysinfo"
)

// Return information about the current user.
func GetCurrentUser() (&user.User,err) {
	u, err := user.Current()
	if err != nil {
		return &user.User{},err
	}
	return u,nil
}