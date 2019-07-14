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

type HostAllowedOperations string

const (
  // Indicates this host is able to provision another VM
	HostAllowedOperationsProvision HostAllowedOperations = "provision"
  // Indicates this host is evacuating
	HostAllowedOperationsEvacuate HostAllowedOperations = "evacuate"
  // Indicates this host is in the process of shutting itself down
	HostAllowedOperationsShutdown HostAllowedOperations = "shutdown"
  // Indicates this host is in the process of rebooting
	HostAllowedOperationsReboot HostAllowedOperations = "reboot"
  // Indicates this host is in the process of being powered on
	HostAllowedOperationsPowerOn HostAllowedOperations = "power_on"
  // This host is starting a VM
	HostAllowedOperationsVMStart HostAllowedOperations = "vm_start"
  // This host is resuming a VM
	HostAllowedOperationsVMResume HostAllowedOperations = "vm_resume"
  // This host is the migration target of a VM
	HostAllowedOperationsVMMigrate HostAllowedOperations = "vm_migrate"
)

type HostDisplay string

const (
  // This host is outputting its console to a physical display device
	HostDisplayEnabled HostDisplay = "enabled"
  // The host will stop outputting its console to a physical display device on next boot
	HostDisplayDisableOnReboot HostDisplay = "disable_on_reboot"
  // This host is not outputting its console to a physical display device
	HostDisplayDisabled HostDisplay = "disabled"
  // The host will start outputting its console to a physical display device on next boot
	HostDisplayEnableOnReboot HostDisplay = "enable_on_reboot"
)

type HostRecord struct {
  // Unique identifier/object reference
	UUID string
  // a human-readable name
	NameLabel string
  // a notes field containing human-readable description
	NameDescription string
  // Virtualization memory overhead (bytes).
	MemoryOverhead int
  // list of the operations allowed in this state. This list is advisory only and the server state may have changed by the time this field is read by a client.
	AllowedOperations []HostAllowedOperations
  // links each of the running tasks using this object (by reference) to a current_operation enum which describes the nature of the task.
	CurrentOperations map[string]HostAllowedOperations
  // major version number
	APIVersionMajor int
  // minor version number
	APIVersionMinor int
  // identification of vendor
	APIVersionVendor string
  // details of vendor implementation
	APIVersionVendorImplementation map[string]string
  // True if the host is currently enabled
	Enabled bool
  // version strings
	SoftwareVersion map[string]string
  // additional configuration
	OtherConfig map[string]string
  // Xen capabilities
	Capabilities []string
  // The CPU configuration on this host.  May contain keys such as "nr_nodes", "sockets_per_node", "cores_per_socket", or "threads_per_core"
	CPUConfiguration map[string]string
  // Scheduler policy currently in force on this host
	SchedPolicy string
  // a list of the bootloaders installed on the machine
	SupportedBootloaders []string
  // list of VMs currently resident on host
	ResidentVMs []VMRef
  // logging configuration
	Logging map[string]string
  // physical network interfaces
	PIFs []PIFRef
  // The SR in which VDIs for suspend images are created
	SuspendImageSr SRRef
  // The SR in which VDIs for crash dumps are created
	CrashDumpSr SRRef
  // Set of host crash dumps
	Crashdumps []HostCrashdumpRef
  // Set of host patches
	Patches []HostPatchRef
  // Set of updates
	Updates []PoolUpdateRef
  // physical blockdevices
	PBDs []PBDRef
  // The physical CPUs on this host
	HostCPUs []HostCPURef
  // Details about the physical CPUs on this host
	CPUInfo map[string]string
  // The hostname of this host
	Hostname string
  // The address by which this host can be contacted from any other host in the pool
	Address string
  // metrics associated with this host
	Metrics HostMetricsRef
  // State of the current license
	LicenseParams map[string]string
  // The set of statefiles accessible from this host
	HaStatefiles []string
  // The set of hosts visible via the network from this host
	HaNetworkPeers []string
  // Binary blobs associated with this host
	Blobs map[string]BlobRef
  // user-specified tags for categorization purposes
	Tags []string
  // type of external authentication service configured; empty if none configured.
	ExternalAuthType string
  // name of external authentication service configured; empty if none configured.
	ExternalAuthServiceName string
  // configuration specific to external authentication service
	ExternalAuthConfiguration map[string]string
  // Product edition
	Edition string
  // Contact information of the license server
	LicenseServer map[string]string
  // BIOS strings
	BiosStrings map[string]string
  // The power on mode
	PowerOnMode string
  // The power on config
	PowerOnConfig map[string]string
  // The SR that is used as a local cache
	LocalCacheSr SRRef
  // Information about chipset features
	ChipsetInfo map[string]string
  // List of PCI devices in the host
	PCIs []PCIRef
  // List of physical GPUs in the host
	PGPUs []PGPURef
  // List of physical USBs in the host
	PUSBs []PUSBRef
  // Allow SSLv3 protocol and ciphersuites as used by older XenServers. This controls both incoming and outgoing connections. When this is set to a different value, the host immediately restarts its SSL/TLS listening service; typically this takes less than a second but existing connections to it will be broken. XenAPI login sessions will remain valid.
	SslLegacy bool
  // VCPUs params to apply to all resident guests
	GuestVCPUsParams map[string]string
  // indicates whether the host is configured to output its console to a physical display device
	Display HostDisplay
  // The set of versions of the virtual hardware platform that the host can offer to its guests
	VirtualHardwarePlatformVersions []int
  // The control domain (domain 0)
	ControlDomain VMRef
  // List of updates which require reboot
	UpdatesRequiringReboot []PoolUpdateRef
  // List of features available on this host
	Features []FeatureRef
}

type HostRef string

// A physical host
type HostClass struct {
	client *Client
}


var HostClass_GetAllRecordsMockedCallback = func (sessionID SessionRef) (_retval map[HostRef]HostRecord, _err error) {
	log.Println("Host.GetAllRecords not mocked")
	_err = errors.New("Host.GetAllRecords not mocked")
	return
}

func (_class HostClass) GetAllRecordsMock(sessionID SessionRef) (_retval map[HostRef]HostRecord, _err error) {
	return HostClass_GetAllRecordsMockedCallback(sessionID)
}
// Return a map of host references to host records for all hosts known to the system.
func (_class HostClass) GetAllRecords(sessionID SessionRef) (_retval map[HostRef]HostRecord, _err error) {
	if (IsMock) {
		return _class.GetAllRecordsMock(sessionID)
	}	
	_method := "host.get_all_records"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertHostRefToHostRecordMapToGo(_method + " -> ", _result.Value)
	return
}


var HostClass_GetAllMockedCallback = func (sessionID SessionRef) (_retval []HostRef, _err error) {
	log.Println("Host.GetAll not mocked")
	_err = errors.New("Host.GetAll not mocked")
	return
}

func (_class HostClass) GetAllMock(sessionID SessionRef) (_retval []HostRef, _err error) {
	return HostClass_GetAllMockedCallback(sessionID)
}
// Return a list of all the hosts known to the system.
func (_class HostClass) GetAll(sessionID SessionRef) (_retval []HostRef, _err error) {
	if (IsMock) {
		return _class.GetAllMock(sessionID)
	}	
	_method := "host.get_all"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertHostRefSetToGo(_method + " -> ", _result.Value)
	return
}


var HostClass_SetSslLegacyMockedCallback = func (sessionID SessionRef, self HostRef, value bool) (_err error) {
	log.Println("Host.SetSslLegacy not mocked")
	_err = errors.New("Host.SetSslLegacy not mocked")
	return
}

func (_class HostClass) SetSslLegacyMock(sessionID SessionRef, self HostRef, value bool) (_err error) {
	return HostClass_SetSslLegacyMockedCallback(sessionID, self, value)
}
// Enable/disable SSLv3 for interoperability with older versions of XenServer. When this is set to a different value, the host immediately restarts its SSL/TLS listening service; typically this takes less than a second but existing connections to it will be broken. XenAPI login sessions will remain valid.
func (_class HostClass) SetSslLegacy(sessionID SessionRef, self HostRef, value bool) (_err error) {
	if (IsMock) {
		return _class.SetSslLegacyMock(sessionID, self, value)
	}	
	_method := "host.set_ssl_legacy"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var HostClass_DisableDisplayMockedCallback = func (sessionID SessionRef, host HostRef) (_retval HostDisplay, _err error) {
	log.Println("Host.DisableDisplay not mocked")
	_err = errors.New("Host.DisableDisplay not mocked")
	return
}

func (_class HostClass) DisableDisplayMock(sessionID SessionRef, host HostRef) (_retval HostDisplay, _err error) {
	return HostClass_DisableDisplayMockedCallback(sessionID, host)
}
// Disable console output to the physical display device next time this host boots
func (_class HostClass) DisableDisplay(sessionID SessionRef, host HostRef) (_retval HostDisplay, _err error) {
	if (IsMock) {
		return _class.DisableDisplayMock(sessionID, host)
	}	
	_method := "host.disable_display"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_hostArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "host"), host)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _hostArg)
	if _err != nil {
		return
	}
	_retval, _err = convertEnumHostDisplayToGo(_method + " -> ", _result.Value)
	return
}


var HostClass_EnableDisplayMockedCallback = func (sessionID SessionRef, host HostRef) (_retval HostDisplay, _err error) {
	log.Println("Host.EnableDisplay not mocked")
	_err = errors.New("Host.EnableDisplay not mocked")
	return
}

func (_class HostClass) EnableDisplayMock(sessionID SessionRef, host HostRef) (_retval HostDisplay, _err error) {
	return HostClass_EnableDisplayMockedCallback(sessionID, host)
}
// Enable console output to the physical display device next time this host boots
func (_class HostClass) EnableDisplay(sessionID SessionRef, host HostRef) (_retval HostDisplay, _err error) {
	if (IsMock) {
		return _class.EnableDisplayMock(sessionID, host)
	}	
	_method := "host.enable_display"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_hostArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "host"), host)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _hostArg)
	if _err != nil {
		return
	}
	_retval, _err = convertEnumHostDisplayToGo(_method + " -> ", _result.Value)
	return
}


var HostClass_DeclareDeadMockedCallback = func (sessionID SessionRef, host HostRef) (_err error) {
	log.Println("Host.DeclareDead not mocked")
	_err = errors.New("Host.DeclareDead not mocked")
	return
}

func (_class HostClass) DeclareDeadMock(sessionID SessionRef, host HostRef) (_err error) {
	return HostClass_DeclareDeadMockedCallback(sessionID, host)
}
// Declare that a host is dead. This is a dangerous operation, and should only be called if the administrator is absolutely sure the host is definitely dead
func (_class HostClass) DeclareDead(sessionID SessionRef, host HostRef) (_err error) {
	if (IsMock) {
		return _class.DeclareDeadMock(sessionID, host)
	}	
	_method := "host.declare_dead"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_hostArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "host"), host)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _hostArg)
	return
}


var HostClass_MigrateReceiveMockedCallback = func (sessionID SessionRef, host HostRef, network NetworkRef, options map[string]string) (_retval map[string]string, _err error) {
	log.Println("Host.MigrateReceive not mocked")
	_err = errors.New("Host.MigrateReceive not mocked")
	return
}

func (_class HostClass) MigrateReceiveMock(sessionID SessionRef, host HostRef, network NetworkRef, options map[string]string) (_retval map[string]string, _err error) {
	return HostClass_MigrateReceiveMockedCallback(sessionID, host, network, options)
}
// Prepare to receive a VM, returning a token which can be passed to VM.migrate.
func (_class HostClass) MigrateReceive(sessionID SessionRef, host HostRef, network NetworkRef, options map[string]string) (_retval map[string]string, _err error) {
	if (IsMock) {
		return _class.MigrateReceiveMock(sessionID, host, network, options)
	}	
	_method := "host.migrate_receive"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_hostArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "host"), host)
	if _err != nil {
		return
	}
	_networkArg, _err := convertNetworkRefToXen(fmt.Sprintf("%s(%s)", _method, "network"), network)
	if _err != nil {
		return
	}
	_optionsArg, _err := convertStringToStringMapToXen(fmt.Sprintf("%s(%s)", _method, "options"), options)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _hostArg, _networkArg, _optionsArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringToStringMapToGo(_method + " -> ", _result.Value)
	return
}


var HostClass_DisableLocalStorageCachingMockedCallback = func (sessionID SessionRef, host HostRef) (_err error) {
	log.Println("Host.DisableLocalStorageCaching not mocked")
	_err = errors.New("Host.DisableLocalStorageCaching not mocked")
	return
}

func (_class HostClass) DisableLocalStorageCachingMock(sessionID SessionRef, host HostRef) (_err error) {
	return HostClass_DisableLocalStorageCachingMockedCallback(sessionID, host)
}
// Disable the use of a local SR for caching purposes
func (_class HostClass) DisableLocalStorageCaching(sessionID SessionRef, host HostRef) (_err error) {
	if (IsMock) {
		return _class.DisableLocalStorageCachingMock(sessionID, host)
	}	
	_method := "host.disable_local_storage_caching"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_hostArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "host"), host)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _hostArg)
	return
}


var HostClass_EnableLocalStorageCachingMockedCallback = func (sessionID SessionRef, host HostRef, sr SRRef) (_err error) {
	log.Println("Host.EnableLocalStorageCaching not mocked")
	_err = errors.New("Host.EnableLocalStorageCaching not mocked")
	return
}

func (_class HostClass) EnableLocalStorageCachingMock(sessionID SessionRef, host HostRef, sr SRRef) (_err error) {
	return HostClass_EnableLocalStorageCachingMockedCallback(sessionID, host, sr)
}
// Enable the use of a local SR for caching purposes
func (_class HostClass) EnableLocalStorageCaching(sessionID SessionRef, host HostRef, sr SRRef) (_err error) {
	if (IsMock) {
		return _class.EnableLocalStorageCachingMock(sessionID, host, sr)
	}	
	_method := "host.enable_local_storage_caching"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_hostArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "host"), host)
	if _err != nil {
		return
	}
	_srArg, _err := convertSRRefToXen(fmt.Sprintf("%s(%s)", _method, "sr"), sr)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _hostArg, _srArg)
	return
}


var HostClass_ResetCPUFeaturesMockedCallback = func (sessionID SessionRef, host HostRef) (_err error) {
	log.Println("Host.ResetCPUFeatures not mocked")
	_err = errors.New("Host.ResetCPUFeatures not mocked")
	return
}

func (_class HostClass) ResetCPUFeaturesMock(sessionID SessionRef, host HostRef) (_err error) {
	return HostClass_ResetCPUFeaturesMockedCallback(sessionID, host)
}
// Remove the feature mask, such that after a reboot all features of the CPU are enabled.
func (_class HostClass) ResetCPUFeatures(sessionID SessionRef, host HostRef) (_err error) {
	if (IsMock) {
		return _class.ResetCPUFeaturesMock(sessionID, host)
	}	
	_method := "host.reset_cpu_features"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_hostArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "host"), host)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _hostArg)
	return
}


var HostClass_SetCPUFeaturesMockedCallback = func (sessionID SessionRef, host HostRef, features string) (_err error) {
	log.Println("Host.SetCPUFeatures not mocked")
	_err = errors.New("Host.SetCPUFeatures not mocked")
	return
}

