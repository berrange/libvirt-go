package libvirt

import (
	"fmt"
	"os"
	"reflect"
	"sync"
	"unsafe"
)

/*
#cgo LDFLAGS: -lvirt
#include <libvirt/libvirt.h>
#include <libvirt/virterror.h>
#include <stdlib.h>
#include "go_libvirt.h"
*/
import "C"

func init() {
	C.virInitialize()
}

const (
	LIBVIR_VERSION_NUMBER = uint32(C.LIBVIR_VERSION_NUMBER)
)

type VirConnectCloseReason int

const (
	VIR_CONNECT_CLOSE_REASON_ERROR     = VirConnectCloseReason(C.VIR_CONNECT_CLOSE_REASON_ERROR)
	VIR_CONNECT_CLOSE_REASON_EOF       = VirConnectCloseReason(C.VIR_CONNECT_CLOSE_REASON_EOF)
	VIR_CONNECT_CLOSE_REASON_KEEPALIVE = VirConnectCloseReason(C.VIR_CONNECT_CLOSE_REASON_KEEPALIVE)
	VIR_CONNECT_CLOSE_REASON_CLIENT    = VirConnectCloseReason(C.VIR_CONNECT_CLOSE_REASON_CLIENT)
)

type VirConnectListAllDomainsFlags int

const (
	VIR_CONNECT_LIST_DOMAINS_ACTIVE         = VirConnectListAllDomainsFlags(C.VIR_CONNECT_LIST_DOMAINS_ACTIVE)
	VIR_CONNECT_LIST_DOMAINS_INACTIVE       = VirConnectListAllDomainsFlags(C.VIR_CONNECT_LIST_DOMAINS_INACTIVE)
	VIR_CONNECT_LIST_DOMAINS_PERSISTENT     = VirConnectListAllDomainsFlags(C.VIR_CONNECT_LIST_DOMAINS_PERSISTENT)
	VIR_CONNECT_LIST_DOMAINS_TRANSIENT      = VirConnectListAllDomainsFlags(C.VIR_CONNECT_LIST_DOMAINS_TRANSIENT)
	VIR_CONNECT_LIST_DOMAINS_RUNNING        = VirConnectListAllDomainsFlags(C.VIR_CONNECT_LIST_DOMAINS_RUNNING)
	VIR_CONNECT_LIST_DOMAINS_PAUSED         = VirConnectListAllDomainsFlags(C.VIR_CONNECT_LIST_DOMAINS_PAUSED)
	VIR_CONNECT_LIST_DOMAINS_SHUTOFF        = VirConnectListAllDomainsFlags(C.VIR_CONNECT_LIST_DOMAINS_SHUTOFF)
	VIR_CONNECT_LIST_DOMAINS_OTHER          = VirConnectListAllDomainsFlags(C.VIR_CONNECT_LIST_DOMAINS_OTHER)
	VIR_CONNECT_LIST_DOMAINS_MANAGEDSAVE    = VirConnectListAllDomainsFlags(C.VIR_CONNECT_LIST_DOMAINS_MANAGEDSAVE)
	VIR_CONNECT_LIST_DOMAINS_NO_MANAGEDSAVE = VirConnectListAllDomainsFlags(C.VIR_CONNECT_LIST_DOMAINS_NO_MANAGEDSAVE)
	VIR_CONNECT_LIST_DOMAINS_AUTOSTART      = VirConnectListAllDomainsFlags(C.VIR_CONNECT_LIST_DOMAINS_AUTOSTART)
	VIR_CONNECT_LIST_DOMAINS_NO_AUTOSTART   = VirConnectListAllDomainsFlags(C.VIR_CONNECT_LIST_DOMAINS_NO_AUTOSTART)
	VIR_CONNECT_LIST_DOMAINS_HAS_SNAPSHOT   = VirConnectListAllDomainsFlags(C.VIR_CONNECT_LIST_DOMAINS_HAS_SNAPSHOT)
	VIR_CONNECT_LIST_DOMAINS_NO_SNAPSHOT    = VirConnectListAllDomainsFlags(C.VIR_CONNECT_LIST_DOMAINS_NO_SNAPSHOT)
)

type VirConnectListAllNetworksFlags int

const (
	VIR_CONNECT_LIST_NETWORKS_INACTIVE     = VirConnectListAllNetworksFlags(C.VIR_CONNECT_LIST_NETWORKS_INACTIVE)
	VIR_CONNECT_LIST_NETWORKS_ACTIVE       = VirConnectListAllNetworksFlags(C.VIR_CONNECT_LIST_NETWORKS_ACTIVE)
	VIR_CONNECT_LIST_NETWORKS_PERSISTENT   = VirConnectListAllNetworksFlags(C.VIR_CONNECT_LIST_NETWORKS_PERSISTENT)
	VIR_CONNECT_LIST_NETWORKS_TRANSIENT    = VirConnectListAllNetworksFlags(C.VIR_CONNECT_LIST_NETWORKS_TRANSIENT)
	VIR_CONNECT_LIST_NETWORKS_AUTOSTART    = VirConnectListAllNetworksFlags(C.VIR_CONNECT_LIST_NETWORKS_AUTOSTART)
	VIR_CONNECT_LIST_NETWORKS_NO_AUTOSTART = VirConnectListAllNetworksFlags(C.VIR_CONNECT_LIST_NETWORKS_NO_AUTOSTART)
)

type VirConnectListAllStoragePoolsFlags int

const (
	VIR_CONNECT_LIST_STORAGE_POOLS_INACTIVE     = VirConnectListAllStoragePoolsFlags(C.VIR_CONNECT_LIST_STORAGE_POOLS_INACTIVE)
	VIR_CONNECT_LIST_STORAGE_POOLS_ACTIVE       = VirConnectListAllStoragePoolsFlags(C.VIR_CONNECT_LIST_STORAGE_POOLS_ACTIVE)
	VIR_CONNECT_LIST_STORAGE_POOLS_PERSISTENT   = VirConnectListAllStoragePoolsFlags(C.VIR_CONNECT_LIST_STORAGE_POOLS_PERSISTENT)
	VIR_CONNECT_LIST_STORAGE_POOLS_TRANSIENT    = VirConnectListAllStoragePoolsFlags(C.VIR_CONNECT_LIST_STORAGE_POOLS_TRANSIENT)
	VIR_CONNECT_LIST_STORAGE_POOLS_AUTOSTART    = VirConnectListAllStoragePoolsFlags(C.VIR_CONNECT_LIST_STORAGE_POOLS_AUTOSTART)
	VIR_CONNECT_LIST_STORAGE_POOLS_NO_AUTOSTART = VirConnectListAllStoragePoolsFlags(C.VIR_CONNECT_LIST_STORAGE_POOLS_NO_AUTOSTART)
	VIR_CONNECT_LIST_STORAGE_POOLS_DIR          = VirConnectListAllStoragePoolsFlags(C.VIR_CONNECT_LIST_STORAGE_POOLS_DIR)
	VIR_CONNECT_LIST_STORAGE_POOLS_FS           = VirConnectListAllStoragePoolsFlags(C.VIR_CONNECT_LIST_STORAGE_POOLS_FS)
	VIR_CONNECT_LIST_STORAGE_POOLS_NETFS        = VirConnectListAllStoragePoolsFlags(C.VIR_CONNECT_LIST_STORAGE_POOLS_NETFS)
	VIR_CONNECT_LIST_STORAGE_POOLS_LOGICAL      = VirConnectListAllStoragePoolsFlags(C.VIR_CONNECT_LIST_STORAGE_POOLS_LOGICAL)
	VIR_CONNECT_LIST_STORAGE_POOLS_DISK         = VirConnectListAllStoragePoolsFlags(C.VIR_CONNECT_LIST_STORAGE_POOLS_DISK)
	VIR_CONNECT_LIST_STORAGE_POOLS_ISCSI        = VirConnectListAllStoragePoolsFlags(C.VIR_CONNECT_LIST_STORAGE_POOLS_ISCSI)
	VIR_CONNECT_LIST_STORAGE_POOLS_SCSI         = VirConnectListAllStoragePoolsFlags(C.VIR_CONNECT_LIST_STORAGE_POOLS_SCSI)
	VIR_CONNECT_LIST_STORAGE_POOLS_MPATH        = VirConnectListAllStoragePoolsFlags(C.VIR_CONNECT_LIST_STORAGE_POOLS_MPATH)
	VIR_CONNECT_LIST_STORAGE_POOLS_RBD          = VirConnectListAllStoragePoolsFlags(C.VIR_CONNECT_LIST_STORAGE_POOLS_RBD)
	VIR_CONNECT_LIST_STORAGE_POOLS_SHEEPDOG     = VirConnectListAllStoragePoolsFlags(C.VIR_CONNECT_LIST_STORAGE_POOLS_SHEEPDOG)
	VIR_CONNECT_LIST_STORAGE_POOLS_GLUSTER      = VirConnectListAllStoragePoolsFlags(C.VIR_CONNECT_LIST_STORAGE_POOLS_GLUSTER)
	VIR_CONNECT_LIST_STORAGE_POOLS_ZFS          = VirConnectListAllStoragePoolsFlags(C.VIR_CONNECT_LIST_STORAGE_POOLS_ZFS)
)

type VirConnectBaselineCPUFlags int

const (
	VIR_CONNECT_BASELINE_CPU_EXPAND_FEATURES = VirConnectBaselineCPUFlags(C.VIR_CONNECT_BASELINE_CPU_EXPAND_FEATURES)
	VIR_CONNECT_BASELINE_CPU_MIGRATABLE      = VirConnectBaselineCPUFlags(C.VIR_CONNECT_BASELINE_CPU_MIGRATABLE)
)

type VirConnectCompareCPUFlags int

const (
	VIR_CONNECT_COMPARE_CPU_FAIL_INCOMPATIBLE = VirConnectCompareCPUFlags(C.VIR_CONNECT_COMPARE_CPU_FAIL_INCOMPATIBLE)
)

