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

type StorageVolCreateFlags int

const (
	VIR_STORAGE_VOL_CREATE_PREALLOC_METADATA = StorageVolCreateFlags(C.VIR_STORAGE_VOL_CREATE_PREALLOC_METADATA)
	VIR_STORAGE_VOL_CREATE_REFLINK           = StorageVolCreateFlags(C.VIR_STORAGE_VOL_CREATE_REFLINK)
)

type StorageVolDeleteFlags int

const (
	VIR_STORAGE_VOL_DELETE_NORMAL         = StorageVolDeleteFlags(C.VIR_STORAGE_VOL_DELETE_NORMAL)         // Delete metadata only (fast)
	VIR_STORAGE_VOL_DELETE_ZEROED         = StorageVolDeleteFlags(C.VIR_STORAGE_VOL_DELETE_ZEROED)         // Clear all data to zeros (slow)
	VIR_STORAGE_VOL_DELETE_WITH_SNAPSHOTS = StorageVolDeleteFlags(C.VIR_STORAGE_VOL_DELETE_WITH_SNAPSHOTS) // Force removal of volume, even if in use
)

type StorageVolResizeFlags int

const (
	VIR_STORAGE_VOL_RESIZE_ALLOCATE = StorageVolResizeFlags(C.VIR_STORAGE_VOL_RESIZE_ALLOCATE) // force allocation of new size
	VIR_STORAGE_VOL_RESIZE_DELTA    = StorageVolResizeFlags(C.VIR_STORAGE_VOL_RESIZE_DELTA)    // size is relative to current
	VIR_STORAGE_VOL_RESIZE_SHRINK   = StorageVolResizeFlags(C.VIR_STORAGE_VOL_RESIZE_SHRINK)   // allow decrease in capacity
)

type StorageVolType int

const (
	VIR_STORAGE_VOL_FILE    = StorageVolType(C.VIR_STORAGE_VOL_FILE)    // Regular file based volumes
	VIR_STORAGE_VOL_BLOCK   = StorageVolType(C.VIR_STORAGE_VOL_BLOCK)   // Block based volumes
	VIR_STORAGE_VOL_DIR     = StorageVolType(C.VIR_STORAGE_VOL_DIR)     // Directory-passthrough based volume
	VIR_STORAGE_VOL_NETWORK = StorageVolType(C.VIR_STORAGE_VOL_NETWORK) //Network volumes like RBD (RADOS Block Device)
	VIR_STORAGE_VOL_NETDIR  = StorageVolType(C.VIR_STORAGE_VOL_NETDIR)  // Network accessible directory that can contain other network volumes
	VIR_STORAGE_VOL_PLOOP   = StorageVolType(C.VIR_STORAGE_VOL_PLOOP)   // Ploop directory based volumes
)

type StorageVolWipeAlgorithm int

const (
	VIR_STORAGE_VOL_WIPE_ALG_ZERO       = StorageVolWipeAlgorithm(C.VIR_STORAGE_VOL_WIPE_ALG_ZERO)       // 1-pass, all zeroes
	VIR_STORAGE_VOL_WIPE_ALG_NNSA       = StorageVolWipeAlgorithm(C.VIR_STORAGE_VOL_WIPE_ALG_NNSA)       // 4-pass NNSA Policy Letter NAP-14.1-C (XVI-8)
	VIR_STORAGE_VOL_WIPE_ALG_DOD        = StorageVolWipeAlgorithm(C.VIR_STORAGE_VOL_WIPE_ALG_DOD)        // 4-pass DoD 5220.22-M section 8-306 procedure
	VIR_STORAGE_VOL_WIPE_ALG_BSI        = StorageVolWipeAlgorithm(C.VIR_STORAGE_VOL_WIPE_ALG_BSI)        // 9-pass method recommended by the German Center of Security in Information Technologies
	VIR_STORAGE_VOL_WIPE_ALG_GUTMANN    = StorageVolWipeAlgorithm(C.VIR_STORAGE_VOL_WIPE_ALG_GUTMANN)    // The canonical 35-pass sequence
	VIR_STORAGE_VOL_WIPE_ALG_SCHNEIER   = StorageVolWipeAlgorithm(C.VIR_STORAGE_VOL_WIPE_ALG_SCHNEIER)   // 7-pass method described by Bruce Schneier in "Applied Cryptography" (1996)
	VIR_STORAGE_VOL_WIPE_ALG_PFITZNER7  = StorageVolWipeAlgorithm(C.VIR_STORAGE_VOL_WIPE_ALG_PFITZNER7)  // 7-pass random
	VIR_STORAGE_VOL_WIPE_ALG_PFITZNER33 = StorageVolWipeAlgorithm(C.VIR_STORAGE_VOL_WIPE_ALG_PFITZNER33) // 33-pass random
	VIR_STORAGE_VOL_WIPE_ALG_RANDOM     = StorageVolWipeAlgorithm(C.VIR_STORAGE_VOL_WIPE_ALG_RANDOM)     // 1-pass random
	VIR_STORAGE_VOL_WIPE_ALG_TRIM       = StorageVolWipeAlgorithm(C.VIR_STORAGE_VOL_WIPE_ALG_TRIM)       // Trim the underlying storage
)

