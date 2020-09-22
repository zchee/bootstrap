package clang

// #include "./clang-c/Index.h"
// #include "go-clang.h"
import "C"

type IdxContainerInfo struct {
	c C.CXIdxContainerInfo
}

func (ici IdxContainerInfo) Cursor() Cursor {
	return Cursor{ici.c.cursor}
}
