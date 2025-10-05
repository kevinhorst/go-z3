package z3

// #include "go-z3.h"
import "C"

// Sort represents a sort in Z3.
type Sort struct {
	rawCtx  C.Z3_context
	rawSort C.Z3_sort
}

// BoolSort returns the boolean type.
func (c *Context) BoolSort() *Sort {
	return &Sort{
		rawCtx:  c.raw,
		rawSort: C.Z3_mk_bool_sort(c.raw),
	}
}

// IntSort returns the int type.
func (c *Context) IntSort() *Sort {
	return &Sort{
		rawCtx:  c.raw,
		rawSort: C.Z3_mk_int_sort(c.raw),
	}
}

func (c *Context) Float32Sort() *Sort {
	return &Sort{
		rawCtx:  c.raw,
		rawSort: C.Z3_mk_fpa_sort_single(c.raw),
	}
}

func (c *Context) Float64Sort() *Sort {
	return &Sort{
		rawCtx:  c.raw,
		rawSort: C.Z3_mk_fpa_sort_double(c.raw),
	}
}

func (c *Context) StringSort() *Sort {
	return &Sort{
		rawCtx:  c.raw,
		rawSort: C.Z3_mk_string_sort(c.raw),
	}
}
