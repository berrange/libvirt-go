package libvirt

/*
#cgo LDFLAGS: -lvirt
#include <libvirt/libvirt.h>
#include <libvirt/virterror.h>
*/
import "C"

import (
	"fmt"
	"unsafe"
)

type typedParamsFieldInfo struct {
	set *bool
	i   *int
	ui  *uint
	l   *int64
	ul  *uint64
	b   *bool
	d   *float64
	s   *string
}

func typedParamsUnpack(cparams []C.virTypedParameter, infomap map[string]typedParamsFieldInfo) error {
	for _, cparam := range cparams {
		name := C.GoString((*C.char)(unsafe.Pointer(&cparam.field)))
		info, ok := infomap[name]
		if !ok {
			continue
		}
		switch cparam._type {
		case C.VIR_TYPED_PARAM_INT:
			if info.i == nil {
				return fmt.Errorf("field %s expects an int", name)
			}
			*info.i = int(*(*C.int)(unsafe.Pointer(&cparam.value)))
			*info.set = true
		case C.VIR_TYPED_PARAM_UINT:
			if info.ui == nil {
				return fmt.Errorf("field %s expects a uint", name)
			}
			*info.ui = uint(*(*C.uint)(unsafe.Pointer(&cparam.value)))
			*info.set = true
		case C.VIR_TYPED_PARAM_LLONG:
			if info.l == nil {
				return fmt.Errorf("field %s expects an int64", name)
			}
			*info.l = int64(*(*C.longlong)(unsafe.Pointer(&cparam.value)))
			*info.set = true
		case C.VIR_TYPED_PARAM_ULLONG:
			if info.ul == nil {
				return fmt.Errorf("field %s expects a uint64", name)
			}
			*info.ul = uint64(*(*C.ulonglong)(unsafe.Pointer(&cparam.value)))
			*info.set = true
		case C.VIR_TYPED_PARAM_DOUBLE:
			if info.d == nil {
				return fmt.Errorf("field %s expects a float64", name)
			}
			*info.d = float64(*(*C.double)(unsafe.Pointer(&cparam.value)))
			*info.set = true
		case C.VIR_TYPED_PARAM_BOOLEAN:
			if info.b == nil {
				return fmt.Errorf("field %s expects a bool", name)
			}
			*info.b = *(*C.char)(unsafe.Pointer(&cparam.value)) == 1
			*info.set = true
		case C.VIR_TYPED_PARAM_STRING:
			if info.s == nil {
				return fmt.Errorf("field %s expects a string", name)
			}
			*info.s = C.GoString(*(**C.char)(unsafe.Pointer(&cparam.value)))
			*info.set = true
		}
	}

	return nil
}

func typedParamsPack(cparams []C.virTypedParameter, infomap map[string]typedParamsFieldInfo) error {
	for _, cparam := range cparams {
		name := C.GoString((*C.char)(unsafe.Pointer(&cparam.field)))
		info, ok := infomap[name]
		if !ok {
			continue
		}
		if !*info.set {
			continue
		}
		switch cparam._type {
		case C.VIR_TYPED_PARAM_INT:
			if info.i == nil {
				return fmt.Errorf("field %s expects an int", name)
			}
			*(*C.int)(unsafe.Pointer(&cparam.value)) = C.int(*info.i)
		case C.VIR_TYPED_PARAM_UINT:
			if info.ui == nil {
				return fmt.Errorf("field %s expects a uint", name)
			}
			*(*C.uint)(unsafe.Pointer(&cparam.value)) = C.uint(*info.ui)
		case C.VIR_TYPED_PARAM_LLONG:
			if info.l == nil {
				return fmt.Errorf("field %s expects an int64", name)
			}
			*(*C.longlong)(unsafe.Pointer(&cparam.value)) = C.longlong(*info.l)
		case C.VIR_TYPED_PARAM_ULLONG:
			if info.ul == nil {
				return fmt.Errorf("field %s expects a uint64", name)
			}
			*(*C.ulonglong)(unsafe.Pointer(&cparam.value)) = C.ulonglong(*info.ul)
		case C.VIR_TYPED_PARAM_DOUBLE:
			if info.d == nil {
				return fmt.Errorf("field %s expects a float64", name)
			}
			*(*C.double)(unsafe.Pointer(&cparam.value)) = C.double(*info.d)
		case C.VIR_TYPED_PARAM_BOOLEAN:
			if info.b == nil {
				return fmt.Errorf("field %s expects a bool", name)
			}
			if *info.b {
				*(*C.char)(unsafe.Pointer(&cparam.value)) = 1
			} else {
				*(*C.char)(unsafe.Pointer(&cparam.value)) = 0
			}
		case C.VIR_TYPED_PARAM_STRING:
			if info.s == nil {
				return fmt.Errorf("field %s expects a string", name)
			}

			*(**C.char)(unsafe.Pointer(&cparam.value)) = C.CString(*info.s)
		}
	}

	return nil
}
