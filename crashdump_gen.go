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

type CrashdumpRecord struct {
  // Unique identifier/object reference
	UUID string
  // the virtual machine
	VM VMRef
  // the virtual disk
	VDI VDIRef
  // additional configuration
	OtherConfig map[string]string
}

type CrashdumpRef string

// A VM crashdump
type CrashdumpClass struct {
	client *Client
}

func (_class CrashdumpClass) GetAllRecords__mock(sessionID SessionRef) (_retval map[CrashdumpRef]CrashdumpRecord, _err error) {
	log.Println("Crashdump.GetAllRecords not mocked")
	_err = errors.New("Crashdump.GetAllRecords not mocked")
	return
}
// Return a map of crashdump references to crashdump records for all crashdumps known to the system.
func (_class CrashdumpClass) GetAllRecords(sessionID SessionRef) (_retval map[CrashdumpRef]CrashdumpRecord, _err error) {
	if (IsMock) {
		return _class.GetAllRecords__mock(sessionID)
	}	
	_method := "crashdump.get_all_records"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertCrashdumpRefToCrashdumpRecordMapToGo(_method + " -> ", _result.Value)
	return
}

func (_class CrashdumpClass) GetAll__mock(sessionID SessionRef) (_retval []CrashdumpRef, _err error) {
	log.Println("Crashdump.GetAll not mocked")
	_err = errors.New("Crashdump.GetAll not mocked")
	return
}
// Return a list of all the crashdumps known to the system.
func (_class CrashdumpClass) GetAll(sessionID SessionRef) (_retval []CrashdumpRef, _err error) {
	if (IsMock) {
		return _class.GetAll__mock(sessionID)
	}	
	_method := "crashdump.get_all"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertCrashdumpRefSetToGo(_method + " -> ", _result.Value)
	return
}

func (_class CrashdumpClass) Destroy__mock(sessionID SessionRef, self CrashdumpRef) (_err error) {
	log.Println("Crashdump.Destroy not mocked")
	_err = errors.New("Crashdump.Destroy not mocked")
	return
}
// Destroy the specified crashdump
func (_class CrashdumpClass) Destroy(sessionID SessionRef, self CrashdumpRef) (_err error) {
	if (IsMock) {
		return _class.Destroy__mock(sessionID, self)
	}	
	_method := "crashdump.destroy"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertCrashdumpRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg)
	return
}

