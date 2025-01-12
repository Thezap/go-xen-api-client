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

type ConsoleProtocol string

const (
  // VT100 terminal
	ConsoleProtocolVt100 ConsoleProtocol = "vt100"
  // Remote FrameBuffer protocol (as used in VNC)
	ConsoleProtocolRfb ConsoleProtocol = "rfb"
  // Remote Desktop Protocol
	ConsoleProtocolRdp ConsoleProtocol = "rdp"
)

type ConsoleRecord struct {
  // Unique identifier/object reference
	UUID string
  // the protocol used by this console
	Protocol ConsoleProtocol
  // URI for the console service
	Location string
  // VM to which this console is attached
	VM VMRef
  // additional configuration
	OtherConfig map[string]string
}

type ConsoleRef string

// A console
type ConsoleClass struct {
	client *Client
}


func ConsoleClassGetAllRecordsMockDefault(sessionID SessionRef) (_retval map[ConsoleRef]ConsoleRecord, _err error) {
	log.Println("Console.GetAllRecords not mocked")
	_err = errors.New("Console.GetAllRecords not mocked")
	return
}

var ConsoleClassGetAllRecordsMockedCallback = ConsoleClassGetAllRecordsMockDefault

func (_class ConsoleClass) GetAllRecordsMock(sessionID SessionRef) (_retval map[ConsoleRef]ConsoleRecord, _err error) {
	return ConsoleClassGetAllRecordsMockedCallback(sessionID)
}
// Return a map of console references to console records for all consoles known to the system.
func (_class ConsoleClass) GetAllRecords(sessionID SessionRef) (_retval map[ConsoleRef]ConsoleRecord, _err error) {
	if IsMock {
		return _class.GetAllRecordsMock(sessionID)
	}	
	_method := "console.get_all_records"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertConsoleRefToConsoleRecordMapToGo(_method + " -> ", _result.Value)
	return
}


func ConsoleClassGetAllMockDefault(sessionID SessionRef) (_retval []ConsoleRef, _err error) {
	log.Println("Console.GetAll not mocked")
	_err = errors.New("Console.GetAll not mocked")
	return
}

var ConsoleClassGetAllMockedCallback = ConsoleClassGetAllMockDefault

func (_class ConsoleClass) GetAllMock(sessionID SessionRef) (_retval []ConsoleRef, _err error) {
	return ConsoleClassGetAllMockedCallback(sessionID)
}
// Return a list of all the consoles known to the system.
func (_class ConsoleClass) GetAll(sessionID SessionRef) (_retval []ConsoleRef, _err error) {
	if IsMock {
		return _class.GetAllMock(sessionID)
	}	
	_method := "console.get_all"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertConsoleRefSetToGo(_method + " -> ", _result.Value)
	return
}


func ConsoleClassRemoveFromOtherConfigMockDefault(sessionID SessionRef, self ConsoleRef, key string) (_err error) {
	log.Println("Console.RemoveFromOtherConfig not mocked")
	_err = errors.New("Console.RemoveFromOtherConfig not mocked")
	return
}

var ConsoleClassRemoveFromOtherConfigMockedCallback = ConsoleClassRemoveFromOtherConfigMockDefault

func (_class ConsoleClass) RemoveFromOtherConfigMock(sessionID SessionRef, self ConsoleRef, key string) (_err error) {
	return ConsoleClassRemoveFromOtherConfigMockedCallback(sessionID, self, key)
}
// Remove the given key and its corresponding value from the other_config field of the given console.  If the key is not in that Map, then do nothing.
func (_class ConsoleClass) RemoveFromOtherConfig(sessionID SessionRef, self ConsoleRef, key string) (_err error) {
	if IsMock {
		return _class.RemoveFromOtherConfigMock(sessionID, self, key)
	}	
	_method := "console.remove_from_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertConsoleRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func ConsoleClassAddToOtherConfigMockDefault(sessionID SessionRef, self ConsoleRef, key string, value string) (_err error) {
	log.Println("Console.AddToOtherConfig not mocked")
	_err = errors.New("Console.AddToOtherConfig not mocked")
	return
}

var ConsoleClassAddToOtherConfigMockedCallback = ConsoleClassAddToOtherConfigMockDefault

func (_class ConsoleClass) AddToOtherConfigMock(sessionID SessionRef, self ConsoleRef, key string, value string) (_err error) {
	return ConsoleClassAddToOtherConfigMockedCallback(sessionID, self, key, value)
}
// Add the given key-value pair to the other_config field of the given console.
func (_class ConsoleClass) AddToOtherConfig(sessionID SessionRef, self ConsoleRef, key string, value string) (_err error) {
	if IsMock {
		return _class.AddToOtherConfigMock(sessionID, self, key, value)
	}	
	_method := "console.add_to_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertConsoleRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func ConsoleClassSetOtherConfigMockDefault(sessionID SessionRef, self ConsoleRef, value map[string]string) (_err error) {
	log.Println("Console.SetOtherConfig not mocked")
	_err = errors.New("Console.SetOtherConfig not mocked")
	return
}

