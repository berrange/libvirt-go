package libvirt

/*
#cgo LDFLAGS: -lvirt-qemu -lvirt
#include <libvirt/libvirt.h>
#include <libvirt/libvirt-qemu.h>
#include <libvirt/virterror.h>
#include <stdlib.h>
*/
import "C"

import (
	"reflect"
	"strings"
	"unsafe"
)

type VirDomainState int

const (
	VIR_DOMAIN_NOSTATE     = VirDomainState(C.VIR_DOMAIN_NOSTATE)
	VIR_DOMAIN_RUNNING     = VirDomainState(C.VIR_DOMAIN_RUNNING)
	VIR_DOMAIN_BLOCKED     = VirDomainState(C.VIR_DOMAIN_BLOCKED)
	VIR_DOMAIN_PAUSED      = VirDomainState(C.VIR_DOMAIN_PAUSED)
	VIR_DOMAIN_SHUTDOWN    = VirDomainState(C.VIR_DOMAIN_SHUTDOWN)
	VIR_DOMAIN_CRASHED     = VirDomainState(C.VIR_DOMAIN_CRASHED)
	VIR_DOMAIN_PMSUSPENDED = VirDomainState(C.VIR_DOMAIN_PMSUSPENDED)
	VIR_DOMAIN_SHUTOFF     = VirDomainState(C.VIR_DOMAIN_SHUTOFF)
)

type VirDomainMetadataType int

const (
	VIR_DOMAIN_METADATA_DESCRIPTION = VirDomainMetadataType(C.VIR_DOMAIN_METADATA_DESCRIPTION)
	VIR_DOMAIN_METADATA_TITLE       = VirDomainMetadataType(C.VIR_DOMAIN_METADATA_TITLE)
	VIR_DOMAIN_METADATA_ELEMENT     = VirDomainMetadataType(C.VIR_DOMAIN_METADATA_ELEMENT)
)

type VirDomainVcpuFlags int

const (
	VIR_DOMAIN_VCPU_CONFIG  = VirDomainVcpuFlags(C.VIR_DOMAIN_VCPU_CONFIG)
	VIR_DOMAIN_VCPU_CURRENT = VirDomainVcpuFlags(C.VIR_DOMAIN_VCPU_CURRENT)
	VIR_DOMAIN_VCPU_LIVE    = VirDomainVcpuFlags(C.VIR_DOMAIN_VCPU_LIVE)
	VIR_DOMAIN_VCPU_MAXIMUM = VirDomainVcpuFlags(C.VIR_DOMAIN_VCPU_MAXIMUM)
	VIR_DOMAIN_VCPU_GUEST   = VirDomainVcpuFlags(C.VIR_DOMAIN_VCPU_GUEST)
)

type VirDomainModificationImpact int

const (
	VIR_DOMAIN_AFFECT_CONFIG  = VirDomainModificationImpact(C.VIR_DOMAIN_AFFECT_CONFIG)
	VIR_DOMAIN_AFFECT_CURRENT = VirDomainModificationImpact(C.VIR_DOMAIN_AFFECT_CURRENT)
	VIR_DOMAIN_AFFECT_LIVE    = VirDomainModificationImpact(C.VIR_DOMAIN_AFFECT_LIVE)
)

type VirDomainMemoryModFlags int

const (
	VIR_DOMAIN_MEM_CONFIG  = VirDomainMemoryModFlags(C.VIR_DOMAIN_AFFECT_CONFIG)
	VIR_DOMAIN_MEM_CURRENT = VirDomainMemoryModFlags(C.VIR_DOMAIN_AFFECT_CURRENT)
	VIR_DOMAIN_MEM_LIVE    = VirDomainMemoryModFlags(C.VIR_DOMAIN_AFFECT_LIVE)
	VIR_DOMAIN_MEM_MAXIMUM = VirDomainMemoryModFlags(C.VIR_DOMAIN_MEM_MAXIMUM)
)

type VirDomainDestroyFlags int

const (
	VIR_DOMAIN_DESTROY_DEFAULT  = VirDomainDestroyFlags(C.VIR_DOMAIN_DESTROY_DEFAULT)
	VIR_DOMAIN_DESTROY_GRACEFUL = VirDomainDestroyFlags(C.VIR_DOMAIN_DESTROY_GRACEFUL)
)

type VirDomainShutdownFlags int

const (
	VIR_DOMAIN_SHUTDOWN_DEFAULT        = VirDomainShutdownFlags(C.VIR_DOMAIN_SHUTDOWN_DEFAULT)
	VIR_DOMAIN_SHUTDOWN_ACPI_POWER_BTN = VirDomainShutdownFlags(C.VIR_DOMAIN_SHUTDOWN_ACPI_POWER_BTN)
	VIR_DOMAIN_SHUTDOWN_GUEST_AGENT    = VirDomainShutdownFlags(C.VIR_DOMAIN_SHUTDOWN_GUEST_AGENT)
	VIR_DOMAIN_SHUTDOWN_INITCTL        = VirDomainShutdownFlags(C.VIR_DOMAIN_SHUTDOWN_INITCTL)
	VIR_DOMAIN_SHUTDOWN_SIGNAL         = VirDomainShutdownFlags(C.VIR_DOMAIN_SHUTDOWN_SIGNAL)
)

type VirDomainUndefineFlags int

const (
	VIR_DOMAIN_UNDEFINE_MANAGED_SAVE       = VirDomainUndefineFlags(C.VIR_DOMAIN_UNDEFINE_MANAGED_SAVE)       // Also remove any managed save
	VIR_DOMAIN_UNDEFINE_SNAPSHOTS_METADATA = VirDomainUndefineFlags(C.VIR_DOMAIN_UNDEFINE_SNAPSHOTS_METADATA) // If last use of domain, then also remove any snapshot metadata
	VIR_DOMAIN_UNDEFINE_NVRAM              = VirDomainUndefineFlags(C.VIR_DOMAIN_UNDEFINE_NVRAM)              // Also remove any nvram file
)

type VirDomainAttachDeviceFlags int

const (
	VIR_DOMAIN_DEVICE_MODIFY_CONFIG  = VirDomainAttachDeviceFlags(C.VIR_DOMAIN_AFFECT_CONFIG)
	VIR_DOMAIN_DEVICE_MODIFY_CURRENT = VirDomainAttachDeviceFlags(C.VIR_DOMAIN_AFFECT_CURRENT)
	VIR_DOMAIN_DEVICE_MODIFY_LIVE    = VirDomainAttachDeviceFlags(C.VIR_DOMAIN_AFFECT_LIVE)
	VIR_DOMAIN_DEVICE_MODIFY_FORCE   = VirDomainAttachDeviceFlags(C.VIR_DOMAIN_DEVICE_MODIFY_FORCE)
)

type VirDomainCreateFlags int

const (
	VIR_DOMAIN_NONE               = VirDomainCreateFlags(C.VIR_DOMAIN_NONE)
	VIR_DOMAIN_START_PAUSED       = VirDomainCreateFlags(C.VIR_DOMAIN_START_PAUSED)
	VIR_DOMAIN_START_AUTODESTROY  = VirDomainCreateFlags(C.VIR_DOMAIN_START_AUTODESTROY)
	VIR_DOMAIN_START_BYPASS_CACHE = VirDomainCreateFlags(C.VIR_DOMAIN_START_BYPASS_CACHE)
	VIR_DOMAIN_START_FORCE_BOOT   = VirDomainCreateFlags(C.VIR_DOMAIN_START_FORCE_BOOT)
)

const VIR_DOMAIN_MEMORY_PARAM_UNLIMITED = C.VIR_DOMAIN_MEMORY_PARAM_UNLIMITED

type VirDomainEventID int

