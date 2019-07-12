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

func (_class USBGroupClass) GetAllRecords__mock(sessionID SessionRef) (_retval map[USBGroupRef]USBGroupRecord, _err error) {
	log.Println("USBGroup.GetAllRecords not mocked")
	_err = errors.New("USBGroup.GetAllRecords not mocked")
	return
}
// Return a map of USB_group references to USB_group records for all USB_groups known to the system.
func (_class USBGroupClass) GetAllRecords(sessionID SessionRef) (_retval map[USBGroupRef]USBGroupRecord, _err error) {
	if (IsMock) {
		return _class.GetAllRecords__mock(sessionID)
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

func (_class USBGroupClass) GetAll__mock(sessionID SessionRef) (_retval []USBGroupRef, _err error) {
	log.Println("USBGroup.GetAll not mocked")
	_err = errors.New("USBGroup.GetAll not mocked")
	return
}
// Return a list of all the USB_groups known to the system.
func (_class USBGroupClass) GetAll(sessionID SessionRef) (_retval []USBGroupRef, _err error) {
	if (IsMock) {
		return _class.GetAll__mock(sessionID)
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

func (_class USBGroupClass) Destroy__mock(sessionID SessionRef, self USBGroupRef) (_err error) {
	log.Println("USBGroup.Destroy not mocked")
	_err = errors.New("USBGroup.Destroy not mocked")
	return
}
// 
func (_class USBGroupClass) Destroy(sessionID SessionRef, self USBGroupRef) (_err error) {
	if (IsMock) {
		return _class.Destroy__mock(sessionID, self)
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

func (_class USBGroupClass) Create__mock(sessionID SessionRef, nameLabel string, nameDescription string, otherConfig map[string]string) (_retval USBGroupRef, _err error) {
	log.Println("USBGroup.Create not mocked")
	_err = errors.New("USBGroup.Create not mocked")
	return
}
// 
func (_class USBGroupClass) Create(sessionID SessionRef, nameLabel string, nameDescription string, otherConfig map[string]string) (_retval USBGroupRef, _err error) {
	if (IsMock) {
		return _class.Create__mock(sessionID, nameLabel, nameDescription, otherConfig)
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

func (_class USBGroupClass) RemoveFromOtherConfig__mock(sessionID SessionRef, self USBGroupRef, key string) (_err error) {
	log.Println("USBGroup.RemoveFromOtherConfig not mocked")
	_err = errors.New("USBGroup.RemoveFromOtherConfig not mocked")
	return
}
// Remove the given key and its corresponding value from the other_config field of the given USB_group.  If the key is not in that Map, then do nothing.
func (_class USBGroupClass) RemoveFromOtherConfig(sessionID SessionRef, self USBGroupRef, key string) (_err error) {
	if (IsMock) {
		return _class.RemoveFromOtherConfig__mock(sessionID, self, key)
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

func (_class USBGroupClass) AddToOtherConfig__mock(sessionID SessionRef, self USBGroupRef, key string, value string) (_err error) {
	log.Println("USBGroup.AddToOtherConfig not mocked")
	_err = errors.New("USBGroup.AddToOtherConfig not mocked")
	return
}
// Add the given key-value pair to the other_config field of the given USB_group.
func (_class USBGroupClass) AddToOtherConfig(sessionID SessionRef, self USBGroupRef, key string, value string) (_err error) {
	if (IsMock) {
		return _class.AddToOtherConfig__mock(sessionID, self, key, value)
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

func (_class USBGroupClass) SetOtherConfig__mock(sessionID SessionRef, self USBGroupRef, value map[string]string) (_err error) {
	log.Println("USBGroup.SetOtherConfig not mocked")
	_err = errors.New("USBGroup.SetOtherConfig not mocked")
	return
}
// Set the other_config field of the given USB_group.
func (_class USBGroupClass) SetOtherConfig(sessionID SessionRef, self USBGroupRef, value map[string]string) (_err error) {
	if (IsMock) {
		return _class.SetOtherConfig__mock(sessionID, self, value)
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

func (_class USBGroupClass) SetNameDescription__mock(sessionID SessionRef, self USBGroupRef, value string) (_err error) {
	log.Println("USBGroup.SetNameDescription not mocked")
	_err = errors.New("USBGroup.SetNameDescription not mocked")
	return
}
// Set the name/description field of the given USB_group.
func (_class USBGroupClass) SetNameDescription(sessionID SessionRef, self USBGroupRef, value string) (_err error) {
	if (IsMock) {
		return _class.SetNameDescription__mock(sessionID, self, value)
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

func (_class USBGroupClass) SetNameLabel__mock(sessionID SessionRef, self USBGroupRef, value string) (_err error) {
	log.Println("USBGroup.SetNameLabel not mocked")
	_err = errors.New("USBGroup.SetNameLabel not mocked")
	return
}
// Set the name/label field of the given USB_group.
func (_class USBGroupClass) SetNameLabel(sessionID SessionRef, self USBGroupRef, value string) (_err error) {
	if (IsMock) {
		return _class.SetNameLabel__mock(sessionID, self, value)
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

func (_class USBGroupClass) GetOtherConfig__mock(sessionID SessionRef, self USBGroupRef) (_retval map[string]string, _err error) {
	log.Println("USBGroup.GetOtherConfig not mocked")
	_err = errors.New("USBGroup.GetOtherConfig not mocked")
	return
}
// Get the other_config field of the given USB_group.
func (_class USBGroupClass) GetOtherConfig(sessionID SessionRef, self USBGroupRef) (_retval map[string]string, _err error) {
	if (IsMock) {
		return _class.GetOtherConfig__mock(sessionID, self)
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

func (_class USBGroupClass) GetVUSBs__mock(sessionID SessionRef, self USBGroupRef) (_retval []VUSBRef, _err error) {
	log.Println("USBGroup.GetVUSBs not mocked")
	_err = errors.New("USBGroup.GetVUSBs not mocked")
	return
}
// Get the VUSBs field of the given USB_group.
func (_class USBGroupClass) GetVUSBs(sessionID SessionRef, self USBGroupRef) (_retval []VUSBRef, _err error) {
	if (IsMock) {
		return _class.GetVUSBs__mock(sessionID, self)
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

func (_class USBGroupClass) GetPUSBs__mock(sessionID SessionRef, self USBGroupRef) (_retval []PUSBRef, _err error) {
	log.Println("USBGroup.GetPUSBs not mocked")
	_err = errors.New("USBGroup.GetPUSBs not mocked")
	return
}
// Get the PUSBs field of the given USB_group.
func (_class USBGroupClass) GetPUSBs(sessionID SessionRef, self USBGroupRef) (_retval []PUSBRef, _err error) {
	if (IsMock) {
		return _class.GetPUSBs__mock(sessionID, self)
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

func (_class USBGroupClass) GetNameDescription__mock(sessionID SessionRef, self USBGroupRef) (_retval string, _err error) {
	log.Println("USBGroup.GetNameDescription not mocked")
	_err = errors.New("USBGroup.GetNameDescription not mocked")
	return
}
// Get the name/description field of the given USB_group.
func (_class USBGroupClass) GetNameDescription(sessionID SessionRef, self USBGroupRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetNameDescription__mock(sessionID, self)
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

func (_class USBGroupClass) GetNameLabel__mock(sessionID SessionRef, self USBGroupRef) (_retval string, _err error) {
	log.Println("USBGroup.GetNameLabel not mocked")
	_err = errors.New("USBGroup.GetNameLabel not mocked")
	return
}
// Get the name/label field of the given USB_group.
func (_class USBGroupClass) GetNameLabel(sessionID SessionRef, self USBGroupRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetNameLabel__mock(sessionID, self)
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

func (_class USBGroupClass) GetUUID__mock(sessionID SessionRef, self USBGroupRef) (_retval string, _err error) {
	log.Println("USBGroup.GetUUID not mocked")
	_err = errors.New("USBGroup.GetUUID not mocked")
	return
}
// Get the uuid field of the given USB_group.
func (_class USBGroupClass) GetUUID(sessionID SessionRef, self USBGroupRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetUUID__mock(sessionID, self)
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

func (_class USBGroupClass) GetByNameLabel__mock(sessionID SessionRef, label string) (_retval []USBGroupRef, _err error) {
	log.Println("USBGroup.GetByNameLabel not mocked")
	_err = errors.New("USBGroup.GetByNameLabel not mocked")
	return
}
// Get all the USB_group instances with the given label.
func (_class USBGroupClass) GetByNameLabel(sessionID SessionRef, label string) (_retval []USBGroupRef, _err error) {
	if (IsMock) {
		return _class.GetByNameLabel__mock(sessionID, label)
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

func (_class USBGroupClass) GetByUUID__mock(sessionID SessionRef, uuid string) (_retval USBGroupRef, _err error) {
	log.Println("USBGroup.GetByUUID not mocked")
	_err = errors.New("USBGroup.GetByUUID not mocked")
	return
}
// Get a reference to the USB_group instance with the specified UUID.
func (_class USBGroupClass) GetByUUID(sessionID SessionRef, uuid string) (_retval USBGroupRef, _err error) {
	if (IsMock) {
		return _class.GetByUUID__mock(sessionID, uuid)
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

func (_class USBGroupClass) GetRecord__mock(sessionID SessionRef, self USBGroupRef) (_retval USBGroupRecord, _err error) {
	log.Println("USBGroup.GetRecord not mocked")
	_err = errors.New("USBGroup.GetRecord not mocked")
	return
}
// Get a record containing the current state of the given USB_group.
func (_class USBGroupClass) GetRecord(sessionID SessionRef, self USBGroupRef) (_retval USBGroupRecord, _err error) {
	if (IsMock) {
		return _class.GetRecord__mock(sessionID, self)
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
