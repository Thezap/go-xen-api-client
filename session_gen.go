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

type SessionRecord struct {
  // Unique identifier/object reference
	UUID string
  // Currently connected host
	ThisHost HostRef
  // Currently connected user
	ThisUser UserRef
  // Timestamp for last time session was active
	LastActive time.Time
  // True if this session relates to a intra-pool login, false otherwise
	Pool bool
  // additional configuration
	OtherConfig map[string]string
  // true iff this session was created using local superuser credentials
	IsLocalSuperuser bool
  // references the subject instance that created the session. If a session instance has is_local_superuser set, then the value of this field is undefined.
	Subject SubjectRef
  // time when session was last validated
	ValidationTime time.Time
  // the subject identifier of the user that was externally authenticated. If a session instance has is_local_superuser set, then the value of this field is undefined.
	AuthUserSid string
  // the subject name of the user that was externally authenticated. If a session instance has is_local_superuser set, then the value of this field is undefined.
	AuthUserName string
  // list with all RBAC permissions for this session
	RbacPermissions []string
  // list of tasks created using the current session
	Tasks []TaskRef
  // references the parent session that created this session
	Parent SessionRef
  // a key string provided by a API user to distinguish itself from other users sharing the same login name
	Originator string
}

type SessionRef string

// A session
type SessionClass struct {
	client *Client
}


func SessionClassLogoutSubjectIdentifierMockDefault(sessionID SessionRef, subjectIdentifier string) (_err error) {
	log.Println("Session.LogoutSubjectIdentifier not mocked")
	_err = errors.New("Session.LogoutSubjectIdentifier not mocked")
	return
}

var SessionClassLogoutSubjectIdentifierMockedCallback = SessionClassLogoutSubjectIdentifierMockDefault

func (_class SessionClass) LogoutSubjectIdentifierMock(sessionID SessionRef, subjectIdentifier string) (_err error) {
	return SessionClassLogoutSubjectIdentifierMockedCallback(sessionID, subjectIdentifier)
}
// Log out all sessions associated to a user subject-identifier, except the session associated with the context calling this function
func (_class SessionClass) LogoutSubjectIdentifier(sessionID SessionRef, subjectIdentifier string) (_err error) {
	if IsMock {
		return _class.LogoutSubjectIdentifierMock(sessionID, subjectIdentifier)
	}	
	_method := "session.logout_subject_identifier"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_subjectIdentifierArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "subject_identifier"), subjectIdentifier)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _subjectIdentifierArg)
	return
}


func SessionClassGetAllSubjectIdentifiersMockDefault(sessionID SessionRef) (_retval []string, _err error) {
	log.Println("Session.GetAllSubjectIdentifiers not mocked")
	_err = errors.New("Session.GetAllSubjectIdentifiers not mocked")
	return
}

var SessionClassGetAllSubjectIdentifiersMockedCallback = SessionClassGetAllSubjectIdentifiersMockDefault

func (_class SessionClass) GetAllSubjectIdentifiersMock(sessionID SessionRef) (_retval []string, _err error) {
	return SessionClassGetAllSubjectIdentifiersMockedCallback(sessionID)
}
// Return a list of all the user subject-identifiers of all existing sessions
func (_class SessionClass) GetAllSubjectIdentifiers(sessionID SessionRef) (_retval []string, _err error) {
	if IsMock {
		return _class.GetAllSubjectIdentifiersMock(sessionID)
	}	
	_method := "session.get_all_subject_identifiers"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringSetToGo(_method + " -> ", _result.Value)
	return
}


func SessionClassLocalLogoutMockDefault(sessionID SessionRef) (_err error) {
	log.Println("Session.LocalLogout not mocked")
	_err = errors.New("Session.LocalLogout not mocked")
	return
}

var SessionClassLocalLogoutMockedCallback = SessionClassLocalLogoutMockDefault

