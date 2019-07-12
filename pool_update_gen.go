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

type UpdateAfterApplyGuidance string

const (
  // This update requires HVM guests to be restarted once applied.
	UpdateAfterApplyGuidanceRestartHVM UpdateAfterApplyGuidance = "restartHVM"
  // This update requires PV guests to be restarted once applied.
	UpdateAfterApplyGuidanceRestartPV UpdateAfterApplyGuidance = "restartPV"
  // This update requires the host to be restarted once applied.
	UpdateAfterApplyGuidanceRestartHost UpdateAfterApplyGuidance = "restartHost"
  // This update requires XAPI to be restarted once applied.
	UpdateAfterApplyGuidanceRestartXAPI UpdateAfterApplyGuidance = "restartXAPI"
)

type LivepatchStatus string

const (
  // An applicable live patch exists for every required component
	LivepatchStatusOkLivepatchComplete LivepatchStatus = "ok_livepatch_complete"
  // An applicable live patch exists but it is not sufficient
	LivepatchStatusOkLivepatchIncomplete LivepatchStatus = "ok_livepatch_incomplete"
  // There is no applicable live patch
	LivepatchStatusOk LivepatchStatus = "ok"
)

type PoolUpdateRecord struct {
  // Unique identifier/object reference
	UUID string
  // a human-readable name
	NameLabel string
  // a notes field containing human-readable description
	NameDescription string
  // Update version number
	Version string
  // Size of the update in bytes
	InstallationSize int
  // GPG key of the update
	Key string
  // What the client should do after this update has been applied.
	AfterApplyGuidance []UpdateAfterApplyGuidance
  // VDI the update was uploaded to
	Vdi VDIRef
  // The hosts that have applied this update.
	Hosts []HostRef
  // additional configuration
	OtherConfig map[string]string
  // Flag - if true, all hosts in a pool must apply this update
	EnforceHomogeneity bool
}

type PoolUpdateRef string

// Pool-wide updates to the host software
type PoolUpdateClass struct {
	client *Client
}

func (_class PoolUpdateClass) GetAllRecords__mock(sessionID SessionRef) (_retval map[PoolUpdateRef]PoolUpdateRecord, _err error) {
	log.Println("PoolUpdate.GetAllRecords not mocked")
	_err = errors.New("PoolUpdate.GetAllRecords not mocked")
	return
}
// Return a map of pool_update references to pool_update records for all pool_updates known to the system.
func (_class PoolUpdateClass) GetAllRecords(sessionID SessionRef) (_retval map[PoolUpdateRef]PoolUpdateRecord, _err error) {
	if (IsMock) {
		return _class.GetAllRecords__mock(sessionID)
	}	
	_method := "pool_update.get_all_records"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertPoolUpdateRefToPoolUpdateRecordMapToGo(_method + " -> ", _result.Value)
	return
}

func (_class PoolUpdateClass) GetAll__mock(sessionID SessionRef) (_retval []PoolUpdateRef, _err error) {
	log.Println("PoolUpdate.GetAll not mocked")
	_err = errors.New("PoolUpdate.GetAll not mocked")
	return
}
// Return a list of all the pool_updates known to the system.
func (_class PoolUpdateClass) GetAll(sessionID SessionRef) (_retval []PoolUpdateRef, _err error) {
	if (IsMock) {
		return _class.GetAll__mock(sessionID)
	}	
	_method := "pool_update.get_all"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertPoolUpdateRefSetToGo(_method + " -> ", _result.Value)
	return
}

func (_class PoolUpdateClass) Destroy__mock(sessionID SessionRef, self PoolUpdateRef) (_err error) {
	log.Println("PoolUpdate.Destroy not mocked")
	_err = errors.New("PoolUpdate.Destroy not mocked")
	return
}
// Removes the database entry. Only works on unapplied update.
func (_class PoolUpdateClass) Destroy(sessionID SessionRef, self PoolUpdateRef) (_err error) {
	if (IsMock) {
		return _class.Destroy__mock(sessionID, self)
	}	
	_method := "pool_update.destroy"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolUpdateRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg)
	return
}

func (_class PoolUpdateClass) PoolClean__mock(sessionID SessionRef, self PoolUpdateRef) (_err error) {
	log.Println("PoolUpdate.PoolClean not mocked")
	_err = errors.New("PoolUpdate.PoolClean not mocked")
	return
}
// Removes the update's files from all hosts in the pool, but does not revert the update
func (_class PoolUpdateClass) PoolClean(sessionID SessionRef, self PoolUpdateRef) (_err error) {
	if (IsMock) {
		return _class.PoolClean__mock(sessionID, self)
	}	
	_method := "pool_update.pool_clean"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolUpdateRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg)
	return
}

