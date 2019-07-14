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

type PUSBRecord struct {
  // Unique identifier/object reference
	UUID string
  // USB group the PUSB is contained in
	USBGroup USBGroupRef
  // Physical machine that owns the USB device
	Host HostRef
  // port path of USB device
	Path string
  // vendor id of the USB device
	VendorID string
  // vendor description of the USB device
	VendorDesc string
  // product id of the USB device
	ProductID string
  // product description of the USB device
	ProductDesc string
  // serial of the USB device
	Serial string
  // USB device version
	Version string
  // USB device description
	Description string
  // enabled for passthrough
	PassthroughEnabled bool
  // additional configuration
	OtherConfig map[string]string
}

type PUSBRef string

// A physical USB device
type PUSBClass struct {
	client *Client
}


func PUSBClassGetAllRecordsMockDefault(sessionID SessionRef) (_retval map[PUSBRef]PUSBRecord, _err error) {
	log.Println("PUSB.GetAllRecords not mocked")
	_err = errors.New("PUSB.GetAllRecords not mocked")
	return
}

var PUSBClassGetAllRecordsMockedCallback = PUSBClassGetAllRecordsMockDefault

func (_class PUSBClass) GetAllRecordsMock(sessionID SessionRef) (_retval map[PUSBRef]PUSBRecord, _err error) {
	return PUSBClassGetAllRecordsMockedCallback(sessionID)
}
// Return a map of PUSB references to PUSB records for all PUSBs known to the system.
func (_class PUSBClass) GetAllRecords(sessionID SessionRef) (_retval map[PUSBRef]PUSBRecord, _err error) {
	if IsMock {
		return _class.GetAllRecordsMock(sessionID)
	}	
	_method := "PUSB.get_all_records"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertPUSBRefToPUSBRecordMapToGo(_method + " -> ", _result.Value)
	return
}


func PUSBClassGetAllMockDefault(sessionID SessionRef) (_retval []PUSBRef, _err error) {
	log.Println("PUSB.GetAll not mocked")
	_err = errors.New("PUSB.GetAll not mocked")
	return
}

var PUSBClassGetAllMockedCallback = PUSBClassGetAllMockDefault

func (_class PUSBClass) GetAllMock(sessionID SessionRef) (_retval []PUSBRef, _err error) {
	return PUSBClassGetAllMockedCallback(sessionID)
}
// Return a list of all the PUSBs known to the system.
func (_class PUSBClass) GetAll(sessionID SessionRef) (_retval []PUSBRef, _err error) {
	if IsMock {
		return _class.GetAllMock(sessionID)
	}	
	_method := "PUSB.get_all"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertPUSBRefSetToGo(_method + " -> ", _result.Value)
	return
}


func PUSBClassSetPassthroughEnabledMockDefault(sessionID SessionRef, self PUSBRef, value bool) (_err error) {
	log.Println("PUSB.SetPassthroughEnabled not mocked")
	_err = errors.New("PUSB.SetPassthroughEnabled not mocked")
	return
}

var PUSBClassSetPassthroughEnabledMockedCallback = PUSBClassSetPassthroughEnabledMockDefault

func (_class PUSBClass) SetPassthroughEnabledMock(sessionID SessionRef, self PUSBRef, value bool) (_err error) {
	return PUSBClassSetPassthroughEnabledMockedCallback(sessionID, self, value)
}
// 
func (_class PUSBClass) SetPassthroughEnabled(sessionID SessionRef, self PUSBRef, value bool) (_err error) {
	if IsMock {
		return _class.SetPassthroughEnabledMock(sessionID, self, value)
	}	
	_method := "PUSB.set_passthrough_enabled"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPUSBRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertBoolToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}


func PUSBClassScanMockDefault(sessionID SessionRef, host HostRef) (_err error) {
	log.Println("PUSB.Scan not mocked")
	_err = errors.New("PUSB.Scan not mocked")
	return
}

var PUSBClassScanMockedCallback = PUSBClassScanMockDefault

