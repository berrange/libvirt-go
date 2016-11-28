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
		"VIR_CONNECT_BASELINE_CPU_EXPAND_FEATURES",
		"VIR_CONNECT_BASELINE_CPU_MIGRATABLE",
		"VIR_CONNECT_COMPARE_CPU_FAIL_INCOMPATIBLE",
		"VIR_CONNECT_DOMAIN_EVENT_AGENT_LIFECYCLE_REASON_CHANNEL",
		"VIR_CONNECT_DOMAIN_EVENT_AGENT_LIFECYCLE_REASON_DOMAIN_STARTED",
		"VIR_CONNECT_DOMAIN_EVENT_AGENT_LIFECYCLE_REASON_UNKNOWN",
		"VIR_CONNECT_DOMAIN_EVENT_AGENT_LIFECYCLE_STATE_CONNECTED",
		"VIR_CONNECT_DOMAIN_EVENT_AGENT_LIFECYCLE_STATE_DISCONNECTED",
		"VIR_CONNECT_GET_ALL_DOMAINS_STATS_ACTIVE",
		"VIR_CONNECT_GET_ALL_DOMAINS_STATS_BACKING",
		"VIR_CONNECT_GET_ALL_DOMAINS_STATS_ENFORCE_STATS",
		"VIR_CONNECT_GET_ALL_DOMAINS_STATS_INACTIVE",
		"VIR_CONNECT_GET_ALL_DOMAINS_STATS_OTHER",
		"VIR_CONNECT_GET_ALL_DOMAINS_STATS_PAUSED",
		"VIR_CONNECT_GET_ALL_DOMAINS_STATS_PERSISTENT",
		"VIR_CONNECT_GET_ALL_DOMAINS_STATS_RUNNING",
		"VIR_CONNECT_GET_ALL_DOMAINS_STATS_SHUTOFF",
		"VIR_CONNECT_GET_ALL_DOMAINS_STATS_TRANSIENT",
		"VIR_CONNECT_LIST_INTERFACES_ACTIVE",
		"VIR_CONNECT_LIST_INTERFACES_INACTIVE",
		"VIR_CONNECT_LIST_NODE_DEVICES_CAP_FC_HOST",
		"VIR_CONNECT_LIST_NODE_DEVICES_CAP_NET",
		"VIR_CONNECT_LIST_NODE_DEVICES_CAP_PCI_DEV",
		"VIR_CONNECT_LIST_NODE_DEVICES_CAP_SCSI",
		"VIR_CONNECT_LIST_NODE_DEVICES_CAP_SCSI_GENERIC",
		"VIR_CONNECT_LIST_NODE_DEVICES_CAP_SCSI_HOST",
		"VIR_CONNECT_LIST_NODE_DEVICES_CAP_SCSI_TARGET",
		"VIR_CONNECT_LIST_NODE_DEVICES_CAP_STORAGE",
		"VIR_CONNECT_LIST_NODE_DEVICES_CAP_SYSTEM",
		"VIR_CONNECT_LIST_NODE_DEVICES_CAP_USB_DEV",
		"VIR_CONNECT_LIST_NODE_DEVICES_CAP_USB_INTERFACE",
		"VIR_CONNECT_LIST_NODE_DEVICES_CAP_VPORTS",
		"VIR_CONNECT_LIST_SECRETS_EPHEMERAL",
		"VIR_CONNECT_LIST_SECRETS_NO_EPHEMERAL",
		"VIR_CONNECT_LIST_SECRETS_NO_PRIVATE",
		"VIR_CONNECT_LIST_SECRETS_PRIVATE",
		"VIR_CONNECT_LIST_STORAGE_POOLS_ZFS",
		"VIR_CONNECT_NO_ALIASES",
		"VIR_CONNECT_RO",
		"VIR_CPU_COMPARE_ERROR",
		"VIR_CPU_COMPARE_IDENTICAL",
		"VIR_CPU_COMPARE_INCOMPATIBLE",
		"VIR_CPU_COMPARE_SUPERSET",
		"VIR_CRED_AUTHNAME",
		"VIR_CRED_CNONCE",
		"VIR_CRED_ECHOPROMPT",
		"VIR_CRED_EXTERNAL",
		"VIR_CRED_LANGUAGE",
		"VIR_CRED_NOECHOPROMPT",
		"VIR_CRED_PASSPHRASE",
		"VIR_CRED_REALM",
		"VIR_CRED_USERNAME",
		"VIR_DOMAIN_BLKIO_PARAM_BOOLEAN",
		"VIR_DOMAIN_BLKIO_PARAM_DOUBLE",
		"VIR_DOMAIN_BLKIO_PARAM_INT",
		"VIR_DOMAIN_BLKIO_PARAM_LLONG",
		"VIR_DOMAIN_BLKIO_PARAM_UINT",
		"VIR_DOMAIN_BLKIO_PARAM_ULLONG",
		"VIR_DOMAIN_BLOCKED_UNKNOWN",
		"VIR_DOMAIN_BLOCK_COMMIT_ACTIVE",
		"VIR_DOMAIN_BLOCK_COMMIT_BANDWIDTH_BYTES",
		"VIR_DOMAIN_BLOCK_COMMIT_DELETE",
		"VIR_DOMAIN_BLOCK_COMMIT_RELATIVE",
		"VIR_DOMAIN_BLOCK_COMMIT_SHALLOW",
		"VIR_DOMAIN_BLOCK_COPY_REUSE_EXT",
		"VIR_DOMAIN_BLOCK_COPY_SHALLOW",
		"VIR_DOMAIN_BLOCK_JOB_ABORT_ASYNC",
		"VIR_DOMAIN_BLOCK_JOB_ABORT_PIVOT",
		"VIR_DOMAIN_BLOCK_JOB_INFO_BANDWIDTH_BYTES",
		"VIR_DOMAIN_BLOCK_JOB_SPEED_BANDWIDTH_BYTES",
		"VIR_DOMAIN_BLOCK_PULL_BANDWIDTH_BYTES",
		"VIR_DOMAIN_BLOCK_REBASE_BANDWIDTH_BYTES",
		"VIR_DOMAIN_BLOCK_REBASE_COPY",
		"VIR_DOMAIN_BLOCK_REBASE_COPY_DEV",
		"VIR_DOMAIN_BLOCK_REBASE_COPY_RAW",
		"VIR_DOMAIN_BLOCK_REBASE_RELATIVE",
		"VIR_DOMAIN_BLOCK_REBASE_REUSE_EXT",
		"VIR_DOMAIN_BLOCK_REBASE_SHALLOW",
		"VIR_DOMAIN_BLOCK_RESIZE_BYTES",
		"VIR_DOMAIN_CHANNEL_FORCE",
		"VIR_DOMAIN_CONSOLE_FORCE",
		"VIR_DOMAIN_CONSOLE_SAFE",
		"VIR_DOMAIN_CONTROL_ERROR",
		"VIR_DOMAIN_CONTROL_ERROR_REASON_INTERNAL",
		"VIR_DOMAIN_CONTROL_ERROR_REASON_MONITOR",
		"VIR_DOMAIN_CONTROL_ERROR_REASON_NONE",
		"VIR_DOMAIN_CONTROL_ERROR_REASON_UNKNOWN",
		"VIR_DOMAIN_CONTROL_JOB",
		"VIR_DOMAIN_CONTROL_OCCUPIED",
		"VIR_DOMAIN_CONTROL_OK",
		"VIR_DOMAIN_CORE_DUMP_FORMAT_KDUMP_LZO",
		"VIR_DOMAIN_CORE_DUMP_FORMAT_KDUMP_SNAPPY",
		"VIR_DOMAIN_CORE_DUMP_FORMAT_KDUMP_ZLIB",
		"VIR_DOMAIN_CORE_DUMP_FORMAT_RAW",
		"VIR_DOMAIN_CRASHED_PANICKED",
		"VIR_DOMAIN_CRASHED_UNKNOWN",
		"VIR_DOMAIN_DEFINE_VALIDATE",
		"VIR_DOMAIN_DEVICE_MODIFY_CONFIG",
		"VIR_DOMAIN_DEVICE_MODIFY_CURRENT",
		"VIR_DOMAIN_DEVICE_MODIFY_LIVE",
		"VIR_DOMAIN_DISK_ERROR_NONE",
		"VIR_DOMAIN_DISK_ERROR_NO_SPACE",
		"VIR_DOMAIN_DISK_ERROR_UNSPEC",
		"VIR_DOMAIN_EVENT_CRASHED_PANICKED",
		"VIR_DOMAIN_EVENT_DEFINED_FROM_SNAPSHOT",
		"VIR_DOMAIN_EVENT_DEFINED_RENAMED",
		"VIR_DOMAIN_EVENT_ID_AGENT_LIFECYCLE",
		"VIR_DOMAIN_EVENT_ID_DEVICE_ADDED",
		"VIR_DOMAIN_EVENT_ID_DEVICE_REMOVAL_FAILED",
		"VIR_DOMAIN_EVENT_ID_JOB_COMPLETED",
		"VIR_DOMAIN_EVENT_ID_MIGRATION_ITERATION",
		"VIR_DOMAIN_EVENT_ID_TUNABLE",
		"VIR_DOMAIN_EVENT_PMSUSPENDED_DISK",
		"VIR_DOMAIN_EVENT_PMSUSPENDED_MEMORY",
		"VIR_DOMAIN_EVENT_RESUMED_POSTCOPY",
		"VIR_DOMAIN_EVENT_SUSPENDED_POSTCOPY",
		"VIR_DOMAIN_EVENT_SUSPENDED_POSTCOPY_FAILED",
		"VIR_DOMAIN_EVENT_UNDEFINED_RENAMED",
		"VIR_DOMAIN_EVENT_WATCHDOG_INJECTNMI",
		"VIR_DOMAIN_INTERFACE_ADDRESSES_SRC_AGENT",
		"VIR_DOMAIN_INTERFACE_ADDRESSES_SRC_LEASE",
		"VIR_DOMAIN_JOB_BOUNDED",
		"VIR_DOMAIN_JOB_CANCELLED",
		"VIR_DOMAIN_JOB_COMPLETED",
		"VIR_DOMAIN_JOB_FAILED",
		"VIR_DOMAIN_JOB_NONE",
		"VIR_DOMAIN_JOB_STATS_COMPLETED",
		"VIR_DOMAIN_JOB_UNBOUNDED",
		"VIR_DOMAIN_MEMORY_PARAM_BOOLEAN",
		"VIR_DOMAIN_MEMORY_PARAM_DOUBLE",
		"VIR_DOMAIN_MEMORY_PARAM_INT",
		"VIR_DOMAIN_MEMORY_PARAM_LLONG",
		"VIR_DOMAIN_MEMORY_PARAM_UINT",
		"VIR_DOMAIN_MEMORY_PARAM_ULLONG",
		"VIR_DOMAIN_MEMORY_STAT_LAST_UPDATE",
		"VIR_DOMAIN_MEMORY_STAT_USABLE",
		"VIR_DOMAIN_MEM_CONFIG",
		"VIR_DOMAIN_MEM_CURRENT",
		"VIR_DOMAIN_MEM_LIVE",
		"VIR_DOMAIN_NOSTATE_UNKNOWN",
		"VIR_DOMAIN_NUMATUNE_MEM_INTERLEAVE",
		"VIR_DOMAIN_NUMATUNE_MEM_PREFERRED",
		"VIR_DOMAIN_NUMATUNE_MEM_STRICT",
		"VIR_DOMAIN_OPEN_GRAPHICS_SKIPAUTH",
		"VIR_DOMAIN_PASSWORD_ENCRYPTED",
		"VIR_DOMAIN_PAUSED_POSTCOPY",
		"VIR_DOMAIN_PAUSED_POSTCOPY_FAILED",
		"VIR_DOMAIN_PAUSED_STARTING_UP",
		"VIR_DOMAIN_PMSUSPENDED_DISK_UNKNOWN",
		"VIR_DOMAIN_PMSUSPENDED_UNKNOWN",
		"VIR_DOMAIN_REBOOT_ACPI_POWER_BTN",
		"VIR_DOMAIN_REBOOT_DEFAULT",
		"VIR_DOMAIN_REBOOT_GUEST_AGENT",
		"VIR_DOMAIN_REBOOT_INITCTL",
		"VIR_DOMAIN_REBOOT_PARAVIRT",
		"VIR_DOMAIN_REBOOT_SIGNAL",
		"VIR_DOMAIN_RUNNING_POSTCOPY",
		"VIR_DOMAIN_SAVE_BYPASS_CACHE",
		"VIR_DOMAIN_SAVE_PAUSED",
		"VIR_DOMAIN_SAVE_RUNNING",
		"VIR_DOMAIN_SCHED_FIELD_BOOLEAN",
		"VIR_DOMAIN_SCHED_FIELD_DOUBLE",
		"VIR_DOMAIN_SCHED_FIELD_INT",
		"VIR_DOMAIN_SCHED_FIELD_LLONG",
		"VIR_DOMAIN_SCHED_FIELD_UINT",
		"VIR_DOMAIN_SCHED_FIELD_ULLONG",
		"VIR_DOMAIN_SHUTDOWN_PARAVIRT",
		"VIR_DOMAIN_SHUTDOWN_UNKNOWN",
		"VIR_DOMAIN_SHUTDOWN_USER",
		"VIR_DOMAIN_SHUTOFF_CRASHED",
		"VIR_DOMAIN_SHUTOFF_DESTROYED",
		"VIR_DOMAIN_SHUTOFF_FAILED",
		"VIR_DOMAIN_SHUTOFF_FROM_SNAPSHOT",
		"VIR_DOMAIN_SHUTOFF_MIGRATED",
		"VIR_DOMAIN_SHUTOFF_SAVED",
		"VIR_DOMAIN_SHUTOFF_SHUTDOWN",
		"VIR_DOMAIN_SHUTOFF_UNKNOWN",
		"VIR_DOMAIN_SNAPSHOT_CREATE_ATOMIC",
		"VIR_DOMAIN_SNAPSHOT_CREATE_CURRENT",
		"VIR_DOMAIN_SNAPSHOT_CREATE_DISK_ONLY",
		"VIR_DOMAIN_SNAPSHOT_CREATE_HALT",
		"VIR_DOMAIN_SNAPSHOT_CREATE_LIVE",
		"VIR_DOMAIN_SNAPSHOT_CREATE_NO_METADATA",
		"VIR_DOMAIN_SNAPSHOT_CREATE_QUIESCE",
		"VIR_DOMAIN_SNAPSHOT_CREATE_REDEFINE",
		"VIR_DOMAIN_SNAPSHOT_CREATE_REUSE_EXT",
		"VIR_DOMAIN_SNAPSHOT_DELETE_CHILDREN",
		"VIR_DOMAIN_SNAPSHOT_DELETE_CHILDREN_ONLY",
		"VIR_DOMAIN_SNAPSHOT_DELETE_METADATA_ONLY",
		"VIR_DOMAIN_SNAPSHOT_LIST_ACTIVE",
		"VIR_DOMAIN_SNAPSHOT_LIST_DESCENDANTS",
		"VIR_DOMAIN_SNAPSHOT_LIST_DISK_ONLY",
		"VIR_DOMAIN_SNAPSHOT_LIST_EXTERNAL",
		"VIR_DOMAIN_SNAPSHOT_LIST_INACTIVE",
		"VIR_DOMAIN_SNAPSHOT_LIST_INTERNAL",
		"VIR_DOMAIN_SNAPSHOT_LIST_LEAVES",
		"VIR_DOMAIN_SNAPSHOT_LIST_METADATA",
		"VIR_DOMAIN_SNAPSHOT_LIST_NO_LEAVES",
		"VIR_DOMAIN_SNAPSHOT_LIST_NO_METADATA",
		"VIR_DOMAIN_SNAPSHOT_LIST_ROOTS",
		"VIR_DOMAIN_SNAPSHOT_REVERT_FORCE",
		"VIR_DOMAIN_SNAPSHOT_REVERT_PAUSED",
		"VIR_DOMAIN_SNAPSHOT_REVERT_RUNNING",
		"VIR_DOMAIN_START_VALIDATE",
		"VIR_DOMAIN_STATS_BALLOON",
		"VIR_DOMAIN_STATS_BLOCK",
		"VIR_DOMAIN_STATS_CPU_TOTAL",
		"VIR_DOMAIN_STATS_INTERFACE",
		"VIR_DOMAIN_STATS_PERF",
		"VIR_DOMAIN_STATS_STATE",
		"VIR_DOMAIN_STATS_VCPU",
		"VIR_DOMAIN_TIME_SYNC",
		"VIR_DOMAIN_UNDEFINE_KEEP_NVRAM",
		"VIR_DOMAIN_VCPU_HOTPLUGGABLE",
		"VIR_DUMP_BYPASS_CACHE",
		"VIR_DUMP_CRASH",
		"VIR_DUMP_LIVE",
		"VIR_DUMP_MEMORY_ONLY",
		"VIR_DUMP_RESET",
		"VIR_ERR_AGENT_UNSYNCED",
		"VIR_ERR_AUTH_UNAVAILABLE",
		"VIR_ERR_MIGRATE_FINISH_OK",
		"VIR_ERR_NO_CLIENT",
		"VIR_ERR_NO_SERVER",
		"VIR_ERR_XML_INVALID_SCHEMA",
		"VIR_EVENT_HANDLE_ERROR",
		"VIR_EVENT_HANDLE_HANGUP",
		"VIR_EVENT_HANDLE_READABLE",
		"VIR_EVENT_HANDLE_WRITABLE",
		"VIR_FROM_ADMIN",
		"VIR_FROM_LOGGING",
		"VIR_FROM_PERF",
		"VIR_FROM_POLKIT",
		"VIR_FROM_THREAD",
		"VIR_FROM_XENXL",
		"VIR_INTERFACE_XML_INACTIVE",
		"VIR_MEMORY_PHYSICAL",
		"VIR_MEMORY_VIRTUAL",
		"VIR_MIGRATE_ABORT_ON_ERROR",
		"VIR_MIGRATE_AUTO_CONVERGE",
		"VIR_MIGRATE_CHANGE_PROTECTION",
		"VIR_MIGRATE_COMPRESSED",
		"VIR_MIGRATE_LIVE",
		"VIR_MIGRATE_NON_SHARED_DISK",
		"VIR_MIGRATE_NON_SHARED_INC",
		"VIR_MIGRATE_OFFLINE",
		"VIR_MIGRATE_PAUSED",
		"VIR_MIGRATE_PEER2PEER",
		"VIR_MIGRATE_PERSIST_DEST",
		"VIR_MIGRATE_POSTCOPY",
		"VIR_MIGRATE_RDMA_PIN_ALL",
		"VIR_MIGRATE_TUNNELLED",
		"VIR_MIGRATE_UNDEFINE_SOURCE",
		"VIR_MIGRATE_UNSAFE",
		"VIR_NETWORK_EVENT_DEFINED",
		"VIR_NETWORK_EVENT_ID_LIFECYCLE",
		"VIR_NETWORK_EVENT_STARTED",
		"VIR_NETWORK_EVENT_STOPPED",
		"VIR_NETWORK_EVENT_UNDEFINED",
		"VIR_NETWORK_SECTION_BRIDGE",
		"VIR_NETWORK_SECTION_DNS_HOST",
		"VIR_NETWORK_SECTION_DNS_SRV",
		"VIR_NETWORK_SECTION_DNS_TXT",
		"VIR_NETWORK_SECTION_DOMAIN",
		"VIR_NETWORK_SECTION_FORWARD",
		"VIR_NETWORK_SECTION_FORWARD_INTERFACE",
		"VIR_NETWORK_SECTION_FORWARD_PF",
		"VIR_NETWORK_SECTION_IP",
		"VIR_NETWORK_SECTION_IP_DHCP_HOST",
		"VIR_NETWORK_SECTION_IP_DHCP_RANGE",
		"VIR_NETWORK_SECTION_NONE",
		"VIR_NETWORK_SECTION_PORTGROUP",
		"VIR_NETWORK_UPDATE_AFFECT_CONFIG",
		"VIR_NETWORK_UPDATE_AFFECT_CURRENT",
		"VIR_NETWORK_UPDATE_AFFECT_LIVE",
		"VIR_NETWORK_UPDATE_COMMAND_ADD_FIRST",
		"VIR_NETWORK_UPDATE_COMMAND_DELETE",
		"VIR_NETWORK_UPDATE_COMMAND_MODIFY",
		"VIR_NETWORK_UPDATE_COMMAND_NONE",
		"VIR_NETWORK_XML_INACTIVE",
		"VIR_NODE_ALLOC_PAGES_ADD",
		"VIR_NODE_ALLOC_PAGES_SET",
		"VIR_NODE_CPU_STATS_ALL_CPUS",
		"VIR_NODE_DEVICE_EVENT_CREATED",
		"VIR_NODE_DEVICE_EVENT_DELETED",
		"VIR_NODE_DEVICE_EVENT_ID_LIFECYCLE",
		"VIR_NODE_DEVICE_EVENT_ID_UPDATE",
		"VIR_NODE_MEMORY_STATS_ALL_CELLS",
		"VIR_NODE_SUSPEND_TARGET_DISK",
		"VIR_NODE_SUSPEND_TARGET_HYBRID",
		"VIR_NODE_SUSPEND_TARGET_MEM",
		"VIR_SECRET_USAGE_TYPE_TLS",
		"VIR_STORAGE_POOL_CREATE_NORMAL",
		"VIR_STORAGE_POOL_CREATE_WITH_BUILD",
		"VIR_STORAGE_POOL_CREATE_WITH_BUILD_NO_OVERWRITE",
		"VIR_STORAGE_POOL_CREATE_WITH_BUILD_OVERWRITE",
		"VIR_STORAGE_POOL_DELETE_NORMAL",
		"VIR_STORAGE_POOL_DELETE_ZEROED",
		"VIR_STORAGE_POOL_EVENT_DEFINED",
		"VIR_STORAGE_POOL_EVENT_ID_LIFECYCLE",
		"VIR_STORAGE_POOL_EVENT_ID_REFRESH",
		"VIR_STORAGE_POOL_EVENT_STARTED",
		"VIR_STORAGE_POOL_EVENT_STOPPED",
		"VIR_STORAGE_POOL_EVENT_UNDEFINED",
		"VIR_STORAGE_VOL_CREATE_REFLINK",
		"VIR_STORAGE_VOL_DELETE_WITH_SNAPSHOTS",
		"VIR_STORAGE_VOL_PLOOP",
		"VIR_STORAGE_VOL_WIPE_ALG_TRIM",
		"VIR_STORAGE_XML_INACTIVE",
		"VIR_STREAM_EVENT_ERROR",
		"VIR_STREAM_EVENT_HANGUP",
		"VIR_STREAM_EVENT_READABLE",
		"VIR_STREAM_EVENT_WRITABLE",
		"VIR_TYPED_PARAM_STRING_OKAY",
		"VIR_VCPU_BLOCKED",
		"VIR_VCPU_OFFLINE",
		"VIR_VCPU_RUNNING",
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
