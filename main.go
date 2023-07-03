package main

import (
	"fmt"
	"os"
)

func main() {
	FlagC := FBool("c", false, "treba li se ispisati")
	FlagI := FBool("i", false, "Input necega")

	AddCommand("commit", Commit, FlagC, FlagI)
	AddCommand("add", Commit, FlagC, FlagI)

	Parser()
}

type fnCommand func()

type Flag struct {
	Name     string
	DefValue interface{}
	Usage    string
}

type FExecutable struct {
	Args         []string
	FlagsFn      map[string][]*Flag
	FnExecutable map[string]fnCommand
}

func NewFlag(flag string, defValue interface{}, usage string) *Flag {
	return &Flag{
		Name:     flag,
		DefValue: defValue,
		Usage:    usage,
	}
}

func FBool(flag string, defValue bool, usage string) *Flag {
	return NewFlag(flag, defValue, usage)
}

func FInt(flag string, defValue int, usage string) *Flag {
	return NewFlag(flag, defValue, usage)
}

var Executable *FExecutable = NewFExecutable()

func NewFExecutable() *FExecutable {
	return &FExecutable{
		FlagsFn:      make(map[string][]*Flag),
		FnExecutable: make(map[string]fnCommand),
	}
}

func (f *FExecutable) AddCommand(command string, fn fnCommand, flags ...*Flag) {
	f.FlagsFn[command], f.FnExecutable[command] = flags, fn
}

// ovo je jedna te ista komanda, ali valjda je vako ljepse za napisati, tako su developeri koju su pravili flag package napisali slicno
func AddCommand(command string, fn fnCommand, flags ...*Flag) {
	Executable.AddCommand(command, fn, flags...)
}

// ===============================================================================
func (f *FExecutable) Parser() {
	arguments := os.Args
	f.Args = arguments[2:]
	f.FnExecutable[arguments[1]]()
}

// manje koda za korisnika koji bude koristio ovaj tool, isti je kod, nema razlike
func Parser() {
	Executable.Parser()
}

// ===============================================================================

func Commit() {
	//c,i := parser()

	fmt.Println(Executable.Args)
}