const (
	// event parameter in the callback is of type DomainLifecycleEvent
	VIR_DOMAIN_EVENT_ID_LIFECYCLE = VirDomainEventID(C.VIR_DOMAIN_EVENT_ID_LIFECYCLE)

	// event parameter in the callback is nil
	VIR_DOMAIN_EVENT_ID_REBOOT = VirDomainEventID(C.VIR_DOMAIN_EVENT_ID_REBOOT)

	// event parameter in the callback is of type DomainRTCChangeEvent
	VIR_DOMAIN_EVENT_ID_RTC_CHANGE = VirDomainEventID(C.VIR_DOMAIN_EVENT_ID_RTC_CHANGE)

	// event parameter in the callback is of type DomainWatchdogEvent
	VIR_DOMAIN_EVENT_ID_WATCHDOG = VirDomainEventID(C.VIR_DOMAIN_EVENT_ID_WATCHDOG)

	// event parameter in the callback is of type DomainIOErrorEvent
	VIR_DOMAIN_EVENT_ID_IO_ERROR = VirDomainEventID(C.VIR_DOMAIN_EVENT_ID_IO_ERROR)

	// event parameter in the callback is of type DomainGraphicsEvent
	VIR_DOMAIN_EVENT_ID_GRAPHICS = VirDomainEventID(C.VIR_DOMAIN_EVENT_ID_GRAPHICS)

	// virConnectDomainEventIOErrorReasonCallback
	VIR_DOMAIN_EVENT_ID_IO_ERROR_REASON = VirDomainEventID(C.VIR_DOMAIN_EVENT_ID_IO_ERROR_REASON)

	// event parameter in the callback is nil
	VIR_DOMAIN_EVENT_ID_CONTROL_ERROR = VirDomainEventID(C.VIR_DOMAIN_EVENT_ID_CONTROL_ERROR)

	// event parameter in the callback is of type DomainBlockJobEvent
	VIR_DOMAIN_EVENT_ID_BLOCK_JOB = VirDomainEventID(C.VIR_DOMAIN_EVENT_ID_BLOCK_JOB)

	// event parameter in the callback is of type DomainDiskChangeEvent
	VIR_DOMAIN_EVENT_ID_DISK_CHANGE = VirDomainEventID(C.VIR_DOMAIN_EVENT_ID_DISK_CHANGE)

	// event parameter in the callback is of type DomainTrayChangeEvent
	VIR_DOMAIN_EVENT_ID_TRAY_CHANGE = VirDomainEventID(C.VIR_DOMAIN_EVENT_ID_TRAY_CHANGE)

	// event parameter in the callback is of type DomainReasonEvent
	VIR_DOMAIN_EVENT_ID_PMWAKEUP = VirDomainEventID(C.VIR_DOMAIN_EVENT_ID_PMWAKEUP)

	// event parameter in the callback is of type DomainReasonEvent
	VIR_DOMAIN_EVENT_ID_PMSUSPEND = VirDomainEventID(C.VIR_DOMAIN_EVENT_ID_PMSUSPEND)

	// event parameter in the callback is of type DomainBalloonChangeEvent
	VIR_DOMAIN_EVENT_ID_BALLOON_CHANGE = VirDomainEventID(C.VIR_DOMAIN_EVENT_ID_BALLOON_CHANGE)

	// event parameter in the callback is of type DomainReasonEvent
	VIR_DOMAIN_EVENT_ID_PMSUSPEND_DISK = VirDomainEventID(C.VIR_DOMAIN_EVENT_ID_PMSUSPEND_DISK)

	// event parameter in the callback is of type DomainDeviceRemovedEvent
	VIR_DOMAIN_EVENT_ID_DEVICE_REMOVED = VirDomainEventID(C.VIR_DOMAIN_EVENT_ID_DEVICE_REMOVED)

	// event parameter in the callback is of type DomainBlockJobEvent
	VIR_DOMAIN_EVENT_ID_BLOCK_JOB_2 = VirDomainEventID(C.VIR_DOMAIN_EVENT_ID_BLOCK_JOB_2)
)

type VirDomainEventType int

const (
	VIR_DOMAIN_EVENT_DEFINED     = VirDomainEventType(C.VIR_DOMAIN_EVENT_DEFINED)
	VIR_DOMAIN_EVENT_UNDEFINED   = VirDomainEventType(C.VIR_DOMAIN_EVENT_UNDEFINED)
	VIR_DOMAIN_EVENT_STARTED     = VirDomainEventType(C.VIR_DOMAIN_EVENT_STARTED)
	VIR_DOMAIN_EVENT_SUSPENDED   = VirDomainEventType(C.VIR_DOMAIN_EVENT_SUSPENDED)
	VIR_DOMAIN_EVENT_RESUMED     = VirDomainEventType(C.VIR_DOMAIN_EVENT_RESUMED)
	VIR_DOMAIN_EVENT_STOPPED     = VirDomainEventType(C.VIR_DOMAIN_EVENT_STOPPED)
	VIR_DOMAIN_EVENT_SHUTDOWN    = VirDomainEventType(C.VIR_DOMAIN_EVENT_SHUTDOWN)
	VIR_DOMAIN_EVENT_PMSUSPENDED = VirDomainEventType(C.VIR_DOMAIN_EVENT_PMSUSPENDED)
	VIR_DOMAIN_EVENT_CRASHED     = VirDomainEventType(C.VIR_DOMAIN_EVENT_CRASHED)
)

type VirDomainEventWatchdogAction int

// The action that is to be taken due to the watchdog device firing
const (
	// No action, watchdog ignored
	VIR_DOMAIN_EVENT_WATCHDOG_NONE = VirDomainEventWatchdogAction(C.VIR_DOMAIN_EVENT_WATCHDOG_NONE)

	// Guest CPUs are paused
	VIR_DOMAIN_EVENT_WATCHDOG_PAUSE = VirDomainEventWatchdogAction(C.VIR_DOMAIN_EVENT_WATCHDOG_PAUSE)

	// Guest CPUs are reset
	VIR_DOMAIN_EVENT_WATCHDOG_RESET = VirDomainEventWatchdogAction(C.VIR_DOMAIN_EVENT_WATCHDOG_RESET)

	// Guest is forcibly powered off
	VIR_DOMAIN_EVENT_WATCHDOG_POWEROFF = VirDomainEventWatchdogAction(C.VIR_DOMAIN_EVENT_WATCHDOG_POWEROFF)

	// Guest is requested to gracefully shutdown
	VIR_DOMAIN_EVENT_WATCHDOG_SHUTDOWN = VirDomainEventWatchdogAction(C.VIR_DOMAIN_EVENT_WATCHDOG_SHUTDOWN)

	// No action, a debug message logged
	VIR_DOMAIN_EVENT_WATCHDOG_DEBUG = VirDomainEventWatchdogAction(C.VIR_DOMAIN_EVENT_WATCHDOG_DEBUG)
)

type VirDomainEventIOErrorAction int

// The action that is to be taken due to an IO error occurring
const (
	// No action, IO error ignored
	VIR_DOMAIN_EVENT_IO_ERROR_NONE = VirDomainEventIOErrorAction(C.VIR_DOMAIN_EVENT_IO_ERROR_NONE)

	// Guest CPUs are paused
	VIR_DOMAIN_EVENT_IO_ERROR_PAUSE = VirDomainEventIOErrorAction(C.VIR_DOMAIN_EVENT_IO_ERROR_PAUSE)

	// IO error reported to guest OS
	VIR_DOMAIN_EVENT_IO_ERROR_REPORT = VirDomainEventIOErrorAction(C.VIR_DOMAIN_EVENT_IO_ERROR_REPORT)
)

type VirDomainEventGraphicsPhase int

// The phase of the graphics client connection
const (
	// Initial socket connection established
	VIR_DOMAIN_EVENT_GRAPHICS_CONNECT = VirDomainEventGraphicsPhase(C.VIR_DOMAIN_EVENT_GRAPHICS_CONNECT)

	// Authentication & setup completed
	VIR_DOMAIN_EVENT_GRAPHICS_INITIALIZE = VirDomainEventGraphicsPhase(C.VIR_DOMAIN_EVENT_GRAPHICS_INITIALIZE)

	// Final socket disconnection
	VIR_DOMAIN_EVENT_GRAPHICS_DISCONNECT = VirDomainEventGraphicsPhase(C.VIR_DOMAIN_EVENT_GRAPHICS_DISCONNECT)
)

type VirDomainEventGraphicsAddressType int

const (
	// IPv4 address
	VIR_DOMAIN_EVENT_GRAPHICS_ADDRESS_IPV4 = VirDomainEventGraphicsAddressType(C.VIR_DOMAIN_EVENT_GRAPHICS_ADDRESS_IPV4)

	// IPv6 address
	VIR_DOMAIN_EVENT_GRAPHICS_ADDRESS_IPV6 = VirDomainEventGraphicsAddressType(C.VIR_DOMAIN_EVENT_GRAPHICS_ADDRESS_IPV6)

	// UNIX socket path
	VIR_DOMAIN_EVENT_GRAPHICS_ADDRESS_UNIX = VirDomainEventGraphicsAddressType(C.VIR_DOMAIN_EVENT_GRAPHICS_ADDRESS_UNIX)
)

type VirDomainBlockJobType int

const (
	// Placeholder
	VIR_DOMAIN_BLOCK_JOB_TYPE_UNKNOWN = VirDomainBlockJobType(C.VIR_DOMAIN_BLOCK_JOB_TYPE_UNKNOWN)

	// Block Pull (virDomainBlockPull, or virDomainBlockRebase without
	// flags), job ends on completion
	VIR_DOMAIN_BLOCK_JOB_TYPE_PULL = VirDomainBlockJobType(C.VIR_DOMAIN_BLOCK_JOB_TYPE_PULL)

	// Block Copy (virDomainBlockCopy, or virDomainBlockRebase with
	// flags), job exists as long as mirroring is active
	VIR_DOMAIN_BLOCK_JOB_TYPE_COPY = VirDomainBlockJobType(C.VIR_DOMAIN_BLOCK_JOB_TYPE_COPY)

	// Block Commit (virDomainBlockCommit without flags), job ends on
	// completion
	VIR_DOMAIN_BLOCK_JOB_TYPE_COMMIT = VirDomainBlockJobType(C.VIR_DOMAIN_BLOCK_JOB_TYPE_COMMIT)

	// Active Block Commit (virDomainBlockCommit with flags), job
	// exists as long as sync is active
	VIR_DOMAIN_BLOCK_JOB_TYPE_ACTIVE_COMMIT = VirDomainBlockJobType(C.VIR_DOMAIN_BLOCK_JOB_TYPE_ACTIVE_COMMIT)
)

type VirDomainRunningReason int

