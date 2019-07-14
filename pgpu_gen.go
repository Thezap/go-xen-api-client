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

type PgpuDom0Access string

const (
  // dom0 can access this device as normal
	PgpuDom0AccessEnabled PgpuDom0Access = "enabled"
  // On host reboot dom0 will be blocked from accessing this device
	PgpuDom0AccessDisableOnReboot PgpuDom0Access = "disable_on_reboot"
  // dom0 cannot access this device
	PgpuDom0AccessDisabled PgpuDom0Access = "disabled"
  // On host reboot dom0 will be allowed to access this device
	PgpuDom0AccessEnableOnReboot PgpuDom0Access = "enable_on_reboot"
)

type PGPURecord struct {
  // Unique identifier/object reference
	UUID string
  // Link to underlying PCI device
	PCI PCIRef
  // GPU group the pGPU is contained in
	GPUGroup GPUGroupRef
  // Host that owns the GPU
	Host HostRef
  // Additional configuration
	OtherConfig map[string]string
  // List of VGPU types supported by the underlying hardware
	SupportedVGPUTypes []VGPUTypeRef
  // List of VGPU types which have been enabled for this PGPU
	EnabledVGPUTypes []VGPUTypeRef
  // List of VGPUs running on this PGPU
	ResidentVGPUs []VGPURef
  // A map relating each VGPU type supported on this GPU to the maximum number of VGPUs of that type which can run simultaneously on this GPU
	SupportedVGPUMaxCapacities map[VGPUTypeRef]int
  // The accessibility of this device from dom0
	Dom0Access PgpuDom0Access
  // Is this device the system display device
	IsSystemDisplayDevice bool
  // PGPU metadata to determine whether a VGPU can migrate between two PGPUs
	CompatibilityMetadata map[string]string
}

type PGPURef string

// A physical GPU (pGPU)
type PGPUClass struct {
	client *Client
}


var PGPUClass_GetAllRecordsMockedCallback = func (sessionID SessionRef) (_retval map[PGPURef]PGPURecord, _err error) {
	log.Println("PGPU.GetAllRecords not mocked")
	_err = errors.New("PGPU.GetAllRecords not mocked")
	return
}

func (_class PGPUClass) GetAllRecordsMock(sessionID SessionRef) (_retval map[PGPURef]PGPURecord, _err error) {
	return PGPUClass_GetAllRecordsMockedCallback(sessionID)
}
// Return a map of PGPU references to PGPU records for all PGPUs known to the system.
func (_class PGPUClass) GetAllRecords(sessionID SessionRef) (_retval map[PGPURef]PGPURecord, _err error) {
	if (IsMock) {
		return _class.GetAllRecordsMock(sessionID)
	}	
	_method := "PGPU.get_all_records"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertPGPURefToPGPURecordMapToGo(_method + " -> ", _result.Value)
	return
}


var PGPUClass_GetAllMockedCallback = func (sessionID SessionRef) (_retval []PGPURef, _err error) {
	log.Println("PGPU.GetAll not mocked")
	_err = errors.New("PGPU.GetAll not mocked")
	return
}

func (_class PGPUClass) GetAllMock(sessionID SessionRef) (_retval []PGPURef, _err error) {
	return PGPUClass_GetAllMockedCallback(sessionID)
}
// Return a list of all the PGPUs known to the system.
func (_class PGPUClass) GetAll(sessionID SessionRef) (_retval []PGPURef, _err error) {
	if (IsMock) {
		return _class.GetAllMock(sessionID)
	}	
	_method := "PGPU.get_all"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertPGPURefSetToGo(_method + " -> ", _result.Value)
	return
}


var PGPUClass_DisableDom0AccessMockedCallback = func (sessionID SessionRef, self PGPURef) (_retval PgpuDom0Access, _err error) {
	log.Println("PGPU.DisableDom0Access not mocked")
	_err = errors.New("PGPU.DisableDom0Access not mocked")
	return
}

func (_class PGPUClass) DisableDom0AccessMock(sessionID SessionRef, self PGPURef) (_retval PgpuDom0Access, _err error) {
	return PGPUClass_DisableDom0AccessMockedCallback(sessionID, self)
}
// 
func (_class PGPUClass) DisableDom0Access(sessionID SessionRef, self PGPURef) (_retval PgpuDom0Access, _err error) {
	if (IsMock) {
		return _class.DisableDom0AccessMock(sessionID, self)
	}	
	_method := "PGPU.disable_dom0_access"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPGPURefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertEnumPgpuDom0AccessToGo(_method + " -> ", _result.Value)
	return
}


