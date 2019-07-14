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

type HostCPURecord struct {
  // Unique identifier/object reference
	UUID string
  // the host the CPU is in
	Host HostRef
  // the number of the physical CPU within the host
	Number int
  // the vendor of the physical CPU
	Vendor string
  // the speed of the physical CPU
	Speed int
  // the model name of the physical CPU
	Modelname string
  // the family (number) of the physical CPU
	Family int
  // the model number of the physical CPU
	Model int
  // the stepping of the physical CPU
	Stepping string
  // the flags of the physical CPU (a decoded version of the features field)
	Flags string
  // the physical CPU feature bitmap
	Features string
  // the current CPU utilisation
	Utilisation float64
  // additional configuration
	OtherConfig map[string]string
}

type HostCPURef string

// A physical CPU
type HostCPUClass struct {
	client *Client
}


func HostCPUClassGetAllRecordsMockDefault(sessionID SessionRef) (_retval map[HostCPURef]HostCPURecord, _err error) {
	log.Println("HostCPU.GetAllRecords not mocked")
	_err = errors.New("HostCPU.GetAllRecords not mocked")
	return
}

var HostCPUClassGetAllRecordsMockedCallback = HostCPUClassGetAllRecordsMockDefault

func (_class HostCPUClass) GetAllRecordsMock(sessionID SessionRef) (_retval map[HostCPURef]HostCPURecord, _err error) {
	return HostCPUClassGetAllRecordsMockedCallback(sessionID)
}
// Return a map of host_cpu references to host_cpu records for all host_cpus known to the system.
func (_class HostCPUClass) GetAllRecords(sessionID SessionRef) (_retval map[HostCPURef]HostCPURecord, _err error) {
	if IsMock {
		return _class.GetAllRecordsMock(sessionID)
	}	
	_method := "host_cpu.get_all_records"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertHostCPURefToHostCPURecordMapToGo(_method + " -> ", _result.Value)
	return
}


func HostCPUClassGetAllMockDefault(sessionID SessionRef) (_retval []HostCPURef, _err error) {
	log.Println("HostCPU.GetAll not mocked")
	_err = errors.New("HostCPU.GetAll not mocked")
	return
}

var HostCPUClassGetAllMockedCallback = HostCPUClassGetAllMockDefault

func (_class HostCPUClass) GetAllMock(sessionID SessionRef) (_retval []HostCPURef, _err error) {
	return HostCPUClassGetAllMockedCallback(sessionID)
}
// Return a list of all the host_cpus known to the system.
func (_class HostCPUClass) GetAll(sessionID SessionRef) (_retval []HostCPURef, _err error) {
	if IsMock {
		return _class.GetAllMock(sessionID)
	}	
	_method := "host_cpu.get_all"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertHostCPURefSetToGo(_method + " -> ", _result.Value)
	return
}


func HostCPUClassRemoveFromOtherConfigMockDefault(sessionID SessionRef, self HostCPURef, key string) (_err error) {
	log.Println("HostCPU.RemoveFromOtherConfig not mocked")
	_err = errors.New("HostCPU.RemoveFromOtherConfig not mocked")
	return
}

var HostCPUClassRemoveFromOtherConfigMockedCallback = HostCPUClassRemoveFromOtherConfigMockDefault

func (_class HostCPUClass) RemoveFromOtherConfigMock(sessionID SessionRef, self HostCPURef, key string) (_err error) {
	return HostCPUClassRemoveFromOtherConfigMockedCallback(sessionID, self, key)
}
// Remove the given key and its corresponding value from the other_config field of the given host_cpu.  If the key is not in that Map, then do nothing.
func (_class HostCPUClass) RemoveFromOtherConfig(sessionID SessionRef, self HostCPURef, key string) (_err error) {
	if IsMock {
		return _class.RemoveFromOtherConfigMock(sessionID, self, key)
	}	
	_method := "host_cpu.remove_from_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertHostCPURefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func HostCPUClassAddToOtherConfigMockDefault(sessionID SessionRef, self HostCPURef, key string, value string) (_err error) {
	log.Println("HostCPU.AddToOtherConfig not mocked")
	_err = errors.New("HostCPU.AddToOtherConfig not mocked")
	return
}

