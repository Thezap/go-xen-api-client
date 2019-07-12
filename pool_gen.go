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

type PoolAllowedOperations string

const (
  // Indicates this pool is in the process of enabling HA
	PoolAllowedOperationsHaEnable PoolAllowedOperations = "ha_enable"
  // Indicates this pool is in the process of disabling HA
	PoolAllowedOperationsHaDisable PoolAllowedOperations = "ha_disable"
)

type PoolRecord struct {
  // Unique identifier/object reference
	UUID string
  // Short name
	NameLabel string
  // Description
	NameDescription string
  // The host that is pool master
	Master HostRef
  // Default SR for VDIs
	DefaultSR SRRef
  // The SR in which VDIs for suspend images are created
	SuspendImageSR SRRef
  // The SR in which VDIs for crash dumps are created
	CrashDumpSR SRRef
  // additional configuration
	OtherConfig map[string]string
  // true if HA is enabled on the pool, false otherwise
	HaEnabled bool
  // The current HA configuration
	HaConfiguration map[string]string
  // HA statefile VDIs in use
	HaStatefiles []string
  // Number of host failures to tolerate before the Pool is declared to be overcommitted
	HaHostFailuresToTolerate int
  // Number of future host failures we have managed to find a plan for. Once this reaches zero any future host failures will cause the failure of protected VMs.
	HaPlanExistsFor int
  // If set to false then operations which would cause the Pool to become overcommitted will be blocked.
	HaAllowOvercommit bool
  // True if the Pool is considered to be overcommitted i.e. if there exist insufficient physical resources to tolerate the configured number of host failures
	HaOvercommitted bool
  // Binary blobs associated with this pool
	Blobs map[string]BlobRef
  // user-specified tags for categorization purposes
	Tags []string
  // gui-specific configuration for pool
	GuiConfig map[string]string
  // Configuration for the automatic health check feature
	HealthCheckConfig map[string]string
  // Url for the configured workload balancing host
	WlbURL string
  // Username for accessing the workload balancing host
	WlbUsername string
  // true if workload balancing is enabled on the pool, false otherwise
	WlbEnabled bool
  // true if communication with the WLB server should enforce SSL certificate verification.
	WlbVerifyCert bool
  // true a redo-log is to be used other than when HA is enabled, false otherwise
	RedoLogEnabled bool
  // indicates the VDI to use for the redo-log other than when HA is enabled
	RedoLogVdi VDIRef
  // address of the vswitch controller
	VswitchController string
  // Pool-wide restrictions currently in effect
	Restrictions map[string]string
  // The set of currently known metadata VDIs for this pool
	MetadataVDIs []VDIRef
  // The HA cluster stack that is currently in use. Only valid when HA is enabled.
	HaClusterStack string
  // list of the operations allowed in this state. This list is advisory only and the server state may have changed by the time this field is read by a client.
	AllowedOperations []PoolAllowedOperations
  // links each of the running tasks using this object (by reference) to a current_operation enum which describes the nature of the task.
	CurrentOperations map[string]PoolAllowedOperations
  // Pool-wide guest agent configuration information
	GuestAgentConfig map[string]string
  // Details about the physical CPUs on the pool
	CPUInfo map[string]string
  // The pool-wide policy for clients on whether to use the vendor device or not on newly created VMs. This field will also be consulted if the 'has_vendor_device' field is not specified in the VM.create call.
	PolicyNoVendorDevice bool
  // The pool-wide flag to show if the live patching feauture is disabled or not.
	LivePatchingDisabled bool
  // true if IGMP snooping is enabled in the pool, false otherwise.
	IgmpSnoopingEnabled bool
}

type PoolRef string

// Pool-wide information
type PoolClass struct {
	client *Client
}

func (_class PoolClass) GetAllRecords__mock(sessionID SessionRef) (_retval map[PoolRef]PoolRecord, _err error) {
	log.Println("Pool.GetAllRecords not mocked")
	_err = errors.New("Pool.GetAllRecords not mocked")
	return
}
// Return a map of pool references to pool records for all pools known to the system.
func (_class PoolClass) GetAllRecords(sessionID SessionRef) (_retval map[PoolRef]PoolRecord, _err error) {
	if (IsMock) {
		return _class.GetAllRecords__mock(sessionID)
	}	
	_method := "pool.get_all_records"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertPoolRefToPoolRecordMapToGo(_method + " -> ", _result.Value)
	return
}

func (_class PoolClass) GetAll__mock(sessionID SessionRef) (_retval []PoolRef, _err error) {
	log.Println("Pool.GetAll not mocked")
	_err = errors.New("Pool.GetAll not mocked")
	return
}
// Return a list of all the pools known to the system.
func (_class PoolClass) GetAll(sessionID SessionRef) (_retval []PoolRef, _err error) {
	if (IsMock) {
		return _class.GetAll__mock(sessionID)
	}	
	_method := "pool.get_all"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertPoolRefSetToGo(_method + " -> ", _result.Value)
	return
}

func (_class PoolClass) RemoveFromGuestAgentConfig__mock(sessionID SessionRef, self PoolRef, key string) (_err error) {
	log.Println("Pool.RemoveFromGuestAgentConfig not mocked")
	_err = errors.New("Pool.RemoveFromGuestAgentConfig not mocked")
	return
}
// Remove a key-value pair from the pool-wide guest agent configuration
func (_class PoolClass) RemoveFromGuestAgentConfig(sessionID SessionRef, self PoolRef, key string) (_err error) {
	if (IsMock) {
		return _class.RemoveFromGuestAgentConfig__mock(sessionID, self, key)
	}	
	_method := "pool.remove_from_guest_agent_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PoolClass) AddToGuestAgentConfig__mock(sessionID SessionRef, self PoolRef, key string, value string) (_err error) {
	log.Println("Pool.AddToGuestAgentConfig not mocked")
	_err = errors.New("Pool.AddToGuestAgentConfig not mocked")
	return
}
// Add a key-value pair to the pool-wide guest agent configuration
func (_class PoolClass) AddToGuestAgentConfig(sessionID SessionRef, self PoolRef, key string, value string) (_err error) {
	if (IsMock) {
		return _class.AddToGuestAgentConfig__mock(sessionID, self, key, value)
	}	
	_method := "pool.add_to_guest_agent_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PoolClass) HasExtension__mock(sessionID SessionRef, self PoolRef, name string) (_retval bool, _err error) {
	log.Println("Pool.HasExtension not mocked")
	_err = errors.New("Pool.HasExtension not mocked")
	return
}
// Return true if the extension is available on the pool
func (_class PoolClass) HasExtension(sessionID SessionRef, self PoolRef, name string) (_retval bool, _err error) {
	if (IsMock) {
		return _class.HasExtension__mock(sessionID, self, name)
	}	
	_method := "pool.has_extension"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_nameArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "name"), name)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg, _nameArg)
	if _err != nil {
		return
	}
	_retval, _err = convertBoolToGo(_method + " -> ", _result.Value)
	return
}

func (_class PoolClass) SetIgmpSnoopingEnabled__mock(sessionID SessionRef, self PoolRef, value bool) (_err error) {
	log.Println("Pool.SetIgmpSnoopingEnabled not mocked")
	_err = errors.New("Pool.SetIgmpSnoopingEnabled not mocked")
	return
}
// Enable or disable IGMP Snooping on the pool.
func (_class PoolClass) SetIgmpSnoopingEnabled(sessionID SessionRef, self PoolRef, value bool) (_err error) {
	if (IsMock) {
		return _class.SetIgmpSnoopingEnabled__mock(sessionID, self, value)
	}	
	_method := "pool.set_igmp_snooping_enabled"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PoolClass) DisableSslLegacy__mock(sessionID SessionRef, self PoolRef) (_err error) {
	log.Println("Pool.DisableSslLegacy not mocked")
	_err = errors.New("Pool.DisableSslLegacy not mocked")
	return
}
// Sets ssl_legacy true on each host, pool-master last. See Host.ssl_legacy and Host.set_ssl_legacy.
func (_class PoolClass) DisableSslLegacy(sessionID SessionRef, self PoolRef) (_err error) {
	if (IsMock) {
		return _class.DisableSslLegacy__mock(sessionID, self)
	}	
	_method := "pool.disable_ssl_legacy"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg)
	return
}

func (_class PoolClass) EnableSslLegacy__mock(sessionID SessionRef, self PoolRef) (_err error) {
	log.Println("Pool.EnableSslLegacy not mocked")
	_err = errors.New("Pool.EnableSslLegacy not mocked")
	return
}
// Sets ssl_legacy true on each host, pool-master last. See Host.ssl_legacy and Host.set_ssl_legacy.
func (_class PoolClass) EnableSslLegacy(sessionID SessionRef, self PoolRef) (_err error) {
	if (IsMock) {
		return _class.EnableSslLegacy__mock(sessionID, self)
	}	
	_method := "pool.enable_ssl_legacy"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg)
	return
}

func (_class PoolClass) ApplyEdition__mock(sessionID SessionRef, self PoolRef, edition string) (_err error) {
	log.Println("Pool.ApplyEdition not mocked")
	_err = errors.New("Pool.ApplyEdition not mocked")
	return
}
// Apply an edition to all hosts in the pool
func (_class PoolClass) ApplyEdition(sessionID SessionRef, self PoolRef, edition string) (_err error) {
	if (IsMock) {
		return _class.ApplyEdition__mock(sessionID, self, edition)
	}	
	_method := "pool.apply_edition"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_editionArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "edition"), edition)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _editionArg)
	return
}