type VirConnectListAllInterfacesFlags int

const (
	VIR_CONNECT_LIST_INTERFACES_INACTIVE = VirConnectListAllInterfacesFlags(C.VIR_CONNECT_LIST_INTERFACES_INACTIVE)
	VIR_CONNECT_LIST_INTERFACES_ACTIVE   = VirConnectListAllInterfacesFlags(C.VIR_CONNECT_LIST_INTERFACES_ACTIVE)
)

type VirConnectListAllNodeDeviceFlags int

const (
	VIR_CONNECT_LIST_NODE_DEVICES_CAP_SYSTEM        = VirConnectListAllNodeDeviceFlags(C.VIR_CONNECT_LIST_NODE_DEVICES_CAP_SYSTEM)
	VIR_CONNECT_LIST_NODE_DEVICES_CAP_PCI_DEV       = VirConnectListAllNodeDeviceFlags(C.VIR_CONNECT_LIST_NODE_DEVICES_CAP_PCI_DEV)
	VIR_CONNECT_LIST_NODE_DEVICES_CAP_USB_DEV       = VirConnectListAllNodeDeviceFlags(C.VIR_CONNECT_LIST_NODE_DEVICES_CAP_USB_DEV)
	VIR_CONNECT_LIST_NODE_DEVICES_CAP_USB_INTERFACE = VirConnectListAllNodeDeviceFlags(C.VIR_CONNECT_LIST_NODE_DEVICES_CAP_USB_INTERFACE)
	VIR_CONNECT_LIST_NODE_DEVICES_CAP_NET           = VirConnectListAllNodeDeviceFlags(C.VIR_CONNECT_LIST_NODE_DEVICES_CAP_NET)
	VIR_CONNECT_LIST_NODE_DEVICES_CAP_SCSI_HOST     = VirConnectListAllNodeDeviceFlags(C.VIR_CONNECT_LIST_NODE_DEVICES_CAP_SCSI_HOST)
	VIR_CONNECT_LIST_NODE_DEVICES_CAP_SCSI_TARGET   = VirConnectListAllNodeDeviceFlags(C.VIR_CONNECT_LIST_NODE_DEVICES_CAP_SCSI_TARGET)
	VIR_CONNECT_LIST_NODE_DEVICES_CAP_SCSI          = VirConnectListAllNodeDeviceFlags(C.VIR_CONNECT_LIST_NODE_DEVICES_CAP_SCSI)
	VIR_CONNECT_LIST_NODE_DEVICES_CAP_STORAGE       = VirConnectListAllNodeDeviceFlags(C.VIR_CONNECT_LIST_NODE_DEVICES_CAP_STORAGE)
	VIR_CONNECT_LIST_NODE_DEVICES_CAP_FC_HOST       = VirConnectListAllNodeDeviceFlags(C.VIR_CONNECT_LIST_NODE_DEVICES_CAP_FC_HOST)
	VIR_CONNECT_LIST_NODE_DEVICES_CAP_VPORTS        = VirConnectListAllNodeDeviceFlags(C.VIR_CONNECT_LIST_NODE_DEVICES_CAP_VPORTS)
	VIR_CONNECT_LIST_NODE_DEVICES_CAP_SCSI_GENERIC  = VirConnectListAllNodeDeviceFlags(C.VIR_CONNECT_LIST_NODE_DEVICES_CAP_SCSI_GENERIC)
)

type VirConnectListAllSecretsFlags int

const (
	VIR_CONNECT_LIST_SECRETS_EPHEMERAL    = VirConnectListAllSecretsFlags(C.VIR_CONNECT_LIST_SECRETS_EPHEMERAL)
	VIR_CONNECT_LIST_SECRETS_NO_EPHEMERAL = VirConnectListAllSecretsFlags(C.VIR_CONNECT_LIST_SECRETS_NO_EPHEMERAL)
	VIR_CONNECT_LIST_SECRETS_PRIVATE      = VirConnectListAllSecretsFlags(C.VIR_CONNECT_LIST_SECRETS_PRIVATE)
	VIR_CONNECT_LIST_SECRETS_NO_PRIVATE   = VirConnectListAllSecretsFlags(C.VIR_CONNECT_LIST_SECRETS_NO_PRIVATE)
)

type VirConnectGetAllDomainStatsFlags int

const (
	VIR_CONNECT_GET_ALL_DOMAINS_STATS_ACTIVE        = VirConnectGetAllDomainStatsFlags(C.VIR_CONNECT_GET_ALL_DOMAINS_STATS_ACTIVE)
	VIR_CONNECT_GET_ALL_DOMAINS_STATS_INACTIVE      = VirConnectGetAllDomainStatsFlags(C.VIR_CONNECT_GET_ALL_DOMAINS_STATS_INACTIVE)
	VIR_CONNECT_GET_ALL_DOMAINS_STATS_PERSISTENT    = VirConnectGetAllDomainStatsFlags(C.VIR_CONNECT_GET_ALL_DOMAINS_STATS_PERSISTENT)
	VIR_CONNECT_GET_ALL_DOMAINS_STATS_TRANSIENT     = VirConnectGetAllDomainStatsFlags(C.VIR_CONNECT_GET_ALL_DOMAINS_STATS_TRANSIENT)
	VIR_CONNECT_GET_ALL_DOMAINS_STATS_RUNNING       = VirConnectGetAllDomainStatsFlags(C.VIR_CONNECT_GET_ALL_DOMAINS_STATS_RUNNING)
	VIR_CONNECT_GET_ALL_DOMAINS_STATS_PAUSED        = VirConnectGetAllDomainStatsFlags(C.VIR_CONNECT_GET_ALL_DOMAINS_STATS_PAUSED)
	VIR_CONNECT_GET_ALL_DOMAINS_STATS_SHUTOFF       = VirConnectGetAllDomainStatsFlags(C.VIR_CONNECT_GET_ALL_DOMAINS_STATS_SHUTOFF)
	VIR_CONNECT_GET_ALL_DOMAINS_STATS_OTHER         = VirConnectGetAllDomainStatsFlags(C.VIR_CONNECT_GET_ALL_DOMAINS_STATS_OTHER)
	VIR_CONNECT_GET_ALL_DOMAINS_STATS_BACKING       = VirConnectGetAllDomainStatsFlags(C.VIR_CONNECT_GET_ALL_DOMAINS_STATS_BACKING)
	VIR_CONNECT_GET_ALL_DOMAINS_STATS_ENFORCE_STATS = VirConnectGetAllDomainStatsFlags(C.VIR_CONNECT_GET_ALL_DOMAINS_STATS_ENFORCE_STATS)
)

type VirConnectFlags int

const (
	VIR_CONNECT_RO         = VirConnectFlags(C.VIR_CONNECT_RO)
	VIR_CONNECT_NO_ALIASES = VirConnectFlags(C.VIR_CONNECT_NO_ALIASES)
)

type VirConnectDomainEventAgentLifecycleState int

const (
	VIR_CONNECT_DOMAIN_EVENT_AGENT_LIFECYCLE_STATE_CONNECTED    = VirConnectDomainEventAgentLifecycleState(C.VIR_CONNECT_DOMAIN_EVENT_AGENT_LIFECYCLE_STATE_CONNECTED)
	VIR_CONNECT_DOMAIN_EVENT_AGENT_LIFECYCLE_STATE_DISCONNECTED = VirConnectDomainEventAgentLifecycleState(C.VIR_CONNECT_DOMAIN_EVENT_AGENT_LIFECYCLE_STATE_DISCONNECTED)
)

type VirConnectDomainEventAgentLifecycleReason int

const (
	VIR_CONNECT_DOMAIN_EVENT_AGENT_LIFECYCLE_REASON_UNKNOWN        = VirConnectDomainEventAgentLifecycleReason(C.VIR_CONNECT_DOMAIN_EVENT_AGENT_LIFECYCLE_REASON_UNKNOWN)
	VIR_CONNECT_DOMAIN_EVENT_AGENT_LIFECYCLE_REASON_DOMAIN_STARTED = VirConnectDomainEventAgentLifecycleReason(C.VIR_CONNECT_DOMAIN_EVENT_AGENT_LIFECYCLE_REASON_DOMAIN_STARTED)
	VIR_CONNECT_DOMAIN_EVENT_AGENT_LIFECYCLE_REASON_CHANNEL        = VirConnectDomainEventAgentLifecycleReason(C.VIR_CONNECT_DOMAIN_EVENT_AGENT_LIFECYCLE_REASON_CHANNEL)
)

type VirCPUCompareResult int

const (
	VIR_CPU_COMPARE_ERROR        = VirCPUCompareResult(C.VIR_CPU_COMPARE_ERROR)
	VIR_CPU_COMPARE_INCOMPATIBLE = VirCPUCompareResult(C.VIR_CPU_COMPARE_INCOMPATIBLE)
	VIR_CPU_COMPARE_IDENTICAL    = VirCPUCompareResult(C.VIR_CPU_COMPARE_IDENTICAL)
	VIR_CPU_COMPARE_SUPERSET     = VirCPUCompareResult(C.VIR_CPU_COMPARE_SUPERSET)
)

type VirNodeAllocPagesFlags int

const (
	VIR_NODE_ALLOC_PAGES_ADD = VirNodeAllocPagesFlags(C.VIR_NODE_ALLOC_PAGES_ADD)
	VIR_NODE_ALLOC_PAGES_SET = VirNodeAllocPagesFlags(C.VIR_NODE_ALLOC_PAGES_SET)
)

type VirNodeSuspendTarget int

