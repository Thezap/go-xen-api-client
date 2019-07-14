//
// This file is generated. To change the content of this file, please do not
// apply the change to this file because it will get overwritten. Instead,
// change xenapi.go and execute 'go generate'.
//

package xenAPI

import (
	"errors"
	"fmt"
	"log"
	"github.com/amfranz/go-xmlrpc-client"
	"reflect"
	"strconv"
	"time"
)


var _ = errors.New
var _ = log.Println
var _ = fmt.Errorf
var _ = xmlrpc.NewClient
var _ = reflect.TypeOf
var _ = strconv.Atoi
var _ = time.UTC

type VMPowerState string

const (
  // VM is offline and not using any resources
	VMPowerStateHalted VMPowerState = "Halted"
  // All resources have been allocated but the VM itself is paused and its vCPUs are not running
	VMPowerStatePaused VMPowerState = "Paused"
  // Running
	VMPowerStateRunning VMPowerState = "Running"
  // VM state has been saved to disk and it is nolonger running. Note that disks remain in-use while the VM is suspended.
	VMPowerStateSuspended VMPowerState = "Suspended"
)

type OnNormalExit string

const (
  // destroy the VM state
	OnNormalExitDestroy OnNormalExit = "destroy"
  // restart the VM
	OnNormalExitRestart OnNormalExit = "restart"
)

type VMOperations string

const (
  // refers to the operation "snapshot"
	VMOperationsSnapshot VMOperations = "snapshot"
  // refers to the operation "clone"
	VMOperationsClone VMOperations = "clone"
  // refers to the operation "copy"
	VMOperationsCopy VMOperations = "copy"
  // refers to the operation "create_template"
	VMOperationsCreateTemplate VMOperations = "create_template"
  // refers to the operation "revert"
	VMOperationsRevert VMOperations = "revert"
  // refers to the operation "checkpoint"
	VMOperationsCheckpoint VMOperations = "checkpoint"
  // refers to the operation "snapshot_with_quiesce"
	VMOperationsSnapshotWithQuiesce VMOperations = "snapshot_with_quiesce"
  // refers to the operation "provision"
	VMOperationsProvision VMOperations = "provision"
  // refers to the operation "start"
	VMOperationsStart VMOperations = "start"
  // refers to the operation "start_on"
	VMOperationsStartOn VMOperations = "start_on"
  // refers to the operation "pause"
	VMOperationsPause VMOperations = "pause"
  // refers to the operation "unpause"
	VMOperationsUnpause VMOperations = "unpause"
  // refers to the operation "clean_shutdown"
	VMOperationsCleanShutdown VMOperations = "clean_shutdown"
  // refers to the operation "clean_reboot"
	VMOperationsCleanReboot VMOperations = "clean_reboot"
  // refers to the operation "hard_shutdown"
	VMOperationsHardShutdown VMOperations = "hard_shutdown"
  // refers to the operation "power_state_reset"
	VMOperationsPowerStateReset VMOperations = "power_state_reset"
  // refers to the operation "hard_reboot"
	VMOperationsHardReboot VMOperations = "hard_reboot"
  // refers to the operation "suspend"
	VMOperationsSuspend VMOperations = "suspend"
  // refers to the operation "csvm"
	VMOperationsCsvm VMOperations = "csvm"
  // refers to the operation "resume"
	VMOperationsResume VMOperations = "resume"
  // refers to the operation "resume_on"
	VMOperationsResumeOn VMOperations = "resume_on"
  // refers to the operation "pool_migrate"
	VMOperationsPoolMigrate VMOperations = "pool_migrate"
  // refers to the operation "migrate_send"
	VMOperationsMigrateSend VMOperations = "migrate_send"
  // refers to the operation "get_boot_record"
	VMOperationsGetBootRecord VMOperations = "get_boot_record"
  // refers to the operation "send_sysrq"
	VMOperationsSendSysrq VMOperations = "send_sysrq"
  // refers to the operation "send_trigger"
	VMOperationsSendTrigger VMOperations = "send_trigger"
  // refers to the operation "query_services"
	VMOperationsQueryServices VMOperations = "query_services"
  // refers to the operation "shutdown"
	VMOperationsShutdown VMOperations = "shutdown"
  // refers to the operation "call_plugin"
	VMOperationsCallPlugin VMOperations = "call_plugin"
  // Changing the memory settings
	VMOperationsChangingMemoryLive VMOperations = "changing_memory_live"
  // Waiting for the memory settings to change
	VMOperationsAwaitingMemoryLive VMOperations = "awaiting_memory_live"
  // Changing the memory dynamic range
	VMOperationsChangingDynamicRange VMOperations = "changing_dynamic_range"
  // Changing the memory static range
	VMOperationsChangingStaticRange VMOperations = "changing_static_range"
  // Changing the memory limits
	VMOperationsChangingMemoryLimits VMOperations = "changing_memory_limits"
  // Changing the shadow memory for a halted VM.
	VMOperationsChangingShadowMemory VMOperations = "changing_shadow_memory"
  // Changing the shadow memory for a running VM.
	VMOperationsChangingShadowMemoryLive VMOperations = "changing_shadow_memory_live"
  // Changing VCPU settings for a halted VM.
	VMOperationsChangingVCPUs VMOperations = "changing_VCPUs"
  // Changing VCPU settings for a running VM.
	VMOperationsChangingVCPUsLive VMOperations = "changing_VCPUs_live"
  // 
	VMOperationsAssertOperationValid VMOperations = "assert_operation_valid"
  // Add, remove, query or list data sources
	VMOperationsDataSourceOp VMOperations = "data_source_op"
  // 
	VMOperationsUpdateAllowedOperations VMOperations = "update_allowed_operations"
  // Turning this VM into a template
	VMOperationsMakeIntoTemplate VMOperations = "make_into_template"
  // importing a VM from a network stream
	VMOperationsImport VMOperations = "import"
  // exporting a VM to a network stream
	VMOperationsExport VMOperations = "export"
  // exporting VM metadata to a network stream
	VMOperationsMetadataExport VMOperations = "metadata_export"
  // Reverting the VM to a previous snapshotted state
	VMOperationsReverting VMOperations = "reverting"
  // refers to the act of uninstalling the VM
	VMOperationsDestroy VMOperations = "destroy"
)

type OnCrashBehaviour string

const (
  // destroy the VM state
	OnCrashBehaviourDestroy OnCrashBehaviour = "destroy"
  // record a coredump and then destroy the VM state
	OnCrashBehaviourCoredumpAndDestroy OnCrashBehaviour = "coredump_and_destroy"
  // restart the VM
	OnCrashBehaviourRestart OnCrashBehaviour = "restart"
  // record a coredump and then restart the VM
	OnCrashBehaviourCoredumpAndRestart OnCrashBehaviour = "coredump_and_restart"
  // leave the crashed VM paused
	OnCrashBehaviourPreserve OnCrashBehaviour = "preserve"
  // rename the crashed VM and start a new copy
	OnCrashBehaviourRenameRestart OnCrashBehaviour = "rename_restart"
)

type VMRecord struct {
  // Unique identifier/object reference
	UUID string
  // list of the operations allowed in this state. This list is advisory only and the server state may have changed by the time this field is read by a client.
	AllowedOperations []VMOperations
  // links each of the running tasks using this object (by reference) to a current_operation enum which describes the nature of the task.
	CurrentOperations map[string]VMOperations
  // Current power state of the machine
	PowerState VMPowerState
  // a human-readable name
	NameLabel string
  // a notes field containing human-readable description
	NameDescription string
  // Creators of VMs and templates may store version information here.
	UserVersion int
  // true if this is a template. Template VMs can never be started, they are used only for cloning other VMs
	IsATemplate bool
  // true if this is a default template. Default template VMs can never be started or migrated, they are used only for cloning other VMs
	IsDefaultTemplate bool
  // The VDI that a suspend image is stored on. (Only has meaning if VM is currently suspended)
	SuspendVDI VDIRef
  // the host the VM is currently resident on
	ResidentOn HostRef
  // A host which the VM has some affinity for (or NULL). This is used as a hint to the start call when it decides where to run the VM. Resource constraints may cause the VM to be started elsewhere.
	Affinity HostRef
  // Virtualization memory overhead (bytes).
	MemoryOverhead int
  // Dynamically-set memory target (bytes). The value of this field indicates the current target for memory available to this VM.
	MemoryTarget int
  // Statically-set (i.e. absolute) maximum (bytes). The value of this field at VM start time acts as a hard limit of the amount of memory a guest can use. New values only take effect on reboot.
	MemoryStaticMax int
  // Dynamic maximum (bytes)
	MemoryDynamicMax int
  // Dynamic minimum (bytes)
	MemoryDynamicMin int
  // Statically-set (i.e. absolute) mininum (bytes). The value of this field indicates the least amount of memory this VM can boot with without crashing.
	MemoryStaticMin int
  // configuration parameters for the selected VCPU policy
	VCPUsParams map[string]string
  // Max number of VCPUs
	VCPUsMax int
  // Boot number of VCPUs
	VCPUsAtStartup int
  // action to take after the guest has shutdown itself
	ActionsAfterShutdown OnNormalExit
  // action to take after the guest has rebooted itself
	ActionsAfterReboot OnNormalExit
  // action to take if the guest crashes
	ActionsAfterCrash OnCrashBehaviour
  // virtual console devices
	Consoles []ConsoleRef
  // virtual network interfaces
	VIFs []VIFRef
  // virtual block devices
	VBDs []VBDRef
  // vitual usb devices
	VUSBs []VUSBRef
  // crash dumps associated with this VM
	CrashDumps []CrashdumpRef
  // virtual TPMs
	VTPMs []VTPMRef
  // name of or path to bootloader
	PVBootloader string
  // path to the kernel
	PVKernel string
  // path to the initrd
	PVRamdisk string
  // kernel command-line arguments
	PVArgs string
  // miscellaneous arguments for the bootloader
	PVBootloaderArgs string
  // to make Zurich guests boot
	PVLegacyArgs string
  // HVM boot policy
	HVMBootPolicy string
  // HVM boot params
	HVMBootParams map[string]string
  // multiplier applied to the amount of shadow that will be made available to the guest
	HVMShadowMultiplier float64
  // platform-specific configuration
	Platform map[string]string
  // PCI bus path for pass-through devices
	PCIBus string
  // additional configuration
	OtherConfig map[string]string
  // domain ID (if available, -1 otherwise)
	Domid int
  // Domain architecture (if available, null string otherwise)
	Domarch string
  // describes the CPU flags on which the VM was last booted
	LastBootCPUFlags map[string]string
  // true if this is a control domain (domain 0 or a driver domain)
	IsControlDomain bool
  // metrics associated with this VM
	Metrics VMMetricsRef
  // metrics associated with the running guest
	GuestMetrics VMGuestMetricsRef
  // marshalled value containing VM record at time of last boot, updated dynamically to reflect the runtime state of the domain
	LastBootedRecord string
  // An XML specification of recommended values and ranges for properties of this VM
	Recommendations string
  // data to be inserted into the xenstore tree (/local/domain/<domid>/vm-data) after the VM is created.
	XenstoreData map[string]string
  // if true then the system will attempt to keep the VM running as much as possible.
	HaAlwaysRun bool
  // has possible values: "best-effort" meaning "try to restart this VM if possible but don't consider the Pool to be overcommitted if this is not possible"; "restart" meaning "this VM should be restarted"; "" meaning "do not try to restart this VM"
	HaRestartPriority string
  // true if this is a snapshot. Snapshotted VMs can never be started, they are used only for cloning other VMs
	IsASnapshot bool
  // Ref pointing to the VM this snapshot is of.
	SnapshotOf VMRef
  // List pointing to all the VM snapshots.
	Snapshots []VMRef
  // Date/time when this snapshot was created.
	SnapshotTime time.Time
  // Transportable ID of the snapshot VM
	TransportableSnapshotID string
  // Binary blobs associated with this VM
	Blobs map[string]BlobRef
  // user-specified tags for categorization purposes
	Tags []string
  // List of operations which have been explicitly blocked and an error code
	BlockedOperations map[VMOperations]string
  // Human-readable information concerning this snapshot
	SnapshotInfo map[string]string
  // Encoded information about the VM's metadata this is a snapshot of
	SnapshotMetadata string
  // Ref pointing to the parent of this VM
	Parent VMRef
  // List pointing to all the children of this VM
	Children []VMRef
  // BIOS strings
	BiosStrings map[string]string
  // Ref pointing to a protection policy for this VM
	ProtectionPolicy VMPPRef
  // true if this snapshot was created by the protection policy
	IsSnapshotFromVmpp bool
  // Ref pointing to a snapshot schedule for this VM
	SnapshotSchedule VMSSRef
  // true if this snapshot was created by the snapshot schedule
	IsVmssSnapshot bool
  // the appliance to which this VM belongs
	Appliance VMApplianceRef
  // The delay to wait before proceeding to the next order in the startup sequence (seconds)
	StartDelay int
  // The delay to wait before proceeding to the next order in the shutdown sequence (seconds)
	ShutdownDelay int
  // The point in the startup or shutdown sequence at which this VM will be started
	Order int
  // Virtual GPUs
	VGPUs []VGPURef
  // Currently passed-through PCI devices
	AttachedPCIs []PCIRef
  // The SR on which a suspend image is stored
	SuspendSR SRRef
  // The number of times this VM has been recovered
	Version int
  // Generation ID of the VM
	GenerationID string
  // The host virtual hardware platform version the VM can run on
	HardwarePlatformVersion int
  // When an HVM guest starts, this controls the presence of the emulated C000 PCI device which triggers Windows Update to fetch or update PV drivers.
	HasVendorDevice bool
  // Indicates whether a VM requires a reboot in order to update its configuration, e.g. its memory allocation.
	RequiresReboot bool
  // Textual reference to the template used to create a VM. This can be used by clients in need of an immutable reference to the template since the latter's uuid and name_label may change, for example, after a package installation or upgrade.
	ReferenceLabel string
}

type VMRef string

// A virtual machine (or 'guest').
type VMClass struct {
	client *Client
}


func VMClassGetAllRecordsMockDefault(sessionID SessionRef) (_retval map[VMRef]VMRecord, _err error) {
	log.Println("VM.GetAllRecords not mocked")
	_err = errors.New("VM.GetAllRecords not mocked")
	return
}

var VMClassGetAllRecordsMockedCallback = VMClassGetAllRecordsMockDefault

func (_class VMClass) GetAllRecordsMock(sessionID SessionRef) (_retval map[VMRef]VMRecord, _err error) {
	return VMClassGetAllRecordsMockedCallback(sessionID)
}
// Return a map of VM references to VM records for all VMs known to the system.
func (_class VMClass) GetAllRecords(sessionID SessionRef) (_retval map[VMRef]VMRecord, _err error) {
	if IsMock {
		return _class.GetAllRecordsMock(sessionID)
	}	
	_method := "VM.get_all_records"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVMRefToVMRecordMapToGo(_method + " -> ", _result.Value)
	return
}


func VMClassGetAllMockDefault(sessionID SessionRef) (_retval []VMRef, _err error) {
	log.Println("VM.GetAll not mocked")
	_err = errors.New("VM.GetAll not mocked")
	return
}

var VMClassGetAllMockedCallback = VMClassGetAllMockDefault

func (_class VMClass) GetAllMock(sessionID SessionRef) (_retval []VMRef, _err error) {
	return VMClassGetAllMockedCallback(sessionID)
}
// Return a list of all the VMs known to the system.
func (_class VMClass) GetAll(sessionID SessionRef) (_retval []VMRef, _err error) {
	if IsMock {
		return _class.GetAllMock(sessionID)
	}	
	_method := "VM.get_all"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVMRefSetToGo(_method + " -> ", _result.Value)
	return
}


func VMClassSetActionsAfterCrashMockDefault(sessionID SessionRef, self VMRef, value OnCrashBehaviour) (_err error) {
	log.Println("VM.SetActionsAfterCrash not mocked")
	_err = errors.New("VM.SetActionsAfterCrash not mocked")
	return
}

var VMClassSetActionsAfterCrashMockedCallback = VMClassSetActionsAfterCrashMockDefault

func (_class VMClass) SetActionsAfterCrashMock(sessionID SessionRef, self VMRef, value OnCrashBehaviour) (_err error) {
	return VMClassSetActionsAfterCrashMockedCallback(sessionID, self, value)
}
// Sets the actions_after_crash parameter
func (_class VMClass) SetActionsAfterCrash(sessionID SessionRef, self VMRef, value OnCrashBehaviour) (_err error) {
	if IsMock {
		return _class.SetActionsAfterCrashMock(sessionID, self, value)
	}	
	_method := "VM.set_actions_after_crash"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertEnumOnCrashBehaviourToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}


func VMClassImportMockDefault(sessionID SessionRef, url string, sr SRRef, fullRestore bool, force bool) (_retval []VMRef, _err error) {
	log.Println("VM.Import not mocked")
	_err = errors.New("VM.Import not mocked")
	return
}

var VMClassImportMockedCallback = VMClassImportMockDefault

func (_class VMClass) ImportMock(sessionID SessionRef, url string, sr SRRef, fullRestore bool, force bool) (_retval []VMRef, _err error) {
	return VMClassImportMockedCallback(sessionID, url, sr, fullRestore, force)
}
// Import an XVA from a URI
func (_class VMClass) Import(sessionID SessionRef, url string, sr SRRef, fullRestore bool, force bool) (_retval []VMRef, _err error) {
	if IsMock {
		return _class.ImportMock(sessionID, url, sr, fullRestore, force)
	}	
	_method := "VM.import"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_urlArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "url"), url)
	if _err != nil {
		return
	}
	_srArg, _err := convertSRRefToXen(fmt.Sprintf("%s(%s)", _method, "sr"), sr)
	if _err != nil {
		return
	}
	_fullRestoreArg, _err := convertBoolToXen(fmt.Sprintf("%s(%s)", _method, "full_restore"), fullRestore)
	if _err != nil {
		return
	}
	_forceArg, _err := convertBoolToXen(fmt.Sprintf("%s(%s)", _method, "force"), force)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _urlArg, _srArg, _fullRestoreArg, _forceArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVMRefSetToGo(_method + " -> ", _result.Value)
	return
}


func VMClassSetHasVendorDeviceMockDefault(sessionID SessionRef, self VMRef, value bool) (_err error) {
	log.Println("VM.SetHasVendorDevice not mocked")
	_err = errors.New("VM.SetHasVendorDevice not mocked")
	return
}

var VMClassSetHasVendorDeviceMockedCallback = VMClassSetHasVendorDeviceMockDefault

func (_class VMClass) SetHasVendorDeviceMock(sessionID SessionRef, self VMRef, value bool) (_err error) {
	return VMClassSetHasVendorDeviceMockedCallback(sessionID, self, value)
}
// Controls whether, when the VM starts in HVM mode, its virtual hardware will include the emulated PCI device for which drivers may be available through Windows Update. Usually this should never be changed on a VM on which Windows has been installed: changing it on such a VM is likely to lead to a crash on next start.
func (_class VMClass) SetHasVendorDevice(sessionID SessionRef, self VMRef, value bool) (_err error) {
	if IsMock {
		return _class.SetHasVendorDeviceMock(sessionID, self, value)
	}	
	_method := "VM.set_has_vendor_device"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertBoolToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}


func VMClassCallPluginMockDefault(sessionID SessionRef, vm VMRef, plugin string, fn string, args map[string]string) (_retval string, _err error) {
	log.Println("VM.CallPlugin not mocked")
	_err = errors.New("VM.CallPlugin not mocked")
	return
}

var VMClassCallPluginMockedCallback = VMClassCallPluginMockDefault

func (_class VMClass) CallPluginMock(sessionID SessionRef, vm VMRef, plugin string, fn string, args map[string]string) (_retval string, _err error) {
	return VMClassCallPluginMockedCallback(sessionID, vm, plugin, fn, args)
}
// Call a XenAPI plugin on this vm
func (_class VMClass) CallPlugin(sessionID SessionRef, vm VMRef, plugin string, fn string, args map[string]string) (_retval string, _err error) {
	if IsMock {
		return _class.CallPluginMock(sessionID, vm, plugin, fn, args)
	}	
	_method := "VM.call_plugin"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_vmArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "vm"), vm)
	if _err != nil {
		return
	}
	_pluginArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "plugin"), plugin)
	if _err != nil {
		return
	}
	_fnArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "fn"), fn)
	if _err != nil {
		return
	}
	_argsArg, _err := convertStringToStringMapToXen(fmt.Sprintf("%s(%s)", _method, "args"), args)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _vmArg, _pluginArg, _fnArg, _argsArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringToGo(_method + " -> ", _result.Value)
	return
}


func VMClassQueryServicesMockDefault(sessionID SessionRef, self VMRef) (_retval map[string]string, _err error) {
	log.Println("VM.QueryServices not mocked")
	_err = errors.New("VM.QueryServices not mocked")
	return
}

var VMClassQueryServicesMockedCallback = VMClassQueryServicesMockDefault

func (_class VMClass) QueryServicesMock(sessionID SessionRef, self VMRef) (_retval map[string]string, _err error) {
	return VMClassQueryServicesMockedCallback(sessionID, self)
}
// Query the system services advertised by this VM and register them. This can only be applied to a system domain.
func (_class VMClass) QueryServices(sessionID SessionRef, self VMRef) (_retval map[string]string, _err error) {
	if IsMock {
		return _class.QueryServicesMock(sessionID, self)
	}	
	_method := "VM.query_services"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringToStringMapToGo(_method + " -> ", _result.Value)
	return
}


func VMClassSetApplianceMockDefault(sessionID SessionRef, self VMRef, value VMApplianceRef) (_err error) {
	log.Println("VM.SetAppliance not mocked")
	_err = errors.New("VM.SetAppliance not mocked")
	return
}

var VMClassSetApplianceMockedCallback = VMClassSetApplianceMockDefault

func (_class VMClass) SetApplianceMock(sessionID SessionRef, self VMRef, value VMApplianceRef) (_err error) {
	return VMClassSetApplianceMockedCallback(sessionID, self, value)
}
// Assign this VM to an appliance.
func (_class VMClass) SetAppliance(sessionID SessionRef, self VMRef, value VMApplianceRef) (_err error) {
	if IsMock {
		return _class.SetApplianceMock(sessionID, self, value)
	}	
	_method := "VM.set_appliance"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertVMApplianceRefToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}


func VMClassImportConvertMockDefault(sessionID SessionRef, atype string, username string, password string, sr SRRef, remoteConfig map[string]string) (_err error) {
	log.Println("VM.ImportConvert not mocked")
	_err = errors.New("VM.ImportConvert not mocked")
	return
}

var VMClassImportConvertMockedCallback = VMClassImportConvertMockDefault

func (_class VMClass) ImportConvertMock(sessionID SessionRef, atype string, username string, password string, sr SRRef, remoteConfig map[string]string) (_err error) {
	return VMClassImportConvertMockedCallback(sessionID, atype, username, password, sr, remoteConfig)
}
// Import using a conversion service.
func (_class VMClass) ImportConvert(sessionID SessionRef, atype string, username string, password string, sr SRRef, remoteConfig map[string]string) (_err error) {
	if IsMock {
		return _class.ImportConvertMock(sessionID, atype, username, password, sr, remoteConfig)
	}	
	_method := "VM.import_convert"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_atypeArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "type"), atype)
	if _err != nil {
		return
	}
	_usernameArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "username"), username)
	if _err != nil {
		return
	}
	_passwordArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "password"), password)
	if _err != nil {
		return
	}
	_srArg, _err := convertSRRefToXen(fmt.Sprintf("%s(%s)", _method, "sr"), sr)
	if _err != nil {
		return
	}
	_remoteConfigArg, _err := convertStringToStringMapToXen(fmt.Sprintf("%s(%s)", _method, "remote_config"), remoteConfig)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _atypeArg, _usernameArg, _passwordArg, _srArg, _remoteConfigArg)
	return
}


func VMClassRecoverMockDefault(sessionID SessionRef, self VMRef, sessionTo SessionRef, force bool) (_err error) {
	log.Println("VM.Recover not mocked")
	_err = errors.New("VM.Recover not mocked")
	return
}

var VMClassRecoverMockedCallback = VMClassRecoverMockDefault

func (_class VMClass) RecoverMock(sessionID SessionRef, self VMRef, sessionTo SessionRef, force bool) (_err error) {
	return VMClassRecoverMockedCallback(sessionID, self, sessionTo, force)
}
// Recover the VM
func (_class VMClass) Recover(sessionID SessionRef, self VMRef, sessionTo SessionRef, force bool) (_err error) {
	if IsMock {
		return _class.RecoverMock(sessionID, self, sessionTo, force)
	}	
	_method := "VM.recover"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_sessionToArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_to"), sessionTo)
	if _err != nil {
		return
	}
	_forceArg, _err := convertBoolToXen(fmt.Sprintf("%s(%s)", _method, "force"), force)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _sessionToArg, _forceArg)
	return
}


func VMClassGetSRsRequiredForRecoveryMockDefault(sessionID SessionRef, self VMRef, sessionTo SessionRef) (_retval []SRRef, _err error) {
	log.Println("VM.GetSRsRequiredForRecovery not mocked")
	_err = errors.New("VM.GetSRsRequiredForRecovery not mocked")
	return
}

var VMClassGetSRsRequiredForRecoveryMockedCallback = VMClassGetSRsRequiredForRecoveryMockDefault

func (_class VMClass) GetSRsRequiredForRecoveryMock(sessionID SessionRef, self VMRef, sessionTo SessionRef) (_retval []SRRef, _err error) {
	return VMClassGetSRsRequiredForRecoveryMockedCallback(sessionID, self, sessionTo)
}
// List all the SR's that are required for the VM to be recovered
func (_class VMClass) GetSRsRequiredForRecovery(sessionID SessionRef, self VMRef, sessionTo SessionRef) (_retval []SRRef, _err error) {
	if IsMock {
		return _class.GetSRsRequiredForRecoveryMock(sessionID, self, sessionTo)
	}	
	_method := "VM.get_SRs_required_for_recovery"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_sessionToArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_to"), sessionTo)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg, _sessionToArg)
	if _err != nil {
		return
	}
	_retval, _err = convertSRRefSetToGo(_method + " -> ", _result.Value)
	return
}


func VMClassAssertCanBeRecoveredMockDefault(sessionID SessionRef, self VMRef, sessionTo SessionRef) (_err error) {
	log.Println("VM.AssertCanBeRecovered not mocked")
	_err = errors.New("VM.AssertCanBeRecovered not mocked")
	return
}

var VMClassAssertCanBeRecoveredMockedCallback = VMClassAssertCanBeRecoveredMockDefault

func (_class VMClass) AssertCanBeRecoveredMock(sessionID SessionRef, self VMRef, sessionTo SessionRef) (_err error) {
	return VMClassAssertCanBeRecoveredMockedCallback(sessionID, self, sessionTo)
}
// Assert whether all SRs required to recover this VM are available.
//
// Errors:
//  VM_IS_PART_OF_AN_APPLIANCE - This operation is not allowed as the VM is part of an appliance.
//  VM_REQUIRES_SR - You attempted to run a VM on a host which doesn't have access to an SR needed by the VM. The VM has at least one VBD attached to a VDI in the SR.
func (_class VMClass) AssertCanBeRecovered(sessionID SessionRef, self VMRef, sessionTo SessionRef) (_err error) {
	if IsMock {
		return _class.AssertCanBeRecoveredMock(sessionID, self, sessionTo)
	}	
	_method := "VM.assert_can_be_recovered"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_sessionToArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_to"), sessionTo)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _sessionToArg)
	return
}


func VMClassSetSuspendVDIMockDefault(sessionID SessionRef, self VMRef, value VDIRef) (_err error) {
	log.Println("VM.SetSuspendVDI not mocked")
	_err = errors.New("VM.SetSuspendVDI not mocked")
	return
}

var VMClassSetSuspendVDIMockedCallback = VMClassSetSuspendVDIMockDefault

func (_class VMClass) SetSuspendVDIMock(sessionID SessionRef, self VMRef, value VDIRef) (_err error) {
	return VMClassSetSuspendVDIMockedCallback(sessionID, self, value)
}
// Set this VM's suspend VDI, which must be indentical to its current one
func (_class VMClass) SetSuspendVDI(sessionID SessionRef, self VMRef, value VDIRef) (_err error) {
	if IsMock {
		return _class.SetSuspendVDIMock(sessionID, self, value)
	}	
	_method := "VM.set_suspend_VDI"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}


func VMClassSetOrderMockDefault(sessionID SessionRef, self VMRef, value int) (_err error) {
	log.Println("VM.SetOrder not mocked")
	_err = errors.New("VM.SetOrder not mocked")
	return
}

var VMClassSetOrderMockedCallback = VMClassSetOrderMockDefault

func (_class VMClass) SetOrderMock(sessionID SessionRef, self VMRef, value int) (_err error) {
	return VMClassSetOrderMockedCallback(sessionID, self, value)
}
// Set this VM's boot order
func (_class VMClass) SetOrder(sessionID SessionRef, self VMRef, value int) (_err error) {
	if IsMock {
		return _class.SetOrderMock(sessionID, self, value)
	}	
	_method := "VM.set_order"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertIntToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}