var PGPUClass_EnableDom0AccessMockedCallback = func (sessionID SessionRef, self PGPURef) (_retval PgpuDom0Access, _err error) {
	log.Println("PGPU.EnableDom0Access not mocked")
	_err = errors.New("PGPU.EnableDom0Access not mocked")
	return
}

func (_class PGPUClass) EnableDom0AccessMock(sessionID SessionRef, self PGPURef) (_retval PgpuDom0Access, _err error) {
	return PGPUClass_EnableDom0AccessMockedCallback(sessionID, self)
}
// 
func (_class PGPUClass) EnableDom0Access(sessionID SessionRef, self PGPURef) (_retval PgpuDom0Access, _err error) {
	if (IsMock) {
		return _class.EnableDom0AccessMock(sessionID, self)
	}	
	_method := "PGPU.enable_dom0_access"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPGPURefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertEnumPgpuDom0AccessToGo(_method + " -> ", _result.Value)
	return
}


var PGPUClass_GetRemainingCapacityMockedCallback = func (sessionID SessionRef, self PGPURef, vgpuType VGPUTypeRef) (_retval int, _err error) {
	log.Println("PGPU.GetRemainingCapacity not mocked")
	_err = errors.New("PGPU.GetRemainingCapacity not mocked")
	return
}

