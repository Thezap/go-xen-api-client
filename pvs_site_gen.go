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

type PVSSiteRecord struct {
  // Unique identifier/object reference
	UUID string
  // a human-readable name
	NameLabel string
  // a notes field containing human-readable description
	NameDescription string
  // Unique identifier of the PVS site, as configured in PVS
	PVSUUID string
  // The SR used by PVS proxy for the cache
	CacheStorage []PVSCacheStorageRef
  // The set of PVS servers in the site
	Servers []PVSServerRef
  // The set of proxies associated with the site
	Proxies []PVSProxyRef
}

type PVSSiteRef string

// machines serving blocks of data for provisioning VMs
type PVSSiteClass struct {
	client *Client
}


func PVSSiteClassGetAllRecordsMockDefault(sessionID SessionRef) (_retval map[PVSSiteRef]PVSSiteRecord, _err error) {
	log.Println("PVSSite.GetAllRecords not mocked")
	_err = errors.New("PVSSite.GetAllRecords not mocked")
	return
}

var PVSSiteClassGetAllRecordsMockedCallback = PVSSiteClassGetAllRecordsMockDefault

func (_class PVSSiteClass) GetAllRecordsMock(sessionID SessionRef) (_retval map[PVSSiteRef]PVSSiteRecord, _err error) {
	return PVSSiteClassGetAllRecordsMockedCallback(sessionID)
}
// Return a map of PVS_site references to PVS_site records for all PVS_sites known to the system.
func (_class PVSSiteClass) GetAllRecords(sessionID SessionRef) (_retval map[PVSSiteRef]PVSSiteRecord, _err error) {
	if IsMock {
		return _class.GetAllRecordsMock(sessionID)
	}	
	_method := "PVS_site.get_all_records"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertPVSSiteRefToPVSSiteRecordMapToGo(_method + " -> ", _result.Value)
	return
}


func PVSSiteClassGetAllMockDefault(sessionID SessionRef) (_retval []PVSSiteRef, _err error) {
	log.Println("PVSSite.GetAll not mocked")
	_err = errors.New("PVSSite.GetAll not mocked")
	return
}

var PVSSiteClassGetAllMockedCallback = PVSSiteClassGetAllMockDefault

func (_class PVSSiteClass) GetAllMock(sessionID SessionRef) (_retval []PVSSiteRef, _err error) {
	return PVSSiteClassGetAllMockedCallback(sessionID)
}
// Return a list of all the PVS_sites known to the system.
func (_class PVSSiteClass) GetAll(sessionID SessionRef) (_retval []PVSSiteRef, _err error) {
	if IsMock {
		return _class.GetAllMock(sessionID)
	}	
	_method := "PVS_site.get_all"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertPVSSiteRefSetToGo(_method + " -> ", _result.Value)
	return
}


func PVSSiteClassSetPVSUUIDMockDefault(sessionID SessionRef, self PVSSiteRef, value string) (_err error) {
	log.Println("PVSSite.SetPVSUUID not mocked")
	_err = errors.New("PVSSite.SetPVSUUID not mocked")
	return
}

var PVSSiteClassSetPVSUUIDMockedCallback = PVSSiteClassSetPVSUUIDMockDefault

func (_class PVSSiteClass) SetPVSUUIDMock(sessionID SessionRef, self PVSSiteRef, value string) (_err error) {
	return PVSSiteClassSetPVSUUIDMockedCallback(sessionID, self, value)
}
// Update the PVS UUID of the PVS site
func (_class PVSSiteClass) SetPVSUUID(sessionID SessionRef, self PVSSiteRef, value string) (_err error) {
	if IsMock {
		return _class.SetPVSUUIDMock(sessionID, self, value)
	}	
	_method := "PVS_site.set_PVS_uuid"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPVSSiteRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func PVSSiteClassForgetMockDefault(sessionID SessionRef, self PVSSiteRef) (_err error) {
	log.Println("PVSSite.Forget not mocked")
	_err = errors.New("PVSSite.Forget not mocked")
	return
}

var PVSSiteClassForgetMockedCallback = PVSSiteClassForgetMockDefault

func (_class PVSSiteClass) ForgetMock(sessionID SessionRef, self PVSSiteRef) (_err error) {
	return PVSSiteClassForgetMockedCallback(sessionID, self)
}
// Remove a site's meta data
//
// Errors:
//  PVS_SITE_CONTAINS_RUNNING_PROXIES - The PVS site contains running proxies.
//  PVS_SITE_CONTAINS_SERVERS - The PVS site contains servers and cannot be forgotten.
func (_class PVSSiteClass) Forget(sessionID SessionRef, self PVSSiteRef) (_err error) {
	if IsMock {
		return _class.ForgetMock(sessionID, self)
	}	
	_method := "PVS_site.forget"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPVSSiteRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg)
	return
}


