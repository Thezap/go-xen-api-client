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


func MessageClassGetAllRecordsWhereMockDefault(sessionID SessionRef, expr string) (_retval map[MessageRef]MessageRecord, _err error) {
	log.Println("Message.GetAllRecordsWhere not mocked")
	_err = errors.New("Message.GetAllRecordsWhere not mocked")
	return
}

var MessageClassGetAllRecordsWhereMockedCallback = MessageClassGetAllRecordsWhereMockDefault

func (_class MessageClass) GetAllRecordsWhereMock(sessionID SessionRef, expr string) (_retval map[MessageRef]MessageRecord, _err error) {
	return MessageClassGetAllRecordsWhereMockedCallback(sessionID, expr)
}
// 
func (_class MessageClass) GetAllRecordsWhere(sessionID SessionRef, expr string) (_retval map[MessageRef]MessageRecord, _err error) {
	if IsMock {
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


func MessageClassGetAllRecordsMockDefault(sessionID SessionRef) (_retval map[MessageRef]MessageRecord, _err error) {
	log.Println("Message.GetAllRecords not mocked")
	_err = errors.New("Message.GetAllRecords not mocked")
	return
}

var MessageClassGetAllRecordsMockedCallback = MessageClassGetAllRecordsMockDefault

func (_class MessageClass) GetAllRecordsMock(sessionID SessionRef) (_retval map[MessageRef]MessageRecord, _err error) {
	return MessageClassGetAllRecordsMockedCallback(sessionID)
}
// 
func (_class MessageClass) GetAllRecords(sessionID SessionRef) (_retval map[MessageRef]MessageRecord, _err error) {
	if IsMock {
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


func MessageClassGetByUUIDMockDefault(sessionID SessionRef, uuid string) (_retval MessageRef, _err error) {
	log.Println("Message.GetByUUID not mocked")
	_err = errors.New("Message.GetByUUID not mocked")
	return
}

var MessageClassGetByUUIDMockedCallback = MessageClassGetByUUIDMockDefault

func (_class MessageClass) GetByUUIDMock(sessionID SessionRef, uuid string) (_retval MessageRef, _err error) {
	return MessageClassGetByUUIDMockedCallback(sessionID, uuid)
}
// 
func (_class MessageClass) GetByUUID(sessionID SessionRef, uuid string) (_retval MessageRef, _err error) {
	if IsMock {
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


func MessageClassGetRecordMockDefault(sessionID SessionRef, self MessageRef) (_retval MessageRecord, _err error) {
	log.Println("Message.GetRecord not mocked")
	_err = errors.New("Message.GetRecord not mocked")
	return
}

var MessageClassGetRecordMockedCallback = MessageClassGetRecordMockDefault

func (_class MessageClass) GetRecordMock(sessionID SessionRef, self MessageRef) (_retval MessageRecord, _err error) {
	return MessageClassGetRecordMockedCallback(sessionID, self)
}
// 
func (_class MessageClass) GetRecord(sessionID SessionRef, self MessageRef) (_retval MessageRecord, _err error) {
	if IsMock {
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


func MessageClassGetSinceMockDefault(sessionID SessionRef, since time.Time) (_retval map[MessageRef]MessageRecord, _err error) {
	log.Println("Message.GetSince not mocked")
	_err = errors.New("Message.GetSince not mocked")
	return
}

var MessageClassGetSinceMockedCallback = MessageClassGetSinceMockDefault

func (_class MessageClass) GetSinceMock(sessionID SessionRef, since time.Time) (_retval map[MessageRef]MessageRecord, _err error) {
	return MessageClassGetSinceMockedCallback(sessionID, since)
}
// 
func (_class MessageClass) GetSince(sessionID SessionRef, since time.Time) (_retval map[MessageRef]MessageRecord, _err error) {
	if IsMock {
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


func MessageClassGetAllMockDefault(sessionID SessionRef) (_retval []MessageRef, _err error) {
	log.Println("Message.GetAll not mocked")
	_err = errors.New("Message.GetAll not mocked")
	return
}

var MessageClassGetAllMockedCallback = MessageClassGetAllMockDefault

func (_class MessageClass) GetAllMock(sessionID SessionRef) (_retval []MessageRef, _err error) {
	return MessageClassGetAllMockedCallback(sessionID)
}
// 
func (_class MessageClass) GetAll(sessionID SessionRef) (_retval []MessageRef, _err error) {
	if IsMock {
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


func MessageClassGetMockDefault(sessionID SessionRef, cls Cls, objUUID string, since time.Time) (_retval map[MessageRef]MessageRecord, _err error) {
	log.Println("Message.Get not mocked")
	_err = errors.New("Message.Get not mocked")
	return
}

var MessageClassGetMockedCallback = MessageClassGetMockDefault

func (_class MessageClass) GetMock(sessionID SessionRef, cls Cls, objUUID string, since time.Time) (_retval map[MessageRef]MessageRecord, _err error) {
	return MessageClassGetMockedCallback(sessionID, cls, objUUID, since)
}
// 
func (_class MessageClass) Get(sessionID SessionRef, cls Cls, objUUID string, since time.Time) (_retval map[MessageRef]MessageRecord, _err error) {
	if IsMock {
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


func MessageClassDestroyMockDefault(sessionID SessionRef, self MessageRef) (_err error) {
	log.Println("Message.Destroy not mocked")
	_err = errors.New("Message.Destroy not mocked")
	return
}

var MessageClassDestroyMockedCallback = MessageClassDestroyMockDefault

func (_class MessageClass) DestroyMock(sessionID SessionRef, self MessageRef) (_err error) {
	return MessageClassDestroyMockedCallback(sessionID, self)
}
// 
func (_class MessageClass) Destroy(sessionID SessionRef, self MessageRef) (_err error) {
	if IsMock {
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


func MessageClassCreateMockDefault(sessionID SessionRef, name string, priority int, cls Cls, objUUID string, body string) (_retval MessageRef, _err error) {
	log.Println("Message.Create not mocked")
	_err = errors.New("Message.Create not mocked")
	return
}

var MessageClassCreateMockedCallback = MessageClassCreateMockDefault

func (_class MessageClass) CreateMock(sessionID SessionRef, name string, priority int, cls Cls, objUUID string, body string) (_retval MessageRef, _err error) {
	return MessageClassCreateMockedCallback(sessionID, name, priority, cls, objUUID, body)
}
// 
func (_class MessageClass) Create(sessionID SessionRef, name string, priority int, cls Cls, objUUID string, body string) (_retval MessageRef, _err error) {
	if IsMock {
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
