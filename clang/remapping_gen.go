package clang

// #include "./clang-c/Index.h"
// #include "go-clang.h"
import "C"

// A remapping of original source files and their translated files.
type Remapping struct {
	c C.CXRemapping
}