const (
	VIR_DOMAIN_RUNNING_UNKNOWN            = VirDomainRunningReason(C.VIR_DOMAIN_RUNNING_UNKNOWN)
	VIR_DOMAIN_RUNNING_BOOTED             = VirDomainRunningReason(C.VIR_DOMAIN_RUNNING_BOOTED)             /* normal startup from boot */
	VIR_DOMAIN_RUNNING_MIGRATED           = VirDomainRunningReason(C.VIR_DOMAIN_RUNNING_MIGRATED)           /* migrated from another host */
	VIR_DOMAIN_RUNNING_RESTORED           = VirDomainRunningReason(C.VIR_DOMAIN_RUNNING_RESTORED)           /* restored from a state file */
	VIR_DOMAIN_RUNNING_FROM_SNAPSHOT      = VirDomainRunningReason(C.VIR_DOMAIN_RUNNING_FROM_SNAPSHOT)      /* restored from snapshot */
	VIR_DOMAIN_RUNNING_UNPAUSED           = VirDomainRunningReason(C.VIR_DOMAIN_RUNNING_UNPAUSED)           /* returned from paused state */
	VIR_DOMAIN_RUNNING_MIGRATION_CANCELED = VirDomainRunningReason(C.VIR_DOMAIN_RUNNING_MIGRATION_CANCELED) /* returned from migration */
	VIR_DOMAIN_RUNNING_SAVE_CANCELED      = VirDomainRunningReason(C.VIR_DOMAIN_RUNNING_SAVE_CANCELED)      /* returned from failed save process */
	VIR_DOMAIN_RUNNING_WAKEUP             = VirDomainRunningReason(C.VIR_DOMAIN_RUNNING_WAKEUP)             /* returned from pmsuspended due to wakeup event */
	VIR_DOMAIN_RUNNING_CRASHED            = VirDomainRunningReason(C.VIR_DOMAIN_RUNNING_CRASHED)            /* resumed from crashed */
)

type VirDomainPausedReason int

const (
	VIR_DOMAIN_PAUSED_UNKNOWN       = VirDomainPausedReason(C.VIR_DOMAIN_PAUSED_UNKNOWN)       /* the reason is unknown */
	VIR_DOMAIN_PAUSED_USER          = VirDomainPausedReason(C.VIR_DOMAIN_PAUSED_USER)          /* paused on user request */
	VIR_DOMAIN_PAUSED_MIGRATION     = VirDomainPausedReason(C.VIR_DOMAIN_PAUSED_MIGRATION)     /* paused for offline migration */
	VIR_DOMAIN_PAUSED_SAVE          = VirDomainPausedReason(C.VIR_DOMAIN_PAUSED_SAVE)          /* paused for save */
	VIR_DOMAIN_PAUSED_DUMP          = VirDomainPausedReason(C.VIR_DOMAIN_PAUSED_DUMP)          /* paused for offline core dump */
	VIR_DOMAIN_PAUSED_IOERROR       = VirDomainPausedReason(C.VIR_DOMAIN_PAUSED_IOERROR)       /* paused due to a disk I/O error */
	VIR_DOMAIN_PAUSED_WATCHDOG      = VirDomainPausedReason(C.VIR_DOMAIN_PAUSED_WATCHDOG)      /* paused due to a watchdog event */
	VIR_DOMAIN_PAUSED_FROM_SNAPSHOT = VirDomainPausedReason(C.VIR_DOMAIN_PAUSED_FROM_SNAPSHOT) /* paused after restoring from snapshot */
	VIR_DOMAIN_PAUSED_SHUTTING_DOWN = VirDomainPausedReason(C.VIR_DOMAIN_PAUSED_SHUTTING_DOWN) /* paused during shutdown process */
	VIR_DOMAIN_PAUSED_SNAPSHOT      = VirDomainPausedReason(C.VIR_DOMAIN_PAUSED_SNAPSHOT)      /* paused while creating a snapshot */
	VIR_DOMAIN_PAUSED_CRASHED       = VirDomainPausedReason(C.VIR_DOMAIN_PAUSED_CRASHED)       /* paused due to a guest crash */
)

type VirDomainXMLFlags int

const (
	VIR_DOMAIN_XML_SECURE     = VirDomainXMLFlags(C.VIR_DOMAIN_XML_SECURE)     /* dump security sensitive information too */
	VIR_DOMAIN_XML_INACTIVE   = VirDomainXMLFlags(C.VIR_DOMAIN_XML_INACTIVE)   /* dump inactive domain information */
	VIR_DOMAIN_XML_UPDATE_CPU = VirDomainXMLFlags(C.VIR_DOMAIN_XML_UPDATE_CPU) /* update guest CPU requirements according to host CPU */
	VIR_DOMAIN_XML_MIGRATABLE = VirDomainXMLFlags(C.VIR_DOMAIN_XML_MIGRATABLE) /* dump XML suitable for migration */
)

type VirDomainEventDefinedDetailType int

const (
	VIR_DOMAIN_EVENT_DEFINED_ADDED   = VirDomainEventDefinedDetailType(C.VIR_DOMAIN_EVENT_DEFINED_ADDED)
	VIR_DOMAIN_EVENT_DEFINED_UPDATED = VirDomainEventDefinedDetailType(C.VIR_DOMAIN_EVENT_DEFINED_UPDATED)
)

type VirDomainEventUndefinedDetailType int

const (
	VIR_DOMAIN_EVENT_UNDEFINED_REMOVED = VirDomainEventUndefinedDetailType(C.VIR_DOMAIN_EVENT_UNDEFINED_REMOVED)
)

type VirDomainEventStartedDetailType int

const (
	VIR_DOMAIN_EVENT_STARTED_BOOTED        = VirDomainEventStartedDetailType(C.VIR_DOMAIN_EVENT_STARTED_BOOTED)
	VIR_DOMAIN_EVENT_STARTED_MIGRATED      = VirDomainEventStartedDetailType(C.VIR_DOMAIN_EVENT_STARTED_MIGRATED)
	VIR_DOMAIN_EVENT_STARTED_RESTORED      = VirDomainEventStartedDetailType(C.VIR_DOMAIN_EVENT_STARTED_RESTORED)
	VIR_DOMAIN_EVENT_STARTED_FROM_SNAPSHOT = VirDomainEventStartedDetailType(C.VIR_DOMAIN_EVENT_STARTED_FROM_SNAPSHOT)
	VIR_DOMAIN_EVENT_STARTED_WAKEUP        = VirDomainEventStartedDetailType(C.VIR_DOMAIN_EVENT_STARTED_WAKEUP)
)

type VirDomainEventSuspendedDetailType int

const (
	VIR_DOMAIN_EVENT_SUSPENDED_PAUSED        = VirDomainEventSuspendedDetailType(C.VIR_DOMAIN_EVENT_SUSPENDED_PAUSED)
	VIR_DOMAIN_EVENT_SUSPENDED_MIGRATED      = VirDomainEventSuspendedDetailType(C.VIR_DOMAIN_EVENT_SUSPENDED_MIGRATED)
	VIR_DOMAIN_EVENT_SUSPENDED_IOERROR       = VirDomainEventSuspendedDetailType(C.VIR_DOMAIN_EVENT_SUSPENDED_IOERROR)
	VIR_DOMAIN_EVENT_SUSPENDED_WATCHDOG      = VirDomainEventSuspendedDetailType(C.VIR_DOMAIN_EVENT_SUSPENDED_WATCHDOG)
	VIR_DOMAIN_EVENT_SUSPENDED_RESTORED      = VirDomainEventSuspendedDetailType(C.VIR_DOMAIN_EVENT_SUSPENDED_RESTORED)
	VIR_DOMAIN_EVENT_SUSPENDED_FROM_SNAPSHOT = VirDomainEventSuspendedDetailType(C.VIR_DOMAIN_EVENT_SUSPENDED_FROM_SNAPSHOT)
	VIR_DOMAIN_EVENT_SUSPENDED_API_ERROR     = VirDomainEventSuspendedDetailType(C.VIR_DOMAIN_EVENT_SUSPENDED_API_ERROR)
)

type VirDomainEventResumedDetailType int

const (
	VIR_DOMAIN_EVENT_RESUMED_UNPAUSED      = VirDomainEventResumedDetailType(C.VIR_DOMAIN_EVENT_RESUMED_UNPAUSED)
	VIR_DOMAIN_EVENT_RESUMED_MIGRATED      = VirDomainEventResumedDetailType(C.VIR_DOMAIN_EVENT_RESUMED_MIGRATED)
	VIR_DOMAIN_EVENT_RESUMED_FROM_SNAPSHOT = VirDomainEventResumedDetailType(C.VIR_DOMAIN_EVENT_RESUMED_FROM_SNAPSHOT)
)

type VirDomainEventStoppedDetailType int

const (
	VIR_DOMAIN_EVENT_STOPPED_SHUTDOWN      = VirDomainEventStoppedDetailType(C.VIR_DOMAIN_EVENT_STOPPED_SHUTDOWN)
	VIR_DOMAIN_EVENT_STOPPED_DESTROYED     = VirDomainEventStoppedDetailType(C.VIR_DOMAIN_EVENT_STOPPED_DESTROYED)
	VIR_DOMAIN_EVENT_STOPPED_CRASHED       = VirDomainEventStoppedDetailType(C.VIR_DOMAIN_EVENT_STOPPED_CRASHED)
	VIR_DOMAIN_EVENT_STOPPED_MIGRATED      = VirDomainEventStoppedDetailType(C.VIR_DOMAIN_EVENT_STOPPED_MIGRATED)
	VIR_DOMAIN_EVENT_STOPPED_SAVED         = VirDomainEventStoppedDetailType(C.VIR_DOMAIN_EVENT_STOPPED_SAVED)
	VIR_DOMAIN_EVENT_STOPPED_FAILED        = VirDomainEventStoppedDetailType(C.VIR_DOMAIN_EVENT_STOPPED_FAILED)
	VIR_DOMAIN_EVENT_STOPPED_FROM_SNAPSHOT = VirDomainEventStoppedDetailType(C.VIR_DOMAIN_EVENT_STOPPED_FROM_SNAPSHOT)
)

type VirDomainEventShutdownDetailType int

