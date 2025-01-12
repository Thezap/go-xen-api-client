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

type FeatureRecord struct {
  // Unique identifier/object reference
	UUID string
  // a human-readable name
	NameLabel string
  // a notes field containing human-readable description
	NameDescription string
  // Indicates whether the feature is enabled
	Enabled bool
  // Indicates whether the feature is experimental (as opposed to stable and fully supported)
	Experimental bool
  // The version of this feature
	Version string
  // The host where this feature is available
	Host HostRef
}

type FeatureRef string

// A new piece of functionality
type FeatureClass struct {
	client *Client
}


func FeatureClassGetAllRecordsMockDefault(sessionID SessionRef) (_retval map[FeatureRef]FeatureRecord, _err error) {
	log.Println("Feature.GetAllRecords not mocked")
	_err = errors.New("Feature.GetAllRecords not mocked")
	return
}

var FeatureClassGetAllRecordsMockedCallback = FeatureClassGetAllRecordsMockDefault

func (_class FeatureClass) GetAllRecordsMock(sessionID SessionRef) (_retval map[FeatureRef]FeatureRecord, _err error) {
	return FeatureClassGetAllRecordsMockedCallback(sessionID)
}
// Return a map of Feature references to Feature records for all Features known to the system.
func (_class FeatureClass) GetAllRecords(sessionID SessionRef) (_retval map[FeatureRef]FeatureRecord, _err error) {
	if IsMock {
		return _class.GetAllRecordsMock(sessionID)
	}	
	_method := "Feature.get_all_records"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertFeatureRefToFeatureRecordMapToGo(_method + " -> ", _result.Value)
	return
}


func FeatureClassGetAllMockDefault(sessionID SessionRef) (_retval []FeatureRef, _err error) {
	log.Println("Feature.GetAll not mocked")
	_err = errors.New("Feature.GetAll not mocked")
	return
}

var FeatureClassGetAllMockedCallback = FeatureClassGetAllMockDefault

func (_class FeatureClass) GetAllMock(sessionID SessionRef) (_retval []FeatureRef, _err error) {
	return FeatureClassGetAllMockedCallback(sessionID)
}
// Return a list of all the Features known to the system.
func (_class FeatureClass) GetAll(sessionID SessionRef) (_retval []FeatureRef, _err error) {
	if IsMock {
		return _class.GetAllMock(sessionID)
	}	
	_method := "Feature.get_all"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertFeatureRefSetToGo(_method + " -> ", _result.Value)
	return
}


func FeatureClassGetHostMockDefault(sessionID SessionRef, self FeatureRef) (_retval HostRef, _err error) {
	log.Println("Feature.GetHost not mocked")
	_err = errors.New("Feature.GetHost not mocked")
	return
}

var FeatureClassGetHostMockedCallback = FeatureClassGetHostMockDefault

