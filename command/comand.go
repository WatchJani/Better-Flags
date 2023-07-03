package command

import "os"

type fnCommand func()

type Flag struct {
	Name     string
	DefValue interface{}
	Usage    string
}

type FExecutable struct {
	Args         map[string]string
	FlagsFn      map[string]map[string]*Flag
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
		Args:         make(map[string]string),           //Flag key, Flag value
		FlagsFn:      make(map[string]map[string]*Flag), //function, FlagName, AllFlag
		FnExecutable: make(map[string]fnCommand),        //Function Name, That function
	}
}

func (f *FExecutable) AddCommand(command string, fn fnCommand, flags ...*Flag) {
	f.FnExecutable[command] = fn

	f.FlagsFn[command] = make(map[string]*Flag, len(flags))
	for _, value := range flags {
		f.FlagsFn[command][value.Name] = value
	}
}

// ovo je jedna te ista komanda, ali valjda je vako ljepse za napisati, tako su developeri koju su pravili flag package napisali slicno
func CommandLineFunction(key string, fn fnCommand, flags ...*Flag) {
	Executable.AddCommand(key, fn, flags...)
}

// ===============================================================================
func (f *FExecutable) Parser() {
	arguments := os.Args
	f.Args = Convert(arguments[2:])
	f.FnExecutable[arguments[1]]()
}

// manje koda za korisnika koji bude koristio ovaj tool, isti je kod, nema razlike
func Parser() {
	Executable.Parser()
}

// ===============================================================================

func Convert(args []string) map[string]string {
	systemFlags := make(map[string]string)

	return systemFlags
}

// func GetFlag(name string) interface{} {
// 	return Executable.Args[name]
// }
