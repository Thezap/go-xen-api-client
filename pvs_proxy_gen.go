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

type PvsProxyStatus string

const (
  // The proxy is not currently running
	PvsProxyStatusStopped PvsProxyStatus = "stopped"
  // The proxy is setup but has not yet cached anything
	PvsProxyStatusInitialised PvsProxyStatus = "initialised"
  // The proxy is currently caching data
	PvsProxyStatusCaching PvsProxyStatus = "caching"
  // The PVS device is configured to use an incompatible write-cache mode
	PvsProxyStatusIncompatibleWriteCacheMode PvsProxyStatus = "incompatible_write_cache_mode"
  // The PVS protocol in use is not compatible with the PVS proxy
	PvsProxyStatusIncompatibleProtocolVersion PvsProxyStatus = "incompatible_protocol_version"
)

type PVSProxyRecord struct {
  // Unique identifier/object reference
	UUID string
  // PVS site this proxy is part of
	Site PVSSiteRef
  // VIF of the VM using the proxy
	VIF VIFRef
  // true = VM is currently proxied
	CurrentlyAttached bool
  // The run-time status of the proxy
	Status PvsProxyStatus
}

type PVSProxyRef string

// a proxy connects a VM/VIF with a PVS site
type PVSProxyClass struct {
	client *Client
}


func PVSProxyClassGetAllRecordsMockDefault(sessionID SessionRef) (_retval map[PVSProxyRef]PVSProxyRecord, _err error) {
	log.Println("PVSProxy.GetAllRecords not mocked")
	_err = errors.New("PVSProxy.GetAllRecords not mocked")
	return
}

var PVSProxyClassGetAllRecordsMockedCallback = PVSProxyClassGetAllRecordsMockDefault

func (_class PVSProxyClass) GetAllRecordsMock(sessionID SessionRef) (_retval map[PVSProxyRef]PVSProxyRecord, _err error) {
	return PVSProxyClassGetAllRecordsMockedCallback(sessionID)
}
// Return a map of PVS_proxy references to PVS_proxy records for all PVS_proxys known to the system.
func (_class PVSProxyClass) GetAllRecords(sessionID SessionRef) (_retval map[PVSProxyRef]PVSProxyRecord, _err error) {
	if IsMock {
		return _class.GetAllRecordsMock(sessionID)
	}	
	_method := "PVS_proxy.get_all_records"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertPVSProxyRefToPVSProxyRecordMapToGo(_method + " -> ", _result.Value)
	return
}


func PVSProxyClassGetAllMockDefault(sessionID SessionRef) (_retval []PVSProxyRef, _err error) {
	log.Println("PVSProxy.GetAll not mocked")
	_err = errors.New("PVSProxy.GetAll not mocked")
	return
}

var PVSProxyClassGetAllMockedCallback = PVSProxyClassGetAllMockDefault

func (_class PVSProxyClass) GetAllMock(sessionID SessionRef) (_retval []PVSProxyRef, _err error) {
	return PVSProxyClassGetAllMockedCallback(sessionID)
}
// Return a list of all the PVS_proxys known to the system.
func (_class PVSProxyClass) GetAll(sessionID SessionRef) (_retval []PVSProxyRef, _err error) {
	if IsMock {
		return _class.GetAllMock(sessionID)
	}	
	_method := "PVS_proxy.get_all"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertPVSProxyRefSetToGo(_method + " -> ", _result.Value)
	return
}


func PVSProxyClassDestroyMockDefault(sessionID SessionRef, self PVSProxyRef) (_err error) {
	log.Println("PVSProxy.Destroy not mocked")
	_err = errors.New("PVSProxy.Destroy not mocked")
	return
}

var PVSProxyClassDestroyMockedCallback = PVSProxyClassDestroyMockDefault

func (_class PVSProxyClass) DestroyMock(sessionID SessionRef, self PVSProxyRef) (_err error) {
	return PVSProxyClassDestroyMockedCallback(sessionID, self)
}
// remove (or switch off) a PVS proxy for this VM
func (_class PVSProxyClass) Destroy(sessionID SessionRef, self PVSProxyRef) (_err error) {
	if IsMock {
		return _class.DestroyMock(sessionID, self)
	}	
	_method := "PVS_proxy.destroy"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPVSProxyRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg)
	return
}


