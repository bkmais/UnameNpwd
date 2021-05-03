package main

import (
	"fmt"
	"os"
)

const (
	usage   = "Usage: [username] [password]"
	UserOK  = "User %q logged in."
	errUser = "User %q is invalid."
	errPwd  = "Password is invalid for user:%q"
)

func main1() {
	username, pwd := "BK", "1234"
	uname, pd := "newton", "9988"

	if len(os.Args) != 3 {
		fmt.Println(usage)
		return
	}

	if os.Args[1] == username && os.Args[2] == pwd {
		fmt.Printf(UserOK, os.Args[1])
	} else if os.Args[1] == uname && os.Args[2] == pd {
		fmt.Printf(UserOK, os.Args[1])
	} else if os.Args[1] != username {
		fmt.Printf(errUser, os.Args[1])
	} else if os.Args[2] != pwd {
		fmt.Printf(errPwd, os.Args[1])
	}

}
