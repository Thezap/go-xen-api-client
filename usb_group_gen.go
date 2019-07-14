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

type USBGroupRecord struct {
  // Unique identifier/object reference
	UUID string
  // a human-readable name
	NameLabel string
  // a notes field containing human-readable description
	NameDescription string
  // List of PUSBs in the group
	PUSBs []PUSBRef
  // List of VUSBs using the group
	VUSBs []VUSBRef
  // Additional configuration
	OtherConfig map[string]string
}

type USBGroupRef string

// A group of compatible USBs across the resource pool
type USBGroupClass struct {
	client *Client
}


func USBGroupClassGetAllRecordsMockDefault(sessionID SessionRef) (_retval map[USBGroupRef]USBGroupRecord, _err error) {
	log.Println("USBGroup.GetAllRecords not mocked")
	_err = errors.New("USBGroup.GetAllRecords not mocked")
	return
}

var USBGroupClassGetAllRecordsMockedCallback = USBGroupClassGetAllRecordsMockDefault

func (_class USBGroupClass) GetAllRecordsMock(sessionID SessionRef) (_retval map[USBGroupRef]USBGroupRecord, _err error) {
	return USBGroupClassGetAllRecordsMockedCallback(sessionID)
}
// Return a map of USB_group references to USB_group records for all USB_groups known to the system.
func (_class USBGroupClass) GetAllRecords(sessionID SessionRef) (_retval map[USBGroupRef]USBGroupRecord, _err error) {
	if IsMock {
		return _class.GetAllRecordsMock(sessionID)
	}	
	_method := "USB_group.get_all_records"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertUSBGroupRefToUSBGroupRecordMapToGo(_method + " -> ", _result.Value)
	return
}


func USBGroupClassGetAllMockDefault(sessionID SessionRef) (_retval []USBGroupRef, _err error) {
	log.Println("USBGroup.GetAll not mocked")
	_err = errors.New("USBGroup.GetAll not mocked")
	return
}

var USBGroupClassGetAllMockedCallback = USBGroupClassGetAllMockDefault

func (_class USBGroupClass) GetAllMock(sessionID SessionRef) (_retval []USBGroupRef, _err error) {
	return USBGroupClassGetAllMockedCallback(sessionID)
}
// Return a list of all the USB_groups known to the system.
func (_class USBGroupClass) GetAll(sessionID SessionRef) (_retval []USBGroupRef, _err error) {
	if IsMock {
		return _class.GetAllMock(sessionID)
	}	
	_method := "USB_group.get_all"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertUSBGroupRefSetToGo(_method + " -> ", _result.Value)
	return
}


func USBGroupClassDestroyMockDefault(sessionID SessionRef, self USBGroupRef) (_err error) {
	log.Println("USBGroup.Destroy not mocked")
	_err = errors.New("USBGroup.Destroy not mocked")
	return
}

var USBGroupClassDestroyMockedCallback = USBGroupClassDestroyMockDefault

func (_class USBGroupClass) DestroyMock(sessionID SessionRef, self USBGroupRef) (_err error) {
	return USBGroupClassDestroyMockedCallback(sessionID, self)
}
// 
func (_class USBGroupClass) Destroy(sessionID SessionRef, self USBGroupRef) (_err error) {
	if IsMock {
		return _class.DestroyMock(sessionID, self)
	}	
	_method := "USB_group.destroy"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertUSBGroupRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg)
	return
}


func USBGroupClassCreateMockDefault(sessionID SessionRef, nameLabel string, nameDescription string, otherConfig map[string]string) (_retval USBGroupRef, _err error) {
	log.Println("USBGroup.Create not mocked")
	_err = errors.New("USBGroup.Create not mocked")
	return
}

var USBGroupClassCreateMockedCallback = USBGroupClassCreateMockDefault

func (_class USBGroupClass) CreateMock(sessionID SessionRef, nameLabel string, nameDescription string, otherConfig map[string]string) (_retval USBGroupRef, _err error) {
	return USBGroupClassCreateMockedCallback(sessionID, nameLabel, nameDescription, otherConfig)
}
// 
func (_class USBGroupClass) Create(sessionID SessionRef, nameLabel string, nameDescription string, otherConfig map[string]string) (_retval USBGroupRef, _err error) {
	if IsMock {
		return _class.CreateMock(sessionID, nameLabel, nameDescription, otherConfig)
	}	
	_method := "USB_group.create"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_nameLabelArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "name_label"), nameLabel)
	if _err != nil {
		return
	}
	_nameDescriptionArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "name_description"), nameDescription)
	if _err != nil {
		return
	}
	_otherConfigArg, _err := convertStringToStringMapToXen(fmt.Sprintf("%s(%s)", _method, "other_config"), otherConfig)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _nameLabelArg, _nameDescriptionArg, _otherConfigArg)
	if _err != nil {
		return
	}
	_retval, _err = convertUSBGroupRefToGo(_method + " -> ", _result.Value)
	return
}