const (
	VIR_DOMAIN_EVENT_SHUTDOWN_FINISHED = VirDomainEventShutdownDetailType(C.VIR_DOMAIN_EVENT_SHUTDOWN_FINISHED)
)

type VirDomainMemoryStatTags int

const (
	VIR_DOMAIN_MEMORY_STAT_LAST           = VirDomainMemoryStatTags(C.VIR_DOMAIN_MEMORY_STAT_NR)
	VIR_DOMAIN_MEMORY_STAT_SWAP_IN        = VirDomainMemoryStatTags(C.VIR_DOMAIN_MEMORY_STAT_SWAP_IN)
	VIR_DOMAIN_MEMORY_STAT_SWAP_OUT       = VirDomainMemoryStatTags(C.VIR_DOMAIN_MEMORY_STAT_SWAP_OUT)
	VIR_DOMAIN_MEMORY_STAT_MAJOR_FAULT    = VirDomainMemoryStatTags(C.VIR_DOMAIN_MEMORY_STAT_MAJOR_FAULT)
	VIR_DOMAIN_MEMORY_STAT_MINOR_FAULT    = VirDomainMemoryStatTags(C.VIR_DOMAIN_MEMORY_STAT_MINOR_FAULT)
	VIR_DOMAIN_MEMORY_STAT_UNUSED         = VirDomainMemoryStatTags(C.VIR_DOMAIN_MEMORY_STAT_UNUSED)
	VIR_DOMAIN_MEMORY_STAT_AVAILABLE      = VirDomainMemoryStatTags(C.VIR_DOMAIN_MEMORY_STAT_AVAILABLE)
	VIR_DOMAIN_MEMORY_STAT_ACTUAL_BALLOON = VirDomainMemoryStatTags(C.VIR_DOMAIN_MEMORY_STAT_ACTUAL_BALLOON)
	VIR_DOMAIN_MEMORY_STAT_RSS            = VirDomainMemoryStatTags(C.VIR_DOMAIN_MEMORY_STAT_RSS)
	VIR_DOMAIN_MEMORY_STAT_NR             = VirDomainMemoryStatTags(C.VIR_DOMAIN_MEMORY_STAT_NR)
)

type VirDomainCPUStatsTags string

const (
	VIR_DOMAIN_CPU_STATS_CPUTIME    = VirDomainCPUStatsTags(C.VIR_DOMAIN_CPU_STATS_CPUTIME)
	VIR_DOMAIN_CPU_STATS_SYSTEMTIME = VirDomainCPUStatsTags(C.VIR_DOMAIN_CPU_STATS_SYSTEMTIME)
	VIR_DOMAIN_CPU_STATS_USERTIME   = VirDomainCPUStatsTags(C.VIR_DOMAIN_CPU_STATS_USERTIME)
	VIR_DOMAIN_CPU_STATS_VCPUTIME   = VirDomainCPUStatsTags(C.VIR_DOMAIN_CPU_STATS_VCPUTIME)
)

type VirDomainInterfaceAddressesSource int

const (
	VIR_DOMAIN_INTERFACE_ADDRESSES_SRC_LEASE = VirDomainInterfaceAddressesSource(C.VIR_DOMAIN_INTERFACE_ADDRESSES_SRC_LEASE)
	VIR_DOMAIN_INTERFACE_ADDRESSES_SRC_AGENT = VirDomainInterfaceAddressesSource(C.VIR_DOMAIN_INTERFACE_ADDRESSES_SRC_AGENT)
)

type VirKeycodeSet int

const (
	VIR_KEYCODE_SET_LINUX  = VirKeycodeSet(C.VIR_KEYCODE_SET_LINUX)
	VIR_KEYCODE_SET_XT     = VirKeycodeSet(C.VIR_KEYCODE_SET_XT)
	VIR_KEYCODE_SET_ATSET1 = VirKeycodeSet(C.VIR_KEYCODE_SET_ATSET1)
	VIR_KEYCODE_SET_ATSET2 = VirKeycodeSet(C.VIR_KEYCODE_SET_ATSET2)
	VIR_KEYCODE_SET_ATSET3 = VirKeycodeSet(C.VIR_KEYCODE_SET_ATSET3)
	VIR_KEYCODE_SET_OSX    = VirKeycodeSet(C.VIR_KEYCODE_SET_OSX)
	VIR_KEYCODE_SET_XT_KBD = VirKeycodeSet(C.VIR_KEYCODE_SET_XT_KBD)
	VIR_KEYCODE_SET_USB    = VirKeycodeSet(C.VIR_KEYCODE_SET_USB)
	VIR_KEYCODE_SET_WIN32  = VirKeycodeSet(C.VIR_KEYCODE_SET_WIN32)
	VIR_KEYCODE_SET_RFB    = VirKeycodeSet(C.VIR_KEYCODE_SET_RFB)
)

type VirConnectDomainEventBlockJobStatus int

const (
	VIR_DOMAIN_BLOCK_JOB_COMPLETED = VirConnectDomainEventBlockJobStatus(C.VIR_DOMAIN_BLOCK_JOB_COMPLETED)
	VIR_DOMAIN_BLOCK_JOB_FAILED    = VirConnectDomainEventBlockJobStatus(C.VIR_DOMAIN_BLOCK_JOB_FAILED)
	VIR_DOMAIN_BLOCK_JOB_CANCELED  = VirConnectDomainEventBlockJobStatus(C.VIR_DOMAIN_BLOCK_JOB_CANCELED)
	VIR_DOMAIN_BLOCK_JOB_READY     = VirConnectDomainEventBlockJobStatus(C.VIR_DOMAIN_BLOCK_JOB_READY)
)

type VirConnectDomainEventDiskChangeReason int

const (
	// OldSrcPath is set
	VIR_DOMAIN_EVENT_DISK_CHANGE_MISSING_ON_START = VirConnectDomainEventDiskChangeReason(C.VIR_DOMAIN_EVENT_DISK_CHANGE_MISSING_ON_START)
	VIR_DOMAIN_EVENT_DISK_DROP_MISSING_ON_START   = VirConnectDomainEventDiskChangeReason(C.VIR_DOMAIN_EVENT_DISK_DROP_MISSING_ON_START)
)

type VirConnectDomainEventTrayChangeReason int

const (
	VIR_DOMAIN_EVENT_TRAY_CHANGE_OPEN  = VirConnectDomainEventTrayChangeReason(C.VIR_DOMAIN_EVENT_TRAY_CHANGE_OPEN)
	VIR_DOMAIN_EVENT_TRAY_CHANGE_CLOSE = VirConnectDomainEventTrayChangeReason(C.VIR_DOMAIN_EVENT_TRAY_CHANGE_CLOSE)
)

/*
 * QMP has two different kinds of ways to talk to QEMU. One is legacy (HMP,
 * or 'human' monitor protocol. The default is QMP, which is all-JSON.
 *
 * QMP json commands are of the format:
 * 	{"execute" : "query-cpus"}
 *
 * whereas the same command in 'HMP' would be:
 *	'info cpus'
 */
const (
	VIR_DOMAIN_QEMU_MONITOR_COMMAND_DEFAULT = 0
	VIR_DOMAIN_QEMU_MONITOR_COMMAND_HMP     = (1 << 0)
)

type VirDomainProcessSignal int