var ConsoleClassSetOtherConfigMockedCallback = ConsoleClassSetOtherConfigMockDefault

func (_class ConsoleClass) SetOtherConfigMock(sessionID SessionRef, self ConsoleRef, value map[string]string) (_err error) {
	return ConsoleClassSetOtherConfigMockedCallback(sessionID, self, value)
}
// Set the other_config field of the given console.
func (_class ConsoleClass) SetOtherConfig(sessionID SessionRef, self ConsoleRef, value map[string]string) (_err error) {
	if IsMock {
		return _class.SetOtherConfigMock(sessionID, self, value)
	}	
	_method := "console.set_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertConsoleRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func ConsoleClassGetOtherConfigMockDefault(sessionID SessionRef, self ConsoleRef) (_retval map[string]string, _err error) {
	log.Println("Console.GetOtherConfig not mocked")
	_err = errors.New("Console.GetOtherConfig not mocked")
	return
}

var ConsoleClassGetOtherConfigMockedCallback = ConsoleClassGetOtherConfigMockDefault

func (_class ConsoleClass) GetOtherConfigMock(sessionID SessionRef, self ConsoleRef) (_retval map[string]string, _err error) {
	return ConsoleClassGetOtherConfigMockedCallback(sessionID, self)
}
// Get the other_config field of the given console.
func (_class ConsoleClass) GetOtherConfig(sessionID SessionRef, self ConsoleRef) (_retval map[string]string, _err error) {
	if IsMock {
		return _class.GetOtherConfigMock(sessionID, self)
	}	
	_method := "console.get_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertConsoleRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func ConsoleClassGetVMMockDefault(sessionID SessionRef, self ConsoleRef) (_retval VMRef, _err error) {
	log.Println("Console.GetVM not mocked")
	_err = errors.New("Console.GetVM not mocked")
	return
}

var ConsoleClassGetVMMockedCallback = ConsoleClassGetVMMockDefault

