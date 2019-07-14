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

type VMMetricsRecord struct {
  // Unique identifier/object reference
	UUID string
  // Guest's actual memory (bytes)
	MemoryActual int
  // Current number of VCPUs
	VCPUsNumber int
  // Utilisation for all of guest's current VCPUs
	VCPUsUtilisation map[int]float64
  // VCPU to PCPU map
	VCPUsCPU map[int]int
  // The live equivalent to VM.VCPUs_params
	VCPUsParams map[string]string
  // CPU flags (blocked,online,running)
	VCPUsFlags map[int][]string
  // The state of the guest, eg blocked, dying etc
	State []string
  // Time at which this VM was last booted
	StartTime time.Time
  // Time at which the VM was installed
	InstallTime time.Time
  // Time at which this information was last updated
	LastUpdated time.Time
  // additional configuration
	OtherConfig map[string]string
  // hardware virtual machine
	Hvm bool
  // VM supports nested virtualisation
	NestedVirt bool
  // VM is immobile and can't migrate between hosts
	Nomigrate bool
}

type VMMetricsRef string

// The metrics associated with a VM
type VMMetricsClass struct {
	client *Client
}


var VMMetricsClass_GetAllRecordsMockedCallback = func (sessionID SessionRef) (_retval map[VMMetricsRef]VMMetricsRecord, _err error) {
	log.Println("VMMetrics.GetAllRecords not mocked")
	_err = errors.New("VMMetrics.GetAllRecords not mocked")
	return
}

func (_class VMMetricsClass) GetAllRecordsMock(sessionID SessionRef) (_retval map[VMMetricsRef]VMMetricsRecord, _err error) {
	return VMMetricsClass_GetAllRecordsMockedCallback(sessionID)
}
// Return a map of VM_metrics references to VM_metrics records for all VM_metrics instances known to the system.
func (_class VMMetricsClass) GetAllRecords(sessionID SessionRef) (_retval map[VMMetricsRef]VMMetricsRecord, _err error) {
	if (IsMock) {
		return _class.GetAllRecordsMock(sessionID)
	}	
	_method := "VM_metrics.get_all_records"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVMMetricsRefToVMMetricsRecordMapToGo(_method + " -> ", _result.Value)
	return
}


var VMMetricsClass_GetAllMockedCallback = func (sessionID SessionRef) (_retval []VMMetricsRef, _err error) {
	log.Println("VMMetrics.GetAll not mocked")
	_err = errors.New("VMMetrics.GetAll not mocked")
	return
}

func (_class VMMetricsClass) GetAllMock(sessionID SessionRef) (_retval []VMMetricsRef, _err error) {
	return VMMetricsClass_GetAllMockedCallback(sessionID)
}
// Return a list of all the VM_metrics instances known to the system.
func (_class VMMetricsClass) GetAll(sessionID SessionRef) (_retval []VMMetricsRef, _err error) {
	if (IsMock) {
		return _class.GetAllMock(sessionID)
	}	
	_method := "VM_metrics.get_all"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVMMetricsRefSetToGo(_method + " -> ", _result.Value)
	return
}


var VMMetricsClass_RemoveFromOtherConfigMockedCallback = func (sessionID SessionRef, self VMMetricsRef, key string) (_err error) {
	log.Println("VMMetrics.RemoveFromOtherConfig not mocked")
	_err = errors.New("VMMetrics.RemoveFromOtherConfig not mocked")
	return
}

func (_class VMMetricsClass) RemoveFromOtherConfigMock(sessionID SessionRef, self VMMetricsRef, key string) (_err error) {
	return VMMetricsClass_RemoveFromOtherConfigMockedCallback(sessionID, self, key)
}
// Remove the given key and its corresponding value from the other_config field of the given VM_metrics.  If the key is not in that Map, then do nothing.
func (_class VMMetricsClass) RemoveFromOtherConfig(sessionID SessionRef, self VMMetricsRef, key string) (_err error) {
	if (IsMock) {
		return _class.RemoveFromOtherConfigMock(sessionID, self, key)
	}	
	_method := "VM_metrics.remove_from_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMMetricsRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var VMMetricsClass_AddToOtherConfigMockedCallback = func (sessionID SessionRef, self VMMetricsRef, key string, value string) (_err error) {
	log.Println("VMMetrics.AddToOtherConfig not mocked")
	_err = errors.New("VMMetrics.AddToOtherConfig not mocked")
	return
}