const (
	VIR_DOMAIN_PROCESS_SIGNAL_NOP  = VirDomainProcessSignal(C.VIR_DOMAIN_PROCESS_SIGNAL_NOP)
	VIR_DOMAIN_PROCESS_SIGNAL_HUP  = VirDomainProcessSignal(C.VIR_DOMAIN_PROCESS_SIGNAL_HUP)
	VIR_DOMAIN_PROCESS_SIGNAL_INT  = VirDomainProcessSignal(C.VIR_DOMAIN_PROCESS_SIGNAL_INT)
	VIR_DOMAIN_PROCESS_SIGNAL_QUIT = VirDomainProcessSignal(C.VIR_DOMAIN_PROCESS_SIGNAL_QUIT)
	VIR_DOMAIN_PROCESS_SIGNAL_ILL  = VirDomainProcessSignal(C.VIR_DOMAIN_PROCESS_SIGNAL_ILL)
	VIR_DOMAIN_PROCESS_SIGNAL_TRAP = VirDomainProcessSignal(C.VIR_DOMAIN_PROCESS_SIGNAL_TRAP)
	VIR_DOMAIN_PROCESS_SIGNAL_ABRT = VirDomainProcessSignal(C.VIR_DOMAIN_PROCESS_SIGNAL_ABRT)
	VIR_DOMAIN_PROCESS_SIGNAL_BUS  = VirDomainProcessSignal(C.VIR_DOMAIN_PROCESS_SIGNAL_BUS)
	VIR_DOMAIN_PROCESS_SIGNAL_FPE  = VirDomainProcessSignal(C.VIR_DOMAIN_PROCESS_SIGNAL_FPE)
	VIR_DOMAIN_PROCESS_SIGNAL_KILL = VirDomainProcessSignal(C.VIR_DOMAIN_PROCESS_SIGNAL_KILL)

	VIR_DOMAIN_PROCESS_SIGNAL_USR1   = VirDomainProcessSignal(C.VIR_DOMAIN_PROCESS_SIGNAL_USR1)
	VIR_DOMAIN_PROCESS_SIGNAL_SEGV   = VirDomainProcessSignal(C.VIR_DOMAIN_PROCESS_SIGNAL_SEGV)
	VIR_DOMAIN_PROCESS_SIGNAL_USR2   = VirDomainProcessSignal(C.VIR_DOMAIN_PROCESS_SIGNAL_USR2)
	VIR_DOMAIN_PROCESS_SIGNAL_PIPE   = VirDomainProcessSignal(C.VIR_DOMAIN_PROCESS_SIGNAL_PIPE)
	VIR_DOMAIN_PROCESS_SIGNAL_ALRM   = VirDomainProcessSignal(C.VIR_DOMAIN_PROCESS_SIGNAL_ALRM)
	VIR_DOMAIN_PROCESS_SIGNAL_TERM   = VirDomainProcessSignal(C.VIR_DOMAIN_PROCESS_SIGNAL_TERM)
	VIR_DOMAIN_PROCESS_SIGNAL_STKFLT = VirDomainProcessSignal(C.VIR_DOMAIN_PROCESS_SIGNAL_STKFLT)
	VIR_DOMAIN_PROCESS_SIGNAL_CHLD   = VirDomainProcessSignal(C.VIR_DOMAIN_PROCESS_SIGNAL_CHLD)
	VIR_DOMAIN_PROCESS_SIGNAL_CONT   = VirDomainProcessSignal(C.VIR_DOMAIN_PROCESS_SIGNAL_CONT)
	VIR_DOMAIN_PROCESS_SIGNAL_STOP   = VirDomainProcessSignal(C.VIR_DOMAIN_PROCESS_SIGNAL_STOP)

	VIR_DOMAIN_PROCESS_SIGNAL_TSTP   = VirDomainProcessSignal(C.VIR_DOMAIN_PROCESS_SIGNAL_TSTP)
	VIR_DOMAIN_PROCESS_SIGNAL_TTIN   = VirDomainProcessSignal(C.VIR_DOMAIN_PROCESS_SIGNAL_TTIN)
	VIR_DOMAIN_PROCESS_SIGNAL_TTOU   = VirDomainProcessSignal(C.VIR_DOMAIN_PROCESS_SIGNAL_TTOU)
	VIR_DOMAIN_PROCESS_SIGNAL_URG    = VirDomainProcessSignal(C.VIR_DOMAIN_PROCESS_SIGNAL_URG)
	VIR_DOMAIN_PROCESS_SIGNAL_XCPU   = VirDomainProcessSignal(C.VIR_DOMAIN_PROCESS_SIGNAL_XCPU)
	VIR_DOMAIN_PROCESS_SIGNAL_XFSZ   = VirDomainProcessSignal(C.VIR_DOMAIN_PROCESS_SIGNAL_XFSZ)
	VIR_DOMAIN_PROCESS_SIGNAL_VTALRM = VirDomainProcessSignal(C.VIR_DOMAIN_PROCESS_SIGNAL_VTALRM)
	VIR_DOMAIN_PROCESS_SIGNAL_PROF   = VirDomainProcessSignal(C.VIR_DOMAIN_PROCESS_SIGNAL_PROF)
	VIR_DOMAIN_PROCESS_SIGNAL_WINCH  = VirDomainProcessSignal(C.VIR_DOMAIN_PROCESS_SIGNAL_WINCH)
	VIR_DOMAIN_PROCESS_SIGNAL_POLL   = VirDomainProcessSignal(C.VIR_DOMAIN_PROCESS_SIGNAL_POLL)

	VIR_DOMAIN_PROCESS_SIGNAL_PWR = VirDomainProcessSignal(C.VIR_DOMAIN_PROCESS_SIGNAL_PWR)
	VIR_DOMAIN_PROCESS_SIGNAL_SYS = VirDomainProcessSignal(C.VIR_DOMAIN_PROCESS_SIGNAL_SYS)
	VIR_DOMAIN_PROCESS_SIGNAL_RT0 = VirDomainProcessSignal(C.VIR_DOMAIN_PROCESS_SIGNAL_RT0)
	VIR_DOMAIN_PROCESS_SIGNAL_RT1 = VirDomainProcessSignal(C.VIR_DOMAIN_PROCESS_SIGNAL_RT1)
	VIR_DOMAIN_PROCESS_SIGNAL_RT2 = VirDomainProcessSignal(C.VIR_DOMAIN_PROCESS_SIGNAL_RT2)
	VIR_DOMAIN_PROCESS_SIGNAL_RT3 = VirDomainProcessSignal(C.VIR_DOMAIN_PROCESS_SIGNAL_RT3)
	VIR_DOMAIN_PROCESS_SIGNAL_RT4 = VirDomainProcessSignal(C.VIR_DOMAIN_PROCESS_SIGNAL_RT4)
	VIR_DOMAIN_PROCESS_SIGNAL_RT5 = VirDomainProcessSignal(C.VIR_DOMAIN_PROCESS_SIGNAL_RT5)
	VIR_DOMAIN_PROCESS_SIGNAL_RT6 = VirDomainProcessSignal(C.VIR_DOMAIN_PROCESS_SIGNAL_RT6)
	VIR_DOMAIN_PROCESS_SIGNAL_RT7 = VirDomainProcessSignal(C.VIR_DOMAIN_PROCESS_SIGNAL_RT7)

	VIR_DOMAIN_PROCESS_SIGNAL_RT8  = VirDomainProcessSignal(C.VIR_DOMAIN_PROCESS_SIGNAL_RT8)
	VIR_DOMAIN_PROCESS_SIGNAL_RT9  = VirDomainProcessSignal(C.VIR_DOMAIN_PROCESS_SIGNAL_RT9)
	VIR_DOMAIN_PROCESS_SIGNAL_RT10 = VirDomainProcessSignal(C.VIR_DOMAIN_PROCESS_SIGNAL_RT10)
	VIR_DOMAIN_PROCESS_SIGNAL_RT11 = VirDomainProcessSignal(C.VIR_DOMAIN_PROCESS_SIGNAL_RT11)
	VIR_DOMAIN_PROCESS_SIGNAL_RT12 = VirDomainProcessSignal(C.VIR_DOMAIN_PROCESS_SIGNAL_RT12)
	VIR_DOMAIN_PROCESS_SIGNAL_RT13 = VirDomainProcessSignal(C.VIR_DOMAIN_PROCESS_SIGNAL_RT13)
	VIR_DOMAIN_PROCESS_SIGNAL_RT14 = VirDomainProcessSignal(C.VIR_DOMAIN_PROCESS_SIGNAL_RT14)
	VIR_DOMAIN_PROCESS_SIGNAL_RT15 = VirDomainProcessSignal(C.VIR_DOMAIN_PROCESS_SIGNAL_RT15)
	VIR_DOMAIN_PROCESS_SIGNAL_RT16 = VirDomainProcessSignal(C.VIR_DOMAIN_PROCESS_SIGNAL_RT16)
	VIR_DOMAIN_PROCESS_SIGNAL_RT17 = VirDomainProcessSignal(C.VIR_DOMAIN_PROCESS_SIGNAL_RT17)
	VIR_DOMAIN_PROCESS_SIGNAL_RT18 = VirDomainProcessSignal(C.VIR_DOMAIN_PROCESS_SIGNAL_RT18)

	VIR_DOMAIN_PROCESS_SIGNAL_RT19 = VirDomainProcessSignal(C.VIR_DOMAIN_PROCESS_SIGNAL_RT19)
	VIR_DOMAIN_PROCESS_SIGNAL_RT20 = VirDomainProcessSignal(C.VIR_DOMAIN_PROCESS_SIGNAL_RT20)
	VIR_DOMAIN_PROCESS_SIGNAL_RT21 = VirDomainProcessSignal(C.VIR_DOMAIN_PROCESS_SIGNAL_RT21)
	VIR_DOMAIN_PROCESS_SIGNAL_RT22 = VirDomainProcessSignal(C.VIR_DOMAIN_PROCESS_SIGNAL_RT22)
	VIR_DOMAIN_PROCESS_SIGNAL_RT23 = VirDomainProcessSignal(C.VIR_DOMAIN_PROCESS_SIGNAL_RT23)
	VIR_DOMAIN_PROCESS_SIGNAL_RT24 = VirDomainProcessSignal(C.VIR_DOMAIN_PROCESS_SIGNAL_RT24)
	VIR_DOMAIN_PROCESS_SIGNAL_RT25 = VirDomainProcessSignal(C.VIR_DOMAIN_PROCESS_SIGNAL_RT25)
	VIR_DOMAIN_PROCESS_SIGNAL_RT26 = VirDomainProcessSignal(C.VIR_DOMAIN_PROCESS_SIGNAL_RT26)
	VIR_DOMAIN_PROCESS_SIGNAL_RT27 = VirDomainProcessSignal(C.VIR_DOMAIN_PROCESS_SIGNAL_RT27)

	VIR_DOMAIN_PROCESS_SIGNAL_RT28 = VirDomainProcessSignal(C.VIR_DOMAIN_PROCESS_SIGNAL_RT28)
	VIR_DOMAIN_PROCESS_SIGNAL_RT29 = VirDomainProcessSignal(C.VIR_DOMAIN_PROCESS_SIGNAL_RT29)
	VIR_DOMAIN_PROCESS_SIGNAL_RT30 = VirDomainProcessSignal(C.VIR_DOMAIN_PROCESS_SIGNAL_RT30)
	VIR_DOMAIN_PROCESS_SIGNAL_RT31 = VirDomainProcessSignal(C.VIR_DOMAIN_PROCESS_SIGNAL_RT31)
	VIR_DOMAIN_PROCESS_SIGNAL_RT32 = VirDomainProcessSignal(C.VIR_DOMAIN_PROCESS_SIGNAL_RT32)
)

type VirDomain struct {
	ptr C.virDomainPtr
}

