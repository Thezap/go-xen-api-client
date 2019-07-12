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

func (_class HostClass) GetAllRecords__mock(sessionID SessionRef) (_retval map[HostRef]HostRecord, _err error) {
	log.Println("Host.GetAllRecords not mocked")
	_err = errors.New("Host.GetAllRecords not mocked")
	return
}
// Return a map of host references to host records for all hosts known to the system.
func (_class HostClass) GetAllRecords(sessionID SessionRef) (_retval map[HostRef]HostRecord, _err error) {
	if (IsMock) {
		return _class.GetAllRecords__mock(sessionID)
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

func (_class HostClass) GetAll__mock(sessionID SessionRef) (_retval []HostRef, _err error) {
	log.Println("Host.GetAll not mocked")
	_err = errors.New("Host.GetAll not mocked")
	return
}
// Return a list of all the hosts known to the system.
func (_class HostClass) GetAll(sessionID SessionRef) (_retval []HostRef, _err error) {
	if (IsMock) {
		return _class.GetAll__mock(sessionID)
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

func (_class HostClass) SetSslLegacy__mock(sessionID SessionRef, self HostRef, value bool) (_err error) {
	log.Println("Host.SetSslLegacy not mocked")
	_err = errors.New("Host.SetSslLegacy not mocked")
	return
}
// Enable/disable SSLv3 for interoperability with older versions of XenServer. When this is set to a different value, the host immediately restarts its SSL/TLS listening service; typically this takes less than a second but existing connections to it will be broken. XenAPI login sessions will remain valid.
func (_class HostClass) SetSslLegacy(sessionID SessionRef, self HostRef, value bool) (_err error) {
	if (IsMock) {
		return _class.SetSslLegacy__mock(sessionID, self, value)
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

func (_class HostClass) DisableDisplay__mock(sessionID SessionRef, host HostRef) (_retval HostDisplay, _err error) {
	log.Println("Host.DisableDisplay not mocked")
	_err = errors.New("Host.DisableDisplay not mocked")
	return
}
// Disable console output to the physical display device next time this host boots
func (_class HostClass) DisableDisplay(sessionID SessionRef, host HostRef) (_retval HostDisplay, _err error) {
	if (IsMock) {
		return _class.DisableDisplay__mock(sessionID, host)
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

func (_class HostClass) EnableDisplay__mock(sessionID SessionRef, host HostRef) (_retval HostDisplay, _err error) {
	log.Println("Host.EnableDisplay not mocked")
	_err = errors.New("Host.EnableDisplay not mocked")
	return
}
// Enable console output to the physical display device next time this host boots
func (_class HostClass) EnableDisplay(sessionID SessionRef, host HostRef) (_retval HostDisplay, _err error) {
	if (IsMock) {
		return _class.EnableDisplay__mock(sessionID, host)
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

func (_class HostClass) DeclareDead__mock(sessionID SessionRef, host HostRef) (_err error) {
	log.Println("Host.DeclareDead not mocked")
	_err = errors.New("Host.DeclareDead not mocked")
	return
}
// Declare that a host is dead. This is a dangerous operation, and should only be called if the administrator is absolutely sure the host is definitely dead
func (_class HostClass) DeclareDead(sessionID SessionRef, host HostRef) (_err error) {
	if (IsMock) {
		return _class.DeclareDead__mock(sessionID, host)
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

func (_class HostClass) MigrateReceive__mock(sessionID SessionRef, host HostRef, network NetworkRef, options map[string]string) (_retval map[string]string, _err error) {
	log.Println("Host.MigrateReceive not mocked")
	_err = errors.New("Host.MigrateReceive not mocked")
	return
}
// Prepare to receive a VM, returning a token which can be passed to VM.migrate.
func (_class HostClass) MigrateReceive(sessionID SessionRef, host HostRef, network NetworkRef, options map[string]string) (_retval map[string]string, _err error) {
	if (IsMock) {
		return _class.MigrateReceive__mock(sessionID, host, network, options)
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

func (_class HostClass) DisableLocalStorageCaching__mock(sessionID SessionRef, host HostRef) (_err error) {
	log.Println("Host.DisableLocalStorageCaching not mocked")
	_err = errors.New("Host.DisableLocalStorageCaching not mocked")
	return
}
// Disable the use of a local SR for caching purposes
func (_class HostClass) DisableLocalStorageCaching(sessionID SessionRef, host HostRef) (_err error) {
	if (IsMock) {
		return _class.DisableLocalStorageCaching__mock(sessionID, host)
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

func (_class HostClass) EnableLocalStorageCaching__mock(sessionID SessionRef, host HostRef, sr SRRef) (_err error) {
	log.Println("Host.EnableLocalStorageCaching not mocked")
	_err = errors.New("Host.EnableLocalStorageCaching not mocked")
	return
}
// Enable the use of a local SR for caching purposes
func (_class HostClass) EnableLocalStorageCaching(sessionID SessionRef, host HostRef, sr SRRef) (_err error) {
	if (IsMock) {
		return _class.EnableLocalStorageCaching__mock(sessionID, host, sr)
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

func (_class HostClass) ResetCPUFeatures__mock(sessionID SessionRef, host HostRef) (_err error) {
	log.Println("Host.ResetCPUFeatures not mocked")
	_err = errors.New("Host.ResetCPUFeatures not mocked")
	return
}
// Remove the feature mask, such that after a reboot all features of the CPU are enabled.
func (_class HostClass) ResetCPUFeatures(sessionID SessionRef, host HostRef) (_err error) {
	if (IsMock) {
		return _class.ResetCPUFeatures__mock(sessionID, host)
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

func (_class HostClass) SetCPUFeatures__mock(sessionID SessionRef, host HostRef, features string) (_err error) {
	log.Println("Host.SetCPUFeatures not mocked")
	_err = errors.New("Host.SetCPUFeatures not mocked")
	return
}
// Set the CPU features to be used after a reboot, if the given features string is valid.
func (_class HostClass) SetCPUFeatures(sessionID SessionRef, host HostRef, features string) (_err error) {
	if (IsMock) {
		return _class.SetCPUFeatures__mock(sessionID, host, features)
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

func (_class HostClass) SetPowerOnMode__mock(sessionID SessionRef, self HostRef, powerOnMode string, powerOnConfig map[string]string) (_err error) {
	log.Println("Host.SetPowerOnMode not mocked")
	_err = errors.New("Host.SetPowerOnMode not mocked")
	return
}
// Set the power-on-mode, host, user and password 
func (_class HostClass) SetPowerOnMode(sessionID SessionRef, self HostRef, powerOnMode string, powerOnConfig map[string]string) (_err error) {
	if (IsMock) {
		return _class.SetPowerOnMode__mock(sessionID, self, powerOnMode, powerOnConfig)
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

func (_class HostClass) RefreshPackInfo__mock(sessionID SessionRef, host HostRef) (_err error) {
	log.Println("Host.RefreshPackInfo not mocked")
	_err = errors.New("Host.RefreshPackInfo not mocked")
	return
}
// Refresh the list of installed Supplemental Packs.
func (_class HostClass) RefreshPackInfo(sessionID SessionRef, host HostRef) (_err error) {
	if (IsMock) {
		return _class.RefreshPackInfo__mock(sessionID, host)
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

func (_class HostClass) ApplyEdition__mock(sessionID SessionRef, host HostRef, edition string, force bool) (_err error) {
	log.Println("Host.ApplyEdition not mocked")
	_err = errors.New("Host.ApplyEdition not mocked")
	return
}
// Change to another edition, or reactivate the current edition after a license has expired. This may be subject to the successful checkout of an appropriate license.
func (_class HostClass) ApplyEdition(sessionID SessionRef, host HostRef, edition string, force bool) (_err error) {
	if (IsMock) {
		return _class.ApplyEdition__mock(sessionID, host, edition, force)
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

func (_class HostClass) GetServerCertificate__mock(sessionID SessionRef, host HostRef) (_retval string, _err error) {
	log.Println("Host.GetServerCertificate not mocked")
	_err = errors.New("Host.GetServerCertificate not mocked")
	return
}
// Get the installed server public TLS certificate.
func (_class HostClass) GetServerCertificate(sessionID SessionRef, host HostRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetServerCertificate__mock(sessionID, host)
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

func (_class HostClass) RetrieveWlbEvacuateRecommendations__mock(sessionID SessionRef, self HostRef) (_retval map[VMRef][]string, _err error) {
	log.Println("Host.RetrieveWlbEvacuateRecommendations not mocked")
	_err = errors.New("Host.RetrieveWlbEvacuateRecommendations not mocked")
	return
}
// Retrieves recommended host migrations to perform when evacuating the host from the wlb server. If a VM cannot be migrated from the host the reason is listed instead of a recommendation.
func (_class HostClass) RetrieveWlbEvacuateRecommendations(sessionID SessionRef, self HostRef) (_retval map[VMRef][]string, _err error) {
	if (IsMock) {
		return _class.RetrieveWlbEvacuateRecommendations__mock(sessionID, self)
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

func (_class HostClass) DisableExternalAuth__mock(sessionID SessionRef, host HostRef, config map[string]string) (_err error) {
	log.Println("Host.DisableExternalAuth not mocked")
	_err = errors.New("Host.DisableExternalAuth not mocked")
	return
}
// This call disables external authentication on the local host
func (_class HostClass) DisableExternalAuth(sessionID SessionRef, host HostRef, config map[string]string) (_err error) {
	if (IsMock) {
		return _class.DisableExternalAuth__mock(sessionID, host, config)
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

func (_class HostClass) EnableExternalAuth__mock(sessionID SessionRef, host HostRef, config map[string]string, serviceName string, authType string) (_err error) {
	log.Println("Host.EnableExternalAuth not mocked")
	_err = errors.New("Host.EnableExternalAuth not mocked")
	return
}
// This call enables external authentication on a host
func (_class HostClass) EnableExternalAuth(sessionID SessionRef, host HostRef, config map[string]string, serviceName string, authType string) (_err error) {
	if (IsMock) {
		return _class.EnableExternalAuth__mock(sessionID, host, config, serviceName, authType)
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

func (_class HostClass) GetServerLocaltime__mock(sessionID SessionRef, host HostRef) (_retval time.Time, _err error) {
	log.Println("Host.GetServerLocaltime not mocked")
	_err = errors.New("Host.GetServerLocaltime not mocked")
	return
}
// This call queries the host's clock for the current time in the host's local timezone
func (_class HostClass) GetServerLocaltime(sessionID SessionRef, host HostRef) (_retval time.Time, _err error) {
	if (IsMock) {
		return _class.GetServerLocaltime__mock(sessionID, host)
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

func (_class HostClass) GetServertime__mock(sessionID SessionRef, host HostRef) (_retval time.Time, _err error) {
	log.Println("Host.GetServertime not mocked")
	_err = errors.New("Host.GetServertime not mocked")
	return
}
// This call queries the host's clock for the current time
func (_class HostClass) GetServertime(sessionID SessionRef, host HostRef) (_retval time.Time, _err error) {
	if (IsMock) {
		return _class.GetServertime__mock(sessionID, host)
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

func (_class HostClass) CallExtension__mock(sessionID SessionRef, host HostRef, call string) (_retval string, _err error) {
	log.Println("Host.CallExtension not mocked")
	_err = errors.New("Host.CallExtension not mocked")
	return
}
// Call a XenAPI extension on this host
func (_class HostClass) CallExtension(sessionID SessionRef, host HostRef, call string) (_retval string, _err error) {
	if (IsMock) {
		return _class.CallExtension__mock(sessionID, host, call)
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

func (_class HostClass) HasExtension__mock(sessionID SessionRef, host HostRef, name string) (_retval bool, _err error) {
	log.Println("Host.HasExtension not mocked")
	_err = errors.New("Host.HasExtension not mocked")
	return
}
// Return true if the extension is available on the host
func (_class HostClass) HasExtension(sessionID SessionRef, host HostRef, name string) (_retval bool, _err error) {
	if (IsMock) {
		return _class.HasExtension__mock(sessionID, host, name)
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

func (_class HostClass) CallPlugin__mock(sessionID SessionRef, host HostRef, plugin string, fn string, args map[string]string) (_retval string, _err error) {
	log.Println("Host.CallPlugin not mocked")
	_err = errors.New("Host.CallPlugin not mocked")
	return
}
// Call a XenAPI plugin on this host
func (_class HostClass) CallPlugin(sessionID SessionRef, host HostRef, plugin string, fn string, args map[string]string) (_retval string, _err error) {
	if (IsMock) {
		return _class.CallPlugin__mock(sessionID, host, plugin, fn, args)
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

func (_class HostClass) CreateNewBlob__mock(sessionID SessionRef, host HostRef, name string, mimeType string, public bool) (_retval BlobRef, _err error) {
	log.Println("Host.CreateNewBlob not mocked")
	_err = errors.New("Host.CreateNewBlob not mocked")
	return
}
// Create a placeholder for a named binary blob of data that is associated with this host
func (_class HostClass) CreateNewBlob(sessionID SessionRef, host HostRef, name string, mimeType string, public bool) (_retval BlobRef, _err error) {
	if (IsMock) {
		return _class.CreateNewBlob__mock(sessionID, host, name, mimeType, public)
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

func (_class HostClass) BackupRrds__mock(sessionID SessionRef, host HostRef, delay float64) (_err error) {
	log.Println("Host.BackupRrds not mocked")
	_err = errors.New("Host.BackupRrds not mocked")
	return
}
// This causes the RRDs to be backed up to the master
func (_class HostClass) BackupRrds(sessionID SessionRef, host HostRef, delay float64) (_err error) {
	if (IsMock) {
		return _class.BackupRrds__mock(sessionID, host, delay)
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

func (_class HostClass) SyncData__mock(sessionID SessionRef, host HostRef) (_err error) {
	log.Println("Host.SyncData not mocked")
	_err = errors.New("Host.SyncData not mocked")
	return
}
// This causes the synchronisation of the non-database data (messages, RRDs and so on) stored on the master to be synchronised with the host
func (_class HostClass) SyncData(sessionID SessionRef, host HostRef) (_err error) {
	if (IsMock) {
		return _class.SyncData__mock(sessionID, host)
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

func (_class HostClass) ComputeMemoryOverhead__mock(sessionID SessionRef, host HostRef) (_retval int, _err error) {
	log.Println("Host.ComputeMemoryOverhead not mocked")
	_err = errors.New("Host.ComputeMemoryOverhead not mocked")
	return
}
// Computes the virtualization memory overhead of a host.
func (_class HostClass) ComputeMemoryOverhead(sessionID SessionRef, host HostRef) (_retval int, _err error) {
	if (IsMock) {
		return _class.ComputeMemoryOverhead__mock(sessionID, host)
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

func (_class HostClass) ComputeFreeMemory__mock(sessionID SessionRef, host HostRef) (_retval int, _err error) {
	log.Println("Host.ComputeFreeMemory not mocked")
	_err = errors.New("Host.ComputeFreeMemory not mocked")
	return
}
// Computes the amount of free memory on the host.
func (_class HostClass) ComputeFreeMemory(sessionID SessionRef, host HostRef) (_retval int, _err error) {
	if (IsMock) {
		return _class.ComputeFreeMemory__mock(sessionID, host)
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

func (_class HostClass) SetHostnameLive__mock(sessionID SessionRef, host HostRef, hostname string) (_err error) {
	log.Println("Host.SetHostnameLive not mocked")
	_err = errors.New("Host.SetHostnameLive not mocked")
	return
}
// Sets the host name to the specified string.  Both the API and lower-level system hostname are changed immediately.
//
// Errors:
//  HOST_NAME_INVALID - The host name is invalid.
func (_class HostClass) SetHostnameLive(sessionID SessionRef, host HostRef, hostname string) (_err error) {
	if (IsMock) {
		return _class.SetHostnameLive__mock(sessionID, host, hostname)
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

func (_class HostClass) ShutdownAgent__mock(sessionID SessionRef) (_err error) {
	log.Println("Host.ShutdownAgent not mocked")
	_err = errors.New("Host.ShutdownAgent not mocked")
	return
}
// Shuts the agent down after a 10 second pause. WARNING: this is a dangerous operation. Any operations in progress will be aborted, and unrecoverable data loss may occur. The caller is responsible for ensuring that there are no operations in progress when this method is called.
func (_class HostClass) ShutdownAgent(sessionID SessionRef) (_err error) {
	if (IsMock) {
		return _class.ShutdownAgent__mock(sessionID)
	}	
	_method := "host.shutdown_agent"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg)
	return
}

func (_class HostClass) RestartAgent__mock(sessionID SessionRef, host HostRef) (_err error) {
	log.Println("Host.RestartAgent not mocked")
	_err = errors.New("Host.RestartAgent not mocked")
	return
}
// Restarts the agent after a 10 second pause. WARNING: this is a dangerous operation. Any operations in progress will be aborted, and unrecoverable data loss may occur. The caller is responsible for ensuring that there are no operations in progress when this method is called.
func (_class HostClass) RestartAgent(sessionID SessionRef, host HostRef) (_err error) {
	if (IsMock) {
		return _class.RestartAgent__mock(sessionID, host)
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

func (_class HostClass) GetSystemStatusCapabilities__mock(sessionID SessionRef, host HostRef) (_retval string, _err error) {
	log.Println("Host.GetSystemStatusCapabilities not mocked")
	_err = errors.New("Host.GetSystemStatusCapabilities not mocked")
	return
}
// 
func (_class HostClass) GetSystemStatusCapabilities(sessionID SessionRef, host HostRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetSystemStatusCapabilities__mock(sessionID, host)
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

func (_class HostClass) GetManagementInterface__mock(sessionID SessionRef, host HostRef) (_retval PIFRef, _err error) {
	log.Println("Host.GetManagementInterface not mocked")
	_err = errors.New("Host.GetManagementInterface not mocked")
	return
}
// Returns the management interface for the specified host
func (_class HostClass) GetManagementInterface(sessionID SessionRef, host HostRef) (_retval PIFRef, _err error) {
	if (IsMock) {
		return _class.GetManagementInterface__mock(sessionID, host)
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

func (_class HostClass) ManagementDisable__mock(sessionID SessionRef) (_err error) {
	log.Println("Host.ManagementDisable not mocked")
	_err = errors.New("Host.ManagementDisable not mocked")
	return
}
// Disable the management network interface
func (_class HostClass) ManagementDisable(sessionID SessionRef) (_err error) {
	if (IsMock) {
		return _class.ManagementDisable__mock(sessionID)
	}	
	_method := "host.management_disable"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg)
	return
}

func (_class HostClass) LocalManagementReconfigure__mock(sessionID SessionRef, iface string) (_err error) {
	log.Println("Host.LocalManagementReconfigure not mocked")
	_err = errors.New("Host.LocalManagementReconfigure not mocked")
	return
}
// Reconfigure the management network interface. Should only be used if Host.management_reconfigure is impossible because the network configuration is broken.
func (_class HostClass) LocalManagementReconfigure(sessionID SessionRef, iface string) (_err error) {
	if (IsMock) {
		return _class.LocalManagementReconfigure__mock(sessionID, iface)
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

func (_class HostClass) ManagementReconfigure__mock(sessionID SessionRef, pif PIFRef) (_err error) {
	log.Println("Host.ManagementReconfigure not mocked")
	_err = errors.New("Host.ManagementReconfigure not mocked")
	return
}
// Reconfigure the management network interface
func (_class HostClass) ManagementReconfigure(sessionID SessionRef, pif PIFRef) (_err error) {
	if (IsMock) {
		return _class.ManagementReconfigure__mock(sessionID, pif)
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

func (_class HostClass) SyslogReconfigure__mock(sessionID SessionRef, host HostRef) (_err error) {
	log.Println("Host.SyslogReconfigure not mocked")
	_err = errors.New("Host.SyslogReconfigure not mocked")
	return
}
// Re-configure syslog logging
func (_class HostClass) SyslogReconfigure(sessionID SessionRef, host HostRef) (_err error) {
	if (IsMock) {
		return _class.SyslogReconfigure__mock(sessionID, host)
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

func (_class HostClass) Evacuate__mock(sessionID SessionRef, host HostRef) (_err error) {
	log.Println("Host.Evacuate not mocked")
	_err = errors.New("Host.Evacuate not mocked")
	return
}
// Migrate all VMs off of this host, where possible.
func (_class HostClass) Evacuate(sessionID SessionRef, host HostRef) (_err error) {
	if (IsMock) {
		return _class.Evacuate__mock(sessionID, host)
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

func (_class HostClass) GetUncooperativeResidentVMs__mock(sessionID SessionRef, self HostRef) (_retval []VMRef, _err error) {
	log.Println("Host.GetUncooperativeResidentVMs not mocked")
	_err = errors.New("Host.GetUncooperativeResidentVMs not mocked")
	return
}
// Return a set of VMs which are not co-operating with the host's memory control system
func (_class HostClass) GetUncooperativeResidentVMs(sessionID SessionRef, self HostRef) (_retval []VMRef, _err error) {
	if (IsMock) {
		return _class.GetUncooperativeResidentVMs__mock(sessionID, self)
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

func (_class HostClass) GetVmsWhichPreventEvacuation__mock(sessionID SessionRef, self HostRef) (_retval map[VMRef][]string, _err error) {
	log.Println("Host.GetVmsWhichPreventEvacuation not mocked")
	_err = errors.New("Host.GetVmsWhichPreventEvacuation not mocked")
	return
}
// Return a set of VMs which prevent the host being evacuated, with per-VM error codes
func (_class HostClass) GetVmsWhichPreventEvacuation(sessionID SessionRef, self HostRef) (_retval map[VMRef][]string, _err error) {
	if (IsMock) {
		return _class.GetVmsWhichPreventEvacuation__mock(sessionID, self)
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

func (_class HostClass) AssertCanEvacuate__mock(sessionID SessionRef, host HostRef) (_err error) {
	log.Println("Host.AssertCanEvacuate not mocked")
	_err = errors.New("Host.AssertCanEvacuate not mocked")
	return
}
// Check this host can be evacuated.
func (_class HostClass) AssertCanEvacuate(sessionID SessionRef, host HostRef) (_err error) {
	if (IsMock) {
		return _class.AssertCanEvacuate__mock(sessionID, host)
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

func (_class HostClass) ForgetDataSourceArchives__mock(sessionID SessionRef, host HostRef, dataSource string) (_err error) {
	log.Println("Host.ForgetDataSourceArchives not mocked")
	_err = errors.New("Host.ForgetDataSourceArchives not mocked")
	return
}
// Forget the recorded statistics related to the specified data source
func (_class HostClass) ForgetDataSourceArchives(sessionID SessionRef, host HostRef, dataSource string) (_err error) {
	if (IsMock) {
		return _class.ForgetDataSourceArchives__mock(sessionID, host, dataSource)
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

func (_class HostClass) QueryDataSource__mock(sessionID SessionRef, host HostRef, dataSource string) (_retval float64, _err error) {
	log.Println("Host.QueryDataSource not mocked")
	_err = errors.New("Host.QueryDataSource not mocked")
	return
}
// Query the latest value of the specified data source
func (_class HostClass) QueryDataSource(sessionID SessionRef, host HostRef, dataSource string) (_retval float64, _err error) {
	if (IsMock) {
		return _class.QueryDataSource__mock(sessionID, host, dataSource)
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

func (_class HostClass) RecordDataSource__mock(sessionID SessionRef, host HostRef, dataSource string) (_err error) {
	log.Println("Host.RecordDataSource not mocked")
	_err = errors.New("Host.RecordDataSource not mocked")
	return
}
// Start recording the specified data source
func (_class HostClass) RecordDataSource(sessionID SessionRef, host HostRef, dataSource string) (_err error) {
	if (IsMock) {
		return _class.RecordDataSource__mock(sessionID, host, dataSource)
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

func (_class HostClass) GetDataSources__mock(sessionID SessionRef, host HostRef) (_retval []DataSourceRecord, _err error) {
	log.Println("Host.GetDataSources not mocked")
	_err = errors.New("Host.GetDataSources not mocked")
	return
}
// 
func (_class HostClass) GetDataSources(sessionID SessionRef, host HostRef) (_retval []DataSourceRecord, _err error) {
	if (IsMock) {
		return _class.GetDataSources__mock(sessionID, host)
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

func (_class HostClass) EmergencyHaDisable__mock(sessionID SessionRef, soft bool) (_err error) {
	log.Println("Host.EmergencyHaDisable not mocked")
	_err = errors.New("Host.EmergencyHaDisable not mocked")
	return
}
// This call disables HA on the local host. This should only be used with extreme care.
func (_class HostClass) EmergencyHaDisable(sessionID SessionRef, soft bool) (_err error) {
	if (IsMock) {
		return _class.EmergencyHaDisable__mock(sessionID, soft)
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

func (_class HostClass) PowerOn__mock(sessionID SessionRef, host HostRef) (_err error) {
	log.Println("Host.PowerOn not mocked")
	_err = errors.New("Host.PowerOn not mocked")
	return
}
// Attempt to power-on the host (if the capability exists).
func (_class HostClass) PowerOn(sessionID SessionRef, host HostRef) (_err error) {
	if (IsMock) {
		return _class.PowerOn__mock(sessionID, host)
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

func (_class HostClass) Destroy__mock(sessionID SessionRef, self HostRef) (_err error) {
	log.Println("Host.Destroy not mocked")
	_err = errors.New("Host.Destroy not mocked")
	return
}
// Destroy specified host record in database
func (_class HostClass) Destroy(sessionID SessionRef, self HostRef) (_err error) {
	if (IsMock) {
		return _class.Destroy__mock(sessionID, self)
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

func (_class HostClass) LicenseRemove__mock(sessionID SessionRef, host HostRef) (_err error) {
	log.Println("Host.LicenseRemove not mocked")
	_err = errors.New("Host.LicenseRemove not mocked")
	return
}
// Remove any license file from the specified host, and switch that host to the unlicensed edition
func (_class HostClass) LicenseRemove(sessionID SessionRef, host HostRef) (_err error) {
	if (IsMock) {
		return _class.LicenseRemove__mock(sessionID, host)
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

func (_class HostClass) LicenseAdd__mock(sessionID SessionRef, host HostRef, contents string) (_err error) {
	log.Println("Host.LicenseAdd not mocked")
	_err = errors.New("Host.LicenseAdd not mocked")
	return
}
// Apply a new license to a host
//
// Errors:
//  LICENSE_PROCESSING_ERROR - There was an error processing your license.  Please contact your support representative.
func (_class HostClass) LicenseAdd(sessionID SessionRef, host HostRef, contents string) (_err error) {
	if (IsMock) {
		return _class.LicenseAdd__mock(sessionID, host, contents)
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

func (_class HostClass) LicenseApply__mock(sessionID SessionRef, host HostRef, contents string) (_err error) {
	log.Println("Host.LicenseApply not mocked")
	_err = errors.New("Host.LicenseApply not mocked")
	return
}
// Apply a new license to a host
//
// Errors:
//  LICENSE_PROCESSING_ERROR - There was an error processing your license.  Please contact your support representative.
func (_class HostClass) LicenseApply(sessionID SessionRef, host HostRef, contents string) (_err error) {
	if (IsMock) {
		return _class.LicenseApply__mock(sessionID, host, contents)
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

func (_class HostClass) ListMethods__mock(sessionID SessionRef) (_retval []string, _err error) {
	log.Println("Host.ListMethods not mocked")
	_err = errors.New("Host.ListMethods not mocked")
	return
}
// List all supported methods
func (_class HostClass) ListMethods(sessionID SessionRef) (_retval []string, _err error) {
	if (IsMock) {
		return _class.ListMethods__mock(sessionID)
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

func (_class HostClass) BugreportUpload__mock(sessionID SessionRef, host HostRef, url string, options map[string]string) (_err error) {
	log.Println("Host.BugreportUpload not mocked")
	_err = errors.New("Host.BugreportUpload not mocked")
	return
}
// Run xen-bugtool --yestoall and upload the output to support
func (_class HostClass) BugreportUpload(sessionID SessionRef, host HostRef, url string, options map[string]string) (_err error) {
	if (IsMock) {
		return _class.BugreportUpload__mock(sessionID, host, url, options)
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

func (_class HostClass) SendDebugKeys__mock(sessionID SessionRef, host HostRef, keys string) (_err error) {
	log.Println("Host.SendDebugKeys not mocked")
	_err = errors.New("Host.SendDebugKeys not mocked")
	return
}
// Inject the given string as debugging keys into Xen
func (_class HostClass) SendDebugKeys(sessionID SessionRef, host HostRef, keys string) (_err error) {
	if (IsMock) {
		return _class.SendDebugKeys__mock(sessionID, host, keys)
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

func (_class HostClass) GetLog__mock(sessionID SessionRef, host HostRef) (_retval string, _err error) {
	log.Println("Host.GetLog not mocked")
	_err = errors.New("Host.GetLog not mocked")
	return
}
// Get the host's log file
func (_class HostClass) GetLog(sessionID SessionRef, host HostRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetLog__mock(sessionID, host)
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

func (_class HostClass) DmesgClear__mock(sessionID SessionRef, host HostRef) (_retval string, _err error) {
	log.Println("Host.DmesgClear not mocked")
	_err = errors.New("Host.DmesgClear not mocked")
	return
}
// Get the host xen dmesg, and clear the buffer.
func (_class HostClass) DmesgClear(sessionID SessionRef, host HostRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.DmesgClear__mock(sessionID, host)
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

func (_class HostClass) Dmesg__mock(sessionID SessionRef, host HostRef) (_retval string, _err error) {
	log.Println("Host.Dmesg not mocked")
	_err = errors.New("Host.Dmesg not mocked")
	return
}
// Get the host xen dmesg.
func (_class HostClass) Dmesg(sessionID SessionRef, host HostRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.Dmesg__mock(sessionID, host)
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

func (_class HostClass) Reboot__mock(sessionID SessionRef, host HostRef) (_err error) {
	log.Println("Host.Reboot not mocked")
	_err = errors.New("Host.Reboot not mocked")
	return
}
// Reboot the host. (This function can only be called if there are no currently running VMs on the host and it is disabled.)
func (_class HostClass) Reboot(sessionID SessionRef, host HostRef) (_err error) {
	if (IsMock) {
		return _class.Reboot__mock(sessionID, host)
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

func (_class HostClass) Shutdown__mock(sessionID SessionRef, host HostRef) (_err error) {
	log.Println("Host.Shutdown not mocked")
	_err = errors.New("Host.Shutdown not mocked")
	return
}
// Shutdown the host. (This function can only be called if there are no currently running VMs on the host and it is disabled.)
func (_class HostClass) Shutdown(sessionID SessionRef, host HostRef) (_err error) {
	if (IsMock) {
		return _class.Shutdown__mock(sessionID, host)
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

func (_class HostClass) Enable__mock(sessionID SessionRef, host HostRef) (_err error) {
	log.Println("Host.Enable not mocked")
	_err = errors.New("Host.Enable not mocked")
	return
}
// Puts the host into a state in which new VMs can be started.
func (_class HostClass) Enable(sessionID SessionRef, host HostRef) (_err error) {
	if (IsMock) {
		return _class.Enable__mock(sessionID, host)
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

func (_class HostClass) Disable__mock(sessionID SessionRef, host HostRef) (_err error) {
	log.Println("Host.Disable not mocked")
	_err = errors.New("Host.Disable not mocked")
	return
}
// Puts the host into a state in which no new VMs can be started. Currently active VMs on the host continue to execute.
func (_class HostClass) Disable(sessionID SessionRef, host HostRef) (_err error) {
	if (IsMock) {
		return _class.Disable__mock(sessionID, host)
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

func (_class HostClass) SetDisplay__mock(sessionID SessionRef, self HostRef, value HostDisplay) (_err error) {
	log.Println("Host.SetDisplay not mocked")
	_err = errors.New("Host.SetDisplay not mocked")
	return
}
// Set the display field of the given host.
func (_class HostClass) SetDisplay(sessionID SessionRef, self HostRef, value HostDisplay) (_err error) {
	if (IsMock) {
		return _class.SetDisplay__mock(sessionID, self, value)
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

func (_class HostClass) RemoveFromGuestVCPUsParams__mock(sessionID SessionRef, self HostRef, key string) (_err error) {
	log.Println("Host.RemoveFromGuestVCPUsParams not mocked")
	_err = errors.New("Host.RemoveFromGuestVCPUsParams not mocked")
	return
}
// Remove the given key and its corresponding value from the guest_VCPUs_params field of the given host.  If the key is not in that Map, then do nothing.
func (_class HostClass) RemoveFromGuestVCPUsParams(sessionID SessionRef, self HostRef, key string) (_err error) {
	if (IsMock) {
		return _class.RemoveFromGuestVCPUsParams__mock(sessionID, self, key)
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

func (_class HostClass) AddToGuestVCPUsParams__mock(sessionID SessionRef, self HostRef, key string, value string) (_err error) {
	log.Println("Host.AddToGuestVCPUsParams not mocked")
	_err = errors.New("Host.AddToGuestVCPUsParams not mocked")
	return
}
// Add the given key-value pair to the guest_VCPUs_params field of the given host.
func (_class HostClass) AddToGuestVCPUsParams(sessionID SessionRef, self HostRef, key string, value string) (_err error) {
	if (IsMock) {
		return _class.AddToGuestVCPUsParams__mock(sessionID, self, key, value)
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

func (_class HostClass) SetGuestVCPUsParams__mock(sessionID SessionRef, self HostRef, value map[string]string) (_err error) {
	log.Println("Host.SetGuestVCPUsParams not mocked")
	_err = errors.New("Host.SetGuestVCPUsParams not mocked")
	return
}
// Set the guest_VCPUs_params field of the given host.
func (_class HostClass) SetGuestVCPUsParams(sessionID SessionRef, self HostRef, value map[string]string) (_err error) {
	if (IsMock) {
		return _class.SetGuestVCPUsParams__mock(sessionID, self, value)
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

func (_class HostClass) RemoveFromLicenseServer__mock(sessionID SessionRef, self HostRef, key string) (_err error) {
	log.Println("Host.RemoveFromLicenseServer not mocked")
	_err = errors.New("Host.RemoveFromLicenseServer not mocked")
	return
}
// Remove the given key and its corresponding value from the license_server field of the given host.  If the key is not in that Map, then do nothing.
func (_class HostClass) RemoveFromLicenseServer(sessionID SessionRef, self HostRef, key string) (_err error) {
	if (IsMock) {
		return _class.RemoveFromLicenseServer__mock(sessionID, self, key)
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

func (_class HostClass) AddToLicenseServer__mock(sessionID SessionRef, self HostRef, key string, value string) (_err error) {
	log.Println("Host.AddToLicenseServer not mocked")
	_err = errors.New("Host.AddToLicenseServer not mocked")
	return
}
// Add the given key-value pair to the license_server field of the given host.
func (_class HostClass) AddToLicenseServer(sessionID SessionRef, self HostRef, key string, value string) (_err error) {
	if (IsMock) {
		return _class.AddToLicenseServer__mock(sessionID, self, key, value)
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

func (_class HostClass) SetLicenseServer__mock(sessionID SessionRef, self HostRef, value map[string]string) (_err error) {
	log.Println("Host.SetLicenseServer not mocked")
	_err = errors.New("Host.SetLicenseServer not mocked")
	return
}
// Set the license_server field of the given host.
func (_class HostClass) SetLicenseServer(sessionID SessionRef, self HostRef, value map[string]string) (_err error) {
	if (IsMock) {
		return _class.SetLicenseServer__mock(sessionID, self, value)
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

func (_class HostClass) RemoveTags__mock(sessionID SessionRef, self HostRef, value string) (_err error) {
	log.Println("Host.RemoveTags not mocked")
	_err = errors.New("Host.RemoveTags not mocked")
	return
}
// Remove the given value from the tags field of the given host.  If the value is not in that Set, then do nothing.
func (_class HostClass) RemoveTags(sessionID SessionRef, self HostRef, value string) (_err error) {
	if (IsMock) {
		return _class.RemoveTags__mock(sessionID, self, value)
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

func (_class HostClass) AddTags__mock(sessionID SessionRef, self HostRef, value string) (_err error) {
	log.Println("Host.AddTags not mocked")
	_err = errors.New("Host.AddTags not mocked")
	return
}
// Add the given value to the tags field of the given host.  If the value is already in that Set, then do nothing.
func (_class HostClass) AddTags(sessionID SessionRef, self HostRef, value string) (_err error) {
	if (IsMock) {
		return _class.AddTags__mock(sessionID, self, value)
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

func (_class HostClass) SetTags__mock(sessionID SessionRef, self HostRef, value []string) (_err error) {
	log.Println("Host.SetTags not mocked")
	_err = errors.New("Host.SetTags not mocked")
	return
}
// Set the tags field of the given host.
func (_class HostClass) SetTags(sessionID SessionRef, self HostRef, value []string) (_err error) {
	if (IsMock) {
		return _class.SetTags__mock(sessionID, self, value)
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

func (_class HostClass) SetAddress__mock(sessionID SessionRef, self HostRef, value string) (_err error) {
	log.Println("Host.SetAddress not mocked")
	_err = errors.New("Host.SetAddress not mocked")
	return
}
// Set the address field of the given host.
func (_class HostClass) SetAddress(sessionID SessionRef, self HostRef, value string) (_err error) {
	if (IsMock) {
		return _class.SetAddress__mock(sessionID, self, value)
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

func (_class HostClass) SetHostname__mock(sessionID SessionRef, self HostRef, value string) (_err error) {
	log.Println("Host.SetHostname not mocked")
	_err = errors.New("Host.SetHostname not mocked")
	return
}
// Set the hostname field of the given host.
func (_class HostClass) SetHostname(sessionID SessionRef, self HostRef, value string) (_err error) {
	if (IsMock) {
		return _class.SetHostname__mock(sessionID, self, value)
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

func (_class HostClass) SetCrashDumpSr__mock(sessionID SessionRef, self HostRef, value SRRef) (_err error) {
	log.Println("Host.SetCrashDumpSr not mocked")
	_err = errors.New("Host.SetCrashDumpSr not mocked")
	return
}
// Set the crash_dump_sr field of the given host.
func (_class HostClass) SetCrashDumpSr(sessionID SessionRef, self HostRef, value SRRef) (_err error) {
	if (IsMock) {
		return _class.SetCrashDumpSr__mock(sessionID, self, value)
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

func (_class HostClass) SetSuspendImageSr__mock(sessionID SessionRef, self HostRef, value SRRef) (_err error) {
	log.Println("Host.SetSuspendImageSr not mocked")
	_err = errors.New("Host.SetSuspendImageSr not mocked")
	return
}
// Set the suspend_image_sr field of the given host.
func (_class HostClass) SetSuspendImageSr(sessionID SessionRef, self HostRef, value SRRef) (_err error) {
	if (IsMock) {
		return _class.SetSuspendImageSr__mock(sessionID, self, value)
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

func (_class HostClass) RemoveFromLogging__mock(sessionID SessionRef, self HostRef, key string) (_err error) {
	log.Println("Host.RemoveFromLogging not mocked")
	_err = errors.New("Host.RemoveFromLogging not mocked")
	return
}
// Remove the given key and its corresponding value from the logging field of the given host.  If the key is not in that Map, then do nothing.
func (_class HostClass) RemoveFromLogging(sessionID SessionRef, self HostRef, key string) (_err error) {
	if (IsMock) {
		return _class.RemoveFromLogging__mock(sessionID, self, key)
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

func (_class HostClass) AddToLogging__mock(sessionID SessionRef, self HostRef, key string, value string) (_err error) {
	log.Println("Host.AddToLogging not mocked")
	_err = errors.New("Host.AddToLogging not mocked")
	return
}
// Add the given key-value pair to the logging field of the given host.
func (_class HostClass) AddToLogging(sessionID SessionRef, self HostRef, key string, value string) (_err error) {
	if (IsMock) {
		return _class.AddToLogging__mock(sessionID, self, key, value)
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

func (_class HostClass) SetLogging__mock(sessionID SessionRef, self HostRef, value map[string]string) (_err error) {
	log.Println("Host.SetLogging not mocked")
	_err = errors.New("Host.SetLogging not mocked")
	return
}
// Set the logging field of the given host.
func (_class HostClass) SetLogging(sessionID SessionRef, self HostRef, value map[string]string) (_err error) {
	if (IsMock) {
		return _class.SetLogging__mock(sessionID, self, value)
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

func (_class HostClass) RemoveFromOtherConfig__mock(sessionID SessionRef, self HostRef, key string) (_err error) {
	log.Println("Host.RemoveFromOtherConfig not mocked")
	_err = errors.New("Host.RemoveFromOtherConfig not mocked")
	return
}
// Remove the given key and its corresponding value from the other_config field of the given host.  If the key is not in that Map, then do nothing.
func (_class HostClass) RemoveFromOtherConfig(sessionID SessionRef, self HostRef, key string) (_err error) {
	if (IsMock) {
		return _class.RemoveFromOtherConfig__mock(sessionID, self, key)
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

func (_class HostClass) AddToOtherConfig__mock(sessionID SessionRef, self HostRef, key string, value string) (_err error) {
	log.Println("Host.AddToOtherConfig not mocked")
	_err = errors.New("Host.AddToOtherConfig not mocked")
	return
}
// Add the given key-value pair to the other_config field of the given host.
func (_class HostClass) AddToOtherConfig(sessionID SessionRef, self HostRef, key string, value string) (_err error) {
	if (IsMock) {
		return _class.AddToOtherConfig__mock(sessionID, self, key, value)
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

func (_class HostClass) SetOtherConfig__mock(sessionID SessionRef, self HostRef, value map[string]string) (_err error) {
	log.Println("Host.SetOtherConfig not mocked")
	_err = errors.New("Host.SetOtherConfig not mocked")
	return
}
// Set the other_config field of the given host.
func (_class HostClass) SetOtherConfig(sessionID SessionRef, self HostRef, value map[string]string) (_err error) {
	if (IsMock) {
		return _class.SetOtherConfig__mock(sessionID, self, value)
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

func (_class HostClass) SetNameDescription__mock(sessionID SessionRef, self HostRef, value string) (_err error) {
	log.Println("Host.SetNameDescription not mocked")
	_err = errors.New("Host.SetNameDescription not mocked")
	return
}
// Set the name/description field of the given host.
func (_class HostClass) SetNameDescription(sessionID SessionRef, self HostRef, value string) (_err error) {
	if (IsMock) {
		return _class.SetNameDescription__mock(sessionID, self, value)
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

func (_class HostClass) SetNameLabel__mock(sessionID SessionRef, self HostRef, value string) (_err error) {
	log.Println("Host.SetNameLabel not mocked")
	_err = errors.New("Host.SetNameLabel not mocked")
	return
}
// Set the name/label field of the given host.
func (_class HostClass) SetNameLabel(sessionID SessionRef, self HostRef, value string) (_err error) {
	if (IsMock) {
		return _class.SetNameLabel__mock(sessionID, self, value)
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

func (_class HostClass) GetFeatures__mock(sessionID SessionRef, self HostRef) (_retval []FeatureRef, _err error) {
	log.Println("Host.GetFeatures not mocked")
	_err = errors.New("Host.GetFeatures not mocked")
	return
}
// Get the features field of the given host.
func (_class HostClass) GetFeatures(sessionID SessionRef, self HostRef) (_retval []FeatureRef, _err error) {
	if (IsMock) {
		return _class.GetFeatures__mock(sessionID, self)
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

func (_class HostClass) GetUpdatesRequiringReboot__mock(sessionID SessionRef, self HostRef) (_retval []PoolUpdateRef, _err error) {
	log.Println("Host.GetUpdatesRequiringReboot not mocked")
	_err = errors.New("Host.GetUpdatesRequiringReboot not mocked")
	return
}
// Get the updates_requiring_reboot field of the given host.
func (_class HostClass) GetUpdatesRequiringReboot(sessionID SessionRef, self HostRef) (_retval []PoolUpdateRef, _err error) {
	if (IsMock) {
		return _class.GetUpdatesRequiringReboot__mock(sessionID, self)
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

func (_class HostClass) GetControlDomain__mock(sessionID SessionRef, self HostRef) (_retval VMRef, _err error) {
	log.Println("Host.GetControlDomain not mocked")
	_err = errors.New("Host.GetControlDomain not mocked")
	return
}
// Get the control_domain field of the given host.
func (_class HostClass) GetControlDomain(sessionID SessionRef, self HostRef) (_retval VMRef, _err error) {
	if (IsMock) {
		return _class.GetControlDomain__mock(sessionID, self)
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

func (_class HostClass) GetVirtualHardwarePlatformVersions__mock(sessionID SessionRef, self HostRef) (_retval []int, _err error) {
	log.Println("Host.GetVirtualHardwarePlatformVersions not mocked")
	_err = errors.New("Host.GetVirtualHardwarePlatformVersions not mocked")
	return
}
// Get the virtual_hardware_platform_versions field of the given host.
func (_class HostClass) GetVirtualHardwarePlatformVersions(sessionID SessionRef, self HostRef) (_retval []int, _err error) {
	if (IsMock) {
		return _class.GetVirtualHardwarePlatformVersions__mock(sessionID, self)
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

func (_class HostClass) GetDisplay__mock(sessionID SessionRef, self HostRef) (_retval HostDisplay, _err error) {
	log.Println("Host.GetDisplay not mocked")
	_err = errors.New("Host.GetDisplay not mocked")
	return
}
// Get the display field of the given host.
func (_class HostClass) GetDisplay(sessionID SessionRef, self HostRef) (_retval HostDisplay, _err error) {
	if (IsMock) {
		return _class.GetDisplay__mock(sessionID, self)
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

func (_class HostClass) GetGuestVCPUsParams__mock(sessionID SessionRef, self HostRef) (_retval map[string]string, _err error) {
	log.Println("Host.GetGuestVCPUsParams not mocked")
	_err = errors.New("Host.GetGuestVCPUsParams not mocked")
	return
}
// Get the guest_VCPUs_params field of the given host.
func (_class HostClass) GetGuestVCPUsParams(sessionID SessionRef, self HostRef) (_retval map[string]string, _err error) {
	if (IsMock) {
		return _class.GetGuestVCPUsParams__mock(sessionID, self)
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

func (_class HostClass) GetSslLegacy__mock(sessionID SessionRef, self HostRef) (_retval bool, _err error) {
	log.Println("Host.GetSslLegacy not mocked")
	_err = errors.New("Host.GetSslLegacy not mocked")
	return
}
// Get the ssl_legacy field of the given host.
func (_class HostClass) GetSslLegacy(sessionID SessionRef, self HostRef) (_retval bool, _err error) {
	if (IsMock) {
		return _class.GetSslLegacy__mock(sessionID, self)
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

func (_class HostClass) GetPUSBs__mock(sessionID SessionRef, self HostRef) (_retval []PUSBRef, _err error) {
	log.Println("Host.GetPUSBs not mocked")
	_err = errors.New("Host.GetPUSBs not mocked")
	return
}
// Get the PUSBs field of the given host.
func (_class HostClass) GetPUSBs(sessionID SessionRef, self HostRef) (_retval []PUSBRef, _err error) {
	if (IsMock) {
		return _class.GetPUSBs__mock(sessionID, self)
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

func (_class HostClass) GetPGPUs__mock(sessionID SessionRef, self HostRef) (_retval []PGPURef, _err error) {
	log.Println("Host.GetPGPUs not mocked")
	_err = errors.New("Host.GetPGPUs not mocked")
	return
}
// Get the PGPUs field of the given host.
func (_class HostClass) GetPGPUs(sessionID SessionRef, self HostRef) (_retval []PGPURef, _err error) {
	if (IsMock) {
		return _class.GetPGPUs__mock(sessionID, self)
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

func (_class HostClass) GetPCIs__mock(sessionID SessionRef, self HostRef) (_retval []PCIRef, _err error) {
	log.Println("Host.GetPCIs not mocked")
	_err = errors.New("Host.GetPCIs not mocked")
	return
}
// Get the PCIs field of the given host.
func (_class HostClass) GetPCIs(sessionID SessionRef, self HostRef) (_retval []PCIRef, _err error) {
	if (IsMock) {
		return _class.GetPCIs__mock(sessionID, self)
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

func (_class HostClass) GetChipsetInfo__mock(sessionID SessionRef, self HostRef) (_retval map[string]string, _err error) {
	log.Println("Host.GetChipsetInfo not mocked")
	_err = errors.New("Host.GetChipsetInfo not mocked")
	return
}
// Get the chipset_info field of the given host.
func (_class HostClass) GetChipsetInfo(sessionID SessionRef, self HostRef) (_retval map[string]string, _err error) {
	if (IsMock) {
		return _class.GetChipsetInfo__mock(sessionID, self)
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

func (_class HostClass) GetLocalCacheSr__mock(sessionID SessionRef, self HostRef) (_retval SRRef, _err error) {
	log.Println("Host.GetLocalCacheSr not mocked")
	_err = errors.New("Host.GetLocalCacheSr not mocked")
	return
}
// Get the local_cache_sr field of the given host.
func (_class HostClass) GetLocalCacheSr(sessionID SessionRef, self HostRef) (_retval SRRef, _err error) {
	if (IsMock) {
		return _class.GetLocalCacheSr__mock(sessionID, self)
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

func (_class HostClass) GetPowerOnConfig__mock(sessionID SessionRef, self HostRef) (_retval map[string]string, _err error) {
	log.Println("Host.GetPowerOnConfig not mocked")
	_err = errors.New("Host.GetPowerOnConfig not mocked")
	return
}
// Get the power_on_config field of the given host.
func (_class HostClass) GetPowerOnConfig(sessionID SessionRef, self HostRef) (_retval map[string]string, _err error) {
	if (IsMock) {
		return _class.GetPowerOnConfig__mock(sessionID, self)
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

func (_class HostClass) GetPowerOnMode__mock(sessionID SessionRef, self HostRef) (_retval string, _err error) {
	log.Println("Host.GetPowerOnMode not mocked")
	_err = errors.New("Host.GetPowerOnMode not mocked")
	return
}
// Get the power_on_mode field of the given host.
func (_class HostClass) GetPowerOnMode(sessionID SessionRef, self HostRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetPowerOnMode__mock(sessionID, self)
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

func (_class HostClass) GetBiosStrings__mock(sessionID SessionRef, self HostRef) (_retval map[string]string, _err error) {
	log.Println("Host.GetBiosStrings not mocked")
	_err = errors.New("Host.GetBiosStrings not mocked")
	return
}
// Get the bios_strings field of the given host.
func (_class HostClass) GetBiosStrings(sessionID SessionRef, self HostRef) (_retval map[string]string, _err error) {
	if (IsMock) {
		return _class.GetBiosStrings__mock(sessionID, self)
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

func (_class HostClass) GetLicenseServer__mock(sessionID SessionRef, self HostRef) (_retval map[string]string, _err error) {
	log.Println("Host.GetLicenseServer not mocked")
	_err = errors.New("Host.GetLicenseServer not mocked")
	return
}
// Get the license_server field of the given host.
func (_class HostClass) GetLicenseServer(sessionID SessionRef, self HostRef) (_retval map[string]string, _err error) {
	if (IsMock) {
		return _class.GetLicenseServer__mock(sessionID, self)
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

func (_class HostClass) GetEdition__mock(sessionID SessionRef, self HostRef) (_retval string, _err error) {
	log.Println("Host.GetEdition not mocked")
	_err = errors.New("Host.GetEdition not mocked")
	return
}
// Get the edition field of the given host.
func (_class HostClass) GetEdition(sessionID SessionRef, self HostRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetEdition__mock(sessionID, self)
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

func (_class HostClass) GetExternalAuthConfiguration__mock(sessionID SessionRef, self HostRef) (_retval map[string]string, _err error) {
	log.Println("Host.GetExternalAuthConfiguration not mocked")
	_err = errors.New("Host.GetExternalAuthConfiguration not mocked")
	return
}
// Get the external_auth_configuration field of the given host.
func (_class HostClass) GetExternalAuthConfiguration(sessionID SessionRef, self HostRef) (_retval map[string]string, _err error) {
	if (IsMock) {
		return _class.GetExternalAuthConfiguration__mock(sessionID, self)
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

func (_class HostClass) GetExternalAuthServiceName__mock(sessionID SessionRef, self HostRef) (_retval string, _err error) {
	log.Println("Host.GetExternalAuthServiceName not mocked")
	_err = errors.New("Host.GetExternalAuthServiceName not mocked")
	return
}
// Get the external_auth_service_name field of the given host.
func (_class HostClass) GetExternalAuthServiceName(sessionID SessionRef, self HostRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetExternalAuthServiceName__mock(sessionID, self)
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

func (_class HostClass) GetExternalAuthType__mock(sessionID SessionRef, self HostRef) (_retval string, _err error) {
	log.Println("Host.GetExternalAuthType not mocked")
	_err = errors.New("Host.GetExternalAuthType not mocked")
	return
}
// Get the external_auth_type field of the given host.
func (_class HostClass) GetExternalAuthType(sessionID SessionRef, self HostRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetExternalAuthType__mock(sessionID, self)
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

func (_class HostClass) GetTags__mock(sessionID SessionRef, self HostRef) (_retval []string, _err error) {
	log.Println("Host.GetTags not mocked")
	_err = errors.New("Host.GetTags not mocked")
	return
}
// Get the tags field of the given host.
func (_class HostClass) GetTags(sessionID SessionRef, self HostRef) (_retval []string, _err error) {
	if (IsMock) {
		return _class.GetTags__mock(sessionID, self)
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

func (_class HostClass) GetBlobs__mock(sessionID SessionRef, self HostRef) (_retval map[string]BlobRef, _err error) {
	log.Println("Host.GetBlobs not mocked")
	_err = errors.New("Host.GetBlobs not mocked")
	return
}
// Get the blobs field of the given host.
func (_class HostClass) GetBlobs(sessionID SessionRef, self HostRef) (_retval map[string]BlobRef, _err error) {
	if (IsMock) {
		return _class.GetBlobs__mock(sessionID, self)
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

func (_class HostClass) GetHaNetworkPeers__mock(sessionID SessionRef, self HostRef) (_retval []string, _err error) {
	log.Println("Host.GetHaNetworkPeers not mocked")
	_err = errors.New("Host.GetHaNetworkPeers not mocked")
	return
}
// Get the ha_network_peers field of the given host.
func (_class HostClass) GetHaNetworkPeers(sessionID SessionRef, self HostRef) (_retval []string, _err error) {
	if (IsMock) {
		return _class.GetHaNetworkPeers__mock(sessionID, self)
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

func (_class HostClass) GetHaStatefiles__mock(sessionID SessionRef, self HostRef) (_retval []string, _err error) {
	log.Println("Host.GetHaStatefiles not mocked")
	_err = errors.New("Host.GetHaStatefiles not mocked")
	return
}
// Get the ha_statefiles field of the given host.
func (_class HostClass) GetHaStatefiles(sessionID SessionRef, self HostRef) (_retval []string, _err error) {
	if (IsMock) {
		return _class.GetHaStatefiles__mock(sessionID, self)
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

func (_class HostClass) GetLicenseParams__mock(sessionID SessionRef, self HostRef) (_retval map[string]string, _err error) {
	log.Println("Host.GetLicenseParams not mocked")
	_err = errors.New("Host.GetLicenseParams not mocked")
	return
}
// Get the license_params field of the given host.
func (_class HostClass) GetLicenseParams(sessionID SessionRef, self HostRef) (_retval map[string]string, _err error) {
	if (IsMock) {
		return _class.GetLicenseParams__mock(sessionID, self)
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

func (_class HostClass) GetMetrics__mock(sessionID SessionRef, self HostRef) (_retval HostMetricsRef, _err error) {
	log.Println("Host.GetMetrics not mocked")
	_err = errors.New("Host.GetMetrics not mocked")
	return
}
// Get the metrics field of the given host.
func (_class HostClass) GetMetrics(sessionID SessionRef, self HostRef) (_retval HostMetricsRef, _err error) {
	if (IsMock) {
		return _class.GetMetrics__mock(sessionID, self)
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

func (_class HostClass) GetAddress__mock(sessionID SessionRef, self HostRef) (_retval string, _err error) {
	log.Println("Host.GetAddress not mocked")
	_err = errors.New("Host.GetAddress not mocked")
	return
}
// Get the address field of the given host.
func (_class HostClass) GetAddress(sessionID SessionRef, self HostRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetAddress__mock(sessionID, self)
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

func (_class HostClass) GetHostname__mock(sessionID SessionRef, self HostRef) (_retval string, _err error) {
	log.Println("Host.GetHostname not mocked")
	_err = errors.New("Host.GetHostname not mocked")
	return
}
// Get the hostname field of the given host.
func (_class HostClass) GetHostname(sessionID SessionRef, self HostRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetHostname__mock(sessionID, self)
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

func (_class HostClass) GetCPUInfo__mock(sessionID SessionRef, self HostRef) (_retval map[string]string, _err error) {
	log.Println("Host.GetCPUInfo not mocked")
	_err = errors.New("Host.GetCPUInfo not mocked")
	return
}
// Get the cpu_info field of the given host.
func (_class HostClass) GetCPUInfo(sessionID SessionRef, self HostRef) (_retval map[string]string, _err error) {
	if (IsMock) {
		return _class.GetCPUInfo__mock(sessionID, self)
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

func (_class HostClass) GetHostCPUs__mock(sessionID SessionRef, self HostRef) (_retval []HostCPURef, _err error) {
	log.Println("Host.GetHostCPUs not mocked")
	_err = errors.New("Host.GetHostCPUs not mocked")
	return
}
// Get the host_CPUs field of the given host.
func (_class HostClass) GetHostCPUs(sessionID SessionRef, self HostRef) (_retval []HostCPURef, _err error) {
	if (IsMock) {
		return _class.GetHostCPUs__mock(sessionID, self)
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

func (_class HostClass) GetPBDs__mock(sessionID SessionRef, self HostRef) (_retval []PBDRef, _err error) {
	log.Println("Host.GetPBDs not mocked")
	_err = errors.New("Host.GetPBDs not mocked")
	return
}
// Get the PBDs field of the given host.
func (_class HostClass) GetPBDs(sessionID SessionRef, self HostRef) (_retval []PBDRef, _err error) {
	if (IsMock) {
		return _class.GetPBDs__mock(sessionID, self)
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

func (_class HostClass) GetUpdates__mock(sessionID SessionRef, self HostRef) (_retval []PoolUpdateRef, _err error) {
	log.Println("Host.GetUpdates not mocked")
	_err = errors.New("Host.GetUpdates not mocked")
	return
}
// Get the updates field of the given host.
func (_class HostClass) GetUpdates(sessionID SessionRef, self HostRef) (_retval []PoolUpdateRef, _err error) {
	if (IsMock) {
		return _class.GetUpdates__mock(sessionID, self)
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

func (_class HostClass) GetPatches__mock(sessionID SessionRef, self HostRef) (_retval []HostPatchRef, _err error) {
	log.Println("Host.GetPatches not mocked")
	_err = errors.New("Host.GetPatches not mocked")
	return
}
// Get the patches field of the given host.
func (_class HostClass) GetPatches(sessionID SessionRef, self HostRef) (_retval []HostPatchRef, _err error) {
	if (IsMock) {
		return _class.GetPatches__mock(sessionID, self)
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

func (_class HostClass) GetCrashdumps__mock(sessionID SessionRef, self HostRef) (_retval []HostCrashdumpRef, _err error) {
	log.Println("Host.GetCrashdumps not mocked")
	_err = errors.New("Host.GetCrashdumps not mocked")
	return
}
// Get the crashdumps field of the given host.
func (_class HostClass) GetCrashdumps(sessionID SessionRef, self HostRef) (_retval []HostCrashdumpRef, _err error) {
	if (IsMock) {
		return _class.GetCrashdumps__mock(sessionID, self)
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

func (_class HostClass) GetCrashDumpSr__mock(sessionID SessionRef, self HostRef) (_retval SRRef, _err error) {
	log.Println("Host.GetCrashDumpSr not mocked")
	_err = errors.New("Host.GetCrashDumpSr not mocked")
	return
}
// Get the crash_dump_sr field of the given host.
func (_class HostClass) GetCrashDumpSr(sessionID SessionRef, self HostRef) (_retval SRRef, _err error) {
	if (IsMock) {
		return _class.GetCrashDumpSr__mock(sessionID, self)
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

func (_class HostClass) GetSuspendImageSr__mock(sessionID SessionRef, self HostRef) (_retval SRRef, _err error) {
	log.Println("Host.GetSuspendImageSr not mocked")
	_err = errors.New("Host.GetSuspendImageSr not mocked")
	return
}
// Get the suspend_image_sr field of the given host.
func (_class HostClass) GetSuspendImageSr(sessionID SessionRef, self HostRef) (_retval SRRef, _err error) {
	if (IsMock) {
		return _class.GetSuspendImageSr__mock(sessionID, self)
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

func (_class HostClass) GetPIFs__mock(sessionID SessionRef, self HostRef) (_retval []PIFRef, _err error) {
	log.Println("Host.GetPIFs not mocked")
	_err = errors.New("Host.GetPIFs not mocked")
	return
}
// Get the PIFs field of the given host.
func (_class HostClass) GetPIFs(sessionID SessionRef, self HostRef) (_retval []PIFRef, _err error) {
	if (IsMock) {
		return _class.GetPIFs__mock(sessionID, self)
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

func (_class HostClass) GetLogging__mock(sessionID SessionRef, self HostRef) (_retval map[string]string, _err error) {
	log.Println("Host.GetLogging not mocked")
	_err = errors.New("Host.GetLogging not mocked")
	return
}
// Get the logging field of the given host.
func (_class HostClass) GetLogging(sessionID SessionRef, self HostRef) (_retval map[string]string, _err error) {
	if (IsMock) {
		return _class.GetLogging__mock(sessionID, self)
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

func (_class HostClass) GetResidentVMs__mock(sessionID SessionRef, self HostRef) (_retval []VMRef, _err error) {
	log.Println("Host.GetResidentVMs not mocked")
	_err = errors.New("Host.GetResidentVMs not mocked")
	return
}
// Get the resident_VMs field of the given host.
func (_class HostClass) GetResidentVMs(sessionID SessionRef, self HostRef) (_retval []VMRef, _err error) {
	if (IsMock) {
		return _class.GetResidentVMs__mock(sessionID, self)
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

func (_class HostClass) GetSupportedBootloaders__mock(sessionID SessionRef, self HostRef) (_retval []string, _err error) {
	log.Println("Host.GetSupportedBootloaders not mocked")
	_err = errors.New("Host.GetSupportedBootloaders not mocked")
	return
}
// Get the supported_bootloaders field of the given host.
func (_class HostClass) GetSupportedBootloaders(sessionID SessionRef, self HostRef) (_retval []string, _err error) {
	if (IsMock) {
		return _class.GetSupportedBootloaders__mock(sessionID, self)
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

func (_class HostClass) GetSchedPolicy__mock(sessionID SessionRef, self HostRef) (_retval string, _err error) {
	log.Println("Host.GetSchedPolicy not mocked")
	_err = errors.New("Host.GetSchedPolicy not mocked")
	return
}
// Get the sched_policy field of the given host.
func (_class HostClass) GetSchedPolicy(sessionID SessionRef, self HostRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetSchedPolicy__mock(sessionID, self)
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

func (_class HostClass) GetCPUConfiguration__mock(sessionID SessionRef, self HostRef) (_retval map[string]string, _err error) {
	log.Println("Host.GetCPUConfiguration not mocked")
	_err = errors.New("Host.GetCPUConfiguration not mocked")
	return
}
// Get the cpu_configuration field of the given host.
func (_class HostClass) GetCPUConfiguration(sessionID SessionRef, self HostRef) (_retval map[string]string, _err error) {
	if (IsMock) {
		return _class.GetCPUConfiguration__mock(sessionID, self)
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

func (_class HostClass) GetCapabilities__mock(sessionID SessionRef, self HostRef) (_retval []string, _err error) {
	log.Println("Host.GetCapabilities not mocked")
	_err = errors.New("Host.GetCapabilities not mocked")
	return
}
// Get the capabilities field of the given host.
func (_class HostClass) GetCapabilities(sessionID SessionRef, self HostRef) (_retval []string, _err error) {
	if (IsMock) {
		return _class.GetCapabilities__mock(sessionID, self)
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

func (_class HostClass) GetOtherConfig__mock(sessionID SessionRef, self HostRef) (_retval map[string]string, _err error) {
	log.Println("Host.GetOtherConfig not mocked")
	_err = errors.New("Host.GetOtherConfig not mocked")
	return
}
// Get the other_config field of the given host.
func (_class HostClass) GetOtherConfig(sessionID SessionRef, self HostRef) (_retval map[string]string, _err error) {
	if (IsMock) {
		return _class.GetOtherConfig__mock(sessionID, self)
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

func (_class HostClass) GetSoftwareVersion__mock(sessionID SessionRef, self HostRef) (_retval map[string]string, _err error) {
	log.Println("Host.GetSoftwareVersion not mocked")
	_err = errors.New("Host.GetSoftwareVersion not mocked")
	return
}
// Get the software_version field of the given host.
func (_class HostClass) GetSoftwareVersion(sessionID SessionRef, self HostRef) (_retval map[string]string, _err error) {
	if (IsMock) {
		return _class.GetSoftwareVersion__mock(sessionID, self)
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

func (_class HostClass) GetEnabled__mock(sessionID SessionRef, self HostRef) (_retval bool, _err error) {
	log.Println("Host.GetEnabled not mocked")
	_err = errors.New("Host.GetEnabled not mocked")
	return
}
// Get the enabled field of the given host.
func (_class HostClass) GetEnabled(sessionID SessionRef, self HostRef) (_retval bool, _err error) {
	if (IsMock) {
		return _class.GetEnabled__mock(sessionID, self)
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

func (_class HostClass) GetAPIVersionVendorImplementation__mock(sessionID SessionRef, self HostRef) (_retval map[string]string, _err error) {
	log.Println("Host.GetAPIVersionVendorImplementation not mocked")
	_err = errors.New("Host.GetAPIVersionVendorImplementation not mocked")
	return
}
// Get the API_version/vendor_implementation field of the given host.
func (_class HostClass) GetAPIVersionVendorImplementation(sessionID SessionRef, self HostRef) (_retval map[string]string, _err error) {
	if (IsMock) {
		return _class.GetAPIVersionVendorImplementation__mock(sessionID, self)
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

func (_class HostClass) GetAPIVersionVendor__mock(sessionID SessionRef, self HostRef) (_retval string, _err error) {
	log.Println("Host.GetAPIVersionVendor not mocked")
	_err = errors.New("Host.GetAPIVersionVendor not mocked")
	return
}
// Get the API_version/vendor field of the given host.
func (_class HostClass) GetAPIVersionVendor(sessionID SessionRef, self HostRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetAPIVersionVendor__mock(sessionID, self)
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

func (_class HostClass) GetAPIVersionMinor__mock(sessionID SessionRef, self HostRef) (_retval int, _err error) {
	log.Println("Host.GetAPIVersionMinor not mocked")
	_err = errors.New("Host.GetAPIVersionMinor not mocked")
	return
}
// Get the API_version/minor field of the given host.
func (_class HostClass) GetAPIVersionMinor(sessionID SessionRef, self HostRef) (_retval int, _err error) {
	if (IsMock) {
		return _class.GetAPIVersionMinor__mock(sessionID, self)
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

func (_class HostClass) GetAPIVersionMajor__mock(sessionID SessionRef, self HostRef) (_retval int, _err error) {
	log.Println("Host.GetAPIVersionMajor not mocked")
	_err = errors.New("Host.GetAPIVersionMajor not mocked")
	return
}
// Get the API_version/major field of the given host.
func (_class HostClass) GetAPIVersionMajor(sessionID SessionRef, self HostRef) (_retval int, _err error) {
	if (IsMock) {
		return _class.GetAPIVersionMajor__mock(sessionID, self)
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

func (_class HostClass) GetCurrentOperations__mock(sessionID SessionRef, self HostRef) (_retval map[string]HostAllowedOperations, _err error) {
	log.Println("Host.GetCurrentOperations not mocked")
	_err = errors.New("Host.GetCurrentOperations not mocked")
	return
}
// Get the current_operations field of the given host.
func (_class HostClass) GetCurrentOperations(sessionID SessionRef, self HostRef) (_retval map[string]HostAllowedOperations, _err error) {
	if (IsMock) {
		return _class.GetCurrentOperations__mock(sessionID, self)
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

func (_class HostClass) GetAllowedOperations__mock(sessionID SessionRef, self HostRef) (_retval []HostAllowedOperations, _err error) {
	log.Println("Host.GetAllowedOperations not mocked")
	_err = errors.New("Host.GetAllowedOperations not mocked")
	return
}
// Get the allowed_operations field of the given host.
func (_class HostClass) GetAllowedOperations(sessionID SessionRef, self HostRef) (_retval []HostAllowedOperations, _err error) {
	if (IsMock) {
		return _class.GetAllowedOperations__mock(sessionID, self)
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

func (_class HostClass) GetMemoryOverhead__mock(sessionID SessionRef, self HostRef) (_retval int, _err error) {
	log.Println("Host.GetMemoryOverhead not mocked")
	_err = errors.New("Host.GetMemoryOverhead not mocked")
	return
}
// Get the memory/overhead field of the given host.
func (_class HostClass) GetMemoryOverhead(sessionID SessionRef, self HostRef) (_retval int, _err error) {
	if (IsMock) {
		return _class.GetMemoryOverhead__mock(sessionID, self)
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

func (_class HostClass) GetNameDescription__mock(sessionID SessionRef, self HostRef) (_retval string, _err error) {
	log.Println("Host.GetNameDescription not mocked")
	_err = errors.New("Host.GetNameDescription not mocked")
	return
}
// Get the name/description field of the given host.
func (_class HostClass) GetNameDescription(sessionID SessionRef, self HostRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetNameDescription__mock(sessionID, self)
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

func (_class HostClass) GetNameLabel__mock(sessionID SessionRef, self HostRef) (_retval string, _err error) {
	log.Println("Host.GetNameLabel not mocked")
	_err = errors.New("Host.GetNameLabel not mocked")
	return
}
// Get the name/label field of the given host.
func (_class HostClass) GetNameLabel(sessionID SessionRef, self HostRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetNameLabel__mock(sessionID, self)
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

func (_class HostClass) GetUUID__mock(sessionID SessionRef, self HostRef) (_retval string, _err error) {
	log.Println("Host.GetUUID not mocked")
	_err = errors.New("Host.GetUUID not mocked")
	return
}
// Get the uuid field of the given host.
func (_class HostClass) GetUUID(sessionID SessionRef, self HostRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetUUID__mock(sessionID, self)
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

func (_class HostClass) GetByNameLabel__mock(sessionID SessionRef, label string) (_retval []HostRef, _err error) {
	log.Println("Host.GetByNameLabel not mocked")
	_err = errors.New("Host.GetByNameLabel not mocked")
	return
}
// Get all the host instances with the given label.
func (_class HostClass) GetByNameLabel(sessionID SessionRef, label string) (_retval []HostRef, _err error) {
	if (IsMock) {
		return _class.GetByNameLabel__mock(sessionID, label)
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

func (_class HostClass) GetByUUID__mock(sessionID SessionRef, uuid string) (_retval HostRef, _err error) {
	log.Println("Host.GetByUUID not mocked")
	_err = errors.New("Host.GetByUUID not mocked")
	return
}
// Get a reference to the host instance with the specified UUID.
func (_class HostClass) GetByUUID(sessionID SessionRef, uuid string) (_retval HostRef, _err error) {
	if (IsMock) {
		return _class.GetByUUID__mock(sessionID, uuid)
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

func (_class HostClass) GetRecord__mock(sessionID SessionRef, self HostRef) (_retval HostRecord, _err error) {
	log.Println("Host.GetRecord not mocked")
	_err = errors.New("Host.GetRecord not mocked")
	return
}
// Get a record containing the current state of the given host.
func (_class HostClass) GetRecord(sessionID SessionRef, self HostRef) (_retval HostRecord, _err error) {
	if (IsMock) {
		return _class.GetRecord__mock(sessionID, self)
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