func VMClassSetShutdownDelayMockDefault(sessionID SessionRef, self VMRef, value int) (_err error) {
	log.Println("VM.SetShutdownDelay not mocked")
	_err = errors.New("VM.SetShutdownDelay not mocked")
	return
}

var VMClassSetShutdownDelayMockedCallback = VMClassSetShutdownDelayMockDefault

func (_class VMClass) SetShutdownDelayMock(sessionID SessionRef, self VMRef, value int) (_err error) {
	return VMClassSetShutdownDelayMockedCallback(sessionID, self, value)
}
// Set this VM's shutdown delay in seconds
func (_class VMClass) SetShutdownDelay(sessionID SessionRef, self VMRef, value int) (_err error) {
	if IsMock {
		return _class.SetShutdownDelayMock(sessionID, self, value)
	}	
	_method := "VM.set_shutdown_delay"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertIntToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}


func VMClassSetStartDelayMockDefault(sessionID SessionRef, self VMRef, value int) (_err error) {
	log.Println("VM.SetStartDelay not mocked")
	_err = errors.New("VM.SetStartDelay not mocked")
	return
}

var VMClassSetStartDelayMockedCallback = VMClassSetStartDelayMockDefault

func (_class VMClass) SetStartDelayMock(sessionID SessionRef, self VMRef, value int) (_err error) {
	return VMClassSetStartDelayMockedCallback(sessionID, self, value)
}
// Set this VM's start delay in seconds
func (_class VMClass) SetStartDelay(sessionID SessionRef, self VMRef, value int) (_err error) {
	if IsMock {
		return _class.SetStartDelayMock(sessionID, self, value)
	}	
	_method := "VM.set_start_delay"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertIntToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}


func VMClassSetSnapshotScheduleMockDefault(sessionID SessionRef, self VMRef, value VMSSRef) (_err error) {
	log.Println("VM.SetSnapshotSchedule not mocked")
	_err = errors.New("VM.SetSnapshotSchedule not mocked")
	return
}

var VMClassSetSnapshotScheduleMockedCallback = VMClassSetSnapshotScheduleMockDefault

func (_class VMClass) SetSnapshotScheduleMock(sessionID SessionRef, self VMRef, value VMSSRef) (_err error) {
	return VMClassSetSnapshotScheduleMockedCallback(sessionID, self, value)
}
// Set the value of the snapshot schedule field
func (_class VMClass) SetSnapshotSchedule(sessionID SessionRef, self VMRef, value VMSSRef) (_err error) {
	if IsMock {
		return _class.SetSnapshotScheduleMock(sessionID, self, value)
	}	
	_method := "VM.set_snapshot_schedule"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertVMSSRefToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}


func VMClassSetProtectionPolicyMockDefault(sessionID SessionRef, self VMRef, value VMPPRef) (_err error) {
	log.Println("VM.SetProtectionPolicy not mocked")
	_err = errors.New("VM.SetProtectionPolicy not mocked")
	return
}

var VMClassSetProtectionPolicyMockedCallback = VMClassSetProtectionPolicyMockDefault

func (_class VMClass) SetProtectionPolicyMock(sessionID SessionRef, self VMRef, value VMPPRef) (_err error) {
	return VMClassSetProtectionPolicyMockedCallback(sessionID, self, value)
}
// Set the value of the protection_policy field
func (_class VMClass) SetProtectionPolicy(sessionID SessionRef, self VMRef, value VMPPRef) (_err error) {
	if IsMock {
		return _class.SetProtectionPolicyMock(sessionID, self, value)
	}	
	_method := "VM.set_protection_policy"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertVMPPRefToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}


func VMClassCopyBiosStringsMockDefault(sessionID SessionRef, vm VMRef, host HostRef) (_err error) {
	log.Println("VM.CopyBiosStrings not mocked")
	_err = errors.New("VM.CopyBiosStrings not mocked")
	return
}

var VMClassCopyBiosStringsMockedCallback = VMClassCopyBiosStringsMockDefault

func (_class VMClass) CopyBiosStringsMock(sessionID SessionRef, vm VMRef, host HostRef) (_err error) {
	return VMClassCopyBiosStringsMockedCallback(sessionID, vm, host)
}
// Copy the BIOS strings from the given host to this VM
func (_class VMClass) CopyBiosStrings(sessionID SessionRef, vm VMRef, host HostRef) (_err error) {
	if IsMock {
		return _class.CopyBiosStringsMock(sessionID, vm, host)
	}	
	_method := "VM.copy_bios_strings"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_vmArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "vm"), vm)
	if _err != nil {
		return
	}
	_hostArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "host"), host)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _vmArg, _hostArg)
	return
}


func VMClassSetBiosStringsMockDefault(sessionID SessionRef, self VMRef, value map[string]string) (_err error) {
	log.Println("VM.SetBiosStrings not mocked")
	_err = errors.New("VM.SetBiosStrings not mocked")
	return
}

var VMClassSetBiosStringsMockedCallback = VMClassSetBiosStringsMockDefault

func (_class VMClass) SetBiosStringsMock(sessionID SessionRef, self VMRef, value map[string]string) (_err error) {
	return VMClassSetBiosStringsMockedCallback(sessionID, self, value)
}
// Set custom BIOS strings to this VM. VM will be given a default set of BIOS strings, only some of which can be overridden by the supplied values. Allowed keys are: 'bios-vendor', 'bios-version', 'system-manufacturer', 'system-product-name', 'system-version', 'system-serial-number', 'enclosure-asset-tag'
//
// Errors:
//  VM_BIOS_STRINGS_ALREADY_SET - The BIOS strings for this VM have already been set and cannot be changed.
//  INVALID_VALUE - The value given is invalid
func (_class VMClass) SetBiosStrings(sessionID SessionRef, self VMRef, value map[string]string) (_err error) {
	if IsMock {
		return _class.SetBiosStringsMock(sessionID, self, value)
	}	
	_method := "VM.set_bios_strings"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertStringToStringMapToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}


func VMClassRetrieveWlbRecommendationsMockDefault(sessionID SessionRef, vm VMRef) (_retval map[HostRef][]string, _err error) {
	log.Println("VM.RetrieveWlbRecommendations not mocked")
	_err = errors.New("VM.RetrieveWlbRecommendations not mocked")
	return
}

var VMClassRetrieveWlbRecommendationsMockedCallback = VMClassRetrieveWlbRecommendationsMockDefault

func (_class VMClass) RetrieveWlbRecommendationsMock(sessionID SessionRef, vm VMRef) (_retval map[HostRef][]string, _err error) {
	return VMClassRetrieveWlbRecommendationsMockedCallback(sessionID, vm)
}
// Returns mapping of hosts to ratings, indicating the suitability of starting the VM at that location according to wlb. Rating is replaced with an error if the VM cannot boot there.
func (_class VMClass) RetrieveWlbRecommendations(sessionID SessionRef, vm VMRef) (_retval map[HostRef][]string, _err error) {
	if IsMock {
		return _class.RetrieveWlbRecommendationsMock(sessionID, vm)
	}	
	_method := "VM.retrieve_wlb_recommendations"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_vmArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "vm"), vm)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _vmArg)
	if _err != nil {
		return
	}
	_retval, _err = convertHostRefToStringSetMapToGo(_method + " -> ", _result.Value)
	return
}


func VMClassAssertAgileMockDefault(sessionID SessionRef, self VMRef) (_err error) {
	log.Println("VM.AssertAgile not mocked")
	_err = errors.New("VM.AssertAgile not mocked")
	return
}

var VMClassAssertAgileMockedCallback = VMClassAssertAgileMockDefault

func (_class VMClass) AssertAgileMock(sessionID SessionRef, self VMRef) (_err error) {
	return VMClassAssertAgileMockedCallback(sessionID, self)
}
// Returns an error if the VM is not considered agile e.g. because it is tied to a resource local to a host
func (_class VMClass) AssertAgile(sessionID SessionRef, self VMRef) (_err error) {
	if IsMock {
		return _class.AssertAgileMock(sessionID, self)
	}	
	_method := "VM.assert_agile"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg)
	return
}


func VMClassCreateNewBlobMockDefault(sessionID SessionRef, vm VMRef, name string, mimeType string, public bool) (_retval BlobRef, _err error) {
	log.Println("VM.CreateNewBlob not mocked")
	_err = errors.New("VM.CreateNewBlob not mocked")
	return
}

var VMClassCreateNewBlobMockedCallback = VMClassCreateNewBlobMockDefault

func (_class VMClass) CreateNewBlobMock(sessionID SessionRef, vm VMRef, name string, mimeType string, public bool) (_retval BlobRef, _err error) {
	return VMClassCreateNewBlobMockedCallback(sessionID, vm, name, mimeType, public)
}
// Create a placeholder for a named binary blob of data that is associated with this VM
func (_class VMClass) CreateNewBlob(sessionID SessionRef, vm VMRef, name string, mimeType string, public bool) (_retval BlobRef, _err error) {
	if IsMock {
		return _class.CreateNewBlobMock(sessionID, vm, name, mimeType, public)
	}	
	_method := "VM.create_new_blob"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_vmArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "vm"), vm)
	if _err != nil {
		return
	}
	_nameArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "name"), name)
	if _err != nil {
		return
	}
	_mimeTypeArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "mime_type"), mimeType)
	if _err != nil {
		return
	}
	_publicArg, _err := convertBoolToXen(fmt.Sprintf("%s(%s)", _method, "public"), public)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _vmArg, _nameArg, _mimeTypeArg, _publicArg)
	if _err != nil {
		return
	}
	_retval, _err = convertBlobRefToGo(_method + " -> ", _result.Value)
	return
}


func VMClassAssertCanBootHereMockDefault(sessionID SessionRef, self VMRef, host HostRef) (_err error) {
	log.Println("VM.AssertCanBootHere not mocked")
	_err = errors.New("VM.AssertCanBootHere not mocked")
	return
}

var VMClassAssertCanBootHereMockedCallback = VMClassAssertCanBootHereMockDefault

func (_class VMClass) AssertCanBootHereMock(sessionID SessionRef, self VMRef, host HostRef) (_err error) {
	return VMClassAssertCanBootHereMockedCallback(sessionID, self, host)
}
// Returns an error if the VM could not boot on this host for some reason
//
// Errors:
//  HOST_NOT_ENOUGH_FREE_MEMORY - Not enough host memory is available to perform this operation
//  VM_REQUIRES_SR - You attempted to run a VM on a host which doesn't have access to an SR needed by the VM. The VM has at least one VBD attached to a VDI in the SR.
//  VM_HOST_INCOMPATIBLE_VERSION - This VM operation cannot be performed on an older-versioned host during an upgrade.
//  VM_HOST_INCOMPATIBLE_VIRTUAL_HARDWARE_PLATFORM_VERSION - You attempted to run a VM on a host that cannot provide the VM's required Virtual Hardware Platform version.
func (_class VMClass) AssertCanBootHere(sessionID SessionRef, self VMRef, host HostRef) (_err error) {
	if IsMock {
		return _class.AssertCanBootHereMock(sessionID, self, host)
	}	
	_method := "VM.assert_can_boot_here"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_hostArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "host"), host)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _hostArg)
	return
}


func VMClassGetPossibleHostsMockDefault(sessionID SessionRef, vm VMRef) (_retval []HostRef, _err error) {
	log.Println("VM.GetPossibleHosts not mocked")
	_err = errors.New("VM.GetPossibleHosts not mocked")
	return
}

var VMClassGetPossibleHostsMockedCallback = VMClassGetPossibleHostsMockDefault

func (_class VMClass) GetPossibleHostsMock(sessionID SessionRef, vm VMRef) (_retval []HostRef, _err error) {
	return VMClassGetPossibleHostsMockedCallback(sessionID, vm)
}
// Return the list of hosts on which this VM may run.
func (_class VMClass) GetPossibleHosts(sessionID SessionRef, vm VMRef) (_retval []HostRef, _err error) {
	if IsMock {
		return _class.GetPossibleHostsMock(sessionID, vm)
	}	
	_method := "VM.get_possible_hosts"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_vmArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "vm"), vm)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _vmArg)
	if _err != nil {
		return
	}
	_retval, _err = convertHostRefSetToGo(_method + " -> ", _result.Value)
	return
}


func VMClassGetAllowedVIFDevicesMockDefault(sessionID SessionRef, vm VMRef) (_retval []string, _err error) {
	log.Println("VM.GetAllowedVIFDevices not mocked")
	_err = errors.New("VM.GetAllowedVIFDevices not mocked")
	return
}

var VMClassGetAllowedVIFDevicesMockedCallback = VMClassGetAllowedVIFDevicesMockDefault

func (_class VMClass) GetAllowedVIFDevicesMock(sessionID SessionRef, vm VMRef) (_retval []string, _err error) {
	return VMClassGetAllowedVIFDevicesMockedCallback(sessionID, vm)
}
// Returns a list of the allowed values that a VIF device field can take
func (_class VMClass) GetAllowedVIFDevices(sessionID SessionRef, vm VMRef) (_retval []string, _err error) {
	if IsMock {
		return _class.GetAllowedVIFDevicesMock(sessionID, vm)
	}	
	_method := "VM.get_allowed_VIF_devices"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_vmArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "vm"), vm)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _vmArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringSetToGo(_method + " -> ", _result.Value)
	return
}


func VMClassGetAllowedVBDDevicesMockDefault(sessionID SessionRef, vm VMRef) (_retval []string, _err error) {
	log.Println("VM.GetAllowedVBDDevices not mocked")
	_err = errors.New("VM.GetAllowedVBDDevices not mocked")
	return
}

var VMClassGetAllowedVBDDevicesMockedCallback = VMClassGetAllowedVBDDevicesMockDefault

func (_class VMClass) GetAllowedVBDDevicesMock(sessionID SessionRef, vm VMRef) (_retval []string, _err error) {
	return VMClassGetAllowedVBDDevicesMockedCallback(sessionID, vm)
}
// Returns a list of the allowed values that a VBD device field can take
func (_class VMClass) GetAllowedVBDDevices(sessionID SessionRef, vm VMRef) (_retval []string, _err error) {
	if IsMock {
		return _class.GetAllowedVBDDevicesMock(sessionID, vm)
	}	
	_method := "VM.get_allowed_VBD_devices"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_vmArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "vm"), vm)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _vmArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringSetToGo(_method + " -> ", _result.Value)
	return
}


func VMClassUpdateAllowedOperationsMockDefault(sessionID SessionRef, self VMRef) (_err error) {
	log.Println("VM.UpdateAllowedOperations not mocked")
	_err = errors.New("VM.UpdateAllowedOperations not mocked")
	return
}

var VMClassUpdateAllowedOperationsMockedCallback = VMClassUpdateAllowedOperationsMockDefault

func (_class VMClass) UpdateAllowedOperationsMock(sessionID SessionRef, self VMRef) (_err error) {
	return VMClassUpdateAllowedOperationsMockedCallback(sessionID, self)
}
// Recomputes the list of acceptable operations
func (_class VMClass) UpdateAllowedOperations(sessionID SessionRef, self VMRef) (_err error) {
	if IsMock {
		return _class.UpdateAllowedOperationsMock(sessionID, self)
	}	
	_method := "VM.update_allowed_operations"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg)
	return
}


func VMClassAssertOperationValidMockDefault(sessionID SessionRef, self VMRef, op VMOperations) (_err error) {
	log.Println("VM.AssertOperationValid not mocked")
	_err = errors.New("VM.AssertOperationValid not mocked")
	return
}

var VMClassAssertOperationValidMockedCallback = VMClassAssertOperationValidMockDefault

func (_class VMClass) AssertOperationValidMock(sessionID SessionRef, self VMRef, op VMOperations) (_err error) {
	return VMClassAssertOperationValidMockedCallback(sessionID, self, op)
}
// Check to see whether this operation is acceptable in the current state of the system, raising an error if the operation is invalid for some reason
func (_class VMClass) AssertOperationValid(sessionID SessionRef, self VMRef, op VMOperations) (_err error) {
	if IsMock {
		return _class.AssertOperationValidMock(sessionID, self, op)
	}	
	_method := "VM.assert_operation_valid"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_opArg, _err := convertEnumVMOperationsToXen(fmt.Sprintf("%s(%s)", _method, "op"), op)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _opArg)
	return
}


func VMClassForgetDataSourceArchivesMockDefault(sessionID SessionRef, self VMRef, dataSource string) (_err error) {
	log.Println("VM.ForgetDataSourceArchives not mocked")
	_err = errors.New("VM.ForgetDataSourceArchives not mocked")
	return
}

var VMClassForgetDataSourceArchivesMockedCallback = VMClassForgetDataSourceArchivesMockDefault

func (_class VMClass) ForgetDataSourceArchivesMock(sessionID SessionRef, self VMRef, dataSource string) (_err error) {
	return VMClassForgetDataSourceArchivesMockedCallback(sessionID, self, dataSource)
}
// Forget the recorded statistics related to the specified data source
func (_class VMClass) ForgetDataSourceArchives(sessionID SessionRef, self VMRef, dataSource string) (_err error) {
	if IsMock {
		return _class.ForgetDataSourceArchivesMock(sessionID, self, dataSource)
	}	
	_method := "VM.forget_data_source_archives"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_dataSourceArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "data_source"), dataSource)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _dataSourceArg)
	return
}


func VMClassQueryDataSourceMockDefault(sessionID SessionRef, self VMRef, dataSource string) (_retval float64, _err error) {
	log.Println("VM.QueryDataSource not mocked")
	_err = errors.New("VM.QueryDataSource not mocked")
	return
}

var VMClassQueryDataSourceMockedCallback = VMClassQueryDataSourceMockDefault

func (_class VMClass) QueryDataSourceMock(sessionID SessionRef, self VMRef, dataSource string) (_retval float64, _err error) {
	return VMClassQueryDataSourceMockedCallback(sessionID, self, dataSource)
}
// Query the latest value of the specified data source
func (_class VMClass) QueryDataSource(sessionID SessionRef, self VMRef, dataSource string) (_retval float64, _err error) {
	if IsMock {
		return _class.QueryDataSourceMock(sessionID, self, dataSource)
	}	
	_method := "VM.query_data_source"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_dataSourceArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "data_source"), dataSource)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg, _dataSourceArg)
	if _err != nil {
		return
	}
	_retval, _err = convertFloatToGo(_method + " -> ", _result.Value)
	return
}


func VMClassRecordDataSourceMockDefault(sessionID SessionRef, self VMRef, dataSource string) (_err error) {
	log.Println("VM.RecordDataSource not mocked")
	_err = errors.New("VM.RecordDataSource not mocked")
	return
}

var VMClassRecordDataSourceMockedCallback = VMClassRecordDataSourceMockDefault

func (_class VMClass) RecordDataSourceMock(sessionID SessionRef, self VMRef, dataSource string) (_err error) {
	return VMClassRecordDataSourceMockedCallback(sessionID, self, dataSource)
}
// Start recording the specified data source
func (_class VMClass) RecordDataSource(sessionID SessionRef, self VMRef, dataSource string) (_err error) {
	if IsMock {
		return _class.RecordDataSourceMock(sessionID, self, dataSource)
	}	
	_method := "VM.record_data_source"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_dataSourceArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "data_source"), dataSource)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _dataSourceArg)
	return
}


func VMClassGetDataSourcesMockDefault(sessionID SessionRef, self VMRef) (_retval []DataSourceRecord, _err error) {
	log.Println("VM.GetDataSources not mocked")
	_err = errors.New("VM.GetDataSources not mocked")
	return
}

var VMClassGetDataSourcesMockedCallback = VMClassGetDataSourcesMockDefault

func (_class VMClass) GetDataSourcesMock(sessionID SessionRef, self VMRef) (_retval []DataSourceRecord, _err error) {
	return VMClassGetDataSourcesMockedCallback(sessionID, self)
}
// 
func (_class VMClass) GetDataSources(sessionID SessionRef, self VMRef) (_retval []DataSourceRecord, _err error) {
	if IsMock {
		return _class.GetDataSourcesMock(sessionID, self)
	}	
	_method := "VM.get_data_sources"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertDataSourceRecordSetToGo(_method + " -> ", _result.Value)
	return
}


func VMClassGetBootRecordMockDefault(sessionID SessionRef, self VMRef) (_retval VMRecord, _err error) {
	log.Println("VM.GetBootRecord not mocked")
	_err = errors.New("VM.GetBootRecord not mocked")
	return
}

var VMClassGetBootRecordMockedCallback = VMClassGetBootRecordMockDefault

func (_class VMClass) GetBootRecordMock(sessionID SessionRef, self VMRef) (_retval VMRecord, _err error) {
	return VMClassGetBootRecordMockedCallback(sessionID, self)
}
// Returns a record describing the VM's dynamic state, initialised when the VM boots and updated to reflect runtime configuration changes e.g. CPU hotplug
func (_class VMClass) GetBootRecord(sessionID SessionRef, self VMRef) (_retval VMRecord, _err error) {
	if IsMock {
		return _class.GetBootRecordMock(sessionID, self)
	}	
	_method := "VM.get_boot_record"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVMRecordToGo(_method + " -> ", _result.Value)
	return
}


func VMClassAssertCanMigrateMockDefault(sessionID SessionRef, vm VMRef, dest map[string]string, live bool, vdiMap map[VDIRef]SRRef, vifMap map[VIFRef]NetworkRef, options map[string]string, vgpuMap map[VGPURef]GPUGroupRef) (_err error) {
	log.Println("VM.AssertCanMigrate not mocked")
	_err = errors.New("VM.AssertCanMigrate not mocked")
	return
}

var VMClassAssertCanMigrateMockedCallback = VMClassAssertCanMigrateMockDefault

func (_class VMClass) AssertCanMigrateMock(sessionID SessionRef, vm VMRef, dest map[string]string, live bool, vdiMap map[VDIRef]SRRef, vifMap map[VIFRef]NetworkRef, options map[string]string, vgpuMap map[VGPURef]GPUGroupRef) (_err error) {
	return VMClassAssertCanMigrateMockedCallback(sessionID, vm, dest, live, vdiMap, vifMap, options, vgpuMap)
}
// Assert whether a VM can be migrated to the specified destination.
//
// Errors:
//  LICENCE_RESTRICTION - This operation is not allowed because your license lacks a needed feature.  Please contact your support representative.
func (_class VMClass) AssertCanMigrate(sessionID SessionRef, vm VMRef, dest map[string]string, live bool, vdiMap map[VDIRef]SRRef, vifMap map[VIFRef]NetworkRef, options map[string]string, vgpuMap map[VGPURef]GPUGroupRef) (_err error) {
	if IsMock {
		return _class.AssertCanMigrateMock(sessionID, vm, dest, live, vdiMap, vifMap, options, vgpuMap)
	}	
	_method := "VM.assert_can_migrate"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_vmArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "vm"), vm)
	if _err != nil {
		return
	}
	_destArg, _err := convertStringToStringMapToXen(fmt.Sprintf("%s(%s)", _method, "dest"), dest)
	if _err != nil {
		return
	}
	_liveArg, _err := convertBoolToXen(fmt.Sprintf("%s(%s)", _method, "live"), live)
	if _err != nil {
		return
	}
	_vdiMapArg, _err := convertVDIRefToSRRefMapToXen(fmt.Sprintf("%s(%s)", _method, "vdi_map"), vdiMap)
	if _err != nil {
		return
	}
	_vifMapArg, _err := convertVIFRefToNetworkRefMapToXen(fmt.Sprintf("%s(%s)", _method, "vif_map"), vifMap)
	if _err != nil {
		return
	}
	_optionsArg, _err := convertStringToStringMapToXen(fmt.Sprintf("%s(%s)", _method, "options"), options)
	if _err != nil {
		return
	}
	_vgpuMapArg, _err := convertVGPURefToGPUGroupRefMapToXen(fmt.Sprintf("%s(%s)", _method, "vgpu_map"), vgpuMap)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _vmArg, _destArg, _liveArg, _vdiMapArg, _vifMapArg, _optionsArg, _vgpuMapArg)
	return
}


func VMClassMigrateSendMockDefault(sessionID SessionRef, vm VMRef, dest map[string]string, live bool, vdiMap map[VDIRef]SRRef, vifMap map[VIFRef]NetworkRef, options map[string]string, vgpuMap map[VGPURef]GPUGroupRef) (_retval VMRef, _err error) {
	log.Println("VM.MigrateSend not mocked")
	_err = errors.New("VM.MigrateSend not mocked")
	return
}

var VMClassMigrateSendMockedCallback = VMClassMigrateSendMockDefault

func (_class VMClass) MigrateSendMock(sessionID SessionRef, vm VMRef, dest map[string]string, live bool, vdiMap map[VDIRef]SRRef, vifMap map[VIFRef]NetworkRef, options map[string]string, vgpuMap map[VGPURef]GPUGroupRef) (_retval VMRef, _err error) {
	return VMClassMigrateSendMockedCallback(sessionID, vm, dest, live, vdiMap, vifMap, options, vgpuMap)
}
// Migrate the VM to another host.  This can only be called when the specified VM is in the Running state.
//
// Errors:
//  VM_BAD_POWER_STATE - You attempted an operation on a VM that was not in an appropriate power state at the time; for example, you attempted to start a VM that was already running.  The parameters returned are the VM's handle, and the expected and actual VM state at the time of the call.
//  LICENCE_RESTRICTION - This operation is not allowed because your license lacks a needed feature.  Please contact your support representative.
func (_class VMClass) MigrateSend(sessionID SessionRef, vm VMRef, dest map[string]string, live bool, vdiMap map[VDIRef]SRRef, vifMap map[VIFRef]NetworkRef, options map[string]string, vgpuMap map[VGPURef]GPUGroupRef) (_retval VMRef, _err error) {
	if IsMock {
		return _class.MigrateSendMock(sessionID, vm, dest, live, vdiMap, vifMap, options, vgpuMap)
	}	
	_method := "VM.migrate_send"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_vmArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "vm"), vm)
	if _err != nil {
		return
	}
	_destArg, _err := convertStringToStringMapToXen(fmt.Sprintf("%s(%s)", _method, "dest"), dest)
	if _err != nil {
		return
	}
	_liveArg, _err := convertBoolToXen(fmt.Sprintf("%s(%s)", _method, "live"), live)
	if _err != nil {
		return
	}
	_vdiMapArg, _err := convertVDIRefToSRRefMapToXen(fmt.Sprintf("%s(%s)", _method, "vdi_map"), vdiMap)
	if _err != nil {
		return
	}
	_vifMapArg, _err := convertVIFRefToNetworkRefMapToXen(fmt.Sprintf("%s(%s)", _method, "vif_map"), vifMap)
	if _err != nil {
		return
	}
	_optionsArg, _err := convertStringToStringMapToXen(fmt.Sprintf("%s(%s)", _method, "options"), options)
	if _err != nil {
		return
	}
	_vgpuMapArg, _err := convertVGPURefToGPUGroupRefMapToXen(fmt.Sprintf("%s(%s)", _method, "vgpu_map"), vgpuMap)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _vmArg, _destArg, _liveArg, _vdiMapArg, _vifMapArg, _optionsArg, _vgpuMapArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVMRefToGo(_method + " -> ", _result.Value)
	return
}


func VMClassMaximiseMemoryMockDefault(sessionID SessionRef, self VMRef, total int, approximate bool) (_retval int, _err error) {
	log.Println("VM.MaximiseMemory not mocked")
	_err = errors.New("VM.MaximiseMemory not mocked")
	return
}

var VMClassMaximiseMemoryMockedCallback = VMClassMaximiseMemoryMockDefault

func (_class VMClass) MaximiseMemoryMock(sessionID SessionRef, self VMRef, total int, approximate bool) (_retval int, _err error) {
	return VMClassMaximiseMemoryMockedCallback(sessionID, self, total, approximate)
}
// Returns the maximum amount of guest memory which will fit, together with overheads, in the supplied amount of physical memory. If 'exact' is true then an exact calculation is performed using the VM's current settings. If 'exact' is false then a more conservative approximation is used
func (_class VMClass) MaximiseMemory(sessionID SessionRef, self VMRef, total int, approximate bool) (_retval int, _err error) {
	if IsMock {
		return _class.MaximiseMemoryMock(sessionID, self, total, approximate)
	}	
	_method := "VM.maximise_memory"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_totalArg, _err := convertIntToXen(fmt.Sprintf("%s(%s)", _method, "total"), total)
	if _err != nil {
		return
	}
	_approximateArg, _err := convertBoolToXen(fmt.Sprintf("%s(%s)", _method, "approximate"), approximate)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg, _totalArg, _approximateArg)
	if _err != nil {
		return
	}
	_retval, _err = convertIntToGo(_method + " -> ", _result.Value)
	return
}


func VMClassSendTriggerMockDefault(sessionID SessionRef, vm VMRef, trigger string) (_err error) {
	log.Println("VM.SendTrigger not mocked")
	_err = errors.New("VM.SendTrigger not mocked")
	return
}

