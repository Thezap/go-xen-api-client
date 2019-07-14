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

type Cls string

const (
  // VM
	ClsVM Cls = "VM"
  // Host
	ClsHost Cls = "Host"
  // SR
	ClsSR Cls = "SR"
  // Pool
	ClsPool Cls = "Pool"
  // VMPP
	ClsVMPP Cls = "VMPP"
  // VMSS
	ClsVMSS Cls = "VMSS"
  // PVS_proxy
	ClsPVSProxy Cls = "PVS_proxy"
  // VDI
	ClsVDI Cls = "VDI"
)

type MessageRecord struct {
  // Unique identifier/object reference
	UUID string
  // The name of the message
	Name string
  // The message priority, 0 being low priority
	Priority int
  // The class of the object this message is associated with
	Cls Cls
  // The uuid of the object this message is associated with
	ObjUUID string
  // The time at which the message was created
	Timestamp time.Time
  // The body of the message
	Body string
}

type MessageRef string

// An message for the attention of the administrator
type MessageClass struct {
	client *Client
}


var MessageClass_GetAllRecordsWhereMockedCallback = func (sessionID SessionRef, expr string) (_retval map[MessageRef]MessageRecord, _err error) {
	log.Println("Message.GetAllRecordsWhere not mocked")
	_err = errors.New("Message.GetAllRecordsWhere not mocked")
	return
}

func (_class MessageClass) GetAllRecordsWhereMock(sessionID SessionRef, expr string) (_retval map[MessageRef]MessageRecord, _err error) {
	return MessageClass_GetAllRecordsWhereMockedCallback(sessionID, expr)
}
// 
func (_class MessageClass) GetAllRecordsWhere(sessionID SessionRef, expr string) (_retval map[MessageRef]MessageRecord, _err error) {
	if (IsMock) {
		return _class.GetAllRecordsWhereMock(sessionID, expr)
	}	
	_method := "message.get_all_records_where"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_exprArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "expr"), expr)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _exprArg)
	if _err != nil {
		return
	}
	_retval, _err = convertMessageRefToMessageRecordMapToGo(_method + " -> ", _result.Value)
	return
}


var MessageClass_GetAllRecordsMockedCallback = func (sessionID SessionRef) (_retval map[MessageRef]MessageRecord, _err error) {
	log.Println("Message.GetAllRecords not mocked")
	_err = errors.New("Message.GetAllRecords not mocked")
	return
}

func (_class MessageClass) GetAllRecordsMock(sessionID SessionRef) (_retval map[MessageRef]MessageRecord, _err error) {
	return MessageClass_GetAllRecordsMockedCallback(sessionID)
}
// 
func (_class MessageClass) GetAllRecords(sessionID SessionRef) (_retval map[MessageRef]MessageRecord, _err error) {
	if (IsMock) {
		return _class.GetAllRecordsMock(sessionID)
	}	
	_method := "message.get_all_records"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertMessageRefToMessageRecordMapToGo(_method + " -> ", _result.Value)
	return
}


var MessageClass_GetByUUIDMockedCallback = func (sessionID SessionRef, uuid string) (_retval MessageRef, _err error) {
	log.Println("Message.GetByUUID not mocked")
	_err = errors.New("Message.GetByUUID not mocked")
	return
}

func (_class MessageClass) GetByUUIDMock(sessionID SessionRef, uuid string) (_retval MessageRef, _err error) {
	return MessageClass_GetByUUIDMockedCallback(sessionID, uuid)
}
// 
func (_class MessageClass) GetByUUID(sessionID SessionRef, uuid string) (_retval MessageRef, _err error) {
	if (IsMock) {
		return _class.GetByUUIDMock(sessionID, uuid)
	}	
	_method := "message.get_by_uuid"
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
	_retval, _err = convertMessageRefToGo(_method + " -> ", _result.Value)
	return
}


var MessageClass_GetRecordMockedCallback = func (sessionID SessionRef, self MessageRef) (_retval MessageRecord, _err error) {
	log.Println("Message.GetRecord not mocked")
	_err = errors.New("Message.GetRecord not mocked")
	return
}

func (_class MessageClass) GetRecordMock(sessionID SessionRef, self MessageRef) (_retval MessageRecord, _err error) {
	return MessageClass_GetRecordMockedCallback(sessionID, self)
}
// 
func (_class MessageClass) GetRecord(sessionID SessionRef, self MessageRef) (_retval MessageRecord, _err error) {
	if (IsMock) {
		return _class.GetRecordMock(sessionID, self)
	}	
	_method := "message.get_record"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertMessageRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertMessageRecordToGo(_method + " -> ", _result.Value)
	return
}


var MessageClass_GetSinceMockedCallback = func (sessionID SessionRef, since time.Time) (_retval map[MessageRef]MessageRecord, _err error) {
	log.Println("Message.GetSince not mocked")
	_err = errors.New("Message.GetSince not mocked")
	return
}

