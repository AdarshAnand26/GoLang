package main

import (
	"fmt"
	"os/user"
	"pkg/auth"
	"github.com/fatih/color"
)

func main() {
	auth.LoginWith("Adarsh", "suru")
	session := auth.GetSession()

	fmt.Println("Session", session)

	user := user.User{
		Username: "sbisadb2gmail.com",
		Name:     "suru",
	}
	color.Red(user.Username)
	color.Green("suru")
	fmt.Println(user.Username)
}