func USBGroupClassRemoveFromOtherConfigMockDefault(sessionID SessionRef, self USBGroupRef, key string) (_err error) {
	log.Println("USBGroup.RemoveFromOtherConfig not mocked")
	_err = errors.New("USBGroup.RemoveFromOtherConfig not mocked")
	return
}

var USBGroupClassRemoveFromOtherConfigMockedCallback = USBGroupClassRemoveFromOtherConfigMockDefault

func (_class USBGroupClass) RemoveFromOtherConfigMock(sessionID SessionRef, self USBGroupRef, key string) (_err error) {
	return USBGroupClassRemoveFromOtherConfigMockedCallback(sessionID, self, key)
}
// Remove the given key and its corresponding value from the other_config field of the given USB_group.  If the key is not in that Map, then do nothing.
func (_class USBGroupClass) RemoveFromOtherConfig(sessionID SessionRef, self USBGroupRef, key string) (_err error) {
	if IsMock {
		return _class.RemoveFromOtherConfigMock(sessionID, self, key)
	}	
	_method := "USB_group.remove_from_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertUSBGroupRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func USBGroupClassAddToOtherConfigMockDefault(sessionID SessionRef, self USBGroupRef, key string, value string) (_err error) {
	log.Println("USBGroup.AddToOtherConfig not mocked")
	_err = errors.New("USBGroup.AddToOtherConfig not mocked")
	return
}

var USBGroupClassAddToOtherConfigMockedCallback = USBGroupClassAddToOtherConfigMockDefault

func (_class USBGroupClass) AddToOtherConfigMock(sessionID SessionRef, self USBGroupRef, key string, value string) (_err error) {
	return USBGroupClassAddToOtherConfigMockedCallback(sessionID, self, key, value)
}
// Add the given key-value pair to the other_config field of the given USB_group.
func (_class USBGroupClass) AddToOtherConfig(sessionID SessionRef, self USBGroupRef, key string, value string) (_err error) {
	if IsMock {
		return _class.AddToOtherConfigMock(sessionID, self, key, value)
	}	
	_method := "USB_group.add_to_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertUSBGroupRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func USBGroupClassSetOtherConfigMockDefault(sessionID SessionRef, self USBGroupRef, value map[string]string) (_err error) {
	log.Println("USBGroup.SetOtherConfig not mocked")
	_err = errors.New("USBGroup.SetOtherConfig not mocked")
	return
}

var USBGroupClassSetOtherConfigMockedCallback = USBGroupClassSetOtherConfigMockDefault

func (_class USBGroupClass) SetOtherConfigMock(sessionID SessionRef, self USBGroupRef, value map[string]string) (_err error) {
	return USBGroupClassSetOtherConfigMockedCallback(sessionID, self, value)
}
// Set the other_config field of the given USB_group.
func (_class USBGroupClass) SetOtherConfig(sessionID SessionRef, self USBGroupRef, value map[string]string) (_err error) {
	if IsMock {
		return _class.SetOtherConfigMock(sessionID, self, value)
	}	
	_method := "USB_group.set_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertUSBGroupRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func USBGroupClassSetNameDescriptionMockDefault(sessionID SessionRef, self USBGroupRef, value string) (_err error) {
	log.Println("USBGroup.SetNameDescription not mocked")
	_err = errors.New("USBGroup.SetNameDescription not mocked")
	return
}

var USBGroupClassSetNameDescriptionMockedCallback = USBGroupClassSetNameDescriptionMockDefault

func (_class USBGroupClass) SetNameDescriptionMock(sessionID SessionRef, self USBGroupRef, value string) (_err error) {
	return USBGroupClassSetNameDescriptionMockedCallback(sessionID, self, value)
}
// Set the name/description field of the given USB_group.
func (_class USBGroupClass) SetNameDescription(sessionID SessionRef, self USBGroupRef, value string) (_err error) {
	if IsMock {
		return _class.SetNameDescriptionMock(sessionID, self, value)
	}	
	_method := "USB_group.set_name_description"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertUSBGroupRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}


func USBGroupClassSetNameLabelMockDefault(sessionID SessionRef, self USBGroupRef, value string) (_err error) {
	log.Println("USBGroup.SetNameLabel not mocked")
	_err = errors.New("USBGroup.SetNameLabel not mocked")
	return
}

var USBGroupClassSetNameLabelMockedCallback = USBGroupClassSetNameLabelMockDefault