func (_class MessageClass) GetSinceMock(sessionID SessionRef, since time.Time) (_retval map[MessageRef]MessageRecord, _err error) {
	return MessageClass_GetSinceMockedCallback(sessionID, since)
}
// 
func (_class MessageClass) GetSince(sessionID SessionRef, since time.Time) (_retval map[MessageRef]MessageRecord, _err error) {
	if (IsMock) {
		return _class.GetSinceMock(sessionID, since)
	}	
	_method := "message.get_since"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_sinceArg, _err := convertTimeToXen(fmt.Sprintf("%s(%s)", _method, "since"), since)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _sinceArg)
	if _err != nil {
		return
	}
	_retval, _err = convertMessageRefToMessageRecordMapToGo(_method + " -> ", _result.Value)
	return
}


var MessageClass_GetAllMockedCallback = func (sessionID SessionRef) (_retval []MessageRef, _err error) {
	log.Println("Message.GetAll not mocked")
	_err = errors.New("Message.GetAll not mocked")
	return
}

func (_class MessageClass) GetAllMock(sessionID SessionRef) (_retval []MessageRef, _err error) {
	return MessageClass_GetAllMockedCallback(sessionID)
}
// 
func (_class MessageClass) GetAll(sessionID SessionRef) (_retval []MessageRef, _err error) {
	if (IsMock) {
		return _class.GetAllMock(sessionID)
	}	
	_method := "message.get_all"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertMessageRefSetToGo(_method + " -> ", _result.Value)
	return
}


var MessageClass_GetMockedCallback = func (sessionID SessionRef, cls Cls, objUUID string, since time.Time) (_retval map[MessageRef]MessageRecord, _err error) {
	log.Println("Message.Get not mocked")
	_err = errors.New("Message.Get not mocked")
	return
}

func (_class MessageClass) GetMock(sessionID SessionRef, cls Cls, objUUID string, since time.Time) (_retval map[MessageRef]MessageRecord, _err error) {
	return MessageClass_GetMockedCallback(sessionID, cls, objUUID, since)
}
// 
func (_class MessageClass) Get(sessionID SessionRef, cls Cls, objUUID string, since time.Time) (_retval map[MessageRef]MessageRecord, _err error) {
	if (IsMock) {
		return _class.GetMock(sessionID, cls, objUUID, since)
	}	
	_method := "message.get"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_clsArg, _err := convertEnumClsToXen(fmt.Sprintf("%s(%s)", _method, "cls"), cls)
	if _err != nil {
		return
	}
	_objUUIDArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "obj_uuid"), objUUID)
	if _err != nil {
		return
	}
	_sinceArg, _err := convertTimeToXen(fmt.Sprintf("%s(%s)", _method, "since"), since)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _clsArg, _objUUIDArg, _sinceArg)
	if _err != nil {
		return
	}
	_retval, _err = convertMessageRefToMessageRecordMapToGo(_method + " -> ", _result.Value)
	return
}


var MessageClass_DestroyMockedCallback = func (sessionID SessionRef, self MessageRef) (_err error) {
	log.Println("Message.Destroy not mocked")
	_err = errors.New("Message.Destroy not mocked")
	return
}

func (_class MessageClass) DestroyMock(sessionID SessionRef, self MessageRef) (_err error) {
	return MessageClass_DestroyMockedCallback(sessionID, self)
}
// 
func (_class MessageClass) Destroy(sessionID SessionRef, self MessageRef) (_err error) {
	if (IsMock) {
		return _class.DestroyMock(sessionID, self)
	}	
	_method := "message.destroy"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertMessageRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg)
	return
}


var MessageClass_CreateMockedCallback = func (sessionID SessionRef, name string, priority int, cls Cls, objUUID string, body string) (_retval MessageRef, _err error) {
	log.Println("Message.Create not mocked")
	_err = errors.New("Message.Create not mocked")
	return
}

func (_class MessageClass) CreateMock(sessionID SessionRef, name string, priority int, cls Cls, objUUID string, body string) (_retval MessageRef, _err error) {
	return MessageClass_CreateMockedCallback(sessionID, name, priority, cls, objUUID, body)
}
// 
func (_class MessageClass) Create(sessionID SessionRef, name string, priority int, cls Cls, objUUID string, body string) (_retval MessageRef, _err error) {
	if (IsMock) {
		return _class.CreateMock(sessionID, name, priority, cls, objUUID, body)
	}	
	_method := "message.create"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_nameArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "name"), name)
	if _err != nil {
		return
	}
	_priorityArg, _err := convertIntToXen(fmt.Sprintf("%s(%s)", _method, "priority"), priority)
	if _err != nil {
		return
	}
	_clsArg, _err := convertEnumClsToXen(fmt.Sprintf("%s(%s)", _method, "cls"), cls)
	if _err != nil {
		return
	}
	_objUUIDArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "obj_uuid"), objUUID)
	if _err != nil {
		return
	}
	_bodyArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "body"), body)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _nameArg, _priorityArg, _clsArg, _objUUIDArg, _bodyArg)
	if _err != nil {
		return
	}
	_retval, _err = convertMessageRefToGo(_method + " -> ", _result.Value)
	return
}
