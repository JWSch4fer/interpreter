package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/JWSch4fer/interpreter/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s\n", user.Username)
	fmt.Println("starting interpreter...")
	repl.Start(os.Stdin, os.Stdout)
}
