package libvirt

import (
	"bufio"
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"regexp"
	"sort"
	"strings"
	"testing"
)

var (
	ignoreFuncs = []string{
		"virConnCopyLastError",
		"virConnGetLastError",
		"virConnResetLastError",
		"virCopyLastError",
		"virFreeError",
		"virGetLastErrorMessage",
		"virInitialize",
		"virResetLastError",
		"virSaveLastError",

		"virConnectBaselineCPU",
		"virConnectCloseFunc",
		"virConnectCompareCPU",
		"virConnectDomainEventAgentLifecycleCallback",
		"virConnectDomainEventBalloonChangeCallback",
		"virConnectDomainEventBlockJobCallback",
		"virConnectDomainEventCallback",
		"virConnectDomainEventDeregister",
		"virConnectDomainEventDeviceAddedCallback",
		"virConnectDomainEventDeviceRemovalFailedCallback",
		"virConnectDomainEventDeviceRemovedCallback",
		"virConnectDomainEventDiskChangeCallback",
		"virConnectDomainEventGraphicsCallback",
		"virConnectDomainEventIOErrorCallback",
		"virConnectDomainEventIOErrorReasonCallback",
		"virConnectDomainEventJobCompletedCallback",
		"virConnectDomainEventMigrationIterationCallback",
		"virConnectDomainEventPMSuspendCallback",
		"virConnectDomainEventPMSuspendDiskCallback",
		"virConnectDomainEventPMWakeupCallback",
		"virConnectDomainEventRTCChangeCallback",
		"virConnectDomainEventRegister",
		"virConnectDomainEventRegisterAny",
		"virConnectDomainEventTrayChangeCallback",
		"virConnectDomainEventTunableCallback",
		"virConnectDomainEventWatchdogCallback",
		"virConnectDomainXMLFromNative",
		"virConnectDomainXMLToNative",
		"virConnectFindStoragePoolSources",
		"virConnectGetAllDomainStats",
		"virConnectGetCPUModelNames",
		"virConnectGetDomainCapabilities",
		"virConnectGetVersion",
		"virConnectListAllNodeDevices",
		"virConnectListAllSecrets",
		"virConnectListNWFilters",
		"virConnectListSecrets",
		"virConnectNetworkEventDeregisterAny",
		"virConnectNetworkEventGenericCallback",
		"virConnectNetworkEventLifecycleCallback",
		"virConnectNetworkEventRegisterAny",
		"virConnectNodeDeviceEventDeregisterAny",
		"virConnectNodeDeviceEventGenericCallback",
		"virConnectNodeDeviceEventLifecycleCallback",
		"virConnectNodeDeviceEventRegisterAny",
		"virConnectNumOfDefinedDomains",
		"virConnectNumOfStoragePools",
		"virConnectRef",
		"virConnectRegisterCloseCallback",
		"virConnectStoragePoolEventDeregisterAny",
		"virConnectStoragePoolEventGenericCallback",
		"virConnectStoragePoolEventLifecycleCallback",
		"virConnectStoragePoolEventRegisterAny",
		"virDefaultErrorFunc",
		"virDomainAddIOThread",
		"virDomainBlockCommit",
		"virDomainBlockCopy",
		"virDomainBlockJobAbort",
		"virDomainBlockJobSetSpeed",
		"virDomainBlockPeek",
		"virDomainBlockPull",
		"virDomainBlockRebase",
		"virDomainBlockResize",
		"virDomainCoreDump",
		"virDomainCoreDumpWithFormat",
		"virDomainCreateLinux",
		"virDomainCreateWithFiles",
		"virDomainCreateXMLWithFiles",
		"virDomainDefineXMLFlags",
		"virDomainDelIOThread",
		"virDomainFSFreeze",
		"virDomainFSInfoFree",
		"virDomainFSThaw",
		"virDomainFSTrim",
		"virDomainGetBlkioParameters",
		"virDomainGetBlockIoTune",
		"virDomainGetBlockJobInfo",
		"virDomainGetConnect",
		"virDomainGetControlInfo",
		"virDomainGetDiskErrors",
		"virDomainGetEmulatorPinInfo",
		"virDomainGetFSInfo",
		"virDomainGetGuestVcpus",
		"virDomainGetHostname",
		"virDomainGetIOThreadInfo",
		"virDomainGetJobInfo",
		"virDomainGetJobStats",
		"virDomainGetMaxMemory",
		"virDomainGetMaxVcpus",
		"virDomainGetMemoryParameters",
		"virDomainGetNumaParameters",
		"virDomainGetOSType",
		"virDomainGetPerfEvents",
		"virDomainGetSchedulerParameters",
		"virDomainGetSchedulerParametersFlags",
		"virDomainGetSchedulerType",
		"virDomainGetSecurityLabel",
		"virDomainGetSecurityLabelList",
		"virDomainGetTime",
		"virDomainGetVcpuPinInfo",
		"virDomainHasCurrentSnapshot",
		"virDomainHasManagedSaveImage",
		"virDomainIOThreadInfoFree",
		"virDomainInjectNMI",
		"virDomainIsUpdated",
		"virDomainListAllSnapshots",
		"virDomainListGetStats",
		"virDomainLookupByUUID",
		"virDomainManagedSave",
		"virDomainManagedSaveRemove",
		"virDomainMemoryPeek",
		"virDomainMigrate",
		"virDomainMigrate2",
		"virDomainMigrate3",
		"virDomainMigrateGetCompressionCache",
		"virDomainMigrateGetMaxSpeed",
		"virDomainMigrateSetCompressionCache",
		"virDomainMigrateSetMaxDowntime",
		"virDomainMigrateSetMaxSpeed",
		"virDomainMigrateStartPostCopy",
		"virDomainMigrateToURI",
		"virDomainMigrateToURI2",
		"virDomainMigrateToURI3",
		"virDomainOpenChannel",
		"virDomainOpenConsole",
		"virDomainOpenGraphics",
		"virDomainOpenGraphicsFD",
		"virDomainPMSuspendForDuration",
		"virDomainPMWakeup",
		"virDomainPinEmulator",
		"virDomainPinIOThread",
		"virDomainRef",
		"virDomainRename",
		"virDomainReset",
		"virDomainSaveImageDefineXML",
		"virDomainSaveImageGetXMLDesc",
		"virDomainSendProcessSignal",
		"virDomainSetBlkioParameters",
		"virDomainSetBlockIoTune",
		"virDomainSetGuestVcpus",
		"virDomainSetInterfaceParameters",
		"virDomainSetMemoryParameters",
		"virDomainSetNumaParameters",
		"virDomainSetPerfEvents",
		"virDomainSetSchedulerParameters",
		"virDomainSetSchedulerParametersFlags",
		"virDomainSetTime",
		"virDomainSetUserPassword",
		"virDomainSnapshotCurrent",
		"virDomainSnapshotGetConnect",
		"virDomainSnapshotGetDomain",
		"virDomainSnapshotGetName",
		"virDomainSnapshotGetParent",
		"virDomainSnapshotGetXMLDesc",
		"virDomainSnapshotHasMetadata",
		"virDomainSnapshotIsCurrent",
		"virDomainSnapshotListAllChildren",
		"virDomainSnapshotListChildrenNames",
		"virDomainSnapshotListNames",
		"virDomainSnapshotLookupByName",
		"virDomainSnapshotNum",
		"virDomainSnapshotNumChildren",
		"virDomainSnapshotRef",
		"virDomainStatsRecordListFree",
		"virErrorFunc",
		"virEventAddHandle",
		"virEventAddHandleFunc",
		"virEventAddTimeout",
		"virEventAddTimeoutFunc",
		"virEventHandleCallback",
		"virEventRegisterImpl",
		"virEventRemoveHandle",
		"virEventRemoveHandleFunc",
		"virEventRemoveTimeout",
		"virEventRemoveTimeoutFunc",
		"virEventTimeoutCallback",
		"virEventUpdateHandle",
		"virEventUpdateHandleFunc",
		"virEventUpdateTimeout",
		"virEventUpdateTimeoutFunc",
		"virFreeCallback",
		"virInterfaceChangeBegin",
		"virInterfaceChangeCommit",
		"virInterfaceChangeRollback",
		"virInterfaceGetConnect",
		"virInterfaceRef",
		"virNWFilterLookupByUUID",
		"virNWFilterRef",
		"virNetworkDHCPLeaseFree",
		"virNetworkGetConnect",
		"virNetworkLookupByUUID",
		"virNetworkRef",
		"virNetworkUpdate",
		"virNodeAllocPages",
		"virNodeDeviceCreateXML",
		"virNodeDeviceDestroy",
		"virNodeDeviceDetachFlags",
		"virNodeDeviceDettach",
		"virNodeDeviceFree",
		"virNodeDeviceGetName",
		"virNodeDeviceGetParent",
		"virNodeDeviceGetXMLDesc",
		"virNodeDeviceListCaps",
		"virNodeDeviceLookupByName",
		"virNodeDeviceLookupSCSIHostByWWN",
		"virNodeDeviceNumOfCaps",
		"virNodeDeviceReAttach",
		"virNodeDeviceRef",
		"virNodeDeviceReset",
		"virNodeGetCPUMap",
		"virNodeGetCPUStats",
		"virNodeGetCellsFreeMemory",
		"virNodeGetFreeMemory",
		"virNodeGetFreePages",
		"virNodeGetMemoryParameters",
		"virNodeGetMemoryStats",
		"virNodeGetSecurityModel",
		"virNodeListDevices",
		"virNodeNumOfDevices",
		"virNodeSetMemoryParameters",
		"virNodeSuspendForDuration",
		"virSecretGetConnect",
		"virSecretGetValue",
		"virSecretLookupByUUID",
		"virSecretRef",
		"virStoragePoolCreateXML",
		"virStoragePoolGetConnect",
		"virStoragePoolIsPersistent",
		"virStoragePoolListAllVolumes",
		"virStoragePoolListVolumes",
		"virStoragePoolLookupByUUID",
		"virStoragePoolNumOfVolumes",
		"virStoragePoolRef",
		"virStorageVolGetConnect",
		"virStorageVolRef",
		"virStreamEventAddCallback",
		"virStreamEventCallback",
		"virStreamEventRemoveCallback",
		"virStreamEventUpdateCallback",
		"virStreamRecvAll",
		"virStreamRef",
		"virStreamSendAll",
		"virStreamSinkFunc",
		"virStreamSourceFunc",
		"virTypedParamsAddBoolean",
		"virTypedParamsAddDouble",
		"virTypedParamsAddFromString",
		"virTypedParamsAddInt",
		"virTypedParamsAddLLong",
		"virTypedParamsAddString",
		"virTypedParamsAddStringList",
		"virTypedParamsAddUInt",
		"virTypedParamsAddULLong",
		"virTypedParamsClear",
		"virTypedParamsGet",
		"virTypedParamsGetBoolean",
		"virTypedParamsGetDouble",
		"virTypedParamsGetInt",
		"virTypedParamsGetLLong",
		"virTypedParamsGetString",
		"virTypedParamsGetUInt",
		"virTypedParamsGetULLong",
	}

	ignoreMacros = []string{
		"_virBlkioParameter",
		"_virMemoryParameter",
		"_virSchedParameter",
		"LIBVIR_CHECK_VERSION",

		"LIBVIR_VERSION_NUMBER",
		"VIR_COPY_CPUMAP",
		"VIR_CPU_MAPLEN",
		"VIR_CPU_USABLE",
		"VIR_CPU_USED",
		"VIR_DOMAIN_BANDWIDTH_IN_AVERAGE",
		"VIR_DOMAIN_BANDWIDTH_IN_BURST",
		"VIR_DOMAIN_BANDWIDTH_IN_FLOOR",
		"VIR_DOMAIN_BANDWIDTH_IN_PEAK",
		"VIR_DOMAIN_BANDWIDTH_OUT_AVERAGE",
		"VIR_DOMAIN_BANDWIDTH_OUT_BURST",
		"VIR_DOMAIN_BANDWIDTH_OUT_PEAK",
		"VIR_DOMAIN_BLKIO_DEVICE_READ_BPS",
		"VIR_DOMAIN_BLKIO_DEVICE_READ_IOPS",
		"VIR_DOMAIN_BLKIO_DEVICE_WEIGHT",
		"VIR_DOMAIN_BLKIO_DEVICE_WRITE_BPS",
		"VIR_DOMAIN_BLKIO_DEVICE_WRITE_IOPS",
		"VIR_DOMAIN_BLKIO_FIELD_LENGTH",
		"VIR_DOMAIN_BLKIO_WEIGHT",
		"VIR_DOMAIN_BLOCK_COPY_BANDWIDTH",
		"VIR_DOMAIN_BLOCK_COPY_BUF_SIZE",
		"VIR_DOMAIN_BLOCK_COPY_GRANULARITY",
		"VIR_DOMAIN_BLOCK_IOTUNE_READ_BYTES_SEC",
		"VIR_DOMAIN_BLOCK_IOTUNE_READ_BYTES_SEC_MAX",
		"VIR_DOMAIN_BLOCK_IOTUNE_READ_BYTES_SEC_MAX_LENGTH",
		"VIR_DOMAIN_BLOCK_IOTUNE_READ_IOPS_SEC",
		"VIR_DOMAIN_BLOCK_IOTUNE_READ_IOPS_SEC_MAX",
		"VIR_DOMAIN_BLOCK_IOTUNE_READ_IOPS_SEC_MAX_LENGTH",
		"VIR_DOMAIN_BLOCK_IOTUNE_SIZE_IOPS_SEC",
		"VIR_DOMAIN_BLOCK_IOTUNE_TOTAL_BYTES_SEC",
		"VIR_DOMAIN_BLOCK_IOTUNE_TOTAL_BYTES_SEC_MAX",
		"VIR_DOMAIN_BLOCK_IOTUNE_TOTAL_BYTES_SEC_MAX_LENGTH",
		"VIR_DOMAIN_BLOCK_IOTUNE_TOTAL_IOPS_SEC",
		"VIR_DOMAIN_BLOCK_IOTUNE_TOTAL_IOPS_SEC_MAX",
		"VIR_DOMAIN_BLOCK_IOTUNE_TOTAL_IOPS_SEC_MAX_LENGTH",
		"VIR_DOMAIN_BLOCK_IOTUNE_WRITE_BYTES_SEC",
		"VIR_DOMAIN_BLOCK_IOTUNE_WRITE_BYTES_SEC_MAX",
		"VIR_DOMAIN_BLOCK_IOTUNE_WRITE_BYTES_SEC_MAX_LENGTH",
		"VIR_DOMAIN_BLOCK_IOTUNE_WRITE_IOPS_SEC",
		"VIR_DOMAIN_BLOCK_IOTUNE_WRITE_IOPS_SEC_MAX",
		"VIR_DOMAIN_BLOCK_IOTUNE_WRITE_IOPS_SEC_MAX_LENGTH",
		"VIR_DOMAIN_BLOCK_STATS_ERRS",
		"VIR_DOMAIN_BLOCK_STATS_FIELD_LENGTH",
		"VIR_DOMAIN_BLOCK_STATS_FLUSH_REQ",
		"VIR_DOMAIN_BLOCK_STATS_FLUSH_TOTAL_TIMES",
		"VIR_DOMAIN_BLOCK_STATS_READ_BYTES",
		"VIR_DOMAIN_BLOCK_STATS_READ_REQ",
		"VIR_DOMAIN_BLOCK_STATS_READ_TOTAL_TIMES",
		"VIR_DOMAIN_BLOCK_STATS_WRITE_BYTES",
		"VIR_DOMAIN_BLOCK_STATS_WRITE_REQ",
		"VIR_DOMAIN_BLOCK_STATS_WRITE_TOTAL_TIMES",
		"VIR_DOMAIN_EVENT_CALLBACK",
		"VIR_DOMAIN_JOB_AUTO_CONVERGE_THROTTLE",
		"VIR_DOMAIN_JOB_COMPRESSION_BYTES",
		"VIR_DOMAIN_JOB_COMPRESSION_CACHE",
		"VIR_DOMAIN_JOB_COMPRESSION_CACHE_MISSES",
		"VIR_DOMAIN_JOB_COMPRESSION_OVERFLOW",
		"VIR_DOMAIN_JOB_COMPRESSION_PAGES",
		"VIR_DOMAIN_JOB_DATA_PROCESSED",
		"VIR_DOMAIN_JOB_DATA_REMAINING",
		"VIR_DOMAIN_JOB_DATA_TOTAL",
		"VIR_DOMAIN_JOB_DISK_BPS",
		"VIR_DOMAIN_JOB_DISK_PROCESSED",
		"VIR_DOMAIN_JOB_DISK_REMAINING",
		"VIR_DOMAIN_JOB_DISK_TOTAL",
		"VIR_DOMAIN_JOB_DOWNTIME",
		"VIR_DOMAIN_JOB_DOWNTIME_NET",
		"VIR_DOMAIN_JOB_MEMORY_BPS",
		"VIR_DOMAIN_JOB_MEMORY_CONSTANT",
		"VIR_DOMAIN_JOB_MEMORY_DIRTY_RATE",
		"VIR_DOMAIN_JOB_MEMORY_ITERATION",
		"VIR_DOMAIN_JOB_MEMORY_NORMAL",
		"VIR_DOMAIN_JOB_MEMORY_NORMAL_BYTES",
		"VIR_DOMAIN_JOB_MEMORY_PROCESSED",
		"VIR_DOMAIN_JOB_MEMORY_REMAINING",
		"VIR_DOMAIN_JOB_MEMORY_TOTAL",
		"VIR_DOMAIN_JOB_SETUP_TIME",
		"VIR_DOMAIN_JOB_TIME_ELAPSED",
		"VIR_DOMAIN_JOB_TIME_ELAPSED_NET",
		"VIR_DOMAIN_JOB_TIME_REMAINING",
		"VIR_DOMAIN_MEMORY_FIELD_LENGTH",
		"VIR_DOMAIN_MEMORY_HARD_LIMIT",
		"VIR_DOMAIN_MEMORY_MIN_GUARANTEE",
		"VIR_DOMAIN_MEMORY_SOFT_LIMIT",
		"VIR_DOMAIN_MEMORY_SWAP_HARD_LIMIT",
		"VIR_DOMAIN_NUMA_MODE",
		"VIR_DOMAIN_NUMA_NODESET",
		"VIR_DOMAIN_SCHEDULER_CAP",
		"VIR_DOMAIN_SCHEDULER_CPU_SHARES",
		"VIR_DOMAIN_SCHEDULER_EMULATOR_PERIOD",
		"VIR_DOMAIN_SCHEDULER_EMULATOR_QUOTA",
		"VIR_DOMAIN_SCHEDULER_GLOBAL_PERIOD",
		"VIR_DOMAIN_SCHEDULER_GLOBAL_QUOTA",
		"VIR_DOMAIN_SCHEDULER_IOTHREAD_PERIOD",
		"VIR_DOMAIN_SCHEDULER_IOTHREAD_QUOTA",
		"VIR_DOMAIN_SCHEDULER_LIMIT",
		"VIR_DOMAIN_SCHEDULER_RESERVATION",
		"VIR_DOMAIN_SCHEDULER_SHARES",
		"VIR_DOMAIN_SCHEDULER_VCPU_PERIOD",
		"VIR_DOMAIN_SCHEDULER_VCPU_QUOTA",
		"VIR_DOMAIN_SCHEDULER_WEIGHT",
		"VIR_DOMAIN_SCHED_FIELD_LENGTH",
		"VIR_DOMAIN_SEND_KEY_MAX_KEYS",
		"VIR_DOMAIN_TUNABLE_BLKDEV_DISK",
		"VIR_DOMAIN_TUNABLE_BLKDEV_READ_BYTES_SEC",
		"VIR_DOMAIN_TUNABLE_BLKDEV_READ_BYTES_SEC_MAX",
		"VIR_DOMAIN_TUNABLE_BLKDEV_READ_BYTES_SEC_MAX_LENGTH",
		"VIR_DOMAIN_TUNABLE_BLKDEV_READ_IOPS_SEC",
		"VIR_DOMAIN_TUNABLE_BLKDEV_READ_IOPS_SEC_MAX",
		"VIR_DOMAIN_TUNABLE_BLKDEV_READ_IOPS_SEC_MAX_LENGTH",
		"VIR_DOMAIN_TUNABLE_BLKDEV_SIZE_IOPS_SEC",
		"VIR_DOMAIN_TUNABLE_BLKDEV_TOTAL_BYTES_SEC",
		"VIR_DOMAIN_TUNABLE_BLKDEV_TOTAL_BYTES_SEC_MAX",
		"VIR_DOMAIN_TUNABLE_BLKDEV_TOTAL_BYTES_SEC_MAX_LENGTH",
		"VIR_DOMAIN_TUNABLE_BLKDEV_TOTAL_IOPS_SEC",
		"VIR_DOMAIN_TUNABLE_BLKDEV_TOTAL_IOPS_SEC_MAX",
		"VIR_DOMAIN_TUNABLE_BLKDEV_TOTAL_IOPS_SEC_MAX_LENGTH",
		"VIR_DOMAIN_TUNABLE_BLKDEV_WRITE_BYTES_SEC",
		"VIR_DOMAIN_TUNABLE_BLKDEV_WRITE_BYTES_SEC_MAX",
		"VIR_DOMAIN_TUNABLE_BLKDEV_WRITE_BYTES_SEC_MAX_LENGTH",
		"VIR_DOMAIN_TUNABLE_BLKDEV_WRITE_IOPS_SEC",
		"VIR_DOMAIN_TUNABLE_BLKDEV_WRITE_IOPS_SEC_MAX",
		"VIR_DOMAIN_TUNABLE_BLKDEV_WRITE_IOPS_SEC_MAX_LENGTH",
		"VIR_DOMAIN_TUNABLE_CPU_CPU_SHARES",
		"VIR_DOMAIN_TUNABLE_CPU_EMULATORPIN",
		"VIR_DOMAIN_TUNABLE_CPU_EMULATOR_PERIOD",
		"VIR_DOMAIN_TUNABLE_CPU_EMULATOR_QUOTA",
		"VIR_DOMAIN_TUNABLE_CPU_GLOBAL_PERIOD",
		"VIR_DOMAIN_TUNABLE_CPU_GLOBAL_QUOTA",
		"VIR_DOMAIN_TUNABLE_CPU_IOTHREADSPIN",
		"VIR_DOMAIN_TUNABLE_CPU_IOTHREAD_PERIOD",
		"VIR_DOMAIN_TUNABLE_CPU_IOTHREAD_QUOTA",
		"VIR_DOMAIN_TUNABLE_CPU_VCPUPIN",
		"VIR_DOMAIN_TUNABLE_CPU_VCPU_PERIOD",
		"VIR_DOMAIN_TUNABLE_CPU_VCPU_QUOTA",
		"VIR_GET_CPUMAP",
		"VIR_MIGRATE_PARAM_AUTO_CONVERGE_INCREMENT",
		"VIR_MIGRATE_PARAM_AUTO_CONVERGE_INITIAL",
		"VIR_MIGRATE_PARAM_BANDWIDTH",
		"VIR_MIGRATE_PARAM_COMPRESSION",
		"VIR_MIGRATE_PARAM_COMPRESSION_MT_DTHREADS",
		"VIR_MIGRATE_PARAM_COMPRESSION_MT_LEVEL",
		"VIR_MIGRATE_PARAM_COMPRESSION_MT_THREADS",
		"VIR_MIGRATE_PARAM_COMPRESSION_XBZRLE_CACHE",
		"VIR_MIGRATE_PARAM_DEST_NAME",
		"VIR_MIGRATE_PARAM_DEST_XML",
		"VIR_MIGRATE_PARAM_DISKS_PORT",
		"VIR_MIGRATE_PARAM_GRAPHICS_URI",
		"VIR_MIGRATE_PARAM_LISTEN_ADDRESS",
		"VIR_MIGRATE_PARAM_MIGRATE_DISKS",
		"VIR_MIGRATE_PARAM_PERSIST_XML",
		"VIR_MIGRATE_PARAM_URI",
		"VIR_NETWORK_EVENT_CALLBACK",
		"VIR_NODEINFO_MAXCPUS",
		"VIR_NODE_CPU_STATS_FIELD_LENGTH",
		"VIR_NODE_CPU_STATS_IDLE",
		"VIR_NODE_CPU_STATS_INTR",
		"VIR_NODE_CPU_STATS_IOWAIT",
		"VIR_NODE_CPU_STATS_KERNEL",
		"VIR_NODE_CPU_STATS_USER",
		"VIR_NODE_CPU_STATS_UTILIZATION",
		"VIR_NODE_DEVICE_EVENT_CALLBACK",
		"VIR_NODE_MEMORY_SHARED_FULL_SCANS",
		"VIR_NODE_MEMORY_SHARED_MERGE_ACROSS_NODES",
		"VIR_NODE_MEMORY_SHARED_PAGES_SHARED",
		"VIR_NODE_MEMORY_SHARED_PAGES_SHARING",
		"VIR_NODE_MEMORY_SHARED_PAGES_TO_SCAN",
		"VIR_NODE_MEMORY_SHARED_PAGES_UNSHARED",
		"VIR_NODE_MEMORY_SHARED_PAGES_VOLATILE",
		"VIR_NODE_MEMORY_SHARED_SLEEP_MILLISECS",
		"VIR_NODE_MEMORY_STATS_BUFFERS",
		"VIR_NODE_MEMORY_STATS_CACHED",
		"VIR_NODE_MEMORY_STATS_FIELD_LENGTH",
		"VIR_NODE_MEMORY_STATS_FREE",
		"VIR_NODE_MEMORY_STATS_TOTAL",
		"VIR_PERF_PARAM_CACHE_MISSES",
		"VIR_PERF_PARAM_CACHE_REFERENCES",
		"VIR_PERF_PARAM_CMT",
		"VIR_PERF_PARAM_CPU_CYCLES",
		"VIR_PERF_PARAM_INSTRUCTIONS",
		"VIR_PERF_PARAM_MBML",
		"VIR_PERF_PARAM_MBMT",
		"VIR_SECURITY_DOI_BUFLEN",
		"VIR_SECURITY_LABEL_BUFLEN",
		"VIR_SECURITY_MODEL_BUFLEN",
		"VIR_STORAGE_POOL_EVENT_CALLBACK",
		"VIR_UNUSE_CPU",
		"VIR_USE_CPU",
	}

	ignoreEnums = []string{

		// Deprecated in favour of VIR_TYPED_PARAM_*
		"VIR_DOMAIN_BLKIO_PARAM_BOOLEAN",
		"VIR_DOMAIN_BLKIO_PARAM_DOUBLE",
		"VIR_DOMAIN_BLKIO_PARAM_INT",
		"VIR_DOMAIN_BLKIO_PARAM_LLONG",
		"VIR_DOMAIN_BLKIO_PARAM_UINT",
		"VIR_DOMAIN_BLKIO_PARAM_ULLONG",
		"VIR_DOMAIN_MEMORY_PARAM_BOOLEAN",
		"VIR_DOMAIN_MEMORY_PARAM_DOUBLE",
		"VIR_DOMAIN_MEMORY_PARAM_INT",
		"VIR_DOMAIN_MEMORY_PARAM_LLONG",
		"VIR_DOMAIN_MEMORY_PARAM_UINT",
		"VIR_DOMAIN_MEMORY_PARAM_ULLONG",
		"VIR_DOMAIN_SCHED_FIELD_BOOLEAN",
		"VIR_DOMAIN_SCHED_FIELD_DOUBLE",
		"VIR_DOMAIN_SCHED_FIELD_INT",
		"VIR_DOMAIN_SCHED_FIELD_LLONG",
		"VIR_DOMAIN_SCHED_FIELD_UINT",
		"VIR_DOMAIN_SCHED_FIELD_ULLONG",

		"VIR_STREAM_EVENT_ERROR",
		"VIR_STREAM_EVENT_HANGUP",
		"VIR_STREAM_EVENT_READABLE",
		"VIR_STREAM_EVENT_WRITABLE",
		"VIR_TYPED_PARAM_STRING_OKAY",
	}
)