func (_class PoolClass) GetLicenseState__mock(sessionID SessionRef, self PoolRef) (_retval map[string]string, _err error) {
	log.Println("Pool.GetLicenseState not mocked")
	_err = errors.New("Pool.GetLicenseState not mocked")
	return
}
// This call returns the license state for the pool
func (_class PoolClass) GetLicenseState(sessionID SessionRef, self PoolRef) (_retval map[string]string, _err error) {
	if (IsMock) {
		return _class.GetLicenseState__mock(sessionID, self)
	}	
	_method := "pool.get_license_state"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PoolClass) DisableLocalStorageCaching__mock(sessionID SessionRef, self PoolRef) (_err error) {
	log.Println("Pool.DisableLocalStorageCaching not mocked")
	_err = errors.New("Pool.DisableLocalStorageCaching not mocked")
	return
}
// This call disables pool-wide local storage caching
func (_class PoolClass) DisableLocalStorageCaching(sessionID SessionRef, self PoolRef) (_err error) {
	if (IsMock) {
		return _class.DisableLocalStorageCaching__mock(sessionID, self)
	}	
	_method := "pool.disable_local_storage_caching"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg)
	return
}

func (_class PoolClass) EnableLocalStorageCaching__mock(sessionID SessionRef, self PoolRef) (_err error) {
	log.Println("Pool.EnableLocalStorageCaching not mocked")
	_err = errors.New("Pool.EnableLocalStorageCaching not mocked")
	return
}
// This call attempts to enable pool-wide local storage caching
func (_class PoolClass) EnableLocalStorageCaching(sessionID SessionRef, self PoolRef) (_err error) {
	if (IsMock) {
		return _class.EnableLocalStorageCaching__mock(sessionID, self)
	}	
	_method := "pool.enable_local_storage_caching"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg)
	return
}

func (_class PoolClass) TestArchiveTarget__mock(sessionID SessionRef, self PoolRef, config map[string]string) (_retval string, _err error) {
	log.Println("Pool.TestArchiveTarget not mocked")
	_err = errors.New("Pool.TestArchiveTarget not mocked")
	return
}
// This call tests if a location is valid
func (_class PoolClass) TestArchiveTarget(sessionID SessionRef, self PoolRef, config map[string]string) (_retval string, _err error) {
	if (IsMock) {
		return _class.TestArchiveTarget__mock(sessionID, self, config)
	}	
	_method := "pool.test_archive_target"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_configArg, _err := convertStringToStringMapToXen(fmt.Sprintf("%s(%s)", _method, "config"), config)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg, _configArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringToGo(_method + " -> ", _result.Value)
	return
}

func (_class PoolClass) SetVswitchController__mock(sessionID SessionRef, address string) (_err error) {
	log.Println("Pool.SetVswitchController not mocked")
	_err = errors.New("Pool.SetVswitchController not mocked")
	return
}
// Set the IP address of the vswitch controller.
func (_class PoolClass) SetVswitchController(sessionID SessionRef, address string) (_err error) {
	if (IsMock) {
		return _class.SetVswitchController__mock(sessionID, address)
	}	
	_method := "pool.set_vswitch_controller"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_addressArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "address"), address)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _addressArg)
	return
}

func (_class PoolClass) DisableRedoLog__mock(sessionID SessionRef) (_err error) {
	log.Println("Pool.DisableRedoLog not mocked")
	_err = errors.New("Pool.DisableRedoLog not mocked")
	return
}
// Disable the redo log if in use, unless HA is enabled.
func (_class PoolClass) DisableRedoLog(sessionID SessionRef) (_err error) {
	if (IsMock) {
		return _class.DisableRedoLog__mock(sessionID)
	}	
	_method := "pool.disable_redo_log"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg)
	return
}

func (_class PoolClass) EnableRedoLog__mock(sessionID SessionRef, sr SRRef) (_err error) {
	log.Println("Pool.EnableRedoLog not mocked")
	_err = errors.New("Pool.EnableRedoLog not mocked")
	return
}
// Enable the redo log on the given SR and start using it, unless HA is enabled.
func (_class PoolClass) EnableRedoLog(sessionID SessionRef, sr SRRef) (_err error) {
	if (IsMock) {
		return _class.EnableRedoLog__mock(sessionID, sr)
	}	
	_method := "pool.enable_redo_log"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_srArg, _err := convertSRRefToXen(fmt.Sprintf("%s(%s)", _method, "sr"), sr)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _srArg)
	return
}

func (_class PoolClass) CertificateSync__mock(sessionID SessionRef) (_err error) {
	log.Println("Pool.CertificateSync not mocked")
	_err = errors.New("Pool.CertificateSync not mocked")
	return
}
// Sync SSL certificates from master to slaves.
func (_class PoolClass) CertificateSync(sessionID SessionRef) (_err error) {
	if (IsMock) {
		return _class.CertificateSync__mock(sessionID)
	}	
	_method := "pool.certificate_sync"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg)
	return
}

func (_class PoolClass) CrlList__mock(sessionID SessionRef) (_retval []string, _err error) {
	log.Println("Pool.CrlList not mocked")
	_err = errors.New("Pool.CrlList not mocked")
	return
}
// List all installed SSL certificate revocation lists.
func (_class PoolClass) CrlList(sessionID SessionRef) (_retval []string, _err error) {
	if (IsMock) {
		return _class.CrlList__mock(sessionID)
	}	
	_method := "pool.crl_list"
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

func (_class PoolClass) CrlUninstall__mock(sessionID SessionRef, name string) (_err error) {
	log.Println("Pool.CrlUninstall not mocked")
	_err = errors.New("Pool.CrlUninstall not mocked")
	return
}
// Remove an SSL certificate revocation list.
func (_class PoolClass) CrlUninstall(sessionID SessionRef, name string) (_err error) {
	if (IsMock) {
		return _class.CrlUninstall__mock(sessionID, name)
	}	
	_method := "pool.crl_uninstall"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_nameArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "name"), name)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _nameArg)
	return
}

func (_class PoolClass) CrlInstall__mock(sessionID SessionRef, name string, cert string) (_err error) {
	log.Println("Pool.CrlInstall not mocked")
	_err = errors.New("Pool.CrlInstall not mocked")
	return
}
// Install an SSL certificate revocation list, pool-wide.
func (_class PoolClass) CrlInstall(sessionID SessionRef, name string, cert string) (_err error) {
	if (IsMock) {
		return _class.CrlInstall__mock(sessionID, name, cert)
	}	
	_method := "pool.crl_install"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_nameArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "name"), name)
	if _err != nil {
		return
	}
	_certArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "cert"), cert)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _nameArg, _certArg)
	return
}

func (_class PoolClass) CertificateList__mock(sessionID SessionRef) (_retval []string, _err error) {
	log.Println("Pool.CertificateList not mocked")
	_err = errors.New("Pool.CertificateList not mocked")
	return
}
// List all installed SSL certificates.
func (_class PoolClass) CertificateList(sessionID SessionRef) (_retval []string, _err error) {
	if (IsMock) {
		return _class.CertificateList__mock(sessionID)
	}	
	_method := "pool.certificate_list"
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

func (_class PoolClass) CertificateUninstall__mock(sessionID SessionRef, name string) (_err error) {
	log.Println("Pool.CertificateUninstall not mocked")
	_err = errors.New("Pool.CertificateUninstall not mocked")
	return
}
// Remove an SSL certificate.
func (_class PoolClass) CertificateUninstall(sessionID SessionRef, name string) (_err error) {
	if (IsMock) {
		return _class.CertificateUninstall__mock(sessionID, name)
	}	
	_method := "pool.certificate_uninstall"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_nameArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "name"), name)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _nameArg)
	return
}

func (_class PoolClass) CertificateInstall__mock(sessionID SessionRef, name string, cert string) (_err error) {
	log.Println("Pool.CertificateInstall not mocked")
	_err = errors.New("Pool.CertificateInstall not mocked")
	return
}
// Install an SSL certificate pool-wide.
func (_class PoolClass) CertificateInstall(sessionID SessionRef, name string, cert string) (_err error) {
	if (IsMock) {
		return _class.CertificateInstall__mock(sessionID, name, cert)
	}	
	_method := "pool.certificate_install"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_nameArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "name"), name)
	if _err != nil {
		return
	}
	_certArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "cert"), cert)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _nameArg, _certArg)
	return
}

func (_class PoolClass) SendTestPost__mock(sessionID SessionRef, host string, port int, body string) (_retval string, _err error) {
	log.Println("Pool.SendTestPost not mocked")
	_err = errors.New("Pool.SendTestPost not mocked")
	return
}
// Send the given body to the given host and port, using HTTPS, and print the response.  This is used for debugging the SSL layer.
func (_class PoolClass) SendTestPost(sessionID SessionRef, host string, port int, body string) (_retval string, _err error) {
	if (IsMock) {
		return _class.SendTestPost__mock(sessionID, host, port, body)
	}	
	_method := "pool.send_test_post"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_hostArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "host"), host)
	if _err != nil {
		return
	}
	_portArg, _err := convertIntToXen(fmt.Sprintf("%s(%s)", _method, "port"), port)
	if _err != nil {
		return
	}
	_bodyArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "body"), body)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _hostArg, _portArg, _bodyArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringToGo(_method + " -> ", _result.Value)
	return
}

func (_class PoolClass) RetrieveWlbRecommendations__mock(sessionID SessionRef) (_retval map[VMRef][]string, _err error) {
	log.Println("Pool.RetrieveWlbRecommendations not mocked")
	_err = errors.New("Pool.RetrieveWlbRecommendations not mocked")
	return
}
// Retrieves vm migrate recommendations for the pool from the workload balancing server
func (_class PoolClass) RetrieveWlbRecommendations(sessionID SessionRef) (_retval map[VMRef][]string, _err error) {
	if (IsMock) {
		return _class.RetrieveWlbRecommendations__mock(sessionID)
	}	
	_method := "pool.retrieve_wlb_recommendations"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVMRefToStringSetMapToGo(_method + " -> ", _result.Value)
	return
}

