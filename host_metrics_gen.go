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

type HostMetricsRecord struct {
  // Unique identifier/object reference
	UUID string
  // Total host memory (bytes)
	MemoryTotal int
  // Free host memory (bytes)
	MemoryFree int
  // Pool master thinks this host is live
	Live bool
  // Time at which this information was last updated
	LastUpdated time.Time
  // additional configuration
	OtherConfig map[string]string
}

type HostMetricsRef string

// The metrics associated with a host
type HostMetricsClass struct {
	client *Client
}


var HostMetricsClass_GetAllRecordsMockedCallback = func (sessionID SessionRef) (_retval map[HostMetricsRef]HostMetricsRecord, _err error) {
	log.Println("HostMetrics.GetAllRecords not mocked")
	_err = errors.New("HostMetrics.GetAllRecords not mocked")
	return
}

func (_class HostMetricsClass) GetAllRecordsMock(sessionID SessionRef) (_retval map[HostMetricsRef]HostMetricsRecord, _err error) {
	return HostMetricsClass_GetAllRecordsMockedCallback(sessionID)
}
// Return a map of host_metrics references to host_metrics records for all host_metrics instances known to the system.
func (_class HostMetricsClass) GetAllRecords(sessionID SessionRef) (_retval map[HostMetricsRef]HostMetricsRecord, _err error) {
	if (IsMock) {
		return _class.GetAllRecordsMock(sessionID)
	}	
	_method := "host_metrics.get_all_records"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertHostMetricsRefToHostMetricsRecordMapToGo(_method + " -> ", _result.Value)
	return
}


var HostMetricsClass_GetAllMockedCallback = func (sessionID SessionRef) (_retval []HostMetricsRef, _err error) {
	log.Println("HostMetrics.GetAll not mocked")
	_err = errors.New("HostMetrics.GetAll not mocked")
	return
}

func (_class HostMetricsClass) GetAllMock(sessionID SessionRef) (_retval []HostMetricsRef, _err error) {
	return HostMetricsClass_GetAllMockedCallback(sessionID)
}
// Return a list of all the host_metrics instances known to the system.
func (_class HostMetricsClass) GetAll(sessionID SessionRef) (_retval []HostMetricsRef, _err error) {
	if (IsMock) {
		return _class.GetAllMock(sessionID)
	}	
	_method := "host_metrics.get_all"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertHostMetricsRefSetToGo(_method + " -> ", _result.Value)
	return
}


var HostMetricsClass_RemoveFromOtherConfigMockedCallback = func (sessionID SessionRef, self HostMetricsRef, key string) (_err error) {
	log.Println("HostMetrics.RemoveFromOtherConfig not mocked")
	_err = errors.New("HostMetrics.RemoveFromOtherConfig not mocked")
	return
}

func (_class HostMetricsClass) RemoveFromOtherConfigMock(sessionID SessionRef, self HostMetricsRef, key string) (_err error) {
	return HostMetricsClass_RemoveFromOtherConfigMockedCallback(sessionID, self, key)
}
// Remove the given key and its corresponding value from the other_config field of the given host_metrics.  If the key is not in that Map, then do nothing.
func (_class HostMetricsClass) RemoveFromOtherConfig(sessionID SessionRef, self HostMetricsRef, key string) (_err error) {
	if (IsMock) {
		return _class.RemoveFromOtherConfigMock(sessionID, self, key)
	}	
	_method := "host_metrics.remove_from_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertHostMetricsRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var HostMetricsClass_AddToOtherConfigMockedCallback = func (sessionID SessionRef, self HostMetricsRef, key string, value string) (_err error) {
	log.Println("HostMetrics.AddToOtherConfig not mocked")
	_err = errors.New("HostMetrics.AddToOtherConfig not mocked")
	return
}

