package repro

/*
#include "repro.h"
*/
import "C"
import (
	"runtime"
)

type Cstr struct {
	cpointer *C.char
}

func AllocateAddCleanup() *Cstr {
	cstr := &Cstr{C.allocate()}
	runtime.AddCleanup(cstr, func(c *C.char) { C.free_allocated(c) }, cstr.cpointer)
	return cstr
}

func AllocateSetFinalizer() *Cstr {
	cstr := &Cstr{C.allocate()}
	runtime.SetFinalizer(cstr, func(c *Cstr) { C.free_allocated(c.cpointer) })
	return cstr
}

func Allocate() *C.char {
	return C.allocate()
}

func Free(c *C.char) {
	C.free_allocated(c)
}