var VMClassSendTriggerMockedCallback = VMClassSendTriggerMockDefault

func (_class VMClass) SendTriggerMock(sessionID SessionRef, vm VMRef, trigger string) (_err error) {
	return VMClassSendTriggerMockedCallback(sessionID, vm, trigger)
}
// Send the named trigger to this VM.  This can only be called when the specified VM is in the Running state.
//
// Errors:
//  VM_BAD_POWER_STATE - You attempted an operation on a VM that was not in an appropriate power state at the time; for example, you attempted to start a VM that was already running.  The parameters returned are the VM's handle, and the expected and actual VM state at the time of the call.
func (_class VMClass) SendTrigger(sessionID SessionRef, vm VMRef, trigger string) (_err error) {
	if IsMock {
		return _class.SendTriggerMock(sessionID, vm, trigger)
	}	
	_method := "VM.send_trigger"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_vmArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "vm"), vm)
	if _err != nil {
		return
	}
	_triggerArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "trigger"), trigger)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _vmArg, _triggerArg)
	return
}


func VMClassSendSysrqMockDefault(sessionID SessionRef, vm VMRef, key string) (_err error) {
	log.Println("VM.SendSysrq not mocked")
	_err = errors.New("VM.SendSysrq not mocked")
	return
}

var VMClassSendSysrqMockedCallback = VMClassSendSysrqMockDefault

func (_class VMClass) SendSysrqMock(sessionID SessionRef, vm VMRef, key string) (_err error) {
	return VMClassSendSysrqMockedCallback(sessionID, vm, key)
}
// Send the given key as a sysrq to this VM.  The key is specified as a single character (a String of length 1).  This can only be called when the specified VM is in the Running state.
//
// Errors:
//  VM_BAD_POWER_STATE - You attempted an operation on a VM that was not in an appropriate power state at the time; for example, you attempted to start a VM that was already running.  The parameters returned are the VM's handle, and the expected and actual VM state at the time of the call.
func (_class VMClass) SendSysrq(sessionID SessionRef, vm VMRef, key string) (_err error) {
	if IsMock {
		return _class.SendSysrqMock(sessionID, vm, key)
	}	
	_method := "VM.send_sysrq"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_vmArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "vm"), vm)
	if _err != nil {
		return
	}
	_keyArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "key"), key)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _vmArg, _keyArg)
	return
}


func VMClassSetVCPUsAtStartupMockDefault(sessionID SessionRef, self VMRef, value int) (_err error) {
	log.Println("VM.SetVCPUsAtStartup not mocked")
	_err = errors.New("VM.SetVCPUsAtStartup not mocked")
	return
}

var VMClassSetVCPUsAtStartupMockedCallback = VMClassSetVCPUsAtStartupMockDefault

func (_class VMClass) SetVCPUsAtStartupMock(sessionID SessionRef, self VMRef, value int) (_err error) {
	return VMClassSetVCPUsAtStartupMockedCallback(sessionID, self, value)
}
// Set the number of startup VCPUs for a halted VM
func (_class VMClass) SetVCPUsAtStartup(sessionID SessionRef, self VMRef, value int) (_err error) {
	if IsMock {
		return _class.SetVCPUsAtStartupMock(sessionID, self, value)
	}	
	_method := "VM.set_VCPUs_at_startup"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertIntToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}


func VMClassSetVCPUsMaxMockDefault(sessionID SessionRef, self VMRef, value int) (_err error) {
	log.Println("VM.SetVCPUsMax not mocked")
	_err = errors.New("VM.SetVCPUsMax not mocked")
	return
}

var VMClassSetVCPUsMaxMockedCallback = VMClassSetVCPUsMaxMockDefault

func (_class VMClass) SetVCPUsMaxMock(sessionID SessionRef, self VMRef, value int) (_err error) {
	return VMClassSetVCPUsMaxMockedCallback(sessionID, self, value)
}
// Set the maximum number of VCPUs for a halted VM
func (_class VMClass) SetVCPUsMax(sessionID SessionRef, self VMRef, value int) (_err error) {
	if IsMock {
		return _class.SetVCPUsMaxMock(sessionID, self, value)
	}	
	_method := "VM.set_VCPUs_max"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertIntToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}


func VMClassSetShadowMultiplierLiveMockDefault(sessionID SessionRef, self VMRef, multiplier float64) (_err error) {
	log.Println("VM.SetShadowMultiplierLive not mocked")
	_err = errors.New("VM.SetShadowMultiplierLive not mocked")
	return
}

var VMClassSetShadowMultiplierLiveMockedCallback = VMClassSetShadowMultiplierLiveMockDefault

func (_class VMClass) SetShadowMultiplierLiveMock(sessionID SessionRef, self VMRef, multiplier float64) (_err error) {
	return VMClassSetShadowMultiplierLiveMockedCallback(sessionID, self, multiplier)
}
// Set the shadow memory multiplier on a running VM
func (_class VMClass) SetShadowMultiplierLive(sessionID SessionRef, self VMRef, multiplier float64) (_err error) {
	if IsMock {
		return _class.SetShadowMultiplierLiveMock(sessionID, self, multiplier)
	}	
	_method := "VM.set_shadow_multiplier_live"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_multiplierArg, _err := convertFloatToXen(fmt.Sprintf("%s(%s)", _method, "multiplier"), multiplier)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _multiplierArg)
	return
}


func VMClassSetHVMShadowMultiplierMockDefault(sessionID SessionRef, self VMRef, value float64) (_err error) {
	log.Println("VM.SetHVMShadowMultiplier not mocked")
	_err = errors.New("VM.SetHVMShadowMultiplier not mocked")
	return
}

var VMClassSetHVMShadowMultiplierMockedCallback = VMClassSetHVMShadowMultiplierMockDefault

func (_class VMClass) SetHVMShadowMultiplierMock(sessionID SessionRef, self VMRef, value float64) (_err error) {
	return VMClassSetHVMShadowMultiplierMockedCallback(sessionID, self, value)
}
// Set the shadow memory multiplier on a halted VM
func (_class VMClass) SetHVMShadowMultiplier(sessionID SessionRef, self VMRef, value float64) (_err error) {
	if IsMock {
		return _class.SetHVMShadowMultiplierMock(sessionID, self, value)
	}	
	_method := "VM.set_HVM_shadow_multiplier"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertFloatToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}


func VMClassGetCooperativeMockDefault(sessionID SessionRef, self VMRef) (_retval bool, _err error) {
	log.Println("VM.GetCooperative not mocked")
	_err = errors.New("VM.GetCooperative not mocked")
	return
}

var VMClassGetCooperativeMockedCallback = VMClassGetCooperativeMockDefault

func (_class VMClass) GetCooperativeMock(sessionID SessionRef, self VMRef) (_retval bool, _err error) {
	return VMClassGetCooperativeMockedCallback(sessionID, self)
}
// Return true if the VM is currently 'co-operative' i.e. is expected to reach a balloon target and actually has done
func (_class VMClass) GetCooperative(sessionID SessionRef, self VMRef) (_retval bool, _err error) {
	if IsMock {
		return _class.GetCooperativeMock(sessionID, self)
	}	
	_method := "VM.get_cooperative"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertBoolToGo(_method + " -> ", _result.Value)
	return
}


func VMClassWaitMemoryTargetLiveMockDefault(sessionID SessionRef, self VMRef) (_err error) {
	log.Println("VM.WaitMemoryTargetLive not mocked")
	_err = errors.New("VM.WaitMemoryTargetLive not mocked")
	return
}

var VMClassWaitMemoryTargetLiveMockedCallback = VMClassWaitMemoryTargetLiveMockDefault

func (_class VMClass) WaitMemoryTargetLiveMock(sessionID SessionRef, self VMRef) (_err error) {
	return VMClassWaitMemoryTargetLiveMockedCallback(sessionID, self)
}
// Wait for a running VM to reach its current memory target
func (_class VMClass) WaitMemoryTargetLive(sessionID SessionRef, self VMRef) (_err error) {
	if IsMock {
		return _class.WaitMemoryTargetLiveMock(sessionID, self)
	}	
	_method := "VM.wait_memory_target_live"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg)
	return
}


func VMClassSetMemoryTargetLiveMockDefault(sessionID SessionRef, self VMRef, target int) (_err error) {
	log.Println("VM.SetMemoryTargetLive not mocked")
	_err = errors.New("VM.SetMemoryTargetLive not mocked")
	return
}

var VMClassSetMemoryTargetLiveMockedCallback = VMClassSetMemoryTargetLiveMockDefault

func (_class VMClass) SetMemoryTargetLiveMock(sessionID SessionRef, self VMRef, target int) (_err error) {
	return VMClassSetMemoryTargetLiveMockedCallback(sessionID, self, target)
}
// Set the memory target for a running VM
func (_class VMClass) SetMemoryTargetLive(sessionID SessionRef, self VMRef, target int) (_err error) {
	if IsMock {
		return _class.SetMemoryTargetLiveMock(sessionID, self, target)
	}	
	_method := "VM.set_memory_target_live"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_targetArg, _err := convertIntToXen(fmt.Sprintf("%s(%s)", _method, "target"), target)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _targetArg)
	return
}


func VMClassSetMemoryMockDefault(sessionID SessionRef, self VMRef, value int) (_err error) {
	log.Println("VM.SetMemory not mocked")
	_err = errors.New("VM.SetMemory not mocked")
	return
}

var VMClassSetMemoryMockedCallback = VMClassSetMemoryMockDefault

func (_class VMClass) SetMemoryMock(sessionID SessionRef, self VMRef, value int) (_err error) {
	return VMClassSetMemoryMockedCallback(sessionID, self, value)
}
// Set the memory allocation of this VM. Sets all of memory_static_max, memory_dynamic_min, and memory_dynamic_max to the given value, and leaves memory_static_min untouched.
func (_class VMClass) SetMemory(sessionID SessionRef, self VMRef, value int) (_err error) {
	if IsMock {
		return _class.SetMemoryMock(sessionID, self, value)
	}	
	_method := "VM.set_memory"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertIntToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}


func VMClassSetMemoryLimitsMockDefault(sessionID SessionRef, self VMRef, staticMin int, staticMax int, dynamicMin int, dynamicMax int) (_err error) {
	log.Println("VM.SetMemoryLimits not mocked")
	_err = errors.New("VM.SetMemoryLimits not mocked")
	return
}

var VMClassSetMemoryLimitsMockedCallback = VMClassSetMemoryLimitsMockDefault

func (_class VMClass) SetMemoryLimitsMock(sessionID SessionRef, self VMRef, staticMin int, staticMax int, dynamicMin int, dynamicMax int) (_err error) {
	return VMClassSetMemoryLimitsMockedCallback(sessionID, self, staticMin, staticMax, dynamicMin, dynamicMax)
}
// Set the memory limits of this VM.
func (_class VMClass) SetMemoryLimits(sessionID SessionRef, self VMRef, staticMin int, staticMax int, dynamicMin int, dynamicMax int) (_err error) {
	if IsMock {
		return _class.SetMemoryLimitsMock(sessionID, self, staticMin, staticMax, dynamicMin, dynamicMax)
	}	
	_method := "VM.set_memory_limits"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_staticMinArg, _err := convertIntToXen(fmt.Sprintf("%s(%s)", _method, "static_min"), staticMin)
	if _err != nil {
		return
	}
	_staticMaxArg, _err := convertIntToXen(fmt.Sprintf("%s(%s)", _method, "static_max"), staticMax)
	if _err != nil {
		return
	}
	_dynamicMinArg, _err := convertIntToXen(fmt.Sprintf("%s(%s)", _method, "dynamic_min"), dynamicMin)
	if _err != nil {
		return
	}
	_dynamicMaxArg, _err := convertIntToXen(fmt.Sprintf("%s(%s)", _method, "dynamic_max"), dynamicMax)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _staticMinArg, _staticMaxArg, _dynamicMinArg, _dynamicMaxArg)
	return
}


func VMClassSetMemoryStaticRangeMockDefault(sessionID SessionRef, self VMRef, min int, max int) (_err error) {
	log.Println("VM.SetMemoryStaticRange not mocked")
	_err = errors.New("VM.SetMemoryStaticRange not mocked")
	return
}

var VMClassSetMemoryStaticRangeMockedCallback = VMClassSetMemoryStaticRangeMockDefault

func (_class VMClass) SetMemoryStaticRangeMock(sessionID SessionRef, self VMRef, min int, max int) (_err error) {
	return VMClassSetMemoryStaticRangeMockedCallback(sessionID, self, min, max)
}
// Set the static (ie boot-time) range of virtual memory that the VM is allowed to use.
func (_class VMClass) SetMemoryStaticRange(sessionID SessionRef, self VMRef, min int, max int) (_err error) {
	if IsMock {
		return _class.SetMemoryStaticRangeMock(sessionID, self, min, max)
	}	
	_method := "VM.set_memory_static_range"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_minArg, _err := convertIntToXen(fmt.Sprintf("%s(%s)", _method, "min"), min)
	if _err != nil {
		return
	}
	_maxArg, _err := convertIntToXen(fmt.Sprintf("%s(%s)", _method, "max"), max)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _minArg, _maxArg)
	return
}


func VMClassSetMemoryStaticMinMockDefault(sessionID SessionRef, self VMRef, value int) (_err error) {
	log.Println("VM.SetMemoryStaticMin not mocked")
	_err = errors.New("VM.SetMemoryStaticMin not mocked")
	return
}

var VMClassSetMemoryStaticMinMockedCallback = VMClassSetMemoryStaticMinMockDefault

func (_class VMClass) SetMemoryStaticMinMock(sessionID SessionRef, self VMRef, value int) (_err error) {
	return VMClassSetMemoryStaticMinMockedCallback(sessionID, self, value)
}
// Set the value of the memory_static_min field
func (_class VMClass) SetMemoryStaticMin(sessionID SessionRef, self VMRef, value int) (_err error) {
	if IsMock {
		return _class.SetMemoryStaticMinMock(sessionID, self, value)
	}	
	_method := "VM.set_memory_static_min"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertIntToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}


func VMClassSetMemoryStaticMaxMockDefault(sessionID SessionRef, self VMRef, value int) (_err error) {
	log.Println("VM.SetMemoryStaticMax not mocked")
	_err = errors.New("VM.SetMemoryStaticMax not mocked")
	return
}

var VMClassSetMemoryStaticMaxMockedCallback = VMClassSetMemoryStaticMaxMockDefault

func (_class VMClass) SetMemoryStaticMaxMock(sessionID SessionRef, self VMRef, value int) (_err error) {
	return VMClassSetMemoryStaticMaxMockedCallback(sessionID, self, value)
}
// Set the value of the memory_static_max field
//
// Errors:
//  HA_OPERATION_WOULD_BREAK_FAILOVER_PLAN - This operation cannot be performed because it would invalidate VM failover planning such that the system would be unable to guarantee to restart protected VMs after a Host failure.
func (_class VMClass) SetMemoryStaticMax(sessionID SessionRef, self VMRef, value int) (_err error) {
	if IsMock {
		return _class.SetMemoryStaticMaxMock(sessionID, self, value)
	}	
	_method := "VM.set_memory_static_max"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertIntToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}


func VMClassSetMemoryDynamicRangeMockDefault(sessionID SessionRef, self VMRef, min int, max int) (_err error) {
	log.Println("VM.SetMemoryDynamicRange not mocked")
	_err = errors.New("VM.SetMemoryDynamicRange not mocked")
	return
}

var VMClassSetMemoryDynamicRangeMockedCallback = VMClassSetMemoryDynamicRangeMockDefault

func (_class VMClass) SetMemoryDynamicRangeMock(sessionID SessionRef, self VMRef, min int, max int) (_err error) {
	return VMClassSetMemoryDynamicRangeMockedCallback(sessionID, self, min, max)
}
// Set the minimum and maximum amounts of physical memory the VM is allowed to use.
func (_class VMClass) SetMemoryDynamicRange(sessionID SessionRef, self VMRef, min int, max int) (_err error) {
	if IsMock {
		return _class.SetMemoryDynamicRangeMock(sessionID, self, min, max)
	}	
	_method := "VM.set_memory_dynamic_range"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_minArg, _err := convertIntToXen(fmt.Sprintf("%s(%s)", _method, "min"), min)
	if _err != nil {
		return
	}
	_maxArg, _err := convertIntToXen(fmt.Sprintf("%s(%s)", _method, "max"), max)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _minArg, _maxArg)
	return
}


func VMClassSetMemoryDynamicMinMockDefault(sessionID SessionRef, self VMRef, value int) (_err error) {
	log.Println("VM.SetMemoryDynamicMin not mocked")
	_err = errors.New("VM.SetMemoryDynamicMin not mocked")
	return
}

var VMClassSetMemoryDynamicMinMockedCallback = VMClassSetMemoryDynamicMinMockDefault

func (_class VMClass) SetMemoryDynamicMinMock(sessionID SessionRef, self VMRef, value int) (_err error) {
	return VMClassSetMemoryDynamicMinMockedCallback(sessionID, self, value)
}
// Set the value of the memory_dynamic_min field
func (_class VMClass) SetMemoryDynamicMin(sessionID SessionRef, self VMRef, value int) (_err error) {
	if IsMock {
		return _class.SetMemoryDynamicMinMock(sessionID, self, value)
	}	
	_method := "VM.set_memory_dynamic_min"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertIntToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}


func VMClassSetMemoryDynamicMaxMockDefault(sessionID SessionRef, self VMRef, value int) (_err error) {
	log.Println("VM.SetMemoryDynamicMax not mocked")
	_err = errors.New("VM.SetMemoryDynamicMax not mocked")
	return
}

var VMClassSetMemoryDynamicMaxMockedCallback = VMClassSetMemoryDynamicMaxMockDefault

func (_class VMClass) SetMemoryDynamicMaxMock(sessionID SessionRef, self VMRef, value int) (_err error) {
	return VMClassSetMemoryDynamicMaxMockedCallback(sessionID, self, value)
}
// Set the value of the memory_dynamic_max field
func (_class VMClass) SetMemoryDynamicMax(sessionID SessionRef, self VMRef, value int) (_err error) {
	if IsMock {
		return _class.SetMemoryDynamicMaxMock(sessionID, self, value)
	}	
	_method := "VM.set_memory_dynamic_max"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertIntToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}


func VMClassComputeMemoryOverheadMockDefault(sessionID SessionRef, vm VMRef) (_retval int, _err error) {
	log.Println("VM.ComputeMemoryOverhead not mocked")
	_err = errors.New("VM.ComputeMemoryOverhead not mocked")
	return
}

var VMClassComputeMemoryOverheadMockedCallback = VMClassComputeMemoryOverheadMockDefault

func (_class VMClass) ComputeMemoryOverheadMock(sessionID SessionRef, vm VMRef) (_retval int, _err error) {
	return VMClassComputeMemoryOverheadMockedCallback(sessionID, vm)
}
// Computes the virtualization memory overhead of a VM.
func (_class VMClass) ComputeMemoryOverhead(sessionID SessionRef, vm VMRef) (_retval int, _err error) {
	if IsMock {
		return _class.ComputeMemoryOverheadMock(sessionID, vm)
	}	
	_method := "VM.compute_memory_overhead"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_vmArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "vm"), vm)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _vmArg)
	if _err != nil {
		return
	}
	_retval, _err = convertIntToGo(_method + " -> ", _result.Value)
	return
}


func VMClassSetHaAlwaysRunMockDefault(sessionID SessionRef, self VMRef, value bool) (_err error) {
	log.Println("VM.SetHaAlwaysRun not mocked")
	_err = errors.New("VM.SetHaAlwaysRun not mocked")
	return
}

var VMClassSetHaAlwaysRunMockedCallback = VMClassSetHaAlwaysRunMockDefault

func (_class VMClass) SetHaAlwaysRunMock(sessionID SessionRef, self VMRef, value bool) (_err error) {
	return VMClassSetHaAlwaysRunMockedCallback(sessionID, self, value)
}
// Set the value of the ha_always_run
func (_class VMClass) SetHaAlwaysRun(sessionID SessionRef, self VMRef, value bool) (_err error) {
	if IsMock {
		return _class.SetHaAlwaysRunMock(sessionID, self, value)
	}	
	_method := "VM.set_ha_always_run"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertBoolToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}


func VMClassSetHaRestartPriorityMockDefault(sessionID SessionRef, self VMRef, value string) (_err error) {
	log.Println("VM.SetHaRestartPriority not mocked")
	_err = errors.New("VM.SetHaRestartPriority not mocked")
	return
}

var VMClassSetHaRestartPriorityMockedCallback = VMClassSetHaRestartPriorityMockDefault

func (_class VMClass) SetHaRestartPriorityMock(sessionID SessionRef, self VMRef, value string) (_err error) {
	return VMClassSetHaRestartPriorityMockedCallback(sessionID, self, value)
}
// Set the value of the ha_restart_priority field
func (_class VMClass) SetHaRestartPriority(sessionID SessionRef, self VMRef, value string) (_err error) {
	if IsMock {
		return _class.SetHaRestartPriorityMock(sessionID, self, value)
	}	
	_method := "VM.set_ha_restart_priority"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}


func VMClassAddToVCPUsParamsLiveMockDefault(sessionID SessionRef, self VMRef, key string, value string) (_err error) {
	log.Println("VM.AddToVCPUsParamsLive not mocked")
	_err = errors.New("VM.AddToVCPUsParamsLive not mocked")
	return
}

var VMClassAddToVCPUsParamsLiveMockedCallback = VMClassAddToVCPUsParamsLiveMockDefault

func (_class VMClass) AddToVCPUsParamsLiveMock(sessionID SessionRef, self VMRef, key string, value string) (_err error) {
	return VMClassAddToVCPUsParamsLiveMockedCallback(sessionID, self, key, value)
}
// Add the given key-value pair to VM.VCPUs_params, and apply that value on the running VM
func (_class VMClass) AddToVCPUsParamsLive(sessionID SessionRef, self VMRef, key string, value string) (_err error) {
	if IsMock {
		return _class.AddToVCPUsParamsLiveMock(sessionID, self, key, value)
	}	
	_method := "VM.add_to_VCPUs_params_live"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_keyArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "key"), key)
	if _err != nil {
		return
	}
	_valueArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _keyArg, _valueArg)
	return
}


func VMClassSetVCPUsNumberLiveMockDefault(sessionID SessionRef, self VMRef, nvcpu int) (_err error) {
	log.Println("VM.SetVCPUsNumberLive not mocked")
	_err = errors.New("VM.SetVCPUsNumberLive not mocked")
	return
}

var VMClassSetVCPUsNumberLiveMockedCallback = VMClassSetVCPUsNumberLiveMockDefault

func (_class VMClass) SetVCPUsNumberLiveMock(sessionID SessionRef, self VMRef, nvcpu int) (_err error) {
	return VMClassSetVCPUsNumberLiveMockedCallback(sessionID, self, nvcpu)
}
// Set the number of VCPUs for a running VM
//
// Errors:
//  OPERATION_NOT_ALLOWED - You attempted an operation that was not allowed.
//  LICENCE_RESTRICTION - This operation is not allowed because your license lacks a needed feature.  Please contact your support representative.
func (_class VMClass) SetVCPUsNumberLive(sessionID SessionRef, self VMRef, nvcpu int) (_err error) {
	if IsMock {
		return _class.SetVCPUsNumberLiveMock(sessionID, self, nvcpu)
	}	
	_method := "VM.set_VCPUs_number_live"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_nvcpuArg, _err := convertIntToXen(fmt.Sprintf("%s(%s)", _method, "nvcpu"), nvcpu)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _nvcpuArg)
	return
}


func VMClassPoolMigrateMockDefault(sessionID SessionRef, vm VMRef, host HostRef, options map[string]string) (_err error) {
	log.Println("VM.PoolMigrate not mocked")
	_err = errors.New("VM.PoolMigrate not mocked")
	return
}

var VMClassPoolMigrateMockedCallback = VMClassPoolMigrateMockDefault

func (_class VMClass) PoolMigrateMock(sessionID SessionRef, vm VMRef, host HostRef, options map[string]string) (_err error) {
	return VMClassPoolMigrateMockedCallback(sessionID, vm, host, options)
}
// Migrate a VM to another Host.
//
// Errors:
//  VM_BAD_POWER_STATE - You attempted an operation on a VM that was not in an appropriate power state at the time; for example, you attempted to start a VM that was already running.  The parameters returned are the VM's handle, and the expected and actual VM state at the time of the call.
//  OTHER_OPERATION_IN_PROGRESS - Another operation involving the object is currently in progress
//  VM_IS_TEMPLATE - The operation attempted is not valid for a template VM
//  OPERATION_NOT_ALLOWED - You attempted an operation that was not allowed.
//  VM_MIGRATE_FAILED - An error occurred during the migration process.
func (_class VMClass) PoolMigrate(sessionID SessionRef, vm VMRef, host HostRef, options map[string]string) (_err error) {
	if IsMock {
		return _class.PoolMigrateMock(sessionID, vm, host, options)
	}	
	_method := "VM.pool_migrate"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_vmArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "vm"), vm)
	if _err != nil {
		return
	}
	_hostArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "host"), host)
	if _err != nil {
		return
	}
	_optionsArg, _err := convertStringToStringMapToXen(fmt.Sprintf("%s(%s)", _method, "options"), options)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _vmArg, _hostArg, _optionsArg)
	return
}


func VMClassResumeOnMockDefault(sessionID SessionRef, vm VMRef, host HostRef, startPaused bool, force bool) (_err error) {
	log.Println("VM.ResumeOn not mocked")
	_err = errors.New("VM.ResumeOn not mocked")
	return
}

var VMClassResumeOnMockedCallback = VMClassResumeOnMockDefault

func (_class VMClass) ResumeOnMock(sessionID SessionRef, vm VMRef, host HostRef, startPaused bool, force bool) (_err error) {
	return VMClassResumeOnMockedCallback(sessionID, vm, host, startPaused, force)
}
// Awaken the specified VM and resume it on a particular Host.  This can only be called when the specified VM is in the Suspended state.
//
// Errors:
//  VM_BAD_POWER_STATE - You attempted an operation on a VM that was not in an appropriate power state at the time; for example, you attempted to start a VM that was already running.  The parameters returned are the VM's handle, and the expected and actual VM state at the time of the call.
//  OPERATION_NOT_ALLOWED - You attempted an operation that was not allowed.
//  VM_IS_TEMPLATE - The operation attempted is not valid for a template VM
func (_class VMClass) ResumeOn(sessionID SessionRef, vm VMRef, host HostRef, startPaused bool, force bool) (_err error) {
	if IsMock {
		return _class.ResumeOnMock(sessionID, vm, host, startPaused, force)
	}	
	_method := "VM.resume_on"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_vmArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "vm"), vm)
	if _err != nil {
		return
	}
	_hostArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "host"), host)
	if _err != nil {
		return
	}
	_startPausedArg, _err := convertBoolToXen(fmt.Sprintf("%s(%s)", _method, "start_paused"), startPaused)
	if _err != nil {
		return
	}
	_forceArg, _err := convertBoolToXen(fmt.Sprintf("%s(%s)", _method, "force"), force)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _vmArg, _hostArg, _startPausedArg, _forceArg)
	return
}


func VMClassResumeMockDefault(sessionID SessionRef, vm VMRef, startPaused bool, force bool) (_err error) {
	log.Println("VM.Resume not mocked")
	_err = errors.New("VM.Resume not mocked")
	return
}

var VMClassResumeMockedCallback = VMClassResumeMockDefault

func (_class VMClass) ResumeMock(sessionID SessionRef, vm VMRef, startPaused bool, force bool) (_err error) {
	return VMClassResumeMockedCallback(sessionID, vm, startPaused, force)
}
// Awaken the specified VM and resume it.  This can only be called when the specified VM is in the Suspended state.
//
// Errors:
//  VM_BAD_POWER_STATE - You attempted an operation on a VM that was not in an appropriate power state at the time; for example, you attempted to start a VM that was already running.  The parameters returned are the VM's handle, and the expected and actual VM state at the time of the call.
//  OPERATION_NOT_ALLOWED - You attempted an operation that was not allowed.
//  VM_IS_TEMPLATE - The operation attempted is not valid for a template VM
func (_class VMClass) Resume(sessionID SessionRef, vm VMRef, startPaused bool, force bool) (_err error) {
	if IsMock {
		return _class.ResumeMock(sessionID, vm, startPaused, force)
	}	
	_method := "VM.resume"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_vmArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "vm"), vm)
	if _err != nil {
		return
	}
	_startPausedArg, _err := convertBoolToXen(fmt.Sprintf("%s(%s)", _method, "start_paused"), startPaused)
	if _err != nil {
		return
	}
	_forceArg, _err := convertBoolToXen(fmt.Sprintf("%s(%s)", _method, "force"), force)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _vmArg, _startPausedArg, _forceArg)
	return
}


func VMClassSuspendMockDefault(sessionID SessionRef, vm VMRef) (_err error) {
	log.Println("VM.Suspend not mocked")
	_err = errors.New("VM.Suspend not mocked")
	return
}

var VMClassSuspendMockedCallback = VMClassSuspendMockDefault

