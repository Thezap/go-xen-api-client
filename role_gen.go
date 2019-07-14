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

type RoleRecord struct {
  // Unique identifier/object reference
	UUID string
  // a short user-friendly name for the role
	NameLabel string
  // what this role is for
	NameDescription string
  // a list of pointers to other roles or permissions
	Subroles []RoleRef
}

type RoleRef string

// A set of permissions associated with a subject
type RoleClass struct {
	client *Client
}


var RoleClass_GetAllRecordsMockedCallback = func (sessionID SessionRef) (_retval map[RoleRef]RoleRecord, _err error) {
	log.Println("Role.GetAllRecords not mocked")
	_err = errors.New("Role.GetAllRecords not mocked")
	return
}

func (_class RoleClass) GetAllRecordsMock(sessionID SessionRef) (_retval map[RoleRef]RoleRecord, _err error) {
	return RoleClass_GetAllRecordsMockedCallback(sessionID)
}
// Return a map of role references to role records for all roles known to the system.
func (_class RoleClass) GetAllRecords(sessionID SessionRef) (_retval map[RoleRef]RoleRecord, _err error) {
	if (IsMock) {
		return _class.GetAllRecordsMock(sessionID)
	}	
	_method := "role.get_all_records"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertRoleRefToRoleRecordMapToGo(_method + " -> ", _result.Value)
	return
}


var RoleClass_GetAllMockedCallback = func (sessionID SessionRef) (_retval []RoleRef, _err error) {
	log.Println("Role.GetAll not mocked")
	_err = errors.New("Role.GetAll not mocked")
	return
}

func (_class RoleClass) GetAllMock(sessionID SessionRef) (_retval []RoleRef, _err error) {
	return RoleClass_GetAllMockedCallback(sessionID)
}
// Return a list of all the roles known to the system.
func (_class RoleClass) GetAll(sessionID SessionRef) (_retval []RoleRef, _err error) {
	if (IsMock) {
		return _class.GetAllMock(sessionID)
	}	
	_method := "role.get_all"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertRoleRefSetToGo(_method + " -> ", _result.Value)
	return
}


var RoleClass_GetByPermissionNameLabelMockedCallback = func (sessionID SessionRef, label string) (_retval []RoleRef, _err error) {
	log.Println("Role.GetByPermissionNameLabel not mocked")
	_err = errors.New("Role.GetByPermissionNameLabel not mocked")
	return
}

func (_class RoleClass) GetByPermissionNameLabelMock(sessionID SessionRef, label string) (_retval []RoleRef, _err error) {
	return RoleClass_GetByPermissionNameLabelMockedCallback(sessionID, label)
}
// This call returns a list of roles given a permission name
func (_class RoleClass) GetByPermissionNameLabel(sessionID SessionRef, label string) (_retval []RoleRef, _err error) {
	if (IsMock) {
		return _class.GetByPermissionNameLabelMock(sessionID, label)
	}	
	_method := "role.get_by_permission_name_label"
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
	_retval, _err = convertRoleRefSetToGo(_method + " -> ", _result.Value)
	return
}


var RoleClass_GetByPermissionMockedCallback = func (sessionID SessionRef, permission RoleRef) (_retval []RoleRef, _err error) {
	log.Println("Role.GetByPermission not mocked")
	_err = errors.New("Role.GetByPermission not mocked")
	return
}

func (_class RoleClass) GetByPermissionMock(sessionID SessionRef, permission RoleRef) (_retval []RoleRef, _err error) {
	return RoleClass_GetByPermissionMockedCallback(sessionID, permission)
}
// This call returns a list of roles given a permission
func (_class RoleClass) GetByPermission(sessionID SessionRef, permission RoleRef) (_retval []RoleRef, _err error) {
	if (IsMock) {
		return _class.GetByPermissionMock(sessionID, permission)
	}	
	_method := "role.get_by_permission"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_permissionArg, _err := convertRoleRefToXen(fmt.Sprintf("%s(%s)", _method, "permission"), permission)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _permissionArg)
	if _err != nil {
		return
	}
	_retval, _err = convertRoleRefSetToGo(_method + " -> ", _result.Value)
	return
}


var RoleClass_GetPermissionsNameLabelMockedCallback = func (sessionID SessionRef, self RoleRef) (_retval []string, _err error) {
	log.Println("Role.GetPermissionsNameLabel not mocked")
	_err = errors.New("Role.GetPermissionsNameLabel not mocked")
	return
}