func (_class HostMetricsClass) AddToOtherConfigMock(sessionID SessionRef, self HostMetricsRef, key string, value string) (_err error) {
	return HostMetricsClass_AddToOtherConfigMockedCallback(sessionID, self, key, value)
}
// Add the given key-value pair to the other_config field of the given host_metrics.
func (_class HostMetricsClass) AddToOtherConfig(sessionID SessionRef, self HostMetricsRef, key string, value string) (_err error) {
	if (IsMock) {
		return _class.AddToOtherConfigMock(sessionID, self, key, value)
	}	
	_method := "host_metrics.add_to_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertHostMetricsRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var HostMetricsClass_SetOtherConfigMockedCallback = func (sessionID SessionRef, self HostMetricsRef, value map[string]string) (_err error) {
	log.Println("HostMetrics.SetOtherConfig not mocked")
	_err = errors.New("HostMetrics.SetOtherConfig not mocked")
	return
}

func (_class HostMetricsClass) SetOtherConfigMock(sessionID SessionRef, self HostMetricsRef, value map[string]string) (_err error) {
	return HostMetricsClass_SetOtherConfigMockedCallback(sessionID, self, value)
}
// Set the other_config field of the given host_metrics.
func (_class HostMetricsClass) SetOtherConfig(sessionID SessionRef, self HostMetricsRef, value map[string]string) (_err error) {
	if (IsMock) {
		return _class.SetOtherConfigMock(sessionID, self, value)
	}	
	_method := "host_metrics.set_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertHostMetricsRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var HostMetricsClass_GetOtherConfigMockedCallback = func (sessionID SessionRef, self HostMetricsRef) (_retval map[string]string, _err error) {
	log.Println("HostMetrics.GetOtherConfig not mocked")
	_err = errors.New("HostMetrics.GetOtherConfig not mocked")
	return
}

func (_class HostMetricsClass) GetOtherConfigMock(sessionID SessionRef, self HostMetricsRef) (_retval map[string]string, _err error) {
	return HostMetricsClass_GetOtherConfigMockedCallback(sessionID, self)
}
// Get the other_config field of the given host_metrics.
func (_class HostMetricsClass) GetOtherConfig(sessionID SessionRef, self HostMetricsRef) (_retval map[string]string, _err error) {
	if (IsMock) {
		return _class.GetOtherConfigMock(sessionID, self)
	}	
	_method := "host_metrics.get_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertHostMetricsRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var HostMetricsClass_GetLastUpdatedMockedCallback = func (sessionID SessionRef, self HostMetricsRef) (_retval time.Time, _err error) {
	log.Println("HostMetrics.GetLastUpdated not mocked")
	_err = errors.New("HostMetrics.GetLastUpdated not mocked")
	return
}

