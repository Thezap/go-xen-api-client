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

type HostCrashdumpRecord struct {
  // Unique identifier/object reference
	UUID string
  // Host the crashdump relates to
	Host HostRef
  // Time the crash happened
	Timestamp time.Time
  // Size of the crashdump
	Size int
  // additional configuration
	OtherConfig map[string]string
}

type HostCrashdumpRef string

// Represents a host crash dump
type HostCrashdumpClass struct {
	client *Client
}


var HostCrashdumpClass_GetAllRecordsMockedCallback = func (sessionID SessionRef) (_retval map[HostCrashdumpRef]HostCrashdumpRecord, _err error) {
	log.Println("HostCrashdump.GetAllRecords not mocked")
	_err = errors.New("HostCrashdump.GetAllRecords not mocked")
	return
}

func (_class HostCrashdumpClass) GetAllRecordsMock(sessionID SessionRef) (_retval map[HostCrashdumpRef]HostCrashdumpRecord, _err error) {
	return HostCrashdumpClass_GetAllRecordsMockedCallback(sessionID)
}
// Return a map of host_crashdump references to host_crashdump records for all host_crashdumps known to the system.
func (_class HostCrashdumpClass) GetAllRecords(sessionID SessionRef) (_retval map[HostCrashdumpRef]HostCrashdumpRecord, _err error) {
	if (IsMock) {
		return _class.GetAllRecordsMock(sessionID)
	}	
	_method := "host_crashdump.get_all_records"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertHostCrashdumpRefToHostCrashdumpRecordMapToGo(_method + " -> ", _result.Value)
	return
}


var HostCrashdumpClass_GetAllMockedCallback = func (sessionID SessionRef) (_retval []HostCrashdumpRef, _err error) {
	log.Println("HostCrashdump.GetAll not mocked")
	_err = errors.New("HostCrashdump.GetAll not mocked")
	return
}

func (_class HostCrashdumpClass) GetAllMock(sessionID SessionRef) (_retval []HostCrashdumpRef, _err error) {
	return HostCrashdumpClass_GetAllMockedCallback(sessionID)
}
// Return a list of all the host_crashdumps known to the system.
func (_class HostCrashdumpClass) GetAll(sessionID SessionRef) (_retval []HostCrashdumpRef, _err error) {
	if (IsMock) {
		return _class.GetAllMock(sessionID)
	}	
	_method := "host_crashdump.get_all"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertHostCrashdumpRefSetToGo(_method + " -> ", _result.Value)
	return
}


var HostCrashdumpClass_UploadMockedCallback = func (sessionID SessionRef, self HostCrashdumpRef, url string, options map[string]string) (_err error) {
	log.Println("HostCrashdump.Upload not mocked")
	_err = errors.New("HostCrashdump.Upload not mocked")
	return
}

func (_class HostCrashdumpClass) UploadMock(sessionID SessionRef, self HostCrashdumpRef, url string, options map[string]string) (_err error) {
	return HostCrashdumpClass_UploadMockedCallback(sessionID, self, url, options)
}
// Upload the specified host crash dump to a specified URL
func (_class HostCrashdumpClass) Upload(sessionID SessionRef, self HostCrashdumpRef, url string, options map[string]string) (_err error) {
	if (IsMock) {
		return _class.UploadMock(sessionID, self, url, options)
	}	
	_method := "host_crashdump.upload"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertHostCrashdumpRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _urlArg, _optionsArg)
	return
}


var HostCrashdumpClass_DestroyMockedCallback = func (sessionID SessionRef, self HostCrashdumpRef) (_err error) {
	log.Println("HostCrashdump.Destroy not mocked")
	_err = errors.New("HostCrashdump.Destroy not mocked")
	return
}

func (_class HostCrashdumpClass) DestroyMock(sessionID SessionRef, self HostCrashdumpRef) (_err error) {
	return HostCrashdumpClass_DestroyMockedCallback(sessionID, self)
}
// Destroy specified host crash dump, removing it from the disk.
func (_class HostCrashdumpClass) Destroy(sessionID SessionRef, self HostCrashdumpRef) (_err error) {
	if (IsMock) {
		return _class.DestroyMock(sessionID, self)
	}	
	_method := "host_crashdump.destroy"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertHostCrashdumpRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg)
	return
}


var HostCrashdumpClass_RemoveFromOtherConfigMockedCallback = func (sessionID SessionRef, self HostCrashdumpRef, key string) (_err error) {
	log.Println("HostCrashdump.RemoveFromOtherConfig not mocked")
	_err = errors.New("HostCrashdump.RemoveFromOtherConfig not mocked")
	return
}

