package template

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"reflect"
	"text/template"
	"text/template/parse"
)

var _ template.Template
var (
	errContinue = errors.New("continue")
	errBreak    = errors.New("break")
)

// [template.Template]
type Template struct {
	*parse.Tree
	name      string
	tmpls     map[string]*Template
	funcs     map[string]any
	execFuncs map[string]reflect.Value
}

// [template.New]
func New(name string) *Template {
	return &Template{name: name, tmpls: make(map[string]*Template), funcs: make(map[string]any), execFuncs: make(map[string]reflect.Value)}
}

// [template.Template.Name]
func (t *Template) Name() string { return t.name }

// [template.Template.Templates]
func (t *Template) Templates() map[string]*Template { return t.tmpls }

// [template.Template.New]
func (t *Template) New(name string) *Template {
	return &Template{name: name, tmpls: t.tmpls, funcs: t.funcs, execFuncs: t.execFuncs}
}

// [template.Template.Funcs]
func (t *Template) Funcs(funcs map[string]any) {
	for name, fn := range funcs {
		t.funcs[name], t.execFuncs[name] = fn, reflect.ValueOf(fn)
	}
}

// [template.Template.Parse]
func (t *Template) Parse(text string) error {
	trees, err := parse.Parse(t.name, text, "{{", "}}", t.funcs)
	if err != nil {
		panic(err)
	}
	t.Tree = trees[t.name]
	t.tmpls[t.name] = t
	return nil
}

// [template.Template.ParseFiles]
func (t *Template) ParseFiles(filenames ...string) error {
	for _, filename := range filenames {
		buf, err := os.ReadFile(filename)
		if err != nil {
			panic(err)
		}
		if name := filepath.Base(filename); t.Name() == name {
			err = t.Parse(string(buf))
		} else {
			err = t.New(name).Parse(string(buf))
		}
		if err != nil {
			panic(err)
		}
	}
	return nil
}

// [template.Template.ParseGlob]
func (t *Template) ParseGlob(pattern string) error {
	filenames, err := filepath.Glob(pattern)
	if err != nil {
		panic(err)
	}
	return t.ParseFiles(filenames...)
}

// [template.Template.Execute]
func (t *Template) Execute(wr io.Writer, data any) error {
	t.execute(t.Tree.Root, wr, reflect.ValueOf(data))
	return nil
}
func (t *Template) execute(node parse.Node, wr io.Writer, data reflect.Value) {
	switch node := node.(type) {
	case *parse.CommentNode:
	case *parse.IfNode:
		if isTrue(t.eval(node.Pipe, data)) {
			t.execute(node.List, wr, data)
		} else if node.ElseList != nil {
			t.execute(node.ElseList, wr, data)
		}
	case *parse.RangeNode:
		var iter = func(elem reflect.Value) {
			defer func() {
				if err := recover(); err != nil && err != errContinue {
					panic(err)
				}
			}()
			t.execute(node.List, wr, elem)
		}
		defer func() {
			if err := recover(); err != nil && err != errBreak {
				panic(err)
			}
		}()
		switch arg := t.eval(node.Pipe, data); arg.Kind() {
		case reflect.Array, reflect.Slice:
			for i := 0; i < arg.Len(); i++ {
				iter(arg.Index(i))
			}
		default:
			panic("not implemented")
		}
	case *parse.WithNode:
		if val := t.eval(node.Pipe, data); isTrue(val) {
			t.execute(node.List, wr, val)
		} else if node.ElseList != nil {
			t.execute(node.ElseList, wr, data)
		}
	case *parse.ContinueNode:
		panic(errContinue)
	case *parse.BreakNode:
		panic(errBreak)
	case *parse.ListNode:
		for _, node := range node.Nodes {
			t.execute(node, wr, data)
		}
	case *parse.TemplateNode:
		t.tmpls[node.Name].Execute(wr, t.eval(node.Pipe, data))
	case *parse.ActionNode:
		fmt.Fprint(wr, t.eval(node.Pipe.Cmds[0], data).Interface())
	case *parse.TextNode:
		wr.Write([]byte(node.Text))
	default:
		panic(node.Type())
	}
}
func (t *Template) eval(node parse.Node, data reflect.Value) reflect.Value {
	switch node := node.(type) {
	case *parse.PipeNode:
		return t.eval(node.Cmds[0], data)
	case *parse.CommandNode:
		switch arg := node.Args[0].(type) {
		case *parse.IdentifierNode:
			var args = []reflect.Value{}
			for _, arg := range node.Args[1:] {
				args = append(args, t.eval(arg, data))
			}
			return t.execFuncs[arg.Ident].Call(args)[0]
		default:
			return t.eval(arg, data)
		}
	case *parse.FieldNode:
		var field = data.FieldByName(node.Ident[0])
		for _, ident := range node.Ident[1:] {
			field = field.FieldByName(ident)
		}
		return field
	case *parse.NumberNode:
		switch {
		case node.IsFloat:
			return reflect.ValueOf(node.Float64)
		case node.IsInt:
			return reflect.ValueOf(int(node.Int64))
		}
	case *parse.BoolNode:
		return reflect.ValueOf(node.True)
	case *parse.StringNode:
		return reflect.ValueOf(node.Text)
	case *parse.NilNode:
		return reflect.ValueOf(nil)
	case *parse.DotNode:
		return data
	}
	panic(node.Type())
}