func PVSSiteClassIntroduceMockDefault(sessionID SessionRef, nameLabel string, nameDescription string, pvsUUID string) (_retval PVSSiteRef, _err error) {
	log.Println("PVSSite.Introduce not mocked")
	_err = errors.New("PVSSite.Introduce not mocked")
	return
}

var PVSSiteClassIntroduceMockedCallback = PVSSiteClassIntroduceMockDefault

func (_class PVSSiteClass) IntroduceMock(sessionID SessionRef, nameLabel string, nameDescription string, pvsUUID string) (_retval PVSSiteRef, _err error) {
	return PVSSiteClassIntroduceMockedCallback(sessionID, nameLabel, nameDescription, pvsUUID)
}
// Introduce new PVS site
func (_class PVSSiteClass) Introduce(sessionID SessionRef, nameLabel string, nameDescription string, pvsUUID string) (_retval PVSSiteRef, _err error) {
	if IsMock {
		return _class.IntroduceMock(sessionID, nameLabel, nameDescription, pvsUUID)
	}	
	_method := "PVS_site.introduce"
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
	_pvsUUIDArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "PVS_uuid"), pvsUUID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _nameLabelArg, _nameDescriptionArg, _pvsUUIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertPVSSiteRefToGo(_method + " -> ", _result.Value)
	return
}


func PVSSiteClassSetNameDescriptionMockDefault(sessionID SessionRef, self PVSSiteRef, value string) (_err error) {
	log.Println("PVSSite.SetNameDescription not mocked")
	_err = errors.New("PVSSite.SetNameDescription not mocked")
	return
}

var PVSSiteClassSetNameDescriptionMockedCallback = PVSSiteClassSetNameDescriptionMockDefault

