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

func (_class VIFMetricsClass) GetAllRecords__mock(sessionID SessionRef) (_retval map[VIFMetricsRef]VIFMetricsRecord, _err error) {
	log.Println("VIFMetrics.GetAllRecords not mocked")
	_err = errors.New("VIFMetrics.GetAllRecords not mocked")
	return
}
// Return a map of VIF_metrics references to VIF_metrics records for all VIF_metrics instances known to the system.
func (_class VIFMetricsClass) GetAllRecords(sessionID SessionRef) (_retval map[VIFMetricsRef]VIFMetricsRecord, _err error) {
	if (IsMock) {
		return _class.GetAllRecords__mock(sessionID)
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

func (_class VIFMetricsClass) GetAll__mock(sessionID SessionRef) (_retval []VIFMetricsRef, _err error) {
	log.Println("VIFMetrics.GetAll not mocked")
	_err = errors.New("VIFMetrics.GetAll not mocked")
	return
}
// Return a list of all the VIF_metrics instances known to the system.
func (_class VIFMetricsClass) GetAll(sessionID SessionRef) (_retval []VIFMetricsRef, _err error) {
	if (IsMock) {
		return _class.GetAll__mock(sessionID)
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

func (_class VIFMetricsClass) RemoveFromOtherConfig__mock(sessionID SessionRef, self VIFMetricsRef, key string) (_err error) {
	log.Println("VIFMetrics.RemoveFromOtherConfig not mocked")
	_err = errors.New("VIFMetrics.RemoveFromOtherConfig not mocked")
	return
}
// Remove the given key and its corresponding value from the other_config field of the given VIF_metrics.  If the key is not in that Map, then do nothing.
func (_class VIFMetricsClass) RemoveFromOtherConfig(sessionID SessionRef, self VIFMetricsRef, key string) (_err error) {
	if (IsMock) {
		return _class.RemoveFromOtherConfig__mock(sessionID, self, key)
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

func (_class VIFMetricsClass) AddToOtherConfig__mock(sessionID SessionRef, self VIFMetricsRef, key string, value string) (_err error) {
	log.Println("VIFMetrics.AddToOtherConfig not mocked")
	_err = errors.New("VIFMetrics.AddToOtherConfig not mocked")
	return
}
// Add the given key-value pair to the other_config field of the given VIF_metrics.
func (_class VIFMetricsClass) AddToOtherConfig(sessionID SessionRef, self VIFMetricsRef, key string, value string) (_err error) {
	if (IsMock) {
		return _class.AddToOtherConfig__mock(sessionID, self, key, value)
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

func (_class VIFMetricsClass) SetOtherConfig__mock(sessionID SessionRef, self VIFMetricsRef, value map[string]string) (_err error) {
	log.Println("VIFMetrics.SetOtherConfig not mocked")
	_err = errors.New("VIFMetrics.SetOtherConfig not mocked")
	return
}
// Set the other_config field of the given VIF_metrics.
func (_class VIFMetricsClass) SetOtherConfig(sessionID SessionRef, self VIFMetricsRef, value map[string]string) (_err error) {
	if (IsMock) {
		return _class.SetOtherConfig__mock(sessionID, self, value)
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

func (_class VIFMetricsClass) GetOtherConfig__mock(sessionID SessionRef, self VIFMetricsRef) (_retval map[string]string, _err error) {
	log.Println("VIFMetrics.GetOtherConfig not mocked")
	_err = errors.New("VIFMetrics.GetOtherConfig not mocked")
	return
}
// Get the other_config field of the given VIF_metrics.
func (_class VIFMetricsClass) GetOtherConfig(sessionID SessionRef, self VIFMetricsRef) (_retval map[string]string, _err error) {
	if (IsMock) {
		return _class.GetOtherConfig__mock(sessionID, self)
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

func (_class VIFMetricsClass) GetLastUpdated__mock(sessionID SessionRef, self VIFMetricsRef) (_retval time.Time, _err error) {
	log.Println("VIFMetrics.GetLastUpdated not mocked")
	_err = errors.New("VIFMetrics.GetLastUpdated not mocked")
	return
}
// Get the last_updated field of the given VIF_metrics.
func (_class VIFMetricsClass) GetLastUpdated(sessionID SessionRef, self VIFMetricsRef) (_retval time.Time, _err error) {
	if (IsMock) {
		return _class.GetLastUpdated__mock(sessionID, self)
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

func (_class VIFMetricsClass) GetIoWriteKbs__mock(sessionID SessionRef, self VIFMetricsRef) (_retval float64, _err error) {
	log.Println("VIFMetrics.GetIoWriteKbs not mocked")
	_err = errors.New("VIFMetrics.GetIoWriteKbs not mocked")
	return
}
// Get the io/write_kbs field of the given VIF_metrics.
func (_class VIFMetricsClass) GetIoWriteKbs(sessionID SessionRef, self VIFMetricsRef) (_retval float64, _err error) {
	if (IsMock) {
		return _class.GetIoWriteKbs__mock(sessionID, self)
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

func (_class VIFMetricsClass) GetIoReadKbs__mock(sessionID SessionRef, self VIFMetricsRef) (_retval float64, _err error) {
	log.Println("VIFMetrics.GetIoReadKbs not mocked")
	_err = errors.New("VIFMetrics.GetIoReadKbs not mocked")
	return
}
// Get the io/read_kbs field of the given VIF_metrics.
func (_class VIFMetricsClass) GetIoReadKbs(sessionID SessionRef, self VIFMetricsRef) (_retval float64, _err error) {
	if (IsMock) {
		return _class.GetIoReadKbs__mock(sessionID, self)
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

func (_class VIFMetricsClass) GetUUID__mock(sessionID SessionRef, self VIFMetricsRef) (_retval string, _err error) {
	log.Println("VIFMetrics.GetUUID not mocked")
	_err = errors.New("VIFMetrics.GetUUID not mocked")
	return
}
// Get the uuid field of the given VIF_metrics.
func (_class VIFMetricsClass) GetUUID(sessionID SessionRef, self VIFMetricsRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetUUID__mock(sessionID, self)
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

func (_class VIFMetricsClass) GetByUUID__mock(sessionID SessionRef, uuid string) (_retval VIFMetricsRef, _err error) {
	log.Println("VIFMetrics.GetByUUID not mocked")
	_err = errors.New("VIFMetrics.GetByUUID not mocked")
	return
}
// Get a reference to the VIF_metrics instance with the specified UUID.
func (_class VIFMetricsClass) GetByUUID(sessionID SessionRef, uuid string) (_retval VIFMetricsRef, _err error) {
	if (IsMock) {
		return _class.GetByUUID__mock(sessionID, uuid)
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

func (_class VIFMetricsClass) GetRecord__mock(sessionID SessionRef, self VIFMetricsRef) (_retval VIFMetricsRecord, _err error) {
	log.Println("VIFMetrics.GetRecord not mocked")
	_err = errors.New("VIFMetrics.GetRecord not mocked")
	return
}
// Get a record containing the current state of the given VIF_metrics.
func (_class VIFMetricsClass) GetRecord(sessionID SessionRef, self VIFMetricsRef) (_retval VIFMetricsRecord, _err error) {
	if (IsMock) {
		return _class.GetRecord__mock(sessionID, self)
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
