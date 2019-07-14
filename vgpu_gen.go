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

type VGPURecord struct {
  // Unique identifier/object reference
	UUID string
  // VM that owns the vGPU
	VM VMRef
  // GPU group used by the vGPU
	GPUGroup GPUGroupRef
  // Order in which the devices are plugged into the VM
	Device string
  // Reflects whether the virtual device is currently connected to a physical device
	CurrentlyAttached bool
  // Additional configuration
	OtherConfig map[string]string
  // Preset type for this VGPU
	Type VGPUTypeRef
  // The PGPU on which this VGPU is running
	ResidentOn PGPURef
  // The PGPU on which this VGPU is scheduled to run
	ScheduledToBeResidentOn PGPURef
  // VGPU metadata to determine whether a VGPU can migrate between two PGPUs
	CompatibilityMetadata map[string]string
}

type VGPURef string

// A virtual GPU (vGPU)
type VGPUClass struct {
	client *Client
}


func VGPUClassGetAllRecordsMockDefault(sessionID SessionRef) (_retval map[VGPURef]VGPURecord, _err error) {
	log.Println("VGPU.GetAllRecords not mocked")
	_err = errors.New("VGPU.GetAllRecords not mocked")
	return
}

var VGPUClassGetAllRecordsMockedCallback = VGPUClassGetAllRecordsMockDefault

func (_class VGPUClass) GetAllRecordsMock(sessionID SessionRef) (_retval map[VGPURef]VGPURecord, _err error) {
	return VGPUClassGetAllRecordsMockedCallback(sessionID)
}
// Return a map of VGPU references to VGPU records for all VGPUs known to the system.
func (_class VGPUClass) GetAllRecords(sessionID SessionRef) (_retval map[VGPURef]VGPURecord, _err error) {
	if IsMock {
		return _class.GetAllRecordsMock(sessionID)
	}	
	_method := "VGPU.get_all_records"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVGPURefToVGPURecordMapToGo(_method + " -> ", _result.Value)
	return
}


func VGPUClassGetAllMockDefault(sessionID SessionRef) (_retval []VGPURef, _err error) {
	log.Println("VGPU.GetAll not mocked")
	_err = errors.New("VGPU.GetAll not mocked")
	return
}

var VGPUClassGetAllMockedCallback = VGPUClassGetAllMockDefault

func (_class VGPUClass) GetAllMock(sessionID SessionRef) (_retval []VGPURef, _err error) {
	return VGPUClassGetAllMockedCallback(sessionID)
}
// Return a list of all the VGPUs known to the system.
func (_class VGPUClass) GetAll(sessionID SessionRef) (_retval []VGPURef, _err error) {
	if IsMock {
		return _class.GetAllMock(sessionID)
	}	
	_method := "VGPU.get_all"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVGPURefSetToGo(_method + " -> ", _result.Value)
	return
}


func VGPUClassDestroyMockDefault(sessionID SessionRef, self VGPURef) (_err error) {
	log.Println("VGPU.Destroy not mocked")
	_err = errors.New("VGPU.Destroy not mocked")
	return
}

var VGPUClassDestroyMockedCallback = VGPUClassDestroyMockDefault

func (_class VGPUClass) DestroyMock(sessionID SessionRef, self VGPURef) (_err error) {
	return VGPUClassDestroyMockedCallback(sessionID, self)
}
// 
func (_class VGPUClass) Destroy(sessionID SessionRef, self VGPURef) (_err error) {
	if IsMock {
		return _class.DestroyMock(sessionID, self)
	}	
	_method := "VGPU.destroy"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVGPURefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg)
	return
}


func VGPUClassCreateMockDefault(sessionID SessionRef, vm VMRef, gpuGroup GPUGroupRef, device string, otherConfig map[string]string, atype VGPUTypeRef) (_retval VGPURef, _err error) {
	log.Println("VGPU.Create not mocked")
	_err = errors.New("VGPU.Create not mocked")
	return
}

var VGPUClassCreateMockedCallback = VGPUClassCreateMockDefault

