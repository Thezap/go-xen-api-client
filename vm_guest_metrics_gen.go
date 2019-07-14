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

type TristateType string

const (
  // Known to be true
	TristateTypeYes TristateType = "yes"
  // Known to be false
	TristateTypeNo TristateType = "no"
  // Unknown or unspecified
	TristateTypeUnspecified TristateType = "unspecified"
)

type VMGuestMetricsRecord struct {
  // Unique identifier/object reference
	UUID string
  // version of the OS
	OSVersion map[string]string
  // version of the PV drivers
	PVDriversVersion map[string]string
  // Logically equivalent to PV_drivers_detected
	PVDriversUpToDate bool
  // This field exists but has no data. Use the memory and memory_internal_free RRD data-sources instead.
	Memory map[string]string
  // This field exists but has no data.
	Disks map[string]string
  // network configuration
	Networks map[string]string
  // anything else
	Other map[string]string
  // Time at which this information was last updated
	LastUpdated time.Time
  // additional configuration
	OtherConfig map[string]string
  // True if the guest is sending heartbeat messages via the guest agent
	Live bool
  // The guest's statement of whether it supports VBD hotplug, i.e. whether it is capable of responding immediately to instantiation of a new VBD by bringing online a new PV block device. If the guest states that it is not capable, then the VBD plug and unplug operations will not be allowed while the guest is running.
	CanUseHotplugVbd TristateType
  // The guest's statement of whether it supports VIF hotplug, i.e. whether it is capable of responding immediately to instantiation of a new VIF by bringing online a new PV network device. If the guest states that it is not capable, then the VIF plug and unplug operations will not be allowed while the guest is running.
	CanUseHotplugVif TristateType
  // At least one of the guest's devices has successfully connected to the backend.
	PVDriversDetected bool
}

type VMGuestMetricsRef string

// The metrics reported by the guest (as opposed to inferred from outside)
type VMGuestMetricsClass struct {
	client *Client
}


func VMGuestMetricsClassGetAllRecordsMockDefault(sessionID SessionRef) (_retval map[VMGuestMetricsRef]VMGuestMetricsRecord, _err error) {
	log.Println("VMGuestMetrics.GetAllRecords not mocked")
	_err = errors.New("VMGuestMetrics.GetAllRecords not mocked")
	return
}

var VMGuestMetricsClassGetAllRecordsMockedCallback = VMGuestMetricsClassGetAllRecordsMockDefault

func (_class VMGuestMetricsClass) GetAllRecordsMock(sessionID SessionRef) (_retval map[VMGuestMetricsRef]VMGuestMetricsRecord, _err error) {
	return VMGuestMetricsClassGetAllRecordsMockedCallback(sessionID)
}
// Return a map of VM_guest_metrics references to VM_guest_metrics records for all VM_guest_metrics instances known to the system.
func (_class VMGuestMetricsClass) GetAllRecords(sessionID SessionRef) (_retval map[VMGuestMetricsRef]VMGuestMetricsRecord, _err error) {
	if IsMock {
		return _class.GetAllRecordsMock(sessionID)
	}	
	_method := "VM_guest_metrics.get_all_records"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVMGuestMetricsRefToVMGuestMetricsRecordMapToGo(_method + " -> ", _result.Value)
	return
}


func VMGuestMetricsClassGetAllMockDefault(sessionID SessionRef) (_retval []VMGuestMetricsRef, _err error) {
	log.Println("VMGuestMetrics.GetAll not mocked")
	_err = errors.New("VMGuestMetrics.GetAll not mocked")
	return
}

var VMGuestMetricsClassGetAllMockedCallback = VMGuestMetricsClassGetAllMockDefault

func (_class VMGuestMetricsClass) GetAllMock(sessionID SessionRef) (_retval []VMGuestMetricsRef, _err error) {
	return VMGuestMetricsClassGetAllMockedCallback(sessionID)
}
// Return a list of all the VM_guest_metrics instances known to the system.
func (_class VMGuestMetricsClass) GetAll(sessionID SessionRef) (_retval []VMGuestMetricsRef, _err error) {
	if IsMock {
		return _class.GetAllMock(sessionID)
	}	
	_method := "VM_guest_metrics.get_all"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVMGuestMetricsRefSetToGo(_method + " -> ", _result.Value)
	return
}