const (
	VIR_NODE_SUSPEND_TARGET_MEM    = VirNodeSuspendTarget(C.VIR_NODE_SUSPEND_TARGET_MEM)
	VIR_NODE_SUSPEND_TARGET_DISK   = VirNodeSuspendTarget(C.VIR_NODE_SUSPEND_TARGET_DISK)
	VIR_NODE_SUSPEND_TARGET_HYBRID = VirNodeSuspendTarget(C.VIR_NODE_SUSPEND_TARGET_HYBRID)
)

type VirNodeGetCPUStatsAllCPUs int

const (
	VIR_NODE_CPU_STATS_ALL_CPUS = VirNodeGetCPUStatsAllCPUs(C.VIR_NODE_CPU_STATS_ALL_CPUS)
)

const (
	VIR_NODE_MEMORY_STATS_ALL_CELLS = int(C.VIR_NODE_MEMORY_STATS_ALL_CELLS)
)

type VirConnectCredentialType int

const (
	VIR_CRED_USERNAME     = VirConnectCredentialType(C.VIR_CRED_USERNAME)
	VIR_CRED_AUTHNAME     = VirConnectCredentialType(C.VIR_CRED_AUTHNAME)
	VIR_CRED_LANGUAGE     = VirConnectCredentialType(C.VIR_CRED_LANGUAGE)
	VIR_CRED_CNONCE       = VirConnectCredentialType(C.VIR_CRED_CNONCE)
	VIR_CRED_PASSPHRASE   = VirConnectCredentialType(C.VIR_CRED_PASSPHRASE)
	VIR_CRED_ECHOPROMPT   = VirConnectCredentialType(C.VIR_CRED_ECHOPROMPT)
	VIR_CRED_NOECHOPROMPT = VirConnectCredentialType(C.VIR_CRED_NOECHOPROMPT)
	VIR_CRED_REALM        = VirConnectCredentialType(C.VIR_CRED_REALM)
	VIR_CRED_EXTERNAL     = VirConnectCredentialType(C.VIR_CRED_EXTERNAL)
)

type VirConnection struct {
	ptr C.virConnectPtr
}

type VirNodeInfo struct {
	Model   string
	Memory  uint64
	Cpus    uint
	MHz     uint
	Nodes   uint
	Sockets uint
	Cores   uint
	Threads uint
}

// Additional data associated to the connection.
type virConnectionData struct {
	errCallbackId   *int
	closeCallbackId *int
}

var connections map[C.virConnectPtr]*virConnectionData
var connectionsLock sync.RWMutex

func init() {
	connections = make(map[C.virConnectPtr]*virConnectionData)
}

func saveConnectionData(c *VirConnection, d *virConnectionData) {
	if c.ptr == nil {
		return // Or panic?
	}
	connectionsLock.Lock()
	defer connectionsLock.Unlock()
	connections[c.ptr] = d
}

func getConnectionData(c *VirConnection) *virConnectionData {
	connectionsLock.RLock()
	d := connections[c.ptr]
	connectionsLock.RUnlock()
	if d != nil {
		return d
	}
	d = &virConnectionData{}
	saveConnectionData(c, d)
	return d
}

func releaseConnectionData(c *VirConnection) {
	if c.ptr == nil {
		return
	}
	connectionsLock.Lock()
	defer connectionsLock.Unlock()
	delete(connections, c.ptr)
}

func GetVersion() (uint32, error) {
	var version C.ulong
	if err := C.virGetVersion(&version, nil, nil); err < 0 {
		return 0, GetLastError()
	}
	return uint32(version), nil
}

func NewVirConnection(uri string) (*VirConnection, error) {
	var cUri *C.char
	if uri != "" {
		cUri = C.CString(uri)
		defer C.free(unsafe.Pointer(cUri))
	}
	ptr := C.virConnectOpen(cUri)
	if ptr == nil {
		return nil, GetLastError()
	}
	return &VirConnection{ptr: ptr}, nil
}

func NewVirConnectionWithAuth(uri string, username string, password string) (*VirConnection, error) {
	var cUri *C.char

	authMechs := C.authMechs()
	defer C.free(unsafe.Pointer(authMechs))
	cUsername := C.CString(username)
	defer C.free(unsafe.Pointer(cUsername))
	cPassword := C.CString(password)
	defer C.free(unsafe.Pointer(cPassword))
	cbData := C.authData(cUsername, C.uint(len(username)), cPassword, C.uint(len(password)))
	defer C.free(unsafe.Pointer(cbData))

	auth := C.virConnectAuth{
		credtype:  authMechs,
		ncredtype: C.uint(2),
		cb:        C.virConnectAuthCallbackPtr(unsafe.Pointer(C.authCb)),
		cbdata:    unsafe.Pointer(cbData),
	}

	if uri != "" {
		cUri = C.CString(uri)
		defer C.free(unsafe.Pointer(cUri))
	}
	ptr := C.virConnectOpenAuth(cUri, (*C.struct__virConnectAuth)(unsafe.Pointer(&auth)), C.uint(0))
	if ptr == nil {
		return nil, GetLastError()
	}
	return &VirConnection{ptr: ptr}, nil
}

func NewVirConnectionReadOnly(uri string) (*VirConnection, error) {
	var cUri *C.char
	if uri != "" {
		cUri = C.CString(uri)
		defer C.free(unsafe.Pointer(cUri))
	}
	ptr := C.virConnectOpenReadOnly(cUri)
	if ptr == nil {
		return nil, GetLastError()
	}
	return &VirConnection{ptr: ptr}, nil
}

func (c *VirConnection) CloseConnection() (int, error) {
	result := int(C.virConnectClose(c.ptr))
	if result == -1 {
		return result, GetLastError()
	}
	if result == 0 {
		// No more reference to this connection, release data.
		releaseConnectionData(c)
	}
	return result, nil
}

type CloseCallback func(conn VirConnection, reason int, opaque func())
type closeContext struct {
	cb CloseCallback
	f  func()
}

// Register a close callback for the given destination. Only one
// callback per connection is allowed. Setting a callback will remove
// the previous one.
func (c *VirConnection) RegisterCloseCallback(cb CloseCallback, opaque func()) error {
	c.UnregisterCloseCallback()
	context := &closeContext{
		cb: cb,
		f:  opaque,
	}
	goCallbackId := registerCallbackId(context)
	callbackPtr := unsafe.Pointer(C.closeCallback_cgo)
	res := C.virConnectRegisterCloseCallback_cgo(c.ptr, C.virConnectCloseFunc(callbackPtr), C.long(goCallbackId))
	if res != 0 {
		freeCallbackId(goCallbackId)
		return GetLastError()
	}
	connData := getConnectionData(c)
	connData.closeCallbackId = &goCallbackId
	return nil
}

func (c *VirConnection) UnregisterCloseCallback() error {
	connData := getConnectionData(c)
	if connData.closeCallbackId == nil {
		return nil
	}
	callbackPtr := unsafe.Pointer(C.closeCallback_cgo)
	res := C.virConnectUnregisterCloseCallback(c.ptr, C.virConnectCloseFunc(callbackPtr))
	if res != 0 {
		return GetLastError()
	}
	connData.closeCallbackId = nil
	return nil
}

//export closeCallback
func closeCallback(conn C.virConnectPtr, reason int, goCallbackId int) {
	ctx := getCallbackId(goCallbackId)
	switch cctx := ctx.(type) {
	case *closeContext:
		cctx.cb(VirConnection{ptr: conn}, reason, cctx.f)
	default:
		panic("Inappropriate callback type called")
	}
}

func (c *VirConnection) GetCapabilities() (string, error) {
	str := C.virConnectGetCapabilities(c.ptr)
	if str == nil {
		return "", GetLastError()
	}
	capabilities := C.GoString(str)
	C.free(unsafe.Pointer(str))
	return capabilities, nil
}

func (c *VirConnection) GetNodeInfo() (*VirNodeInfo, error) {
	var cinfo C.virNodeInfo
	result := C.virNodeGetInfo(c.ptr, &cinfo)
	if result == -1 {
		return nil, GetLastError()
	}
	return &VirNodeInfo{
		Model:   C.GoString((*C.char)(unsafe.Pointer(&cinfo.model[0]))),
		Memory:  uint64(cinfo.memory),
		Cpus:    uint(cinfo.cpus),
		MHz:     uint(cinfo.mhz),
		Nodes:   uint(cinfo.nodes),
		Sockets: uint(cinfo.sockets),
		Cores:   uint(cinfo.cores),
		Threads: uint(cinfo.threads),
	}, nil
}

func (c *VirConnection) GetHostname() (string, error) {
	str := C.virConnectGetHostname(c.ptr)
	if str == nil {
		return "", GetLastError()
	}
	hostname := C.GoString(str)
	C.free(unsafe.Pointer(str))
	return hostname, nil
}

func (c *VirConnection) GetLibVersion() (uint32, error) {
	var version C.ulong
	if err := C.virConnectGetLibVersion(c.ptr, &version); err < 0 {
		return 0, GetLastError()
	}
	return uint32(version), nil
}

func (c *VirConnection) GetType() (string, error) {
	str := C.virConnectGetType(c.ptr)
	if str == nil {
		return "", GetLastError()
	}
	hypDriver := C.GoString(str)
	return hypDriver, nil
}

func (c *VirConnection) IsAlive() (bool, error) {
	result := C.virConnectIsAlive(c.ptr)
	if result == -1 {
		return false, GetLastError()
	}
	if result == 1 {
		return true, nil
	}
	return false, nil
}

func (c *VirConnection) IsEncrypted() (bool, error) {
	result := C.virConnectIsEncrypted(c.ptr)
	if result == -1 {
		return false, GetLastError()
	}
	if result == 1 {
		return true, nil
	}
	return false, nil
}

func (c *VirConnection) IsSecure() (bool, error) {
	result := C.virConnectIsSecure(c.ptr)
	if result == -1 {
		return false, GetLastError()
	}
	if result == 1 {
		return true, nil
	}
	return false, nil
}

