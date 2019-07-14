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

type HostPatchRecord struct {
  // Unique identifier/object reference
	UUID string
  // a human-readable name
	NameLabel string
  // a notes field containing human-readable description
	NameDescription string
  // Patch version number
	Version string
  // Host the patch relates to
	Host HostRef
  // True if the patch has been applied
	Applied bool
  // Time the patch was applied
	TimestampApplied time.Time
  // Size of the patch
	Size int
  // The patch applied
	PoolPatch PoolPatchRef
  // additional configuration
	OtherConfig map[string]string
}

type HostPatchRef string

// Represents a patch stored on a server
type HostPatchClass struct {
	client *Client
}


var HostPatchClass_GetAllRecordsMockedCallback = func (sessionID SessionRef) (_retval map[HostPatchRef]HostPatchRecord, _err error) {
	log.Println("HostPatch.GetAllRecords not mocked")
	_err = errors.New("HostPatch.GetAllRecords not mocked")
	return
}

func (_class HostPatchClass) GetAllRecordsMock(sessionID SessionRef) (_retval map[HostPatchRef]HostPatchRecord, _err error) {
	return HostPatchClass_GetAllRecordsMockedCallback(sessionID)
}
// Return a map of host_patch references to host_patch records for all host_patchs known to the system.
func (_class HostPatchClass) GetAllRecords(sessionID SessionRef) (_retval map[HostPatchRef]HostPatchRecord, _err error) {
	if (IsMock) {
		return _class.GetAllRecordsMock(sessionID)
	}	
	_method := "host_patch.get_all_records"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertHostPatchRefToHostPatchRecordMapToGo(_method + " -> ", _result.Value)
	return
}


var HostPatchClass_GetAllMockedCallback = func (sessionID SessionRef) (_retval []HostPatchRef, _err error) {
	log.Println("HostPatch.GetAll not mocked")
	_err = errors.New("HostPatch.GetAll not mocked")
	return
}

func (_class HostPatchClass) GetAllMock(sessionID SessionRef) (_retval []HostPatchRef, _err error) {
	return HostPatchClass_GetAllMockedCallback(sessionID)
}
// Return a list of all the host_patchs known to the system.
func (_class HostPatchClass) GetAll(sessionID SessionRef) (_retval []HostPatchRef, _err error) {
	if (IsMock) {
		return _class.GetAllMock(sessionID)
	}	
	_method := "host_patch.get_all"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertHostPatchRefSetToGo(_method + " -> ", _result.Value)
	return
}


var HostPatchClass_ApplyMockedCallback = func (sessionID SessionRef, self HostPatchRef) (_retval string, _err error) {
	log.Println("HostPatch.Apply not mocked")
	_err = errors.New("HostPatch.Apply not mocked")
	return
}

func (_class HostPatchClass) ApplyMock(sessionID SessionRef, self HostPatchRef) (_retval string, _err error) {
	return HostPatchClass_ApplyMockedCallback(sessionID, self)
}
// Apply the selected patch and return its output
func (_class HostPatchClass) Apply(sessionID SessionRef, self HostPatchRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.ApplyMock(sessionID, self)
	}	
	_method := "host_patch.apply"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertHostPatchRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var HostPatchClass_DestroyMockedCallback = func (sessionID SessionRef, self HostPatchRef) (_err error) {
	log.Println("HostPatch.Destroy not mocked")
	_err = errors.New("HostPatch.Destroy not mocked")
	return
}

func (_class HostPatchClass) DestroyMock(sessionID SessionRef, self HostPatchRef) (_err error) {
	return HostPatchClass_DestroyMockedCallback(sessionID, self)
}
// Destroy the specified host patch, removing it from the disk. This does NOT reverse the patch
func (_class HostPatchClass) Destroy(sessionID SessionRef, self HostPatchRef) (_err error) {
	if (IsMock) {
		return _class.DestroyMock(sessionID, self)
	}	
	_method := "host_patch.destroy"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertHostPatchRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg)
	return
}


var HostPatchClass_RemoveFromOtherConfigMockedCallback = func (sessionID SessionRef, self HostPatchRef, key string) (_err error) {
	log.Println("HostPatch.RemoveFromOtherConfig not mocked")
	_err = errors.New("HostPatch.RemoveFromOtherConfig not mocked")
	return
}

