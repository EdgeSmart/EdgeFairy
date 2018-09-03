package user

import (
	"fmt"

	"golang.org/x/crypto/ssh/terminal"
)

func Login() (string, error) {
	var (
		loginType = ""
		username  = ""
		password  = []byte{}
	)
	fmt.Print("Please select your login type, [test|oauth2]: ")
	fmt.Scan(&loginType)
	fmt.Print("Please input your username: ")
	fmt.Scan(&username)
	fmt.Print("Please input your password: ")
	password, err := terminal.ReadPassword(0)
	fmt.Println(string(password), err)

	return "", nil
}
