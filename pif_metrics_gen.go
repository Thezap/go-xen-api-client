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

type PIFMetricsRecord struct {
  // Unique identifier/object reference
	UUID string
  // Read bandwidth (KiB/s)
	IoReadKbs float64
  // Write bandwidth (KiB/s)
	IoWriteKbs float64
  // Report if the PIF got a carrier or not
	Carrier bool
  // Report vendor ID
	VendorID string
  // Report vendor name
	VendorName string
  // Report device ID
	DeviceID string
  // Report device name
	DeviceName string
  // Speed of the link (if available)
	Speed int
  // Full duplex capability of the link (if available)
	Duplex bool
  // PCI bus path of the pif (if available)
	PciBusPath string
  // Time at which this information was last updated
	LastUpdated time.Time
  // additional configuration
	OtherConfig map[string]string
}

type PIFMetricsRef string

// The metrics associated with a physical network interface
type PIFMetricsClass struct {
	client *Client
}

func (_class PIFMetricsClass) GetAllRecords__mock(sessionID SessionRef) (_retval map[PIFMetricsRef]PIFMetricsRecord, _err error) {
	log.Println("PIFMetrics.GetAllRecords not mocked")
	_err = errors.New("PIFMetrics.GetAllRecords not mocked")
	return
}
// Return a map of PIF_metrics references to PIF_metrics records for all PIF_metrics instances known to the system.
func (_class PIFMetricsClass) GetAllRecords(sessionID SessionRef) (_retval map[PIFMetricsRef]PIFMetricsRecord, _err error) {
	if (IsMock) {
		return _class.GetAllRecords__mock(sessionID)
	}	
	_method := "PIF_metrics.get_all_records"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertPIFMetricsRefToPIFMetricsRecordMapToGo(_method + " -> ", _result.Value)
	return
}

func (_class PIFMetricsClass) GetAll__mock(sessionID SessionRef) (_retval []PIFMetricsRef, _err error) {
	log.Println("PIFMetrics.GetAll not mocked")
	_err = errors.New("PIFMetrics.GetAll not mocked")
	return
}
// Return a list of all the PIF_metrics instances known to the system.
func (_class PIFMetricsClass) GetAll(sessionID SessionRef) (_retval []PIFMetricsRef, _err error) {
	if (IsMock) {
		return _class.GetAll__mock(sessionID)
	}	
	_method := "PIF_metrics.get_all"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertPIFMetricsRefSetToGo(_method + " -> ", _result.Value)
	return
}