func (_class VMClass) SuspendMock(sessionID SessionRef, vm VMRef) (_err error) {
	return VMClassSuspendMockedCallback(sessionID, vm)
}
// Suspend the specified VM to disk.  This can only be called when the specified VM is in the Running state.
//
// Errors:
//  VM_BAD_POWER_STATE - You attempted an operation on a VM that was not in an appropriate power state at the time; for example, you attempted to start a VM that was already running.  The parameters returned are the VM's handle, and the expected and actual VM state at the time of the call.
//  OTHER_OPERATION_IN_PROGRESS - Another operation involving the object is currently in progress
//  OPERATION_NOT_ALLOWED - You attempted an operation that was not allowed.
//  VM_IS_TEMPLATE - The operation attempted is not valid for a template VM
func (_class VMClass) Suspend(sessionID SessionRef, vm VMRef) (_err error) {
	if IsMock {
		return _class.SuspendMock(sessionID, vm)
	}	
	_method := "VM.suspend"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_vmArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "vm"), vm)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _vmArg)
	return
}


func VMClassHardRebootMockDefault(sessionID SessionRef, vm VMRef) (_err error) {
	log.Println("VM.HardReboot not mocked")
	_err = errors.New("VM.HardReboot not mocked")
	return
}

var VMClassHardRebootMockedCallback = VMClassHardRebootMockDefault

func (_class VMClass) HardRebootMock(sessionID SessionRef, vm VMRef) (_err error) {
	return VMClassHardRebootMockedCallback(sessionID, vm)
}
// Stop executing the specified VM without attempting a clean shutdown and immediately restart the VM.
//
// Errors:
//  VM_BAD_POWER_STATE - You attempted an operation on a VM that was not in an appropriate power state at the time; for example, you attempted to start a VM that was already running.  The parameters returned are the VM's handle, and the expected and actual VM state at the time of the call.
//  OTHER_OPERATION_IN_PROGRESS - Another operation involving the object is currently in progress
//  OPERATION_NOT_ALLOWED - You attempted an operation that was not allowed.
//  VM_IS_TEMPLATE - The operation attempted is not valid for a template VM
func (_class VMClass) HardReboot(sessionID SessionRef, vm VMRef) (_err error) {
	if IsMock {
		return _class.HardRebootMock(sessionID, vm)
	}	
	_method := "VM.hard_reboot"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_vmArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "vm"), vm)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _vmArg)
	return
}


func VMClassPowerStateResetMockDefault(sessionID SessionRef, vm VMRef) (_err error) {
	log.Println("VM.PowerStateReset not mocked")
	_err = errors.New("VM.PowerStateReset not mocked")
	return
}

var VMClassPowerStateResetMockedCallback = VMClassPowerStateResetMockDefault

func (_class VMClass) PowerStateResetMock(sessionID SessionRef, vm VMRef) (_err error) {
	return VMClassPowerStateResetMockedCallback(sessionID, vm)
}
// Reset the power-state of the VM to halted in the database only. (Used to recover from slave failures in pooling scenarios by resetting the power-states of VMs running on dead slaves to halted.) This is a potentially dangerous operation; use with care.
func (_class VMClass) PowerStateReset(sessionID SessionRef, vm VMRef) (_err error) {
	if IsMock {
		return _class.PowerStateResetMock(sessionID, vm)
	}	
	_method := "VM.power_state_reset"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_vmArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "vm"), vm)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _vmArg)
	return
}


func VMClassHardShutdownMockDefault(sessionID SessionRef, vm VMRef) (_err error) {
	log.Println("VM.HardShutdown not mocked")
	_err = errors.New("VM.HardShutdown not mocked")
	return
}

var VMClassHardShutdownMockedCallback = VMClassHardShutdownMockDefault

func (_class VMClass) HardShutdownMock(sessionID SessionRef, vm VMRef) (_err error) {
	return VMClassHardShutdownMockedCallback(sessionID, vm)
}
// Stop executing the specified VM without attempting a clean shutdown.
//
// Errors:
//  VM_BAD_POWER_STATE - You attempted an operation on a VM that was not in an appropriate power state at the time; for example, you attempted to start a VM that was already running.  The parameters returned are the VM's handle, and the expected and actual VM state at the time of the call.
//  OTHER_OPERATION_IN_PROGRESS - Another operation involving the object is currently in progress
//  OPERATION_NOT_ALLOWED - You attempted an operation that was not allowed.
//  VM_IS_TEMPLATE - The operation attempted is not valid for a template VM
func (_class VMClass) HardShutdown(sessionID SessionRef, vm VMRef) (_err error) {
	if IsMock {
		return _class.HardShutdownMock(sessionID, vm)
	}	
	_method := "VM.hard_shutdown"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_vmArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "vm"), vm)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _vmArg)
	return
}


func VMClassCleanRebootMockDefault(sessionID SessionRef, vm VMRef) (_err error) {
	log.Println("VM.CleanReboot not mocked")
	_err = errors.New("VM.CleanReboot not mocked")
	return
}

var VMClassCleanRebootMockedCallback = VMClassCleanRebootMockDefault

func (_class VMClass) CleanRebootMock(sessionID SessionRef, vm VMRef) (_err error) {
	return VMClassCleanRebootMockedCallback(sessionID, vm)
}
// Attempt to cleanly shutdown the specified VM (Note: this may not be supported---e.g. if a guest agent is not installed). This can only be called when the specified VM is in the Running state.
//
// Errors:
//  VM_BAD_POWER_STATE - You attempted an operation on a VM that was not in an appropriate power state at the time; for example, you attempted to start a VM that was already running.  The parameters returned are the VM's handle, and the expected and actual VM state at the time of the call.
//  OTHER_OPERATION_IN_PROGRESS - Another operation involving the object is currently in progress
//  OPERATION_NOT_ALLOWED - You attempted an operation that was not allowed.
//  VM_IS_TEMPLATE - The operation attempted is not valid for a template VM
func (_class VMClass) CleanReboot(sessionID SessionRef, vm VMRef) (_err error) {
	if IsMock {
		return _class.CleanRebootMock(sessionID, vm)
	}	
	_method := "VM.clean_reboot"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_vmArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "vm"), vm)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _vmArg)
	return
}


func VMClassShutdownMockDefault(sessionID SessionRef, vm VMRef) (_err error) {
	log.Println("VM.Shutdown not mocked")
	_err = errors.New("VM.Shutdown not mocked")
	return
}

var VMClassShutdownMockedCallback = VMClassShutdownMockDefault

func (_class VMClass) ShutdownMock(sessionID SessionRef, vm VMRef) (_err error) {
	return VMClassShutdownMockedCallback(sessionID, vm)
}
// Attempts to first clean shutdown a VM and if it should fail then perform a hard shutdown on it.
//
// Errors:
//  VM_BAD_POWER_STATE - You attempted an operation on a VM that was not in an appropriate power state at the time; for example, you attempted to start a VM that was already running.  The parameters returned are the VM's handle, and the expected and actual VM state at the time of the call.
//  OTHER_OPERATION_IN_PROGRESS - Another operation involving the object is currently in progress
//  OPERATION_NOT_ALLOWED - You attempted an operation that was not allowed.
//  VM_IS_TEMPLATE - The operation attempted is not valid for a template VM
func (_class VMClass) Shutdown(sessionID SessionRef, vm VMRef) (_err error) {
	if IsMock {
		return _class.ShutdownMock(sessionID, vm)
	}	
	_method := "VM.shutdown"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_vmArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "vm"), vm)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _vmArg)
	return
}


func VMClassCleanShutdownMockDefault(sessionID SessionRef, vm VMRef) (_err error) {
	log.Println("VM.CleanShutdown not mocked")
	_err = errors.New("VM.CleanShutdown not mocked")
	return
}

var VMClassCleanShutdownMockedCallback = VMClassCleanShutdownMockDefault

func (_class VMClass) CleanShutdownMock(sessionID SessionRef, vm VMRef) (_err error) {
	return VMClassCleanShutdownMockedCallback(sessionID, vm)
}
// Attempt to cleanly shutdown the specified VM. (Note: this may not be supported---e.g. if a guest agent is not installed). This can only be called when the specified VM is in the Running state.
//
// Errors:
//  VM_BAD_POWER_STATE - You attempted an operation on a VM that was not in an appropriate power state at the time; for example, you attempted to start a VM that was already running.  The parameters returned are the VM's handle, and the expected and actual VM state at the time of the call.
//  OTHER_OPERATION_IN_PROGRESS - Another operation involving the object is currently in progress
//  OPERATION_NOT_ALLOWED - You attempted an operation that was not allowed.
//  VM_IS_TEMPLATE - The operation attempted is not valid for a template VM
func (_class VMClass) CleanShutdown(sessionID SessionRef, vm VMRef) (_err error) {
	if IsMock {
		return _class.CleanShutdownMock(sessionID, vm)
	}	
	_method := "VM.clean_shutdown"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_vmArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "vm"), vm)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _vmArg)
	return
}


func VMClassUnpauseMockDefault(sessionID SessionRef, vm VMRef) (_err error) {
	log.Println("VM.Unpause not mocked")
	_err = errors.New("VM.Unpause not mocked")
	return
}

var VMClassUnpauseMockedCallback = VMClassUnpauseMockDefault

func (_class VMClass) UnpauseMock(sessionID SessionRef, vm VMRef) (_err error) {
	return VMClassUnpauseMockedCallback(sessionID, vm)
}
// Resume the specified VM. This can only be called when the specified VM is in the Paused state.
//
// Errors:
//  VM_BAD_POWER_STATE - You attempted an operation on a VM that was not in an appropriate power state at the time; for example, you attempted to start a VM that was already running.  The parameters returned are the VM's handle, and the expected and actual VM state at the time of the call.
//  OPERATION_NOT_ALLOWED - You attempted an operation that was not allowed.
//  VM_IS_TEMPLATE - The operation attempted is not valid for a template VM
func (_class VMClass) Unpause(sessionID SessionRef, vm VMRef) (_err error) {
	if IsMock {
		return _class.UnpauseMock(sessionID, vm)
	}	
	_method := "VM.unpause"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_vmArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "vm"), vm)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _vmArg)
	return
}


func VMClassPauseMockDefault(sessionID SessionRef, vm VMRef) (_err error) {
	log.Println("VM.Pause not mocked")
	_err = errors.New("VM.Pause not mocked")
	return
}

var VMClassPauseMockedCallback = VMClassPauseMockDefault

func (_class VMClass) PauseMock(sessionID SessionRef, vm VMRef) (_err error) {
	return VMClassPauseMockedCallback(sessionID, vm)
}
// Pause the specified VM. This can only be called when the specified VM is in the Running state.
//
// Errors:
//  VM_BAD_POWER_STATE - You attempted an operation on a VM that was not in an appropriate power state at the time; for example, you attempted to start a VM that was already running.  The parameters returned are the VM's handle, and the expected and actual VM state at the time of the call.
//  OTHER_OPERATION_IN_PROGRESS - Another operation involving the object is currently in progress
//  OPERATION_NOT_ALLOWED - You attempted an operation that was not allowed.
//  VM_IS_TEMPLATE - The operation attempted is not valid for a template VM
func (_class VMClass) Pause(sessionID SessionRef, vm VMRef) (_err error) {
	if IsMock {
		return _class.PauseMock(sessionID, vm)
	}	
	_method := "VM.pause"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_vmArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "vm"), vm)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _vmArg)
	return
}


func VMClassStartOnMockDefault(sessionID SessionRef, vm VMRef, host HostRef, startPaused bool, force bool) (_err error) {
	log.Println("VM.StartOn not mocked")
	_err = errors.New("VM.StartOn not mocked")
	return
}

var VMClassStartOnMockedCallback = VMClassStartOnMockDefault

func (_class VMClass) StartOnMock(sessionID SessionRef, vm VMRef, host HostRef, startPaused bool, force bool) (_err error) {
	return VMClassStartOnMockedCallback(sessionID, vm, host, startPaused, force)
}
// Start the specified VM on a particular host.  This function can only be called with the VM is in the Halted State.
//
// Errors:
//  VM_BAD_POWER_STATE - You attempted an operation on a VM that was not in an appropriate power state at the time; for example, you attempted to start a VM that was already running.  The parameters returned are the VM's handle, and the expected and actual VM state at the time of the call.
//  VM_IS_TEMPLATE - The operation attempted is not valid for a template VM
//  OTHER_OPERATION_IN_PROGRESS - Another operation involving the object is currently in progress
//  OPERATION_NOT_ALLOWED - You attempted an operation that was not allowed.
//  BOOTLOADER_FAILED - The bootloader returned an error
//  UNKNOWN_BOOTLOADER - The requested bootloader is unknown
func (_class VMClass) StartOn(sessionID SessionRef, vm VMRef, host HostRef, startPaused bool, force bool) (_err error) {
	if IsMock {
		return _class.StartOnMock(sessionID, vm, host, startPaused, force)
	}	
	_method := "VM.start_on"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_vmArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "vm"), vm)
	if _err != nil {
		return
	}
	_hostArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "host"), host)
	if _err != nil {
		return
	}
	_startPausedArg, _err := convertBoolToXen(fmt.Sprintf("%s(%s)", _method, "start_paused"), startPaused)
	if _err != nil {
		return
	}
	_forceArg, _err := convertBoolToXen(fmt.Sprintf("%s(%s)", _method, "force"), force)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _vmArg, _hostArg, _startPausedArg, _forceArg)
	return
}


func VMClassStartMockDefault(sessionID SessionRef, vm VMRef, startPaused bool, force bool) (_err error) {
	log.Println("VM.Start not mocked")
	_err = errors.New("VM.Start not mocked")
	return
}

var VMClassStartMockedCallback = VMClassStartMockDefault

func (_class VMClass) StartMock(sessionID SessionRef, vm VMRef, startPaused bool, force bool) (_err error) {
	return VMClassStartMockedCallback(sessionID, vm, startPaused, force)
}
// Start the specified VM.  This function can only be called with the VM is in the Halted State.
//
// Errors:
//  VM_BAD_POWER_STATE - You attempted an operation on a VM that was not in an appropriate power state at the time; for example, you attempted to start a VM that was already running.  The parameters returned are the VM's handle, and the expected and actual VM state at the time of the call.
//  VM_HVM_REQUIRED - HVM is required for this operation
//  VM_IS_TEMPLATE - The operation attempted is not valid for a template VM
//  OTHER_OPERATION_IN_PROGRESS - Another operation involving the object is currently in progress
//  OPERATION_NOT_ALLOWED - You attempted an operation that was not allowed.
//  BOOTLOADER_FAILED - The bootloader returned an error
//  UNKNOWN_BOOTLOADER - The requested bootloader is unknown
//  NO_HOSTS_AVAILABLE - There were no hosts available to complete the specified operation.
//  LICENCE_RESTRICTION - This operation is not allowed because your license lacks a needed feature.  Please contact your support representative.
func (_class VMClass) Start(sessionID SessionRef, vm VMRef, startPaused bool, force bool) (_err error) {
	if IsMock {
		return _class.StartMock(sessionID, vm, startPaused, force)
	}	
	_method := "VM.start"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_vmArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "vm"), vm)
	if _err != nil {
		return
	}
	_startPausedArg, _err := convertBoolToXen(fmt.Sprintf("%s(%s)", _method, "start_paused"), startPaused)
	if _err != nil {
		return
	}
	_forceArg, _err := convertBoolToXen(fmt.Sprintf("%s(%s)", _method, "force"), force)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _vmArg, _startPausedArg, _forceArg)
	return
}


func VMClassProvisionMockDefault(sessionID SessionRef, vm VMRef) (_err error) {
	log.Println("VM.Provision not mocked")
	_err = errors.New("VM.Provision not mocked")
	return
}

var VMClassProvisionMockedCallback = VMClassProvisionMockDefault

func (_class VMClass) ProvisionMock(sessionID SessionRef, vm VMRef) (_err error) {
	return VMClassProvisionMockedCallback(sessionID, vm)
}
// Inspects the disk configuration contained within the VM's other_config, creates VDIs and VBDs and then executes any applicable post-install script.
//
// Errors:
//  VM_BAD_POWER_STATE - You attempted an operation on a VM that was not in an appropriate power state at the time; for example, you attempted to start a VM that was already running.  The parameters returned are the VM's handle, and the expected and actual VM state at the time of the call.
//  SR_FULL - The SR is full. Requested new size exceeds the maximum size
//  OPERATION_NOT_ALLOWED - You attempted an operation that was not allowed.
//  LICENCE_RESTRICTION - This operation is not allowed because your license lacks a needed feature.  Please contact your support representative.
func (_class VMClass) Provision(sessionID SessionRef, vm VMRef) (_err error) {
	if IsMock {
		return _class.ProvisionMock(sessionID, vm)
	}	
	_method := "VM.provision"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_vmArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "vm"), vm)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _vmArg)
	return
}


func VMClassCheckpointMockDefault(sessionID SessionRef, vm VMRef, newName string) (_retval VMRef, _err error) {
	log.Println("VM.Checkpoint not mocked")
	_err = errors.New("VM.Checkpoint not mocked")
	return
}

var VMClassCheckpointMockedCallback = VMClassCheckpointMockDefault

func (_class VMClass) CheckpointMock(sessionID SessionRef, vm VMRef, newName string) (_retval VMRef, _err error) {
	return VMClassCheckpointMockedCallback(sessionID, vm, newName)
}
// Checkpoints the specified VM, making a new VM. Checkpoint automatically exploits the capabilities of the underlying storage repository in which the VM's disk images are stored (e.g. Copy on Write) and saves the memory image as well.
//
// Errors:
//  VM_BAD_POWER_STATE - You attempted an operation on a VM that was not in an appropriate power state at the time; for example, you attempted to start a VM that was already running.  The parameters returned are the VM's handle, and the expected and actual VM state at the time of the call.
//  SR_FULL - The SR is full. Requested new size exceeds the maximum size
//  OPERATION_NOT_ALLOWED - You attempted an operation that was not allowed.
//  VM_CHECKPOINT_SUSPEND_FAILED - An error occured while saving the memory image of the specified virtual machine
//  VM_CHECKPOINT_RESUME_FAILED - An error occured while restoring the memory image of the specified virtual machine
func (_class VMClass) Checkpoint(sessionID SessionRef, vm VMRef, newName string) (_retval VMRef, _err error) {
	if IsMock {
		return _class.CheckpointMock(sessionID, vm, newName)
	}	
	_method := "VM.checkpoint"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_vmArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "vm"), vm)
	if _err != nil {
		return
	}
	_newNameArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "new_name"), newName)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _vmArg, _newNameArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVMRefToGo(_method + " -> ", _result.Value)
	return
}


func VMClassRevertMockDefault(sessionID SessionRef, snapshot VMRef) (_err error) {
	log.Println("VM.Revert not mocked")
	_err = errors.New("VM.Revert not mocked")
	return
}

var VMClassRevertMockedCallback = VMClassRevertMockDefault

func (_class VMClass) RevertMock(sessionID SessionRef, snapshot VMRef) (_err error) {
	return VMClassRevertMockedCallback(sessionID, snapshot)
}
// Reverts the specified VM to a previous state.
//
// Errors:
//  VM_BAD_POWER_STATE - You attempted an operation on a VM that was not in an appropriate power state at the time; for example, you attempted to start a VM that was already running.  The parameters returned are the VM's handle, and the expected and actual VM state at the time of the call.
//  OPERATION_NOT_ALLOWED - You attempted an operation that was not allowed.
//  SR_FULL - The SR is full. Requested new size exceeds the maximum size
//  VM_REVERT_FAILED - An error occured while reverting the specified virtual machine to the specified snapshot
func (_class VMClass) Revert(sessionID SessionRef, snapshot VMRef) (_err error) {
	if IsMock {
		return _class.RevertMock(sessionID, snapshot)
	}	
	_method := "VM.revert"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_snapshotArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "snapshot"), snapshot)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _snapshotArg)
	return
}


func VMClassCopyMockDefault(sessionID SessionRef, vm VMRef, newName string, sr SRRef) (_retval VMRef, _err error) {
	log.Println("VM.Copy not mocked")
	_err = errors.New("VM.Copy not mocked")
	return
}

var VMClassCopyMockedCallback = VMClassCopyMockDefault

func (_class VMClass) CopyMock(sessionID SessionRef, vm VMRef, newName string, sr SRRef) (_retval VMRef, _err error) {
	return VMClassCopyMockedCallback(sessionID, vm, newName, sr)
}
// Copied the specified VM, making a new VM. Unlike clone, copy does not exploits the capabilities of the underlying storage repository in which the VM's disk images are stored. Instead, copy guarantees that the disk images of the newly created VM will be 'full disks' - i.e. not part of a CoW chain.  This function can only be called when the VM is in the Halted State.
//
// Errors:
//  VM_BAD_POWER_STATE - You attempted an operation on a VM that was not in an appropriate power state at the time; for example, you attempted to start a VM that was already running.  The parameters returned are the VM's handle, and the expected and actual VM state at the time of the call.
//  SR_FULL - The SR is full. Requested new size exceeds the maximum size
//  OPERATION_NOT_ALLOWED - You attempted an operation that was not allowed.
//  LICENCE_RESTRICTION - This operation is not allowed because your license lacks a needed feature.  Please contact your support representative.
func (_class VMClass) Copy(sessionID SessionRef, vm VMRef, newName string, sr SRRef) (_retval VMRef, _err error) {
	if IsMock {
		return _class.CopyMock(sessionID, vm, newName, sr)
	}	
	_method := "VM.copy"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_vmArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "vm"), vm)
	if _err != nil {
		return
	}
	_newNameArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "new_name"), newName)
	if _err != nil {
		return
	}
	_srArg, _err := convertSRRefToXen(fmt.Sprintf("%s(%s)", _method, "sr"), sr)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _vmArg, _newNameArg, _srArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVMRefToGo(_method + " -> ", _result.Value)
	return
}


func VMClassCloneMockDefault(sessionID SessionRef, vm VMRef, newName string) (_retval VMRef, _err error) {
	log.Println("VM.Clone not mocked")
	_err = errors.New("VM.Clone not mocked")
	return
}

var VMClassCloneMockedCallback = VMClassCloneMockDefault

func (_class VMClass) CloneMock(sessionID SessionRef, vm VMRef, newName string) (_retval VMRef, _err error) {
	return VMClassCloneMockedCallback(sessionID, vm, newName)
}
// Clones the specified VM, making a new VM. Clone automatically exploits the capabilities of the underlying storage repository in which the VM's disk images are stored (e.g. Copy on Write).   This function can only be called when the VM is in the Halted State.
//
// Errors:
//  VM_BAD_POWER_STATE - You attempted an operation on a VM that was not in an appropriate power state at the time; for example, you attempted to start a VM that was already running.  The parameters returned are the VM's handle, and the expected and actual VM state at the time of the call.
//  SR_FULL - The SR is full. Requested new size exceeds the maximum size
//  OPERATION_NOT_ALLOWED - You attempted an operation that was not allowed.
//  LICENCE_RESTRICTION - This operation is not allowed because your license lacks a needed feature.  Please contact your support representative.
func (_class VMClass) Clone(sessionID SessionRef, vm VMRef, newName string) (_retval VMRef, _err error) {
	if IsMock {
		return _class.CloneMock(sessionID, vm, newName)
	}	
	_method := "VM.clone"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_vmArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "vm"), vm)
	if _err != nil {
		return
	}
	_newNameArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "new_name"), newName)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _vmArg, _newNameArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVMRefToGo(_method + " -> ", _result.Value)
	return
}


func VMClassSnapshotWithQuiesceMockDefault(sessionID SessionRef, vm VMRef, newName string) (_retval VMRef, _err error) {
	log.Println("VM.SnapshotWithQuiesce not mocked")
	_err = errors.New("VM.SnapshotWithQuiesce not mocked")
	return
}

var VMClassSnapshotWithQuiesceMockedCallback = VMClassSnapshotWithQuiesceMockDefault

func (_class VMClass) SnapshotWithQuiesceMock(sessionID SessionRef, vm VMRef, newName string) (_retval VMRef, _err error) {
	return VMClassSnapshotWithQuiesceMockedCallback(sessionID, vm, newName)
}
// Snapshots the specified VM with quiesce, making a new VM. Snapshot automatically exploits the capabilities of the underlying storage repository in which the VM's disk images are stored (e.g. Copy on Write).
//
// Errors:
//  VM_BAD_POWER_STATE - You attempted an operation on a VM that was not in an appropriate power state at the time; for example, you attempted to start a VM that was already running.  The parameters returned are the VM's handle, and the expected and actual VM state at the time of the call.
//  SR_FULL - The SR is full. Requested new size exceeds the maximum size
//  OPERATION_NOT_ALLOWED - You attempted an operation that was not allowed.
//  VM_SNAPSHOT_WITH_QUIESCE_FAILED - The quiesced-snapshot operation failed for an unexpected reason
//  VM_SNAPSHOT_WITH_QUIESCE_TIMEOUT - The VSS plug-in has timed out
//  VM_SNAPSHOT_WITH_QUIESCE_PLUGIN_DEOS_NOT_RESPOND - The VSS plug-in cannot be contacted
//  VM_SNAPSHOT_WITH_QUIESCE_NOT_SUPPORTED - The VSS plug-in is not installed on this virtual machine
func (_class VMClass) SnapshotWithQuiesce(sessionID SessionRef, vm VMRef, newName string) (_retval VMRef, _err error) {
	if IsMock {
		return _class.SnapshotWithQuiesceMock(sessionID, vm, newName)
	}	
	_method := "VM.snapshot_with_quiesce"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_vmArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "vm"), vm)
	if _err != nil {
		return
	}
	_newNameArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "new_name"), newName)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _vmArg, _newNameArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVMRefToGo(_method + " -> ", _result.Value)
	return
}


func VMClassSnapshotMockDefault(sessionID SessionRef, vm VMRef, newName string) (_retval VMRef, _err error) {
	log.Println("VM.Snapshot not mocked")
	_err = errors.New("VM.Snapshot not mocked")
	return
}

var VMClassSnapshotMockedCallback = VMClassSnapshotMockDefault

func (_class VMClass) SnapshotMock(sessionID SessionRef, vm VMRef, newName string) (_retval VMRef, _err error) {
	return VMClassSnapshotMockedCallback(sessionID, vm, newName)
}
// Snapshots the specified VM, making a new VM. Snapshot automatically exploits the capabilities of the underlying storage repository in which the VM's disk images are stored (e.g. Copy on Write).
//
// Errors:
//  VM_BAD_POWER_STATE - You attempted an operation on a VM that was not in an appropriate power state at the time; for example, you attempted to start a VM that was already running.  The parameters returned are the VM's handle, and the expected and actual VM state at the time of the call.
//  SR_FULL - The SR is full. Requested new size exceeds the maximum size
//  OPERATION_NOT_ALLOWED - You attempted an operation that was not allowed.
func (_class VMClass) Snapshot(sessionID SessionRef, vm VMRef, newName string) (_retval VMRef, _err error) {
	if IsMock {
		return _class.SnapshotMock(sessionID, vm, newName)
	}	
	_method := "VM.snapshot"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_vmArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "vm"), vm)
	if _err != nil {
		return
	}
	_newNameArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "new_name"), newName)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _vmArg, _newNameArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVMRefToGo(_method + " -> ", _result.Value)
	return
}


func VMClassSetHardwarePlatformVersionMockDefault(sessionID SessionRef, self VMRef, value int) (_err error) {
	log.Println("VM.SetHardwarePlatformVersion not mocked")
	_err = errors.New("VM.SetHardwarePlatformVersion not mocked")
	return
}

var VMClassSetHardwarePlatformVersionMockedCallback = VMClassSetHardwarePlatformVersionMockDefault

func (_class VMClass) SetHardwarePlatformVersionMock(sessionID SessionRef, self VMRef, value int) (_err error) {
	return VMClassSetHardwarePlatformVersionMockedCallback(sessionID, self, value)
}
// Set the hardware_platform_version field of the given VM.
func (_class VMClass) SetHardwarePlatformVersion(sessionID SessionRef, self VMRef, value int) (_err error) {
	if IsMock {
		return _class.SetHardwarePlatformVersionMock(sessionID, self, value)
	}	
	_method := "VM.set_hardware_platform_version"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertIntToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}


func VMClassSetSuspendSRMockDefault(sessionID SessionRef, self VMRef, value SRRef) (_err error) {
	log.Println("VM.SetSuspendSR not mocked")
	_err = errors.New("VM.SetSuspendSR not mocked")
	return
}

var VMClassSetSuspendSRMockedCallback = VMClassSetSuspendSRMockDefault

func (_class VMClass) SetSuspendSRMock(sessionID SessionRef, self VMRef, value SRRef) (_err error) {
	return VMClassSetSuspendSRMockedCallback(sessionID, self, value)
}
// Set the suspend_SR field of the given VM.
func (_class VMClass) SetSuspendSR(sessionID SessionRef, self VMRef, value SRRef) (_err error) {
	if IsMock {
		return _class.SetSuspendSRMock(sessionID, self, value)
	}	
	_method := "VM.set_suspend_SR"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertSRRefToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}