func VMGuestMetricsClassRemoveFromOtherConfigMockDefault(sessionID SessionRef, self VMGuestMetricsRef, key string) (_err error) {
	log.Println("VMGuestMetrics.RemoveFromOtherConfig not mocked")
	_err = errors.New("VMGuestMetrics.RemoveFromOtherConfig not mocked")
	return
}

var VMGuestMetricsClassRemoveFromOtherConfigMockedCallback = VMGuestMetricsClassRemoveFromOtherConfigMockDefault

func (_class VMGuestMetricsClass) RemoveFromOtherConfigMock(sessionID SessionRef, self VMGuestMetricsRef, key string) (_err error) {
	return VMGuestMetricsClassRemoveFromOtherConfigMockedCallback(sessionID, self, key)
}
// Remove the given key and its corresponding value from the other_config field of the given VM_guest_metrics.  If the key is not in that Map, then do nothing.
func (_class VMGuestMetricsClass) RemoveFromOtherConfig(sessionID SessionRef, self VMGuestMetricsRef, key string) (_err error) {
	if IsMock {
		return _class.RemoveFromOtherConfigMock(sessionID, self, key)
	}	
	_method := "VM_guest_metrics.remove_from_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMGuestMetricsRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func VMGuestMetricsClassAddToOtherConfigMockDefault(sessionID SessionRef, self VMGuestMetricsRef, key string, value string) (_err error) {
	log.Println("VMGuestMetrics.AddToOtherConfig not mocked")
	_err = errors.New("VMGuestMetrics.AddToOtherConfig not mocked")
	return
}

var VMGuestMetricsClassAddToOtherConfigMockedCallback = VMGuestMetricsClassAddToOtherConfigMockDefault

func (_class VMGuestMetricsClass) AddToOtherConfigMock(sessionID SessionRef, self VMGuestMetricsRef, key string, value string) (_err error) {
	return VMGuestMetricsClassAddToOtherConfigMockedCallback(sessionID, self, key, value)
}
// Add the given key-value pair to the other_config field of the given VM_guest_metrics.
func (_class VMGuestMetricsClass) AddToOtherConfig(sessionID SessionRef, self VMGuestMetricsRef, key string, value string) (_err error) {
	if IsMock {
		return _class.AddToOtherConfigMock(sessionID, self, key, value)
	}	
	_method := "VM_guest_metrics.add_to_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMGuestMetricsRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func VMGuestMetricsClassSetOtherConfigMockDefault(sessionID SessionRef, self VMGuestMetricsRef, value map[string]string) (_err error) {
	log.Println("VMGuestMetrics.SetOtherConfig not mocked")
	_err = errors.New("VMGuestMetrics.SetOtherConfig not mocked")
	return
}

var VMGuestMetricsClassSetOtherConfigMockedCallback = VMGuestMetricsClassSetOtherConfigMockDefault

func (_class VMGuestMetricsClass) SetOtherConfigMock(sessionID SessionRef, self VMGuestMetricsRef, value map[string]string) (_err error) {
	return VMGuestMetricsClassSetOtherConfigMockedCallback(sessionID, self, value)
}
// Set the other_config field of the given VM_guest_metrics.
func (_class VMGuestMetricsClass) SetOtherConfig(sessionID SessionRef, self VMGuestMetricsRef, value map[string]string) (_err error) {
	if IsMock {
		return _class.SetOtherConfigMock(sessionID, self, value)
	}	
	_method := "VM_guest_metrics.set_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMGuestMetricsRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func VMGuestMetricsClassGetPVDriversDetectedMockDefault(sessionID SessionRef, self VMGuestMetricsRef) (_retval bool, _err error) {
	log.Println("VMGuestMetrics.GetPVDriversDetected not mocked")
	_err = errors.New("VMGuestMetrics.GetPVDriversDetected not mocked")
	return
}

var VMGuestMetricsClassGetPVDriversDetectedMockedCallback = VMGuestMetricsClassGetPVDriversDetectedMockDefault