func (_class PIFMetricsClass) RemoveFromOtherConfig__mock(sessionID SessionRef, self PIFMetricsRef, key string) (_err error) {
	log.Println("PIFMetrics.RemoveFromOtherConfig not mocked")
	_err = errors.New("PIFMetrics.RemoveFromOtherConfig not mocked")
	return
}
// Remove the given key and its corresponding value from the other_config field of the given PIF_metrics.  If the key is not in that Map, then do nothing.
func (_class PIFMetricsClass) RemoveFromOtherConfig(sessionID SessionRef, self PIFMetricsRef, key string) (_err error) {
	if (IsMock) {
		return _class.RemoveFromOtherConfig__mock(sessionID, self, key)
	}	
	_method := "PIF_metrics.remove_from_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPIFMetricsRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PIFMetricsClass) AddToOtherConfig__mock(sessionID SessionRef, self PIFMetricsRef, key string, value string) (_err error) {
	log.Println("PIFMetrics.AddToOtherConfig not mocked")
	_err = errors.New("PIFMetrics.AddToOtherConfig not mocked")
	return
}
// Add the given key-value pair to the other_config field of the given PIF_metrics.
func (_class PIFMetricsClass) AddToOtherConfig(sessionID SessionRef, self PIFMetricsRef, key string, value string) (_err error) {
	if (IsMock) {
		return _class.AddToOtherConfig__mock(sessionID, self, key, value)
	}	
	_method := "PIF_metrics.add_to_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPIFMetricsRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PIFMetricsClass) SetOtherConfig__mock(sessionID SessionRef, self PIFMetricsRef, value map[string]string) (_err error) {
	log.Println("PIFMetrics.SetOtherConfig not mocked")
	_err = errors.New("PIFMetrics.SetOtherConfig not mocked")
	return
}
// Set the other_config field of the given PIF_metrics.
func (_class PIFMetricsClass) SetOtherConfig(sessionID SessionRef, self PIFMetricsRef, value map[string]string) (_err error) {
	if (IsMock) {
		return _class.SetOtherConfig__mock(sessionID, self, value)
	}	
	_method := "PIF_metrics.set_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPIFMetricsRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PIFMetricsClass) GetOtherConfig__mock(sessionID SessionRef, self PIFMetricsRef) (_retval map[string]string, _err error) {
	log.Println("PIFMetrics.GetOtherConfig not mocked")
	_err = errors.New("PIFMetrics.GetOtherConfig not mocked")
	return
}
// Get the other_config field of the given PIF_metrics.
func (_class PIFMetricsClass) GetOtherConfig(sessionID SessionRef, self PIFMetricsRef) (_retval map[string]string, _err error) {
	if (IsMock) {
		return _class.GetOtherConfig__mock(sessionID, self)
	}	
	_method := "PIF_metrics.get_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPIFMetricsRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PIFMetricsClass) GetLastUpdated__mock(sessionID SessionRef, self PIFMetricsRef) (_retval time.Time, _err error) {
	log.Println("PIFMetrics.GetLastUpdated not mocked")
	_err = errors.New("PIFMetrics.GetLastUpdated not mocked")
	return
}
// Get the last_updated field of the given PIF_metrics.
func (_class PIFMetricsClass) GetLastUpdated(sessionID SessionRef, self PIFMetricsRef) (_retval time.Time, _err error) {
	if (IsMock) {
		return _class.GetLastUpdated__mock(sessionID, self)
	}	
	_method := "PIF_metrics.get_last_updated"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPIFMetricsRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PIFMetricsClass) GetPciBusPath__mock(sessionID SessionRef, self PIFMetricsRef) (_retval string, _err error) {
	log.Println("PIFMetrics.GetPciBusPath not mocked")
	_err = errors.New("PIFMetrics.GetPciBusPath not mocked")
	return
}
// Get the pci_bus_path field of the given PIF_metrics.
func (_class PIFMetricsClass) GetPciBusPath(sessionID SessionRef, self PIFMetricsRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetPciBusPath__mock(sessionID, self)
	}	
	_method := "PIF_metrics.get_pci_bus_path"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPIFMetricsRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PIFMetricsClass) GetDuplex__mock(sessionID SessionRef, self PIFMetricsRef) (_retval bool, _err error) {
	log.Println("PIFMetrics.GetDuplex not mocked")
	_err = errors.New("PIFMetrics.GetDuplex not mocked")
	return
}
// Get the duplex field of the given PIF_metrics.
func (_class PIFMetricsClass) GetDuplex(sessionID SessionRef, self PIFMetricsRef) (_retval bool, _err error) {
	if (IsMock) {
		return _class.GetDuplex__mock(sessionID, self)
	}	
	_method := "PIF_metrics.get_duplex"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPIFMetricsRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PIFMetricsClass) GetSpeed__mock(sessionID SessionRef, self PIFMetricsRef) (_retval int, _err error) {
	log.Println("PIFMetrics.GetSpeed not mocked")
	_err = errors.New("PIFMetrics.GetSpeed not mocked")
	return
}
// Get the speed field of the given PIF_metrics.
func (_class PIFMetricsClass) GetSpeed(sessionID SessionRef, self PIFMetricsRef) (_retval int, _err error) {
	if (IsMock) {
		return _class.GetSpeed__mock(sessionID, self)
	}	
	_method := "PIF_metrics.get_speed"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPIFMetricsRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PIFMetricsClass) GetDeviceName__mock(sessionID SessionRef, self PIFMetricsRef) (_retval string, _err error) {
	log.Println("PIFMetrics.GetDeviceName not mocked")
	_err = errors.New("PIFMetrics.GetDeviceName not mocked")
	return
}
// Get the device_name field of the given PIF_metrics.
func (_class PIFMetricsClass) GetDeviceName(sessionID SessionRef, self PIFMetricsRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetDeviceName__mock(sessionID, self)
	}	
	_method := "PIF_metrics.get_device_name"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPIFMetricsRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PIFMetricsClass) GetDeviceID__mock(sessionID SessionRef, self PIFMetricsRef) (_retval string, _err error) {
	log.Println("PIFMetrics.GetDeviceID not mocked")
	_err = errors.New("PIFMetrics.GetDeviceID not mocked")
	return
}
// Get the device_id field of the given PIF_metrics.
func (_class PIFMetricsClass) GetDeviceID(sessionID SessionRef, self PIFMetricsRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetDeviceID__mock(sessionID, self)
	}	
	_method := "PIF_metrics.get_device_id"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPIFMetricsRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PIFMetricsClass) GetVendorName__mock(sessionID SessionRef, self PIFMetricsRef) (_retval string, _err error) {
	log.Println("PIFMetrics.GetVendorName not mocked")
	_err = errors.New("PIFMetrics.GetVendorName not mocked")
	return
}
// Get the vendor_name field of the given PIF_metrics.
func (_class PIFMetricsClass) GetVendorName(sessionID SessionRef, self PIFMetricsRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetVendorName__mock(sessionID, self)
	}	
	_method := "PIF_metrics.get_vendor_name"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPIFMetricsRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PIFMetricsClass) GetVendorID__mock(sessionID SessionRef, self PIFMetricsRef) (_retval string, _err error) {
	log.Println("PIFMetrics.GetVendorID not mocked")
	_err = errors.New("PIFMetrics.GetVendorID not mocked")
	return
}
// Get the vendor_id field of the given PIF_metrics.
func (_class PIFMetricsClass) GetVendorID(sessionID SessionRef, self PIFMetricsRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetVendorID__mock(sessionID, self)
	}	
	_method := "PIF_metrics.get_vendor_id"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPIFMetricsRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PIFMetricsClass) GetCarrier__mock(sessionID SessionRef, self PIFMetricsRef) (_retval bool, _err error) {
	log.Println("PIFMetrics.GetCarrier not mocked")
	_err = errors.New("PIFMetrics.GetCarrier not mocked")
	return
}
// Get the carrier field of the given PIF_metrics.
func (_class PIFMetricsClass) GetCarrier(sessionID SessionRef, self PIFMetricsRef) (_retval bool, _err error) {
	if (IsMock) {
		return _class.GetCarrier__mock(sessionID, self)
	}	
	_method := "PIF_metrics.get_carrier"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPIFMetricsRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PIFMetricsClass) GetIoWriteKbs__mock(sessionID SessionRef, self PIFMetricsRef) (_retval float64, _err error) {
	log.Println("PIFMetrics.GetIoWriteKbs not mocked")
	_err = errors.New("PIFMetrics.GetIoWriteKbs not mocked")
	return
}
// Get the io/write_kbs field of the given PIF_metrics.
func (_class PIFMetricsClass) GetIoWriteKbs(sessionID SessionRef, self PIFMetricsRef) (_retval float64, _err error) {
	if (IsMock) {
		return _class.GetIoWriteKbs__mock(sessionID, self)
	}	
	_method := "PIF_metrics.get_io_write_kbs"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPIFMetricsRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PIFMetricsClass) GetIoReadKbs__mock(sessionID SessionRef, self PIFMetricsRef) (_retval float64, _err error) {
	log.Println("PIFMetrics.GetIoReadKbs not mocked")
	_err = errors.New("PIFMetrics.GetIoReadKbs not mocked")
	return
}
// Get the io/read_kbs field of the given PIF_metrics.
func (_class PIFMetricsClass) GetIoReadKbs(sessionID SessionRef, self PIFMetricsRef) (_retval float64, _err error) {
	if (IsMock) {
		return _class.GetIoReadKbs__mock(sessionID, self)
	}	
	_method := "PIF_metrics.get_io_read_kbs"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPIFMetricsRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PIFMetricsClass) GetUUID__mock(sessionID SessionRef, self PIFMetricsRef) (_retval string, _err error) {
	log.Println("PIFMetrics.GetUUID not mocked")
	_err = errors.New("PIFMetrics.GetUUID not mocked")
	return
}
// Get the uuid field of the given PIF_metrics.
func (_class PIFMetricsClass) GetUUID(sessionID SessionRef, self PIFMetricsRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetUUID__mock(sessionID, self)
	}	
	_method := "PIF_metrics.get_uuid"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPIFMetricsRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PIFMetricsClass) GetByUUID__mock(sessionID SessionRef, uuid string) (_retval PIFMetricsRef, _err error) {
	log.Println("PIFMetrics.GetByUUID not mocked")
	_err = errors.New("PIFMetrics.GetByUUID not mocked")
	return
}
// Get a reference to the PIF_metrics instance with the specified UUID.
func (_class PIFMetricsClass) GetByUUID(sessionID SessionRef, uuid string) (_retval PIFMetricsRef, _err error) {
	if (IsMock) {
		return _class.GetByUUID__mock(sessionID, uuid)
	}	
	_method := "PIF_metrics.get_by_uuid"
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
	_retval, _err = convertPIFMetricsRefToGo(_method + " -> ", _result.Value)
	return
}

func (_class PIFMetricsClass) GetRecord__mock(sessionID SessionRef, self PIFMetricsRef) (_retval PIFMetricsRecord, _err error) {
	log.Println("PIFMetrics.GetRecord not mocked")
	_err = errors.New("PIFMetrics.GetRecord not mocked")
	return
}
// Get a record containing the current state of the given PIF_metrics.
func (_class PIFMetricsClass) GetRecord(sessionID SessionRef, self PIFMetricsRef) (_retval PIFMetricsRecord, _err error) {
	if (IsMock) {
		return _class.GetRecord__mock(sessionID, self)
	}	
	_method := "PIF_metrics.get_record"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPIFMetricsRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertPIFMetricsRecordToGo(_method + " -> ", _result.Value)
	return
}
