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

type SdnControllerProtocol string

const (
  // Active ssl connection
	SdnControllerProtocolSsl SdnControllerProtocol = "ssl"
  // Passive ssl connection
	SdnControllerProtocolPssl SdnControllerProtocol = "pssl"
)

type SDNControllerRecord struct {
  // Unique identifier/object reference
	UUID string
  // Protocol to connect with SDN controller
	Protocol SdnControllerProtocol
  // IP address of the controller
	Address string
  // TCP port of the controller
	Port int
}

type SDNControllerRef string

// Describes the SDN controller that is to connect with the pool
type SDNControllerClass struct {
	client *Client
}


var SDNControllerClass_GetAllRecordsMockedCallback = func (sessionID SessionRef) (_retval map[SDNControllerRef]SDNControllerRecord, _err error) {
	log.Println("SDNController.GetAllRecords not mocked")
	_err = errors.New("SDNController.GetAllRecords not mocked")
	return
}

func (_class SDNControllerClass) GetAllRecordsMock(sessionID SessionRef) (_retval map[SDNControllerRef]SDNControllerRecord, _err error) {
	return SDNControllerClass_GetAllRecordsMockedCallback(sessionID)
}
// Return a map of SDN_controller references to SDN_controller records for all SDN_controllers known to the system.
func (_class SDNControllerClass) GetAllRecords(sessionID SessionRef) (_retval map[SDNControllerRef]SDNControllerRecord, _err error) {
	if (IsMock) {
		return _class.GetAllRecordsMock(sessionID)
	}	
	_method := "SDN_controller.get_all_records"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertSDNControllerRefToSDNControllerRecordMapToGo(_method + " -> ", _result.Value)
	return
}


var SDNControllerClass_GetAllMockedCallback = func (sessionID SessionRef) (_retval []SDNControllerRef, _err error) {
	log.Println("SDNController.GetAll not mocked")
	_err = errors.New("SDNController.GetAll not mocked")
	return
}

func (_class SDNControllerClass) GetAllMock(sessionID SessionRef) (_retval []SDNControllerRef, _err error) {
	return SDNControllerClass_GetAllMockedCallback(sessionID)
}
// Return a list of all the SDN_controllers known to the system.
func (_class SDNControllerClass) GetAll(sessionID SessionRef) (_retval []SDNControllerRef, _err error) {
	if (IsMock) {
		return _class.GetAllMock(sessionID)
	}	
	_method := "SDN_controller.get_all"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertSDNControllerRefSetToGo(_method + " -> ", _result.Value)
	return
}


var SDNControllerClass_ForgetMockedCallback = func (sessionID SessionRef, self SDNControllerRef) (_err error) {
	log.Println("SDNController.Forget not mocked")
	_err = errors.New("SDNController.Forget not mocked")
	return
}

func (_class SDNControllerClass) ForgetMock(sessionID SessionRef, self SDNControllerRef) (_err error) {
	return SDNControllerClass_ForgetMockedCallback(sessionID, self)
}
// Remove the OVS manager of the pool and destroy the db record.
func (_class SDNControllerClass) Forget(sessionID SessionRef, self SDNControllerRef) (_err error) {
	if (IsMock) {
		return _class.ForgetMock(sessionID, self)
	}	
	_method := "SDN_controller.forget"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertSDNControllerRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg)
	return
}


var SDNControllerClass_IntroduceMockedCallback = func (sessionID SessionRef, protocol SdnControllerProtocol, address string, port int) (_retval SDNControllerRef, _err error) {
	log.Println("SDNController.Introduce not mocked")
	_err = errors.New("SDNController.Introduce not mocked")
	return
}

func (_class SDNControllerClass) IntroduceMock(sessionID SessionRef, protocol SdnControllerProtocol, address string, port int) (_retval SDNControllerRef, _err error) {
	return SDNControllerClass_IntroduceMockedCallback(sessionID, protocol, address, port)
}
// Introduce an SDN controller to the pool.
func (_class SDNControllerClass) Introduce(sessionID SessionRef, protocol SdnControllerProtocol, address string, port int) (_retval SDNControllerRef, _err error) {
	if (IsMock) {
		return _class.IntroduceMock(sessionID, protocol, address, port)
	}	
	_method := "SDN_controller.introduce"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_protocolArg, _err := convertEnumSdnControllerProtocolToXen(fmt.Sprintf("%s(%s)", _method, "protocol"), protocol)
	if _err != nil {
		return
	}
	_addressArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "address"), address)
	if _err != nil {
		return
	}
	_portArg, _err := convertIntToXen(fmt.Sprintf("%s(%s)", _method, "port"), port)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _protocolArg, _addressArg, _portArg)
	if _err != nil {
		return
	}
	_retval, _err = convertSDNControllerRefToGo(_method + " -> ", _result.Value)
	return
}


var SDNControllerClass_GetPortMockedCallback = func (sessionID SessionRef, self SDNControllerRef) (_retval int, _err error) {
	log.Println("SDNController.GetPort not mocked")
	_err = errors.New("SDNController.GetPort not mocked")
	return
}

