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

type AfterApplyGuidance string

const (
  // This patch requires HVM guests to be restarted once applied.
	AfterApplyGuidanceRestartHVM AfterApplyGuidance = "restartHVM"
  // This patch requires PV guests to be restarted once applied.
	AfterApplyGuidanceRestartPV AfterApplyGuidance = "restartPV"
  // This patch requires the host to be restarted once applied.
	AfterApplyGuidanceRestartHost AfterApplyGuidance = "restartHost"
  // This patch requires XAPI to be restarted once applied.
	AfterApplyGuidanceRestartXAPI AfterApplyGuidance = "restartXAPI"
)

type PoolPatchRecord struct {
  // Unique identifier/object reference
	UUID string
  // a human-readable name
	NameLabel string
  // a notes field containing human-readable description
	NameDescription string
  // Patch version number
	Version string
  // Size of the patch
	Size int
  // This patch should be applied across the entire pool
	PoolApplied bool
  // This hosts this patch is applied to.
	HostPatches []HostPatchRef
  // What the client should do after this patch has been applied.
	AfterApplyGuidance []AfterApplyGuidance
  // A reference to the associated pool_update object
	PoolUpdate PoolUpdateRef
  // additional configuration
	OtherConfig map[string]string
}

type PoolPatchRef string

// Pool-wide patches
type PoolPatchClass struct {
	client *Client
}

func (_class PoolPatchClass) GetAllRecords__mock(sessionID SessionRef) (_retval map[PoolPatchRef]PoolPatchRecord, _err error) {
	log.Println("PoolPatch.GetAllRecords not mocked")
	_err = errors.New("PoolPatch.GetAllRecords not mocked")
	return
}
// Return a map of pool_patch references to pool_patch records for all pool_patchs known to the system.
func (_class PoolPatchClass) GetAllRecords(sessionID SessionRef) (_retval map[PoolPatchRef]PoolPatchRecord, _err error) {
	if (IsMock) {
		return _class.GetAllRecords__mock(sessionID)
	}	
	_method := "pool_patch.get_all_records"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertPoolPatchRefToPoolPatchRecordMapToGo(_method + " -> ", _result.Value)
	return
}

func (_class PoolPatchClass) GetAll__mock(sessionID SessionRef) (_retval []PoolPatchRef, _err error) {
	log.Println("PoolPatch.GetAll not mocked")
	_err = errors.New("PoolPatch.GetAll not mocked")
	return
}
// Return a list of all the pool_patchs known to the system.
func (_class PoolPatchClass) GetAll(sessionID SessionRef) (_retval []PoolPatchRef, _err error) {
	if (IsMock) {
		return _class.GetAll__mock(sessionID)
	}	
	_method := "pool_patch.get_all"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertPoolPatchRefSetToGo(_method + " -> ", _result.Value)
	return
}

func (_class PoolPatchClass) CleanOnHost__mock(sessionID SessionRef, self PoolPatchRef, host HostRef) (_err error) {
	log.Println("PoolPatch.CleanOnHost not mocked")
	_err = errors.New("PoolPatch.CleanOnHost not mocked")
	return
}
// Removes the patch's files from the specified host
func (_class PoolPatchClass) CleanOnHost(sessionID SessionRef, self PoolPatchRef, host HostRef) (_err error) {
	if (IsMock) {
		return _class.CleanOnHost__mock(sessionID, self, host)
	}	
	_method := "pool_patch.clean_on_host"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolPatchRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PoolPatchClass) Destroy__mock(sessionID SessionRef, self PoolPatchRef) (_err error) {
	log.Println("PoolPatch.Destroy not mocked")
	_err = errors.New("PoolPatch.Destroy not mocked")
	return
}
// Removes the patch's files from all hosts in the pool, and removes the database entries.  Only works on unapplied patches.
func (_class PoolPatchClass) Destroy(sessionID SessionRef, self PoolPatchRef) (_err error) {
	if (IsMock) {
		return _class.Destroy__mock(sessionID, self)
	}	
	_method := "pool_patch.destroy"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolPatchRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg)
	return
}

func (_class PoolPatchClass) PoolClean__mock(sessionID SessionRef, self PoolPatchRef) (_err error) {
	log.Println("PoolPatch.PoolClean not mocked")
	_err = errors.New("PoolPatch.PoolClean not mocked")
	return
}
// Removes the patch's files from all hosts in the pool, but does not remove the database entries
func (_class PoolPatchClass) PoolClean(sessionID SessionRef, self PoolPatchRef) (_err error) {
	if (IsMock) {
		return _class.PoolClean__mock(sessionID, self)
	}	
	_method := "pool_patch.pool_clean"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolPatchRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg)
	return
}