func (_class HostClass) SetCPUFeaturesMock(sessionID SessionRef, host HostRef, features string) (_err error) {
	return HostClass_SetCPUFeaturesMockedCallback(sessionID, host, features)
}
// Set the CPU features to be used after a reboot, if the given features string is valid.
func (_class HostClass) SetCPUFeatures(sessionID SessionRef, host HostRef, features string) (_err error) {
	if (IsMock) {
		return _class.SetCPUFeaturesMock(sessionID, host, features)
	}	
	_method := "host.set_cpu_features"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_hostArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "host"), host)
	if _err != nil {
		return
	}
	_featuresArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "features"), features)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _hostArg, _featuresArg)
	return
}


var HostClass_SetPowerOnModeMockedCallback = func (sessionID SessionRef, self HostRef, powerOnMode string, powerOnConfig map[string]string) (_err error) {
	log.Println("Host.SetPowerOnMode not mocked")
	_err = errors.New("Host.SetPowerOnMode not mocked")
	return
}

func (_class HostClass) SetPowerOnModeMock(sessionID SessionRef, self HostRef, powerOnMode string, powerOnConfig map[string]string) (_err error) {
	return HostClass_SetPowerOnModeMockedCallback(sessionID, self, powerOnMode, powerOnConfig)
}
// Set the power-on-mode, host, user and password 
func (_class HostClass) SetPowerOnMode(sessionID SessionRef, self HostRef, powerOnMode string, powerOnConfig map[string]string) (_err error) {
	if (IsMock) {
		return _class.SetPowerOnModeMock(sessionID, self, powerOnMode, powerOnConfig)
	}	
	_method := "host.set_power_on_mode"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_powerOnModeArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "power_on_mode"), powerOnMode)
	if _err != nil {
		return
	}
	_powerOnConfigArg, _err := convertStringToStringMapToXen(fmt.Sprintf("%s(%s)", _method, "power_on_config"), powerOnConfig)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _powerOnModeArg, _powerOnConfigArg)
	return
}


var HostClass_RefreshPackInfoMockedCallback = func (sessionID SessionRef, host HostRef) (_err error) {
	log.Println("Host.RefreshPackInfo not mocked")
	_err = errors.New("Host.RefreshPackInfo not mocked")
	return
}

func (_class HostClass) RefreshPackInfoMock(sessionID SessionRef, host HostRef) (_err error) {
	return HostClass_RefreshPackInfoMockedCallback(sessionID, host)
}
// Refresh the list of installed Supplemental Packs.
func (_class HostClass) RefreshPackInfo(sessionID SessionRef, host HostRef) (_err error) {
	if (IsMock) {
		return _class.RefreshPackInfoMock(sessionID, host)
	}	
	_method := "host.refresh_pack_info"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_hostArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "host"), host)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _hostArg)
	return
}


var HostClass_ApplyEditionMockedCallback = func (sessionID SessionRef, host HostRef, edition string, force bool) (_err error) {
	log.Println("Host.ApplyEdition not mocked")
	_err = errors.New("Host.ApplyEdition not mocked")
	return
}

func (_class HostClass) ApplyEditionMock(sessionID SessionRef, host HostRef, edition string, force bool) (_err error) {
	return HostClass_ApplyEditionMockedCallback(sessionID, host, edition, force)
}
// Change to another edition, or reactivate the current edition after a license has expired. This may be subject to the successful checkout of an appropriate license.
func (_class HostClass) ApplyEdition(sessionID SessionRef, host HostRef, edition string, force bool) (_err error) {
	if (IsMock) {
		return _class.ApplyEditionMock(sessionID, host, edition, force)
	}	
	_method := "host.apply_edition"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_hostArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "host"), host)
	if _err != nil {
		return
	}
	_editionArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "edition"), edition)
	if _err != nil {
		return
	}
	_forceArg, _err := convertBoolToXen(fmt.Sprintf("%s(%s)", _method, "force"), force)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _hostArg, _editionArg, _forceArg)
	return
}


var HostClass_GetServerCertificateMockedCallback = func (sessionID SessionRef, host HostRef) (_retval string, _err error) {
	log.Println("Host.GetServerCertificate not mocked")
	_err = errors.New("Host.GetServerCertificate not mocked")
	return
}

func (_class HostClass) GetServerCertificateMock(sessionID SessionRef, host HostRef) (_retval string, _err error) {
	return HostClass_GetServerCertificateMockedCallback(sessionID, host)
}
// Get the installed server public TLS certificate.
func (_class HostClass) GetServerCertificate(sessionID SessionRef, host HostRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetServerCertificateMock(sessionID, host)
	}	
	_method := "host.get_server_certificate"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_hostArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "host"), host)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _hostArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringToGo(_method + " -> ", _result.Value)
	return
}


var HostClass_RetrieveWlbEvacuateRecommendationsMockedCallback = func (sessionID SessionRef, self HostRef) (_retval map[VMRef][]string, _err error) {
	log.Println("Host.RetrieveWlbEvacuateRecommendations not mocked")
	_err = errors.New("Host.RetrieveWlbEvacuateRecommendations not mocked")
	return
}

func (_class HostClass) RetrieveWlbEvacuateRecommendationsMock(sessionID SessionRef, self HostRef) (_retval map[VMRef][]string, _err error) {
	return HostClass_RetrieveWlbEvacuateRecommendationsMockedCallback(sessionID, self)
}
// Retrieves recommended host migrations to perform when evacuating the host from the wlb server. If a VM cannot be migrated from the host the reason is listed instead of a recommendation.
func (_class HostClass) RetrieveWlbEvacuateRecommendations(sessionID SessionRef, self HostRef) (_retval map[VMRef][]string, _err error) {
	if (IsMock) {
		return _class.RetrieveWlbEvacuateRecommendationsMock(sessionID, self)
	}	
	_method := "host.retrieve_wlb_evacuate_recommendations"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVMRefToStringSetMapToGo(_method + " -> ", _result.Value)
	return
}


var HostClass_DisableExternalAuthMockedCallback = func (sessionID SessionRef, host HostRef, config map[string]string) (_err error) {
	log.Println("Host.DisableExternalAuth not mocked")
	_err = errors.New("Host.DisableExternalAuth not mocked")
	return
}

func (_class HostClass) DisableExternalAuthMock(sessionID SessionRef, host HostRef, config map[string]string) (_err error) {
	return HostClass_DisableExternalAuthMockedCallback(sessionID, host, config)
}
// This call disables external authentication on the local host
func (_class HostClass) DisableExternalAuth(sessionID SessionRef, host HostRef, config map[string]string) (_err error) {
	if (IsMock) {
		return _class.DisableExternalAuthMock(sessionID, host, config)
	}	
	_method := "host.disable_external_auth"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_hostArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "host"), host)
	if _err != nil {
		return
	}
	_configArg, _err := convertStringToStringMapToXen(fmt.Sprintf("%s(%s)", _method, "config"), config)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _hostArg, _configArg)
	return
}


var HostClass_EnableExternalAuthMockedCallback = func (sessionID SessionRef, host HostRef, config map[string]string, serviceName string, authType string) (_err error) {
	log.Println("Host.EnableExternalAuth not mocked")
	_err = errors.New("Host.EnableExternalAuth not mocked")
	return
}

func (_class HostClass) EnableExternalAuthMock(sessionID SessionRef, host HostRef, config map[string]string, serviceName string, authType string) (_err error) {
	return HostClass_EnableExternalAuthMockedCallback(sessionID, host, config, serviceName, authType)
}
// This call enables external authentication on a host
func (_class HostClass) EnableExternalAuth(sessionID SessionRef, host HostRef, config map[string]string, serviceName string, authType string) (_err error) {
	if (IsMock) {
		return _class.EnableExternalAuthMock(sessionID, host, config, serviceName, authType)
	}	
	_method := "host.enable_external_auth"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_hostArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "host"), host)
	if _err != nil {
		return
	}
	_configArg, _err := convertStringToStringMapToXen(fmt.Sprintf("%s(%s)", _method, "config"), config)
	if _err != nil {
		return
	}
	_serviceNameArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "service_name"), serviceName)
	if _err != nil {
		return
	}
	_authTypeArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "auth_type"), authType)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _hostArg, _configArg, _serviceNameArg, _authTypeArg)
	return
}


var HostClass_GetServerLocaltimeMockedCallback = func (sessionID SessionRef, host HostRef) (_retval time.Time, _err error) {
	log.Println("Host.GetServerLocaltime not mocked")
	_err = errors.New("Host.GetServerLocaltime not mocked")
	return
}

func (_class HostClass) GetServerLocaltimeMock(sessionID SessionRef, host HostRef) (_retval time.Time, _err error) {
	return HostClass_GetServerLocaltimeMockedCallback(sessionID, host)
}
// This call queries the host's clock for the current time in the host's local timezone
func (_class HostClass) GetServerLocaltime(sessionID SessionRef, host HostRef) (_retval time.Time, _err error) {
	if (IsMock) {
		return _class.GetServerLocaltimeMock(sessionID, host)
	}	
	_method := "host.get_server_localtime"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_hostArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "host"), host)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _hostArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTimeToGo(_method + " -> ", _result.Value)
	return
}


var HostClass_GetServertimeMockedCallback = func (sessionID SessionRef, host HostRef) (_retval time.Time, _err error) {
	log.Println("Host.GetServertime not mocked")
	_err = errors.New("Host.GetServertime not mocked")
	return
}

func (_class HostClass) GetServertimeMock(sessionID SessionRef, host HostRef) (_retval time.Time, _err error) {
	return HostClass_GetServertimeMockedCallback(sessionID, host)
}
// This call queries the host's clock for the current time
func (_class HostClass) GetServertime(sessionID SessionRef, host HostRef) (_retval time.Time, _err error) {
	if (IsMock) {
		return _class.GetServertimeMock(sessionID, host)
	}	
	_method := "host.get_servertime"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_hostArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "host"), host)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _hostArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTimeToGo(_method + " -> ", _result.Value)
	return
}


var HostClass_CallExtensionMockedCallback = func (sessionID SessionRef, host HostRef, call string) (_retval string, _err error) {
	log.Println("Host.CallExtension not mocked")
	_err = errors.New("Host.CallExtension not mocked")
	return
}

func (_class HostClass) CallExtensionMock(sessionID SessionRef, host HostRef, call string) (_retval string, _err error) {
	return HostClass_CallExtensionMockedCallback(sessionID, host, call)
}
// Call a XenAPI extension on this host
func (_class HostClass) CallExtension(sessionID SessionRef, host HostRef, call string) (_retval string, _err error) {
	if (IsMock) {
		return _class.CallExtensionMock(sessionID, host, call)
	}	
	_method := "host.call_extension"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_hostArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "host"), host)
	if _err != nil {
		return
	}
	_callArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "call"), call)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _hostArg, _callArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringToGo(_method + " -> ", _result.Value)
	return
}


var HostClass_HasExtensionMockedCallback = func (sessionID SessionRef, host HostRef, name string) (_retval bool, _err error) {
	log.Println("Host.HasExtension not mocked")
	_err = errors.New("Host.HasExtension not mocked")
	return
}

func (_class HostClass) HasExtensionMock(sessionID SessionRef, host HostRef, name string) (_retval bool, _err error) {
	return HostClass_HasExtensionMockedCallback(sessionID, host, name)
}
// Return true if the extension is available on the host
func (_class HostClass) HasExtension(sessionID SessionRef, host HostRef, name string) (_retval bool, _err error) {
	if (IsMock) {
		return _class.HasExtensionMock(sessionID, host, name)
	}	
	_method := "host.has_extension"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_hostArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "host"), host)
	if _err != nil {
		return
	}
	_nameArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "name"), name)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _hostArg, _nameArg)
	if _err != nil {
		return
	}
	_retval, _err = convertBoolToGo(_method + " -> ", _result.Value)
	return
}


var HostClass_CallPluginMockedCallback = func (sessionID SessionRef, host HostRef, plugin string, fn string, args map[string]string) (_retval string, _err error) {
	log.Println("Host.CallPlugin not mocked")
	_err = errors.New("Host.CallPlugin not mocked")
	return
}

func (_class HostClass) CallPluginMock(sessionID SessionRef, host HostRef, plugin string, fn string, args map[string]string) (_retval string, _err error) {
	return HostClass_CallPluginMockedCallback(sessionID, host, plugin, fn, args)
}
// Call a XenAPI plugin on this host
func (_class HostClass) CallPlugin(sessionID SessionRef, host HostRef, plugin string, fn string, args map[string]string) (_retval string, _err error) {
	if (IsMock) {
		return _class.CallPluginMock(sessionID, host, plugin, fn, args)
	}	
	_method := "host.call_plugin"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_hostArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "host"), host)
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
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _hostArg, _pluginArg, _fnArg, _argsArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringToGo(_method + " -> ", _result.Value)
	return
}


var HostClass_CreateNewBlobMockedCallback = func (sessionID SessionRef, host HostRef, name string, mimeType string, public bool) (_retval BlobRef, _err error) {
	log.Println("Host.CreateNewBlob not mocked")
	_err = errors.New("Host.CreateNewBlob not mocked")
	return
}

func (_class HostClass) CreateNewBlobMock(sessionID SessionRef, host HostRef, name string, mimeType string, public bool) (_retval BlobRef, _err error) {
	return HostClass_CreateNewBlobMockedCallback(sessionID, host, name, mimeType, public)
}
// Create a placeholder for a named binary blob of data that is associated with this host
func (_class HostClass) CreateNewBlob(sessionID SessionRef, host HostRef, name string, mimeType string, public bool) (_retval BlobRef, _err error) {
	if (IsMock) {
		return _class.CreateNewBlobMock(sessionID, host, name, mimeType, public)
	}	
	_method := "host.create_new_blob"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_hostArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "host"), host)
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
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _hostArg, _nameArg, _mimeTypeArg, _publicArg)
	if _err != nil {
		return
	}
	_retval, _err = convertBlobRefToGo(_method + " -> ", _result.Value)
	return
}


var HostClass_BackupRrdsMockedCallback = func (sessionID SessionRef, host HostRef, delay float64) (_err error) {
	log.Println("Host.BackupRrds not mocked")
	_err = errors.New("Host.BackupRrds not mocked")
	return
}

func (_class HostClass) BackupRrdsMock(sessionID SessionRef, host HostRef, delay float64) (_err error) {
	return HostClass_BackupRrdsMockedCallback(sessionID, host, delay)
}
// This causes the RRDs to be backed up to the master
func (_class HostClass) BackupRrds(sessionID SessionRef, host HostRef, delay float64) (_err error) {
	if (IsMock) {
		return _class.BackupRrdsMock(sessionID, host, delay)
	}	
	_method := "host.backup_rrds"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_hostArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "host"), host)
	if _err != nil {
		return
	}
	_delayArg, _err := convertFloatToXen(fmt.Sprintf("%s(%s)", _method, "delay"), delay)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _hostArg, _delayArg)
	return
}


var HostClass_SyncDataMockedCallback = func (sessionID SessionRef, host HostRef) (_err error) {
	log.Println("Host.SyncData not mocked")
	_err = errors.New("Host.SyncData not mocked")
	return
}

func (_class HostClass) SyncDataMock(sessionID SessionRef, host HostRef) (_err error) {
	return HostClass_SyncDataMockedCallback(sessionID, host)
}
// This causes the synchronisation of the non-database data (messages, RRDs and so on) stored on the master to be synchronised with the host
func (_class HostClass) SyncData(sessionID SessionRef, host HostRef) (_err error) {
	if (IsMock) {
		return _class.SyncDataMock(sessionID, host)
	}	
	_method := "host.sync_data"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_hostArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "host"), host)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _hostArg)
	return
}


