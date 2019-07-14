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

type SMRecord struct {
  // Unique identifier/object reference
	UUID string
  // a human-readable name
	NameLabel string
  // a notes field containing human-readable description
	NameDescription string
  // SR.type
	Type string
  // Vendor who created this plugin
	Vendor string
  // Entity which owns the copyright of this plugin
	Copyright string
  // Version of the plugin
	Version string
  // Minimum SM API version required on the server
	RequiredAPIVersion string
  // names and descriptions of device config keys
	Configuration map[string]string
  // capabilities of the SM plugin
	Capabilities []string
  // capabilities of the SM plugin, with capability version numbers
	Features map[string]int
  // additional configuration
	OtherConfig map[string]string
  // filename of the storage driver
	DriverFilename string
  // The storage plugin requires that one of these cluster stacks is configured and running.
	RequiredClusterStack []string
}

type SMRef string

// A storage manager plugin
type SMClass struct {
	client *Client
}


func SMClassGetAllRecordsMockDefault(sessionID SessionRef) (_retval map[SMRef]SMRecord, _err error) {
	log.Println("SM.GetAllRecords not mocked")
	_err = errors.New("SM.GetAllRecords not mocked")
	return
}

var SMClassGetAllRecordsMockedCallback = SMClassGetAllRecordsMockDefault

func (_class SMClass) GetAllRecordsMock(sessionID SessionRef) (_retval map[SMRef]SMRecord, _err error) {
	return SMClassGetAllRecordsMockedCallback(sessionID)
}
// Return a map of SM references to SM records for all SMs known to the system.
func (_class SMClass) GetAllRecords(sessionID SessionRef) (_retval map[SMRef]SMRecord, _err error) {
	if IsMock {
		return _class.GetAllRecordsMock(sessionID)
	}	
	_method := "SM.get_all_records"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertSMRefToSMRecordMapToGo(_method + " -> ", _result.Value)
	return
}


func SMClassGetAllMockDefault(sessionID SessionRef) (_retval []SMRef, _err error) {
	log.Println("SM.GetAll not mocked")
	_err = errors.New("SM.GetAll not mocked")
	return
}

var SMClassGetAllMockedCallback = SMClassGetAllMockDefault

func (_class SMClass) GetAllMock(sessionID SessionRef) (_retval []SMRef, _err error) {
	return SMClassGetAllMockedCallback(sessionID)
}
// Return a list of all the SMs known to the system.
func (_class SMClass) GetAll(sessionID SessionRef) (_retval []SMRef, _err error) {
	if IsMock {
		return _class.GetAllMock(sessionID)
	}	
	_method := "SM.get_all"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertSMRefSetToGo(_method + " -> ", _result.Value)
	return
}


func SMClassRemoveFromOtherConfigMockDefault(sessionID SessionRef, self SMRef, key string) (_err error) {
	log.Println("SM.RemoveFromOtherConfig not mocked")
	_err = errors.New("SM.RemoveFromOtherConfig not mocked")
	return
}

var SMClassRemoveFromOtherConfigMockedCallback = SMClassRemoveFromOtherConfigMockDefault

func (_class SMClass) RemoveFromOtherConfigMock(sessionID SessionRef, self SMRef, key string) (_err error) {
	return SMClassRemoveFromOtherConfigMockedCallback(sessionID, self, key)
}
// Remove the given key and its corresponding value from the other_config field of the given SM.  If the key is not in that Map, then do nothing.
func (_class SMClass) RemoveFromOtherConfig(sessionID SessionRef, self SMRef, key string) (_err error) {
	if IsMock {
		return _class.RemoveFromOtherConfigMock(sessionID, self, key)
	}	
	_method := "SM.remove_from_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertSMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func SMClassAddToOtherConfigMockDefault(sessionID SessionRef, self SMRef, key string, value string) (_err error) {
	log.Println("SM.AddToOtherConfig not mocked")
	_err = errors.New("SM.AddToOtherConfig not mocked")
	return
}