func (c *VirConnection) ListDefinedDomains() ([]string, error) {
	var names [1024](*C.char)
	namesPtr := unsafe.Pointer(&names)
	numDomains := C.virConnectListDefinedDomains(
		c.ptr,
		(**C.char)(namesPtr),
		1024)
	if numDomains == -1 {
		return nil, GetLastError()
	}
	goNames := make([]string, numDomains)
	for k := 0; k < int(numDomains); k++ {
		goNames[k] = C.GoString(names[k])
		C.free(unsafe.Pointer(names[k]))
	}
	return goNames, nil
}

func (c *VirConnection) ListDomains() ([]uint32, error) {
	var cDomainsIds [512](uint32)
	cDomainsPointer := unsafe.Pointer(&cDomainsIds)
	numDomains := C.virConnectListDomains(c.ptr, (*C.int)(cDomainsPointer), 512)
	if numDomains == -1 {
		return nil, GetLastError()
	}

	return cDomainsIds[:numDomains], nil
}

func (c *VirConnection) ListInterfaces() ([]string, error) {
	const maxIfaces = 1024
	var names [maxIfaces](*C.char)
	namesPtr := unsafe.Pointer(&names)
	numIfaces := C.virConnectListInterfaces(
		c.ptr,
		(**C.char)(namesPtr),
		maxIfaces)
	if numIfaces == -1 {
		return nil, GetLastError()
	}
	goNames := make([]string, numIfaces)
	for k := 0; k < int(numIfaces); k++ {
		goNames[k] = C.GoString(names[k])
		C.free(unsafe.Pointer(names[k]))
	}
	return goNames, nil
}

func (c *VirConnection) ListNetworks() ([]string, error) {
	const maxNets = 1024
	var names [maxNets](*C.char)
	namesPtr := unsafe.Pointer(&names)
	numNetworks := C.virConnectListNetworks(
		c.ptr,
		(**C.char)(namesPtr),
		maxNets)
	if numNetworks == -1 {
		return nil, GetLastError()
	}
	goNames := make([]string, numNetworks)
	for k := 0; k < int(numNetworks); k++ {
		goNames[k] = C.GoString(names[k])
		C.free(unsafe.Pointer(names[k]))
	}
	return goNames, nil
}

func (c *VirConnection) ListNWFilters() ([]string, error) {
	const maxFilters = 1024
	var names [maxFilters](*C.char)
	namesPtr := unsafe.Pointer(&names)
	numNWFilters := C.virConnectListNWFilters(
		c.ptr,
		(**C.char)(namesPtr),
		maxFilters)
	if numNWFilters == -1 {
		return nil, GetLastError()
	}
	goNames := make([]string, numNWFilters)
	for k := 0; k < int(numNWFilters); k++ {
		goNames[k] = C.GoString(names[k])
		C.free(unsafe.Pointer(names[k]))
	}
	return goNames, nil
}

func (c *VirConnection) ListStoragePools() ([]string, error) {
	const maxPools = 1024
	var names [maxPools](*C.char)
	namesPtr := unsafe.Pointer(&names)
	numStoragePools := C.virConnectListStoragePools(
		c.ptr,
		(**C.char)(namesPtr),
		maxPools)
	if numStoragePools == -1 {
		return nil, GetLastError()
	}
	goNames := make([]string, numStoragePools)
	for k := 0; k < int(numStoragePools); k++ {
		goNames[k] = C.GoString(names[k])
		C.free(unsafe.Pointer(names[k]))
	}
	return goNames, nil
}

func (c *VirConnection) ListSecrets() ([]string, error) {
	const maxSecrets = 1024
	var uuids [maxSecrets](*C.char)
	uuidsPtr := unsafe.Pointer(&uuids)
	numSecrets := C.virConnectListSecrets(
		c.ptr,
		(**C.char)(uuidsPtr),
		maxSecrets)
	if numSecrets == -1 {
		return nil, GetLastError()
	}
	goUuids := make([]string, numSecrets)
	for k := 0; k < int(numSecrets); k++ {
		goUuids[k] = C.GoString(uuids[k])
		C.free(unsafe.Pointer(uuids[k]))
	}
	return goUuids, nil
}

func (c *VirConnection) ListDevices(cap string, flags uint32) ([]string, error) {
	ccap := C.CString(cap)
	defer C.free(ccap)
	const maxNodeDevices = 1024
	var uuids [maxNodeDevices](*C.char)
	uuidsPtr := unsafe.Pointer(&uuids)
	numNodeDevices := C.virNodeListDevices(
		c.ptr, ccap,
		(**C.char)(uuidsPtr),
		maxNodeDevices, C.uint(flags))
	if numNodeDevices == -1 {
		return nil, GetLastError()
	}
	goUuids := make([]string, numNodeDevices)
	for k := 0; k < int(numNodeDevices); k++ {
		goUuids[k] = C.GoString(uuids[k])
		C.free(unsafe.Pointer(uuids[k]))
	}
	return goUuids, nil
}

func (c *VirConnection) LookupDomainById(id uint32) (*Domain, error) {
	ptr := C.virDomainLookupByID(c.ptr, C.int(id))
	if ptr == nil {
		return nil, GetLastError()
	}
	return &Domain{ptr: ptr}, nil
}

func (c *VirConnection) LookupDomainByName(id string) (*Domain, error) {
	cName := C.CString(id)
	defer C.free(unsafe.Pointer(cName))
	ptr := C.virDomainLookupByName(c.ptr, cName)
	if ptr == nil {
		return nil, GetLastError()
	}
	return &Domain{ptr: ptr}, nil
}

func (c *VirConnection) LookupDomainByUUIDString(uuid string) (*Domain, error) {
	cUuid := C.CString(uuid)
	defer C.free(unsafe.Pointer(cUuid))
	ptr := C.virDomainLookupByUUIDString(c.ptr, cUuid)
	if ptr == nil {
		return nil, GetLastError()
	}
	return &Domain{ptr: ptr}, nil
}

func (c *VirConnection) LookupDomainByUUID(uuid []byte) (*Domain, error) {
	if len(uuid) != C.VIR_UUID_BUFLEN {
		return nil, fmt.Errorf("UUID must be exactly %d bytes in size",
			int(C.VIR_UUID_BUFLEN))
	}
	cUuid := C.CBytes(uuid)
	defer C.free(unsafe.Pointer(cUuid))
	ptr := C.virDomainLookupByUUID(c.ptr, (*C.uchar)(cUuid))
	if ptr == nil {
		return nil, GetLastError()
	}
	return &Domain{ptr: ptr}, nil
}

func (c *VirConnection) DomainCreateXML(xmlConfig string, flags DomainCreateFlags) (*Domain, error) {
	cXml := C.CString(string(xmlConfig))
	defer C.free(unsafe.Pointer(cXml))
	ptr := C.virDomainCreateXML(c.ptr, cXml, C.uint(flags))
	if ptr == nil {
		return nil, GetLastError()
	}
	return &Domain{ptr: ptr}, nil
}

func (c *VirConnection) DomainCreateXMLWithFiles(xmlConfig string, files []os.File, flags DomainCreateFlags) (*Domain, error) {
	cXml := C.CString(string(xmlConfig))
	defer C.free(unsafe.Pointer(cXml))
	cfiles := make([]C.int, len(files))
	for i := 0; i < len(files); i++ {
		cfiles[i] = C.int(files[i].Fd())
	}
	ptr := C.virDomainCreateXMLWithFiles(c.ptr, cXml, C.uint(len(files)), (&cfiles[0]), C.uint(flags))
	if ptr == nil {
		return nil, GetLastError()
	}
	return &Domain{ptr: ptr}, nil
}

func (c *VirConnection) DomainDefineXML(xmlConfig string) (*Domain, error) {
	cXml := C.CString(string(xmlConfig))
	defer C.free(unsafe.Pointer(cXml))
	ptr := C.virDomainDefineXML(c.ptr, cXml)
	if ptr == nil {
		return nil, GetLastError()
	}
	return &Domain{ptr: ptr}, nil
}

func (c *VirConnection) DomainDefineXMLFlags(xmlConfig string, flags DomainDefineFlags) (*Domain, error) {
	cXml := C.CString(string(xmlConfig))
	defer C.free(unsafe.Pointer(cXml))
	ptr := C.virDomainDefineXMLFlags(c.ptr, cXml, C.uint(flags))
	if ptr == nil {
		return nil, GetLastError()
	}
	return &Domain{ptr: ptr}, nil
}

func (c *VirConnection) ListDefinedInterfaces() ([]string, error) {
	const maxIfaces = 1024
	var names [maxIfaces](*C.char)
	namesPtr := unsafe.Pointer(&names)
	numIfaces := C.virConnectListDefinedInterfaces(
		c.ptr,
		(**C.char)(namesPtr),
		maxIfaces)
	if numIfaces == -1 {
		return nil, GetLastError()
	}
	goNames := make([]string, numIfaces)
	for k := 0; k < int(numIfaces); k++ {
		goNames[k] = C.GoString(names[k])
		C.free(unsafe.Pointer(names[k]))
	}
	return goNames, nil
}

func (c *VirConnection) ListDefinedNetworks() ([]string, error) {
	const maxNets = 1024
	var names [maxNets](*C.char)
	namesPtr := unsafe.Pointer(&names)
	numNetworks := C.virConnectListDefinedNetworks(
		c.ptr,
		(**C.char)(namesPtr),
		maxNets)
	if numNetworks == -1 {
		return nil, GetLastError()
	}
	goNames := make([]string, numNetworks)
	for k := 0; k < int(numNetworks); k++ {
		goNames[k] = C.GoString(names[k])
		C.free(unsafe.Pointer(names[k]))
	}
	return goNames, nil
}