func VMClassRemoveFromBlockedOperationsMockDefault(sessionID SessionRef, self VMRef, key VMOperations) (_err error) {
	log.Println("VM.RemoveFromBlockedOperations not mocked")
	_err = errors.New("VM.RemoveFromBlockedOperations not mocked")
	return
}

var VMClassRemoveFromBlockedOperationsMockedCallback = VMClassRemoveFromBlockedOperationsMockDefault

func (_class VMClass) RemoveFromBlockedOperationsMock(sessionID SessionRef, self VMRef, key VMOperations) (_err error) {
	return VMClassRemoveFromBlockedOperationsMockedCallback(sessionID, self, key)
}
// Remove the given key and its corresponding value from the blocked_operations field of the given VM.  If the key is not in that Map, then do nothing.
func (_class VMClass) RemoveFromBlockedOperations(sessionID SessionRef, self VMRef, key VMOperations) (_err error) {
	if IsMock {
		return _class.RemoveFromBlockedOperationsMock(sessionID, self, key)
	}	
	_method := "VM.remove_from_blocked_operations"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_keyArg, _err := convertEnumVMOperationsToXen(fmt.Sprintf("%s(%s)", _method, "key"), key)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _keyArg)
	return
}


func VMClassAddToBlockedOperationsMockDefault(sessionID SessionRef, self VMRef, key VMOperations, value string) (_err error) {
	log.Println("VM.AddToBlockedOperations not mocked")
	_err = errors.New("VM.AddToBlockedOperations not mocked")
	return
}

var VMClassAddToBlockedOperationsMockedCallback = VMClassAddToBlockedOperationsMockDefault

func (_class VMClass) AddToBlockedOperationsMock(sessionID SessionRef, self VMRef, key VMOperations, value string) (_err error) {
	return VMClassAddToBlockedOperationsMockedCallback(sessionID, self, key, value)
}
// Add the given key-value pair to the blocked_operations field of the given VM.
func (_class VMClass) AddToBlockedOperations(sessionID SessionRef, self VMRef, key VMOperations, value string) (_err error) {
	if IsMock {
		return _class.AddToBlockedOperationsMock(sessionID, self, key, value)
	}	
	_method := "VM.add_to_blocked_operations"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_keyArg, _err := convertEnumVMOperationsToXen(fmt.Sprintf("%s(%s)", _method, "key"), key)
	if _err != nil {
		return
	}
	_valueArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _keyArg, _valueArg)
	return
}


func VMClassSetBlockedOperationsMockDefault(sessionID SessionRef, self VMRef, value map[VMOperations]string) (_err error) {
	log.Println("VM.SetBlockedOperations not mocked")
	_err = errors.New("VM.SetBlockedOperations not mocked")
	return
}

var VMClassSetBlockedOperationsMockedCallback = VMClassSetBlockedOperationsMockDefault

func (_class VMClass) SetBlockedOperationsMock(sessionID SessionRef, self VMRef, value map[VMOperations]string) (_err error) {
	return VMClassSetBlockedOperationsMockedCallback(sessionID, self, value)
}
// Set the blocked_operations field of the given VM.
func (_class VMClass) SetBlockedOperations(sessionID SessionRef, self VMRef, value map[VMOperations]string) (_err error) {
	if IsMock {
		return _class.SetBlockedOperationsMock(sessionID, self, value)
	}	
	_method := "VM.set_blocked_operations"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertEnumVMOperationsToStringMapToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}


func VMClassRemoveTagsMockDefault(sessionID SessionRef, self VMRef, value string) (_err error) {
	log.Println("VM.RemoveTags not mocked")
	_err = errors.New("VM.RemoveTags not mocked")
	return
}

var VMClassRemoveTagsMockedCallback = VMClassRemoveTagsMockDefault

func (_class VMClass) RemoveTagsMock(sessionID SessionRef, self VMRef, value string) (_err error) {
	return VMClassRemoveTagsMockedCallback(sessionID, self, value)
}
// Remove the given value from the tags field of the given VM.  If the value is not in that Set, then do nothing.
func (_class VMClass) RemoveTags(sessionID SessionRef, self VMRef, value string) (_err error) {
	if IsMock {
		return _class.RemoveTagsMock(sessionID, self, value)
	}	
	_method := "VM.remove_tags"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}


func VMClassAddTagsMockDefault(sessionID SessionRef, self VMRef, value string) (_err error) {
	log.Println("VM.AddTags not mocked")
	_err = errors.New("VM.AddTags not mocked")
	return
}

var VMClassAddTagsMockedCallback = VMClassAddTagsMockDefault

func (_class VMClass) AddTagsMock(sessionID SessionRef, self VMRef, value string) (_err error) {
	return VMClassAddTagsMockedCallback(sessionID, self, value)
}
// Add the given value to the tags field of the given VM.  If the value is already in that Set, then do nothing.
func (_class VMClass) AddTags(sessionID SessionRef, self VMRef, value string) (_err error) {
	if IsMock {
		return _class.AddTagsMock(sessionID, self, value)
	}	
	_method := "VM.add_tags"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}


func VMClassSetTagsMockDefault(sessionID SessionRef, self VMRef, value []string) (_err error) {
	log.Println("VM.SetTags not mocked")
	_err = errors.New("VM.SetTags not mocked")
	return
}

var VMClassSetTagsMockedCallback = VMClassSetTagsMockDefault

func (_class VMClass) SetTagsMock(sessionID SessionRef, self VMRef, value []string) (_err error) {
	return VMClassSetTagsMockedCallback(sessionID, self, value)
}
// Set the tags field of the given VM.
func (_class VMClass) SetTags(sessionID SessionRef, self VMRef, value []string) (_err error) {
	if IsMock {
		return _class.SetTagsMock(sessionID, self, value)
	}	
	_method := "VM.set_tags"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertStringSetToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}


func VMClassRemoveFromXenstoreDataMockDefault(sessionID SessionRef, self VMRef, key string) (_err error) {
	log.Println("VM.RemoveFromXenstoreData not mocked")
	_err = errors.New("VM.RemoveFromXenstoreData not mocked")
	return
}

var VMClassRemoveFromXenstoreDataMockedCallback = VMClassRemoveFromXenstoreDataMockDefault

func (_class VMClass) RemoveFromXenstoreDataMock(sessionID SessionRef, self VMRef, key string) (_err error) {
	return VMClassRemoveFromXenstoreDataMockedCallback(sessionID, self, key)
}
// Remove the given key and its corresponding value from the xenstore_data field of the given VM.  If the key is not in that Map, then do nothing.
func (_class VMClass) RemoveFromXenstoreData(sessionID SessionRef, self VMRef, key string) (_err error) {
	if IsMock {
		return _class.RemoveFromXenstoreDataMock(sessionID, self, key)
	}	
	_method := "VM.remove_from_xenstore_data"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_keyArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "key"), key)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _keyArg)
	return
}


func VMClassAddToXenstoreDataMockDefault(sessionID SessionRef, self VMRef, key string, value string) (_err error) {
	log.Println("VM.AddToXenstoreData not mocked")
	_err = errors.New("VM.AddToXenstoreData not mocked")
	return
}

var VMClassAddToXenstoreDataMockedCallback = VMClassAddToXenstoreDataMockDefault

func (_class VMClass) AddToXenstoreDataMock(sessionID SessionRef, self VMRef, key string, value string) (_err error) {
	return VMClassAddToXenstoreDataMockedCallback(sessionID, self, key, value)
}
// Add the given key-value pair to the xenstore_data field of the given VM.
func (_class VMClass) AddToXenstoreData(sessionID SessionRef, self VMRef, key string, value string) (_err error) {
	if IsMock {
		return _class.AddToXenstoreDataMock(sessionID, self, key, value)
	}	
	_method := "VM.add_to_xenstore_data"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_keyArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "key"), key)
	if _err != nil {
		return
	}
	_valueArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _keyArg, _valueArg)
	return
}


func VMClassSetXenstoreDataMockDefault(sessionID SessionRef, self VMRef, value map[string]string) (_err error) {
	log.Println("VM.SetXenstoreData not mocked")
	_err = errors.New("VM.SetXenstoreData not mocked")
	return
}

var VMClassSetXenstoreDataMockedCallback = VMClassSetXenstoreDataMockDefault

func (_class VMClass) SetXenstoreDataMock(sessionID SessionRef, self VMRef, value map[string]string) (_err error) {
	return VMClassSetXenstoreDataMockedCallback(sessionID, self, value)
}
// Set the xenstore_data field of the given VM.
func (_class VMClass) SetXenstoreData(sessionID SessionRef, self VMRef, value map[string]string) (_err error) {
	if IsMock {
		return _class.SetXenstoreDataMock(sessionID, self, value)
	}	
	_method := "VM.set_xenstore_data"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertStringToStringMapToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}


func VMClassSetRecommendationsMockDefault(sessionID SessionRef, self VMRef, value string) (_err error) {
	log.Println("VM.SetRecommendations not mocked")
	_err = errors.New("VM.SetRecommendations not mocked")
	return
}

var VMClassSetRecommendationsMockedCallback = VMClassSetRecommendationsMockDefault

func (_class VMClass) SetRecommendationsMock(sessionID SessionRef, self VMRef, value string) (_err error) {
	return VMClassSetRecommendationsMockedCallback(sessionID, self, value)
}
// Set the recommendations field of the given VM.
func (_class VMClass) SetRecommendations(sessionID SessionRef, self VMRef, value string) (_err error) {
	if IsMock {
		return _class.SetRecommendationsMock(sessionID, self, value)
	}	
	_method := "VM.set_recommendations"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}


func VMClassRemoveFromOtherConfigMockDefault(sessionID SessionRef, self VMRef, key string) (_err error) {
	log.Println("VM.RemoveFromOtherConfig not mocked")
	_err = errors.New("VM.RemoveFromOtherConfig not mocked")
	return
}

var VMClassRemoveFromOtherConfigMockedCallback = VMClassRemoveFromOtherConfigMockDefault

func (_class VMClass) RemoveFromOtherConfigMock(sessionID SessionRef, self VMRef, key string) (_err error) {
	return VMClassRemoveFromOtherConfigMockedCallback(sessionID, self, key)
}
// Remove the given key and its corresponding value from the other_config field of the given VM.  If the key is not in that Map, then do nothing.
func (_class VMClass) RemoveFromOtherConfig(sessionID SessionRef, self VMRef, key string) (_err error) {
	if IsMock {
		return _class.RemoveFromOtherConfigMock(sessionID, self, key)
	}	
	_method := "VM.remove_from_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_keyArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "key"), key)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _keyArg)
	return
}


func VMClassAddToOtherConfigMockDefault(sessionID SessionRef, self VMRef, key string, value string) (_err error) {
	log.Println("VM.AddToOtherConfig not mocked")
	_err = errors.New("VM.AddToOtherConfig not mocked")
	return
}

var VMClassAddToOtherConfigMockedCallback = VMClassAddToOtherConfigMockDefault

func (_class VMClass) AddToOtherConfigMock(sessionID SessionRef, self VMRef, key string, value string) (_err error) {
	return VMClassAddToOtherConfigMockedCallback(sessionID, self, key, value)
}
// Add the given key-value pair to the other_config field of the given VM.
func (_class VMClass) AddToOtherConfig(sessionID SessionRef, self VMRef, key string, value string) (_err error) {
	if IsMock {
		return _class.AddToOtherConfigMock(sessionID, self, key, value)
	}	
	_method := "VM.add_to_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_keyArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "key"), key)
	if _err != nil {
		return
	}
	_valueArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _keyArg, _valueArg)
	return
}


func VMClassSetOtherConfigMockDefault(sessionID SessionRef, self VMRef, value map[string]string) (_err error) {
	log.Println("VM.SetOtherConfig not mocked")
	_err = errors.New("VM.SetOtherConfig not mocked")
	return
}

var VMClassSetOtherConfigMockedCallback = VMClassSetOtherConfigMockDefault

func (_class VMClass) SetOtherConfigMock(sessionID SessionRef, self VMRef, value map[string]string) (_err error) {
	return VMClassSetOtherConfigMockedCallback(sessionID, self, value)
}
// Set the other_config field of the given VM.
func (_class VMClass) SetOtherConfig(sessionID SessionRef, self VMRef, value map[string]string) (_err error) {
	if IsMock {
		return _class.SetOtherConfigMock(sessionID, self, value)
	}	
	_method := "VM.set_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertStringToStringMapToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}


func VMClassSetPCIBusMockDefault(sessionID SessionRef, self VMRef, value string) (_err error) {
	log.Println("VM.SetPCIBus not mocked")
	_err = errors.New("VM.SetPCIBus not mocked")
	return
}

var VMClassSetPCIBusMockedCallback = VMClassSetPCIBusMockDefault

func (_class VMClass) SetPCIBusMock(sessionID SessionRef, self VMRef, value string) (_err error) {
	return VMClassSetPCIBusMockedCallback(sessionID, self, value)
}
// Set the PCI_bus field of the given VM.
func (_class VMClass) SetPCIBus(sessionID SessionRef, self VMRef, value string) (_err error) {
	if IsMock {
		return _class.SetPCIBusMock(sessionID, self, value)
	}	
	_method := "VM.set_PCI_bus"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}


func VMClassRemoveFromPlatformMockDefault(sessionID SessionRef, self VMRef, key string) (_err error) {
	log.Println("VM.RemoveFromPlatform not mocked")
	_err = errors.New("VM.RemoveFromPlatform not mocked")
	return
}

var VMClassRemoveFromPlatformMockedCallback = VMClassRemoveFromPlatformMockDefault

func (_class VMClass) RemoveFromPlatformMock(sessionID SessionRef, self VMRef, key string) (_err error) {
	return VMClassRemoveFromPlatformMockedCallback(sessionID, self, key)
}
// Remove the given key and its corresponding value from the platform field of the given VM.  If the key is not in that Map, then do nothing.
func (_class VMClass) RemoveFromPlatform(sessionID SessionRef, self VMRef, key string) (_err error) {
	if IsMock {
		return _class.RemoveFromPlatformMock(sessionID, self, key)
	}	
	_method := "VM.remove_from_platform"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_keyArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "key"), key)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _keyArg)
	return
}


func VMClassAddToPlatformMockDefault(sessionID SessionRef, self VMRef, key string, value string) (_err error) {
	log.Println("VM.AddToPlatform not mocked")
	_err = errors.New("VM.AddToPlatform not mocked")
	return
}

var VMClassAddToPlatformMockedCallback = VMClassAddToPlatformMockDefault

func (_class VMClass) AddToPlatformMock(sessionID SessionRef, self VMRef, key string, value string) (_err error) {
	return VMClassAddToPlatformMockedCallback(sessionID, self, key, value)
}
// Add the given key-value pair to the platform field of the given VM.
func (_class VMClass) AddToPlatform(sessionID SessionRef, self VMRef, key string, value string) (_err error) {
	if IsMock {
		return _class.AddToPlatformMock(sessionID, self, key, value)
	}	
	_method := "VM.add_to_platform"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_keyArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "key"), key)
	if _err != nil {
		return
	}
	_valueArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _keyArg, _valueArg)
	return
}


func VMClassSetPlatformMockDefault(sessionID SessionRef, self VMRef, value map[string]string) (_err error) {
	log.Println("VM.SetPlatform not mocked")
	_err = errors.New("VM.SetPlatform not mocked")
	return
}

var VMClassSetPlatformMockedCallback = VMClassSetPlatformMockDefault

func (_class VMClass) SetPlatformMock(sessionID SessionRef, self VMRef, value map[string]string) (_err error) {
	return VMClassSetPlatformMockedCallback(sessionID, self, value)
}
// Set the platform field of the given VM.
func (_class VMClass) SetPlatform(sessionID SessionRef, self VMRef, value map[string]string) (_err error) {
	if IsMock {
		return _class.SetPlatformMock(sessionID, self, value)
	}	
	_method := "VM.set_platform"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertStringToStringMapToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}


func VMClassRemoveFromHVMBootParamsMockDefault(sessionID SessionRef, self VMRef, key string) (_err error) {
	log.Println("VM.RemoveFromHVMBootParams not mocked")
	_err = errors.New("VM.RemoveFromHVMBootParams not mocked")
	return
}

var VMClassRemoveFromHVMBootParamsMockedCallback = VMClassRemoveFromHVMBootParamsMockDefault

func (_class VMClass) RemoveFromHVMBootParamsMock(sessionID SessionRef, self VMRef, key string) (_err error) {
	return VMClassRemoveFromHVMBootParamsMockedCallback(sessionID, self, key)
}
// Remove the given key and its corresponding value from the HVM/boot_params field of the given VM.  If the key is not in that Map, then do nothing.
func (_class VMClass) RemoveFromHVMBootParams(sessionID SessionRef, self VMRef, key string) (_err error) {
	if IsMock {
		return _class.RemoveFromHVMBootParamsMock(sessionID, self, key)
	}	
	_method := "VM.remove_from_HVM_boot_params"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_keyArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "key"), key)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _keyArg)
	return
}


func VMClassAddToHVMBootParamsMockDefault(sessionID SessionRef, self VMRef, key string, value string) (_err error) {
	log.Println("VM.AddToHVMBootParams not mocked")
	_err = errors.New("VM.AddToHVMBootParams not mocked")
	return
}

var VMClassAddToHVMBootParamsMockedCallback = VMClassAddToHVMBootParamsMockDefault

func (_class VMClass) AddToHVMBootParamsMock(sessionID SessionRef, self VMRef, key string, value string) (_err error) {
	return VMClassAddToHVMBootParamsMockedCallback(sessionID, self, key, value)
}
// Add the given key-value pair to the HVM/boot_params field of the given VM.
func (_class VMClass) AddToHVMBootParams(sessionID SessionRef, self VMRef, key string, value string) (_err error) {
	if IsMock {
		return _class.AddToHVMBootParamsMock(sessionID, self, key, value)
	}	
	_method := "VM.add_to_HVM_boot_params"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_keyArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "key"), key)
	if _err != nil {
		return
	}
	_valueArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _keyArg, _valueArg)
	return
}


func VMClassSetHVMBootParamsMockDefault(sessionID SessionRef, self VMRef, value map[string]string) (_err error) {
	log.Println("VM.SetHVMBootParams not mocked")
	_err = errors.New("VM.SetHVMBootParams not mocked")
	return
}

var VMClassSetHVMBootParamsMockedCallback = VMClassSetHVMBootParamsMockDefault

func (_class VMClass) SetHVMBootParamsMock(sessionID SessionRef, self VMRef, value map[string]string) (_err error) {
	return VMClassSetHVMBootParamsMockedCallback(sessionID, self, value)
}
// Set the HVM/boot_params field of the given VM.
func (_class VMClass) SetHVMBootParams(sessionID SessionRef, self VMRef, value map[string]string) (_err error) {
	if IsMock {
		return _class.SetHVMBootParamsMock(sessionID, self, value)
	}	
	_method := "VM.set_HVM_boot_params"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertStringToStringMapToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}


func VMClassSetHVMBootPolicyMockDefault(sessionID SessionRef, self VMRef, value string) (_err error) {
	log.Println("VM.SetHVMBootPolicy not mocked")
	_err = errors.New("VM.SetHVMBootPolicy not mocked")
	return
}

var VMClassSetHVMBootPolicyMockedCallback = VMClassSetHVMBootPolicyMockDefault

func (_class VMClass) SetHVMBootPolicyMock(sessionID SessionRef, self VMRef, value string) (_err error) {
	return VMClassSetHVMBootPolicyMockedCallback(sessionID, self, value)
}
// Set the HVM/boot_policy field of the given VM.
func (_class VMClass) SetHVMBootPolicy(sessionID SessionRef, self VMRef, value string) (_err error) {
	if IsMock {
		return _class.SetHVMBootPolicyMock(sessionID, self, value)
	}	
	_method := "VM.set_HVM_boot_policy"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}


func VMClassSetPVLegacyArgsMockDefault(sessionID SessionRef, self VMRef, value string) (_err error) {
	log.Println("VM.SetPVLegacyArgs not mocked")
	_err = errors.New("VM.SetPVLegacyArgs not mocked")
	return
}

var VMClassSetPVLegacyArgsMockedCallback = VMClassSetPVLegacyArgsMockDefault

func (_class VMClass) SetPVLegacyArgsMock(sessionID SessionRef, self VMRef, value string) (_err error) {
	return VMClassSetPVLegacyArgsMockedCallback(sessionID, self, value)
}
// Set the PV/legacy_args field of the given VM.
func (_class VMClass) SetPVLegacyArgs(sessionID SessionRef, self VMRef, value string) (_err error) {
	if IsMock {
		return _class.SetPVLegacyArgsMock(sessionID, self, value)
	}	
	_method := "VM.set_PV_legacy_args"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}


func VMClassSetPVBootloaderArgsMockDefault(sessionID SessionRef, self VMRef, value string) (_err error) {
	log.Println("VM.SetPVBootloaderArgs not mocked")
	_err = errors.New("VM.SetPVBootloaderArgs not mocked")
	return
}

var VMClassSetPVBootloaderArgsMockedCallback = VMClassSetPVBootloaderArgsMockDefault

func (_class VMClass) SetPVBootloaderArgsMock(sessionID SessionRef, self VMRef, value string) (_err error) {
	return VMClassSetPVBootloaderArgsMockedCallback(sessionID, self, value)
}
// Set the PV/bootloader_args field of the given VM.
func (_class VMClass) SetPVBootloaderArgs(sessionID SessionRef, self VMRef, value string) (_err error) {
	if IsMock {
		return _class.SetPVBootloaderArgsMock(sessionID, self, value)
	}	
	_method := "VM.set_PV_bootloader_args"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}


func VMClassSetPVArgsMockDefault(sessionID SessionRef, self VMRef, value string) (_err error) {
	log.Println("VM.SetPVArgs not mocked")
	_err = errors.New("VM.SetPVArgs not mocked")
	return
}

var VMClassSetPVArgsMockedCallback = VMClassSetPVArgsMockDefault

func (_class VMClass) SetPVArgsMock(sessionID SessionRef, self VMRef, value string) (_err error) {
	return VMClassSetPVArgsMockedCallback(sessionID, self, value)
}
// Set the PV/args field of the given VM.
func (_class VMClass) SetPVArgs(sessionID SessionRef, self VMRef, value string) (_err error) {
	if IsMock {
		return _class.SetPVArgsMock(sessionID, self, value)
	}	
	_method := "VM.set_PV_args"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}


func VMClassSetPVRamdiskMockDefault(sessionID SessionRef, self VMRef, value string) (_err error) {
	log.Println("VM.SetPVRamdisk not mocked")
	_err = errors.New("VM.SetPVRamdisk not mocked")
	return
}

var VMClassSetPVRamdiskMockedCallback = VMClassSetPVRamdiskMockDefault

func (_class VMClass) SetPVRamdiskMock(sessionID SessionRef, self VMRef, value string) (_err error) {
	return VMClassSetPVRamdiskMockedCallback(sessionID, self, value)
}
// Set the PV/ramdisk field of the given VM.
func (_class VMClass) SetPVRamdisk(sessionID SessionRef, self VMRef, value string) (_err error) {
	if IsMock {
		return _class.SetPVRamdiskMock(sessionID, self, value)
	}	
	_method := "VM.set_PV_ramdisk"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}


func VMClassSetPVKernelMockDefault(sessionID SessionRef, self VMRef, value string) (_err error) {
	log.Println("VM.SetPVKernel not mocked")
	_err = errors.New("VM.SetPVKernel not mocked")
	return
}

var VMClassSetPVKernelMockedCallback = VMClassSetPVKernelMockDefault

func (_class VMClass) SetPVKernelMock(sessionID SessionRef, self VMRef, value string) (_err error) {
	return VMClassSetPVKernelMockedCallback(sessionID, self, value)
}
// Set the PV/kernel field of the given VM.
func (_class VMClass) SetPVKernel(sessionID SessionRef, self VMRef, value string) (_err error) {
	if IsMock {
		return _class.SetPVKernelMock(sessionID, self, value)
	}	
	_method := "VM.set_PV_kernel"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}


func VMClassSetPVBootloaderMockDefault(sessionID SessionRef, self VMRef, value string) (_err error) {
	log.Println("VM.SetPVBootloader not mocked")
	_err = errors.New("VM.SetPVBootloader not mocked")
	return
}

var VMClassSetPVBootloaderMockedCallback = VMClassSetPVBootloaderMockDefault

func (_class VMClass) SetPVBootloaderMock(sessionID SessionRef, self VMRef, value string) (_err error) {
	return VMClassSetPVBootloaderMockedCallback(sessionID, self, value)
}
// Set the PV/bootloader field of the given VM.
func (_class VMClass) SetPVBootloader(sessionID SessionRef, self VMRef, value string) (_err error) {
	if IsMock {
		return _class.SetPVBootloaderMock(sessionID, self, value)
	}	
	_method := "VM.set_PV_bootloader"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}


func VMClassSetActionsAfterRebootMockDefault(sessionID SessionRef, self VMRef, value OnNormalExit) (_err error) {
	log.Println("VM.SetActionsAfterReboot not mocked")
	_err = errors.New("VM.SetActionsAfterReboot not mocked")
	return
}

var VMClassSetActionsAfterRebootMockedCallback = VMClassSetActionsAfterRebootMockDefault

func (_class VMClass) SetActionsAfterRebootMock(sessionID SessionRef, self VMRef, value OnNormalExit) (_err error) {
	return VMClassSetActionsAfterRebootMockedCallback(sessionID, self, value)
}
// Set the actions/after_reboot field of the given VM.
func (_class VMClass) SetActionsAfterReboot(sessionID SessionRef, self VMRef, value OnNormalExit) (_err error) {
	if IsMock {
		return _class.SetActionsAfterRebootMock(sessionID, self, value)
	}	
	_method := "VM.set_actions_after_reboot"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertEnumOnNormalExitToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}


func VMClassSetActionsAfterShutdownMockDefault(sessionID SessionRef, self VMRef, value OnNormalExit) (_err error) {
	log.Println("VM.SetActionsAfterShutdown not mocked")
	_err = errors.New("VM.SetActionsAfterShutdown not mocked")
	return
}

var VMClassSetActionsAfterShutdownMockedCallback = VMClassSetActionsAfterShutdownMockDefault

func (_class VMClass) SetActionsAfterShutdownMock(sessionID SessionRef, self VMRef, value OnNormalExit) (_err error) {
	return VMClassSetActionsAfterShutdownMockedCallback(sessionID, self, value)
}
// Set the actions/after_shutdown field of the given VM.
func (_class VMClass) SetActionsAfterShutdown(sessionID SessionRef, self VMRef, value OnNormalExit) (_err error) {
	if IsMock {
		return _class.SetActionsAfterShutdownMock(sessionID, self, value)
	}	
	_method := "VM.set_actions_after_shutdown"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertEnumOnNormalExitToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}


func VMClassRemoveFromVCPUsParamsMockDefault(sessionID SessionRef, self VMRef, key string) (_err error) {
	log.Println("VM.RemoveFromVCPUsParams not mocked")
	_err = errors.New("VM.RemoveFromVCPUsParams not mocked")
	return
}

var VMClassRemoveFromVCPUsParamsMockedCallback = VMClassRemoveFromVCPUsParamsMockDefault

func (_class VMClass) RemoveFromVCPUsParamsMock(sessionID SessionRef, self VMRef, key string) (_err error) {
	return VMClassRemoveFromVCPUsParamsMockedCallback(sessionID, self, key)
}
// Remove the given key and its corresponding value from the VCPUs/params field of the given VM.  If the key is not in that Map, then do nothing.
func (_class VMClass) RemoveFromVCPUsParams(sessionID SessionRef, self VMRef, key string) (_err error) {
	if IsMock {
		return _class.RemoveFromVCPUsParamsMock(sessionID, self, key)
	}	
	_method := "VM.remove_from_VCPUs_params"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_keyArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "key"), key)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _keyArg)
	return
}


func VMClassAddToVCPUsParamsMockDefault(sessionID SessionRef, self VMRef, key string, value string) (_err error) {
	log.Println("VM.AddToVCPUsParams not mocked")
	_err = errors.New("VM.AddToVCPUsParams not mocked")
	return
}

var VMClassAddToVCPUsParamsMockedCallback = VMClassAddToVCPUsParamsMockDefault

func (_class VMClass) AddToVCPUsParamsMock(sessionID SessionRef, self VMRef, key string, value string) (_err error) {
	return VMClassAddToVCPUsParamsMockedCallback(sessionID, self, key, value)
}
// Add the given key-value pair to the VCPUs/params field of the given VM.
func (_class VMClass) AddToVCPUsParams(sessionID SessionRef, self VMRef, key string, value string) (_err error) {
	if IsMock {
		return _class.AddToVCPUsParamsMock(sessionID, self, key, value)
	}	
	_method := "VM.add_to_VCPUs_params"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_keyArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "key"), key)
	if _err != nil {
		return
	}
	_valueArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _keyArg, _valueArg)
	return
}


