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

type VIFMetricsRecord struct {
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

type VIFMetricsRef string

// The metrics associated with a virtual network device
type VIFMetricsClass struct {
	client *Client
}


var VIFMetricsClass_GetAllRecordsMockedCallback = func (sessionID SessionRef) (_retval map[VIFMetricsRef]VIFMetricsRecord, _err error) {
	log.Println("VIFMetrics.GetAllRecords not mocked")
	_err = errors.New("VIFMetrics.GetAllRecords not mocked")
	return
}

func (_class VIFMetricsClass) GetAllRecordsMock(sessionID SessionRef) (_retval map[VIFMetricsRef]VIFMetricsRecord, _err error) {
	return VIFMetricsClass_GetAllRecordsMockedCallback(sessionID)
}
// Return a map of VIF_metrics references to VIF_metrics records for all VIF_metrics instances known to the system.
func (_class VIFMetricsClass) GetAllRecords(sessionID SessionRef) (_retval map[VIFMetricsRef]VIFMetricsRecord, _err error) {
	if (IsMock) {
		return _class.GetAllRecordsMock(sessionID)
	}	
	_method := "VIF_metrics.get_all_records"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVIFMetricsRefToVIFMetricsRecordMapToGo(_method + " -> ", _result.Value)
	return
}


var VIFMetricsClass_GetAllMockedCallback = func (sessionID SessionRef) (_retval []VIFMetricsRef, _err error) {
	log.Println("VIFMetrics.GetAll not mocked")
	_err = errors.New("VIFMetrics.GetAll not mocked")
	return
}

func (_class VIFMetricsClass) GetAllMock(sessionID SessionRef) (_retval []VIFMetricsRef, _err error) {
	return VIFMetricsClass_GetAllMockedCallback(sessionID)
}
// Return a list of all the VIF_metrics instances known to the system.
func (_class VIFMetricsClass) GetAll(sessionID SessionRef) (_retval []VIFMetricsRef, _err error) {
	if (IsMock) {
		return _class.GetAllMock(sessionID)
	}	
	_method := "VIF_metrics.get_all"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVIFMetricsRefSetToGo(_method + " -> ", _result.Value)
	return
}


var VIFMetricsClass_RemoveFromOtherConfigMockedCallback = func (sessionID SessionRef, self VIFMetricsRef, key string) (_err error) {
	log.Println("VIFMetrics.RemoveFromOtherConfig not mocked")
	_err = errors.New("VIFMetrics.RemoveFromOtherConfig not mocked")
	return
}

func (_class VIFMetricsClass) RemoveFromOtherConfigMock(sessionID SessionRef, self VIFMetricsRef, key string) (_err error) {
	return VIFMetricsClass_RemoveFromOtherConfigMockedCallback(sessionID, self, key)
}
// Remove the given key and its corresponding value from the other_config field of the given VIF_metrics.  If the key is not in that Map, then do nothing.
func (_class VIFMetricsClass) RemoveFromOtherConfig(sessionID SessionRef, self VIFMetricsRef, key string) (_err error) {
	if (IsMock) {
		return _class.RemoveFromOtherConfigMock(sessionID, self, key)
	}	
	_method := "VIF_metrics.remove_from_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVIFMetricsRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var VIFMetricsClass_AddToOtherConfigMockedCallback = func (sessionID SessionRef, self VIFMetricsRef, key string, value string) (_err error) {
	log.Println("VIFMetrics.AddToOtherConfig not mocked")
	_err = errors.New("VIFMetrics.AddToOtherConfig not mocked")
	return
}

func (_class VIFMetricsClass) AddToOtherConfigMock(sessionID SessionRef, self VIFMetricsRef, key string, value string) (_err error) {
	return VIFMetricsClass_AddToOtherConfigMockedCallback(sessionID, self, key, value)
}
// Add the given key-value pair to the other_config field of the given VIF_metrics.
func (_class VIFMetricsClass) AddToOtherConfig(sessionID SessionRef, self VIFMetricsRef, key string, value string) (_err error) {
	if (IsMock) {
		return _class.AddToOtherConfigMock(sessionID, self, key, value)
	}	
	_method := "VIF_metrics.add_to_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVIFMetricsRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var VIFMetricsClass_SetOtherConfigMockedCallback = func (sessionID SessionRef, self VIFMetricsRef, value map[string]string) (_err error) {
	log.Println("VIFMetrics.SetOtherConfig not mocked")
	_err = errors.New("VIFMetrics.SetOtherConfig not mocked")
	return
}

func (_class VIFMetricsClass) SetOtherConfigMock(sessionID SessionRef, self VIFMetricsRef, value map[string]string) (_err error) {
	return VIFMetricsClass_SetOtherConfigMockedCallback(sessionID, self, value)
}
// Set the other_config field of the given VIF_metrics.
func (_class VIFMetricsClass) SetOtherConfig(sessionID SessionRef, self VIFMetricsRef, value map[string]string) (_err error) {
	if (IsMock) {
		return _class.SetOtherConfigMock(sessionID, self, value)
	}	
	_method := "VIF_metrics.set_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVIFMetricsRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var VIFMetricsClass_GetOtherConfigMockedCallback = func (sessionID SessionRef, self VIFMetricsRef) (_retval map[string]string, _err error) {
	log.Println("VIFMetrics.GetOtherConfig not mocked")
	_err = errors.New("VIFMetrics.GetOtherConfig not mocked")
	return
}

func (_class VIFMetricsClass) GetOtherConfigMock(sessionID SessionRef, self VIFMetricsRef) (_retval map[string]string, _err error) {
	return VIFMetricsClass_GetOtherConfigMockedCallback(sessionID, self)
}
// Get the other_config field of the given VIF_metrics.
func (_class VIFMetricsClass) GetOtherConfig(sessionID SessionRef, self VIFMetricsRef) (_retval map[string]string, _err error) {
	if (IsMock) {
		return _class.GetOtherConfigMock(sessionID, self)
	}	
	_method := "VIF_metrics.get_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVIFMetricsRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var VIFMetricsClass_GetLastUpdatedMockedCallback = func (sessionID SessionRef, self VIFMetricsRef) (_retval time.Time, _err error) {
	log.Println("VIFMetrics.GetLastUpdated not mocked")
	_err = errors.New("VIFMetrics.GetLastUpdated not mocked")
	return
}

func (_class VIFMetricsClass) GetLastUpdatedMock(sessionID SessionRef, self VIFMetricsRef) (_retval time.Time, _err error) {
	return VIFMetricsClass_GetLastUpdatedMockedCallback(sessionID, self)
}
// Get the last_updated field of the given VIF_metrics.
func (_class VIFMetricsClass) GetLastUpdated(sessionID SessionRef, self VIFMetricsRef) (_retval time.Time, _err error) {
	if (IsMock) {
		return _class.GetLastUpdatedMock(sessionID, self)
	}	
	_method := "VIF_metrics.get_last_updated"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVIFMetricsRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var VIFMetricsClass_GetIoWriteKbsMockedCallback = func (sessionID SessionRef, self VIFMetricsRef) (_retval float64, _err error) {
	log.Println("VIFMetrics.GetIoWriteKbs not mocked")
	_err = errors.New("VIFMetrics.GetIoWriteKbs not mocked")
	return
}

func (_class VIFMetricsClass) GetIoWriteKbsMock(sessionID SessionRef, self VIFMetricsRef) (_retval float64, _err error) {
	return VIFMetricsClass_GetIoWriteKbsMockedCallback(sessionID, self)
}
// Get the io/write_kbs field of the given VIF_metrics.
func (_class VIFMetricsClass) GetIoWriteKbs(sessionID SessionRef, self VIFMetricsRef) (_retval float64, _err error) {
	if (IsMock) {
		return _class.GetIoWriteKbsMock(sessionID, self)
	}	
	_method := "VIF_metrics.get_io_write_kbs"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVIFMetricsRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var VIFMetricsClass_GetIoReadKbsMockedCallback = func (sessionID SessionRef, self VIFMetricsRef) (_retval float64, _err error) {
	log.Println("VIFMetrics.GetIoReadKbs not mocked")
	_err = errors.New("VIFMetrics.GetIoReadKbs not mocked")
	return
}

func (_class VIFMetricsClass) GetIoReadKbsMock(sessionID SessionRef, self VIFMetricsRef) (_retval float64, _err error) {
	return VIFMetricsClass_GetIoReadKbsMockedCallback(sessionID, self)
}
// Get the io/read_kbs field of the given VIF_metrics.
func (_class VIFMetricsClass) GetIoReadKbs(sessionID SessionRef, self VIFMetricsRef) (_retval float64, _err error) {
	if (IsMock) {
		return _class.GetIoReadKbsMock(sessionID, self)
	}	
	_method := "VIF_metrics.get_io_read_kbs"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVIFMetricsRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var VIFMetricsClass_GetUUIDMockedCallback = func (sessionID SessionRef, self VIFMetricsRef) (_retval string, _err error) {
	log.Println("VIFMetrics.GetUUID not mocked")
	_err = errors.New("VIFMetrics.GetUUID not mocked")
	return
}

func (_class VIFMetricsClass) GetUUIDMock(sessionID SessionRef, self VIFMetricsRef) (_retval string, _err error) {
	return VIFMetricsClass_GetUUIDMockedCallback(sessionID, self)
}
// Get the uuid field of the given VIF_metrics.
func (_class VIFMetricsClass) GetUUID(sessionID SessionRef, self VIFMetricsRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetUUIDMock(sessionID, self)
	}	
	_method := "VIF_metrics.get_uuid"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVIFMetricsRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var VIFMetricsClass_GetByUUIDMockedCallback = func (sessionID SessionRef, uuid string) (_retval VIFMetricsRef, _err error) {
	log.Println("VIFMetrics.GetByUUID not mocked")
	_err = errors.New("VIFMetrics.GetByUUID not mocked")
	return
}

func (_class VIFMetricsClass) GetByUUIDMock(sessionID SessionRef, uuid string) (_retval VIFMetricsRef, _err error) {
	return VIFMetricsClass_GetByUUIDMockedCallback(sessionID, uuid)
}
// Get a reference to the VIF_metrics instance with the specified UUID.
func (_class VIFMetricsClass) GetByUUID(sessionID SessionRef, uuid string) (_retval VIFMetricsRef, _err error) {
	if (IsMock) {
		return _class.GetByUUIDMock(sessionID, uuid)
	}	
	_method := "VIF_metrics.get_by_uuid"
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
	_retval, _err = convertVIFMetricsRefToGo(_method + " -> ", _result.Value)
	return
}


var VIFMetricsClass_GetRecordMockedCallback = func (sessionID SessionRef, self VIFMetricsRef) (_retval VIFMetricsRecord, _err error) {
	log.Println("VIFMetrics.GetRecord not mocked")
	_err = errors.New("VIFMetrics.GetRecord not mocked")
	return
}

func (_class VIFMetricsClass) GetRecordMock(sessionID SessionRef, self VIFMetricsRef) (_retval VIFMetricsRecord, _err error) {
	return VIFMetricsClass_GetRecordMockedCallback(sessionID, self)
}
// Get a record containing the current state of the given VIF_metrics.
func (_class VIFMetricsClass) GetRecord(sessionID SessionRef, self VIFMetricsRef) (_retval VIFMetricsRecord, _err error) {
	if (IsMock) {
		return _class.GetRecordMock(sessionID, self)
	}	
	_method := "VIF_metrics.get_record"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVIFMetricsRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVIFMetricsRecordToGo(_method + " -> ", _result.Value)
	return
}
