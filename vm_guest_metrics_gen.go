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

func (_class VMGuestMetricsClass) GetAllRecords__mock(sessionID SessionRef) (_retval map[VMGuestMetricsRef]VMGuestMetricsRecord, _err error) {
	log.Println("VMGuestMetrics.GetAllRecords not mocked")
	_err = errors.New("VMGuestMetrics.GetAllRecords not mocked")
	return
}
// Return a map of VM_guest_metrics references to VM_guest_metrics records for all VM_guest_metrics instances known to the system.
func (_class VMGuestMetricsClass) GetAllRecords(sessionID SessionRef) (_retval map[VMGuestMetricsRef]VMGuestMetricsRecord, _err error) {
	if (IsMock) {
		return _class.GetAllRecords__mock(sessionID)
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

func (_class VMGuestMetricsClass) GetAll__mock(sessionID SessionRef) (_retval []VMGuestMetricsRef, _err error) {
	log.Println("VMGuestMetrics.GetAll not mocked")
	_err = errors.New("VMGuestMetrics.GetAll not mocked")
	return
}
// Return a list of all the VM_guest_metrics instances known to the system.
func (_class VMGuestMetricsClass) GetAll(sessionID SessionRef) (_retval []VMGuestMetricsRef, _err error) {
	if (IsMock) {
		return _class.GetAll__mock(sessionID)
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

func (_class VMGuestMetricsClass) RemoveFromOtherConfig__mock(sessionID SessionRef, self VMGuestMetricsRef, key string) (_err error) {
	log.Println("VMGuestMetrics.RemoveFromOtherConfig not mocked")
	_err = errors.New("VMGuestMetrics.RemoveFromOtherConfig not mocked")
	return
}
// Remove the given key and its corresponding value from the other_config field of the given VM_guest_metrics.  If the key is not in that Map, then do nothing.
func (_class VMGuestMetricsClass) RemoveFromOtherConfig(sessionID SessionRef, self VMGuestMetricsRef, key string) (_err error) {
	if (IsMock) {
		return _class.RemoveFromOtherConfig__mock(sessionID, self, key)
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

func (_class VMGuestMetricsClass) AddToOtherConfig__mock(sessionID SessionRef, self VMGuestMetricsRef, key string, value string) (_err error) {
	log.Println("VMGuestMetrics.AddToOtherConfig not mocked")
	_err = errors.New("VMGuestMetrics.AddToOtherConfig not mocked")
	return
}
// Add the given key-value pair to the other_config field of the given VM_guest_metrics.
func (_class VMGuestMetricsClass) AddToOtherConfig(sessionID SessionRef, self VMGuestMetricsRef, key string, value string) (_err error) {
	if (IsMock) {
		return _class.AddToOtherConfig__mock(sessionID, self, key, value)
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

func (_class VMGuestMetricsClass) SetOtherConfig__mock(sessionID SessionRef, self VMGuestMetricsRef, value map[string]string) (_err error) {
	log.Println("VMGuestMetrics.SetOtherConfig not mocked")
	_err = errors.New("VMGuestMetrics.SetOtherConfig not mocked")
	return
}
// Set the other_config field of the given VM_guest_metrics.
func (_class VMGuestMetricsClass) SetOtherConfig(sessionID SessionRef, self VMGuestMetricsRef, value map[string]string) (_err error) {
	if (IsMock) {
		return _class.SetOtherConfig__mock(sessionID, self, value)
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

func (_class VMGuestMetricsClass) GetPVDriversDetected__mock(sessionID SessionRef, self VMGuestMetricsRef) (_retval bool, _err error) {
	log.Println("VMGuestMetrics.GetPVDriversDetected not mocked")
	_err = errors.New("VMGuestMetrics.GetPVDriversDetected not mocked")
	return
}
// Get the PV_drivers_detected field of the given VM_guest_metrics.
func (_class VMGuestMetricsClass) GetPVDriversDetected(sessionID SessionRef, self VMGuestMetricsRef) (_retval bool, _err error) {
	if (IsMock) {
		return _class.GetPVDriversDetected__mock(sessionID, self)
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

func (_class VMGuestMetricsClass) GetCanUseHotplugVif__mock(sessionID SessionRef, self VMGuestMetricsRef) (_retval TristateType, _err error) {
	log.Println("VMGuestMetrics.GetCanUseHotplugVif not mocked")
	_err = errors.New("VMGuestMetrics.GetCanUseHotplugVif not mocked")
	return
}
// Get the can_use_hotplug_vif field of the given VM_guest_metrics.
func (_class VMGuestMetricsClass) GetCanUseHotplugVif(sessionID SessionRef, self VMGuestMetricsRef) (_retval TristateType, _err error) {
	if (IsMock) {
		return _class.GetCanUseHotplugVif__mock(sessionID, self)
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

func (_class VMGuestMetricsClass) GetCanUseHotplugVbd__mock(sessionID SessionRef, self VMGuestMetricsRef) (_retval TristateType, _err error) {
	log.Println("VMGuestMetrics.GetCanUseHotplugVbd not mocked")
	_err = errors.New("VMGuestMetrics.GetCanUseHotplugVbd not mocked")
	return
}
// Get the can_use_hotplug_vbd field of the given VM_guest_metrics.
func (_class VMGuestMetricsClass) GetCanUseHotplugVbd(sessionID SessionRef, self VMGuestMetricsRef) (_retval TristateType, _err error) {
	if (IsMock) {
		return _class.GetCanUseHotplugVbd__mock(sessionID, self)
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

func (_class VMGuestMetricsClass) GetLive__mock(sessionID SessionRef, self VMGuestMetricsRef) (_retval bool, _err error) {
	log.Println("VMGuestMetrics.GetLive not mocked")
	_err = errors.New("VMGuestMetrics.GetLive not mocked")
	return
}
// Get the live field of the given VM_guest_metrics.
func (_class VMGuestMetricsClass) GetLive(sessionID SessionRef, self VMGuestMetricsRef) (_retval bool, _err error) {
	if (IsMock) {
		return _class.GetLive__mock(sessionID, self)
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

func (_class VMGuestMetricsClass) GetOtherConfig__mock(sessionID SessionRef, self VMGuestMetricsRef) (_retval map[string]string, _err error) {
	log.Println("VMGuestMetrics.GetOtherConfig not mocked")
	_err = errors.New("VMGuestMetrics.GetOtherConfig not mocked")
	return
}
// Get the other_config field of the given VM_guest_metrics.
func (_class VMGuestMetricsClass) GetOtherConfig(sessionID SessionRef, self VMGuestMetricsRef) (_retval map[string]string, _err error) {
	if (IsMock) {
		return _class.GetOtherConfig__mock(sessionID, self)
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

func (_class VMGuestMetricsClass) GetLastUpdated__mock(sessionID SessionRef, self VMGuestMetricsRef) (_retval time.Time, _err error) {
	log.Println("VMGuestMetrics.GetLastUpdated not mocked")
	_err = errors.New("VMGuestMetrics.GetLastUpdated not mocked")
	return
}
// Get the last_updated field of the given VM_guest_metrics.
func (_class VMGuestMetricsClass) GetLastUpdated(sessionID SessionRef, self VMGuestMetricsRef) (_retval time.Time, _err error) {
	if (IsMock) {
		return _class.GetLastUpdated__mock(sessionID, self)
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

func (_class VMGuestMetricsClass) GetOther__mock(sessionID SessionRef, self VMGuestMetricsRef) (_retval map[string]string, _err error) {
	log.Println("VMGuestMetrics.GetOther not mocked")
	_err = errors.New("VMGuestMetrics.GetOther not mocked")
	return
}
// Get the other field of the given VM_guest_metrics.
func (_class VMGuestMetricsClass) GetOther(sessionID SessionRef, self VMGuestMetricsRef) (_retval map[string]string, _err error) {
	if (IsMock) {
		return _class.GetOther__mock(sessionID, self)
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

func (_class VMGuestMetricsClass) GetNetworks__mock(sessionID SessionRef, self VMGuestMetricsRef) (_retval map[string]string, _err error) {
	log.Println("VMGuestMetrics.GetNetworks not mocked")
	_err = errors.New("VMGuestMetrics.GetNetworks not mocked")
	return
}
// Get the networks field of the given VM_guest_metrics.
func (_class VMGuestMetricsClass) GetNetworks(sessionID SessionRef, self VMGuestMetricsRef) (_retval map[string]string, _err error) {
	if (IsMock) {
		return _class.GetNetworks__mock(sessionID, self)
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

func (_class VMGuestMetricsClass) GetDisks__mock(sessionID SessionRef, self VMGuestMetricsRef) (_retval map[string]string, _err error) {
	log.Println("VMGuestMetrics.GetDisks not mocked")
	_err = errors.New("VMGuestMetrics.GetDisks not mocked")
	return
}
// Get the disks field of the given VM_guest_metrics.
func (_class VMGuestMetricsClass) GetDisks(sessionID SessionRef, self VMGuestMetricsRef) (_retval map[string]string, _err error) {
	if (IsMock) {
		return _class.GetDisks__mock(sessionID, self)
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

func (_class VMGuestMetricsClass) GetMemory__mock(sessionID SessionRef, self VMGuestMetricsRef) (_retval map[string]string, _err error) {
	log.Println("VMGuestMetrics.GetMemory not mocked")
	_err = errors.New("VMGuestMetrics.GetMemory not mocked")
	return
}
// Get the memory field of the given VM_guest_metrics.
func (_class VMGuestMetricsClass) GetMemory(sessionID SessionRef, self VMGuestMetricsRef) (_retval map[string]string, _err error) {
	if (IsMock) {
		return _class.GetMemory__mock(sessionID, self)
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

func (_class VMGuestMetricsClass) GetPVDriversUpToDate__mock(sessionID SessionRef, self VMGuestMetricsRef) (_retval bool, _err error) {
	log.Println("VMGuestMetrics.GetPVDriversUpToDate not mocked")
	_err = errors.New("VMGuestMetrics.GetPVDriversUpToDate not mocked")
	return
}
// Get the PV_drivers_up_to_date field of the given VM_guest_metrics.
func (_class VMGuestMetricsClass) GetPVDriversUpToDate(sessionID SessionRef, self VMGuestMetricsRef) (_retval bool, _err error) {
	if (IsMock) {
		return _class.GetPVDriversUpToDate__mock(sessionID, self)
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

func (_class VMGuestMetricsClass) GetPVDriversVersion__mock(sessionID SessionRef, self VMGuestMetricsRef) (_retval map[string]string, _err error) {
	log.Println("VMGuestMetrics.GetPVDriversVersion not mocked")
	_err = errors.New("VMGuestMetrics.GetPVDriversVersion not mocked")
	return
}
// Get the PV_drivers_version field of the given VM_guest_metrics.
func (_class VMGuestMetricsClass) GetPVDriversVersion(sessionID SessionRef, self VMGuestMetricsRef) (_retval map[string]string, _err error) {
	if (IsMock) {
		return _class.GetPVDriversVersion__mock(sessionID, self)
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

func (_class VMGuestMetricsClass) GetOSVersion__mock(sessionID SessionRef, self VMGuestMetricsRef) (_retval map[string]string, _err error) {
	log.Println("VMGuestMetrics.GetOSVersion not mocked")
	_err = errors.New("VMGuestMetrics.GetOSVersion not mocked")
	return
}
// Get the os_version field of the given VM_guest_metrics.
func (_class VMGuestMetricsClass) GetOSVersion(sessionID SessionRef, self VMGuestMetricsRef) (_retval map[string]string, _err error) {
	if (IsMock) {
		return _class.GetOSVersion__mock(sessionID, self)
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

func (_class VMGuestMetricsClass) GetUUID__mock(sessionID SessionRef, self VMGuestMetricsRef) (_retval string, _err error) {
	log.Println("VMGuestMetrics.GetUUID not mocked")
	_err = errors.New("VMGuestMetrics.GetUUID not mocked")
	return
}
// Get the uuid field of the given VM_guest_metrics.
func (_class VMGuestMetricsClass) GetUUID(sessionID SessionRef, self VMGuestMetricsRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetUUID__mock(sessionID, self)
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

func (_class VMGuestMetricsClass) GetByUUID__mock(sessionID SessionRef, uuid string) (_retval VMGuestMetricsRef, _err error) {
	log.Println("VMGuestMetrics.GetByUUID not mocked")
	_err = errors.New("VMGuestMetrics.GetByUUID not mocked")
	return
}
// Get a reference to the VM_guest_metrics instance with the specified UUID.
func (_class VMGuestMetricsClass) GetByUUID(sessionID SessionRef, uuid string) (_retval VMGuestMetricsRef, _err error) {
	if (IsMock) {
		return _class.GetByUUID__mock(sessionID, uuid)
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

func (_class VMGuestMetricsClass) GetRecord__mock(sessionID SessionRef, self VMGuestMetricsRef) (_retval VMGuestMetricsRecord, _err error) {
	log.Println("VMGuestMetrics.GetRecord not mocked")
	_err = errors.New("VMGuestMetrics.GetRecord not mocked")
	return
}
// Get a record containing the current state of the given VM_guest_metrics.
func (_class VMGuestMetricsClass) GetRecord(sessionID SessionRef, self VMGuestMetricsRef) (_retval VMGuestMetricsRecord, _err error) {
	if (IsMock) {
		return _class.GetRecord__mock(sessionID, self)
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