var HostClass_ComputeMemoryOverheadMockedCallback = func (sessionID SessionRef, host HostRef) (_retval int, _err error) {
	log.Println("Host.ComputeMemoryOverhead not mocked")
	_err = errors.New("Host.ComputeMemoryOverhead not mocked")
	return
}

func (_class HostClass) ComputeMemoryOverheadMock(sessionID SessionRef, host HostRef) (_retval int, _err error) {
	return HostClass_ComputeMemoryOverheadMockedCallback(sessionID, host)
}
// Computes the virtualization memory overhead of a host.
func (_class HostClass) ComputeMemoryOverhead(sessionID SessionRef, host HostRef) (_retval int, _err error) {
	if (IsMock) {
		return _class.ComputeMemoryOverheadMock(sessionID, host)
	}	
	_method := "host.compute_memory_overhead"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_hostArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "host"), host)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _hostArg)
	if _err != nil {
		return
	}
	_retval, _err = convertIntToGo(_method + " -> ", _result.Value)
	return
}


var HostClass_ComputeFreeMemoryMockedCallback = func (sessionID SessionRef, host HostRef) (_retval int, _err error) {
	log.Println("Host.ComputeFreeMemory not mocked")
	_err = errors.New("Host.ComputeFreeMemory not mocked")
	return
}

func (_class HostClass) ComputeFreeMemoryMock(sessionID SessionRef, host HostRef) (_retval int, _err error) {
	return HostClass_ComputeFreeMemoryMockedCallback(sessionID, host)
}
// Computes the amount of free memory on the host.
func (_class HostClass) ComputeFreeMemory(sessionID SessionRef, host HostRef) (_retval int, _err error) {
	if (IsMock) {
		return _class.ComputeFreeMemoryMock(sessionID, host)
	}	
	_method := "host.compute_free_memory"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_hostArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "host"), host)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _hostArg)
	if _err != nil {
		return
	}
	_retval, _err = convertIntToGo(_method + " -> ", _result.Value)
	return
}


var HostClass_SetHostnameLiveMockedCallback = func (sessionID SessionRef, host HostRef, hostname string) (_err error) {
	log.Println("Host.SetHostnameLive not mocked")
	_err = errors.New("Host.SetHostnameLive not mocked")
	return
}

func (_class HostClass) SetHostnameLiveMock(sessionID SessionRef, host HostRef, hostname string) (_err error) {
	return HostClass_SetHostnameLiveMockedCallback(sessionID, host, hostname)
}
// Sets the host name to the specified string.  Both the API and lower-level system hostname are changed immediately.
//
// Errors:
//  HOST_NAME_INVALID - The host name is invalid.
func (_class HostClass) SetHostnameLive(sessionID SessionRef, host HostRef, hostname string) (_err error) {
	if (IsMock) {
		return _class.SetHostnameLiveMock(sessionID, host, hostname)
	}	
	_method := "host.set_hostname_live"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_hostArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "host"), host)
	if _err != nil {
		return
	}
	_hostnameArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "hostname"), hostname)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _hostArg, _hostnameArg)
	return
}


var HostClass_ShutdownAgentMockedCallback = func (sessionID SessionRef) (_err error) {
	log.Println("Host.ShutdownAgent not mocked")
	_err = errors.New("Host.ShutdownAgent not mocked")
	return
}

func (_class HostClass) ShutdownAgentMock(sessionID SessionRef) (_err error) {
	return HostClass_ShutdownAgentMockedCallback(sessionID)
}
// Shuts the agent down after a 10 second pause. WARNING: this is a dangerous operation. Any operations in progress will be aborted, and unrecoverable data loss may occur. The caller is responsible for ensuring that there are no operations in progress when this method is called.
func (_class HostClass) ShutdownAgent(sessionID SessionRef) (_err error) {
	if (IsMock) {
		return _class.ShutdownAgentMock(sessionID)
	}	
	_method := "host.shutdown_agent"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg)
	return
}


var HostClass_RestartAgentMockedCallback = func (sessionID SessionRef, host HostRef) (_err error) {
	log.Println("Host.RestartAgent not mocked")
	_err = errors.New("Host.RestartAgent not mocked")
	return
}

func (_class HostClass) RestartAgentMock(sessionID SessionRef, host HostRef) (_err error) {
	return HostClass_RestartAgentMockedCallback(sessionID, host)
}
// Restarts the agent after a 10 second pause. WARNING: this is a dangerous operation. Any operations in progress will be aborted, and unrecoverable data loss may occur. The caller is responsible for ensuring that there are no operations in progress when this method is called.
func (_class HostClass) RestartAgent(sessionID SessionRef, host HostRef) (_err error) {
	if (IsMock) {
		return _class.RestartAgentMock(sessionID, host)
	}	
	_method := "host.restart_agent"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_hostArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "host"), host)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _hostArg)
	return
}


var HostClass_GetSystemStatusCapabilitiesMockedCallback = func (sessionID SessionRef, host HostRef) (_retval string, _err error) {
	log.Println("Host.GetSystemStatusCapabilities not mocked")
	_err = errors.New("Host.GetSystemStatusCapabilities not mocked")
	return
}

func (_class HostClass) GetSystemStatusCapabilitiesMock(sessionID SessionRef, host HostRef) (_retval string, _err error) {
	return HostClass_GetSystemStatusCapabilitiesMockedCallback(sessionID, host)
}
// 
func (_class HostClass) GetSystemStatusCapabilities(sessionID SessionRef, host HostRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetSystemStatusCapabilitiesMock(sessionID, host)
	}	
	_method := "host.get_system_status_capabilities"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_hostArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "host"), host)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _hostArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringToGo(_method + " -> ", _result.Value)
	return
}


var HostClass_GetManagementInterfaceMockedCallback = func (sessionID SessionRef, host HostRef) (_retval PIFRef, _err error) {
	log.Println("Host.GetManagementInterface not mocked")
	_err = errors.New("Host.GetManagementInterface not mocked")
	return
}

func (_class HostClass) GetManagementInterfaceMock(sessionID SessionRef, host HostRef) (_retval PIFRef, _err error) {
	return HostClass_GetManagementInterfaceMockedCallback(sessionID, host)
}
// Returns the management interface for the specified host
func (_class HostClass) GetManagementInterface(sessionID SessionRef, host HostRef) (_retval PIFRef, _err error) {
	if (IsMock) {
		return _class.GetManagementInterfaceMock(sessionID, host)
	}	
	_method := "host.get_management_interface"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_hostArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "host"), host)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _hostArg)
	if _err != nil {
		return
	}
	_retval, _err = convertPIFRefToGo(_method + " -> ", _result.Value)
	return
}


var HostClass_ManagementDisableMockedCallback = func (sessionID SessionRef) (_err error) {
	log.Println("Host.ManagementDisable not mocked")
	_err = errors.New("Host.ManagementDisable not mocked")
	return
}

func (_class HostClass) ManagementDisableMock(sessionID SessionRef) (_err error) {
	return HostClass_ManagementDisableMockedCallback(sessionID)
}
// Disable the management network interface
func (_class HostClass) ManagementDisable(sessionID SessionRef) (_err error) {
	if (IsMock) {
		return _class.ManagementDisableMock(sessionID)
	}	
	_method := "host.management_disable"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg)
	return
}


var HostClass_LocalManagementReconfigureMockedCallback = func (sessionID SessionRef, iface string) (_err error) {
	log.Println("Host.LocalManagementReconfigure not mocked")
	_err = errors.New("Host.LocalManagementReconfigure not mocked")
	return
}

func (_class HostClass) LocalManagementReconfigureMock(sessionID SessionRef, iface string) (_err error) {
	return HostClass_LocalManagementReconfigureMockedCallback(sessionID, iface)
}
// Reconfigure the management network interface. Should only be used if Host.management_reconfigure is impossible because the network configuration is broken.
func (_class HostClass) LocalManagementReconfigure(sessionID SessionRef, iface string) (_err error) {
	if (IsMock) {
		return _class.LocalManagementReconfigureMock(sessionID, iface)
	}	
	_method := "host.local_management_reconfigure"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_ifaceArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "interface"), iface)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _ifaceArg)
	return
}


var HostClass_ManagementReconfigureMockedCallback = func (sessionID SessionRef, pif PIFRef) (_err error) {
	log.Println("Host.ManagementReconfigure not mocked")
	_err = errors.New("Host.ManagementReconfigure not mocked")
	return
}

func (_class HostClass) ManagementReconfigureMock(sessionID SessionRef, pif PIFRef) (_err error) {
	return HostClass_ManagementReconfigureMockedCallback(sessionID, pif)
}
// Reconfigure the management network interface
func (_class HostClass) ManagementReconfigure(sessionID SessionRef, pif PIFRef) (_err error) {
	if (IsMock) {
		return _class.ManagementReconfigureMock(sessionID, pif)
	}	
	_method := "host.management_reconfigure"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_pifArg, _err := convertPIFRefToXen(fmt.Sprintf("%s(%s)", _method, "pif"), pif)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _pifArg)
	return
}


var HostClass_SyslogReconfigureMockedCallback = func (sessionID SessionRef, host HostRef) (_err error) {
	log.Println("Host.SyslogReconfigure not mocked")
	_err = errors.New("Host.SyslogReconfigure not mocked")
	return
}

func (_class HostClass) SyslogReconfigureMock(sessionID SessionRef, host HostRef) (_err error) {
	return HostClass_SyslogReconfigureMockedCallback(sessionID, host)
}
// Re-configure syslog logging
func (_class HostClass) SyslogReconfigure(sessionID SessionRef, host HostRef) (_err error) {
	if (IsMock) {
		return _class.SyslogReconfigureMock(sessionID, host)
	}	
	_method := "host.syslog_reconfigure"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_hostArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "host"), host)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _hostArg)
	return
}


var HostClass_EvacuateMockedCallback = func (sessionID SessionRef, host HostRef) (_err error) {
	log.Println("Host.Evacuate not mocked")
	_err = errors.New("Host.Evacuate not mocked")
	return
}

func (_class HostClass) EvacuateMock(sessionID SessionRef, host HostRef) (_err error) {
	return HostClass_EvacuateMockedCallback(sessionID, host)
}
// Migrate all VMs off of this host, where possible.
func (_class HostClass) Evacuate(sessionID SessionRef, host HostRef) (_err error) {
	if (IsMock) {
		return _class.EvacuateMock(sessionID, host)
	}	
	_method := "host.evacuate"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_hostArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "host"), host)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _hostArg)
	return
}


var HostClass_GetUncooperativeResidentVMsMockedCallback = func (sessionID SessionRef, self HostRef) (_retval []VMRef, _err error) {
	log.Println("Host.GetUncooperativeResidentVMs not mocked")
	_err = errors.New("Host.GetUncooperativeResidentVMs not mocked")
	return
}

func (_class HostClass) GetUncooperativeResidentVMsMock(sessionID SessionRef, self HostRef) (_retval []VMRef, _err error) {
	return HostClass_GetUncooperativeResidentVMsMockedCallback(sessionID, self)
}
// Return a set of VMs which are not co-operating with the host's memory control system
func (_class HostClass) GetUncooperativeResidentVMs(sessionID SessionRef, self HostRef) (_retval []VMRef, _err error) {
	if (IsMock) {
		return _class.GetUncooperativeResidentVMsMock(sessionID, self)
	}	
	_method := "host.get_uncooperative_resident_VMs"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var HostClass_GetVmsWhichPreventEvacuationMockedCallback = func (sessionID SessionRef, self HostRef) (_retval map[VMRef][]string, _err error) {
	log.Println("Host.GetVmsWhichPreventEvacuation not mocked")
	_err = errors.New("Host.GetVmsWhichPreventEvacuation not mocked")
	return
}

func (_class HostClass) GetVmsWhichPreventEvacuationMock(sessionID SessionRef, self HostRef) (_retval map[VMRef][]string, _err error) {
	return HostClass_GetVmsWhichPreventEvacuationMockedCallback(sessionID, self)
}
// Return a set of VMs which prevent the host being evacuated, with per-VM error codes
func (_class HostClass) GetVmsWhichPreventEvacuation(sessionID SessionRef, self HostRef) (_retval map[VMRef][]string, _err error) {
	if (IsMock) {
		return _class.GetVmsWhichPreventEvacuationMock(sessionID, self)
	}	
	_method := "host.get_vms_which_prevent_evacuation"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVMRefToStringSetMapToGo(_method + " -> ", _result.Value)
	return
}


var HostClass_AssertCanEvacuateMockedCallback = func (sessionID SessionRef, host HostRef) (_err error) {
	log.Println("Host.AssertCanEvacuate not mocked")
	_err = errors.New("Host.AssertCanEvacuate not mocked")
	return
}

func (_class HostClass) AssertCanEvacuateMock(sessionID SessionRef, host HostRef) (_err error) {
	return HostClass_AssertCanEvacuateMockedCallback(sessionID, host)
}
// Check this host can be evacuated.
func (_class HostClass) AssertCanEvacuate(sessionID SessionRef, host HostRef) (_err error) {
	if (IsMock) {
		return _class.AssertCanEvacuateMock(sessionID, host)
	}	
	_method := "host.assert_can_evacuate"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_hostArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "host"), host)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _hostArg)
	return
}


var HostClass_ForgetDataSourceArchivesMockedCallback = func (sessionID SessionRef, host HostRef, dataSource string) (_err error) {
	log.Println("Host.ForgetDataSourceArchives not mocked")
	_err = errors.New("Host.ForgetDataSourceArchives not mocked")
	return
}

func (_class HostClass) ForgetDataSourceArchivesMock(sessionID SessionRef, host HostRef, dataSource string) (_err error) {
	return HostClass_ForgetDataSourceArchivesMockedCallback(sessionID, host, dataSource)
}
// Forget the recorded statistics related to the specified data source
func (_class HostClass) ForgetDataSourceArchives(sessionID SessionRef, host HostRef, dataSource string) (_err error) {
	if (IsMock) {
		return _class.ForgetDataSourceArchivesMock(sessionID, host, dataSource)
	}	
	_method := "host.forget_data_source_archives"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_hostArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "host"), host)
	if _err != nil {
		return
	}
	_dataSourceArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "data_source"), dataSource)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _hostArg, _dataSourceArg)
	return
}


var HostClass_QueryDataSourceMockedCallback = func (sessionID SessionRef, host HostRef, dataSource string) (_retval float64, _err error) {
	log.Println("Host.QueryDataSource not mocked")
	_err = errors.New("Host.QueryDataSource not mocked")
	return
}

func (_class HostClass) QueryDataSourceMock(sessionID SessionRef, host HostRef, dataSource string) (_retval float64, _err error) {
	return HostClass_QueryDataSourceMockedCallback(sessionID, host, dataSource)
}
// Query the latest value of the specified data source
func (_class HostClass) QueryDataSource(sessionID SessionRef, host HostRef, dataSource string) (_retval float64, _err error) {
	if (IsMock) {
		return _class.QueryDataSourceMock(sessionID, host, dataSource)
	}	
	_method := "host.query_data_source"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_hostArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "host"), host)
	if _err != nil {
		return
	}
	_dataSourceArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "data_source"), dataSource)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _hostArg, _dataSourceArg)
	if _err != nil {
		return
	}
	_retval, _err = convertFloatToGo(_method + " -> ", _result.Value)
	return
}


var HostClass_RecordDataSourceMockedCallback = func (sessionID SessionRef, host HostRef, dataSource string) (_err error) {
	log.Println("Host.RecordDataSource not mocked")
	_err = errors.New("Host.RecordDataSource not mocked")
	return
}

