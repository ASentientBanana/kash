package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/ASentientBanana/kash/definitions"
	"github.com/ASentientBanana/kash/globals"
	"github.com/ASentientBanana/kash/handlers"
	"github.com/logrusorgru/aurora/v4"
)

func main() {

	// get current working dir
	wd, err := os.Getwd()
	if err != nil {
		return
	}

	globals.Wd = wd

	// LIBS::
	// box library
	//  go get github.com/Delta456/box-cli-maker/v2
	// Color library
	//go get -u github.com/logrusorgru/aurora/v4
	// Initializing custom definitions not handeled by a third party program but the shell
	definitions.InitializeCommandDefinitions()
	// Reader taking input from the keyboard
	reader := bufio.NewReader(os.Stdin)

	// Game loop
	for {
		// TODO(petar):
		// add graphics and a more complex system
		fmt.Print(aurora.Magenta(globals.Wd + " >> "))

		//Seperating input by new line
		input, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println(err)
		}
		if err = handlers.ExecInput(input); err != nil {
			fmt.Println(err)
		}
	}
}
