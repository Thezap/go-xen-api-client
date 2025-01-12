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

type AllocationAlgorithm string

const (
  // vGPUs of a given type are allocated evenly across supporting pGPUs.
	AllocationAlgorithmBreadthFirst AllocationAlgorithm = "breadth_first"
  // vGPUs of a given type are allocated on supporting pGPUs until they are full.
	AllocationAlgorithmDepthFirst AllocationAlgorithm = "depth_first"
)

type GPUGroupRecord struct {
  // Unique identifier/object reference
	UUID string
  // a human-readable name
	NameLabel string
  // a notes field containing human-readable description
	NameDescription string
  // List of pGPUs in the group
	PGPUs []PGPURef
  // List of vGPUs using the group
	VGPUs []VGPURef
  // List of GPU types (vendor+device ID) that can be in this group
	GPUTypes []string
  // Additional configuration
	OtherConfig map[string]string
  // Current allocation of vGPUs to pGPUs for this group
	AllocationAlgorithm AllocationAlgorithm
  // vGPU types supported on at least one of the pGPUs in this group
	SupportedVGPUTypes []VGPUTypeRef
  // vGPU types supported on at least one of the pGPUs in this group
	EnabledVGPUTypes []VGPUTypeRef
}

type GPUGroupRef string

// A group of compatible GPUs across the resource pool
type GPUGroupClass struct {
	client *Client
}


func GPUGroupClassGetAllRecordsMockDefault(sessionID SessionRef) (_retval map[GPUGroupRef]GPUGroupRecord, _err error) {
	log.Println("GPUGroup.GetAllRecords not mocked")
	_err = errors.New("GPUGroup.GetAllRecords not mocked")
	return
}

var GPUGroupClassGetAllRecordsMockedCallback = GPUGroupClassGetAllRecordsMockDefault

func (_class GPUGroupClass) GetAllRecordsMock(sessionID SessionRef) (_retval map[GPUGroupRef]GPUGroupRecord, _err error) {
	return GPUGroupClassGetAllRecordsMockedCallback(sessionID)
}
// Return a map of GPU_group references to GPU_group records for all GPU_groups known to the system.
func (_class GPUGroupClass) GetAllRecords(sessionID SessionRef) (_retval map[GPUGroupRef]GPUGroupRecord, _err error) {
	if IsMock {
		return _class.GetAllRecordsMock(sessionID)
	}	
	_method := "GPU_group.get_all_records"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertGPUGroupRefToGPUGroupRecordMapToGo(_method + " -> ", _result.Value)
	return
}


func GPUGroupClassGetAllMockDefault(sessionID SessionRef) (_retval []GPUGroupRef, _err error) {
	log.Println("GPUGroup.GetAll not mocked")
	_err = errors.New("GPUGroup.GetAll not mocked")
	return
}

var GPUGroupClassGetAllMockedCallback = GPUGroupClassGetAllMockDefault

func (_class GPUGroupClass) GetAllMock(sessionID SessionRef) (_retval []GPUGroupRef, _err error) {
	return GPUGroupClassGetAllMockedCallback(sessionID)
}
// Return a list of all the GPU_groups known to the system.
func (_class GPUGroupClass) GetAll(sessionID SessionRef) (_retval []GPUGroupRef, _err error) {
	if IsMock {
		return _class.GetAllMock(sessionID)
	}	
	_method := "GPU_group.get_all"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertGPUGroupRefSetToGo(_method + " -> ", _result.Value)
	return
}


func GPUGroupClassGetRemainingCapacityMockDefault(sessionID SessionRef, self GPUGroupRef, vgpuType VGPUTypeRef) (_retval int, _err error) {
	log.Println("GPUGroup.GetRemainingCapacity not mocked")
	_err = errors.New("GPUGroup.GetRemainingCapacity not mocked")
	return
}

var GPUGroupClassGetRemainingCapacityMockedCallback = GPUGroupClassGetRemainingCapacityMockDefault