type CharsetISO88591er struct {
	r   io.ByteReader
	buf *bytes.Buffer
}

func NewCharsetISO88591(r io.Reader) *CharsetISO88591er {
	buf := bytes.Buffer{}
	return &CharsetISO88591er{r.(io.ByteReader), &buf}
}

func (cs *CharsetISO88591er) Read(p []byte) (n int, err error) {
	for _ = range p {
		if r, err := cs.r.ReadByte(); err != nil {
			break
		} else {
			cs.buf.WriteRune(rune(r))
		}
	}
	return cs.buf.Read(p)
}

func isCharset(charset string, names []string) bool {
	charset = strings.ToLower(charset)
	for _, n := range names {
		if charset == strings.ToLower(n) {
			return true
		}
	}
	return false
}

func IsCharsetISO88591(charset string) bool {
	// http://www.iana.org/assignments/character-sets
	// (last updated 2010-11-04)
	names := []string{
		// Name
		"ISO_8859-1:1987",
		// Alias (preferred MIME name)
		"ISO-8859-1",
		// Aliases
		"iso-ir-100",
		"ISO_8859-1",
		"latin1",
		"l1",
		"IBM819",
		"CP819",
		"csISOLatin1",
	}
	return isCharset(charset, names)
}

func CharsetReader(charset string, input io.Reader) (io.Reader, error) {
	if IsCharsetISO88591(charset) {
		return NewCharsetISO88591(input), nil
	}
	return input, nil
}