func (c *VirConnection) ListDefinedStoragePools() ([]string, error) {
	const maxPools = 1024
	var names [maxPools](*C.char)
	namesPtr := unsafe.Pointer(&names)
	numStoragePools := C.virConnectListDefinedStoragePools(
		c.ptr,
		(**C.char)(namesPtr),
		maxPools)
	if numStoragePools == -1 {
		return nil, GetLastError()
	}
	goNames := make([]string, numStoragePools)
	for k := 0; k < int(numStoragePools); k++ {
		goNames[k] = C.GoString(names[k])
		C.free(unsafe.Pointer(names[k]))
	}
	return goNames, nil
}

func (c *VirConnection) NumOfDefinedDomains() (int, error) {
	result := int(C.virConnectNumOfDefinedDomains(c.ptr))
	if result == -1 {
		return 0, GetLastError()
	}
	return result, nil
}

func (c *VirConnection) NumOfDefinedInterfaces() (int, error) {
	result := int(C.virConnectNumOfDefinedInterfaces(c.ptr))
	if result == -1 {
		return 0, GetLastError()
	}
	return result, nil
}

func (c *VirConnection) NumOfDefinedNetworks() (int, error) {
	result := int(C.virConnectNumOfDefinedNetworks(c.ptr))
	if result == -1 {
		return 0, GetLastError()
	}
	return result, nil
}

func (c *VirConnection) NumOfDefinedStoragePools() (int, error) {
	result := int(C.virConnectNumOfDefinedStoragePools(c.ptr))
	if result == -1 {
		return 0, GetLastError()
	}
	return result, nil
}

func (c *VirConnection) NumOfDomains() (int, error) {
	result := int(C.virConnectNumOfDomains(c.ptr))
	if result == -1 {
		return 0, GetLastError()
	}
	return result, nil
}

func (c *VirConnection) NumOfStoragePools() (int, error) {
	result := int(C.virConnectNumOfStoragePools(c.ptr))
	if result == -1 {
		return 0, GetLastError()
	}
	return result, nil
}

func (c *VirConnection) NumOfInterfaces() (int, error) {
	result := int(C.virConnectNumOfInterfaces(c.ptr))
	if result == -1 {
		return 0, GetLastError()
	}
	return result, nil
}

func (c *VirConnection) NumOfNetworks() (int, error) {
	result := int(C.virConnectNumOfNetworks(c.ptr))
	if result == -1 {
		return 0, GetLastError()
	}
	return result, nil
}

func (c *VirConnection) NumOfNWFilters() (int, error) {
	result := int(C.virConnectNumOfNWFilters(c.ptr))
	if result == -1 {
		return 0, GetLastError()
	}
	return result, nil
}

func (c *VirConnection) NumOfSecrets() (int, error) {
	result := int(C.virConnectNumOfSecrets(c.ptr))
	if result == -1 {
		return 0, GetLastError()
	}
	return result, nil
}

func (c *VirConnection) NumOfDevices(cap string, flags uint32) (int, error) {
	ccap := C.CString(cap)
	defer C.free(ccap)
	result := int(C.virNodeNumOfDevices(c.ptr, ccap, C.uint(flags)))
	if result == -1 {
		return 0, GetLastError()
	}
	return result, nil
}

func (c *VirConnection) NetworkDefineXML(xmlConfig string) (*Network, error) {
	cXml := C.CString(string(xmlConfig))
	defer C.free(unsafe.Pointer(cXml))
	ptr := C.virNetworkDefineXML(c.ptr, cXml)
	if ptr == nil {
		return nil, GetLastError()
	}
	return &Network{ptr: ptr}, nil
}

func (c *VirConnection) NetworkCreateXML(xmlConfig string) (*Network, error) {
	cXml := C.CString(string(xmlConfig))
	defer C.free(unsafe.Pointer(cXml))
	ptr := C.virNetworkCreateXML(c.ptr, cXml)
	if ptr == nil {
		return nil, GetLastError()
	}
	return &Network{ptr: ptr}, nil
}

func (c *VirConnection) LookupNetworkByName(name string) (*Network, error) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	ptr := C.virNetworkLookupByName(c.ptr, cName)
	if ptr == nil {
		return nil, GetLastError()
	}
	return &Network{ptr: ptr}, nil
}

func (c *VirConnection) LookupNetworkByUUIDString(uuid string) (*Network, error) {
	cUuid := C.CString(uuid)
	defer C.free(unsafe.Pointer(cUuid))
	ptr := C.virNetworkLookupByUUIDString(c.ptr, cUuid)
	if ptr == nil {
		return nil, GetLastError()
	}
	return &Network{ptr: ptr}, nil
}

func (c *VirConnection) LookupNetworkByUUID(uuid []byte) (*Network, error) {
	if len(uuid) != C.VIR_UUID_BUFLEN {
		return nil, fmt.Errorf("UUID must be exactly %d bytes in size",
			int(C.VIR_UUID_BUFLEN))
	}
	cUuid := C.CBytes(uuid)
	defer C.free(unsafe.Pointer(cUuid))
	ptr := C.virNetworkLookupByUUID(c.ptr, (*C.uchar)(cUuid))
	if ptr == nil {
		return nil, GetLastError()
	}
	return &Network{ptr: ptr}, nil
}

func (c *VirConnection) SetKeepAlive(interval int, count uint) error {
	res := int(C.virConnectSetKeepAlive(c.ptr, C.int(interval), C.uint(count)))
	switch res {
	case 0:
		return nil
	default:
		return GetLastError()
	}
}

func (c *VirConnection) GetSysinfo(flags uint) (string, error) {
	cStr := C.virConnectGetSysinfo(c.ptr, C.uint(flags))
	if cStr == nil {
		return "", GetLastError()
	}
	info := C.GoString(cStr)
	C.free(unsafe.Pointer(cStr))
	return info, nil
}

func (c *VirConnection) GetURI() (string, error) {
	cStr := C.virConnectGetURI(c.ptr)
	if cStr == nil {
		return "", GetLastError()
	}
	uri := C.GoString(cStr)
	C.free(unsafe.Pointer(cStr))
	return uri, nil
}

func (c *VirConnection) GetMaxVcpus(typeAttr string) (int, error) {
	var cTypeAttr *C.char
	if typeAttr != "" {
		cTypeAttr = C.CString(typeAttr)
		defer C.free(unsafe.Pointer(cTypeAttr))
	}
	result := int(C.virConnectGetMaxVcpus(c.ptr, cTypeAttr))
	if result == -1 {
		return 0, GetLastError()
	}
	return result, nil
}

func (c *VirConnection) InterfaceDefineXML(xmlConfig string, flags uint32) (*VirInterface, error) {
	cXml := C.CString(string(xmlConfig))
	defer C.free(unsafe.Pointer(cXml))
	ptr := C.virInterfaceDefineXML(c.ptr, cXml, C.uint(flags))
	if ptr == nil {
		return nil, GetLastError()
	}
	return &VirInterface{ptr: ptr}, nil
}

func (c *VirConnection) LookupInterfaceByName(name string) (*VirInterface, error) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	ptr := C.virInterfaceLookupByName(c.ptr, cName)
	if ptr == nil {
		return nil, GetLastError()
	}
	return &VirInterface{ptr: ptr}, nil
}

func (c *VirConnection) LookupInterfaceByMACString(mac string) (*VirInterface, error) {
	cName := C.CString(mac)
	defer C.free(unsafe.Pointer(cName))
	ptr := C.virInterfaceLookupByMACString(c.ptr, cName)
	if ptr == nil {
		return nil, GetLastError()
	}
	return &VirInterface{ptr: ptr}, nil
}

func (c *VirConnection) StoragePoolDefineXML(xmlConfig string, flags uint32) (*StoragePool, error) {
	cXml := C.CString(string(xmlConfig))
	defer C.free(unsafe.Pointer(cXml))
	ptr := C.virStoragePoolDefineXML(c.ptr, cXml, C.uint(flags))
	if ptr == nil {
		return nil, GetLastError()
	}
	return &StoragePool{ptr: ptr}, nil
}

func (c *VirConnection) StoragePoolCreateXML(xmlConfig string, flags uint32) (*StoragePool, error) {
	cXml := C.CString(string(xmlConfig))
	defer C.free(unsafe.Pointer(cXml))
	ptr := C.virStoragePoolCreateXML(c.ptr, cXml, C.uint(flags))
	if ptr == nil {
		return nil, GetLastError()
	}
	return &StoragePool{ptr: ptr}, nil
}

func (c *VirConnection) LookupStoragePoolByName(name string) (*StoragePool, error) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	ptr := C.virStoragePoolLookupByName(c.ptr, cName)
	if ptr == nil {
		return nil, GetLastError()
	}
	return &StoragePool{ptr: ptr}, nil
}

func (c *VirConnection) LookupStoragePoolByUUIDString(uuid string) (*StoragePool, error) {
	cUuid := C.CString(uuid)
	defer C.free(unsafe.Pointer(cUuid))
	ptr := C.virStoragePoolLookupByUUIDString(c.ptr, cUuid)
	if ptr == nil {
		return nil, GetLastError()
	}
	return &StoragePool{ptr: ptr}, nil
}

func (c *VirConnection) LookupStoragePoolByUUID(uuid []byte) (*StoragePool, error) {
	if len(uuid) != C.VIR_UUID_BUFLEN {
		return nil, fmt.Errorf("UUID must be exactly %d bytes in size",
			int(C.VIR_UUID_BUFLEN))
	}
	cUuid := C.CBytes(uuid)
	defer C.free(unsafe.Pointer(cUuid))
	ptr := C.virStoragePoolLookupByUUID(c.ptr, (*C.uchar)(cUuid))
	if ptr == nil {
		return nil, GetLastError()
	}
	return &StoragePool{ptr: ptr}, nil
}

func (c *VirConnection) NWFilterDefineXML(xmlConfig string) (*NWFilter, error) {
	cXml := C.CString(string(xmlConfig))
	defer C.free(unsafe.Pointer(cXml))
	ptr := C.virNWFilterDefineXML(c.ptr, cXml)
	if ptr == nil {
		return nil, GetLastError()
	}
	return &NWFilter{ptr: ptr}, nil
}