func (_class RoleClass) GetPermissionsNameLabelMock(sessionID SessionRef, self RoleRef) (_retval []string, _err error) {
	return RoleClass_GetPermissionsNameLabelMockedCallback(sessionID, self)
}
// This call returns a list of permission names given a role
func (_class RoleClass) GetPermissionsNameLabel(sessionID SessionRef, self RoleRef) (_retval []string, _err error) {
	if (IsMock) {
		return _class.GetPermissionsNameLabelMock(sessionID, self)
	}	
	_method := "role.get_permissions_name_label"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertRoleRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var RoleClass_GetPermissionsMockedCallback = func (sessionID SessionRef, self RoleRef) (_retval []RoleRef, _err error) {
	log.Println("Role.GetPermissions not mocked")
	_err = errors.New("Role.GetPermissions not mocked")
	return
}

func (_class RoleClass) GetPermissionsMock(sessionID SessionRef, self RoleRef) (_retval []RoleRef, _err error) {
	return RoleClass_GetPermissionsMockedCallback(sessionID, self)
}
// This call returns a list of permissions given a role
func (_class RoleClass) GetPermissions(sessionID SessionRef, self RoleRef) (_retval []RoleRef, _err error) {
	if (IsMock) {
		return _class.GetPermissionsMock(sessionID, self)
	}	
	_method := "role.get_permissions"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertRoleRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var RoleClass_GetSubrolesMockedCallback = func (sessionID SessionRef, self RoleRef) (_retval []RoleRef, _err error) {
	log.Println("Role.GetSubroles not mocked")
	_err = errors.New("Role.GetSubroles not mocked")
	return
}

func (_class RoleClass) GetSubrolesMock(sessionID SessionRef, self RoleRef) (_retval []RoleRef, _err error) {
	return RoleClass_GetSubrolesMockedCallback(sessionID, self)
}
// Get the subroles field of the given role.
func (_class RoleClass) GetSubroles(sessionID SessionRef, self RoleRef) (_retval []RoleRef, _err error) {
	if (IsMock) {
		return _class.GetSubrolesMock(sessionID, self)
	}	
	_method := "role.get_subroles"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertRoleRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var RoleClass_GetNameDescriptionMockedCallback = func (sessionID SessionRef, self RoleRef) (_retval string, _err error) {
	log.Println("Role.GetNameDescription not mocked")
	_err = errors.New("Role.GetNameDescription not mocked")
	return
}

func (_class RoleClass) GetNameDescriptionMock(sessionID SessionRef, self RoleRef) (_retval string, _err error) {
	return RoleClass_GetNameDescriptionMockedCallback(sessionID, self)
}
// Get the name/description field of the given role.
func (_class RoleClass) GetNameDescription(sessionID SessionRef, self RoleRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetNameDescriptionMock(sessionID, self)
	}	
	_method := "role.get_name_description"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertRoleRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var RoleClass_GetNameLabelMockedCallback = func (sessionID SessionRef, self RoleRef) (_retval string, _err error) {
	log.Println("Role.GetNameLabel not mocked")
	_err = errors.New("Role.GetNameLabel not mocked")
	return
}

func (_class RoleClass) GetNameLabelMock(sessionID SessionRef, self RoleRef) (_retval string, _err error) {
	return RoleClass_GetNameLabelMockedCallback(sessionID, self)
}
// Get the name/label field of the given role.
func (_class RoleClass) GetNameLabel(sessionID SessionRef, self RoleRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetNameLabelMock(sessionID, self)
	}	
	_method := "role.get_name_label"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertRoleRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var RoleClass_GetUUIDMockedCallback = func (sessionID SessionRef, self RoleRef) (_retval string, _err error) {
	log.Println("Role.GetUUID not mocked")
	_err = errors.New("Role.GetUUID not mocked")
	return
}

func (_class RoleClass) GetUUIDMock(sessionID SessionRef, self RoleRef) (_retval string, _err error) {
	return RoleClass_GetUUIDMockedCallback(sessionID, self)
}
// Get the uuid field of the given role.
func (_class RoleClass) GetUUID(sessionID SessionRef, self RoleRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetUUIDMock(sessionID, self)
	}	
	_method := "role.get_uuid"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertRoleRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var RoleClass_GetByNameLabelMockedCallback = func (sessionID SessionRef, label string) (_retval []RoleRef, _err error) {
	log.Println("Role.GetByNameLabel not mocked")
	_err = errors.New("Role.GetByNameLabel not mocked")
	return
}

func (_class RoleClass) GetByNameLabelMock(sessionID SessionRef, label string) (_retval []RoleRef, _err error) {
	return RoleClass_GetByNameLabelMockedCallback(sessionID, label)
}
// Get all the role instances with the given label.
func (_class RoleClass) GetByNameLabel(sessionID SessionRef, label string) (_retval []RoleRef, _err error) {
	if (IsMock) {
		return _class.GetByNameLabelMock(sessionID, label)
	}	
	_method := "role.get_by_name_label"
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
	_retval, _err = convertRoleRefSetToGo(_method + " -> ", _result.Value)
	return
}


var RoleClass_GetByUUIDMockedCallback = func (sessionID SessionRef, uuid string) (_retval RoleRef, _err error) {
	log.Println("Role.GetByUUID not mocked")
	_err = errors.New("Role.GetByUUID not mocked")
	return
}

func (_class RoleClass) GetByUUIDMock(sessionID SessionRef, uuid string) (_retval RoleRef, _err error) {
	return RoleClass_GetByUUIDMockedCallback(sessionID, uuid)
}
// Get a reference to the role instance with the specified UUID.
func (_class RoleClass) GetByUUID(sessionID SessionRef, uuid string) (_retval RoleRef, _err error) {
	if (IsMock) {
		return _class.GetByUUIDMock(sessionID, uuid)
	}	
	_method := "role.get_by_uuid"
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
	_retval, _err = convertRoleRefToGo(_method + " -> ", _result.Value)
	return
}


var RoleClass_GetRecordMockedCallback = func (sessionID SessionRef, self RoleRef) (_retval RoleRecord, _err error) {
	log.Println("Role.GetRecord not mocked")
	_err = errors.New("Role.GetRecord not mocked")
	return
}

func (_class RoleClass) GetRecordMock(sessionID SessionRef, self RoleRef) (_retval RoleRecord, _err error) {
	return RoleClass_GetRecordMockedCallback(sessionID, self)
}
// Get a record containing the current state of the given role.
func (_class RoleClass) GetRecord(sessionID SessionRef, self RoleRef) (_retval RoleRecord, _err error) {
	if (IsMock) {
		return _class.GetRecordMock(sessionID, self)
	}	
	_method := "role.get_record"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertRoleRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertRoleRecordToGo(_method + " -> ", _result.Value)
	return
}