type VirDomainBlockInfo struct {
	ptr C.virDomainBlockInfo
}

type VirDomainInfo struct {
	ptr C.virDomainInfo
}

type VirTypedParameter struct {
	Name  string
	Value interface{}
}

type VirDomainMemoryStat struct {
	Tag int32
	Val uint64
}

type VirVcpuInfo struct {
	Number  uint32
	State   int32
	CpuTime uint64
	Cpu     int32
	CpuMap  []uint32
}

type VirTypedParameters []VirTypedParameter

func (dest *VirTypedParameters) loadFromCPtr(params C.virTypedParameterPtr, nParams int) {
	// reset slice
	*dest = VirTypedParameters{}

	// transform that C array to a go slice
	hdr := reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(params)),
		Len:  int(nParams),
		Cap:  int(nParams),
	}
	rawParams := *(*[]C.struct__virTypedParameter)(unsafe.Pointer(&hdr))

	// there is probably a more elegant way to deal with that union
	for _, rawParam := range rawParams {
		name := C.GoStringN(&rawParam.field[0], C.VIR_TYPED_PARAM_FIELD_LENGTH)
		if nbIdx := strings.Index(name, "\x00"); nbIdx != -1 {
			name = name[:nbIdx]
		}
		switch rawParam._type {
		case C.VIR_TYPED_PARAM_INT:
			*dest = append(*dest, VirTypedParameter{name, int(*(*C.int)(unsafe.Pointer(&rawParam.value[0])))})
		case C.VIR_TYPED_PARAM_UINT:
			*dest = append(*dest, VirTypedParameter{name, uint32(*(*C.uint)(unsafe.Pointer(&rawParam.value[0])))})
		case C.VIR_TYPED_PARAM_LLONG:
			*dest = append(*dest, VirTypedParameter{name, int64(*(*C.longlong)(unsafe.Pointer(&rawParam.value[0])))})
		case C.VIR_TYPED_PARAM_ULLONG:
			*dest = append(*dest, VirTypedParameter{name, uint64(*(*C.ulonglong)(unsafe.Pointer(&rawParam.value[0])))})
		case C.VIR_TYPED_PARAM_DOUBLE:
			*dest = append(*dest, VirTypedParameter{name, float64(*(*C.double)(unsafe.Pointer(&rawParam.value[0])))})
		case C.VIR_TYPED_PARAM_BOOLEAN:
			if int(*(*C.char)(unsafe.Pointer(&rawParam.value[0]))) == 1 {
				*dest = append(*dest, VirTypedParameter{name, true})
			} else {
				*dest = append(*dest, VirTypedParameter{name, false})
			}
		case C.VIR_TYPED_PARAM_STRING:
			*dest = append(*dest, VirTypedParameter{name, C.GoString((*C.char)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(&rawParam.value[0])))))})
		}
	}
}

func (d *VirDomain) Free() error {
	if result := C.virDomainFree(d.ptr); result != 0 {
		return GetLastError()
	}
	return nil
}

func (d *VirDomain) Create() error {
	result := C.virDomainCreate(d.ptr)
	if result == -1 {
		return GetLastError()
	}
	return nil
}

func (d *VirDomain) CreateWithFlags(flags uint) error {
	result := C.virDomainCreateWithFlags(d.ptr, C.uint(flags))
	if result == -1 {
		return GetLastError()
	}
	return nil
}

func (d *VirDomain) Destroy() error {
	result := C.virDomainDestroy(d.ptr)
	if result == -1 {
		return GetLastError()
	}
	return nil
}

func (d *VirDomain) Shutdown() error {
	result := C.virDomainShutdown(d.ptr)
	if result == -1 {
		return GetLastError()
	}
	return nil
}

func (d *VirDomain) Reboot(flags uint) error {
	result := C.virDomainReboot(d.ptr, C.uint(flags))
	if result == -1 {
		return GetLastError()
	}
	return nil
}

func (d *VirDomain) IsActive() (bool, error) {
	result := C.virDomainIsActive(d.ptr)
	if result == -1 {
		return false, GetLastError()
	}
	if result == 1 {
		return true, nil
	}
	return false, nil
}

func (d *VirDomain) IsPersistent() (bool, error) {
	result := C.virDomainIsPersistent(d.ptr)
	if result == -1 {
		return false, GetLastError()
	}
	if result == 1 {
		return true, nil
	}
	return false, nil
}

func (d *VirDomain) SetAutostart(autostart bool) error {
	var cAutostart C.int
	switch autostart {
	case true:
		cAutostart = 1
	default:
		cAutostart = 0
	}
	result := C.virDomainSetAutostart(d.ptr, cAutostart)
	if result == -1 {
		return GetLastError()
	}
	return nil
}

func (d *VirDomain) GetAutostart() (bool, error) {
	var out C.int
	result := C.virDomainGetAutostart(d.ptr, (*C.int)(unsafe.Pointer(&out)))
	if result == -1 {
		return false, GetLastError()
	}
	switch out {
	case 1:
		return true, nil
	default:
		return false, nil
	}
}

func (d *VirDomain) GetBlockInfo(disk string, flag uint) (VirDomainBlockInfo, error) {
	bi := VirDomainBlockInfo{}
	var ptr C.virDomainBlockInfo
	cDisk := C.CString(disk)
	defer C.free(unsafe.Pointer(cDisk))
	result := C.virDomainGetBlockInfo(d.ptr, cDisk, (*C.virDomainBlockInfo)(unsafe.Pointer(&ptr)), C.uint(flag))
	if result == -1 {
		return bi, GetLastError()
	}
	bi.ptr = ptr
	return bi, nil
}

func (b *VirDomainBlockInfo) Allocation() uint64 {
	return uint64(b.ptr.allocation)
}

func (b *VirDomainBlockInfo) Capacity() uint64 {
	return uint64(b.ptr.capacity)
}

func (b *VirDomainBlockInfo) Physical() uint64 {
	return uint64(b.ptr.physical)
}

func (d *VirDomain) GetName() (string, error) {
	name := C.virDomainGetName(d.ptr)
	if name == nil {
		return "", GetLastError()
	}
	return C.GoString(name), nil
}

func (d *VirDomain) GetState() ([]int, error) {
	var cState C.int
	var cReason C.int
	result := C.virDomainGetState(d.ptr,
		(*C.int)(unsafe.Pointer(&cState)),
		(*C.int)(unsafe.Pointer(&cReason)),
		0)
	if int(result) == -1 {
		return []int{}, GetLastError()
	}
	return []int{int(cState), int(cReason)}, nil
}

func (d *VirDomain) GetID() (uint, error) {
	id := uint(C.virDomainGetID(d.ptr))
	if id == ^uint(0) {
		return id, GetLastError()
	}
	return id, nil
}

func (d *VirDomain) GetUUID() ([]byte, error) {
	var cUuid [C.VIR_UUID_BUFLEN](byte)
	cuidPtr := unsafe.Pointer(&cUuid)
	result := C.virDomainGetUUID(d.ptr, (*C.uchar)(cuidPtr))
	if result != 0 {
		return []byte{}, GetLastError()
	}
	return C.GoBytes(cuidPtr, C.VIR_UUID_BUFLEN), nil
}

func (d *VirDomain) GetUUIDString() (string, error) {
	var cUuid [C.VIR_UUID_STRING_BUFLEN](C.char)
	cuidPtr := unsafe.Pointer(&cUuid)
	result := C.virDomainGetUUIDString(d.ptr, (*C.char)(cuidPtr))
	if result != 0 {
		return "", GetLastError()
	}
	return C.GoString((*C.char)(cuidPtr)), nil
}

func (d *VirDomain) GetInfo() (VirDomainInfo, error) {
	di := VirDomainInfo{}
	var ptr C.virDomainInfo
	result := C.virDomainGetInfo(d.ptr, (*C.virDomainInfo)(unsafe.Pointer(&ptr)))
	if result == -1 {
		return di, GetLastError()
	}
	di.ptr = ptr
	return di, nil
}

func (d *VirDomain) GetXMLDesc(flags uint32) (string, error) {
	result := C.virDomainGetXMLDesc(d.ptr, C.uint(flags))
	if result == nil {
		return "", GetLastError()
	}
	xml := C.GoString(result)
	C.free(unsafe.Pointer(result))
	return xml, nil
}

func (i *VirDomainInfo) GetState() uint8 {
	return uint8(i.ptr.state)
}

func (i *VirDomainInfo) GetMaxMem() uint64 {
	return uint64(i.ptr.maxMem)
}

func (i *VirDomainInfo) GetMemory() uint64 {
	return uint64(i.ptr.memory)
}

func (i *VirDomainInfo) GetNrVirtCpu() uint16 {
	return uint16(i.ptr.nrVirtCpu)
}

func (i *VirDomainInfo) GetCpuTime() uint64 {
	return uint64(i.ptr.cpuTime)
}

func (d *VirDomain) GetCPUStats(params *VirTypedParameters, nParams int, startCpu int, nCpus uint32, flags uint32) (int, error) {
	var cParams C.virTypedParameterPtr
	var cParamsLen int

	cParamsLen = int(nCpus) * nParams

	if params != nil && cParamsLen > 0 {
		cParams = (C.virTypedParameterPtr)(C.calloc(C.size_t(cParamsLen), C.size_t(unsafe.Sizeof(C.struct__virTypedParameter{}))))
		defer C.virTypedParamsFree(cParams, C.int(cParamsLen))
	} else {
		cParamsLen = 0
		cParams = nil
	}

	result := int(C.virDomainGetCPUStats(d.ptr, (C.virTypedParameterPtr)(cParams), C.uint(nParams), C.int(startCpu), C.uint(nCpus), C.uint(flags)))
	if result == -1 {
		return result, GetLastError()
	}

	if cParamsLen > 0 {
		params.loadFromCPtr(cParams, cParamsLen)
	}

	return result, nil
}