func (c *VirConnection) LookupNWFilterByName(name string) (*NWFilter, error) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	ptr := C.virNWFilterLookupByName(c.ptr, cName)
	if ptr == nil {
		return nil, GetLastError()
	}
	return &NWFilter{ptr: ptr}, nil
}

func (c *VirConnection) LookupNWFilterByUUIDString(uuid string) (*NWFilter, error) {
	cUuid := C.CString(uuid)
	defer C.free(unsafe.Pointer(cUuid))
	ptr := C.virNWFilterLookupByUUIDString(c.ptr, cUuid)
	if ptr == nil {
		return nil, GetLastError()
	}
	return &NWFilter{ptr: ptr}, nil
}

func (c *VirConnection) LookupNWFilterByUUID(uuid []byte) (*NWFilter, error) {
	if len(uuid) != C.VIR_UUID_BUFLEN {
		return nil, fmt.Errorf("UUID must be exactly %d bytes in size",
			int(C.VIR_UUID_BUFLEN))
	}
	cUuid := C.CBytes(uuid)
	defer C.free(unsafe.Pointer(cUuid))
	ptr := C.virNWFilterLookupByUUID(c.ptr, (*C.uchar)(cUuid))
	if ptr == nil {
		return nil, GetLastError()
	}
	return &NWFilter{ptr: ptr}, nil
}

func (c *VirConnection) LookupStorageVolByKey(key string) (*StorageVol, error) {
	cKey := C.CString(key)
	defer C.free(unsafe.Pointer(cKey))
	ptr := C.virStorageVolLookupByKey(c.ptr, cKey)
	if ptr == nil {
		return nil, GetLastError()
	}
	return &StorageVol{ptr: ptr}, nil
}

func (c *VirConnection) LookupStorageVolByPath(path string) (*StorageVol, error) {
	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))
	ptr := C.virStorageVolLookupByPath(c.ptr, cPath)
	if ptr == nil {
		return nil, GetLastError()
	}
	return &StorageVol{ptr: ptr}, nil
}

func (c *VirConnection) SecretDefineXML(xmlConfig string, flags uint32) (*Secret, error) {
	cXml := C.CString(string(xmlConfig))
	defer C.free(unsafe.Pointer(cXml))
	ptr := C.virSecretDefineXML(c.ptr, cXml, C.uint(flags))
	if ptr == nil {
		return nil, GetLastError()
	}
	return &Secret{ptr: ptr}, nil
}

func (c *VirConnection) LookupSecretByUUID(uuid []byte) (*Secret, error) {
	if len(uuid) != C.VIR_UUID_BUFLEN {
		return nil, fmt.Errorf("UUID must be exactly %d bytes in size",
			int(C.VIR_UUID_BUFLEN))
	}
	cUuid := C.CBytes(uuid)
	defer C.free(unsafe.Pointer(cUuid))
	ptr := C.virSecretLookupByUUID(c.ptr, (*C.uchar)(cUuid))
	if ptr == nil {
		return nil, GetLastError()
	}
	return &Secret{ptr: ptr}, nil
}

func (c *VirConnection) LookupSecretByUUIDString(uuid string) (*Secret, error) {
	cUuid := C.CString(uuid)
	defer C.free(unsafe.Pointer(cUuid))
	ptr := C.virSecretLookupByUUIDString(c.ptr, cUuid)
	if ptr == nil {
		return nil, GetLastError()
	}
	return &Secret{ptr: ptr}, nil
}

func (c *VirConnection) LookupSecretByUsage(usageType int, usageID string) (*Secret, error) {
	cUsageID := C.CString(usageID)
	defer C.free(unsafe.Pointer(cUsageID))
	ptr := C.virSecretLookupByUsage(c.ptr, C.int(usageType), cUsageID)
	if ptr == nil {
		return nil, GetLastError()
	}
	return &Secret{ptr: ptr}, nil
}

func (c *VirConnection) LookupDeviceByName(id string) (*VirNodeDevice, error) {
	cName := C.CString(id)
	defer C.free(unsafe.Pointer(cName))
	ptr := C.virNodeDeviceLookupByName(c.ptr, cName)
	if ptr == nil {
		return nil, GetLastError()
	}
	return &VirNodeDevice{ptr: ptr}, nil
}

func (c *VirConnection) LookupDeviceSCSIHostByWWN(wwnn, wwpn string, flags uint32) (*VirNodeDevice, error) {
	cWwnn := C.CString(wwnn)
	cWwpn := C.CString(wwpn)
	defer C.free(unsafe.Pointer(cWwnn))
	defer C.free(unsafe.Pointer(cWwpn))
	ptr := C.virNodeDeviceLookupSCSIHostByWWN(c.ptr, cWwnn, cWwpn, C.uint(flags))
	if ptr == nil {
		return nil, GetLastError()
	}
	return &VirNodeDevice{ptr: ptr}, nil
}

func (c *VirConnection) DeviceCreateXML(xmlConfig string, flags uint32) (*VirNodeDevice, error) {
	cXml := C.CString(string(xmlConfig))
	defer C.free(unsafe.Pointer(cXml))
	ptr := C.virNodeDeviceCreateXML(c.ptr, cXml, C.uint(flags))
	if ptr == nil {
		return nil, GetLastError()
	}
	return &VirNodeDevice{ptr: ptr}, nil
}

func (c *VirConnection) ListAllInterfaces(flags uint32) ([]VirInterface, error) {
	var cList *C.virInterfacePtr
	numIfaces := C.virConnectListAllInterfaces(c.ptr, (**C.virInterfacePtr)(&cList), C.uint(flags))
	if numIfaces == -1 {
		return nil, GetLastError()
	}
	hdr := reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(cList)),
		Len:  int(numIfaces),
		Cap:  int(numIfaces),
	}
	var ifaces []VirInterface
	slice := *(*[]C.virInterfacePtr)(unsafe.Pointer(&hdr))
	for _, ptr := range slice {
		ifaces = append(ifaces, VirInterface{ptr})
	}
	C.free(unsafe.Pointer(cList))
	return ifaces, nil
}

func (c *VirConnection) ListAllNetworks(flags VirConnectListAllNetworksFlags) ([]Network, error) {
	var cList *C.virNetworkPtr
	numNets := C.virConnectListAllNetworks(c.ptr, (**C.virNetworkPtr)(&cList), C.uint(flags))
	if numNets == -1 {
		return nil, GetLastError()
	}
	hdr := reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(cList)),
		Len:  int(numNets),
		Cap:  int(numNets),
	}
	var nets []Network
	slice := *(*[]C.virNetworkPtr)(unsafe.Pointer(&hdr))
	for _, ptr := range slice {
		nets = append(nets, Network{ptr})
	}
	C.free(unsafe.Pointer(cList))
	return nets, nil
}

func (c *VirConnection) ListAllDomains(flags VirConnectListAllDomainsFlags) ([]Domain, error) {
	var cList *C.virDomainPtr
	numDomains := C.virConnectListAllDomains(c.ptr, (**C.virDomainPtr)(&cList), C.uint(flags))
	if numDomains == -1 {
		return nil, GetLastError()
	}
	hdr := reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(cList)),
		Len:  int(numDomains),
		Cap:  int(numDomains),
	}
	var domains []Domain
	slice := *(*[]C.virDomainPtr)(unsafe.Pointer(&hdr))
	for _, ptr := range slice {
		domains = append(domains, Domain{ptr})
	}
	C.free(unsafe.Pointer(cList))
	return domains, nil
}

func (c *VirConnection) ListAllNWFilters(flags uint32) ([]NWFilter, error) {
	var cList *C.virNWFilterPtr
	numNWFilters := C.virConnectListAllNWFilters(c.ptr, (**C.virNWFilterPtr)(&cList), C.uint(flags))
	if numNWFilters == -1 {
		return nil, GetLastError()
	}
	hdr := reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(cList)),
		Len:  int(numNWFilters),
		Cap:  int(numNWFilters),
	}
	var filters []NWFilter
	slice := *(*[]C.virNWFilterPtr)(unsafe.Pointer(&hdr))
	for _, ptr := range slice {
		filters = append(filters, NWFilter{ptr})
	}
	C.free(unsafe.Pointer(cList))
	return filters, nil
}

func (c *VirConnection) ListAllStoragePools(flags VirConnectListAllStoragePoolsFlags) ([]StoragePool, error) {
	var cList *C.virStoragePoolPtr
	numPools := C.virConnectListAllStoragePools(c.ptr, (**C.virStoragePoolPtr)(&cList), C.uint(flags))
	if numPools == -1 {
		return nil, GetLastError()
	}
	hdr := reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(cList)),
		Len:  int(numPools),
		Cap:  int(numPools),
	}
	var pools []StoragePool
	slice := *(*[]C.virStoragePoolPtr)(unsafe.Pointer(&hdr))
	for _, ptr := range slice {
		pools = append(pools, StoragePool{ptr})
	}
	C.free(unsafe.Pointer(cList))
	return pools, nil
}

func (c *VirConnection) ListAllSecrets(flags VirConnectListAllSecretsFlags) ([]Secret, error) {
	var cList *C.virSecretPtr
	numPools := C.virConnectListAllSecrets(c.ptr, (**C.virSecretPtr)(&cList), C.uint(flags))
	if numPools == -1 {
		return nil, GetLastError()
	}
	hdr := reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(cList)),
		Len:  int(numPools),
		Cap:  int(numPools),
	}
	var pools []Secret
	slice := *(*[]C.virSecretPtr)(unsafe.Pointer(&hdr))
	for _, ptr := range slice {
		pools = append(pools, Secret{ptr})
	}
	C.free(unsafe.Pointer(cList))
	return pools, nil
}

