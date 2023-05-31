package main

// This is the package declaration. Every Go program must start with a package declaration.
// The package name tells the compiler where to find the code for the functions and types used in the program.

import "fmt"

// This is the import declaration. The import statement tells the compiler to load the fmt package into the program.
// The fmt package provides functions for formatting and printing text.

func main() {
  // This is the main function. The main function is the entry point for all Go programs.

  // This line prints the message "Hello, World!" to the console.
  fmt.Println("Hello, World!")
}
