package flag

import "flag"

// [flag.FlagSet]
type FlagSet struct {
	name  string
	flags map[string]*flag.Flag
	args  []string
	Usage func()
}

// [flag.NewFlagSet]
func NewFlagSet(name string) *FlagSet {
	return &FlagSet{name: name, flags: map[string]*flag.Flag{}, Usage: func() {}}
}

// [flag.FlagSet.Name]
func (f *FlagSet) Name() string {
	return f.name
}

// [flag.FlagSet.Var]
func (f *FlagSet) Var(value flag.Value, name string, usage string) {
	f.flags[name] = &flag.Flag{Value: value, Name: name, Usage: usage, DefValue: value.String()}
}
func (f *FlagSet) Parse(arguments []string) error {
	f.args = arguments
	for {
		if len(f.args) == 0 {
			break
		}
		var arg = f.args[0]
		if len(arg) < 2 || arg[0] != '-' {
			break
		}
		f.args = f.args[1:]
		if arg == "--" {
			break
		}
		var hyphen = 1
		if arg[1] == '-' {
			hyphen = 2
		}
		var name = arg[hyphen:]
		if name == "help" {
			f.Usage()
			break
		}
		switch v := f.flags[name].Value.(type) {
		case interface {
			flag.Value
			IsBoolFlag() bool
		}:
			v.Set("true")
		default:
			v.Set(f.args[0])
			f.args = f.args[1:]
		}
	}
	return nil
}

// [flag.FlagSet.Args]
func (f *FlagSet) Args() []string {
	return f.args
}
