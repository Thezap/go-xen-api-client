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

type DRTaskRecord struct {
  // Unique identifier/object reference
	UUID string
  // All SRs introduced by this appliance
	IntroducedSRs []SRRef
}

type DRTaskRef string

// DR task
type DRTaskClass struct {
	client *Client
}


var DRTaskClass_GetAllRecordsMockedCallback = func (sessionID SessionRef) (_retval map[DRTaskRef]DRTaskRecord, _err error) {
	log.Println("DRTask.GetAllRecords not mocked")
	_err = errors.New("DRTask.GetAllRecords not mocked")
	return
}

func (_class DRTaskClass) GetAllRecordsMock(sessionID SessionRef) (_retval map[DRTaskRef]DRTaskRecord, _err error) {
	return DRTaskClass_GetAllRecordsMockedCallback(sessionID)
}
// Return a map of DR_task references to DR_task records for all DR_tasks known to the system.
func (_class DRTaskClass) GetAllRecords(sessionID SessionRef) (_retval map[DRTaskRef]DRTaskRecord, _err error) {
	if (IsMock) {
		return _class.GetAllRecordsMock(sessionID)
	}	
	_method := "DR_task.get_all_records"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertDRTaskRefToDRTaskRecordMapToGo(_method + " -> ", _result.Value)
	return
}


var DRTaskClass_GetAllMockedCallback = func (sessionID SessionRef) (_retval []DRTaskRef, _err error) {
	log.Println("DRTask.GetAll not mocked")
	_err = errors.New("DRTask.GetAll not mocked")
	return
}

func (_class DRTaskClass) GetAllMock(sessionID SessionRef) (_retval []DRTaskRef, _err error) {
	return DRTaskClass_GetAllMockedCallback(sessionID)
}
// Return a list of all the DR_tasks known to the system.
func (_class DRTaskClass) GetAll(sessionID SessionRef) (_retval []DRTaskRef, _err error) {
	if (IsMock) {
		return _class.GetAllMock(sessionID)
	}	
	_method := "DR_task.get_all"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertDRTaskRefSetToGo(_method + " -> ", _result.Value)
	return
}


var DRTaskClass_DestroyMockedCallback = func (sessionID SessionRef, self DRTaskRef) (_err error) {
	log.Println("DRTask.Destroy not mocked")
	_err = errors.New("DRTask.Destroy not mocked")
	return
}

func (_class DRTaskClass) DestroyMock(sessionID SessionRef, self DRTaskRef) (_err error) {
	return DRTaskClass_DestroyMockedCallback(sessionID, self)
}
// Destroy the disaster recovery task, detaching and forgetting any SRs introduced which are no longer required
func (_class DRTaskClass) Destroy(sessionID SessionRef, self DRTaskRef) (_err error) {
	if (IsMock) {
		return _class.DestroyMock(sessionID, self)
	}	
	_method := "DR_task.destroy"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertDRTaskRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg)
	return
}


var DRTaskClass_CreateMockedCallback = func (sessionID SessionRef, atype string, deviceConfig map[string]string, whitelist []string) (_retval DRTaskRef, _err error) {
	log.Println("DRTask.Create not mocked")
	_err = errors.New("DRTask.Create not mocked")
	return
}

func (_class DRTaskClass) CreateMock(sessionID SessionRef, atype string, deviceConfig map[string]string, whitelist []string) (_retval DRTaskRef, _err error) {
	return DRTaskClass_CreateMockedCallback(sessionID, atype, deviceConfig, whitelist)
}
// Create a disaster recovery task which will query the supplied list of devices
func (_class DRTaskClass) Create(sessionID SessionRef, atype string, deviceConfig map[string]string, whitelist []string) (_retval DRTaskRef, _err error) {
	if (IsMock) {
		return _class.CreateMock(sessionID, atype, deviceConfig, whitelist)
	}	
	_method := "DR_task.create"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_atypeArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "type"), atype)
	if _err != nil {
		return
	}
	_deviceConfigArg, _err := convertStringToStringMapToXen(fmt.Sprintf("%s(%s)", _method, "device_config"), deviceConfig)
	if _err != nil {
		return
	}
	_whitelistArg, _err := convertStringSetToXen(fmt.Sprintf("%s(%s)", _method, "whitelist"), whitelist)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _atypeArg, _deviceConfigArg, _whitelistArg)
	if _err != nil {
		return
	}
	_retval, _err = convertDRTaskRefToGo(_method + " -> ", _result.Value)
	return
}


