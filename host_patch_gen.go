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

func (_class HostPatchClass) GetAllRecords__mock(sessionID SessionRef) (_retval map[HostPatchRef]HostPatchRecord, _err error) {
	log.Println("HostPatch.GetAllRecords not mocked")
	_err = errors.New("HostPatch.GetAllRecords not mocked")
	return
}
// Return a map of host_patch references to host_patch records for all host_patchs known to the system.
func (_class HostPatchClass) GetAllRecords(sessionID SessionRef) (_retval map[HostPatchRef]HostPatchRecord, _err error) {
	if (IsMock) {
		return _class.GetAllRecords__mock(sessionID)
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

func (_class HostPatchClass) GetAll__mock(sessionID SessionRef) (_retval []HostPatchRef, _err error) {
	log.Println("HostPatch.GetAll not mocked")
	_err = errors.New("HostPatch.GetAll not mocked")
	return
}
// Return a list of all the host_patchs known to the system.
func (_class HostPatchClass) GetAll(sessionID SessionRef) (_retval []HostPatchRef, _err error) {
	if (IsMock) {
		return _class.GetAll__mock(sessionID)
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

func (_class HostPatchClass) Apply__mock(sessionID SessionRef, self HostPatchRef) (_retval string, _err error) {
	log.Println("HostPatch.Apply not mocked")
	_err = errors.New("HostPatch.Apply not mocked")
	return
}
// Apply the selected patch and return its output
func (_class HostPatchClass) Apply(sessionID SessionRef, self HostPatchRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.Apply__mock(sessionID, self)
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

func (_class HostPatchClass) Destroy__mock(sessionID SessionRef, self HostPatchRef) (_err error) {
	log.Println("HostPatch.Destroy not mocked")
	_err = errors.New("HostPatch.Destroy not mocked")
	return
}
// Destroy the specified host patch, removing it from the disk. This does NOT reverse the patch
func (_class HostPatchClass) Destroy(sessionID SessionRef, self HostPatchRef) (_err error) {
	if (IsMock) {
		return _class.Destroy__mock(sessionID, self)
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

func (_class HostPatchClass) RemoveFromOtherConfig__mock(sessionID SessionRef, self HostPatchRef, key string) (_err error) {
	log.Println("HostPatch.RemoveFromOtherConfig not mocked")
	_err = errors.New("HostPatch.RemoveFromOtherConfig not mocked")
	return
}
// Remove the given key and its corresponding value from the other_config field of the given host_patch.  If the key is not in that Map, then do nothing.
func (_class HostPatchClass) RemoveFromOtherConfig(sessionID SessionRef, self HostPatchRef, key string) (_err error) {
	if (IsMock) {
		return _class.RemoveFromOtherConfig__mock(sessionID, self, key)
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

func (_class HostPatchClass) AddToOtherConfig__mock(sessionID SessionRef, self HostPatchRef, key string, value string) (_err error) {
	log.Println("HostPatch.AddToOtherConfig not mocked")
	_err = errors.New("HostPatch.AddToOtherConfig not mocked")
	return
}
// Add the given key-value pair to the other_config field of the given host_patch.
func (_class HostPatchClass) AddToOtherConfig(sessionID SessionRef, self HostPatchRef, key string, value string) (_err error) {
	if (IsMock) {
		return _class.AddToOtherConfig__mock(sessionID, self, key, value)
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

func (_class HostPatchClass) SetOtherConfig__mock(sessionID SessionRef, self HostPatchRef, value map[string]string) (_err error) {
	log.Println("HostPatch.SetOtherConfig not mocked")
	_err = errors.New("HostPatch.SetOtherConfig not mocked")
	return
}
// Set the other_config field of the given host_patch.
func (_class HostPatchClass) SetOtherConfig(sessionID SessionRef, self HostPatchRef, value map[string]string) (_err error) {
	if (IsMock) {
		return _class.SetOtherConfig__mock(sessionID, self, value)
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

func (_class HostPatchClass) GetOtherConfig__mock(sessionID SessionRef, self HostPatchRef) (_retval map[string]string, _err error) {
	log.Println("HostPatch.GetOtherConfig not mocked")
	_err = errors.New("HostPatch.GetOtherConfig not mocked")
	return
}
// Get the other_config field of the given host_patch.
func (_class HostPatchClass) GetOtherConfig(sessionID SessionRef, self HostPatchRef) (_retval map[string]string, _err error) {
	if (IsMock) {
		return _class.GetOtherConfig__mock(sessionID, self)
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

func (_class HostPatchClass) GetPoolPatch__mock(sessionID SessionRef, self HostPatchRef) (_retval PoolPatchRef, _err error) {
	log.Println("HostPatch.GetPoolPatch not mocked")
	_err = errors.New("HostPatch.GetPoolPatch not mocked")
	return
}
// Get the pool_patch field of the given host_patch.
func (_class HostPatchClass) GetPoolPatch(sessionID SessionRef, self HostPatchRef) (_retval PoolPatchRef, _err error) {
	if (IsMock) {
		return _class.GetPoolPatch__mock(sessionID, self)
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

func (_class HostPatchClass) GetSize__mock(sessionID SessionRef, self HostPatchRef) (_retval int, _err error) {
	log.Println("HostPatch.GetSize not mocked")
	_err = errors.New("HostPatch.GetSize not mocked")
	return
}
// Get the size field of the given host_patch.
func (_class HostPatchClass) GetSize(sessionID SessionRef, self HostPatchRef) (_retval int, _err error) {
	if (IsMock) {
		return _class.GetSize__mock(sessionID, self)
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

func (_class HostPatchClass) GetTimestampApplied__mock(sessionID SessionRef, self HostPatchRef) (_retval time.Time, _err error) {
	log.Println("HostPatch.GetTimestampApplied not mocked")
	_err = errors.New("HostPatch.GetTimestampApplied not mocked")
	return
}
// Get the timestamp_applied field of the given host_patch.
func (_class HostPatchClass) GetTimestampApplied(sessionID SessionRef, self HostPatchRef) (_retval time.Time, _err error) {
	if (IsMock) {
		return _class.GetTimestampApplied__mock(sessionID, self)
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

func (_class HostPatchClass) GetApplied__mock(sessionID SessionRef, self HostPatchRef) (_retval bool, _err error) {
	log.Println("HostPatch.GetApplied not mocked")
	_err = errors.New("HostPatch.GetApplied not mocked")
	return
}
// Get the applied field of the given host_patch.
func (_class HostPatchClass) GetApplied(sessionID SessionRef, self HostPatchRef) (_retval bool, _err error) {
	if (IsMock) {
		return _class.GetApplied__mock(sessionID, self)
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

func (_class HostPatchClass) GetHost__mock(sessionID SessionRef, self HostPatchRef) (_retval HostRef, _err error) {
	log.Println("HostPatch.GetHost not mocked")
	_err = errors.New("HostPatch.GetHost not mocked")
	return
}
// Get the host field of the given host_patch.
func (_class HostPatchClass) GetHost(sessionID SessionRef, self HostPatchRef) (_retval HostRef, _err error) {
	if (IsMock) {
		return _class.GetHost__mock(sessionID, self)
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

func (_class HostPatchClass) GetVersion__mock(sessionID SessionRef, self HostPatchRef) (_retval string, _err error) {
	log.Println("HostPatch.GetVersion not mocked")
	_err = errors.New("HostPatch.GetVersion not mocked")
	return
}
// Get the version field of the given host_patch.
func (_class HostPatchClass) GetVersion(sessionID SessionRef, self HostPatchRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetVersion__mock(sessionID, self)
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

func (_class HostPatchClass) GetNameDescription__mock(sessionID SessionRef, self HostPatchRef) (_retval string, _err error) {
	log.Println("HostPatch.GetNameDescription not mocked")
	_err = errors.New("HostPatch.GetNameDescription not mocked")
	return
}
// Get the name/description field of the given host_patch.
func (_class HostPatchClass) GetNameDescription(sessionID SessionRef, self HostPatchRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetNameDescription__mock(sessionID, self)
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

func (_class HostPatchClass) GetNameLabel__mock(sessionID SessionRef, self HostPatchRef) (_retval string, _err error) {
	log.Println("HostPatch.GetNameLabel not mocked")
	_err = errors.New("HostPatch.GetNameLabel not mocked")
	return
}
// Get the name/label field of the given host_patch.
func (_class HostPatchClass) GetNameLabel(sessionID SessionRef, self HostPatchRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetNameLabel__mock(sessionID, self)
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

func (_class HostPatchClass) GetUUID__mock(sessionID SessionRef, self HostPatchRef) (_retval string, _err error) {
	log.Println("HostPatch.GetUUID not mocked")
	_err = errors.New("HostPatch.GetUUID not mocked")
	return
}
// Get the uuid field of the given host_patch.
func (_class HostPatchClass) GetUUID(sessionID SessionRef, self HostPatchRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetUUID__mock(sessionID, self)
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

func (_class HostPatchClass) GetByNameLabel__mock(sessionID SessionRef, label string) (_retval []HostPatchRef, _err error) {
	log.Println("HostPatch.GetByNameLabel not mocked")
	_err = errors.New("HostPatch.GetByNameLabel not mocked")
	return
}
// Get all the host_patch instances with the given label.
func (_class HostPatchClass) GetByNameLabel(sessionID SessionRef, label string) (_retval []HostPatchRef, _err error) {
	if (IsMock) {
		return _class.GetByNameLabel__mock(sessionID, label)
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

func (_class HostPatchClass) GetByUUID__mock(sessionID SessionRef, uuid string) (_retval HostPatchRef, _err error) {
	log.Println("HostPatch.GetByUUID not mocked")
	_err = errors.New("HostPatch.GetByUUID not mocked")
	return
}
// Get a reference to the host_patch instance with the specified UUID.
func (_class HostPatchClass) GetByUUID(sessionID SessionRef, uuid string) (_retval HostPatchRef, _err error) {
	if (IsMock) {
		return _class.GetByUUID__mock(sessionID, uuid)
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

func (_class HostPatchClass) GetRecord__mock(sessionID SessionRef, self HostPatchRef) (_retval HostPatchRecord, _err error) {
	log.Println("HostPatch.GetRecord not mocked")
	_err = errors.New("HostPatch.GetRecord not mocked")
	return
}
// Get a record containing the current state of the given host_patch.
func (_class HostPatchClass) GetRecord(sessionID SessionRef, self HostPatchRef) (_retval HostPatchRecord, _err error) {
	if (IsMock) {
		return _class.GetRecord__mock(sessionID, self)
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
