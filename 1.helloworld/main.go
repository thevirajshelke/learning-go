/*
This is package declaration. This is the name of the package that we are going to create.
There are two types of packages,
 1. The name is 'main' - This generates a executable file out of the code. This type of package will
    need a main function to be defined.
 2. The name is not 'main' - This generates a non executable file & this package is more like a
    library or reusable code.
*/
package main

/*
This package 'fmt' is a default package availble in Go as a standard library
*/
import "fmt"

func main() {
	fmt.Println("Hello World")
}

/*
Basic file structure of go is,
1. Package declaration
2. Imports (if needed)
3. Defining functions that will be executed by go
*/
