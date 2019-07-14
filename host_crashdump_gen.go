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


func HostCrashdumpClassGetAllRecordsMockDefault(sessionID SessionRef) (_retval map[HostCrashdumpRef]HostCrashdumpRecord, _err error) {
	log.Println("HostCrashdump.GetAllRecords not mocked")
	_err = errors.New("HostCrashdump.GetAllRecords not mocked")
	return
}

var HostCrashdumpClassGetAllRecordsMockedCallback = HostCrashdumpClassGetAllRecordsMockDefault

func (_class HostCrashdumpClass) GetAllRecordsMock(sessionID SessionRef) (_retval map[HostCrashdumpRef]HostCrashdumpRecord, _err error) {
	return HostCrashdumpClassGetAllRecordsMockedCallback(sessionID)
}
// Return a map of host_crashdump references to host_crashdump records for all host_crashdumps known to the system.
func (_class HostCrashdumpClass) GetAllRecords(sessionID SessionRef) (_retval map[HostCrashdumpRef]HostCrashdumpRecord, _err error) {
	if IsMock {
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


func HostCrashdumpClassGetAllMockDefault(sessionID SessionRef) (_retval []HostCrashdumpRef, _err error) {
	log.Println("HostCrashdump.GetAll not mocked")
	_err = errors.New("HostCrashdump.GetAll not mocked")
	return
}

var HostCrashdumpClassGetAllMockedCallback = HostCrashdumpClassGetAllMockDefault

func (_class HostCrashdumpClass) GetAllMock(sessionID SessionRef) (_retval []HostCrashdumpRef, _err error) {
	return HostCrashdumpClassGetAllMockedCallback(sessionID)
}
// Return a list of all the host_crashdumps known to the system.
func (_class HostCrashdumpClass) GetAll(sessionID SessionRef) (_retval []HostCrashdumpRef, _err error) {
	if IsMock {
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


func HostCrashdumpClassUploadMockDefault(sessionID SessionRef, self HostCrashdumpRef, url string, options map[string]string) (_err error) {
	log.Println("HostCrashdump.Upload not mocked")
	_err = errors.New("HostCrashdump.Upload not mocked")
	return
}

var HostCrashdumpClassUploadMockedCallback = HostCrashdumpClassUploadMockDefault

func (_class HostCrashdumpClass) UploadMock(sessionID SessionRef, self HostCrashdumpRef, url string, options map[string]string) (_err error) {
	return HostCrashdumpClassUploadMockedCallback(sessionID, self, url, options)
}
// Upload the specified host crash dump to a specified URL
func (_class HostCrashdumpClass) Upload(sessionID SessionRef, self HostCrashdumpRef, url string, options map[string]string) (_err error) {
	if IsMock {
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


func HostCrashdumpClassDestroyMockDefault(sessionID SessionRef, self HostCrashdumpRef) (_err error) {
	log.Println("HostCrashdump.Destroy not mocked")
	_err = errors.New("HostCrashdump.Destroy not mocked")
	return
}

var HostCrashdumpClassDestroyMockedCallback = HostCrashdumpClassDestroyMockDefault

func (_class HostCrashdumpClass) DestroyMock(sessionID SessionRef, self HostCrashdumpRef) (_err error) {
	return HostCrashdumpClassDestroyMockedCallback(sessionID, self)
}
// Destroy specified host crash dump, removing it from the disk.
func (_class HostCrashdumpClass) Destroy(sessionID SessionRef, self HostCrashdumpRef) (_err error) {
	if IsMock {
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


func HostCrashdumpClassRemoveFromOtherConfigMockDefault(sessionID SessionRef, self HostCrashdumpRef, key string) (_err error) {
	log.Println("HostCrashdump.RemoveFromOtherConfig not mocked")
	_err = errors.New("HostCrashdump.RemoveFromOtherConfig not mocked")
	return
}

var HostCrashdumpClassRemoveFromOtherConfigMockedCallback = HostCrashdumpClassRemoveFromOtherConfigMockDefault

func (_class HostCrashdumpClass) RemoveFromOtherConfigMock(sessionID SessionRef, self HostCrashdumpRef, key string) (_err error) {
	return HostCrashdumpClassRemoveFromOtherConfigMockedCallback(sessionID, self, key)
}
// Remove the given key and its corresponding value from the other_config field of the given host_crashdump.  If the key is not in that Map, then do nothing.
func (_class HostCrashdumpClass) RemoveFromOtherConfig(sessionID SessionRef, self HostCrashdumpRef, key string) (_err error) {
	if IsMock {
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


func HostCrashdumpClassAddToOtherConfigMockDefault(sessionID SessionRef, self HostCrashdumpRef, key string, value string) (_err error) {
	log.Println("HostCrashdump.AddToOtherConfig not mocked")
	_err = errors.New("HostCrashdump.AddToOtherConfig not mocked")
	return
}

var HostCrashdumpClassAddToOtherConfigMockedCallback = HostCrashdumpClassAddToOtherConfigMockDefault

func (_class HostCrashdumpClass) AddToOtherConfigMock(sessionID SessionRef, self HostCrashdumpRef, key string, value string) (_err error) {
	return HostCrashdumpClassAddToOtherConfigMockedCallback(sessionID, self, key, value)
}
// Add the given key-value pair to the other_config field of the given host_crashdump.
func (_class HostCrashdumpClass) AddToOtherConfig(sessionID SessionRef, self HostCrashdumpRef, key string, value string) (_err error) {
	if IsMock {
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


func HostCrashdumpClassSetOtherConfigMockDefault(sessionID SessionRef, self HostCrashdumpRef, value map[string]string) (_err error) {
	log.Println("HostCrashdump.SetOtherConfig not mocked")
	_err = errors.New("HostCrashdump.SetOtherConfig not mocked")
	return
}

var HostCrashdumpClassSetOtherConfigMockedCallback = HostCrashdumpClassSetOtherConfigMockDefault

func (_class HostCrashdumpClass) SetOtherConfigMock(sessionID SessionRef, self HostCrashdumpRef, value map[string]string) (_err error) {
	return HostCrashdumpClassSetOtherConfigMockedCallback(sessionID, self, value)
}
// Set the other_config field of the given host_crashdump.
func (_class HostCrashdumpClass) SetOtherConfig(sessionID SessionRef, self HostCrashdumpRef, value map[string]string) (_err error) {
	if IsMock {
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


func HostCrashdumpClassGetOtherConfigMockDefault(sessionID SessionRef, self HostCrashdumpRef) (_retval map[string]string, _err error) {
	log.Println("HostCrashdump.GetOtherConfig not mocked")
	_err = errors.New("HostCrashdump.GetOtherConfig not mocked")
	return
}

var HostCrashdumpClassGetOtherConfigMockedCallback = HostCrashdumpClassGetOtherConfigMockDefault

func (_class HostCrashdumpClass) GetOtherConfigMock(sessionID SessionRef, self HostCrashdumpRef) (_retval map[string]string, _err error) {
	return HostCrashdumpClassGetOtherConfigMockedCallback(sessionID, self)
}
// Get the other_config field of the given host_crashdump.
func (_class HostCrashdumpClass) GetOtherConfig(sessionID SessionRef, self HostCrashdumpRef) (_retval map[string]string, _err error) {
	if IsMock {
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


func HostCrashdumpClassGetSizeMockDefault(sessionID SessionRef, self HostCrashdumpRef) (_retval int, _err error) {
	log.Println("HostCrashdump.GetSize not mocked")
	_err = errors.New("HostCrashdump.GetSize not mocked")
	return
}

var HostCrashdumpClassGetSizeMockedCallback = HostCrashdumpClassGetSizeMockDefault

func (_class HostCrashdumpClass) GetSizeMock(sessionID SessionRef, self HostCrashdumpRef) (_retval int, _err error) {
	return HostCrashdumpClassGetSizeMockedCallback(sessionID, self)
}
// Get the size field of the given host_crashdump.
func (_class HostCrashdumpClass) GetSize(sessionID SessionRef, self HostCrashdumpRef) (_retval int, _err error) {
	if IsMock {
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


func HostCrashdumpClassGetTimestampMockDefault(sessionID SessionRef, self HostCrashdumpRef) (_retval time.Time, _err error) {
	log.Println("HostCrashdump.GetTimestamp not mocked")
	_err = errors.New("HostCrashdump.GetTimestamp not mocked")
	return
}

var HostCrashdumpClassGetTimestampMockedCallback = HostCrashdumpClassGetTimestampMockDefault

func (_class HostCrashdumpClass) GetTimestampMock(sessionID SessionRef, self HostCrashdumpRef) (_retval time.Time, _err error) {
	return HostCrashdumpClassGetTimestampMockedCallback(sessionID, self)
}
// Get the timestamp field of the given host_crashdump.
func (_class HostCrashdumpClass) GetTimestamp(sessionID SessionRef, self HostCrashdumpRef) (_retval time.Time, _err error) {
	if IsMock {
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


func HostCrashdumpClassGetHostMockDefault(sessionID SessionRef, self HostCrashdumpRef) (_retval HostRef, _err error) {
	log.Println("HostCrashdump.GetHost not mocked")
	_err = errors.New("HostCrashdump.GetHost not mocked")
	return
}

var HostCrashdumpClassGetHostMockedCallback = HostCrashdumpClassGetHostMockDefault

func (_class HostCrashdumpClass) GetHostMock(sessionID SessionRef, self HostCrashdumpRef) (_retval HostRef, _err error) {
	return HostCrashdumpClassGetHostMockedCallback(sessionID, self)
}
// Get the host field of the given host_crashdump.
func (_class HostCrashdumpClass) GetHost(sessionID SessionRef, self HostCrashdumpRef) (_retval HostRef, _err error) {
	if IsMock {
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


func HostCrashdumpClassGetUUIDMockDefault(sessionID SessionRef, self HostCrashdumpRef) (_retval string, _err error) {
	log.Println("HostCrashdump.GetUUID not mocked")
	_err = errors.New("HostCrashdump.GetUUID not mocked")
	return
}

var HostCrashdumpClassGetUUIDMockedCallback = HostCrashdumpClassGetUUIDMockDefault

func (_class HostCrashdumpClass) GetUUIDMock(sessionID SessionRef, self HostCrashdumpRef) (_retval string, _err error) {
	return HostCrashdumpClassGetUUIDMockedCallback(sessionID, self)
}
// Get the uuid field of the given host_crashdump.
func (_class HostCrashdumpClass) GetUUID(sessionID SessionRef, self HostCrashdumpRef) (_retval string, _err error) {
	if IsMock {
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


func HostCrashdumpClassGetByUUIDMockDefault(sessionID SessionRef, uuid string) (_retval HostCrashdumpRef, _err error) {
	log.Println("HostCrashdump.GetByUUID not mocked")
	_err = errors.New("HostCrashdump.GetByUUID not mocked")
	return
}

var HostCrashdumpClassGetByUUIDMockedCallback = HostCrashdumpClassGetByUUIDMockDefault

func (_class HostCrashdumpClass) GetByUUIDMock(sessionID SessionRef, uuid string) (_retval HostCrashdumpRef, _err error) {
	return HostCrashdumpClassGetByUUIDMockedCallback(sessionID, uuid)
}
// Get a reference to the host_crashdump instance with the specified UUID.
func (_class HostCrashdumpClass) GetByUUID(sessionID SessionRef, uuid string) (_retval HostCrashdumpRef, _err error) {
	if IsMock {
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


func HostCrashdumpClassGetRecordMockDefault(sessionID SessionRef, self HostCrashdumpRef) (_retval HostCrashdumpRecord, _err error) {
	log.Println("HostCrashdump.GetRecord not mocked")
	_err = errors.New("HostCrashdump.GetRecord not mocked")
	return
}

var HostCrashdumpClassGetRecordMockedCallback = HostCrashdumpClassGetRecordMockDefault

func (_class HostCrashdumpClass) GetRecordMock(sessionID SessionRef, self HostCrashdumpRef) (_retval HostCrashdumpRecord, _err error) {
	return HostCrashdumpClassGetRecordMockedCallback(sessionID, self)
}
// Get a record containing the current state of the given host_crashdump.
func (_class HostCrashdumpClass) GetRecord(sessionID SessionRef, self HostCrashdumpRef) (_retval HostCrashdumpRecord, _err error) {
	if IsMock {
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