func VMClassSetVCPUsParamsMockDefault(sessionID SessionRef, self VMRef, value map[string]string) (_err error) {
	log.Println("VM.SetVCPUsParams not mocked")
	_err = errors.New("VM.SetVCPUsParams not mocked")
	return
}

var VMClassSetVCPUsParamsMockedCallback = VMClassSetVCPUsParamsMockDefault

func (_class VMClass) SetVCPUsParamsMock(sessionID SessionRef, self VMRef, value map[string]string) (_err error) {
	return VMClassSetVCPUsParamsMockedCallback(sessionID, self, value)
}
// Set the VCPUs/params field of the given VM.
func (_class VMClass) SetVCPUsParams(sessionID SessionRef, self VMRef, value map[string]string) (_err error) {
	if IsMock {
		return _class.SetVCPUsParamsMock(sessionID, self, value)
	}	
	_method := "VM.set_VCPUs_params"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertStringToStringMapToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}


func VMClassSetAffinityMockDefault(sessionID SessionRef, self VMRef, value HostRef) (_err error) {
	log.Println("VM.SetAffinity not mocked")
	_err = errors.New("VM.SetAffinity not mocked")
	return
}

var VMClassSetAffinityMockedCallback = VMClassSetAffinityMockDefault

func (_class VMClass) SetAffinityMock(sessionID SessionRef, self VMRef, value HostRef) (_err error) {
	return VMClassSetAffinityMockedCallback(sessionID, self, value)
}
// Set the affinity field of the given VM.
func (_class VMClass) SetAffinity(sessionID SessionRef, self VMRef, value HostRef) (_err error) {
	if IsMock {
		return _class.SetAffinityMock(sessionID, self, value)
	}	
	_method := "VM.set_affinity"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}


func VMClassSetIsATemplateMockDefault(sessionID SessionRef, self VMRef, value bool) (_err error) {
	log.Println("VM.SetIsATemplate not mocked")
	_err = errors.New("VM.SetIsATemplate not mocked")
	return
}

var VMClassSetIsATemplateMockedCallback = VMClassSetIsATemplateMockDefault

func (_class VMClass) SetIsATemplateMock(sessionID SessionRef, self VMRef, value bool) (_err error) {
	return VMClassSetIsATemplateMockedCallback(sessionID, self, value)
}
// Set the is_a_template field of the given VM.
func (_class VMClass) SetIsATemplate(sessionID SessionRef, self VMRef, value bool) (_err error) {
	if IsMock {
		return _class.SetIsATemplateMock(sessionID, self, value)
	}	
	_method := "VM.set_is_a_template"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertBoolToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}


func VMClassSetUserVersionMockDefault(sessionID SessionRef, self VMRef, value int) (_err error) {
	log.Println("VM.SetUserVersion not mocked")
	_err = errors.New("VM.SetUserVersion not mocked")
	return
}

var VMClassSetUserVersionMockedCallback = VMClassSetUserVersionMockDefault

func (_class VMClass) SetUserVersionMock(sessionID SessionRef, self VMRef, value int) (_err error) {
	return VMClassSetUserVersionMockedCallback(sessionID, self, value)
}
// Set the user_version field of the given VM.
func (_class VMClass) SetUserVersion(sessionID SessionRef, self VMRef, value int) (_err error) {
	if IsMock {
		return _class.SetUserVersionMock(sessionID, self, value)
	}	
	_method := "VM.set_user_version"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertIntToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}


func VMClassSetNameDescriptionMockDefault(sessionID SessionRef, self VMRef, value string) (_err error) {
	log.Println("VM.SetNameDescription not mocked")
	_err = errors.New("VM.SetNameDescription not mocked")
	return
}

var VMClassSetNameDescriptionMockedCallback = VMClassSetNameDescriptionMockDefault

func (_class VMClass) SetNameDescriptionMock(sessionID SessionRef, self VMRef, value string) (_err error) {
	return VMClassSetNameDescriptionMockedCallback(sessionID, self, value)
}
// Set the name/description field of the given VM.
func (_class VMClass) SetNameDescription(sessionID SessionRef, self VMRef, value string) (_err error) {
	if IsMock {
		return _class.SetNameDescriptionMock(sessionID, self, value)
	}	
	_method := "VM.set_name_description"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}


func VMClassSetNameLabelMockDefault(sessionID SessionRef, self VMRef, value string) (_err error) {
	log.Println("VM.SetNameLabel not mocked")
	_err = errors.New("VM.SetNameLabel not mocked")
	return
}

var VMClassSetNameLabelMockedCallback = VMClassSetNameLabelMockDefault

func (_class VMClass) SetNameLabelMock(sessionID SessionRef, self VMRef, value string) (_err error) {
	return VMClassSetNameLabelMockedCallback(sessionID, self, value)
}
// Set the name/label field of the given VM.
func (_class VMClass) SetNameLabel(sessionID SessionRef, self VMRef, value string) (_err error) {
	if IsMock {
		return _class.SetNameLabelMock(sessionID, self, value)
	}	
	_method := "VM.set_name_label"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}


func VMClassGetReferenceLabelMockDefault(sessionID SessionRef, self VMRef) (_retval string, _err error) {
	log.Println("VM.GetReferenceLabel not mocked")
	_err = errors.New("VM.GetReferenceLabel not mocked")
	return
}

var VMClassGetReferenceLabelMockedCallback = VMClassGetReferenceLabelMockDefault

func (_class VMClass) GetReferenceLabelMock(sessionID SessionRef, self VMRef) (_retval string, _err error) {
	return VMClassGetReferenceLabelMockedCallback(sessionID, self)
}
// Get the reference_label field of the given VM.
func (_class VMClass) GetReferenceLabel(sessionID SessionRef, self VMRef) (_retval string, _err error) {
	if IsMock {
		return _class.GetReferenceLabelMock(sessionID, self)
	}	
	_method := "VM.get_reference_label"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringToGo(_method + " -> ", _result.Value)
	return
}


func VMClassGetRequiresRebootMockDefault(sessionID SessionRef, self VMRef) (_retval bool, _err error) {
	log.Println("VM.GetRequiresReboot not mocked")
	_err = errors.New("VM.GetRequiresReboot not mocked")
	return
}

var VMClassGetRequiresRebootMockedCallback = VMClassGetRequiresRebootMockDefault

func (_class VMClass) GetRequiresRebootMock(sessionID SessionRef, self VMRef) (_retval bool, _err error) {
	return VMClassGetRequiresRebootMockedCallback(sessionID, self)
}
// Get the requires_reboot field of the given VM.
func (_class VMClass) GetRequiresReboot(sessionID SessionRef, self VMRef) (_retval bool, _err error) {
	if IsMock {
		return _class.GetRequiresRebootMock(sessionID, self)
	}	
	_method := "VM.get_requires_reboot"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertBoolToGo(_method + " -> ", _result.Value)
	return
}


func VMClassGetHasVendorDeviceMockDefault(sessionID SessionRef, self VMRef) (_retval bool, _err error) {
	log.Println("VM.GetHasVendorDevice not mocked")
	_err = errors.New("VM.GetHasVendorDevice not mocked")
	return
}

var VMClassGetHasVendorDeviceMockedCallback = VMClassGetHasVendorDeviceMockDefault

func (_class VMClass) GetHasVendorDeviceMock(sessionID SessionRef, self VMRef) (_retval bool, _err error) {
	return VMClassGetHasVendorDeviceMockedCallback(sessionID, self)
}
// Get the has_vendor_device field of the given VM.
func (_class VMClass) GetHasVendorDevice(sessionID SessionRef, self VMRef) (_retval bool, _err error) {
	if IsMock {
		return _class.GetHasVendorDeviceMock(sessionID, self)
	}	
	_method := "VM.get_has_vendor_device"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertBoolToGo(_method + " -> ", _result.Value)
	return
}


func VMClassGetHardwarePlatformVersionMockDefault(sessionID SessionRef, self VMRef) (_retval int, _err error) {
	log.Println("VM.GetHardwarePlatformVersion not mocked")
	_err = errors.New("VM.GetHardwarePlatformVersion not mocked")
	return
}

var VMClassGetHardwarePlatformVersionMockedCallback = VMClassGetHardwarePlatformVersionMockDefault

func (_class VMClass) GetHardwarePlatformVersionMock(sessionID SessionRef, self VMRef) (_retval int, _err error) {
	return VMClassGetHardwarePlatformVersionMockedCallback(sessionID, self)
}
// Get the hardware_platform_version field of the given VM.
func (_class VMClass) GetHardwarePlatformVersion(sessionID SessionRef, self VMRef) (_retval int, _err error) {
	if IsMock {
		return _class.GetHardwarePlatformVersionMock(sessionID, self)
	}	
	_method := "VM.get_hardware_platform_version"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertIntToGo(_method + " -> ", _result.Value)
	return
}


func VMClassGetGenerationIDMockDefault(sessionID SessionRef, self VMRef) (_retval string, _err error) {
	log.Println("VM.GetGenerationID not mocked")
	_err = errors.New("VM.GetGenerationID not mocked")
	return
}

var VMClassGetGenerationIDMockedCallback = VMClassGetGenerationIDMockDefault

func (_class VMClass) GetGenerationIDMock(sessionID SessionRef, self VMRef) (_retval string, _err error) {
	return VMClassGetGenerationIDMockedCallback(sessionID, self)
}
// Get the generation_id field of the given VM.
func (_class VMClass) GetGenerationID(sessionID SessionRef, self VMRef) (_retval string, _err error) {
	if IsMock {
		return _class.GetGenerationIDMock(sessionID, self)
	}	
	_method := "VM.get_generation_id"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringToGo(_method + " -> ", _result.Value)
	return
}


func VMClassGetVersionMockDefault(sessionID SessionRef, self VMRef) (_retval int, _err error) {
	log.Println("VM.GetVersion not mocked")
	_err = errors.New("VM.GetVersion not mocked")
	return
}

var VMClassGetVersionMockedCallback = VMClassGetVersionMockDefault

func (_class VMClass) GetVersionMock(sessionID SessionRef, self VMRef) (_retval int, _err error) {
	return VMClassGetVersionMockedCallback(sessionID, self)
}
// Get the version field of the given VM.
func (_class VMClass) GetVersion(sessionID SessionRef, self VMRef) (_retval int, _err error) {
	if IsMock {
		return _class.GetVersionMock(sessionID, self)
	}	
	_method := "VM.get_version"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertIntToGo(_method + " -> ", _result.Value)
	return
}


func VMClassGetSuspendSRMockDefault(sessionID SessionRef, self VMRef) (_retval SRRef, _err error) {
	log.Println("VM.GetSuspendSR not mocked")
	_err = errors.New("VM.GetSuspendSR not mocked")
	return
}

var VMClassGetSuspendSRMockedCallback = VMClassGetSuspendSRMockDefault

func (_class VMClass) GetSuspendSRMock(sessionID SessionRef, self VMRef) (_retval SRRef, _err error) {
	return VMClassGetSuspendSRMockedCallback(sessionID, self)
}
// Get the suspend_SR field of the given VM.
func (_class VMClass) GetSuspendSR(sessionID SessionRef, self VMRef) (_retval SRRef, _err error) {
	if IsMock {
		return _class.GetSuspendSRMock(sessionID, self)
	}	
	_method := "VM.get_suspend_SR"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertSRRefToGo(_method + " -> ", _result.Value)
	return
}


func VMClassGetAttachedPCIsMockDefault(sessionID SessionRef, self VMRef) (_retval []PCIRef, _err error) {
	log.Println("VM.GetAttachedPCIs not mocked")
	_err = errors.New("VM.GetAttachedPCIs not mocked")
	return
}

var VMClassGetAttachedPCIsMockedCallback = VMClassGetAttachedPCIsMockDefault

func (_class VMClass) GetAttachedPCIsMock(sessionID SessionRef, self VMRef) (_retval []PCIRef, _err error) {
	return VMClassGetAttachedPCIsMockedCallback(sessionID, self)
}
// Get the attached_PCIs field of the given VM.
func (_class VMClass) GetAttachedPCIs(sessionID SessionRef, self VMRef) (_retval []PCIRef, _err error) {
	if IsMock {
		return _class.GetAttachedPCIsMock(sessionID, self)
	}	
	_method := "VM.get_attached_PCIs"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertPCIRefSetToGo(_method + " -> ", _result.Value)
	return
}


func VMClassGetVGPUsMockDefault(sessionID SessionRef, self VMRef) (_retval []VGPURef, _err error) {
	log.Println("VM.GetVGPUs not mocked")
	_err = errors.New("VM.GetVGPUs not mocked")
	return
}

var VMClassGetVGPUsMockedCallback = VMClassGetVGPUsMockDefault

func (_class VMClass) GetVGPUsMock(sessionID SessionRef, self VMRef) (_retval []VGPURef, _err error) {
	return VMClassGetVGPUsMockedCallback(sessionID, self)
}
// Get the VGPUs field of the given VM.
func (_class VMClass) GetVGPUs(sessionID SessionRef, self VMRef) (_retval []VGPURef, _err error) {
	if IsMock {
		return _class.GetVGPUsMock(sessionID, self)
	}	
	_method := "VM.get_VGPUs"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVGPURefSetToGo(_method + " -> ", _result.Value)
	return
}


func VMClassGetOrderMockDefault(sessionID SessionRef, self VMRef) (_retval int, _err error) {
	log.Println("VM.GetOrder not mocked")
	_err = errors.New("VM.GetOrder not mocked")
	return
}

var VMClassGetOrderMockedCallback = VMClassGetOrderMockDefault

func (_class VMClass) GetOrderMock(sessionID SessionRef, self VMRef) (_retval int, _err error) {
	return VMClassGetOrderMockedCallback(sessionID, self)
}
// Get the order field of the given VM.
func (_class VMClass) GetOrder(sessionID SessionRef, self VMRef) (_retval int, _err error) {
	if IsMock {
		return _class.GetOrderMock(sessionID, self)
	}	
	_method := "VM.get_order"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertIntToGo(_method + " -> ", _result.Value)
	return
}


func VMClassGetShutdownDelayMockDefault(sessionID SessionRef, self VMRef) (_retval int, _err error) {
	log.Println("VM.GetShutdownDelay not mocked")
	_err = errors.New("VM.GetShutdownDelay not mocked")
	return
}

var VMClassGetShutdownDelayMockedCallback = VMClassGetShutdownDelayMockDefault

func (_class VMClass) GetShutdownDelayMock(sessionID SessionRef, self VMRef) (_retval int, _err error) {
	return VMClassGetShutdownDelayMockedCallback(sessionID, self)
}
// Get the shutdown_delay field of the given VM.
func (_class VMClass) GetShutdownDelay(sessionID SessionRef, self VMRef) (_retval int, _err error) {
	if IsMock {
		return _class.GetShutdownDelayMock(sessionID, self)
	}	
	_method := "VM.get_shutdown_delay"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertIntToGo(_method + " -> ", _result.Value)
	return
}


func VMClassGetStartDelayMockDefault(sessionID SessionRef, self VMRef) (_retval int, _err error) {
	log.Println("VM.GetStartDelay not mocked")
	_err = errors.New("VM.GetStartDelay not mocked")
	return
}

var VMClassGetStartDelayMockedCallback = VMClassGetStartDelayMockDefault

func (_class VMClass) GetStartDelayMock(sessionID SessionRef, self VMRef) (_retval int, _err error) {
	return VMClassGetStartDelayMockedCallback(sessionID, self)
}
// Get the start_delay field of the given VM.
func (_class VMClass) GetStartDelay(sessionID SessionRef, self VMRef) (_retval int, _err error) {
	if IsMock {
		return _class.GetStartDelayMock(sessionID, self)
	}	
	_method := "VM.get_start_delay"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertIntToGo(_method + " -> ", _result.Value)
	return
}


func VMClassGetApplianceMockDefault(sessionID SessionRef, self VMRef) (_retval VMApplianceRef, _err error) {
	log.Println("VM.GetAppliance not mocked")
	_err = errors.New("VM.GetAppliance not mocked")
	return
}

var VMClassGetApplianceMockedCallback = VMClassGetApplianceMockDefault

func (_class VMClass) GetApplianceMock(sessionID SessionRef, self VMRef) (_retval VMApplianceRef, _err error) {
	return VMClassGetApplianceMockedCallback(sessionID, self)
}
// Get the appliance field of the given VM.
func (_class VMClass) GetAppliance(sessionID SessionRef, self VMRef) (_retval VMApplianceRef, _err error) {
	if IsMock {
		return _class.GetApplianceMock(sessionID, self)
	}	
	_method := "VM.get_appliance"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVMApplianceRefToGo(_method + " -> ", _result.Value)
	return
}


func VMClassGetIsVmssSnapshotMockDefault(sessionID SessionRef, self VMRef) (_retval bool, _err error) {
	log.Println("VM.GetIsVmssSnapshot not mocked")
	_err = errors.New("VM.GetIsVmssSnapshot not mocked")
	return
}

var VMClassGetIsVmssSnapshotMockedCallback = VMClassGetIsVmssSnapshotMockDefault

func (_class VMClass) GetIsVmssSnapshotMock(sessionID SessionRef, self VMRef) (_retval bool, _err error) {
	return VMClassGetIsVmssSnapshotMockedCallback(sessionID, self)
}
// Get the is_vmss_snapshot field of the given VM.
func (_class VMClass) GetIsVmssSnapshot(sessionID SessionRef, self VMRef) (_retval bool, _err error) {
	if IsMock {
		return _class.GetIsVmssSnapshotMock(sessionID, self)
	}	
	_method := "VM.get_is_vmss_snapshot"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertBoolToGo(_method + " -> ", _result.Value)
	return
}


func VMClassGetSnapshotScheduleMockDefault(sessionID SessionRef, self VMRef) (_retval VMSSRef, _err error) {
	log.Println("VM.GetSnapshotSchedule not mocked")
	_err = errors.New("VM.GetSnapshotSchedule not mocked")
	return
}

var VMClassGetSnapshotScheduleMockedCallback = VMClassGetSnapshotScheduleMockDefault

func (_class VMClass) GetSnapshotScheduleMock(sessionID SessionRef, self VMRef) (_retval VMSSRef, _err error) {
	return VMClassGetSnapshotScheduleMockedCallback(sessionID, self)
}
// Get the snapshot_schedule field of the given VM.
func (_class VMClass) GetSnapshotSchedule(sessionID SessionRef, self VMRef) (_retval VMSSRef, _err error) {
	if IsMock {
		return _class.GetSnapshotScheduleMock(sessionID, self)
	}	
	_method := "VM.get_snapshot_schedule"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVMSSRefToGo(_method + " -> ", _result.Value)
	return
}


func VMClassGetIsSnapshotFromVmppMockDefault(sessionID SessionRef, self VMRef) (_retval bool, _err error) {
	log.Println("VM.GetIsSnapshotFromVmpp not mocked")
	_err = errors.New("VM.GetIsSnapshotFromVmpp not mocked")
	return
}

var VMClassGetIsSnapshotFromVmppMockedCallback = VMClassGetIsSnapshotFromVmppMockDefault

func (_class VMClass) GetIsSnapshotFromVmppMock(sessionID SessionRef, self VMRef) (_retval bool, _err error) {
	return VMClassGetIsSnapshotFromVmppMockedCallback(sessionID, self)
}
// Get the is_snapshot_from_vmpp field of the given VM.
func (_class VMClass) GetIsSnapshotFromVmpp(sessionID SessionRef, self VMRef) (_retval bool, _err error) {
	if IsMock {
		return _class.GetIsSnapshotFromVmppMock(sessionID, self)
	}	
	_method := "VM.get_is_snapshot_from_vmpp"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertBoolToGo(_method + " -> ", _result.Value)
	return
}


func VMClassGetProtectionPolicyMockDefault(sessionID SessionRef, self VMRef) (_retval VMPPRef, _err error) {
	log.Println("VM.GetProtectionPolicy not mocked")
	_err = errors.New("VM.GetProtectionPolicy not mocked")
	return
}

var VMClassGetProtectionPolicyMockedCallback = VMClassGetProtectionPolicyMockDefault

func (_class VMClass) GetProtectionPolicyMock(sessionID SessionRef, self VMRef) (_retval VMPPRef, _err error) {
	return VMClassGetProtectionPolicyMockedCallback(sessionID, self)
}
// Get the protection_policy field of the given VM.
func (_class VMClass) GetProtectionPolicy(sessionID SessionRef, self VMRef) (_retval VMPPRef, _err error) {
	if IsMock {
		return _class.GetProtectionPolicyMock(sessionID, self)
	}	
	_method := "VM.get_protection_policy"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVMPPRefToGo(_method + " -> ", _result.Value)
	return
}


func VMClassGetBiosStringsMockDefault(sessionID SessionRef, self VMRef) (_retval map[string]string, _err error) {
	log.Println("VM.GetBiosStrings not mocked")
	_err = errors.New("VM.GetBiosStrings not mocked")
	return
}

var VMClassGetBiosStringsMockedCallback = VMClassGetBiosStringsMockDefault

func (_class VMClass) GetBiosStringsMock(sessionID SessionRef, self VMRef) (_retval map[string]string, _err error) {
	return VMClassGetBiosStringsMockedCallback(sessionID, self)
}
// Get the bios_strings field of the given VM.
func (_class VMClass) GetBiosStrings(sessionID SessionRef, self VMRef) (_retval map[string]string, _err error) {
	if IsMock {
		return _class.GetBiosStringsMock(sessionID, self)
	}	
	_method := "VM.get_bios_strings"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringToStringMapToGo(_method + " -> ", _result.Value)
	return
}


func VMClassGetChildrenMockDefault(sessionID SessionRef, self VMRef) (_retval []VMRef, _err error) {
	log.Println("VM.GetChildren not mocked")
	_err = errors.New("VM.GetChildren not mocked")
	return
}

var VMClassGetChildrenMockedCallback = VMClassGetChildrenMockDefault

func (_class VMClass) GetChildrenMock(sessionID SessionRef, self VMRef) (_retval []VMRef, _err error) {
	return VMClassGetChildrenMockedCallback(sessionID, self)
}
// Get the children field of the given VM.
func (_class VMClass) GetChildren(sessionID SessionRef, self VMRef) (_retval []VMRef, _err error) {
	if IsMock {
		return _class.GetChildrenMock(sessionID, self)
	}	
	_method := "VM.get_children"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVMRefSetToGo(_method + " -> ", _result.Value)
	return
}


func VMClassGetParentMockDefault(sessionID SessionRef, self VMRef) (_retval VMRef, _err error) {
	log.Println("VM.GetParent not mocked")
	_err = errors.New("VM.GetParent not mocked")
	return
}

var VMClassGetParentMockedCallback = VMClassGetParentMockDefault

func (_class VMClass) GetParentMock(sessionID SessionRef, self VMRef) (_retval VMRef, _err error) {
	return VMClassGetParentMockedCallback(sessionID, self)
}
// Get the parent field of the given VM.
func (_class VMClass) GetParent(sessionID SessionRef, self VMRef) (_retval VMRef, _err error) {
	if IsMock {
		return _class.GetParentMock(sessionID, self)
	}	
	_method := "VM.get_parent"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVMRefToGo(_method + " -> ", _result.Value)
	return
}


func VMClassGetSnapshotMetadataMockDefault(sessionID SessionRef, self VMRef) (_retval string, _err error) {
	log.Println("VM.GetSnapshotMetadata not mocked")
	_err = errors.New("VM.GetSnapshotMetadata not mocked")
	return
}

var VMClassGetSnapshotMetadataMockedCallback = VMClassGetSnapshotMetadataMockDefault

func (_class VMClass) GetSnapshotMetadataMock(sessionID SessionRef, self VMRef) (_retval string, _err error) {
	return VMClassGetSnapshotMetadataMockedCallback(sessionID, self)
}
// Get the snapshot_metadata field of the given VM.
func (_class VMClass) GetSnapshotMetadata(sessionID SessionRef, self VMRef) (_retval string, _err error) {
	if IsMock {
		return _class.GetSnapshotMetadataMock(sessionID, self)
	}	
	_method := "VM.get_snapshot_metadata"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringToGo(_method + " -> ", _result.Value)
	return
}


func VMClassGetSnapshotInfoMockDefault(sessionID SessionRef, self VMRef) (_retval map[string]string, _err error) {
	log.Println("VM.GetSnapshotInfo not mocked")
	_err = errors.New("VM.GetSnapshotInfo not mocked")
	return
}

var VMClassGetSnapshotInfoMockedCallback = VMClassGetSnapshotInfoMockDefault

func (_class VMClass) GetSnapshotInfoMock(sessionID SessionRef, self VMRef) (_retval map[string]string, _err error) {
	return VMClassGetSnapshotInfoMockedCallback(sessionID, self)
}
// Get the snapshot_info field of the given VM.
func (_class VMClass) GetSnapshotInfo(sessionID SessionRef, self VMRef) (_retval map[string]string, _err error) {
	if IsMock {
		return _class.GetSnapshotInfoMock(sessionID, self)
	}	
	_method := "VM.get_snapshot_info"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringToStringMapToGo(_method + " -> ", _result.Value)
	return
}


func VMClassGetBlockedOperationsMockDefault(sessionID SessionRef, self VMRef) (_retval map[VMOperations]string, _err error) {
	log.Println("VM.GetBlockedOperations not mocked")
	_err = errors.New("VM.GetBlockedOperations not mocked")
	return
}

var VMClassGetBlockedOperationsMockedCallback = VMClassGetBlockedOperationsMockDefault

func (_class VMClass) GetBlockedOperationsMock(sessionID SessionRef, self VMRef) (_retval map[VMOperations]string, _err error) {
	return VMClassGetBlockedOperationsMockedCallback(sessionID, self)
}
// Get the blocked_operations field of the given VM.
func (_class VMClass) GetBlockedOperations(sessionID SessionRef, self VMRef) (_retval map[VMOperations]string, _err error) {
	if IsMock {
		return _class.GetBlockedOperationsMock(sessionID, self)
	}	
	_method := "VM.get_blocked_operations"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertEnumVMOperationsToStringMapToGo(_method + " -> ", _result.Value)
	return
}


func VMClassGetTagsMockDefault(sessionID SessionRef, self VMRef) (_retval []string, _err error) {
	log.Println("VM.GetTags not mocked")
	_err = errors.New("VM.GetTags not mocked")
	return
}

var VMClassGetTagsMockedCallback = VMClassGetTagsMockDefault

func (_class VMClass) GetTagsMock(sessionID SessionRef, self VMRef) (_retval []string, _err error) {
	return VMClassGetTagsMockedCallback(sessionID, self)
}
// Get the tags field of the given VM.
func (_class VMClass) GetTags(sessionID SessionRef, self VMRef) (_retval []string, _err error) {
	if IsMock {
		return _class.GetTagsMock(sessionID, self)
	}	
	_method := "VM.get_tags"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringSetToGo(_method + " -> ", _result.Value)
	return
}


func VMClassGetBlobsMockDefault(sessionID SessionRef, self VMRef) (_retval map[string]BlobRef, _err error) {
	log.Println("VM.GetBlobs not mocked")
	_err = errors.New("VM.GetBlobs not mocked")
	return
}

var VMClassGetBlobsMockedCallback = VMClassGetBlobsMockDefault

func (_class VMClass) GetBlobsMock(sessionID SessionRef, self VMRef) (_retval map[string]BlobRef, _err error) {
	return VMClassGetBlobsMockedCallback(sessionID, self)
}
// Get the blobs field of the given VM.
func (_class VMClass) GetBlobs(sessionID SessionRef, self VMRef) (_retval map[string]BlobRef, _err error) {
	if IsMock {
		return _class.GetBlobsMock(sessionID, self)
	}	
	_method := "VM.get_blobs"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringToBlobRefMapToGo(_method + " -> ", _result.Value)
	return
}


func VMClassGetTransportableSnapshotIDMockDefault(sessionID SessionRef, self VMRef) (_retval string, _err error) {
	log.Println("VM.GetTransportableSnapshotID not mocked")
	_err = errors.New("VM.GetTransportableSnapshotID not mocked")
	return
}

var VMClassGetTransportableSnapshotIDMockedCallback = VMClassGetTransportableSnapshotIDMockDefault

func (_class VMClass) GetTransportableSnapshotIDMock(sessionID SessionRef, self VMRef) (_retval string, _err error) {
	return VMClassGetTransportableSnapshotIDMockedCallback(sessionID, self)
}
// Get the transportable_snapshot_id field of the given VM.
func (_class VMClass) GetTransportableSnapshotID(sessionID SessionRef, self VMRef) (_retval string, _err error) {
	if IsMock {
		return _class.GetTransportableSnapshotIDMock(sessionID, self)
	}	
	_method := "VM.get_transportable_snapshot_id"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringToGo(_method + " -> ", _result.Value)
	return
}


func VMClassGetSnapshotTimeMockDefault(sessionID SessionRef, self VMRef) (_retval time.Time, _err error) {
	log.Println("VM.GetSnapshotTime not mocked")
	_err = errors.New("VM.GetSnapshotTime not mocked")
	return
}

var VMClassGetSnapshotTimeMockedCallback = VMClassGetSnapshotTimeMockDefault