func (_class PVSSiteClass) SetNameDescriptionMock(sessionID SessionRef, self PVSSiteRef, value string) (_err error) {
	return PVSSiteClassSetNameDescriptionMockedCallback(sessionID, self, value)
}
// Set the name/description field of the given PVS_site.
func (_class PVSSiteClass) SetNameDescription(sessionID SessionRef, self PVSSiteRef, value string) (_err error) {
	if IsMock {
		return _class.SetNameDescriptionMock(sessionID, self, value)
	}	
	_method := "PVS_site.set_name_description"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPVSSiteRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func PVSSiteClassSetNameLabelMockDefault(sessionID SessionRef, self PVSSiteRef, value string) (_err error) {
	log.Println("PVSSite.SetNameLabel not mocked")
	_err = errors.New("PVSSite.SetNameLabel not mocked")
	return
}

var PVSSiteClassSetNameLabelMockedCallback = PVSSiteClassSetNameLabelMockDefault

func (_class PVSSiteClass) SetNameLabelMock(sessionID SessionRef, self PVSSiteRef, value string) (_err error) {
	return PVSSiteClassSetNameLabelMockedCallback(sessionID, self, value)
}
// Set the name/label field of the given PVS_site.
func (_class PVSSiteClass) SetNameLabel(sessionID SessionRef, self PVSSiteRef, value string) (_err error) {
	if IsMock {
		return _class.SetNameLabelMock(sessionID, self, value)
	}	
	_method := "PVS_site.set_name_label"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPVSSiteRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func PVSSiteClassGetProxiesMockDefault(sessionID SessionRef, self PVSSiteRef) (_retval []PVSProxyRef, _err error) {
	log.Println("PVSSite.GetProxies not mocked")
	_err = errors.New("PVSSite.GetProxies not mocked")
	return
}

var PVSSiteClassGetProxiesMockedCallback = PVSSiteClassGetProxiesMockDefault

func (_class PVSSiteClass) GetProxiesMock(sessionID SessionRef, self PVSSiteRef) (_retval []PVSProxyRef, _err error) {
	return PVSSiteClassGetProxiesMockedCallback(sessionID, self)
}
// Get the proxies field of the given PVS_site.
func (_class PVSSiteClass) GetProxies(sessionID SessionRef, self PVSSiteRef) (_retval []PVSProxyRef, _err error) {
	if IsMock {
		return _class.GetProxiesMock(sessionID, self)
	}	
	_method := "PVS_site.get_proxies"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPVSSiteRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertPVSProxyRefSetToGo(_method + " -> ", _result.Value)
	return
}


func PVSSiteClassGetServersMockDefault(sessionID SessionRef, self PVSSiteRef) (_retval []PVSServerRef, _err error) {
	log.Println("PVSSite.GetServers not mocked")
	_err = errors.New("PVSSite.GetServers not mocked")
	return
}

var PVSSiteClassGetServersMockedCallback = PVSSiteClassGetServersMockDefault

func (_class PVSSiteClass) GetServersMock(sessionID SessionRef, self PVSSiteRef) (_retval []PVSServerRef, _err error) {
	return PVSSiteClassGetServersMockedCallback(sessionID, self)
}
// Get the servers field of the given PVS_site.
func (_class PVSSiteClass) GetServers(sessionID SessionRef, self PVSSiteRef) (_retval []PVSServerRef, _err error) {
	if IsMock {
		return _class.GetServersMock(sessionID, self)
	}	
	_method := "PVS_site.get_servers"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPVSSiteRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertPVSServerRefSetToGo(_method + " -> ", _result.Value)
	return
}


func PVSSiteClassGetCacheStorageMockDefault(sessionID SessionRef, self PVSSiteRef) (_retval []PVSCacheStorageRef, _err error) {
	log.Println("PVSSite.GetCacheStorage not mocked")
	_err = errors.New("PVSSite.GetCacheStorage not mocked")
	return
}

var PVSSiteClassGetCacheStorageMockedCallback = PVSSiteClassGetCacheStorageMockDefault

func (_class PVSSiteClass) GetCacheStorageMock(sessionID SessionRef, self PVSSiteRef) (_retval []PVSCacheStorageRef, _err error) {
	return PVSSiteClassGetCacheStorageMockedCallback(sessionID, self)
}
// Get the cache_storage field of the given PVS_site.
func (_class PVSSiteClass) GetCacheStorage(sessionID SessionRef, self PVSSiteRef) (_retval []PVSCacheStorageRef, _err error) {
	if IsMock {
		return _class.GetCacheStorageMock(sessionID, self)
	}	
	_method := "PVS_site.get_cache_storage"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPVSSiteRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertPVSCacheStorageRefSetToGo(_method + " -> ", _result.Value)
	return
}


func PVSSiteClassGetPVSUUIDMockDefault(sessionID SessionRef, self PVSSiteRef) (_retval string, _err error) {
	log.Println("PVSSite.GetPVSUUID not mocked")
	_err = errors.New("PVSSite.GetPVSUUID not mocked")
	return
}

var PVSSiteClassGetPVSUUIDMockedCallback = PVSSiteClassGetPVSUUIDMockDefault

func (_class PVSSiteClass) GetPVSUUIDMock(sessionID SessionRef, self PVSSiteRef) (_retval string, _err error) {
	return PVSSiteClassGetPVSUUIDMockedCallback(sessionID, self)
}
// Get the PVS_uuid field of the given PVS_site.
func (_class PVSSiteClass) GetPVSUUID(sessionID SessionRef, self PVSSiteRef) (_retval string, _err error) {
	if IsMock {
		return _class.GetPVSUUIDMock(sessionID, self)
	}	
	_method := "PVS_site.get_PVS_uuid"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPVSSiteRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func PVSSiteClassGetNameDescriptionMockDefault(sessionID SessionRef, self PVSSiteRef) (_retval string, _err error) {
	log.Println("PVSSite.GetNameDescription not mocked")
	_err = errors.New("PVSSite.GetNameDescription not mocked")
	return
}

var PVSSiteClassGetNameDescriptionMockedCallback = PVSSiteClassGetNameDescriptionMockDefault

func (_class PVSSiteClass) GetNameDescriptionMock(sessionID SessionRef, self PVSSiteRef) (_retval string, _err error) {
	return PVSSiteClassGetNameDescriptionMockedCallback(sessionID, self)
}
// Get the name/description field of the given PVS_site.
func (_class PVSSiteClass) GetNameDescription(sessionID SessionRef, self PVSSiteRef) (_retval string, _err error) {
	if IsMock {
		return _class.GetNameDescriptionMock(sessionID, self)
	}	
	_method := "PVS_site.get_name_description"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPVSSiteRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func PVSSiteClassGetNameLabelMockDefault(sessionID SessionRef, self PVSSiteRef) (_retval string, _err error) {
	log.Println("PVSSite.GetNameLabel not mocked")
	_err = errors.New("PVSSite.GetNameLabel not mocked")
	return
}

var PVSSiteClassGetNameLabelMockedCallback = PVSSiteClassGetNameLabelMockDefault

func (_class PVSSiteClass) GetNameLabelMock(sessionID SessionRef, self PVSSiteRef) (_retval string, _err error) {
	return PVSSiteClassGetNameLabelMockedCallback(sessionID, self)
}
// Get the name/label field of the given PVS_site.
func (_class PVSSiteClass) GetNameLabel(sessionID SessionRef, self PVSSiteRef) (_retval string, _err error) {
	if IsMock {
		return _class.GetNameLabelMock(sessionID, self)
	}	
	_method := "PVS_site.get_name_label"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPVSSiteRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func PVSSiteClassGetUUIDMockDefault(sessionID SessionRef, self PVSSiteRef) (_retval string, _err error) {
	log.Println("PVSSite.GetUUID not mocked")
	_err = errors.New("PVSSite.GetUUID not mocked")
	return
}

var PVSSiteClassGetUUIDMockedCallback = PVSSiteClassGetUUIDMockDefault

func (_class PVSSiteClass) GetUUIDMock(sessionID SessionRef, self PVSSiteRef) (_retval string, _err error) {
	return PVSSiteClassGetUUIDMockedCallback(sessionID, self)
}
// Get the uuid field of the given PVS_site.
func (_class PVSSiteClass) GetUUID(sessionID SessionRef, self PVSSiteRef) (_retval string, _err error) {
	if IsMock {
		return _class.GetUUIDMock(sessionID, self)
	}	
	_method := "PVS_site.get_uuid"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPVSSiteRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func PVSSiteClassGetByNameLabelMockDefault(sessionID SessionRef, label string) (_retval []PVSSiteRef, _err error) {
	log.Println("PVSSite.GetByNameLabel not mocked")
	_err = errors.New("PVSSite.GetByNameLabel not mocked")
	return
}

var PVSSiteClassGetByNameLabelMockedCallback = PVSSiteClassGetByNameLabelMockDefault

func (_class PVSSiteClass) GetByNameLabelMock(sessionID SessionRef, label string) (_retval []PVSSiteRef, _err error) {
	return PVSSiteClassGetByNameLabelMockedCallback(sessionID, label)
}
// Get all the PVS_site instances with the given label.
func (_class PVSSiteClass) GetByNameLabel(sessionID SessionRef, label string) (_retval []PVSSiteRef, _err error) {
	if IsMock {
		return _class.GetByNameLabelMock(sessionID, label)
	}	
	_method := "PVS_site.get_by_name_label"
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
	_retval, _err = convertPVSSiteRefSetToGo(_method + " -> ", _result.Value)
	return
}


func PVSSiteClassGetByUUIDMockDefault(sessionID SessionRef, uuid string) (_retval PVSSiteRef, _err error) {
	log.Println("PVSSite.GetByUUID not mocked")
	_err = errors.New("PVSSite.GetByUUID not mocked")
	return
}

var PVSSiteClassGetByUUIDMockedCallback = PVSSiteClassGetByUUIDMockDefault

func (_class PVSSiteClass) GetByUUIDMock(sessionID SessionRef, uuid string) (_retval PVSSiteRef, _err error) {
	return PVSSiteClassGetByUUIDMockedCallback(sessionID, uuid)
}
// Get a reference to the PVS_site instance with the specified UUID.
func (_class PVSSiteClass) GetByUUID(sessionID SessionRef, uuid string) (_retval PVSSiteRef, _err error) {
	if IsMock {
		return _class.GetByUUIDMock(sessionID, uuid)
	}	
	_method := "PVS_site.get_by_uuid"
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
	_retval, _err = convertPVSSiteRefToGo(_method + " -> ", _result.Value)
	return
}


func PVSSiteClassGetRecordMockDefault(sessionID SessionRef, self PVSSiteRef) (_retval PVSSiteRecord, _err error) {
	log.Println("PVSSite.GetRecord not mocked")
	_err = errors.New("PVSSite.GetRecord not mocked")
	return
}

var PVSSiteClassGetRecordMockedCallback = PVSSiteClassGetRecordMockDefault

func (_class PVSSiteClass) GetRecordMock(sessionID SessionRef, self PVSSiteRef) (_retval PVSSiteRecord, _err error) {
	return PVSSiteClassGetRecordMockedCallback(sessionID, self)
}
// Get a record containing the current state of the given PVS_site.
func (_class PVSSiteClass) GetRecord(sessionID SessionRef, self PVSSiteRef) (_retval PVSSiteRecord, _err error) {
	if IsMock {
		return _class.GetRecordMock(sessionID, self)
	}	
	_method := "PVS_site.get_record"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPVSSiteRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertPVSSiteRecordToGo(_method + " -> ", _result.Value)
	return
}
