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

type VusbOperations string

const (
  // Attempting to attach this VUSB to a VM
	VusbOperationsAttach VusbOperations = "attach"
  // Attempting to plug this VUSB into a VM
	VusbOperationsPlug VusbOperations = "plug"
  // Attempting to hot unplug this VUSB
	VusbOperationsUnplug VusbOperations = "unplug"
)

type VUSBRecord struct {
  // Unique identifier/object reference
	UUID string
  // list of the operations allowed in this state. This list is advisory only and the server state may have changed by the time this field is read by a client.
	AllowedOperations []VusbOperations
  // links each of the running tasks using this object (by reference) to a current_operation enum which describes the nature of the task.
	CurrentOperations map[string]VusbOperations
  // VM that owns the VUSB
	VM VMRef
  // USB group used by the VUSB
	USBGroup USBGroupRef
  // Additional configuration
	OtherConfig map[string]string
  // is the device currently attached
	CurrentlyAttached bool
}

type VUSBRef string

// Describes the vusb device
type VUSBClass struct {
	client *Client
}


func VUSBClassGetAllRecordsMockDefault(sessionID SessionRef) (_retval map[VUSBRef]VUSBRecord, _err error) {
	log.Println("VUSB.GetAllRecords not mocked")
	_err = errors.New("VUSB.GetAllRecords not mocked")
	return
}

var VUSBClassGetAllRecordsMockedCallback = VUSBClassGetAllRecordsMockDefault

func (_class VUSBClass) GetAllRecordsMock(sessionID SessionRef) (_retval map[VUSBRef]VUSBRecord, _err error) {
	return VUSBClassGetAllRecordsMockedCallback(sessionID)
}
// Return a map of VUSB references to VUSB records for all VUSBs known to the system.
func (_class VUSBClass) GetAllRecords(sessionID SessionRef) (_retval map[VUSBRef]VUSBRecord, _err error) {
	if IsMock {
		return _class.GetAllRecordsMock(sessionID)
	}	
	_method := "VUSB.get_all_records"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVUSBRefToVUSBRecordMapToGo(_method + " -> ", _result.Value)
	return
}


func VUSBClassGetAllMockDefault(sessionID SessionRef) (_retval []VUSBRef, _err error) {
	log.Println("VUSB.GetAll not mocked")
	_err = errors.New("VUSB.GetAll not mocked")
	return
}

var VUSBClassGetAllMockedCallback = VUSBClassGetAllMockDefault

func (_class VUSBClass) GetAllMock(sessionID SessionRef) (_retval []VUSBRef, _err error) {
	return VUSBClassGetAllMockedCallback(sessionID)
}
// Return a list of all the VUSBs known to the system.
func (_class VUSBClass) GetAll(sessionID SessionRef) (_retval []VUSBRef, _err error) {
	if IsMock {
		return _class.GetAllMock(sessionID)
	}	
	_method := "VUSB.get_all"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVUSBRefSetToGo(_method + " -> ", _result.Value)
	return
}


func VUSBClassDestroyMockDefault(sessionID SessionRef, self VUSBRef) (_err error) {
	log.Println("VUSB.Destroy not mocked")
	_err = errors.New("VUSB.Destroy not mocked")
	return
}

var VUSBClassDestroyMockedCallback = VUSBClassDestroyMockDefault

func (_class VUSBClass) DestroyMock(sessionID SessionRef, self VUSBRef) (_err error) {
	return VUSBClassDestroyMockedCallback(sessionID, self)
}
// Removes a VUSB record from the database
func (_class VUSBClass) Destroy(sessionID SessionRef, self VUSBRef) (_err error) {
	if IsMock {
		return _class.DestroyMock(sessionID, self)
	}	
	_method := "VUSB.destroy"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVUSBRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg)
	return
}


func VUSBClassUnplugMockDefault(sessionID SessionRef, self VUSBRef) (_err error) {
	log.Println("VUSB.Unplug not mocked")
	_err = errors.New("VUSB.Unplug not mocked")
	return
}

var VUSBClassUnplugMockedCallback = VUSBClassUnplugMockDefault

func (_class VUSBClass) UnplugMock(sessionID SessionRef, self VUSBRef) (_err error) {
	return VUSBClassUnplugMockedCallback(sessionID, self)
}
// Unplug the vusb device from the vm.
func (_class VUSBClass) Unplug(sessionID SessionRef, self VUSBRef) (_err error) {
	if IsMock {
		return _class.UnplugMock(sessionID, self)
	}	
	_method := "VUSB.unplug"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVUSBRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg)
	return
}