func (_class VMClass) GetSnapshotTimeMock(sessionID SessionRef, self VMRef) (_retval time.Time, _err error) {
	return VMClassGetSnapshotTimeMockedCallback(sessionID, self)
}
// Get the snapshot_time field of the given VM.
func (_class VMClass) GetSnapshotTime(sessionID SessionRef, self VMRef) (_retval time.Time, _err error) {
	if IsMock {
		return _class.GetSnapshotTimeMock(sessionID, self)
	}	
	_method := "VM.get_snapshot_time"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTimeToGo(_method + " -> ", _result.Value)
	return
}


func VMClassGetSnapshotsMockDefault(sessionID SessionRef, self VMRef) (_retval []VMRef, _err error) {
	log.Println("VM.GetSnapshots not mocked")
	_err = errors.New("VM.GetSnapshots not mocked")
	return
}

var VMClassGetSnapshotsMockedCallback = VMClassGetSnapshotsMockDefault

func (_class VMClass) GetSnapshotsMock(sessionID SessionRef, self VMRef) (_retval []VMRef, _err error) {
	return VMClassGetSnapshotsMockedCallback(sessionID, self)
}
// Get the snapshots field of the given VM.
func (_class VMClass) GetSnapshots(sessionID SessionRef, self VMRef) (_retval []VMRef, _err error) {
	if IsMock {
		return _class.GetSnapshotsMock(sessionID, self)
	}	
	_method := "VM.get_snapshots"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVMRefSetToGo(_method + " -> ", _result.Value)
	return
}


func VMClassGetSnapshotOfMockDefault(sessionID SessionRef, self VMRef) (_retval VMRef, _err error) {
	log.Println("VM.GetSnapshotOf not mocked")
	_err = errors.New("VM.GetSnapshotOf not mocked")
	return
}

var VMClassGetSnapshotOfMockedCallback = VMClassGetSnapshotOfMockDefault

func (_class VMClass) GetSnapshotOfMock(sessionID SessionRef, self VMRef) (_retval VMRef, _err error) {
	return VMClassGetSnapshotOfMockedCallback(sessionID, self)
}
// Get the snapshot_of field of the given VM.
func (_class VMClass) GetSnapshotOf(sessionID SessionRef, self VMRef) (_retval VMRef, _err error) {
	if IsMock {
		return _class.GetSnapshotOfMock(sessionID, self)
	}	
	_method := "VM.get_snapshot_of"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVMRefToGo(_method + " -> ", _result.Value)
	return
}


func VMClassGetIsASnapshotMockDefault(sessionID SessionRef, self VMRef) (_retval bool, _err error) {
	log.Println("VM.GetIsASnapshot not mocked")
	_err = errors.New("VM.GetIsASnapshot not mocked")
	return
}

var VMClassGetIsASnapshotMockedCallback = VMClassGetIsASnapshotMockDefault

func (_class VMClass) GetIsASnapshotMock(sessionID SessionRef, self VMRef) (_retval bool, _err error) {
	return VMClassGetIsASnapshotMockedCallback(sessionID, self)
}
// Get the is_a_snapshot field of the given VM.
func (_class VMClass) GetIsASnapshot(sessionID SessionRef, self VMRef) (_retval bool, _err error) {
	if IsMock {
		return _class.GetIsASnapshotMock(sessionID, self)
	}	
	_method := "VM.get_is_a_snapshot"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertBoolToGo(_method + " -> ", _result.Value)
	return
}


func VMClassGetHaRestartPriorityMockDefault(sessionID SessionRef, self VMRef) (_retval string, _err error) {
	log.Println("VM.GetHaRestartPriority not mocked")
	_err = errors.New("VM.GetHaRestartPriority not mocked")
	return
}

var VMClassGetHaRestartPriorityMockedCallback = VMClassGetHaRestartPriorityMockDefault

func (_class VMClass) GetHaRestartPriorityMock(sessionID SessionRef, self VMRef) (_retval string, _err error) {
	return VMClassGetHaRestartPriorityMockedCallback(sessionID, self)
}
// Get the ha_restart_priority field of the given VM.
func (_class VMClass) GetHaRestartPriority(sessionID SessionRef, self VMRef) (_retval string, _err error) {
	if IsMock {
		return _class.GetHaRestartPriorityMock(sessionID, self)
	}	
	_method := "VM.get_ha_restart_priority"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringToGo(_method + " -> ", _result.Value)
	return
}


func VMClassGetHaAlwaysRunMockDefault(sessionID SessionRef, self VMRef) (_retval bool, _err error) {
	log.Println("VM.GetHaAlwaysRun not mocked")
	_err = errors.New("VM.GetHaAlwaysRun not mocked")
	return
}

var VMClassGetHaAlwaysRunMockedCallback = VMClassGetHaAlwaysRunMockDefault

func (_class VMClass) GetHaAlwaysRunMock(sessionID SessionRef, self VMRef) (_retval bool, _err error) {
	return VMClassGetHaAlwaysRunMockedCallback(sessionID, self)
}
// Get the ha_always_run field of the given VM.
func (_class VMClass) GetHaAlwaysRun(sessionID SessionRef, self VMRef) (_retval bool, _err error) {
	if IsMock {
		return _class.GetHaAlwaysRunMock(sessionID, self)
	}	
	_method := "VM.get_ha_always_run"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertBoolToGo(_method + " -> ", _result.Value)
	return
}


func VMClassGetXenstoreDataMockDefault(sessionID SessionRef, self VMRef) (_retval map[string]string, _err error) {
	log.Println("VM.GetXenstoreData not mocked")
	_err = errors.New("VM.GetXenstoreData not mocked")
	return
}

var VMClassGetXenstoreDataMockedCallback = VMClassGetXenstoreDataMockDefault

func (_class VMClass) GetXenstoreDataMock(sessionID SessionRef, self VMRef) (_retval map[string]string, _err error) {
	return VMClassGetXenstoreDataMockedCallback(sessionID, self)
}
// Get the xenstore_data field of the given VM.
func (_class VMClass) GetXenstoreData(sessionID SessionRef, self VMRef) (_retval map[string]string, _err error) {
	if IsMock {
		return _class.GetXenstoreDataMock(sessionID, self)
	}	
	_method := "VM.get_xenstore_data"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringToStringMapToGo(_method + " -> ", _result.Value)
	return
}


func VMClassGetRecommendationsMockDefault(sessionID SessionRef, self VMRef) (_retval string, _err error) {
	log.Println("VM.GetRecommendations not mocked")
	_err = errors.New("VM.GetRecommendations not mocked")
	return
}

var VMClassGetRecommendationsMockedCallback = VMClassGetRecommendationsMockDefault

func (_class VMClass) GetRecommendationsMock(sessionID SessionRef, self VMRef) (_retval string, _err error) {
	return VMClassGetRecommendationsMockedCallback(sessionID, self)
}
// Get the recommendations field of the given VM.
func (_class VMClass) GetRecommendations(sessionID SessionRef, self VMRef) (_retval string, _err error) {
	if IsMock {
		return _class.GetRecommendationsMock(sessionID, self)
	}	
	_method := "VM.get_recommendations"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringToGo(_method + " -> ", _result.Value)
	return
}


func VMClassGetLastBootedRecordMockDefault(sessionID SessionRef, self VMRef) (_retval string, _err error) {
	log.Println("VM.GetLastBootedRecord not mocked")
	_err = errors.New("VM.GetLastBootedRecord not mocked")
	return
}

var VMClassGetLastBootedRecordMockedCallback = VMClassGetLastBootedRecordMockDefault

func (_class VMClass) GetLastBootedRecordMock(sessionID SessionRef, self VMRef) (_retval string, _err error) {
	return VMClassGetLastBootedRecordMockedCallback(sessionID, self)
}
// Get the last_booted_record field of the given VM.
func (_class VMClass) GetLastBootedRecord(sessionID SessionRef, self VMRef) (_retval string, _err error) {
	if IsMock {
		return _class.GetLastBootedRecordMock(sessionID, self)
	}	
	_method := "VM.get_last_booted_record"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringToGo(_method + " -> ", _result.Value)
	return
}


func VMClassGetGuestMetricsMockDefault(sessionID SessionRef, self VMRef) (_retval VMGuestMetricsRef, _err error) {
	log.Println("VM.GetGuestMetrics not mocked")
	_err = errors.New("VM.GetGuestMetrics not mocked")
	return
}

var VMClassGetGuestMetricsMockedCallback = VMClassGetGuestMetricsMockDefault

func (_class VMClass) GetGuestMetricsMock(sessionID SessionRef, self VMRef) (_retval VMGuestMetricsRef, _err error) {
	return VMClassGetGuestMetricsMockedCallback(sessionID, self)
}
// Get the guest_metrics field of the given VM.
func (_class VMClass) GetGuestMetrics(sessionID SessionRef, self VMRef) (_retval VMGuestMetricsRef, _err error) {
	if IsMock {
		return _class.GetGuestMetricsMock(sessionID, self)
	}	
	_method := "VM.get_guest_metrics"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVMGuestMetricsRefToGo(_method + " -> ", _result.Value)
	return
}


func VMClassGetMetricsMockDefault(sessionID SessionRef, self VMRef) (_retval VMMetricsRef, _err error) {
	log.Println("VM.GetMetrics not mocked")
	_err = errors.New("VM.GetMetrics not mocked")
	return
}

var VMClassGetMetricsMockedCallback = VMClassGetMetricsMockDefault

func (_class VMClass) GetMetricsMock(sessionID SessionRef, self VMRef) (_retval VMMetricsRef, _err error) {
	return VMClassGetMetricsMockedCallback(sessionID, self)
}
// Get the metrics field of the given VM.
func (_class VMClass) GetMetrics(sessionID SessionRef, self VMRef) (_retval VMMetricsRef, _err error) {
	if IsMock {
		return _class.GetMetricsMock(sessionID, self)
	}	
	_method := "VM.get_metrics"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVMMetricsRefToGo(_method + " -> ", _result.Value)
	return
}


func VMClassGetIsControlDomainMockDefault(sessionID SessionRef, self VMRef) (_retval bool, _err error) {
	log.Println("VM.GetIsControlDomain not mocked")
	_err = errors.New("VM.GetIsControlDomain not mocked")
	return
}

var VMClassGetIsControlDomainMockedCallback = VMClassGetIsControlDomainMockDefault

func (_class VMClass) GetIsControlDomainMock(sessionID SessionRef, self VMRef) (_retval bool, _err error) {
	return VMClassGetIsControlDomainMockedCallback(sessionID, self)
}
// Get the is_control_domain field of the given VM.
func (_class VMClass) GetIsControlDomain(sessionID SessionRef, self VMRef) (_retval bool, _err error) {
	if IsMock {
		return _class.GetIsControlDomainMock(sessionID, self)
	}	
	_method := "VM.get_is_control_domain"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertBoolToGo(_method + " -> ", _result.Value)
	return
}


func VMClassGetLastBootCPUFlagsMockDefault(sessionID SessionRef, self VMRef) (_retval map[string]string, _err error) {
	log.Println("VM.GetLastBootCPUFlags not mocked")
	_err = errors.New("VM.GetLastBootCPUFlags not mocked")
	return
}

var VMClassGetLastBootCPUFlagsMockedCallback = VMClassGetLastBootCPUFlagsMockDefault

func (_class VMClass) GetLastBootCPUFlagsMock(sessionID SessionRef, self VMRef) (_retval map[string]string, _err error) {
	return VMClassGetLastBootCPUFlagsMockedCallback(sessionID, self)
}
// Get the last_boot_CPU_flags field of the given VM.
func (_class VMClass) GetLastBootCPUFlags(sessionID SessionRef, self VMRef) (_retval map[string]string, _err error) {
	if IsMock {
		return _class.GetLastBootCPUFlagsMock(sessionID, self)
	}	
	_method := "VM.get_last_boot_CPU_flags"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringToStringMapToGo(_method + " -> ", _result.Value)
	return
}


func VMClassGetDomarchMockDefault(sessionID SessionRef, self VMRef) (_retval string, _err error) {
	log.Println("VM.GetDomarch not mocked")
	_err = errors.New("VM.GetDomarch not mocked")
	return
}

var VMClassGetDomarchMockedCallback = VMClassGetDomarchMockDefault

func (_class VMClass) GetDomarchMock(sessionID SessionRef, self VMRef) (_retval string, _err error) {
	return VMClassGetDomarchMockedCallback(sessionID, self)
}
// Get the domarch field of the given VM.
func (_class VMClass) GetDomarch(sessionID SessionRef, self VMRef) (_retval string, _err error) {
	if IsMock {
		return _class.GetDomarchMock(sessionID, self)
	}	
	_method := "VM.get_domarch"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringToGo(_method + " -> ", _result.Value)
	return
}


func VMClassGetDomidMockDefault(sessionID SessionRef, self VMRef) (_retval int, _err error) {
	log.Println("VM.GetDomid not mocked")
	_err = errors.New("VM.GetDomid not mocked")
	return
}

var VMClassGetDomidMockedCallback = VMClassGetDomidMockDefault

func (_class VMClass) GetDomidMock(sessionID SessionRef, self VMRef) (_retval int, _err error) {
	return VMClassGetDomidMockedCallback(sessionID, self)
}
// Get the domid field of the given VM.
func (_class VMClass) GetDomid(sessionID SessionRef, self VMRef) (_retval int, _err error) {
	if IsMock {
		return _class.GetDomidMock(sessionID, self)
	}	
	_method := "VM.get_domid"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertIntToGo(_method + " -> ", _result.Value)
	return
}


func VMClassGetOtherConfigMockDefault(sessionID SessionRef, self VMRef) (_retval map[string]string, _err error) {
	log.Println("VM.GetOtherConfig not mocked")
	_err = errors.New("VM.GetOtherConfig not mocked")
	return
}

var VMClassGetOtherConfigMockedCallback = VMClassGetOtherConfigMockDefault

func (_class VMClass) GetOtherConfigMock(sessionID SessionRef, self VMRef) (_retval map[string]string, _err error) {
	return VMClassGetOtherConfigMockedCallback(sessionID, self)
}
// Get the other_config field of the given VM.
func (_class VMClass) GetOtherConfig(sessionID SessionRef, self VMRef) (_retval map[string]string, _err error) {
	if IsMock {
		return _class.GetOtherConfigMock(sessionID, self)
	}	
	_method := "VM.get_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringToStringMapToGo(_method + " -> ", _result.Value)
	return
}


func VMClassGetPCIBusMockDefault(sessionID SessionRef, self VMRef) (_retval string, _err error) {
	log.Println("VM.GetPCIBus not mocked")
	_err = errors.New("VM.GetPCIBus not mocked")
	return
}

var VMClassGetPCIBusMockedCallback = VMClassGetPCIBusMockDefault

func (_class VMClass) GetPCIBusMock(sessionID SessionRef, self VMRef) (_retval string, _err error) {
	return VMClassGetPCIBusMockedCallback(sessionID, self)
}
// Get the PCI_bus field of the given VM.
func (_class VMClass) GetPCIBus(sessionID SessionRef, self VMRef) (_retval string, _err error) {
	if IsMock {
		return _class.GetPCIBusMock(sessionID, self)
	}	
	_method := "VM.get_PCI_bus"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringToGo(_method + " -> ", _result.Value)
	return
}


func VMClassGetPlatformMockDefault(sessionID SessionRef, self VMRef) (_retval map[string]string, _err error) {
	log.Println("VM.GetPlatform not mocked")
	_err = errors.New("VM.GetPlatform not mocked")
	return
}

var VMClassGetPlatformMockedCallback = VMClassGetPlatformMockDefault

func (_class VMClass) GetPlatformMock(sessionID SessionRef, self VMRef) (_retval map[string]string, _err error) {
	return VMClassGetPlatformMockedCallback(sessionID, self)
}
// Get the platform field of the given VM.
func (_class VMClass) GetPlatform(sessionID SessionRef, self VMRef) (_retval map[string]string, _err error) {
	if IsMock {
		return _class.GetPlatformMock(sessionID, self)
	}	
	_method := "VM.get_platform"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringToStringMapToGo(_method + " -> ", _result.Value)
	return
}


func VMClassGetHVMShadowMultiplierMockDefault(sessionID SessionRef, self VMRef) (_retval float64, _err error) {
	log.Println("VM.GetHVMShadowMultiplier not mocked")
	_err = errors.New("VM.GetHVMShadowMultiplier not mocked")
	return
}

var VMClassGetHVMShadowMultiplierMockedCallback = VMClassGetHVMShadowMultiplierMockDefault

func (_class VMClass) GetHVMShadowMultiplierMock(sessionID SessionRef, self VMRef) (_retval float64, _err error) {
	return VMClassGetHVMShadowMultiplierMockedCallback(sessionID, self)
}
// Get the HVM/shadow_multiplier field of the given VM.
func (_class VMClass) GetHVMShadowMultiplier(sessionID SessionRef, self VMRef) (_retval float64, _err error) {
	if IsMock {
		return _class.GetHVMShadowMultiplierMock(sessionID, self)
	}	
	_method := "VM.get_HVM_shadow_multiplier"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertFloatToGo(_method + " -> ", _result.Value)
	return
}


func VMClassGetHVMBootParamsMockDefault(sessionID SessionRef, self VMRef) (_retval map[string]string, _err error) {
	log.Println("VM.GetHVMBootParams not mocked")
	_err = errors.New("VM.GetHVMBootParams not mocked")
	return
}

var VMClassGetHVMBootParamsMockedCallback = VMClassGetHVMBootParamsMockDefault

func (_class VMClass) GetHVMBootParamsMock(sessionID SessionRef, self VMRef) (_retval map[string]string, _err error) {
	return VMClassGetHVMBootParamsMockedCallback(sessionID, self)
}
// Get the HVM/boot_params field of the given VM.
func (_class VMClass) GetHVMBootParams(sessionID SessionRef, self VMRef) (_retval map[string]string, _err error) {
	if IsMock {
		return _class.GetHVMBootParamsMock(sessionID, self)
	}	
	_method := "VM.get_HVM_boot_params"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringToStringMapToGo(_method + " -> ", _result.Value)
	return
}


func VMClassGetHVMBootPolicyMockDefault(sessionID SessionRef, self VMRef) (_retval string, _err error) {
	log.Println("VM.GetHVMBootPolicy not mocked")
	_err = errors.New("VM.GetHVMBootPolicy not mocked")
	return
}

var VMClassGetHVMBootPolicyMockedCallback = VMClassGetHVMBootPolicyMockDefault

func (_class VMClass) GetHVMBootPolicyMock(sessionID SessionRef, self VMRef) (_retval string, _err error) {
	return VMClassGetHVMBootPolicyMockedCallback(sessionID, self)
}
// Get the HVM/boot_policy field of the given VM.
func (_class VMClass) GetHVMBootPolicy(sessionID SessionRef, self VMRef) (_retval string, _err error) {
	if IsMock {
		return _class.GetHVMBootPolicyMock(sessionID, self)
	}	
	_method := "VM.get_HVM_boot_policy"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringToGo(_method + " -> ", _result.Value)
	return
}


func VMClassGetPVLegacyArgsMockDefault(sessionID SessionRef, self VMRef) (_retval string, _err error) {
	log.Println("VM.GetPVLegacyArgs not mocked")
	_err = errors.New("VM.GetPVLegacyArgs not mocked")
	return
}

var VMClassGetPVLegacyArgsMockedCallback = VMClassGetPVLegacyArgsMockDefault

func (_class VMClass) GetPVLegacyArgsMock(sessionID SessionRef, self VMRef) (_retval string, _err error) {
	return VMClassGetPVLegacyArgsMockedCallback(sessionID, self)
}
// Get the PV/legacy_args field of the given VM.
func (_class VMClass) GetPVLegacyArgs(sessionID SessionRef, self VMRef) (_retval string, _err error) {
	if IsMock {
		return _class.GetPVLegacyArgsMock(sessionID, self)
	}	
	_method := "VM.get_PV_legacy_args"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringToGo(_method + " -> ", _result.Value)
	return
}


func VMClassGetPVBootloaderArgsMockDefault(sessionID SessionRef, self VMRef) (_retval string, _err error) {
	log.Println("VM.GetPVBootloaderArgs not mocked")
	_err = errors.New("VM.GetPVBootloaderArgs not mocked")
	return
}

var VMClassGetPVBootloaderArgsMockedCallback = VMClassGetPVBootloaderArgsMockDefault

func (_class VMClass) GetPVBootloaderArgsMock(sessionID SessionRef, self VMRef) (_retval string, _err error) {
	return VMClassGetPVBootloaderArgsMockedCallback(sessionID, self)
}
// Get the PV/bootloader_args field of the given VM.
func (_class VMClass) GetPVBootloaderArgs(sessionID SessionRef, self VMRef) (_retval string, _err error) {
	if IsMock {
		return _class.GetPVBootloaderArgsMock(sessionID, self)
	}	
	_method := "VM.get_PV_bootloader_args"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringToGo(_method + " -> ", _result.Value)
	return
}


func VMClassGetPVArgsMockDefault(sessionID SessionRef, self VMRef) (_retval string, _err error) {
	log.Println("VM.GetPVArgs not mocked")
	_err = errors.New("VM.GetPVArgs not mocked")
	return
}

var VMClassGetPVArgsMockedCallback = VMClassGetPVArgsMockDefault

func (_class VMClass) GetPVArgsMock(sessionID SessionRef, self VMRef) (_retval string, _err error) {
	return VMClassGetPVArgsMockedCallback(sessionID, self)
}
// Get the PV/args field of the given VM.
func (_class VMClass) GetPVArgs(sessionID SessionRef, self VMRef) (_retval string, _err error) {
	if IsMock {
		return _class.GetPVArgsMock(sessionID, self)
	}	
	_method := "VM.get_PV_args"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringToGo(_method + " -> ", _result.Value)
	return
}


func VMClassGetPVRamdiskMockDefault(sessionID SessionRef, self VMRef) (_retval string, _err error) {
	log.Println("VM.GetPVRamdisk not mocked")
	_err = errors.New("VM.GetPVRamdisk not mocked")
	return
}

var VMClassGetPVRamdiskMockedCallback = VMClassGetPVRamdiskMockDefault

func (_class VMClass) GetPVRamdiskMock(sessionID SessionRef, self VMRef) (_retval string, _err error) {
	return VMClassGetPVRamdiskMockedCallback(sessionID, self)
}
// Get the PV/ramdisk field of the given VM.
func (_class VMClass) GetPVRamdisk(sessionID SessionRef, self VMRef) (_retval string, _err error) {
	if IsMock {
		return _class.GetPVRamdiskMock(sessionID, self)
	}	
	_method := "VM.get_PV_ramdisk"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringToGo(_method + " -> ", _result.Value)
	return
}


func VMClassGetPVKernelMockDefault(sessionID SessionRef, self VMRef) (_retval string, _err error) {
	log.Println("VM.GetPVKernel not mocked")
	_err = errors.New("VM.GetPVKernel not mocked")
	return
}

var VMClassGetPVKernelMockedCallback = VMClassGetPVKernelMockDefault

func (_class VMClass) GetPVKernelMock(sessionID SessionRef, self VMRef) (_retval string, _err error) {
	return VMClassGetPVKernelMockedCallback(sessionID, self)
}
// Get the PV/kernel field of the given VM.
func (_class VMClass) GetPVKernel(sessionID SessionRef, self VMRef) (_retval string, _err error) {
	if IsMock {
		return _class.GetPVKernelMock(sessionID, self)
	}	
	_method := "VM.get_PV_kernel"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringToGo(_method + " -> ", _result.Value)
	return
}


func VMClassGetPVBootloaderMockDefault(sessionID SessionRef, self VMRef) (_retval string, _err error) {
	log.Println("VM.GetPVBootloader not mocked")
	_err = errors.New("VM.GetPVBootloader not mocked")
	return
}

var VMClassGetPVBootloaderMockedCallback = VMClassGetPVBootloaderMockDefault

func (_class VMClass) GetPVBootloaderMock(sessionID SessionRef, self VMRef) (_retval string, _err error) {
	return VMClassGetPVBootloaderMockedCallback(sessionID, self)
}
// Get the PV/bootloader field of the given VM.
func (_class VMClass) GetPVBootloader(sessionID SessionRef, self VMRef) (_retval string, _err error) {
	if IsMock {
		return _class.GetPVBootloaderMock(sessionID, self)
	}	
	_method := "VM.get_PV_bootloader"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringToGo(_method + " -> ", _result.Value)
	return
}


func VMClassGetVTPMsMockDefault(sessionID SessionRef, self VMRef) (_retval []VTPMRef, _err error) {
	log.Println("VM.GetVTPMs not mocked")
	_err = errors.New("VM.GetVTPMs not mocked")
	return
}

var VMClassGetVTPMsMockedCallback = VMClassGetVTPMsMockDefault

func (_class VMClass) GetVTPMsMock(sessionID SessionRef, self VMRef) (_retval []VTPMRef, _err error) {
	return VMClassGetVTPMsMockedCallback(sessionID, self)
}
// Get the VTPMs field of the given VM.
func (_class VMClass) GetVTPMs(sessionID SessionRef, self VMRef) (_retval []VTPMRef, _err error) {
	if IsMock {
		return _class.GetVTPMsMock(sessionID, self)
	}	
	_method := "VM.get_VTPMs"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVTPMRefSetToGo(_method + " -> ", _result.Value)
	return
}


func VMClassGetCrashDumpsMockDefault(sessionID SessionRef, self VMRef) (_retval []CrashdumpRef, _err error) {
	log.Println("VM.GetCrashDumps not mocked")
	_err = errors.New("VM.GetCrashDumps not mocked")
	return
}

var VMClassGetCrashDumpsMockedCallback = VMClassGetCrashDumpsMockDefault

func (_class VMClass) GetCrashDumpsMock(sessionID SessionRef, self VMRef) (_retval []CrashdumpRef, _err error) {
	return VMClassGetCrashDumpsMockedCallback(sessionID, self)
}
// Get the crash_dumps field of the given VM.
func (_class VMClass) GetCrashDumps(sessionID SessionRef, self VMRef) (_retval []CrashdumpRef, _err error) {
	if IsMock {
		return _class.GetCrashDumpsMock(sessionID, self)
	}	
	_method := "VM.get_crash_dumps"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertCrashdumpRefSetToGo(_method + " -> ", _result.Value)
	return
}


func VMClassGetVUSBsMockDefault(sessionID SessionRef, self VMRef) (_retval []VUSBRef, _err error) {
	log.Println("VM.GetVUSBs not mocked")
	_err = errors.New("VM.GetVUSBs not mocked")
	return
}

var VMClassGetVUSBsMockedCallback = VMClassGetVUSBsMockDefault

func (_class VMClass) GetVUSBsMock(sessionID SessionRef, self VMRef) (_retval []VUSBRef, _err error) {
	return VMClassGetVUSBsMockedCallback(sessionID, self)
}
// Get the VUSBs field of the given VM.
func (_class VMClass) GetVUSBs(sessionID SessionRef, self VMRef) (_retval []VUSBRef, _err error) {
	if IsMock {
		return _class.GetVUSBsMock(sessionID, self)
	}	
	_method := "VM.get_VUSBs"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVUSBRefSetToGo(_method + " -> ", _result.Value)
	return
}


func VMClassGetVBDsMockDefault(sessionID SessionRef, self VMRef) (_retval []VBDRef, _err error) {
	log.Println("VM.GetVBDs not mocked")
	_err = errors.New("VM.GetVBDs not mocked")
	return
}

var VMClassGetVBDsMockedCallback = VMClassGetVBDsMockDefault

func (_class VMClass) GetVBDsMock(sessionID SessionRef, self VMRef) (_retval []VBDRef, _err error) {
	return VMClassGetVBDsMockedCallback(sessionID, self)
}
// Get the VBDs field of the given VM.
func (_class VMClass) GetVBDs(sessionID SessionRef, self VMRef) (_retval []VBDRef, _err error) {
	if IsMock {
		return _class.GetVBDsMock(sessionID, self)
	}	
	_method := "VM.get_VBDs"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVBDRefSetToGo(_method + " -> ", _result.Value)
	return
}


func VMClassGetVIFsMockDefault(sessionID SessionRef, self VMRef) (_retval []VIFRef, _err error) {
	log.Println("VM.GetVIFs not mocked")
	_err = errors.New("VM.GetVIFs not mocked")
	return
}

var VMClassGetVIFsMockedCallback = VMClassGetVIFsMockDefault

func (_class VMClass) GetVIFsMock(sessionID SessionRef, self VMRef) (_retval []VIFRef, _err error) {
	return VMClassGetVIFsMockedCallback(sessionID, self)
}
// Get the VIFs field of the given VM.
func (_class VMClass) GetVIFs(sessionID SessionRef, self VMRef) (_retval []VIFRef, _err error) {
	if IsMock {
		return _class.GetVIFsMock(sessionID, self)
	}	
	_method := "VM.get_VIFs"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVIFRefSetToGo(_method + " -> ", _result.Value)
	return
}


func VMClassGetConsolesMockDefault(sessionID SessionRef, self VMRef) (_retval []ConsoleRef, _err error) {
	log.Println("VM.GetConsoles not mocked")
	_err = errors.New("VM.GetConsoles not mocked")
	return
}