func (c *VirConnection) ListAllNodeDevices(flags VirConnectListAllNodeDeviceFlags) ([]VirNodeDevice, error) {
	var cList *C.virNodeDevicePtr
	numPools := C.virConnectListAllNodeDevices(c.ptr, (**C.virNodeDevicePtr)(&cList), C.uint(flags))
	if numPools == -1 {
		return nil, GetLastError()
	}
	hdr := reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(cList)),
		Len:  int(numPools),
		Cap:  int(numPools),
	}
	var pools []VirNodeDevice
	slice := *(*[]C.virNodeDevicePtr)(unsafe.Pointer(&hdr))
	for _, ptr := range slice {
		pools = append(pools, VirNodeDevice{ptr})
	}
	C.free(unsafe.Pointer(cList))
	return pools, nil
}

func (c *VirConnection) InterfaceChangeBegin(flags uint) error {
	ret := C.virInterfaceChangeBegin(c.ptr, C.uint(flags))
	if ret == -1 {
		return GetLastError()
	}
	return nil
}

func (c *VirConnection) InterfaceChangeCommit(flags uint) error {
	ret := C.virInterfaceChangeCommit(c.ptr, C.uint(flags))
	if ret == -1 {
		return GetLastError()
	}
	return nil
}

func (c *VirConnection) InterfaceChangeRollback(flags uint) error {
	ret := C.virInterfaceChangeRollback(c.ptr, C.uint(flags))
	if ret == -1 {
		return GetLastError()
	}
	return nil
}

func (c *VirConnection) AllocPages(pageSizes map[int]int64, startCell int, cellCount uint, flags VirNodeAllocPagesFlags) (int, error) {
	cpages := make([]C.uint, len(pageSizes))
	ccounts := make([]C.ulonglong, len(pageSizes))

	i := 0
	for key, val := range pageSizes {
		cpages[i] = C.uint(key)
		ccounts[i] = C.ulonglong(val)
		i++
	}

	ret := C.virNodeAllocPages(c.ptr, C.uint(len(pageSizes)), (*C.uint)(unsafe.Pointer(&cpages)),
		(*C.ulonglong)(unsafe.Pointer(&ccounts)), C.int(startCell), C.uint(cellCount), C.uint(flags))
	if ret == -1 {
		return 0, GetLastError()
	}

	return int(ret), nil
}

func (c *VirConnection) GetCPUMap(flags uint32) (map[int]bool, uint, error) {
	var ccpumap *C.uchar
	var conline C.uint
	ret := C.virNodeGetCPUMap(c.ptr, &ccpumap, &conline, C.uint(flags))
	if ret == -1 {
		return map[int]bool{}, 0, GetLastError()
	}
	defer C.free(ccpumap)

	cpumapbytes := C.GoBytes(unsafe.Pointer(ccpumap), C.int(ret/8))

	cpumap := make(map[int]bool, 0)
	for i := 0; i < int(ret); i++ {
		idx := int(i / 8)
		val := byte(cpumapbytes[idx])
		shift := i % 8
		cpumap[i] = (val & (1 << uint(shift))) == 1
	}

	return cpumap, uint(conline), nil
}

type VirNodeCPUStats struct {
	KernelSet      bool
	Kernel         uint64
	UserSet        bool
	User           uint64
	IdleSet        bool
	Idle           uint64
	IowaitSet      bool
	Iowait         uint64
	IntrSet        bool
	Intr           uint64
	UtilizationSet bool
	Utilization    uint64
}

func (c *VirConnection) GetCPUStats(cpuNum int, flags uint32) (*VirNodeCPUStats, error) {
	var nparams C.int

	ret := C.virNodeGetCPUStats(c.ptr, C.int(cpuNum), nil, &nparams, C.uint(0))
	if ret == -1 {
		return nil, GetLastError()
	}

	params := make([]C.virNodeCPUStats, nparams)
	ret = C.virNodeGetCPUStats(c.ptr, C.int(cpuNum), (*C.virNodeCPUStats)(unsafe.Pointer(&params)), &nparams, C.uint(flags))
	if ret == -1 {
		return nil, GetLastError()
	}

	stats := &VirNodeCPUStats{}
	for i := 0; i < int(nparams); i++ {
		param := params[i]
		field := C.GoString((*C.char)(unsafe.Pointer(&param.field)))
		switch field {
		case C.VIR_NODE_CPU_STATS_KERNEL:
			stats.KernelSet = true
			stats.Kernel = uint64(param.value)
		case C.VIR_NODE_CPU_STATS_USER:
			stats.UserSet = true
			stats.User = uint64(param.value)
		case C.VIR_NODE_CPU_STATS_IDLE:
			stats.IdleSet = true
			stats.Idle = uint64(param.value)
		case C.VIR_NODE_CPU_STATS_IOWAIT:
			stats.IowaitSet = true
			stats.Iowait = uint64(param.value)
		case C.VIR_NODE_CPU_STATS_INTR:
			stats.IntrSet = true
			stats.Intr = uint64(param.value)
		case C.VIR_NODE_CPU_STATS_UTILIZATION:
			stats.UtilizationSet = true
			stats.Utilization = uint64(param.value)
		}
	}

	return stats, nil
}

func (c *VirConnection) GetCellsFreeMemory(startCell int, maxCells int) ([]uint64, error) {
	cmem := make([]C.ulonglong, maxCells)
	ret := C.virNodeGetCellsFreeMemory(c.ptr, (*C.ulonglong)(unsafe.Pointer(&cmem[0])), C.int(startCell), C.int(maxCells))
	if ret == -1 {
		return []uint64{}, GetLastError()
	}

	mem := make([]uint64, ret)
	for i := 0; i < int(ret); i++ {
		mem[i] = uint64(cmem[i])
	}

	return mem, nil
}

func (c *VirConnection) GetFreeMemory() (uint64, error) {
	ret := C.virNodeGetFreeMemory(c.ptr)
	if ret == 0 {
		return 0, GetLastError()
	}

	return (uint64)(ret), nil
}

func (c *VirConnection) GetFreePages(pageSizes []uint64, startCell int, maxCells uint, flags uint32) ([]uint64, error) {
	cpageSizes := make([]C.uint, len(pageSizes))
	ccounts := make([]C.ulonglong, len(pageSizes)*int(maxCells))

	for i := 0; i < len(pageSizes); i++ {
		cpageSizes[i] = C.uint(pageSizes[i])
	}

	ret := C.virNodeGetFreePages(c.ptr, C.uint(len(pageSizes)), (*C.uint)(unsafe.Pointer(&cpageSizes)), C.int(startCell),
		C.uint(maxCells), (*C.ulonglong)(unsafe.Pointer(&ccounts)), C.uint(flags))
	if ret == -1 {
		return []uint64{}, GetLastError()
	}

	counts := make([]uint64, ret)
	for i := 0; i < int(ret); i++ {
		counts[i] = uint64(ccounts[i])
	}

	return counts, nil
}

type VirNodeMemoryParameters struct {
	ShmPagesToScanSet      bool
	ShmPagesToScan         uint
	ShmSleepMillisecsSet   bool
	ShmSleepMillisecs      uint
	ShmPagesSharedSet      bool
	ShmPagesShared         uint64
	ShmPagesSharingSet     bool
	ShmPagesSharing        uint64
	ShmPagesUnsharedSet    bool
	ShmPagesUnshared       uint64
	ShmPagesVolatileSet    bool
	ShmPagesVolatile       uint64
	ShmFullScansSet        bool
	ShmFullScans           uint64
	ShmMergeAcrossNodesSet bool
	ShmMergeAcrossNodes    uint
}

func getMemoryParameterFieldInfo(params *VirNodeMemoryParameters) map[string]typedParamsFieldInfo {
	return map[string]typedParamsFieldInfo{
		C.VIR_NODE_MEMORY_SHARED_PAGES_TO_SCAN: typedParamsFieldInfo{
			set: &params.ShmPagesToScanSet,
			ui:  &params.ShmPagesToScan,
		},
		C.VIR_NODE_MEMORY_SHARED_SLEEP_MILLISECS: typedParamsFieldInfo{
			set: &params.ShmSleepMillisecsSet,
			ui:  &params.ShmSleepMillisecs,
		},
		C.VIR_NODE_MEMORY_SHARED_MERGE_ACROSS_NODES: typedParamsFieldInfo{
			set: &params.ShmMergeAcrossNodesSet,
			ui:  &params.ShmMergeAcrossNodes,
		},
		C.VIR_NODE_MEMORY_SHARED_PAGES_SHARED: typedParamsFieldInfo{
			set: &params.ShmPagesSharedSet,
			ul:  &params.ShmPagesShared,
		},
		C.VIR_NODE_MEMORY_SHARED_PAGES_SHARING: typedParamsFieldInfo{
			set: &params.ShmPagesSharingSet,
			ul:  &params.ShmPagesSharing,
		},
		C.VIR_NODE_MEMORY_SHARED_PAGES_UNSHARED: typedParamsFieldInfo{
			set: &params.ShmPagesUnsharedSet,
			ul:  &params.ShmPagesUnshared,
		},
		C.VIR_NODE_MEMORY_SHARED_PAGES_VOLATILE: typedParamsFieldInfo{
			set: &params.ShmPagesVolatileSet,
			ul:  &params.ShmPagesVolatile,
		},
		C.VIR_NODE_MEMORY_SHARED_FULL_SCANS: typedParamsFieldInfo{
			set: &params.ShmFullScansSet,
			ul:  &params.ShmFullScans,
		},
	}
}

