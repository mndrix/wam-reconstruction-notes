package l0

import "fmt"

type HeapAddress int

type Cell interface {
	RenderCell() string
}

type Machine0 struct {
	Heap      []Cell
	Registers []Cell
}

type Structure struct {
	TermAddress HeapAddress
}

func (s *Structure) RenderCell() string {
	return fmt.Sprintf("STR | %d", s.TermAddress)
}

type Functor struct {
	Name  string
	Arity int
}

func (f *Functor) RenderCell() string {
	return fmt.Sprintf("%s/%d", f.Name, f.Arity)
}

type Variable struct {
	TermAddress HeapAddress
}

func (v *Variable) RenderCell() string {
	return fmt.Sprintf("REF | %d", v.TermAddress)
}
