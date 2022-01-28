package main

import (
	"git.bytecode.nl/bytecode/genesis/init/cleanup"
	"git.bytecode.nl/bytecode/genesis/init/listfiles"
	"git.bytecode.nl/bytecode/genesis/init/mover"
	"git.bytecode.nl/bytecode/genesis/init/opengenesis"
	"git.bytecode.nl/bytecode/genesis/init/prompt"
	"git.bytecode.nl/bytecode/genesis/init/replacer"
)

func main() {
	// Change the working directory to the Genesis root
	opengenesis.ChangeToGenesisRootProject()

	// Show prompts to user to get new values that should be replaced throughout the project
	replaceValues, err := prompt.GetReplaceValues()
	if err != nil {
		panic(err)
	}

	// Generate a list of all files (excluding directories, including hidden files) that should be search-replaced
	projectFiles, err := listfiles.RecursivelyFromWorkingDirectory()
	if err != nil {
		panic(err)
	}

	// Replace values in all project files
	err = replacer.Replace(replaceValues, projectFiles)
	if err != nil {
		panic(err)
	}

	// Move files and directories that have the name Genesis in them
	newCapitalized := replaceValues.FetchByID(prompt.ReplaceValueIdNameCapitalized).NewValue
	newLowercase := replaceValues.FetchByID(prompt.ReplaceValueIdNameLowercase).NewValue
	err = mover.MoveFilesAndDirectories(newCapitalized, newLowercase)
	if err != nil {
		panic(err)
	}

	// Clean up files that are particular to the Genesis scaffold stage that should not be kept in the project
	err = cleanup.Cleanup(newLowercase)
	if err != nil {
		panic(err)
	}
}
