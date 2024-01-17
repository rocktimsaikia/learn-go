package main

import f "fmt"
import m "math"
import t "time"

func main() {
	f.Println("hello world")
	f.Println(m.Sqrt(100))
	f.Println(t.Now())
}

// This function used to demostrate how to comment stuff in Go
// And you can know what this function does by running go doc Comment
func Comment() {
	// f.Println("This does not print")
	/*
		f.Prinln("This does not print too")
		f.Prinln("Because every under this block comment is ignored by the compiler")
	*/
	f.Println("This prints")
}