func PVSProxyClassCreateMockDefault(sessionID SessionRef, site PVSSiteRef, vif VIFRef) (_retval PVSProxyRef, _err error) {
	log.Println("PVSProxy.Create not mocked")
	_err = errors.New("PVSProxy.Create not mocked")
	return
}

var PVSProxyClassCreateMockedCallback = PVSProxyClassCreateMockDefault

func (_class PVSProxyClass) CreateMock(sessionID SessionRef, site PVSSiteRef, vif VIFRef) (_retval PVSProxyRef, _err error) {
	return PVSProxyClassCreateMockedCallback(sessionID, site, vif)
}
// Configure a VM/VIF to use a PVS proxy
func (_class PVSProxyClass) Create(sessionID SessionRef, site PVSSiteRef, vif VIFRef) (_retval PVSProxyRef, _err error) {
	if IsMock {
		return _class.CreateMock(sessionID, site, vif)
	}	
	_method := "PVS_proxy.create"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_siteArg, _err := convertPVSSiteRefToXen(fmt.Sprintf("%s(%s)", _method, "site"), site)
	if _err != nil {
		return
	}
	_vifArg, _err := convertVIFRefToXen(fmt.Sprintf("%s(%s)", _method, "VIF"), vif)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _siteArg, _vifArg)
	if _err != nil {
		return
	}
	_retval, _err = convertPVSProxyRefToGo(_method + " -> ", _result.Value)
	return
}


func PVSProxyClassGetStatusMockDefault(sessionID SessionRef, self PVSProxyRef) (_retval PvsProxyStatus, _err error) {
	log.Println("PVSProxy.GetStatus not mocked")
	_err = errors.New("PVSProxy.GetStatus not mocked")
	return
}

var PVSProxyClassGetStatusMockedCallback = PVSProxyClassGetStatusMockDefault

func (_class PVSProxyClass) GetStatusMock(sessionID SessionRef, self PVSProxyRef) (_retval PvsProxyStatus, _err error) {
	return PVSProxyClassGetStatusMockedCallback(sessionID, self)
}
// Get the status field of the given PVS_proxy.
func (_class PVSProxyClass) GetStatus(sessionID SessionRef, self PVSProxyRef) (_retval PvsProxyStatus, _err error) {
	if IsMock {
		return _class.GetStatusMock(sessionID, self)
	}	
	_method := "PVS_proxy.get_status"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPVSProxyRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertEnumPvsProxyStatusToGo(_method + " -> ", _result.Value)
	return
}


func PVSProxyClassGetCurrentlyAttachedMockDefault(sessionID SessionRef, self PVSProxyRef) (_retval bool, _err error) {
	log.Println("PVSProxy.GetCurrentlyAttached not mocked")
	_err = errors.New("PVSProxy.GetCurrentlyAttached not mocked")
	return
}

var PVSProxyClassGetCurrentlyAttachedMockedCallback = PVSProxyClassGetCurrentlyAttachedMockDefault