var SMClassAddToOtherConfigMockedCallback = SMClassAddToOtherConfigMockDefault

func (_class SMClass) AddToOtherConfigMock(sessionID SessionRef, self SMRef, key string, value string) (_err error) {
	return SMClassAddToOtherConfigMockedCallback(sessionID, self, key, value)
}
// Add the given key-value pair to the other_config field of the given SM.
func (_class SMClass) AddToOtherConfig(sessionID SessionRef, self SMRef, key string, value string) (_err error) {
	if IsMock {
		return _class.AddToOtherConfigMock(sessionID, self, key, value)
	}	
	_method := "SM.add_to_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertSMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func SMClassSetOtherConfigMockDefault(sessionID SessionRef, self SMRef, value map[string]string) (_err error) {
	log.Println("SM.SetOtherConfig not mocked")
	_err = errors.New("SM.SetOtherConfig not mocked")
	return
}

var SMClassSetOtherConfigMockedCallback = SMClassSetOtherConfigMockDefault

func (_class SMClass) SetOtherConfigMock(sessionID SessionRef, self SMRef, value map[string]string) (_err error) {
	return SMClassSetOtherConfigMockedCallback(sessionID, self, value)
}
// Set the other_config field of the given SM.
func (_class SMClass) SetOtherConfig(sessionID SessionRef, self SMRef, value map[string]string) (_err error) {
	if IsMock {
		return _class.SetOtherConfigMock(sessionID, self, value)
	}	
	_method := "SM.set_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertSMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func SMClassGetRequiredClusterStackMockDefault(sessionID SessionRef, self SMRef) (_retval []string, _err error) {
	log.Println("SM.GetRequiredClusterStack not mocked")
	_err = errors.New("SM.GetRequiredClusterStack not mocked")
	return
}

var SMClassGetRequiredClusterStackMockedCallback = SMClassGetRequiredClusterStackMockDefault