func (_class PoolUpdateClass) PoolApply__mock(sessionID SessionRef, self PoolUpdateRef) (_err error) {
	log.Println("PoolUpdate.PoolApply not mocked")
	_err = errors.New("PoolUpdate.PoolApply not mocked")
	return
}
// Apply the selected update to all hosts in the pool
func (_class PoolUpdateClass) PoolApply(sessionID SessionRef, self PoolUpdateRef) (_err error) {
	if (IsMock) {
		return _class.PoolApply__mock(sessionID, self)
	}	
	_method := "pool_update.pool_apply"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolUpdateRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg)
	return
}

func (_class PoolUpdateClass) Apply__mock(sessionID SessionRef, self PoolUpdateRef, host HostRef) (_err error) {
	log.Println("PoolUpdate.Apply not mocked")
	_err = errors.New("PoolUpdate.Apply not mocked")
	return
}
// Apply the selected update to a host
func (_class PoolUpdateClass) Apply(sessionID SessionRef, self PoolUpdateRef, host HostRef) (_err error) {
	if (IsMock) {
		return _class.Apply__mock(sessionID, self, host)
	}	
	_method := "pool_update.apply"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolUpdateRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PoolUpdateClass) Precheck__mock(sessionID SessionRef, self PoolUpdateRef, host HostRef) (_retval LivepatchStatus, _err error) {
	log.Println("PoolUpdate.Precheck not mocked")
	_err = errors.New("PoolUpdate.Precheck not mocked")
	return
}
// Execute the precheck stage of the selected update on a host
func (_class PoolUpdateClass) Precheck(sessionID SessionRef, self PoolUpdateRef, host HostRef) (_retval LivepatchStatus, _err error) {
	if (IsMock) {
		return _class.Precheck__mock(sessionID, self, host)
	}	
	_method := "pool_update.precheck"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolUpdateRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_hostArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "host"), host)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg, _hostArg)
	if _err != nil {
		return
	}
	_retval, _err = convertEnumLivepatchStatusToGo(_method + " -> ", _result.Value)
	return
}

func (_class PoolUpdateClass) Introduce__mock(sessionID SessionRef, vdi VDIRef) (_retval PoolUpdateRef, _err error) {
	log.Println("PoolUpdate.Introduce not mocked")
	_err = errors.New("PoolUpdate.Introduce not mocked")
	return
}
// Introduce update VDI
func (_class PoolUpdateClass) Introduce(sessionID SessionRef, vdi VDIRef) (_retval PoolUpdateRef, _err error) {
	if (IsMock) {
		return _class.Introduce__mock(sessionID, vdi)
	}	
	_method := "pool_update.introduce"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_vdiArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "vdi"), vdi)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _vdiArg)
	if _err != nil {
		return
	}
	_retval, _err = convertPoolUpdateRefToGo(_method + " -> ", _result.Value)
	return
}

