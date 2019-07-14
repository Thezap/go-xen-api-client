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

type PVSServerRecord struct {
  // Unique identifier/object reference
	UUID string
  // IPv4 addresses of this server
	Addresses []string
  // First UDP port accepted by this server
	FirstPort int
  // Last UDP port accepted by this server
	LastPort int
  // PVS site this server is part of
	Site PVSSiteRef
}

type PVSServerRef string

// individual machine serving provisioning (block) data
type PVSServerClass struct {
	client *Client
}


var PVSServerClass_GetAllRecordsMockedCallback = func (sessionID SessionRef) (_retval map[PVSServerRef]PVSServerRecord, _err error) {
	log.Println("PVSServer.GetAllRecords not mocked")
	_err = errors.New("PVSServer.GetAllRecords not mocked")
	return
}

func (_class PVSServerClass) GetAllRecordsMock(sessionID SessionRef) (_retval map[PVSServerRef]PVSServerRecord, _err error) {
	return PVSServerClass_GetAllRecordsMockedCallback(sessionID)
}
// Return a map of PVS_server references to PVS_server records for all PVS_servers known to the system.
func (_class PVSServerClass) GetAllRecords(sessionID SessionRef) (_retval map[PVSServerRef]PVSServerRecord, _err error) {
	if (IsMock) {
		return _class.GetAllRecordsMock(sessionID)
	}	
	_method := "PVS_server.get_all_records"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertPVSServerRefToPVSServerRecordMapToGo(_method + " -> ", _result.Value)
	return
}


var PVSServerClass_GetAllMockedCallback = func (sessionID SessionRef) (_retval []PVSServerRef, _err error) {
	log.Println("PVSServer.GetAll not mocked")
	_err = errors.New("PVSServer.GetAll not mocked")
	return
}

func (_class PVSServerClass) GetAllMock(sessionID SessionRef) (_retval []PVSServerRef, _err error) {
	return PVSServerClass_GetAllMockedCallback(sessionID)
}
// Return a list of all the PVS_servers known to the system.
func (_class PVSServerClass) GetAll(sessionID SessionRef) (_retval []PVSServerRef, _err error) {
	if (IsMock) {
		return _class.GetAllMock(sessionID)
	}	
	_method := "PVS_server.get_all"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertPVSServerRefSetToGo(_method + " -> ", _result.Value)
	return
}


var PVSServerClass_ForgetMockedCallback = func (sessionID SessionRef, self PVSServerRef) (_err error) {
	log.Println("PVSServer.Forget not mocked")
	_err = errors.New("PVSServer.Forget not mocked")
	return
}

func (_class PVSServerClass) ForgetMock(sessionID SessionRef, self PVSServerRef) (_err error) {
	return PVSServerClass_ForgetMockedCallback(sessionID, self)
}
// forget a PVS server
func (_class PVSServerClass) Forget(sessionID SessionRef, self PVSServerRef) (_err error) {
	if (IsMock) {
		return _class.ForgetMock(sessionID, self)
	}	
	_method := "PVS_server.forget"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPVSServerRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg)
	return
}


var PVSServerClass_IntroduceMockedCallback = func (sessionID SessionRef, addresses []string, firstPort int, lastPort int, site PVSSiteRef) (_retval PVSServerRef, _err error) {
	log.Println("PVSServer.Introduce not mocked")
	_err = errors.New("PVSServer.Introduce not mocked")
	return
}

func (_class PVSServerClass) IntroduceMock(sessionID SessionRef, addresses []string, firstPort int, lastPort int, site PVSSiteRef) (_retval PVSServerRef, _err error) {
	return PVSServerClass_IntroduceMockedCallback(sessionID, addresses, firstPort, lastPort, site)
}
// introduce new PVS server
func (_class PVSServerClass) Introduce(sessionID SessionRef, addresses []string, firstPort int, lastPort int, site PVSSiteRef) (_retval PVSServerRef, _err error) {
	if (IsMock) {
		return _class.IntroduceMock(sessionID, addresses, firstPort, lastPort, site)
	}	
	_method := "PVS_server.introduce"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_addressesArg, _err := convertStringSetToXen(fmt.Sprintf("%s(%s)", _method, "addresses"), addresses)
	if _err != nil {
		return
	}
	_firstPortArg, _err := convertIntToXen(fmt.Sprintf("%s(%s)", _method, "first_port"), firstPort)
	if _err != nil {
		return
	}
	_lastPortArg, _err := convertIntToXen(fmt.Sprintf("%s(%s)", _method, "last_port"), lastPort)
	if _err != nil {
		return
	}
	_siteArg, _err := convertPVSSiteRefToXen(fmt.Sprintf("%s(%s)", _method, "site"), site)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _addressesArg, _firstPortArg, _lastPortArg, _siteArg)
	if _err != nil {
		return
	}
	_retval, _err = convertPVSServerRefToGo(_method + " -> ", _result.Value)
	return
}


var PVSServerClass_GetSiteMockedCallback = func (sessionID SessionRef, self PVSServerRef) (_retval PVSSiteRef, _err error) {
	log.Println("PVSServer.GetSite not mocked")
	_err = errors.New("PVSServer.GetSite not mocked")
	return
}