func (_class SessionClass) LocalLogoutMock(sessionID SessionRef) (_err error) {
	return SessionClassLocalLogoutMockedCallback(sessionID)
}
// Log out of local session.
func (_class SessionClass) LocalLogout(sessionID SessionRef) (_err error) {
	if IsMock {
		return _class.LocalLogoutMock(sessionID)
	}	
	_method := "session.local_logout"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg)
	return
}


func SessionClassCreateFromDbFileMockDefault(sessionID SessionRef, filename string) (_retval SessionRef, _err error) {
	log.Println("Session.CreateFromDbFile not mocked")
	_err = errors.New("Session.CreateFromDbFile not mocked")
	return
}

var SessionClassCreateFromDbFileMockedCallback = SessionClassCreateFromDbFileMockDefault

func (_class SessionClass) CreateFromDbFileMock(sessionID SessionRef, filename string) (_retval SessionRef, _err error) {
	return SessionClassCreateFromDbFileMockedCallback(sessionID, filename)
}
// 
func (_class SessionClass) CreateFromDbFile(sessionID SessionRef, filename string) (_retval SessionRef, _err error) {
	if IsMock {
		return _class.CreateFromDbFileMock(sessionID, filename)
	}	
	_method := "session.create_from_db_file"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_filenameArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "filename"), filename)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _filenameArg)
	if _err != nil {
		return
	}
	_retval, _err = convertSessionRefToGo(_method + " -> ", _result.Value)
	return
}


func SessionClassSlaveLocalLoginWithPasswordMockDefault(uname string, pwd string) (_retval SessionRef, _err error) {
	log.Println("Session.SlaveLocalLoginWithPassword not mocked")
	_err = errors.New("Session.SlaveLocalLoginWithPassword not mocked")
	return
}

var SessionClassSlaveLocalLoginWithPasswordMockedCallback = SessionClassSlaveLocalLoginWithPasswordMockDefault

func (_class SessionClass) SlaveLocalLoginWithPasswordMock(uname string, pwd string) (_retval SessionRef, _err error) {
	return SessionClassSlaveLocalLoginWithPasswordMockedCallback(uname, pwd)
}
// Authenticate locally against a slave in emergency mode. Note the resulting sessions are only good for use on this host.
func (_class SessionClass) SlaveLocalLoginWithPassword(uname string, pwd string) (_retval SessionRef, _err error) {
	if IsMock {
		return _class.SlaveLocalLoginWithPasswordMock(uname, pwd)
	}	
	_method := "session.slave_local_login_with_password"
	_unameArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "uname"), uname)
	if _err != nil {
		return
	}
	_pwdArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "pwd"), pwd)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _unameArg, _pwdArg)
	if _err != nil {
		return
	}
	_retval, _err = convertSessionRefToGo(_method + " -> ", _result.Value)
	return
}


func SessionClassChangePasswordMockDefault(sessionID SessionRef, oldPwd string, newPwd string) (_err error) {
	log.Println("Session.ChangePassword not mocked")
	_err = errors.New("Session.ChangePassword not mocked")
	return
}

var SessionClassChangePasswordMockedCallback = SessionClassChangePasswordMockDefault

func (_class SessionClass) ChangePasswordMock(sessionID SessionRef, oldPwd string, newPwd string) (_err error) {
	return SessionClassChangePasswordMockedCallback(sessionID, oldPwd, newPwd)
}
// Change the account password; if your session is authenticated with root priviledges then the old_pwd is validated and the new_pwd is set regardless
func (_class SessionClass) ChangePassword(sessionID SessionRef, oldPwd string, newPwd string) (_err error) {
	if IsMock {
		return _class.ChangePasswordMock(sessionID, oldPwd, newPwd)
	}	
	_method := "session.change_password"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_oldPwdArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "old_pwd"), oldPwd)
	if _err != nil {
		return
	}
	_newPwdArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "new_pwd"), newPwd)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _oldPwdArg, _newPwdArg)
	return
}


func SessionClassLogoutMockDefault(sessionID SessionRef) (_err error) {
	log.Println("Session.Logout not mocked")
	_err = errors.New("Session.Logout not mocked")
	return
}

