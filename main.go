package main

import (
	"fmt"
	clf "root/command" //Command Line Function
)

func main() {
	FlagC := clf.FBool("c", false, "treba li se ispisati")
	FlagI := clf.FBool("i", false, "Input necega")

	clf.CommandLineFunction("commit", Commit, FlagC, FlagI)
	clf.CommandLineFunction("add", Commit, FlagC, FlagI)

	clf.Parser()
}

func Commit() {
	// i := GetFlag("i")
	fmt.Println("work")
	// fmt.Println(i)
}
