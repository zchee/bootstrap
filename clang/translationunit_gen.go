package clang

// #include "./clang-c/Index.h"
// #include "go-clang.h"
import "C"

// A single translation unit, which resides in an index.
type TranslationUnit struct {
	c C.CXTranslationUnit
}