func (_class GPUGroupClass) GetRemainingCapacityMock(sessionID SessionRef, self GPUGroupRef, vgpuType VGPUTypeRef) (_retval int, _err error) {
	return GPUGroupClassGetRemainingCapacityMockedCallback(sessionID, self, vgpuType)
}
// 
func (_class GPUGroupClass) GetRemainingCapacity(sessionID SessionRef, self GPUGroupRef, vgpuType VGPUTypeRef) (_retval int, _err error) {
	if IsMock {
		return _class.GetRemainingCapacityMock(sessionID, self, vgpuType)
	}	
	_method := "GPU_group.get_remaining_capacity"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertGPUGroupRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_vgpuTypeArg, _err := convertVGPUTypeRefToXen(fmt.Sprintf("%s(%s)", _method, "vgpu_type"), vgpuType)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg, _vgpuTypeArg)
	if _err != nil {
		return
	}
	_retval, _err = convertIntToGo(_method + " -> ", _result.Value)
	return
}


func GPUGroupClassDestroyMockDefault(sessionID SessionRef, self GPUGroupRef) (_err error) {
	log.Println("GPUGroup.Destroy not mocked")
	_err = errors.New("GPUGroup.Destroy not mocked")
	return
}

var GPUGroupClassDestroyMockedCallback = GPUGroupClassDestroyMockDefault

func (_class GPUGroupClass) DestroyMock(sessionID SessionRef, self GPUGroupRef) (_err error) {
	return GPUGroupClassDestroyMockedCallback(sessionID, self)
}
// 
func (_class GPUGroupClass) Destroy(sessionID SessionRef, self GPUGroupRef) (_err error) {
	if IsMock {
		return _class.DestroyMock(sessionID, self)
	}	
	_method := "GPU_group.destroy"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertGPUGroupRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg)
	return
}


func GPUGroupClassCreateMockDefault(sessionID SessionRef, nameLabel string, nameDescription string, otherConfig map[string]string) (_retval GPUGroupRef, _err error) {
	log.Println("GPUGroup.Create not mocked")
	_err = errors.New("GPUGroup.Create not mocked")
	return
}

var GPUGroupClassCreateMockedCallback = GPUGroupClassCreateMockDefault

func (_class GPUGroupClass) CreateMock(sessionID SessionRef, nameLabel string, nameDescription string, otherConfig map[string]string) (_retval GPUGroupRef, _err error) {
	return GPUGroupClassCreateMockedCallback(sessionID, nameLabel, nameDescription, otherConfig)
}
// 
func (_class GPUGroupClass) Create(sessionID SessionRef, nameLabel string, nameDescription string, otherConfig map[string]string) (_retval GPUGroupRef, _err error) {
	if IsMock {
		return _class.CreateMock(sessionID, nameLabel, nameDescription, otherConfig)
	}	
	_method := "GPU_group.create"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_nameLabelArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "name_label"), nameLabel)
	if _err != nil {
		return
	}
	_nameDescriptionArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "name_description"), nameDescription)
	if _err != nil {
		return
	}
	_otherConfigArg, _err := convertStringToStringMapToXen(fmt.Sprintf("%s(%s)", _method, "other_config"), otherConfig)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _nameLabelArg, _nameDescriptionArg, _otherConfigArg)
	if _err != nil {
		return
	}
	_retval, _err = convertGPUGroupRefToGo(_method + " -> ", _result.Value)
	return
}


func GPUGroupClassSetAllocationAlgorithmMockDefault(sessionID SessionRef, self GPUGroupRef, value AllocationAlgorithm) (_err error) {
	log.Println("GPUGroup.SetAllocationAlgorithm not mocked")
	_err = errors.New("GPUGroup.SetAllocationAlgorithm not mocked")
	return
}

var GPUGroupClassSetAllocationAlgorithmMockedCallback = GPUGroupClassSetAllocationAlgorithmMockDefault