var VMClassGetConsolesMockedCallback = VMClassGetConsolesMockDefault

func (_class VMClass) GetConsolesMock(sessionID SessionRef, self VMRef) (_retval []ConsoleRef, _err error) {
	return VMClassGetConsolesMockedCallback(sessionID, self)
}
// Get the consoles field of the given VM.
func (_class VMClass) GetConsoles(sessionID SessionRef, self VMRef) (_retval []ConsoleRef, _err error) {
	if IsMock {
		return _class.GetConsolesMock(sessionID, self)
	}	
	_method := "VM.get_consoles"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertConsoleRefSetToGo(_method + " -> ", _result.Value)
	return
}


func VMClassGetActionsAfterCrashMockDefault(sessionID SessionRef, self VMRef) (_retval OnCrashBehaviour, _err error) {
	log.Println("VM.GetActionsAfterCrash not mocked")
	_err = errors.New("VM.GetActionsAfterCrash not mocked")
	return
}

var VMClassGetActionsAfterCrashMockedCallback = VMClassGetActionsAfterCrashMockDefault

func (_class VMClass) GetActionsAfterCrashMock(sessionID SessionRef, self VMRef) (_retval OnCrashBehaviour, _err error) {
	return VMClassGetActionsAfterCrashMockedCallback(sessionID, self)
}
// Get the actions/after_crash field of the given VM.
func (_class VMClass) GetActionsAfterCrash(sessionID SessionRef, self VMRef) (_retval OnCrashBehaviour, _err error) {
	if IsMock {
		return _class.GetActionsAfterCrashMock(sessionID, self)
	}	
	_method := "VM.get_actions_after_crash"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertEnumOnCrashBehaviourToGo(_method + " -> ", _result.Value)
	return
}


func VMClassGetActionsAfterRebootMockDefault(sessionID SessionRef, self VMRef) (_retval OnNormalExit, _err error) {
	log.Println("VM.GetActionsAfterReboot not mocked")
	_err = errors.New("VM.GetActionsAfterReboot not mocked")
	return
}

var VMClassGetActionsAfterRebootMockedCallback = VMClassGetActionsAfterRebootMockDefault

func (_class VMClass) GetActionsAfterRebootMock(sessionID SessionRef, self VMRef) (_retval OnNormalExit, _err error) {
	return VMClassGetActionsAfterRebootMockedCallback(sessionID, self)
}
// Get the actions/after_reboot field of the given VM.
func (_class VMClass) GetActionsAfterReboot(sessionID SessionRef, self VMRef) (_retval OnNormalExit, _err error) {
	if IsMock {
		return _class.GetActionsAfterRebootMock(sessionID, self)
	}	
	_method := "VM.get_actions_after_reboot"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertEnumOnNormalExitToGo(_method + " -> ", _result.Value)
	return
}


func VMClassGetActionsAfterShutdownMockDefault(sessionID SessionRef, self VMRef) (_retval OnNormalExit, _err error) {
	log.Println("VM.GetActionsAfterShutdown not mocked")
	_err = errors.New("VM.GetActionsAfterShutdown not mocked")
	return
}

var VMClassGetActionsAfterShutdownMockedCallback = VMClassGetActionsAfterShutdownMockDefault

func (_class VMClass) GetActionsAfterShutdownMock(sessionID SessionRef, self VMRef) (_retval OnNormalExit, _err error) {
	return VMClassGetActionsAfterShutdownMockedCallback(sessionID, self)
}
// Get the actions/after_shutdown field of the given VM.
func (_class VMClass) GetActionsAfterShutdown(sessionID SessionRef, self VMRef) (_retval OnNormalExit, _err error) {
	if IsMock {
		return _class.GetActionsAfterShutdownMock(sessionID, self)
	}	
	_method := "VM.get_actions_after_shutdown"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertEnumOnNormalExitToGo(_method + " -> ", _result.Value)
	return
}


func VMClassGetVCPUsAtStartupMockDefault(sessionID SessionRef, self VMRef) (_retval int, _err error) {
	log.Println("VM.GetVCPUsAtStartup not mocked")
	_err = errors.New("VM.GetVCPUsAtStartup not mocked")
	return
}

var VMClassGetVCPUsAtStartupMockedCallback = VMClassGetVCPUsAtStartupMockDefault

func (_class VMClass) GetVCPUsAtStartupMock(sessionID SessionRef, self VMRef) (_retval int, _err error) {
	return VMClassGetVCPUsAtStartupMockedCallback(sessionID, self)
}
// Get the VCPUs/at_startup field of the given VM.
func (_class VMClass) GetVCPUsAtStartup(sessionID SessionRef, self VMRef) (_retval int, _err error) {
	if IsMock {
		return _class.GetVCPUsAtStartupMock(sessionID, self)
	}	
	_method := "VM.get_VCPUs_at_startup"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertIntToGo(_method + " -> ", _result.Value)
	return
}


func VMClassGetVCPUsMaxMockDefault(sessionID SessionRef, self VMRef) (_retval int, _err error) {
	log.Println("VM.GetVCPUsMax not mocked")
	_err = errors.New("VM.GetVCPUsMax not mocked")
	return
}

var VMClassGetVCPUsMaxMockedCallback = VMClassGetVCPUsMaxMockDefault

func (_class VMClass) GetVCPUsMaxMock(sessionID SessionRef, self VMRef) (_retval int, _err error) {
	return VMClassGetVCPUsMaxMockedCallback(sessionID, self)
}
// Get the VCPUs/max field of the given VM.
func (_class VMClass) GetVCPUsMax(sessionID SessionRef, self VMRef) (_retval int, _err error) {
	if IsMock {
		return _class.GetVCPUsMaxMock(sessionID, self)
	}	
	_method := "VM.get_VCPUs_max"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertIntToGo(_method + " -> ", _result.Value)
	return
}


func VMClassGetVCPUsParamsMockDefault(sessionID SessionRef, self VMRef) (_retval map[string]string, _err error) {
	log.Println("VM.GetVCPUsParams not mocked")
	_err = errors.New("VM.GetVCPUsParams not mocked")
	return
}

var VMClassGetVCPUsParamsMockedCallback = VMClassGetVCPUsParamsMockDefault

func (_class VMClass) GetVCPUsParamsMock(sessionID SessionRef, self VMRef) (_retval map[string]string, _err error) {
	return VMClassGetVCPUsParamsMockedCallback(sessionID, self)
}
// Get the VCPUs/params field of the given VM.
func (_class VMClass) GetVCPUsParams(sessionID SessionRef, self VMRef) (_retval map[string]string, _err error) {
	if IsMock {
		return _class.GetVCPUsParamsMock(sessionID, self)
	}	
	_method := "VM.get_VCPUs_params"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringToStringMapToGo(_method + " -> ", _result.Value)
	return
}


func VMClassGetMemoryStaticMinMockDefault(sessionID SessionRef, self VMRef) (_retval int, _err error) {
	log.Println("VM.GetMemoryStaticMin not mocked")
	_err = errors.New("VM.GetMemoryStaticMin not mocked")
	return
}

var VMClassGetMemoryStaticMinMockedCallback = VMClassGetMemoryStaticMinMockDefault

func (_class VMClass) GetMemoryStaticMinMock(sessionID SessionRef, self VMRef) (_retval int, _err error) {
	return VMClassGetMemoryStaticMinMockedCallback(sessionID, self)
}
// Get the memory/static_min field of the given VM.
func (_class VMClass) GetMemoryStaticMin(sessionID SessionRef, self VMRef) (_retval int, _err error) {
	if IsMock {
		return _class.GetMemoryStaticMinMock(sessionID, self)
	}	
	_method := "VM.get_memory_static_min"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertIntToGo(_method + " -> ", _result.Value)
	return
}


func VMClassGetMemoryDynamicMinMockDefault(sessionID SessionRef, self VMRef) (_retval int, _err error) {
	log.Println("VM.GetMemoryDynamicMin not mocked")
	_err = errors.New("VM.GetMemoryDynamicMin not mocked")
	return
}

var VMClassGetMemoryDynamicMinMockedCallback = VMClassGetMemoryDynamicMinMockDefault

func (_class VMClass) GetMemoryDynamicMinMock(sessionID SessionRef, self VMRef) (_retval int, _err error) {
	return VMClassGetMemoryDynamicMinMockedCallback(sessionID, self)
}
// Get the memory/dynamic_min field of the given VM.
func (_class VMClass) GetMemoryDynamicMin(sessionID SessionRef, self VMRef) (_retval int, _err error) {
	if IsMock {
		return _class.GetMemoryDynamicMinMock(sessionID, self)
	}	
	_method := "VM.get_memory_dynamic_min"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertIntToGo(_method + " -> ", _result.Value)
	return
}


func VMClassGetMemoryDynamicMaxMockDefault(sessionID SessionRef, self VMRef) (_retval int, _err error) {
	log.Println("VM.GetMemoryDynamicMax not mocked")
	_err = errors.New("VM.GetMemoryDynamicMax not mocked")
	return
}

var VMClassGetMemoryDynamicMaxMockedCallback = VMClassGetMemoryDynamicMaxMockDefault

func (_class VMClass) GetMemoryDynamicMaxMock(sessionID SessionRef, self VMRef) (_retval int, _err error) {
	return VMClassGetMemoryDynamicMaxMockedCallback(sessionID, self)
}
// Get the memory/dynamic_max field of the given VM.
func (_class VMClass) GetMemoryDynamicMax(sessionID SessionRef, self VMRef) (_retval int, _err error) {
	if IsMock {
		return _class.GetMemoryDynamicMaxMock(sessionID, self)
	}	
	_method := "VM.get_memory_dynamic_max"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertIntToGo(_method + " -> ", _result.Value)
	return
}


func VMClassGetMemoryStaticMaxMockDefault(sessionID SessionRef, self VMRef) (_retval int, _err error) {
	log.Println("VM.GetMemoryStaticMax not mocked")
	_err = errors.New("VM.GetMemoryStaticMax not mocked")
	return
}

var VMClassGetMemoryStaticMaxMockedCallback = VMClassGetMemoryStaticMaxMockDefault

func (_class VMClass) GetMemoryStaticMaxMock(sessionID SessionRef, self VMRef) (_retval int, _err error) {
	return VMClassGetMemoryStaticMaxMockedCallback(sessionID, self)
}
// Get the memory/static_max field of the given VM.
func (_class VMClass) GetMemoryStaticMax(sessionID SessionRef, self VMRef) (_retval int, _err error) {
	if IsMock {
		return _class.GetMemoryStaticMaxMock(sessionID, self)
	}	
	_method := "VM.get_memory_static_max"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertIntToGo(_method + " -> ", _result.Value)
	return
}


func VMClassGetMemoryTargetMockDefault(sessionID SessionRef, self VMRef) (_retval int, _err error) {
	log.Println("VM.GetMemoryTarget not mocked")
	_err = errors.New("VM.GetMemoryTarget not mocked")
	return
}

var VMClassGetMemoryTargetMockedCallback = VMClassGetMemoryTargetMockDefault

func (_class VMClass) GetMemoryTargetMock(sessionID SessionRef, self VMRef) (_retval int, _err error) {
	return VMClassGetMemoryTargetMockedCallback(sessionID, self)
}
// Get the memory/target field of the given VM.
func (_class VMClass) GetMemoryTarget(sessionID SessionRef, self VMRef) (_retval int, _err error) {
	if IsMock {
		return _class.GetMemoryTargetMock(sessionID, self)
	}	
	_method := "VM.get_memory_target"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertIntToGo(_method + " -> ", _result.Value)
	return
}


func VMClassGetMemoryOverheadMockDefault(sessionID SessionRef, self VMRef) (_retval int, _err error) {
	log.Println("VM.GetMemoryOverhead not mocked")
	_err = errors.New("VM.GetMemoryOverhead not mocked")
	return
}

var VMClassGetMemoryOverheadMockedCallback = VMClassGetMemoryOverheadMockDefault

func (_class VMClass) GetMemoryOverheadMock(sessionID SessionRef, self VMRef) (_retval int, _err error) {
	return VMClassGetMemoryOverheadMockedCallback(sessionID, self)
}
// Get the memory/overhead field of the given VM.
func (_class VMClass) GetMemoryOverhead(sessionID SessionRef, self VMRef) (_retval int, _err error) {
	if IsMock {
		return _class.GetMemoryOverheadMock(sessionID, self)
	}	
	_method := "VM.get_memory_overhead"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertIntToGo(_method + " -> ", _result.Value)
	return
}


func VMClassGetAffinityMockDefault(sessionID SessionRef, self VMRef) (_retval HostRef, _err error) {
	log.Println("VM.GetAffinity not mocked")
	_err = errors.New("VM.GetAffinity not mocked")
	return
}

var VMClassGetAffinityMockedCallback = VMClassGetAffinityMockDefault

func (_class VMClass) GetAffinityMock(sessionID SessionRef, self VMRef) (_retval HostRef, _err error) {
	return VMClassGetAffinityMockedCallback(sessionID, self)
}
// Get the affinity field of the given VM.
func (_class VMClass) GetAffinity(sessionID SessionRef, self VMRef) (_retval HostRef, _err error) {
	if IsMock {
		return _class.GetAffinityMock(sessionID, self)
	}	
	_method := "VM.get_affinity"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertHostRefToGo(_method + " -> ", _result.Value)
	return
}


func VMClassGetResidentOnMockDefault(sessionID SessionRef, self VMRef) (_retval HostRef, _err error) {
	log.Println("VM.GetResidentOn not mocked")
	_err = errors.New("VM.GetResidentOn not mocked")
	return
}

var VMClassGetResidentOnMockedCallback = VMClassGetResidentOnMockDefault

func (_class VMClass) GetResidentOnMock(sessionID SessionRef, self VMRef) (_retval HostRef, _err error) {
	return VMClassGetResidentOnMockedCallback(sessionID, self)
}
// Get the resident_on field of the given VM.
func (_class VMClass) GetResidentOn(sessionID SessionRef, self VMRef) (_retval HostRef, _err error) {
	if IsMock {
		return _class.GetResidentOnMock(sessionID, self)
	}	
	_method := "VM.get_resident_on"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertHostRefToGo(_method + " -> ", _result.Value)
	return
}


func VMClassGetSuspendVDIMockDefault(sessionID SessionRef, self VMRef) (_retval VDIRef, _err error) {
	log.Println("VM.GetSuspendVDI not mocked")
	_err = errors.New("VM.GetSuspendVDI not mocked")
	return
}

var VMClassGetSuspendVDIMockedCallback = VMClassGetSuspendVDIMockDefault

func (_class VMClass) GetSuspendVDIMock(sessionID SessionRef, self VMRef) (_retval VDIRef, _err error) {
	return VMClassGetSuspendVDIMockedCallback(sessionID, self)
}
// Get the suspend_VDI field of the given VM.
func (_class VMClass) GetSuspendVDI(sessionID SessionRef, self VMRef) (_retval VDIRef, _err error) {
	if IsMock {
		return _class.GetSuspendVDIMock(sessionID, self)
	}	
	_method := "VM.get_suspend_VDI"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVDIRefToGo(_method + " -> ", _result.Value)
	return
}


func VMClassGetIsDefaultTemplateMockDefault(sessionID SessionRef, self VMRef) (_retval bool, _err error) {
	log.Println("VM.GetIsDefaultTemplate not mocked")
	_err = errors.New("VM.GetIsDefaultTemplate not mocked")
	return
}

var VMClassGetIsDefaultTemplateMockedCallback = VMClassGetIsDefaultTemplateMockDefault

func (_class VMClass) GetIsDefaultTemplateMock(sessionID SessionRef, self VMRef) (_retval bool, _err error) {
	return VMClassGetIsDefaultTemplateMockedCallback(sessionID, self)
}
// Get the is_default_template field of the given VM.
func (_class VMClass) GetIsDefaultTemplate(sessionID SessionRef, self VMRef) (_retval bool, _err error) {
	if IsMock {
		return _class.GetIsDefaultTemplateMock(sessionID, self)
	}	
	_method := "VM.get_is_default_template"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertBoolToGo(_method + " -> ", _result.Value)
	return
}


func VMClassGetIsATemplateMockDefault(sessionID SessionRef, self VMRef) (_retval bool, _err error) {
	log.Println("VM.GetIsATemplate not mocked")
	_err = errors.New("VM.GetIsATemplate not mocked")
	return
}

var VMClassGetIsATemplateMockedCallback = VMClassGetIsATemplateMockDefault

func (_class VMClass) GetIsATemplateMock(sessionID SessionRef, self VMRef) (_retval bool, _err error) {
	return VMClassGetIsATemplateMockedCallback(sessionID, self)
}
// Get the is_a_template field of the given VM.
func (_class VMClass) GetIsATemplate(sessionID SessionRef, self VMRef) (_retval bool, _err error) {
	if IsMock {
		return _class.GetIsATemplateMock(sessionID, self)
	}	
	_method := "VM.get_is_a_template"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertBoolToGo(_method + " -> ", _result.Value)
	return
}


func VMClassGetUserVersionMockDefault(sessionID SessionRef, self VMRef) (_retval int, _err error) {
	log.Println("VM.GetUserVersion not mocked")
	_err = errors.New("VM.GetUserVersion not mocked")
	return
}

var VMClassGetUserVersionMockedCallback = VMClassGetUserVersionMockDefault

func (_class VMClass) GetUserVersionMock(sessionID SessionRef, self VMRef) (_retval int, _err error) {
	return VMClassGetUserVersionMockedCallback(sessionID, self)
}
// Get the user_version field of the given VM.
func (_class VMClass) GetUserVersion(sessionID SessionRef, self VMRef) (_retval int, _err error) {
	if IsMock {
		return _class.GetUserVersionMock(sessionID, self)
	}	
	_method := "VM.get_user_version"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertIntToGo(_method + " -> ", _result.Value)
	return
}


func VMClassGetNameDescriptionMockDefault(sessionID SessionRef, self VMRef) (_retval string, _err error) {
	log.Println("VM.GetNameDescription not mocked")
	_err = errors.New("VM.GetNameDescription not mocked")
	return
}

var VMClassGetNameDescriptionMockedCallback = VMClassGetNameDescriptionMockDefault

func (_class VMClass) GetNameDescriptionMock(sessionID SessionRef, self VMRef) (_retval string, _err error) {
	return VMClassGetNameDescriptionMockedCallback(sessionID, self)
}
// Get the name/description field of the given VM.
func (_class VMClass) GetNameDescription(sessionID SessionRef, self VMRef) (_retval string, _err error) {
	if IsMock {
		return _class.GetNameDescriptionMock(sessionID, self)
	}	
	_method := "VM.get_name_description"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringToGo(_method + " -> ", _result.Value)
	return
}


func VMClassGetNameLabelMockDefault(sessionID SessionRef, self VMRef) (_retval string, _err error) {
	log.Println("VM.GetNameLabel not mocked")
	_err = errors.New("VM.GetNameLabel not mocked")
	return
}

var VMClassGetNameLabelMockedCallback = VMClassGetNameLabelMockDefault

func (_class VMClass) GetNameLabelMock(sessionID SessionRef, self VMRef) (_retval string, _err error) {
	return VMClassGetNameLabelMockedCallback(sessionID, self)
}
// Get the name/label field of the given VM.
func (_class VMClass) GetNameLabel(sessionID SessionRef, self VMRef) (_retval string, _err error) {
	if IsMock {
		return _class.GetNameLabelMock(sessionID, self)
	}	
	_method := "VM.get_name_label"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringToGo(_method + " -> ", _result.Value)
	return
}


func VMClassGetPowerStateMockDefault(sessionID SessionRef, self VMRef) (_retval VMPowerState, _err error) {
	log.Println("VM.GetPowerState not mocked")
	_err = errors.New("VM.GetPowerState not mocked")
	return
}

var VMClassGetPowerStateMockedCallback = VMClassGetPowerStateMockDefault

func (_class VMClass) GetPowerStateMock(sessionID SessionRef, self VMRef) (_retval VMPowerState, _err error) {
	return VMClassGetPowerStateMockedCallback(sessionID, self)
}
// Get the power_state field of the given VM.
func (_class VMClass) GetPowerState(sessionID SessionRef, self VMRef) (_retval VMPowerState, _err error) {
	if IsMock {
		return _class.GetPowerStateMock(sessionID, self)
	}	
	_method := "VM.get_power_state"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertEnumVMPowerStateToGo(_method + " -> ", _result.Value)
	return
}


func VMClassGetCurrentOperationsMockDefault(sessionID SessionRef, self VMRef) (_retval map[string]VMOperations, _err error) {
	log.Println("VM.GetCurrentOperations not mocked")
	_err = errors.New("VM.GetCurrentOperations not mocked")
	return
}

var VMClassGetCurrentOperationsMockedCallback = VMClassGetCurrentOperationsMockDefault

func (_class VMClass) GetCurrentOperationsMock(sessionID SessionRef, self VMRef) (_retval map[string]VMOperations, _err error) {
	return VMClassGetCurrentOperationsMockedCallback(sessionID, self)
}
// Get the current_operations field of the given VM.
func (_class VMClass) GetCurrentOperations(sessionID SessionRef, self VMRef) (_retval map[string]VMOperations, _err error) {
	if IsMock {
		return _class.GetCurrentOperationsMock(sessionID, self)
	}	
	_method := "VM.get_current_operations"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringToEnumVMOperationsMapToGo(_method + " -> ", _result.Value)
	return
}


func VMClassGetAllowedOperationsMockDefault(sessionID SessionRef, self VMRef) (_retval []VMOperations, _err error) {
	log.Println("VM.GetAllowedOperations not mocked")
	_err = errors.New("VM.GetAllowedOperations not mocked")
	return
}

var VMClassGetAllowedOperationsMockedCallback = VMClassGetAllowedOperationsMockDefault

func (_class VMClass) GetAllowedOperationsMock(sessionID SessionRef, self VMRef) (_retval []VMOperations, _err error) {
	return VMClassGetAllowedOperationsMockedCallback(sessionID, self)
}
// Get the allowed_operations field of the given VM.
func (_class VMClass) GetAllowedOperations(sessionID SessionRef, self VMRef) (_retval []VMOperations, _err error) {
	if IsMock {
		return _class.GetAllowedOperationsMock(sessionID, self)
	}	
	_method := "VM.get_allowed_operations"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertEnumVMOperationsSetToGo(_method + " -> ", _result.Value)
	return
}


func VMClassGetUUIDMockDefault(sessionID SessionRef, self VMRef) (_retval string, _err error) {
	log.Println("VM.GetUUID not mocked")
	_err = errors.New("VM.GetUUID not mocked")
	return
}

var VMClassGetUUIDMockedCallback = VMClassGetUUIDMockDefault

func (_class VMClass) GetUUIDMock(sessionID SessionRef, self VMRef) (_retval string, _err error) {
	return VMClassGetUUIDMockedCallback(sessionID, self)
}
// Get the uuid field of the given VM.
func (_class VMClass) GetUUID(sessionID SessionRef, self VMRef) (_retval string, _err error) {
	if IsMock {
		return _class.GetUUIDMock(sessionID, self)
	}	
	_method := "VM.get_uuid"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringToGo(_method + " -> ", _result.Value)
	return
}


func VMClassGetByNameLabelMockDefault(sessionID SessionRef, label string) (_retval []VMRef, _err error) {
	log.Println("VM.GetByNameLabel not mocked")
	_err = errors.New("VM.GetByNameLabel not mocked")
	return
}

var VMClassGetByNameLabelMockedCallback = VMClassGetByNameLabelMockDefault

func (_class VMClass) GetByNameLabelMock(sessionID SessionRef, label string) (_retval []VMRef, _err error) {
	return VMClassGetByNameLabelMockedCallback(sessionID, label)
}
// Get all the VM instances with the given label.
func (_class VMClass) GetByNameLabel(sessionID SessionRef, label string) (_retval []VMRef, _err error) {
	if IsMock {
		return _class.GetByNameLabelMock(sessionID, label)
	}	
	_method := "VM.get_by_name_label"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_labelArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "label"), label)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _labelArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVMRefSetToGo(_method + " -> ", _result.Value)
	return
}


func VMClassDestroyMockDefault(sessionID SessionRef, self VMRef) (_err error) {
	log.Println("VM.Destroy not mocked")
	_err = errors.New("VM.Destroy not mocked")
	return
}

var VMClassDestroyMockedCallback = VMClassDestroyMockDefault

func (_class VMClass) DestroyMock(sessionID SessionRef, self VMRef) (_err error) {
	return VMClassDestroyMockedCallback(sessionID, self)
}
// Destroy the specified VM.  The VM is completely removed from the system.  This function can only be called when the VM is in the Halted State.
func (_class VMClass) Destroy(sessionID SessionRef, self VMRef) (_err error) {
	if IsMock {
		return _class.DestroyMock(sessionID, self)
	}	
	_method := "VM.destroy"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg)
	return
}


func VMClassCreateMockDefault(sessionID SessionRef, args VMRecord) (_retval VMRef, _err error) {
	log.Println("VM.Create not mocked")
	_err = errors.New("VM.Create not mocked")
	return
}

var VMClassCreateMockedCallback = VMClassCreateMockDefault

func (_class VMClass) CreateMock(sessionID SessionRef, args VMRecord) (_retval VMRef, _err error) {
	return VMClassCreateMockedCallback(sessionID, args)
}
// NOT RECOMMENDED! VM.clone or VM.copy (or VM.import) is a better choice in almost all situations. The standard way to obtain a new VM is to call VM.clone on a template VM, then call VM.provision on the new clone. Caution: if VM.create is used and then the new VM is attached to a virtual disc that has an operating system already installed, then there is no guarantee that the operating system will boot and run. Any software that calls VM.create on a future version of this API may fail or give unexpected results. For example this could happen if an additional parameter were added to VM.create. VM.create is intended only for use in the automatic creation of the system VM templates. It creates a new VM instance, and returns its handle.
// The constructor args are: name_label, name_description, user_version*, is_a_template*, affinity*, memory_target, memory_static_max*, memory_dynamic_max*, memory_dynamic_min*, memory_static_min*, VCPUs_params*, VCPUs_max*, VCPUs_at_startup*, actions_after_shutdown*, actions_after_reboot*, actions_after_crash*, PV_bootloader*, PV_kernel*, PV_ramdisk*, PV_args*, PV_bootloader_args*, PV_legacy_args*, HVM_boot_policy*, HVM_boot_params*, HVM_shadow_multiplier, platform*, PCI_bus*, other_config*, recommendations*, xenstore_data, ha_always_run, ha_restart_priority, tags, blocked_operations, protection_policy, is_snapshot_from_vmpp, snapshot_schedule, is_vmss_snapshot, appliance, start_delay, shutdown_delay, order, suspend_SR, version, generation_id, hardware_platform_version, has_vendor_device, reference_label (* = non-optional).
func (_class VMClass) Create(sessionID SessionRef, args VMRecord) (_retval VMRef, _err error) {
	if IsMock {
		return _class.CreateMock(sessionID, args)
	}	
	_method := "VM.create"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_argsArg, _err := convertVMRecordToXen(fmt.Sprintf("%s(%s)", _method, "args"), args)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _argsArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVMRefToGo(_method + " -> ", _result.Value)
	return
}


func VMClassGetByUUIDMockDefault(sessionID SessionRef, uuid string) (_retval VMRef, _err error) {
	log.Println("VM.GetByUUID not mocked")
	_err = errors.New("VM.GetByUUID not mocked")
	return
}

var VMClassGetByUUIDMockedCallback = VMClassGetByUUIDMockDefault

func (_class VMClass) GetByUUIDMock(sessionID SessionRef, uuid string) (_retval VMRef, _err error) {
	return VMClassGetByUUIDMockedCallback(sessionID, uuid)
}
// Get a reference to the VM instance with the specified UUID.
func (_class VMClass) GetByUUID(sessionID SessionRef, uuid string) (_retval VMRef, _err error) {
	if IsMock {
		return _class.GetByUUIDMock(sessionID, uuid)
	}	
	_method := "VM.get_by_uuid"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_uuidArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "uuid"), uuid)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _uuidArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVMRefToGo(_method + " -> ", _result.Value)
	return
}


func VMClassGetRecordMockDefault(sessionID SessionRef, self VMRef) (_retval VMRecord, _err error) {
	log.Println("VM.GetRecord not mocked")
	_err = errors.New("VM.GetRecord not mocked")
	return
}

var VMClassGetRecordMockedCallback = VMClassGetRecordMockDefault

func (_class VMClass) GetRecordMock(sessionID SessionRef, self VMRef) (_retval VMRecord, _err error) {
	return VMClassGetRecordMockedCallback(sessionID, self)
}
// Get a record containing the current state of the given VM.
func (_class VMClass) GetRecord(sessionID SessionRef, self VMRef) (_retval VMRecord, _err error) {
	if IsMock {
		return _class.GetRecordMock(sessionID, self)
	}	
	_method := "VM.get_record"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVMRecordToGo(_method + " -> ", _result.Value)
	return
}