func (_class PUSBClass) ScanMock(sessionID SessionRef, host HostRef) (_err error) {
	return PUSBClassScanMockedCallback(sessionID, host)
}
// 
func (_class PUSBClass) Scan(sessionID SessionRef, host HostRef) (_err error) {
	if IsMock {
		return _class.ScanMock(sessionID, host)
	}	
	_method := "PUSB.scan"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_hostArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "host"), host)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _hostArg)
	return
}


func PUSBClassRemoveFromOtherConfigMockDefault(sessionID SessionRef, self PUSBRef, key string) (_err error) {
	log.Println("PUSB.RemoveFromOtherConfig not mocked")
	_err = errors.New("PUSB.RemoveFromOtherConfig not mocked")
	return
}

var PUSBClassRemoveFromOtherConfigMockedCallback = PUSBClassRemoveFromOtherConfigMockDefault

func (_class PUSBClass) RemoveFromOtherConfigMock(sessionID SessionRef, self PUSBRef, key string) (_err error) {
	return PUSBClassRemoveFromOtherConfigMockedCallback(sessionID, self, key)
}
// Remove the given key and its corresponding value from the other_config field of the given PUSB.  If the key is not in that Map, then do nothing.
func (_class PUSBClass) RemoveFromOtherConfig(sessionID SessionRef, self PUSBRef, key string) (_err error) {
	if IsMock {
		return _class.RemoveFromOtherConfigMock(sessionID, self, key)
	}	
	_method := "PUSB.remove_from_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPUSBRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func PUSBClassAddToOtherConfigMockDefault(sessionID SessionRef, self PUSBRef, key string, value string) (_err error) {
	log.Println("PUSB.AddToOtherConfig not mocked")
	_err = errors.New("PUSB.AddToOtherConfig not mocked")
	return
}

var PUSBClassAddToOtherConfigMockedCallback = PUSBClassAddToOtherConfigMockDefault

func (_class PUSBClass) AddToOtherConfigMock(sessionID SessionRef, self PUSBRef, key string, value string) (_err error) {
	return PUSBClassAddToOtherConfigMockedCallback(sessionID, self, key, value)
}
// Add the given key-value pair to the other_config field of the given PUSB.
func (_class PUSBClass) AddToOtherConfig(sessionID SessionRef, self PUSBRef, key string, value string) (_err error) {
	if IsMock {
		return _class.AddToOtherConfigMock(sessionID, self, key, value)
	}	
	_method := "PUSB.add_to_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPUSBRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func PUSBClassSetOtherConfigMockDefault(sessionID SessionRef, self PUSBRef, value map[string]string) (_err error) {
	log.Println("PUSB.SetOtherConfig not mocked")
	_err = errors.New("PUSB.SetOtherConfig not mocked")
	return
}

var PUSBClassSetOtherConfigMockedCallback = PUSBClassSetOtherConfigMockDefault

func (_class PUSBClass) SetOtherConfigMock(sessionID SessionRef, self PUSBRef, value map[string]string) (_err error) {
	return PUSBClassSetOtherConfigMockedCallback(sessionID, self, value)
}
// Set the other_config field of the given PUSB.
func (_class PUSBClass) SetOtherConfig(sessionID SessionRef, self PUSBRef, value map[string]string) (_err error) {
	if IsMock {
		return _class.SetOtherConfigMock(sessionID, self, value)
	}	
	_method := "PUSB.set_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPUSBRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func PUSBClassGetOtherConfigMockDefault(sessionID SessionRef, self PUSBRef) (_retval map[string]string, _err error) {
	log.Println("PUSB.GetOtherConfig not mocked")
	_err = errors.New("PUSB.GetOtherConfig not mocked")
	return
}

var PUSBClassGetOtherConfigMockedCallback = PUSBClassGetOtherConfigMockDefault

func (_class PUSBClass) GetOtherConfigMock(sessionID SessionRef, self PUSBRef) (_retval map[string]string, _err error) {
	return PUSBClassGetOtherConfigMockedCallback(sessionID, self)
}
// Get the other_config field of the given PUSB.
func (_class PUSBClass) GetOtherConfig(sessionID SessionRef, self PUSBRef) (_retval map[string]string, _err error) {
	if IsMock {
		return _class.GetOtherConfigMock(sessionID, self)
	}	
	_method := "PUSB.get_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPUSBRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func PUSBClassGetPassthroughEnabledMockDefault(sessionID SessionRef, self PUSBRef) (_retval bool, _err error) {
	log.Println("PUSB.GetPassthroughEnabled not mocked")
	_err = errors.New("PUSB.GetPassthroughEnabled not mocked")
	return
}