func (_class GPUGroupClass) SetAllocationAlgorithmMock(sessionID SessionRef, self GPUGroupRef, value AllocationAlgorithm) (_err error) {
	return GPUGroupClassSetAllocationAlgorithmMockedCallback(sessionID, self, value)
}
// Set the allocation_algorithm field of the given GPU_group.
func (_class GPUGroupClass) SetAllocationAlgorithm(sessionID SessionRef, self GPUGroupRef, value AllocationAlgorithm) (_err error) {
	if IsMock {
		return _class.SetAllocationAlgorithmMock(sessionID, self, value)
	}	
	_method := "GPU_group.set_allocation_algorithm"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertGPUGroupRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertEnumAllocationAlgorithmToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}


func GPUGroupClassRemoveFromOtherConfigMockDefault(sessionID SessionRef, self GPUGroupRef, key string) (_err error) {
	log.Println("GPUGroup.RemoveFromOtherConfig not mocked")
	_err = errors.New("GPUGroup.RemoveFromOtherConfig not mocked")
	return
}

var GPUGroupClassRemoveFromOtherConfigMockedCallback = GPUGroupClassRemoveFromOtherConfigMockDefault

func (_class GPUGroupClass) RemoveFromOtherConfigMock(sessionID SessionRef, self GPUGroupRef, key string) (_err error) {
	return GPUGroupClassRemoveFromOtherConfigMockedCallback(sessionID, self, key)
}
// Remove the given key and its corresponding value from the other_config field of the given GPU_group.  If the key is not in that Map, then do nothing.
func (_class GPUGroupClass) RemoveFromOtherConfig(sessionID SessionRef, self GPUGroupRef, key string) (_err error) {
	if IsMock {
		return _class.RemoveFromOtherConfigMock(sessionID, self, key)
	}	
	_method := "GPU_group.remove_from_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertGPUGroupRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func GPUGroupClassAddToOtherConfigMockDefault(sessionID SessionRef, self GPUGroupRef, key string, value string) (_err error) {
	log.Println("GPUGroup.AddToOtherConfig not mocked")
	_err = errors.New("GPUGroup.AddToOtherConfig not mocked")
	return
}

var GPUGroupClassAddToOtherConfigMockedCallback = GPUGroupClassAddToOtherConfigMockDefault

func (_class GPUGroupClass) AddToOtherConfigMock(sessionID SessionRef, self GPUGroupRef, key string, value string) (_err error) {
	return GPUGroupClassAddToOtherConfigMockedCallback(sessionID, self, key, value)
}
// Add the given key-value pair to the other_config field of the given GPU_group.
func (_class GPUGroupClass) AddToOtherConfig(sessionID SessionRef, self GPUGroupRef, key string, value string) (_err error) {
	if IsMock {
		return _class.AddToOtherConfigMock(sessionID, self, key, value)
	}	
	_method := "GPU_group.add_to_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertGPUGroupRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func GPUGroupClassSetOtherConfigMockDefault(sessionID SessionRef, self GPUGroupRef, value map[string]string) (_err error) {
	log.Println("GPUGroup.SetOtherConfig not mocked")
	_err = errors.New("GPUGroup.SetOtherConfig not mocked")
	return
}

var GPUGroupClassSetOtherConfigMockedCallback = GPUGroupClassSetOtherConfigMockDefault

func (_class GPUGroupClass) SetOtherConfigMock(sessionID SessionRef, self GPUGroupRef, value map[string]string) (_err error) {
	return GPUGroupClassSetOtherConfigMockedCallback(sessionID, self, value)
}
// Set the other_config field of the given GPU_group.
func (_class GPUGroupClass) SetOtherConfig(sessionID SessionRef, self GPUGroupRef, value map[string]string) (_err error) {
	if IsMock {
		return _class.SetOtherConfigMock(sessionID, self, value)
	}	
	_method := "GPU_group.set_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertGPUGroupRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func GPUGroupClassSetNameDescriptionMockDefault(sessionID SessionRef, self GPUGroupRef, value string) (_err error) {
	log.Println("GPUGroup.SetNameDescription not mocked")
	_err = errors.New("GPUGroup.SetNameDescription not mocked")
	return
}

