package clang

// #include "./clang-c/Index.h"
// #include "go-clang.h"
import "C"

// Describes a single preprocessing token.
type Token struct {
	c C.CXToken
}
