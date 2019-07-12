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

func (_class HostCPUClass) GetAllRecords__mock(sessionID SessionRef) (_retval map[HostCPURef]HostCPURecord, _err error) {
	log.Println("HostCPU.GetAllRecords not mocked")
	_err = errors.New("HostCPU.GetAllRecords not mocked")
	return
}
// Return a map of host_cpu references to host_cpu records for all host_cpus known to the system.
func (_class HostCPUClass) GetAllRecords(sessionID SessionRef) (_retval map[HostCPURef]HostCPURecord, _err error) {
	if (IsMock) {
		return _class.GetAllRecords__mock(sessionID)
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

func (_class HostCPUClass) GetAll__mock(sessionID SessionRef) (_retval []HostCPURef, _err error) {
	log.Println("HostCPU.GetAll not mocked")
	_err = errors.New("HostCPU.GetAll not mocked")
	return
}
// Return a list of all the host_cpus known to the system.
func (_class HostCPUClass) GetAll(sessionID SessionRef) (_retval []HostCPURef, _err error) {
	if (IsMock) {
		return _class.GetAll__mock(sessionID)
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

func (_class HostCPUClass) RemoveFromOtherConfig__mock(sessionID SessionRef, self HostCPURef, key string) (_err error) {
	log.Println("HostCPU.RemoveFromOtherConfig not mocked")
	_err = errors.New("HostCPU.RemoveFromOtherConfig not mocked")
	return
}
// Remove the given key and its corresponding value from the other_config field of the given host_cpu.  If the key is not in that Map, then do nothing.
func (_class HostCPUClass) RemoveFromOtherConfig(sessionID SessionRef, self HostCPURef, key string) (_err error) {
	if (IsMock) {
		return _class.RemoveFromOtherConfig__mock(sessionID, self, key)
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

func (_class HostCPUClass) AddToOtherConfig__mock(sessionID SessionRef, self HostCPURef, key string, value string) (_err error) {
	log.Println("HostCPU.AddToOtherConfig not mocked")
	_err = errors.New("HostCPU.AddToOtherConfig not mocked")
	return
}
// Add the given key-value pair to the other_config field of the given host_cpu.
func (_class HostCPUClass) AddToOtherConfig(sessionID SessionRef, self HostCPURef, key string, value string) (_err error) {
	if (IsMock) {
		return _class.AddToOtherConfig__mock(sessionID, self, key, value)
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

func (_class HostCPUClass) SetOtherConfig__mock(sessionID SessionRef, self HostCPURef, value map[string]string) (_err error) {
	log.Println("HostCPU.SetOtherConfig not mocked")
	_err = errors.New("HostCPU.SetOtherConfig not mocked")
	return
}
// Set the other_config field of the given host_cpu.
func (_class HostCPUClass) SetOtherConfig(sessionID SessionRef, self HostCPURef, value map[string]string) (_err error) {
	if (IsMock) {
		return _class.SetOtherConfig__mock(sessionID, self, value)
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

func (_class HostCPUClass) GetOtherConfig__mock(sessionID SessionRef, self HostCPURef) (_retval map[string]string, _err error) {
	log.Println("HostCPU.GetOtherConfig not mocked")
	_err = errors.New("HostCPU.GetOtherConfig not mocked")
	return
}
// Get the other_config field of the given host_cpu.
func (_class HostCPUClass) GetOtherConfig(sessionID SessionRef, self HostCPURef) (_retval map[string]string, _err error) {
	if (IsMock) {
		return _class.GetOtherConfig__mock(sessionID, self)
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

func (_class HostCPUClass) GetUtilisation__mock(sessionID SessionRef, self HostCPURef) (_retval float64, _err error) {
	log.Println("HostCPU.GetUtilisation not mocked")
	_err = errors.New("HostCPU.GetUtilisation not mocked")
	return
}
// Get the utilisation field of the given host_cpu.
func (_class HostCPUClass) GetUtilisation(sessionID SessionRef, self HostCPURef) (_retval float64, _err error) {
	if (IsMock) {
		return _class.GetUtilisation__mock(sessionID, self)
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

func (_class HostCPUClass) GetFeatures__mock(sessionID SessionRef, self HostCPURef) (_retval string, _err error) {
	log.Println("HostCPU.GetFeatures not mocked")
	_err = errors.New("HostCPU.GetFeatures not mocked")
	return
}
// Get the features field of the given host_cpu.
func (_class HostCPUClass) GetFeatures(sessionID SessionRef, self HostCPURef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetFeatures__mock(sessionID, self)
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

func (_class HostCPUClass) GetFlags__mock(sessionID SessionRef, self HostCPURef) (_retval string, _err error) {
	log.Println("HostCPU.GetFlags not mocked")
	_err = errors.New("HostCPU.GetFlags not mocked")
	return
}
// Get the flags field of the given host_cpu.
func (_class HostCPUClass) GetFlags(sessionID SessionRef, self HostCPURef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetFlags__mock(sessionID, self)
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

func (_class HostCPUClass) GetStepping__mock(sessionID SessionRef, self HostCPURef) (_retval string, _err error) {
	log.Println("HostCPU.GetStepping not mocked")
	_err = errors.New("HostCPU.GetStepping not mocked")
	return
}
// Get the stepping field of the given host_cpu.
func (_class HostCPUClass) GetStepping(sessionID SessionRef, self HostCPURef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetStepping__mock(sessionID, self)
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

func (_class HostCPUClass) GetModel__mock(sessionID SessionRef, self HostCPURef) (_retval int, _err error) {
	log.Println("HostCPU.GetModel not mocked")
	_err = errors.New("HostCPU.GetModel not mocked")
	return
}
// Get the model field of the given host_cpu.
func (_class HostCPUClass) GetModel(sessionID SessionRef, self HostCPURef) (_retval int, _err error) {
	if (IsMock) {
		return _class.GetModel__mock(sessionID, self)
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

func (_class HostCPUClass) GetFamily__mock(sessionID SessionRef, self HostCPURef) (_retval int, _err error) {
	log.Println("HostCPU.GetFamily not mocked")
	_err = errors.New("HostCPU.GetFamily not mocked")
	return
}
// Get the family field of the given host_cpu.
func (_class HostCPUClass) GetFamily(sessionID SessionRef, self HostCPURef) (_retval int, _err error) {
	if (IsMock) {
		return _class.GetFamily__mock(sessionID, self)
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

func (_class HostCPUClass) GetModelname__mock(sessionID SessionRef, self HostCPURef) (_retval string, _err error) {
	log.Println("HostCPU.GetModelname not mocked")
	_err = errors.New("HostCPU.GetModelname not mocked")
	return
}
// Get the modelname field of the given host_cpu.
func (_class HostCPUClass) GetModelname(sessionID SessionRef, self HostCPURef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetModelname__mock(sessionID, self)
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

func (_class HostCPUClass) GetSpeed__mock(sessionID SessionRef, self HostCPURef) (_retval int, _err error) {
	log.Println("HostCPU.GetSpeed not mocked")
	_err = errors.New("HostCPU.GetSpeed not mocked")
	return
}
// Get the speed field of the given host_cpu.
func (_class HostCPUClass) GetSpeed(sessionID SessionRef, self HostCPURef) (_retval int, _err error) {
	if (IsMock) {
		return _class.GetSpeed__mock(sessionID, self)
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

func (_class HostCPUClass) GetVendor__mock(sessionID SessionRef, self HostCPURef) (_retval string, _err error) {
	log.Println("HostCPU.GetVendor not mocked")
	_err = errors.New("HostCPU.GetVendor not mocked")
	return
}
// Get the vendor field of the given host_cpu.
func (_class HostCPUClass) GetVendor(sessionID SessionRef, self HostCPURef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetVendor__mock(sessionID, self)
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

func (_class HostCPUClass) GetNumber__mock(sessionID SessionRef, self HostCPURef) (_retval int, _err error) {
	log.Println("HostCPU.GetNumber not mocked")
	_err = errors.New("HostCPU.GetNumber not mocked")
	return
}
// Get the number field of the given host_cpu.
func (_class HostCPUClass) GetNumber(sessionID SessionRef, self HostCPURef) (_retval int, _err error) {
	if (IsMock) {
		return _class.GetNumber__mock(sessionID, self)
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

func (_class HostCPUClass) GetHost__mock(sessionID SessionRef, self HostCPURef) (_retval HostRef, _err error) {
	log.Println("HostCPU.GetHost not mocked")
	_err = errors.New("HostCPU.GetHost not mocked")
	return
}
// Get the host field of the given host_cpu.
func (_class HostCPUClass) GetHost(sessionID SessionRef, self HostCPURef) (_retval HostRef, _err error) {
	if (IsMock) {
		return _class.GetHost__mock(sessionID, self)
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

func (_class HostCPUClass) GetUUID__mock(sessionID SessionRef, self HostCPURef) (_retval string, _err error) {
	log.Println("HostCPU.GetUUID not mocked")
	_err = errors.New("HostCPU.GetUUID not mocked")
	return
}
// Get the uuid field of the given host_cpu.
func (_class HostCPUClass) GetUUID(sessionID SessionRef, self HostCPURef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetUUID__mock(sessionID, self)
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

func (_class HostCPUClass) GetByUUID__mock(sessionID SessionRef, uuid string) (_retval HostCPURef, _err error) {
	log.Println("HostCPU.GetByUUID not mocked")
	_err = errors.New("HostCPU.GetByUUID not mocked")
	return
}
// Get a reference to the host_cpu instance with the specified UUID.
func (_class HostCPUClass) GetByUUID(sessionID SessionRef, uuid string) (_retval HostCPURef, _err error) {
	if (IsMock) {
		return _class.GetByUUID__mock(sessionID, uuid)
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

func (_class HostCPUClass) GetRecord__mock(sessionID SessionRef, self HostCPURef) (_retval HostCPURecord, _err error) {
	log.Println("HostCPU.GetRecord not mocked")
	_err = errors.New("HostCPU.GetRecord not mocked")
	return
}
// Get a record containing the current state of the given host_cpu.
func (_class HostCPUClass) GetRecord(sessionID SessionRef, self HostCPURef) (_retval HostCPURecord, _err error) {
	if (IsMock) {
		return _class.GetRecord__mock(sessionID, self)
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
