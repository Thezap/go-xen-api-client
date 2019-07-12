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

type VBDMetricsRecord struct {
  // Unique identifier/object reference
	UUID string
  // Read bandwidth (KiB/s)
	IoReadKbs float64
  // Write bandwidth (KiB/s)
	IoWriteKbs float64
  // Time at which this information was last updated
	LastUpdated time.Time
  // additional configuration
	OtherConfig map[string]string
}

type VBDMetricsRef string

// The metrics associated with a virtual block device
type VBDMetricsClass struct {
	client *Client
}

func (_class VBDMetricsClass) GetAllRecords__mock(sessionID SessionRef) (_retval map[VBDMetricsRef]VBDMetricsRecord, _err error) {
	log.Println("VBDMetrics.GetAllRecords not mocked")
	_err = errors.New("VBDMetrics.GetAllRecords not mocked")
	return
}
// Return a map of VBD_metrics references to VBD_metrics records for all VBD_metrics instances known to the system.
func (_class VBDMetricsClass) GetAllRecords(sessionID SessionRef) (_retval map[VBDMetricsRef]VBDMetricsRecord, _err error) {
	if (IsMock) {
		return _class.GetAllRecords__mock(sessionID)
	}	
	_method := "VBD_metrics.get_all_records"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVBDMetricsRefToVBDMetricsRecordMapToGo(_method + " -> ", _result.Value)
	return
}

func (_class VBDMetricsClass) GetAll__mock(sessionID SessionRef) (_retval []VBDMetricsRef, _err error) {
	log.Println("VBDMetrics.GetAll not mocked")
	_err = errors.New("VBDMetrics.GetAll not mocked")
	return
}
// Return a list of all the VBD_metrics instances known to the system.
func (_class VBDMetricsClass) GetAll(sessionID SessionRef) (_retval []VBDMetricsRef, _err error) {
	if (IsMock) {
		return _class.GetAll__mock(sessionID)
	}	
	_method := "VBD_metrics.get_all"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVBDMetricsRefSetToGo(_method + " -> ", _result.Value)
	return
}

func (_class VBDMetricsClass) RemoveFromOtherConfig__mock(sessionID SessionRef, self VBDMetricsRef, key string) (_err error) {
	log.Println("VBDMetrics.RemoveFromOtherConfig not mocked")
	_err = errors.New("VBDMetrics.RemoveFromOtherConfig not mocked")
	return
}
// Remove the given key and its corresponding value from the other_config field of the given VBD_metrics.  If the key is not in that Map, then do nothing.
func (_class VBDMetricsClass) RemoveFromOtherConfig(sessionID SessionRef, self VBDMetricsRef, key string) (_err error) {
	if (IsMock) {
		return _class.RemoveFromOtherConfig__mock(sessionID, self, key)
	}	
	_method := "VBD_metrics.remove_from_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVBDMetricsRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class VBDMetricsClass) AddToOtherConfig__mock(sessionID SessionRef, self VBDMetricsRef, key string, value string) (_err error) {
	log.Println("VBDMetrics.AddToOtherConfig not mocked")
	_err = errors.New("VBDMetrics.AddToOtherConfig not mocked")
	return
}
// Add the given key-value pair to the other_config field of the given VBD_metrics.
func (_class VBDMetricsClass) AddToOtherConfig(sessionID SessionRef, self VBDMetricsRef, key string, value string) (_err error) {
	if (IsMock) {
		return _class.AddToOtherConfig__mock(sessionID, self, key, value)
	}	
	_method := "VBD_metrics.add_to_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVBDMetricsRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class VBDMetricsClass) SetOtherConfig__mock(sessionID SessionRef, self VBDMetricsRef, value map[string]string) (_err error) {
	log.Println("VBDMetrics.SetOtherConfig not mocked")
	_err = errors.New("VBDMetrics.SetOtherConfig not mocked")
	return
}
// Set the other_config field of the given VBD_metrics.
func (_class VBDMetricsClass) SetOtherConfig(sessionID SessionRef, self VBDMetricsRef, value map[string]string) (_err error) {
	if (IsMock) {
		return _class.SetOtherConfig__mock(sessionID, self, value)
	}	
	_method := "VBD_metrics.set_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVBDMetricsRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class VBDMetricsClass) GetOtherConfig__mock(sessionID SessionRef, self VBDMetricsRef) (_retval map[string]string, _err error) {
	log.Println("VBDMetrics.GetOtherConfig not mocked")
	_err = errors.New("VBDMetrics.GetOtherConfig not mocked")
	return
}
// Get the other_config field of the given VBD_metrics.
func (_class VBDMetricsClass) GetOtherConfig(sessionID SessionRef, self VBDMetricsRef) (_retval map[string]string, _err error) {
	if (IsMock) {
		return _class.GetOtherConfig__mock(sessionID, self)
	}	
	_method := "VBD_metrics.get_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVBDMetricsRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class VBDMetricsClass) GetLastUpdated__mock(sessionID SessionRef, self VBDMetricsRef) (_retval time.Time, _err error) {
	log.Println("VBDMetrics.GetLastUpdated not mocked")
	_err = errors.New("VBDMetrics.GetLastUpdated not mocked")
	return
}
// Get the last_updated field of the given VBD_metrics.
func (_class VBDMetricsClass) GetLastUpdated(sessionID SessionRef, self VBDMetricsRef) (_retval time.Time, _err error) {
	if (IsMock) {
		return _class.GetLastUpdated__mock(sessionID, self)
	}	
	_method := "VBD_metrics.get_last_updated"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVBDMetricsRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class VBDMetricsClass) GetIoWriteKbs__mock(sessionID SessionRef, self VBDMetricsRef) (_retval float64, _err error) {
	log.Println("VBDMetrics.GetIoWriteKbs not mocked")
	_err = errors.New("VBDMetrics.GetIoWriteKbs not mocked")
	return
}
// Get the io/write_kbs field of the given VBD_metrics.
func (_class VBDMetricsClass) GetIoWriteKbs(sessionID SessionRef, self VBDMetricsRef) (_retval float64, _err error) {
	if (IsMock) {
		return _class.GetIoWriteKbs__mock(sessionID, self)
	}	
	_method := "VBD_metrics.get_io_write_kbs"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVBDMetricsRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertFloatToGo(_method + " -> ", _result.Value)
	return
}