type APIExport struct {
	Type   string `xml:"type,attr"`
	Symbol string `xml:"symbol,attr"`
}

type APIFile struct {
	Name    string      `xml:"name,attr"`
	Exports []APIExport `xml:"exports"`
}

type API struct {
	XMLName xml.Name  `xml:"api"`
	Files   []APIFile `xml:"files>file"`
}

func GetAPIPath() string {
	cmd := exec.Command("pkg-config", "--variable=libvirt_api", "libvirt")

	cmdOutput := &bytes.Buffer{}
	cmd.Stdout = cmdOutput

	err := cmd.Run()
	if err != nil {
		panic(err)
	}

	return strings.TrimSpace(string(cmdOutput.Bytes()))
}

func GetAPI(path string) *API {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	decoder := xml.NewDecoder(file)
	decoder.CharsetReader = CharsetReader
	var api API
	err = decoder.Decode(&api)
	if err != nil {
		panic(err)
	}

	return &api
}

func GetSourceFiles() []string {
	files, _ := ioutil.ReadDir(".")

	src := make([]string, 0)
	for _, f := range files {
		if strings.HasSuffix(f.Name(), ".go") &&
			!strings.HasSuffix(f.Name(), "_test.go") {
			src = append(src, f.Name())
		}
	}

	return src
}

