package mover

import (
	"fmt"
	"os"
	"os/exec"
)

func findCommandArgs(old, new string) []string {
	nameFind := fmt.Sprintf("*%s*", old)
	sed := fmt.Sprintf(`"s/%s/%s/g"`, old, new)
	cArg := "mv {} $(echo {} | sed " + sed + ")"
	return []string{".", "-depth", "-name", nameFind, "-execdir", "sh", "-c", cArg, ";"}
}

func MoveFilesAndDirectories(newCapitalized, newLowercase string) error {
	cmdCap := exec.Command("find", findCommandArgs("Genesis", newCapitalized)...)
	cmdCap.Stdout = os.Stdout
	cmdCap.Stderr = os.Stderr

	cmdLow := exec.Command("find", findCommandArgs("genesis", newLowercase)...)
	cmdLow.Stdout = os.Stdout
	cmdLow.Stderr = os.Stderr

	err := cmdCap.Run()
	if err != nil {
		return err
	}

	err = cmdLow.Run()
	return err
}