var HostCPUClassAddToOtherConfigMockedCallback = HostCPUClassAddToOtherConfigMockDefault

func (_class HostCPUClass) AddToOtherConfigMock(sessionID SessionRef, self HostCPURef, key string, value string) (_err error) {
	return HostCPUClassAddToOtherConfigMockedCallback(sessionID, self, key, value)
}
// Add the given key-value pair to the other_config field of the given host_cpu.
func (_class HostCPUClass) AddToOtherConfig(sessionID SessionRef, self HostCPURef, key string, value string) (_err error) {
	if IsMock {
		return _class.AddToOtherConfigMock(sessionID, self, key, value)
	}	
	_method := "host_cpu.add_to_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertHostCPURefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func HostCPUClassSetOtherConfigMockDefault(sessionID SessionRef, self HostCPURef, value map[string]string) (_err error) {
	log.Println("HostCPU.SetOtherConfig not mocked")
	_err = errors.New("HostCPU.SetOtherConfig not mocked")
	return
}

var HostCPUClassSetOtherConfigMockedCallback = HostCPUClassSetOtherConfigMockDefault

func (_class HostCPUClass) SetOtherConfigMock(sessionID SessionRef, self HostCPURef, value map[string]string) (_err error) {
	return HostCPUClassSetOtherConfigMockedCallback(sessionID, self, value)
}
// Set the other_config field of the given host_cpu.
func (_class HostCPUClass) SetOtherConfig(sessionID SessionRef, self HostCPURef, value map[string]string) (_err error) {
	if IsMock {
		return _class.SetOtherConfigMock(sessionID, self, value)
	}	
	_method := "host_cpu.set_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertHostCPURefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func HostCPUClassGetOtherConfigMockDefault(sessionID SessionRef, self HostCPURef) (_retval map[string]string, _err error) {
	log.Println("HostCPU.GetOtherConfig not mocked")
	_err = errors.New("HostCPU.GetOtherConfig not mocked")
	return
}

var HostCPUClassGetOtherConfigMockedCallback = HostCPUClassGetOtherConfigMockDefault

func (_class HostCPUClass) GetOtherConfigMock(sessionID SessionRef, self HostCPURef) (_retval map[string]string, _err error) {
	return HostCPUClassGetOtherConfigMockedCallback(sessionID, self)
}
// Get the other_config field of the given host_cpu.
func (_class HostCPUClass) GetOtherConfig(sessionID SessionRef, self HostCPURef) (_retval map[string]string, _err error) {
	if IsMock {
		return _class.GetOtherConfigMock(sessionID, self)
	}	
	_method := "host_cpu.get_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertHostCPURefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func HostCPUClassGetUtilisationMockDefault(sessionID SessionRef, self HostCPURef) (_retval float64, _err error) {
	log.Println("HostCPU.GetUtilisation not mocked")
	_err = errors.New("HostCPU.GetUtilisation not mocked")
	return
}

var HostCPUClassGetUtilisationMockedCallback = HostCPUClassGetUtilisationMockDefault

func (_class HostCPUClass) GetUtilisationMock(sessionID SessionRef, self HostCPURef) (_retval float64, _err error) {
	return HostCPUClassGetUtilisationMockedCallback(sessionID, self)
}
// Get the utilisation field of the given host_cpu.
func (_class HostCPUClass) GetUtilisation(sessionID SessionRef, self HostCPURef) (_retval float64, _err error) {
	if IsMock {
		return _class.GetUtilisationMock(sessionID, self)
	}	
	_method := "host_cpu.get_utilisation"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertHostCPURefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertFloatToGo(_method + " -> ", _result.Value)
	return
}


