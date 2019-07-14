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


var PIFMetricsClass_GetAllRecordsMockedCallback = func (sessionID SessionRef) (_retval map[PIFMetricsRef]PIFMetricsRecord, _err error) {
	log.Println("PIFMetrics.GetAllRecords not mocked")
	_err = errors.New("PIFMetrics.GetAllRecords not mocked")
	return
}

func (_class PIFMetricsClass) GetAllRecordsMock(sessionID SessionRef) (_retval map[PIFMetricsRef]PIFMetricsRecord, _err error) {
	return PIFMetricsClass_GetAllRecordsMockedCallback(sessionID)
}
// Return a map of PIF_metrics references to PIF_metrics records for all PIF_metrics instances known to the system.
func (_class PIFMetricsClass) GetAllRecords(sessionID SessionRef) (_retval map[PIFMetricsRef]PIFMetricsRecord, _err error) {
	if (IsMock) {
		return _class.GetAllRecordsMock(sessionID)
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


var PIFMetricsClass_GetAllMockedCallback = func (sessionID SessionRef) (_retval []PIFMetricsRef, _err error) {
	log.Println("PIFMetrics.GetAll not mocked")
	_err = errors.New("PIFMetrics.GetAll not mocked")
	return
}

func (_class PIFMetricsClass) GetAllMock(sessionID SessionRef) (_retval []PIFMetricsRef, _err error) {
	return PIFMetricsClass_GetAllMockedCallback(sessionID)
}
// Return a list of all the PIF_metrics instances known to the system.
func (_class PIFMetricsClass) GetAll(sessionID SessionRef) (_retval []PIFMetricsRef, _err error) {
	if (IsMock) {
		return _class.GetAllMock(sessionID)
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


var PIFMetricsClass_RemoveFromOtherConfigMockedCallback = func (sessionID SessionRef, self PIFMetricsRef, key string) (_err error) {
	log.Println("PIFMetrics.RemoveFromOtherConfig not mocked")
	_err = errors.New("PIFMetrics.RemoveFromOtherConfig not mocked")
	return
}

func (_class PIFMetricsClass) RemoveFromOtherConfigMock(sessionID SessionRef, self PIFMetricsRef, key string) (_err error) {
	return PIFMetricsClass_RemoveFromOtherConfigMockedCallback(sessionID, self, key)
}
// Remove the given key and its corresponding value from the other_config field of the given PIF_metrics.  If the key is not in that Map, then do nothing.
func (_class PIFMetricsClass) RemoveFromOtherConfig(sessionID SessionRef, self PIFMetricsRef, key string) (_err error) {
	if (IsMock) {
		return _class.RemoveFromOtherConfigMock(sessionID, self, key)
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


var PIFMetricsClass_AddToOtherConfigMockedCallback = func (sessionID SessionRef, self PIFMetricsRef, key string, value string) (_err error) {
	log.Println("PIFMetrics.AddToOtherConfig not mocked")
	_err = errors.New("PIFMetrics.AddToOtherConfig not mocked")
	return
}

func (_class PIFMetricsClass) AddToOtherConfigMock(sessionID SessionRef, self PIFMetricsRef, key string, value string) (_err error) {
	return PIFMetricsClass_AddToOtherConfigMockedCallback(sessionID, self, key, value)
}
// Add the given key-value pair to the other_config field of the given PIF_metrics.
func (_class PIFMetricsClass) AddToOtherConfig(sessionID SessionRef, self PIFMetricsRef, key string, value string) (_err error) {
	if (IsMock) {
		return _class.AddToOtherConfigMock(sessionID, self, key, value)
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


var PIFMetricsClass_SetOtherConfigMockedCallback = func (sessionID SessionRef, self PIFMetricsRef, value map[string]string) (_err error) {
	log.Println("PIFMetrics.SetOtherConfig not mocked")
	_err = errors.New("PIFMetrics.SetOtherConfig not mocked")
	return
}

func (_class PIFMetricsClass) SetOtherConfigMock(sessionID SessionRef, self PIFMetricsRef, value map[string]string) (_err error) {
	return PIFMetricsClass_SetOtherConfigMockedCallback(sessionID, self, value)
}
// Set the other_config field of the given PIF_metrics.
func (_class PIFMetricsClass) SetOtherConfig(sessionID SessionRef, self PIFMetricsRef, value map[string]string) (_err error) {
	if (IsMock) {
		return _class.SetOtherConfigMock(sessionID, self, value)
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


var PIFMetricsClass_GetOtherConfigMockedCallback = func (sessionID SessionRef, self PIFMetricsRef) (_retval map[string]string, _err error) {
	log.Println("PIFMetrics.GetOtherConfig not mocked")
	_err = errors.New("PIFMetrics.GetOtherConfig not mocked")
	return
}

func (_class PIFMetricsClass) GetOtherConfigMock(sessionID SessionRef, self PIFMetricsRef) (_retval map[string]string, _err error) {
	return PIFMetricsClass_GetOtherConfigMockedCallback(sessionID, self)
}
// Get the other_config field of the given PIF_metrics.
func (_class PIFMetricsClass) GetOtherConfig(sessionID SessionRef, self PIFMetricsRef) (_retval map[string]string, _err error) {
	if (IsMock) {
		return _class.GetOtherConfigMock(sessionID, self)
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


var PIFMetricsClass_GetLastUpdatedMockedCallback = func (sessionID SessionRef, self PIFMetricsRef) (_retval time.Time, _err error) {
	log.Println("PIFMetrics.GetLastUpdated not mocked")
	_err = errors.New("PIFMetrics.GetLastUpdated not mocked")
	return
}

func (_class PIFMetricsClass) GetLastUpdatedMock(sessionID SessionRef, self PIFMetricsRef) (_retval time.Time, _err error) {
	return PIFMetricsClass_GetLastUpdatedMockedCallback(sessionID, self)
}
// Get the last_updated field of the given PIF_metrics.
func (_class PIFMetricsClass) GetLastUpdated(sessionID SessionRef, self PIFMetricsRef) (_retval time.Time, _err error) {
	if (IsMock) {
		return _class.GetLastUpdatedMock(sessionID, self)
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


var PIFMetricsClass_GetPciBusPathMockedCallback = func (sessionID SessionRef, self PIFMetricsRef) (_retval string, _err error) {
	log.Println("PIFMetrics.GetPciBusPath not mocked")
	_err = errors.New("PIFMetrics.GetPciBusPath not mocked")
	return
}

func (_class PIFMetricsClass) GetPciBusPathMock(sessionID SessionRef, self PIFMetricsRef) (_retval string, _err error) {
	return PIFMetricsClass_GetPciBusPathMockedCallback(sessionID, self)
}
// Get the pci_bus_path field of the given PIF_metrics.
func (_class PIFMetricsClass) GetPciBusPath(sessionID SessionRef, self PIFMetricsRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetPciBusPathMock(sessionID, self)
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


var PIFMetricsClass_GetDuplexMockedCallback = func (sessionID SessionRef, self PIFMetricsRef) (_retval bool, _err error) {
	log.Println("PIFMetrics.GetDuplex not mocked")
	_err = errors.New("PIFMetrics.GetDuplex not mocked")
	return
}

func (_class PIFMetricsClass) GetDuplexMock(sessionID SessionRef, self PIFMetricsRef) (_retval bool, _err error) {
	return PIFMetricsClass_GetDuplexMockedCallback(sessionID, self)
}
// Get the duplex field of the given PIF_metrics.
func (_class PIFMetricsClass) GetDuplex(sessionID SessionRef, self PIFMetricsRef) (_retval bool, _err error) {
	if (IsMock) {
		return _class.GetDuplexMock(sessionID, self)
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


var PIFMetricsClass_GetSpeedMockedCallback = func (sessionID SessionRef, self PIFMetricsRef) (_retval int, _err error) {
	log.Println("PIFMetrics.GetSpeed not mocked")
	_err = errors.New("PIFMetrics.GetSpeed not mocked")
	return
}

func (_class PIFMetricsClass) GetSpeedMock(sessionID SessionRef, self PIFMetricsRef) (_retval int, _err error) {
	return PIFMetricsClass_GetSpeedMockedCallback(sessionID, self)
}
// Get the speed field of the given PIF_metrics.
func (_class PIFMetricsClass) GetSpeed(sessionID SessionRef, self PIFMetricsRef) (_retval int, _err error) {
	if (IsMock) {
		return _class.GetSpeedMock(sessionID, self)
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


var PIFMetricsClass_GetDeviceNameMockedCallback = func (sessionID SessionRef, self PIFMetricsRef) (_retval string, _err error) {
	log.Println("PIFMetrics.GetDeviceName not mocked")
	_err = errors.New("PIFMetrics.GetDeviceName not mocked")
	return
}

func (_class PIFMetricsClass) GetDeviceNameMock(sessionID SessionRef, self PIFMetricsRef) (_retval string, _err error) {
	return PIFMetricsClass_GetDeviceNameMockedCallback(sessionID, self)
}
// Get the device_name field of the given PIF_metrics.
func (_class PIFMetricsClass) GetDeviceName(sessionID SessionRef, self PIFMetricsRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetDeviceNameMock(sessionID, self)
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


var PIFMetricsClass_GetDeviceIDMockedCallback = func (sessionID SessionRef, self PIFMetricsRef) (_retval string, _err error) {
	log.Println("PIFMetrics.GetDeviceID not mocked")
	_err = errors.New("PIFMetrics.GetDeviceID not mocked")
	return
}

func (_class PIFMetricsClass) GetDeviceIDMock(sessionID SessionRef, self PIFMetricsRef) (_retval string, _err error) {
	return PIFMetricsClass_GetDeviceIDMockedCallback(sessionID, self)
}
// Get the device_id field of the given PIF_metrics.
func (_class PIFMetricsClass) GetDeviceID(sessionID SessionRef, self PIFMetricsRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetDeviceIDMock(sessionID, self)
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


var PIFMetricsClass_GetVendorNameMockedCallback = func (sessionID SessionRef, self PIFMetricsRef) (_retval string, _err error) {
	log.Println("PIFMetrics.GetVendorName not mocked")
	_err = errors.New("PIFMetrics.GetVendorName not mocked")
	return
}

func (_class PIFMetricsClass) GetVendorNameMock(sessionID SessionRef, self PIFMetricsRef) (_retval string, _err error) {
	return PIFMetricsClass_GetVendorNameMockedCallback(sessionID, self)
}
// Get the vendor_name field of the given PIF_metrics.
func (_class PIFMetricsClass) GetVendorName(sessionID SessionRef, self PIFMetricsRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetVendorNameMock(sessionID, self)
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


var PIFMetricsClass_GetVendorIDMockedCallback = func (sessionID SessionRef, self PIFMetricsRef) (_retval string, _err error) {
	log.Println("PIFMetrics.GetVendorID not mocked")
	_err = errors.New("PIFMetrics.GetVendorID not mocked")
	return
}

func (_class PIFMetricsClass) GetVendorIDMock(sessionID SessionRef, self PIFMetricsRef) (_retval string, _err error) {
	return PIFMetricsClass_GetVendorIDMockedCallback(sessionID, self)
}
// Get the vendor_id field of the given PIF_metrics.
func (_class PIFMetricsClass) GetVendorID(sessionID SessionRef, self PIFMetricsRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetVendorIDMock(sessionID, self)
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


var PIFMetricsClass_GetCarrierMockedCallback = func (sessionID SessionRef, self PIFMetricsRef) (_retval bool, _err error) {
	log.Println("PIFMetrics.GetCarrier not mocked")
	_err = errors.New("PIFMetrics.GetCarrier not mocked")
	return
}

func (_class PIFMetricsClass) GetCarrierMock(sessionID SessionRef, self PIFMetricsRef) (_retval bool, _err error) {
	return PIFMetricsClass_GetCarrierMockedCallback(sessionID, self)
}
// Get the carrier field of the given PIF_metrics.
func (_class PIFMetricsClass) GetCarrier(sessionID SessionRef, self PIFMetricsRef) (_retval bool, _err error) {
	if (IsMock) {
		return _class.GetCarrierMock(sessionID, self)
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


var PIFMetricsClass_GetIoWriteKbsMockedCallback = func (sessionID SessionRef, self PIFMetricsRef) (_retval float64, _err error) {
	log.Println("PIFMetrics.GetIoWriteKbs not mocked")
	_err = errors.New("PIFMetrics.GetIoWriteKbs not mocked")
	return
}

func (_class PIFMetricsClass) GetIoWriteKbsMock(sessionID SessionRef, self PIFMetricsRef) (_retval float64, _err error) {
	return PIFMetricsClass_GetIoWriteKbsMockedCallback(sessionID, self)
}
// Get the io/write_kbs field of the given PIF_metrics.
func (_class PIFMetricsClass) GetIoWriteKbs(sessionID SessionRef, self PIFMetricsRef) (_retval float64, _err error) {
	if (IsMock) {
		return _class.GetIoWriteKbsMock(sessionID, self)
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


var PIFMetricsClass_GetIoReadKbsMockedCallback = func (sessionID SessionRef, self PIFMetricsRef) (_retval float64, _err error) {
	log.Println("PIFMetrics.GetIoReadKbs not mocked")
	_err = errors.New("PIFMetrics.GetIoReadKbs not mocked")
	return
}

func (_class PIFMetricsClass) GetIoReadKbsMock(sessionID SessionRef, self PIFMetricsRef) (_retval float64, _err error) {
	return PIFMetricsClass_GetIoReadKbsMockedCallback(sessionID, self)
}
// Get the io/read_kbs field of the given PIF_metrics.
func (_class PIFMetricsClass) GetIoReadKbs(sessionID SessionRef, self PIFMetricsRef) (_retval float64, _err error) {
	if (IsMock) {
		return _class.GetIoReadKbsMock(sessionID, self)
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


var PIFMetricsClass_GetUUIDMockedCallback = func (sessionID SessionRef, self PIFMetricsRef) (_retval string, _err error) {
	log.Println("PIFMetrics.GetUUID not mocked")
	_err = errors.New("PIFMetrics.GetUUID not mocked")
	return
}

func (_class PIFMetricsClass) GetUUIDMock(sessionID SessionRef, self PIFMetricsRef) (_retval string, _err error) {
	return PIFMetricsClass_GetUUIDMockedCallback(sessionID, self)
}
// Get the uuid field of the given PIF_metrics.
func (_class PIFMetricsClass) GetUUID(sessionID SessionRef, self PIFMetricsRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetUUIDMock(sessionID, self)
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


var PIFMetricsClass_GetByUUIDMockedCallback = func (sessionID SessionRef, uuid string) (_retval PIFMetricsRef, _err error) {
	log.Println("PIFMetrics.GetByUUID not mocked")
	_err = errors.New("PIFMetrics.GetByUUID not mocked")
	return
}

func (_class PIFMetricsClass) GetByUUIDMock(sessionID SessionRef, uuid string) (_retval PIFMetricsRef, _err error) {
	return PIFMetricsClass_GetByUUIDMockedCallback(sessionID, uuid)
}
// Get a reference to the PIF_metrics instance with the specified UUID.
func (_class PIFMetricsClass) GetByUUID(sessionID SessionRef, uuid string) (_retval PIFMetricsRef, _err error) {
	if (IsMock) {
		return _class.GetByUUIDMock(sessionID, uuid)
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


var PIFMetricsClass_GetRecordMockedCallback = func (sessionID SessionRef, self PIFMetricsRef) (_retval PIFMetricsRecord, _err error) {
	log.Println("PIFMetrics.GetRecord not mocked")
	_err = errors.New("PIFMetrics.GetRecord not mocked")
	return
}

func (_class PIFMetricsClass) GetRecordMock(sessionID SessionRef, self PIFMetricsRef) (_retval PIFMetricsRecord, _err error) {
	return PIFMetricsClass_GetRecordMockedCallback(sessionID, self)
}
// Get a record containing the current state of the given PIF_metrics.
func (_class PIFMetricsClass) GetRecord(sessionID SessionRef, self PIFMetricsRef) (_retval PIFMetricsRecord, _err error) {
	if (IsMock) {
		return _class.GetRecordMock(sessionID, self)
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
