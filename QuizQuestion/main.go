/**
Here's a tough one!  Imagine we have two files in the same folder with the same package name.  Will the following code successfully compile?  This might take a little experimentation on your side.  If you do try this out yourself, run your code with the command go run main.go state.go .

In main.go:

package main

func main() {
    printState()
}


In a separate file called state.go:

package main

import "fmt"

func printState() {
    fmt.Println("California")
}
Will this compile and run successfully ?
**/

package main

func main() {
	printState()
}
