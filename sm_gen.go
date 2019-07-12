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

func (_class SMClass) GetAllRecords__mock(sessionID SessionRef) (_retval map[SMRef]SMRecord, _err error) {
	log.Println("SM.GetAllRecords not mocked")
	_err = errors.New("SM.GetAllRecords not mocked")
	return
}
// Return a map of SM references to SM records for all SMs known to the system.
func (_class SMClass) GetAllRecords(sessionID SessionRef) (_retval map[SMRef]SMRecord, _err error) {
	if (IsMock) {
		return _class.GetAllRecords__mock(sessionID)
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

func (_class SMClass) GetAll__mock(sessionID SessionRef) (_retval []SMRef, _err error) {
	log.Println("SM.GetAll not mocked")
	_err = errors.New("SM.GetAll not mocked")
	return
}
// Return a list of all the SMs known to the system.
func (_class SMClass) GetAll(sessionID SessionRef) (_retval []SMRef, _err error) {
	if (IsMock) {
		return _class.GetAll__mock(sessionID)
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

func (_class SMClass) RemoveFromOtherConfig__mock(sessionID SessionRef, self SMRef, key string) (_err error) {
	log.Println("SM.RemoveFromOtherConfig not mocked")
	_err = errors.New("SM.RemoveFromOtherConfig not mocked")
	return
}
// Remove the given key and its corresponding value from the other_config field of the given SM.  If the key is not in that Map, then do nothing.
func (_class SMClass) RemoveFromOtherConfig(sessionID SessionRef, self SMRef, key string) (_err error) {
	if (IsMock) {
		return _class.RemoveFromOtherConfig__mock(sessionID, self, key)
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

func (_class SMClass) AddToOtherConfig__mock(sessionID SessionRef, self SMRef, key string, value string) (_err error) {
	log.Println("SM.AddToOtherConfig not mocked")
	_err = errors.New("SM.AddToOtherConfig not mocked")
	return
}
// Add the given key-value pair to the other_config field of the given SM.
func (_class SMClass) AddToOtherConfig(sessionID SessionRef, self SMRef, key string, value string) (_err error) {
	if (IsMock) {
		return _class.AddToOtherConfig__mock(sessionID, self, key, value)
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

func (_class SMClass) SetOtherConfig__mock(sessionID SessionRef, self SMRef, value map[string]string) (_err error) {
	log.Println("SM.SetOtherConfig not mocked")
	_err = errors.New("SM.SetOtherConfig not mocked")
	return
}
// Set the other_config field of the given SM.
func (_class SMClass) SetOtherConfig(sessionID SessionRef, self SMRef, value map[string]string) (_err error) {
	if (IsMock) {
		return _class.SetOtherConfig__mock(sessionID, self, value)
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

func (_class SMClass) GetRequiredClusterStack__mock(sessionID SessionRef, self SMRef) (_retval []string, _err error) {
	log.Println("SM.GetRequiredClusterStack not mocked")
	_err = errors.New("SM.GetRequiredClusterStack not mocked")
	return
}
// Get the required_cluster_stack field of the given SM.
func (_class SMClass) GetRequiredClusterStack(sessionID SessionRef, self SMRef) (_retval []string, _err error) {
	if (IsMock) {
		return _class.GetRequiredClusterStack__mock(sessionID, self)
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

func (_class SMClass) GetDriverFilename__mock(sessionID SessionRef, self SMRef) (_retval string, _err error) {
	log.Println("SM.GetDriverFilename not mocked")
	_err = errors.New("SM.GetDriverFilename not mocked")
	return
}
// Get the driver_filename field of the given SM.
func (_class SMClass) GetDriverFilename(sessionID SessionRef, self SMRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetDriverFilename__mock(sessionID, self)
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

func (_class SMClass) GetOtherConfig__mock(sessionID SessionRef, self SMRef) (_retval map[string]string, _err error) {
	log.Println("SM.GetOtherConfig not mocked")
	_err = errors.New("SM.GetOtherConfig not mocked")
	return
}
// Get the other_config field of the given SM.
func (_class SMClass) GetOtherConfig(sessionID SessionRef, self SMRef) (_retval map[string]string, _err error) {
	if (IsMock) {
		return _class.GetOtherConfig__mock(sessionID, self)
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

func (_class SMClass) GetFeatures__mock(sessionID SessionRef, self SMRef) (_retval map[string]int, _err error) {
	log.Println("SM.GetFeatures not mocked")
	_err = errors.New("SM.GetFeatures not mocked")
	return
}
// Get the features field of the given SM.
func (_class SMClass) GetFeatures(sessionID SessionRef, self SMRef) (_retval map[string]int, _err error) {
	if (IsMock) {
		return _class.GetFeatures__mock(sessionID, self)
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

func (_class SMClass) GetCapabilities__mock(sessionID SessionRef, self SMRef) (_retval []string, _err error) {
	log.Println("SM.GetCapabilities not mocked")
	_err = errors.New("SM.GetCapabilities not mocked")
	return
}
// Get the capabilities field of the given SM.
func (_class SMClass) GetCapabilities(sessionID SessionRef, self SMRef) (_retval []string, _err error) {
	if (IsMock) {
		return _class.GetCapabilities__mock(sessionID, self)
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

func (_class SMClass) GetConfiguration__mock(sessionID SessionRef, self SMRef) (_retval map[string]string, _err error) {
	log.Println("SM.GetConfiguration not mocked")
	_err = errors.New("SM.GetConfiguration not mocked")
	return
}
// Get the configuration field of the given SM.
func (_class SMClass) GetConfiguration(sessionID SessionRef, self SMRef) (_retval map[string]string, _err error) {
	if (IsMock) {
		return _class.GetConfiguration__mock(sessionID, self)
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

func (_class SMClass) GetRequiredAPIVersion__mock(sessionID SessionRef, self SMRef) (_retval string, _err error) {
	log.Println("SM.GetRequiredAPIVersion not mocked")
	_err = errors.New("SM.GetRequiredAPIVersion not mocked")
	return
}
// Get the required_api_version field of the given SM.
func (_class SMClass) GetRequiredAPIVersion(sessionID SessionRef, self SMRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetRequiredAPIVersion__mock(sessionID, self)
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

func (_class SMClass) GetVersion__mock(sessionID SessionRef, self SMRef) (_retval string, _err error) {
	log.Println("SM.GetVersion not mocked")
	_err = errors.New("SM.GetVersion not mocked")
	return
}
// Get the version field of the given SM.
func (_class SMClass) GetVersion(sessionID SessionRef, self SMRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetVersion__mock(sessionID, self)
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

func (_class SMClass) GetCopyright__mock(sessionID SessionRef, self SMRef) (_retval string, _err error) {
	log.Println("SM.GetCopyright not mocked")
	_err = errors.New("SM.GetCopyright not mocked")
	return
}
// Get the copyright field of the given SM.
func (_class SMClass) GetCopyright(sessionID SessionRef, self SMRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetCopyright__mock(sessionID, self)
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

func (_class SMClass) GetVendor__mock(sessionID SessionRef, self SMRef) (_retval string, _err error) {
	log.Println("SM.GetVendor not mocked")
	_err = errors.New("SM.GetVendor not mocked")
	return
}
// Get the vendor field of the given SM.
func (_class SMClass) GetVendor(sessionID SessionRef, self SMRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetVendor__mock(sessionID, self)
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

func (_class SMClass) GetType__mock(sessionID SessionRef, self SMRef) (_retval string, _err error) {
	log.Println("SM.GetType not mocked")
	_err = errors.New("SM.GetType not mocked")
	return
}
// Get the type field of the given SM.
func (_class SMClass) GetType(sessionID SessionRef, self SMRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetType__mock(sessionID, self)
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

func (_class SMClass) GetNameDescription__mock(sessionID SessionRef, self SMRef) (_retval string, _err error) {
	log.Println("SM.GetNameDescription not mocked")
	_err = errors.New("SM.GetNameDescription not mocked")
	return
}
// Get the name/description field of the given SM.
func (_class SMClass) GetNameDescription(sessionID SessionRef, self SMRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetNameDescription__mock(sessionID, self)
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

func (_class SMClass) GetNameLabel__mock(sessionID SessionRef, self SMRef) (_retval string, _err error) {
	log.Println("SM.GetNameLabel not mocked")
	_err = errors.New("SM.GetNameLabel not mocked")
	return
}
// Get the name/label field of the given SM.
func (_class SMClass) GetNameLabel(sessionID SessionRef, self SMRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetNameLabel__mock(sessionID, self)
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

func (_class SMClass) GetUUID__mock(sessionID SessionRef, self SMRef) (_retval string, _err error) {
	log.Println("SM.GetUUID not mocked")
	_err = errors.New("SM.GetUUID not mocked")
	return
}
// Get the uuid field of the given SM.
func (_class SMClass) GetUUID(sessionID SessionRef, self SMRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetUUID__mock(sessionID, self)
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

func (_class SMClass) GetByNameLabel__mock(sessionID SessionRef, label string) (_retval []SMRef, _err error) {
	log.Println("SM.GetByNameLabel not mocked")
	_err = errors.New("SM.GetByNameLabel not mocked")
	return
}
// Get all the SM instances with the given label.
func (_class SMClass) GetByNameLabel(sessionID SessionRef, label string) (_retval []SMRef, _err error) {
	if (IsMock) {
		return _class.GetByNameLabel__mock(sessionID, label)
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

func (_class SMClass) GetByUUID__mock(sessionID SessionRef, uuid string) (_retval SMRef, _err error) {
	log.Println("SM.GetByUUID not mocked")
	_err = errors.New("SM.GetByUUID not mocked")
	return
}
// Get a reference to the SM instance with the specified UUID.
func (_class SMClass) GetByUUID(sessionID SessionRef, uuid string) (_retval SMRef, _err error) {
	if (IsMock) {
		return _class.GetByUUID__mock(sessionID, uuid)
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

func (_class SMClass) GetRecord__mock(sessionID SessionRef, self SMRef) (_retval SMRecord, _err error) {
	log.Println("SM.GetRecord not mocked")
	_err = errors.New("SM.GetRecord not mocked")
	return
}
// Get a record containing the current state of the given SM.
func (_class SMClass) GetRecord(sessionID SessionRef, self SMRef) (_retval SMRecord, _err error) {
	if (IsMock) {
		return _class.GetRecord__mock(sessionID, self)
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
