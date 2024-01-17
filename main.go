package main

import f "fmt"
import m "math"
import t "time"

func main() {
	f.Println("hello world")
	f.Println(m.Sqrt(100))
	f.Println(t.Now())

	var name string
	var age int8
	var isEngineer bool

	name = "Rocktim Saikia"
	age = 26
	isEngineer = true

	f.Println("Name:", name)
	f.Println("Age:", age)
	f.Println("Is Engineer:", isEngineer)
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
