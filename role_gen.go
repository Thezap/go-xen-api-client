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

func (_class RoleClass) GetAllRecords__mock(sessionID SessionRef) (_retval map[RoleRef]RoleRecord, _err error) {
	log.Println("Role.GetAllRecords not mocked")
	_err = errors.New("Role.GetAllRecords not mocked")
	return
}
// Return a map of role references to role records for all roles known to the system.
func (_class RoleClass) GetAllRecords(sessionID SessionRef) (_retval map[RoleRef]RoleRecord, _err error) {
	if (IsMock) {
		return _class.GetAllRecords__mock(sessionID)
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

func (_class RoleClass) GetAll__mock(sessionID SessionRef) (_retval []RoleRef, _err error) {
	log.Println("Role.GetAll not mocked")
	_err = errors.New("Role.GetAll not mocked")
	return
}
// Return a list of all the roles known to the system.
func (_class RoleClass) GetAll(sessionID SessionRef) (_retval []RoleRef, _err error) {
	if (IsMock) {
		return _class.GetAll__mock(sessionID)
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

func (_class RoleClass) GetByPermissionNameLabel__mock(sessionID SessionRef, label string) (_retval []RoleRef, _err error) {
	log.Println("Role.GetByPermissionNameLabel not mocked")
	_err = errors.New("Role.GetByPermissionNameLabel not mocked")
	return
}
// This call returns a list of roles given a permission name
func (_class RoleClass) GetByPermissionNameLabel(sessionID SessionRef, label string) (_retval []RoleRef, _err error) {
	if (IsMock) {
		return _class.GetByPermissionNameLabel__mock(sessionID, label)
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

func (_class RoleClass) GetByPermission__mock(sessionID SessionRef, permission RoleRef) (_retval []RoleRef, _err error) {
	log.Println("Role.GetByPermission not mocked")
	_err = errors.New("Role.GetByPermission not mocked")
	return
}
// This call returns a list of roles given a permission
func (_class RoleClass) GetByPermission(sessionID SessionRef, permission RoleRef) (_retval []RoleRef, _err error) {
	if (IsMock) {
		return _class.GetByPermission__mock(sessionID, permission)
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

func (_class RoleClass) GetPermissionsNameLabel__mock(sessionID SessionRef, self RoleRef) (_retval []string, _err error) {
	log.Println("Role.GetPermissionsNameLabel not mocked")
	_err = errors.New("Role.GetPermissionsNameLabel not mocked")
	return
}
// This call returns a list of permission names given a role
func (_class RoleClass) GetPermissionsNameLabel(sessionID SessionRef, self RoleRef) (_retval []string, _err error) {
	if (IsMock) {
		return _class.GetPermissionsNameLabel__mock(sessionID, self)
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

func (_class RoleClass) GetPermissions__mock(sessionID SessionRef, self RoleRef) (_retval []RoleRef, _err error) {
	log.Println("Role.GetPermissions not mocked")
	_err = errors.New("Role.GetPermissions not mocked")
	return
}
// This call returns a list of permissions given a role
func (_class RoleClass) GetPermissions(sessionID SessionRef, self RoleRef) (_retval []RoleRef, _err error) {
	if (IsMock) {
		return _class.GetPermissions__mock(sessionID, self)
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

func (_class RoleClass) GetSubroles__mock(sessionID SessionRef, self RoleRef) (_retval []RoleRef, _err error) {
	log.Println("Role.GetSubroles not mocked")
	_err = errors.New("Role.GetSubroles not mocked")
	return
}
// Get the subroles field of the given role.
func (_class RoleClass) GetSubroles(sessionID SessionRef, self RoleRef) (_retval []RoleRef, _err error) {
	if (IsMock) {
		return _class.GetSubroles__mock(sessionID, self)
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

func (_class RoleClass) GetNameDescription__mock(sessionID SessionRef, self RoleRef) (_retval string, _err error) {
	log.Println("Role.GetNameDescription not mocked")
	_err = errors.New("Role.GetNameDescription not mocked")
	return
}
// Get the name/description field of the given role.
func (_class RoleClass) GetNameDescription(sessionID SessionRef, self RoleRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetNameDescription__mock(sessionID, self)
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

func (_class RoleClass) GetNameLabel__mock(sessionID SessionRef, self RoleRef) (_retval string, _err error) {
	log.Println("Role.GetNameLabel not mocked")
	_err = errors.New("Role.GetNameLabel not mocked")
	return
}
// Get the name/label field of the given role.
func (_class RoleClass) GetNameLabel(sessionID SessionRef, self RoleRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetNameLabel__mock(sessionID, self)
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

func (_class RoleClass) GetUUID__mock(sessionID SessionRef, self RoleRef) (_retval string, _err error) {
	log.Println("Role.GetUUID not mocked")
	_err = errors.New("Role.GetUUID not mocked")
	return
}
// Get the uuid field of the given role.
func (_class RoleClass) GetUUID(sessionID SessionRef, self RoleRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetUUID__mock(sessionID, self)
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

func (_class RoleClass) GetByNameLabel__mock(sessionID SessionRef, label string) (_retval []RoleRef, _err error) {
	log.Println("Role.GetByNameLabel not mocked")
	_err = errors.New("Role.GetByNameLabel not mocked")
	return
}
// Get all the role instances with the given label.
func (_class RoleClass) GetByNameLabel(sessionID SessionRef, label string) (_retval []RoleRef, _err error) {
	if (IsMock) {
		return _class.GetByNameLabel__mock(sessionID, label)
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

func (_class RoleClass) GetByUUID__mock(sessionID SessionRef, uuid string) (_retval RoleRef, _err error) {
	log.Println("Role.GetByUUID not mocked")
	_err = errors.New("Role.GetByUUID not mocked")
	return
}
// Get a reference to the role instance with the specified UUID.
func (_class RoleClass) GetByUUID(sessionID SessionRef, uuid string) (_retval RoleRef, _err error) {
	if (IsMock) {
		return _class.GetByUUID__mock(sessionID, uuid)
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

func (_class RoleClass) GetRecord__mock(sessionID SessionRef, self RoleRef) (_retval RoleRecord, _err error) {
	log.Println("Role.GetRecord not mocked")
	_err = errors.New("Role.GetRecord not mocked")
	return
}
// Get a record containing the current state of the given role.
func (_class RoleClass) GetRecord(sessionID SessionRef, self RoleRef) (_retval RoleRecord, _err error) {
	if (IsMock) {
		return _class.GetRecord__mock(sessionID, self)
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
