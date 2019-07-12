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

type PVSCacheStorageRecord struct {
  // Unique identifier/object reference
	UUID string
  // The host on which this object defines PVS cache storage
	Host HostRef
  // SR providing storage for the PVS cache
	SR SRRef
  // The PVS_site for which this object defines the storage
	Site PVSSiteRef
  // The size of the cache VDI (in bytes)
	Size int
  // The VDI used for caching
	VDI VDIRef
}

type PVSCacheStorageRef string

// Describes the storage that is available to a PVS site for caching purposes
type PVSCacheStorageClass struct {
	client *Client
}

func (_class PVSCacheStorageClass) GetAllRecords__mock(sessionID SessionRef) (_retval map[PVSCacheStorageRef]PVSCacheStorageRecord, _err error) {
	log.Println("PVSCacheStorage.GetAllRecords not mocked")
	_err = errors.New("PVSCacheStorage.GetAllRecords not mocked")
	return
}
// Return a map of PVS_cache_storage references to PVS_cache_storage records for all PVS_cache_storages known to the system.
func (_class PVSCacheStorageClass) GetAllRecords(sessionID SessionRef) (_retval map[PVSCacheStorageRef]PVSCacheStorageRecord, _err error) {
	if (IsMock) {
		return _class.GetAllRecords__mock(sessionID)
	}	
	_method := "PVS_cache_storage.get_all_records"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertPVSCacheStorageRefToPVSCacheStorageRecordMapToGo(_method + " -> ", _result.Value)
	return
}

func (_class PVSCacheStorageClass) GetAll__mock(sessionID SessionRef) (_retval []PVSCacheStorageRef, _err error) {
	log.Println("PVSCacheStorage.GetAll not mocked")
	_err = errors.New("PVSCacheStorage.GetAll not mocked")
	return
}
// Return a list of all the PVS_cache_storages known to the system.
func (_class PVSCacheStorageClass) GetAll(sessionID SessionRef) (_retval []PVSCacheStorageRef, _err error) {
	if (IsMock) {
		return _class.GetAll__mock(sessionID)
	}	
	_method := "PVS_cache_storage.get_all"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertPVSCacheStorageRefSetToGo(_method + " -> ", _result.Value)
	return
}

func (_class PVSCacheStorageClass) GetVDI__mock(sessionID SessionRef, self PVSCacheStorageRef) (_retval VDIRef, _err error) {
	log.Println("PVSCacheStorage.GetVDI not mocked")
	_err = errors.New("PVSCacheStorage.GetVDI not mocked")
	return
}
// Get the VDI field of the given PVS_cache_storage.
func (_class PVSCacheStorageClass) GetVDI(sessionID SessionRef, self PVSCacheStorageRef) (_retval VDIRef, _err error) {
	if (IsMock) {
		return _class.GetVDI__mock(sessionID, self)
	}	
	_method := "PVS_cache_storage.get_VDI"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPVSCacheStorageRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVDIRefToGo(_method + " -> ", _result.Value)
	return
}