func (_class HostPatchClass) RemoveFromOtherConfigMock(sessionID SessionRef, self HostPatchRef, key string) (_err error) {
	return HostPatchClass_RemoveFromOtherConfigMockedCallback(sessionID, self, key)
}
// Remove the given key and its corresponding value from the other_config field of the given host_patch.  If the key is not in that Map, then do nothing.
func (_class HostPatchClass) RemoveFromOtherConfig(sessionID SessionRef, self HostPatchRef, key string) (_err error) {
	if (IsMock) {
		return _class.RemoveFromOtherConfigMock(sessionID, self, key)
	}	
	_method := "host_patch.remove_from_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertHostPatchRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var HostPatchClass_AddToOtherConfigMockedCallback = func (sessionID SessionRef, self HostPatchRef, key string, value string) (_err error) {
	log.Println("HostPatch.AddToOtherConfig not mocked")
	_err = errors.New("HostPatch.AddToOtherConfig not mocked")
	return
}

func (_class HostPatchClass) AddToOtherConfigMock(sessionID SessionRef, self HostPatchRef, key string, value string) (_err error) {
	return HostPatchClass_AddToOtherConfigMockedCallback(sessionID, self, key, value)
}
// Add the given key-value pair to the other_config field of the given host_patch.
func (_class HostPatchClass) AddToOtherConfig(sessionID SessionRef, self HostPatchRef, key string, value string) (_err error) {
	if (IsMock) {
		return _class.AddToOtherConfigMock(sessionID, self, key, value)
	}	
	_method := "host_patch.add_to_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertHostPatchRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var HostPatchClass_SetOtherConfigMockedCallback = func (sessionID SessionRef, self HostPatchRef, value map[string]string) (_err error) {
	log.Println("HostPatch.SetOtherConfig not mocked")
	_err = errors.New("HostPatch.SetOtherConfig not mocked")
	return
}

func (_class HostPatchClass) SetOtherConfigMock(sessionID SessionRef, self HostPatchRef, value map[string]string) (_err error) {
	return HostPatchClass_SetOtherConfigMockedCallback(sessionID, self, value)
}
// Set the other_config field of the given host_patch.
func (_class HostPatchClass) SetOtherConfig(sessionID SessionRef, self HostPatchRef, value map[string]string) (_err error) {
	if (IsMock) {
		return _class.SetOtherConfigMock(sessionID, self, value)
	}	
	_method := "host_patch.set_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertHostPatchRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var HostPatchClass_GetOtherConfigMockedCallback = func (sessionID SessionRef, self HostPatchRef) (_retval map[string]string, _err error) {
	log.Println("HostPatch.GetOtherConfig not mocked")
	_err = errors.New("HostPatch.GetOtherConfig not mocked")
	return
}

func (_class HostPatchClass) GetOtherConfigMock(sessionID SessionRef, self HostPatchRef) (_retval map[string]string, _err error) {
	return HostPatchClass_GetOtherConfigMockedCallback(sessionID, self)
}
// Get the other_config field of the given host_patch.
func (_class HostPatchClass) GetOtherConfig(sessionID SessionRef, self HostPatchRef) (_retval map[string]string, _err error) {
	if (IsMock) {
		return _class.GetOtherConfigMock(sessionID, self)
	}	
	_method := "host_patch.get_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertHostPatchRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var HostPatchClass_GetPoolPatchMockedCallback = func (sessionID SessionRef, self HostPatchRef) (_retval PoolPatchRef, _err error) {
	log.Println("HostPatch.GetPoolPatch not mocked")
	_err = errors.New("HostPatch.GetPoolPatch not mocked")
	return
}

func (_class HostPatchClass) GetPoolPatchMock(sessionID SessionRef, self HostPatchRef) (_retval PoolPatchRef, _err error) {
	return HostPatchClass_GetPoolPatchMockedCallback(sessionID, self)
}
// Get the pool_patch field of the given host_patch.
func (_class HostPatchClass) GetPoolPatch(sessionID SessionRef, self HostPatchRef) (_retval PoolPatchRef, _err error) {
	if (IsMock) {
		return _class.GetPoolPatchMock(sessionID, self)
	}	
	_method := "host_patch.get_pool_patch"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertHostPatchRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertPoolPatchRefToGo(_method + " -> ", _result.Value)
	return
}


var HostPatchClass_GetSizeMockedCallback = func (sessionID SessionRef, self HostPatchRef) (_retval int, _err error) {
	log.Println("HostPatch.GetSize not mocked")
	_err = errors.New("HostPatch.GetSize not mocked")
	return
}