func (_class PoolPatchClass) Clean__mock(sessionID SessionRef, self PoolPatchRef) (_err error) {
	log.Println("PoolPatch.Clean not mocked")
	_err = errors.New("PoolPatch.Clean not mocked")
	return
}
// Removes the patch's files from the server
func (_class PoolPatchClass) Clean(sessionID SessionRef, self PoolPatchRef) (_err error) {
	if (IsMock) {
		return _class.Clean__mock(sessionID, self)
	}	
	_method := "pool_patch.clean"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolPatchRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg)
	return
}

func (_class PoolPatchClass) Precheck__mock(sessionID SessionRef, self PoolPatchRef, host HostRef) (_retval string, _err error) {
	log.Println("PoolPatch.Precheck not mocked")
	_err = errors.New("PoolPatch.Precheck not mocked")
	return
}
// Execute the precheck stage of the selected patch on a host and return its output
func (_class PoolPatchClass) Precheck(sessionID SessionRef, self PoolPatchRef, host HostRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.Precheck__mock(sessionID, self, host)
	}	
	_method := "pool_patch.precheck"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolPatchRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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
	_retval, _err = convertStringToGo(_method + " -> ", _result.Value)
	return
}

func (_class PoolPatchClass) PoolApply__mock(sessionID SessionRef, self PoolPatchRef) (_err error) {
	log.Println("PoolPatch.PoolApply not mocked")
	_err = errors.New("PoolPatch.PoolApply not mocked")
	return
}
// Apply the selected patch to all hosts in the pool and return a map of host_ref -> patch output
func (_class PoolPatchClass) PoolApply(sessionID SessionRef, self PoolPatchRef) (_err error) {
	if (IsMock) {
		return _class.PoolApply__mock(sessionID, self)
	}	
	_method := "pool_patch.pool_apply"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolPatchRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg)
	return
}

func (_class PoolPatchClass) Apply__mock(sessionID SessionRef, self PoolPatchRef, host HostRef) (_retval string, _err error) {
	log.Println("PoolPatch.Apply not mocked")
	_err = errors.New("PoolPatch.Apply not mocked")
	return
}
// Apply the selected patch to a host and return its output
func (_class PoolPatchClass) Apply(sessionID SessionRef, self PoolPatchRef, host HostRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.Apply__mock(sessionID, self, host)
	}	
	_method := "pool_patch.apply"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolPatchRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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
	_retval, _err = convertStringToGo(_method + " -> ", _result.Value)
	return
}

