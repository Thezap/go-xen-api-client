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

type SecretRecord struct {
  // Unique identifier/object reference
	UUID string
  // the secret
	Value string
  // other_config
	OtherConfig map[string]string
}

type SecretRef string

// A secret
type SecretClass struct {
	client *Client
}

func (_class SecretClass) GetAllRecords__mock(sessionID SessionRef) (_retval map[SecretRef]SecretRecord, _err error) {
	log.Println("Secret.GetAllRecords not mocked")
	_err = errors.New("Secret.GetAllRecords not mocked")
	return
}
// Return a map of secret references to secret records for all secrets known to the system.
func (_class SecretClass) GetAllRecords(sessionID SessionRef) (_retval map[SecretRef]SecretRecord, _err error) {
	if (IsMock) {
		return _class.GetAllRecords__mock(sessionID)
	}	
	_method := "secret.get_all_records"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertSecretRefToSecretRecordMapToGo(_method + " -> ", _result.Value)
	return
}

func (_class SecretClass) GetAll__mock(sessionID SessionRef) (_retval []SecretRef, _err error) {
	log.Println("Secret.GetAll not mocked")
	_err = errors.New("Secret.GetAll not mocked")
	return
}
// Return a list of all the secrets known to the system.
func (_class SecretClass) GetAll(sessionID SessionRef) (_retval []SecretRef, _err error) {
	if (IsMock) {
		return _class.GetAll__mock(sessionID)
	}	
	_method := "secret.get_all"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertSecretRefSetToGo(_method + " -> ", _result.Value)
	return
}

func (_class SecretClass) RemoveFromOtherConfig__mock(sessionID SessionRef, self SecretRef, key string) (_err error) {
	log.Println("Secret.RemoveFromOtherConfig not mocked")
	_err = errors.New("Secret.RemoveFromOtherConfig not mocked")
	return
}
// Remove the given key and its corresponding value from the other_config field of the given secret.  If the key is not in that Map, then do nothing.
func (_class SecretClass) RemoveFromOtherConfig(sessionID SessionRef, self SecretRef, key string) (_err error) {
	if (IsMock) {
		return _class.RemoveFromOtherConfig__mock(sessionID, self, key)
	}	
	_method := "secret.remove_from_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertSecretRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class SecretClass) AddToOtherConfig__mock(sessionID SessionRef, self SecretRef, key string, value string) (_err error) {
	log.Println("Secret.AddToOtherConfig not mocked")
	_err = errors.New("Secret.AddToOtherConfig not mocked")
	return
}
// Add the given key-value pair to the other_config field of the given secret.
func (_class SecretClass) AddToOtherConfig(sessionID SessionRef, self SecretRef, key string, value string) (_err error) {
	if (IsMock) {
		return _class.AddToOtherConfig__mock(sessionID, self, key, value)
	}	
	_method := "secret.add_to_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertSecretRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class SecretClass) SetOtherConfig__mock(sessionID SessionRef, self SecretRef, value map[string]string) (_err error) {
	log.Println("Secret.SetOtherConfig not mocked")
	_err = errors.New("Secret.SetOtherConfig not mocked")
	return
}
// Set the other_config field of the given secret.
func (_class SecretClass) SetOtherConfig(sessionID SessionRef, self SecretRef, value map[string]string) (_err error) {
	if (IsMock) {
		return _class.SetOtherConfig__mock(sessionID, self, value)
	}	
	_method := "secret.set_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertSecretRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class SecretClass) SetValue__mock(sessionID SessionRef, self SecretRef, value string) (_err error) {
	log.Println("Secret.SetValue not mocked")
	_err = errors.New("Secret.SetValue not mocked")
	return
}
// Set the value field of the given secret.
func (_class SecretClass) SetValue(sessionID SessionRef, self SecretRef, value string) (_err error) {
	if (IsMock) {
		return _class.SetValue__mock(sessionID, self, value)
	}	
	_method := "secret.set_value"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertSecretRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class SecretClass) GetOtherConfig__mock(sessionID SessionRef, self SecretRef) (_retval map[string]string, _err error) {
	log.Println("Secret.GetOtherConfig not mocked")
	_err = errors.New("Secret.GetOtherConfig not mocked")
	return
}
// Get the other_config field of the given secret.
func (_class SecretClass) GetOtherConfig(sessionID SessionRef, self SecretRef) (_retval map[string]string, _err error) {
	if (IsMock) {
		return _class.GetOtherConfig__mock(sessionID, self)
	}	
	_method := "secret.get_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertSecretRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class SecretClass) GetValue__mock(sessionID SessionRef, self SecretRef) (_retval string, _err error) {
	log.Println("Secret.GetValue not mocked")
	_err = errors.New("Secret.GetValue not mocked")
	return
}
// Get the value field of the given secret.
func (_class SecretClass) GetValue(sessionID SessionRef, self SecretRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetValue__mock(sessionID, self)
	}	
	_method := "secret.get_value"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertSecretRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class SecretClass) GetUUID__mock(sessionID SessionRef, self SecretRef) (_retval string, _err error) {
	log.Println("Secret.GetUUID not mocked")
	_err = errors.New("Secret.GetUUID not mocked")
	return
}
// Get the uuid field of the given secret.
func (_class SecretClass) GetUUID(sessionID SessionRef, self SecretRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetUUID__mock(sessionID, self)
	}	
	_method := "secret.get_uuid"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertSecretRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class SecretClass) Destroy__mock(sessionID SessionRef, self SecretRef) (_err error) {
	log.Println("Secret.Destroy not mocked")
	_err = errors.New("Secret.Destroy not mocked")
	return
}
// Destroy the specified secret instance.
func (_class SecretClass) Destroy(sessionID SessionRef, self SecretRef) (_err error) {
	if (IsMock) {
		return _class.Destroy__mock(sessionID, self)
	}	
	_method := "secret.destroy"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertSecretRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg)
	return
}

