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

type VMApplianceOperation string

const (
  // Start
	VMApplianceOperationStart VMApplianceOperation = "start"
  // Clean shutdown
	VMApplianceOperationCleanShutdown VMApplianceOperation = "clean_shutdown"
  // Hard shutdown
	VMApplianceOperationHardShutdown VMApplianceOperation = "hard_shutdown"
  // Shutdown
	VMApplianceOperationShutdown VMApplianceOperation = "shutdown"
)

type VMApplianceRecord struct {
  // Unique identifier/object reference
	UUID string
  // a human-readable name
	NameLabel string
  // a notes field containing human-readable description
	NameDescription string
  // list of the operations allowed in this state. This list is advisory only and the server state may have changed by the time this field is read by a client.
	AllowedOperations []VMApplianceOperation
  // links each of the running tasks using this object (by reference) to a current_operation enum which describes the nature of the task.
	CurrentOperations map[string]VMApplianceOperation
  // all VMs in this appliance
	VMs []VMRef
}

type VMApplianceRef string

// VM appliance
type VMApplianceClass struct {
	client *Client
}


var VMApplianceClass_GetAllRecordsMockedCallback = func (sessionID SessionRef) (_retval map[VMApplianceRef]VMApplianceRecord, _err error) {
	log.Println("VMAppliance.GetAllRecords not mocked")
	_err = errors.New("VMAppliance.GetAllRecords not mocked")
	return
}

func (_class VMApplianceClass) GetAllRecordsMock(sessionID SessionRef) (_retval map[VMApplianceRef]VMApplianceRecord, _err error) {
	return VMApplianceClass_GetAllRecordsMockedCallback(sessionID)
}
// Return a map of VM_appliance references to VM_appliance records for all VM_appliances known to the system.
func (_class VMApplianceClass) GetAllRecords(sessionID SessionRef) (_retval map[VMApplianceRef]VMApplianceRecord, _err error) {
	if (IsMock) {
		return _class.GetAllRecordsMock(sessionID)
	}	
	_method := "VM_appliance.get_all_records"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVMApplianceRefToVMApplianceRecordMapToGo(_method + " -> ", _result.Value)
	return
}


var VMApplianceClass_GetAllMockedCallback = func (sessionID SessionRef) (_retval []VMApplianceRef, _err error) {
	log.Println("VMAppliance.GetAll not mocked")
	_err = errors.New("VMAppliance.GetAll not mocked")
	return
}

func (_class VMApplianceClass) GetAllMock(sessionID SessionRef) (_retval []VMApplianceRef, _err error) {
	return VMApplianceClass_GetAllMockedCallback(sessionID)
}
// Return a list of all the VM_appliances known to the system.
func (_class VMApplianceClass) GetAll(sessionID SessionRef) (_retval []VMApplianceRef, _err error) {
	if (IsMock) {
		return _class.GetAllMock(sessionID)
	}	
	_method := "VM_appliance.get_all"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVMApplianceRefSetToGo(_method + " -> ", _result.Value)
	return
}


var VMApplianceClass_RecoverMockedCallback = func (sessionID SessionRef, self VMApplianceRef, sessionTo SessionRef, force bool) (_err error) {
	log.Println("VMAppliance.Recover not mocked")
	_err = errors.New("VMAppliance.Recover not mocked")
	return
}

func (_class VMApplianceClass) RecoverMock(sessionID SessionRef, self VMApplianceRef, sessionTo SessionRef, force bool) (_err error) {
	return VMApplianceClass_RecoverMockedCallback(sessionID, self, sessionTo, force)
}
// Recover the VM appliance
//
// Errors:
//  VM_REQUIRES_SR - You attempted to run a VM on a host which doesn't have access to an SR needed by the VM. The VM has at least one VBD attached to a VDI in the SR.
func (_class VMApplianceClass) Recover(sessionID SessionRef, self VMApplianceRef, sessionTo SessionRef, force bool) (_err error) {
	if (IsMock) {
		return _class.RecoverMock(sessionID, self, sessionTo, force)
	}	
	_method := "VM_appliance.recover"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMApplianceRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_sessionToArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_to"), sessionTo)
	if _err != nil {
		return
	}
	_forceArg, _err := convertBoolToXen(fmt.Sprintf("%s(%s)", _method, "force"), force)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _sessionToArg, _forceArg)
	return
}


var VMApplianceClass_GetSRsRequiredForRecoveryMockedCallback = func (sessionID SessionRef, self VMApplianceRef, sessionTo SessionRef) (_retval []SRRef, _err error) {
	log.Println("VMAppliance.GetSRsRequiredForRecovery not mocked")
	_err = errors.New("VMAppliance.GetSRsRequiredForRecovery not mocked")
	return
}