func (_class HostPatchClass) GetSizeMock(sessionID SessionRef, self HostPatchRef) (_retval int, _err error) {
	return HostPatchClass_GetSizeMockedCallback(sessionID, self)
}
// Get the size field of the given host_patch.
func (_class HostPatchClass) GetSize(sessionID SessionRef, self HostPatchRef) (_retval int, _err error) {
	if (IsMock) {
		return _class.GetSizeMock(sessionID, self)
	}	
	_method := "host_patch.get_size"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertHostPatchRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var HostPatchClass_GetTimestampAppliedMockedCallback = func (sessionID SessionRef, self HostPatchRef) (_retval time.Time, _err error) {
	log.Println("HostPatch.GetTimestampApplied not mocked")
	_err = errors.New("HostPatch.GetTimestampApplied not mocked")
	return
}

func (_class HostPatchClass) GetTimestampAppliedMock(sessionID SessionRef, self HostPatchRef) (_retval time.Time, _err error) {
	return HostPatchClass_GetTimestampAppliedMockedCallback(sessionID, self)
}
// Get the timestamp_applied field of the given host_patch.
func (_class HostPatchClass) GetTimestampApplied(sessionID SessionRef, self HostPatchRef) (_retval time.Time, _err error) {
	if (IsMock) {
		return _class.GetTimestampAppliedMock(sessionID, self)
	}	
	_method := "host_patch.get_timestamp_applied"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertHostPatchRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var HostPatchClass_GetAppliedMockedCallback = func (sessionID SessionRef, self HostPatchRef) (_retval bool, _err error) {
	log.Println("HostPatch.GetApplied not mocked")
	_err = errors.New("HostPatch.GetApplied not mocked")
	return
}

func (_class HostPatchClass) GetAppliedMock(sessionID SessionRef, self HostPatchRef) (_retval bool, _err error) {
	return HostPatchClass_GetAppliedMockedCallback(sessionID, self)
}
// Get the applied field of the given host_patch.
func (_class HostPatchClass) GetApplied(sessionID SessionRef, self HostPatchRef) (_retval bool, _err error) {
	if (IsMock) {
		return _class.GetAppliedMock(sessionID, self)
	}	
	_method := "host_patch.get_applied"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertHostPatchRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var HostPatchClass_GetHostMockedCallback = func (sessionID SessionRef, self HostPatchRef) (_retval HostRef, _err error) {
	log.Println("HostPatch.GetHost not mocked")
	_err = errors.New("HostPatch.GetHost not mocked")
	return
}

func (_class HostPatchClass) GetHostMock(sessionID SessionRef, self HostPatchRef) (_retval HostRef, _err error) {
	return HostPatchClass_GetHostMockedCallback(sessionID, self)
}
// Get the host field of the given host_patch.
func (_class HostPatchClass) GetHost(sessionID SessionRef, self HostPatchRef) (_retval HostRef, _err error) {
	if (IsMock) {
		return _class.GetHostMock(sessionID, self)
	}	
	_method := "host_patch.get_host"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertHostPatchRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var HostPatchClass_GetVersionMockedCallback = func (sessionID SessionRef, self HostPatchRef) (_retval string, _err error) {
	log.Println("HostPatch.GetVersion not mocked")
	_err = errors.New("HostPatch.GetVersion not mocked")
	return
}

func (_class HostPatchClass) GetVersionMock(sessionID SessionRef, self HostPatchRef) (_retval string, _err error) {
	return HostPatchClass_GetVersionMockedCallback(sessionID, self)
}
// Get the version field of the given host_patch.
func (_class HostPatchClass) GetVersion(sessionID SessionRef, self HostPatchRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetVersionMock(sessionID, self)
	}	
	_method := "host_patch.get_version"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertHostPatchRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var HostPatchClass_GetNameDescriptionMockedCallback = func (sessionID SessionRef, self HostPatchRef) (_retval string, _err error) {
	log.Println("HostPatch.GetNameDescription not mocked")
	_err = errors.New("HostPatch.GetNameDescription not mocked")
	return
}

func (_class HostPatchClass) GetNameDescriptionMock(sessionID SessionRef, self HostPatchRef) (_retval string, _err error) {
	return HostPatchClass_GetNameDescriptionMockedCallback(sessionID, self)
}
// Get the name/description field of the given host_patch.
func (_class HostPatchClass) GetNameDescription(sessionID SessionRef, self HostPatchRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetNameDescriptionMock(sessionID, self)
	}	
	_method := "host_patch.get_name_description"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertHostPatchRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var HostPatchClass_GetNameLabelMockedCallback = func (sessionID SessionRef, self HostPatchRef) (_retval string, _err error) {
	log.Println("HostPatch.GetNameLabel not mocked")
	_err = errors.New("HostPatch.GetNameLabel not mocked")
	return
}

func (_class HostPatchClass) GetNameLabelMock(sessionID SessionRef, self HostPatchRef) (_retval string, _err error) {
	return HostPatchClass_GetNameLabelMockedCallback(sessionID, self)
}
// Get the name/label field of the given host_patch.
func (_class HostPatchClass) GetNameLabel(sessionID SessionRef, self HostPatchRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetNameLabelMock(sessionID, self)
	}	
	_method := "host_patch.get_name_label"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertHostPatchRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var HostPatchClass_GetUUIDMockedCallback = func (sessionID SessionRef, self HostPatchRef) (_retval string, _err error) {
	log.Println("HostPatch.GetUUID not mocked")
	_err = errors.New("HostPatch.GetUUID not mocked")
	return
}

func (_class HostPatchClass) GetUUIDMock(sessionID SessionRef, self HostPatchRef) (_retval string, _err error) {
	return HostPatchClass_GetUUIDMockedCallback(sessionID, self)
}
// Get the uuid field of the given host_patch.
func (_class HostPatchClass) GetUUID(sessionID SessionRef, self HostPatchRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetUUIDMock(sessionID, self)
	}	
	_method := "host_patch.get_uuid"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertHostPatchRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var HostPatchClass_GetByNameLabelMockedCallback = func (sessionID SessionRef, label string) (_retval []HostPatchRef, _err error) {
	log.Println("HostPatch.GetByNameLabel not mocked")
	_err = errors.New("HostPatch.GetByNameLabel not mocked")
	return
}

func (_class HostPatchClass) GetByNameLabelMock(sessionID SessionRef, label string) (_retval []HostPatchRef, _err error) {
	return HostPatchClass_GetByNameLabelMockedCallback(sessionID, label)
}
// Get all the host_patch instances with the given label.
func (_class HostPatchClass) GetByNameLabel(sessionID SessionRef, label string) (_retval []HostPatchRef, _err error) {
	if (IsMock) {
		return _class.GetByNameLabelMock(sessionID, label)
	}	
	_method := "host_patch.get_by_name_label"
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
	_retval, _err = convertHostPatchRefSetToGo(_method + " -> ", _result.Value)
	return
}


var HostPatchClass_GetByUUIDMockedCallback = func (sessionID SessionRef, uuid string) (_retval HostPatchRef, _err error) {
	log.Println("HostPatch.GetByUUID not mocked")
	_err = errors.New("HostPatch.GetByUUID not mocked")
	return
}

func (_class HostPatchClass) GetByUUIDMock(sessionID SessionRef, uuid string) (_retval HostPatchRef, _err error) {
	return HostPatchClass_GetByUUIDMockedCallback(sessionID, uuid)
}
// Get a reference to the host_patch instance with the specified UUID.
func (_class HostPatchClass) GetByUUID(sessionID SessionRef, uuid string) (_retval HostPatchRef, _err error) {
	if (IsMock) {
		return _class.GetByUUIDMock(sessionID, uuid)
	}	
	_method := "host_patch.get_by_uuid"
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
	_retval, _err = convertHostPatchRefToGo(_method + " -> ", _result.Value)
	return
}


var HostPatchClass_GetRecordMockedCallback = func (sessionID SessionRef, self HostPatchRef) (_retval HostPatchRecord, _err error) {
	log.Println("HostPatch.GetRecord not mocked")
	_err = errors.New("HostPatch.GetRecord not mocked")
	return
}

func (_class HostPatchClass) GetRecordMock(sessionID SessionRef, self HostPatchRef) (_retval HostPatchRecord, _err error) {
	return HostPatchClass_GetRecordMockedCallback(sessionID, self)
}
// Get a record containing the current state of the given host_patch.
func (_class HostPatchClass) GetRecord(sessionID SessionRef, self HostPatchRef) (_retval HostPatchRecord, _err error) {
	if (IsMock) {
		return _class.GetRecordMock(sessionID, self)
	}	
	_method := "host_patch.get_record"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertHostPatchRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertHostPatchRecordToGo(_method + " -> ", _result.Value)
	return
}