func (_class PoolClass) RetrieveWlbConfiguration__mock(sessionID SessionRef) (_retval map[string]string, _err error) {
	log.Println("Pool.RetrieveWlbConfiguration not mocked")
	_err = errors.New("Pool.RetrieveWlbConfiguration not mocked")
	return
}
// Retrieves the pool optimization criteria from the workload balancing server
func (_class PoolClass) RetrieveWlbConfiguration(sessionID SessionRef) (_retval map[string]string, _err error) {
	if (IsMock) {
		return _class.RetrieveWlbConfiguration__mock(sessionID)
	}	
	_method := "pool.retrieve_wlb_configuration"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringToStringMapToGo(_method + " -> ", _result.Value)
	return
}

func (_class PoolClass) SendWlbConfiguration__mock(sessionID SessionRef, config map[string]string) (_err error) {
	log.Println("Pool.SendWlbConfiguration not mocked")
	_err = errors.New("Pool.SendWlbConfiguration not mocked")
	return
}
// Sets the pool optimization criteria for the workload balancing server
func (_class PoolClass) SendWlbConfiguration(sessionID SessionRef, config map[string]string) (_err error) {
	if (IsMock) {
		return _class.SendWlbConfiguration__mock(sessionID, config)
	}	
	_method := "pool.send_wlb_configuration"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_configArg, _err := convertStringToStringMapToXen(fmt.Sprintf("%s(%s)", _method, "config"), config)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _configArg)
	return
}

func (_class PoolClass) DeconfigureWlb__mock(sessionID SessionRef) (_err error) {
	log.Println("Pool.DeconfigureWlb not mocked")
	_err = errors.New("Pool.DeconfigureWlb not mocked")
	return
}
// Permanently deconfigures workload balancing monitoring on this pool
func (_class PoolClass) DeconfigureWlb(sessionID SessionRef) (_err error) {
	if (IsMock) {
		return _class.DeconfigureWlb__mock(sessionID)
	}	
	_method := "pool.deconfigure_wlb"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg)
	return
}

func (_class PoolClass) InitializeWlb__mock(sessionID SessionRef, wlbURL string, wlbUsername string, wlbPassword string, xenserverUsername string, xenserverPassword string) (_err error) {
	log.Println("Pool.InitializeWlb not mocked")
	_err = errors.New("Pool.InitializeWlb not mocked")
	return
}
// Initializes workload balancing monitoring on this pool with the specified wlb server
func (_class PoolClass) InitializeWlb(sessionID SessionRef, wlbURL string, wlbUsername string, wlbPassword string, xenserverUsername string, xenserverPassword string) (_err error) {
	if (IsMock) {
		return _class.InitializeWlb__mock(sessionID, wlbURL, wlbUsername, wlbPassword, xenserverUsername, xenserverPassword)
	}	
	_method := "pool.initialize_wlb"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_wlbURLArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "wlb_url"), wlbURL)
	if _err != nil {
		return
	}
	_wlbUsernameArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "wlb_username"), wlbUsername)
	if _err != nil {
		return
	}
	_wlbPasswordArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "wlb_password"), wlbPassword)
	if _err != nil {
		return
	}
	_xenserverUsernameArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "xenserver_username"), xenserverUsername)
	if _err != nil {
		return
	}
	_xenserverPasswordArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "xenserver_password"), xenserverPassword)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _wlbURLArg, _wlbUsernameArg, _wlbPasswordArg, _xenserverUsernameArg, _xenserverPasswordArg)
	return
}

func (_class PoolClass) DetectNonhomogeneousExternalAuth__mock(sessionID SessionRef, pool PoolRef) (_err error) {
	log.Println("Pool.DetectNonhomogeneousExternalAuth not mocked")
	_err = errors.New("Pool.DetectNonhomogeneousExternalAuth not mocked")
	return
}
// This call asynchronously detects if the external authentication configuration in any slave is different from that in the master and raises appropriate alerts
func (_class PoolClass) DetectNonhomogeneousExternalAuth(sessionID SessionRef, pool PoolRef) (_err error) {
	if (IsMock) {
		return _class.DetectNonhomogeneousExternalAuth__mock(sessionID, pool)
	}	
	_method := "pool.detect_nonhomogeneous_external_auth"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_poolArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "pool"), pool)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _poolArg)
	return
}

func (_class PoolClass) DisableExternalAuth__mock(sessionID SessionRef, pool PoolRef, config map[string]string) (_err error) {
	log.Println("Pool.DisableExternalAuth not mocked")
	_err = errors.New("Pool.DisableExternalAuth not mocked")
	return
}
// This call disables external authentication on all the hosts of the pool
func (_class PoolClass) DisableExternalAuth(sessionID SessionRef, pool PoolRef, config map[string]string) (_err error) {
	if (IsMock) {
		return _class.DisableExternalAuth__mock(sessionID, pool, config)
	}	
	_method := "pool.disable_external_auth"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_poolArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "pool"), pool)
	if _err != nil {
		return
	}
	_configArg, _err := convertStringToStringMapToXen(fmt.Sprintf("%s(%s)", _method, "config"), config)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _poolArg, _configArg)
	return
}

func (_class PoolClass) EnableExternalAuth__mock(sessionID SessionRef, pool PoolRef, config map[string]string, serviceName string, authType string) (_err error) {
	log.Println("Pool.EnableExternalAuth not mocked")
	_err = errors.New("Pool.EnableExternalAuth not mocked")
	return
}
// This call enables external authentication on all the hosts of the pool
func (_class PoolClass) EnableExternalAuth(sessionID SessionRef, pool PoolRef, config map[string]string, serviceName string, authType string) (_err error) {
	if (IsMock) {
		return _class.EnableExternalAuth__mock(sessionID, pool, config, serviceName, authType)
	}	
	_method := "pool.enable_external_auth"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_poolArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "pool"), pool)
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
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _poolArg, _configArg, _serviceNameArg, _authTypeArg)
	return
}

func (_class PoolClass) CreateNewBlob__mock(sessionID SessionRef, pool PoolRef, name string, mimeType string, public bool) (_retval BlobRef, _err error) {
	log.Println("Pool.CreateNewBlob not mocked")
	_err = errors.New("Pool.CreateNewBlob not mocked")
	return
}
// Create a placeholder for a named binary blob of data that is associated with this pool
func (_class PoolClass) CreateNewBlob(sessionID SessionRef, pool PoolRef, name string, mimeType string, public bool) (_retval BlobRef, _err error) {
	if (IsMock) {
		return _class.CreateNewBlob__mock(sessionID, pool, name, mimeType, public)
	}	
	_method := "pool.create_new_blob"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_poolArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "pool"), pool)
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
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _poolArg, _nameArg, _mimeTypeArg, _publicArg)
	if _err != nil {
		return
	}
	_retval, _err = convertBlobRefToGo(_method + " -> ", _result.Value)
	return
}

func (_class PoolClass) SetHaHostFailuresToTolerate__mock(sessionID SessionRef, self PoolRef, value int) (_err error) {
	log.Println("Pool.SetHaHostFailuresToTolerate not mocked")
	_err = errors.New("Pool.SetHaHostFailuresToTolerate not mocked")
	return
}
// Set the maximum number of host failures to consider in the HA VM restart planner
func (_class PoolClass) SetHaHostFailuresToTolerate(sessionID SessionRef, self PoolRef, value int) (_err error) {
	if (IsMock) {
		return _class.SetHaHostFailuresToTolerate__mock(sessionID, self, value)
	}	
	_method := "pool.set_ha_host_failures_to_tolerate"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PoolClass) HaComputeVMFailoverPlan__mock(sessionID SessionRef, failedHosts []HostRef, failedVms []VMRef) (_retval map[VMRef]map[string]string, _err error) {
	log.Println("Pool.HaComputeVMFailoverPlan not mocked")
	_err = errors.New("Pool.HaComputeVMFailoverPlan not mocked")
	return
}
// Return a VM failover plan assuming a given subset of hosts fail
func (_class PoolClass) HaComputeVMFailoverPlan(sessionID SessionRef, failedHosts []HostRef, failedVms []VMRef) (_retval map[VMRef]map[string]string, _err error) {
	if (IsMock) {
		return _class.HaComputeVMFailoverPlan__mock(sessionID, failedHosts, failedVms)
	}	
	_method := "pool.ha_compute_vm_failover_plan"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_failedHostsArg, _err := convertHostRefSetToXen(fmt.Sprintf("%s(%s)", _method, "failed_hosts"), failedHosts)
	if _err != nil {
		return
	}
	_failedVmsArg, _err := convertVMRefSetToXen(fmt.Sprintf("%s(%s)", _method, "failed_vms"), failedVms)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _failedHostsArg, _failedVmsArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVMRefToStringToStringMapMapToGo(_method + " -> ", _result.Value)
	return
}

func (_class PoolClass) HaComputeHypotheticalMaxHostFailuresToTolerate__mock(sessionID SessionRef, configuration map[VMRef]string) (_retval int, _err error) {
	log.Println("Pool.HaComputeHypotheticalMaxHostFailuresToTolerate not mocked")
	_err = errors.New("Pool.HaComputeHypotheticalMaxHostFailuresToTolerate not mocked")
	return
}
// Returns the maximum number of host failures we could tolerate before we would be unable to restart the provided VMs
func (_class PoolClass) HaComputeHypotheticalMaxHostFailuresToTolerate(sessionID SessionRef, configuration map[VMRef]string) (_retval int, _err error) {
	if (IsMock) {
		return _class.HaComputeHypotheticalMaxHostFailuresToTolerate__mock(sessionID, configuration)
	}	
	_method := "pool.ha_compute_hypothetical_max_host_failures_to_tolerate"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_configurationArg, _err := convertVMRefToStringMapToXen(fmt.Sprintf("%s(%s)", _method, "configuration"), configuration)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _configurationArg)
	if _err != nil {
		return
	}
	_retval, _err = convertIntToGo(_method + " -> ", _result.Value)
	return
}

