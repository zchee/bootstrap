package clang

// #include "./clang-c/Index.h"
// #include "go-clang.h"
import "C"

type Module struct {
	c C.CXModule
}