func (c *VirConnection) GetMemoryParameters(flags uint32) (*VirNodeMemoryParameters, error) {
	params := &VirNodeMemoryParameters{}
	info := getMemoryParameterFieldInfo(params)

	var nparams C.int

	ret := C.virNodeGetMemoryParameters(c.ptr, nil, &nparams, C.uint(0))
	if ret == -1 {
		return nil, GetLastError()
	}

	cparams := make([]C.virTypedParameter, nparams)
	ret = C.virNodeGetMemoryParameters(c.ptr, (*C.virTypedParameter)(unsafe.Pointer(&cparams[0])), &nparams, C.uint(flags))
	if ret == -1 {
		return nil, GetLastError()
	}

	defer C.virTypedParamsClear((*C.virTypedParameter)(unsafe.Pointer(&cparams[0])), nparams)

	err := typedParamsUnpack(cparams, info)
	if err != nil {
		return nil, err
	}

	return params, nil
}

type VirNodeMemoryStats struct {
	TotalSet   bool
	Total      uint64
	FreeSet    bool
	Free       uint64
	BuffersSet bool
	Buffers    uint64
	CachedSet  bool
	Cached     uint64
}

func (c *VirConnection) GetMemoryStats(cellNum int, flags uint32) (*VirNodeMemoryStats, error) {
	var nparams C.int

	ret := C.virNodeGetMemoryStats(c.ptr, C.int(cellNum), nil, &nparams, 0)
	if ret == -1 {
		return nil, GetLastError()
	}

	params := make([]C.virNodeMemoryStats, nparams)
	ret = C.virNodeGetMemoryStats(c.ptr, C.int(cellNum), (*C.virNodeMemoryStats)(unsafe.Pointer(&params)), &nparams, C.uint(flags))
	if ret == -1 {
		return nil, GetLastError()
	}

	stats := &VirNodeMemoryStats{}
	for i := 0; i < int(nparams); i++ {
		param := params[i]
		field := C.GoString((*C.char)(unsafe.Pointer(&param.field)))
		switch field {
		case C.VIR_NODE_MEMORY_STATS_TOTAL:
			stats.TotalSet = true
			stats.Total = uint64(param.value)
		case C.VIR_NODE_MEMORY_STATS_FREE:
			stats.FreeSet = true
			stats.Free = uint64(param.value)
		case C.VIR_NODE_MEMORY_STATS_BUFFERS:
			stats.BuffersSet = true
			stats.Buffers = uint64(param.value)
		case C.VIR_NODE_MEMORY_STATS_CACHED:
			stats.CachedSet = true
			stats.Cached = uint64(param.value)
		}
	}

	return stats, nil
}

type VirNodeSecurityModel struct {
	Model string
	Doi   string
}

func (c *VirConnection) GetSecurityModel() (*VirNodeSecurityModel, error) {
	var cmodel C.virSecurityModel
	ret := C.virNodeGetSecurityModel(c.ptr, &cmodel)
	if ret == -1 {
		return nil, GetLastError()
	}

	return &VirNodeSecurityModel{
		Model: C.GoString((*C.char)(unsafe.Pointer(&cmodel.model))),
		Doi:   C.GoString((*C.char)(unsafe.Pointer(&cmodel.doi))),
	}, nil
}

func (c *VirConnection) SetMemoryParameters(params *VirNodeMemoryParameters, flags uint32) error {
	info := getMemoryParameterFieldInfo(params)

	var nparams C.int

	ret := C.virNodeGetMemoryParameters(c.ptr, nil, &nparams, 0)
	if ret == -1 {
		return GetLastError()
	}

	cparams := make([]C.virTypedParameter, nparams)
	ret = C.virNodeGetMemoryParameters(c.ptr, (*C.virTypedParameter)(unsafe.Pointer(&cparams[0])), &nparams, 0)
	if ret == -1 {
		return GetLastError()
	}

	defer C.virTypedParamsClear((*C.virTypedParameter)(unsafe.Pointer(&cparams[0])), nparams)

	err := typedParamsPack(cparams, info)
	if err != nil {
		return err
	}

	ret = C.virNodeSetMemoryParameters(c.ptr, (*C.virTypedParameter)(unsafe.Pointer(&cparams[0])), nparams, C.uint(flags))

	return nil
}

func (c *VirConnection) SuspendForDuration(target VirNodeSuspendTarget, duration uint64, flags uint32) error {
	ret := C.virNodeSuspendForDuration(c.ptr, C.uint(target), C.ulonglong(duration), C.uint(flags))
	if ret == -1 {
		return GetLastError()
	}
	return nil
}

func (c *VirConnection) DomainSaveImageDefineXML(file string, xml string, flags DomainSaveRestoreFlags) error {
	cfile := C.CString(file)
	defer C.free(cfile)
	cxml := C.CString(xml)
	defer C.free(cxml)

	ret := C.virDomainSaveImageDefineXML(c.ptr, cfile, cxml, C.uint(flags))

	if ret == -1 {
		return GetLastError()
	}

	return nil
}

func (c *VirConnection) DomainSaveImageGetXMLDesc(file string, flags DomainXMLFlags) (string, error) {
	cfile := C.CString(file)
	defer C.free(cfile)

	ret := C.virDomainSaveImageGetXMLDesc(c.ptr, cfile, C.uint(flags))

	if ret == nil {
		return "", GetLastError()
	}

	defer C.free(ret)

	return C.GoString(ret), nil
}

func (c *VirConnection) BaselineCPU(xmlCPUs []string, flags VirConnectBaselineCPUFlags) (string, error) {
	cxmlCPUs := make([]*C.char, len(xmlCPUs))
	for i := 0; i < len(xmlCPUs); i++ {
		cxmlCPUs[i] = C.CString(xmlCPUs[i])
		defer C.free(cxmlCPUs[i])
	}

	ret := C.virConnectBaselineCPU(c.ptr, &cxmlCPUs[0], C.uint(len(xmlCPUs)), C.uint(flags))
	if ret == nil {
		return "", GetLastError()
	}

	defer C.free(ret)

	return C.GoString(ret), nil
}

func (c *VirConnection) CompareCPU(xmlDesc string, flags VirConnectCompareCPUFlags) (VirCPUCompareResult, error) {
	cxmlDesc := C.CString(xmlDesc)
	defer C.free(cxmlDesc)

	ret := C.virConnectCompareCPU(c.ptr, cxmlDesc, C.uint(flags))
	if ret == C.VIR_CPU_COMPARE_ERROR {
		return VIR_CPU_COMPARE_ERROR, GetLastError()
	}

	return VirCPUCompareResult(ret), nil
}

func (c *VirConnection) DomainXMLFromNative(nativeFormat string, nativeConfig string, flags uint32) (string, error) {
	cnativeFormat := C.CString(nativeFormat)
	defer C.free(cnativeFormat)
	cnativeConfig := C.CString(nativeConfig)
	defer C.free(cnativeConfig)

	ret := C.virConnectDomainXMLFromNative(c.ptr, cnativeFormat, cnativeConfig, C.uint(flags))
	if ret == nil {
		return "", GetLastError()
	}

	defer C.free(ret)

	return C.GoString(ret), nil
}

func (c *VirConnection) DomainXMLToNative(nativeFormat string, domainXml string, flags uint32) (string, error) {
	cnativeFormat := C.CString(nativeFormat)
	defer C.free(cnativeFormat)
	cdomainXml := C.CString(domainXml)
	defer C.free(cdomainXml)

	ret := C.virConnectDomainXMLToNative(c.ptr, cnativeFormat, cdomainXml, C.uint(flags))
	if ret == nil {
		return "", GetLastError()
	}

	defer C.free(ret)

	return C.GoString(ret), nil
}

func (c *VirConnection) GetCPUModelNames(arch string, flags uint32) ([]string, error) {
	carch := C.CString(arch)
	defer C.free(carch)

	var cmodels **C.char
	ret := C.virConnectGetCPUModelNames(c.ptr, carch, &cmodels, C.uint(flags))
	if ret == -1 {
		return []string{}, GetLastError()
	}

	models := make([]string, int(ret))
	for i := 0; i < int(ret); i++ {
		cmodel := *(**C.char)(unsafe.Pointer(uintptr(unsafe.Pointer(cmodels)) + (unsafe.Sizeof(*cmodels) * uintptr(i))))

		defer C.free(cmodel)
		models[i] = C.GoString(cmodel)
	}
	defer C.free(cmodels)

	return models, nil
}

func (c *VirConnection) GetDomainCapabilities(emulatorbin string, arch string, machine string, virttype string, flags uint32) (string, error) {
	var cemulatorbin *C.char
	if emulatorbin != "" {
		cemulatorbin = C.CString(emulatorbin)
		defer C.free(cemulatorbin)
	}
	var carch *C.char
	if arch != "" {
		carch = C.CString(arch)
		defer C.free(carch)
	}
	var cmachine *C.char
	if machine != "" {
		cmachine = C.CString(machine)
		defer C.free(cmachine)
	}
	var cvirttype *C.char
	if virttype != "" {
		cvirttype = C.CString(virttype)
		defer C.free(cvirttype)
	}

	ret := C.virConnectGetDomainCapabilities(c.ptr, cemulatorbin, carch, cmachine, cvirttype, C.uint(flags))
	if ret == nil {
		return "", GetLastError()
	}

	defer C.free(ret)

	return C.GoString(ret), nil
}

func (c *VirConnection) GetVersion() (uint32, error) {
	var hvVer C.ulong
	ret := C.virConnectGetVersion(c.ptr, &hvVer)
	if ret == -1 {
		return 0, GetLastError()
	}

	return uint32(hvVer), nil
}

func (c *VirConnection) FindStoragePoolSources(pooltype string, srcSpec string, flags uint32) (string, error) {
	cpooltype := C.CString(pooltype)
	defer C.free(cpooltype)
	var csrcSpec *C.char
	if srcSpec != "" {
		csrcSpec := C.CString(srcSpec)
		defer C.free(csrcSpec)
	}
	ret := C.virConnectFindStoragePoolSources(c.ptr, cpooltype, csrcSpec, C.uint(flags))
	if ret == nil {
		return "", GetLastError()
	}

	defer C.free(ret)

	return C.GoString(ret), nil
}
