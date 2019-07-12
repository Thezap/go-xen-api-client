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

func (_class VGPUClass) GetAllRecords__mock(sessionID SessionRef) (_retval map[VGPURef]VGPURecord, _err error) {
	log.Println("VGPU.GetAllRecords not mocked")
	_err = errors.New("VGPU.GetAllRecords not mocked")
	return
}
// Return a map of VGPU references to VGPU records for all VGPUs known to the system.
func (_class VGPUClass) GetAllRecords(sessionID SessionRef) (_retval map[VGPURef]VGPURecord, _err error) {
	if (IsMock) {
		return _class.GetAllRecords__mock(sessionID)
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

func (_class VGPUClass) GetAll__mock(sessionID SessionRef) (_retval []VGPURef, _err error) {
	log.Println("VGPU.GetAll not mocked")
	_err = errors.New("VGPU.GetAll not mocked")
	return
}
// Return a list of all the VGPUs known to the system.
func (_class VGPUClass) GetAll(sessionID SessionRef) (_retval []VGPURef, _err error) {
	if (IsMock) {
		return _class.GetAll__mock(sessionID)
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

func (_class VGPUClass) Destroy__mock(sessionID SessionRef, self VGPURef) (_err error) {
	log.Println("VGPU.Destroy not mocked")
	_err = errors.New("VGPU.Destroy not mocked")
	return
}
// 
func (_class VGPUClass) Destroy(sessionID SessionRef, self VGPURef) (_err error) {
	if (IsMock) {
		return _class.Destroy__mock(sessionID, self)
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

func (_class VGPUClass) Create__mock(sessionID SessionRef, vm VMRef, gpuGroup GPUGroupRef, device string, otherConfig map[string]string, atype VGPUTypeRef) (_retval VGPURef, _err error) {
	log.Println("VGPU.Create not mocked")
	_err = errors.New("VGPU.Create not mocked")
	return
}
// 
func (_class VGPUClass) Create(sessionID SessionRef, vm VMRef, gpuGroup GPUGroupRef, device string, otherConfig map[string]string, atype VGPUTypeRef) (_retval VGPURef, _err error) {
	if (IsMock) {
		return _class.Create__mock(sessionID, vm, gpuGroup, device, otherConfig, atype)
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

func (_class VGPUClass) RemoveFromOtherConfig__mock(sessionID SessionRef, self VGPURef, key string) (_err error) {
	log.Println("VGPU.RemoveFromOtherConfig not mocked")
	_err = errors.New("VGPU.RemoveFromOtherConfig not mocked")
	return
}
// Remove the given key and its corresponding value from the other_config field of the given VGPU.  If the key is not in that Map, then do nothing.
func (_class VGPUClass) RemoveFromOtherConfig(sessionID SessionRef, self VGPURef, key string) (_err error) {
	if (IsMock) {
		return _class.RemoveFromOtherConfig__mock(sessionID, self, key)
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

func (_class VGPUClass) AddToOtherConfig__mock(sessionID SessionRef, self VGPURef, key string, value string) (_err error) {
	log.Println("VGPU.AddToOtherConfig not mocked")
	_err = errors.New("VGPU.AddToOtherConfig not mocked")
	return
}
// Add the given key-value pair to the other_config field of the given VGPU.
func (_class VGPUClass) AddToOtherConfig(sessionID SessionRef, self VGPURef, key string, value string) (_err error) {
	if (IsMock) {
		return _class.AddToOtherConfig__mock(sessionID, self, key, value)
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

func (_class VGPUClass) SetOtherConfig__mock(sessionID SessionRef, self VGPURef, value map[string]string) (_err error) {
	log.Println("VGPU.SetOtherConfig not mocked")
	_err = errors.New("VGPU.SetOtherConfig not mocked")
	return
}
// Set the other_config field of the given VGPU.
func (_class VGPUClass) SetOtherConfig(sessionID SessionRef, self VGPURef, value map[string]string) (_err error) {
	if (IsMock) {
		return _class.SetOtherConfig__mock(sessionID, self, value)
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

func (_class VGPUClass) GetCompatibilityMetadata__mock(sessionID SessionRef, self VGPURef) (_retval map[string]string, _err error) {
	log.Println("VGPU.GetCompatibilityMetadata not mocked")
	_err = errors.New("VGPU.GetCompatibilityMetadata not mocked")
	return
}
// Get the compatibility_metadata field of the given VGPU.
func (_class VGPUClass) GetCompatibilityMetadata(sessionID SessionRef, self VGPURef) (_retval map[string]string, _err error) {
	if (IsMock) {
		return _class.GetCompatibilityMetadata__mock(sessionID, self)
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

func (_class VGPUClass) GetScheduledToBeResidentOn__mock(sessionID SessionRef, self VGPURef) (_retval PGPURef, _err error) {
	log.Println("VGPU.GetScheduledToBeResidentOn not mocked")
	_err = errors.New("VGPU.GetScheduledToBeResidentOn not mocked")
	return
}
// Get the scheduled_to_be_resident_on field of the given VGPU.
func (_class VGPUClass) GetScheduledToBeResidentOn(sessionID SessionRef, self VGPURef) (_retval PGPURef, _err error) {
	if (IsMock) {
		return _class.GetScheduledToBeResidentOn__mock(sessionID, self)
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

func (_class VGPUClass) GetResidentOn__mock(sessionID SessionRef, self VGPURef) (_retval PGPURef, _err error) {
	log.Println("VGPU.GetResidentOn not mocked")
	_err = errors.New("VGPU.GetResidentOn not mocked")
	return
}
// Get the resident_on field of the given VGPU.
func (_class VGPUClass) GetResidentOn(sessionID SessionRef, self VGPURef) (_retval PGPURef, _err error) {
	if (IsMock) {
		return _class.GetResidentOn__mock(sessionID, self)
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

func (_class VGPUClass) GetType__mock(sessionID SessionRef, self VGPURef) (_retval VGPUTypeRef, _err error) {
	log.Println("VGPU.GetType not mocked")
	_err = errors.New("VGPU.GetType not mocked")
	return
}
// Get the type field of the given VGPU.
func (_class VGPUClass) GetType(sessionID SessionRef, self VGPURef) (_retval VGPUTypeRef, _err error) {
	if (IsMock) {
		return _class.GetType__mock(sessionID, self)
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

func (_class VGPUClass) GetOtherConfig__mock(sessionID SessionRef, self VGPURef) (_retval map[string]string, _err error) {
	log.Println("VGPU.GetOtherConfig not mocked")
	_err = errors.New("VGPU.GetOtherConfig not mocked")
	return
}
// Get the other_config field of the given VGPU.
func (_class VGPUClass) GetOtherConfig(sessionID SessionRef, self VGPURef) (_retval map[string]string, _err error) {
	if (IsMock) {
		return _class.GetOtherConfig__mock(sessionID, self)
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

func (_class VGPUClass) GetCurrentlyAttached__mock(sessionID SessionRef, self VGPURef) (_retval bool, _err error) {
	log.Println("VGPU.GetCurrentlyAttached not mocked")
	_err = errors.New("VGPU.GetCurrentlyAttached not mocked")
	return
}
// Get the currently_attached field of the given VGPU.
func (_class VGPUClass) GetCurrentlyAttached(sessionID SessionRef, self VGPURef) (_retval bool, _err error) {
	if (IsMock) {
		return _class.GetCurrentlyAttached__mock(sessionID, self)
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

func (_class VGPUClass) GetDevice__mock(sessionID SessionRef, self VGPURef) (_retval string, _err error) {
	log.Println("VGPU.GetDevice not mocked")
	_err = errors.New("VGPU.GetDevice not mocked")
	return
}
// Get the device field of the given VGPU.
func (_class VGPUClass) GetDevice(sessionID SessionRef, self VGPURef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetDevice__mock(sessionID, self)
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

func (_class VGPUClass) GetGPUGroup__mock(sessionID SessionRef, self VGPURef) (_retval GPUGroupRef, _err error) {
	log.Println("VGPU.GetGPUGroup not mocked")
	_err = errors.New("VGPU.GetGPUGroup not mocked")
	return
}
// Get the GPU_group field of the given VGPU.
func (_class VGPUClass) GetGPUGroup(sessionID SessionRef, self VGPURef) (_retval GPUGroupRef, _err error) {
	if (IsMock) {
		return _class.GetGPUGroup__mock(sessionID, self)
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

func (_class VGPUClass) GetVM__mock(sessionID SessionRef, self VGPURef) (_retval VMRef, _err error) {
	log.Println("VGPU.GetVM not mocked")
	_err = errors.New("VGPU.GetVM not mocked")
	return
}
// Get the VM field of the given VGPU.
func (_class VGPUClass) GetVM(sessionID SessionRef, self VGPURef) (_retval VMRef, _err error) {
	if (IsMock) {
		return _class.GetVM__mock(sessionID, self)
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

func (_class VGPUClass) GetUUID__mock(sessionID SessionRef, self VGPURef) (_retval string, _err error) {
	log.Println("VGPU.GetUUID not mocked")
	_err = errors.New("VGPU.GetUUID not mocked")
	return
}
// Get the uuid field of the given VGPU.
func (_class VGPUClass) GetUUID(sessionID SessionRef, self VGPURef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetUUID__mock(sessionID, self)
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

func (_class VGPUClass) GetByUUID__mock(sessionID SessionRef, uuid string) (_retval VGPURef, _err error) {
	log.Println("VGPU.GetByUUID not mocked")
	_err = errors.New("VGPU.GetByUUID not mocked")
	return
}
// Get a reference to the VGPU instance with the specified UUID.
func (_class VGPUClass) GetByUUID(sessionID SessionRef, uuid string) (_retval VGPURef, _err error) {
	if (IsMock) {
		return _class.GetByUUID__mock(sessionID, uuid)
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

func (_class VGPUClass) GetRecord__mock(sessionID SessionRef, self VGPURef) (_retval VGPURecord, _err error) {
	log.Println("VGPU.GetRecord not mocked")
	_err = errors.New("VGPU.GetRecord not mocked")
	return
}
// Get a record containing the current state of the given VGPU.
func (_class VGPUClass) GetRecord(sessionID SessionRef, self VGPURef) (_retval VGPURecord, _err error) {
	if (IsMock) {
		return _class.GetRecord__mock(sessionID, self)
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
