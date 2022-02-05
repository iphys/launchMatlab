# MATLAB Launcher for Windows

January 2022

## Introduction

`launchMatlab` is a Go program to launch MATLAB available
in your 64-bit Windows machine.

`launchMatlab` assumes that MATLAB is installed under
the default folder `C:\Program Files\MATLAB`.

Running `launchMatlab` without any arguments returns
the list of folder names under `C:\Program Files\MATLAB`.

```powershell
powershell> launchMatlab.exe
MATLAB Launcher in Go.

3 MATLAB install folders were found under C:\Program Files\MATLAB\
R2020b
R2021aU
R2021b
```

You can launch a specific version of available MATLABs
by passing the folder name without "R20".
This is an example to launch from R2021aU:

```powershell
launchMatlab.exe 21aU
```

## Additional Information

This program is written in Go language.
To build a binary from the source code, install Go,
and run `go build` in the `src` folder as follows.

```powershell
cd src
go build .
```
