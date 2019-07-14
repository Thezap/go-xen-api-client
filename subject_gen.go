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

type SubjectRecord struct {
  // Unique identifier/object reference
	UUID string
  // the subject identifier, unique in the external directory service
	SubjectIdentifier string
  // additional configuration
	OtherConfig map[string]string
  // the roles associated with this subject
	Roles []RoleRef
}

type SubjectRef string

// A user or group that can log in xapi
type SubjectClass struct {
	client *Client
}


var SubjectClass_GetAllRecordsMockedCallback = func (sessionID SessionRef) (_retval map[SubjectRef]SubjectRecord, _err error) {
	log.Println("Subject.GetAllRecords not mocked")
	_err = errors.New("Subject.GetAllRecords not mocked")
	return
}

func (_class SubjectClass) GetAllRecordsMock(sessionID SessionRef) (_retval map[SubjectRef]SubjectRecord, _err error) {
	return SubjectClass_GetAllRecordsMockedCallback(sessionID)
}
// Return a map of subject references to subject records for all subjects known to the system.
func (_class SubjectClass) GetAllRecords(sessionID SessionRef) (_retval map[SubjectRef]SubjectRecord, _err error) {
	if (IsMock) {
		return _class.GetAllRecordsMock(sessionID)
	}	
	_method := "subject.get_all_records"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertSubjectRefToSubjectRecordMapToGo(_method + " -> ", _result.Value)
	return
}


var SubjectClass_GetAllMockedCallback = func (sessionID SessionRef) (_retval []SubjectRef, _err error) {
	log.Println("Subject.GetAll not mocked")
	_err = errors.New("Subject.GetAll not mocked")
	return
}

func (_class SubjectClass) GetAllMock(sessionID SessionRef) (_retval []SubjectRef, _err error) {
	return SubjectClass_GetAllMockedCallback(sessionID)
}
// Return a list of all the subjects known to the system.
func (_class SubjectClass) GetAll(sessionID SessionRef) (_retval []SubjectRef, _err error) {
	if (IsMock) {
		return _class.GetAllMock(sessionID)
	}	
	_method := "subject.get_all"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertSubjectRefSetToGo(_method + " -> ", _result.Value)
	return
}


var SubjectClass_GetPermissionsNameLabelMockedCallback = func (sessionID SessionRef, self SubjectRef) (_retval []string, _err error) {
	log.Println("Subject.GetPermissionsNameLabel not mocked")
	_err = errors.New("Subject.GetPermissionsNameLabel not mocked")
	return
}

func (_class SubjectClass) GetPermissionsNameLabelMock(sessionID SessionRef, self SubjectRef) (_retval []string, _err error) {
	return SubjectClass_GetPermissionsNameLabelMockedCallback(sessionID, self)
}
// This call returns a list of permission names given a subject
func (_class SubjectClass) GetPermissionsNameLabel(sessionID SessionRef, self SubjectRef) (_retval []string, _err error) {
	if (IsMock) {
		return _class.GetPermissionsNameLabelMock(sessionID, self)
	}	
	_method := "subject.get_permissions_name_label"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertSubjectRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var SubjectClass_RemoveFromRolesMockedCallback = func (sessionID SessionRef, self SubjectRef, role RoleRef) (_err error) {
	log.Println("Subject.RemoveFromRoles not mocked")
	_err = errors.New("Subject.RemoveFromRoles not mocked")
	return
}

func (_class SubjectClass) RemoveFromRolesMock(sessionID SessionRef, self SubjectRef, role RoleRef) (_err error) {
	return SubjectClass_RemoveFromRolesMockedCallback(sessionID, self, role)
}
// This call removes a role from a subject
func (_class SubjectClass) RemoveFromRoles(sessionID SessionRef, self SubjectRef, role RoleRef) (_err error) {
	if (IsMock) {
		return _class.RemoveFromRolesMock(sessionID, self, role)
	}	
	_method := "subject.remove_from_roles"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertSubjectRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_roleArg, _err := convertRoleRefToXen(fmt.Sprintf("%s(%s)", _method, "role"), role)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _roleArg)
	return
}