func (_class FeatureClass) GetHostMock(sessionID SessionRef, self FeatureRef) (_retval HostRef, _err error) {
	return FeatureClassGetHostMockedCallback(sessionID, self)
}
// Get the host field of the given Feature.
func (_class FeatureClass) GetHost(sessionID SessionRef, self FeatureRef) (_retval HostRef, _err error) {
	if IsMock {
		return _class.GetHostMock(sessionID, self)
	}	
	_method := "Feature.get_host"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertFeatureRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func FeatureClassGetVersionMockDefault(sessionID SessionRef, self FeatureRef) (_retval string, _err error) {
	log.Println("Feature.GetVersion not mocked")
	_err = errors.New("Feature.GetVersion not mocked")
	return
}

var FeatureClassGetVersionMockedCallback = FeatureClassGetVersionMockDefault

func (_class FeatureClass) GetVersionMock(sessionID SessionRef, self FeatureRef) (_retval string, _err error) {
	return FeatureClassGetVersionMockedCallback(sessionID, self)
}
// Get the version field of the given Feature.
func (_class FeatureClass) GetVersion(sessionID SessionRef, self FeatureRef) (_retval string, _err error) {
	if IsMock {
		return _class.GetVersionMock(sessionID, self)
	}	
	_method := "Feature.get_version"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertFeatureRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func FeatureClassGetExperimentalMockDefault(sessionID SessionRef, self FeatureRef) (_retval bool, _err error) {
	log.Println("Feature.GetExperimental not mocked")
	_err = errors.New("Feature.GetExperimental not mocked")
	return
}

var FeatureClassGetExperimentalMockedCallback = FeatureClassGetExperimentalMockDefault

func (_class FeatureClass) GetExperimentalMock(sessionID SessionRef, self FeatureRef) (_retval bool, _err error) {
	return FeatureClassGetExperimentalMockedCallback(sessionID, self)
}
// Get the experimental field of the given Feature.
func (_class FeatureClass) GetExperimental(sessionID SessionRef, self FeatureRef) (_retval bool, _err error) {
	if IsMock {
		return _class.GetExperimentalMock(sessionID, self)
	}	
	_method := "Feature.get_experimental"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertFeatureRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func FeatureClassGetEnabledMockDefault(sessionID SessionRef, self FeatureRef) (_retval bool, _err error) {
	log.Println("Feature.GetEnabled not mocked")
	_err = errors.New("Feature.GetEnabled not mocked")
	return
}

var FeatureClassGetEnabledMockedCallback = FeatureClassGetEnabledMockDefault

func (_class FeatureClass) GetEnabledMock(sessionID SessionRef, self FeatureRef) (_retval bool, _err error) {
	return FeatureClassGetEnabledMockedCallback(sessionID, self)
}
// Get the enabled field of the given Feature.
func (_class FeatureClass) GetEnabled(sessionID SessionRef, self FeatureRef) (_retval bool, _err error) {
	if IsMock {
		return _class.GetEnabledMock(sessionID, self)
	}	
	_method := "Feature.get_enabled"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertFeatureRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func FeatureClassGetNameDescriptionMockDefault(sessionID SessionRef, self FeatureRef) (_retval string, _err error) {
	log.Println("Feature.GetNameDescription not mocked")
	_err = errors.New("Feature.GetNameDescription not mocked")
	return
}

var FeatureClassGetNameDescriptionMockedCallback = FeatureClassGetNameDescriptionMockDefault

func (_class FeatureClass) GetNameDescriptionMock(sessionID SessionRef, self FeatureRef) (_retval string, _err error) {
	return FeatureClassGetNameDescriptionMockedCallback(sessionID, self)
}
// Get the name/description field of the given Feature.
func (_class FeatureClass) GetNameDescription(sessionID SessionRef, self FeatureRef) (_retval string, _err error) {
	if IsMock {
		return _class.GetNameDescriptionMock(sessionID, self)
	}	
	_method := "Feature.get_name_description"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertFeatureRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func FeatureClassGetNameLabelMockDefault(sessionID SessionRef, self FeatureRef) (_retval string, _err error) {
	log.Println("Feature.GetNameLabel not mocked")
	_err = errors.New("Feature.GetNameLabel not mocked")
	return
}

var FeatureClassGetNameLabelMockedCallback = FeatureClassGetNameLabelMockDefault

func (_class FeatureClass) GetNameLabelMock(sessionID SessionRef, self FeatureRef) (_retval string, _err error) {
	return FeatureClassGetNameLabelMockedCallback(sessionID, self)
}
// Get the name/label field of the given Feature.
func (_class FeatureClass) GetNameLabel(sessionID SessionRef, self FeatureRef) (_retval string, _err error) {
	if IsMock {
		return _class.GetNameLabelMock(sessionID, self)
	}	
	_method := "Feature.get_name_label"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertFeatureRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func FeatureClassGetUUIDMockDefault(sessionID SessionRef, self FeatureRef) (_retval string, _err error) {
	log.Println("Feature.GetUUID not mocked")
	_err = errors.New("Feature.GetUUID not mocked")
	return
}

var FeatureClassGetUUIDMockedCallback = FeatureClassGetUUIDMockDefault

func (_class FeatureClass) GetUUIDMock(sessionID SessionRef, self FeatureRef) (_retval string, _err error) {
	return FeatureClassGetUUIDMockedCallback(sessionID, self)
}
// Get the uuid field of the given Feature.
func (_class FeatureClass) GetUUID(sessionID SessionRef, self FeatureRef) (_retval string, _err error) {
	if IsMock {
		return _class.GetUUIDMock(sessionID, self)
	}	
	_method := "Feature.get_uuid"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertFeatureRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func FeatureClassGetByNameLabelMockDefault(sessionID SessionRef, label string) (_retval []FeatureRef, _err error) {
	log.Println("Feature.GetByNameLabel not mocked")
	_err = errors.New("Feature.GetByNameLabel not mocked")
	return
}

var FeatureClassGetByNameLabelMockedCallback = FeatureClassGetByNameLabelMockDefault

func (_class FeatureClass) GetByNameLabelMock(sessionID SessionRef, label string) (_retval []FeatureRef, _err error) {
	return FeatureClassGetByNameLabelMockedCallback(sessionID, label)
}
// Get all the Feature instances with the given label.
func (_class FeatureClass) GetByNameLabel(sessionID SessionRef, label string) (_retval []FeatureRef, _err error) {
	if IsMock {
		return _class.GetByNameLabelMock(sessionID, label)
	}	
	_method := "Feature.get_by_name_label"
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
	_retval, _err = convertFeatureRefSetToGo(_method + " -> ", _result.Value)
	return
}


func FeatureClassGetByUUIDMockDefault(sessionID SessionRef, uuid string) (_retval FeatureRef, _err error) {
	log.Println("Feature.GetByUUID not mocked")
	_err = errors.New("Feature.GetByUUID not mocked")
	return
}

var FeatureClassGetByUUIDMockedCallback = FeatureClassGetByUUIDMockDefault

func (_class FeatureClass) GetByUUIDMock(sessionID SessionRef, uuid string) (_retval FeatureRef, _err error) {
	return FeatureClassGetByUUIDMockedCallback(sessionID, uuid)
}
// Get a reference to the Feature instance with the specified UUID.
func (_class FeatureClass) GetByUUID(sessionID SessionRef, uuid string) (_retval FeatureRef, _err error) {
	if IsMock {
		return _class.GetByUUIDMock(sessionID, uuid)
	}	
	_method := "Feature.get_by_uuid"
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
	_retval, _err = convertFeatureRefToGo(_method + " -> ", _result.Value)
	return
}


func FeatureClassGetRecordMockDefault(sessionID SessionRef, self FeatureRef) (_retval FeatureRecord, _err error) {
	log.Println("Feature.GetRecord not mocked")
	_err = errors.New("Feature.GetRecord not mocked")
	return
}

var FeatureClassGetRecordMockedCallback = FeatureClassGetRecordMockDefault

func (_class FeatureClass) GetRecordMock(sessionID SessionRef, self FeatureRef) (_retval FeatureRecord, _err error) {
	return FeatureClassGetRecordMockedCallback(sessionID, self)
}
// Get a record containing the current state of the given Feature.
func (_class FeatureClass) GetRecord(sessionID SessionRef, self FeatureRef) (_retval FeatureRecord, _err error) {
	if IsMock {
		return _class.GetRecordMock(sessionID, self)
	}	
	_method := "Feature.get_record"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertFeatureRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertFeatureRecordToGo(_method + " -> ", _result.Value)
	return
}
