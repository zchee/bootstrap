package clang

// #include "./clang-c/Index.h"
// #include "go-clang.h"
import "C"

type IdxAttrInfo struct {
	c C.CXIdxAttrInfo
}

func (iai IdxAttrInfo) Kind() IdxAttrKind {
	return IdxAttrKind(iai.c.kind)
}

func (iai IdxAttrInfo) Cursor() Cursor {
	return Cursor{iai.c.cursor}
}

func (iai IdxAttrInfo) Loc() IdxLoc {
	return IdxLoc{iai.c.loc}
}
