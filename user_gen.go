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

type UserRecord struct {
  // Unique identifier/object reference
	UUID string
  // short name (e.g. userid)
	ShortName string
  // full name
	Fullname string
  // additional configuration
	OtherConfig map[string]string
}

type UserRef string

// A user of the system
type UserClass struct {
	client *Client
}

func (_class UserClass) RemoveFromOtherConfig__mock(sessionID SessionRef, self UserRef, key string) (_err error) {
	log.Println("User.RemoveFromOtherConfig not mocked")
	_err = errors.New("User.RemoveFromOtherConfig not mocked")
	return
}
// Remove the given key and its corresponding value from the other_config field of the given user.  If the key is not in that Map, then do nothing.
func (_class UserClass) RemoveFromOtherConfig(sessionID SessionRef, self UserRef, key string) (_err error) {
	if (IsMock) {
		return _class.RemoveFromOtherConfig__mock(sessionID, self, key)
	}	
	_method := "user.remove_from_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertUserRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class UserClass) AddToOtherConfig__mock(sessionID SessionRef, self UserRef, key string, value string) (_err error) {
	log.Println("User.AddToOtherConfig not mocked")
	_err = errors.New("User.AddToOtherConfig not mocked")
	return
}
// Add the given key-value pair to the other_config field of the given user.
func (_class UserClass) AddToOtherConfig(sessionID SessionRef, self UserRef, key string, value string) (_err error) {
	if (IsMock) {
		return _class.AddToOtherConfig__mock(sessionID, self, key, value)
	}	
	_method := "user.add_to_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertUserRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class UserClass) SetOtherConfig__mock(sessionID SessionRef, self UserRef, value map[string]string) (_err error) {
	log.Println("User.SetOtherConfig not mocked")
	_err = errors.New("User.SetOtherConfig not mocked")
	return
}
// Set the other_config field of the given user.
func (_class UserClass) SetOtherConfig(sessionID SessionRef, self UserRef, value map[string]string) (_err error) {
	if (IsMock) {
		return _class.SetOtherConfig__mock(sessionID, self, value)
	}	
	_method := "user.set_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertUserRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class UserClass) SetFullname__mock(sessionID SessionRef, self UserRef, value string) (_err error) {
	log.Println("User.SetFullname not mocked")
	_err = errors.New("User.SetFullname not mocked")
	return
}
// Set the fullname field of the given user.
func (_class UserClass) SetFullname(sessionID SessionRef, self UserRef, value string) (_err error) {
	if (IsMock) {
		return _class.SetFullname__mock(sessionID, self, value)
	}	
	_method := "user.set_fullname"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertUserRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}

func (_class UserClass) GetOtherConfig__mock(sessionID SessionRef, self UserRef) (_retval map[string]string, _err error) {
	log.Println("User.GetOtherConfig not mocked")
	_err = errors.New("User.GetOtherConfig not mocked")
	return
}
// Get the other_config field of the given user.
func (_class UserClass) GetOtherConfig(sessionID SessionRef, self UserRef) (_retval map[string]string, _err error) {
	if (IsMock) {
		return _class.GetOtherConfig__mock(sessionID, self)
	}	
	_method := "user.get_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertUserRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class UserClass) GetFullname__mock(sessionID SessionRef, self UserRef) (_retval string, _err error) {
	log.Println("User.GetFullname not mocked")
	_err = errors.New("User.GetFullname not mocked")
	return
}
// Get the fullname field of the given user.
func (_class UserClass) GetFullname(sessionID SessionRef, self UserRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetFullname__mock(sessionID, self)
	}	
	_method := "user.get_fullname"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertUserRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class UserClass) GetShortName__mock(sessionID SessionRef, self UserRef) (_retval string, _err error) {
	log.Println("User.GetShortName not mocked")
	_err = errors.New("User.GetShortName not mocked")
	return
}
// Get the short_name field of the given user.
func (_class UserClass) GetShortName(sessionID SessionRef, self UserRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetShortName__mock(sessionID, self)
	}	
	_method := "user.get_short_name"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertUserRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class UserClass) GetUUID__mock(sessionID SessionRef, self UserRef) (_retval string, _err error) {
	log.Println("User.GetUUID not mocked")
	_err = errors.New("User.GetUUID not mocked")
	return
}
// Get the uuid field of the given user.
func (_class UserClass) GetUUID(sessionID SessionRef, self UserRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetUUID__mock(sessionID, self)
	}	
	_method := "user.get_uuid"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertUserRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class UserClass) Destroy__mock(sessionID SessionRef, self UserRef) (_err error) {
	log.Println("User.Destroy not mocked")
	_err = errors.New("User.Destroy not mocked")
	return
}
// Destroy the specified user instance.
func (_class UserClass) Destroy(sessionID SessionRef, self UserRef) (_err error) {
	if (IsMock) {
		return _class.Destroy__mock(sessionID, self)
	}	
	_method := "user.destroy"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertUserRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg)
	return
}

func (_class UserClass) Create__mock(sessionID SessionRef, args UserRecord) (_retval UserRef, _err error) {
	log.Println("User.Create not mocked")
	_err = errors.New("User.Create not mocked")
	return
}
// Create a new user instance, and return its handle.
// The constructor args are: short_name*, fullname*, other_config (* = non-optional).
func (_class UserClass) Create(sessionID SessionRef, args UserRecord) (_retval UserRef, _err error) {
	if (IsMock) {
		return _class.Create__mock(sessionID, args)
	}	
	_method := "user.create"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_argsArg, _err := convertUserRecordToXen(fmt.Sprintf("%s(%s)", _method, "args"), args)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _argsArg)
	if _err != nil {
		return
	}
	_retval, _err = convertUserRefToGo(_method + " -> ", _result.Value)
	return
}

func (_class UserClass) GetByUUID__mock(sessionID SessionRef, uuid string) (_retval UserRef, _err error) {
	log.Println("User.GetByUUID not mocked")
	_err = errors.New("User.GetByUUID not mocked")
	return
}
// Get a reference to the user instance with the specified UUID.
func (_class UserClass) GetByUUID(sessionID SessionRef, uuid string) (_retval UserRef, _err error) {
	if (IsMock) {
		return _class.GetByUUID__mock(sessionID, uuid)
	}	
	_method := "user.get_by_uuid"
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
	_retval, _err = convertUserRefToGo(_method + " -> ", _result.Value)
	return
}

func (_class UserClass) GetRecord__mock(sessionID SessionRef, self UserRef) (_retval UserRecord, _err error) {
	log.Println("User.GetRecord not mocked")
	_err = errors.New("User.GetRecord not mocked")
	return
}
// Get a record containing the current state of the given user.
func (_class UserClass) GetRecord(sessionID SessionRef, self UserRef) (_retval UserRecord, _err error) {
	if (IsMock) {
		return _class.GetRecord__mock(sessionID, self)
	}	
	_method := "user.get_record"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertUserRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertUserRecordToGo(_method + " -> ", _result.Value)
	return
}
