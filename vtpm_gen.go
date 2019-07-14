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

type VTPMRecord struct {
  // Unique identifier/object reference
	UUID string
  // the virtual machine
	VM VMRef
  // the domain where the backend is located
	Backend VMRef
}

type VTPMRef string

// A virtual TPM device
type VTPMClass struct {
	client *Client
}


var VTPMClass_GetBackendMockedCallback = func (sessionID SessionRef, self VTPMRef) (_retval VMRef, _err error) {
	log.Println("VTPM.GetBackend not mocked")
	_err = errors.New("VTPM.GetBackend not mocked")
	return
}

func (_class VTPMClass) GetBackendMock(sessionID SessionRef, self VTPMRef) (_retval VMRef, _err error) {
	return VTPMClass_GetBackendMockedCallback(sessionID, self)
}
// Get the backend field of the given VTPM.
func (_class VTPMClass) GetBackend(sessionID SessionRef, self VTPMRef) (_retval VMRef, _err error) {
	if (IsMock) {
		return _class.GetBackendMock(sessionID, self)
	}	
	_method := "VTPM.get_backend"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVTPMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var VTPMClass_GetVMMockedCallback = func (sessionID SessionRef, self VTPMRef) (_retval VMRef, _err error) {
	log.Println("VTPM.GetVM not mocked")
	_err = errors.New("VTPM.GetVM not mocked")
	return
}

func (_class VTPMClass) GetVMMock(sessionID SessionRef, self VTPMRef) (_retval VMRef, _err error) {
	return VTPMClass_GetVMMockedCallback(sessionID, self)
}
// Get the VM field of the given VTPM.
func (_class VTPMClass) GetVM(sessionID SessionRef, self VTPMRef) (_retval VMRef, _err error) {
	if (IsMock) {
		return _class.GetVMMock(sessionID, self)
	}	
	_method := "VTPM.get_VM"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVTPMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var VTPMClass_GetUUIDMockedCallback = func (sessionID SessionRef, self VTPMRef) (_retval string, _err error) {
	log.Println("VTPM.GetUUID not mocked")
	_err = errors.New("VTPM.GetUUID not mocked")
	return
}

func (_class VTPMClass) GetUUIDMock(sessionID SessionRef, self VTPMRef) (_retval string, _err error) {
	return VTPMClass_GetUUIDMockedCallback(sessionID, self)
}
// Get the uuid field of the given VTPM.
func (_class VTPMClass) GetUUID(sessionID SessionRef, self VTPMRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetUUIDMock(sessionID, self)
	}	
	_method := "VTPM.get_uuid"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVTPMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var VTPMClass_DestroyMockedCallback = func (sessionID SessionRef, self VTPMRef) (_err error) {
	log.Println("VTPM.Destroy not mocked")
	_err = errors.New("VTPM.Destroy not mocked")
	return
}

func (_class VTPMClass) DestroyMock(sessionID SessionRef, self VTPMRef) (_err error) {
	return VTPMClass_DestroyMockedCallback(sessionID, self)
}
// Destroy the specified VTPM instance.
func (_class VTPMClass) Destroy(sessionID SessionRef, self VTPMRef) (_err error) {
	if (IsMock) {
		return _class.DestroyMock(sessionID, self)
	}	
	_method := "VTPM.destroy"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVTPMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg)
	return
}


var VTPMClass_CreateMockedCallback = func (sessionID SessionRef, args VTPMRecord) (_retval VTPMRef, _err error) {
	log.Println("VTPM.Create not mocked")
	_err = errors.New("VTPM.Create not mocked")
	return
}

func (_class VTPMClass) CreateMock(sessionID SessionRef, args VTPMRecord) (_retval VTPMRef, _err error) {
	return VTPMClass_CreateMockedCallback(sessionID, args)
}
// Create a new VTPM instance, and return its handle.
// The constructor args are: VM*, backend* (* = non-optional).
func (_class VTPMClass) Create(sessionID SessionRef, args VTPMRecord) (_retval VTPMRef, _err error) {
	if (IsMock) {
		return _class.CreateMock(sessionID, args)
	}	
	_method := "VTPM.create"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_argsArg, _err := convertVTPMRecordToXen(fmt.Sprintf("%s(%s)", _method, "args"), args)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _argsArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVTPMRefToGo(_method + " -> ", _result.Value)
	return
}


var VTPMClass_GetByUUIDMockedCallback = func (sessionID SessionRef, uuid string) (_retval VTPMRef, _err error) {
	log.Println("VTPM.GetByUUID not mocked")
	_err = errors.New("VTPM.GetByUUID not mocked")
	return
}

func (_class VTPMClass) GetByUUIDMock(sessionID SessionRef, uuid string) (_retval VTPMRef, _err error) {
	return VTPMClass_GetByUUIDMockedCallback(sessionID, uuid)
}
// Get a reference to the VTPM instance with the specified UUID.
func (_class VTPMClass) GetByUUID(sessionID SessionRef, uuid string) (_retval VTPMRef, _err error) {
	if (IsMock) {
		return _class.GetByUUIDMock(sessionID, uuid)
	}	
	_method := "VTPM.get_by_uuid"
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
	_retval, _err = convertVTPMRefToGo(_method + " -> ", _result.Value)
	return
}


var VTPMClass_GetRecordMockedCallback = func (sessionID SessionRef, self VTPMRef) (_retval VTPMRecord, _err error) {
	log.Println("VTPM.GetRecord not mocked")
	_err = errors.New("VTPM.GetRecord not mocked")
	return
}

func (_class VTPMClass) GetRecordMock(sessionID SessionRef, self VTPMRef) (_retval VTPMRecord, _err error) {
	return VTPMClass_GetRecordMockedCallback(sessionID, self)
}
// Get a record containing the current state of the given VTPM.
func (_class VTPMClass) GetRecord(sessionID SessionRef, self VTPMRef) (_retval VTPMRecord, _err error) {
	if (IsMock) {
		return _class.GetRecordMock(sessionID, self)
	}	
	_method := "VTPM.get_record"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVTPMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVTPMRecordToGo(_method + " -> ", _result.Value)
	return
}