var SessionClassLogoutMockedCallback = SessionClassLogoutMockDefault

func (_class SessionClass) LogoutMock(sessionID SessionRef) (_err error) {
	return SessionClassLogoutMockedCallback(sessionID)
}
// Log out of a session
func (_class SessionClass) Logout(sessionID SessionRef) (_err error) {
	if IsMock {
		return _class.LogoutMock(sessionID)
	}	
	_method := "session.logout"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg)
	return
}


func SessionClassLoginWithPasswordMockDefault(uname string, pwd string, version string, originator string) (_retval SessionRef, _err error) {
	log.Println("Session.LoginWithPassword not mocked")
	_err = errors.New("Session.LoginWithPassword not mocked")
	return
}

var SessionClassLoginWithPasswordMockedCallback = SessionClassLoginWithPasswordMockDefault

func (_class SessionClass) LoginWithPasswordMock(uname string, pwd string, version string, originator string) (_retval SessionRef, _err error) {
	return SessionClassLoginWithPasswordMockedCallback(uname, pwd, version, originator)
}
// Attempt to authenticate the user, returning a session reference if successful
//
// Errors:
//  SESSION_AUTHENTICATION_FAILED - The credentials given by the user are incorrect, so access has been denied, and you have not been issued a session handle.
//  HOST_IS_SLAVE - You cannot make regular API calls directly on a slave. Please pass API calls via the master host.
func (_class SessionClass) LoginWithPassword(uname string, pwd string, version string, originator string) (_retval SessionRef, _err error) {
	if IsMock {
		return _class.LoginWithPasswordMock(uname, pwd, version, originator)
	}	
	_method := "session.login_with_password"
	_unameArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "uname"), uname)
	if _err != nil {
		return
	}
	_pwdArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "pwd"), pwd)
	if _err != nil {
		return
	}
	_versionArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "version"), version)
	if _err != nil {
		return
	}
	_originatorArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "originator"), originator)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _unameArg, _pwdArg, _versionArg, _originatorArg)
	if _err != nil {
		return
	}
	_retval, _err = convertSessionRefToGo(_method + " -> ", _result.Value)
	return
}


func SessionClassRemoveFromOtherConfigMockDefault(sessionID SessionRef, self SessionRef, key string) (_err error) {
	log.Println("Session.RemoveFromOtherConfig not mocked")
	_err = errors.New("Session.RemoveFromOtherConfig not mocked")
	return
}

var SessionClassRemoveFromOtherConfigMockedCallback = SessionClassRemoveFromOtherConfigMockDefault

func (_class SessionClass) RemoveFromOtherConfigMock(sessionID SessionRef, self SessionRef, key string) (_err error) {
	return SessionClassRemoveFromOtherConfigMockedCallback(sessionID, self, key)
}
// Remove the given key and its corresponding value from the other_config field of the given session.  If the key is not in that Map, then do nothing.
func (_class SessionClass) RemoveFromOtherConfig(sessionID SessionRef, self SessionRef, key string) (_err error) {
	if IsMock {
		return _class.RemoveFromOtherConfigMock(sessionID, self, key)
	}	
	_method := "session.remove_from_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func SessionClassAddToOtherConfigMockDefault(sessionID SessionRef, self SessionRef, key string, value string) (_err error) {
	log.Println("Session.AddToOtherConfig not mocked")
	_err = errors.New("Session.AddToOtherConfig not mocked")
	return
}

var SessionClassAddToOtherConfigMockedCallback = SessionClassAddToOtherConfigMockDefault

func (_class SessionClass) AddToOtherConfigMock(sessionID SessionRef, self SessionRef, key string, value string) (_err error) {
	return SessionClassAddToOtherConfigMockedCallback(sessionID, self, key, value)
}
// Add the given key-value pair to the other_config field of the given session.
func (_class SessionClass) AddToOtherConfig(sessionID SessionRef, self SessionRef, key string, value string) (_err error) {
	if IsMock {
		return _class.AddToOtherConfigMock(sessionID, self, key, value)
	}	
	_method := "session.add_to_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func SessionClassSetOtherConfigMockDefault(sessionID SessionRef, self SessionRef, value map[string]string) (_err error) {
	log.Println("Session.SetOtherConfig not mocked")
	_err = errors.New("Session.SetOtherConfig not mocked")
	return
}

