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


var PoolPatchClass_GetAllRecordsMockedCallback = func (sessionID SessionRef) (_retval map[PoolPatchRef]PoolPatchRecord, _err error) {
	log.Println("PoolPatch.GetAllRecords not mocked")
	_err = errors.New("PoolPatch.GetAllRecords not mocked")
	return
}

func (_class PoolPatchClass) GetAllRecordsMock(sessionID SessionRef) (_retval map[PoolPatchRef]PoolPatchRecord, _err error) {
	return PoolPatchClass_GetAllRecordsMockedCallback(sessionID)
}
// Return a map of pool_patch references to pool_patch records for all pool_patchs known to the system.
func (_class PoolPatchClass) GetAllRecords(sessionID SessionRef) (_retval map[PoolPatchRef]PoolPatchRecord, _err error) {
	if (IsMock) {
		return _class.GetAllRecordsMock(sessionID)
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


var PoolPatchClass_GetAllMockedCallback = func (sessionID SessionRef) (_retval []PoolPatchRef, _err error) {
	log.Println("PoolPatch.GetAll not mocked")
	_err = errors.New("PoolPatch.GetAll not mocked")
	return
}

func (_class PoolPatchClass) GetAllMock(sessionID SessionRef) (_retval []PoolPatchRef, _err error) {
	return PoolPatchClass_GetAllMockedCallback(sessionID)
}
// Return a list of all the pool_patchs known to the system.
func (_class PoolPatchClass) GetAll(sessionID SessionRef) (_retval []PoolPatchRef, _err error) {
	if (IsMock) {
		return _class.GetAllMock(sessionID)
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


var PoolPatchClass_CleanOnHostMockedCallback = func (sessionID SessionRef, self PoolPatchRef, host HostRef) (_err error) {
	log.Println("PoolPatch.CleanOnHost not mocked")
	_err = errors.New("PoolPatch.CleanOnHost not mocked")
	return
}

func (_class PoolPatchClass) CleanOnHostMock(sessionID SessionRef, self PoolPatchRef, host HostRef) (_err error) {
	return PoolPatchClass_CleanOnHostMockedCallback(sessionID, self, host)
}
// Removes the patch's files from the specified host
func (_class PoolPatchClass) CleanOnHost(sessionID SessionRef, self PoolPatchRef, host HostRef) (_err error) {
	if (IsMock) {
		return _class.CleanOnHostMock(sessionID, self, host)
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


var PoolPatchClass_DestroyMockedCallback = func (sessionID SessionRef, self PoolPatchRef) (_err error) {
	log.Println("PoolPatch.Destroy not mocked")
	_err = errors.New("PoolPatch.Destroy not mocked")
	return
}

func (_class PoolPatchClass) DestroyMock(sessionID SessionRef, self PoolPatchRef) (_err error) {
	return PoolPatchClass_DestroyMockedCallback(sessionID, self)
}
// Removes the patch's files from all hosts in the pool, and removes the database entries.  Only works on unapplied patches.
func (_class PoolPatchClass) Destroy(sessionID SessionRef, self PoolPatchRef) (_err error) {
	if (IsMock) {
		return _class.DestroyMock(sessionID, self)
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


var PoolPatchClass_PoolCleanMockedCallback = func (sessionID SessionRef, self PoolPatchRef) (_err error) {
	log.Println("PoolPatch.PoolClean not mocked")
	_err = errors.New("PoolPatch.PoolClean not mocked")
	return
}

func (_class PoolPatchClass) PoolCleanMock(sessionID SessionRef, self PoolPatchRef) (_err error) {
	return PoolPatchClass_PoolCleanMockedCallback(sessionID, self)
}
// Removes the patch's files from all hosts in the pool, but does not remove the database entries
func (_class PoolPatchClass) PoolClean(sessionID SessionRef, self PoolPatchRef) (_err error) {
	if (IsMock) {
		return _class.PoolCleanMock(sessionID, self)
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


var PoolPatchClass_CleanMockedCallback = func (sessionID SessionRef, self PoolPatchRef) (_err error) {
	log.Println("PoolPatch.Clean not mocked")
	_err = errors.New("PoolPatch.Clean not mocked")
	return
}

func (_class PoolPatchClass) CleanMock(sessionID SessionRef, self PoolPatchRef) (_err error) {
	return PoolPatchClass_CleanMockedCallback(sessionID, self)
}
// Removes the patch's files from the server
func (_class PoolPatchClass) Clean(sessionID SessionRef, self PoolPatchRef) (_err error) {
	if (IsMock) {
		return _class.CleanMock(sessionID, self)
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


var PoolPatchClass_PrecheckMockedCallback = func (sessionID SessionRef, self PoolPatchRef, host HostRef) (_retval string, _err error) {
	log.Println("PoolPatch.Precheck not mocked")
	_err = errors.New("PoolPatch.Precheck not mocked")
	return
}

func (_class PoolPatchClass) PrecheckMock(sessionID SessionRef, self PoolPatchRef, host HostRef) (_retval string, _err error) {
	return PoolPatchClass_PrecheckMockedCallback(sessionID, self, host)
}
// Execute the precheck stage of the selected patch on a host and return its output
func (_class PoolPatchClass) Precheck(sessionID SessionRef, self PoolPatchRef, host HostRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.PrecheckMock(sessionID, self, host)
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


var PoolPatchClass_PoolApplyMockedCallback = func (sessionID SessionRef, self PoolPatchRef) (_err error) {
	log.Println("PoolPatch.PoolApply not mocked")
	_err = errors.New("PoolPatch.PoolApply not mocked")
	return
}

func (_class PoolPatchClass) PoolApplyMock(sessionID SessionRef, self PoolPatchRef) (_err error) {
	return PoolPatchClass_PoolApplyMockedCallback(sessionID, self)
}
// Apply the selected patch to all hosts in the pool and return a map of host_ref -> patch output
func (_class PoolPatchClass) PoolApply(sessionID SessionRef, self PoolPatchRef) (_err error) {
	if (IsMock) {
		return _class.PoolApplyMock(sessionID, self)
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


var PoolPatchClass_ApplyMockedCallback = func (sessionID SessionRef, self PoolPatchRef, host HostRef) (_retval string, _err error) {
	log.Println("PoolPatch.Apply not mocked")
	_err = errors.New("PoolPatch.Apply not mocked")
	return
}

func (_class PoolPatchClass) ApplyMock(sessionID SessionRef, self PoolPatchRef, host HostRef) (_retval string, _err error) {
	return PoolPatchClass_ApplyMockedCallback(sessionID, self, host)
}
// Apply the selected patch to a host and return its output
func (_class PoolPatchClass) Apply(sessionID SessionRef, self PoolPatchRef, host HostRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.ApplyMock(sessionID, self, host)
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


var PoolPatchClass_RemoveFromOtherConfigMockedCallback = func (sessionID SessionRef, self PoolPatchRef, key string) (_err error) {
	log.Println("PoolPatch.RemoveFromOtherConfig not mocked")
	_err = errors.New("PoolPatch.RemoveFromOtherConfig not mocked")
	return
}

func (_class PoolPatchClass) RemoveFromOtherConfigMock(sessionID SessionRef, self PoolPatchRef, key string) (_err error) {
	return PoolPatchClass_RemoveFromOtherConfigMockedCallback(sessionID, self, key)
}
// Remove the given key and its corresponding value from the other_config field of the given pool_patch.  If the key is not in that Map, then do nothing.
func (_class PoolPatchClass) RemoveFromOtherConfig(sessionID SessionRef, self PoolPatchRef, key string) (_err error) {
	if (IsMock) {
		return _class.RemoveFromOtherConfigMock(sessionID, self, key)
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


var PoolPatchClass_AddToOtherConfigMockedCallback = func (sessionID SessionRef, self PoolPatchRef, key string, value string) (_err error) {
	log.Println("PoolPatch.AddToOtherConfig not mocked")
	_err = errors.New("PoolPatch.AddToOtherConfig not mocked")
	return
}

func (_class PoolPatchClass) AddToOtherConfigMock(sessionID SessionRef, self PoolPatchRef, key string, value string) (_err error) {
	return PoolPatchClass_AddToOtherConfigMockedCallback(sessionID, self, key, value)
}
// Add the given key-value pair to the other_config field of the given pool_patch.
func (_class PoolPatchClass) AddToOtherConfig(sessionID SessionRef, self PoolPatchRef, key string, value string) (_err error) {
	if (IsMock) {
		return _class.AddToOtherConfigMock(sessionID, self, key, value)
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


var PoolPatchClass_SetOtherConfigMockedCallback = func (sessionID SessionRef, self PoolPatchRef, value map[string]string) (_err error) {
	log.Println("PoolPatch.SetOtherConfig not mocked")
	_err = errors.New("PoolPatch.SetOtherConfig not mocked")
	return
}

func (_class PoolPatchClass) SetOtherConfigMock(sessionID SessionRef, self PoolPatchRef, value map[string]string) (_err error) {
	return PoolPatchClass_SetOtherConfigMockedCallback(sessionID, self, value)
}
// Set the other_config field of the given pool_patch.
func (_class PoolPatchClass) SetOtherConfig(sessionID SessionRef, self PoolPatchRef, value map[string]string) (_err error) {
	if (IsMock) {
		return _class.SetOtherConfigMock(sessionID, self, value)
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


var PoolPatchClass_GetOtherConfigMockedCallback = func (sessionID SessionRef, self PoolPatchRef) (_retval map[string]string, _err error) {
	log.Println("PoolPatch.GetOtherConfig not mocked")
	_err = errors.New("PoolPatch.GetOtherConfig not mocked")
	return
}

func (_class PoolPatchClass) GetOtherConfigMock(sessionID SessionRef, self PoolPatchRef) (_retval map[string]string, _err error) {
	return PoolPatchClass_GetOtherConfigMockedCallback(sessionID, self)
}
// Get the other_config field of the given pool_patch.
func (_class PoolPatchClass) GetOtherConfig(sessionID SessionRef, self PoolPatchRef) (_retval map[string]string, _err error) {
	if (IsMock) {
		return _class.GetOtherConfigMock(sessionID, self)
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


var PoolPatchClass_GetPoolUpdateMockedCallback = func (sessionID SessionRef, self PoolPatchRef) (_retval PoolUpdateRef, _err error) {
	log.Println("PoolPatch.GetPoolUpdate not mocked")
	_err = errors.New("PoolPatch.GetPoolUpdate not mocked")
	return
}

func (_class PoolPatchClass) GetPoolUpdateMock(sessionID SessionRef, self PoolPatchRef) (_retval PoolUpdateRef, _err error) {
	return PoolPatchClass_GetPoolUpdateMockedCallback(sessionID, self)
}
// Get the pool_update field of the given pool_patch.
func (_class PoolPatchClass) GetPoolUpdate(sessionID SessionRef, self PoolPatchRef) (_retval PoolUpdateRef, _err error) {
	if (IsMock) {
		return _class.GetPoolUpdateMock(sessionID, self)
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


var PoolPatchClass_GetAfterApplyGuidanceMockedCallback = func (sessionID SessionRef, self PoolPatchRef) (_retval []AfterApplyGuidance, _err error) {
	log.Println("PoolPatch.GetAfterApplyGuidance not mocked")
	_err = errors.New("PoolPatch.GetAfterApplyGuidance not mocked")
	return
}

func (_class PoolPatchClass) GetAfterApplyGuidanceMock(sessionID SessionRef, self PoolPatchRef) (_retval []AfterApplyGuidance, _err error) {
	return PoolPatchClass_GetAfterApplyGuidanceMockedCallback(sessionID, self)
}
// Get the after_apply_guidance field of the given pool_patch.
func (_class PoolPatchClass) GetAfterApplyGuidance(sessionID SessionRef, self PoolPatchRef) (_retval []AfterApplyGuidance, _err error) {
	if (IsMock) {
		return _class.GetAfterApplyGuidanceMock(sessionID, self)
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


var PoolPatchClass_GetHostPatchesMockedCallback = func (sessionID SessionRef, self PoolPatchRef) (_retval []HostPatchRef, _err error) {
	log.Println("PoolPatch.GetHostPatches not mocked")
	_err = errors.New("PoolPatch.GetHostPatches not mocked")
	return
}

func (_class PoolPatchClass) GetHostPatchesMock(sessionID SessionRef, self PoolPatchRef) (_retval []HostPatchRef, _err error) {
	return PoolPatchClass_GetHostPatchesMockedCallback(sessionID, self)
}
// Get the host_patches field of the given pool_patch.
func (_class PoolPatchClass) GetHostPatches(sessionID SessionRef, self PoolPatchRef) (_retval []HostPatchRef, _err error) {
	if (IsMock) {
		return _class.GetHostPatchesMock(sessionID, self)
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


var PoolPatchClass_GetPoolAppliedMockedCallback = func (sessionID SessionRef, self PoolPatchRef) (_retval bool, _err error) {
	log.Println("PoolPatch.GetPoolApplied not mocked")
	_err = errors.New("PoolPatch.GetPoolApplied not mocked")
	return
}

func (_class PoolPatchClass) GetPoolAppliedMock(sessionID SessionRef, self PoolPatchRef) (_retval bool, _err error) {
	return PoolPatchClass_GetPoolAppliedMockedCallback(sessionID, self)
}
// Get the pool_applied field of the given pool_patch.
func (_class PoolPatchClass) GetPoolApplied(sessionID SessionRef, self PoolPatchRef) (_retval bool, _err error) {
	if (IsMock) {
		return _class.GetPoolAppliedMock(sessionID, self)
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


var PoolPatchClass_GetSizeMockedCallback = func (sessionID SessionRef, self PoolPatchRef) (_retval int, _err error) {
	log.Println("PoolPatch.GetSize not mocked")
	_err = errors.New("PoolPatch.GetSize not mocked")
	return
}

func (_class PoolPatchClass) GetSizeMock(sessionID SessionRef, self PoolPatchRef) (_retval int, _err error) {
	return PoolPatchClass_GetSizeMockedCallback(sessionID, self)
}
// Get the size field of the given pool_patch.
func (_class PoolPatchClass) GetSize(sessionID SessionRef, self PoolPatchRef) (_retval int, _err error) {
	if (IsMock) {
		return _class.GetSizeMock(sessionID, self)
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


var PoolPatchClass_GetVersionMockedCallback = func (sessionID SessionRef, self PoolPatchRef) (_retval string, _err error) {
	log.Println("PoolPatch.GetVersion not mocked")
	_err = errors.New("PoolPatch.GetVersion not mocked")
	return
}

func (_class PoolPatchClass) GetVersionMock(sessionID SessionRef, self PoolPatchRef) (_retval string, _err error) {
	return PoolPatchClass_GetVersionMockedCallback(sessionID, self)
}
// Get the version field of the given pool_patch.
func (_class PoolPatchClass) GetVersion(sessionID SessionRef, self PoolPatchRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetVersionMock(sessionID, self)
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


var PoolPatchClass_GetNameDescriptionMockedCallback = func (sessionID SessionRef, self PoolPatchRef) (_retval string, _err error) {
	log.Println("PoolPatch.GetNameDescription not mocked")
	_err = errors.New("PoolPatch.GetNameDescription not mocked")
	return
}

func (_class PoolPatchClass) GetNameDescriptionMock(sessionID SessionRef, self PoolPatchRef) (_retval string, _err error) {
	return PoolPatchClass_GetNameDescriptionMockedCallback(sessionID, self)
}
// Get the name/description field of the given pool_patch.
func (_class PoolPatchClass) GetNameDescription(sessionID SessionRef, self PoolPatchRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetNameDescriptionMock(sessionID, self)
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


var PoolPatchClass_GetNameLabelMockedCallback = func (sessionID SessionRef, self PoolPatchRef) (_retval string, _err error) {
	log.Println("PoolPatch.GetNameLabel not mocked")
	_err = errors.New("PoolPatch.GetNameLabel not mocked")
	return
}

func (_class PoolPatchClass) GetNameLabelMock(sessionID SessionRef, self PoolPatchRef) (_retval string, _err error) {
	return PoolPatchClass_GetNameLabelMockedCallback(sessionID, self)
}
// Get the name/label field of the given pool_patch.
func (_class PoolPatchClass) GetNameLabel(sessionID SessionRef, self PoolPatchRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetNameLabelMock(sessionID, self)
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


var PoolPatchClass_GetUUIDMockedCallback = func (sessionID SessionRef, self PoolPatchRef) (_retval string, _err error) {
	log.Println("PoolPatch.GetUUID not mocked")
	_err = errors.New("PoolPatch.GetUUID not mocked")
	return
}

func (_class PoolPatchClass) GetUUIDMock(sessionID SessionRef, self PoolPatchRef) (_retval string, _err error) {
	return PoolPatchClass_GetUUIDMockedCallback(sessionID, self)
}
// Get the uuid field of the given pool_patch.
func (_class PoolPatchClass) GetUUID(sessionID SessionRef, self PoolPatchRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetUUIDMock(sessionID, self)
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


var PoolPatchClass_GetByNameLabelMockedCallback = func (sessionID SessionRef, label string) (_retval []PoolPatchRef, _err error) {
	log.Println("PoolPatch.GetByNameLabel not mocked")
	_err = errors.New("PoolPatch.GetByNameLabel not mocked")
	return
}

func (_class PoolPatchClass) GetByNameLabelMock(sessionID SessionRef, label string) (_retval []PoolPatchRef, _err error) {
	return PoolPatchClass_GetByNameLabelMockedCallback(sessionID, label)
}
// Get all the pool_patch instances with the given label.
func (_class PoolPatchClass) GetByNameLabel(sessionID SessionRef, label string) (_retval []PoolPatchRef, _err error) {
	if (IsMock) {
		return _class.GetByNameLabelMock(sessionID, label)
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


var PoolPatchClass_GetByUUIDMockedCallback = func (sessionID SessionRef, uuid string) (_retval PoolPatchRef, _err error) {
	log.Println("PoolPatch.GetByUUID not mocked")
	_err = errors.New("PoolPatch.GetByUUID not mocked")
	return
}

func (_class PoolPatchClass) GetByUUIDMock(sessionID SessionRef, uuid string) (_retval PoolPatchRef, _err error) {
	return PoolPatchClass_GetByUUIDMockedCallback(sessionID, uuid)
}
// Get a reference to the pool_patch instance with the specified UUID.
func (_class PoolPatchClass) GetByUUID(sessionID SessionRef, uuid string) (_retval PoolPatchRef, _err error) {
	if (IsMock) {
		return _class.GetByUUIDMock(sessionID, uuid)
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


var PoolPatchClass_GetRecordMockedCallback = func (sessionID SessionRef, self PoolPatchRef) (_retval PoolPatchRecord, _err error) {
	log.Println("PoolPatch.GetRecord not mocked")
	_err = errors.New("PoolPatch.GetRecord not mocked")
	return
}

func (_class PoolPatchClass) GetRecordMock(sessionID SessionRef, self PoolPatchRef) (_retval PoolPatchRecord, _err error) {
	return PoolPatchClass_GetRecordMockedCallback(sessionID, self)
}
// Get a record containing the current state of the given pool_patch.
func (_class PoolPatchClass) GetRecord(sessionID SessionRef, self PoolPatchRef) (_retval PoolPatchRecord, _err error) {
	if (IsMock) {
		return _class.GetRecordMock(sessionID, self)
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