func (_class USBGroupClass) SetNameLabelMock(sessionID SessionRef, self USBGroupRef, value string) (_err error) {
	return USBGroupClassSetNameLabelMockedCallback(sessionID, self, value)
}
// Set the name/label field of the given USB_group.
func (_class USBGroupClass) SetNameLabel(sessionID SessionRef, self USBGroupRef, value string) (_err error) {
	if IsMock {
		return _class.SetNameLabelMock(sessionID, self, value)
	}	
	_method := "USB_group.set_name_label"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertUSBGroupRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}


func USBGroupClassGetOtherConfigMockDefault(sessionID SessionRef, self USBGroupRef) (_retval map[string]string, _err error) {
	log.Println("USBGroup.GetOtherConfig not mocked")
	_err = errors.New("USBGroup.GetOtherConfig not mocked")
	return
}

var USBGroupClassGetOtherConfigMockedCallback = USBGroupClassGetOtherConfigMockDefault

func (_class USBGroupClass) GetOtherConfigMock(sessionID SessionRef, self USBGroupRef) (_retval map[string]string, _err error) {
	return USBGroupClassGetOtherConfigMockedCallback(sessionID, self)
}
// Get the other_config field of the given USB_group.
func (_class USBGroupClass) GetOtherConfig(sessionID SessionRef, self USBGroupRef) (_retval map[string]string, _err error) {
	if IsMock {
		return _class.GetOtherConfigMock(sessionID, self)
	}	
	_method := "USB_group.get_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertUSBGroupRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func USBGroupClassGetVUSBsMockDefault(sessionID SessionRef, self USBGroupRef) (_retval []VUSBRef, _err error) {
	log.Println("USBGroup.GetVUSBs not mocked")
	_err = errors.New("USBGroup.GetVUSBs not mocked")
	return
}

var USBGroupClassGetVUSBsMockedCallback = USBGroupClassGetVUSBsMockDefault

func (_class USBGroupClass) GetVUSBsMock(sessionID SessionRef, self USBGroupRef) (_retval []VUSBRef, _err error) {
	return USBGroupClassGetVUSBsMockedCallback(sessionID, self)
}
// Get the VUSBs field of the given USB_group.
func (_class USBGroupClass) GetVUSBs(sessionID SessionRef, self USBGroupRef) (_retval []VUSBRef, _err error) {
	if IsMock {
		return _class.GetVUSBsMock(sessionID, self)
	}	
	_method := "USB_group.get_VUSBs"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertUSBGroupRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVUSBRefSetToGo(_method + " -> ", _result.Value)
	return
}


func USBGroupClassGetPUSBsMockDefault(sessionID SessionRef, self USBGroupRef) (_retval []PUSBRef, _err error) {
	log.Println("USBGroup.GetPUSBs not mocked")
	_err = errors.New("USBGroup.GetPUSBs not mocked")
	return
}

var USBGroupClassGetPUSBsMockedCallback = USBGroupClassGetPUSBsMockDefault

func (_class USBGroupClass) GetPUSBsMock(sessionID SessionRef, self USBGroupRef) (_retval []PUSBRef, _err error) {
	return USBGroupClassGetPUSBsMockedCallback(sessionID, self)
}
// Get the PUSBs field of the given USB_group.
func (_class USBGroupClass) GetPUSBs(sessionID SessionRef, self USBGroupRef) (_retval []PUSBRef, _err error) {
	if IsMock {
		return _class.GetPUSBsMock(sessionID, self)
	}	
	_method := "USB_group.get_PUSBs"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertUSBGroupRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertPUSBRefSetToGo(_method + " -> ", _result.Value)
	return
}


func USBGroupClassGetNameDescriptionMockDefault(sessionID SessionRef, self USBGroupRef) (_retval string, _err error) {
	log.Println("USBGroup.GetNameDescription not mocked")
	_err = errors.New("USBGroup.GetNameDescription not mocked")
	return
}

var USBGroupClassGetNameDescriptionMockedCallback = USBGroupClassGetNameDescriptionMockDefault