func (_class PoolClass) HaComputeMaxHostFailuresToTolerate__mock(sessionID SessionRef) (_retval int, _err error) {
	log.Println("Pool.HaComputeMaxHostFailuresToTolerate not mocked")
	_err = errors.New("Pool.HaComputeMaxHostFailuresToTolerate not mocked")
	return
}
// Returns the maximum number of host failures we could tolerate before we would be unable to restart configured VMs
func (_class PoolClass) HaComputeMaxHostFailuresToTolerate(sessionID SessionRef) (_retval int, _err error) {
	if (IsMock) {
		return _class.HaComputeMaxHostFailuresToTolerate__mock(sessionID)
	}	
	_method := "pool.ha_compute_max_host_failures_to_tolerate"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertIntToGo(_method + " -> ", _result.Value)
	return
}

func (_class PoolClass) HaFailoverPlanExists__mock(sessionID SessionRef, n int) (_retval bool, _err error) {
	log.Println("Pool.HaFailoverPlanExists not mocked")
	_err = errors.New("Pool.HaFailoverPlanExists not mocked")
	return
}
// Returns true if a VM failover plan exists for up to 'n' host failures
func (_class PoolClass) HaFailoverPlanExists(sessionID SessionRef, n int) (_retval bool, _err error) {
	if (IsMock) {
		return _class.HaFailoverPlanExists__mock(sessionID, n)
	}	
	_method := "pool.ha_failover_plan_exists"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_nArg, _err := convertIntToXen(fmt.Sprintf("%s(%s)", _method, "n"), n)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _nArg)
	if _err != nil {
		return
	}
	_retval, _err = convertBoolToGo(_method + " -> ", _result.Value)
	return
}

func (_class PoolClass) HaPreventRestartsFor__mock(sessionID SessionRef, seconds int) (_err error) {
	log.Println("Pool.HaPreventRestartsFor not mocked")
	_err = errors.New("Pool.HaPreventRestartsFor not mocked")
	return
}
// When this call returns the VM restart logic will not run for the requested number of seconds. If the argument is zero then the restart thread is immediately unblocked
func (_class PoolClass) HaPreventRestartsFor(sessionID SessionRef, seconds int) (_err error) {
	if (IsMock) {
		return _class.HaPreventRestartsFor__mock(sessionID, seconds)
	}	
	_method := "pool.ha_prevent_restarts_for"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_secondsArg, _err := convertIntToXen(fmt.Sprintf("%s(%s)", _method, "seconds"), seconds)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _secondsArg)
	return
}

func (_class PoolClass) DesignateNewMaster__mock(sessionID SessionRef, host HostRef) (_err error) {
	log.Println("Pool.DesignateNewMaster not mocked")
	_err = errors.New("Pool.DesignateNewMaster not mocked")
	return
}
// Perform an orderly handover of the role of master to the referenced host.
func (_class PoolClass) DesignateNewMaster(sessionID SessionRef, host HostRef) (_err error) {
	if (IsMock) {
		return _class.DesignateNewMaster__mock(sessionID, host)
	}	
	_method := "pool.designate_new_master"
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

func (_class PoolClass) SyncDatabase__mock(sessionID SessionRef) (_err error) {
	log.Println("Pool.SyncDatabase not mocked")
	_err = errors.New("Pool.SyncDatabase not mocked")
	return
}
// Forcibly synchronise the database now
func (_class PoolClass) SyncDatabase(sessionID SessionRef) (_err error) {
	if (IsMock) {
		return _class.SyncDatabase__mock(sessionID)
	}	
	_method := "pool.sync_database"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg)
	return
}

func (_class PoolClass) DisableHa__mock(sessionID SessionRef) (_err error) {
	log.Println("Pool.DisableHa not mocked")
	_err = errors.New("Pool.DisableHa not mocked")
	return
}
// Turn off High Availability mode
func (_class PoolClass) DisableHa(sessionID SessionRef) (_err error) {
	if (IsMock) {
		return _class.DisableHa__mock(sessionID)
	}	
	_method := "pool.disable_ha"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg)
	return
}

func (_class PoolClass) EnableHa__mock(sessionID SessionRef, heartbeatSrs []SRRef, configuration map[string]string) (_err error) {
	log.Println("Pool.EnableHa not mocked")
	_err = errors.New("Pool.EnableHa not mocked")
	return
}
// Turn on High Availability mode
func (_class PoolClass) EnableHa(sessionID SessionRef, heartbeatSrs []SRRef, configuration map[string]string) (_err error) {
	if (IsMock) {
		return _class.EnableHa__mock(sessionID, heartbeatSrs, configuration)
	}	
	_method := "pool.enable_ha"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_heartbeatSrsArg, _err := convertSRRefSetToXen(fmt.Sprintf("%s(%s)", _method, "heartbeat_srs"), heartbeatSrs)
	if _err != nil {
		return
	}
	_configurationArg, _err := convertStringToStringMapToXen(fmt.Sprintf("%s(%s)", _method, "configuration"), configuration)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _heartbeatSrsArg, _configurationArg)
	return
}

func (_class PoolClass) CreateVLANFromPIF__mock(sessionID SessionRef, pif PIFRef, network NetworkRef, vlan int) (_retval []PIFRef, _err error) {
	log.Println("Pool.CreateVLANFromPIF not mocked")
	_err = errors.New("Pool.CreateVLANFromPIF not mocked")
	return
}
// Create a pool-wide VLAN by taking the PIF.
//
// Errors:
//  VLAN_TAG_INVALID - You tried to create a VLAN, but the tag you gave was invalid -- it must be between 0 and 4094.  The parameter echoes the VLAN tag you gave.
func (_class PoolClass) CreateVLANFromPIF(sessionID SessionRef, pif PIFRef, network NetworkRef, vlan int) (_retval []PIFRef, _err error) {
	if (IsMock) {
		return _class.CreateVLANFromPIF__mock(sessionID, pif, network, vlan)
	}	
	_method := "pool.create_VLAN_from_PIF"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_pifArg, _err := convertPIFRefToXen(fmt.Sprintf("%s(%s)", _method, "pif"), pif)
	if _err != nil {
		return
	}
	_networkArg, _err := convertNetworkRefToXen(fmt.Sprintf("%s(%s)", _method, "network"), network)
	if _err != nil {
		return
	}
	_vlanArg, _err := convertIntToXen(fmt.Sprintf("%s(%s)", _method, "VLAN"), vlan)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _pifArg, _networkArg, _vlanArg)
	if _err != nil {
		return
	}
	_retval, _err = convertPIFRefSetToGo(_method + " -> ", _result.Value)
	return
}

func (_class PoolClass) ManagementReconfigure__mock(sessionID SessionRef, network NetworkRef) (_err error) {
	log.Println("Pool.ManagementReconfigure not mocked")
	_err = errors.New("Pool.ManagementReconfigure not mocked")
	return
}
// Reconfigure the management network interface for all Hosts in the Pool
//
// Errors:
//  HA_IS_ENABLED - The operation could not be performed because HA is enabled on the Pool
//  PIF_NOT_PRESENT - This host has no PIF on the given network.
//  CANNOT_PLUG_BOND_SLAVE - This PIF is a bond slave and cannot be plugged.
//  PIF_INCOMPATIBLE_PRIMARY_ADDRESS_TYPE - The primary address types are not compatible
//  PIF_HAS_NO_NETWORK_CONFIGURATION - PIF has no IP configuration (mode currently set to 'none')
//  PIF_HAS_NO_V6_NETWORK_CONFIGURATION - PIF has no IPv6 configuration (mode currently set to 'none')
func (_class PoolClass) ManagementReconfigure(sessionID SessionRef, network NetworkRef) (_err error) {
	if (IsMock) {
		return _class.ManagementReconfigure__mock(sessionID, network)
	}	
	_method := "pool.management_reconfigure"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_networkArg, _err := convertNetworkRefToXen(fmt.Sprintf("%s(%s)", _method, "network"), network)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _networkArg)
	return
}

func (_class PoolClass) CreateVLAN__mock(sessionID SessionRef, device string, network NetworkRef, vlan int) (_retval []PIFRef, _err error) {
	log.Println("Pool.CreateVLAN not mocked")
	_err = errors.New("Pool.CreateVLAN not mocked")
	return
}
// Create PIFs, mapping a network to the same physical interface/VLAN on each host. This call is deprecated: use Pool.create_VLAN_from_PIF instead.
//
// Errors:
//  VLAN_TAG_INVALID - You tried to create a VLAN, but the tag you gave was invalid -- it must be between 0 and 4094.  The parameter echoes the VLAN tag you gave.
func (_class PoolClass) CreateVLAN(sessionID SessionRef, device string, network NetworkRef, vlan int) (_retval []PIFRef, _err error) {
	if (IsMock) {
		return _class.CreateVLAN__mock(sessionID, device, network, vlan)
	}	
	_method := "pool.create_VLAN"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_deviceArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "device"), device)
	if _err != nil {
		return
	}
	_networkArg, _err := convertNetworkRefToXen(fmt.Sprintf("%s(%s)", _method, "network"), network)
	if _err != nil {
		return
	}
	_vlanArg, _err := convertIntToXen(fmt.Sprintf("%s(%s)", _method, "VLAN"), vlan)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _deviceArg, _networkArg, _vlanArg)
	if _err != nil {
		return
	}
	_retval, _err = convertPIFRefSetToGo(_method + " -> ", _result.Value)
	return
}