func (_class HostClass) RecordDataSourceMock(sessionID SessionRef, host HostRef, dataSource string) (_err error) {
	return HostClass_RecordDataSourceMockedCallback(sessionID, host, dataSource)
}
// Start recording the specified data source
func (_class HostClass) RecordDataSource(sessionID SessionRef, host HostRef, dataSource string) (_err error) {
	if (IsMock) {
		return _class.RecordDataSourceMock(sessionID, host, dataSource)
	}	
	_method := "host.record_data_source"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_hostArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "host"), host)
	if _err != nil {
		return
	}
	_dataSourceArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "data_source"), dataSource)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _hostArg, _dataSourceArg)
	return
}


var HostClass_GetDataSourcesMockedCallback = func (sessionID SessionRef, host HostRef) (_retval []DataSourceRecord, _err error) {
	log.Println("Host.GetDataSources not mocked")
	_err = errors.New("Host.GetDataSources not mocked")
	return
}

func (_class HostClass) GetDataSourcesMock(sessionID SessionRef, host HostRef) (_retval []DataSourceRecord, _err error) {
	return HostClass_GetDataSourcesMockedCallback(sessionID, host)
}
// 
func (_class HostClass) GetDataSources(sessionID SessionRef, host HostRef) (_retval []DataSourceRecord, _err error) {
	if (IsMock) {
		return _class.GetDataSourcesMock(sessionID, host)
	}	
	_method := "host.get_data_sources"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_hostArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "host"), host)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _hostArg)
	if _err != nil {
		return
	}
	_retval, _err = convertDataSourceRecordSetToGo(_method + " -> ", _result.Value)
	return
}


var HostClass_EmergencyHaDisableMockedCallback = func (sessionID SessionRef, soft bool) (_err error) {
	log.Println("Host.EmergencyHaDisable not mocked")
	_err = errors.New("Host.EmergencyHaDisable not mocked")
	return
}

func (_class HostClass) EmergencyHaDisableMock(sessionID SessionRef, soft bool) (_err error) {
	return HostClass_EmergencyHaDisableMockedCallback(sessionID, soft)
}
// This call disables HA on the local host. This should only be used with extreme care.
func (_class HostClass) EmergencyHaDisable(sessionID SessionRef, soft bool) (_err error) {
	if (IsMock) {
		return _class.EmergencyHaDisableMock(sessionID, soft)
	}	
	_method := "host.emergency_ha_disable"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_softArg, _err := convertBoolToXen(fmt.Sprintf("%s(%s)", _method, "soft"), soft)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _softArg)
	return
}


var HostClass_PowerOnMockedCallback = func (sessionID SessionRef, host HostRef) (_err error) {
	log.Println("Host.PowerOn not mocked")
	_err = errors.New("Host.PowerOn not mocked")
	return
}

func (_class HostClass) PowerOnMock(sessionID SessionRef, host HostRef) (_err error) {
	return HostClass_PowerOnMockedCallback(sessionID, host)
}
// Attempt to power-on the host (if the capability exists).
func (_class HostClass) PowerOn(sessionID SessionRef, host HostRef) (_err error) {
	if (IsMock) {
		return _class.PowerOnMock(sessionID, host)
	}	
	_method := "host.power_on"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_hostArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "host"), host)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _hostArg)
	return
}


var HostClass_DestroyMockedCallback = func (sessionID SessionRef, self HostRef) (_err error) {
	log.Println("Host.Destroy not mocked")
	_err = errors.New("Host.Destroy not mocked")
	return
}

func (_class HostClass) DestroyMock(sessionID SessionRef, self HostRef) (_err error) {
	return HostClass_DestroyMockedCallback(sessionID, self)
}
// Destroy specified host record in database
func (_class HostClass) Destroy(sessionID SessionRef, self HostRef) (_err error) {
	if (IsMock) {
		return _class.DestroyMock(sessionID, self)
	}	
	_method := "host.destroy"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg)
	return
}


var HostClass_LicenseRemoveMockedCallback = func (sessionID SessionRef, host HostRef) (_err error) {
	log.Println("Host.LicenseRemove not mocked")
	_err = errors.New("Host.LicenseRemove not mocked")
	return
}

func (_class HostClass) LicenseRemoveMock(sessionID SessionRef, host HostRef) (_err error) {
	return HostClass_LicenseRemoveMockedCallback(sessionID, host)
}
// Remove any license file from the specified host, and switch that host to the unlicensed edition
func (_class HostClass) LicenseRemove(sessionID SessionRef, host HostRef) (_err error) {
	if (IsMock) {
		return _class.LicenseRemoveMock(sessionID, host)
	}	
	_method := "host.license_remove"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_hostArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "host"), host)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _hostArg)
	return
}


var HostClass_LicenseAddMockedCallback = func (sessionID SessionRef, host HostRef, contents string) (_err error) {
	log.Println("Host.LicenseAdd not mocked")
	_err = errors.New("Host.LicenseAdd not mocked")
	return
}

func (_class HostClass) LicenseAddMock(sessionID SessionRef, host HostRef, contents string) (_err error) {
	return HostClass_LicenseAddMockedCallback(sessionID, host, contents)
}
// Apply a new license to a host
//
// Errors:
//  LICENSE_PROCESSING_ERROR - There was an error processing your license.  Please contact your support representative.
func (_class HostClass) LicenseAdd(sessionID SessionRef, host HostRef, contents string) (_err error) {
	if (IsMock) {
		return _class.LicenseAddMock(sessionID, host, contents)
	}	
	_method := "host.license_add"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_hostArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "host"), host)
	if _err != nil {
		return
	}
	_contentsArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "contents"), contents)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _hostArg, _contentsArg)
	return
}


var HostClass_LicenseApplyMockedCallback = func (sessionID SessionRef, host HostRef, contents string) (_err error) {
	log.Println("Host.LicenseApply not mocked")
	_err = errors.New("Host.LicenseApply not mocked")
	return
}

func (_class HostClass) LicenseApplyMock(sessionID SessionRef, host HostRef, contents string) (_err error) {
	return HostClass_LicenseApplyMockedCallback(sessionID, host, contents)
}
// Apply a new license to a host
//
// Errors:
//  LICENSE_PROCESSING_ERROR - There was an error processing your license.  Please contact your support representative.
func (_class HostClass) LicenseApply(sessionID SessionRef, host HostRef, contents string) (_err error) {
	if (IsMock) {
		return _class.LicenseApplyMock(sessionID, host, contents)
	}	
	_method := "host.license_apply"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_hostArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "host"), host)
	if _err != nil {
		return
	}
	_contentsArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "contents"), contents)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _hostArg, _contentsArg)
	return
}


var HostClass_ListMethodsMockedCallback = func (sessionID SessionRef) (_retval []string, _err error) {
	log.Println("Host.ListMethods not mocked")
	_err = errors.New("Host.ListMethods not mocked")
	return
}

func (_class HostClass) ListMethodsMock(sessionID SessionRef) (_retval []string, _err error) {
	return HostClass_ListMethodsMockedCallback(sessionID)
}
// List all supported methods
func (_class HostClass) ListMethods(sessionID SessionRef) (_retval []string, _err error) {
	if (IsMock) {
		return _class.ListMethodsMock(sessionID)
	}	
	_method := "host.list_methods"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringSetToGo(_method + " -> ", _result.Value)
	return
}


var HostClass_BugreportUploadMockedCallback = func (sessionID SessionRef, host HostRef, url string, options map[string]string) (_err error) {
	log.Println("Host.BugreportUpload not mocked")
	_err = errors.New("Host.BugreportUpload not mocked")
	return
}

func (_class HostClass) BugreportUploadMock(sessionID SessionRef, host HostRef, url string, options map[string]string) (_err error) {
	return HostClass_BugreportUploadMockedCallback(sessionID, host, url, options)
}
// Run xen-bugtool --yestoall and upload the output to support
func (_class HostClass) BugreportUpload(sessionID SessionRef, host HostRef, url string, options map[string]string) (_err error) {
	if (IsMock) {
		return _class.BugreportUploadMock(sessionID, host, url, options)
	}	
	_method := "host.bugreport_upload"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_hostArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "host"), host)
	if _err != nil {
		return
	}
	_urlArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "url"), url)
	if _err != nil {
		return
	}
	_optionsArg, _err := convertStringToStringMapToXen(fmt.Sprintf("%s(%s)", _method, "options"), options)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _hostArg, _urlArg, _optionsArg)
	return
}


var HostClass_SendDebugKeysMockedCallback = func (sessionID SessionRef, host HostRef, keys string) (_err error) {
	log.Println("Host.SendDebugKeys not mocked")
	_err = errors.New("Host.SendDebugKeys not mocked")
	return
}

func (_class HostClass) SendDebugKeysMock(sessionID SessionRef, host HostRef, keys string) (_err error) {
	return HostClass_SendDebugKeysMockedCallback(sessionID, host, keys)
}
// Inject the given string as debugging keys into Xen
func (_class HostClass) SendDebugKeys(sessionID SessionRef, host HostRef, keys string) (_err error) {
	if (IsMock) {
		return _class.SendDebugKeysMock(sessionID, host, keys)
	}	
	_method := "host.send_debug_keys"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_hostArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "host"), host)
	if _err != nil {
		return
	}
	_keysArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "keys"), keys)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _hostArg, _keysArg)
	return
}


var HostClass_GetLogMockedCallback = func (sessionID SessionRef, host HostRef) (_retval string, _err error) {
	log.Println("Host.GetLog not mocked")
	_err = errors.New("Host.GetLog not mocked")
	return
}

func (_class HostClass) GetLogMock(sessionID SessionRef, host HostRef) (_retval string, _err error) {
	return HostClass_GetLogMockedCallback(sessionID, host)
}
// Get the host's log file
func (_class HostClass) GetLog(sessionID SessionRef, host HostRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetLogMock(sessionID, host)
	}	
	_method := "host.get_log"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_hostArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "host"), host)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _hostArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringToGo(_method + " -> ", _result.Value)
	return
}


var HostClass_DmesgClearMockedCallback = func (sessionID SessionRef, host HostRef) (_retval string, _err error) {
	log.Println("Host.DmesgClear not mocked")
	_err = errors.New("Host.DmesgClear not mocked")
	return
}

func (_class HostClass) DmesgClearMock(sessionID SessionRef, host HostRef) (_retval string, _err error) {
	return HostClass_DmesgClearMockedCallback(sessionID, host)
}
// Get the host xen dmesg, and clear the buffer.
func (_class HostClass) DmesgClear(sessionID SessionRef, host HostRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.DmesgClearMock(sessionID, host)
	}	
	_method := "host.dmesg_clear"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_hostArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "host"), host)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _hostArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringToGo(_method + " -> ", _result.Value)
	return
}


var HostClass_DmesgMockedCallback = func (sessionID SessionRef, host HostRef) (_retval string, _err error) {
	log.Println("Host.Dmesg not mocked")
	_err = errors.New("Host.Dmesg not mocked")
	return
}

func (_class HostClass) DmesgMock(sessionID SessionRef, host HostRef) (_retval string, _err error) {
	return HostClass_DmesgMockedCallback(sessionID, host)
}
// Get the host xen dmesg.
func (_class HostClass) Dmesg(sessionID SessionRef, host HostRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.DmesgMock(sessionID, host)
	}	
	_method := "host.dmesg"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_hostArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "host"), host)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _hostArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringToGo(_method + " -> ", _result.Value)
	return
}


var HostClass_RebootMockedCallback = func (sessionID SessionRef, host HostRef) (_err error) {
	log.Println("Host.Reboot not mocked")
	_err = errors.New("Host.Reboot not mocked")
	return
}

func (_class HostClass) RebootMock(sessionID SessionRef, host HostRef) (_err error) {
	return HostClass_RebootMockedCallback(sessionID, host)
}
// Reboot the host. (This function can only be called if there are no currently running VMs on the host and it is disabled.)
func (_class HostClass) Reboot(sessionID SessionRef, host HostRef) (_err error) {
	if (IsMock) {
		return _class.RebootMock(sessionID, host)
	}	
	_method := "host.reboot"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_hostArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "host"), host)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _hostArg)
	return
}


var HostClass_ShutdownMockedCallback = func (sessionID SessionRef, host HostRef) (_err error) {
	log.Println("Host.Shutdown not mocked")
	_err = errors.New("Host.Shutdown not mocked")
	return
}

func (_class HostClass) ShutdownMock(sessionID SessionRef, host HostRef) (_err error) {
	return HostClass_ShutdownMockedCallback(sessionID, host)
}
// Shutdown the host. (This function can only be called if there are no currently running VMs on the host and it is disabled.)
func (_class HostClass) Shutdown(sessionID SessionRef, host HostRef) (_err error) {
	if (IsMock) {
		return _class.ShutdownMock(sessionID, host)
	}	
	_method := "host.shutdown"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_hostArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "host"), host)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _hostArg)
	return
}


var HostClass_EnableMockedCallback = func (sessionID SessionRef, host HostRef) (_err error) {
	log.Println("Host.Enable not mocked")
	_err = errors.New("Host.Enable not mocked")
	return
}

func (_class HostClass) EnableMock(sessionID SessionRef, host HostRef) (_err error) {
	return HostClass_EnableMockedCallback(sessionID, host)
}
// Puts the host into a state in which new VMs can be started.
func (_class HostClass) Enable(sessionID SessionRef, host HostRef) (_err error) {
	if (IsMock) {
		return _class.EnableMock(sessionID, host)
	}	
	_method := "host.enable"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_hostArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "host"), host)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _hostArg)
	return
}


var HostClass_DisableMockedCallback = func (sessionID SessionRef, host HostRef) (_err error) {
	log.Println("Host.Disable not mocked")
	_err = errors.New("Host.Disable not mocked")
	return
}

func (_class HostClass) DisableMock(sessionID SessionRef, host HostRef) (_err error) {
	return HostClass_DisableMockedCallback(sessionID, host)
}
// Puts the host into a state in which no new VMs can be started. Currently active VMs on the host continue to execute.
func (_class HostClass) Disable(sessionID SessionRef, host HostRef) (_err error) {
	if (IsMock) {
		return _class.DisableMock(sessionID, host)
	}	
	_method := "host.disable"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_hostArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "host"), host)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _hostArg)
	return
}


var HostClass_SetDisplayMockedCallback = func (sessionID SessionRef, self HostRef, value HostDisplay) (_err error) {
	log.Println("Host.SetDisplay not mocked")
	_err = errors.New("Host.SetDisplay not mocked")
	return
}

func (_class HostClass) SetDisplayMock(sessionID SessionRef, self HostRef, value HostDisplay) (_err error) {
	return HostClass_SetDisplayMockedCallback(sessionID, self, value)
}
// Set the display field of the given host.
func (_class HostClass) SetDisplay(sessionID SessionRef, self HostRef, value HostDisplay) (_err error) {
	if (IsMock) {
		return _class.SetDisplayMock(sessionID, self, value)
	}	
	_method := "host.set_display"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertEnumHostDisplayToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}


var HostClass_RemoveFromGuestVCPUsParamsMockedCallback = func (sessionID SessionRef, self HostRef, key string) (_err error) {
	log.Println("Host.RemoveFromGuestVCPUsParams not mocked")
	_err = errors.New("Host.RemoveFromGuestVCPUsParams not mocked")
	return
}