var SubjectClass_AddToRolesMockedCallback = func (sessionID SessionRef, self SubjectRef, role RoleRef) (_err error) {
	log.Println("Subject.AddToRoles not mocked")
	_err = errors.New("Subject.AddToRoles not mocked")
	return
}

func (_class SubjectClass) AddToRolesMock(sessionID SessionRef, self SubjectRef, role RoleRef) (_err error) {
	return SubjectClass_AddToRolesMockedCallback(sessionID, self, role)
}
// This call adds a new role to a subject
func (_class SubjectClass) AddToRoles(sessionID SessionRef, self SubjectRef, role RoleRef) (_err error) {
	if (IsMock) {
		return _class.AddToRolesMock(sessionID, self, role)
	}	
	_method := "subject.add_to_roles"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertSubjectRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_roleArg, _err := convertRoleRefToXen(fmt.Sprintf("%s(%s)", _method, "role"), role)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _roleArg)
	return
}


var SubjectClass_GetRolesMockedCallback = func (sessionID SessionRef, self SubjectRef) (_retval []RoleRef, _err error) {
	log.Println("Subject.GetRoles not mocked")
	_err = errors.New("Subject.GetRoles not mocked")
	return
}

func (_class SubjectClass) GetRolesMock(sessionID SessionRef, self SubjectRef) (_retval []RoleRef, _err error) {
	return SubjectClass_GetRolesMockedCallback(sessionID, self)
}
// Get the roles field of the given subject.
func (_class SubjectClass) GetRoles(sessionID SessionRef, self SubjectRef) (_retval []RoleRef, _err error) {
	if (IsMock) {
		return _class.GetRolesMock(sessionID, self)
	}	
	_method := "subject.get_roles"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertSubjectRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertRoleRefSetToGo(_method + " -> ", _result.Value)
	return
}


var SubjectClass_GetOtherConfigMockedCallback = func (sessionID SessionRef, self SubjectRef) (_retval map[string]string, _err error) {
	log.Println("Subject.GetOtherConfig not mocked")
	_err = errors.New("Subject.GetOtherConfig not mocked")
	return
}

func (_class SubjectClass) GetOtherConfigMock(sessionID SessionRef, self SubjectRef) (_retval map[string]string, _err error) {
	return SubjectClass_GetOtherConfigMockedCallback(sessionID, self)
}
// Get the other_config field of the given subject.
func (_class SubjectClass) GetOtherConfig(sessionID SessionRef, self SubjectRef) (_retval map[string]string, _err error) {
	if (IsMock) {
		return _class.GetOtherConfigMock(sessionID, self)
	}	
	_method := "subject.get_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertSubjectRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var SubjectClass_GetSubjectIdentifierMockedCallback = func (sessionID SessionRef, self SubjectRef) (_retval string, _err error) {
	log.Println("Subject.GetSubjectIdentifier not mocked")
	_err = errors.New("Subject.GetSubjectIdentifier not mocked")
	return
}

func (_class SubjectClass) GetSubjectIdentifierMock(sessionID SessionRef, self SubjectRef) (_retval string, _err error) {
	return SubjectClass_GetSubjectIdentifierMockedCallback(sessionID, self)
}
// Get the subject_identifier field of the given subject.
func (_class SubjectClass) GetSubjectIdentifier(sessionID SessionRef, self SubjectRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetSubjectIdentifierMock(sessionID, self)
	}	
	_method := "subject.get_subject_identifier"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertSubjectRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var SubjectClass_GetUUIDMockedCallback = func (sessionID SessionRef, self SubjectRef) (_retval string, _err error) {
	log.Println("Subject.GetUUID not mocked")
	_err = errors.New("Subject.GetUUID not mocked")
	return
}