func (_class VGPUClass) CreateMock(sessionID SessionRef, vm VMRef, gpuGroup GPUGroupRef, device string, otherConfig map[string]string, atype VGPUTypeRef) (_retval VGPURef, _err error) {
	return VGPUClassCreateMockedCallback(sessionID, vm, gpuGroup, device, otherConfig, atype)
}
// 
func (_class VGPUClass) Create(sessionID SessionRef, vm VMRef, gpuGroup GPUGroupRef, device string, otherConfig map[string]string, atype VGPUTypeRef) (_retval VGPURef, _err error) {
	if IsMock {
		return _class.CreateMock(sessionID, vm, gpuGroup, device, otherConfig, atype)
	}	
	_method := "VGPU.create"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_vmArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "VM"), vm)
	if _err != nil {
		return
	}
	_gpuGroupArg, _err := convertGPUGroupRefToXen(fmt.Sprintf("%s(%s)", _method, "GPU_group"), gpuGroup)
	if _err != nil {
		return
	}
	_deviceArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "device"), device)
	if _err != nil {
		return
	}
	_otherConfigArg, _err := convertStringToStringMapToXen(fmt.Sprintf("%s(%s)", _method, "other_config"), otherConfig)
	if _err != nil {
		return
	}
	_atypeArg, _err := convertVGPUTypeRefToXen(fmt.Sprintf("%s(%s)", _method, "type"), atype)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _vmArg, _gpuGroupArg, _deviceArg, _otherConfigArg, _atypeArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVGPURefToGo(_method + " -> ", _result.Value)
	return
}


func VGPUClassRemoveFromOtherConfigMockDefault(sessionID SessionRef, self VGPURef, key string) (_err error) {
	log.Println("VGPU.RemoveFromOtherConfig not mocked")
	_err = errors.New("VGPU.RemoveFromOtherConfig not mocked")
	return
}

var VGPUClassRemoveFromOtherConfigMockedCallback = VGPUClassRemoveFromOtherConfigMockDefault

func (_class VGPUClass) RemoveFromOtherConfigMock(sessionID SessionRef, self VGPURef, key string) (_err error) {
	return VGPUClassRemoveFromOtherConfigMockedCallback(sessionID, self, key)
}
// Remove the given key and its corresponding value from the other_config field of the given VGPU.  If the key is not in that Map, then do nothing.
func (_class VGPUClass) RemoveFromOtherConfig(sessionID SessionRef, self VGPURef, key string) (_err error) {
	if IsMock {
		return _class.RemoveFromOtherConfigMock(sessionID, self, key)
	}	
	_method := "VGPU.remove_from_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVGPURefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func VGPUClassAddToOtherConfigMockDefault(sessionID SessionRef, self VGPURef, key string, value string) (_err error) {
	log.Println("VGPU.AddToOtherConfig not mocked")
	_err = errors.New("VGPU.AddToOtherConfig not mocked")
	return
}

var VGPUClassAddToOtherConfigMockedCallback = VGPUClassAddToOtherConfigMockDefault

func (_class VGPUClass) AddToOtherConfigMock(sessionID SessionRef, self VGPURef, key string, value string) (_err error) {
	return VGPUClassAddToOtherConfigMockedCallback(sessionID, self, key, value)
}
// Add the given key-value pair to the other_config field of the given VGPU.
func (_class VGPUClass) AddToOtherConfig(sessionID SessionRef, self VGPURef, key string, value string) (_err error) {
	if IsMock {
		return _class.AddToOtherConfigMock(sessionID, self, key, value)
	}	
	_method := "VGPU.add_to_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVGPURefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func VGPUClassSetOtherConfigMockDefault(sessionID SessionRef, self VGPURef, value map[string]string) (_err error) {
	log.Println("VGPU.SetOtherConfig not mocked")
	_err = errors.New("VGPU.SetOtherConfig not mocked")
	return
}

var VGPUClassSetOtherConfigMockedCallback = VGPUClassSetOtherConfigMockDefault

