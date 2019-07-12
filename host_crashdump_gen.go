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

func (_class HostCrashdumpClass) GetAllRecords__mock(sessionID SessionRef) (_retval map[HostCrashdumpRef]HostCrashdumpRecord, _err error) {
	log.Println("HostCrashdump.GetAllRecords not mocked")
	_err = errors.New("HostCrashdump.GetAllRecords not mocked")
	return
}
// Return a map of host_crashdump references to host_crashdump records for all host_crashdumps known to the system.
func (_class HostCrashdumpClass) GetAllRecords(sessionID SessionRef) (_retval map[HostCrashdumpRef]HostCrashdumpRecord, _err error) {
	if (IsMock) {
		return _class.GetAllRecords__mock(sessionID)
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

func (_class HostCrashdumpClass) GetAll__mock(sessionID SessionRef) (_retval []HostCrashdumpRef, _err error) {
	log.Println("HostCrashdump.GetAll not mocked")
	_err = errors.New("HostCrashdump.GetAll not mocked")
	return
}
// Return a list of all the host_crashdumps known to the system.
func (_class HostCrashdumpClass) GetAll(sessionID SessionRef) (_retval []HostCrashdumpRef, _err error) {
	if (IsMock) {
		return _class.GetAll__mock(sessionID)
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

func (_class HostCrashdumpClass) Upload__mock(sessionID SessionRef, self HostCrashdumpRef, url string, options map[string]string) (_err error) {
	log.Println("HostCrashdump.Upload not mocked")
	_err = errors.New("HostCrashdump.Upload not mocked")
	return
}
// Upload the specified host crash dump to a specified URL
func (_class HostCrashdumpClass) Upload(sessionID SessionRef, self HostCrashdumpRef, url string, options map[string]string) (_err error) {
	if (IsMock) {
		return _class.Upload__mock(sessionID, self, url, options)
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

func (_class HostCrashdumpClass) Destroy__mock(sessionID SessionRef, self HostCrashdumpRef) (_err error) {
	log.Println("HostCrashdump.Destroy not mocked")
	_err = errors.New("HostCrashdump.Destroy not mocked")
	return
}
// Destroy specified host crash dump, removing it from the disk.
func (_class HostCrashdumpClass) Destroy(sessionID SessionRef, self HostCrashdumpRef) (_err error) {
	if (IsMock) {
		return _class.Destroy__mock(sessionID, self)
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

func (_class HostCrashdumpClass) RemoveFromOtherConfig__mock(sessionID SessionRef, self HostCrashdumpRef, key string) (_err error) {
	log.Println("HostCrashdump.RemoveFromOtherConfig not mocked")
	_err = errors.New("HostCrashdump.RemoveFromOtherConfig not mocked")
	return
}
// Remove the given key and its corresponding value from the other_config field of the given host_crashdump.  If the key is not in that Map, then do nothing.
func (_class HostCrashdumpClass) RemoveFromOtherConfig(sessionID SessionRef, self HostCrashdumpRef, key string) (_err error) {
	if (IsMock) {
		return _class.RemoveFromOtherConfig__mock(sessionID, self, key)
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

func (_class HostCrashdumpClass) AddToOtherConfig__mock(sessionID SessionRef, self HostCrashdumpRef, key string, value string) (_err error) {
	log.Println("HostCrashdump.AddToOtherConfig not mocked")
	_err = errors.New("HostCrashdump.AddToOtherConfig not mocked")
	return
}
// Add the given key-value pair to the other_config field of the given host_crashdump.
func (_class HostCrashdumpClass) AddToOtherConfig(sessionID SessionRef, self HostCrashdumpRef, key string, value string) (_err error) {
	if (IsMock) {
		return _class.AddToOtherConfig__mock(sessionID, self, key, value)
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

func (_class HostCrashdumpClass) SetOtherConfig__mock(sessionID SessionRef, self HostCrashdumpRef, value map[string]string) (_err error) {
	log.Println("HostCrashdump.SetOtherConfig not mocked")
	_err = errors.New("HostCrashdump.SetOtherConfig not mocked")
	return
}
// Set the other_config field of the given host_crashdump.
func (_class HostCrashdumpClass) SetOtherConfig(sessionID SessionRef, self HostCrashdumpRef, value map[string]string) (_err error) {
	if (IsMock) {
		return _class.SetOtherConfig__mock(sessionID, self, value)
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

func (_class HostCrashdumpClass) GetOtherConfig__mock(sessionID SessionRef, self HostCrashdumpRef) (_retval map[string]string, _err error) {
	log.Println("HostCrashdump.GetOtherConfig not mocked")
	_err = errors.New("HostCrashdump.GetOtherConfig not mocked")
	return
}
// Get the other_config field of the given host_crashdump.
func (_class HostCrashdumpClass) GetOtherConfig(sessionID SessionRef, self HostCrashdumpRef) (_retval map[string]string, _err error) {
	if (IsMock) {
		return _class.GetOtherConfig__mock(sessionID, self)
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

func (_class HostCrashdumpClass) GetSize__mock(sessionID SessionRef, self HostCrashdumpRef) (_retval int, _err error) {
	log.Println("HostCrashdump.GetSize not mocked")
	_err = errors.New("HostCrashdump.GetSize not mocked")
	return
}
// Get the size field of the given host_crashdump.
func (_class HostCrashdumpClass) GetSize(sessionID SessionRef, self HostCrashdumpRef) (_retval int, _err error) {
	if (IsMock) {
		return _class.GetSize__mock(sessionID, self)
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

func (_class HostCrashdumpClass) GetTimestamp__mock(sessionID SessionRef, self HostCrashdumpRef) (_retval time.Time, _err error) {
	log.Println("HostCrashdump.GetTimestamp not mocked")
	_err = errors.New("HostCrashdump.GetTimestamp not mocked")
	return
}
// Get the timestamp field of the given host_crashdump.
func (_class HostCrashdumpClass) GetTimestamp(sessionID SessionRef, self HostCrashdumpRef) (_retval time.Time, _err error) {
	if (IsMock) {
		return _class.GetTimestamp__mock(sessionID, self)
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

func (_class HostCrashdumpClass) GetHost__mock(sessionID SessionRef, self HostCrashdumpRef) (_retval HostRef, _err error) {
	log.Println("HostCrashdump.GetHost not mocked")
	_err = errors.New("HostCrashdump.GetHost not mocked")
	return
}
// Get the host field of the given host_crashdump.
func (_class HostCrashdumpClass) GetHost(sessionID SessionRef, self HostCrashdumpRef) (_retval HostRef, _err error) {
	if (IsMock) {
		return _class.GetHost__mock(sessionID, self)
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

func (_class HostCrashdumpClass) GetUUID__mock(sessionID SessionRef, self HostCrashdumpRef) (_retval string, _err error) {
	log.Println("HostCrashdump.GetUUID not mocked")
	_err = errors.New("HostCrashdump.GetUUID not mocked")
	return
}
// Get the uuid field of the given host_crashdump.
func (_class HostCrashdumpClass) GetUUID(sessionID SessionRef, self HostCrashdumpRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetUUID__mock(sessionID, self)
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

func (_class HostCrashdumpClass) GetByUUID__mock(sessionID SessionRef, uuid string) (_retval HostCrashdumpRef, _err error) {
	log.Println("HostCrashdump.GetByUUID not mocked")
	_err = errors.New("HostCrashdump.GetByUUID not mocked")
	return
}
// Get a reference to the host_crashdump instance with the specified UUID.
func (_class HostCrashdumpClass) GetByUUID(sessionID SessionRef, uuid string) (_retval HostCrashdumpRef, _err error) {
	if (IsMock) {
		return _class.GetByUUID__mock(sessionID, uuid)
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

func (_class HostCrashdumpClass) GetRecord__mock(sessionID SessionRef, self HostCrashdumpRef) (_retval HostCrashdumpRecord, _err error) {
	log.Println("HostCrashdump.GetRecord not mocked")
	_err = errors.New("HostCrashdump.GetRecord not mocked")
	return
}
// Get a record containing the current state of the given host_crashdump.
func (_class HostCrashdumpClass) GetRecord(sessionID SessionRef, self HostCrashdumpRef) (_retval HostCrashdumpRecord, _err error) {
	if (IsMock) {
		return _class.GetRecord__mock(sessionID, self)
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
