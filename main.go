package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/glassrye/tim/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Well, %s. Some people call me.... TIM?\n", user.Username)
	fmt.Printf("Tell me what to do, Daddy...\n")
	repl.Start(os.Stdin, os.Stdout)
}