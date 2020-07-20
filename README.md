# Go_Learning_Repo
This is my Go Learnig Repository for all fellow Go noobs. Focused on InfoSec.

## Why Go ?
- Cross Platform
- Fast 
- Easy to write (lets see)

## How to run and compile Go ?
```
$go run hello.go
```

will execute the current programm without needing to compile it

To create a binary we need this command:
```
$go build hello.go
```
Go binaries are quite big by default we can use these parameters to make them smaller:
```
$go build -ldflags "-w -s"
```
## How do we cross compile in Go ?

This command will build an executable for linux 64 bit systems:
```
$GOOS="linux" GOARCH="amd64" go build hello.go
```
Nice !

### Other useful Go Commands

- The "go doc" Command will show you Guidance Papers on any part of Go ($go doc fmt.Println )
- The "go get" Command is an equivalent to python pip 
- The "go fmt" Command will format your code for you, by enforcing the use of proper line breaks, indentation and brace allignment 
- Last but not least the "go vet" Command which will do some heuristic stuff on your code and tell you why your code suckx :)

## Go Syntax

### Primitive Data Types

primitive types include: bool, string, int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, uintptr, byte, rune, float32, float64, complex64, complex128

Declaring a variable:
```
var x = "Hello World"
or
y := "Hello World"

var x = int(42)
or 
y := int(42)
```

### Slices and Maps

