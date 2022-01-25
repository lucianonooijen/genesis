package replacer

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"git.bytecode.nl/bytecode/genesis/init/prompt"
)

func Replace(replaceValues prompt.ReplaceValues, files []string) error {
	sortedReplaceValues := replaceValues.Sort()

	for _, file := range files {
		err := replaceInFile(file, sortedReplaceValues)
		if err != nil {
			return fmt.Errorf("error replacing values in file '%s': %s", file, err)
		}
	}

	return nil
}

func replaceInFile(filename string, sortedReplaceValues prompt.ReplaceValues) error {
	contents, err := ioutil.ReadFile(filename)
	contentString := string(contents)

	for _, sv := range sortedReplaceValues {
		contentString = strings.ReplaceAll(contentString, sv.OldValue, sv.NewValue)
	}

	// Open file
	file, err := os.OpenFile(filename, os.O_RDWR, 0)
	if err != nil {
		return err
	}
	defer file.Close()

	// Get file information (needed for keeping the correct permissions)
	fileInfo, err := file.Stat()
	if err != nil {
		return err
	}

	// Remove the file (needed to avoid loose bits of data at the end of files when writing)
	err = os.Remove(filename)
	if err != nil {
		return err
	}

	// Write data to file, keeping the permissions the same
	err = os.WriteFile(filename, []byte(contentString), fileInfo.Mode())
	if err != nil {
		return err
	}

	return nil
}
