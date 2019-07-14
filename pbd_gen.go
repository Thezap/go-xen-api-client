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

type PBDRecord struct {
  // Unique identifier/object reference
	UUID string
  // physical machine on which the pbd is available
	Host HostRef
  // the storage repository that the pbd realises
	SR SRRef
  // a config string to string map that is provided to the host's SR-backend-driver
	DeviceConfig map[string]string
  // is the SR currently attached on this host?
	CurrentlyAttached bool
  // additional configuration
	OtherConfig map[string]string
}

type PBDRef string

// The physical block devices through which hosts access SRs
type PBDClass struct {
	client *Client
}


func PBDClassGetAllRecordsMockDefault(sessionID SessionRef) (_retval map[PBDRef]PBDRecord, _err error) {
	log.Println("PBD.GetAllRecords not mocked")
	_err = errors.New("PBD.GetAllRecords not mocked")
	return
}

var PBDClassGetAllRecordsMockedCallback = PBDClassGetAllRecordsMockDefault

func (_class PBDClass) GetAllRecordsMock(sessionID SessionRef) (_retval map[PBDRef]PBDRecord, _err error) {
	return PBDClassGetAllRecordsMockedCallback(sessionID)
}
// Return a map of PBD references to PBD records for all PBDs known to the system.
func (_class PBDClass) GetAllRecords(sessionID SessionRef) (_retval map[PBDRef]PBDRecord, _err error) {
	if IsMock {
		return _class.GetAllRecordsMock(sessionID)
	}	
	_method := "PBD.get_all_records"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertPBDRefToPBDRecordMapToGo(_method + " -> ", _result.Value)
	return
}


func PBDClassGetAllMockDefault(sessionID SessionRef) (_retval []PBDRef, _err error) {
	log.Println("PBD.GetAll not mocked")
	_err = errors.New("PBD.GetAll not mocked")
	return
}

var PBDClassGetAllMockedCallback = PBDClassGetAllMockDefault

func (_class PBDClass) GetAllMock(sessionID SessionRef) (_retval []PBDRef, _err error) {
	return PBDClassGetAllMockedCallback(sessionID)
}
// Return a list of all the PBDs known to the system.
func (_class PBDClass) GetAll(sessionID SessionRef) (_retval []PBDRef, _err error) {
	if IsMock {
		return _class.GetAllMock(sessionID)
	}	
	_method := "PBD.get_all"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertPBDRefSetToGo(_method + " -> ", _result.Value)
	return
}


func PBDClassSetDeviceConfigMockDefault(sessionID SessionRef, self PBDRef, value map[string]string) (_err error) {
	log.Println("PBD.SetDeviceConfig not mocked")
	_err = errors.New("PBD.SetDeviceConfig not mocked")
	return
}

var PBDClassSetDeviceConfigMockedCallback = PBDClassSetDeviceConfigMockDefault

func (_class PBDClass) SetDeviceConfigMock(sessionID SessionRef, self PBDRef, value map[string]string) (_err error) {
	return PBDClassSetDeviceConfigMockedCallback(sessionID, self, value)
}
// Sets the PBD's device_config field
func (_class PBDClass) SetDeviceConfig(sessionID SessionRef, self PBDRef, value map[string]string) (_err error) {
	if IsMock {
		return _class.SetDeviceConfigMock(sessionID, self, value)
	}	
	_method := "PBD.set_device_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPBDRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func PBDClassUnplugMockDefault(sessionID SessionRef, self PBDRef) (_err error) {
	log.Println("PBD.Unplug not mocked")
	_err = errors.New("PBD.Unplug not mocked")
	return
}

var PBDClassUnplugMockedCallback = PBDClassUnplugMockDefault

func (_class PBDClass) UnplugMock(sessionID SessionRef, self PBDRef) (_err error) {
	return PBDClassUnplugMockedCallback(sessionID, self)
}
// Deactivate the specified PBD, causing the referenced SR to be detached and nolonger scanned
func (_class PBDClass) Unplug(sessionID SessionRef, self PBDRef) (_err error) {
	if IsMock {
		return _class.UnplugMock(sessionID, self)
	}	
	_method := "PBD.unplug"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPBDRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg)
	return
}