func (_class PoolUpdateClass) RemoveFromOtherConfig__mock(sessionID SessionRef, self PoolUpdateRef, key string) (_err error) {
	log.Println("PoolUpdate.RemoveFromOtherConfig not mocked")
	_err = errors.New("PoolUpdate.RemoveFromOtherConfig not mocked")
	return
}
// Remove the given key and its corresponding value from the other_config field of the given pool_update.  If the key is not in that Map, then do nothing.
func (_class PoolUpdateClass) RemoveFromOtherConfig(sessionID SessionRef, self PoolUpdateRef, key string) (_err error) {
	if (IsMock) {
		return _class.RemoveFromOtherConfig__mock(sessionID, self, key)
	}	
	_method := "pool_update.remove_from_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolUpdateRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PoolUpdateClass) AddToOtherConfig__mock(sessionID SessionRef, self PoolUpdateRef, key string, value string) (_err error) {
	log.Println("PoolUpdate.AddToOtherConfig not mocked")
	_err = errors.New("PoolUpdate.AddToOtherConfig not mocked")
	return
}
// Add the given key-value pair to the other_config field of the given pool_update.
func (_class PoolUpdateClass) AddToOtherConfig(sessionID SessionRef, self PoolUpdateRef, key string, value string) (_err error) {
	if (IsMock) {
		return _class.AddToOtherConfig__mock(sessionID, self, key, value)
	}	
	_method := "pool_update.add_to_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolUpdateRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PoolUpdateClass) SetOtherConfig__mock(sessionID SessionRef, self PoolUpdateRef, value map[string]string) (_err error) {
	log.Println("PoolUpdate.SetOtherConfig not mocked")
	_err = errors.New("PoolUpdate.SetOtherConfig not mocked")
	return
}
// Set the other_config field of the given pool_update.
func (_class PoolUpdateClass) SetOtherConfig(sessionID SessionRef, self PoolUpdateRef, value map[string]string) (_err error) {
	if (IsMock) {
		return _class.SetOtherConfig__mock(sessionID, self, value)
	}	
	_method := "pool_update.set_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolUpdateRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PoolUpdateClass) GetEnforceHomogeneity__mock(sessionID SessionRef, self PoolUpdateRef) (_retval bool, _err error) {
	log.Println("PoolUpdate.GetEnforceHomogeneity not mocked")
	_err = errors.New("PoolUpdate.GetEnforceHomogeneity not mocked")
	return
}
// Get the enforce_homogeneity field of the given pool_update.
func (_class PoolUpdateClass) GetEnforceHomogeneity(sessionID SessionRef, self PoolUpdateRef) (_retval bool, _err error) {
	if (IsMock) {
		return _class.GetEnforceHomogeneity__mock(sessionID, self)
	}	
	_method := "pool_update.get_enforce_homogeneity"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolUpdateRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PoolUpdateClass) GetOtherConfig__mock(sessionID SessionRef, self PoolUpdateRef) (_retval map[string]string, _err error) {
	log.Println("PoolUpdate.GetOtherConfig not mocked")
	_err = errors.New("PoolUpdate.GetOtherConfig not mocked")
	return
}
// Get the other_config field of the given pool_update.
func (_class PoolUpdateClass) GetOtherConfig(sessionID SessionRef, self PoolUpdateRef) (_retval map[string]string, _err error) {
	if (IsMock) {
		return _class.GetOtherConfig__mock(sessionID, self)
	}	
	_method := "pool_update.get_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolUpdateRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PoolUpdateClass) GetHosts__mock(sessionID SessionRef, self PoolUpdateRef) (_retval []HostRef, _err error) {
	log.Println("PoolUpdate.GetHosts not mocked")
	_err = errors.New("PoolUpdate.GetHosts not mocked")
	return
}
// Get the hosts field of the given pool_update.
func (_class PoolUpdateClass) GetHosts(sessionID SessionRef, self PoolUpdateRef) (_retval []HostRef, _err error) {
	if (IsMock) {
		return _class.GetHosts__mock(sessionID, self)
	}	
	_method := "pool_update.get_hosts"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolUpdateRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertHostRefSetToGo(_method + " -> ", _result.Value)
	return
}

func (_class PoolUpdateClass) GetVdi__mock(sessionID SessionRef, self PoolUpdateRef) (_retval VDIRef, _err error) {
	log.Println("PoolUpdate.GetVdi not mocked")
	_err = errors.New("PoolUpdate.GetVdi not mocked")
	return
}
// Get the vdi field of the given pool_update.
func (_class PoolUpdateClass) GetVdi(sessionID SessionRef, self PoolUpdateRef) (_retval VDIRef, _err error) {
	if (IsMock) {
		return _class.GetVdi__mock(sessionID, self)
	}	
	_method := "pool_update.get_vdi"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolUpdateRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PoolUpdateClass) GetAfterApplyGuidance__mock(sessionID SessionRef, self PoolUpdateRef) (_retval []UpdateAfterApplyGuidance, _err error) {
	log.Println("PoolUpdate.GetAfterApplyGuidance not mocked")
	_err = errors.New("PoolUpdate.GetAfterApplyGuidance not mocked")
	return
}
// Get the after_apply_guidance field of the given pool_update.
func (_class PoolUpdateClass) GetAfterApplyGuidance(sessionID SessionRef, self PoolUpdateRef) (_retval []UpdateAfterApplyGuidance, _err error) {
	if (IsMock) {
		return _class.GetAfterApplyGuidance__mock(sessionID, self)
	}	
	_method := "pool_update.get_after_apply_guidance"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolUpdateRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertEnumUpdateAfterApplyGuidanceSetToGo(_method + " -> ", _result.Value)
	return
}

