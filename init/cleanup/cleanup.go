package cleanup

import (
	"fmt"
	"os"
)

func Cleanup(newLowercase string) error {
	genesisRootFileName := fmt.Sprintf("./.%s", newLowercase)
	err := os.Remove(genesisRootFileName)
	if err != nil {
		return err
	}

	err = os.RemoveAll("./init")
	if err != nil {
		return err
	}

	return nil
}