func (_class VMApplianceClass) GetSRsRequiredForRecoveryMock(sessionID SessionRef, self VMApplianceRef, sessionTo SessionRef) (_retval []SRRef, _err error) {
	return VMApplianceClass_GetSRsRequiredForRecoveryMockedCallback(sessionID, self, sessionTo)
}
// Get the list of SRs required by the VM appliance to recover.
func (_class VMApplianceClass) GetSRsRequiredForRecovery(sessionID SessionRef, self VMApplianceRef, sessionTo SessionRef) (_retval []SRRef, _err error) {
	if (IsMock) {
		return _class.GetSRsRequiredForRecoveryMock(sessionID, self, sessionTo)
	}	
	_method := "VM_appliance.get_SRs_required_for_recovery"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMApplianceRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_sessionToArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_to"), sessionTo)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg, _sessionToArg)
	if _err != nil {
		return
	}
	_retval, _err = convertSRRefSetToGo(_method + " -> ", _result.Value)
	return
}


var VMApplianceClass_AssertCanBeRecoveredMockedCallback = func (sessionID SessionRef, self VMApplianceRef, sessionTo SessionRef) (_err error) {
	log.Println("VMAppliance.AssertCanBeRecovered not mocked")
	_err = errors.New("VMAppliance.AssertCanBeRecovered not mocked")
	return
}

func (_class VMApplianceClass) AssertCanBeRecoveredMock(sessionID SessionRef, self VMApplianceRef, sessionTo SessionRef) (_err error) {
	return VMApplianceClass_AssertCanBeRecoveredMockedCallback(sessionID, self, sessionTo)
}
// Assert whether all SRs required to recover this VM appliance are available.
//
// Errors:
//  VM_REQUIRES_SR - You attempted to run a VM on a host which doesn't have access to an SR needed by the VM. The VM has at least one VBD attached to a VDI in the SR.
func (_class VMApplianceClass) AssertCanBeRecovered(sessionID SessionRef, self VMApplianceRef, sessionTo SessionRef) (_err error) {
	if (IsMock) {
		return _class.AssertCanBeRecoveredMock(sessionID, self, sessionTo)
	}	
	_method := "VM_appliance.assert_can_be_recovered"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMApplianceRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_sessionToArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_to"), sessionTo)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _sessionToArg)
	return
}


var VMApplianceClass_ShutdownMockedCallback = func (sessionID SessionRef, self VMApplianceRef) (_err error) {
	log.Println("VMAppliance.Shutdown not mocked")
	_err = errors.New("VMAppliance.Shutdown not mocked")
	return
}

func (_class VMApplianceClass) ShutdownMock(sessionID SessionRef, self VMApplianceRef) (_err error) {
	return VMApplianceClass_ShutdownMockedCallback(sessionID, self)
}
// For each VM in the appliance, try to shut it down cleanly. If this fails, perform a hard shutdown of the VM.
//
// Errors:
//  OPERATION_PARTIALLY_FAILED - Some VMs belonging to the appliance threw an exception while carrying out the specified operation
func (_class VMApplianceClass) Shutdown(sessionID SessionRef, self VMApplianceRef) (_err error) {
	if (IsMock) {
		return _class.ShutdownMock(sessionID, self)
	}	
	_method := "VM_appliance.shutdown"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMApplianceRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg)
	return
}


var VMApplianceClass_HardShutdownMockedCallback = func (sessionID SessionRef, self VMApplianceRef) (_err error) {
	log.Println("VMAppliance.HardShutdown not mocked")
	_err = errors.New("VMAppliance.HardShutdown not mocked")
	return
}

func (_class VMApplianceClass) HardShutdownMock(sessionID SessionRef, self VMApplianceRef) (_err error) {
	return VMApplianceClass_HardShutdownMockedCallback(sessionID, self)
}
// Perform a hard shutdown of all the VMs in the appliance
//
// Errors:
//  OPERATION_PARTIALLY_FAILED - Some VMs belonging to the appliance threw an exception while carrying out the specified operation
func (_class VMApplianceClass) HardShutdown(sessionID SessionRef, self VMApplianceRef) (_err error) {
	if (IsMock) {
		return _class.HardShutdownMock(sessionID, self)
	}	
	_method := "VM_appliance.hard_shutdown"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMApplianceRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg)
	return
}


var VMApplianceClass_CleanShutdownMockedCallback = func (sessionID SessionRef, self VMApplianceRef) (_err error) {
	log.Println("VMAppliance.CleanShutdown not mocked")
	_err = errors.New("VMAppliance.CleanShutdown not mocked")
	return
}

