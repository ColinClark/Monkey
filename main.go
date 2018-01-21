// colin
//
// 1/21/18
//
// 7:57 PM

package main

import (
	"fmt"
	"monkey/repl"
	"os"
	"os/user"

	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetFormatter(&log.JSONFormatter{})

	usr, err := user.Current()
	if err != nil {
		panic(err)
	}

	log.WithFields(log.Fields{
		"User": usr.Username,
	}).Info("Starting Monkey REPL")

	fmt.Printf("Hello %s! This is the Monkey programming language!\n", usr.Username)
	fmt.Printf("Feel free to type in commands\n")

	repl.Start(os.Stdin, os.Stdout)
}
