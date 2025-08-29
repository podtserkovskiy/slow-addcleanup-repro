package repro

import (
	"runtime"
	"testing"
)

func BenchmarkAllocateFree(b *testing.B) {
	for b.Loop() {
		cstr := Allocate()
		Free(cstr)
		runtime.KeepAlive(cstr)
	}
}

func BenchmarkAllocateDefer(b *testing.B) {
	for b.Loop() {
		func() {
			cstr := Allocate()
			defer Free(cstr)
			runtime.KeepAlive(cstr)
		}()
	}
}

func BenchmarkAllocateAddCleanup(b *testing.B) {
	for b.Loop() {
		cstr := AllocateAddCleanup()
		runtime.KeepAlive(cstr)
	}
}

func BenchmarkAllocateWithFinalizer(b *testing.B) {
	for b.Loop() {
		cstr := AllocateSetFinalizer()
		runtime.KeepAlive(cstr)
	}
}