func (_class VMGuestMetricsClass) GetPVDriversDetectedMock(sessionID SessionRef, self VMGuestMetricsRef) (_retval bool, _err error) {
	return VMGuestMetricsClassGetPVDriversDetectedMockedCallback(sessionID, self)
}
// Get the PV_drivers_detected field of the given VM_guest_metrics.
func (_class VMGuestMetricsClass) GetPVDriversDetected(sessionID SessionRef, self VMGuestMetricsRef) (_retval bool, _err error) {
	if IsMock {
		return _class.GetPVDriversDetectedMock(sessionID, self)
	}	
	_method := "VM_guest_metrics.get_PV_drivers_detected"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMGuestMetricsRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func VMGuestMetricsClassGetCanUseHotplugVifMockDefault(sessionID SessionRef, self VMGuestMetricsRef) (_retval TristateType, _err error) {
	log.Println("VMGuestMetrics.GetCanUseHotplugVif not mocked")
	_err = errors.New("VMGuestMetrics.GetCanUseHotplugVif not mocked")
	return
}

var VMGuestMetricsClassGetCanUseHotplugVifMockedCallback = VMGuestMetricsClassGetCanUseHotplugVifMockDefault

func (_class VMGuestMetricsClass) GetCanUseHotplugVifMock(sessionID SessionRef, self VMGuestMetricsRef) (_retval TristateType, _err error) {
	return VMGuestMetricsClassGetCanUseHotplugVifMockedCallback(sessionID, self)
}
// Get the can_use_hotplug_vif field of the given VM_guest_metrics.
func (_class VMGuestMetricsClass) GetCanUseHotplugVif(sessionID SessionRef, self VMGuestMetricsRef) (_retval TristateType, _err error) {
	if IsMock {
		return _class.GetCanUseHotplugVifMock(sessionID, self)
	}	
	_method := "VM_guest_metrics.get_can_use_hotplug_vif"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMGuestMetricsRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertEnumTristateTypeToGo(_method + " -> ", _result.Value)
	return
}


func VMGuestMetricsClassGetCanUseHotplugVbdMockDefault(sessionID SessionRef, self VMGuestMetricsRef) (_retval TristateType, _err error) {
	log.Println("VMGuestMetrics.GetCanUseHotplugVbd not mocked")
	_err = errors.New("VMGuestMetrics.GetCanUseHotplugVbd not mocked")
	return
}

var VMGuestMetricsClassGetCanUseHotplugVbdMockedCallback = VMGuestMetricsClassGetCanUseHotplugVbdMockDefault

func (_class VMGuestMetricsClass) GetCanUseHotplugVbdMock(sessionID SessionRef, self VMGuestMetricsRef) (_retval TristateType, _err error) {
	return VMGuestMetricsClassGetCanUseHotplugVbdMockedCallback(sessionID, self)
}
// Get the can_use_hotplug_vbd field of the given VM_guest_metrics.
func (_class VMGuestMetricsClass) GetCanUseHotplugVbd(sessionID SessionRef, self VMGuestMetricsRef) (_retval TristateType, _err error) {
	if IsMock {
		return _class.GetCanUseHotplugVbdMock(sessionID, self)
	}	
	_method := "VM_guest_metrics.get_can_use_hotplug_vbd"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMGuestMetricsRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertEnumTristateTypeToGo(_method + " -> ", _result.Value)
	return
}


func VMGuestMetricsClassGetLiveMockDefault(sessionID SessionRef, self VMGuestMetricsRef) (_retval bool, _err error) {
	log.Println("VMGuestMetrics.GetLive not mocked")
	_err = errors.New("VMGuestMetrics.GetLive not mocked")
	return
}

var VMGuestMetricsClassGetLiveMockedCallback = VMGuestMetricsClassGetLiveMockDefault

func (_class VMGuestMetricsClass) GetLiveMock(sessionID SessionRef, self VMGuestMetricsRef) (_retval bool, _err error) {
	return VMGuestMetricsClassGetLiveMockedCallback(sessionID, self)
}
// Get the live field of the given VM_guest_metrics.
func (_class VMGuestMetricsClass) GetLive(sessionID SessionRef, self VMGuestMetricsRef) (_retval bool, _err error) {
	if IsMock {
		return _class.GetLiveMock(sessionID, self)
	}	
	_method := "VM_guest_metrics.get_live"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMGuestMetricsRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func VMGuestMetricsClassGetOtherConfigMockDefault(sessionID SessionRef, self VMGuestMetricsRef) (_retval map[string]string, _err error) {
	log.Println("VMGuestMetrics.GetOtherConfig not mocked")
	_err = errors.New("VMGuestMetrics.GetOtherConfig not mocked")
	return
}

var VMGuestMetricsClassGetOtherConfigMockedCallback = VMGuestMetricsClassGetOtherConfigMockDefault

func (_class VMGuestMetricsClass) GetOtherConfigMock(sessionID SessionRef, self VMGuestMetricsRef) (_retval map[string]string, _err error) {
	return VMGuestMetricsClassGetOtherConfigMockedCallback(sessionID, self)
}
// Get the other_config field of the given VM_guest_metrics.
func (_class VMGuestMetricsClass) GetOtherConfig(sessionID SessionRef, self VMGuestMetricsRef) (_retval map[string]string, _err error) {
	if IsMock {
		return _class.GetOtherConfigMock(sessionID, self)
	}	
	_method := "VM_guest_metrics.get_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMGuestMetricsRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func VMGuestMetricsClassGetLastUpdatedMockDefault(sessionID SessionRef, self VMGuestMetricsRef) (_retval time.Time, _err error) {
	log.Println("VMGuestMetrics.GetLastUpdated not mocked")
	_err = errors.New("VMGuestMetrics.GetLastUpdated not mocked")
	return
}

var VMGuestMetricsClassGetLastUpdatedMockedCallback = VMGuestMetricsClassGetLastUpdatedMockDefault

func (_class VMGuestMetricsClass) GetLastUpdatedMock(sessionID SessionRef, self VMGuestMetricsRef) (_retval time.Time, _err error) {
	return VMGuestMetricsClassGetLastUpdatedMockedCallback(sessionID, self)
}
// Get the last_updated field of the given VM_guest_metrics.
func (_class VMGuestMetricsClass) GetLastUpdated(sessionID SessionRef, self VMGuestMetricsRef) (_retval time.Time, _err error) {
	if IsMock {
		return _class.GetLastUpdatedMock(sessionID, self)
	}	
	_method := "VM_guest_metrics.get_last_updated"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMGuestMetricsRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func VMGuestMetricsClassGetOtherMockDefault(sessionID SessionRef, self VMGuestMetricsRef) (_retval map[string]string, _err error) {
	log.Println("VMGuestMetrics.GetOther not mocked")
	_err = errors.New("VMGuestMetrics.GetOther not mocked")
	return
}

var VMGuestMetricsClassGetOtherMockedCallback = VMGuestMetricsClassGetOtherMockDefault

func (_class VMGuestMetricsClass) GetOtherMock(sessionID SessionRef, self VMGuestMetricsRef) (_retval map[string]string, _err error) {
	return VMGuestMetricsClassGetOtherMockedCallback(sessionID, self)
}
// Get the other field of the given VM_guest_metrics.
func (_class VMGuestMetricsClass) GetOther(sessionID SessionRef, self VMGuestMetricsRef) (_retval map[string]string, _err error) {
	if IsMock {
		return _class.GetOtherMock(sessionID, self)
	}	
	_method := "VM_guest_metrics.get_other"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMGuestMetricsRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func VMGuestMetricsClassGetNetworksMockDefault(sessionID SessionRef, self VMGuestMetricsRef) (_retval map[string]string, _err error) {
	log.Println("VMGuestMetrics.GetNetworks not mocked")
	_err = errors.New("VMGuestMetrics.GetNetworks not mocked")
	return
}

var VMGuestMetricsClassGetNetworksMockedCallback = VMGuestMetricsClassGetNetworksMockDefault

func (_class VMGuestMetricsClass) GetNetworksMock(sessionID SessionRef, self VMGuestMetricsRef) (_retval map[string]string, _err error) {
	return VMGuestMetricsClassGetNetworksMockedCallback(sessionID, self)
}
// Get the networks field of the given VM_guest_metrics.
func (_class VMGuestMetricsClass) GetNetworks(sessionID SessionRef, self VMGuestMetricsRef) (_retval map[string]string, _err error) {
	if IsMock {
		return _class.GetNetworksMock(sessionID, self)
	}	
	_method := "VM_guest_metrics.get_networks"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMGuestMetricsRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func VMGuestMetricsClassGetDisksMockDefault(sessionID SessionRef, self VMGuestMetricsRef) (_retval map[string]string, _err error) {
	log.Println("VMGuestMetrics.GetDisks not mocked")
	_err = errors.New("VMGuestMetrics.GetDisks not mocked")
	return
}

var VMGuestMetricsClassGetDisksMockedCallback = VMGuestMetricsClassGetDisksMockDefault

func (_class VMGuestMetricsClass) GetDisksMock(sessionID SessionRef, self VMGuestMetricsRef) (_retval map[string]string, _err error) {
	return VMGuestMetricsClassGetDisksMockedCallback(sessionID, self)
}
// Get the disks field of the given VM_guest_metrics.
func (_class VMGuestMetricsClass) GetDisks(sessionID SessionRef, self VMGuestMetricsRef) (_retval map[string]string, _err error) {
	if IsMock {
		return _class.GetDisksMock(sessionID, self)
	}	
	_method := "VM_guest_metrics.get_disks"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMGuestMetricsRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func VMGuestMetricsClassGetMemoryMockDefault(sessionID SessionRef, self VMGuestMetricsRef) (_retval map[string]string, _err error) {
	log.Println("VMGuestMetrics.GetMemory not mocked")
	_err = errors.New("VMGuestMetrics.GetMemory not mocked")
	return
}

var VMGuestMetricsClassGetMemoryMockedCallback = VMGuestMetricsClassGetMemoryMockDefault

func (_class VMGuestMetricsClass) GetMemoryMock(sessionID SessionRef, self VMGuestMetricsRef) (_retval map[string]string, _err error) {
	return VMGuestMetricsClassGetMemoryMockedCallback(sessionID, self)
}
// Get the memory field of the given VM_guest_metrics.
func (_class VMGuestMetricsClass) GetMemory(sessionID SessionRef, self VMGuestMetricsRef) (_retval map[string]string, _err error) {
	if IsMock {
		return _class.GetMemoryMock(sessionID, self)
	}	
	_method := "VM_guest_metrics.get_memory"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMGuestMetricsRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func VMGuestMetricsClassGetPVDriversUpToDateMockDefault(sessionID SessionRef, self VMGuestMetricsRef) (_retval bool, _err error) {
	log.Println("VMGuestMetrics.GetPVDriversUpToDate not mocked")
	_err = errors.New("VMGuestMetrics.GetPVDriversUpToDate not mocked")
	return
}

var VMGuestMetricsClassGetPVDriversUpToDateMockedCallback = VMGuestMetricsClassGetPVDriversUpToDateMockDefault

func (_class VMGuestMetricsClass) GetPVDriversUpToDateMock(sessionID SessionRef, self VMGuestMetricsRef) (_retval bool, _err error) {
	return VMGuestMetricsClassGetPVDriversUpToDateMockedCallback(sessionID, self)
}
// Get the PV_drivers_up_to_date field of the given VM_guest_metrics.
func (_class VMGuestMetricsClass) GetPVDriversUpToDate(sessionID SessionRef, self VMGuestMetricsRef) (_retval bool, _err error) {
	if IsMock {
		return _class.GetPVDriversUpToDateMock(sessionID, self)
	}	
	_method := "VM_guest_metrics.get_PV_drivers_up_to_date"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMGuestMetricsRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func VMGuestMetricsClassGetPVDriversVersionMockDefault(sessionID SessionRef, self VMGuestMetricsRef) (_retval map[string]string, _err error) {
	log.Println("VMGuestMetrics.GetPVDriversVersion not mocked")
	_err = errors.New("VMGuestMetrics.GetPVDriversVersion not mocked")
	return
}

var VMGuestMetricsClassGetPVDriversVersionMockedCallback = VMGuestMetricsClassGetPVDriversVersionMockDefault

func (_class VMGuestMetricsClass) GetPVDriversVersionMock(sessionID SessionRef, self VMGuestMetricsRef) (_retval map[string]string, _err error) {
	return VMGuestMetricsClassGetPVDriversVersionMockedCallback(sessionID, self)
}
// Get the PV_drivers_version field of the given VM_guest_metrics.
func (_class VMGuestMetricsClass) GetPVDriversVersion(sessionID SessionRef, self VMGuestMetricsRef) (_retval map[string]string, _err error) {
	if IsMock {
		return _class.GetPVDriversVersionMock(sessionID, self)
	}	
	_method := "VM_guest_metrics.get_PV_drivers_version"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMGuestMetricsRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func VMGuestMetricsClassGetOSVersionMockDefault(sessionID SessionRef, self VMGuestMetricsRef) (_retval map[string]string, _err error) {
	log.Println("VMGuestMetrics.GetOSVersion not mocked")
	_err = errors.New("VMGuestMetrics.GetOSVersion not mocked")
	return
}

var VMGuestMetricsClassGetOSVersionMockedCallback = VMGuestMetricsClassGetOSVersionMockDefault

func (_class VMGuestMetricsClass) GetOSVersionMock(sessionID SessionRef, self VMGuestMetricsRef) (_retval map[string]string, _err error) {
	return VMGuestMetricsClassGetOSVersionMockedCallback(sessionID, self)
}
// Get the os_version field of the given VM_guest_metrics.
func (_class VMGuestMetricsClass) GetOSVersion(sessionID SessionRef, self VMGuestMetricsRef) (_retval map[string]string, _err error) {
	if IsMock {
		return _class.GetOSVersionMock(sessionID, self)
	}	
	_method := "VM_guest_metrics.get_os_version"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMGuestMetricsRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func VMGuestMetricsClassGetUUIDMockDefault(sessionID SessionRef, self VMGuestMetricsRef) (_retval string, _err error) {
	log.Println("VMGuestMetrics.GetUUID not mocked")
	_err = errors.New("VMGuestMetrics.GetUUID not mocked")
	return
}

var VMGuestMetricsClassGetUUIDMockedCallback = VMGuestMetricsClassGetUUIDMockDefault

func (_class VMGuestMetricsClass) GetUUIDMock(sessionID SessionRef, self VMGuestMetricsRef) (_retval string, _err error) {
	return VMGuestMetricsClassGetUUIDMockedCallback(sessionID, self)
}
// Get the uuid field of the given VM_guest_metrics.
func (_class VMGuestMetricsClass) GetUUID(sessionID SessionRef, self VMGuestMetricsRef) (_retval string, _err error) {
	if IsMock {
		return _class.GetUUIDMock(sessionID, self)
	}	
	_method := "VM_guest_metrics.get_uuid"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMGuestMetricsRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func VMGuestMetricsClassGetByUUIDMockDefault(sessionID SessionRef, uuid string) (_retval VMGuestMetricsRef, _err error) {
	log.Println("VMGuestMetrics.GetByUUID not mocked")
	_err = errors.New("VMGuestMetrics.GetByUUID not mocked")
	return
}

var VMGuestMetricsClassGetByUUIDMockedCallback = VMGuestMetricsClassGetByUUIDMockDefault

func (_class VMGuestMetricsClass) GetByUUIDMock(sessionID SessionRef, uuid string) (_retval VMGuestMetricsRef, _err error) {
	return VMGuestMetricsClassGetByUUIDMockedCallback(sessionID, uuid)
}
// Get a reference to the VM_guest_metrics instance with the specified UUID.
func (_class VMGuestMetricsClass) GetByUUID(sessionID SessionRef, uuid string) (_retval VMGuestMetricsRef, _err error) {
	if IsMock {
		return _class.GetByUUIDMock(sessionID, uuid)
	}	
	_method := "VM_guest_metrics.get_by_uuid"
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
	_retval, _err = convertVMGuestMetricsRefToGo(_method + " -> ", _result.Value)
	return
}


func VMGuestMetricsClassGetRecordMockDefault(sessionID SessionRef, self VMGuestMetricsRef) (_retval VMGuestMetricsRecord, _err error) {
	log.Println("VMGuestMetrics.GetRecord not mocked")
	_err = errors.New("VMGuestMetrics.GetRecord not mocked")
	return
}

var VMGuestMetricsClassGetRecordMockedCallback = VMGuestMetricsClassGetRecordMockDefault

func (_class VMGuestMetricsClass) GetRecordMock(sessionID SessionRef, self VMGuestMetricsRef) (_retval VMGuestMetricsRecord, _err error) {
	return VMGuestMetricsClassGetRecordMockedCallback(sessionID, self)
}
// Get a record containing the current state of the given VM_guest_metrics.
func (_class VMGuestMetricsClass) GetRecord(sessionID SessionRef, self VMGuestMetricsRef) (_retval VMGuestMetricsRecord, _err error) {
	if IsMock {
		return _class.GetRecordMock(sessionID, self)
	}	
	_method := "VM_guest_metrics.get_record"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMGuestMetricsRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVMGuestMetricsRecordToGo(_method + " -> ", _result.Value)
	return
}