func (_class VMMetricsClass) AddToOtherConfigMock(sessionID SessionRef, self VMMetricsRef, key string, value string) (_err error) {
	return VMMetricsClass_AddToOtherConfigMockedCallback(sessionID, self, key, value)
}
// Add the given key-value pair to the other_config field of the given VM_metrics.
func (_class VMMetricsClass) AddToOtherConfig(sessionID SessionRef, self VMMetricsRef, key string, value string) (_err error) {
	if (IsMock) {
		return _class.AddToOtherConfigMock(sessionID, self, key, value)
	}	
	_method := "VM_metrics.add_to_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMMetricsRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var VMMetricsClass_SetOtherConfigMockedCallback = func (sessionID SessionRef, self VMMetricsRef, value map[string]string) (_err error) {
	log.Println("VMMetrics.SetOtherConfig not mocked")
	_err = errors.New("VMMetrics.SetOtherConfig not mocked")
	return
}

func (_class VMMetricsClass) SetOtherConfigMock(sessionID SessionRef, self VMMetricsRef, value map[string]string) (_err error) {
	return VMMetricsClass_SetOtherConfigMockedCallback(sessionID, self, value)
}
// Set the other_config field of the given VM_metrics.
func (_class VMMetricsClass) SetOtherConfig(sessionID SessionRef, self VMMetricsRef, value map[string]string) (_err error) {
	if (IsMock) {
		return _class.SetOtherConfigMock(sessionID, self, value)
	}	
	_method := "VM_metrics.set_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMMetricsRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var VMMetricsClass_GetNomigrateMockedCallback = func (sessionID SessionRef, self VMMetricsRef) (_retval bool, _err error) {
	log.Println("VMMetrics.GetNomigrate not mocked")
	_err = errors.New("VMMetrics.GetNomigrate not mocked")
	return
}