func (_class PoolUpdateClass) GetKey__mock(sessionID SessionRef, self PoolUpdateRef) (_retval string, _err error) {
	log.Println("PoolUpdate.GetKey not mocked")
	_err = errors.New("PoolUpdate.GetKey not mocked")
	return
}
// Get the key field of the given pool_update.
func (_class PoolUpdateClass) GetKey(sessionID SessionRef, self PoolUpdateRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetKey__mock(sessionID, self)
	}	
	_method := "pool_update.get_key"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolUpdateRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PoolUpdateClass) GetInstallationSize__mock(sessionID SessionRef, self PoolUpdateRef) (_retval int, _err error) {
	log.Println("PoolUpdate.GetInstallationSize not mocked")
	_err = errors.New("PoolUpdate.GetInstallationSize not mocked")
	return
}
// Get the installation_size field of the given pool_update.
func (_class PoolUpdateClass) GetInstallationSize(sessionID SessionRef, self PoolUpdateRef) (_retval int, _err error) {
	if (IsMock) {
		return _class.GetInstallationSize__mock(sessionID, self)
	}	
	_method := "pool_update.get_installation_size"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolUpdateRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PoolUpdateClass) GetVersion__mock(sessionID SessionRef, self PoolUpdateRef) (_retval string, _err error) {
	log.Println("PoolUpdate.GetVersion not mocked")
	_err = errors.New("PoolUpdate.GetVersion not mocked")
	return
}
// Get the version field of the given pool_update.
func (_class PoolUpdateClass) GetVersion(sessionID SessionRef, self PoolUpdateRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetVersion__mock(sessionID, self)
	}	
	_method := "pool_update.get_version"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolUpdateRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PoolUpdateClass) GetNameDescription__mock(sessionID SessionRef, self PoolUpdateRef) (_retval string, _err error) {
	log.Println("PoolUpdate.GetNameDescription not mocked")
	_err = errors.New("PoolUpdate.GetNameDescription not mocked")
	return
}
// Get the name/description field of the given pool_update.
func (_class PoolUpdateClass) GetNameDescription(sessionID SessionRef, self PoolUpdateRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetNameDescription__mock(sessionID, self)
	}	
	_method := "pool_update.get_name_description"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolUpdateRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PoolUpdateClass) GetNameLabel__mock(sessionID SessionRef, self PoolUpdateRef) (_retval string, _err error) {
	log.Println("PoolUpdate.GetNameLabel not mocked")
	_err = errors.New("PoolUpdate.GetNameLabel not mocked")
	return
}
// Get the name/label field of the given pool_update.
func (_class PoolUpdateClass) GetNameLabel(sessionID SessionRef, self PoolUpdateRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetNameLabel__mock(sessionID, self)
	}	
	_method := "pool_update.get_name_label"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolUpdateRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PoolUpdateClass) GetUUID__mock(sessionID SessionRef, self PoolUpdateRef) (_retval string, _err error) {
	log.Println("PoolUpdate.GetUUID not mocked")
	_err = errors.New("PoolUpdate.GetUUID not mocked")
	return
}
// Get the uuid field of the given pool_update.
func (_class PoolUpdateClass) GetUUID(sessionID SessionRef, self PoolUpdateRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetUUID__mock(sessionID, self)
	}	
	_method := "pool_update.get_uuid"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolUpdateRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PoolUpdateClass) GetByNameLabel__mock(sessionID SessionRef, label string) (_retval []PoolUpdateRef, _err error) {
	log.Println("PoolUpdate.GetByNameLabel not mocked")
	_err = errors.New("PoolUpdate.GetByNameLabel not mocked")
	return
}
// Get all the pool_update instances with the given label.
func (_class PoolUpdateClass) GetByNameLabel(sessionID SessionRef, label string) (_retval []PoolUpdateRef, _err error) {
	if (IsMock) {
		return _class.GetByNameLabel__mock(sessionID, label)
	}	
	_method := "pool_update.get_by_name_label"
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
	_retval, _err = convertPoolUpdateRefSetToGo(_method + " -> ", _result.Value)
	return
}

func (_class PoolUpdateClass) GetByUUID__mock(sessionID SessionRef, uuid string) (_retval PoolUpdateRef, _err error) {
	log.Println("PoolUpdate.GetByUUID not mocked")
	_err = errors.New("PoolUpdate.GetByUUID not mocked")
	return
}
// Get a reference to the pool_update instance with the specified UUID.
func (_class PoolUpdateClass) GetByUUID(sessionID SessionRef, uuid string) (_retval PoolUpdateRef, _err error) {
	if (IsMock) {
		return _class.GetByUUID__mock(sessionID, uuid)
	}	
	_method := "pool_update.get_by_uuid"
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
	_retval, _err = convertPoolUpdateRefToGo(_method + " -> ", _result.Value)
	return
}

func (_class PoolUpdateClass) GetRecord__mock(sessionID SessionRef, self PoolUpdateRef) (_retval PoolUpdateRecord, _err error) {
	log.Println("PoolUpdate.GetRecord not mocked")
	_err = errors.New("PoolUpdate.GetRecord not mocked")
	return
}
// Get a record containing the current state of the given pool_update.
func (_class PoolUpdateClass) GetRecord(sessionID SessionRef, self PoolUpdateRef) (_retval PoolUpdateRecord, _err error) {
	if (IsMock) {
		return _class.GetRecord__mock(sessionID, self)
	}	
	_method := "pool_update.get_record"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolUpdateRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertPoolUpdateRecordToGo(_method + " -> ", _result.Value)
	return
}