func (_class PoolClass) RecoverSlaves__mock(sessionID SessionRef) (_retval []HostRef, _err error) {
	log.Println("Pool.RecoverSlaves not mocked")
	_err = errors.New("Pool.RecoverSlaves not mocked")
	return
}
// Instruct a pool master, M, to try and contact its slaves and, if slaves are in emergency mode, reset their master address to M.
func (_class PoolClass) RecoverSlaves(sessionID SessionRef) (_retval []HostRef, _err error) {
	if (IsMock) {
		return _class.RecoverSlaves__mock(sessionID)
	}	
	_method := "pool.recover_slaves"
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

func (_class PoolClass) EmergencyResetMaster__mock(sessionID SessionRef, masterAddress string) (_err error) {
	log.Println("Pool.EmergencyResetMaster not mocked")
	_err = errors.New("Pool.EmergencyResetMaster not mocked")
	return
}
// Instruct a slave already in a pool that the master has changed
func (_class PoolClass) EmergencyResetMaster(sessionID SessionRef, masterAddress string) (_err error) {
	if (IsMock) {
		return _class.EmergencyResetMaster__mock(sessionID, masterAddress)
	}	
	_method := "pool.emergency_reset_master"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_masterAddressArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "master_address"), masterAddress)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _masterAddressArg)
	return
}

func (_class PoolClass) EmergencyTransitionToMaster__mock(sessionID SessionRef) (_err error) {
	log.Println("Pool.EmergencyTransitionToMaster not mocked")
	_err = errors.New("Pool.EmergencyTransitionToMaster not mocked")
	return
}
// Instruct host that's currently a slave to transition to being master
func (_class PoolClass) EmergencyTransitionToMaster(sessionID SessionRef) (_err error) {
	if (IsMock) {
		return _class.EmergencyTransitionToMaster__mock(sessionID)
	}	
	_method := "pool.emergency_transition_to_master"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg)
	return
}

func (_class PoolClass) Eject__mock(sessionID SessionRef, host HostRef) (_err error) {
	log.Println("Pool.Eject not mocked")
	_err = errors.New("Pool.Eject not mocked")
	return
}
// Instruct a pool master to eject a host from the pool
func (_class PoolClass) Eject(sessionID SessionRef, host HostRef) (_err error) {
	if (IsMock) {
		return _class.Eject__mock(sessionID, host)
	}	
	_method := "pool.eject"
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

func (_class PoolClass) JoinForce__mock(sessionID SessionRef, masterAddress string, masterUsername string, masterPassword string) (_err error) {
	log.Println("Pool.JoinForce not mocked")
	_err = errors.New("Pool.JoinForce not mocked")
	return
}
// Instruct host to join a new pool
func (_class PoolClass) JoinForce(sessionID SessionRef, masterAddress string, masterUsername string, masterPassword string) (_err error) {
	if (IsMock) {
		return _class.JoinForce__mock(sessionID, masterAddress, masterUsername, masterPassword)
	}	
	_method := "pool.join_force"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_masterAddressArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "master_address"), masterAddress)
	if _err != nil {
		return
	}
	_masterUsernameArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "master_username"), masterUsername)
	if _err != nil {
		return
	}
	_masterPasswordArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "master_password"), masterPassword)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _masterAddressArg, _masterUsernameArg, _masterPasswordArg)
	return
}

func (_class PoolClass) Join__mock(sessionID SessionRef, masterAddress string, masterUsername string, masterPassword string) (_err error) {
	log.Println("Pool.Join not mocked")
	_err = errors.New("Pool.Join not mocked")
	return
}
// Instruct host to join a new pool
//
// Errors:
//  JOINING_HOST_CANNOT_CONTAIN_SHARED_SRS - The host joining the pool cannot contain any shared storage.
func (_class PoolClass) Join(sessionID SessionRef, masterAddress string, masterUsername string, masterPassword string) (_err error) {
	if (IsMock) {
		return _class.Join__mock(sessionID, masterAddress, masterUsername, masterPassword)
	}	
	_method := "pool.join"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_masterAddressArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "master_address"), masterAddress)
	if _err != nil {
		return
	}
	_masterUsernameArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "master_username"), masterUsername)
	if _err != nil {
		return
	}
	_masterPasswordArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "master_password"), masterPassword)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _masterAddressArg, _masterUsernameArg, _masterPasswordArg)
	return
}