func (_class PVSServerClass) GetSiteMock(sessionID SessionRef, self PVSServerRef) (_retval PVSSiteRef, _err error) {
	return PVSServerClass_GetSiteMockedCallback(sessionID, self)
}
// Get the site field of the given PVS_server.
func (_class PVSServerClass) GetSite(sessionID SessionRef, self PVSServerRef) (_retval PVSSiteRef, _err error) {
	if (IsMock) {
		return _class.GetSiteMock(sessionID, self)
	}	
	_method := "PVS_server.get_site"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPVSServerRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertPVSSiteRefToGo(_method + " -> ", _result.Value)
	return
}


var PVSServerClass_GetLastPortMockedCallback = func (sessionID SessionRef, self PVSServerRef) (_retval int, _err error) {
	log.Println("PVSServer.GetLastPort not mocked")
	_err = errors.New("PVSServer.GetLastPort not mocked")
	return
}

func (_class PVSServerClass) GetLastPortMock(sessionID SessionRef, self PVSServerRef) (_retval int, _err error) {
	return PVSServerClass_GetLastPortMockedCallback(sessionID, self)
}
// Get the last_port field of the given PVS_server.
func (_class PVSServerClass) GetLastPort(sessionID SessionRef, self PVSServerRef) (_retval int, _err error) {
	if (IsMock) {
		return _class.GetLastPortMock(sessionID, self)
	}	
	_method := "PVS_server.get_last_port"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPVSServerRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var PVSServerClass_GetFirstPortMockedCallback = func (sessionID SessionRef, self PVSServerRef) (_retval int, _err error) {
	log.Println("PVSServer.GetFirstPort not mocked")
	_err = errors.New("PVSServer.GetFirstPort not mocked")
	return
}

func (_class PVSServerClass) GetFirstPortMock(sessionID SessionRef, self PVSServerRef) (_retval int, _err error) {
	return PVSServerClass_GetFirstPortMockedCallback(sessionID, self)
}
// Get the first_port field of the given PVS_server.
func (_class PVSServerClass) GetFirstPort(sessionID SessionRef, self PVSServerRef) (_retval int, _err error) {
	if (IsMock) {
		return _class.GetFirstPortMock(sessionID, self)
	}	
	_method := "PVS_server.get_first_port"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPVSServerRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var PVSServerClass_GetAddressesMockedCallback = func (sessionID SessionRef, self PVSServerRef) (_retval []string, _err error) {
	log.Println("PVSServer.GetAddresses not mocked")
	_err = errors.New("PVSServer.GetAddresses not mocked")
	return
}

func (_class PVSServerClass) GetAddressesMock(sessionID SessionRef, self PVSServerRef) (_retval []string, _err error) {
	return PVSServerClass_GetAddressesMockedCallback(sessionID, self)
}
// Get the addresses field of the given PVS_server.
func (_class PVSServerClass) GetAddresses(sessionID SessionRef, self PVSServerRef) (_retval []string, _err error) {
	if (IsMock) {
		return _class.GetAddressesMock(sessionID, self)
	}	
	_method := "PVS_server.get_addresses"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPVSServerRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringSetToGo(_method + " -> ", _result.Value)
	return
}


var PVSServerClass_GetUUIDMockedCallback = func (sessionID SessionRef, self PVSServerRef) (_retval string, _err error) {
	log.Println("PVSServer.GetUUID not mocked")
	_err = errors.New("PVSServer.GetUUID not mocked")
	return
}

func (_class PVSServerClass) GetUUIDMock(sessionID SessionRef, self PVSServerRef) (_retval string, _err error) {
	return PVSServerClass_GetUUIDMockedCallback(sessionID, self)
}
// Get the uuid field of the given PVS_server.
func (_class PVSServerClass) GetUUID(sessionID SessionRef, self PVSServerRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetUUIDMock(sessionID, self)
	}	
	_method := "PVS_server.get_uuid"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPVSServerRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var PVSServerClass_GetByUUIDMockedCallback = func (sessionID SessionRef, uuid string) (_retval PVSServerRef, _err error) {
	log.Println("PVSServer.GetByUUID not mocked")
	_err = errors.New("PVSServer.GetByUUID not mocked")
	return
}

func (_class PVSServerClass) GetByUUIDMock(sessionID SessionRef, uuid string) (_retval PVSServerRef, _err error) {
	return PVSServerClass_GetByUUIDMockedCallback(sessionID, uuid)
}
// Get a reference to the PVS_server instance with the specified UUID.
func (_class PVSServerClass) GetByUUID(sessionID SessionRef, uuid string) (_retval PVSServerRef, _err error) {
	if (IsMock) {
		return _class.GetByUUIDMock(sessionID, uuid)
	}	
	_method := "PVS_server.get_by_uuid"
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
	_retval, _err = convertPVSServerRefToGo(_method + " -> ", _result.Value)
	return
}


var PVSServerClass_GetRecordMockedCallback = func (sessionID SessionRef, self PVSServerRef) (_retval PVSServerRecord, _err error) {
	log.Println("PVSServer.GetRecord not mocked")
	_err = errors.New("PVSServer.GetRecord not mocked")
	return
}

func (_class PVSServerClass) GetRecordMock(sessionID SessionRef, self PVSServerRef) (_retval PVSServerRecord, _err error) {
	return PVSServerClass_GetRecordMockedCallback(sessionID, self)
}
// Get a record containing the current state of the given PVS_server.
func (_class PVSServerClass) GetRecord(sessionID SessionRef, self PVSServerRef) (_retval PVSServerRecord, _err error) {
	if (IsMock) {
		return _class.GetRecordMock(sessionID, self)
	}	
	_method := "PVS_server.get_record"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPVSServerRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertPVSServerRecordToGo(_method + " -> ", _result.Value)
	return
}
