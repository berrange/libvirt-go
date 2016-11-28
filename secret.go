package libvirt

/*
#cgo LDFLAGS: -lvirt
#include <libvirt/libvirt.h>
#include <libvirt/virterror.h>
#include <stdlib.h>
*/
import "C"

import (
	"unsafe"
)

// virSecretUsageType
const (
	VIR_SECRET_USAGE_TYPE_NONE   = C.VIR_SECRET_USAGE_TYPE_NONE
	VIR_SECRET_USAGE_TYPE_VOLUME = C.VIR_SECRET_USAGE_TYPE_VOLUME
	VIR_SECRET_USAGE_TYPE_CEPH   = C.VIR_SECRET_USAGE_TYPE_CEPH
	VIR_SECRET_USAGE_TYPE_ISCSI  = C.VIR_SECRET_USAGE_TYPE_ISCSI
)

type VirSecret struct {
	ptr C.virSecretPtr
}

func (s *VirSecret) Free() error {
	if result := C.virSecretFree(s.ptr); result != 0 {
		return GetLastError()
	}
	return nil
}

func (s *VirSecret) Undefine() error {
	result := C.virSecretUndefine(s.ptr)
	if result == -1 {
		return GetLastError()
	}
	return nil
}

func (s *VirSecret) GetUUID() ([]byte, error) {
	var cUuid [C.VIR_UUID_BUFLEN](byte)
	cuidPtr := unsafe.Pointer(&cUuid)
	result := C.virSecretGetUUID(s.ptr, (*C.uchar)(cuidPtr))
	if result != 0 {
		return []byte{}, GetLastError()
	}
	return C.GoBytes(cuidPtr, C.VIR_UUID_BUFLEN), nil
}

func (s *VirSecret) GetUUIDString() (string, error) {
	var cUuid [C.VIR_UUID_STRING_BUFLEN](C.char)
	cuidPtr := unsafe.Pointer(&cUuid)
	result := C.virSecretGetUUIDString(s.ptr, (*C.char)(cuidPtr))
	if result != 0 {
		return "", GetLastError()
	}
	return C.GoString((*C.char)(cuidPtr)), nil
}

func (s *VirSecret) GetUsageID() (string, error) {
	result := C.virSecretGetUsageID(s.ptr)
	if result == nil {
		return "", GetLastError()
	}
	return C.GoString(result), nil
}

func (s *VirSecret) GetUsageType() (int, error) {
	result := int(C.virSecretGetUsageType(s.ptr))
	if result == -1 {
		return 0, GetLastError()
	}
	return result, nil
}

func (s *VirSecret) GetXMLDesc(flags uint32) (string, error) {
	result := C.virSecretGetXMLDesc(s.ptr, C.uint(flags))
	if result == nil {
		return "", GetLastError()
	}
	xml := C.GoString(result)
	C.free(unsafe.Pointer(result))
	return xml, nil
}