func (_class VMMetricsClass) GetNomigrateMock(sessionID SessionRef, self VMMetricsRef) (_retval bool, _err error) {
	return VMMetricsClass_GetNomigrateMockedCallback(sessionID, self)
}
// Get the nomigrate field of the given VM_metrics.
func (_class VMMetricsClass) GetNomigrate(sessionID SessionRef, self VMMetricsRef) (_retval bool, _err error) {
	if (IsMock) {
		return _class.GetNomigrateMock(sessionID, self)
	}	
	_method := "VM_metrics.get_nomigrate"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMMetricsRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var VMMetricsClass_GetNestedVirtMockedCallback = func (sessionID SessionRef, self VMMetricsRef) (_retval bool, _err error) {
	log.Println("VMMetrics.GetNestedVirt not mocked")
	_err = errors.New("VMMetrics.GetNestedVirt not mocked")
	return
}

func (_class VMMetricsClass) GetNestedVirtMock(sessionID SessionRef, self VMMetricsRef) (_retval bool, _err error) {
	return VMMetricsClass_GetNestedVirtMockedCallback(sessionID, self)
}
// Get the nested_virt field of the given VM_metrics.
func (_class VMMetricsClass) GetNestedVirt(sessionID SessionRef, self VMMetricsRef) (_retval bool, _err error) {
	if (IsMock) {
		return _class.GetNestedVirtMock(sessionID, self)
	}	
	_method := "VM_metrics.get_nested_virt"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMMetricsRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var VMMetricsClass_GetHvmMockedCallback = func (sessionID SessionRef, self VMMetricsRef) (_retval bool, _err error) {
	log.Println("VMMetrics.GetHvm not mocked")
	_err = errors.New("VMMetrics.GetHvm not mocked")
	return
}

func (_class VMMetricsClass) GetHvmMock(sessionID SessionRef, self VMMetricsRef) (_retval bool, _err error) {
	return VMMetricsClass_GetHvmMockedCallback(sessionID, self)
}
// Get the hvm field of the given VM_metrics.
func (_class VMMetricsClass) GetHvm(sessionID SessionRef, self VMMetricsRef) (_retval bool, _err error) {
	if (IsMock) {
		return _class.GetHvmMock(sessionID, self)
	}	
	_method := "VM_metrics.get_hvm"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMMetricsRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var VMMetricsClass_GetOtherConfigMockedCallback = func (sessionID SessionRef, self VMMetricsRef) (_retval map[string]string, _err error) {
	log.Println("VMMetrics.GetOtherConfig not mocked")
	_err = errors.New("VMMetrics.GetOtherConfig not mocked")
	return
}

func (_class VMMetricsClass) GetOtherConfigMock(sessionID SessionRef, self VMMetricsRef) (_retval map[string]string, _err error) {
	return VMMetricsClass_GetOtherConfigMockedCallback(sessionID, self)
}
// Get the other_config field of the given VM_metrics.
func (_class VMMetricsClass) GetOtherConfig(sessionID SessionRef, self VMMetricsRef) (_retval map[string]string, _err error) {
	if (IsMock) {
		return _class.GetOtherConfigMock(sessionID, self)
	}	
	_method := "VM_metrics.get_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMMetricsRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var VMMetricsClass_GetLastUpdatedMockedCallback = func (sessionID SessionRef, self VMMetricsRef) (_retval time.Time, _err error) {
	log.Println("VMMetrics.GetLastUpdated not mocked")
	_err = errors.New("VMMetrics.GetLastUpdated not mocked")
	return
}

func (_class VMMetricsClass) GetLastUpdatedMock(sessionID SessionRef, self VMMetricsRef) (_retval time.Time, _err error) {
	return VMMetricsClass_GetLastUpdatedMockedCallback(sessionID, self)
}
// Get the last_updated field of the given VM_metrics.
func (_class VMMetricsClass) GetLastUpdated(sessionID SessionRef, self VMMetricsRef) (_retval time.Time, _err error) {
	if (IsMock) {
		return _class.GetLastUpdatedMock(sessionID, self)
	}	
	_method := "VM_metrics.get_last_updated"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMMetricsRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var VMMetricsClass_GetInstallTimeMockedCallback = func (sessionID SessionRef, self VMMetricsRef) (_retval time.Time, _err error) {
	log.Println("VMMetrics.GetInstallTime not mocked")
	_err = errors.New("VMMetrics.GetInstallTime not mocked")
	return
}

func (_class VMMetricsClass) GetInstallTimeMock(sessionID SessionRef, self VMMetricsRef) (_retval time.Time, _err error) {
	return VMMetricsClass_GetInstallTimeMockedCallback(sessionID, self)
}
// Get the install_time field of the given VM_metrics.
func (_class VMMetricsClass) GetInstallTime(sessionID SessionRef, self VMMetricsRef) (_retval time.Time, _err error) {
	if (IsMock) {
		return _class.GetInstallTimeMock(sessionID, self)
	}	
	_method := "VM_metrics.get_install_time"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMMetricsRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var VMMetricsClass_GetStartTimeMockedCallback = func (sessionID SessionRef, self VMMetricsRef) (_retval time.Time, _err error) {
	log.Println("VMMetrics.GetStartTime not mocked")
	_err = errors.New("VMMetrics.GetStartTime not mocked")
	return
}

func (_class VMMetricsClass) GetStartTimeMock(sessionID SessionRef, self VMMetricsRef) (_retval time.Time, _err error) {
	return VMMetricsClass_GetStartTimeMockedCallback(sessionID, self)
}
// Get the start_time field of the given VM_metrics.
func (_class VMMetricsClass) GetStartTime(sessionID SessionRef, self VMMetricsRef) (_retval time.Time, _err error) {
	if (IsMock) {
		return _class.GetStartTimeMock(sessionID, self)
	}	
	_method := "VM_metrics.get_start_time"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMMetricsRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var VMMetricsClass_GetStateMockedCallback = func (sessionID SessionRef, self VMMetricsRef) (_retval []string, _err error) {
	log.Println("VMMetrics.GetState not mocked")
	_err = errors.New("VMMetrics.GetState not mocked")
	return
}

func (_class VMMetricsClass) GetStateMock(sessionID SessionRef, self VMMetricsRef) (_retval []string, _err error) {
	return VMMetricsClass_GetStateMockedCallback(sessionID, self)
}
// Get the state field of the given VM_metrics.
func (_class VMMetricsClass) GetState(sessionID SessionRef, self VMMetricsRef) (_retval []string, _err error) {
	if (IsMock) {
		return _class.GetStateMock(sessionID, self)
	}	
	_method := "VM_metrics.get_state"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMMetricsRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var VMMetricsClass_GetVCPUsFlagsMockedCallback = func (sessionID SessionRef, self VMMetricsRef) (_retval map[int][]string, _err error) {
	log.Println("VMMetrics.GetVCPUsFlags not mocked")
	_err = errors.New("VMMetrics.GetVCPUsFlags not mocked")
	return
}

func (_class VMMetricsClass) GetVCPUsFlagsMock(sessionID SessionRef, self VMMetricsRef) (_retval map[int][]string, _err error) {
	return VMMetricsClass_GetVCPUsFlagsMockedCallback(sessionID, self)
}
// Get the VCPUs/flags field of the given VM_metrics.
func (_class VMMetricsClass) GetVCPUsFlags(sessionID SessionRef, self VMMetricsRef) (_retval map[int][]string, _err error) {
	if (IsMock) {
		return _class.GetVCPUsFlagsMock(sessionID, self)
	}	
	_method := "VM_metrics.get_VCPUs_flags"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMMetricsRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertIntToStringSetMapToGo(_method + " -> ", _result.Value)
	return
}


var VMMetricsClass_GetVCPUsParamsMockedCallback = func (sessionID SessionRef, self VMMetricsRef) (_retval map[string]string, _err error) {
	log.Println("VMMetrics.GetVCPUsParams not mocked")
	_err = errors.New("VMMetrics.GetVCPUsParams not mocked")
	return
}

func (_class VMMetricsClass) GetVCPUsParamsMock(sessionID SessionRef, self VMMetricsRef) (_retval map[string]string, _err error) {
	return VMMetricsClass_GetVCPUsParamsMockedCallback(sessionID, self)
}
// Get the VCPUs/params field of the given VM_metrics.
func (_class VMMetricsClass) GetVCPUsParams(sessionID SessionRef, self VMMetricsRef) (_retval map[string]string, _err error) {
	if (IsMock) {
		return _class.GetVCPUsParamsMock(sessionID, self)
	}	
	_method := "VM_metrics.get_VCPUs_params"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMMetricsRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var VMMetricsClass_GetVCPUsCPUMockedCallback = func (sessionID SessionRef, self VMMetricsRef) (_retval map[int]int, _err error) {
	log.Println("VMMetrics.GetVCPUsCPU not mocked")
	_err = errors.New("VMMetrics.GetVCPUsCPU not mocked")
	return
}

func (_class VMMetricsClass) GetVCPUsCPUMock(sessionID SessionRef, self VMMetricsRef) (_retval map[int]int, _err error) {
	return VMMetricsClass_GetVCPUsCPUMockedCallback(sessionID, self)
}
// Get the VCPUs/CPU field of the given VM_metrics.
func (_class VMMetricsClass) GetVCPUsCPU(sessionID SessionRef, self VMMetricsRef) (_retval map[int]int, _err error) {
	if (IsMock) {
		return _class.GetVCPUsCPUMock(sessionID, self)
	}	
	_method := "VM_metrics.get_VCPUs_CPU"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMMetricsRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertIntToIntMapToGo(_method + " -> ", _result.Value)
	return
}


var VMMetricsClass_GetVCPUsUtilisationMockedCallback = func (sessionID SessionRef, self VMMetricsRef) (_retval map[int]float64, _err error) {
	log.Println("VMMetrics.GetVCPUsUtilisation not mocked")
	_err = errors.New("VMMetrics.GetVCPUsUtilisation not mocked")
	return
}

func (_class VMMetricsClass) GetVCPUsUtilisationMock(sessionID SessionRef, self VMMetricsRef) (_retval map[int]float64, _err error) {
	return VMMetricsClass_GetVCPUsUtilisationMockedCallback(sessionID, self)
}
// Get the VCPUs/utilisation field of the given VM_metrics.
func (_class VMMetricsClass) GetVCPUsUtilisation(sessionID SessionRef, self VMMetricsRef) (_retval map[int]float64, _err error) {
	if (IsMock) {
		return _class.GetVCPUsUtilisationMock(sessionID, self)
	}	
	_method := "VM_metrics.get_VCPUs_utilisation"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMMetricsRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertIntToFloatMapToGo(_method + " -> ", _result.Value)
	return
}


var VMMetricsClass_GetVCPUsNumberMockedCallback = func (sessionID SessionRef, self VMMetricsRef) (_retval int, _err error) {
	log.Println("VMMetrics.GetVCPUsNumber not mocked")
	_err = errors.New("VMMetrics.GetVCPUsNumber not mocked")
	return
}

func (_class VMMetricsClass) GetVCPUsNumberMock(sessionID SessionRef, self VMMetricsRef) (_retval int, _err error) {
	return VMMetricsClass_GetVCPUsNumberMockedCallback(sessionID, self)
}
// Get the VCPUs/number field of the given VM_metrics.
func (_class VMMetricsClass) GetVCPUsNumber(sessionID SessionRef, self VMMetricsRef) (_retval int, _err error) {
	if (IsMock) {
		return _class.GetVCPUsNumberMock(sessionID, self)
	}	
	_method := "VM_metrics.get_VCPUs_number"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMMetricsRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var VMMetricsClass_GetMemoryActualMockedCallback = func (sessionID SessionRef, self VMMetricsRef) (_retval int, _err error) {
	log.Println("VMMetrics.GetMemoryActual not mocked")
	_err = errors.New("VMMetrics.GetMemoryActual not mocked")
	return
}

func (_class VMMetricsClass) GetMemoryActualMock(sessionID SessionRef, self VMMetricsRef) (_retval int, _err error) {
	return VMMetricsClass_GetMemoryActualMockedCallback(sessionID, self)
}
// Get the memory/actual field of the given VM_metrics.
func (_class VMMetricsClass) GetMemoryActual(sessionID SessionRef, self VMMetricsRef) (_retval int, _err error) {
	if (IsMock) {
		return _class.GetMemoryActualMock(sessionID, self)
	}	
	_method := "VM_metrics.get_memory_actual"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMMetricsRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var VMMetricsClass_GetUUIDMockedCallback = func (sessionID SessionRef, self VMMetricsRef) (_retval string, _err error) {
	log.Println("VMMetrics.GetUUID not mocked")
	_err = errors.New("VMMetrics.GetUUID not mocked")
	return
}

func (_class VMMetricsClass) GetUUIDMock(sessionID SessionRef, self VMMetricsRef) (_retval string, _err error) {
	return VMMetricsClass_GetUUIDMockedCallback(sessionID, self)
}
// Get the uuid field of the given VM_metrics.
func (_class VMMetricsClass) GetUUID(sessionID SessionRef, self VMMetricsRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetUUIDMock(sessionID, self)
	}	
	_method := "VM_metrics.get_uuid"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMMetricsRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var VMMetricsClass_GetByUUIDMockedCallback = func (sessionID SessionRef, uuid string) (_retval VMMetricsRef, _err error) {
	log.Println("VMMetrics.GetByUUID not mocked")
	_err = errors.New("VMMetrics.GetByUUID not mocked")
	return
}

func (_class VMMetricsClass) GetByUUIDMock(sessionID SessionRef, uuid string) (_retval VMMetricsRef, _err error) {
	return VMMetricsClass_GetByUUIDMockedCallback(sessionID, uuid)
}
// Get a reference to the VM_metrics instance with the specified UUID.
func (_class VMMetricsClass) GetByUUID(sessionID SessionRef, uuid string) (_retval VMMetricsRef, _err error) {
	if (IsMock) {
		return _class.GetByUUIDMock(sessionID, uuid)
	}	
	_method := "VM_metrics.get_by_uuid"
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
	_retval, _err = convertVMMetricsRefToGo(_method + " -> ", _result.Value)
	return
}


var VMMetricsClass_GetRecordMockedCallback = func (sessionID SessionRef, self VMMetricsRef) (_retval VMMetricsRecord, _err error) {
	log.Println("VMMetrics.GetRecord not mocked")
	_err = errors.New("VMMetrics.GetRecord not mocked")
	return
}

func (_class VMMetricsClass) GetRecordMock(sessionID SessionRef, self VMMetricsRef) (_retval VMMetricsRecord, _err error) {
	return VMMetricsClass_GetRecordMockedCallback(sessionID, self)
}
// Get a record containing the current state of the given VM_metrics.
func (_class VMMetricsClass) GetRecord(sessionID SessionRef, self VMMetricsRef) (_retval VMMetricsRecord, _err error) {
	if (IsMock) {
		return _class.GetRecordMock(sessionID, self)
	}	
	_method := "VM_metrics.get_record"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMMetricsRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVMMetricsRecordToGo(_method + " -> ", _result.Value)
	return
}