func (_class HostCrashdumpClass) RemoveFromOtherConfigMock(sessionID SessionRef, self HostCrashdumpRef, key string) (_err error) {
	return HostCrashdumpClass_RemoveFromOtherConfigMockedCallback(sessionID, self, key)
}
// Remove the given key and its corresponding value from the other_config field of the given host_crashdump.  If the key is not in that Map, then do nothing.
func (_class HostCrashdumpClass) RemoveFromOtherConfig(sessionID SessionRef, self HostCrashdumpRef, key string) (_err error) {
	if (IsMock) {
		return _class.RemoveFromOtherConfigMock(sessionID, self, key)
	}	
	_method := "host_crashdump.remove_from_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertHostCrashdumpRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var HostCrashdumpClass_AddToOtherConfigMockedCallback = func (sessionID SessionRef, self HostCrashdumpRef, key string, value string) (_err error) {
	log.Println("HostCrashdump.AddToOtherConfig not mocked")
	_err = errors.New("HostCrashdump.AddToOtherConfig not mocked")
	return
}

func (_class HostCrashdumpClass) AddToOtherConfigMock(sessionID SessionRef, self HostCrashdumpRef, key string, value string) (_err error) {
	return HostCrashdumpClass_AddToOtherConfigMockedCallback(sessionID, self, key, value)
}
// Add the given key-value pair to the other_config field of the given host_crashdump.
func (_class HostCrashdumpClass) AddToOtherConfig(sessionID SessionRef, self HostCrashdumpRef, key string, value string) (_err error) {
	if (IsMock) {
		return _class.AddToOtherConfigMock(sessionID, self, key, value)
	}	
	_method := "host_crashdump.add_to_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertHostCrashdumpRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var HostCrashdumpClass_SetOtherConfigMockedCallback = func (sessionID SessionRef, self HostCrashdumpRef, value map[string]string) (_err error) {
	log.Println("HostCrashdump.SetOtherConfig not mocked")
	_err = errors.New("HostCrashdump.SetOtherConfig not mocked")
	return
}

func (_class HostCrashdumpClass) SetOtherConfigMock(sessionID SessionRef, self HostCrashdumpRef, value map[string]string) (_err error) {
	return HostCrashdumpClass_SetOtherConfigMockedCallback(sessionID, self, value)
}
// Set the other_config field of the given host_crashdump.
func (_class HostCrashdumpClass) SetOtherConfig(sessionID SessionRef, self HostCrashdumpRef, value map[string]string) (_err error) {
	if (IsMock) {
		return _class.SetOtherConfigMock(sessionID, self, value)
	}	
	_method := "host_crashdump.set_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertHostCrashdumpRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var HostCrashdumpClass_GetOtherConfigMockedCallback = func (sessionID SessionRef, self HostCrashdumpRef) (_retval map[string]string, _err error) {
	log.Println("HostCrashdump.GetOtherConfig not mocked")
	_err = errors.New("HostCrashdump.GetOtherConfig not mocked")
	return
}

func (_class HostCrashdumpClass) GetOtherConfigMock(sessionID SessionRef, self HostCrashdumpRef) (_retval map[string]string, _err error) {
	return HostCrashdumpClass_GetOtherConfigMockedCallback(sessionID, self)
}
// Get the other_config field of the given host_crashdump.
func (_class HostCrashdumpClass) GetOtherConfig(sessionID SessionRef, self HostCrashdumpRef) (_retval map[string]string, _err error) {
	if (IsMock) {
		return _class.GetOtherConfigMock(sessionID, self)
	}	
	_method := "host_crashdump.get_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertHostCrashdumpRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var HostCrashdumpClass_GetSizeMockedCallback = func (sessionID SessionRef, self HostCrashdumpRef) (_retval int, _err error) {
	log.Println("HostCrashdump.GetSize not mocked")
	_err = errors.New("HostCrashdump.GetSize not mocked")
	return
}