var DRTaskClass_GetIntroducedSRsMockedCallback = func (sessionID SessionRef, self DRTaskRef) (_retval []SRRef, _err error) {
	log.Println("DRTask.GetIntroducedSRs not mocked")
	_err = errors.New("DRTask.GetIntroducedSRs not mocked")
	return
}

func (_class DRTaskClass) GetIntroducedSRsMock(sessionID SessionRef, self DRTaskRef) (_retval []SRRef, _err error) {
	return DRTaskClass_GetIntroducedSRsMockedCallback(sessionID, self)
}
// Get the introduced_SRs field of the given DR_task.
func (_class DRTaskClass) GetIntroducedSRs(sessionID SessionRef, self DRTaskRef) (_retval []SRRef, _err error) {
	if (IsMock) {
		return _class.GetIntroducedSRsMock(sessionID, self)
	}	
	_method := "DR_task.get_introduced_SRs"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertDRTaskRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertSRRefSetToGo(_method + " -> ", _result.Value)
	return
}


var DRTaskClass_GetUUIDMockedCallback = func (sessionID SessionRef, self DRTaskRef) (_retval string, _err error) {
	log.Println("DRTask.GetUUID not mocked")
	_err = errors.New("DRTask.GetUUID not mocked")
	return
}

func (_class DRTaskClass) GetUUIDMock(sessionID SessionRef, self DRTaskRef) (_retval string, _err error) {
	return DRTaskClass_GetUUIDMockedCallback(sessionID, self)
}
// Get the uuid field of the given DR_task.
func (_class DRTaskClass) GetUUID(sessionID SessionRef, self DRTaskRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetUUIDMock(sessionID, self)
	}	
	_method := "DR_task.get_uuid"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertDRTaskRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var DRTaskClass_GetByUUIDMockedCallback = func (sessionID SessionRef, uuid string) (_retval DRTaskRef, _err error) {
	log.Println("DRTask.GetByUUID not mocked")
	_err = errors.New("DRTask.GetByUUID not mocked")
	return
}

func (_class DRTaskClass) GetByUUIDMock(sessionID SessionRef, uuid string) (_retval DRTaskRef, _err error) {
	return DRTaskClass_GetByUUIDMockedCallback(sessionID, uuid)
}
// Get a reference to the DR_task instance with the specified UUID.
func (_class DRTaskClass) GetByUUID(sessionID SessionRef, uuid string) (_retval DRTaskRef, _err error) {
	if (IsMock) {
		return _class.GetByUUIDMock(sessionID, uuid)
	}	
	_method := "DR_task.get_by_uuid"
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
	_retval, _err = convertDRTaskRefToGo(_method + " -> ", _result.Value)
	return
}


var DRTaskClass_GetRecordMockedCallback = func (sessionID SessionRef, self DRTaskRef) (_retval DRTaskRecord, _err error) {
	log.Println("DRTask.GetRecord not mocked")
	_err = errors.New("DRTask.GetRecord not mocked")
	return
}

func (_class DRTaskClass) GetRecordMock(sessionID SessionRef, self DRTaskRef) (_retval DRTaskRecord, _err error) {
	return DRTaskClass_GetRecordMockedCallback(sessionID, self)
}
// Get a record containing the current state of the given DR_task.
func (_class DRTaskClass) GetRecord(sessionID SessionRef, self DRTaskRef) (_retval DRTaskRecord, _err error) {
	if (IsMock) {
		return _class.GetRecordMock(sessionID, self)
	}	
	_method := "DR_task.get_record"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertDRTaskRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertDRTaskRecordToGo(_method + " -> ", _result.Value)
	return
}
