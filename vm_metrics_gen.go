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


func VMMetricsClassGetAllRecordsMockDefault(sessionID SessionRef) (_retval map[VMMetricsRef]VMMetricsRecord, _err error) {
	log.Println("VMMetrics.GetAllRecords not mocked")
	_err = errors.New("VMMetrics.GetAllRecords not mocked")
	return
}

var VMMetricsClassGetAllRecordsMockedCallback = VMMetricsClassGetAllRecordsMockDefault

func (_class VMMetricsClass) GetAllRecordsMock(sessionID SessionRef) (_retval map[VMMetricsRef]VMMetricsRecord, _err error) {
	return VMMetricsClassGetAllRecordsMockedCallback(sessionID)
}
// Return a map of VM_metrics references to VM_metrics records for all VM_metrics instances known to the system.
func (_class VMMetricsClass) GetAllRecords(sessionID SessionRef) (_retval map[VMMetricsRef]VMMetricsRecord, _err error) {
	if IsMock {
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


func VMMetricsClassGetAllMockDefault(sessionID SessionRef) (_retval []VMMetricsRef, _err error) {
	log.Println("VMMetrics.GetAll not mocked")
	_err = errors.New("VMMetrics.GetAll not mocked")
	return
}

var VMMetricsClassGetAllMockedCallback = VMMetricsClassGetAllMockDefault

func (_class VMMetricsClass) GetAllMock(sessionID SessionRef) (_retval []VMMetricsRef, _err error) {
	return VMMetricsClassGetAllMockedCallback(sessionID)
}
// Return a list of all the VM_metrics instances known to the system.
func (_class VMMetricsClass) GetAll(sessionID SessionRef) (_retval []VMMetricsRef, _err error) {
	if IsMock {
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


func VMMetricsClassRemoveFromOtherConfigMockDefault(sessionID SessionRef, self VMMetricsRef, key string) (_err error) {
	log.Println("VMMetrics.RemoveFromOtherConfig not mocked")
	_err = errors.New("VMMetrics.RemoveFromOtherConfig not mocked")
	return
}

var VMMetricsClassRemoveFromOtherConfigMockedCallback = VMMetricsClassRemoveFromOtherConfigMockDefault

func (_class VMMetricsClass) RemoveFromOtherConfigMock(sessionID SessionRef, self VMMetricsRef, key string) (_err error) {
	return VMMetricsClassRemoveFromOtherConfigMockedCallback(sessionID, self, key)
}
// Remove the given key and its corresponding value from the other_config field of the given VM_metrics.  If the key is not in that Map, then do nothing.
func (_class VMMetricsClass) RemoveFromOtherConfig(sessionID SessionRef, self VMMetricsRef, key string) (_err error) {
	if IsMock {
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


func VMMetricsClassAddToOtherConfigMockDefault(sessionID SessionRef, self VMMetricsRef, key string, value string) (_err error) {
	log.Println("VMMetrics.AddToOtherConfig not mocked")
	_err = errors.New("VMMetrics.AddToOtherConfig not mocked")
	return
}

var VMMetricsClassAddToOtherConfigMockedCallback = VMMetricsClassAddToOtherConfigMockDefault

func (_class VMMetricsClass) AddToOtherConfigMock(sessionID SessionRef, self VMMetricsRef, key string, value string) (_err error) {
	return VMMetricsClassAddToOtherConfigMockedCallback(sessionID, self, key, value)
}
// Add the given key-value pair to the other_config field of the given VM_metrics.
func (_class VMMetricsClass) AddToOtherConfig(sessionID SessionRef, self VMMetricsRef, key string, value string) (_err error) {
	if IsMock {
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


func VMMetricsClassSetOtherConfigMockDefault(sessionID SessionRef, self VMMetricsRef, value map[string]string) (_err error) {
	log.Println("VMMetrics.SetOtherConfig not mocked")
	_err = errors.New("VMMetrics.SetOtherConfig not mocked")
	return
}

var VMMetricsClassSetOtherConfigMockedCallback = VMMetricsClassSetOtherConfigMockDefault

func (_class VMMetricsClass) SetOtherConfigMock(sessionID SessionRef, self VMMetricsRef, value map[string]string) (_err error) {
	return VMMetricsClassSetOtherConfigMockedCallback(sessionID, self, value)
}
// Set the other_config field of the given VM_metrics.
func (_class VMMetricsClass) SetOtherConfig(sessionID SessionRef, self VMMetricsRef, value map[string]string) (_err error) {
	if IsMock {
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


func VMMetricsClassGetNomigrateMockDefault(sessionID SessionRef, self VMMetricsRef) (_retval bool, _err error) {
	log.Println("VMMetrics.GetNomigrate not mocked")
	_err = errors.New("VMMetrics.GetNomigrate not mocked")
	return
}

var VMMetricsClassGetNomigrateMockedCallback = VMMetricsClassGetNomigrateMockDefault

func (_class VMMetricsClass) GetNomigrateMock(sessionID SessionRef, self VMMetricsRef) (_retval bool, _err error) {
	return VMMetricsClassGetNomigrateMockedCallback(sessionID, self)
}
// Get the nomigrate field of the given VM_metrics.
func (_class VMMetricsClass) GetNomigrate(sessionID SessionRef, self VMMetricsRef) (_retval bool, _err error) {
	if IsMock {
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


func VMMetricsClassGetNestedVirtMockDefault(sessionID SessionRef, self VMMetricsRef) (_retval bool, _err error) {
	log.Println("VMMetrics.GetNestedVirt not mocked")
	_err = errors.New("VMMetrics.GetNestedVirt not mocked")
	return
}

var VMMetricsClassGetNestedVirtMockedCallback = VMMetricsClassGetNestedVirtMockDefault

func (_class VMMetricsClass) GetNestedVirtMock(sessionID SessionRef, self VMMetricsRef) (_retval bool, _err error) {
	return VMMetricsClassGetNestedVirtMockedCallback(sessionID, self)
}
// Get the nested_virt field of the given VM_metrics.
func (_class VMMetricsClass) GetNestedVirt(sessionID SessionRef, self VMMetricsRef) (_retval bool, _err error) {
	if IsMock {
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


func VMMetricsClassGetHvmMockDefault(sessionID SessionRef, self VMMetricsRef) (_retval bool, _err error) {
	log.Println("VMMetrics.GetHvm not mocked")
	_err = errors.New("VMMetrics.GetHvm not mocked")
	return
}

var VMMetricsClassGetHvmMockedCallback = VMMetricsClassGetHvmMockDefault

func (_class VMMetricsClass) GetHvmMock(sessionID SessionRef, self VMMetricsRef) (_retval bool, _err error) {
	return VMMetricsClassGetHvmMockedCallback(sessionID, self)
}
// Get the hvm field of the given VM_metrics.
func (_class VMMetricsClass) GetHvm(sessionID SessionRef, self VMMetricsRef) (_retval bool, _err error) {
	if IsMock {
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


func VMMetricsClassGetOtherConfigMockDefault(sessionID SessionRef, self VMMetricsRef) (_retval map[string]string, _err error) {
	log.Println("VMMetrics.GetOtherConfig not mocked")
	_err = errors.New("VMMetrics.GetOtherConfig not mocked")
	return
}

var VMMetricsClassGetOtherConfigMockedCallback = VMMetricsClassGetOtherConfigMockDefault

func (_class VMMetricsClass) GetOtherConfigMock(sessionID SessionRef, self VMMetricsRef) (_retval map[string]string, _err error) {
	return VMMetricsClassGetOtherConfigMockedCallback(sessionID, self)
}
// Get the other_config field of the given VM_metrics.
func (_class VMMetricsClass) GetOtherConfig(sessionID SessionRef, self VMMetricsRef) (_retval map[string]string, _err error) {
	if IsMock {
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


func VMMetricsClassGetLastUpdatedMockDefault(sessionID SessionRef, self VMMetricsRef) (_retval time.Time, _err error) {
	log.Println("VMMetrics.GetLastUpdated not mocked")
	_err = errors.New("VMMetrics.GetLastUpdated not mocked")
	return
}

var VMMetricsClassGetLastUpdatedMockedCallback = VMMetricsClassGetLastUpdatedMockDefault

func (_class VMMetricsClass) GetLastUpdatedMock(sessionID SessionRef, self VMMetricsRef) (_retval time.Time, _err error) {
	return VMMetricsClassGetLastUpdatedMockedCallback(sessionID, self)
}
// Get the last_updated field of the given VM_metrics.
func (_class VMMetricsClass) GetLastUpdated(sessionID SessionRef, self VMMetricsRef) (_retval time.Time, _err error) {
	if IsMock {
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


func VMMetricsClassGetInstallTimeMockDefault(sessionID SessionRef, self VMMetricsRef) (_retval time.Time, _err error) {
	log.Println("VMMetrics.GetInstallTime not mocked")
	_err = errors.New("VMMetrics.GetInstallTime not mocked")
	return
}

var VMMetricsClassGetInstallTimeMockedCallback = VMMetricsClassGetInstallTimeMockDefault

func (_class VMMetricsClass) GetInstallTimeMock(sessionID SessionRef, self VMMetricsRef) (_retval time.Time, _err error) {
	return VMMetricsClassGetInstallTimeMockedCallback(sessionID, self)
}
// Get the install_time field of the given VM_metrics.
func (_class VMMetricsClass) GetInstallTime(sessionID SessionRef, self VMMetricsRef) (_retval time.Time, _err error) {
	if IsMock {
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


func VMMetricsClassGetStartTimeMockDefault(sessionID SessionRef, self VMMetricsRef) (_retval time.Time, _err error) {
	log.Println("VMMetrics.GetStartTime not mocked")
	_err = errors.New("VMMetrics.GetStartTime not mocked")
	return
}

var VMMetricsClassGetStartTimeMockedCallback = VMMetricsClassGetStartTimeMockDefault

func (_class VMMetricsClass) GetStartTimeMock(sessionID SessionRef, self VMMetricsRef) (_retval time.Time, _err error) {
	return VMMetricsClassGetStartTimeMockedCallback(sessionID, self)
}
// Get the start_time field of the given VM_metrics.
func (_class VMMetricsClass) GetStartTime(sessionID SessionRef, self VMMetricsRef) (_retval time.Time, _err error) {
	if IsMock {
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


func VMMetricsClassGetStateMockDefault(sessionID SessionRef, self VMMetricsRef) (_retval []string, _err error) {
	log.Println("VMMetrics.GetState not mocked")
	_err = errors.New("VMMetrics.GetState not mocked")
	return
}

var VMMetricsClassGetStateMockedCallback = VMMetricsClassGetStateMockDefault

func (_class VMMetricsClass) GetStateMock(sessionID SessionRef, self VMMetricsRef) (_retval []string, _err error) {
	return VMMetricsClassGetStateMockedCallback(sessionID, self)
}
// Get the state field of the given VM_metrics.
func (_class VMMetricsClass) GetState(sessionID SessionRef, self VMMetricsRef) (_retval []string, _err error) {
	if IsMock {
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


func VMMetricsClassGetVCPUsFlagsMockDefault(sessionID SessionRef, self VMMetricsRef) (_retval map[int][]string, _err error) {
	log.Println("VMMetrics.GetVCPUsFlags not mocked")
	_err = errors.New("VMMetrics.GetVCPUsFlags not mocked")
	return
}

var VMMetricsClassGetVCPUsFlagsMockedCallback = VMMetricsClassGetVCPUsFlagsMockDefault

func (_class VMMetricsClass) GetVCPUsFlagsMock(sessionID SessionRef, self VMMetricsRef) (_retval map[int][]string, _err error) {
	return VMMetricsClassGetVCPUsFlagsMockedCallback(sessionID, self)
}
// Get the VCPUs/flags field of the given VM_metrics.
func (_class VMMetricsClass) GetVCPUsFlags(sessionID SessionRef, self VMMetricsRef) (_retval map[int][]string, _err error) {
	if IsMock {
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


func VMMetricsClassGetVCPUsParamsMockDefault(sessionID SessionRef, self VMMetricsRef) (_retval map[string]string, _err error) {
	log.Println("VMMetrics.GetVCPUsParams not mocked")
	_err = errors.New("VMMetrics.GetVCPUsParams not mocked")
	return
}

var VMMetricsClassGetVCPUsParamsMockedCallback = VMMetricsClassGetVCPUsParamsMockDefault

func (_class VMMetricsClass) GetVCPUsParamsMock(sessionID SessionRef, self VMMetricsRef) (_retval map[string]string, _err error) {
	return VMMetricsClassGetVCPUsParamsMockedCallback(sessionID, self)
}
// Get the VCPUs/params field of the given VM_metrics.
func (_class VMMetricsClass) GetVCPUsParams(sessionID SessionRef, self VMMetricsRef) (_retval map[string]string, _err error) {
	if IsMock {
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


func VMMetricsClassGetVCPUsCPUMockDefault(sessionID SessionRef, self VMMetricsRef) (_retval map[int]int, _err error) {
	log.Println("VMMetrics.GetVCPUsCPU not mocked")
	_err = errors.New("VMMetrics.GetVCPUsCPU not mocked")
	return
}

var VMMetricsClassGetVCPUsCPUMockedCallback = VMMetricsClassGetVCPUsCPUMockDefault

func (_class VMMetricsClass) GetVCPUsCPUMock(sessionID SessionRef, self VMMetricsRef) (_retval map[int]int, _err error) {
	return VMMetricsClassGetVCPUsCPUMockedCallback(sessionID, self)
}
// Get the VCPUs/CPU field of the given VM_metrics.
func (_class VMMetricsClass) GetVCPUsCPU(sessionID SessionRef, self VMMetricsRef) (_retval map[int]int, _err error) {
	if IsMock {
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


func VMMetricsClassGetVCPUsUtilisationMockDefault(sessionID SessionRef, self VMMetricsRef) (_retval map[int]float64, _err error) {
	log.Println("VMMetrics.GetVCPUsUtilisation not mocked")
	_err = errors.New("VMMetrics.GetVCPUsUtilisation not mocked")
	return
}

var VMMetricsClassGetVCPUsUtilisationMockedCallback = VMMetricsClassGetVCPUsUtilisationMockDefault

func (_class VMMetricsClass) GetVCPUsUtilisationMock(sessionID SessionRef, self VMMetricsRef) (_retval map[int]float64, _err error) {
	return VMMetricsClassGetVCPUsUtilisationMockedCallback(sessionID, self)
}
// Get the VCPUs/utilisation field of the given VM_metrics.
func (_class VMMetricsClass) GetVCPUsUtilisation(sessionID SessionRef, self VMMetricsRef) (_retval map[int]float64, _err error) {
	if IsMock {
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


func VMMetricsClassGetVCPUsNumberMockDefault(sessionID SessionRef, self VMMetricsRef) (_retval int, _err error) {
	log.Println("VMMetrics.GetVCPUsNumber not mocked")
	_err = errors.New("VMMetrics.GetVCPUsNumber not mocked")
	return
}

var VMMetricsClassGetVCPUsNumberMockedCallback = VMMetricsClassGetVCPUsNumberMockDefault

func (_class VMMetricsClass) GetVCPUsNumberMock(sessionID SessionRef, self VMMetricsRef) (_retval int, _err error) {
	return VMMetricsClassGetVCPUsNumberMockedCallback(sessionID, self)
}
// Get the VCPUs/number field of the given VM_metrics.
func (_class VMMetricsClass) GetVCPUsNumber(sessionID SessionRef, self VMMetricsRef) (_retval int, _err error) {
	if IsMock {
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


func VMMetricsClassGetMemoryActualMockDefault(sessionID SessionRef, self VMMetricsRef) (_retval int, _err error) {
	log.Println("VMMetrics.GetMemoryActual not mocked")
	_err = errors.New("VMMetrics.GetMemoryActual not mocked")
	return
}

var VMMetricsClassGetMemoryActualMockedCallback = VMMetricsClassGetMemoryActualMockDefault

func (_class VMMetricsClass) GetMemoryActualMock(sessionID SessionRef, self VMMetricsRef) (_retval int, _err error) {
	return VMMetricsClassGetMemoryActualMockedCallback(sessionID, self)
}
// Get the memory/actual field of the given VM_metrics.
func (_class VMMetricsClass) GetMemoryActual(sessionID SessionRef, self VMMetricsRef) (_retval int, _err error) {
	if IsMock {
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


func VMMetricsClassGetUUIDMockDefault(sessionID SessionRef, self VMMetricsRef) (_retval string, _err error) {
	log.Println("VMMetrics.GetUUID not mocked")
	_err = errors.New("VMMetrics.GetUUID not mocked")
	return
}

var VMMetricsClassGetUUIDMockedCallback = VMMetricsClassGetUUIDMockDefault

func (_class VMMetricsClass) GetUUIDMock(sessionID SessionRef, self VMMetricsRef) (_retval string, _err error) {
	return VMMetricsClassGetUUIDMockedCallback(sessionID, self)
}
// Get the uuid field of the given VM_metrics.
func (_class VMMetricsClass) GetUUID(sessionID SessionRef, self VMMetricsRef) (_retval string, _err error) {
	if IsMock {
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


func VMMetricsClassGetByUUIDMockDefault(sessionID SessionRef, uuid string) (_retval VMMetricsRef, _err error) {
	log.Println("VMMetrics.GetByUUID not mocked")
	_err = errors.New("VMMetrics.GetByUUID not mocked")
	return
}

var VMMetricsClassGetByUUIDMockedCallback = VMMetricsClassGetByUUIDMockDefault

func (_class VMMetricsClass) GetByUUIDMock(sessionID SessionRef, uuid string) (_retval VMMetricsRef, _err error) {
	return VMMetricsClassGetByUUIDMockedCallback(sessionID, uuid)
}
// Get a reference to the VM_metrics instance with the specified UUID.
func (_class VMMetricsClass) GetByUUID(sessionID SessionRef, uuid string) (_retval VMMetricsRef, _err error) {
	if IsMock {
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


func VMMetricsClassGetRecordMockDefault(sessionID SessionRef, self VMMetricsRef) (_retval VMMetricsRecord, _err error) {
	log.Println("VMMetrics.GetRecord not mocked")
	_err = errors.New("VMMetrics.GetRecord not mocked")
	return
}

var VMMetricsClassGetRecordMockedCallback = VMMetricsClassGetRecordMockDefault

func (_class VMMetricsClass) GetRecordMock(sessionID SessionRef, self VMMetricsRef) (_retval VMMetricsRecord, _err error) {
	return VMMetricsClassGetRecordMockedCallback(sessionID, self)
}
// Get a record containing the current state of the given VM_metrics.
func (_class VMMetricsClass) GetRecord(sessionID SessionRef, self VMMetricsRef) (_retval VMMetricsRecord, _err error) {
	if IsMock {
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