func (_class HostMetricsClass) GetLastUpdatedMock(sessionID SessionRef, self HostMetricsRef) (_retval time.Time, _err error) {
	return HostMetricsClass_GetLastUpdatedMockedCallback(sessionID, self)
}
// Get the last_updated field of the given host_metrics.
func (_class HostMetricsClass) GetLastUpdated(sessionID SessionRef, self HostMetricsRef) (_retval time.Time, _err error) {
	if (IsMock) {
		return _class.GetLastUpdatedMock(sessionID, self)
	}	
	_method := "host_metrics.get_last_updated"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertHostMetricsRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var HostMetricsClass_GetLiveMockedCallback = func (sessionID SessionRef, self HostMetricsRef) (_retval bool, _err error) {
	log.Println("HostMetrics.GetLive not mocked")
	_err = errors.New("HostMetrics.GetLive not mocked")
	return
}

func (_class HostMetricsClass) GetLiveMock(sessionID SessionRef, self HostMetricsRef) (_retval bool, _err error) {
	return HostMetricsClass_GetLiveMockedCallback(sessionID, self)
}
// Get the live field of the given host_metrics.
func (_class HostMetricsClass) GetLive(sessionID SessionRef, self HostMetricsRef) (_retval bool, _err error) {
	if (IsMock) {
		return _class.GetLiveMock(sessionID, self)
	}	
	_method := "host_metrics.get_live"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertHostMetricsRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var HostMetricsClass_GetMemoryFreeMockedCallback = func (sessionID SessionRef, self HostMetricsRef) (_retval int, _err error) {
	log.Println("HostMetrics.GetMemoryFree not mocked")
	_err = errors.New("HostMetrics.GetMemoryFree not mocked")
	return
}

func (_class HostMetricsClass) GetMemoryFreeMock(sessionID SessionRef, self HostMetricsRef) (_retval int, _err error) {
	return HostMetricsClass_GetMemoryFreeMockedCallback(sessionID, self)
}
// Get the memory/free field of the given host_metrics.
func (_class HostMetricsClass) GetMemoryFree(sessionID SessionRef, self HostMetricsRef) (_retval int, _err error) {
	if (IsMock) {
		return _class.GetMemoryFreeMock(sessionID, self)
	}	
	_method := "host_metrics.get_memory_free"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertHostMetricsRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var HostMetricsClass_GetMemoryTotalMockedCallback = func (sessionID SessionRef, self HostMetricsRef) (_retval int, _err error) {
	log.Println("HostMetrics.GetMemoryTotal not mocked")
	_err = errors.New("HostMetrics.GetMemoryTotal not mocked")
	return
}

func (_class HostMetricsClass) GetMemoryTotalMock(sessionID SessionRef, self HostMetricsRef) (_retval int, _err error) {
	return HostMetricsClass_GetMemoryTotalMockedCallback(sessionID, self)
}
// Get the memory/total field of the given host_metrics.
func (_class HostMetricsClass) GetMemoryTotal(sessionID SessionRef, self HostMetricsRef) (_retval int, _err error) {
	if (IsMock) {
		return _class.GetMemoryTotalMock(sessionID, self)
	}	
	_method := "host_metrics.get_memory_total"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertHostMetricsRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var HostMetricsClass_GetUUIDMockedCallback = func (sessionID SessionRef, self HostMetricsRef) (_retval string, _err error) {
	log.Println("HostMetrics.GetUUID not mocked")
	_err = errors.New("HostMetrics.GetUUID not mocked")
	return
}

func (_class HostMetricsClass) GetUUIDMock(sessionID SessionRef, self HostMetricsRef) (_retval string, _err error) {
	return HostMetricsClass_GetUUIDMockedCallback(sessionID, self)
}
// Get the uuid field of the given host_metrics.
func (_class HostMetricsClass) GetUUID(sessionID SessionRef, self HostMetricsRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetUUIDMock(sessionID, self)
	}	
	_method := "host_metrics.get_uuid"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertHostMetricsRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var HostMetricsClass_GetByUUIDMockedCallback = func (sessionID SessionRef, uuid string) (_retval HostMetricsRef, _err error) {
	log.Println("HostMetrics.GetByUUID not mocked")
	_err = errors.New("HostMetrics.GetByUUID not mocked")
	return
}

func (_class HostMetricsClass) GetByUUIDMock(sessionID SessionRef, uuid string) (_retval HostMetricsRef, _err error) {
	return HostMetricsClass_GetByUUIDMockedCallback(sessionID, uuid)
}
// Get a reference to the host_metrics instance with the specified UUID.
func (_class HostMetricsClass) GetByUUID(sessionID SessionRef, uuid string) (_retval HostMetricsRef, _err error) {
	if (IsMock) {
		return _class.GetByUUIDMock(sessionID, uuid)
	}	
	_method := "host_metrics.get_by_uuid"
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
	_retval, _err = convertHostMetricsRefToGo(_method + " -> ", _result.Value)
	return
}


var HostMetricsClass_GetRecordMockedCallback = func (sessionID SessionRef, self HostMetricsRef) (_retval HostMetricsRecord, _err error) {
	log.Println("HostMetrics.GetRecord not mocked")
	_err = errors.New("HostMetrics.GetRecord not mocked")
	return
}

func (_class HostMetricsClass) GetRecordMock(sessionID SessionRef, self HostMetricsRef) (_retval HostMetricsRecord, _err error) {
	return HostMetricsClass_GetRecordMockedCallback(sessionID, self)
}
// Get a record containing the current state of the given host_metrics.
func (_class HostMetricsClass) GetRecord(sessionID SessionRef, self HostMetricsRef) (_retval HostMetricsRecord, _err error) {
	if (IsMock) {
		return _class.GetRecordMock(sessionID, self)
	}	
	_method := "host_metrics.get_record"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertHostMetricsRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertHostMetricsRecordToGo(_method + " -> ", _result.Value)
	return
}