func PBDClassPlugMockDefault(sessionID SessionRef, self PBDRef) (_err error) {
	log.Println("PBD.Plug not mocked")
	_err = errors.New("PBD.Plug not mocked")
	return
}

var PBDClassPlugMockedCallback = PBDClassPlugMockDefault

func (_class PBDClass) PlugMock(sessionID SessionRef, self PBDRef) (_err error) {
	return PBDClassPlugMockedCallback(sessionID, self)
}
// Activate the specified PBD, causing the referenced SR to be attached and scanned
//
// Errors:
//  SR_UNKNOWN_DRIVER - The SR could not be connected because the driver was not recognised.
func (_class PBDClass) Plug(sessionID SessionRef, self PBDRef) (_err error) {
	if IsMock {
		return _class.PlugMock(sessionID, self)
	}	
	_method := "PBD.plug"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPBDRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg)
	return
}


func PBDClassRemoveFromOtherConfigMockDefault(sessionID SessionRef, self PBDRef, key string) (_err error) {
	log.Println("PBD.RemoveFromOtherConfig not mocked")
	_err = errors.New("PBD.RemoveFromOtherConfig not mocked")
	return
}

var PBDClassRemoveFromOtherConfigMockedCallback = PBDClassRemoveFromOtherConfigMockDefault

func (_class PBDClass) RemoveFromOtherConfigMock(sessionID SessionRef, self PBDRef, key string) (_err error) {
	return PBDClassRemoveFromOtherConfigMockedCallback(sessionID, self, key)
}
// Remove the given key and its corresponding value from the other_config field of the given PBD.  If the key is not in that Map, then do nothing.
func (_class PBDClass) RemoveFromOtherConfig(sessionID SessionRef, self PBDRef, key string) (_err error) {
	if IsMock {
		return _class.RemoveFromOtherConfigMock(sessionID, self, key)
	}	
	_method := "PBD.remove_from_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPBDRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func PBDClassAddToOtherConfigMockDefault(sessionID SessionRef, self PBDRef, key string, value string) (_err error) {
	log.Println("PBD.AddToOtherConfig not mocked")
	_err = errors.New("PBD.AddToOtherConfig not mocked")
	return
}

var PBDClassAddToOtherConfigMockedCallback = PBDClassAddToOtherConfigMockDefault

func (_class PBDClass) AddToOtherConfigMock(sessionID SessionRef, self PBDRef, key string, value string) (_err error) {
	return PBDClassAddToOtherConfigMockedCallback(sessionID, self, key, value)
}
// Add the given key-value pair to the other_config field of the given PBD.
func (_class PBDClass) AddToOtherConfig(sessionID SessionRef, self PBDRef, key string, value string) (_err error) {
	if IsMock {
		return _class.AddToOtherConfigMock(sessionID, self, key, value)
	}	
	_method := "PBD.add_to_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPBDRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func PBDClassSetOtherConfigMockDefault(sessionID SessionRef, self PBDRef, value map[string]string) (_err error) {
	log.Println("PBD.SetOtherConfig not mocked")
	_err = errors.New("PBD.SetOtherConfig not mocked")
	return
}

var PBDClassSetOtherConfigMockedCallback = PBDClassSetOtherConfigMockDefault

func (_class PBDClass) SetOtherConfigMock(sessionID SessionRef, self PBDRef, value map[string]string) (_err error) {
	return PBDClassSetOtherConfigMockedCallback(sessionID, self, value)
}
// Set the other_config field of the given PBD.
func (_class PBDClass) SetOtherConfig(sessionID SessionRef, self PBDRef, value map[string]string) (_err error) {
	if IsMock {
		return _class.SetOtherConfigMock(sessionID, self, value)
	}	
	_method := "PBD.set_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPBDRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func PBDClassGetOtherConfigMockDefault(sessionID SessionRef, self PBDRef) (_retval map[string]string, _err error) {
	log.Println("PBD.GetOtherConfig not mocked")
	_err = errors.New("PBD.GetOtherConfig not mocked")
	return
}

var PBDClassGetOtherConfigMockedCallback = PBDClassGetOtherConfigMockDefault