func (_class HostCrashdumpClass) GetSizeMock(sessionID SessionRef, self HostCrashdumpRef) (_retval int, _err error) {
	return HostCrashdumpClass_GetSizeMockedCallback(sessionID, self)
}
// Get the size field of the given host_crashdump.
func (_class HostCrashdumpClass) GetSize(sessionID SessionRef, self HostCrashdumpRef) (_retval int, _err error) {
	if (IsMock) {
		return _class.GetSizeMock(sessionID, self)
	}	
	_method := "host_crashdump.get_size"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertHostCrashdumpRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var HostCrashdumpClass_GetTimestampMockedCallback = func (sessionID SessionRef, self HostCrashdumpRef) (_retval time.Time, _err error) {
	log.Println("HostCrashdump.GetTimestamp not mocked")
	_err = errors.New("HostCrashdump.GetTimestamp not mocked")
	return
}

func (_class HostCrashdumpClass) GetTimestampMock(sessionID SessionRef, self HostCrashdumpRef) (_retval time.Time, _err error) {
	return HostCrashdumpClass_GetTimestampMockedCallback(sessionID, self)
}
// Get the timestamp field of the given host_crashdump.
func (_class HostCrashdumpClass) GetTimestamp(sessionID SessionRef, self HostCrashdumpRef) (_retval time.Time, _err error) {
	if (IsMock) {
		return _class.GetTimestampMock(sessionID, self)
	}	
	_method := "host_crashdump.get_timestamp"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertHostCrashdumpRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var HostCrashdumpClass_GetHostMockedCallback = func (sessionID SessionRef, self HostCrashdumpRef) (_retval HostRef, _err error) {
	log.Println("HostCrashdump.GetHost not mocked")
	_err = errors.New("HostCrashdump.GetHost not mocked")
	return
}

func (_class HostCrashdumpClass) GetHostMock(sessionID SessionRef, self HostCrashdumpRef) (_retval HostRef, _err error) {
	return HostCrashdumpClass_GetHostMockedCallback(sessionID, self)
}
// Get the host field of the given host_crashdump.
func (_class HostCrashdumpClass) GetHost(sessionID SessionRef, self HostCrashdumpRef) (_retval HostRef, _err error) {
	if (IsMock) {
		return _class.GetHostMock(sessionID, self)
	}	
	_method := "host_crashdump.get_host"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertHostCrashdumpRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var HostCrashdumpClass_GetUUIDMockedCallback = func (sessionID SessionRef, self HostCrashdumpRef) (_retval string, _err error) {
	log.Println("HostCrashdump.GetUUID not mocked")
	_err = errors.New("HostCrashdump.GetUUID not mocked")
	return
}

func (_class HostCrashdumpClass) GetUUIDMock(sessionID SessionRef, self HostCrashdumpRef) (_retval string, _err error) {
	return HostCrashdumpClass_GetUUIDMockedCallback(sessionID, self)
}
// Get the uuid field of the given host_crashdump.
func (_class HostCrashdumpClass) GetUUID(sessionID SessionRef, self HostCrashdumpRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetUUIDMock(sessionID, self)
	}	
	_method := "host_crashdump.get_uuid"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertHostCrashdumpRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var HostCrashdumpClass_GetByUUIDMockedCallback = func (sessionID SessionRef, uuid string) (_retval HostCrashdumpRef, _err error) {
	log.Println("HostCrashdump.GetByUUID not mocked")
	_err = errors.New("HostCrashdump.GetByUUID not mocked")
	return
}

func (_class HostCrashdumpClass) GetByUUIDMock(sessionID SessionRef, uuid string) (_retval HostCrashdumpRef, _err error) {
	return HostCrashdumpClass_GetByUUIDMockedCallback(sessionID, uuid)
}
// Get a reference to the host_crashdump instance with the specified UUID.
func (_class HostCrashdumpClass) GetByUUID(sessionID SessionRef, uuid string) (_retval HostCrashdumpRef, _err error) {
	if (IsMock) {
		return _class.GetByUUIDMock(sessionID, uuid)
	}	
	_method := "host_crashdump.get_by_uuid"
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
	_retval, _err = convertHostCrashdumpRefToGo(_method + " -> ", _result.Value)
	return
}


var HostCrashdumpClass_GetRecordMockedCallback = func (sessionID SessionRef, self HostCrashdumpRef) (_retval HostCrashdumpRecord, _err error) {
	log.Println("HostCrashdump.GetRecord not mocked")
	_err = errors.New("HostCrashdump.GetRecord not mocked")
	return
}

func (_class HostCrashdumpClass) GetRecordMock(sessionID SessionRef, self HostCrashdumpRef) (_retval HostCrashdumpRecord, _err error) {
	return HostCrashdumpClass_GetRecordMockedCallback(sessionID, self)
}
// Get a record containing the current state of the given host_crashdump.
func (_class HostCrashdumpClass) GetRecord(sessionID SessionRef, self HostCrashdumpRef) (_retval HostCrashdumpRecord, _err error) {
	if (IsMock) {
		return _class.GetRecordMock(sessionID, self)
	}	
	_method := "host_crashdump.get_record"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertHostCrashdumpRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertHostCrashdumpRecordToGo(_method + " -> ", _result.Value)
	return
}