func (_class SDNControllerClass) GetPortMock(sessionID SessionRef, self SDNControllerRef) (_retval int, _err error) {
	return SDNControllerClass_GetPortMockedCallback(sessionID, self)
}
// Get the port field of the given SDN_controller.
func (_class SDNControllerClass) GetPort(sessionID SessionRef, self SDNControllerRef) (_retval int, _err error) {
	if (IsMock) {
		return _class.GetPortMock(sessionID, self)
	}	
	_method := "SDN_controller.get_port"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertSDNControllerRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var SDNControllerClass_GetAddressMockedCallback = func (sessionID SessionRef, self SDNControllerRef) (_retval string, _err error) {
	log.Println("SDNController.GetAddress not mocked")
	_err = errors.New("SDNController.GetAddress not mocked")
	return
}

func (_class SDNControllerClass) GetAddressMock(sessionID SessionRef, self SDNControllerRef) (_retval string, _err error) {
	return SDNControllerClass_GetAddressMockedCallback(sessionID, self)
}
// Get the address field of the given SDN_controller.
func (_class SDNControllerClass) GetAddress(sessionID SessionRef, self SDNControllerRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetAddressMock(sessionID, self)
	}	
	_method := "SDN_controller.get_address"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertSDNControllerRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var SDNControllerClass_GetProtocolMockedCallback = func (sessionID SessionRef, self SDNControllerRef) (_retval SdnControllerProtocol, _err error) {
	log.Println("SDNController.GetProtocol not mocked")
	_err = errors.New("SDNController.GetProtocol not mocked")
	return
}

func (_class SDNControllerClass) GetProtocolMock(sessionID SessionRef, self SDNControllerRef) (_retval SdnControllerProtocol, _err error) {
	return SDNControllerClass_GetProtocolMockedCallback(sessionID, self)
}
// Get the protocol field of the given SDN_controller.
func (_class SDNControllerClass) GetProtocol(sessionID SessionRef, self SDNControllerRef) (_retval SdnControllerProtocol, _err error) {
	if (IsMock) {
		return _class.GetProtocolMock(sessionID, self)
	}	
	_method := "SDN_controller.get_protocol"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertSDNControllerRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertEnumSdnControllerProtocolToGo(_method + " -> ", _result.Value)
	return
}


var SDNControllerClass_GetUUIDMockedCallback = func (sessionID SessionRef, self SDNControllerRef) (_retval string, _err error) {
	log.Println("SDNController.GetUUID not mocked")
	_err = errors.New("SDNController.GetUUID not mocked")
	return
}

func (_class SDNControllerClass) GetUUIDMock(sessionID SessionRef, self SDNControllerRef) (_retval string, _err error) {
	return SDNControllerClass_GetUUIDMockedCallback(sessionID, self)
}
// Get the uuid field of the given SDN_controller.
func (_class SDNControllerClass) GetUUID(sessionID SessionRef, self SDNControllerRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetUUIDMock(sessionID, self)
	}	
	_method := "SDN_controller.get_uuid"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertSDNControllerRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var SDNControllerClass_GetByUUIDMockedCallback = func (sessionID SessionRef, uuid string) (_retval SDNControllerRef, _err error) {
	log.Println("SDNController.GetByUUID not mocked")
	_err = errors.New("SDNController.GetByUUID not mocked")
	return
}

func (_class SDNControllerClass) GetByUUIDMock(sessionID SessionRef, uuid string) (_retval SDNControllerRef, _err error) {
	return SDNControllerClass_GetByUUIDMockedCallback(sessionID, uuid)
}
// Get a reference to the SDN_controller instance with the specified UUID.
func (_class SDNControllerClass) GetByUUID(sessionID SessionRef, uuid string) (_retval SDNControllerRef, _err error) {
	if (IsMock) {
		return _class.GetByUUIDMock(sessionID, uuid)
	}	
	_method := "SDN_controller.get_by_uuid"
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
	_retval, _err = convertSDNControllerRefToGo(_method + " -> ", _result.Value)
	return
}


var SDNControllerClass_GetRecordMockedCallback = func (sessionID SessionRef, self SDNControllerRef) (_retval SDNControllerRecord, _err error) {
	log.Println("SDNController.GetRecord not mocked")
	_err = errors.New("SDNController.GetRecord not mocked")
	return
}

func (_class SDNControllerClass) GetRecordMock(sessionID SessionRef, self SDNControllerRef) (_retval SDNControllerRecord, _err error) {
	return SDNControllerClass_GetRecordMockedCallback(sessionID, self)
}
// Get a record containing the current state of the given SDN_controller.
func (_class SDNControllerClass) GetRecord(sessionID SessionRef, self SDNControllerRef) (_retval SDNControllerRecord, _err error) {
	if (IsMock) {
		return _class.GetRecordMock(sessionID, self)
	}	
	_method := "SDN_controller.get_record"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertSDNControllerRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertSDNControllerRecordToGo(_method + " -> ", _result.Value)
	return
}