func HostCPUClassGetFeaturesMockDefault(sessionID SessionRef, self HostCPURef) (_retval string, _err error) {
	log.Println("HostCPU.GetFeatures not mocked")
	_err = errors.New("HostCPU.GetFeatures not mocked")
	return
}

var HostCPUClassGetFeaturesMockedCallback = HostCPUClassGetFeaturesMockDefault

func (_class HostCPUClass) GetFeaturesMock(sessionID SessionRef, self HostCPURef) (_retval string, _err error) {
	return HostCPUClassGetFeaturesMockedCallback(sessionID, self)
}
// Get the features field of the given host_cpu.
func (_class HostCPUClass) GetFeatures(sessionID SessionRef, self HostCPURef) (_retval string, _err error) {
	if IsMock {
		return _class.GetFeaturesMock(sessionID, self)
	}	
	_method := "host_cpu.get_features"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertHostCPURefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func HostCPUClassGetFlagsMockDefault(sessionID SessionRef, self HostCPURef) (_retval string, _err error) {
	log.Println("HostCPU.GetFlags not mocked")
	_err = errors.New("HostCPU.GetFlags not mocked")
	return
}

var HostCPUClassGetFlagsMockedCallback = HostCPUClassGetFlagsMockDefault

func (_class HostCPUClass) GetFlagsMock(sessionID SessionRef, self HostCPURef) (_retval string, _err error) {
	return HostCPUClassGetFlagsMockedCallback(sessionID, self)
}
// Get the flags field of the given host_cpu.
func (_class HostCPUClass) GetFlags(sessionID SessionRef, self HostCPURef) (_retval string, _err error) {
	if IsMock {
		return _class.GetFlagsMock(sessionID, self)
	}	
	_method := "host_cpu.get_flags"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertHostCPURefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func HostCPUClassGetSteppingMockDefault(sessionID SessionRef, self HostCPURef) (_retval string, _err error) {
	log.Println("HostCPU.GetStepping not mocked")
	_err = errors.New("HostCPU.GetStepping not mocked")
	return
}

var HostCPUClassGetSteppingMockedCallback = HostCPUClassGetSteppingMockDefault

func (_class HostCPUClass) GetSteppingMock(sessionID SessionRef, self HostCPURef) (_retval string, _err error) {
	return HostCPUClassGetSteppingMockedCallback(sessionID, self)
}
// Get the stepping field of the given host_cpu.
func (_class HostCPUClass) GetStepping(sessionID SessionRef, self HostCPURef) (_retval string, _err error) {
	if IsMock {
		return _class.GetSteppingMock(sessionID, self)
	}	
	_method := "host_cpu.get_stepping"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertHostCPURefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func HostCPUClassGetModelMockDefault(sessionID SessionRef, self HostCPURef) (_retval int, _err error) {
	log.Println("HostCPU.GetModel not mocked")
	_err = errors.New("HostCPU.GetModel not mocked")
	return
}

var HostCPUClassGetModelMockedCallback = HostCPUClassGetModelMockDefault

func (_class HostCPUClass) GetModelMock(sessionID SessionRef, self HostCPURef) (_retval int, _err error) {
	return HostCPUClassGetModelMockedCallback(sessionID, self)
}
// Get the model field of the given host_cpu.
func (_class HostCPUClass) GetModel(sessionID SessionRef, self HostCPURef) (_retval int, _err error) {
	if IsMock {
		return _class.GetModelMock(sessionID, self)
	}	
	_method := "host_cpu.get_model"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertHostCPURefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertIntToGo(_method + " -> ", _result.Value)
	return
}


func HostCPUClassGetFamilyMockDefault(sessionID SessionRef, self HostCPURef) (_retval int, _err error) {
	log.Println("HostCPU.GetFamily not mocked")
	_err = errors.New("HostCPU.GetFamily not mocked")
	return
}

var HostCPUClassGetFamilyMockedCallback = HostCPUClassGetFamilyMockDefault

func (_class HostCPUClass) GetFamilyMock(sessionID SessionRef, self HostCPURef) (_retval int, _err error) {
	return HostCPUClassGetFamilyMockedCallback(sessionID, self)
}
// Get the family field of the given host_cpu.
func (_class HostCPUClass) GetFamily(sessionID SessionRef, self HostCPURef) (_retval int, _err error) {
	if IsMock {
		return _class.GetFamilyMock(sessionID, self)
	}	
	_method := "host_cpu.get_family"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertHostCPURefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertIntToGo(_method + " -> ", _result.Value)
	return
}


func HostCPUClassGetModelnameMockDefault(sessionID SessionRef, self HostCPURef) (_retval string, _err error) {
	log.Println("HostCPU.GetModelname not mocked")
	_err = errors.New("HostCPU.GetModelname not mocked")
	return
}

var HostCPUClassGetModelnameMockedCallback = HostCPUClassGetModelnameMockDefault

func (_class HostCPUClass) GetModelnameMock(sessionID SessionRef, self HostCPURef) (_retval string, _err error) {
	return HostCPUClassGetModelnameMockedCallback(sessionID, self)
}
// Get the modelname field of the given host_cpu.
func (_class HostCPUClass) GetModelname(sessionID SessionRef, self HostCPURef) (_retval string, _err error) {
	if IsMock {
		return _class.GetModelnameMock(sessionID, self)
	}	
	_method := "host_cpu.get_modelname"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertHostCPURefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func HostCPUClassGetSpeedMockDefault(sessionID SessionRef, self HostCPURef) (_retval int, _err error) {
	log.Println("HostCPU.GetSpeed not mocked")
	_err = errors.New("HostCPU.GetSpeed not mocked")
	return
}

var HostCPUClassGetSpeedMockedCallback = HostCPUClassGetSpeedMockDefault

func (_class HostCPUClass) GetSpeedMock(sessionID SessionRef, self HostCPURef) (_retval int, _err error) {
	return HostCPUClassGetSpeedMockedCallback(sessionID, self)
}
// Get the speed field of the given host_cpu.
func (_class HostCPUClass) GetSpeed(sessionID SessionRef, self HostCPURef) (_retval int, _err error) {
	if IsMock {
		return _class.GetSpeedMock(sessionID, self)
	}	
	_method := "host_cpu.get_speed"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertHostCPURefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertIntToGo(_method + " -> ", _result.Value)
	return
}


func HostCPUClassGetVendorMockDefault(sessionID SessionRef, self HostCPURef) (_retval string, _err error) {
	log.Println("HostCPU.GetVendor not mocked")
	_err = errors.New("HostCPU.GetVendor not mocked")
	return
}

var HostCPUClassGetVendorMockedCallback = HostCPUClassGetVendorMockDefault

func (_class HostCPUClass) GetVendorMock(sessionID SessionRef, self HostCPURef) (_retval string, _err error) {
	return HostCPUClassGetVendorMockedCallback(sessionID, self)
}
// Get the vendor field of the given host_cpu.
func (_class HostCPUClass) GetVendor(sessionID SessionRef, self HostCPURef) (_retval string, _err error) {
	if IsMock {
		return _class.GetVendorMock(sessionID, self)
	}	
	_method := "host_cpu.get_vendor"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertHostCPURefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func HostCPUClassGetNumberMockDefault(sessionID SessionRef, self HostCPURef) (_retval int, _err error) {
	log.Println("HostCPU.GetNumber not mocked")
	_err = errors.New("HostCPU.GetNumber not mocked")
	return
}

var HostCPUClassGetNumberMockedCallback = HostCPUClassGetNumberMockDefault

func (_class HostCPUClass) GetNumberMock(sessionID SessionRef, self HostCPURef) (_retval int, _err error) {
	return HostCPUClassGetNumberMockedCallback(sessionID, self)
}
// Get the number field of the given host_cpu.
func (_class HostCPUClass) GetNumber(sessionID SessionRef, self HostCPURef) (_retval int, _err error) {
	if IsMock {
		return _class.GetNumberMock(sessionID, self)
	}	
	_method := "host_cpu.get_number"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertHostCPURefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertIntToGo(_method + " -> ", _result.Value)
	return
}


func HostCPUClassGetHostMockDefault(sessionID SessionRef, self HostCPURef) (_retval HostRef, _err error) {
	log.Println("HostCPU.GetHost not mocked")
	_err = errors.New("HostCPU.GetHost not mocked")
	return
}

var HostCPUClassGetHostMockedCallback = HostCPUClassGetHostMockDefault

func (_class HostCPUClass) GetHostMock(sessionID SessionRef, self HostCPURef) (_retval HostRef, _err error) {
	return HostCPUClassGetHostMockedCallback(sessionID, self)
}
// Get the host field of the given host_cpu.
func (_class HostCPUClass) GetHost(sessionID SessionRef, self HostCPURef) (_retval HostRef, _err error) {
	if IsMock {
		return _class.GetHostMock(sessionID, self)
	}	
	_method := "host_cpu.get_host"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertHostCPURefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func HostCPUClassGetUUIDMockDefault(sessionID SessionRef, self HostCPURef) (_retval string, _err error) {
	log.Println("HostCPU.GetUUID not mocked")
	_err = errors.New("HostCPU.GetUUID not mocked")
	return
}

var HostCPUClassGetUUIDMockedCallback = HostCPUClassGetUUIDMockDefault

func (_class HostCPUClass) GetUUIDMock(sessionID SessionRef, self HostCPURef) (_retval string, _err error) {
	return HostCPUClassGetUUIDMockedCallback(sessionID, self)
}
// Get the uuid field of the given host_cpu.
func (_class HostCPUClass) GetUUID(sessionID SessionRef, self HostCPURef) (_retval string, _err error) {
	if IsMock {
		return _class.GetUUIDMock(sessionID, self)
	}	
	_method := "host_cpu.get_uuid"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertHostCPURefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func HostCPUClassGetByUUIDMockDefault(sessionID SessionRef, uuid string) (_retval HostCPURef, _err error) {
	log.Println("HostCPU.GetByUUID not mocked")
	_err = errors.New("HostCPU.GetByUUID not mocked")
	return
}

var HostCPUClassGetByUUIDMockedCallback = HostCPUClassGetByUUIDMockDefault

func (_class HostCPUClass) GetByUUIDMock(sessionID SessionRef, uuid string) (_retval HostCPURef, _err error) {
	return HostCPUClassGetByUUIDMockedCallback(sessionID, uuid)
}
// Get a reference to the host_cpu instance with the specified UUID.
func (_class HostCPUClass) GetByUUID(sessionID SessionRef, uuid string) (_retval HostCPURef, _err error) {
	if IsMock {
		return _class.GetByUUIDMock(sessionID, uuid)
	}	
	_method := "host_cpu.get_by_uuid"
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
	_retval, _err = convertHostCPURefToGo(_method + " -> ", _result.Value)
	return
}


func HostCPUClassGetRecordMockDefault(sessionID SessionRef, self HostCPURef) (_retval HostCPURecord, _err error) {
	log.Println("HostCPU.GetRecord not mocked")
	_err = errors.New("HostCPU.GetRecord not mocked")
	return
}

var HostCPUClassGetRecordMockedCallback = HostCPUClassGetRecordMockDefault

func (_class HostCPUClass) GetRecordMock(sessionID SessionRef, self HostCPURef) (_retval HostCPURecord, _err error) {
	return HostCPUClassGetRecordMockedCallback(sessionID, self)
}
// Get a record containing the current state of the given host_cpu.
func (_class HostCPUClass) GetRecord(sessionID SessionRef, self HostCPURef) (_retval HostCPURecord, _err error) {
	if IsMock {
		return _class.GetRecordMock(sessionID, self)
	}	
	_method := "host_cpu.get_record"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertHostCPURefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertHostCPURecordToGo(_method + " -> ", _result.Value)
	return
}