func (_class SecretClass) Create__mock(sessionID SessionRef, args SecretRecord) (_retval SecretRef, _err error) {
	log.Println("Secret.Create not mocked")
	_err = errors.New("Secret.Create not mocked")
	return
}
// Create a new secret instance, and return its handle.
// The constructor args are: value*, other_config (* = non-optional).
func (_class SecretClass) Create(sessionID SessionRef, args SecretRecord) (_retval SecretRef, _err error) {
	if (IsMock) {
		return _class.Create__mock(sessionID, args)
	}	
	_method := "secret.create"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_argsArg, _err := convertSecretRecordToXen(fmt.Sprintf("%s(%s)", _method, "args"), args)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _argsArg)
	if _err != nil {
		return
	}
	_retval, _err = convertSecretRefToGo(_method + " -> ", _result.Value)
	return
}

func (_class SecretClass) GetByUUID__mock(sessionID SessionRef, uuid string) (_retval SecretRef, _err error) {
	log.Println("Secret.GetByUUID not mocked")
	_err = errors.New("Secret.GetByUUID not mocked")
	return
}
// Get a reference to the secret instance with the specified UUID.
func (_class SecretClass) GetByUUID(sessionID SessionRef, uuid string) (_retval SecretRef, _err error) {
	if (IsMock) {
		return _class.GetByUUID__mock(sessionID, uuid)
	}	
	_method := "secret.get_by_uuid"
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
	_retval, _err = convertSecretRefToGo(_method + " -> ", _result.Value)
	return
}

func (_class SecretClass) GetRecord__mock(sessionID SessionRef, self SecretRef) (_retval SecretRecord, _err error) {
	log.Println("Secret.GetRecord not mocked")
	_err = errors.New("Secret.GetRecord not mocked")
	return
}
// Get a record containing the current state of the given secret.
func (_class SecretClass) GetRecord(sessionID SessionRef, self SecretRef) (_retval SecretRecord, _err error) {
	if (IsMock) {
		return _class.GetRecord__mock(sessionID, self)
	}	
	_method := "secret.get_record"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertSecretRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertSecretRecordToGo(_method + " -> ", _result.Value)
	return
}
