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


func VBDMetricsClassGetAllRecordsMockDefault(sessionID SessionRef) (_retval map[VBDMetricsRef]VBDMetricsRecord, _err error) {
	log.Println("VBDMetrics.GetAllRecords not mocked")
	_err = errors.New("VBDMetrics.GetAllRecords not mocked")
	return
}

var VBDMetricsClassGetAllRecordsMockedCallback = VBDMetricsClassGetAllRecordsMockDefault

func (_class VBDMetricsClass) GetAllRecordsMock(sessionID SessionRef) (_retval map[VBDMetricsRef]VBDMetricsRecord, _err error) {
	return VBDMetricsClassGetAllRecordsMockedCallback(sessionID)
}
// Return a map of VBD_metrics references to VBD_metrics records for all VBD_metrics instances known to the system.
func (_class VBDMetricsClass) GetAllRecords(sessionID SessionRef) (_retval map[VBDMetricsRef]VBDMetricsRecord, _err error) {
	if IsMock {
		return _class.GetAllRecordsMock(sessionID)
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


func VBDMetricsClassGetAllMockDefault(sessionID SessionRef) (_retval []VBDMetricsRef, _err error) {
	log.Println("VBDMetrics.GetAll not mocked")
	_err = errors.New("VBDMetrics.GetAll not mocked")
	return
}

var VBDMetricsClassGetAllMockedCallback = VBDMetricsClassGetAllMockDefault

func (_class VBDMetricsClass) GetAllMock(sessionID SessionRef) (_retval []VBDMetricsRef, _err error) {
	return VBDMetricsClassGetAllMockedCallback(sessionID)
}
// Return a list of all the VBD_metrics instances known to the system.
func (_class VBDMetricsClass) GetAll(sessionID SessionRef) (_retval []VBDMetricsRef, _err error) {
	if IsMock {
		return _class.GetAllMock(sessionID)
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


func VBDMetricsClassRemoveFromOtherConfigMockDefault(sessionID SessionRef, self VBDMetricsRef, key string) (_err error) {
	log.Println("VBDMetrics.RemoveFromOtherConfig not mocked")
	_err = errors.New("VBDMetrics.RemoveFromOtherConfig not mocked")
	return
}

var VBDMetricsClassRemoveFromOtherConfigMockedCallback = VBDMetricsClassRemoveFromOtherConfigMockDefault

func (_class VBDMetricsClass) RemoveFromOtherConfigMock(sessionID SessionRef, self VBDMetricsRef, key string) (_err error) {
	return VBDMetricsClassRemoveFromOtherConfigMockedCallback(sessionID, self, key)
}
// Remove the given key and its corresponding value from the other_config field of the given VBD_metrics.  If the key is not in that Map, then do nothing.
func (_class VBDMetricsClass) RemoveFromOtherConfig(sessionID SessionRef, self VBDMetricsRef, key string) (_err error) {
	if IsMock {
		return _class.RemoveFromOtherConfigMock(sessionID, self, key)
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


func VBDMetricsClassAddToOtherConfigMockDefault(sessionID SessionRef, self VBDMetricsRef, key string, value string) (_err error) {
	log.Println("VBDMetrics.AddToOtherConfig not mocked")
	_err = errors.New("VBDMetrics.AddToOtherConfig not mocked")
	return
}

var VBDMetricsClassAddToOtherConfigMockedCallback = VBDMetricsClassAddToOtherConfigMockDefault

func (_class VBDMetricsClass) AddToOtherConfigMock(sessionID SessionRef, self VBDMetricsRef, key string, value string) (_err error) {
	return VBDMetricsClassAddToOtherConfigMockedCallback(sessionID, self, key, value)
}
// Add the given key-value pair to the other_config field of the given VBD_metrics.
func (_class VBDMetricsClass) AddToOtherConfig(sessionID SessionRef, self VBDMetricsRef, key string, value string) (_err error) {
	if IsMock {
		return _class.AddToOtherConfigMock(sessionID, self, key, value)
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


func VBDMetricsClassSetOtherConfigMockDefault(sessionID SessionRef, self VBDMetricsRef, value map[string]string) (_err error) {
	log.Println("VBDMetrics.SetOtherConfig not mocked")
	_err = errors.New("VBDMetrics.SetOtherConfig not mocked")
	return
}

var VBDMetricsClassSetOtherConfigMockedCallback = VBDMetricsClassSetOtherConfigMockDefault

func (_class VBDMetricsClass) SetOtherConfigMock(sessionID SessionRef, self VBDMetricsRef, value map[string]string) (_err error) {
	return VBDMetricsClassSetOtherConfigMockedCallback(sessionID, self, value)
}
// Set the other_config field of the given VBD_metrics.
func (_class VBDMetricsClass) SetOtherConfig(sessionID SessionRef, self VBDMetricsRef, value map[string]string) (_err error) {
	if IsMock {
		return _class.SetOtherConfigMock(sessionID, self, value)
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


func VBDMetricsClassGetOtherConfigMockDefault(sessionID SessionRef, self VBDMetricsRef) (_retval map[string]string, _err error) {
	log.Println("VBDMetrics.GetOtherConfig not mocked")
	_err = errors.New("VBDMetrics.GetOtherConfig not mocked")
	return
}

var VBDMetricsClassGetOtherConfigMockedCallback = VBDMetricsClassGetOtherConfigMockDefault

func (_class VBDMetricsClass) GetOtherConfigMock(sessionID SessionRef, self VBDMetricsRef) (_retval map[string]string, _err error) {
	return VBDMetricsClassGetOtherConfigMockedCallback(sessionID, self)
}
// Get the other_config field of the given VBD_metrics.
func (_class VBDMetricsClass) GetOtherConfig(sessionID SessionRef, self VBDMetricsRef) (_retval map[string]string, _err error) {
	if IsMock {
		return _class.GetOtherConfigMock(sessionID, self)
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


func VBDMetricsClassGetLastUpdatedMockDefault(sessionID SessionRef, self VBDMetricsRef) (_retval time.Time, _err error) {
	log.Println("VBDMetrics.GetLastUpdated not mocked")
	_err = errors.New("VBDMetrics.GetLastUpdated not mocked")
	return
}

var VBDMetricsClassGetLastUpdatedMockedCallback = VBDMetricsClassGetLastUpdatedMockDefault

func (_class VBDMetricsClass) GetLastUpdatedMock(sessionID SessionRef, self VBDMetricsRef) (_retval time.Time, _err error) {
	return VBDMetricsClassGetLastUpdatedMockedCallback(sessionID, self)
}
// Get the last_updated field of the given VBD_metrics.
func (_class VBDMetricsClass) GetLastUpdated(sessionID SessionRef, self VBDMetricsRef) (_retval time.Time, _err error) {
	if IsMock {
		return _class.GetLastUpdatedMock(sessionID, self)
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


func VBDMetricsClassGetIoWriteKbsMockDefault(sessionID SessionRef, self VBDMetricsRef) (_retval float64, _err error) {
	log.Println("VBDMetrics.GetIoWriteKbs not mocked")
	_err = errors.New("VBDMetrics.GetIoWriteKbs not mocked")
	return
}

var VBDMetricsClassGetIoWriteKbsMockedCallback = VBDMetricsClassGetIoWriteKbsMockDefault

func (_class VBDMetricsClass) GetIoWriteKbsMock(sessionID SessionRef, self VBDMetricsRef) (_retval float64, _err error) {
	return VBDMetricsClassGetIoWriteKbsMockedCallback(sessionID, self)
}
// Get the io/write_kbs field of the given VBD_metrics.
func (_class VBDMetricsClass) GetIoWriteKbs(sessionID SessionRef, self VBDMetricsRef) (_retval float64, _err error) {
	if IsMock {
		return _class.GetIoWriteKbsMock(sessionID, self)
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


func VBDMetricsClassGetIoReadKbsMockDefault(sessionID SessionRef, self VBDMetricsRef) (_retval float64, _err error) {
	log.Println("VBDMetrics.GetIoReadKbs not mocked")
	_err = errors.New("VBDMetrics.GetIoReadKbs not mocked")
	return
}

var VBDMetricsClassGetIoReadKbsMockedCallback = VBDMetricsClassGetIoReadKbsMockDefault

func (_class VBDMetricsClass) GetIoReadKbsMock(sessionID SessionRef, self VBDMetricsRef) (_retval float64, _err error) {
	return VBDMetricsClassGetIoReadKbsMockedCallback(sessionID, self)
}
// Get the io/read_kbs field of the given VBD_metrics.
func (_class VBDMetricsClass) GetIoReadKbs(sessionID SessionRef, self VBDMetricsRef) (_retval float64, _err error) {
	if IsMock {
		return _class.GetIoReadKbsMock(sessionID, self)
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


func VBDMetricsClassGetUUIDMockDefault(sessionID SessionRef, self VBDMetricsRef) (_retval string, _err error) {
	log.Println("VBDMetrics.GetUUID not mocked")
	_err = errors.New("VBDMetrics.GetUUID not mocked")
	return
}

var VBDMetricsClassGetUUIDMockedCallback = VBDMetricsClassGetUUIDMockDefault

func (_class VBDMetricsClass) GetUUIDMock(sessionID SessionRef, self VBDMetricsRef) (_retval string, _err error) {
	return VBDMetricsClassGetUUIDMockedCallback(sessionID, self)
}
// Get the uuid field of the given VBD_metrics.
func (_class VBDMetricsClass) GetUUID(sessionID SessionRef, self VBDMetricsRef) (_retval string, _err error) {
	if IsMock {
		return _class.GetUUIDMock(sessionID, self)
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


func VBDMetricsClassGetByUUIDMockDefault(sessionID SessionRef, uuid string) (_retval VBDMetricsRef, _err error) {
	log.Println("VBDMetrics.GetByUUID not mocked")
	_err = errors.New("VBDMetrics.GetByUUID not mocked")
	return
}

var VBDMetricsClassGetByUUIDMockedCallback = VBDMetricsClassGetByUUIDMockDefault

func (_class VBDMetricsClass) GetByUUIDMock(sessionID SessionRef, uuid string) (_retval VBDMetricsRef, _err error) {
	return VBDMetricsClassGetByUUIDMockedCallback(sessionID, uuid)
}
// Get a reference to the VBD_metrics instance with the specified UUID.
func (_class VBDMetricsClass) GetByUUID(sessionID SessionRef, uuid string) (_retval VBDMetricsRef, _err error) {
	if IsMock {
		return _class.GetByUUIDMock(sessionID, uuid)
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


func VBDMetricsClassGetRecordMockDefault(sessionID SessionRef, self VBDMetricsRef) (_retval VBDMetricsRecord, _err error) {
	log.Println("VBDMetrics.GetRecord not mocked")
	_err = errors.New("VBDMetrics.GetRecord not mocked")
	return
}

var VBDMetricsClassGetRecordMockedCallback = VBDMetricsClassGetRecordMockDefault

func (_class VBDMetricsClass) GetRecordMock(sessionID SessionRef, self VBDMetricsRef) (_retval VBDMetricsRecord, _err error) {
	return VBDMetricsClassGetRecordMockedCallback(sessionID, self)
}
// Get a record containing the current state of the given VBD_metrics.
func (_class VBDMetricsClass) GetRecord(sessionID SessionRef, self VBDMetricsRef) (_retval VBDMetricsRecord, _err error) {
	if IsMock {
		return _class.GetRecordMock(sessionID, self)
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