func (_class SMClass) GetRequiredClusterStackMock(sessionID SessionRef, self SMRef) (_retval []string, _err error) {
	return SMClassGetRequiredClusterStackMockedCallback(sessionID, self)
}
// Get the required_cluster_stack field of the given SM.
func (_class SMClass) GetRequiredClusterStack(sessionID SessionRef, self SMRef) (_retval []string, _err error) {
	if IsMock {
		return _class.GetRequiredClusterStackMock(sessionID, self)
	}	
	_method := "SM.get_required_cluster_stack"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertSMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func SMClassGetDriverFilenameMockDefault(sessionID SessionRef, self SMRef) (_retval string, _err error) {
	log.Println("SM.GetDriverFilename not mocked")
	_err = errors.New("SM.GetDriverFilename not mocked")
	return
}

var SMClassGetDriverFilenameMockedCallback = SMClassGetDriverFilenameMockDefault

func (_class SMClass) GetDriverFilenameMock(sessionID SessionRef, self SMRef) (_retval string, _err error) {
	return SMClassGetDriverFilenameMockedCallback(sessionID, self)
}
// Get the driver_filename field of the given SM.
func (_class SMClass) GetDriverFilename(sessionID SessionRef, self SMRef) (_retval string, _err error) {
	if IsMock {
		return _class.GetDriverFilenameMock(sessionID, self)
	}	
	_method := "SM.get_driver_filename"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertSMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func SMClassGetOtherConfigMockDefault(sessionID SessionRef, self SMRef) (_retval map[string]string, _err error) {
	log.Println("SM.GetOtherConfig not mocked")
	_err = errors.New("SM.GetOtherConfig not mocked")
	return
}

var SMClassGetOtherConfigMockedCallback = SMClassGetOtherConfigMockDefault

func (_class SMClass) GetOtherConfigMock(sessionID SessionRef, self SMRef) (_retval map[string]string, _err error) {
	return SMClassGetOtherConfigMockedCallback(sessionID, self)
}
// Get the other_config field of the given SM.
func (_class SMClass) GetOtherConfig(sessionID SessionRef, self SMRef) (_retval map[string]string, _err error) {
	if IsMock {
		return _class.GetOtherConfigMock(sessionID, self)
	}	
	_method := "SM.get_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertSMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func SMClassGetFeaturesMockDefault(sessionID SessionRef, self SMRef) (_retval map[string]int, _err error) {
	log.Println("SM.GetFeatures not mocked")
	_err = errors.New("SM.GetFeatures not mocked")
	return
}

var SMClassGetFeaturesMockedCallback = SMClassGetFeaturesMockDefault

func (_class SMClass) GetFeaturesMock(sessionID SessionRef, self SMRef) (_retval map[string]int, _err error) {
	return SMClassGetFeaturesMockedCallback(sessionID, self)
}
// Get the features field of the given SM.
func (_class SMClass) GetFeatures(sessionID SessionRef, self SMRef) (_retval map[string]int, _err error) {
	if IsMock {
		return _class.GetFeaturesMock(sessionID, self)
	}	
	_method := "SM.get_features"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertSMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringToIntMapToGo(_method + " -> ", _result.Value)
	return
}


func SMClassGetCapabilitiesMockDefault(sessionID SessionRef, self SMRef) (_retval []string, _err error) {
	log.Println("SM.GetCapabilities not mocked")
	_err = errors.New("SM.GetCapabilities not mocked")
	return
}

var SMClassGetCapabilitiesMockedCallback = SMClassGetCapabilitiesMockDefault

func (_class SMClass) GetCapabilitiesMock(sessionID SessionRef, self SMRef) (_retval []string, _err error) {
	return SMClassGetCapabilitiesMockedCallback(sessionID, self)
}
// Get the capabilities field of the given SM.
func (_class SMClass) GetCapabilities(sessionID SessionRef, self SMRef) (_retval []string, _err error) {
	if IsMock {
		return _class.GetCapabilitiesMock(sessionID, self)
	}	
	_method := "SM.get_capabilities"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertSMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func SMClassGetConfigurationMockDefault(sessionID SessionRef, self SMRef) (_retval map[string]string, _err error) {
	log.Println("SM.GetConfiguration not mocked")
	_err = errors.New("SM.GetConfiguration not mocked")
	return
}

var SMClassGetConfigurationMockedCallback = SMClassGetConfigurationMockDefault

func (_class SMClass) GetConfigurationMock(sessionID SessionRef, self SMRef) (_retval map[string]string, _err error) {
	return SMClassGetConfigurationMockedCallback(sessionID, self)
}
// Get the configuration field of the given SM.
func (_class SMClass) GetConfiguration(sessionID SessionRef, self SMRef) (_retval map[string]string, _err error) {
	if IsMock {
		return _class.GetConfigurationMock(sessionID, self)
	}	
	_method := "SM.get_configuration"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertSMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func SMClassGetRequiredAPIVersionMockDefault(sessionID SessionRef, self SMRef) (_retval string, _err error) {
	log.Println("SM.GetRequiredAPIVersion not mocked")
	_err = errors.New("SM.GetRequiredAPIVersion not mocked")
	return
}

var SMClassGetRequiredAPIVersionMockedCallback = SMClassGetRequiredAPIVersionMockDefault

func (_class SMClass) GetRequiredAPIVersionMock(sessionID SessionRef, self SMRef) (_retval string, _err error) {
	return SMClassGetRequiredAPIVersionMockedCallback(sessionID, self)
}
// Get the required_api_version field of the given SM.
func (_class SMClass) GetRequiredAPIVersion(sessionID SessionRef, self SMRef) (_retval string, _err error) {
	if IsMock {
		return _class.GetRequiredAPIVersionMock(sessionID, self)
	}	
	_method := "SM.get_required_api_version"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertSMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func SMClassGetVersionMockDefault(sessionID SessionRef, self SMRef) (_retval string, _err error) {
	log.Println("SM.GetVersion not mocked")
	_err = errors.New("SM.GetVersion not mocked")
	return
}

var SMClassGetVersionMockedCallback = SMClassGetVersionMockDefault

func (_class SMClass) GetVersionMock(sessionID SessionRef, self SMRef) (_retval string, _err error) {
	return SMClassGetVersionMockedCallback(sessionID, self)
}
// Get the version field of the given SM.
func (_class SMClass) GetVersion(sessionID SessionRef, self SMRef) (_retval string, _err error) {
	if IsMock {
		return _class.GetVersionMock(sessionID, self)
	}	
	_method := "SM.get_version"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertSMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func SMClassGetCopyrightMockDefault(sessionID SessionRef, self SMRef) (_retval string, _err error) {
	log.Println("SM.GetCopyright not mocked")
	_err = errors.New("SM.GetCopyright not mocked")
	return
}

var SMClassGetCopyrightMockedCallback = SMClassGetCopyrightMockDefault

func (_class SMClass) GetCopyrightMock(sessionID SessionRef, self SMRef) (_retval string, _err error) {
	return SMClassGetCopyrightMockedCallback(sessionID, self)
}
// Get the copyright field of the given SM.
func (_class SMClass) GetCopyright(sessionID SessionRef, self SMRef) (_retval string, _err error) {
	if IsMock {
		return _class.GetCopyrightMock(sessionID, self)
	}	
	_method := "SM.get_copyright"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertSMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func SMClassGetVendorMockDefault(sessionID SessionRef, self SMRef) (_retval string, _err error) {
	log.Println("SM.GetVendor not mocked")
	_err = errors.New("SM.GetVendor not mocked")
	return
}

var SMClassGetVendorMockedCallback = SMClassGetVendorMockDefault

func (_class SMClass) GetVendorMock(sessionID SessionRef, self SMRef) (_retval string, _err error) {
	return SMClassGetVendorMockedCallback(sessionID, self)
}
// Get the vendor field of the given SM.
func (_class SMClass) GetVendor(sessionID SessionRef, self SMRef) (_retval string, _err error) {
	if IsMock {
		return _class.GetVendorMock(sessionID, self)
	}	
	_method := "SM.get_vendor"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertSMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func SMClassGetTypeMockDefault(sessionID SessionRef, self SMRef) (_retval string, _err error) {
	log.Println("SM.GetType not mocked")
	_err = errors.New("SM.GetType not mocked")
	return
}

var SMClassGetTypeMockedCallback = SMClassGetTypeMockDefault

func (_class SMClass) GetTypeMock(sessionID SessionRef, self SMRef) (_retval string, _err error) {
	return SMClassGetTypeMockedCallback(sessionID, self)
}
// Get the type field of the given SM.
func (_class SMClass) GetType(sessionID SessionRef, self SMRef) (_retval string, _err error) {
	if IsMock {
		return _class.GetTypeMock(sessionID, self)
	}	
	_method := "SM.get_type"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertSMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func SMClassGetNameDescriptionMockDefault(sessionID SessionRef, self SMRef) (_retval string, _err error) {
	log.Println("SM.GetNameDescription not mocked")
	_err = errors.New("SM.GetNameDescription not mocked")
	return
}

var SMClassGetNameDescriptionMockedCallback = SMClassGetNameDescriptionMockDefault

func (_class SMClass) GetNameDescriptionMock(sessionID SessionRef, self SMRef) (_retval string, _err error) {
	return SMClassGetNameDescriptionMockedCallback(sessionID, self)
}
// Get the name/description field of the given SM.
func (_class SMClass) GetNameDescription(sessionID SessionRef, self SMRef) (_retval string, _err error) {
	if IsMock {
		return _class.GetNameDescriptionMock(sessionID, self)
	}	
	_method := "SM.get_name_description"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertSMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func SMClassGetNameLabelMockDefault(sessionID SessionRef, self SMRef) (_retval string, _err error) {
	log.Println("SM.GetNameLabel not mocked")
	_err = errors.New("SM.GetNameLabel not mocked")
	return
}

var SMClassGetNameLabelMockedCallback = SMClassGetNameLabelMockDefault

func (_class SMClass) GetNameLabelMock(sessionID SessionRef, self SMRef) (_retval string, _err error) {
	return SMClassGetNameLabelMockedCallback(sessionID, self)
}
// Get the name/label field of the given SM.
func (_class SMClass) GetNameLabel(sessionID SessionRef, self SMRef) (_retval string, _err error) {
	if IsMock {
		return _class.GetNameLabelMock(sessionID, self)
	}	
	_method := "SM.get_name_label"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertSMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func SMClassGetUUIDMockDefault(sessionID SessionRef, self SMRef) (_retval string, _err error) {
	log.Println("SM.GetUUID not mocked")
	_err = errors.New("SM.GetUUID not mocked")
	return
}

var SMClassGetUUIDMockedCallback = SMClassGetUUIDMockDefault

func (_class SMClass) GetUUIDMock(sessionID SessionRef, self SMRef) (_retval string, _err error) {
	return SMClassGetUUIDMockedCallback(sessionID, self)
}
// Get the uuid field of the given SM.
func (_class SMClass) GetUUID(sessionID SessionRef, self SMRef) (_retval string, _err error) {
	if IsMock {
		return _class.GetUUIDMock(sessionID, self)
	}	
	_method := "SM.get_uuid"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertSMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func SMClassGetByNameLabelMockDefault(sessionID SessionRef, label string) (_retval []SMRef, _err error) {
	log.Println("SM.GetByNameLabel not mocked")
	_err = errors.New("SM.GetByNameLabel not mocked")
	return
}

var SMClassGetByNameLabelMockedCallback = SMClassGetByNameLabelMockDefault

func (_class SMClass) GetByNameLabelMock(sessionID SessionRef, label string) (_retval []SMRef, _err error) {
	return SMClassGetByNameLabelMockedCallback(sessionID, label)
}
// Get all the SM instances with the given label.
func (_class SMClass) GetByNameLabel(sessionID SessionRef, label string) (_retval []SMRef, _err error) {
	if IsMock {
		return _class.GetByNameLabelMock(sessionID, label)
	}	
	_method := "SM.get_by_name_label"
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
	_retval, _err = convertSMRefSetToGo(_method + " -> ", _result.Value)
	return
}


func SMClassGetByUUIDMockDefault(sessionID SessionRef, uuid string) (_retval SMRef, _err error) {
	log.Println("SM.GetByUUID not mocked")
	_err = errors.New("SM.GetByUUID not mocked")
	return
}

var SMClassGetByUUIDMockedCallback = SMClassGetByUUIDMockDefault

func (_class SMClass) GetByUUIDMock(sessionID SessionRef, uuid string) (_retval SMRef, _err error) {
	return SMClassGetByUUIDMockedCallback(sessionID, uuid)
}
// Get a reference to the SM instance with the specified UUID.
func (_class SMClass) GetByUUID(sessionID SessionRef, uuid string) (_retval SMRef, _err error) {
	if IsMock {
		return _class.GetByUUIDMock(sessionID, uuid)
	}	
	_method := "SM.get_by_uuid"
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
	_retval, _err = convertSMRefToGo(_method + " -> ", _result.Value)
	return
}


func SMClassGetRecordMockDefault(sessionID SessionRef, self SMRef) (_retval SMRecord, _err error) {
	log.Println("SM.GetRecord not mocked")
	_err = errors.New("SM.GetRecord not mocked")
	return
}

var SMClassGetRecordMockedCallback = SMClassGetRecordMockDefault

func (_class SMClass) GetRecordMock(sessionID SessionRef, self SMRef) (_retval SMRecord, _err error) {
	return SMClassGetRecordMockedCallback(sessionID, self)
}
// Get a record containing the current state of the given SM.
func (_class SMClass) GetRecord(sessionID SessionRef, self SMRef) (_retval SMRecord, _err error) {
	if IsMock {
		return _class.GetRecordMock(sessionID, self)
	}	
	_method := "SM.get_record"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertSMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertSMRecordToGo(_method + " -> ", _result.Value)
	return
}