func (_class PGPUClass) GetRemainingCapacityMock(sessionID SessionRef, self PGPURef, vgpuType VGPUTypeRef) (_retval int, _err error) {
	return PGPUClass_GetRemainingCapacityMockedCallback(sessionID, self, vgpuType)
}
// 
func (_class PGPUClass) GetRemainingCapacity(sessionID SessionRef, self PGPURef, vgpuType VGPUTypeRef) (_retval int, _err error) {
	if (IsMock) {
		return _class.GetRemainingCapacityMock(sessionID, self, vgpuType)
	}	
	_method := "PGPU.get_remaining_capacity"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPGPURefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var PGPUClass_SetGPUGroupMockedCallback = func (sessionID SessionRef, self PGPURef, value GPUGroupRef) (_err error) {
	log.Println("PGPU.SetGPUGroup not mocked")
	_err = errors.New("PGPU.SetGPUGroup not mocked")
	return
}

func (_class PGPUClass) SetGPUGroupMock(sessionID SessionRef, self PGPURef, value GPUGroupRef) (_err error) {
	return PGPUClass_SetGPUGroupMockedCallback(sessionID, self, value)
}
// 
func (_class PGPUClass) SetGPUGroup(sessionID SessionRef, self PGPURef, value GPUGroupRef) (_err error) {
	if (IsMock) {
		return _class.SetGPUGroupMock(sessionID, self, value)
	}	
	_method := "PGPU.set_GPU_group"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPGPURefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertGPUGroupRefToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}


var PGPUClass_SetEnabledVGPUTypesMockedCallback = func (sessionID SessionRef, self PGPURef, value []VGPUTypeRef) (_err error) {
	log.Println("PGPU.SetEnabledVGPUTypes not mocked")
	_err = errors.New("PGPU.SetEnabledVGPUTypes not mocked")
	return
}

func (_class PGPUClass) SetEnabledVGPUTypesMock(sessionID SessionRef, self PGPURef, value []VGPUTypeRef) (_err error) {
	return PGPUClass_SetEnabledVGPUTypesMockedCallback(sessionID, self, value)
}
// 
func (_class PGPUClass) SetEnabledVGPUTypes(sessionID SessionRef, self PGPURef, value []VGPUTypeRef) (_err error) {
	if (IsMock) {
		return _class.SetEnabledVGPUTypesMock(sessionID, self, value)
	}	
	_method := "PGPU.set_enabled_VGPU_types"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPGPURefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertVGPUTypeRefSetToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}


var PGPUClass_RemoveEnabledVGPUTypesMockedCallback = func (sessionID SessionRef, self PGPURef, value VGPUTypeRef) (_err error) {
	log.Println("PGPU.RemoveEnabledVGPUTypes not mocked")
	_err = errors.New("PGPU.RemoveEnabledVGPUTypes not mocked")
	return
}

func (_class PGPUClass) RemoveEnabledVGPUTypesMock(sessionID SessionRef, self PGPURef, value VGPUTypeRef) (_err error) {
	return PGPUClass_RemoveEnabledVGPUTypesMockedCallback(sessionID, self, value)
}
// 
func (_class PGPUClass) RemoveEnabledVGPUTypes(sessionID SessionRef, self PGPURef, value VGPUTypeRef) (_err error) {
	if (IsMock) {
		return _class.RemoveEnabledVGPUTypesMock(sessionID, self, value)
	}	
	_method := "PGPU.remove_enabled_VGPU_types"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPGPURefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertVGPUTypeRefToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}


var PGPUClass_AddEnabledVGPUTypesMockedCallback = func (sessionID SessionRef, self PGPURef, value VGPUTypeRef) (_err error) {
	log.Println("PGPU.AddEnabledVGPUTypes not mocked")
	_err = errors.New("PGPU.AddEnabledVGPUTypes not mocked")
	return
}

func (_class PGPUClass) AddEnabledVGPUTypesMock(sessionID SessionRef, self PGPURef, value VGPUTypeRef) (_err error) {
	return PGPUClass_AddEnabledVGPUTypesMockedCallback(sessionID, self, value)
}
// 
func (_class PGPUClass) AddEnabledVGPUTypes(sessionID SessionRef, self PGPURef, value VGPUTypeRef) (_err error) {
	if (IsMock) {
		return _class.AddEnabledVGPUTypesMock(sessionID, self, value)
	}	
	_method := "PGPU.add_enabled_VGPU_types"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPGPURefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertVGPUTypeRefToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}


var PGPUClass_RemoveFromOtherConfigMockedCallback = func (sessionID SessionRef, self PGPURef, key string) (_err error) {
	log.Println("PGPU.RemoveFromOtherConfig not mocked")
	_err = errors.New("PGPU.RemoveFromOtherConfig not mocked")
	return
}

func (_class PGPUClass) RemoveFromOtherConfigMock(sessionID SessionRef, self PGPURef, key string) (_err error) {
	return PGPUClass_RemoveFromOtherConfigMockedCallback(sessionID, self, key)
}
// Remove the given key and its corresponding value from the other_config field of the given PGPU.  If the key is not in that Map, then do nothing.
func (_class PGPUClass) RemoveFromOtherConfig(sessionID SessionRef, self PGPURef, key string) (_err error) {
	if (IsMock) {
		return _class.RemoveFromOtherConfigMock(sessionID, self, key)
	}	
	_method := "PGPU.remove_from_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPGPURefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var PGPUClass_AddToOtherConfigMockedCallback = func (sessionID SessionRef, self PGPURef, key string, value string) (_err error) {
	log.Println("PGPU.AddToOtherConfig not mocked")
	_err = errors.New("PGPU.AddToOtherConfig not mocked")
	return
}

func (_class PGPUClass) AddToOtherConfigMock(sessionID SessionRef, self PGPURef, key string, value string) (_err error) {
	return PGPUClass_AddToOtherConfigMockedCallback(sessionID, self, key, value)
}
// Add the given key-value pair to the other_config field of the given PGPU.
func (_class PGPUClass) AddToOtherConfig(sessionID SessionRef, self PGPURef, key string, value string) (_err error) {
	if (IsMock) {
		return _class.AddToOtherConfigMock(sessionID, self, key, value)
	}	
	_method := "PGPU.add_to_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPGPURefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var PGPUClass_SetOtherConfigMockedCallback = func (sessionID SessionRef, self PGPURef, value map[string]string) (_err error) {
	log.Println("PGPU.SetOtherConfig not mocked")
	_err = errors.New("PGPU.SetOtherConfig not mocked")
	return
}

func (_class PGPUClass) SetOtherConfigMock(sessionID SessionRef, self PGPURef, value map[string]string) (_err error) {
	return PGPUClass_SetOtherConfigMockedCallback(sessionID, self, value)
}
// Set the other_config field of the given PGPU.
func (_class PGPUClass) SetOtherConfig(sessionID SessionRef, self PGPURef, value map[string]string) (_err error) {
	if (IsMock) {
		return _class.SetOtherConfigMock(sessionID, self, value)
	}	
	_method := "PGPU.set_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPGPURefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var PGPUClass_GetCompatibilityMetadataMockedCallback = func (sessionID SessionRef, self PGPURef) (_retval map[string]string, _err error) {
	log.Println("PGPU.GetCompatibilityMetadata not mocked")
	_err = errors.New("PGPU.GetCompatibilityMetadata not mocked")
	return
}

func (_class PGPUClass) GetCompatibilityMetadataMock(sessionID SessionRef, self PGPURef) (_retval map[string]string, _err error) {
	return PGPUClass_GetCompatibilityMetadataMockedCallback(sessionID, self)
}
// Get the compatibility_metadata field of the given PGPU.
func (_class PGPUClass) GetCompatibilityMetadata(sessionID SessionRef, self PGPURef) (_retval map[string]string, _err error) {
	if (IsMock) {
		return _class.GetCompatibilityMetadataMock(sessionID, self)
	}	
	_method := "PGPU.get_compatibility_metadata"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPGPURefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var PGPUClass_GetIsSystemDisplayDeviceMockedCallback = func (sessionID SessionRef, self PGPURef) (_retval bool, _err error) {
	log.Println("PGPU.GetIsSystemDisplayDevice not mocked")
	_err = errors.New("PGPU.GetIsSystemDisplayDevice not mocked")
	return
}

func (_class PGPUClass) GetIsSystemDisplayDeviceMock(sessionID SessionRef, self PGPURef) (_retval bool, _err error) {
	return PGPUClass_GetIsSystemDisplayDeviceMockedCallback(sessionID, self)
}
// Get the is_system_display_device field of the given PGPU.
func (_class PGPUClass) GetIsSystemDisplayDevice(sessionID SessionRef, self PGPURef) (_retval bool, _err error) {
	if (IsMock) {
		return _class.GetIsSystemDisplayDeviceMock(sessionID, self)
	}	
	_method := "PGPU.get_is_system_display_device"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPGPURefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var PGPUClass_GetDom0AccessMockedCallback = func (sessionID SessionRef, self PGPURef) (_retval PgpuDom0Access, _err error) {
	log.Println("PGPU.GetDom0Access not mocked")
	_err = errors.New("PGPU.GetDom0Access not mocked")
	return
}

func (_class PGPUClass) GetDom0AccessMock(sessionID SessionRef, self PGPURef) (_retval PgpuDom0Access, _err error) {
	return PGPUClass_GetDom0AccessMockedCallback(sessionID, self)
}
// Get the dom0_access field of the given PGPU.
func (_class PGPUClass) GetDom0Access(sessionID SessionRef, self PGPURef) (_retval PgpuDom0Access, _err error) {
	if (IsMock) {
		return _class.GetDom0AccessMock(sessionID, self)
	}	
	_method := "PGPU.get_dom0_access"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPGPURefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertEnumPgpuDom0AccessToGo(_method + " -> ", _result.Value)
	return
}


var PGPUClass_GetSupportedVGPUMaxCapacitiesMockedCallback = func (sessionID SessionRef, self PGPURef) (_retval map[VGPUTypeRef]int, _err error) {
	log.Println("PGPU.GetSupportedVGPUMaxCapacities not mocked")
	_err = errors.New("PGPU.GetSupportedVGPUMaxCapacities not mocked")
	return
}

func (_class PGPUClass) GetSupportedVGPUMaxCapacitiesMock(sessionID SessionRef, self PGPURef) (_retval map[VGPUTypeRef]int, _err error) {
	return PGPUClass_GetSupportedVGPUMaxCapacitiesMockedCallback(sessionID, self)
}
// Get the supported_VGPU_max_capacities field of the given PGPU.
func (_class PGPUClass) GetSupportedVGPUMaxCapacities(sessionID SessionRef, self PGPURef) (_retval map[VGPUTypeRef]int, _err error) {
	if (IsMock) {
		return _class.GetSupportedVGPUMaxCapacitiesMock(sessionID, self)
	}	
	_method := "PGPU.get_supported_VGPU_max_capacities"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPGPURefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVGPUTypeRefToIntMapToGo(_method + " -> ", _result.Value)
	return
}


var PGPUClass_GetResidentVGPUsMockedCallback = func (sessionID SessionRef, self PGPURef) (_retval []VGPURef, _err error) {
	log.Println("PGPU.GetResidentVGPUs not mocked")
	_err = errors.New("PGPU.GetResidentVGPUs not mocked")
	return
}

func (_class PGPUClass) GetResidentVGPUsMock(sessionID SessionRef, self PGPURef) (_retval []VGPURef, _err error) {
	return PGPUClass_GetResidentVGPUsMockedCallback(sessionID, self)
}
// Get the resident_VGPUs field of the given PGPU.
func (_class PGPUClass) GetResidentVGPUs(sessionID SessionRef, self PGPURef) (_retval []VGPURef, _err error) {
	if (IsMock) {
		return _class.GetResidentVGPUsMock(sessionID, self)
	}	
	_method := "PGPU.get_resident_VGPUs"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPGPURefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var PGPUClass_GetEnabledVGPUTypesMockedCallback = func (sessionID SessionRef, self PGPURef) (_retval []VGPUTypeRef, _err error) {
	log.Println("PGPU.GetEnabledVGPUTypes not mocked")
	_err = errors.New("PGPU.GetEnabledVGPUTypes not mocked")
	return
}

func (_class PGPUClass) GetEnabledVGPUTypesMock(sessionID SessionRef, self PGPURef) (_retval []VGPUTypeRef, _err error) {
	return PGPUClass_GetEnabledVGPUTypesMockedCallback(sessionID, self)
}
// Get the enabled_VGPU_types field of the given PGPU.
func (_class PGPUClass) GetEnabledVGPUTypes(sessionID SessionRef, self PGPURef) (_retval []VGPUTypeRef, _err error) {
	if (IsMock) {
		return _class.GetEnabledVGPUTypesMock(sessionID, self)
	}	
	_method := "PGPU.get_enabled_VGPU_types"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPGPURefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var PGPUClass_GetSupportedVGPUTypesMockedCallback = func (sessionID SessionRef, self PGPURef) (_retval []VGPUTypeRef, _err error) {
	log.Println("PGPU.GetSupportedVGPUTypes not mocked")
	_err = errors.New("PGPU.GetSupportedVGPUTypes not mocked")
	return
}

func (_class PGPUClass) GetSupportedVGPUTypesMock(sessionID SessionRef, self PGPURef) (_retval []VGPUTypeRef, _err error) {
	return PGPUClass_GetSupportedVGPUTypesMockedCallback(sessionID, self)
}
// Get the supported_VGPU_types field of the given PGPU.
func (_class PGPUClass) GetSupportedVGPUTypes(sessionID SessionRef, self PGPURef) (_retval []VGPUTypeRef, _err error) {
	if (IsMock) {
		return _class.GetSupportedVGPUTypesMock(sessionID, self)
	}	
	_method := "PGPU.get_supported_VGPU_types"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPGPURefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var PGPUClass_GetOtherConfigMockedCallback = func (sessionID SessionRef, self PGPURef) (_retval map[string]string, _err error) {
	log.Println("PGPU.GetOtherConfig not mocked")
	_err = errors.New("PGPU.GetOtherConfig not mocked")
	return
}

func (_class PGPUClass) GetOtherConfigMock(sessionID SessionRef, self PGPURef) (_retval map[string]string, _err error) {
	return PGPUClass_GetOtherConfigMockedCallback(sessionID, self)
}
// Get the other_config field of the given PGPU.
func (_class PGPUClass) GetOtherConfig(sessionID SessionRef, self PGPURef) (_retval map[string]string, _err error) {
	if (IsMock) {
		return _class.GetOtherConfigMock(sessionID, self)
	}	
	_method := "PGPU.get_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPGPURefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var PGPUClass_GetHostMockedCallback = func (sessionID SessionRef, self PGPURef) (_retval HostRef, _err error) {
	log.Println("PGPU.GetHost not mocked")
	_err = errors.New("PGPU.GetHost not mocked")
	return
}

func (_class PGPUClass) GetHostMock(sessionID SessionRef, self PGPURef) (_retval HostRef, _err error) {
	return PGPUClass_GetHostMockedCallback(sessionID, self)
}
// Get the host field of the given PGPU.
func (_class PGPUClass) GetHost(sessionID SessionRef, self PGPURef) (_retval HostRef, _err error) {
	if (IsMock) {
		return _class.GetHostMock(sessionID, self)
	}	
	_method := "PGPU.get_host"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPGPURefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var PGPUClass_GetGPUGroupMockedCallback = func (sessionID SessionRef, self PGPURef) (_retval GPUGroupRef, _err error) {
	log.Println("PGPU.GetGPUGroup not mocked")
	_err = errors.New("PGPU.GetGPUGroup not mocked")
	return
}

func (_class PGPUClass) GetGPUGroupMock(sessionID SessionRef, self PGPURef) (_retval GPUGroupRef, _err error) {
	return PGPUClass_GetGPUGroupMockedCallback(sessionID, self)
}
// Get the GPU_group field of the given PGPU.
func (_class PGPUClass) GetGPUGroup(sessionID SessionRef, self PGPURef) (_retval GPUGroupRef, _err error) {
	if (IsMock) {
		return _class.GetGPUGroupMock(sessionID, self)
	}	
	_method := "PGPU.get_GPU_group"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPGPURefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var PGPUClass_GetPCIMockedCallback = func (sessionID SessionRef, self PGPURef) (_retval PCIRef, _err error) {
	log.Println("PGPU.GetPCI not mocked")
	_err = errors.New("PGPU.GetPCI not mocked")
	return
}

func (_class PGPUClass) GetPCIMock(sessionID SessionRef, self PGPURef) (_retval PCIRef, _err error) {
	return PGPUClass_GetPCIMockedCallback(sessionID, self)
}
// Get the PCI field of the given PGPU.
func (_class PGPUClass) GetPCI(sessionID SessionRef, self PGPURef) (_retval PCIRef, _err error) {
	if (IsMock) {
		return _class.GetPCIMock(sessionID, self)
	}	
	_method := "PGPU.get_PCI"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPGPURefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertPCIRefToGo(_method + " -> ", _result.Value)
	return
}


var PGPUClass_GetUUIDMockedCallback = func (sessionID SessionRef, self PGPURef) (_retval string, _err error) {
	log.Println("PGPU.GetUUID not mocked")
	_err = errors.New("PGPU.GetUUID not mocked")
	return
}

func (_class PGPUClass) GetUUIDMock(sessionID SessionRef, self PGPURef) (_retval string, _err error) {
	return PGPUClass_GetUUIDMockedCallback(sessionID, self)
}
// Get the uuid field of the given PGPU.
func (_class PGPUClass) GetUUID(sessionID SessionRef, self PGPURef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetUUIDMock(sessionID, self)
	}	
	_method := "PGPU.get_uuid"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPGPURefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var PGPUClass_GetByUUIDMockedCallback = func (sessionID SessionRef, uuid string) (_retval PGPURef, _err error) {
	log.Println("PGPU.GetByUUID not mocked")
	_err = errors.New("PGPU.GetByUUID not mocked")
	return
}

func (_class PGPUClass) GetByUUIDMock(sessionID SessionRef, uuid string) (_retval PGPURef, _err error) {
	return PGPUClass_GetByUUIDMockedCallback(sessionID, uuid)
}
// Get a reference to the PGPU instance with the specified UUID.
func (_class PGPUClass) GetByUUID(sessionID SessionRef, uuid string) (_retval PGPURef, _err error) {
	if (IsMock) {
		return _class.GetByUUIDMock(sessionID, uuid)
	}	
	_method := "PGPU.get_by_uuid"
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
	_retval, _err = convertPGPURefToGo(_method + " -> ", _result.Value)
	return
}


var PGPUClass_GetRecordMockedCallback = func (sessionID SessionRef, self PGPURef) (_retval PGPURecord, _err error) {
	log.Println("PGPU.GetRecord not mocked")
	_err = errors.New("PGPU.GetRecord not mocked")
	return
}

func (_class PGPUClass) GetRecordMock(sessionID SessionRef, self PGPURef) (_retval PGPURecord, _err error) {
	return PGPUClass_GetRecordMockedCallback(sessionID, self)
}
// Get a record containing the current state of the given PGPU.
func (_class PGPUClass) GetRecord(sessionID SessionRef, self PGPURef) (_retval PGPURecord, _err error) {
	if (IsMock) {
		return _class.GetRecordMock(sessionID, self)
	}	
	_method := "PGPU.get_record"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPGPURefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertPGPURecordToGo(_method + " -> ", _result.Value)
	return
}