func (_class VGPUClass) SetOtherConfigMock(sessionID SessionRef, self VGPURef, value map[string]string) (_err error) {
	return VGPUClassSetOtherConfigMockedCallback(sessionID, self, value)
}
// Set the other_config field of the given VGPU.
func (_class VGPUClass) SetOtherConfig(sessionID SessionRef, self VGPURef, value map[string]string) (_err error) {
	if IsMock {
		return _class.SetOtherConfigMock(sessionID, self, value)
	}	
	_method := "VGPU.set_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVGPURefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func VGPUClassGetCompatibilityMetadataMockDefault(sessionID SessionRef, self VGPURef) (_retval map[string]string, _err error) {
	log.Println("VGPU.GetCompatibilityMetadata not mocked")
	_err = errors.New("VGPU.GetCompatibilityMetadata not mocked")
	return
}

var VGPUClassGetCompatibilityMetadataMockedCallback = VGPUClassGetCompatibilityMetadataMockDefault

func (_class VGPUClass) GetCompatibilityMetadataMock(sessionID SessionRef, self VGPURef) (_retval map[string]string, _err error) {
	return VGPUClassGetCompatibilityMetadataMockedCallback(sessionID, self)
}
// Get the compatibility_metadata field of the given VGPU.
func (_class VGPUClass) GetCompatibilityMetadata(sessionID SessionRef, self VGPURef) (_retval map[string]string, _err error) {
	if IsMock {
		return _class.GetCompatibilityMetadataMock(sessionID, self)
	}	
	_method := "VGPU.get_compatibility_metadata"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVGPURefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func VGPUClassGetScheduledToBeResidentOnMockDefault(sessionID SessionRef, self VGPURef) (_retval PGPURef, _err error) {
	log.Println("VGPU.GetScheduledToBeResidentOn not mocked")
	_err = errors.New("VGPU.GetScheduledToBeResidentOn not mocked")
	return
}

var VGPUClassGetScheduledToBeResidentOnMockedCallback = VGPUClassGetScheduledToBeResidentOnMockDefault

func (_class VGPUClass) GetScheduledToBeResidentOnMock(sessionID SessionRef, self VGPURef) (_retval PGPURef, _err error) {
	return VGPUClassGetScheduledToBeResidentOnMockedCallback(sessionID, self)
}
// Get the scheduled_to_be_resident_on field of the given VGPU.
func (_class VGPUClass) GetScheduledToBeResidentOn(sessionID SessionRef, self VGPURef) (_retval PGPURef, _err error) {
	if IsMock {
		return _class.GetScheduledToBeResidentOnMock(sessionID, self)
	}	
	_method := "VGPU.get_scheduled_to_be_resident_on"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVGPURefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertPGPURefToGo(_method + " -> ", _result.Value)
	return
}


func VGPUClassGetResidentOnMockDefault(sessionID SessionRef, self VGPURef) (_retval PGPURef, _err error) {
	log.Println("VGPU.GetResidentOn not mocked")
	_err = errors.New("VGPU.GetResidentOn not mocked")
	return
}

var VGPUClassGetResidentOnMockedCallback = VGPUClassGetResidentOnMockDefault

func (_class VGPUClass) GetResidentOnMock(sessionID SessionRef, self VGPURef) (_retval PGPURef, _err error) {
	return VGPUClassGetResidentOnMockedCallback(sessionID, self)
}
// Get the resident_on field of the given VGPU.
func (_class VGPUClass) GetResidentOn(sessionID SessionRef, self VGPURef) (_retval PGPURef, _err error) {
	if IsMock {
		return _class.GetResidentOnMock(sessionID, self)
	}	
	_method := "VGPU.get_resident_on"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVGPURefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertPGPURefToGo(_method + " -> ", _result.Value)
	return
}


func VGPUClassGetTypeMockDefault(sessionID SessionRef, self VGPURef) (_retval VGPUTypeRef, _err error) {
	log.Println("VGPU.GetType not mocked")
	_err = errors.New("VGPU.GetType not mocked")
	return
}

var VGPUClassGetTypeMockedCallback = VGPUClassGetTypeMockDefault

func (_class VGPUClass) GetTypeMock(sessionID SessionRef, self VGPURef) (_retval VGPUTypeRef, _err error) {
	return VGPUClassGetTypeMockedCallback(sessionID, self)
}
// Get the type field of the given VGPU.
func (_class VGPUClass) GetType(sessionID SessionRef, self VGPURef) (_retval VGPUTypeRef, _err error) {
	if IsMock {
		return _class.GetTypeMock(sessionID, self)
	}	
	_method := "VGPU.get_type"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVGPURefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVGPUTypeRefToGo(_method + " -> ", _result.Value)
	return
}


func VGPUClassGetOtherConfigMockDefault(sessionID SessionRef, self VGPURef) (_retval map[string]string, _err error) {
	log.Println("VGPU.GetOtherConfig not mocked")
	_err = errors.New("VGPU.GetOtherConfig not mocked")
	return
}

var VGPUClassGetOtherConfigMockedCallback = VGPUClassGetOtherConfigMockDefault

func (_class VGPUClass) GetOtherConfigMock(sessionID SessionRef, self VGPURef) (_retval map[string]string, _err error) {
	return VGPUClassGetOtherConfigMockedCallback(sessionID, self)
}
// Get the other_config field of the given VGPU.
func (_class VGPUClass) GetOtherConfig(sessionID SessionRef, self VGPURef) (_retval map[string]string, _err error) {
	if IsMock {
		return _class.GetOtherConfigMock(sessionID, self)
	}	
	_method := "VGPU.get_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVGPURefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func VGPUClassGetCurrentlyAttachedMockDefault(sessionID SessionRef, self VGPURef) (_retval bool, _err error) {
	log.Println("VGPU.GetCurrentlyAttached not mocked")
	_err = errors.New("VGPU.GetCurrentlyAttached not mocked")
	return
}

var VGPUClassGetCurrentlyAttachedMockedCallback = VGPUClassGetCurrentlyAttachedMockDefault

func (_class VGPUClass) GetCurrentlyAttachedMock(sessionID SessionRef, self VGPURef) (_retval bool, _err error) {
	return VGPUClassGetCurrentlyAttachedMockedCallback(sessionID, self)
}
// Get the currently_attached field of the given VGPU.
func (_class VGPUClass) GetCurrentlyAttached(sessionID SessionRef, self VGPURef) (_retval bool, _err error) {
	if IsMock {
		return _class.GetCurrentlyAttachedMock(sessionID, self)
	}	
	_method := "VGPU.get_currently_attached"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVGPURefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func VGPUClassGetDeviceMockDefault(sessionID SessionRef, self VGPURef) (_retval string, _err error) {
	log.Println("VGPU.GetDevice not mocked")
	_err = errors.New("VGPU.GetDevice not mocked")
	return
}

var VGPUClassGetDeviceMockedCallback = VGPUClassGetDeviceMockDefault

func (_class VGPUClass) GetDeviceMock(sessionID SessionRef, self VGPURef) (_retval string, _err error) {
	return VGPUClassGetDeviceMockedCallback(sessionID, self)
}
// Get the device field of the given VGPU.
func (_class VGPUClass) GetDevice(sessionID SessionRef, self VGPURef) (_retval string, _err error) {
	if IsMock {
		return _class.GetDeviceMock(sessionID, self)
	}	
	_method := "VGPU.get_device"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVGPURefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func VGPUClassGetGPUGroupMockDefault(sessionID SessionRef, self VGPURef) (_retval GPUGroupRef, _err error) {
	log.Println("VGPU.GetGPUGroup not mocked")
	_err = errors.New("VGPU.GetGPUGroup not mocked")
	return
}

var VGPUClassGetGPUGroupMockedCallback = VGPUClassGetGPUGroupMockDefault

func (_class VGPUClass) GetGPUGroupMock(sessionID SessionRef, self VGPURef) (_retval GPUGroupRef, _err error) {
	return VGPUClassGetGPUGroupMockedCallback(sessionID, self)
}
// Get the GPU_group field of the given VGPU.
func (_class VGPUClass) GetGPUGroup(sessionID SessionRef, self VGPURef) (_retval GPUGroupRef, _err error) {
	if IsMock {
		return _class.GetGPUGroupMock(sessionID, self)
	}	
	_method := "VGPU.get_GPU_group"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVGPURefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertGPUGroupRefToGo(_method + " -> ", _result.Value)
	return
}


func VGPUClassGetVMMockDefault(sessionID SessionRef, self VGPURef) (_retval VMRef, _err error) {
	log.Println("VGPU.GetVM not mocked")
	_err = errors.New("VGPU.GetVM not mocked")
	return
}

var VGPUClassGetVMMockedCallback = VGPUClassGetVMMockDefault

func (_class VGPUClass) GetVMMock(sessionID SessionRef, self VGPURef) (_retval VMRef, _err error) {
	return VGPUClassGetVMMockedCallback(sessionID, self)
}
// Get the VM field of the given VGPU.
func (_class VGPUClass) GetVM(sessionID SessionRef, self VGPURef) (_retval VMRef, _err error) {
	if IsMock {
		return _class.GetVMMock(sessionID, self)
	}	
	_method := "VGPU.get_VM"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVGPURefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVMRefToGo(_method + " -> ", _result.Value)
	return
}


func VGPUClassGetUUIDMockDefault(sessionID SessionRef, self VGPURef) (_retval string, _err error) {
	log.Println("VGPU.GetUUID not mocked")
	_err = errors.New("VGPU.GetUUID not mocked")
	return
}

var VGPUClassGetUUIDMockedCallback = VGPUClassGetUUIDMockDefault

func (_class VGPUClass) GetUUIDMock(sessionID SessionRef, self VGPURef) (_retval string, _err error) {
	return VGPUClassGetUUIDMockedCallback(sessionID, self)
}
// Get the uuid field of the given VGPU.
func (_class VGPUClass) GetUUID(sessionID SessionRef, self VGPURef) (_retval string, _err error) {
	if IsMock {
		return _class.GetUUIDMock(sessionID, self)
	}	
	_method := "VGPU.get_uuid"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVGPURefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func VGPUClassGetByUUIDMockDefault(sessionID SessionRef, uuid string) (_retval VGPURef, _err error) {
	log.Println("VGPU.GetByUUID not mocked")
	_err = errors.New("VGPU.GetByUUID not mocked")
	return
}

var VGPUClassGetByUUIDMockedCallback = VGPUClassGetByUUIDMockDefault

func (_class VGPUClass) GetByUUIDMock(sessionID SessionRef, uuid string) (_retval VGPURef, _err error) {
	return VGPUClassGetByUUIDMockedCallback(sessionID, uuid)
}
// Get a reference to the VGPU instance with the specified UUID.
func (_class VGPUClass) GetByUUID(sessionID SessionRef, uuid string) (_retval VGPURef, _err error) {
	if IsMock {
		return _class.GetByUUIDMock(sessionID, uuid)
	}	
	_method := "VGPU.get_by_uuid"
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
	_retval, _err = convertVGPURefToGo(_method + " -> ", _result.Value)
	return
}


func VGPUClassGetRecordMockDefault(sessionID SessionRef, self VGPURef) (_retval VGPURecord, _err error) {
	log.Println("VGPU.GetRecord not mocked")
	_err = errors.New("VGPU.GetRecord not mocked")
	return
}

var VGPUClassGetRecordMockedCallback = VGPUClassGetRecordMockDefault

func (_class VGPUClass) GetRecordMock(sessionID SessionRef, self VGPURef) (_retval VGPURecord, _err error) {
	return VGPUClassGetRecordMockedCallback(sessionID, self)
}
// Get a record containing the current state of the given VGPU.
func (_class VGPUClass) GetRecord(sessionID SessionRef, self VGPURef) (_retval VGPURecord, _err error) {
	if IsMock {
		return _class.GetRecordMock(sessionID, self)
	}	
	_method := "VGPU.get_record"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVGPURefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVGPURecordToGo(_method + " -> ", _result.Value)
	return
}