var SessionClassSetOtherConfigMockedCallback = SessionClassSetOtherConfigMockDefault

func (_class SessionClass) SetOtherConfigMock(sessionID SessionRef, self SessionRef, value map[string]string) (_err error) {
	return SessionClassSetOtherConfigMockedCallback(sessionID, self, value)
}
// Set the other_config field of the given session.
func (_class SessionClass) SetOtherConfig(sessionID SessionRef, self SessionRef, value map[string]string) (_err error) {
	if IsMock {
		return _class.SetOtherConfigMock(sessionID, self, value)
	}	
	_method := "session.set_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func SessionClassGetOriginatorMockDefault(sessionID SessionRef, self SessionRef) (_retval string, _err error) {
	log.Println("Session.GetOriginator not mocked")
	_err = errors.New("Session.GetOriginator not mocked")
	return
}

var SessionClassGetOriginatorMockedCallback = SessionClassGetOriginatorMockDefault

func (_class SessionClass) GetOriginatorMock(sessionID SessionRef, self SessionRef) (_retval string, _err error) {
	return SessionClassGetOriginatorMockedCallback(sessionID, self)
}
// Get the originator field of the given session.
func (_class SessionClass) GetOriginator(sessionID SessionRef, self SessionRef) (_retval string, _err error) {
	if IsMock {
		return _class.GetOriginatorMock(sessionID, self)
	}	
	_method := "session.get_originator"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func SessionClassGetParentMockDefault(sessionID SessionRef, self SessionRef) (_retval SessionRef, _err error) {
	log.Println("Session.GetParent not mocked")
	_err = errors.New("Session.GetParent not mocked")
	return
}

var SessionClassGetParentMockedCallback = SessionClassGetParentMockDefault

func (_class SessionClass) GetParentMock(sessionID SessionRef, self SessionRef) (_retval SessionRef, _err error) {
	return SessionClassGetParentMockedCallback(sessionID, self)
}
// Get the parent field of the given session.
func (_class SessionClass) GetParent(sessionID SessionRef, self SessionRef) (_retval SessionRef, _err error) {
	if IsMock {
		return _class.GetParentMock(sessionID, self)
	}	
	_method := "session.get_parent"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertSessionRefToGo(_method + " -> ", _result.Value)
	return
}


func SessionClassGetTasksMockDefault(sessionID SessionRef, self SessionRef) (_retval []TaskRef, _err error) {
	log.Println("Session.GetTasks not mocked")
	_err = errors.New("Session.GetTasks not mocked")
	return
}

var SessionClassGetTasksMockedCallback = SessionClassGetTasksMockDefault

func (_class SessionClass) GetTasksMock(sessionID SessionRef, self SessionRef) (_retval []TaskRef, _err error) {
	return SessionClassGetTasksMockedCallback(sessionID, self)
}
// Get the tasks field of the given session.
func (_class SessionClass) GetTasks(sessionID SessionRef, self SessionRef) (_retval []TaskRef, _err error) {
	if IsMock {
		return _class.GetTasksMock(sessionID, self)
	}	
	_method := "session.get_tasks"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefSetToGo(_method + " -> ", _result.Value)
	return
}


func SessionClassGetRbacPermissionsMockDefault(sessionID SessionRef, self SessionRef) (_retval []string, _err error) {
	log.Println("Session.GetRbacPermissions not mocked")
	_err = errors.New("Session.GetRbacPermissions not mocked")
	return
}

var SessionClassGetRbacPermissionsMockedCallback = SessionClassGetRbacPermissionsMockDefault

func (_class SessionClass) GetRbacPermissionsMock(sessionID SessionRef, self SessionRef) (_retval []string, _err error) {
	return SessionClassGetRbacPermissionsMockedCallback(sessionID, self)
}
// Get the rbac_permissions field of the given session.
func (_class SessionClass) GetRbacPermissions(sessionID SessionRef, self SessionRef) (_retval []string, _err error) {
	if IsMock {
		return _class.GetRbacPermissionsMock(sessionID, self)
	}	
	_method := "session.get_rbac_permissions"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func SessionClassGetAuthUserNameMockDefault(sessionID SessionRef, self SessionRef) (_retval string, _err error) {
	log.Println("Session.GetAuthUserName not mocked")
	_err = errors.New("Session.GetAuthUserName not mocked")
	return
}

var SessionClassGetAuthUserNameMockedCallback = SessionClassGetAuthUserNameMockDefault

func (_class SessionClass) GetAuthUserNameMock(sessionID SessionRef, self SessionRef) (_retval string, _err error) {
	return SessionClassGetAuthUserNameMockedCallback(sessionID, self)
}
// Get the auth_user_name field of the given session.
func (_class SessionClass) GetAuthUserName(sessionID SessionRef, self SessionRef) (_retval string, _err error) {
	if IsMock {
		return _class.GetAuthUserNameMock(sessionID, self)
	}	
	_method := "session.get_auth_user_name"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func SessionClassGetAuthUserSidMockDefault(sessionID SessionRef, self SessionRef) (_retval string, _err error) {
	log.Println("Session.GetAuthUserSid not mocked")
	_err = errors.New("Session.GetAuthUserSid not mocked")
	return
}

var SessionClassGetAuthUserSidMockedCallback = SessionClassGetAuthUserSidMockDefault

func (_class SessionClass) GetAuthUserSidMock(sessionID SessionRef, self SessionRef) (_retval string, _err error) {
	return SessionClassGetAuthUserSidMockedCallback(sessionID, self)
}
// Get the auth_user_sid field of the given session.
func (_class SessionClass) GetAuthUserSid(sessionID SessionRef, self SessionRef) (_retval string, _err error) {
	if IsMock {
		return _class.GetAuthUserSidMock(sessionID, self)
	}	
	_method := "session.get_auth_user_sid"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func SessionClassGetValidationTimeMockDefault(sessionID SessionRef, self SessionRef) (_retval time.Time, _err error) {
	log.Println("Session.GetValidationTime not mocked")
	_err = errors.New("Session.GetValidationTime not mocked")
	return
}

var SessionClassGetValidationTimeMockedCallback = SessionClassGetValidationTimeMockDefault

func (_class SessionClass) GetValidationTimeMock(sessionID SessionRef, self SessionRef) (_retval time.Time, _err error) {
	return SessionClassGetValidationTimeMockedCallback(sessionID, self)
}
// Get the validation_time field of the given session.
func (_class SessionClass) GetValidationTime(sessionID SessionRef, self SessionRef) (_retval time.Time, _err error) {
	if IsMock {
		return _class.GetValidationTimeMock(sessionID, self)
	}	
	_method := "session.get_validation_time"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTimeToGo(_method + " -> ", _result.Value)
	return
}


func SessionClassGetSubjectMockDefault(sessionID SessionRef, self SessionRef) (_retval SubjectRef, _err error) {
	log.Println("Session.GetSubject not mocked")
	_err = errors.New("Session.GetSubject not mocked")
	return
}

var SessionClassGetSubjectMockedCallback = SessionClassGetSubjectMockDefault

func (_class SessionClass) GetSubjectMock(sessionID SessionRef, self SessionRef) (_retval SubjectRef, _err error) {
	return SessionClassGetSubjectMockedCallback(sessionID, self)
}
// Get the subject field of the given session.
func (_class SessionClass) GetSubject(sessionID SessionRef, self SessionRef) (_retval SubjectRef, _err error) {
	if IsMock {
		return _class.GetSubjectMock(sessionID, self)
	}	
	_method := "session.get_subject"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertSubjectRefToGo(_method + " -> ", _result.Value)
	return
}


func SessionClassGetIsLocalSuperuserMockDefault(sessionID SessionRef, self SessionRef) (_retval bool, _err error) {
	log.Println("Session.GetIsLocalSuperuser not mocked")
	_err = errors.New("Session.GetIsLocalSuperuser not mocked")
	return
}

var SessionClassGetIsLocalSuperuserMockedCallback = SessionClassGetIsLocalSuperuserMockDefault

func (_class SessionClass) GetIsLocalSuperuserMock(sessionID SessionRef, self SessionRef) (_retval bool, _err error) {
	return SessionClassGetIsLocalSuperuserMockedCallback(sessionID, self)
}
// Get the is_local_superuser field of the given session.
func (_class SessionClass) GetIsLocalSuperuser(sessionID SessionRef, self SessionRef) (_retval bool, _err error) {
	if IsMock {
		return _class.GetIsLocalSuperuserMock(sessionID, self)
	}	
	_method := "session.get_is_local_superuser"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertBoolToGo(_method + " -> ", _result.Value)
	return
}


func SessionClassGetOtherConfigMockDefault(sessionID SessionRef, self SessionRef) (_retval map[string]string, _err error) {
	log.Println("Session.GetOtherConfig not mocked")
	_err = errors.New("Session.GetOtherConfig not mocked")
	return
}

var SessionClassGetOtherConfigMockedCallback = SessionClassGetOtherConfigMockDefault

func (_class SessionClass) GetOtherConfigMock(sessionID SessionRef, self SessionRef) (_retval map[string]string, _err error) {
	return SessionClassGetOtherConfigMockedCallback(sessionID, self)
}
// Get the other_config field of the given session.
func (_class SessionClass) GetOtherConfig(sessionID SessionRef, self SessionRef) (_retval map[string]string, _err error) {
	if IsMock {
		return _class.GetOtherConfigMock(sessionID, self)
	}	
	_method := "session.get_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func SessionClassGetPoolMockDefault(sessionID SessionRef, self SessionRef) (_retval bool, _err error) {
	log.Println("Session.GetPool not mocked")
	_err = errors.New("Session.GetPool not mocked")
	return
}

var SessionClassGetPoolMockedCallback = SessionClassGetPoolMockDefault

func (_class SessionClass) GetPoolMock(sessionID SessionRef, self SessionRef) (_retval bool, _err error) {
	return SessionClassGetPoolMockedCallback(sessionID, self)
}
// Get the pool field of the given session.
func (_class SessionClass) GetPool(sessionID SessionRef, self SessionRef) (_retval bool, _err error) {
	if IsMock {
		return _class.GetPoolMock(sessionID, self)
	}	
	_method := "session.get_pool"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertBoolToGo(_method + " -> ", _result.Value)
	return
}


func SessionClassGetLastActiveMockDefault(sessionID SessionRef, self SessionRef) (_retval time.Time, _err error) {
	log.Println("Session.GetLastActive not mocked")
	_err = errors.New("Session.GetLastActive not mocked")
	return
}

var SessionClassGetLastActiveMockedCallback = SessionClassGetLastActiveMockDefault

func (_class SessionClass) GetLastActiveMock(sessionID SessionRef, self SessionRef) (_retval time.Time, _err error) {
	return SessionClassGetLastActiveMockedCallback(sessionID, self)
}
// Get the last_active field of the given session.
func (_class SessionClass) GetLastActive(sessionID SessionRef, self SessionRef) (_retval time.Time, _err error) {
	if IsMock {
		return _class.GetLastActiveMock(sessionID, self)
	}	
	_method := "session.get_last_active"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTimeToGo(_method + " -> ", _result.Value)
	return
}


func SessionClassGetThisUserMockDefault(sessionID SessionRef, self SessionRef) (_retval UserRef, _err error) {
	log.Println("Session.GetThisUser not mocked")
	_err = errors.New("Session.GetThisUser not mocked")
	return
}

var SessionClassGetThisUserMockedCallback = SessionClassGetThisUserMockDefault

func (_class SessionClass) GetThisUserMock(sessionID SessionRef, self SessionRef) (_retval UserRef, _err error) {
	return SessionClassGetThisUserMockedCallback(sessionID, self)
}
// Get the this_user field of the given session.
func (_class SessionClass) GetThisUser(sessionID SessionRef, self SessionRef) (_retval UserRef, _err error) {
	if IsMock {
		return _class.GetThisUserMock(sessionID, self)
	}	
	_method := "session.get_this_user"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertUserRefToGo(_method + " -> ", _result.Value)
	return
}


func SessionClassGetThisHostMockDefault(sessionID SessionRef, self SessionRef) (_retval HostRef, _err error) {
	log.Println("Session.GetThisHost not mocked")
	_err = errors.New("Session.GetThisHost not mocked")
	return
}

var SessionClassGetThisHostMockedCallback = SessionClassGetThisHostMockDefault

func (_class SessionClass) GetThisHostMock(sessionID SessionRef, self SessionRef) (_retval HostRef, _err error) {
	return SessionClassGetThisHostMockedCallback(sessionID, self)
}
// Get the this_host field of the given session.
func (_class SessionClass) GetThisHost(sessionID SessionRef, self SessionRef) (_retval HostRef, _err error) {
	if IsMock {
		return _class.GetThisHostMock(sessionID, self)
	}	
	_method := "session.get_this_host"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func SessionClassGetUUIDMockDefault(sessionID SessionRef, self SessionRef) (_retval string, _err error) {
	log.Println("Session.GetUUID not mocked")
	_err = errors.New("Session.GetUUID not mocked")
	return
}

var SessionClassGetUUIDMockedCallback = SessionClassGetUUIDMockDefault

func (_class SessionClass) GetUUIDMock(sessionID SessionRef, self SessionRef) (_retval string, _err error) {
	return SessionClassGetUUIDMockedCallback(sessionID, self)
}
// Get the uuid field of the given session.
func (_class SessionClass) GetUUID(sessionID SessionRef, self SessionRef) (_retval string, _err error) {
	if IsMock {
		return _class.GetUUIDMock(sessionID, self)
	}	
	_method := "session.get_uuid"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func SessionClassGetByUUIDMockDefault(sessionID SessionRef, uuid string) (_retval SessionRef, _err error) {
	log.Println("Session.GetByUUID not mocked")
	_err = errors.New("Session.GetByUUID not mocked")
	return
}

var SessionClassGetByUUIDMockedCallback = SessionClassGetByUUIDMockDefault

func (_class SessionClass) GetByUUIDMock(sessionID SessionRef, uuid string) (_retval SessionRef, _err error) {
	return SessionClassGetByUUIDMockedCallback(sessionID, uuid)
}
// Get a reference to the session instance with the specified UUID.
func (_class SessionClass) GetByUUID(sessionID SessionRef, uuid string) (_retval SessionRef, _err error) {
	if IsMock {
		return _class.GetByUUIDMock(sessionID, uuid)
	}	
	_method := "session.get_by_uuid"
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
	_retval, _err = convertSessionRefToGo(_method + " -> ", _result.Value)
	return
}


func SessionClassGetRecordMockDefault(sessionID SessionRef, self SessionRef) (_retval SessionRecord, _err error) {
	log.Println("Session.GetRecord not mocked")
	_err = errors.New("Session.GetRecord not mocked")
	return
}

var SessionClassGetRecordMockedCallback = SessionClassGetRecordMockDefault

func (_class SessionClass) GetRecordMock(sessionID SessionRef, self SessionRef) (_retval SessionRecord, _err error) {
	return SessionClassGetRecordMockedCallback(sessionID, self)
}
// Get a record containing the current state of the given session.
func (_class SessionClass) GetRecord(sessionID SessionRef, self SessionRef) (_retval SessionRecord, _err error) {
	if IsMock {
		return _class.GetRecordMock(sessionID, self)
	}	
	_method := "session.get_record"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertSessionRecordToGo(_method + " -> ", _result.Value)
	return
}
