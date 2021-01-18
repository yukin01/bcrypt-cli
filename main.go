package main

import (
	"flag"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	flag.Parse()
	password := []byte(flag.Arg(0))

	hashed, _ := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	fmt.Println(string(hashed))
}
