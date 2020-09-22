package clang

// #include "./clang-c/Index.h"
// #include "go-clang.h"
import "C"
import "fmt"

// Describes the kind of type
type TypeKind uint32

const (
	// Represents an invalid type (e.g., where no type is available).
	Type_Invalid TypeKind = C.CXType_Invalid
	// A type whose specific kind is not exposed via this interface.
	Type_Unexposed = C.CXType_Unexposed
	// A type whose specific kind is not exposed via this interface.
	Type_Void = C.CXType_Void
	// A type whose specific kind is not exposed via this interface.
	Type_Bool = C.CXType_Bool
	// A type whose specific kind is not exposed via this interface.
	Type_Char_U = C.CXType_Char_U
	// A type whose specific kind is not exposed via this interface.
	Type_UChar = C.CXType_UChar
	// A type whose specific kind is not exposed via this interface.
	Type_Char16 = C.CXType_Char16
	// A type whose specific kind is not exposed via this interface.
	Type_Char32 = C.CXType_Char32
	// A type whose specific kind is not exposed via this interface.
	Type_UShort = C.CXType_UShort
	// A type whose specific kind is not exposed via this interface.
	Type_UInt = C.CXType_UInt
	// A type whose specific kind is not exposed via this interface.
	Type_ULong = C.CXType_ULong
	// A type whose specific kind is not exposed via this interface.
	Type_ULongLong = C.CXType_ULongLong
	// A type whose specific kind is not exposed via this interface.
	Type_UInt128 = C.CXType_UInt128
	// A type whose specific kind is not exposed via this interface.
	Type_Char_S = C.CXType_Char_S
	// A type whose specific kind is not exposed via this interface.
	Type_SChar = C.CXType_SChar
	// A type whose specific kind is not exposed via this interface.
	Type_WChar = C.CXType_WChar
	// A type whose specific kind is not exposed via this interface.
	Type_Short = C.CXType_Short
	// A type whose specific kind is not exposed via this interface.
	Type_Int = C.CXType_Int
	// A type whose specific kind is not exposed via this interface.
	Type_Long = C.CXType_Long
	// A type whose specific kind is not exposed via this interface.
	Type_LongLong = C.CXType_LongLong
	// A type whose specific kind is not exposed via this interface.
	Type_Int128 = C.CXType_Int128
	// A type whose specific kind is not exposed via this interface.
	Type_Float = C.CXType_Float
	// A type whose specific kind is not exposed via this interface.
	Type_Double = C.CXType_Double
	// A type whose specific kind is not exposed via this interface.
	Type_LongDouble = C.CXType_LongDouble
	// A type whose specific kind is not exposed via this interface.
	Type_NullPtr = C.CXType_NullPtr
	// A type whose specific kind is not exposed via this interface.
	Type_Overload = C.CXType_Overload
	// A type whose specific kind is not exposed via this interface.
	Type_Dependent = C.CXType_Dependent
	// A type whose specific kind is not exposed via this interface.
	Type_ObjCId = C.CXType_ObjCId
	// A type whose specific kind is not exposed via this interface.
	Type_ObjCClass = C.CXType_ObjCClass
	// A type whose specific kind is not exposed via this interface.
	Type_ObjCSel = C.CXType_ObjCSel
	// A type whose specific kind is not exposed via this interface.
	Type_Float128 = C.CXType_Float128
	// A type whose specific kind is not exposed via this interface.
	Type_FirstBuiltin = C.CXType_FirstBuiltin
	// A type whose specific kind is not exposed via this interface.
	Type_LastBuiltin = C.CXType_LastBuiltin
	// A type whose specific kind is not exposed via this interface.
	Type_Complex = C.CXType_Complex
	// A type whose specific kind is not exposed via this interface.
	Type_Pointer = C.CXType_Pointer
	// A type whose specific kind is not exposed via this interface.
	Type_BlockPointer = C.CXType_BlockPointer
	// A type whose specific kind is not exposed via this interface.
	Type_LValueReference = C.CXType_LValueReference
	// A type whose specific kind is not exposed via this interface.
	Type_RValueReference = C.CXType_RValueReference
	// A type whose specific kind is not exposed via this interface.
	Type_Record = C.CXType_Record
	// A type whose specific kind is not exposed via this interface.
	Type_Enum = C.CXType_Enum
	// A type whose specific kind is not exposed via this interface.
	Type_Typedef = C.CXType_Typedef
	// A type whose specific kind is not exposed via this interface.
	Type_ObjCInterface = C.CXType_ObjCInterface
	// A type whose specific kind is not exposed via this interface.
	Type_ObjCObjectPointer = C.CXType_ObjCObjectPointer
	// A type whose specific kind is not exposed via this interface.
	Type_FunctionNoProto = C.CXType_FunctionNoProto
	// A type whose specific kind is not exposed via this interface.
	Type_FunctionProto = C.CXType_FunctionProto
	// A type whose specific kind is not exposed via this interface.
	Type_ConstantArray = C.CXType_ConstantArray
	// A type whose specific kind is not exposed via this interface.
	Type_Vector = C.CXType_Vector
	// A type whose specific kind is not exposed via this interface.
	Type_IncompleteArray = C.CXType_IncompleteArray
	// A type whose specific kind is not exposed via this interface.
	Type_VariableArray = C.CXType_VariableArray
	// A type whose specific kind is not exposed via this interface.
	Type_DependentSizedArray = C.CXType_DependentSizedArray
	// A type whose specific kind is not exposed via this interface.
	Type_MemberPointer = C.CXType_MemberPointer
	// A type whose specific kind is not exposed via this interface.
	Type_Auto = C.CXType_Auto
	/*
		Represents a type that was referred to using an elaborated type keyword.

		E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	Type_Elaborated = C.CXType_Elaborated
)

func (tk TypeKind) Spelling() string {
	switch tk {
	case Type_Invalid:
		return "Type=Invalid"
	case Type_Unexposed:
		return "Type=Unexposed"
	case Type_Void:
		return "Type=Void, FirstBuiltin"
	case Type_Bool:
		return "Type=Bool"
	case Type_Char_U:
		return "Type=Char_U"
	case Type_UChar:
		return "Type=UChar"
	case Type_Char16:
		return "Type=Char16"
	case Type_Char32:
		return "Type=Char32"
	case Type_UShort:
		return "Type=UShort"
	case Type_UInt:
		return "Type=UInt"
	case Type_ULong:
		return "Type=ULong"
	case Type_ULongLong:
		return "Type=ULongLong"
	case Type_UInt128:
		return "Type=UInt128"
	case Type_Char_S:
		return "Type=Char_S"
	case Type_SChar:
		return "Type=SChar"
	case Type_WChar:
		return "Type=WChar"
	case Type_Short:
		return "Type=Short"
	case Type_Int:
		return "Type=Int"
	case Type_Long:
		return "Type=Long"
	case Type_LongLong:
		return "Type=LongLong"
	case Type_Int128:
		return "Type=Int128"
	case Type_Float:
		return "Type=Float"
	case Type_Double:
		return "Type=Double"
	case Type_LongDouble:
		return "Type=LongDouble"
	case Type_NullPtr:
		return "Type=NullPtr"
	case Type_Overload:
		return "Type=Overload"
	case Type_Dependent:
		return "Type=Dependent"
	case Type_ObjCId:
		return "Type=ObjCId"
	case Type_ObjCClass:
		return "Type=ObjCClass"
	case Type_ObjCSel:
		return "Type=ObjCSel, LastBuiltin"
	case Type_Float128:
		return "Type=Float128"
	case Type_Complex:
		return "Type=Complex"
	case Type_Pointer:
		return "Type=Pointer"
	case Type_BlockPointer:
		return "Type=BlockPointer"
	case Type_LValueReference:
		return "Type=LValueReference"
	case Type_RValueReference:
		return "Type=RValueReference"
	case Type_Record:
		return "Type=Record"
	case Type_Enum:
		return "Type=Enum"
	case Type_Typedef:
		return "Type=Typedef"
	case Type_ObjCInterface:
		return "Type=ObjCInterface"
	case Type_ObjCObjectPointer:
		return "Type=ObjCObjectPointer"
	case Type_FunctionNoProto:
		return "Type=FunctionNoProto"
	case Type_FunctionProto:
		return "Type=FunctionProto"
	case Type_ConstantArray:
		return "Type=ConstantArray"
	case Type_Vector:
		return "Type=Vector"
	case Type_IncompleteArray:
		return "Type=IncompleteArray"
	case Type_VariableArray:
		return "Type=VariableArray"
	case Type_DependentSizedArray:
		return "Type=DependentSizedArray"
	case Type_MemberPointer:
		return "Type=MemberPointer"
	case Type_Auto:
		return "Type=Auto"
	case Type_Elaborated:
		return "Type=Elaborated"
	}

	return fmt.Sprintf("TypeKind unkown %d", int(tk))
}

func (tk TypeKind) String() string {
	return tk.Spelling()
}