func (_class VMApplianceClass) CleanShutdownMock(sessionID SessionRef, self VMApplianceRef) (_err error) {
	return VMApplianceClass_CleanShutdownMockedCallback(sessionID, self)
}
// Perform a clean shutdown of all the VMs in the appliance
//
// Errors:
//  OPERATION_PARTIALLY_FAILED - Some VMs belonging to the appliance threw an exception while carrying out the specified operation
func (_class VMApplianceClass) CleanShutdown(sessionID SessionRef, self VMApplianceRef) (_err error) {
	if (IsMock) {
		return _class.CleanShutdownMock(sessionID, self)
	}	
	_method := "VM_appliance.clean_shutdown"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMApplianceRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg)
	return
}


var VMApplianceClass_StartMockedCallback = func (sessionID SessionRef, self VMApplianceRef, paused bool) (_err error) {
	log.Println("VMAppliance.Start not mocked")
	_err = errors.New("VMAppliance.Start not mocked")
	return
}

func (_class VMApplianceClass) StartMock(sessionID SessionRef, self VMApplianceRef, paused bool) (_err error) {
	return VMApplianceClass_StartMockedCallback(sessionID, self, paused)
}
// Start all VMs in the appliance
//
// Errors:
//  OPERATION_PARTIALLY_FAILED - Some VMs belonging to the appliance threw an exception while carrying out the specified operation
func (_class VMApplianceClass) Start(sessionID SessionRef, self VMApplianceRef, paused bool) (_err error) {
	if (IsMock) {
		return _class.StartMock(sessionID, self, paused)
	}	
	_method := "VM_appliance.start"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMApplianceRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_pausedArg, _err := convertBoolToXen(fmt.Sprintf("%s(%s)", _method, "paused"), paused)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _pausedArg)
	return
}


var VMApplianceClass_SetNameDescriptionMockedCallback = func (sessionID SessionRef, self VMApplianceRef, value string) (_err error) {
	log.Println("VMAppliance.SetNameDescription not mocked")
	_err = errors.New("VMAppliance.SetNameDescription not mocked")
	return
}

