package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
)

const InstallPathRoot = "C:\\Program Files\\MATLAB\\"
const RelPathToProgram = "\\bin\\win64\\matlab.exe"

func main() {

	args := os.Args

	var installedFolders, err = os.ReadDir(InstallPathRoot)
	if err != nil {
		log.Fatal(err)
	}

	var numFolders int = len(installedFolders)

	var matlabFolders []string
	for _, file := range installedFolders {
		matlabFolders = append(matlabFolders, file.Name())
	}

	var fullPathToProgram []string
	for i := range matlabFolders {
		tmpFullPath := InstallPathRoot + matlabFolders[i] + RelPathToProgram
		fullPathToProgram = append(fullPathToProgram, tmpFullPath)
	}

	switch len(args) {
	case 1:
		fmt.Println("MATLAB Launcher in Go. January 2022.")
		fmt.Println()
		if len(installedFolders) == 0 {
			fmt.Println("No MATLAB install folder was found under %v\n", numFolders, InstallPathRoot)
			return
		}
		fmt.Printf("%d MATLAB install folders were found under %v\n", numFolders, InstallPathRoot)
		for i := range installedFolders {
			fmt.Println(matlabFolders[i])
		}
		fmt.Println()
		fmt.Println("To start the specific release of MATLAB,")
		fmt.Println("pass the lower two digits followed by letter(s), e.g. 21b,")
		fmt.Println("as the first argument to this command.")
		fmt.Printf("Then %v for that release is launched.", RelPathToProgram)

	case 2:
		commandStr := InstallPathRoot + "R20" + args[1] + RelPathToProgram
		fmt.Printf("Starting MATLAB at %v", commandStr)
		cmd := exec.Command(commandStr)
		var out bytes.Buffer
		cmd.Stdout = &out
		err := cmd.Start()
		if err != nil {
			log.Fatal(err)
		}

	default:
		log.Println("Too many arguments")
	}

}
