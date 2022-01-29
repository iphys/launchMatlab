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
		fmt.Printf("Found MATLAB install folders: %d \n", numFolders)
		for i := range installedFolders {
			fmt.Println(matlabFolders[i])
		}

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

	/*
		cmd := exec.Command(full)
		//cmd.Stdin = strings.NewReader("some input")
		var out bytes.Buffer
		cmd.Stdout = &out
		err := cmd.Start()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%q\n", out.String())
	*/
}