func (_class ConsoleClass) GetVMMock(sessionID SessionRef, self ConsoleRef) (_retval VMRef, _err error) {
	return ConsoleClassGetVMMockedCallback(sessionID, self)
}
// Get the VM field of the given console.
func (_class ConsoleClass) GetVM(sessionID SessionRef, self ConsoleRef) (_retval VMRef, _err error) {
	if IsMock {
		return _class.GetVMMock(sessionID, self)
	}	
	_method := "console.get_VM"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertConsoleRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func ConsoleClassGetLocationMockDefault(sessionID SessionRef, self ConsoleRef) (_retval string, _err error) {
	log.Println("Console.GetLocation not mocked")
	_err = errors.New("Console.GetLocation not mocked")
	return
}

var ConsoleClassGetLocationMockedCallback = ConsoleClassGetLocationMockDefault

func (_class ConsoleClass) GetLocationMock(sessionID SessionRef, self ConsoleRef) (_retval string, _err error) {
	return ConsoleClassGetLocationMockedCallback(sessionID, self)
}
// Get the location field of the given console.
func (_class ConsoleClass) GetLocation(sessionID SessionRef, self ConsoleRef) (_retval string, _err error) {
	if IsMock {
		return _class.GetLocationMock(sessionID, self)
	}	
	_method := "console.get_location"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertConsoleRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func ConsoleClassGetProtocolMockDefault(sessionID SessionRef, self ConsoleRef) (_retval ConsoleProtocol, _err error) {
	log.Println("Console.GetProtocol not mocked")
	_err = errors.New("Console.GetProtocol not mocked")
	return
}

var ConsoleClassGetProtocolMockedCallback = ConsoleClassGetProtocolMockDefault

func (_class ConsoleClass) GetProtocolMock(sessionID SessionRef, self ConsoleRef) (_retval ConsoleProtocol, _err error) {
	return ConsoleClassGetProtocolMockedCallback(sessionID, self)
}
// Get the protocol field of the given console.
func (_class ConsoleClass) GetProtocol(sessionID SessionRef, self ConsoleRef) (_retval ConsoleProtocol, _err error) {
	if IsMock {
		return _class.GetProtocolMock(sessionID, self)
	}	
	_method := "console.get_protocol"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertConsoleRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertEnumConsoleProtocolToGo(_method + " -> ", _result.Value)
	return
}


func ConsoleClassGetUUIDMockDefault(sessionID SessionRef, self ConsoleRef) (_retval string, _err error) {
	log.Println("Console.GetUUID not mocked")
	_err = errors.New("Console.GetUUID not mocked")
	return
}

var ConsoleClassGetUUIDMockedCallback = ConsoleClassGetUUIDMockDefault

func (_class ConsoleClass) GetUUIDMock(sessionID SessionRef, self ConsoleRef) (_retval string, _err error) {
	return ConsoleClassGetUUIDMockedCallback(sessionID, self)
}
// Get the uuid field of the given console.
func (_class ConsoleClass) GetUUID(sessionID SessionRef, self ConsoleRef) (_retval string, _err error) {
	if IsMock {
		return _class.GetUUIDMock(sessionID, self)
	}	
	_method := "console.get_uuid"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertConsoleRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func ConsoleClassDestroyMockDefault(sessionID SessionRef, self ConsoleRef) (_err error) {
	log.Println("Console.Destroy not mocked")
	_err = errors.New("Console.Destroy not mocked")
	return
}

var ConsoleClassDestroyMockedCallback = ConsoleClassDestroyMockDefault

func (_class ConsoleClass) DestroyMock(sessionID SessionRef, self ConsoleRef) (_err error) {
	return ConsoleClassDestroyMockedCallback(sessionID, self)
}
// Destroy the specified console instance.
func (_class ConsoleClass) Destroy(sessionID SessionRef, self ConsoleRef) (_err error) {
	if IsMock {
		return _class.DestroyMock(sessionID, self)
	}	
	_method := "console.destroy"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertConsoleRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg)
	return
}


func ConsoleClassCreateMockDefault(sessionID SessionRef, args ConsoleRecord) (_retval ConsoleRef, _err error) {
	log.Println("Console.Create not mocked")
	_err = errors.New("Console.Create not mocked")
	return
}

var ConsoleClassCreateMockedCallback = ConsoleClassCreateMockDefault

func (_class ConsoleClass) CreateMock(sessionID SessionRef, args ConsoleRecord) (_retval ConsoleRef, _err error) {
	return ConsoleClassCreateMockedCallback(sessionID, args)
}
// Create a new console instance, and return its handle.
// The constructor args are: other_config* (* = non-optional).
func (_class ConsoleClass) Create(sessionID SessionRef, args ConsoleRecord) (_retval ConsoleRef, _err error) {
	if IsMock {
		return _class.CreateMock(sessionID, args)
	}	
	_method := "console.create"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_argsArg, _err := convertConsoleRecordToXen(fmt.Sprintf("%s(%s)", _method, "args"), args)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _argsArg)
	if _err != nil {
		return
	}
	_retval, _err = convertConsoleRefToGo(_method + " -> ", _result.Value)
	return
}


func ConsoleClassGetByUUIDMockDefault(sessionID SessionRef, uuid string) (_retval ConsoleRef, _err error) {
	log.Println("Console.GetByUUID not mocked")
	_err = errors.New("Console.GetByUUID not mocked")
	return
}

var ConsoleClassGetByUUIDMockedCallback = ConsoleClassGetByUUIDMockDefault

func (_class ConsoleClass) GetByUUIDMock(sessionID SessionRef, uuid string) (_retval ConsoleRef, _err error) {
	return ConsoleClassGetByUUIDMockedCallback(sessionID, uuid)
}
// Get a reference to the console instance with the specified UUID.
func (_class ConsoleClass) GetByUUID(sessionID SessionRef, uuid string) (_retval ConsoleRef, _err error) {
	if IsMock {
		return _class.GetByUUIDMock(sessionID, uuid)
	}	
	_method := "console.get_by_uuid"
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
	_retval, _err = convertConsoleRefToGo(_method + " -> ", _result.Value)
	return
}


func ConsoleClassGetRecordMockDefault(sessionID SessionRef, self ConsoleRef) (_retval ConsoleRecord, _err error) {
	log.Println("Console.GetRecord not mocked")
	_err = errors.New("Console.GetRecord not mocked")
	return
}

var ConsoleClassGetRecordMockedCallback = ConsoleClassGetRecordMockDefault

func (_class ConsoleClass) GetRecordMock(sessionID SessionRef, self ConsoleRef) (_retval ConsoleRecord, _err error) {
	return ConsoleClassGetRecordMockedCallback(sessionID, self)
}
// Get a record containing the current state of the given console.
func (_class ConsoleClass) GetRecord(sessionID SessionRef, self ConsoleRef) (_retval ConsoleRecord, _err error) {
	if IsMock {
		return _class.GetRecordMock(sessionID, self)
	}	
	_method := "console.get_record"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertConsoleRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertConsoleRecordToGo(_method + " -> ", _result.Value)
	return
}