var GPUGroupClassSetNameDescriptionMockedCallback = GPUGroupClassSetNameDescriptionMockDefault

func (_class GPUGroupClass) SetNameDescriptionMock(sessionID SessionRef, self GPUGroupRef, value string) (_err error) {
	return GPUGroupClassSetNameDescriptionMockedCallback(sessionID, self, value)
}
// Set the name/description field of the given GPU_group.
func (_class GPUGroupClass) SetNameDescription(sessionID SessionRef, self GPUGroupRef, value string) (_err error) {
	if IsMock {
		return _class.SetNameDescriptionMock(sessionID, self, value)
	}	
	_method := "GPU_group.set_name_description"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertGPUGroupRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func GPUGroupClassSetNameLabelMockDefault(sessionID SessionRef, self GPUGroupRef, value string) (_err error) {
	log.Println("GPUGroup.SetNameLabel not mocked")
	_err = errors.New("GPUGroup.SetNameLabel not mocked")
	return
}

var GPUGroupClassSetNameLabelMockedCallback = GPUGroupClassSetNameLabelMockDefault

func (_class GPUGroupClass) SetNameLabelMock(sessionID SessionRef, self GPUGroupRef, value string) (_err error) {
	return GPUGroupClassSetNameLabelMockedCallback(sessionID, self, value)
}
// Set the name/label field of the given GPU_group.
func (_class GPUGroupClass) SetNameLabel(sessionID SessionRef, self GPUGroupRef, value string) (_err error) {
	if IsMock {
		return _class.SetNameLabelMock(sessionID, self, value)
	}	
	_method := "GPU_group.set_name_label"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertGPUGroupRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func GPUGroupClassGetEnabledVGPUTypesMockDefault(sessionID SessionRef, self GPUGroupRef) (_retval []VGPUTypeRef, _err error) {
	log.Println("GPUGroup.GetEnabledVGPUTypes not mocked")
	_err = errors.New("GPUGroup.GetEnabledVGPUTypes not mocked")
	return
}

var GPUGroupClassGetEnabledVGPUTypesMockedCallback = GPUGroupClassGetEnabledVGPUTypesMockDefault

func (_class GPUGroupClass) GetEnabledVGPUTypesMock(sessionID SessionRef, self GPUGroupRef) (_retval []VGPUTypeRef, _err error) {
	return GPUGroupClassGetEnabledVGPUTypesMockedCallback(sessionID, self)
}
// Get the enabled_VGPU_types field of the given GPU_group.
func (_class GPUGroupClass) GetEnabledVGPUTypes(sessionID SessionRef, self GPUGroupRef) (_retval []VGPUTypeRef, _err error) {
	if IsMock {
		return _class.GetEnabledVGPUTypesMock(sessionID, self)
	}	
	_method := "GPU_group.get_enabled_VGPU_types"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertGPUGroupRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVGPUTypeRefSetToGo(_method + " -> ", _result.Value)
	return
}


func GPUGroupClassGetSupportedVGPUTypesMockDefault(sessionID SessionRef, self GPUGroupRef) (_retval []VGPUTypeRef, _err error) {
	log.Println("GPUGroup.GetSupportedVGPUTypes not mocked")
	_err = errors.New("GPUGroup.GetSupportedVGPUTypes not mocked")
	return
}

var GPUGroupClassGetSupportedVGPUTypesMockedCallback = GPUGroupClassGetSupportedVGPUTypesMockDefault

func (_class GPUGroupClass) GetSupportedVGPUTypesMock(sessionID SessionRef, self GPUGroupRef) (_retval []VGPUTypeRef, _err error) {
	return GPUGroupClassGetSupportedVGPUTypesMockedCallback(sessionID, self)
}
// Get the supported_VGPU_types field of the given GPU_group.
func (_class GPUGroupClass) GetSupportedVGPUTypes(sessionID SessionRef, self GPUGroupRef) (_retval []VGPUTypeRef, _err error) {
	if IsMock {
		return _class.GetSupportedVGPUTypesMock(sessionID, self)
	}	
	_method := "GPU_group.get_supported_VGPU_types"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertGPUGroupRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVGPUTypeRefSetToGo(_method + " -> ", _result.Value)
	return
}


func GPUGroupClassGetAllocationAlgorithmMockDefault(sessionID SessionRef, self GPUGroupRef) (_retval AllocationAlgorithm, _err error) {
	log.Println("GPUGroup.GetAllocationAlgorithm not mocked")
	_err = errors.New("GPUGroup.GetAllocationAlgorithm not mocked")
	return
}

var GPUGroupClassGetAllocationAlgorithmMockedCallback = GPUGroupClassGetAllocationAlgorithmMockDefault

func (_class GPUGroupClass) GetAllocationAlgorithmMock(sessionID SessionRef, self GPUGroupRef) (_retval AllocationAlgorithm, _err error) {
	return GPUGroupClassGetAllocationAlgorithmMockedCallback(sessionID, self)
}
// Get the allocation_algorithm field of the given GPU_group.
func (_class GPUGroupClass) GetAllocationAlgorithm(sessionID SessionRef, self GPUGroupRef) (_retval AllocationAlgorithm, _err error) {
	if IsMock {
		return _class.GetAllocationAlgorithmMock(sessionID, self)
	}	
	_method := "GPU_group.get_allocation_algorithm"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertGPUGroupRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertEnumAllocationAlgorithmToGo(_method + " -> ", _result.Value)
	return
}


func GPUGroupClassGetOtherConfigMockDefault(sessionID SessionRef, self GPUGroupRef) (_retval map[string]string, _err error) {
	log.Println("GPUGroup.GetOtherConfig not mocked")
	_err = errors.New("GPUGroup.GetOtherConfig not mocked")
	return
}

var GPUGroupClassGetOtherConfigMockedCallback = GPUGroupClassGetOtherConfigMockDefault

func (_class GPUGroupClass) GetOtherConfigMock(sessionID SessionRef, self GPUGroupRef) (_retval map[string]string, _err error) {
	return GPUGroupClassGetOtherConfigMockedCallback(sessionID, self)
}
// Get the other_config field of the given GPU_group.
func (_class GPUGroupClass) GetOtherConfig(sessionID SessionRef, self GPUGroupRef) (_retval map[string]string, _err error) {
	if IsMock {
		return _class.GetOtherConfigMock(sessionID, self)
	}	
	_method := "GPU_group.get_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertGPUGroupRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func GPUGroupClassGetGPUTypesMockDefault(sessionID SessionRef, self GPUGroupRef) (_retval []string, _err error) {
	log.Println("GPUGroup.GetGPUTypes not mocked")
	_err = errors.New("GPUGroup.GetGPUTypes not mocked")
	return
}

var GPUGroupClassGetGPUTypesMockedCallback = GPUGroupClassGetGPUTypesMockDefault

func (_class GPUGroupClass) GetGPUTypesMock(sessionID SessionRef, self GPUGroupRef) (_retval []string, _err error) {
	return GPUGroupClassGetGPUTypesMockedCallback(sessionID, self)
}
// Get the GPU_types field of the given GPU_group.
func (_class GPUGroupClass) GetGPUTypes(sessionID SessionRef, self GPUGroupRef) (_retval []string, _err error) {
	if IsMock {
		return _class.GetGPUTypesMock(sessionID, self)
	}	
	_method := "GPU_group.get_GPU_types"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertGPUGroupRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func GPUGroupClassGetVGPUsMockDefault(sessionID SessionRef, self GPUGroupRef) (_retval []VGPURef, _err error) {
	log.Println("GPUGroup.GetVGPUs not mocked")
	_err = errors.New("GPUGroup.GetVGPUs not mocked")
	return
}

var GPUGroupClassGetVGPUsMockedCallback = GPUGroupClassGetVGPUsMockDefault

func (_class GPUGroupClass) GetVGPUsMock(sessionID SessionRef, self GPUGroupRef) (_retval []VGPURef, _err error) {
	return GPUGroupClassGetVGPUsMockedCallback(sessionID, self)
}
// Get the VGPUs field of the given GPU_group.
func (_class GPUGroupClass) GetVGPUs(sessionID SessionRef, self GPUGroupRef) (_retval []VGPURef, _err error) {
	if IsMock {
		return _class.GetVGPUsMock(sessionID, self)
	}	
	_method := "GPU_group.get_VGPUs"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertGPUGroupRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVGPURefSetToGo(_method + " -> ", _result.Value)
	return
}


func GPUGroupClassGetPGPUsMockDefault(sessionID SessionRef, self GPUGroupRef) (_retval []PGPURef, _err error) {
	log.Println("GPUGroup.GetPGPUs not mocked")
	_err = errors.New("GPUGroup.GetPGPUs not mocked")
	return
}

var GPUGroupClassGetPGPUsMockedCallback = GPUGroupClassGetPGPUsMockDefault

func (_class GPUGroupClass) GetPGPUsMock(sessionID SessionRef, self GPUGroupRef) (_retval []PGPURef, _err error) {
	return GPUGroupClassGetPGPUsMockedCallback(sessionID, self)
}
// Get the PGPUs field of the given GPU_group.
func (_class GPUGroupClass) GetPGPUs(sessionID SessionRef, self GPUGroupRef) (_retval []PGPURef, _err error) {
	if IsMock {
		return _class.GetPGPUsMock(sessionID, self)
	}	
	_method := "GPU_group.get_PGPUs"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertGPUGroupRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertPGPURefSetToGo(_method + " -> ", _result.Value)
	return
}


func GPUGroupClassGetNameDescriptionMockDefault(sessionID SessionRef, self GPUGroupRef) (_retval string, _err error) {
	log.Println("GPUGroup.GetNameDescription not mocked")
	_err = errors.New("GPUGroup.GetNameDescription not mocked")
	return
}

var GPUGroupClassGetNameDescriptionMockedCallback = GPUGroupClassGetNameDescriptionMockDefault

func (_class GPUGroupClass) GetNameDescriptionMock(sessionID SessionRef, self GPUGroupRef) (_retval string, _err error) {
	return GPUGroupClassGetNameDescriptionMockedCallback(sessionID, self)
}
// Get the name/description field of the given GPU_group.
func (_class GPUGroupClass) GetNameDescription(sessionID SessionRef, self GPUGroupRef) (_retval string, _err error) {
	if IsMock {
		return _class.GetNameDescriptionMock(sessionID, self)
	}	
	_method := "GPU_group.get_name_description"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertGPUGroupRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func GPUGroupClassGetNameLabelMockDefault(sessionID SessionRef, self GPUGroupRef) (_retval string, _err error) {
	log.Println("GPUGroup.GetNameLabel not mocked")
	_err = errors.New("GPUGroup.GetNameLabel not mocked")
	return
}

var GPUGroupClassGetNameLabelMockedCallback = GPUGroupClassGetNameLabelMockDefault

func (_class GPUGroupClass) GetNameLabelMock(sessionID SessionRef, self GPUGroupRef) (_retval string, _err error) {
	return GPUGroupClassGetNameLabelMockedCallback(sessionID, self)
}
// Get the name/label field of the given GPU_group.
func (_class GPUGroupClass) GetNameLabel(sessionID SessionRef, self GPUGroupRef) (_retval string, _err error) {
	if IsMock {
		return _class.GetNameLabelMock(sessionID, self)
	}	
	_method := "GPU_group.get_name_label"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertGPUGroupRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func GPUGroupClassGetUUIDMockDefault(sessionID SessionRef, self GPUGroupRef) (_retval string, _err error) {
	log.Println("GPUGroup.GetUUID not mocked")
	_err = errors.New("GPUGroup.GetUUID not mocked")
	return
}

var GPUGroupClassGetUUIDMockedCallback = GPUGroupClassGetUUIDMockDefault

func (_class GPUGroupClass) GetUUIDMock(sessionID SessionRef, self GPUGroupRef) (_retval string, _err error) {
	return GPUGroupClassGetUUIDMockedCallback(sessionID, self)
}
// Get the uuid field of the given GPU_group.
func (_class GPUGroupClass) GetUUID(sessionID SessionRef, self GPUGroupRef) (_retval string, _err error) {
	if IsMock {
		return _class.GetUUIDMock(sessionID, self)
	}	
	_method := "GPU_group.get_uuid"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertGPUGroupRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func GPUGroupClassGetByNameLabelMockDefault(sessionID SessionRef, label string) (_retval []GPUGroupRef, _err error) {
	log.Println("GPUGroup.GetByNameLabel not mocked")
	_err = errors.New("GPUGroup.GetByNameLabel not mocked")
	return
}

var GPUGroupClassGetByNameLabelMockedCallback = GPUGroupClassGetByNameLabelMockDefault

func (_class GPUGroupClass) GetByNameLabelMock(sessionID SessionRef, label string) (_retval []GPUGroupRef, _err error) {
	return GPUGroupClassGetByNameLabelMockedCallback(sessionID, label)
}
// Get all the GPU_group instances with the given label.
func (_class GPUGroupClass) GetByNameLabel(sessionID SessionRef, label string) (_retval []GPUGroupRef, _err error) {
	if IsMock {
		return _class.GetByNameLabelMock(sessionID, label)
	}	
	_method := "GPU_group.get_by_name_label"
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
	_retval, _err = convertGPUGroupRefSetToGo(_method + " -> ", _result.Value)
	return
}


func GPUGroupClassGetByUUIDMockDefault(sessionID SessionRef, uuid string) (_retval GPUGroupRef, _err error) {
	log.Println("GPUGroup.GetByUUID not mocked")
	_err = errors.New("GPUGroup.GetByUUID not mocked")
	return
}

var GPUGroupClassGetByUUIDMockedCallback = GPUGroupClassGetByUUIDMockDefault

func (_class GPUGroupClass) GetByUUIDMock(sessionID SessionRef, uuid string) (_retval GPUGroupRef, _err error) {
	return GPUGroupClassGetByUUIDMockedCallback(sessionID, uuid)
}
// Get a reference to the GPU_group instance with the specified UUID.
func (_class GPUGroupClass) GetByUUID(sessionID SessionRef, uuid string) (_retval GPUGroupRef, _err error) {
	if IsMock {
		return _class.GetByUUIDMock(sessionID, uuid)
	}	
	_method := "GPU_group.get_by_uuid"
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
	_retval, _err = convertGPUGroupRefToGo(_method + " -> ", _result.Value)
	return
}


func GPUGroupClassGetRecordMockDefault(sessionID SessionRef, self GPUGroupRef) (_retval GPUGroupRecord, _err error) {
	log.Println("GPUGroup.GetRecord not mocked")
	_err = errors.New("GPUGroup.GetRecord not mocked")
	return
}

var GPUGroupClassGetRecordMockedCallback = GPUGroupClassGetRecordMockDefault

func (_class GPUGroupClass) GetRecordMock(sessionID SessionRef, self GPUGroupRef) (_retval GPUGroupRecord, _err error) {
	return GPUGroupClassGetRecordMockedCallback(sessionID, self)
}
// Get a record containing the current state of the given GPU_group.
func (_class GPUGroupClass) GetRecord(sessionID SessionRef, self GPUGroupRef) (_retval GPUGroupRecord, _err error) {
	if IsMock {
		return _class.GetRecordMock(sessionID, self)
	}	
	_method := "GPU_group.get_record"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertGPUGroupRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertGPUGroupRecordToGo(_method + " -> ", _result.Value)
	return
}