func (_class HostClass) RemoveFromGuestVCPUsParamsMock(sessionID SessionRef, self HostRef, key string) (_err error) {
	return HostClass_RemoveFromGuestVCPUsParamsMockedCallback(sessionID, self, key)
}
// Remove the given key and its corresponding value from the guest_VCPUs_params field of the given host.  If the key is not in that Map, then do nothing.
func (_class HostClass) RemoveFromGuestVCPUsParams(sessionID SessionRef, self HostRef, key string) (_err error) {
	if (IsMock) {
		return _class.RemoveFromGuestVCPUsParamsMock(sessionID, self, key)
	}	
	_method := "host.remove_from_guest_VCPUs_params"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var HostClass_AddToGuestVCPUsParamsMockedCallback = func (sessionID SessionRef, self HostRef, key string, value string) (_err error) {
	log.Println("Host.AddToGuestVCPUsParams not mocked")
	_err = errors.New("Host.AddToGuestVCPUsParams not mocked")
	return
}

func (_class HostClass) AddToGuestVCPUsParamsMock(sessionID SessionRef, self HostRef, key string, value string) (_err error) {
	return HostClass_AddToGuestVCPUsParamsMockedCallback(sessionID, self, key, value)
}
// Add the given key-value pair to the guest_VCPUs_params field of the given host.
func (_class HostClass) AddToGuestVCPUsParams(sessionID SessionRef, self HostRef, key string, value string) (_err error) {
	if (IsMock) {
		return _class.AddToGuestVCPUsParamsMock(sessionID, self, key, value)
	}	
	_method := "host.add_to_guest_VCPUs_params"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var HostClass_SetGuestVCPUsParamsMockedCallback = func (sessionID SessionRef, self HostRef, value map[string]string) (_err error) {
	log.Println("Host.SetGuestVCPUsParams not mocked")
	_err = errors.New("Host.SetGuestVCPUsParams not mocked")
	return
}

func (_class HostClass) SetGuestVCPUsParamsMock(sessionID SessionRef, self HostRef, value map[string]string) (_err error) {
	return HostClass_SetGuestVCPUsParamsMockedCallback(sessionID, self, value)
}
// Set the guest_VCPUs_params field of the given host.
func (_class HostClass) SetGuestVCPUsParams(sessionID SessionRef, self HostRef, value map[string]string) (_err error) {
	if (IsMock) {
		return _class.SetGuestVCPUsParamsMock(sessionID, self, value)
	}	
	_method := "host.set_guest_VCPUs_params"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var HostClass_RemoveFromLicenseServerMockedCallback = func (sessionID SessionRef, self HostRef, key string) (_err error) {
	log.Println("Host.RemoveFromLicenseServer not mocked")
	_err = errors.New("Host.RemoveFromLicenseServer not mocked")
	return
}

func (_class HostClass) RemoveFromLicenseServerMock(sessionID SessionRef, self HostRef, key string) (_err error) {
	return HostClass_RemoveFromLicenseServerMockedCallback(sessionID, self, key)
}
// Remove the given key and its corresponding value from the license_server field of the given host.  If the key is not in that Map, then do nothing.
func (_class HostClass) RemoveFromLicenseServer(sessionID SessionRef, self HostRef, key string) (_err error) {
	if (IsMock) {
		return _class.RemoveFromLicenseServerMock(sessionID, self, key)
	}	
	_method := "host.remove_from_license_server"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var HostClass_AddToLicenseServerMockedCallback = func (sessionID SessionRef, self HostRef, key string, value string) (_err error) {
	log.Println("Host.AddToLicenseServer not mocked")
	_err = errors.New("Host.AddToLicenseServer not mocked")
	return
}

func (_class HostClass) AddToLicenseServerMock(sessionID SessionRef, self HostRef, key string, value string) (_err error) {
	return HostClass_AddToLicenseServerMockedCallback(sessionID, self, key, value)
}
// Add the given key-value pair to the license_server field of the given host.
func (_class HostClass) AddToLicenseServer(sessionID SessionRef, self HostRef, key string, value string) (_err error) {
	if (IsMock) {
		return _class.AddToLicenseServerMock(sessionID, self, key, value)
	}	
	_method := "host.add_to_license_server"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var HostClass_SetLicenseServerMockedCallback = func (sessionID SessionRef, self HostRef, value map[string]string) (_err error) {
	log.Println("Host.SetLicenseServer not mocked")
	_err = errors.New("Host.SetLicenseServer not mocked")
	return
}

func (_class HostClass) SetLicenseServerMock(sessionID SessionRef, self HostRef, value map[string]string) (_err error) {
	return HostClass_SetLicenseServerMockedCallback(sessionID, self, value)
}
// Set the license_server field of the given host.
func (_class HostClass) SetLicenseServer(sessionID SessionRef, self HostRef, value map[string]string) (_err error) {
	if (IsMock) {
		return _class.SetLicenseServerMock(sessionID, self, value)
	}	
	_method := "host.set_license_server"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var HostClass_RemoveTagsMockedCallback = func (sessionID SessionRef, self HostRef, value string) (_err error) {
	log.Println("Host.RemoveTags not mocked")
	_err = errors.New("Host.RemoveTags not mocked")
	return
}

func (_class HostClass) RemoveTagsMock(sessionID SessionRef, self HostRef, value string) (_err error) {
	return HostClass_RemoveTagsMockedCallback(sessionID, self, value)
}
// Remove the given value from the tags field of the given host.  If the value is not in that Set, then do nothing.
func (_class HostClass) RemoveTags(sessionID SessionRef, self HostRef, value string) (_err error) {
	if (IsMock) {
		return _class.RemoveTagsMock(sessionID, self, value)
	}	
	_method := "host.remove_tags"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var HostClass_AddTagsMockedCallback = func (sessionID SessionRef, self HostRef, value string) (_err error) {
	log.Println("Host.AddTags not mocked")
	_err = errors.New("Host.AddTags not mocked")
	return
}

func (_class HostClass) AddTagsMock(sessionID SessionRef, self HostRef, value string) (_err error) {
	return HostClass_AddTagsMockedCallback(sessionID, self, value)
}
// Add the given value to the tags field of the given host.  If the value is already in that Set, then do nothing.
func (_class HostClass) AddTags(sessionID SessionRef, self HostRef, value string) (_err error) {
	if (IsMock) {
		return _class.AddTagsMock(sessionID, self, value)
	}	
	_method := "host.add_tags"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var HostClass_SetTagsMockedCallback = func (sessionID SessionRef, self HostRef, value []string) (_err error) {
	log.Println("Host.SetTags not mocked")
	_err = errors.New("Host.SetTags not mocked")
	return
}

func (_class HostClass) SetTagsMock(sessionID SessionRef, self HostRef, value []string) (_err error) {
	return HostClass_SetTagsMockedCallback(sessionID, self, value)
}
// Set the tags field of the given host.
func (_class HostClass) SetTags(sessionID SessionRef, self HostRef, value []string) (_err error) {
	if (IsMock) {
		return _class.SetTagsMock(sessionID, self, value)
	}	
	_method := "host.set_tags"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var HostClass_SetAddressMockedCallback = func (sessionID SessionRef, self HostRef, value string) (_err error) {
	log.Println("Host.SetAddress not mocked")
	_err = errors.New("Host.SetAddress not mocked")
	return
}

func (_class HostClass) SetAddressMock(sessionID SessionRef, self HostRef, value string) (_err error) {
	return HostClass_SetAddressMockedCallback(sessionID, self, value)
}
// Set the address field of the given host.
func (_class HostClass) SetAddress(sessionID SessionRef, self HostRef, value string) (_err error) {
	if (IsMock) {
		return _class.SetAddressMock(sessionID, self, value)
	}	
	_method := "host.set_address"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var HostClass_SetHostnameMockedCallback = func (sessionID SessionRef, self HostRef, value string) (_err error) {
	log.Println("Host.SetHostname not mocked")
	_err = errors.New("Host.SetHostname not mocked")
	return
}

func (_class HostClass) SetHostnameMock(sessionID SessionRef, self HostRef, value string) (_err error) {
	return HostClass_SetHostnameMockedCallback(sessionID, self, value)
}
// Set the hostname field of the given host.
func (_class HostClass) SetHostname(sessionID SessionRef, self HostRef, value string) (_err error) {
	if (IsMock) {
		return _class.SetHostnameMock(sessionID, self, value)
	}	
	_method := "host.set_hostname"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var HostClass_SetCrashDumpSrMockedCallback = func (sessionID SessionRef, self HostRef, value SRRef) (_err error) {
	log.Println("Host.SetCrashDumpSr not mocked")
	_err = errors.New("Host.SetCrashDumpSr not mocked")
	return
}

func (_class HostClass) SetCrashDumpSrMock(sessionID SessionRef, self HostRef, value SRRef) (_err error) {
	return HostClass_SetCrashDumpSrMockedCallback(sessionID, self, value)
}
// Set the crash_dump_sr field of the given host.
func (_class HostClass) SetCrashDumpSr(sessionID SessionRef, self HostRef, value SRRef) (_err error) {
	if (IsMock) {
		return _class.SetCrashDumpSrMock(sessionID, self, value)
	}	
	_method := "host.set_crash_dump_sr"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var HostClass_SetSuspendImageSrMockedCallback = func (sessionID SessionRef, self HostRef, value SRRef) (_err error) {
	log.Println("Host.SetSuspendImageSr not mocked")
	_err = errors.New("Host.SetSuspendImageSr not mocked")
	return
}

func (_class HostClass) SetSuspendImageSrMock(sessionID SessionRef, self HostRef, value SRRef) (_err error) {
	return HostClass_SetSuspendImageSrMockedCallback(sessionID, self, value)
}
// Set the suspend_image_sr field of the given host.
func (_class HostClass) SetSuspendImageSr(sessionID SessionRef, self HostRef, value SRRef) (_err error) {
	if (IsMock) {
		return _class.SetSuspendImageSrMock(sessionID, self, value)
	}	
	_method := "host.set_suspend_image_sr"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var HostClass_RemoveFromLoggingMockedCallback = func (sessionID SessionRef, self HostRef, key string) (_err error) {
	log.Println("Host.RemoveFromLogging not mocked")
	_err = errors.New("Host.RemoveFromLogging not mocked")
	return
}

func (_class HostClass) RemoveFromLoggingMock(sessionID SessionRef, self HostRef, key string) (_err error) {
	return HostClass_RemoveFromLoggingMockedCallback(sessionID, self, key)
}
// Remove the given key and its corresponding value from the logging field of the given host.  If the key is not in that Map, then do nothing.
func (_class HostClass) RemoveFromLogging(sessionID SessionRef, self HostRef, key string) (_err error) {
	if (IsMock) {
		return _class.RemoveFromLoggingMock(sessionID, self, key)
	}	
	_method := "host.remove_from_logging"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var HostClass_AddToLoggingMockedCallback = func (sessionID SessionRef, self HostRef, key string, value string) (_err error) {
	log.Println("Host.AddToLogging not mocked")
	_err = errors.New("Host.AddToLogging not mocked")
	return
}

func (_class HostClass) AddToLoggingMock(sessionID SessionRef, self HostRef, key string, value string) (_err error) {
	return HostClass_AddToLoggingMockedCallback(sessionID, self, key, value)
}
// Add the given key-value pair to the logging field of the given host.
func (_class HostClass) AddToLogging(sessionID SessionRef, self HostRef, key string, value string) (_err error) {
	if (IsMock) {
		return _class.AddToLoggingMock(sessionID, self, key, value)
	}	
	_method := "host.add_to_logging"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var HostClass_SetLoggingMockedCallback = func (sessionID SessionRef, self HostRef, value map[string]string) (_err error) {
	log.Println("Host.SetLogging not mocked")
	_err = errors.New("Host.SetLogging not mocked")
	return
}

func (_class HostClass) SetLoggingMock(sessionID SessionRef, self HostRef, value map[string]string) (_err error) {
	return HostClass_SetLoggingMockedCallback(sessionID, self, value)
}
// Set the logging field of the given host.
func (_class HostClass) SetLogging(sessionID SessionRef, self HostRef, value map[string]string) (_err error) {
	if (IsMock) {
		return _class.SetLoggingMock(sessionID, self, value)
	}	
	_method := "host.set_logging"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var HostClass_RemoveFromOtherConfigMockedCallback = func (sessionID SessionRef, self HostRef, key string) (_err error) {
	log.Println("Host.RemoveFromOtherConfig not mocked")
	_err = errors.New("Host.RemoveFromOtherConfig not mocked")
	return
}

func (_class HostClass) RemoveFromOtherConfigMock(sessionID SessionRef, self HostRef, key string) (_err error) {
	return HostClass_RemoveFromOtherConfigMockedCallback(sessionID, self, key)
}
// Remove the given key and its corresponding value from the other_config field of the given host.  If the key is not in that Map, then do nothing.
func (_class HostClass) RemoveFromOtherConfig(sessionID SessionRef, self HostRef, key string) (_err error) {
	if (IsMock) {
		return _class.RemoveFromOtherConfigMock(sessionID, self, key)
	}	
	_method := "host.remove_from_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var HostClass_AddToOtherConfigMockedCallback = func (sessionID SessionRef, self HostRef, key string, value string) (_err error) {
	log.Println("Host.AddToOtherConfig not mocked")
	_err = errors.New("Host.AddToOtherConfig not mocked")
	return
}

func (_class HostClass) AddToOtherConfigMock(sessionID SessionRef, self HostRef, key string, value string) (_err error) {
	return HostClass_AddToOtherConfigMockedCallback(sessionID, self, key, value)
}
// Add the given key-value pair to the other_config field of the given host.
func (_class HostClass) AddToOtherConfig(sessionID SessionRef, self HostRef, key string, value string) (_err error) {
	if (IsMock) {
		return _class.AddToOtherConfigMock(sessionID, self, key, value)
	}	
	_method := "host.add_to_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var HostClass_SetOtherConfigMockedCallback = func (sessionID SessionRef, self HostRef, value map[string]string) (_err error) {
	log.Println("Host.SetOtherConfig not mocked")
	_err = errors.New("Host.SetOtherConfig not mocked")
	return
}

func (_class HostClass) SetOtherConfigMock(sessionID SessionRef, self HostRef, value map[string]string) (_err error) {
	return HostClass_SetOtherConfigMockedCallback(sessionID, self, value)
}
// Set the other_config field of the given host.
func (_class HostClass) SetOtherConfig(sessionID SessionRef, self HostRef, value map[string]string) (_err error) {
	if (IsMock) {
		return _class.SetOtherConfigMock(sessionID, self, value)
	}	
	_method := "host.set_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var HostClass_SetNameDescriptionMockedCallback = func (sessionID SessionRef, self HostRef, value string) (_err error) {
	log.Println("Host.SetNameDescription not mocked")
	_err = errors.New("Host.SetNameDescription not mocked")
	return
}

func (_class HostClass) SetNameDescriptionMock(sessionID SessionRef, self HostRef, value string) (_err error) {
	return HostClass_SetNameDescriptionMockedCallback(sessionID, self, value)
}
// Set the name/description field of the given host.
func (_class HostClass) SetNameDescription(sessionID SessionRef, self HostRef, value string) (_err error) {
	if (IsMock) {
		return _class.SetNameDescriptionMock(sessionID, self, value)
	}	
	_method := "host.set_name_description"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var HostClass_SetNameLabelMockedCallback = func (sessionID SessionRef, self HostRef, value string) (_err error) {
	log.Println("Host.SetNameLabel not mocked")
	_err = errors.New("Host.SetNameLabel not mocked")
	return
}

func (_class HostClass) SetNameLabelMock(sessionID SessionRef, self HostRef, value string) (_err error) {
	return HostClass_SetNameLabelMockedCallback(sessionID, self, value)
}
// Set the name/label field of the given host.
func (_class HostClass) SetNameLabel(sessionID SessionRef, self HostRef, value string) (_err error) {
	if (IsMock) {
		return _class.SetNameLabelMock(sessionID, self, value)
	}	
	_method := "host.set_name_label"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var HostClass_GetFeaturesMockedCallback = func (sessionID SessionRef, self HostRef) (_retval []FeatureRef, _err error) {
	log.Println("Host.GetFeatures not mocked")
	_err = errors.New("Host.GetFeatures not mocked")
	return
}

func (_class HostClass) GetFeaturesMock(sessionID SessionRef, self HostRef) (_retval []FeatureRef, _err error) {
	return HostClass_GetFeaturesMockedCallback(sessionID, self)
}
// Get the features field of the given host.
func (_class HostClass) GetFeatures(sessionID SessionRef, self HostRef) (_retval []FeatureRef, _err error) {
	if (IsMock) {
		return _class.GetFeaturesMock(sessionID, self)
	}	
	_method := "host.get_features"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertFeatureRefSetToGo(_method + " -> ", _result.Value)
	return
}


var HostClass_GetUpdatesRequiringRebootMockedCallback = func (sessionID SessionRef, self HostRef) (_retval []PoolUpdateRef, _err error) {
	log.Println("Host.GetUpdatesRequiringReboot not mocked")
	_err = errors.New("Host.GetUpdatesRequiringReboot not mocked")
	return
}

func (_class HostClass) GetUpdatesRequiringRebootMock(sessionID SessionRef, self HostRef) (_retval []PoolUpdateRef, _err error) {
	return HostClass_GetUpdatesRequiringRebootMockedCallback(sessionID, self)
}
// Get the updates_requiring_reboot field of the given host.
func (_class HostClass) GetUpdatesRequiringReboot(sessionID SessionRef, self HostRef) (_retval []PoolUpdateRef, _err error) {
	if (IsMock) {
		return _class.GetUpdatesRequiringRebootMock(sessionID, self)
	}	
	_method := "host.get_updates_requiring_reboot"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertPoolUpdateRefSetToGo(_method + " -> ", _result.Value)
	return
}


var HostClass_GetControlDomainMockedCallback = func (sessionID SessionRef, self HostRef) (_retval VMRef, _err error) {
	log.Println("Host.GetControlDomain not mocked")
	_err = errors.New("Host.GetControlDomain not mocked")
	return
}

func (_class HostClass) GetControlDomainMock(sessionID SessionRef, self HostRef) (_retval VMRef, _err error) {
	return HostClass_GetControlDomainMockedCallback(sessionID, self)
}
// Get the control_domain field of the given host.
func (_class HostClass) GetControlDomain(sessionID SessionRef, self HostRef) (_retval VMRef, _err error) {
	if (IsMock) {
		return _class.GetControlDomainMock(sessionID, self)
	}	
	_method := "host.get_control_domain"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var HostClass_GetVirtualHardwarePlatformVersionsMockedCallback = func (sessionID SessionRef, self HostRef) (_retval []int, _err error) {
	log.Println("Host.GetVirtualHardwarePlatformVersions not mocked")
	_err = errors.New("Host.GetVirtualHardwarePlatformVersions not mocked")
	return
}

func (_class HostClass) GetVirtualHardwarePlatformVersionsMock(sessionID SessionRef, self HostRef) (_retval []int, _err error) {
	return HostClass_GetVirtualHardwarePlatformVersionsMockedCallback(sessionID, self)
}
// Get the virtual_hardware_platform_versions field of the given host.
func (_class HostClass) GetVirtualHardwarePlatformVersions(sessionID SessionRef, self HostRef) (_retval []int, _err error) {
	if (IsMock) {
		return _class.GetVirtualHardwarePlatformVersionsMock(sessionID, self)
	}	
	_method := "host.get_virtual_hardware_platform_versions"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertIntSetToGo(_method + " -> ", _result.Value)
	return
}


var HostClass_GetDisplayMockedCallback = func (sessionID SessionRef, self HostRef) (_retval HostDisplay, _err error) {
	log.Println("Host.GetDisplay not mocked")
	_err = errors.New("Host.GetDisplay not mocked")
	return
}

func (_class HostClass) GetDisplayMock(sessionID SessionRef, self HostRef) (_retval HostDisplay, _err error) {
	return HostClass_GetDisplayMockedCallback(sessionID, self)
}
// Get the display field of the given host.
func (_class HostClass) GetDisplay(sessionID SessionRef, self HostRef) (_retval HostDisplay, _err error) {
	if (IsMock) {
		return _class.GetDisplayMock(sessionID, self)
	}	
	_method := "host.get_display"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertEnumHostDisplayToGo(_method + " -> ", _result.Value)
	return
}


var HostClass_GetGuestVCPUsParamsMockedCallback = func (sessionID SessionRef, self HostRef) (_retval map[string]string, _err error) {
	log.Println("Host.GetGuestVCPUsParams not mocked")
	_err = errors.New("Host.GetGuestVCPUsParams not mocked")
	return
}

func (_class HostClass) GetGuestVCPUsParamsMock(sessionID SessionRef, self HostRef) (_retval map[string]string, _err error) {
	return HostClass_GetGuestVCPUsParamsMockedCallback(sessionID, self)
}
// Get the guest_VCPUs_params field of the given host.
func (_class HostClass) GetGuestVCPUsParams(sessionID SessionRef, self HostRef) (_retval map[string]string, _err error) {
	if (IsMock) {
		return _class.GetGuestVCPUsParamsMock(sessionID, self)
	}	
	_method := "host.get_guest_VCPUs_params"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var HostClass_GetSslLegacyMockedCallback = func (sessionID SessionRef, self HostRef) (_retval bool, _err error) {
	log.Println("Host.GetSslLegacy not mocked")
	_err = errors.New("Host.GetSslLegacy not mocked")
	return
}

func (_class HostClass) GetSslLegacyMock(sessionID SessionRef, self HostRef) (_retval bool, _err error) {
	return HostClass_GetSslLegacyMockedCallback(sessionID, self)
}
// Get the ssl_legacy field of the given host.
func (_class HostClass) GetSslLegacy(sessionID SessionRef, self HostRef) (_retval bool, _err error) {
	if (IsMock) {
		return _class.GetSslLegacyMock(sessionID, self)
	}	
	_method := "host.get_ssl_legacy"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var HostClass_GetPUSBsMockedCallback = func (sessionID SessionRef, self HostRef) (_retval []PUSBRef, _err error) {
	log.Println("Host.GetPUSBs not mocked")
	_err = errors.New("Host.GetPUSBs not mocked")
	return
}

func (_class HostClass) GetPUSBsMock(sessionID SessionRef, self HostRef) (_retval []PUSBRef, _err error) {
	return HostClass_GetPUSBsMockedCallback(sessionID, self)
}
// Get the PUSBs field of the given host.
func (_class HostClass) GetPUSBs(sessionID SessionRef, self HostRef) (_retval []PUSBRef, _err error) {
	if (IsMock) {
		return _class.GetPUSBsMock(sessionID, self)
	}	
	_method := "host.get_PUSBs"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertPUSBRefSetToGo(_method + " -> ", _result.Value)
	return
}


var HostClass_GetPGPUsMockedCallback = func (sessionID SessionRef, self HostRef) (_retval []PGPURef, _err error) {
	log.Println("Host.GetPGPUs not mocked")
	_err = errors.New("Host.GetPGPUs not mocked")
	return
}

func (_class HostClass) GetPGPUsMock(sessionID SessionRef, self HostRef) (_retval []PGPURef, _err error) {
	return HostClass_GetPGPUsMockedCallback(sessionID, self)
}
// Get the PGPUs field of the given host.
func (_class HostClass) GetPGPUs(sessionID SessionRef, self HostRef) (_retval []PGPURef, _err error) {
	if (IsMock) {
		return _class.GetPGPUsMock(sessionID, self)
	}	
	_method := "host.get_PGPUs"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertPGPURefSetToGo(_method + " -> ", _result.Value)
	return
}


var HostClass_GetPCIsMockedCallback = func (sessionID SessionRef, self HostRef) (_retval []PCIRef, _err error) {
	log.Println("Host.GetPCIs not mocked")
	_err = errors.New("Host.GetPCIs not mocked")
	return
}

func (_class HostClass) GetPCIsMock(sessionID SessionRef, self HostRef) (_retval []PCIRef, _err error) {
	return HostClass_GetPCIsMockedCallback(sessionID, self)
}
// Get the PCIs field of the given host.
func (_class HostClass) GetPCIs(sessionID SessionRef, self HostRef) (_retval []PCIRef, _err error) {
	if (IsMock) {
		return _class.GetPCIsMock(sessionID, self)
	}	
	_method := "host.get_PCIs"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var HostClass_GetChipsetInfoMockedCallback = func (sessionID SessionRef, self HostRef) (_retval map[string]string, _err error) {
	log.Println("Host.GetChipsetInfo not mocked")
	_err = errors.New("Host.GetChipsetInfo not mocked")
	return
}

func (_class HostClass) GetChipsetInfoMock(sessionID SessionRef, self HostRef) (_retval map[string]string, _err error) {
	return HostClass_GetChipsetInfoMockedCallback(sessionID, self)
}
// Get the chipset_info field of the given host.
func (_class HostClass) GetChipsetInfo(sessionID SessionRef, self HostRef) (_retval map[string]string, _err error) {
	if (IsMock) {
		return _class.GetChipsetInfoMock(sessionID, self)
	}	
	_method := "host.get_chipset_info"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var HostClass_GetLocalCacheSrMockedCallback = func (sessionID SessionRef, self HostRef) (_retval SRRef, _err error) {
	log.Println("Host.GetLocalCacheSr not mocked")
	_err = errors.New("Host.GetLocalCacheSr not mocked")
	return
}

func (_class HostClass) GetLocalCacheSrMock(sessionID SessionRef, self HostRef) (_retval SRRef, _err error) {
	return HostClass_GetLocalCacheSrMockedCallback(sessionID, self)
}
// Get the local_cache_sr field of the given host.
func (_class HostClass) GetLocalCacheSr(sessionID SessionRef, self HostRef) (_retval SRRef, _err error) {
	if (IsMock) {
		return _class.GetLocalCacheSrMock(sessionID, self)
	}	
	_method := "host.get_local_cache_sr"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var HostClass_GetPowerOnConfigMockedCallback = func (sessionID SessionRef, self HostRef) (_retval map[string]string, _err error) {
	log.Println("Host.GetPowerOnConfig not mocked")
	_err = errors.New("Host.GetPowerOnConfig not mocked")
	return
}

func (_class HostClass) GetPowerOnConfigMock(sessionID SessionRef, self HostRef) (_retval map[string]string, _err error) {
	return HostClass_GetPowerOnConfigMockedCallback(sessionID, self)
}
// Get the power_on_config field of the given host.
func (_class HostClass) GetPowerOnConfig(sessionID SessionRef, self HostRef) (_retval map[string]string, _err error) {
	if (IsMock) {
		return _class.GetPowerOnConfigMock(sessionID, self)
	}	
	_method := "host.get_power_on_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var HostClass_GetPowerOnModeMockedCallback = func (sessionID SessionRef, self HostRef) (_retval string, _err error) {
	log.Println("Host.GetPowerOnMode not mocked")
	_err = errors.New("Host.GetPowerOnMode not mocked")
	return
}

func (_class HostClass) GetPowerOnModeMock(sessionID SessionRef, self HostRef) (_retval string, _err error) {
	return HostClass_GetPowerOnModeMockedCallback(sessionID, self)
}
// Get the power_on_mode field of the given host.
func (_class HostClass) GetPowerOnMode(sessionID SessionRef, self HostRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetPowerOnModeMock(sessionID, self)
	}	
	_method := "host.get_power_on_mode"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var HostClass_GetBiosStringsMockedCallback = func (sessionID SessionRef, self HostRef) (_retval map[string]string, _err error) {
	log.Println("Host.GetBiosStrings not mocked")
	_err = errors.New("Host.GetBiosStrings not mocked")
	return
}

func (_class HostClass) GetBiosStringsMock(sessionID SessionRef, self HostRef) (_retval map[string]string, _err error) {
	return HostClass_GetBiosStringsMockedCallback(sessionID, self)
}
// Get the bios_strings field of the given host.
func (_class HostClass) GetBiosStrings(sessionID SessionRef, self HostRef) (_retval map[string]string, _err error) {
	if (IsMock) {
		return _class.GetBiosStringsMock(sessionID, self)
	}	
	_method := "host.get_bios_strings"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var HostClass_GetLicenseServerMockedCallback = func (sessionID SessionRef, self HostRef) (_retval map[string]string, _err error) {
	log.Println("Host.GetLicenseServer not mocked")
	_err = errors.New("Host.GetLicenseServer not mocked")
	return
}

func (_class HostClass) GetLicenseServerMock(sessionID SessionRef, self HostRef) (_retval map[string]string, _err error) {
	return HostClass_GetLicenseServerMockedCallback(sessionID, self)
}
// Get the license_server field of the given host.
func (_class HostClass) GetLicenseServer(sessionID SessionRef, self HostRef) (_retval map[string]string, _err error) {
	if (IsMock) {
		return _class.GetLicenseServerMock(sessionID, self)
	}	
	_method := "host.get_license_server"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var HostClass_GetEditionMockedCallback = func (sessionID SessionRef, self HostRef) (_retval string, _err error) {
	log.Println("Host.GetEdition not mocked")
	_err = errors.New("Host.GetEdition not mocked")
	return
}

func (_class HostClass) GetEditionMock(sessionID SessionRef, self HostRef) (_retval string, _err error) {
	return HostClass_GetEditionMockedCallback(sessionID, self)
}
// Get the edition field of the given host.
func (_class HostClass) GetEdition(sessionID SessionRef, self HostRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetEditionMock(sessionID, self)
	}	
	_method := "host.get_edition"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var HostClass_GetExternalAuthConfigurationMockedCallback = func (sessionID SessionRef, self HostRef) (_retval map[string]string, _err error) {
	log.Println("Host.GetExternalAuthConfiguration not mocked")
	_err = errors.New("Host.GetExternalAuthConfiguration not mocked")
	return
}

func (_class HostClass) GetExternalAuthConfigurationMock(sessionID SessionRef, self HostRef) (_retval map[string]string, _err error) {
	return HostClass_GetExternalAuthConfigurationMockedCallback(sessionID, self)
}
// Get the external_auth_configuration field of the given host.
func (_class HostClass) GetExternalAuthConfiguration(sessionID SessionRef, self HostRef) (_retval map[string]string, _err error) {
	if (IsMock) {
		return _class.GetExternalAuthConfigurationMock(sessionID, self)
	}	
	_method := "host.get_external_auth_configuration"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var HostClass_GetExternalAuthServiceNameMockedCallback = func (sessionID SessionRef, self HostRef) (_retval string, _err error) {
	log.Println("Host.GetExternalAuthServiceName not mocked")
	_err = errors.New("Host.GetExternalAuthServiceName not mocked")
	return
}

func (_class HostClass) GetExternalAuthServiceNameMock(sessionID SessionRef, self HostRef) (_retval string, _err error) {
	return HostClass_GetExternalAuthServiceNameMockedCallback(sessionID, self)
}
// Get the external_auth_service_name field of the given host.
func (_class HostClass) GetExternalAuthServiceName(sessionID SessionRef, self HostRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetExternalAuthServiceNameMock(sessionID, self)
	}	
	_method := "host.get_external_auth_service_name"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var HostClass_GetExternalAuthTypeMockedCallback = func (sessionID SessionRef, self HostRef) (_retval string, _err error) {
	log.Println("Host.GetExternalAuthType not mocked")
	_err = errors.New("Host.GetExternalAuthType not mocked")
	return
}

func (_class HostClass) GetExternalAuthTypeMock(sessionID SessionRef, self HostRef) (_retval string, _err error) {
	return HostClass_GetExternalAuthTypeMockedCallback(sessionID, self)
}
// Get the external_auth_type field of the given host.
func (_class HostClass) GetExternalAuthType(sessionID SessionRef, self HostRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetExternalAuthTypeMock(sessionID, self)
	}	
	_method := "host.get_external_auth_type"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var HostClass_GetTagsMockedCallback = func (sessionID SessionRef, self HostRef) (_retval []string, _err error) {
	log.Println("Host.GetTags not mocked")
	_err = errors.New("Host.GetTags not mocked")
	return
}

func (_class HostClass) GetTagsMock(sessionID SessionRef, self HostRef) (_retval []string, _err error) {
	return HostClass_GetTagsMockedCallback(sessionID, self)
}
// Get the tags field of the given host.
func (_class HostClass) GetTags(sessionID SessionRef, self HostRef) (_retval []string, _err error) {
	if (IsMock) {
		return _class.GetTagsMock(sessionID, self)
	}	
	_method := "host.get_tags"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var HostClass_GetBlobsMockedCallback = func (sessionID SessionRef, self HostRef) (_retval map[string]BlobRef, _err error) {
	log.Println("Host.GetBlobs not mocked")
	_err = errors.New("Host.GetBlobs not mocked")
	return
}

func (_class HostClass) GetBlobsMock(sessionID SessionRef, self HostRef) (_retval map[string]BlobRef, _err error) {
	return HostClass_GetBlobsMockedCallback(sessionID, self)
}
// Get the blobs field of the given host.
func (_class HostClass) GetBlobs(sessionID SessionRef, self HostRef) (_retval map[string]BlobRef, _err error) {
	if (IsMock) {
		return _class.GetBlobsMock(sessionID, self)
	}	
	_method := "host.get_blobs"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var HostClass_GetHaNetworkPeersMockedCallback = func (sessionID SessionRef, self HostRef) (_retval []string, _err error) {
	log.Println("Host.GetHaNetworkPeers not mocked")
	_err = errors.New("Host.GetHaNetworkPeers not mocked")
	return
}

func (_class HostClass) GetHaNetworkPeersMock(sessionID SessionRef, self HostRef) (_retval []string, _err error) {
	return HostClass_GetHaNetworkPeersMockedCallback(sessionID, self)
}
// Get the ha_network_peers field of the given host.
func (_class HostClass) GetHaNetworkPeers(sessionID SessionRef, self HostRef) (_retval []string, _err error) {
	if (IsMock) {
		return _class.GetHaNetworkPeersMock(sessionID, self)
	}	
	_method := "host.get_ha_network_peers"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var HostClass_GetHaStatefilesMockedCallback = func (sessionID SessionRef, self HostRef) (_retval []string, _err error) {
	log.Println("Host.GetHaStatefiles not mocked")
	_err = errors.New("Host.GetHaStatefiles not mocked")
	return
}

func (_class HostClass) GetHaStatefilesMock(sessionID SessionRef, self HostRef) (_retval []string, _err error) {
	return HostClass_GetHaStatefilesMockedCallback(sessionID, self)
}
// Get the ha_statefiles field of the given host.
func (_class HostClass) GetHaStatefiles(sessionID SessionRef, self HostRef) (_retval []string, _err error) {
	if (IsMock) {
		return _class.GetHaStatefilesMock(sessionID, self)
	}	
	_method := "host.get_ha_statefiles"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var HostClass_GetLicenseParamsMockedCallback = func (sessionID SessionRef, self HostRef) (_retval map[string]string, _err error) {
	log.Println("Host.GetLicenseParams not mocked")
	_err = errors.New("Host.GetLicenseParams not mocked")
	return
}

func (_class HostClass) GetLicenseParamsMock(sessionID SessionRef, self HostRef) (_retval map[string]string, _err error) {
	return HostClass_GetLicenseParamsMockedCallback(sessionID, self)
}
// Get the license_params field of the given host.
func (_class HostClass) GetLicenseParams(sessionID SessionRef, self HostRef) (_retval map[string]string, _err error) {
	if (IsMock) {
		return _class.GetLicenseParamsMock(sessionID, self)
	}	
	_method := "host.get_license_params"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var HostClass_GetMetricsMockedCallback = func (sessionID SessionRef, self HostRef) (_retval HostMetricsRef, _err error) {
	log.Println("Host.GetMetrics not mocked")
	_err = errors.New("Host.GetMetrics not mocked")
	return
}

func (_class HostClass) GetMetricsMock(sessionID SessionRef, self HostRef) (_retval HostMetricsRef, _err error) {
	return HostClass_GetMetricsMockedCallback(sessionID, self)
}
// Get the metrics field of the given host.
func (_class HostClass) GetMetrics(sessionID SessionRef, self HostRef) (_retval HostMetricsRef, _err error) {
	if (IsMock) {
		return _class.GetMetricsMock(sessionID, self)
	}	
	_method := "host.get_metrics"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertHostMetricsRefToGo(_method + " -> ", _result.Value)
	return
}


var HostClass_GetAddressMockedCallback = func (sessionID SessionRef, self HostRef) (_retval string, _err error) {
	log.Println("Host.GetAddress not mocked")
	_err = errors.New("Host.GetAddress not mocked")
	return
}

func (_class HostClass) GetAddressMock(sessionID SessionRef, self HostRef) (_retval string, _err error) {
	return HostClass_GetAddressMockedCallback(sessionID, self)
}
// Get the address field of the given host.
func (_class HostClass) GetAddress(sessionID SessionRef, self HostRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetAddressMock(sessionID, self)
	}	
	_method := "host.get_address"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var HostClass_GetHostnameMockedCallback = func (sessionID SessionRef, self HostRef) (_retval string, _err error) {
	log.Println("Host.GetHostname not mocked")
	_err = errors.New("Host.GetHostname not mocked")
	return
}

func (_class HostClass) GetHostnameMock(sessionID SessionRef, self HostRef) (_retval string, _err error) {
	return HostClass_GetHostnameMockedCallback(sessionID, self)
}
// Get the hostname field of the given host.
func (_class HostClass) GetHostname(sessionID SessionRef, self HostRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetHostnameMock(sessionID, self)
	}	
	_method := "host.get_hostname"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var HostClass_GetCPUInfoMockedCallback = func (sessionID SessionRef, self HostRef) (_retval map[string]string, _err error) {
	log.Println("Host.GetCPUInfo not mocked")
	_err = errors.New("Host.GetCPUInfo not mocked")
	return
}

func (_class HostClass) GetCPUInfoMock(sessionID SessionRef, self HostRef) (_retval map[string]string, _err error) {
	return HostClass_GetCPUInfoMockedCallback(sessionID, self)
}
// Get the cpu_info field of the given host.
func (_class HostClass) GetCPUInfo(sessionID SessionRef, self HostRef) (_retval map[string]string, _err error) {
	if (IsMock) {
		return _class.GetCPUInfoMock(sessionID, self)
	}	
	_method := "host.get_cpu_info"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var HostClass_GetHostCPUsMockedCallback = func (sessionID SessionRef, self HostRef) (_retval []HostCPURef, _err error) {
	log.Println("Host.GetHostCPUs not mocked")
	_err = errors.New("Host.GetHostCPUs not mocked")
	return
}

func (_class HostClass) GetHostCPUsMock(sessionID SessionRef, self HostRef) (_retval []HostCPURef, _err error) {
	return HostClass_GetHostCPUsMockedCallback(sessionID, self)
}
// Get the host_CPUs field of the given host.
func (_class HostClass) GetHostCPUs(sessionID SessionRef, self HostRef) (_retval []HostCPURef, _err error) {
	if (IsMock) {
		return _class.GetHostCPUsMock(sessionID, self)
	}	
	_method := "host.get_host_CPUs"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertHostCPURefSetToGo(_method + " -> ", _result.Value)
	return
}


var HostClass_GetPBDsMockedCallback = func (sessionID SessionRef, self HostRef) (_retval []PBDRef, _err error) {
	log.Println("Host.GetPBDs not mocked")
	_err = errors.New("Host.GetPBDs not mocked")
	return
}

func (_class HostClass) GetPBDsMock(sessionID SessionRef, self HostRef) (_retval []PBDRef, _err error) {
	return HostClass_GetPBDsMockedCallback(sessionID, self)
}
// Get the PBDs field of the given host.
func (_class HostClass) GetPBDs(sessionID SessionRef, self HostRef) (_retval []PBDRef, _err error) {
	if (IsMock) {
		return _class.GetPBDsMock(sessionID, self)
	}	
	_method := "host.get_PBDs"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertPBDRefSetToGo(_method + " -> ", _result.Value)
	return
}


var HostClass_GetUpdatesMockedCallback = func (sessionID SessionRef, self HostRef) (_retval []PoolUpdateRef, _err error) {
	log.Println("Host.GetUpdates not mocked")
	_err = errors.New("Host.GetUpdates not mocked")
	return
}

func (_class HostClass) GetUpdatesMock(sessionID SessionRef, self HostRef) (_retval []PoolUpdateRef, _err error) {
	return HostClass_GetUpdatesMockedCallback(sessionID, self)
}
// Get the updates field of the given host.
func (_class HostClass) GetUpdates(sessionID SessionRef, self HostRef) (_retval []PoolUpdateRef, _err error) {
	if (IsMock) {
		return _class.GetUpdatesMock(sessionID, self)
	}	
	_method := "host.get_updates"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertPoolUpdateRefSetToGo(_method + " -> ", _result.Value)
	return
}


var HostClass_GetPatchesMockedCallback = func (sessionID SessionRef, self HostRef) (_retval []HostPatchRef, _err error) {
	log.Println("Host.GetPatches not mocked")
	_err = errors.New("Host.GetPatches not mocked")
	return
}

func (_class HostClass) GetPatchesMock(sessionID SessionRef, self HostRef) (_retval []HostPatchRef, _err error) {
	return HostClass_GetPatchesMockedCallback(sessionID, self)
}
// Get the patches field of the given host.
func (_class HostClass) GetPatches(sessionID SessionRef, self HostRef) (_retval []HostPatchRef, _err error) {
	if (IsMock) {
		return _class.GetPatchesMock(sessionID, self)
	}	
	_method := "host.get_patches"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertHostPatchRefSetToGo(_method + " -> ", _result.Value)
	return
}


var HostClass_GetCrashdumpsMockedCallback = func (sessionID SessionRef, self HostRef) (_retval []HostCrashdumpRef, _err error) {
	log.Println("Host.GetCrashdumps not mocked")
	_err = errors.New("Host.GetCrashdumps not mocked")
	return
}

func (_class HostClass) GetCrashdumpsMock(sessionID SessionRef, self HostRef) (_retval []HostCrashdumpRef, _err error) {
	return HostClass_GetCrashdumpsMockedCallback(sessionID, self)
}
// Get the crashdumps field of the given host.
func (_class HostClass) GetCrashdumps(sessionID SessionRef, self HostRef) (_retval []HostCrashdumpRef, _err error) {
	if (IsMock) {
		return _class.GetCrashdumpsMock(sessionID, self)
	}	
	_method := "host.get_crashdumps"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertHostCrashdumpRefSetToGo(_method + " -> ", _result.Value)
	return
}


var HostClass_GetCrashDumpSrMockedCallback = func (sessionID SessionRef, self HostRef) (_retval SRRef, _err error) {
	log.Println("Host.GetCrashDumpSr not mocked")
	_err = errors.New("Host.GetCrashDumpSr not mocked")
	return
}

func (_class HostClass) GetCrashDumpSrMock(sessionID SessionRef, self HostRef) (_retval SRRef, _err error) {
	return HostClass_GetCrashDumpSrMockedCallback(sessionID, self)
}
// Get the crash_dump_sr field of the given host.
func (_class HostClass) GetCrashDumpSr(sessionID SessionRef, self HostRef) (_retval SRRef, _err error) {
	if (IsMock) {
		return _class.GetCrashDumpSrMock(sessionID, self)
	}	
	_method := "host.get_crash_dump_sr"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var HostClass_GetSuspendImageSrMockedCallback = func (sessionID SessionRef, self HostRef) (_retval SRRef, _err error) {
	log.Println("Host.GetSuspendImageSr not mocked")
	_err = errors.New("Host.GetSuspendImageSr not mocked")
	return
}

func (_class HostClass) GetSuspendImageSrMock(sessionID SessionRef, self HostRef) (_retval SRRef, _err error) {
	return HostClass_GetSuspendImageSrMockedCallback(sessionID, self)
}
// Get the suspend_image_sr field of the given host.
func (_class HostClass) GetSuspendImageSr(sessionID SessionRef, self HostRef) (_retval SRRef, _err error) {
	if (IsMock) {
		return _class.GetSuspendImageSrMock(sessionID, self)
	}	
	_method := "host.get_suspend_image_sr"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var HostClass_GetPIFsMockedCallback = func (sessionID SessionRef, self HostRef) (_retval []PIFRef, _err error) {
	log.Println("Host.GetPIFs not mocked")
	_err = errors.New("Host.GetPIFs not mocked")
	return
}

func (_class HostClass) GetPIFsMock(sessionID SessionRef, self HostRef) (_retval []PIFRef, _err error) {
	return HostClass_GetPIFsMockedCallback(sessionID, self)
}
// Get the PIFs field of the given host.
func (_class HostClass) GetPIFs(sessionID SessionRef, self HostRef) (_retval []PIFRef, _err error) {
	if (IsMock) {
		return _class.GetPIFsMock(sessionID, self)
	}	
	_method := "host.get_PIFs"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertPIFRefSetToGo(_method + " -> ", _result.Value)
	return
}


var HostClass_GetLoggingMockedCallback = func (sessionID SessionRef, self HostRef) (_retval map[string]string, _err error) {
	log.Println("Host.GetLogging not mocked")
	_err = errors.New("Host.GetLogging not mocked")
	return
}

func (_class HostClass) GetLoggingMock(sessionID SessionRef, self HostRef) (_retval map[string]string, _err error) {
	return HostClass_GetLoggingMockedCallback(sessionID, self)
}
// Get the logging field of the given host.
func (_class HostClass) GetLogging(sessionID SessionRef, self HostRef) (_retval map[string]string, _err error) {
	if (IsMock) {
		return _class.GetLoggingMock(sessionID, self)
	}	
	_method := "host.get_logging"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var HostClass_GetResidentVMsMockedCallback = func (sessionID SessionRef, self HostRef) (_retval []VMRef, _err error) {
	log.Println("Host.GetResidentVMs not mocked")
	_err = errors.New("Host.GetResidentVMs not mocked")
	return
}

func (_class HostClass) GetResidentVMsMock(sessionID SessionRef, self HostRef) (_retval []VMRef, _err error) {
	return HostClass_GetResidentVMsMockedCallback(sessionID, self)
}
// Get the resident_VMs field of the given host.
func (_class HostClass) GetResidentVMs(sessionID SessionRef, self HostRef) (_retval []VMRef, _err error) {
	if (IsMock) {
		return _class.GetResidentVMsMock(sessionID, self)
	}	
	_method := "host.get_resident_VMs"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var HostClass_GetSupportedBootloadersMockedCallback = func (sessionID SessionRef, self HostRef) (_retval []string, _err error) {
	log.Println("Host.GetSupportedBootloaders not mocked")
	_err = errors.New("Host.GetSupportedBootloaders not mocked")
	return
}

func (_class HostClass) GetSupportedBootloadersMock(sessionID SessionRef, self HostRef) (_retval []string, _err error) {
	return HostClass_GetSupportedBootloadersMockedCallback(sessionID, self)
}
// Get the supported_bootloaders field of the given host.
func (_class HostClass) GetSupportedBootloaders(sessionID SessionRef, self HostRef) (_retval []string, _err error) {
	if (IsMock) {
		return _class.GetSupportedBootloadersMock(sessionID, self)
	}	
	_method := "host.get_supported_bootloaders"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var HostClass_GetSchedPolicyMockedCallback = func (sessionID SessionRef, self HostRef) (_retval string, _err error) {
	log.Println("Host.GetSchedPolicy not mocked")
	_err = errors.New("Host.GetSchedPolicy not mocked")
	return
}

func (_class HostClass) GetSchedPolicyMock(sessionID SessionRef, self HostRef) (_retval string, _err error) {
	return HostClass_GetSchedPolicyMockedCallback(sessionID, self)
}
// Get the sched_policy field of the given host.
func (_class HostClass) GetSchedPolicy(sessionID SessionRef, self HostRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetSchedPolicyMock(sessionID, self)
	}	
	_method := "host.get_sched_policy"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var HostClass_GetCPUConfigurationMockedCallback = func (sessionID SessionRef, self HostRef) (_retval map[string]string, _err error) {
	log.Println("Host.GetCPUConfiguration not mocked")
	_err = errors.New("Host.GetCPUConfiguration not mocked")
	return
}

func (_class HostClass) GetCPUConfigurationMock(sessionID SessionRef, self HostRef) (_retval map[string]string, _err error) {
	return HostClass_GetCPUConfigurationMockedCallback(sessionID, self)
}
// Get the cpu_configuration field of the given host.
func (_class HostClass) GetCPUConfiguration(sessionID SessionRef, self HostRef) (_retval map[string]string, _err error) {
	if (IsMock) {
		return _class.GetCPUConfigurationMock(sessionID, self)
	}	
	_method := "host.get_cpu_configuration"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var HostClass_GetCapabilitiesMockedCallback = func (sessionID SessionRef, self HostRef) (_retval []string, _err error) {
	log.Println("Host.GetCapabilities not mocked")
	_err = errors.New("Host.GetCapabilities not mocked")
	return
}

func (_class HostClass) GetCapabilitiesMock(sessionID SessionRef, self HostRef) (_retval []string, _err error) {
	return HostClass_GetCapabilitiesMockedCallback(sessionID, self)
}
// Get the capabilities field of the given host.
func (_class HostClass) GetCapabilities(sessionID SessionRef, self HostRef) (_retval []string, _err error) {
	if (IsMock) {
		return _class.GetCapabilitiesMock(sessionID, self)
	}	
	_method := "host.get_capabilities"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var HostClass_GetOtherConfigMockedCallback = func (sessionID SessionRef, self HostRef) (_retval map[string]string, _err error) {
	log.Println("Host.GetOtherConfig not mocked")
	_err = errors.New("Host.GetOtherConfig not mocked")
	return
}

func (_class HostClass) GetOtherConfigMock(sessionID SessionRef, self HostRef) (_retval map[string]string, _err error) {
	return HostClass_GetOtherConfigMockedCallback(sessionID, self)
}
// Get the other_config field of the given host.
func (_class HostClass) GetOtherConfig(sessionID SessionRef, self HostRef) (_retval map[string]string, _err error) {
	if (IsMock) {
		return _class.GetOtherConfigMock(sessionID, self)
	}	
	_method := "host.get_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var HostClass_GetSoftwareVersionMockedCallback = func (sessionID SessionRef, self HostRef) (_retval map[string]string, _err error) {
	log.Println("Host.GetSoftwareVersion not mocked")
	_err = errors.New("Host.GetSoftwareVersion not mocked")
	return
}

func (_class HostClass) GetSoftwareVersionMock(sessionID SessionRef, self HostRef) (_retval map[string]string, _err error) {
	return HostClass_GetSoftwareVersionMockedCallback(sessionID, self)
}
// Get the software_version field of the given host.
func (_class HostClass) GetSoftwareVersion(sessionID SessionRef, self HostRef) (_retval map[string]string, _err error) {
	if (IsMock) {
		return _class.GetSoftwareVersionMock(sessionID, self)
	}	
	_method := "host.get_software_version"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var HostClass_GetEnabledMockedCallback = func (sessionID SessionRef, self HostRef) (_retval bool, _err error) {
	log.Println("Host.GetEnabled not mocked")
	_err = errors.New("Host.GetEnabled not mocked")
	return
}

func (_class HostClass) GetEnabledMock(sessionID SessionRef, self HostRef) (_retval bool, _err error) {
	return HostClass_GetEnabledMockedCallback(sessionID, self)
}
// Get the enabled field of the given host.
func (_class HostClass) GetEnabled(sessionID SessionRef, self HostRef) (_retval bool, _err error) {
	if (IsMock) {
		return _class.GetEnabledMock(sessionID, self)
	}	
	_method := "host.get_enabled"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var HostClass_GetAPIVersionVendorImplementationMockedCallback = func (sessionID SessionRef, self HostRef) (_retval map[string]string, _err error) {
	log.Println("Host.GetAPIVersionVendorImplementation not mocked")
	_err = errors.New("Host.GetAPIVersionVendorImplementation not mocked")
	return
}

func (_class HostClass) GetAPIVersionVendorImplementationMock(sessionID SessionRef, self HostRef) (_retval map[string]string, _err error) {
	return HostClass_GetAPIVersionVendorImplementationMockedCallback(sessionID, self)
}
// Get the API_version/vendor_implementation field of the given host.
func (_class HostClass) GetAPIVersionVendorImplementation(sessionID SessionRef, self HostRef) (_retval map[string]string, _err error) {
	if (IsMock) {
		return _class.GetAPIVersionVendorImplementationMock(sessionID, self)
	}	
	_method := "host.get_API_version_vendor_implementation"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var HostClass_GetAPIVersionVendorMockedCallback = func (sessionID SessionRef, self HostRef) (_retval string, _err error) {
	log.Println("Host.GetAPIVersionVendor not mocked")
	_err = errors.New("Host.GetAPIVersionVendor not mocked")
	return
}

func (_class HostClass) GetAPIVersionVendorMock(sessionID SessionRef, self HostRef) (_retval string, _err error) {
	return HostClass_GetAPIVersionVendorMockedCallback(sessionID, self)
}
// Get the API_version/vendor field of the given host.
func (_class HostClass) GetAPIVersionVendor(sessionID SessionRef, self HostRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetAPIVersionVendorMock(sessionID, self)
	}	
	_method := "host.get_API_version_vendor"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var HostClass_GetAPIVersionMinorMockedCallback = func (sessionID SessionRef, self HostRef) (_retval int, _err error) {
	log.Println("Host.GetAPIVersionMinor not mocked")
	_err = errors.New("Host.GetAPIVersionMinor not mocked")
	return
}

func (_class HostClass) GetAPIVersionMinorMock(sessionID SessionRef, self HostRef) (_retval int, _err error) {
	return HostClass_GetAPIVersionMinorMockedCallback(sessionID, self)
}
// Get the API_version/minor field of the given host.
func (_class HostClass) GetAPIVersionMinor(sessionID SessionRef, self HostRef) (_retval int, _err error) {
	if (IsMock) {
		return _class.GetAPIVersionMinorMock(sessionID, self)
	}	
	_method := "host.get_API_version_minor"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var HostClass_GetAPIVersionMajorMockedCallback = func (sessionID SessionRef, self HostRef) (_retval int, _err error) {
	log.Println("Host.GetAPIVersionMajor not mocked")
	_err = errors.New("Host.GetAPIVersionMajor not mocked")
	return
}

func (_class HostClass) GetAPIVersionMajorMock(sessionID SessionRef, self HostRef) (_retval int, _err error) {
	return HostClass_GetAPIVersionMajorMockedCallback(sessionID, self)
}
// Get the API_version/major field of the given host.
func (_class HostClass) GetAPIVersionMajor(sessionID SessionRef, self HostRef) (_retval int, _err error) {
	if (IsMock) {
		return _class.GetAPIVersionMajorMock(sessionID, self)
	}	
	_method := "host.get_API_version_major"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var HostClass_GetCurrentOperationsMockedCallback = func (sessionID SessionRef, self HostRef) (_retval map[string]HostAllowedOperations, _err error) {
	log.Println("Host.GetCurrentOperations not mocked")
	_err = errors.New("Host.GetCurrentOperations not mocked")
	return
}

func (_class HostClass) GetCurrentOperationsMock(sessionID SessionRef, self HostRef) (_retval map[string]HostAllowedOperations, _err error) {
	return HostClass_GetCurrentOperationsMockedCallback(sessionID, self)
}
// Get the current_operations field of the given host.
func (_class HostClass) GetCurrentOperations(sessionID SessionRef, self HostRef) (_retval map[string]HostAllowedOperations, _err error) {
	if (IsMock) {
		return _class.GetCurrentOperationsMock(sessionID, self)
	}	
	_method := "host.get_current_operations"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringToEnumHostAllowedOperationsMapToGo(_method + " -> ", _result.Value)
	return
}


var HostClass_GetAllowedOperationsMockedCallback = func (sessionID SessionRef, self HostRef) (_retval []HostAllowedOperations, _err error) {
	log.Println("Host.GetAllowedOperations not mocked")
	_err = errors.New("Host.GetAllowedOperations not mocked")
	return
}

func (_class HostClass) GetAllowedOperationsMock(sessionID SessionRef, self HostRef) (_retval []HostAllowedOperations, _err error) {
	return HostClass_GetAllowedOperationsMockedCallback(sessionID, self)
}
// Get the allowed_operations field of the given host.
func (_class HostClass) GetAllowedOperations(sessionID SessionRef, self HostRef) (_retval []HostAllowedOperations, _err error) {
	if (IsMock) {
		return _class.GetAllowedOperationsMock(sessionID, self)
	}	
	_method := "host.get_allowed_operations"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertEnumHostAllowedOperationsSetToGo(_method + " -> ", _result.Value)
	return
}


var HostClass_GetMemoryOverheadMockedCallback = func (sessionID SessionRef, self HostRef) (_retval int, _err error) {
	log.Println("Host.GetMemoryOverhead not mocked")
	_err = errors.New("Host.GetMemoryOverhead not mocked")
	return
}

func (_class HostClass) GetMemoryOverheadMock(sessionID SessionRef, self HostRef) (_retval int, _err error) {
	return HostClass_GetMemoryOverheadMockedCallback(sessionID, self)
}
// Get the memory/overhead field of the given host.
func (_class HostClass) GetMemoryOverhead(sessionID SessionRef, self HostRef) (_retval int, _err error) {
	if (IsMock) {
		return _class.GetMemoryOverheadMock(sessionID, self)
	}	
	_method := "host.get_memory_overhead"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var HostClass_GetNameDescriptionMockedCallback = func (sessionID SessionRef, self HostRef) (_retval string, _err error) {
	log.Println("Host.GetNameDescription not mocked")
	_err = errors.New("Host.GetNameDescription not mocked")
	return
}

func (_class HostClass) GetNameDescriptionMock(sessionID SessionRef, self HostRef) (_retval string, _err error) {
	return HostClass_GetNameDescriptionMockedCallback(sessionID, self)
}
// Get the name/description field of the given host.
func (_class HostClass) GetNameDescription(sessionID SessionRef, self HostRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetNameDescriptionMock(sessionID, self)
	}	
	_method := "host.get_name_description"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var HostClass_GetNameLabelMockedCallback = func (sessionID SessionRef, self HostRef) (_retval string, _err error) {
	log.Println("Host.GetNameLabel not mocked")
	_err = errors.New("Host.GetNameLabel not mocked")
	return
}

func (_class HostClass) GetNameLabelMock(sessionID SessionRef, self HostRef) (_retval string, _err error) {
	return HostClass_GetNameLabelMockedCallback(sessionID, self)
}
// Get the name/label field of the given host.
func (_class HostClass) GetNameLabel(sessionID SessionRef, self HostRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetNameLabelMock(sessionID, self)
	}	
	_method := "host.get_name_label"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var HostClass_GetUUIDMockedCallback = func (sessionID SessionRef, self HostRef) (_retval string, _err error) {
	log.Println("Host.GetUUID not mocked")
	_err = errors.New("Host.GetUUID not mocked")
	return
}

func (_class HostClass) GetUUIDMock(sessionID SessionRef, self HostRef) (_retval string, _err error) {
	return HostClass_GetUUIDMockedCallback(sessionID, self)
}
// Get the uuid field of the given host.
func (_class HostClass) GetUUID(sessionID SessionRef, self HostRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetUUIDMock(sessionID, self)
	}	
	_method := "host.get_uuid"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var HostClass_GetByNameLabelMockedCallback = func (sessionID SessionRef, label string) (_retval []HostRef, _err error) {
	log.Println("Host.GetByNameLabel not mocked")
	_err = errors.New("Host.GetByNameLabel not mocked")
	return
}

func (_class HostClass) GetByNameLabelMock(sessionID SessionRef, label string) (_retval []HostRef, _err error) {
	return HostClass_GetByNameLabelMockedCallback(sessionID, label)
}
// Get all the host instances with the given label.
func (_class HostClass) GetByNameLabel(sessionID SessionRef, label string) (_retval []HostRef, _err error) {
	if (IsMock) {
		return _class.GetByNameLabelMock(sessionID, label)
	}	
	_method := "host.get_by_name_label"
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
	_retval, _err = convertHostRefSetToGo(_method + " -> ", _result.Value)
	return
}


var HostClass_GetByUUIDMockedCallback = func (sessionID SessionRef, uuid string) (_retval HostRef, _err error) {
	log.Println("Host.GetByUUID not mocked")
	_err = errors.New("Host.GetByUUID not mocked")
	return
}

func (_class HostClass) GetByUUIDMock(sessionID SessionRef, uuid string) (_retval HostRef, _err error) {
	return HostClass_GetByUUIDMockedCallback(sessionID, uuid)
}
// Get a reference to the host instance with the specified UUID.
func (_class HostClass) GetByUUID(sessionID SessionRef, uuid string) (_retval HostRef, _err error) {
	if (IsMock) {
		return _class.GetByUUIDMock(sessionID, uuid)
	}	
	_method := "host.get_by_uuid"
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
	_retval, _err = convertHostRefToGo(_method + " -> ", _result.Value)
	return
}


var HostClass_GetRecordMockedCallback = func (sessionID SessionRef, self HostRef) (_retval HostRecord, _err error) {
	log.Println("Host.GetRecord not mocked")
	_err = errors.New("Host.GetRecord not mocked")
	return
}

func (_class HostClass) GetRecordMock(sessionID SessionRef, self HostRef) (_retval HostRecord, _err error) {
	return HostClass_GetRecordMockedCallback(sessionID, self)
}
// Get a record containing the current state of the given host.
func (_class HostClass) GetRecord(sessionID SessionRef, self HostRef) (_retval HostRecord, _err error) {
	if (IsMock) {
		return _class.GetRecordMock(sessionID, self)
	}	
	_method := "host.get_record"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertHostRecordToGo(_method + " -> ", _result.Value)
	return
}
