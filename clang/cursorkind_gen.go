package clang

// #include "./clang-c/Index.h"
// #include "go-clang.h"
import "C"
import "fmt"

// Describes the kind of entity that a cursor refers to.
type CursorKind uint32

const (
	/*
		A declaration whose specific kind is not exposed via this
		interface.

		Unexposed declarations have the same operations as any other kind
		of declaration; one can extract their location information,
		spelling, find their definitions, etc. However, the specific kind
		of the declaration is not reported.
	*/
	Cursor_UnexposedDecl CursorKind = C.CXCursor_UnexposedDecl
	// A C or C++ struct.
	Cursor_StructDecl = C.CXCursor_StructDecl
	// A C or C++ union.
	Cursor_UnionDecl = C.CXCursor_UnionDecl
	// A C++ class.
	Cursor_ClassDecl = C.CXCursor_ClassDecl
	// An enumeration.
	Cursor_EnumDecl = C.CXCursor_EnumDecl
	// A field (in C) or non-static data member (in C++) in a struct, union, or C++ class.
	Cursor_FieldDecl = C.CXCursor_FieldDecl
	// An enumerator constant.
	Cursor_EnumConstantDecl = C.CXCursor_EnumConstantDecl
	// A function.
	Cursor_FunctionDecl = C.CXCursor_FunctionDecl
	// A variable.
	Cursor_VarDecl = C.CXCursor_VarDecl
	// A function or method parameter.
	Cursor_ParmDecl = C.CXCursor_ParmDecl
	// An Objective-C \@interface.
	Cursor_ObjCInterfaceDecl = C.CXCursor_ObjCInterfaceDecl
	// An Objective-C \@interface for a category.
	Cursor_ObjCCategoryDecl = C.CXCursor_ObjCCategoryDecl
	// An Objective-C \@protocol declaration.
	Cursor_ObjCProtocolDecl = C.CXCursor_ObjCProtocolDecl
	// An Objective-C \@property declaration.
	Cursor_ObjCPropertyDecl = C.CXCursor_ObjCPropertyDecl
	// An Objective-C instance variable.
	Cursor_ObjCIvarDecl = C.CXCursor_ObjCIvarDecl
	// An Objective-C instance method.
	Cursor_ObjCInstanceMethodDecl = C.CXCursor_ObjCInstanceMethodDecl
	// An Objective-C class method.
	Cursor_ObjCClassMethodDecl = C.CXCursor_ObjCClassMethodDecl
	// An Objective-C \@implementation.
	Cursor_ObjCImplementationDecl = C.CXCursor_ObjCImplementationDecl
	// An Objective-C \@implementation for a category.
	Cursor_ObjCCategoryImplDecl = C.CXCursor_ObjCCategoryImplDecl
	// A typedef.
	Cursor_TypedefDecl = C.CXCursor_TypedefDecl
	// A C++ class method.
	Cursor_CXXMethod = C.CXCursor_CXXMethod
	// A C++ namespace.
	Cursor_Namespace = C.CXCursor_Namespace
	// A linkage specification, e.g. 'extern "C"'.
	Cursor_LinkageSpec = C.CXCursor_LinkageSpec
	// A C++ constructor.
	Cursor_Constructor = C.CXCursor_Constructor
	// A C++ destructor.
	Cursor_Destructor = C.CXCursor_Destructor
	// A C++ conversion function.
	Cursor_ConversionFunction = C.CXCursor_ConversionFunction
	// A C++ template type parameter.
	Cursor_TemplateTypeParameter = C.CXCursor_TemplateTypeParameter
	// A C++ non-type template parameter.
	Cursor_NonTypeTemplateParameter = C.CXCursor_NonTypeTemplateParameter
	// A C++ template template parameter.
	Cursor_TemplateTemplateParameter = C.CXCursor_TemplateTemplateParameter
	// A C++ function template.
	Cursor_FunctionTemplate = C.CXCursor_FunctionTemplate
	// A C++ class template.
	Cursor_ClassTemplate = C.CXCursor_ClassTemplate
	// A C++ class template partial specialization.
	Cursor_ClassTemplatePartialSpecialization = C.CXCursor_ClassTemplatePartialSpecialization
	// A C++ namespace alias declaration.
	Cursor_NamespaceAlias = C.CXCursor_NamespaceAlias
	// A C++ using directive.
	Cursor_UsingDirective = C.CXCursor_UsingDirective
	// A C++ using declaration.
	Cursor_UsingDeclaration = C.CXCursor_UsingDeclaration
	// A C++ alias declaration
	Cursor_TypeAliasDecl = C.CXCursor_TypeAliasDecl
	// An Objective-C \@synthesize definition.
	Cursor_ObjCSynthesizeDecl = C.CXCursor_ObjCSynthesizeDecl
	// An Objective-C \@dynamic definition.
	Cursor_ObjCDynamicDecl = C.CXCursor_ObjCDynamicDecl
	// An access specifier.
	Cursor_CXXAccessSpecifier = C.CXCursor_CXXAccessSpecifier
	// An access specifier.
	Cursor_FirstDecl = C.CXCursor_FirstDecl
	// An access specifier.
	Cursor_LastDecl = C.CXCursor_LastDecl
	// An access specifier.
	Cursor_FirstRef = C.CXCursor_FirstRef
	// An access specifier.
	Cursor_ObjCSuperClassRef = C.CXCursor_ObjCSuperClassRef
	// An access specifier.
	Cursor_ObjCProtocolRef = C.CXCursor_ObjCProtocolRef
	// An access specifier.
	Cursor_ObjCClassRef = C.CXCursor_ObjCClassRef
	/*
		A reference to a type declaration.

		A type reference occurs anywhere where a type is named but not
		declared. For example, given:

		\code
		typedef unsigned size_type;
		size_type size;
		\endcode

		The typedef is a declaration of size_type (CXCursor_TypedefDecl),
		while the type of the variable "size" is referenced. The cursor
		referenced by the type of size is the typedef for size_type.
	*/
	Cursor_TypeRef = C.CXCursor_TypeRef
	/*
		A reference to a type declaration.

		A type reference occurs anywhere where a type is named but not
		declared. For example, given:

		\code
		typedef unsigned size_type;
		size_type size;
		\endcode

		The typedef is a declaration of size_type (CXCursor_TypedefDecl),
		while the type of the variable "size" is referenced. The cursor
		referenced by the type of size is the typedef for size_type.
	*/
	Cursor_CXXBaseSpecifier = C.CXCursor_CXXBaseSpecifier
	// A reference to a class template, function template, template template parameter, or class template partial specialization.
	Cursor_TemplateRef = C.CXCursor_TemplateRef
	// A reference to a namespace or namespace alias.
	Cursor_NamespaceRef = C.CXCursor_NamespaceRef
	// A reference to a member of a struct, union, or class that occurs in some non-expression context, e.g., a designated initializer.
	Cursor_MemberRef = C.CXCursor_MemberRef
	/*
		A reference to a labeled statement.

		This cursor kind is used to describe the jump to "start_over" in the
		goto statement in the following example:

		\code
		start_over:
		++counter;

		goto start_over;
		\endcode

		A label reference cursor refers to a label statement.
	*/
	Cursor_LabelRef = C.CXCursor_LabelRef
	/*
		A reference to a set of overloaded functions or function templates
		that has not yet been resolved to a specific function or function template.

		An overloaded declaration reference cursor occurs in C++ templates where
		a dependent name refers to a function. For example:

		\code
		template<typename T> void swap(T&, T&);

		struct X { ... };
		void swap(X&, X&);

		template<typename T>
		void reverse(T* first, T* last) {
		while (first < last - 1) {
		swap(*first, *--last);
		++first;
		}
		}

		struct Y { };
		void swap(Y&, Y&);
		\endcode

		Here, the identifier "swap" is associated with an overloaded declaration
		reference. In the template definition, "swap" refers to either of the two
		"swap" functions declared above, so both results will be available. At
		instantiation time, "swap" may also refer to other functions found via
		argument-dependent lookup (e.g., the "swap" function at the end of the
		example).

		The functions clang_getNumOverloadedDecls() and
		clang_getOverloadedDecl() can be used to retrieve the definitions
		referenced by this cursor.
	*/
	Cursor_OverloadedDeclRef = C.CXCursor_OverloadedDeclRef
	// A reference to a variable that occurs in some non-expression context, e.g., a C++ lambda capture list.
	Cursor_VariableRef = C.CXCursor_VariableRef
	// A reference to a variable that occurs in some non-expression context, e.g., a C++ lambda capture list.
	Cursor_LastRef = C.CXCursor_LastRef
	// A reference to a variable that occurs in some non-expression context, e.g., a C++ lambda capture list.
	Cursor_FirstInvalid = C.CXCursor_FirstInvalid
	// A reference to a variable that occurs in some non-expression context, e.g., a C++ lambda capture list.
	Cursor_InvalidFile = C.CXCursor_InvalidFile
	// A reference to a variable that occurs in some non-expression context, e.g., a C++ lambda capture list.
	Cursor_NoDeclFound = C.CXCursor_NoDeclFound
	// A reference to a variable that occurs in some non-expression context, e.g., a C++ lambda capture list.
	Cursor_NotImplemented = C.CXCursor_NotImplemented
	// A reference to a variable that occurs in some non-expression context, e.g., a C++ lambda capture list.
	Cursor_InvalidCode = C.CXCursor_InvalidCode
	// A reference to a variable that occurs in some non-expression context, e.g., a C++ lambda capture list.
	Cursor_LastInvalid = C.CXCursor_LastInvalid
	// A reference to a variable that occurs in some non-expression context, e.g., a C++ lambda capture list.
	Cursor_FirstExpr = C.CXCursor_FirstExpr
	/*
		An expression whose specific kind is not exposed via this
		interface.

		Unexposed expressions have the same operations as any other kind
		of expression; one can extract their location information,
		spelling, children, etc. However, the specific kind of the
		expression is not reported.
	*/
	Cursor_UnexposedExpr = C.CXCursor_UnexposedExpr
	// An expression that refers to some value declaration, such as a function, variable, or enumerator.
	Cursor_DeclRefExpr = C.CXCursor_DeclRefExpr
	// An expression that refers to a member of a struct, union, class, Objective-C class, etc.
	Cursor_MemberRefExpr = C.CXCursor_MemberRefExpr
	// An expression that calls a function.
	Cursor_CallExpr = C.CXCursor_CallExpr
	// An expression that sends a message to an Objective-C  object or class.
	Cursor_ObjCMessageExpr = C.CXCursor_ObjCMessageExpr
	// An expression that represents a block literal.
	Cursor_BlockExpr = C.CXCursor_BlockExpr
	// An integer literal.
	Cursor_IntegerLiteral = C.CXCursor_IntegerLiteral
	// A floating point number literal.
	Cursor_FloatingLiteral = C.CXCursor_FloatingLiteral
	// An imaginary number literal.
	Cursor_ImaginaryLiteral = C.CXCursor_ImaginaryLiteral
	// A string literal.
	Cursor_StringLiteral = C.CXCursor_StringLiteral
	// A character literal.
	Cursor_CharacterLiteral = C.CXCursor_CharacterLiteral
	/*
		A parenthesized expression, e.g. "(1)".

		This AST node is only formed if full location information is requested.
	*/
	Cursor_ParenExpr = C.CXCursor_ParenExpr
	// This represents the unary-expression's (except sizeof and alignof).
	Cursor_UnaryOperator = C.CXCursor_UnaryOperator
	// [C99 6.5.2.1] Array Subscripting.
	Cursor_ArraySubscriptExpr = C.CXCursor_ArraySubscriptExpr
	// A builtin binary operation expression such as "x + y" or "x <= y".
	Cursor_BinaryOperator = C.CXCursor_BinaryOperator
	// Compound assignment such as "+=".
	Cursor_CompoundAssignOperator = C.CXCursor_CompoundAssignOperator
	// The ?: ternary operator.
	Cursor_ConditionalOperator = C.CXCursor_ConditionalOperator
	/*
		An explicit cast in C (C99 6.5.4) or a C-style cast in C++
		(C++ [expr.cast]), which uses the syntax (Type)expr.

		For example: (int)f.
	*/
	Cursor_CStyleCastExpr = C.CXCursor_CStyleCastExpr
	// [C99 6.5.2.5]
	Cursor_CompoundLiteralExpr = C.CXCursor_CompoundLiteralExpr
	// Describes an C or C++ initializer list.
	Cursor_InitListExpr = C.CXCursor_InitListExpr
	// The GNU address of label extension, representing &&label.
	Cursor_AddrLabelExpr = C.CXCursor_AddrLabelExpr
	// This is the GNU Statement Expression extension: ({int X=4; X;})
	Cursor_StmtExpr = C.CXCursor_StmtExpr
	// Represents a C11 generic selection.
	Cursor_GenericSelectionExpr = C.CXCursor_GenericSelectionExpr
	/*
		Implements the GNU __null extension, which is a name for a null
		pointer constant that has integral type (e.g., int or long) and is the same
		size and alignment as a pointer.

		The __null extension is typically only used by system headers, which define
		NULL as __null in C++ rather than using 0 (which is an integer that may not
		match the size of a pointer).
	*/
	Cursor_GNUNullExpr = C.CXCursor_GNUNullExpr
	// C++'s static_cast<> expression.
	Cursor_CXXStaticCastExpr = C.CXCursor_CXXStaticCastExpr
	// C++'s dynamic_cast<> expression.
	Cursor_CXXDynamicCastExpr = C.CXCursor_CXXDynamicCastExpr
	// C++'s reinterpret_cast<> expression.
	Cursor_CXXReinterpretCastExpr = C.CXCursor_CXXReinterpretCastExpr
	// C++'s const_cast<> expression.
	Cursor_CXXConstCastExpr = C.CXCursor_CXXConstCastExpr
	/*
		Represents an explicit C++ type conversion that uses "functional"
		notion (C++ [expr.type.conv]).

		Example:
		\code
		x = int(0.5);
		\endcode
	*/
	Cursor_CXXFunctionalCastExpr = C.CXCursor_CXXFunctionalCastExpr
	// A C++ typeid expression (C++ [expr.typeid]).
	Cursor_CXXTypeidExpr = C.CXCursor_CXXTypeidExpr
	// [C++ 2.13.5] C++ Boolean Literal.
	Cursor_CXXBoolLiteralExpr = C.CXCursor_CXXBoolLiteralExpr
	// [C++0x 2.14.7] C++ Pointer Literal.
	Cursor_CXXNullPtrLiteralExpr = C.CXCursor_CXXNullPtrLiteralExpr
	// Represents the "this" expression in C++
	Cursor_CXXThisExpr = C.CXCursor_CXXThisExpr
	/*
		[C++ 15] C++ Throw Expression.

		This handles 'throw' and 'throw' assignment-expression. When
		assignment-expression isn't present, Op will be null.
	*/
	Cursor_CXXThrowExpr = C.CXCursor_CXXThrowExpr
	// A new expression for memory allocation and constructor calls, e.g: "new CXXNewExpr(foo)".
	Cursor_CXXNewExpr = C.CXCursor_CXXNewExpr
	// A delete expression for memory deallocation and destructor calls, e.g. "delete[] pArray".
	Cursor_CXXDeleteExpr = C.CXCursor_CXXDeleteExpr
	// A unary expression. (noexcept, sizeof, or other traits)
	Cursor_UnaryExpr = C.CXCursor_UnaryExpr
	// An Objective-C string literal i.e. @"foo".
	Cursor_ObjCStringLiteral = C.CXCursor_ObjCStringLiteral
	// An Objective-C \@encode expression.
	Cursor_ObjCEncodeExpr = C.CXCursor_ObjCEncodeExpr
	// An Objective-C \@selector expression.
	Cursor_ObjCSelectorExpr = C.CXCursor_ObjCSelectorExpr
	// An Objective-C \@protocol expression.
	Cursor_ObjCProtocolExpr = C.CXCursor_ObjCProtocolExpr
	/*
		An Objective-C "bridged" cast expression, which casts between
		Objective-C pointers and C pointers, transferring ownership in the process.

		\code
		NSString *str = (__bridge_transfer NSString *)CFCreateString();
		\endcode
	*/
	Cursor_ObjCBridgedCastExpr = C.CXCursor_ObjCBridgedCastExpr
	/*
		Represents a C++0x pack expansion that produces a sequence of
		expressions.

		A pack expansion expression contains a pattern (which itself is an
		expression) followed by an ellipsis. For example:

		\code
		template<typename F, typename ...Types>
		void forward(F f, Types &&...args) {
		f(static_cast<Types&&>(args)...);
		}
		\endcode
	*/
	Cursor_PackExpansionExpr = C.CXCursor_PackExpansionExpr
	/*
		Represents an expression that computes the length of a parameter
		pack.

		\code
		template<typename ...Types>
		struct count {
		static const unsigned value = sizeof...(Types);
		};
		\endcode
	*/
	Cursor_SizeOfPackExpr = C.CXCursor_SizeOfPackExpr
	Cursor_LambdaExpr     = C.CXCursor_LambdaExpr
	// Objective-c Boolean Literal.
	Cursor_ObjCBoolLiteralExpr = C.CXCursor_ObjCBoolLiteralExpr
	// Represents the "self" expression in an Objective-C method.
	Cursor_ObjCSelfExpr = C.CXCursor_ObjCSelfExpr
	// OpenMP 4.0 [2.4, Array Section].
	Cursor_OMPArraySectionExpr = C.CXCursor_OMPArraySectionExpr
	// Represents an @available(...) check.
	Cursor_ObjCAvailabilityCheckExpr = C.CXCursor_ObjCAvailabilityCheckExpr
	// Represents an @available(...) check.
	Cursor_LastExpr = C.CXCursor_LastExpr
	// Represents an @available(...) check.
	Cursor_FirstStmt = C.CXCursor_FirstStmt
	/*
		A statement whose specific kind is not exposed via this
		interface.

		Unexposed statements have the same operations as any other kind of
		statement; one can extract their location information, spelling,
		children, etc. However, the specific kind of the statement is not
		reported.
	*/
	Cursor_UnexposedStmt = C.CXCursor_UnexposedStmt
	/*
		A labelled statement in a function.

		This cursor kind is used to describe the "start_over:" label statement in
		the following example:

		\code
		start_over:
		++counter;
		\endcode
	*/
	Cursor_LabelStmt = C.CXCursor_LabelStmt
	/*
		A group of statements like { stmt stmt }.

		This cursor kind is used to describe compound statements, e.g. function
		bodies.
	*/
	Cursor_CompoundStmt = C.CXCursor_CompoundStmt
	// A case statement.
	Cursor_CaseStmt = C.CXCursor_CaseStmt
	// A default statement.
	Cursor_DefaultStmt = C.CXCursor_DefaultStmt
	// An if statement
	Cursor_IfStmt = C.CXCursor_IfStmt
	// A switch statement.
	Cursor_SwitchStmt = C.CXCursor_SwitchStmt
	// A while statement.
	Cursor_WhileStmt = C.CXCursor_WhileStmt
	// A do statement.
	Cursor_DoStmt = C.CXCursor_DoStmt
	// A for statement.
	Cursor_ForStmt = C.CXCursor_ForStmt
	// A goto statement.
	Cursor_GotoStmt = C.CXCursor_GotoStmt
	// An indirect goto statement.
	Cursor_IndirectGotoStmt = C.CXCursor_IndirectGotoStmt
	// A continue statement.
	Cursor_ContinueStmt = C.CXCursor_ContinueStmt
	// A break statement.
	Cursor_BreakStmt = C.CXCursor_BreakStmt
	// A return statement.
	Cursor_ReturnStmt = C.CXCursor_ReturnStmt
	// A GCC inline assembly statement extension.
	Cursor_GCCAsmStmt = C.CXCursor_GCCAsmStmt
	// A GCC inline assembly statement extension.
	Cursor_AsmStmt = C.CXCursor_AsmStmt
	// Objective-C's overall \@try-\@catch-\@finally statement.
	Cursor_ObjCAtTryStmt = C.CXCursor_ObjCAtTryStmt
	// Objective-C's \@catch statement.
	Cursor_ObjCAtCatchStmt = C.CXCursor_ObjCAtCatchStmt
	// Objective-C's \@finally statement.
	Cursor_ObjCAtFinallyStmt = C.CXCursor_ObjCAtFinallyStmt
	// Objective-C's \@throw statement.
	Cursor_ObjCAtThrowStmt = C.CXCursor_ObjCAtThrowStmt
	// Objective-C's \@synchronized statement.
	Cursor_ObjCAtSynchronizedStmt = C.CXCursor_ObjCAtSynchronizedStmt
	// Objective-C's autorelease pool statement.
	Cursor_ObjCAutoreleasePoolStmt = C.CXCursor_ObjCAutoreleasePoolStmt
	// Objective-C's collection statement.
	Cursor_ObjCForCollectionStmt = C.CXCursor_ObjCForCollectionStmt
	// C++'s catch statement.
	Cursor_CXXCatchStmt = C.CXCursor_CXXCatchStmt
	// C++'s try statement.
	Cursor_CXXTryStmt = C.CXCursor_CXXTryStmt
	// C++'s for (* : *) statement.
	Cursor_CXXForRangeStmt = C.CXCursor_CXXForRangeStmt
	// Windows Structured Exception Handling's try statement.
	Cursor_SEHTryStmt = C.CXCursor_SEHTryStmt
	// Windows Structured Exception Handling's except statement.
	Cursor_SEHExceptStmt = C.CXCursor_SEHExceptStmt
	// Windows Structured Exception Handling's finally statement.
	Cursor_SEHFinallyStmt = C.CXCursor_SEHFinallyStmt
	// A MS inline assembly statement extension.
	Cursor_MSAsmStmt = C.CXCursor_MSAsmStmt
	/*
		The null statement ";": C99 6.8.3p3.

		This cursor kind is used to describe the null statement.
	*/
	Cursor_NullStmt = C.CXCursor_NullStmt
	// Adaptor class for mixing declarations with statements and expressions.
	Cursor_DeclStmt = C.CXCursor_DeclStmt
	// OpenMP parallel directive.
	Cursor_OMPParallelDirective = C.CXCursor_OMPParallelDirective
	// OpenMP SIMD directive.
	Cursor_OMPSimdDirective = C.CXCursor_OMPSimdDirective
	// OpenMP for directive.
	Cursor_OMPForDirective = C.CXCursor_OMPForDirective
	// OpenMP sections directive.
	Cursor_OMPSectionsDirective = C.CXCursor_OMPSectionsDirective
	// OpenMP section directive.
	Cursor_OMPSectionDirective = C.CXCursor_OMPSectionDirective
	// OpenMP single directive.
	Cursor_OMPSingleDirective = C.CXCursor_OMPSingleDirective
	// OpenMP parallel for directive.
	Cursor_OMPParallelForDirective = C.CXCursor_OMPParallelForDirective
	// OpenMP parallel sections directive.
	Cursor_OMPParallelSectionsDirective = C.CXCursor_OMPParallelSectionsDirective
	// OpenMP task directive.
	Cursor_OMPTaskDirective = C.CXCursor_OMPTaskDirective
	// OpenMP master directive.
	Cursor_OMPMasterDirective = C.CXCursor_OMPMasterDirective
	// OpenMP critical directive.
	Cursor_OMPCriticalDirective = C.CXCursor_OMPCriticalDirective
	// OpenMP taskyield directive.
	Cursor_OMPTaskyieldDirective = C.CXCursor_OMPTaskyieldDirective
	// OpenMP barrier directive.
	Cursor_OMPBarrierDirective = C.CXCursor_OMPBarrierDirective
	// OpenMP taskwait directive.
	Cursor_OMPTaskwaitDirective = C.CXCursor_OMPTaskwaitDirective
	// OpenMP flush directive.
	Cursor_OMPFlushDirective = C.CXCursor_OMPFlushDirective
	// Windows Structured Exception Handling's leave statement.
	Cursor_SEHLeaveStmt = C.CXCursor_SEHLeaveStmt
	// OpenMP ordered directive.
	Cursor_OMPOrderedDirective = C.CXCursor_OMPOrderedDirective
	// OpenMP atomic directive.
	Cursor_OMPAtomicDirective = C.CXCursor_OMPAtomicDirective
	// OpenMP for SIMD directive.
	Cursor_OMPForSimdDirective = C.CXCursor_OMPForSimdDirective
	// OpenMP parallel for SIMD directive.
	Cursor_OMPParallelForSimdDirective = C.CXCursor_OMPParallelForSimdDirective
	// OpenMP target directive.
	Cursor_OMPTargetDirective = C.CXCursor_OMPTargetDirective
	// OpenMP teams directive.
	Cursor_OMPTeamsDirective = C.CXCursor_OMPTeamsDirective
	// OpenMP taskgroup directive.
	Cursor_OMPTaskgroupDirective = C.CXCursor_OMPTaskgroupDirective
	// OpenMP cancellation point directive.
	Cursor_OMPCancellationPointDirective = C.CXCursor_OMPCancellationPointDirective
	// OpenMP cancel directive.
	Cursor_OMPCancelDirective = C.CXCursor_OMPCancelDirective
	// OpenMP target data directive.
	Cursor_OMPTargetDataDirective = C.CXCursor_OMPTargetDataDirective
	// OpenMP taskloop directive.
	Cursor_OMPTaskLoopDirective = C.CXCursor_OMPTaskLoopDirective
	// OpenMP taskloop simd directive.
	Cursor_OMPTaskLoopSimdDirective = C.CXCursor_OMPTaskLoopSimdDirective
	// OpenMP distribute directive.
	Cursor_OMPDistributeDirective = C.CXCursor_OMPDistributeDirective
	// OpenMP target enter data directive.
	Cursor_OMPTargetEnterDataDirective = C.CXCursor_OMPTargetEnterDataDirective
	// OpenMP target exit data directive.
	Cursor_OMPTargetExitDataDirective = C.CXCursor_OMPTargetExitDataDirective
	// OpenMP target parallel directive.
	Cursor_OMPTargetParallelDirective = C.CXCursor_OMPTargetParallelDirective
	// OpenMP target parallel for directive.
	Cursor_OMPTargetParallelForDirective = C.CXCursor_OMPTargetParallelForDirective
	// OpenMP target update directive.
	Cursor_OMPTargetUpdateDirective = C.CXCursor_OMPTargetUpdateDirective
	// OpenMP distribute parallel for directive.
	Cursor_OMPDistributeParallelForDirective = C.CXCursor_OMPDistributeParallelForDirective
	// OpenMP distribute parallel for simd directive.
	Cursor_OMPDistributeParallelForSimdDirective = C.CXCursor_OMPDistributeParallelForSimdDirective
	// OpenMP distribute simd directive.
	Cursor_OMPDistributeSimdDirective = C.CXCursor_OMPDistributeSimdDirective
	// OpenMP target parallel for simd directive.
	Cursor_OMPTargetParallelForSimdDirective = C.CXCursor_OMPTargetParallelForSimdDirective
	// OpenMP target parallel for simd directive.
	Cursor_LastStmt = C.CXCursor_LastStmt
	/*
		Cursor that represents the translation unit itself.

		The translation unit cursor exists primarily to act as the root
		cursor for traversing the contents of a translation unit.
	*/
	Cursor_TranslationUnit = C.CXCursor_TranslationUnit
	/*
		Cursor that represents the translation unit itself.

		The translation unit cursor exists primarily to act as the root
		cursor for traversing the contents of a translation unit.
	*/
	Cursor_FirstAttr = C.CXCursor_FirstAttr
	// An attribute whose specific kind is not exposed via this interface.
	Cursor_UnexposedAttr = C.CXCursor_UnexposedAttr
	// An attribute whose specific kind is not exposed via this interface.
	Cursor_IBActionAttr = C.CXCursor_IBActionAttr
	// An attribute whose specific kind is not exposed via this interface.
	Cursor_IBOutletAttr = C.CXCursor_IBOutletAttr
	// An attribute whose specific kind is not exposed via this interface.
	Cursor_IBOutletCollectionAttr = C.CXCursor_IBOutletCollectionAttr
	// An attribute whose specific kind is not exposed via this interface.
	Cursor_CXXFinalAttr = C.CXCursor_CXXFinalAttr
	// An attribute whose specific kind is not exposed via this interface.
	Cursor_CXXOverrideAttr = C.CXCursor_CXXOverrideAttr
	// An attribute whose specific kind is not exposed via this interface.
	Cursor_AnnotateAttr = C.CXCursor_AnnotateAttr
	// An attribute whose specific kind is not exposed via this interface.
	Cursor_AsmLabelAttr = C.CXCursor_AsmLabelAttr
	// An attribute whose specific kind is not exposed via this interface.
	Cursor_PackedAttr = C.CXCursor_PackedAttr
	// An attribute whose specific kind is not exposed via this interface.
	Cursor_PureAttr = C.CXCursor_PureAttr
	// An attribute whose specific kind is not exposed via this interface.
	Cursor_ConstAttr = C.CXCursor_ConstAttr
	// An attribute whose specific kind is not exposed via this interface.
	Cursor_NoDuplicateAttr = C.CXCursor_NoDuplicateAttr
	// An attribute whose specific kind is not exposed via this interface.
	Cursor_CUDAConstantAttr = C.CXCursor_CUDAConstantAttr
	// An attribute whose specific kind is not exposed via this interface.
	Cursor_CUDADeviceAttr = C.CXCursor_CUDADeviceAttr
	// An attribute whose specific kind is not exposed via this interface.
	Cursor_CUDAGlobalAttr = C.CXCursor_CUDAGlobalAttr
	// An attribute whose specific kind is not exposed via this interface.
	Cursor_CUDAHostAttr = C.CXCursor_CUDAHostAttr
	// An attribute whose specific kind is not exposed via this interface.
	Cursor_CUDASharedAttr = C.CXCursor_CUDASharedAttr
	// An attribute whose specific kind is not exposed via this interface.
	Cursor_VisibilityAttr = C.CXCursor_VisibilityAttr
	// An attribute whose specific kind is not exposed via this interface.
	Cursor_DLLExport = C.CXCursor_DLLExport
	// An attribute whose specific kind is not exposed via this interface.
	Cursor_DLLImport = C.CXCursor_DLLImport
	// An attribute whose specific kind is not exposed via this interface.
	Cursor_LastAttr = C.CXCursor_LastAttr
	// An attribute whose specific kind is not exposed via this interface.
	Cursor_PreprocessingDirective = C.CXCursor_PreprocessingDirective
	// An attribute whose specific kind is not exposed via this interface.
	Cursor_MacroDefinition = C.CXCursor_MacroDefinition
	// An attribute whose specific kind is not exposed via this interface.
	Cursor_MacroExpansion = C.CXCursor_MacroExpansion
	// An attribute whose specific kind is not exposed via this interface.
	Cursor_MacroInstantiation = C.CXCursor_MacroInstantiation
	// An attribute whose specific kind is not exposed via this interface.
	Cursor_InclusionDirective = C.CXCursor_InclusionDirective
	// An attribute whose specific kind is not exposed via this interface.
	Cursor_FirstPreprocessing = C.CXCursor_FirstPreprocessing
	// An attribute whose specific kind is not exposed via this interface.
	Cursor_LastPreprocessing = C.CXCursor_LastPreprocessing
	// A module import declaration.
	Cursor_ModuleImportDecl = C.CXCursor_ModuleImportDecl
	// A module import declaration.
	Cursor_TypeAliasTemplateDecl = C.CXCursor_TypeAliasTemplateDecl
	// A static_assert or _Static_assert node
	Cursor_StaticAssert = C.CXCursor_StaticAssert
	// A static_assert or _Static_assert node
	Cursor_FirstExtraDecl = C.CXCursor_FirstExtraDecl
	// A static_assert or _Static_assert node
	Cursor_LastExtraDecl = C.CXCursor_LastExtraDecl
	// A code completion overload candidate.
	Cursor_OverloadCandidate = C.CXCursor_OverloadCandidate
)

