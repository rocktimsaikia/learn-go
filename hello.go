package main

import f "fmt"
import m "math"
import t "time"

func main() {
	f.Println("hello world")
	f.Println(m.Sqrt(100))
	f.Println(t.Now())
}

func Comment() {
	// f.Println("This does not print")
	/*
		f.Prinln("This does not print too")
		f.Prinln("Because every under this block comment is ignored by the compiler")
	*/
	f.Println("This prints")
}