func (_class PoolPatchClass) RemoveFromOtherConfig__mock(sessionID SessionRef, self PoolPatchRef, key string) (_err error) {
	log.Println("PoolPatch.RemoveFromOtherConfig not mocked")
	_err = errors.New("PoolPatch.RemoveFromOtherConfig not mocked")
	return
}
// Remove the given key and its corresponding value from the other_config field of the given pool_patch.  If the key is not in that Map, then do nothing.
func (_class PoolPatchClass) RemoveFromOtherConfig(sessionID SessionRef, self PoolPatchRef, key string) (_err error) {
	if (IsMock) {
		return _class.RemoveFromOtherConfig__mock(sessionID, self, key)
	}	
	_method := "pool_patch.remove_from_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolPatchRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PoolPatchClass) AddToOtherConfig__mock(sessionID SessionRef, self PoolPatchRef, key string, value string) (_err error) {
	log.Println("PoolPatch.AddToOtherConfig not mocked")
	_err = errors.New("PoolPatch.AddToOtherConfig not mocked")
	return
}
// Add the given key-value pair to the other_config field of the given pool_patch.
func (_class PoolPatchClass) AddToOtherConfig(sessionID SessionRef, self PoolPatchRef, key string, value string) (_err error) {
	if (IsMock) {
		return _class.AddToOtherConfig__mock(sessionID, self, key, value)
	}	
	_method := "pool_patch.add_to_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolPatchRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PoolPatchClass) SetOtherConfig__mock(sessionID SessionRef, self PoolPatchRef, value map[string]string) (_err error) {
	log.Println("PoolPatch.SetOtherConfig not mocked")
	_err = errors.New("PoolPatch.SetOtherConfig not mocked")
	return
}
// Set the other_config field of the given pool_patch.
func (_class PoolPatchClass) SetOtherConfig(sessionID SessionRef, self PoolPatchRef, value map[string]string) (_err error) {
	if (IsMock) {
		return _class.SetOtherConfig__mock(sessionID, self, value)
	}	
	_method := "pool_patch.set_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolPatchRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PoolPatchClass) GetOtherConfig__mock(sessionID SessionRef, self PoolPatchRef) (_retval map[string]string, _err error) {
	log.Println("PoolPatch.GetOtherConfig not mocked")
	_err = errors.New("PoolPatch.GetOtherConfig not mocked")
	return
}
// Get the other_config field of the given pool_patch.
func (_class PoolPatchClass) GetOtherConfig(sessionID SessionRef, self PoolPatchRef) (_retval map[string]string, _err error) {
	if (IsMock) {
		return _class.GetOtherConfig__mock(sessionID, self)
	}	
	_method := "pool_patch.get_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolPatchRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PoolPatchClass) GetPoolUpdate__mock(sessionID SessionRef, self PoolPatchRef) (_retval PoolUpdateRef, _err error) {
	log.Println("PoolPatch.GetPoolUpdate not mocked")
	_err = errors.New("PoolPatch.GetPoolUpdate not mocked")
	return
}
// Get the pool_update field of the given pool_patch.
func (_class PoolPatchClass) GetPoolUpdate(sessionID SessionRef, self PoolPatchRef) (_retval PoolUpdateRef, _err error) {
	if (IsMock) {
		return _class.GetPoolUpdate__mock(sessionID, self)
	}	
	_method := "pool_patch.get_pool_update"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolPatchRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertPoolUpdateRefToGo(_method + " -> ", _result.Value)
	return
}

func (_class PoolPatchClass) GetAfterApplyGuidance__mock(sessionID SessionRef, self PoolPatchRef) (_retval []AfterApplyGuidance, _err error) {
	log.Println("PoolPatch.GetAfterApplyGuidance not mocked")
	_err = errors.New("PoolPatch.GetAfterApplyGuidance not mocked")
	return
}
// Get the after_apply_guidance field of the given pool_patch.
func (_class PoolPatchClass) GetAfterApplyGuidance(sessionID SessionRef, self PoolPatchRef) (_retval []AfterApplyGuidance, _err error) {
	if (IsMock) {
		return _class.GetAfterApplyGuidance__mock(sessionID, self)
	}	
	_method := "pool_patch.get_after_apply_guidance"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolPatchRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertEnumAfterApplyGuidanceSetToGo(_method + " -> ", _result.Value)
	return
}