func GetAPISymbols(api *API) (funcs map[string]bool, macros map[string]bool, enums map[string]bool) {

	funcs = make(map[string]bool)
	macros = make(map[string]bool)
	enums = make(map[string]bool)
	for _, file := range api.Files {
		for _, export := range file.Exports {
			if export.Type == "function" {
				funcs[export.Symbol] = true
			} else if export.Type == "enum" {
				if !strings.HasSuffix(export.Symbol, "_LAST") {
					enums[export.Symbol] = true
				}
			} else if export.Type == "macro" {
				macros[export.Symbol] = true
			}
		}
	}

	return
}

func ProcessFile(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	re, err := regexp.Compile("C\\.((vir|VIR)[a-zA-Z0-9_]+)\\b")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	symbols := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()

		match := re.FindStringSubmatch(line)
		if match != nil {
			symbols = append(symbols, match[1])
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return symbols
}

func RecordUsage(symbols []string, funcs map[string]bool, macros map[string]bool, enums map[string]bool) {

	for _, name := range symbols {
		_, ok := funcs[name]
		if ok {
			funcs[name] = false
			continue
		}

		_, ok = macros[name]
		if ok {
			macros[name] = false
			continue
		}

		_, ok = enums[name]
		if ok {
			enums[name] = false
			continue
		}
	}
}

func ReportMissing(missingNames map[string]bool, symtype string) bool {
	missing := make([]string, 0)
	for key, value := range missingNames {
		if value {
			missing = append(missing, key)
		}
	}

	sort.Strings(missing)

	for _, name := range missing {
		fmt.Println("Missing " + symtype + " '" + name + "'")
	}

	return len(missing) != 0
}

func SetIgnores(ignores []string, symbols map[string]bool) {
	for _, name := range ignores {
		symbols[name] = false
	}
}

func TestAPICoverage(t *testing.T) {
	path := GetAPIPath()
	api := GetAPI(path)

	funcs, macros, enums := GetAPISymbols(api)

	SetIgnores(ignoreFuncs, funcs)
	SetIgnores(ignoreMacros, macros)
	SetIgnores(ignoreEnums, enums)

	src := GetSourceFiles()

	for _, path := range src {
		symbols := ProcessFile(path)

		RecordUsage(symbols, funcs, macros, enums)
	}

	missing := false
	if ReportMissing(funcs, "function") {
		missing = true
	}
	if ReportMissing(macros, "macro") {
		missing = true
	}
	if ReportMissing(enums, "enum") {
		missing = true
	}
	if missing {
		panic("Missing symbols found")
	}
}