func (_class PVSProxyClass) GetCurrentlyAttachedMock(sessionID SessionRef, self PVSProxyRef) (_retval bool, _err error) {
	return PVSProxyClassGetCurrentlyAttachedMockedCallback(sessionID, self)
}
// Get the currently_attached field of the given PVS_proxy.
func (_class PVSProxyClass) GetCurrentlyAttached(sessionID SessionRef, self PVSProxyRef) (_retval bool, _err error) {
	if IsMock {
		return _class.GetCurrentlyAttachedMock(sessionID, self)
	}	
	_method := "PVS_proxy.get_currently_attached"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPVSProxyRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func PVSProxyClassGetVIFMockDefault(sessionID SessionRef, self PVSProxyRef) (_retval VIFRef, _err error) {
	log.Println("PVSProxy.GetVIF not mocked")
	_err = errors.New("PVSProxy.GetVIF not mocked")
	return
}

var PVSProxyClassGetVIFMockedCallback = PVSProxyClassGetVIFMockDefault

func (_class PVSProxyClass) GetVIFMock(sessionID SessionRef, self PVSProxyRef) (_retval VIFRef, _err error) {
	return PVSProxyClassGetVIFMockedCallback(sessionID, self)
}
// Get the VIF field of the given PVS_proxy.
func (_class PVSProxyClass) GetVIF(sessionID SessionRef, self PVSProxyRef) (_retval VIFRef, _err error) {
	if IsMock {
		return _class.GetVIFMock(sessionID, self)
	}	
	_method := "PVS_proxy.get_VIF"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPVSProxyRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVIFRefToGo(_method + " -> ", _result.Value)
	return
}


func PVSProxyClassGetSiteMockDefault(sessionID SessionRef, self PVSProxyRef) (_retval PVSSiteRef, _err error) {
	log.Println("PVSProxy.GetSite not mocked")
	_err = errors.New("PVSProxy.GetSite not mocked")
	return
}

var PVSProxyClassGetSiteMockedCallback = PVSProxyClassGetSiteMockDefault

func (_class PVSProxyClass) GetSiteMock(sessionID SessionRef, self PVSProxyRef) (_retval PVSSiteRef, _err error) {
	return PVSProxyClassGetSiteMockedCallback(sessionID, self)
}
// Get the site field of the given PVS_proxy.
func (_class PVSProxyClass) GetSite(sessionID SessionRef, self PVSProxyRef) (_retval PVSSiteRef, _err error) {
	if IsMock {
		return _class.GetSiteMock(sessionID, self)
	}	
	_method := "PVS_proxy.get_site"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPVSProxyRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func PVSProxyClassGetUUIDMockDefault(sessionID SessionRef, self PVSProxyRef) (_retval string, _err error) {
	log.Println("PVSProxy.GetUUID not mocked")
	_err = errors.New("PVSProxy.GetUUID not mocked")
	return
}

var PVSProxyClassGetUUIDMockedCallback = PVSProxyClassGetUUIDMockDefault

func (_class PVSProxyClass) GetUUIDMock(sessionID SessionRef, self PVSProxyRef) (_retval string, _err error) {
	return PVSProxyClassGetUUIDMockedCallback(sessionID, self)
}
// Get the uuid field of the given PVS_proxy.
func (_class PVSProxyClass) GetUUID(sessionID SessionRef, self PVSProxyRef) (_retval string, _err error) {
	if IsMock {
		return _class.GetUUIDMock(sessionID, self)
	}	
	_method := "PVS_proxy.get_uuid"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPVSProxyRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func PVSProxyClassGetByUUIDMockDefault(sessionID SessionRef, uuid string) (_retval PVSProxyRef, _err error) {
	log.Println("PVSProxy.GetByUUID not mocked")
	_err = errors.New("PVSProxy.GetByUUID not mocked")
	return
}

var PVSProxyClassGetByUUIDMockedCallback = PVSProxyClassGetByUUIDMockDefault

func (_class PVSProxyClass) GetByUUIDMock(sessionID SessionRef, uuid string) (_retval PVSProxyRef, _err error) {
	return PVSProxyClassGetByUUIDMockedCallback(sessionID, uuid)
}
// Get a reference to the PVS_proxy instance with the specified UUID.
func (_class PVSProxyClass) GetByUUID(sessionID SessionRef, uuid string) (_retval PVSProxyRef, _err error) {
	if IsMock {
		return _class.GetByUUIDMock(sessionID, uuid)
	}	
	_method := "PVS_proxy.get_by_uuid"
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
	_retval, _err = convertPVSProxyRefToGo(_method + " -> ", _result.Value)
	return
}


func PVSProxyClassGetRecordMockDefault(sessionID SessionRef, self PVSProxyRef) (_retval PVSProxyRecord, _err error) {
	log.Println("PVSProxy.GetRecord not mocked")
	_err = errors.New("PVSProxy.GetRecord not mocked")
	return
}

var PVSProxyClassGetRecordMockedCallback = PVSProxyClassGetRecordMockDefault

func (_class PVSProxyClass) GetRecordMock(sessionID SessionRef, self PVSProxyRef) (_retval PVSProxyRecord, _err error) {
	return PVSProxyClassGetRecordMockedCallback(sessionID, self)
}
// Get a record containing the current state of the given PVS_proxy.
func (_class PVSProxyClass) GetRecord(sessionID SessionRef, self PVSProxyRef) (_retval PVSProxyRecord, _err error) {
	if IsMock {
		return _class.GetRecordMock(sessionID, self)
	}	
	_method := "PVS_proxy.get_record"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPVSProxyRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertPVSProxyRecordToGo(_method + " -> ", _result.Value)
	return
}