func (_class PBDClass) GetOtherConfigMock(sessionID SessionRef, self PBDRef) (_retval map[string]string, _err error) {
	return PBDClassGetOtherConfigMockedCallback(sessionID, self)
}
// Get the other_config field of the given PBD.
func (_class PBDClass) GetOtherConfig(sessionID SessionRef, self PBDRef) (_retval map[string]string, _err error) {
	if IsMock {
		return _class.GetOtherConfigMock(sessionID, self)
	}	
	_method := "PBD.get_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPBDRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func PBDClassGetCurrentlyAttachedMockDefault(sessionID SessionRef, self PBDRef) (_retval bool, _err error) {
	log.Println("PBD.GetCurrentlyAttached not mocked")
	_err = errors.New("PBD.GetCurrentlyAttached not mocked")
	return
}

var PBDClassGetCurrentlyAttachedMockedCallback = PBDClassGetCurrentlyAttachedMockDefault

func (_class PBDClass) GetCurrentlyAttachedMock(sessionID SessionRef, self PBDRef) (_retval bool, _err error) {
	return PBDClassGetCurrentlyAttachedMockedCallback(sessionID, self)
}
// Get the currently_attached field of the given PBD.
func (_class PBDClass) GetCurrentlyAttached(sessionID SessionRef, self PBDRef) (_retval bool, _err error) {
	if IsMock {
		return _class.GetCurrentlyAttachedMock(sessionID, self)
	}	
	_method := "PBD.get_currently_attached"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPBDRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func PBDClassGetDeviceConfigMockDefault(sessionID SessionRef, self PBDRef) (_retval map[string]string, _err error) {
	log.Println("PBD.GetDeviceConfig not mocked")
	_err = errors.New("PBD.GetDeviceConfig not mocked")
	return
}

var PBDClassGetDeviceConfigMockedCallback = PBDClassGetDeviceConfigMockDefault

func (_class PBDClass) GetDeviceConfigMock(sessionID SessionRef, self PBDRef) (_retval map[string]string, _err error) {
	return PBDClassGetDeviceConfigMockedCallback(sessionID, self)
}
// Get the device_config field of the given PBD.
func (_class PBDClass) GetDeviceConfig(sessionID SessionRef, self PBDRef) (_retval map[string]string, _err error) {
	if IsMock {
		return _class.GetDeviceConfigMock(sessionID, self)
	}	
	_method := "PBD.get_device_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPBDRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func PBDClassGetSRMockDefault(sessionID SessionRef, self PBDRef) (_retval SRRef, _err error) {
	log.Println("PBD.GetSR not mocked")
	_err = errors.New("PBD.GetSR not mocked")
	return
}

var PBDClassGetSRMockedCallback = PBDClassGetSRMockDefault

func (_class PBDClass) GetSRMock(sessionID SessionRef, self PBDRef) (_retval SRRef, _err error) {
	return PBDClassGetSRMockedCallback(sessionID, self)
}
// Get the SR field of the given PBD.
func (_class PBDClass) GetSR(sessionID SessionRef, self PBDRef) (_retval SRRef, _err error) {
	if IsMock {
		return _class.GetSRMock(sessionID, self)
	}	
	_method := "PBD.get_SR"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPBDRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertSRRefToGo(_method + " -> ", _result.Value)
	return
}


func PBDClassGetHostMockDefault(sessionID SessionRef, self PBDRef) (_retval HostRef, _err error) {
	log.Println("PBD.GetHost not mocked")
	_err = errors.New("PBD.GetHost not mocked")
	return
}

var PBDClassGetHostMockedCallback = PBDClassGetHostMockDefault

func (_class PBDClass) GetHostMock(sessionID SessionRef, self PBDRef) (_retval HostRef, _err error) {
	return PBDClassGetHostMockedCallback(sessionID, self)
}
// Get the host field of the given PBD.
func (_class PBDClass) GetHost(sessionID SessionRef, self PBDRef) (_retval HostRef, _err error) {
	if IsMock {
		return _class.GetHostMock(sessionID, self)
	}	
	_method := "PBD.get_host"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPBDRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func PBDClassGetUUIDMockDefault(sessionID SessionRef, self PBDRef) (_retval string, _err error) {
	log.Println("PBD.GetUUID not mocked")
	_err = errors.New("PBD.GetUUID not mocked")
	return
}

var PBDClassGetUUIDMockedCallback = PBDClassGetUUIDMockDefault

func (_class PBDClass) GetUUIDMock(sessionID SessionRef, self PBDRef) (_retval string, _err error) {
	return PBDClassGetUUIDMockedCallback(sessionID, self)
}
// Get the uuid field of the given PBD.
func (_class PBDClass) GetUUID(sessionID SessionRef, self PBDRef) (_retval string, _err error) {
	if IsMock {
		return _class.GetUUIDMock(sessionID, self)
	}	
	_method := "PBD.get_uuid"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPBDRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func PBDClassDestroyMockDefault(sessionID SessionRef, self PBDRef) (_err error) {
	log.Println("PBD.Destroy not mocked")
	_err = errors.New("PBD.Destroy not mocked")
	return
}

var PBDClassDestroyMockedCallback = PBDClassDestroyMockDefault

func (_class PBDClass) DestroyMock(sessionID SessionRef, self PBDRef) (_err error) {
	return PBDClassDestroyMockedCallback(sessionID, self)
}
// Destroy the specified PBD instance.
func (_class PBDClass) Destroy(sessionID SessionRef, self PBDRef) (_err error) {
	if IsMock {
		return _class.DestroyMock(sessionID, self)
	}	
	_method := "PBD.destroy"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPBDRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg)
	return
}


func PBDClassCreateMockDefault(sessionID SessionRef, args PBDRecord) (_retval PBDRef, _err error) {
	log.Println("PBD.Create not mocked")
	_err = errors.New("PBD.Create not mocked")
	return
}

var PBDClassCreateMockedCallback = PBDClassCreateMockDefault

func (_class PBDClass) CreateMock(sessionID SessionRef, args PBDRecord) (_retval PBDRef, _err error) {
	return PBDClassCreateMockedCallback(sessionID, args)
}
// Create a new PBD instance, and return its handle.
// The constructor args are: host*, SR*, device_config*, other_config (* = non-optional).
func (_class PBDClass) Create(sessionID SessionRef, args PBDRecord) (_retval PBDRef, _err error) {
	if IsMock {
		return _class.CreateMock(sessionID, args)
	}	
	_method := "PBD.create"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_argsArg, _err := convertPBDRecordToXen(fmt.Sprintf("%s(%s)", _method, "args"), args)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _argsArg)
	if _err != nil {
		return
	}
	_retval, _err = convertPBDRefToGo(_method + " -> ", _result.Value)
	return
}


func PBDClassGetByUUIDMockDefault(sessionID SessionRef, uuid string) (_retval PBDRef, _err error) {
	log.Println("PBD.GetByUUID not mocked")
	_err = errors.New("PBD.GetByUUID not mocked")
	return
}

var PBDClassGetByUUIDMockedCallback = PBDClassGetByUUIDMockDefault

func (_class PBDClass) GetByUUIDMock(sessionID SessionRef, uuid string) (_retval PBDRef, _err error) {
	return PBDClassGetByUUIDMockedCallback(sessionID, uuid)
}
// Get a reference to the PBD instance with the specified UUID.
func (_class PBDClass) GetByUUID(sessionID SessionRef, uuid string) (_retval PBDRef, _err error) {
	if IsMock {
		return _class.GetByUUIDMock(sessionID, uuid)
	}	
	_method := "PBD.get_by_uuid"
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
	_retval, _err = convertPBDRefToGo(_method + " -> ", _result.Value)
	return
}


func PBDClassGetRecordMockDefault(sessionID SessionRef, self PBDRef) (_retval PBDRecord, _err error) {
	log.Println("PBD.GetRecord not mocked")
	_err = errors.New("PBD.GetRecord not mocked")
	return
}

var PBDClassGetRecordMockedCallback = PBDClassGetRecordMockDefault

func (_class PBDClass) GetRecordMock(sessionID SessionRef, self PBDRef) (_retval PBDRecord, _err error) {
	return PBDClassGetRecordMockedCallback(sessionID, self)
}
// Get a record containing the current state of the given PBD.
func (_class PBDClass) GetRecord(sessionID SessionRef, self PBDRef) (_retval PBDRecord, _err error) {
	if IsMock {
		return _class.GetRecordMock(sessionID, self)
	}	
	_method := "PBD.get_record"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPBDRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertPBDRecordToGo(_method + " -> ", _result.Value)
	return
}
