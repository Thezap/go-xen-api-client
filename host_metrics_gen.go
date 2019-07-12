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

func (_class HostMetricsClass) GetAllRecords__mock(sessionID SessionRef) (_retval map[HostMetricsRef]HostMetricsRecord, _err error) {
	log.Println("HostMetrics.GetAllRecords not mocked")
	_err = errors.New("HostMetrics.GetAllRecords not mocked")
	return
}
// Return a map of host_metrics references to host_metrics records for all host_metrics instances known to the system.
func (_class HostMetricsClass) GetAllRecords(sessionID SessionRef) (_retval map[HostMetricsRef]HostMetricsRecord, _err error) {
	if (IsMock) {
		return _class.GetAllRecords__mock(sessionID)
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

func (_class HostMetricsClass) GetAll__mock(sessionID SessionRef) (_retval []HostMetricsRef, _err error) {
	log.Println("HostMetrics.GetAll not mocked")
	_err = errors.New("HostMetrics.GetAll not mocked")
	return
}
// Return a list of all the host_metrics instances known to the system.
func (_class HostMetricsClass) GetAll(sessionID SessionRef) (_retval []HostMetricsRef, _err error) {
	if (IsMock) {
		return _class.GetAll__mock(sessionID)
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

func (_class HostMetricsClass) RemoveFromOtherConfig__mock(sessionID SessionRef, self HostMetricsRef, key string) (_err error) {
	log.Println("HostMetrics.RemoveFromOtherConfig not mocked")
	_err = errors.New("HostMetrics.RemoveFromOtherConfig not mocked")
	return
}
// Remove the given key and its corresponding value from the other_config field of the given host_metrics.  If the key is not in that Map, then do nothing.
func (_class HostMetricsClass) RemoveFromOtherConfig(sessionID SessionRef, self HostMetricsRef, key string) (_err error) {
	if (IsMock) {
		return _class.RemoveFromOtherConfig__mock(sessionID, self, key)
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

func (_class HostMetricsClass) AddToOtherConfig__mock(sessionID SessionRef, self HostMetricsRef, key string, value string) (_err error) {
	log.Println("HostMetrics.AddToOtherConfig not mocked")
	_err = errors.New("HostMetrics.AddToOtherConfig not mocked")
	return
}
// Add the given key-value pair to the other_config field of the given host_metrics.
func (_class HostMetricsClass) AddToOtherConfig(sessionID SessionRef, self HostMetricsRef, key string, value string) (_err error) {
	if (IsMock) {
		return _class.AddToOtherConfig__mock(sessionID, self, key, value)
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

func (_class HostMetricsClass) SetOtherConfig__mock(sessionID SessionRef, self HostMetricsRef, value map[string]string) (_err error) {
	log.Println("HostMetrics.SetOtherConfig not mocked")
	_err = errors.New("HostMetrics.SetOtherConfig not mocked")
	return
}
// Set the other_config field of the given host_metrics.
func (_class HostMetricsClass) SetOtherConfig(sessionID SessionRef, self HostMetricsRef, value map[string]string) (_err error) {
	if (IsMock) {
		return _class.SetOtherConfig__mock(sessionID, self, value)
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

func (_class HostMetricsClass) GetOtherConfig__mock(sessionID SessionRef, self HostMetricsRef) (_retval map[string]string, _err error) {
	log.Println("HostMetrics.GetOtherConfig not mocked")
	_err = errors.New("HostMetrics.GetOtherConfig not mocked")
	return
}
// Get the other_config field of the given host_metrics.
func (_class HostMetricsClass) GetOtherConfig(sessionID SessionRef, self HostMetricsRef) (_retval map[string]string, _err error) {
	if (IsMock) {
		return _class.GetOtherConfig__mock(sessionID, self)
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

func (_class HostMetricsClass) GetLastUpdated__mock(sessionID SessionRef, self HostMetricsRef) (_retval time.Time, _err error) {
	log.Println("HostMetrics.GetLastUpdated not mocked")
	_err = errors.New("HostMetrics.GetLastUpdated not mocked")
	return
}
// Get the last_updated field of the given host_metrics.
func (_class HostMetricsClass) GetLastUpdated(sessionID SessionRef, self HostMetricsRef) (_retval time.Time, _err error) {
	if (IsMock) {
		return _class.GetLastUpdated__mock(sessionID, self)
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

func (_class HostMetricsClass) GetLive__mock(sessionID SessionRef, self HostMetricsRef) (_retval bool, _err error) {
	log.Println("HostMetrics.GetLive not mocked")
	_err = errors.New("HostMetrics.GetLive not mocked")
	return
}
// Get the live field of the given host_metrics.
func (_class HostMetricsClass) GetLive(sessionID SessionRef, self HostMetricsRef) (_retval bool, _err error) {
	if (IsMock) {
		return _class.GetLive__mock(sessionID, self)
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

func (_class HostMetricsClass) GetMemoryFree__mock(sessionID SessionRef, self HostMetricsRef) (_retval int, _err error) {
	log.Println("HostMetrics.GetMemoryFree not mocked")
	_err = errors.New("HostMetrics.GetMemoryFree not mocked")
	return
}
// Get the memory/free field of the given host_metrics.
func (_class HostMetricsClass) GetMemoryFree(sessionID SessionRef, self HostMetricsRef) (_retval int, _err error) {
	if (IsMock) {
		return _class.GetMemoryFree__mock(sessionID, self)
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

func (_class HostMetricsClass) GetMemoryTotal__mock(sessionID SessionRef, self HostMetricsRef) (_retval int, _err error) {
	log.Println("HostMetrics.GetMemoryTotal not mocked")
	_err = errors.New("HostMetrics.GetMemoryTotal not mocked")
	return
}
// Get the memory/total field of the given host_metrics.
func (_class HostMetricsClass) GetMemoryTotal(sessionID SessionRef, self HostMetricsRef) (_retval int, _err error) {
	if (IsMock) {
		return _class.GetMemoryTotal__mock(sessionID, self)
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

func (_class HostMetricsClass) GetUUID__mock(sessionID SessionRef, self HostMetricsRef) (_retval string, _err error) {
	log.Println("HostMetrics.GetUUID not mocked")
	_err = errors.New("HostMetrics.GetUUID not mocked")
	return
}
// Get the uuid field of the given host_metrics.
func (_class HostMetricsClass) GetUUID(sessionID SessionRef, self HostMetricsRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetUUID__mock(sessionID, self)
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

func (_class HostMetricsClass) GetByUUID__mock(sessionID SessionRef, uuid string) (_retval HostMetricsRef, _err error) {
	log.Println("HostMetrics.GetByUUID not mocked")
	_err = errors.New("HostMetrics.GetByUUID not mocked")
	return
}
// Get a reference to the host_metrics instance with the specified UUID.
func (_class HostMetricsClass) GetByUUID(sessionID SessionRef, uuid string) (_retval HostMetricsRef, _err error) {
	if (IsMock) {
		return _class.GetByUUID__mock(sessionID, uuid)
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

func (_class HostMetricsClass) GetRecord__mock(sessionID SessionRef, self HostMetricsRef) (_retval HostMetricsRecord, _err error) {
	log.Println("HostMetrics.GetRecord not mocked")
	_err = errors.New("HostMetrics.GetRecord not mocked")
	return
}
// Get a record containing the current state of the given host_metrics.
func (_class HostMetricsClass) GetRecord(sessionID SessionRef, self HostMetricsRef) (_retval HostMetricsRecord, _err error) {
	if (IsMock) {
		return _class.GetRecord__mock(sessionID, self)
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
