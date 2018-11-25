package user

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

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
	password, _ = terminal.ReadPassword(0)

	fmt.Println(loginType, username, password)

	req := http.Client{}
	iobuf := strings.NewReader(`
	{
		"type":"test", 
		"username":"peter",
		"password":"test"
	}`)
	res, _ := req.Post("http://127.0.0.1:8080/user/login", "application/json", iobuf)
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))

	return "", nil
}