func (_class VMApplianceClass) SetNameDescriptionMock(sessionID SessionRef, self VMApplianceRef, value string) (_err error) {
	return VMApplianceClass_SetNameDescriptionMockedCallback(sessionID, self, value)
}
// Set the name/description field of the given VM_appliance.
func (_class VMApplianceClass) SetNameDescription(sessionID SessionRef, self VMApplianceRef, value string) (_err error) {
	if (IsMock) {
		return _class.SetNameDescriptionMock(sessionID, self, value)
	}	
	_method := "VM_appliance.set_name_description"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMApplianceRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var VMApplianceClass_SetNameLabelMockedCallback = func (sessionID SessionRef, self VMApplianceRef, value string) (_err error) {
	log.Println("VMAppliance.SetNameLabel not mocked")
	_err = errors.New("VMAppliance.SetNameLabel not mocked")
	return
}

func (_class VMApplianceClass) SetNameLabelMock(sessionID SessionRef, self VMApplianceRef, value string) (_err error) {
	return VMApplianceClass_SetNameLabelMockedCallback(sessionID, self, value)
}
// Set the name/label field of the given VM_appliance.
func (_class VMApplianceClass) SetNameLabel(sessionID SessionRef, self VMApplianceRef, value string) (_err error) {
	if (IsMock) {
		return _class.SetNameLabelMock(sessionID, self, value)
	}	
	_method := "VM_appliance.set_name_label"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMApplianceRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var VMApplianceClass_GetVMsMockedCallback = func (sessionID SessionRef, self VMApplianceRef) (_retval []VMRef, _err error) {
	log.Println("VMAppliance.GetVMs not mocked")
	_err = errors.New("VMAppliance.GetVMs not mocked")
	return
}

func (_class VMApplianceClass) GetVMsMock(sessionID SessionRef, self VMApplianceRef) (_retval []VMRef, _err error) {
	return VMApplianceClass_GetVMsMockedCallback(sessionID, self)
}
// Get the VMs field of the given VM_appliance.
func (_class VMApplianceClass) GetVMs(sessionID SessionRef, self VMApplianceRef) (_retval []VMRef, _err error) {
	if (IsMock) {
		return _class.GetVMsMock(sessionID, self)
	}	
	_method := "VM_appliance.get_VMs"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMApplianceRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVMRefSetToGo(_method + " -> ", _result.Value)
	return
}


var VMApplianceClass_GetCurrentOperationsMockedCallback = func (sessionID SessionRef, self VMApplianceRef) (_retval map[string]VMApplianceOperation, _err error) {
	log.Println("VMAppliance.GetCurrentOperations not mocked")
	_err = errors.New("VMAppliance.GetCurrentOperations not mocked")
	return
}

func (_class VMApplianceClass) GetCurrentOperationsMock(sessionID SessionRef, self VMApplianceRef) (_retval map[string]VMApplianceOperation, _err error) {
	return VMApplianceClass_GetCurrentOperationsMockedCallback(sessionID, self)
}
// Get the current_operations field of the given VM_appliance.
func (_class VMApplianceClass) GetCurrentOperations(sessionID SessionRef, self VMApplianceRef) (_retval map[string]VMApplianceOperation, _err error) {
	if (IsMock) {
		return _class.GetCurrentOperationsMock(sessionID, self)
	}	
	_method := "VM_appliance.get_current_operations"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMApplianceRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringToEnumVMApplianceOperationMapToGo(_method + " -> ", _result.Value)
	return
}


var VMApplianceClass_GetAllowedOperationsMockedCallback = func (sessionID SessionRef, self VMApplianceRef) (_retval []VMApplianceOperation, _err error) {
	log.Println("VMAppliance.GetAllowedOperations not mocked")
	_err = errors.New("VMAppliance.GetAllowedOperations not mocked")
	return
}

func (_class VMApplianceClass) GetAllowedOperationsMock(sessionID SessionRef, self VMApplianceRef) (_retval []VMApplianceOperation, _err error) {
	return VMApplianceClass_GetAllowedOperationsMockedCallback(sessionID, self)
}
// Get the allowed_operations field of the given VM_appliance.
func (_class VMApplianceClass) GetAllowedOperations(sessionID SessionRef, self VMApplianceRef) (_retval []VMApplianceOperation, _err error) {
	if (IsMock) {
		return _class.GetAllowedOperationsMock(sessionID, self)
	}	
	_method := "VM_appliance.get_allowed_operations"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMApplianceRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertEnumVMApplianceOperationSetToGo(_method + " -> ", _result.Value)
	return
}


var VMApplianceClass_GetNameDescriptionMockedCallback = func (sessionID SessionRef, self VMApplianceRef) (_retval string, _err error) {
	log.Println("VMAppliance.GetNameDescription not mocked")
	_err = errors.New("VMAppliance.GetNameDescription not mocked")
	return
}

func (_class VMApplianceClass) GetNameDescriptionMock(sessionID SessionRef, self VMApplianceRef) (_retval string, _err error) {
	return VMApplianceClass_GetNameDescriptionMockedCallback(sessionID, self)
}
// Get the name/description field of the given VM_appliance.
func (_class VMApplianceClass) GetNameDescription(sessionID SessionRef, self VMApplianceRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetNameDescriptionMock(sessionID, self)
	}	
	_method := "VM_appliance.get_name_description"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMApplianceRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var VMApplianceClass_GetNameLabelMockedCallback = func (sessionID SessionRef, self VMApplianceRef) (_retval string, _err error) {
	log.Println("VMAppliance.GetNameLabel not mocked")
	_err = errors.New("VMAppliance.GetNameLabel not mocked")
	return
}

func (_class VMApplianceClass) GetNameLabelMock(sessionID SessionRef, self VMApplianceRef) (_retval string, _err error) {
	return VMApplianceClass_GetNameLabelMockedCallback(sessionID, self)
}
// Get the name/label field of the given VM_appliance.
func (_class VMApplianceClass) GetNameLabel(sessionID SessionRef, self VMApplianceRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetNameLabelMock(sessionID, self)
	}	
	_method := "VM_appliance.get_name_label"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMApplianceRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var VMApplianceClass_GetUUIDMockedCallback = func (sessionID SessionRef, self VMApplianceRef) (_retval string, _err error) {
	log.Println("VMAppliance.GetUUID not mocked")
	_err = errors.New("VMAppliance.GetUUID not mocked")
	return
}

func (_class VMApplianceClass) GetUUIDMock(sessionID SessionRef, self VMApplianceRef) (_retval string, _err error) {
	return VMApplianceClass_GetUUIDMockedCallback(sessionID, self)
}
// Get the uuid field of the given VM_appliance.
func (_class VMApplianceClass) GetUUID(sessionID SessionRef, self VMApplianceRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetUUIDMock(sessionID, self)
	}	
	_method := "VM_appliance.get_uuid"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMApplianceRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var VMApplianceClass_GetByNameLabelMockedCallback = func (sessionID SessionRef, label string) (_retval []VMApplianceRef, _err error) {
	log.Println("VMAppliance.GetByNameLabel not mocked")
	_err = errors.New("VMAppliance.GetByNameLabel not mocked")
	return
}

func (_class VMApplianceClass) GetByNameLabelMock(sessionID SessionRef, label string) (_retval []VMApplianceRef, _err error) {
	return VMApplianceClass_GetByNameLabelMockedCallback(sessionID, label)
}
// Get all the VM_appliance instances with the given label.
func (_class VMApplianceClass) GetByNameLabel(sessionID SessionRef, label string) (_retval []VMApplianceRef, _err error) {
	if (IsMock) {
		return _class.GetByNameLabelMock(sessionID, label)
	}	
	_method := "VM_appliance.get_by_name_label"
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
	_retval, _err = convertVMApplianceRefSetToGo(_method + " -> ", _result.Value)
	return
}


var VMApplianceClass_DestroyMockedCallback = func (sessionID SessionRef, self VMApplianceRef) (_err error) {
	log.Println("VMAppliance.Destroy not mocked")
	_err = errors.New("VMAppliance.Destroy not mocked")
	return
}

func (_class VMApplianceClass) DestroyMock(sessionID SessionRef, self VMApplianceRef) (_err error) {
	return VMApplianceClass_DestroyMockedCallback(sessionID, self)
}
// Destroy the specified VM_appliance instance.
func (_class VMApplianceClass) Destroy(sessionID SessionRef, self VMApplianceRef) (_err error) {
	if (IsMock) {
		return _class.DestroyMock(sessionID, self)
	}	
	_method := "VM_appliance.destroy"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMApplianceRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg)
	return
}


var VMApplianceClass_CreateMockedCallback = func (sessionID SessionRef, args VMApplianceRecord) (_retval VMApplianceRef, _err error) {
	log.Println("VMAppliance.Create not mocked")
	_err = errors.New("VMAppliance.Create not mocked")
	return
}

func (_class VMApplianceClass) CreateMock(sessionID SessionRef, args VMApplianceRecord) (_retval VMApplianceRef, _err error) {
	return VMApplianceClass_CreateMockedCallback(sessionID, args)
}
// Create a new VM_appliance instance, and return its handle.
// The constructor args are: name_label, name_description (* = non-optional).
func (_class VMApplianceClass) Create(sessionID SessionRef, args VMApplianceRecord) (_retval VMApplianceRef, _err error) {
	if (IsMock) {
		return _class.CreateMock(sessionID, args)
	}	
	_method := "VM_appliance.create"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_argsArg, _err := convertVMApplianceRecordToXen(fmt.Sprintf("%s(%s)", _method, "args"), args)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _argsArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVMApplianceRefToGo(_method + " -> ", _result.Value)
	return
}


var VMApplianceClass_GetByUUIDMockedCallback = func (sessionID SessionRef, uuid string) (_retval VMApplianceRef, _err error) {
	log.Println("VMAppliance.GetByUUID not mocked")
	_err = errors.New("VMAppliance.GetByUUID not mocked")
	return
}

func (_class VMApplianceClass) GetByUUIDMock(sessionID SessionRef, uuid string) (_retval VMApplianceRef, _err error) {
	return VMApplianceClass_GetByUUIDMockedCallback(sessionID, uuid)
}
// Get a reference to the VM_appliance instance with the specified UUID.
func (_class VMApplianceClass) GetByUUID(sessionID SessionRef, uuid string) (_retval VMApplianceRef, _err error) {
	if (IsMock) {
		return _class.GetByUUIDMock(sessionID, uuid)
	}	
	_method := "VM_appliance.get_by_uuid"
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
	_retval, _err = convertVMApplianceRefToGo(_method + " -> ", _result.Value)
	return
}


var VMApplianceClass_GetRecordMockedCallback = func (sessionID SessionRef, self VMApplianceRef) (_retval VMApplianceRecord, _err error) {
	log.Println("VMAppliance.GetRecord not mocked")
	_err = errors.New("VMAppliance.GetRecord not mocked")
	return
}

func (_class VMApplianceClass) GetRecordMock(sessionID SessionRef, self VMApplianceRef) (_retval VMApplianceRecord, _err error) {
	return VMApplianceClass_GetRecordMockedCallback(sessionID, self)
}
// Get a record containing the current state of the given VM_appliance.
func (_class VMApplianceClass) GetRecord(sessionID SessionRef, self VMApplianceRef) (_retval VMApplianceRecord, _err error) {
	if (IsMock) {
		return _class.GetRecordMock(sessionID, self)
	}	
	_method := "VM_appliance.get_record"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMApplianceRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVMApplianceRecordToGo(_method + " -> ", _result.Value)
	return
}
