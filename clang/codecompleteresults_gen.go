package clang

// #include "./clang-c/Index.h"
// #include "go-clang.h"
import "C"
import (
	"reflect"
	"unsafe"
)

/*
	Contains the results of code-completion.

	This data structure contains the results of code completion, as
	produced by clang_codeCompleteAt(). Its contents must be freed by
	clang_disposeCodeCompleteResults.
*/
type CodeCompleteResults struct {
	c C.CXCodeCompleteResults
}

// The code-completion results.
func (ccr CodeCompleteResults) Results() []CompletionResult {
	var s []CompletionResult
	gos_s := (*reflect.SliceHeader)(unsafe.Pointer(&s))
	gos_s.Cap = int(ccr.c.NumResults)
	gos_s.Len = int(ccr.c.NumResults)
	gos_s.Data = uintptr(unsafe.Pointer(ccr.c.Results))

	return s
}

// The number of code-completion results stored in the Results array.
func (ccr CodeCompleteResults) NumResults() uint32 {
	return uint32(ccr.c.NumResults)
}
