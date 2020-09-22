package clang

// #include "./clang-c/Index.h"
// #include "go-clang.h"
import "C"

// The type of an element in the abstract syntax tree.
type Type struct {
	c C.CXType
}

func (t Type) Kind() TypeKind {
	return TypeKind(t.c.kind)
}