func (_class PoolClass) SetLivePatchingDisabled__mock(sessionID SessionRef, self PoolRef, value bool) (_err error) {
	log.Println("Pool.SetLivePatchingDisabled not mocked")
	_err = errors.New("Pool.SetLivePatchingDisabled not mocked")
	return
}
// Set the live_patching_disabled field of the given pool.
func (_class PoolClass) SetLivePatchingDisabled(sessionID SessionRef, self PoolRef, value bool) (_err error) {
	if (IsMock) {
		return _class.SetLivePatchingDisabled__mock(sessionID, self, value)
	}	
	_method := "pool.set_live_patching_disabled"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PoolClass) SetPolicyNoVendorDevice__mock(sessionID SessionRef, self PoolRef, value bool) (_err error) {
	log.Println("Pool.SetPolicyNoVendorDevice not mocked")
	_err = errors.New("Pool.SetPolicyNoVendorDevice not mocked")
	return
}
// Set the policy_no_vendor_device field of the given pool.
func (_class PoolClass) SetPolicyNoVendorDevice(sessionID SessionRef, self PoolRef, value bool) (_err error) {
	if (IsMock) {
		return _class.SetPolicyNoVendorDevice__mock(sessionID, self, value)
	}	
	_method := "pool.set_policy_no_vendor_device"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PoolClass) SetWlbVerifyCert__mock(sessionID SessionRef, self PoolRef, value bool) (_err error) {
	log.Println("Pool.SetWlbVerifyCert not mocked")
	_err = errors.New("Pool.SetWlbVerifyCert not mocked")
	return
}
// Set the wlb_verify_cert field of the given pool.
func (_class PoolClass) SetWlbVerifyCert(sessionID SessionRef, self PoolRef, value bool) (_err error) {
	if (IsMock) {
		return _class.SetWlbVerifyCert__mock(sessionID, self, value)
	}	
	_method := "pool.set_wlb_verify_cert"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PoolClass) SetWlbEnabled__mock(sessionID SessionRef, self PoolRef, value bool) (_err error) {
	log.Println("Pool.SetWlbEnabled not mocked")
	_err = errors.New("Pool.SetWlbEnabled not mocked")
	return
}
// Set the wlb_enabled field of the given pool.
func (_class PoolClass) SetWlbEnabled(sessionID SessionRef, self PoolRef, value bool) (_err error) {
	if (IsMock) {
		return _class.SetWlbEnabled__mock(sessionID, self, value)
	}	
	_method := "pool.set_wlb_enabled"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PoolClass) RemoveFromHealthCheckConfig__mock(sessionID SessionRef, self PoolRef, key string) (_err error) {
	log.Println("Pool.RemoveFromHealthCheckConfig not mocked")
	_err = errors.New("Pool.RemoveFromHealthCheckConfig not mocked")
	return
}
// Remove the given key and its corresponding value from the health_check_config field of the given pool.  If the key is not in that Map, then do nothing.
func (_class PoolClass) RemoveFromHealthCheckConfig(sessionID SessionRef, self PoolRef, key string) (_err error) {
	if (IsMock) {
		return _class.RemoveFromHealthCheckConfig__mock(sessionID, self, key)
	}	
	_method := "pool.remove_from_health_check_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PoolClass) AddToHealthCheckConfig__mock(sessionID SessionRef, self PoolRef, key string, value string) (_err error) {
	log.Println("Pool.AddToHealthCheckConfig not mocked")
	_err = errors.New("Pool.AddToHealthCheckConfig not mocked")
	return
}
// Add the given key-value pair to the health_check_config field of the given pool.
func (_class PoolClass) AddToHealthCheckConfig(sessionID SessionRef, self PoolRef, key string, value string) (_err error) {
	if (IsMock) {
		return _class.AddToHealthCheckConfig__mock(sessionID, self, key, value)
	}	
	_method := "pool.add_to_health_check_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PoolClass) SetHealthCheckConfig__mock(sessionID SessionRef, self PoolRef, value map[string]string) (_err error) {
	log.Println("Pool.SetHealthCheckConfig not mocked")
	_err = errors.New("Pool.SetHealthCheckConfig not mocked")
	return
}
// Set the health_check_config field of the given pool.
func (_class PoolClass) SetHealthCheckConfig(sessionID SessionRef, self PoolRef, value map[string]string) (_err error) {
	if (IsMock) {
		return _class.SetHealthCheckConfig__mock(sessionID, self, value)
	}	
	_method := "pool.set_health_check_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PoolClass) RemoveFromGuiConfig__mock(sessionID SessionRef, self PoolRef, key string) (_err error) {
	log.Println("Pool.RemoveFromGuiConfig not mocked")
	_err = errors.New("Pool.RemoveFromGuiConfig not mocked")
	return
}
// Remove the given key and its corresponding value from the gui_config field of the given pool.  If the key is not in that Map, then do nothing.
func (_class PoolClass) RemoveFromGuiConfig(sessionID SessionRef, self PoolRef, key string) (_err error) {
	if (IsMock) {
		return _class.RemoveFromGuiConfig__mock(sessionID, self, key)
	}	
	_method := "pool.remove_from_gui_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PoolClass) AddToGuiConfig__mock(sessionID SessionRef, self PoolRef, key string, value string) (_err error) {
	log.Println("Pool.AddToGuiConfig not mocked")
	_err = errors.New("Pool.AddToGuiConfig not mocked")
	return
}
// Add the given key-value pair to the gui_config field of the given pool.
func (_class PoolClass) AddToGuiConfig(sessionID SessionRef, self PoolRef, key string, value string) (_err error) {
	if (IsMock) {
		return _class.AddToGuiConfig__mock(sessionID, self, key, value)
	}	
	_method := "pool.add_to_gui_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PoolClass) SetGuiConfig__mock(sessionID SessionRef, self PoolRef, value map[string]string) (_err error) {
	log.Println("Pool.SetGuiConfig not mocked")
	_err = errors.New("Pool.SetGuiConfig not mocked")
	return
}
// Set the gui_config field of the given pool.
func (_class PoolClass) SetGuiConfig(sessionID SessionRef, self PoolRef, value map[string]string) (_err error) {
	if (IsMock) {
		return _class.SetGuiConfig__mock(sessionID, self, value)
	}	
	_method := "pool.set_gui_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PoolClass) RemoveTags__mock(sessionID SessionRef, self PoolRef, value string) (_err error) {
	log.Println("Pool.RemoveTags not mocked")
	_err = errors.New("Pool.RemoveTags not mocked")
	return
}
// Remove the given value from the tags field of the given pool.  If the value is not in that Set, then do nothing.
func (_class PoolClass) RemoveTags(sessionID SessionRef, self PoolRef, value string) (_err error) {
	if (IsMock) {
		return _class.RemoveTags__mock(sessionID, self, value)
	}	
	_method := "pool.remove_tags"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PoolClass) AddTags__mock(sessionID SessionRef, self PoolRef, value string) (_err error) {
	log.Println("Pool.AddTags not mocked")
	_err = errors.New("Pool.AddTags not mocked")
	return
}
// Add the given value to the tags field of the given pool.  If the value is already in that Set, then do nothing.
func (_class PoolClass) AddTags(sessionID SessionRef, self PoolRef, value string) (_err error) {
	if (IsMock) {
		return _class.AddTags__mock(sessionID, self, value)
	}	
	_method := "pool.add_tags"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PoolClass) SetTags__mock(sessionID SessionRef, self PoolRef, value []string) (_err error) {
	log.Println("Pool.SetTags not mocked")
	_err = errors.New("Pool.SetTags not mocked")
	return
}
// Set the tags field of the given pool.
func (_class PoolClass) SetTags(sessionID SessionRef, self PoolRef, value []string) (_err error) {
	if (IsMock) {
		return _class.SetTags__mock(sessionID, self, value)
	}	
	_method := "pool.set_tags"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PoolClass) SetHaAllowOvercommit__mock(sessionID SessionRef, self PoolRef, value bool) (_err error) {
	log.Println("Pool.SetHaAllowOvercommit not mocked")
	_err = errors.New("Pool.SetHaAllowOvercommit not mocked")
	return
}
// Set the ha_allow_overcommit field of the given pool.
func (_class PoolClass) SetHaAllowOvercommit(sessionID SessionRef, self PoolRef, value bool) (_err error) {
	if (IsMock) {
		return _class.SetHaAllowOvercommit__mock(sessionID, self, value)
	}	
	_method := "pool.set_ha_allow_overcommit"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PoolClass) RemoveFromOtherConfig__mock(sessionID SessionRef, self PoolRef, key string) (_err error) {
	log.Println("Pool.RemoveFromOtherConfig not mocked")
	_err = errors.New("Pool.RemoveFromOtherConfig not mocked")
	return
}
// Remove the given key and its corresponding value from the other_config field of the given pool.  If the key is not in that Map, then do nothing.
func (_class PoolClass) RemoveFromOtherConfig(sessionID SessionRef, self PoolRef, key string) (_err error) {
	if (IsMock) {
		return _class.RemoveFromOtherConfig__mock(sessionID, self, key)
	}	
	_method := "pool.remove_from_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PoolClass) AddToOtherConfig__mock(sessionID SessionRef, self PoolRef, key string, value string) (_err error) {
	log.Println("Pool.AddToOtherConfig not mocked")
	_err = errors.New("Pool.AddToOtherConfig not mocked")
	return
}
// Add the given key-value pair to the other_config field of the given pool.
func (_class PoolClass) AddToOtherConfig(sessionID SessionRef, self PoolRef, key string, value string) (_err error) {
	if (IsMock) {
		return _class.AddToOtherConfig__mock(sessionID, self, key, value)
	}	
	_method := "pool.add_to_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PoolClass) SetOtherConfig__mock(sessionID SessionRef, self PoolRef, value map[string]string) (_err error) {
	log.Println("Pool.SetOtherConfig not mocked")
	_err = errors.New("Pool.SetOtherConfig not mocked")
	return
}
// Set the other_config field of the given pool.
func (_class PoolClass) SetOtherConfig(sessionID SessionRef, self PoolRef, value map[string]string) (_err error) {
	if (IsMock) {
		return _class.SetOtherConfig__mock(sessionID, self, value)
	}	
	_method := "pool.set_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PoolClass) SetCrashDumpSR__mock(sessionID SessionRef, self PoolRef, value SRRef) (_err error) {
	log.Println("Pool.SetCrashDumpSR not mocked")
	_err = errors.New("Pool.SetCrashDumpSR not mocked")
	return
}
// Set the crash_dump_SR field of the given pool.
func (_class PoolClass) SetCrashDumpSR(sessionID SessionRef, self PoolRef, value SRRef) (_err error) {
	if (IsMock) {
		return _class.SetCrashDumpSR__mock(sessionID, self, value)
	}	
	_method := "pool.set_crash_dump_SR"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PoolClass) SetSuspendImageSR__mock(sessionID SessionRef, self PoolRef, value SRRef) (_err error) {
	log.Println("Pool.SetSuspendImageSR not mocked")
	_err = errors.New("Pool.SetSuspendImageSR not mocked")
	return
}
// Set the suspend_image_SR field of the given pool.
func (_class PoolClass) SetSuspendImageSR(sessionID SessionRef, self PoolRef, value SRRef) (_err error) {
	if (IsMock) {
		return _class.SetSuspendImageSR__mock(sessionID, self, value)
	}	
	_method := "pool.set_suspend_image_SR"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PoolClass) SetDefaultSR__mock(sessionID SessionRef, self PoolRef, value SRRef) (_err error) {
	log.Println("Pool.SetDefaultSR not mocked")
	_err = errors.New("Pool.SetDefaultSR not mocked")
	return
}
// Set the default_SR field of the given pool.
func (_class PoolClass) SetDefaultSR(sessionID SessionRef, self PoolRef, value SRRef) (_err error) {
	if (IsMock) {
		return _class.SetDefaultSR__mock(sessionID, self, value)
	}	
	_method := "pool.set_default_SR"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PoolClass) SetNameDescription__mock(sessionID SessionRef, self PoolRef, value string) (_err error) {
	log.Println("Pool.SetNameDescription not mocked")
	_err = errors.New("Pool.SetNameDescription not mocked")
	return
}
// Set the name_description field of the given pool.
func (_class PoolClass) SetNameDescription(sessionID SessionRef, self PoolRef, value string) (_err error) {
	if (IsMock) {
		return _class.SetNameDescription__mock(sessionID, self, value)
	}	
	_method := "pool.set_name_description"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PoolClass) SetNameLabel__mock(sessionID SessionRef, self PoolRef, value string) (_err error) {
	log.Println("Pool.SetNameLabel not mocked")
	_err = errors.New("Pool.SetNameLabel not mocked")
	return
}
// Set the name_label field of the given pool.
func (_class PoolClass) SetNameLabel(sessionID SessionRef, self PoolRef, value string) (_err error) {
	if (IsMock) {
		return _class.SetNameLabel__mock(sessionID, self, value)
	}	
	_method := "pool.set_name_label"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PoolClass) GetIgmpSnoopingEnabled__mock(sessionID SessionRef, self PoolRef) (_retval bool, _err error) {
	log.Println("Pool.GetIgmpSnoopingEnabled not mocked")
	_err = errors.New("Pool.GetIgmpSnoopingEnabled not mocked")
	return
}
// Get the igmp_snooping_enabled field of the given pool.
func (_class PoolClass) GetIgmpSnoopingEnabled(sessionID SessionRef, self PoolRef) (_retval bool, _err error) {
	if (IsMock) {
		return _class.GetIgmpSnoopingEnabled__mock(sessionID, self)
	}	
	_method := "pool.get_igmp_snooping_enabled"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PoolClass) GetLivePatchingDisabled__mock(sessionID SessionRef, self PoolRef) (_retval bool, _err error) {
	log.Println("Pool.GetLivePatchingDisabled not mocked")
	_err = errors.New("Pool.GetLivePatchingDisabled not mocked")
	return
}
// Get the live_patching_disabled field of the given pool.
func (_class PoolClass) GetLivePatchingDisabled(sessionID SessionRef, self PoolRef) (_retval bool, _err error) {
	if (IsMock) {
		return _class.GetLivePatchingDisabled__mock(sessionID, self)
	}	
	_method := "pool.get_live_patching_disabled"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PoolClass) GetPolicyNoVendorDevice__mock(sessionID SessionRef, self PoolRef) (_retval bool, _err error) {
	log.Println("Pool.GetPolicyNoVendorDevice not mocked")
	_err = errors.New("Pool.GetPolicyNoVendorDevice not mocked")
	return
}
// Get the policy_no_vendor_device field of the given pool.
func (_class PoolClass) GetPolicyNoVendorDevice(sessionID SessionRef, self PoolRef) (_retval bool, _err error) {
	if (IsMock) {
		return _class.GetPolicyNoVendorDevice__mock(sessionID, self)
	}	
	_method := "pool.get_policy_no_vendor_device"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PoolClass) GetCPUInfo__mock(sessionID SessionRef, self PoolRef) (_retval map[string]string, _err error) {
	log.Println("Pool.GetCPUInfo not mocked")
	_err = errors.New("Pool.GetCPUInfo not mocked")
	return
}
// Get the cpu_info field of the given pool.
func (_class PoolClass) GetCPUInfo(sessionID SessionRef, self PoolRef) (_retval map[string]string, _err error) {
	if (IsMock) {
		return _class.GetCPUInfo__mock(sessionID, self)
	}	
	_method := "pool.get_cpu_info"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PoolClass) GetGuestAgentConfig__mock(sessionID SessionRef, self PoolRef) (_retval map[string]string, _err error) {
	log.Println("Pool.GetGuestAgentConfig not mocked")
	_err = errors.New("Pool.GetGuestAgentConfig not mocked")
	return
}
// Get the guest_agent_config field of the given pool.
func (_class PoolClass) GetGuestAgentConfig(sessionID SessionRef, self PoolRef) (_retval map[string]string, _err error) {
	if (IsMock) {
		return _class.GetGuestAgentConfig__mock(sessionID, self)
	}	
	_method := "pool.get_guest_agent_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PoolClass) GetCurrentOperations__mock(sessionID SessionRef, self PoolRef) (_retval map[string]PoolAllowedOperations, _err error) {
	log.Println("Pool.GetCurrentOperations not mocked")
	_err = errors.New("Pool.GetCurrentOperations not mocked")
	return
}
// Get the current_operations field of the given pool.
func (_class PoolClass) GetCurrentOperations(sessionID SessionRef, self PoolRef) (_retval map[string]PoolAllowedOperations, _err error) {
	if (IsMock) {
		return _class.GetCurrentOperations__mock(sessionID, self)
	}	
	_method := "pool.get_current_operations"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringToEnumPoolAllowedOperationsMapToGo(_method + " -> ", _result.Value)
	return
}