func (_class SubjectClass) GetUUIDMock(sessionID SessionRef, self SubjectRef) (_retval string, _err error) {
	return SubjectClass_GetUUIDMockedCallback(sessionID, self)
}
// Get the uuid field of the given subject.
func (_class SubjectClass) GetUUID(sessionID SessionRef, self SubjectRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetUUIDMock(sessionID, self)
	}	
	_method := "subject.get_uuid"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertSubjectRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var SubjectClass_DestroyMockedCallback = func (sessionID SessionRef, self SubjectRef) (_err error) {
	log.Println("Subject.Destroy not mocked")
	_err = errors.New("Subject.Destroy not mocked")
	return
}

func (_class SubjectClass) DestroyMock(sessionID SessionRef, self SubjectRef) (_err error) {
	return SubjectClass_DestroyMockedCallback(sessionID, self)
}
// Destroy the specified subject instance.
func (_class SubjectClass) Destroy(sessionID SessionRef, self SubjectRef) (_err error) {
	if (IsMock) {
		return _class.DestroyMock(sessionID, self)
	}	
	_method := "subject.destroy"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertSubjectRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg)
	return
}


var SubjectClass_CreateMockedCallback = func (sessionID SessionRef, args SubjectRecord) (_retval SubjectRef, _err error) {
	log.Println("Subject.Create not mocked")
	_err = errors.New("Subject.Create not mocked")
	return
}

func (_class SubjectClass) CreateMock(sessionID SessionRef, args SubjectRecord) (_retval SubjectRef, _err error) {
	return SubjectClass_CreateMockedCallback(sessionID, args)
}
// Create a new subject instance, and return its handle.
// The constructor args are: subject_identifier, other_config (* = non-optional).
func (_class SubjectClass) Create(sessionID SessionRef, args SubjectRecord) (_retval SubjectRef, _err error) {
	if (IsMock) {
		return _class.CreateMock(sessionID, args)
	}	
	_method := "subject.create"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_argsArg, _err := convertSubjectRecordToXen(fmt.Sprintf("%s(%s)", _method, "args"), args)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _argsArg)
	if _err != nil {
		return
	}
	_retval, _err = convertSubjectRefToGo(_method + " -> ", _result.Value)
	return
}


var SubjectClass_GetByUUIDMockedCallback = func (sessionID SessionRef, uuid string) (_retval SubjectRef, _err error) {
	log.Println("Subject.GetByUUID not mocked")
	_err = errors.New("Subject.GetByUUID not mocked")
	return
}

func (_class SubjectClass) GetByUUIDMock(sessionID SessionRef, uuid string) (_retval SubjectRef, _err error) {
	return SubjectClass_GetByUUIDMockedCallback(sessionID, uuid)
}
// Get a reference to the subject instance with the specified UUID.
func (_class SubjectClass) GetByUUID(sessionID SessionRef, uuid string) (_retval SubjectRef, _err error) {
	if (IsMock) {
		return _class.GetByUUIDMock(sessionID, uuid)
	}	
	_method := "subject.get_by_uuid"
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
	_retval, _err = convertSubjectRefToGo(_method + " -> ", _result.Value)
	return
}


var SubjectClass_GetRecordMockedCallback = func (sessionID SessionRef, self SubjectRef) (_retval SubjectRecord, _err error) {
	log.Println("Subject.GetRecord not mocked")
	_err = errors.New("Subject.GetRecord not mocked")
	return
}

func (_class SubjectClass) GetRecordMock(sessionID SessionRef, self SubjectRef) (_retval SubjectRecord, _err error) {
	return SubjectClass_GetRecordMockedCallback(sessionID, self)
}
// Get a record containing the current state of the given subject.
func (_class SubjectClass) GetRecord(sessionID SessionRef, self SubjectRef) (_retval SubjectRecord, _err error) {
	if (IsMock) {
		return _class.GetRecordMock(sessionID, self)
	}	
	_method := "subject.get_record"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertSubjectRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertSubjectRecordToGo(_method + " -> ", _result.Value)
	return
}