var PUSBClassGetPassthroughEnabledMockedCallback = PUSBClassGetPassthroughEnabledMockDefault

func (_class PUSBClass) GetPassthroughEnabledMock(sessionID SessionRef, self PUSBRef) (_retval bool, _err error) {
	return PUSBClassGetPassthroughEnabledMockedCallback(sessionID, self)
}
// Get the passthrough_enabled field of the given PUSB.
func (_class PUSBClass) GetPassthroughEnabled(sessionID SessionRef, self PUSBRef) (_retval bool, _err error) {
	if IsMock {
		return _class.GetPassthroughEnabledMock(sessionID, self)
	}	
	_method := "PUSB.get_passthrough_enabled"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPUSBRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func PUSBClassGetDescriptionMockDefault(sessionID SessionRef, self PUSBRef) (_retval string, _err error) {
	log.Println("PUSB.GetDescription not mocked")
	_err = errors.New("PUSB.GetDescription not mocked")
	return
}

var PUSBClassGetDescriptionMockedCallback = PUSBClassGetDescriptionMockDefault

func (_class PUSBClass) GetDescriptionMock(sessionID SessionRef, self PUSBRef) (_retval string, _err error) {
	return PUSBClassGetDescriptionMockedCallback(sessionID, self)
}
// Get the description field of the given PUSB.
func (_class PUSBClass) GetDescription(sessionID SessionRef, self PUSBRef) (_retval string, _err error) {
	if IsMock {
		return _class.GetDescriptionMock(sessionID, self)
	}	
	_method := "PUSB.get_description"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPUSBRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func PUSBClassGetVersionMockDefault(sessionID SessionRef, self PUSBRef) (_retval string, _err error) {
	log.Println("PUSB.GetVersion not mocked")
	_err = errors.New("PUSB.GetVersion not mocked")
	return
}

var PUSBClassGetVersionMockedCallback = PUSBClassGetVersionMockDefault

func (_class PUSBClass) GetVersionMock(sessionID SessionRef, self PUSBRef) (_retval string, _err error) {
	return PUSBClassGetVersionMockedCallback(sessionID, self)
}
// Get the version field of the given PUSB.
func (_class PUSBClass) GetVersion(sessionID SessionRef, self PUSBRef) (_retval string, _err error) {
	if IsMock {
		return _class.GetVersionMock(sessionID, self)
	}	
	_method := "PUSB.get_version"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPUSBRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func PUSBClassGetSerialMockDefault(sessionID SessionRef, self PUSBRef) (_retval string, _err error) {
	log.Println("PUSB.GetSerial not mocked")
	_err = errors.New("PUSB.GetSerial not mocked")
	return
}

var PUSBClassGetSerialMockedCallback = PUSBClassGetSerialMockDefault

func (_class PUSBClass) GetSerialMock(sessionID SessionRef, self PUSBRef) (_retval string, _err error) {
	return PUSBClassGetSerialMockedCallback(sessionID, self)
}
// Get the serial field of the given PUSB.
func (_class PUSBClass) GetSerial(sessionID SessionRef, self PUSBRef) (_retval string, _err error) {
	if IsMock {
		return _class.GetSerialMock(sessionID, self)
	}	
	_method := "PUSB.get_serial"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPUSBRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func PUSBClassGetProductDescMockDefault(sessionID SessionRef, self PUSBRef) (_retval string, _err error) {
	log.Println("PUSB.GetProductDesc not mocked")
	_err = errors.New("PUSB.GetProductDesc not mocked")
	return
}

var PUSBClassGetProductDescMockedCallback = PUSBClassGetProductDescMockDefault

func (_class PUSBClass) GetProductDescMock(sessionID SessionRef, self PUSBRef) (_retval string, _err error) {
	return PUSBClassGetProductDescMockedCallback(sessionID, self)
}
// Get the product_desc field of the given PUSB.
func (_class PUSBClass) GetProductDesc(sessionID SessionRef, self PUSBRef) (_retval string, _err error) {
	if IsMock {
		return _class.GetProductDescMock(sessionID, self)
	}	
	_method := "PUSB.get_product_desc"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPUSBRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func PUSBClassGetProductIDMockDefault(sessionID SessionRef, self PUSBRef) (_retval string, _err error) {
	log.Println("PUSB.GetProductID not mocked")
	_err = errors.New("PUSB.GetProductID not mocked")
	return
}

var PUSBClassGetProductIDMockedCallback = PUSBClassGetProductIDMockDefault

func (_class PUSBClass) GetProductIDMock(sessionID SessionRef, self PUSBRef) (_retval string, _err error) {
	return PUSBClassGetProductIDMockedCallback(sessionID, self)
}
// Get the product_id field of the given PUSB.
func (_class PUSBClass) GetProductID(sessionID SessionRef, self PUSBRef) (_retval string, _err error) {
	if IsMock {
		return _class.GetProductIDMock(sessionID, self)
	}	
	_method := "PUSB.get_product_id"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPUSBRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func PUSBClassGetVendorDescMockDefault(sessionID SessionRef, self PUSBRef) (_retval string, _err error) {
	log.Println("PUSB.GetVendorDesc not mocked")
	_err = errors.New("PUSB.GetVendorDesc not mocked")
	return
}

var PUSBClassGetVendorDescMockedCallback = PUSBClassGetVendorDescMockDefault

func (_class PUSBClass) GetVendorDescMock(sessionID SessionRef, self PUSBRef) (_retval string, _err error) {
	return PUSBClassGetVendorDescMockedCallback(sessionID, self)
}
// Get the vendor_desc field of the given PUSB.
func (_class PUSBClass) GetVendorDesc(sessionID SessionRef, self PUSBRef) (_retval string, _err error) {
	if IsMock {
		return _class.GetVendorDescMock(sessionID, self)
	}	
	_method := "PUSB.get_vendor_desc"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPUSBRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func PUSBClassGetVendorIDMockDefault(sessionID SessionRef, self PUSBRef) (_retval string, _err error) {
	log.Println("PUSB.GetVendorID not mocked")
	_err = errors.New("PUSB.GetVendorID not mocked")
	return
}

var PUSBClassGetVendorIDMockedCallback = PUSBClassGetVendorIDMockDefault

func (_class PUSBClass) GetVendorIDMock(sessionID SessionRef, self PUSBRef) (_retval string, _err error) {
	return PUSBClassGetVendorIDMockedCallback(sessionID, self)
}
// Get the vendor_id field of the given PUSB.
func (_class PUSBClass) GetVendorID(sessionID SessionRef, self PUSBRef) (_retval string, _err error) {
	if IsMock {
		return _class.GetVendorIDMock(sessionID, self)
	}	
	_method := "PUSB.get_vendor_id"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPUSBRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func PUSBClassGetPathMockDefault(sessionID SessionRef, self PUSBRef) (_retval string, _err error) {
	log.Println("PUSB.GetPath not mocked")
	_err = errors.New("PUSB.GetPath not mocked")
	return
}

var PUSBClassGetPathMockedCallback = PUSBClassGetPathMockDefault

func (_class PUSBClass) GetPathMock(sessionID SessionRef, self PUSBRef) (_retval string, _err error) {
	return PUSBClassGetPathMockedCallback(sessionID, self)
}
// Get the path field of the given PUSB.
func (_class PUSBClass) GetPath(sessionID SessionRef, self PUSBRef) (_retval string, _err error) {
	if IsMock {
		return _class.GetPathMock(sessionID, self)
	}	
	_method := "PUSB.get_path"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPUSBRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func PUSBClassGetHostMockDefault(sessionID SessionRef, self PUSBRef) (_retval HostRef, _err error) {
	log.Println("PUSB.GetHost not mocked")
	_err = errors.New("PUSB.GetHost not mocked")
	return
}

var PUSBClassGetHostMockedCallback = PUSBClassGetHostMockDefault

func (_class PUSBClass) GetHostMock(sessionID SessionRef, self PUSBRef) (_retval HostRef, _err error) {
	return PUSBClassGetHostMockedCallback(sessionID, self)
}
// Get the host field of the given PUSB.
func (_class PUSBClass) GetHost(sessionID SessionRef, self PUSBRef) (_retval HostRef, _err error) {
	if IsMock {
		return _class.GetHostMock(sessionID, self)
	}	
	_method := "PUSB.get_host"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPUSBRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func PUSBClassGetUSBGroupMockDefault(sessionID SessionRef, self PUSBRef) (_retval USBGroupRef, _err error) {
	log.Println("PUSB.GetUSBGroup not mocked")
	_err = errors.New("PUSB.GetUSBGroup not mocked")
	return
}

var PUSBClassGetUSBGroupMockedCallback = PUSBClassGetUSBGroupMockDefault

func (_class PUSBClass) GetUSBGroupMock(sessionID SessionRef, self PUSBRef) (_retval USBGroupRef, _err error) {
	return PUSBClassGetUSBGroupMockedCallback(sessionID, self)
}
// Get the USB_group field of the given PUSB.
func (_class PUSBClass) GetUSBGroup(sessionID SessionRef, self PUSBRef) (_retval USBGroupRef, _err error) {
	if IsMock {
		return _class.GetUSBGroupMock(sessionID, self)
	}	
	_method := "PUSB.get_USB_group"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPUSBRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertUSBGroupRefToGo(_method + " -> ", _result.Value)
	return
}


func PUSBClassGetUUIDMockDefault(sessionID SessionRef, self PUSBRef) (_retval string, _err error) {
	log.Println("PUSB.GetUUID not mocked")
	_err = errors.New("PUSB.GetUUID not mocked")
	return
}

var PUSBClassGetUUIDMockedCallback = PUSBClassGetUUIDMockDefault

func (_class PUSBClass) GetUUIDMock(sessionID SessionRef, self PUSBRef) (_retval string, _err error) {
	return PUSBClassGetUUIDMockedCallback(sessionID, self)
}
// Get the uuid field of the given PUSB.
func (_class PUSBClass) GetUUID(sessionID SessionRef, self PUSBRef) (_retval string, _err error) {
	if IsMock {
		return _class.GetUUIDMock(sessionID, self)
	}	
	_method := "PUSB.get_uuid"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPUSBRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func PUSBClassGetByUUIDMockDefault(sessionID SessionRef, uuid string) (_retval PUSBRef, _err error) {
	log.Println("PUSB.GetByUUID not mocked")
	_err = errors.New("PUSB.GetByUUID not mocked")
	return
}

var PUSBClassGetByUUIDMockedCallback = PUSBClassGetByUUIDMockDefault

func (_class PUSBClass) GetByUUIDMock(sessionID SessionRef, uuid string) (_retval PUSBRef, _err error) {
	return PUSBClassGetByUUIDMockedCallback(sessionID, uuid)
}
// Get a reference to the PUSB instance with the specified UUID.
func (_class PUSBClass) GetByUUID(sessionID SessionRef, uuid string) (_retval PUSBRef, _err error) {
	if IsMock {
		return _class.GetByUUIDMock(sessionID, uuid)
	}	
	_method := "PUSB.get_by_uuid"
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
	_retval, _err = convertPUSBRefToGo(_method + " -> ", _result.Value)
	return
}


func PUSBClassGetRecordMockDefault(sessionID SessionRef, self PUSBRef) (_retval PUSBRecord, _err error) {
	log.Println("PUSB.GetRecord not mocked")
	_err = errors.New("PUSB.GetRecord not mocked")
	return
}

var PUSBClassGetRecordMockedCallback = PUSBClassGetRecordMockDefault

func (_class PUSBClass) GetRecordMock(sessionID SessionRef, self PUSBRef) (_retval PUSBRecord, _err error) {
	return PUSBClassGetRecordMockedCallback(sessionID, self)
}
// Get a record containing the current state of the given PUSB.
func (_class PUSBClass) GetRecord(sessionID SessionRef, self PUSBRef) (_retval PUSBRecord, _err error) {
	if IsMock {
		return _class.GetRecordMock(sessionID, self)
	}	
	_method := "PUSB.get_record"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPUSBRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertPUSBRecordToGo(_method + " -> ", _result.Value)
	return
}