func (_class CrashdumpClass) RemoveFromOtherConfig__mock(sessionID SessionRef, self CrashdumpRef, key string) (_err error) {
	log.Println("Crashdump.RemoveFromOtherConfig not mocked")
	_err = errors.New("Crashdump.RemoveFromOtherConfig not mocked")
	return
}
// Remove the given key and its corresponding value from the other_config field of the given crashdump.  If the key is not in that Map, then do nothing.
func (_class CrashdumpClass) RemoveFromOtherConfig(sessionID SessionRef, self CrashdumpRef, key string) (_err error) {
	if (IsMock) {
		return _class.RemoveFromOtherConfig__mock(sessionID, self, key)
	}	
	_method := "crashdump.remove_from_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertCrashdumpRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class CrashdumpClass) AddToOtherConfig__mock(sessionID SessionRef, self CrashdumpRef, key string, value string) (_err error) {
	log.Println("Crashdump.AddToOtherConfig not mocked")
	_err = errors.New("Crashdump.AddToOtherConfig not mocked")
	return
}
// Add the given key-value pair to the other_config field of the given crashdump.
func (_class CrashdumpClass) AddToOtherConfig(sessionID SessionRef, self CrashdumpRef, key string, value string) (_err error) {
	if (IsMock) {
		return _class.AddToOtherConfig__mock(sessionID, self, key, value)
	}	
	_method := "crashdump.add_to_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertCrashdumpRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class CrashdumpClass) SetOtherConfig__mock(sessionID SessionRef, self CrashdumpRef, value map[string]string) (_err error) {
	log.Println("Crashdump.SetOtherConfig not mocked")
	_err = errors.New("Crashdump.SetOtherConfig not mocked")
	return
}
// Set the other_config field of the given crashdump.
func (_class CrashdumpClass) SetOtherConfig(sessionID SessionRef, self CrashdumpRef, value map[string]string) (_err error) {
	if (IsMock) {
		return _class.SetOtherConfig__mock(sessionID, self, value)
	}	
	_method := "crashdump.set_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertCrashdumpRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class CrashdumpClass) GetOtherConfig__mock(sessionID SessionRef, self CrashdumpRef) (_retval map[string]string, _err error) {
	log.Println("Crashdump.GetOtherConfig not mocked")
	_err = errors.New("Crashdump.GetOtherConfig not mocked")
	return
}
// Get the other_config field of the given crashdump.
func (_class CrashdumpClass) GetOtherConfig(sessionID SessionRef, self CrashdumpRef) (_retval map[string]string, _err error) {
	if (IsMock) {
		return _class.GetOtherConfig__mock(sessionID, self)
	}	
	_method := "crashdump.get_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertCrashdumpRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class CrashdumpClass) GetVDI__mock(sessionID SessionRef, self CrashdumpRef) (_retval VDIRef, _err error) {
	log.Println("Crashdump.GetVDI not mocked")
	_err = errors.New("Crashdump.GetVDI not mocked")
	return
}
// Get the VDI field of the given crashdump.
func (_class CrashdumpClass) GetVDI(sessionID SessionRef, self CrashdumpRef) (_retval VDIRef, _err error) {
	if (IsMock) {
		return _class.GetVDI__mock(sessionID, self)
	}	
	_method := "crashdump.get_VDI"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertCrashdumpRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class CrashdumpClass) GetVM__mock(sessionID SessionRef, self CrashdumpRef) (_retval VMRef, _err error) {
	log.Println("Crashdump.GetVM not mocked")
	_err = errors.New("Crashdump.GetVM not mocked")
	return
}
// Get the VM field of the given crashdump.
func (_class CrashdumpClass) GetVM(sessionID SessionRef, self CrashdumpRef) (_retval VMRef, _err error) {
	if (IsMock) {
		return _class.GetVM__mock(sessionID, self)
	}	
	_method := "crashdump.get_VM"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertCrashdumpRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class CrashdumpClass) GetUUID__mock(sessionID SessionRef, self CrashdumpRef) (_retval string, _err error) {
	log.Println("Crashdump.GetUUID not mocked")
	_err = errors.New("Crashdump.GetUUID not mocked")
	return
}
// Get the uuid field of the given crashdump.
func (_class CrashdumpClass) GetUUID(sessionID SessionRef, self CrashdumpRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetUUID__mock(sessionID, self)
	}	
	_method := "crashdump.get_uuid"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertCrashdumpRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class CrashdumpClass) GetByUUID__mock(sessionID SessionRef, uuid string) (_retval CrashdumpRef, _err error) {
	log.Println("Crashdump.GetByUUID not mocked")
	_err = errors.New("Crashdump.GetByUUID not mocked")
	return
}
// Get a reference to the crashdump instance with the specified UUID.
func (_class CrashdumpClass) GetByUUID(sessionID SessionRef, uuid string) (_retval CrashdumpRef, _err error) {
	if (IsMock) {
		return _class.GetByUUID__mock(sessionID, uuid)
	}	
	_method := "crashdump.get_by_uuid"
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
	_retval, _err = convertCrashdumpRefToGo(_method + " -> ", _result.Value)
	return
}

func (_class CrashdumpClass) GetRecord__mock(sessionID SessionRef, self CrashdumpRef) (_retval CrashdumpRecord, _err error) {
	log.Println("Crashdump.GetRecord not mocked")
	_err = errors.New("Crashdump.GetRecord not mocked")
	return
}
// Get a record containing the current state of the given crashdump.
func (_class CrashdumpClass) GetRecord(sessionID SessionRef, self CrashdumpRef) (_retval CrashdumpRecord, _err error) {
	if (IsMock) {
		return _class.GetRecord__mock(sessionID, self)
	}	
	_method := "crashdump.get_record"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertCrashdumpRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertCrashdumpRecordToGo(_method + " -> ", _result.Value)
	return
}
