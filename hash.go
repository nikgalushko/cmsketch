// based on https://github.com/puzpuzpuz/xsync/blob/main/util_hash.go
package cmsketch

import (
	"reflect"
	"unsafe"
)

func makeSeed() uint64 {
	var s1 uint32
	for {
		s1 = runtime_fastrand()
		if s1 != 0 {
			break
		}
	}
	s2 := runtime_fastrand()
	return uint64(s1)<<32 | uint64(s2)
}

func hasher[T comparable]() func(T) uint64 {
	var zero T

	seed := makeSeed()

	if reflect.TypeOf(&zero).Elem().Kind() == reflect.Interface {
		return func(value T) uint64 {
			iValue := any(value)
			i := (*iface)(unsafe.Pointer(&iValue))
			return runtime_typehash64(i.typ, i.word, uint64(seed))
		}
	} else {
		var iZero any = zero
		i := (*iface)(unsafe.Pointer(&iZero))
		return func(value T) uint64 {
			return runtime_typehash64(i.typ, unsafe.Pointer(&value), uint64(seed))
		}
	}
}

type iface struct {
	typ  uintptr
	word unsafe.Pointer
}

func runtime_typehash64(t uintptr, p unsafe.Pointer, seed uint64) uint64 {
	if unsafe.Sizeof(uintptr(0)) == 8 {
		return uint64(runtime_typehash(t, p, uintptr(seed)))
	}

	lo := runtime_typehash(t, p, uintptr(seed))
	hi := runtime_typehash(t, p, uintptr(seed>>32))
	return uint64(hi)<<32 | uint64(lo)
}

//go:noescape
//go:linkname runtime_typehash runtime.typehash
func runtime_typehash(t uintptr, p unsafe.Pointer, h uintptr) uintptr

//go:noescape
//go:linkname runtime_fastrand runtime.fastrand
func runtime_fastrand() uint32