func (_class PoolClass) GetAllowedOperations__mock(sessionID SessionRef, self PoolRef) (_retval []PoolAllowedOperations, _err error) {
	log.Println("Pool.GetAllowedOperations not mocked")
	_err = errors.New("Pool.GetAllowedOperations not mocked")
	return
}
// Get the allowed_operations field of the given pool.
func (_class PoolClass) GetAllowedOperations(sessionID SessionRef, self PoolRef) (_retval []PoolAllowedOperations, _err error) {
	if (IsMock) {
		return _class.GetAllowedOperations__mock(sessionID, self)
	}	
	_method := "pool.get_allowed_operations"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertEnumPoolAllowedOperationsSetToGo(_method + " -> ", _result.Value)
	return
}

func (_class PoolClass) GetHaClusterStack__mock(sessionID SessionRef, self PoolRef) (_retval string, _err error) {
	log.Println("Pool.GetHaClusterStack not mocked")
	_err = errors.New("Pool.GetHaClusterStack not mocked")
	return
}
// Get the ha_cluster_stack field of the given pool.
func (_class PoolClass) GetHaClusterStack(sessionID SessionRef, self PoolRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetHaClusterStack__mock(sessionID, self)
	}	
	_method := "pool.get_ha_cluster_stack"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PoolClass) GetMetadataVDIs__mock(sessionID SessionRef, self PoolRef) (_retval []VDIRef, _err error) {
	log.Println("Pool.GetMetadataVDIs not mocked")
	_err = errors.New("Pool.GetMetadataVDIs not mocked")
	return
}
// Get the metadata_VDIs field of the given pool.
func (_class PoolClass) GetMetadataVDIs(sessionID SessionRef, self PoolRef) (_retval []VDIRef, _err error) {
	if (IsMock) {
		return _class.GetMetadataVDIs__mock(sessionID, self)
	}	
	_method := "pool.get_metadata_VDIs"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVDIRefSetToGo(_method + " -> ", _result.Value)
	return
}

