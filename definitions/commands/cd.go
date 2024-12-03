package commands

import (
	"errors"
	"os"

	"github.com/ASentientBanana/kash/globals"
)

func Cd(args []string) error {

	if len(args) != 1 {
		return errors.New("Ivalid number of arguments")
	}

	err := os.Chdir(args[0])

	if err != nil {
		return err
	}

	dir, err := os.Getwd()

	if err != nil {
		return err
	}

	globals.Wd = dir

	return nil
}
