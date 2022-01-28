package opengenesis

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

const GenesisRootIndicator = "root=true"

func isGenesisProjectRoot() bool {
	f, err := ioutil.ReadFile("./.genesis")
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return false
		}
		panic(err)
	}

	if strings.Contains(string(f), GenesisRootIndicator) {
		return true
	}
	return false
}

func ChangeToGenesisRootProject() {
	for {
		wd, err := os.Getwd()
		if err != nil {
			panic(err)
		}

		if isGenesisProjectRoot() {
			fmt.Printf("Root directory found: %s\n", wd)
			return
		}

		if wd == "/" {
			panic("arrived in '/' directory - Genesis root not found")
		}

		fmt.Printf("%s is not the root directory\n", wd)
		err = os.Chdir("..")
		if err != nil {
			panic(err)
		}
	}
}