func (_class PoolClass) GetRestrictions__mock(sessionID SessionRef, self PoolRef) (_retval map[string]string, _err error) {
	log.Println("Pool.GetRestrictions not mocked")
	_err = errors.New("Pool.GetRestrictions not mocked")
	return
}
// Get the restrictions field of the given pool.
func (_class PoolClass) GetRestrictions(sessionID SessionRef, self PoolRef) (_retval map[string]string, _err error) {
	if (IsMock) {
		return _class.GetRestrictions__mock(sessionID, self)
	}	
	_method := "pool.get_restrictions"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PoolClass) GetVswitchController__mock(sessionID SessionRef, self PoolRef) (_retval string, _err error) {
	log.Println("Pool.GetVswitchController not mocked")
	_err = errors.New("Pool.GetVswitchController not mocked")
	return
}
// Get the vswitch_controller field of the given pool.
func (_class PoolClass) GetVswitchController(sessionID SessionRef, self PoolRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetVswitchController__mock(sessionID, self)
	}	
	_method := "pool.get_vswitch_controller"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PoolClass) GetRedoLogVdi__mock(sessionID SessionRef, self PoolRef) (_retval VDIRef, _err error) {
	log.Println("Pool.GetRedoLogVdi not mocked")
	_err = errors.New("Pool.GetRedoLogVdi not mocked")
	return
}
// Get the redo_log_vdi field of the given pool.
func (_class PoolClass) GetRedoLogVdi(sessionID SessionRef, self PoolRef) (_retval VDIRef, _err error) {
	if (IsMock) {
		return _class.GetRedoLogVdi__mock(sessionID, self)
	}	
	_method := "pool.get_redo_log_vdi"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PoolClass) GetRedoLogEnabled__mock(sessionID SessionRef, self PoolRef) (_retval bool, _err error) {
	log.Println("Pool.GetRedoLogEnabled not mocked")
	_err = errors.New("Pool.GetRedoLogEnabled not mocked")
	return
}
// Get the redo_log_enabled field of the given pool.
func (_class PoolClass) GetRedoLogEnabled(sessionID SessionRef, self PoolRef) (_retval bool, _err error) {
	if (IsMock) {
		return _class.GetRedoLogEnabled__mock(sessionID, self)
	}	
	_method := "pool.get_redo_log_enabled"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PoolClass) GetWlbVerifyCert__mock(sessionID SessionRef, self PoolRef) (_retval bool, _err error) {
	log.Println("Pool.GetWlbVerifyCert not mocked")
	_err = errors.New("Pool.GetWlbVerifyCert not mocked")
	return
}
// Get the wlb_verify_cert field of the given pool.
func (_class PoolClass) GetWlbVerifyCert(sessionID SessionRef, self PoolRef) (_retval bool, _err error) {
	if (IsMock) {
		return _class.GetWlbVerifyCert__mock(sessionID, self)
	}	
	_method := "pool.get_wlb_verify_cert"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PoolClass) GetWlbEnabled__mock(sessionID SessionRef, self PoolRef) (_retval bool, _err error) {
	log.Println("Pool.GetWlbEnabled not mocked")
	_err = errors.New("Pool.GetWlbEnabled not mocked")
	return
}
// Get the wlb_enabled field of the given pool.
func (_class PoolClass) GetWlbEnabled(sessionID SessionRef, self PoolRef) (_retval bool, _err error) {
	if (IsMock) {
		return _class.GetWlbEnabled__mock(sessionID, self)
	}	
	_method := "pool.get_wlb_enabled"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PoolClass) GetWlbUsername__mock(sessionID SessionRef, self PoolRef) (_retval string, _err error) {
	log.Println("Pool.GetWlbUsername not mocked")
	_err = errors.New("Pool.GetWlbUsername not mocked")
	return
}
// Get the wlb_username field of the given pool.
func (_class PoolClass) GetWlbUsername(sessionID SessionRef, self PoolRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetWlbUsername__mock(sessionID, self)
	}	
	_method := "pool.get_wlb_username"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PoolClass) GetWlbURL__mock(sessionID SessionRef, self PoolRef) (_retval string, _err error) {
	log.Println("Pool.GetWlbURL not mocked")
	_err = errors.New("Pool.GetWlbURL not mocked")
	return
}
// Get the wlb_url field of the given pool.
func (_class PoolClass) GetWlbURL(sessionID SessionRef, self PoolRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetWlbURL__mock(sessionID, self)
	}	
	_method := "pool.get_wlb_url"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PoolClass) GetHealthCheckConfig__mock(sessionID SessionRef, self PoolRef) (_retval map[string]string, _err error) {
	log.Println("Pool.GetHealthCheckConfig not mocked")
	_err = errors.New("Pool.GetHealthCheckConfig not mocked")
	return
}
// Get the health_check_config field of the given pool.
func (_class PoolClass) GetHealthCheckConfig(sessionID SessionRef, self PoolRef) (_retval map[string]string, _err error) {
	if (IsMock) {
		return _class.GetHealthCheckConfig__mock(sessionID, self)
	}	
	_method := "pool.get_health_check_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PoolClass) GetGuiConfig__mock(sessionID SessionRef, self PoolRef) (_retval map[string]string, _err error) {
	log.Println("Pool.GetGuiConfig not mocked")
	_err = errors.New("Pool.GetGuiConfig not mocked")
	return
}
// Get the gui_config field of the given pool.
func (_class PoolClass) GetGuiConfig(sessionID SessionRef, self PoolRef) (_retval map[string]string, _err error) {
	if (IsMock) {
		return _class.GetGuiConfig__mock(sessionID, self)
	}	
	_method := "pool.get_gui_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PoolClass) GetTags__mock(sessionID SessionRef, self PoolRef) (_retval []string, _err error) {
	log.Println("Pool.GetTags not mocked")
	_err = errors.New("Pool.GetTags not mocked")
	return
}
// Get the tags field of the given pool.
func (_class PoolClass) GetTags(sessionID SessionRef, self PoolRef) (_retval []string, _err error) {
	if (IsMock) {
		return _class.GetTags__mock(sessionID, self)
	}	
	_method := "pool.get_tags"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PoolClass) GetBlobs__mock(sessionID SessionRef, self PoolRef) (_retval map[string]BlobRef, _err error) {
	log.Println("Pool.GetBlobs not mocked")
	_err = errors.New("Pool.GetBlobs not mocked")
	return
}
// Get the blobs field of the given pool.
func (_class PoolClass) GetBlobs(sessionID SessionRef, self PoolRef) (_retval map[string]BlobRef, _err error) {
	if (IsMock) {
		return _class.GetBlobs__mock(sessionID, self)
	}	
	_method := "pool.get_blobs"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PoolClass) GetHaOvercommitted__mock(sessionID SessionRef, self PoolRef) (_retval bool, _err error) {
	log.Println("Pool.GetHaOvercommitted not mocked")
	_err = errors.New("Pool.GetHaOvercommitted not mocked")
	return
}
// Get the ha_overcommitted field of the given pool.
func (_class PoolClass) GetHaOvercommitted(sessionID SessionRef, self PoolRef) (_retval bool, _err error) {
	if (IsMock) {
		return _class.GetHaOvercommitted__mock(sessionID, self)
	}	
	_method := "pool.get_ha_overcommitted"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PoolClass) GetHaAllowOvercommit__mock(sessionID SessionRef, self PoolRef) (_retval bool, _err error) {
	log.Println("Pool.GetHaAllowOvercommit not mocked")
	_err = errors.New("Pool.GetHaAllowOvercommit not mocked")
	return
}
// Get the ha_allow_overcommit field of the given pool.
func (_class PoolClass) GetHaAllowOvercommit(sessionID SessionRef, self PoolRef) (_retval bool, _err error) {
	if (IsMock) {
		return _class.GetHaAllowOvercommit__mock(sessionID, self)
	}	
	_method := "pool.get_ha_allow_overcommit"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PoolClass) GetHaPlanExistsFor__mock(sessionID SessionRef, self PoolRef) (_retval int, _err error) {
	log.Println("Pool.GetHaPlanExistsFor not mocked")
	_err = errors.New("Pool.GetHaPlanExistsFor not mocked")
	return
}
// Get the ha_plan_exists_for field of the given pool.
func (_class PoolClass) GetHaPlanExistsFor(sessionID SessionRef, self PoolRef) (_retval int, _err error) {
	if (IsMock) {
		return _class.GetHaPlanExistsFor__mock(sessionID, self)
	}	
	_method := "pool.get_ha_plan_exists_for"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PoolClass) GetHaHostFailuresToTolerate__mock(sessionID SessionRef, self PoolRef) (_retval int, _err error) {
	log.Println("Pool.GetHaHostFailuresToTolerate not mocked")
	_err = errors.New("Pool.GetHaHostFailuresToTolerate not mocked")
	return
}
// Get the ha_host_failures_to_tolerate field of the given pool.
func (_class PoolClass) GetHaHostFailuresToTolerate(sessionID SessionRef, self PoolRef) (_retval int, _err error) {
	if (IsMock) {
		return _class.GetHaHostFailuresToTolerate__mock(sessionID, self)
	}	
	_method := "pool.get_ha_host_failures_to_tolerate"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PoolClass) GetHaStatefiles__mock(sessionID SessionRef, self PoolRef) (_retval []string, _err error) {
	log.Println("Pool.GetHaStatefiles not mocked")
	_err = errors.New("Pool.GetHaStatefiles not mocked")
	return
}
// Get the ha_statefiles field of the given pool.
func (_class PoolClass) GetHaStatefiles(sessionID SessionRef, self PoolRef) (_retval []string, _err error) {
	if (IsMock) {
		return _class.GetHaStatefiles__mock(sessionID, self)
	}	
	_method := "pool.get_ha_statefiles"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PoolClass) GetHaConfiguration__mock(sessionID SessionRef, self PoolRef) (_retval map[string]string, _err error) {
	log.Println("Pool.GetHaConfiguration not mocked")
	_err = errors.New("Pool.GetHaConfiguration not mocked")
	return
}
// Get the ha_configuration field of the given pool.
func (_class PoolClass) GetHaConfiguration(sessionID SessionRef, self PoolRef) (_retval map[string]string, _err error) {
	if (IsMock) {
		return _class.GetHaConfiguration__mock(sessionID, self)
	}	
	_method := "pool.get_ha_configuration"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PoolClass) GetHaEnabled__mock(sessionID SessionRef, self PoolRef) (_retval bool, _err error) {
	log.Println("Pool.GetHaEnabled not mocked")
	_err = errors.New("Pool.GetHaEnabled not mocked")
	return
}
// Get the ha_enabled field of the given pool.
func (_class PoolClass) GetHaEnabled(sessionID SessionRef, self PoolRef) (_retval bool, _err error) {
	if (IsMock) {
		return _class.GetHaEnabled__mock(sessionID, self)
	}	
	_method := "pool.get_ha_enabled"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PoolClass) GetOtherConfig__mock(sessionID SessionRef, self PoolRef) (_retval map[string]string, _err error) {
	log.Println("Pool.GetOtherConfig not mocked")
	_err = errors.New("Pool.GetOtherConfig not mocked")
	return
}
// Get the other_config field of the given pool.
func (_class PoolClass) GetOtherConfig(sessionID SessionRef, self PoolRef) (_retval map[string]string, _err error) {
	if (IsMock) {
		return _class.GetOtherConfig__mock(sessionID, self)
	}	
	_method := "pool.get_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PoolClass) GetCrashDumpSR__mock(sessionID SessionRef, self PoolRef) (_retval SRRef, _err error) {
	log.Println("Pool.GetCrashDumpSR not mocked")
	_err = errors.New("Pool.GetCrashDumpSR not mocked")
	return
}
// Get the crash_dump_SR field of the given pool.
func (_class PoolClass) GetCrashDumpSR(sessionID SessionRef, self PoolRef) (_retval SRRef, _err error) {
	if (IsMock) {
		return _class.GetCrashDumpSR__mock(sessionID, self)
	}	
	_method := "pool.get_crash_dump_SR"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PoolClass) GetSuspendImageSR__mock(sessionID SessionRef, self PoolRef) (_retval SRRef, _err error) {
	log.Println("Pool.GetSuspendImageSR not mocked")
	_err = errors.New("Pool.GetSuspendImageSR not mocked")
	return
}
// Get the suspend_image_SR field of the given pool.
func (_class PoolClass) GetSuspendImageSR(sessionID SessionRef, self PoolRef) (_retval SRRef, _err error) {
	if (IsMock) {
		return _class.GetSuspendImageSR__mock(sessionID, self)
	}	
	_method := "pool.get_suspend_image_SR"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PoolClass) GetDefaultSR__mock(sessionID SessionRef, self PoolRef) (_retval SRRef, _err error) {
	log.Println("Pool.GetDefaultSR not mocked")
	_err = errors.New("Pool.GetDefaultSR not mocked")
	return
}
// Get the default_SR field of the given pool.
func (_class PoolClass) GetDefaultSR(sessionID SessionRef, self PoolRef) (_retval SRRef, _err error) {
	if (IsMock) {
		return _class.GetDefaultSR__mock(sessionID, self)
	}	
	_method := "pool.get_default_SR"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PoolClass) GetMaster__mock(sessionID SessionRef, self PoolRef) (_retval HostRef, _err error) {
	log.Println("Pool.GetMaster not mocked")
	_err = errors.New("Pool.GetMaster not mocked")
	return
}
// Get the master field of the given pool.
func (_class PoolClass) GetMaster(sessionID SessionRef, self PoolRef) (_retval HostRef, _err error) {
	if (IsMock) {
		return _class.GetMaster__mock(sessionID, self)
	}	
	_method := "pool.get_master"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PoolClass) GetNameDescription__mock(sessionID SessionRef, self PoolRef) (_retval string, _err error) {
	log.Println("Pool.GetNameDescription not mocked")
	_err = errors.New("Pool.GetNameDescription not mocked")
	return
}
// Get the name_description field of the given pool.
func (_class PoolClass) GetNameDescription(sessionID SessionRef, self PoolRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetNameDescription__mock(sessionID, self)
	}	
	_method := "pool.get_name_description"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PoolClass) GetNameLabel__mock(sessionID SessionRef, self PoolRef) (_retval string, _err error) {
	log.Println("Pool.GetNameLabel not mocked")
	_err = errors.New("Pool.GetNameLabel not mocked")
	return
}
// Get the name_label field of the given pool.
func (_class PoolClass) GetNameLabel(sessionID SessionRef, self PoolRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetNameLabel__mock(sessionID, self)
	}	
	_method := "pool.get_name_label"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PoolClass) GetUUID__mock(sessionID SessionRef, self PoolRef) (_retval string, _err error) {
	log.Println("Pool.GetUUID not mocked")
	_err = errors.New("Pool.GetUUID not mocked")
	return
}
// Get the uuid field of the given pool.
func (_class PoolClass) GetUUID(sessionID SessionRef, self PoolRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetUUID__mock(sessionID, self)
	}	
	_method := "pool.get_uuid"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PoolClass) GetByUUID__mock(sessionID SessionRef, uuid string) (_retval PoolRef, _err error) {
	log.Println("Pool.GetByUUID not mocked")
	_err = errors.New("Pool.GetByUUID not mocked")
	return
}
// Get a reference to the pool instance with the specified UUID.
func (_class PoolClass) GetByUUID(sessionID SessionRef, uuid string) (_retval PoolRef, _err error) {
	if (IsMock) {
		return _class.GetByUUID__mock(sessionID, uuid)
	}	
	_method := "pool.get_by_uuid"
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
	_retval, _err = convertPoolRefToGo(_method + " -> ", _result.Value)
	return
}

func (_class PoolClass) GetRecord__mock(sessionID SessionRef, self PoolRef) (_retval PoolRecord, _err error) {
	log.Println("Pool.GetRecord not mocked")
	_err = errors.New("Pool.GetRecord not mocked")
	return
}
// Get a record containing the current state of the given pool.
func (_class PoolClass) GetRecord(sessionID SessionRef, self PoolRef) (_retval PoolRecord, _err error) {
	if (IsMock) {
		return _class.GetRecord__mock(sessionID, self)
	}	
	_method := "pool.get_record"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertPoolRecordToGo(_method + " -> ", _result.Value)
	return
}