func (_class PVSCacheStorageClass) GetSize__mock(sessionID SessionRef, self PVSCacheStorageRef) (_retval int, _err error) {
	log.Println("PVSCacheStorage.GetSize not mocked")
	_err = errors.New("PVSCacheStorage.GetSize not mocked")
	return
}
// Get the size field of the given PVS_cache_storage.
func (_class PVSCacheStorageClass) GetSize(sessionID SessionRef, self PVSCacheStorageRef) (_retval int, _err error) {
	if (IsMock) {
		return _class.GetSize__mock(sessionID, self)
	}	
	_method := "PVS_cache_storage.get_size"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPVSCacheStorageRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PVSCacheStorageClass) GetSite__mock(sessionID SessionRef, self PVSCacheStorageRef) (_retval PVSSiteRef, _err error) {
	log.Println("PVSCacheStorage.GetSite not mocked")
	_err = errors.New("PVSCacheStorage.GetSite not mocked")
	return
}
// Get the site field of the given PVS_cache_storage.
func (_class PVSCacheStorageClass) GetSite(sessionID SessionRef, self PVSCacheStorageRef) (_retval PVSSiteRef, _err error) {
	if (IsMock) {
		return _class.GetSite__mock(sessionID, self)
	}	
	_method := "PVS_cache_storage.get_site"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPVSCacheStorageRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PVSCacheStorageClass) GetSR__mock(sessionID SessionRef, self PVSCacheStorageRef) (_retval SRRef, _err error) {
	log.Println("PVSCacheStorage.GetSR not mocked")
	_err = errors.New("PVSCacheStorage.GetSR not mocked")
	return
}
// Get the SR field of the given PVS_cache_storage.
func (_class PVSCacheStorageClass) GetSR(sessionID SessionRef, self PVSCacheStorageRef) (_retval SRRef, _err error) {
	if (IsMock) {
		return _class.GetSR__mock(sessionID, self)
	}	
	_method := "PVS_cache_storage.get_SR"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPVSCacheStorageRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PVSCacheStorageClass) GetHost__mock(sessionID SessionRef, self PVSCacheStorageRef) (_retval HostRef, _err error) {
	log.Println("PVSCacheStorage.GetHost not mocked")
	_err = errors.New("PVSCacheStorage.GetHost not mocked")
	return
}
// Get the host field of the given PVS_cache_storage.
func (_class PVSCacheStorageClass) GetHost(sessionID SessionRef, self PVSCacheStorageRef) (_retval HostRef, _err error) {
	if (IsMock) {
		return _class.GetHost__mock(sessionID, self)
	}	
	_method := "PVS_cache_storage.get_host"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPVSCacheStorageRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PVSCacheStorageClass) GetUUID__mock(sessionID SessionRef, self PVSCacheStorageRef) (_retval string, _err error) {
	log.Println("PVSCacheStorage.GetUUID not mocked")
	_err = errors.New("PVSCacheStorage.GetUUID not mocked")
	return
}
// Get the uuid field of the given PVS_cache_storage.
func (_class PVSCacheStorageClass) GetUUID(sessionID SessionRef, self PVSCacheStorageRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetUUID__mock(sessionID, self)
	}	
	_method := "PVS_cache_storage.get_uuid"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPVSCacheStorageRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PVSCacheStorageClass) Destroy__mock(sessionID SessionRef, self PVSCacheStorageRef) (_err error) {
	log.Println("PVSCacheStorage.Destroy not mocked")
	_err = errors.New("PVSCacheStorage.Destroy not mocked")
	return
}
// Destroy the specified PVS_cache_storage instance.
func (_class PVSCacheStorageClass) Destroy(sessionID SessionRef, self PVSCacheStorageRef) (_err error) {
	if (IsMock) {
		return _class.Destroy__mock(sessionID, self)
	}	
	_method := "PVS_cache_storage.destroy"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPVSCacheStorageRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg)
	return
}

func (_class PVSCacheStorageClass) Create__mock(sessionID SessionRef, args PVSCacheStorageRecord) (_retval PVSCacheStorageRef, _err error) {
	log.Println("PVSCacheStorage.Create not mocked")
	_err = errors.New("PVSCacheStorage.Create not mocked")
	return
}
// Create a new PVS_cache_storage instance, and return its handle.
// The constructor args are: host, SR, site, size (* = non-optional).
func (_class PVSCacheStorageClass) Create(sessionID SessionRef, args PVSCacheStorageRecord) (_retval PVSCacheStorageRef, _err error) {
	if (IsMock) {
		return _class.Create__mock(sessionID, args)
	}	
	_method := "PVS_cache_storage.create"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_argsArg, _err := convertPVSCacheStorageRecordToXen(fmt.Sprintf("%s(%s)", _method, "args"), args)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _argsArg)
	if _err != nil {
		return
	}
	_retval, _err = convertPVSCacheStorageRefToGo(_method + " -> ", _result.Value)
	return
}

func (_class PVSCacheStorageClass) GetByUUID__mock(sessionID SessionRef, uuid string) (_retval PVSCacheStorageRef, _err error) {
	log.Println("PVSCacheStorage.GetByUUID not mocked")
	_err = errors.New("PVSCacheStorage.GetByUUID not mocked")
	return
}
// Get a reference to the PVS_cache_storage instance with the specified UUID.
func (_class PVSCacheStorageClass) GetByUUID(sessionID SessionRef, uuid string) (_retval PVSCacheStorageRef, _err error) {
	if (IsMock) {
		return _class.GetByUUID__mock(sessionID, uuid)
	}	
	_method := "PVS_cache_storage.get_by_uuid"
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
	_retval, _err = convertPVSCacheStorageRefToGo(_method + " -> ", _result.Value)
	return
}

func (_class PVSCacheStorageClass) GetRecord__mock(sessionID SessionRef, self PVSCacheStorageRef) (_retval PVSCacheStorageRecord, _err error) {
	log.Println("PVSCacheStorage.GetRecord not mocked")
	_err = errors.New("PVSCacheStorage.GetRecord not mocked")
	return
}
// Get a record containing the current state of the given PVS_cache_storage.
func (_class PVSCacheStorageClass) GetRecord(sessionID SessionRef, self PVSCacheStorageRef) (_retval PVSCacheStorageRecord, _err error) {
	if (IsMock) {
		return _class.GetRecord__mock(sessionID, self)
	}	
	_method := "PVS_cache_storage.get_record"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPVSCacheStorageRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertPVSCacheStorageRecordToGo(_method + " -> ", _result.Value)
	return
}
