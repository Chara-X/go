package expvar

import (
	"encoding/json"
	"expvar"
	"math"
	"sync"
	"sync/atomic"
)

var _ expvar.Var

// [expvar.Var]
type Var interface{ String() string }

// [expvar.Int]
type Int struct{ i atomic.Int64 }

// [expvar.NewInt]
func NewInt(name string) *Int {
	i := &Int{}
	Publish(name, i)
	return i
}

// [expvar.Int.Add]
func (i *Int) Add(delta int64) { i.i.Add(delta) }

// [expvar.Int.Set]
func (i *Int) Set(value int64) { i.i.Store(value) }

// [expvar.Int.Value]
func (i *Int) Value() int64 { return i.i.Load() }

// [expvar.Int.String]
func (i *Int) String() string {
	b, _ := json.Marshal(i.Value())
	return string(b)
}

// [expvar.Float]
type Float struct{ f atomic.Uint64 }

// [expvar.NewFloat]
func NewFloat(name string) *Float {
	f := &Float{}
	Publish(name, f)
	return f
}

// [expvar.Float.Set]
func (f *Float) Set(value float64) { f.f.Store(math.Float64bits(value)) }

// [expvar.Float.Value]
func (f *Float) Value() float64 { return math.Float64frombits(f.f.Load()) }

// [expvar.Float.String]
func (v *Float) String() string {
	b, _ := json.Marshal(v.Value())
	return string(b)
}

// [expvar.String]
type String struct{ s atomic.Value }

// [expvar.NewString]
func NewString(name string) *String {
	s := &String{}
	Publish(name, s)
	return s
}

// [expvar.String.Set]
func (s *String) Set(value string) { s.s.Store(value) }

// [expvar.String.Value]
func (s *String) Value() string { return s.s.Load().(string) }

// [expvar.String.String]
func (s *String) String() string {
	b, _ := json.Marshal(s.Value())
	return string(b)
}

// [expvar.Map]
type Map struct{ m sync.Map }

// [expvar.NewMap]
func NewMap(name string) *Map {
	m := &Map{}
	Publish(name, m)
	return m
}

// [expvar.Map.Add]
func (m *Map) Add(key string, delta int64) {
	v, _ := m.m.LoadOrStore(key, new(Int))
	(v.(*Int)).Add(delta)
}

// [expvar.Map.Set]
func (m *Map) Set(key string, value Var) { m.m.Store(key, value) }

// [expvar.Map.Delete]
func (m *Map) Delete(key string) { m.m.Delete(key) }

// [expvar.Map.Get]
func (m *Map) Get(key string) Var {
	v, _ := m.m.Load(key)
	return v.(Var)
}

// [expvar.Map.Do]
func (m *Map) Do(f func(key string, value Var)) {
	m.m.Range(func(key, value any) bool {
		f(key.(string), value.(Var))
		return true
	})
}

// [expvar.Map.String]
func (m *Map) String() string {
	v := make(map[string]json.RawMessage)
	m.Do(func(key string, value Var) {
		v[key] = json.RawMessage(value.String())
	})
	b, _ := json.Marshal(v)
	return string(b)
}

// [expvar.Func]
type Func func() any

// [expvar.Func.Value]
func (f Func) Value() any { return f() }

// [expvar.Func.String]
func (f Func) String() string {
	b, _ := json.Marshal(f())
	return string(b)
}
