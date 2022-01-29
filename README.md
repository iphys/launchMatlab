# MATLAB Launcher for Windows

January 2022

## Introduction

`launchMatlab` is a Go program to launch MATLAB available
in your 64-bit Windows machine.

Running `launchMatlab` without any arguments returns the list of
MATLAB found under `C:\Program Fiels\MATLAB` folder.

You can launch a specific version of available MATLABs
by passing the lower two digits of MATLAB release string
followed by letter(s).

Example:

```cmd
launchMatlab.exe 21b
```

## Additional Information

This program is written in Go language.
To build a binary from the source code, install Go,
and run `go build` in the `src` folder as follows.

```powershell
cd src
go build .
```

To modify the source code, use GoLand IDE.
You can use any text editor, but GoLand provides
many convenient features to develop in Go language.