func VUSBClassCreateMockDefault(sessionID SessionRef, vm VMRef, usbGroup USBGroupRef, otherConfig map[string]string) (_retval VUSBRef, _err error) {
	log.Println("VUSB.Create not mocked")
	_err = errors.New("VUSB.Create not mocked")
	return
}

var VUSBClassCreateMockedCallback = VUSBClassCreateMockDefault

func (_class VUSBClass) CreateMock(sessionID SessionRef, vm VMRef, usbGroup USBGroupRef, otherConfig map[string]string) (_retval VUSBRef, _err error) {
	return VUSBClassCreateMockedCallback(sessionID, vm, usbGroup, otherConfig)
}
// Create a new VUSB record in the database only
func (_class VUSBClass) Create(sessionID SessionRef, vm VMRef, usbGroup USBGroupRef, otherConfig map[string]string) (_retval VUSBRef, _err error) {
	if IsMock {
		return _class.CreateMock(sessionID, vm, usbGroup, otherConfig)
	}	
	_method := "VUSB.create"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_vmArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "VM"), vm)
	if _err != nil {
		return
	}
	_usbGroupArg, _err := convertUSBGroupRefToXen(fmt.Sprintf("%s(%s)", _method, "USB_group"), usbGroup)
	if _err != nil {
		return
	}
	_otherConfigArg, _err := convertStringToStringMapToXen(fmt.Sprintf("%s(%s)", _method, "other_config"), otherConfig)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _vmArg, _usbGroupArg, _otherConfigArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVUSBRefToGo(_method + " -> ", _result.Value)
	return
}


func VUSBClassRemoveFromOtherConfigMockDefault(sessionID SessionRef, self VUSBRef, key string) (_err error) {
	log.Println("VUSB.RemoveFromOtherConfig not mocked")
	_err = errors.New("VUSB.RemoveFromOtherConfig not mocked")
	return
}

var VUSBClassRemoveFromOtherConfigMockedCallback = VUSBClassRemoveFromOtherConfigMockDefault

func (_class VUSBClass) RemoveFromOtherConfigMock(sessionID SessionRef, self VUSBRef, key string) (_err error) {
	return VUSBClassRemoveFromOtherConfigMockedCallback(sessionID, self, key)
}
// Remove the given key and its corresponding value from the other_config field of the given VUSB.  If the key is not in that Map, then do nothing.
func (_class VUSBClass) RemoveFromOtherConfig(sessionID SessionRef, self VUSBRef, key string) (_err error) {
	if IsMock {
		return _class.RemoveFromOtherConfigMock(sessionID, self, key)
	}	
	_method := "VUSB.remove_from_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVUSBRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func VUSBClassAddToOtherConfigMockDefault(sessionID SessionRef, self VUSBRef, key string, value string) (_err error) {
	log.Println("VUSB.AddToOtherConfig not mocked")
	_err = errors.New("VUSB.AddToOtherConfig not mocked")
	return
}

var VUSBClassAddToOtherConfigMockedCallback = VUSBClassAddToOtherConfigMockDefault

func (_class VUSBClass) AddToOtherConfigMock(sessionID SessionRef, self VUSBRef, key string, value string) (_err error) {
	return VUSBClassAddToOtherConfigMockedCallback(sessionID, self, key, value)
}
// Add the given key-value pair to the other_config field of the given VUSB.
func (_class VUSBClass) AddToOtherConfig(sessionID SessionRef, self VUSBRef, key string, value string) (_err error) {
	if IsMock {
		return _class.AddToOtherConfigMock(sessionID, self, key, value)
	}	
	_method := "VUSB.add_to_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVUSBRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func VUSBClassSetOtherConfigMockDefault(sessionID SessionRef, self VUSBRef, value map[string]string) (_err error) {
	log.Println("VUSB.SetOtherConfig not mocked")
	_err = errors.New("VUSB.SetOtherConfig not mocked")
	return
}

var VUSBClassSetOtherConfigMockedCallback = VUSBClassSetOtherConfigMockDefault