func (_class USBGroupClass) GetNameDescriptionMock(sessionID SessionRef, self USBGroupRef) (_retval string, _err error) {
	return USBGroupClassGetNameDescriptionMockedCallback(sessionID, self)
}
// Get the name/description field of the given USB_group.
func (_class USBGroupClass) GetNameDescription(sessionID SessionRef, self USBGroupRef) (_retval string, _err error) {
	if IsMock {
		return _class.GetNameDescriptionMock(sessionID, self)
	}	
	_method := "USB_group.get_name_description"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertUSBGroupRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func USBGroupClassGetNameLabelMockDefault(sessionID SessionRef, self USBGroupRef) (_retval string, _err error) {
	log.Println("USBGroup.GetNameLabel not mocked")
	_err = errors.New("USBGroup.GetNameLabel not mocked")
	return
}

var USBGroupClassGetNameLabelMockedCallback = USBGroupClassGetNameLabelMockDefault

func (_class USBGroupClass) GetNameLabelMock(sessionID SessionRef, self USBGroupRef) (_retval string, _err error) {
	return USBGroupClassGetNameLabelMockedCallback(sessionID, self)
}
// Get the name/label field of the given USB_group.
func (_class USBGroupClass) GetNameLabel(sessionID SessionRef, self USBGroupRef) (_retval string, _err error) {
	if IsMock {
		return _class.GetNameLabelMock(sessionID, self)
	}	
	_method := "USB_group.get_name_label"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertUSBGroupRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func USBGroupClassGetUUIDMockDefault(sessionID SessionRef, self USBGroupRef) (_retval string, _err error) {
	log.Println("USBGroup.GetUUID not mocked")
	_err = errors.New("USBGroup.GetUUID not mocked")
	return
}

var USBGroupClassGetUUIDMockedCallback = USBGroupClassGetUUIDMockDefault

func (_class USBGroupClass) GetUUIDMock(sessionID SessionRef, self USBGroupRef) (_retval string, _err error) {
	return USBGroupClassGetUUIDMockedCallback(sessionID, self)
}
// Get the uuid field of the given USB_group.
func (_class USBGroupClass) GetUUID(sessionID SessionRef, self USBGroupRef) (_retval string, _err error) {
	if IsMock {
		return _class.GetUUIDMock(sessionID, self)
	}	
	_method := "USB_group.get_uuid"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertUSBGroupRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func USBGroupClassGetByNameLabelMockDefault(sessionID SessionRef, label string) (_retval []USBGroupRef, _err error) {
	log.Println("USBGroup.GetByNameLabel not mocked")
	_err = errors.New("USBGroup.GetByNameLabel not mocked")
	return
}

var USBGroupClassGetByNameLabelMockedCallback = USBGroupClassGetByNameLabelMockDefault

func (_class USBGroupClass) GetByNameLabelMock(sessionID SessionRef, label string) (_retval []USBGroupRef, _err error) {
	return USBGroupClassGetByNameLabelMockedCallback(sessionID, label)
}
// Get all the USB_group instances with the given label.
func (_class USBGroupClass) GetByNameLabel(sessionID SessionRef, label string) (_retval []USBGroupRef, _err error) {
	if IsMock {
		return _class.GetByNameLabelMock(sessionID, label)
	}	
	_method := "USB_group.get_by_name_label"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_labelArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "label"), label)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _labelArg)
	if _err != nil {
		return
	}
	_retval, _err = convertUSBGroupRefSetToGo(_method + " -> ", _result.Value)
	return
}


func USBGroupClassGetByUUIDMockDefault(sessionID SessionRef, uuid string) (_retval USBGroupRef, _err error) {
	log.Println("USBGroup.GetByUUID not mocked")
	_err = errors.New("USBGroup.GetByUUID not mocked")
	return
}

var USBGroupClassGetByUUIDMockedCallback = USBGroupClassGetByUUIDMockDefault

func (_class USBGroupClass) GetByUUIDMock(sessionID SessionRef, uuid string) (_retval USBGroupRef, _err error) {
	return USBGroupClassGetByUUIDMockedCallback(sessionID, uuid)
}
// Get a reference to the USB_group instance with the specified UUID.
func (_class USBGroupClass) GetByUUID(sessionID SessionRef, uuid string) (_retval USBGroupRef, _err error) {
	if IsMock {
		return _class.GetByUUIDMock(sessionID, uuid)
	}	
	_method := "USB_group.get_by_uuid"
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
	_retval, _err = convertUSBGroupRefToGo(_method + " -> ", _result.Value)
	return
}


func USBGroupClassGetRecordMockDefault(sessionID SessionRef, self USBGroupRef) (_retval USBGroupRecord, _err error) {
	log.Println("USBGroup.GetRecord not mocked")
	_err = errors.New("USBGroup.GetRecord not mocked")
	return
}

var USBGroupClassGetRecordMockedCallback = USBGroupClassGetRecordMockDefault

func (_class USBGroupClass) GetRecordMock(sessionID SessionRef, self USBGroupRef) (_retval USBGroupRecord, _err error) {
	return USBGroupClassGetRecordMockedCallback(sessionID, self)
}
// Get a record containing the current state of the given USB_group.
func (_class USBGroupClass) GetRecord(sessionID SessionRef, self USBGroupRef) (_retval USBGroupRecord, _err error) {
	if IsMock {
		return _class.GetRecordMock(sessionID, self)
	}	
	_method := "USB_group.get_record"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertUSBGroupRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertUSBGroupRecordToGo(_method + " -> ", _result.Value)
	return
}