func (ck CursorKind) Spelling() string {
	switch ck {
	case Cursor_UnexposedDecl:
		return "Cursor=UnexposedDecl, FirstDecl"
	case Cursor_StructDecl:
		return "Cursor=StructDecl"
	case Cursor_UnionDecl:
		return "Cursor=UnionDecl"
	case Cursor_ClassDecl:
		return "Cursor=ClassDecl"
	case Cursor_EnumDecl:
		return "Cursor=EnumDecl"
	case Cursor_FieldDecl:
		return "Cursor=FieldDecl"
	case Cursor_EnumConstantDecl:
		return "Cursor=EnumConstantDecl"
	case Cursor_FunctionDecl:
		return "Cursor=FunctionDecl"
	case Cursor_VarDecl:
		return "Cursor=VarDecl"
	case Cursor_ParmDecl:
		return "Cursor=ParmDecl"
	case Cursor_ObjCInterfaceDecl:
		return "Cursor=ObjCInterfaceDecl"
	case Cursor_ObjCCategoryDecl:
		return "Cursor=ObjCCategoryDecl"
	case Cursor_ObjCProtocolDecl:
		return "Cursor=ObjCProtocolDecl"
	case Cursor_ObjCPropertyDecl:
		return "Cursor=ObjCPropertyDecl"
	case Cursor_ObjCIvarDecl:
		return "Cursor=ObjCIvarDecl"
	case Cursor_ObjCInstanceMethodDecl:
		return "Cursor=ObjCInstanceMethodDecl"
	case Cursor_ObjCClassMethodDecl:
		return "Cursor=ObjCClassMethodDecl"
	case Cursor_ObjCImplementationDecl:
		return "Cursor=ObjCImplementationDecl"
	case Cursor_ObjCCategoryImplDecl:
		return "Cursor=ObjCCategoryImplDecl"
	case Cursor_TypedefDecl:
		return "Cursor=TypedefDecl"
	case Cursor_CXXMethod:
		return "Cursor=CXXMethod"
	case Cursor_Namespace:
		return "Cursor=Namespace"
	case Cursor_LinkageSpec:
		return "Cursor=LinkageSpec"
	case Cursor_Constructor:
		return "Cursor=Constructor"
	case Cursor_Destructor:
		return "Cursor=Destructor"
	case Cursor_ConversionFunction:
		return "Cursor=ConversionFunction"
	case Cursor_TemplateTypeParameter:
		return "Cursor=TemplateTypeParameter"
	case Cursor_NonTypeTemplateParameter:
		return "Cursor=NonTypeTemplateParameter"
	case Cursor_TemplateTemplateParameter:
		return "Cursor=TemplateTemplateParameter"
	case Cursor_FunctionTemplate:
		return "Cursor=FunctionTemplate"
	case Cursor_ClassTemplate:
		return "Cursor=ClassTemplate"
	case Cursor_ClassTemplatePartialSpecialization:
		return "Cursor=ClassTemplatePartialSpecialization"
	case Cursor_NamespaceAlias:
		return "Cursor=NamespaceAlias"
	case Cursor_UsingDirective:
		return "Cursor=UsingDirective"
	case Cursor_UsingDeclaration:
		return "Cursor=UsingDeclaration"
	case Cursor_TypeAliasDecl:
		return "Cursor=TypeAliasDecl"
	case Cursor_ObjCSynthesizeDecl:
		return "Cursor=ObjCSynthesizeDecl"
	case Cursor_ObjCDynamicDecl:
		return "Cursor=ObjCDynamicDecl"
	case Cursor_CXXAccessSpecifier:
		return "Cursor=CXXAccessSpecifier, LastDecl"
	case Cursor_FirstRef:
		return "Cursor=FirstRef, ObjCSuperClassRef"
	case Cursor_ObjCProtocolRef:
		return "Cursor=ObjCProtocolRef"
	case Cursor_ObjCClassRef:
		return "Cursor=ObjCClassRef"
	case Cursor_TypeRef:
		return "Cursor=TypeRef"
	case Cursor_CXXBaseSpecifier:
		return "Cursor=CXXBaseSpecifier"
	case Cursor_TemplateRef:
		return "Cursor=TemplateRef"
	case Cursor_NamespaceRef:
		return "Cursor=NamespaceRef"
	case Cursor_MemberRef:
		return "Cursor=MemberRef"
	case Cursor_LabelRef:
		return "Cursor=LabelRef"
	case Cursor_OverloadedDeclRef:
		return "Cursor=OverloadedDeclRef"
	case Cursor_VariableRef:
		return "Cursor=VariableRef, LastRef"
	case Cursor_FirstInvalid:
		return "Cursor=FirstInvalid, InvalidFile"
	case Cursor_NoDeclFound:
		return "Cursor=NoDeclFound"
	case Cursor_NotImplemented:
		return "Cursor=NotImplemented"
	case Cursor_InvalidCode:
		return "Cursor=InvalidCode, LastInvalid"
	case Cursor_FirstExpr:
		return "Cursor=FirstExpr, UnexposedExpr"
	case Cursor_DeclRefExpr:
		return "Cursor=DeclRefExpr"
	case Cursor_MemberRefExpr:
		return "Cursor=MemberRefExpr"
	case Cursor_CallExpr:
		return "Cursor=CallExpr"
	case Cursor_ObjCMessageExpr:
		return "Cursor=ObjCMessageExpr"
	case Cursor_BlockExpr:
		return "Cursor=BlockExpr"
	case Cursor_IntegerLiteral:
		return "Cursor=IntegerLiteral"
	case Cursor_FloatingLiteral:
		return "Cursor=FloatingLiteral"
	case Cursor_ImaginaryLiteral:
		return "Cursor=ImaginaryLiteral"
	case Cursor_StringLiteral:
		return "Cursor=StringLiteral"
	case Cursor_CharacterLiteral:
		return "Cursor=CharacterLiteral"
	case Cursor_ParenExpr:
		return "Cursor=ParenExpr"
	case Cursor_UnaryOperator:
		return "Cursor=UnaryOperator"
	case Cursor_ArraySubscriptExpr:
		return "Cursor=ArraySubscriptExpr"
	case Cursor_BinaryOperator:
		return "Cursor=BinaryOperator"
	case Cursor_CompoundAssignOperator:
		return "Cursor=CompoundAssignOperator"
	case Cursor_ConditionalOperator:
		return "Cursor=ConditionalOperator"
	case Cursor_CStyleCastExpr:
		return "Cursor=CStyleCastExpr"
	case Cursor_CompoundLiteralExpr:
		return "Cursor=CompoundLiteralExpr"
	case Cursor_InitListExpr:
		return "Cursor=InitListExpr"
	case Cursor_AddrLabelExpr:
		return "Cursor=AddrLabelExpr"
	case Cursor_StmtExpr:
		return "Cursor=StmtExpr"
	case Cursor_GenericSelectionExpr:
		return "Cursor=GenericSelectionExpr"
	case Cursor_GNUNullExpr:
		return "Cursor=GNUNullExpr"
	case Cursor_CXXStaticCastExpr:
		return "Cursor=CXXStaticCastExpr"
	case Cursor_CXXDynamicCastExpr:
		return "Cursor=CXXDynamicCastExpr"
	case Cursor_CXXReinterpretCastExpr:
		return "Cursor=CXXReinterpretCastExpr"
	case Cursor_CXXConstCastExpr:
		return "Cursor=CXXConstCastExpr"
	case Cursor_CXXFunctionalCastExpr:
		return "Cursor=CXXFunctionalCastExpr"
	case Cursor_CXXTypeidExpr:
		return "Cursor=CXXTypeidExpr"
	case Cursor_CXXBoolLiteralExpr:
		return "Cursor=CXXBoolLiteralExpr"
	case Cursor_CXXNullPtrLiteralExpr:
		return "Cursor=CXXNullPtrLiteralExpr"
	case Cursor_CXXThisExpr:
		return "Cursor=CXXThisExpr"
	case Cursor_CXXThrowExpr:
		return "Cursor=CXXThrowExpr"
	case Cursor_CXXNewExpr:
		return "Cursor=CXXNewExpr"
	case Cursor_CXXDeleteExpr:
		return "Cursor=CXXDeleteExpr"
	case Cursor_UnaryExpr:
		return "Cursor=UnaryExpr"
	case Cursor_ObjCStringLiteral:
		return "Cursor=ObjCStringLiteral"
	case Cursor_ObjCEncodeExpr:
		return "Cursor=ObjCEncodeExpr"
	case Cursor_ObjCSelectorExpr:
		return "Cursor=ObjCSelectorExpr"
	case Cursor_ObjCProtocolExpr:
		return "Cursor=ObjCProtocolExpr"
	case Cursor_ObjCBridgedCastExpr:
		return "Cursor=ObjCBridgedCastExpr"
	case Cursor_PackExpansionExpr:
		return "Cursor=PackExpansionExpr"
	case Cursor_SizeOfPackExpr:
		return "Cursor=SizeOfPackExpr"
	case Cursor_LambdaExpr:
		return "Cursor=LambdaExpr"
	case Cursor_ObjCBoolLiteralExpr:
		return "Cursor=ObjCBoolLiteralExpr"
	case Cursor_ObjCSelfExpr:
		return "Cursor=ObjCSelfExpr"
	case Cursor_OMPArraySectionExpr:
		return "Cursor=OMPArraySectionExpr"
	case Cursor_ObjCAvailabilityCheckExpr:
		return "Cursor=ObjCAvailabilityCheckExpr, LastExpr"
	case Cursor_FirstStmt:
		return "Cursor=FirstStmt, UnexposedStmt"
	case Cursor_LabelStmt:
		return "Cursor=LabelStmt"
	case Cursor_CompoundStmt:
		return "Cursor=CompoundStmt"
	case Cursor_CaseStmt:
		return "Cursor=CaseStmt"
	case Cursor_DefaultStmt:
		return "Cursor=DefaultStmt"
	case Cursor_IfStmt:
		return "Cursor=IfStmt"
	case Cursor_SwitchStmt:
		return "Cursor=SwitchStmt"
	case Cursor_WhileStmt:
		return "Cursor=WhileStmt"
	case Cursor_DoStmt:
		return "Cursor=DoStmt"
	case Cursor_ForStmt:
		return "Cursor=ForStmt"
	case Cursor_GotoStmt:
		return "Cursor=GotoStmt"
	case Cursor_IndirectGotoStmt:
		return "Cursor=IndirectGotoStmt"
	case Cursor_ContinueStmt:
		return "Cursor=ContinueStmt"
	case Cursor_BreakStmt:
		return "Cursor=BreakStmt"
	case Cursor_ReturnStmt:
		return "Cursor=ReturnStmt"
	case Cursor_GCCAsmStmt:
		return "Cursor=GCCAsmStmt, AsmStmt"
	case Cursor_ObjCAtTryStmt:
		return "Cursor=ObjCAtTryStmt"
	case Cursor_ObjCAtCatchStmt:
		return "Cursor=ObjCAtCatchStmt"
	case Cursor_ObjCAtFinallyStmt:
		return "Cursor=ObjCAtFinallyStmt"
	case Cursor_ObjCAtThrowStmt:
		return "Cursor=ObjCAtThrowStmt"
	case Cursor_ObjCAtSynchronizedStmt:
		return "Cursor=ObjCAtSynchronizedStmt"
	case Cursor_ObjCAutoreleasePoolStmt:
		return "Cursor=ObjCAutoreleasePoolStmt"
	case Cursor_ObjCForCollectionStmt:
		return "Cursor=ObjCForCollectionStmt"
	case Cursor_CXXCatchStmt:
		return "Cursor=CXXCatchStmt"
	case Cursor_CXXTryStmt:
		return "Cursor=CXXTryStmt"
	case Cursor_CXXForRangeStmt:
		return "Cursor=CXXForRangeStmt"
	case Cursor_SEHTryStmt:
		return "Cursor=SEHTryStmt"
	case Cursor_SEHExceptStmt:
		return "Cursor=SEHExceptStmt"
	case Cursor_SEHFinallyStmt:
		return "Cursor=SEHFinallyStmt"
	case Cursor_MSAsmStmt:
		return "Cursor=MSAsmStmt"
	case Cursor_NullStmt:
		return "Cursor=NullStmt"
	case Cursor_DeclStmt:
		return "Cursor=DeclStmt"
	case Cursor_OMPParallelDirective:
		return "Cursor=OMPParallelDirective"
	case Cursor_OMPSimdDirective:
		return "Cursor=OMPSimdDirective"
	case Cursor_OMPForDirective:
		return "Cursor=OMPForDirective"
	case Cursor_OMPSectionsDirective:
		return "Cursor=OMPSectionsDirective"
	case Cursor_OMPSectionDirective:
		return "Cursor=OMPSectionDirective"
	case Cursor_OMPSingleDirective:
		return "Cursor=OMPSingleDirective"
	case Cursor_OMPParallelForDirective:
		return "Cursor=OMPParallelForDirective"
	case Cursor_OMPParallelSectionsDirective:
		return "Cursor=OMPParallelSectionsDirective"
	case Cursor_OMPTaskDirective:
		return "Cursor=OMPTaskDirective"
	case Cursor_OMPMasterDirective:
		return "Cursor=OMPMasterDirective"
	case Cursor_OMPCriticalDirective:
		return "Cursor=OMPCriticalDirective"
	case Cursor_OMPTaskyieldDirective:
		return "Cursor=OMPTaskyieldDirective"
	case Cursor_OMPBarrierDirective:
		return "Cursor=OMPBarrierDirective"
	case Cursor_OMPTaskwaitDirective:
		return "Cursor=OMPTaskwaitDirective"
	case Cursor_OMPFlushDirective:
		return "Cursor=OMPFlushDirective"
	case Cursor_SEHLeaveStmt:
		return "Cursor=SEHLeaveStmt"
	case Cursor_OMPOrderedDirective:
		return "Cursor=OMPOrderedDirective"
	case Cursor_OMPAtomicDirective:
		return "Cursor=OMPAtomicDirective"
	case Cursor_OMPForSimdDirective:
		return "Cursor=OMPForSimdDirective"
	case Cursor_OMPParallelForSimdDirective:
		return "Cursor=OMPParallelForSimdDirective"
	case Cursor_OMPTargetDirective:
		return "Cursor=OMPTargetDirective"
	case Cursor_OMPTeamsDirective:
		return "Cursor=OMPTeamsDirective"
	case Cursor_OMPTaskgroupDirective:
		return "Cursor=OMPTaskgroupDirective"
	case Cursor_OMPCancellationPointDirective:
		return "Cursor=OMPCancellationPointDirective"
	case Cursor_OMPCancelDirective:
		return "Cursor=OMPCancelDirective"
	case Cursor_OMPTargetDataDirective:
		return "Cursor=OMPTargetDataDirective"
	case Cursor_OMPTaskLoopDirective:
		return "Cursor=OMPTaskLoopDirective"
	case Cursor_OMPTaskLoopSimdDirective:
		return "Cursor=OMPTaskLoopSimdDirective"
	case Cursor_OMPDistributeDirective:
		return "Cursor=OMPDistributeDirective"
	case Cursor_OMPTargetEnterDataDirective:
		return "Cursor=OMPTargetEnterDataDirective"
	case Cursor_OMPTargetExitDataDirective:
		return "Cursor=OMPTargetExitDataDirective"
	case Cursor_OMPTargetParallelDirective:
		return "Cursor=OMPTargetParallelDirective"
	case Cursor_OMPTargetParallelForDirective:
		return "Cursor=OMPTargetParallelForDirective"
	case Cursor_OMPTargetUpdateDirective:
		return "Cursor=OMPTargetUpdateDirective"
	case Cursor_OMPDistributeParallelForDirective:
		return "Cursor=OMPDistributeParallelForDirective"
	case Cursor_OMPDistributeParallelForSimdDirective:
		return "Cursor=OMPDistributeParallelForSimdDirective"
	case Cursor_OMPDistributeSimdDirective:
		return "Cursor=OMPDistributeSimdDirective"
	case Cursor_OMPTargetParallelForSimdDirective:
		return "Cursor=OMPTargetParallelForSimdDirective, LastStmt"
	case Cursor_TranslationUnit:
		return "Cursor=TranslationUnit"
	case Cursor_FirstAttr:
		return "Cursor=FirstAttr, UnexposedAttr"
	case Cursor_IBActionAttr:
		return "Cursor=IBActionAttr"
	case Cursor_IBOutletAttr:
		return "Cursor=IBOutletAttr"
	case Cursor_IBOutletCollectionAttr:
		return "Cursor=IBOutletCollectionAttr"
	case Cursor_CXXFinalAttr:
		return "Cursor=CXXFinalAttr"
	case Cursor_CXXOverrideAttr:
		return "Cursor=CXXOverrideAttr"
	case Cursor_AnnotateAttr:
		return "Cursor=AnnotateAttr"
	case Cursor_AsmLabelAttr:
		return "Cursor=AsmLabelAttr"
	case Cursor_PackedAttr:
		return "Cursor=PackedAttr"
	case Cursor_PureAttr:
		return "Cursor=PureAttr"
	case Cursor_ConstAttr:
		return "Cursor=ConstAttr"
	case Cursor_NoDuplicateAttr:
		return "Cursor=NoDuplicateAttr"
	case Cursor_CUDAConstantAttr:
		return "Cursor=CUDAConstantAttr"
	case Cursor_CUDADeviceAttr:
		return "Cursor=CUDADeviceAttr"
	case Cursor_CUDAGlobalAttr:
		return "Cursor=CUDAGlobalAttr"
	case Cursor_CUDAHostAttr:
		return "Cursor=CUDAHostAttr"
	case Cursor_CUDASharedAttr:
		return "Cursor=CUDASharedAttr"
	case Cursor_VisibilityAttr:
		return "Cursor=VisibilityAttr"
	case Cursor_DLLExport:
		return "Cursor=DLLExport"
	case Cursor_DLLImport:
		return "Cursor=DLLImport, LastAttr"
	case Cursor_PreprocessingDirective:
		return "Cursor=PreprocessingDirective, FirstPreprocessing"
	case Cursor_MacroDefinition:
		return "Cursor=MacroDefinition"
	case Cursor_MacroExpansion:
		return "Cursor=MacroExpansion, MacroInstantiation"
	case Cursor_InclusionDirective:
		return "Cursor=InclusionDirective, LastPreprocessing"
	case Cursor_ModuleImportDecl:
		return "Cursor=ModuleImportDecl, FirstExtraDecl"
	case Cursor_TypeAliasTemplateDecl:
		return "Cursor=TypeAliasTemplateDecl"
	case Cursor_StaticAssert:
		return "Cursor=StaticAssert, LastExtraDecl"
	case Cursor_OverloadCandidate:
		return "Cursor=OverloadCandidate"
	}

	return fmt.Sprintf("CursorKind unkown %d", int(ck))
}

func (ck CursorKind) String() string {
	return ck.Spelling()
}
