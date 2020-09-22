package clang

// #include "./clang-c/Index.h"
// #include "go-clang.h"
import "C"

// A particular source file that is part of a translation unit.
type File struct {
	c C.CXFile
}
