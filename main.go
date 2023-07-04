package main

import (
	"flag"
	"fmt"
	clf "root/command" //clf => Command Line Function
)

func main() {
	FlagC := clf.FBool("c", false, "treba li se ispisati")
	FlagI := clf.FBool("i", false, "Input necega")

	clf.CommandLineFunction("commit", Commit, FlagC, FlagI)
	clf.CommandLineFunction("add", Commit, FlagC, FlagI)

	flag.Parse()

	clf.Parse()
}

func Commit() {
	// i := GetFlag("i")
	fmt.Println("work")
	// fmt.Println(i)
}
