# Go_Learning_Repo
This is my Go Learnig Repository for all fellow Go noobs. Focused on InfoSec.

# Why Go ?
- Cross Platform
- Fast 
- Easy to write (lets see)

# How to run and compile Go ?

$go run hello.go 
will execute the current programm without needing to compile it

To create a binary we need this command:
$go build hello.go

Go binaries are quite big by default we can use these parameters to make them smaller:
$go build -ldflags "-w -s"

# How do we cross compile in Go ?

This command will build an executable for linux 64 bit systems:
$GOOS="linux" GOARCH="amd64" go build hello.go

Nice !