func (_class VBDMetricsClass) GetIoReadKbs__mock(sessionID SessionRef, self VBDMetricsRef) (_retval float64, _err error) {
	log.Println("VBDMetrics.GetIoReadKbs not mocked")
	_err = errors.New("VBDMetrics.GetIoReadKbs not mocked")
	return
}
// Get the io/read_kbs field of the given VBD_metrics.
func (_class VBDMetricsClass) GetIoReadKbs(sessionID SessionRef, self VBDMetricsRef) (_retval float64, _err error) {
	if (IsMock) {
		return _class.GetIoReadKbs__mock(sessionID, self)
	}	
	_method := "VBD_metrics.get_io_read_kbs"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVBDMetricsRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertFloatToGo(_method + " -> ", _result.Value)
	return
}

func (_class VBDMetricsClass) GetUUID__mock(sessionID SessionRef, self VBDMetricsRef) (_retval string, _err error) {
	log.Println("VBDMetrics.GetUUID not mocked")
	_err = errors.New("VBDMetrics.GetUUID not mocked")
	return
}
// Get the uuid field of the given VBD_metrics.
func (_class VBDMetricsClass) GetUUID(sessionID SessionRef, self VBDMetricsRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetUUID__mock(sessionID, self)
	}	
	_method := "VBD_metrics.get_uuid"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVBDMetricsRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class VBDMetricsClass) GetByUUID__mock(sessionID SessionRef, uuid string) (_retval VBDMetricsRef, _err error) {
	log.Println("VBDMetrics.GetByUUID not mocked")
	_err = errors.New("VBDMetrics.GetByUUID not mocked")
	return
}
// Get a reference to the VBD_metrics instance with the specified UUID.
func (_class VBDMetricsClass) GetByUUID(sessionID SessionRef, uuid string) (_retval VBDMetricsRef, _err error) {
	if (IsMock) {
		return _class.GetByUUID__mock(sessionID, uuid)
	}	
	_method := "VBD_metrics.get_by_uuid"
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
	_retval, _err = convertVBDMetricsRefToGo(_method + " -> ", _result.Value)
	return
}

func (_class VBDMetricsClass) GetRecord__mock(sessionID SessionRef, self VBDMetricsRef) (_retval VBDMetricsRecord, _err error) {
	log.Println("VBDMetrics.GetRecord not mocked")
	_err = errors.New("VBDMetrics.GetRecord not mocked")
	return
}
// Get a record containing the current state of the given VBD_metrics.
func (_class VBDMetricsClass) GetRecord(sessionID SessionRef, self VBDMetricsRef) (_retval VBDMetricsRecord, _err error) {
	if (IsMock) {
		return _class.GetRecord__mock(sessionID, self)
	}	
	_method := "VBD_metrics.get_record"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVBDMetricsRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVBDMetricsRecordToGo(_method + " -> ", _result.Value)
	return
}