type StorageXMLFlags int

const (
	VIR_STORAGE_XML_INACTIVE = StorageXMLFlags(C.VIR_STORAGE_XML_INACTIVE)
)

type StorageVol struct {
	ptr C.virStorageVolPtr
}

type StorageVolInfo struct {
	Type       StorageVolType
	Capacity   uint64
	Allocation uint64
}

func (v *StorageVol) Delete(flags StorageVolDeleteFlags) error {
	result := C.virStorageVolDelete(v.ptr, C.uint(flags))
	if result == -1 {
		return GetLastError()
	}
	return nil
}

func (v *StorageVol) Free() error {
	if result := C.virStorageVolFree(v.ptr); result != 0 {
		return GetLastError()
	}
	v.ptr = nil
	return nil
}

func (v *StorageVol) GetInfo() (*StorageVolInfo, error) {
	var cinfo C.virStorageVolInfo
	result := C.virStorageVolGetInfo(v.ptr, &cinfo)
	if result == -1 {
		return nil, GetLastError()
	}
	return &StorageVolInfo{
		Type:       StorageVolType(cinfo._type),
		Capacity:   uint64(cinfo.capacity),
		Allocation: uint64(cinfo.allocation),
	}, nil
}

func (v *StorageVol) GetKey() (string, error) {
	key := C.virStorageVolGetKey(v.ptr)
	if key == nil {
		return "", GetLastError()
	}
	return C.GoString(key), nil
}

func (v *StorageVol) GetName() (string, error) {
	name := C.virStorageVolGetName(v.ptr)
	if name == nil {
		return "", GetLastError()
	}
	return C.GoString(name), nil
}

func (v *StorageVol) GetPath() (string, error) {
	result := C.virStorageVolGetPath(v.ptr)
	if result == nil {
		return "", GetLastError()
	}
	path := C.GoString(result)
	C.free(unsafe.Pointer(result))
	return path, nil
}

func (v *StorageVol) GetXMLDesc(flags uint32) (string, error) {
	result := C.virStorageVolGetXMLDesc(v.ptr, C.uint(flags))
	if result == nil {
		return "", GetLastError()
	}
	xml := C.GoString(result)
	C.free(unsafe.Pointer(result))
	return xml, nil
}

func (v *StorageVol) Resize(capacity uint64, flags StorageVolResizeFlags) error {
	result := C.virStorageVolResize(v.ptr, C.ulonglong(capacity), C.uint(flags))
	if result == -1 {
		return GetLastError()
	}
	return nil
}

func (v *StorageVol) Wipe(flags uint32) error {
	result := C.virStorageVolWipe(v.ptr, C.uint(flags))
	if result == -1 {
		return GetLastError()
	}
	return nil
}
func (v *StorageVol) WipePattern(algorithm StorageVolWipeAlgorithm, flags uint32) error {
	result := C.virStorageVolWipePattern(v.ptr, C.uint(algorithm), C.uint(flags))
	if result == -1 {
		return GetLastError()
	}
	return nil
}

func (v *StorageVol) Upload(stream *Stream, offset, length uint64, flags uint32) error {
	if C.virStorageVolUpload(v.ptr, stream.ptr, C.ulonglong(offset),
		C.ulonglong(length), C.uint(flags)) == -1 {
		return GetLastError()
	}
	return nil
}

func (v *StorageVol) Download(stream *Stream, offset, length uint64, flags uint32) error {
	if C.virStorageVolDownload(v.ptr, stream.ptr, C.ulonglong(offset),
		C.ulonglong(length), C.uint(flags)) == -1 {
		return GetLastError()
	}
	return nil
}

func (v *StorageVol) LookupPoolByVolume() (*StoragePool, error) {
	poolPtr := C.virStoragePoolLookupByVolume(v.ptr)
	if poolPtr == nil {
		return nil, GetLastError()
	}
	return &StoragePool{ptr: poolPtr}, nil
}
