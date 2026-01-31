package template_test

import (
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/Chara-X/go/text/template"
)

func ExampleTemplate() {
	type Recipient struct {
		Name, Gift string
		Attended   bool
	}
	t := template.New("letter")
	if err := t.Parse(`Dear {{.Name}}, {{if .Attended}} It was a pleasure to see you at the wedding. {{else}} It is a shame you couldn't make it to the wedding. {{end}} {{with .Gift}} Thank you for the lovely {{.}}. {{end}} Best wishes, Josie`); err != nil {
		panic(err)
	}
	for _, r := range []Recipient{{"Aunt Mildred", "bone china tea set", true}, {"Uncle John", "moleskin pants", false}, {"Cousin Rodney", "", false}} {
		if err := t.Execute(os.Stdout, r); err != nil {
			panic(err)
		}
	}
	// Output:
	// Dear Aunt Mildred,  It was a pleasure to see you at the wedding.   Thank you for the lovely bone china tea set.  Best wishes, JosieDear Uncle John,  It is a shame you couldn't make it to the wedding.   Thank you for the lovely moleskin pants.  Best wishes, JosieDear Cousin Rodney,  It is a shame you couldn't make it to the wedding.   Best wishes, Josie
}
func ExampleTemplate_funcs() {
	t := template.New("t")
	t.Funcs(map[string]any{
		"lower":  strings.ToLower,
		"repeat": func(s string) string { return strings.Repeat(s, 2) },
	})
	if err := t.Parse(`{{ repeat (lower .) }}`); err != nil {
		panic(err)
	}
	if err := t.Execute(os.Stdout, "ABC\n"); err != nil {
		panic(err)
	}
	t.Funcs(map[string]any{
		"repeat": func(s string) string { return strings.Repeat(s, 3) },
	})
	if err := t.Execute(os.Stdout, "DEF\n"); err != nil {
		panic(err)
	}
	// Output:
	// abc
	// abc
	// def
	// def
	// def
}

type templateFile struct {
	name     string
	contents string
}

func createTestDir(files []templateFile) string {
	dir, err := os.MkdirTemp("", "template")
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		f, err := os.Create(filepath.Join(dir, file.name))
		if err != nil {
			panic(err)
		}
		defer f.Close()
		_, err = io.WriteString(f, file.contents)
		if err != nil {
			panic(err)
		}
	}
	return dir
}
func ExampleTemplate_glob() {
	dir := createTestDir([]templateFile{
		{"T0.tmpl", `T0 invokes T1: ({{template "T1.tmpl" nil}})`},
		{"T1.tmpl", `T1 invokes T2: ({{template "T2.tmpl" nil}})`},
		{"T2.tmpl", `This is T2`},
	})
	defer os.RemoveAll(dir)
	t := template.New("T0.tmpl")
	if err := t.ParseGlob(filepath.Join(dir, "*.tmpl")); err != nil {
		panic(err)
	}
	if err := t.Execute(os.Stdout, nil); err != nil {
		panic(err)
	}
	// Output:
	// T0 invokes T1: (T1 invokes T2: (This is T2))
}
func ExampleTemplate_if() {
	type book struct {
		Stars float64
		Name  string
	}
	t := template.New("book")
	t.Funcs(map[string]any{
		"gt": func(a, b float64) bool { return a > b },
	})
	if err := t.Parse(`{{ if (gt .Stars 4.0) }}"{{.Name }}" is a great book.{{ else }}"{{.Name}}" is not a great book.{{ end }}`); err != nil {
		panic(err)
	}
	if err := t.Execute(os.Stdout, book{
		Stars: 4.9,
		Name:  "Good Night, Gopher",
	}); err != nil {
		panic(err)
	}
	// Output:
	// "Good Night, Gopher" is a great book.
}
