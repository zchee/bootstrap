package clang

// #include "./clang-c/Index.h"
// #include "go-clang.h"
import "C"

// Evaluation result of a cursor
type EvalResult struct {
	c C.CXEvalResult
}
