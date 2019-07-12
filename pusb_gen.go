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

func (_class PUSBClass) GetAllRecords__mock(sessionID SessionRef) (_retval map[PUSBRef]PUSBRecord, _err error) {
	log.Println("PUSB.GetAllRecords not mocked")
	_err = errors.New("PUSB.GetAllRecords not mocked")
	return
}
// Return a map of PUSB references to PUSB records for all PUSBs known to the system.
func (_class PUSBClass) GetAllRecords(sessionID SessionRef) (_retval map[PUSBRef]PUSBRecord, _err error) {
	if (IsMock) {
		return _class.GetAllRecords__mock(sessionID)
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

func (_class PUSBClass) GetAll__mock(sessionID SessionRef) (_retval []PUSBRef, _err error) {
	log.Println("PUSB.GetAll not mocked")
	_err = errors.New("PUSB.GetAll not mocked")
	return
}
// Return a list of all the PUSBs known to the system.
func (_class PUSBClass) GetAll(sessionID SessionRef) (_retval []PUSBRef, _err error) {
	if (IsMock) {
		return _class.GetAll__mock(sessionID)
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

func (_class PUSBClass) SetPassthroughEnabled__mock(sessionID SessionRef, self PUSBRef, value bool) (_err error) {
	log.Println("PUSB.SetPassthroughEnabled not mocked")
	_err = errors.New("PUSB.SetPassthroughEnabled not mocked")
	return
}
// 
func (_class PUSBClass) SetPassthroughEnabled(sessionID SessionRef, self PUSBRef, value bool) (_err error) {
	if (IsMock) {
		return _class.SetPassthroughEnabled__mock(sessionID, self, value)
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

func (_class PUSBClass) Scan__mock(sessionID SessionRef, host HostRef) (_err error) {
	log.Println("PUSB.Scan not mocked")
	_err = errors.New("PUSB.Scan not mocked")
	return
}
// 
func (_class PUSBClass) Scan(sessionID SessionRef, host HostRef) (_err error) {
	if (IsMock) {
		return _class.Scan__mock(sessionID, host)
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

func (_class PUSBClass) RemoveFromOtherConfig__mock(sessionID SessionRef, self PUSBRef, key string) (_err error) {
	log.Println("PUSB.RemoveFromOtherConfig not mocked")
	_err = errors.New("PUSB.RemoveFromOtherConfig not mocked")
	return
}
// Remove the given key and its corresponding value from the other_config field of the given PUSB.  If the key is not in that Map, then do nothing.
func (_class PUSBClass) RemoveFromOtherConfig(sessionID SessionRef, self PUSBRef, key string) (_err error) {
	if (IsMock) {
		return _class.RemoveFromOtherConfig__mock(sessionID, self, key)
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

func (_class PUSBClass) AddToOtherConfig__mock(sessionID SessionRef, self PUSBRef, key string, value string) (_err error) {
	log.Println("PUSB.AddToOtherConfig not mocked")
	_err = errors.New("PUSB.AddToOtherConfig not mocked")
	return
}
// Add the given key-value pair to the other_config field of the given PUSB.
func (_class PUSBClass) AddToOtherConfig(sessionID SessionRef, self PUSBRef, key string, value string) (_err error) {
	if (IsMock) {
		return _class.AddToOtherConfig__mock(sessionID, self, key, value)
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

func (_class PUSBClass) SetOtherConfig__mock(sessionID SessionRef, self PUSBRef, value map[string]string) (_err error) {
	log.Println("PUSB.SetOtherConfig not mocked")
	_err = errors.New("PUSB.SetOtherConfig not mocked")
	return
}
// Set the other_config field of the given PUSB.
func (_class PUSBClass) SetOtherConfig(sessionID SessionRef, self PUSBRef, value map[string]string) (_err error) {
	if (IsMock) {
		return _class.SetOtherConfig__mock(sessionID, self, value)
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

func (_class PUSBClass) GetOtherConfig__mock(sessionID SessionRef, self PUSBRef) (_retval map[string]string, _err error) {
	log.Println("PUSB.GetOtherConfig not mocked")
	_err = errors.New("PUSB.GetOtherConfig not mocked")
	return
}
// Get the other_config field of the given PUSB.
func (_class PUSBClass) GetOtherConfig(sessionID SessionRef, self PUSBRef) (_retval map[string]string, _err error) {
	if (IsMock) {
		return _class.GetOtherConfig__mock(sessionID, self)
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

func (_class PUSBClass) GetPassthroughEnabled__mock(sessionID SessionRef, self PUSBRef) (_retval bool, _err error) {
	log.Println("PUSB.GetPassthroughEnabled not mocked")
	_err = errors.New("PUSB.GetPassthroughEnabled not mocked")
	return
}
// Get the passthrough_enabled field of the given PUSB.
func (_class PUSBClass) GetPassthroughEnabled(sessionID SessionRef, self PUSBRef) (_retval bool, _err error) {
	if (IsMock) {
		return _class.GetPassthroughEnabled__mock(sessionID, self)
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

func (_class PUSBClass) GetDescription__mock(sessionID SessionRef, self PUSBRef) (_retval string, _err error) {
	log.Println("PUSB.GetDescription not mocked")
	_err = errors.New("PUSB.GetDescription not mocked")
	return
}
// Get the description field of the given PUSB.
func (_class PUSBClass) GetDescription(sessionID SessionRef, self PUSBRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetDescription__mock(sessionID, self)
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

func (_class PUSBClass) GetVersion__mock(sessionID SessionRef, self PUSBRef) (_retval string, _err error) {
	log.Println("PUSB.GetVersion not mocked")
	_err = errors.New("PUSB.GetVersion not mocked")
	return
}
// Get the version field of the given PUSB.
func (_class PUSBClass) GetVersion(sessionID SessionRef, self PUSBRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetVersion__mock(sessionID, self)
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

func (_class PUSBClass) GetSerial__mock(sessionID SessionRef, self PUSBRef) (_retval string, _err error) {
	log.Println("PUSB.GetSerial not mocked")
	_err = errors.New("PUSB.GetSerial not mocked")
	return
}
// Get the serial field of the given PUSB.
func (_class PUSBClass) GetSerial(sessionID SessionRef, self PUSBRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetSerial__mock(sessionID, self)
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

func (_class PUSBClass) GetProductDesc__mock(sessionID SessionRef, self PUSBRef) (_retval string, _err error) {
	log.Println("PUSB.GetProductDesc not mocked")
	_err = errors.New("PUSB.GetProductDesc not mocked")
	return
}
// Get the product_desc field of the given PUSB.
func (_class PUSBClass) GetProductDesc(sessionID SessionRef, self PUSBRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetProductDesc__mock(sessionID, self)
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

func (_class PUSBClass) GetProductID__mock(sessionID SessionRef, self PUSBRef) (_retval string, _err error) {
	log.Println("PUSB.GetProductID not mocked")
	_err = errors.New("PUSB.GetProductID not mocked")
	return
}
// Get the product_id field of the given PUSB.
func (_class PUSBClass) GetProductID(sessionID SessionRef, self PUSBRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetProductID__mock(sessionID, self)
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

func (_class PUSBClass) GetVendorDesc__mock(sessionID SessionRef, self PUSBRef) (_retval string, _err error) {
	log.Println("PUSB.GetVendorDesc not mocked")
	_err = errors.New("PUSB.GetVendorDesc not mocked")
	return
}
// Get the vendor_desc field of the given PUSB.
func (_class PUSBClass) GetVendorDesc(sessionID SessionRef, self PUSBRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetVendorDesc__mock(sessionID, self)
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

func (_class PUSBClass) GetVendorID__mock(sessionID SessionRef, self PUSBRef) (_retval string, _err error) {
	log.Println("PUSB.GetVendorID not mocked")
	_err = errors.New("PUSB.GetVendorID not mocked")
	return
}
// Get the vendor_id field of the given PUSB.
func (_class PUSBClass) GetVendorID(sessionID SessionRef, self PUSBRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetVendorID__mock(sessionID, self)
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

func (_class PUSBClass) GetPath__mock(sessionID SessionRef, self PUSBRef) (_retval string, _err error) {
	log.Println("PUSB.GetPath not mocked")
	_err = errors.New("PUSB.GetPath not mocked")
	return
}
// Get the path field of the given PUSB.
func (_class PUSBClass) GetPath(sessionID SessionRef, self PUSBRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetPath__mock(sessionID, self)
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

func (_class PUSBClass) GetHost__mock(sessionID SessionRef, self PUSBRef) (_retval HostRef, _err error) {
	log.Println("PUSB.GetHost not mocked")
	_err = errors.New("PUSB.GetHost not mocked")
	return
}
// Get the host field of the given PUSB.
func (_class PUSBClass) GetHost(sessionID SessionRef, self PUSBRef) (_retval HostRef, _err error) {
	if (IsMock) {
		return _class.GetHost__mock(sessionID, self)
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

func (_class PUSBClass) GetUSBGroup__mock(sessionID SessionRef, self PUSBRef) (_retval USBGroupRef, _err error) {
	log.Println("PUSB.GetUSBGroup not mocked")
	_err = errors.New("PUSB.GetUSBGroup not mocked")
	return
}
// Get the USB_group field of the given PUSB.
func (_class PUSBClass) GetUSBGroup(sessionID SessionRef, self PUSBRef) (_retval USBGroupRef, _err error) {
	if (IsMock) {
		return _class.GetUSBGroup__mock(sessionID, self)
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

func (_class PUSBClass) GetUUID__mock(sessionID SessionRef, self PUSBRef) (_retval string, _err error) {
	log.Println("PUSB.GetUUID not mocked")
	_err = errors.New("PUSB.GetUUID not mocked")
	return
}
// Get the uuid field of the given PUSB.
func (_class PUSBClass) GetUUID(sessionID SessionRef, self PUSBRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetUUID__mock(sessionID, self)
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

func (_class PUSBClass) GetByUUID__mock(sessionID SessionRef, uuid string) (_retval PUSBRef, _err error) {
	log.Println("PUSB.GetByUUID not mocked")
	_err = errors.New("PUSB.GetByUUID not mocked")
	return
}
// Get a reference to the PUSB instance with the specified UUID.
func (_class PUSBClass) GetByUUID(sessionID SessionRef, uuid string) (_retval PUSBRef, _err error) {
	if (IsMock) {
		return _class.GetByUUID__mock(sessionID, uuid)
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

func (_class PUSBClass) GetRecord__mock(sessionID SessionRef, self PUSBRef) (_retval PUSBRecord, _err error) {
	log.Println("PUSB.GetRecord not mocked")
	_err = errors.New("PUSB.GetRecord not mocked")
	return
}
// Get a record containing the current state of the given PUSB.
func (_class PUSBClass) GetRecord(sessionID SessionRef, self PUSBRef) (_retval PUSBRecord, _err error) {
	if (IsMock) {
		return _class.GetRecord__mock(sessionID, self)
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