// Warning: No test written for this function
func (d *VirDomain) GetInterfaceParameters(device string, params *VirTypedParameters, nParams *int, flags uint32) (int, error) {
	var cParams C.virTypedParameterPtr

	if params != nil && *nParams > 0 {
		cParams = (C.virTypedParameterPtr)(C.calloc(C.size_t(*nParams), C.size_t(unsafe.Sizeof(C.struct__virTypedParameter{}))))
		defer C.virTypedParamsFree(cParams, C.int(*nParams))
	} else {
		cParams = nil
	}

	cDevice := C.CString(device)
	defer C.free(unsafe.Pointer(cDevice))
	result := int(C.virDomainGetInterfaceParameters(d.ptr, cDevice,
		(C.virTypedParameterPtr)(cParams), (*C.int)(unsafe.Pointer(nParams)), C.uint(flags)))
	if result == -1 {
		return result, GetLastError()
	}

	if params != nil && *nParams > 0 {
		params.loadFromCPtr(cParams, *nParams)
	}

	return result, nil
}

func (d *VirDomain) GetMetadata(tipus int, uri string, flags uint32) (string, error) {
	var cUri *C.char
	if uri != "" {
		cUri = C.CString(uri)
		defer C.free(unsafe.Pointer(cUri))
	}

	result := C.virDomainGetMetadata(d.ptr, C.int(tipus), cUri, C.uint(flags))
	if result == nil {
		return "", GetLastError()

	}
	defer C.free(unsafe.Pointer(result))
	return C.GoString(result), nil
}

func (d *VirDomain) SetMetadata(metaDataType VirDomainMetadataType, metaDataCont, uriKey, uri string, flags uint32) error {
	var cMetaDataCont *C.char
	var cUriKey *C.char
	var cUri *C.char

	cMetaDataCont = C.CString(metaDataCont)
	defer C.free(unsafe.Pointer(cMetaDataCont))

	if metaDataType == VIR_DOMAIN_METADATA_ELEMENT {
		cUriKey = C.CString(uriKey)
		defer C.free(unsafe.Pointer(cUriKey))
		cUri = C.CString(uri)
		defer C.free(unsafe.Pointer(cUri))
	}
	result := C.virDomainSetMetadata(d.ptr, C.int(metaDataType), cMetaDataCont, cUriKey, cUri, C.uint(flags))
	if result == -1 {
		return GetLastError()
	}
	return nil
}

func (d *VirDomain) Undefine() error {
	result := C.virDomainUndefine(d.ptr)
	if result == -1 {
		return GetLastError()
	}
	return nil
}

func (d *VirDomain) UndefineFlags(flags uint) error {
	result := C.virDomainUndefineFlags(d.ptr, C.uint(flags))
	if result == -1 {
		return GetLastError()
	}
	return nil
}

func (d *VirDomain) SetMaxMemory(memory uint) error {
	result := C.virDomainSetMaxMemory(d.ptr, C.ulong(memory))
	if result == -1 {
		return GetLastError()
	}
	return nil
}

func (d *VirDomain) SetMemory(memory uint64) error {
	result := C.virDomainSetMemory(d.ptr, C.ulong(memory))
	if result == -1 {
		return GetLastError()
	}
	return nil
}

func (d *VirDomain) SetMemoryFlags(memory uint64, flags uint32) error {
	result := C.virDomainSetMemoryFlags(d.ptr, C.ulong(memory), C.uint(flags))
	if result == -1 {
		return GetLastError()
	}
	return nil
}

func (d *VirDomain) SetMemoryStatsPeriod(period int, flags uint) error {
	result := C.virDomainSetMemoryStatsPeriod(d.ptr, C.int(period), C.uint(flags))
	if result == -1 {
		return GetLastError()
	}
	return nil
}

func (d *VirDomain) SetVcpus(vcpu uint16) error {
	result := C.virDomainSetVcpus(d.ptr, C.uint(vcpu))
	if result == -1 {
		return GetLastError()
	}
	return nil
}

func (d *VirDomain) SetVcpusFlags(vcpu uint16, flags VirDomainVcpuFlags) error {
	result := C.virDomainSetVcpusFlags(d.ptr, C.uint(vcpu), C.uint(flags))
	if result == -1 {
		return GetLastError()
	}
	return nil
}

func (d *VirDomain) Suspend() error {
	result := C.virDomainSuspend(d.ptr)
	if result == -1 {
		return GetLastError()
	}
	return nil
}

func (d *VirDomain) Resume() error {
	result := C.virDomainResume(d.ptr)
	if result == -1 {
		return GetLastError()
	}
	return nil
}

func (d *VirDomain) AbortJob() error {
	result := C.virDomainAbortJob(d.ptr)
	if result == -1 {
		return GetLastError()
	}
	return nil
}

func (d *VirDomain) DestroyFlags(flags VirDomainDestroyFlags) error {
	result := C.virDomainDestroyFlags(d.ptr, C.uint(flags))
	if result == -1 {
		return GetLastError()
	}
	return nil
}

func (d *VirDomain) ShutdownFlags(flags VirDomainShutdownFlags) error {
	result := C.virDomainShutdownFlags(d.ptr, C.uint(flags))
	if result == -1 {
		return GetLastError()
	}
	return nil
}

func (d *VirDomain) AttachDevice(xml string) error {
	cXml := C.CString(xml)
	defer C.free(unsafe.Pointer(cXml))
	result := C.virDomainAttachDevice(d.ptr, cXml)
	if result == -1 {
		return GetLastError()
	}
	return nil
}

func (d *VirDomain) AttachDeviceFlags(xml string, flags uint) error {
	cXml := C.CString(xml)
	defer C.free(unsafe.Pointer(cXml))
	result := C.virDomainAttachDeviceFlags(d.ptr, cXml, C.uint(flags))
	if result == -1 {
		return GetLastError()
	}
	return nil
}

func (d *VirDomain) DetachDevice(xml string) error {
	cXml := C.CString(xml)
	defer C.free(unsafe.Pointer(cXml))
	result := C.virDomainDetachDevice(d.ptr, cXml)
	if result == -1 {
		return GetLastError()
	}
	return nil
}

func (d *VirDomain) DetachDeviceFlags(xml string, flags uint) error {
	cXml := C.CString(xml)
	defer C.free(unsafe.Pointer(cXml))
	result := C.virDomainDetachDeviceFlags(d.ptr, cXml, C.uint(flags))
	if result == -1 {
		return GetLastError()
	}
	return nil
}

func (d *VirDomain) UpdateDeviceFlags(xml string, flags uint) error {
	cXml := C.CString(xml)
	defer C.free(unsafe.Pointer(cXml))
	result := C.virDomainUpdateDeviceFlags(d.ptr, cXml, C.uint(flags))
	if result == -1 {
		return GetLastError()
	}
	return nil
}

func (d *VirDomain) Screenshot(stream *VirStream, screen, flags uint) (string, error) {
	cType := C.virDomainScreenshot(d.ptr, stream.ptr, C.uint(screen), C.uint(flags))
	if cType == nil {
		return "", GetLastError()
	}
	defer C.free(unsafe.Pointer(cType))

	mimeType := C.GoString(cType)
	return mimeType, nil
}

func (d *VirDomain) SendKey(codeset, holdtime uint, keycodes []uint, flags uint) error {
	result := C.virDomainSendKey(d.ptr, C.uint(codeset), C.uint(holdtime), (*C.uint)(unsafe.Pointer(&keycodes[0])), C.int(len(keycodes)), C.uint(flags))
	if result == -1 {
		return GetLastError()
	}

	return nil
}

func (d *VirDomain) BlockStatsFlags(disk string, params *VirTypedParameters, nParams int, flags uint32) (int, error) {
	var cParams C.virTypedParameterPtr
	cDisk := C.CString(disk)
	defer C.free(unsafe.Pointer(cDisk))

	cParamsLen := C.int(nParams)

	if params != nil && nParams > 0 {
		cParams = (C.virTypedParameterPtr)(C.calloc(C.size_t(nParams), C.size_t(unsafe.Sizeof(C.struct__virTypedParameter{}))))
		defer C.virTypedParamsFree(cParams, cParamsLen)
	} else {
		cParams = nil
	}

	result := int(C.virDomainBlockStatsFlags(d.ptr, cDisk, (C.virTypedParameterPtr)(cParams), &cParamsLen, C.uint(flags)))
	if result == -1 {
		return result, GetLastError()
	}

	if cParamsLen > 0 && params != nil {
		params.loadFromCPtr(cParams, nParams)
	}

	return int(cParamsLen), nil
}

type VirDomainBlockStats struct {
	RdReq   int64
	WrReq   int64
	RdBytes int64
	WrBytes int64
}

type VirDomainInterfaceStats struct {
	RxBytes   int64
	RxPackets int64
	RxErrs    int64
	RxDrop    int64
	TxBytes   int64
	TxPackets int64
	TxErrs    int64
	TxDrop    int64
}

