package z3

// #include "go-z3.h"
import "C"

// Lt creates a "equal" comparison for floating point numbers.
//
// Maps to: Z3_mk_lt
func (a *AST) FpaEq(a2 *AST) *AST {
	return &AST{
		rawCtx: a.rawCtx,
		rawAST: C.Z3_mk_fpa_eq(a.rawCtx, a.rawAST, a2.rawAST),
	}
}

// Lt creates a "less than" comparison.
//
// Maps to: Z3_mk_fpa_lt
func (a *AST) FpaLt(a2 *AST) *AST {
	return &AST{
		rawCtx: a.rawCtx,
		rawAST: C.Z3_mk_fpa_lt(a.rawCtx, a.rawAST, a2.rawAST),
	}
}

// FpaLe creates a "less than" comparison for floating point numbers.
//
// Maps to: Z3_mk_fpa_leq
func (a *AST) FpaLe(a2 *AST) *AST {
	return &AST{
		rawCtx: a.rawCtx,
		rawAST: C.Z3_mk_fpa_leq(a.rawCtx, a.rawAST, a2.rawAST),
	}
}

// FpaGe creates a "greater equal" comparison for floating point numbers.
//
// Maps to: Z3_mk_fpa_geq
func (a *AST) FpaGe(a2 *AST) *AST {
	return &AST{
		rawCtx: a.rawCtx,
		rawAST: C.Z3_mk_fpa_geq(a.rawCtx, a.rawAST, a2.rawAST),
	}
}

// FpaGt creates a "greater than" comparison for floating point numbers.
//
// Maps to: Z3_mk_fpa_gt
func (a *AST) FpaGt(a2 *AST) *AST {
	return &AST{
		rawCtx: a.rawCtx,
		rawAST: C.Z3_mk_fpa_gt(a.rawCtx, a.rawAST, a2.rawAST),
	}
}