func (_class PoolPatchClass) GetHostPatches__mock(sessionID SessionRef, self PoolPatchRef) (_retval []HostPatchRef, _err error) {
	log.Println("PoolPatch.GetHostPatches not mocked")
	_err = errors.New("PoolPatch.GetHostPatches not mocked")
	return
}
// Get the host_patches field of the given pool_patch.
func (_class PoolPatchClass) GetHostPatches(sessionID SessionRef, self PoolPatchRef) (_retval []HostPatchRef, _err error) {
	if (IsMock) {
		return _class.GetHostPatches__mock(sessionID, self)
	}	
	_method := "pool_patch.get_host_patches"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolPatchRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PoolPatchClass) GetPoolApplied__mock(sessionID SessionRef, self PoolPatchRef) (_retval bool, _err error) {
	log.Println("PoolPatch.GetPoolApplied not mocked")
	_err = errors.New("PoolPatch.GetPoolApplied not mocked")
	return
}
// Get the pool_applied field of the given pool_patch.
func (_class PoolPatchClass) GetPoolApplied(sessionID SessionRef, self PoolPatchRef) (_retval bool, _err error) {
	if (IsMock) {
		return _class.GetPoolApplied__mock(sessionID, self)
	}	
	_method := "pool_patch.get_pool_applied"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolPatchRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PoolPatchClass) GetSize__mock(sessionID SessionRef, self PoolPatchRef) (_retval int, _err error) {
	log.Println("PoolPatch.GetSize not mocked")
	_err = errors.New("PoolPatch.GetSize not mocked")
	return
}
// Get the size field of the given pool_patch.
func (_class PoolPatchClass) GetSize(sessionID SessionRef, self PoolPatchRef) (_retval int, _err error) {
	if (IsMock) {
		return _class.GetSize__mock(sessionID, self)
	}	
	_method := "pool_patch.get_size"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolPatchRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PoolPatchClass) GetVersion__mock(sessionID SessionRef, self PoolPatchRef) (_retval string, _err error) {
	log.Println("PoolPatch.GetVersion not mocked")
	_err = errors.New("PoolPatch.GetVersion not mocked")
	return
}
// Get the version field of the given pool_patch.
func (_class PoolPatchClass) GetVersion(sessionID SessionRef, self PoolPatchRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetVersion__mock(sessionID, self)
	}	
	_method := "pool_patch.get_version"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolPatchRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PoolPatchClass) GetNameDescription__mock(sessionID SessionRef, self PoolPatchRef) (_retval string, _err error) {
	log.Println("PoolPatch.GetNameDescription not mocked")
	_err = errors.New("PoolPatch.GetNameDescription not mocked")
	return
}
// Get the name/description field of the given pool_patch.
func (_class PoolPatchClass) GetNameDescription(sessionID SessionRef, self PoolPatchRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetNameDescription__mock(sessionID, self)
	}	
	_method := "pool_patch.get_name_description"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolPatchRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PoolPatchClass) GetNameLabel__mock(sessionID SessionRef, self PoolPatchRef) (_retval string, _err error) {
	log.Println("PoolPatch.GetNameLabel not mocked")
	_err = errors.New("PoolPatch.GetNameLabel not mocked")
	return
}
// Get the name/label field of the given pool_patch.
func (_class PoolPatchClass) GetNameLabel(sessionID SessionRef, self PoolPatchRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetNameLabel__mock(sessionID, self)
	}	
	_method := "pool_patch.get_name_label"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolPatchRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PoolPatchClass) GetUUID__mock(sessionID SessionRef, self PoolPatchRef) (_retval string, _err error) {
	log.Println("PoolPatch.GetUUID not mocked")
	_err = errors.New("PoolPatch.GetUUID not mocked")
	return
}
// Get the uuid field of the given pool_patch.
func (_class PoolPatchClass) GetUUID(sessionID SessionRef, self PoolPatchRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetUUID__mock(sessionID, self)
	}	
	_method := "pool_patch.get_uuid"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolPatchRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PoolPatchClass) GetByNameLabel__mock(sessionID SessionRef, label string) (_retval []PoolPatchRef, _err error) {
	log.Println("PoolPatch.GetByNameLabel not mocked")
	_err = errors.New("PoolPatch.GetByNameLabel not mocked")
	return
}
// Get all the pool_patch instances with the given label.
func (_class PoolPatchClass) GetByNameLabel(sessionID SessionRef, label string) (_retval []PoolPatchRef, _err error) {
	if (IsMock) {
		return _class.GetByNameLabel__mock(sessionID, label)
	}	
	_method := "pool_patch.get_by_name_label"
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
	_retval, _err = convertPoolPatchRefSetToGo(_method + " -> ", _result.Value)
	return
}

func (_class PoolPatchClass) GetByUUID__mock(sessionID SessionRef, uuid string) (_retval PoolPatchRef, _err error) {
	log.Println("PoolPatch.GetByUUID not mocked")
	_err = errors.New("PoolPatch.GetByUUID not mocked")
	return
}
// Get a reference to the pool_patch instance with the specified UUID.
func (_class PoolPatchClass) GetByUUID(sessionID SessionRef, uuid string) (_retval PoolPatchRef, _err error) {
	if (IsMock) {
		return _class.GetByUUID__mock(sessionID, uuid)
	}	
	_method := "pool_patch.get_by_uuid"
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
	_retval, _err = convertPoolPatchRefToGo(_method + " -> ", _result.Value)
	return
}

func (_class PoolPatchClass) GetRecord__mock(sessionID SessionRef, self PoolPatchRef) (_retval PoolPatchRecord, _err error) {
	log.Println("PoolPatch.GetRecord not mocked")
	_err = errors.New("PoolPatch.GetRecord not mocked")
	return
}
// Get a record containing the current state of the given pool_patch.
func (_class PoolPatchClass) GetRecord(sessionID SessionRef, self PoolPatchRef) (_retval PoolPatchRecord, _err error) {
	if (IsMock) {
		return _class.GetRecord__mock(sessionID, self)
	}	
	_method := "pool_patch.get_record"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolPatchRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertPoolPatchRecordToGo(_method + " -> ", _result.Value)
	return
}