func (d *VirDomain) BlockStats(path string) (VirDomainBlockStats, error) {
	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))

	size := C.size_t(unsafe.Sizeof(C.struct__virDomainBlockStats{}))

	cStats := (C.virDomainBlockStatsPtr)(C.malloc(size))
	defer C.free(unsafe.Pointer(cStats))

	result := C.virDomainBlockStats(d.ptr, cPath, (C.virDomainBlockStatsPtr)(cStats), size)

	if result != 0 {
		return VirDomainBlockStats{}, GetLastError()
	}
	return VirDomainBlockStats{
		WrReq:   int64(cStats.wr_req),
		RdReq:   int64(cStats.rd_req),
		RdBytes: int64(cStats.rd_bytes),
		WrBytes: int64(cStats.wr_bytes),
	}, nil
}

func (d *VirDomain) InterfaceStats(path string) (VirDomainInterfaceStats, error) {
	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))

	size := C.size_t(unsafe.Sizeof(C.struct__virDomainInterfaceStats{}))

	cStats := (C.virDomainInterfaceStatsPtr)(C.malloc(size))
	defer C.free(unsafe.Pointer(cStats))

	result := C.virDomainInterfaceStats(d.ptr, cPath, (C.virDomainInterfaceStatsPtr)(cStats), size)

	if result != 0 {
		return VirDomainInterfaceStats{}, GetLastError()
	}
	return VirDomainInterfaceStats{
		RxBytes:   int64(cStats.rx_bytes),
		RxPackets: int64(cStats.rx_packets),
		RxErrs:    int64(cStats.rx_errs),
		RxDrop:    int64(cStats.rx_drop),
		TxBytes:   int64(cStats.tx_bytes),
		TxPackets: int64(cStats.tx_packets),
		TxErrs:    int64(cStats.tx_errs),
		TxDrop:    int64(cStats.tx_drop),
	}, nil
}

func (d *VirDomain) MemoryStats(nrStats uint32, flags uint32) ([]VirDomainMemoryStat, error) {
	ptr := make([]C.virDomainMemoryStatStruct, nrStats)

	result := C.virDomainMemoryStats(
		d.ptr, (C.virDomainMemoryStatPtr)(unsafe.Pointer(&ptr[0])),
		C.uint(nrStats), C.uint(flags))

	if result == -1 {
		return []VirDomainMemoryStat{}, GetLastError()
	}

	out := make([]VirDomainMemoryStat, result)
	for i := 0; i < int(result); i++ {
		out = append(out, VirDomainMemoryStat{
			Tag: int32(ptr[i].tag),
			Val: uint64(ptr[i].val),
		})
	}
	return out, nil
}

func (d *VirDomain) GetVcpus(maxInfo int32) ([]VirVcpuInfo, error) {
	ptr := make([]C.virVcpuInfo, maxInfo)

	result := C.virDomainGetVcpus(
		d.ptr, (C.virVcpuInfoPtr)(unsafe.Pointer(&ptr[0])),
		C.int(maxInfo), nil, C.int(0))

	if result == -1 {
		return []VirVcpuInfo{}, GetLastError()
	}

	out := make([]VirVcpuInfo, 0)
	for i := 0; i < int(result); i++ {
		out = append(out, VirVcpuInfo{
			Number:  uint32(ptr[i].number),
			State:   int32(ptr[i].state),
			CpuTime: uint64(ptr[i].cpuTime),
			Cpu:     int32(ptr[i].cpu),
		})
	}

	return out, nil
}

// libvirt-domain.h: VIR_CPU_MAPLEN
func virCpuMapLen(cpu uint32) C.int {
	return C.int((cpu + 7) / 8)
}

// extractCpuMask extracts an individual cpumask from a slice of cpumasks
// and parses it into a slice of CPU ids
func extractCpuMask(bytesCpuMaps []byte, n, mapLen int) []uint32 {
	const byteSize = uint(8)

	// Repslice the big array to separate only mask number 'n'
	cpuMap := bytesCpuMaps[n*mapLen : (n+1)*mapLen]

	out := make([]uint32, 0)
	for i, b := range cpuMap { // iterate over bytes of the mask
		for j := uint(0); j < byteSize; j++ { // iterate over bits in this byte
			if (b>>j)&0x1 == 1 {
				out = append(out, uint32(j+uint(i)*byteSize))
			}
		}
	}

	return out
}

func (d *VirDomain) GetVcpusCpuMap(maxInfo int, maxCPUs uint32) ([]VirVcpuInfo, error) {
	ptr := make([]C.virVcpuInfo, maxInfo)

	mapLen := virCpuMapLen(maxCPUs)                    // Length of CPUs bitmask in bytes
	bufSize := int(mapLen) * int(maxInfo)              // Length of the array of 'maxinfo' bitmasks
	cpuMaps := (*C.uchar)(C.malloc(C.size_t(bufSize))) // Array itself
	defer C.free(unsafe.Pointer(cpuMaps))

	result := C.virDomainGetVcpus(
		d.ptr, (C.virVcpuInfoPtr)(unsafe.Pointer(&ptr[0])),
		C.int(maxInfo), cpuMaps, mapLen)

	if result == -1 {
		return nil, GetLastError()
	}

	// Convert to golang []byte for easier handling
	bytesCpuMaps := C.GoBytes(unsafe.Pointer(cpuMaps), C.int(bufSize))

	out := make([]VirVcpuInfo, 0)
	for i := 0; i < int(result); i++ {
		out = append(out, VirVcpuInfo{
			Number:  uint32(ptr[i].number),
			State:   int32(ptr[i].state),
			CpuTime: uint64(ptr[i].cpuTime),
			Cpu:     int32(ptr[i].cpu),
			CpuMap:  extractCpuMask(bytesCpuMaps, i, int(mapLen)),
		})
	}

	return out, nil
}

func (d *VirDomain) GetVcpusFlags(flags uint32) (int32, error) {
	result := C.virDomainGetVcpusFlags(d.ptr, C.uint(flags))
	if result == -1 {
		return 0, GetLastError()
	}
	return int32(result), nil
}

func (d *VirDomain) QemuMonitorCommand(flags uint32, command string) (string, error) {
	var cResult *C.char
	cCommand := C.CString(command)
	defer C.free(unsafe.Pointer(cCommand))
	result := C.virDomainQemuMonitorCommand(d.ptr, cCommand, &cResult, C.uint(flags))

	if result != 0 {
		return "", GetLastError()
	}

	rstring := C.GoString(cResult)
	C.free(unsafe.Pointer(cResult))
	return rstring, nil
}

func cpuMask(cpuMap []uint32, maxCPUs uint32) (*C.uchar, C.int) {
	const byteSize = uint(8)

	mapLen := virCpuMapLen(maxCPUs) // Length of CPUs bitmask in bytes
	bytesCpuMap := make([]byte, mapLen)

	for _, c := range cpuMap {
		by := uint(c) / byteSize
		bi := uint(c) % byteSize
		bytesCpuMap[by] |= 1 << bi
	}

	return (*C.uchar)(&bytesCpuMap[0]), mapLen
}

func (d *VirDomain) PinVcpu(vcpu uint, cpuMap []uint32, maxCPUs uint32) error {

	cpumap, maplen := cpuMask(cpuMap, maxCPUs)

	result := C.virDomainPinVcpu(d.ptr, C.uint(vcpu), cpumap, maplen)

	if result == -1 {
		return GetLastError()
	}

	return nil
}

func (d *VirDomain) PinVcpuFlags(vcpu uint, cpuMap []uint32, flags uint, maxCPUs uint32) error {
	cpumap, maplen := cpuMask(cpuMap, maxCPUs)

	result := C.virDomainPinVcpuFlags(d.ptr, C.uint(vcpu), cpumap, maplen, C.uint(flags))

	if result == -1 {
		return GetLastError()
	}

	return nil
}

type VirDomainIPAddress struct {
	Type   int
	Addr   string
	Prefix uint
}

type VirDomainInterface struct {
	Name   string
	Hwaddr string
	Addrs  []VirDomainIPAddress
}

func (d *VirDomain) ListAllInterfaceAddresses(src uint) ([]VirDomainInterface, error) {
	var cList *C.virDomainInterfacePtr
	numIfaces := int(C.virDomainInterfaceAddresses(d.ptr, (**C.virDomainInterfacePtr)(&cList), C.uint(src), 0))
	if numIfaces == -1 {
		return nil, GetLastError()
	}

	hdr := reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(cList)),
		Len:  int(numIfaces),
		Cap:  int(numIfaces),
	}

	ifaces := make([]VirDomainInterface, numIfaces)
	ifaceSlice := *(*[]C.virDomainInterfacePtr)(unsafe.Pointer(&hdr))

	for i := 0; i < numIfaces; i++ {
		ifaces[i].Name = C.GoString(ifaceSlice[i].name)
		ifaces[i].Hwaddr = C.GoString(ifaceSlice[i].hwaddr)

		numAddr := int(ifaceSlice[i].naddrs)
		addrHdr := reflect.SliceHeader{
			Data: uintptr(unsafe.Pointer(&ifaceSlice[i].addrs)),
			Len:  int(numAddr),
			Cap:  int(numAddr),
		}

		ifaces[i].Addrs = make([]VirDomainIPAddress, numAddr)
		addrSlice := *(*[]C.virDomainIPAddressPtr)(unsafe.Pointer(&addrHdr))

		for k := 0; k < numAddr; k++ {
			ifaces[i].Addrs[k] = VirDomainIPAddress{}
			ifaces[i].Addrs[k].Type = int(addrSlice[k]._type)
			ifaces[i].Addrs[k].Addr = C.GoString(addrSlice[k].addr)
			ifaces[i].Addrs[k].Prefix = uint(addrSlice[k].prefix)

		}
		C.virDomainInterfaceFree(ifaceSlice[i])
	}
	C.free(unsafe.Pointer(cList))
	return ifaces, nil
}