func (_class VUSBClass) SetOtherConfigMock(sessionID SessionRef, self VUSBRef, value map[string]string) (_err error) {
	return VUSBClassSetOtherConfigMockedCallback(sessionID, self, value)
}
// Set the other_config field of the given VUSB.
func (_class VUSBClass) SetOtherConfig(sessionID SessionRef, self VUSBRef, value map[string]string) (_err error) {
	if IsMock {
		return _class.SetOtherConfigMock(sessionID, self, value)
	}	
	_method := "VUSB.set_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVUSBRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func VUSBClassGetCurrentlyAttachedMockDefault(sessionID SessionRef, self VUSBRef) (_retval bool, _err error) {
	log.Println("VUSB.GetCurrentlyAttached not mocked")
	_err = errors.New("VUSB.GetCurrentlyAttached not mocked")
	return
}

var VUSBClassGetCurrentlyAttachedMockedCallback = VUSBClassGetCurrentlyAttachedMockDefault

func (_class VUSBClass) GetCurrentlyAttachedMock(sessionID SessionRef, self VUSBRef) (_retval bool, _err error) {
	return VUSBClassGetCurrentlyAttachedMockedCallback(sessionID, self)
}
// Get the currently_attached field of the given VUSB.
func (_class VUSBClass) GetCurrentlyAttached(sessionID SessionRef, self VUSBRef) (_retval bool, _err error) {
	if IsMock {
		return _class.GetCurrentlyAttachedMock(sessionID, self)
	}	
	_method := "VUSB.get_currently_attached"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVUSBRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func VUSBClassGetOtherConfigMockDefault(sessionID SessionRef, self VUSBRef) (_retval map[string]string, _err error) {
	log.Println("VUSB.GetOtherConfig not mocked")
	_err = errors.New("VUSB.GetOtherConfig not mocked")
	return
}

var VUSBClassGetOtherConfigMockedCallback = VUSBClassGetOtherConfigMockDefault

func (_class VUSBClass) GetOtherConfigMock(sessionID SessionRef, self VUSBRef) (_retval map[string]string, _err error) {
	return VUSBClassGetOtherConfigMockedCallback(sessionID, self)
}
// Get the other_config field of the given VUSB.
func (_class VUSBClass) GetOtherConfig(sessionID SessionRef, self VUSBRef) (_retval map[string]string, _err error) {
	if IsMock {
		return _class.GetOtherConfigMock(sessionID, self)
	}	
	_method := "VUSB.get_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVUSBRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func VUSBClassGetUSBGroupMockDefault(sessionID SessionRef, self VUSBRef) (_retval USBGroupRef, _err error) {
	log.Println("VUSB.GetUSBGroup not mocked")
	_err = errors.New("VUSB.GetUSBGroup not mocked")
	return
}

var VUSBClassGetUSBGroupMockedCallback = VUSBClassGetUSBGroupMockDefault

func (_class VUSBClass) GetUSBGroupMock(sessionID SessionRef, self VUSBRef) (_retval USBGroupRef, _err error) {
	return VUSBClassGetUSBGroupMockedCallback(sessionID, self)
}
// Get the USB_group field of the given VUSB.
func (_class VUSBClass) GetUSBGroup(sessionID SessionRef, self VUSBRef) (_retval USBGroupRef, _err error) {
	if IsMock {
		return _class.GetUSBGroupMock(sessionID, self)
	}	
	_method := "VUSB.get_USB_group"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVUSBRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func VUSBClassGetVMMockDefault(sessionID SessionRef, self VUSBRef) (_retval VMRef, _err error) {
	log.Println("VUSB.GetVM not mocked")
	_err = errors.New("VUSB.GetVM not mocked")
	return
}

var VUSBClassGetVMMockedCallback = VUSBClassGetVMMockDefault

func (_class VUSBClass) GetVMMock(sessionID SessionRef, self VUSBRef) (_retval VMRef, _err error) {
	return VUSBClassGetVMMockedCallback(sessionID, self)
}
// Get the VM field of the given VUSB.
func (_class VUSBClass) GetVM(sessionID SessionRef, self VUSBRef) (_retval VMRef, _err error) {
	if IsMock {
		return _class.GetVMMock(sessionID, self)
	}	
	_method := "VUSB.get_VM"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVUSBRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVMRefToGo(_method + " -> ", _result.Value)
	return
}


func VUSBClassGetCurrentOperationsMockDefault(sessionID SessionRef, self VUSBRef) (_retval map[string]VusbOperations, _err error) {
	log.Println("VUSB.GetCurrentOperations not mocked")
	_err = errors.New("VUSB.GetCurrentOperations not mocked")
	return
}

var VUSBClassGetCurrentOperationsMockedCallback = VUSBClassGetCurrentOperationsMockDefault

func (_class VUSBClass) GetCurrentOperationsMock(sessionID SessionRef, self VUSBRef) (_retval map[string]VusbOperations, _err error) {
	return VUSBClassGetCurrentOperationsMockedCallback(sessionID, self)
}
// Get the current_operations field of the given VUSB.
func (_class VUSBClass) GetCurrentOperations(sessionID SessionRef, self VUSBRef) (_retval map[string]VusbOperations, _err error) {
	if IsMock {
		return _class.GetCurrentOperationsMock(sessionID, self)
	}	
	_method := "VUSB.get_current_operations"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVUSBRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringToEnumVusbOperationsMapToGo(_method + " -> ", _result.Value)
	return
}


func VUSBClassGetAllowedOperationsMockDefault(sessionID SessionRef, self VUSBRef) (_retval []VusbOperations, _err error) {
	log.Println("VUSB.GetAllowedOperations not mocked")
	_err = errors.New("VUSB.GetAllowedOperations not mocked")
	return
}

var VUSBClassGetAllowedOperationsMockedCallback = VUSBClassGetAllowedOperationsMockDefault

func (_class VUSBClass) GetAllowedOperationsMock(sessionID SessionRef, self VUSBRef) (_retval []VusbOperations, _err error) {
	return VUSBClassGetAllowedOperationsMockedCallback(sessionID, self)
}
// Get the allowed_operations field of the given VUSB.
func (_class VUSBClass) GetAllowedOperations(sessionID SessionRef, self VUSBRef) (_retval []VusbOperations, _err error) {
	if IsMock {
		return _class.GetAllowedOperationsMock(sessionID, self)
	}	
	_method := "VUSB.get_allowed_operations"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVUSBRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertEnumVusbOperationsSetToGo(_method + " -> ", _result.Value)
	return
}


func VUSBClassGetUUIDMockDefault(sessionID SessionRef, self VUSBRef) (_retval string, _err error) {
	log.Println("VUSB.GetUUID not mocked")
	_err = errors.New("VUSB.GetUUID not mocked")
	return
}

var VUSBClassGetUUIDMockedCallback = VUSBClassGetUUIDMockDefault

func (_class VUSBClass) GetUUIDMock(sessionID SessionRef, self VUSBRef) (_retval string, _err error) {
	return VUSBClassGetUUIDMockedCallback(sessionID, self)
}
// Get the uuid field of the given VUSB.
func (_class VUSBClass) GetUUID(sessionID SessionRef, self VUSBRef) (_retval string, _err error) {
	if IsMock {
		return _class.GetUUIDMock(sessionID, self)
	}	
	_method := "VUSB.get_uuid"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVUSBRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func VUSBClassGetByUUIDMockDefault(sessionID SessionRef, uuid string) (_retval VUSBRef, _err error) {
	log.Println("VUSB.GetByUUID not mocked")
	_err = errors.New("VUSB.GetByUUID not mocked")
	return
}

var VUSBClassGetByUUIDMockedCallback = VUSBClassGetByUUIDMockDefault

func (_class VUSBClass) GetByUUIDMock(sessionID SessionRef, uuid string) (_retval VUSBRef, _err error) {
	return VUSBClassGetByUUIDMockedCallback(sessionID, uuid)
}
// Get a reference to the VUSB instance with the specified UUID.
func (_class VUSBClass) GetByUUID(sessionID SessionRef, uuid string) (_retval VUSBRef, _err error) {
	if IsMock {
		return _class.GetByUUIDMock(sessionID, uuid)
	}	
	_method := "VUSB.get_by_uuid"
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
	_retval, _err = convertVUSBRefToGo(_method + " -> ", _result.Value)
	return
}


func VUSBClassGetRecordMockDefault(sessionID SessionRef, self VUSBRef) (_retval VUSBRecord, _err error) {
	log.Println("VUSB.GetRecord not mocked")
	_err = errors.New("VUSB.GetRecord not mocked")
	return
}

var VUSBClassGetRecordMockedCallback = VUSBClassGetRecordMockDefault

func (_class VUSBClass) GetRecordMock(sessionID SessionRef, self VUSBRef) (_retval VUSBRecord, _err error) {
	return VUSBClassGetRecordMockedCallback(sessionID, self)
}
// Get a record containing the current state of the given VUSB.
func (_class VUSBClass) GetRecord(sessionID SessionRef, self VUSBRef) (_retval VUSBRecord, _err error) {
	if IsMock {
		return _class.GetRecordMock(sessionID, self)
	}	
	_method := "VUSB.get_record"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVUSBRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVUSBRecordToGo(_method + " -> ", _result.Value)
	return
}
