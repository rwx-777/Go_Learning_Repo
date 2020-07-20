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
Go also has more complex data types like slices and maps.
Slices are like arrays that you can dynamically resize and pass to fucntions.
Maps are associative arrays, unordered lists of key/value pairs that lets you quickly look up values for a unique key.

Example of a Slice:
```
var s = make([]string, 0)
s = append(s, "hek")
```

Example of a Map:
```
var m = make(map[string]string)
m["some key"] = "some value"
```

### Pointers, Structs and Interfaces
Pointers in Go work like they do in C. We use the &-operator to retrieve the memory address of some variable, and the * -operator to dereference the address.

Pointer Example:

```
var count = int(42)
var ptr = &count
fmt.Println(ptr)  //prints memory address of count
fmt.Println(*ptr) //prints the value of count
*ptr = 100        //writes new value into count
fmt.Println(count)
```

Struct Example:

```
type Person struct {
  Name string
  Age int
}
func (p *Person) SayHello() {
  fmt.Println("Hello,", p.Name)
}
func main() {
  var guy = new (Person)
  guy.Name = "Dave"
  guy.Age = 21
  guy.SayHello()
}
```

Example of a interface:

```
type Friend interface {
  SayHello()
}

func Greet (f Friend) {
  f.SayHello()
}

func main() {
  var guy = new (Person)
  guy.Name = "Gary"
  Greet(guy)
```


### Control Structures

Basic If-else Statement
```
if z == 1 {
		fmt.Println("Z is equal to 1")
} else {
		fmt.Println("Z is not equal to 1")
}
```
Basic Switch Statement
```
var x = "foo"
switch x {
	case "foo":
		fmt.Println("found foo")
	case "bar":
		fmt.Println("Found bar")
	default:
		fmt.Println("Default case")
}
```
Go doesnt hava a do or while loop. Only a for loop.
For Loop Example:
```
for i := 0; i < 10; i++ {
  fmt.Println(i)
}
```




