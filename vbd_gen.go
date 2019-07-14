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

type VbdOperations string

const (
  // Attempting to attach this VBD to a VM
	VbdOperationsAttach VbdOperations = "attach"
  // Attempting to eject the media from this VBD
	VbdOperationsEject VbdOperations = "eject"
  // Attempting to insert new media into this VBD
	VbdOperationsInsert VbdOperations = "insert"
  // Attempting to hotplug this VBD
	VbdOperationsPlug VbdOperations = "plug"
  // Attempting to hot unplug this VBD
	VbdOperationsUnplug VbdOperations = "unplug"
  // Attempting to forcibly unplug this VBD
	VbdOperationsUnplugForce VbdOperations = "unplug_force"
  // Attempting to pause a block device backend
	VbdOperationsPause VbdOperations = "pause"
  // Attempting to unpause a block device backend
	VbdOperationsUnpause VbdOperations = "unpause"
)

type VbdType string

const (
  // VBD will appear to guest as CD
	VbdTypeCD VbdType = "CD"
  // VBD will appear to guest as disk
	VbdTypeDisk VbdType = "Disk"
  // VBD will appear as a floppy
	VbdTypeFloppy VbdType = "Floppy"
)

type VbdMode string

const (
  // only read-only access will be allowed
	VbdModeRO VbdMode = "RO"
  // read-write access will be allowed
	VbdModeRW VbdMode = "RW"
)

type VBDRecord struct {
  // Unique identifier/object reference
	UUID string
  // list of the operations allowed in this state. This list is advisory only and the server state may have changed by the time this field is read by a client.
	AllowedOperations []VbdOperations
  // links each of the running tasks using this object (by reference) to a current_operation enum which describes the nature of the task.
	CurrentOperations map[string]VbdOperations
  // the virtual machine
	VM VMRef
  // the virtual disk
	VDI VDIRef
  // device seen by the guest e.g. hda1
	Device string
  // user-friendly device name e.g. 0,1,2,etc.
	Userdevice string
  // true if this VBD is bootable
	Bootable bool
  // the mode the VBD should be mounted with
	Mode VbdMode
  // how the VBD will appear to the guest (e.g. disk or CD)
	Type VbdType
  // true if this VBD will support hot-unplug
	Unpluggable bool
  // true if a storage level lock was acquired
	StorageLock bool
  // if true this represents an empty drive
	Empty bool
  // additional configuration
	OtherConfig map[string]string
  // is the device currently attached (erased on reboot)
	CurrentlyAttached bool
  // error/success code associated with last attach-operation (erased on reboot)
	StatusCode int
  // error/success information associated with last attach-operation status (erased on reboot)
	StatusDetail string
  // Device runtime properties
	RuntimeProperties map[string]string
  // QoS algorithm to use
	QosAlgorithmType string
  // parameters for chosen QoS algorithm
	QosAlgorithmParams map[string]string
  // supported QoS algorithms for this VBD
	QosSupportedAlgorithms []string
  // metrics associated with this VBD
	Metrics VBDMetricsRef
}

type VBDRef string

// A virtual block device
type VBDClass struct {
	client *Client
}


func VBDClassGetAllRecordsMockDefault(sessionID SessionRef) (_retval map[VBDRef]VBDRecord, _err error) {
	log.Println("VBD.GetAllRecords not mocked")
	_err = errors.New("VBD.GetAllRecords not mocked")
	return
}

var VBDClassGetAllRecordsMockedCallback = VBDClassGetAllRecordsMockDefault

func (_class VBDClass) GetAllRecordsMock(sessionID SessionRef) (_retval map[VBDRef]VBDRecord, _err error) {
	return VBDClassGetAllRecordsMockedCallback(sessionID)
}
// Return a map of VBD references to VBD records for all VBDs known to the system.
func (_class VBDClass) GetAllRecords(sessionID SessionRef) (_retval map[VBDRef]VBDRecord, _err error) {
	if IsMock {
		return _class.GetAllRecordsMock(sessionID)
	}	
	_method := "VBD.get_all_records"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVBDRefToVBDRecordMapToGo(_method + " -> ", _result.Value)
	return
}


func VBDClassGetAllMockDefault(sessionID SessionRef) (_retval []VBDRef, _err error) {
	log.Println("VBD.GetAll not mocked")
	_err = errors.New("VBD.GetAll not mocked")
	return
}

var VBDClassGetAllMockedCallback = VBDClassGetAllMockDefault

func (_class VBDClass) GetAllMock(sessionID SessionRef) (_retval []VBDRef, _err error) {
	return VBDClassGetAllMockedCallback(sessionID)
}
// Return a list of all the VBDs known to the system.
func (_class VBDClass) GetAll(sessionID SessionRef) (_retval []VBDRef, _err error) {
	if IsMock {
		return _class.GetAllMock(sessionID)
	}	
	_method := "VBD.get_all"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVBDRefSetToGo(_method + " -> ", _result.Value)
	return
}


func VBDClassSetModeMockDefault(sessionID SessionRef, self VBDRef, value VbdMode) (_err error) {
	log.Println("VBD.SetMode not mocked")
	_err = errors.New("VBD.SetMode not mocked")
	return
}

var VBDClassSetModeMockedCallback = VBDClassSetModeMockDefault

func (_class VBDClass) SetModeMock(sessionID SessionRef, self VBDRef, value VbdMode) (_err error) {
	return VBDClassSetModeMockedCallback(sessionID, self, value)
}
// Sets the mode of the VBD. The power_state of the VM must be halted.
func (_class VBDClass) SetMode(sessionID SessionRef, self VBDRef, value VbdMode) (_err error) {
	if IsMock {
		return _class.SetModeMock(sessionID, self, value)
	}	
	_method := "VBD.set_mode"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVBDRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertEnumVbdModeToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}


func VBDClassAssertAttachableMockDefault(sessionID SessionRef, self VBDRef) (_err error) {
	log.Println("VBD.AssertAttachable not mocked")
	_err = errors.New("VBD.AssertAttachable not mocked")
	return
}

var VBDClassAssertAttachableMockedCallback = VBDClassAssertAttachableMockDefault

func (_class VBDClass) AssertAttachableMock(sessionID SessionRef, self VBDRef) (_err error) {
	return VBDClassAssertAttachableMockedCallback(sessionID, self)
}
// Throws an error if this VBD could not be attached to this VM if the VM were running. Intended for debugging.
func (_class VBDClass) AssertAttachable(sessionID SessionRef, self VBDRef) (_err error) {
	if IsMock {
		return _class.AssertAttachableMock(sessionID, self)
	}	
	_method := "VBD.assert_attachable"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVBDRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg)
	return
}


func VBDClassUnplugForceMockDefault(sessionID SessionRef, self VBDRef) (_err error) {
	log.Println("VBD.UnplugForce not mocked")
	_err = errors.New("VBD.UnplugForce not mocked")
	return
}

var VBDClassUnplugForceMockedCallback = VBDClassUnplugForceMockDefault

func (_class VBDClass) UnplugForceMock(sessionID SessionRef, self VBDRef) (_err error) {
	return VBDClassUnplugForceMockedCallback(sessionID, self)
}
// Forcibly unplug the specified VBD
func (_class VBDClass) UnplugForce(sessionID SessionRef, self VBDRef) (_err error) {
	if IsMock {
		return _class.UnplugForceMock(sessionID, self)
	}	
	_method := "VBD.unplug_force"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVBDRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg)
	return
}


func VBDClassUnplugMockDefault(sessionID SessionRef, self VBDRef) (_err error) {
	log.Println("VBD.Unplug not mocked")
	_err = errors.New("VBD.Unplug not mocked")
	return
}

var VBDClassUnplugMockedCallback = VBDClassUnplugMockDefault

func (_class VBDClass) UnplugMock(sessionID SessionRef, self VBDRef) (_err error) {
	return VBDClassUnplugMockedCallback(sessionID, self)
}
// Hot-unplug the specified VBD, dynamically unattaching it from the running VM
//
// Errors:
//  DEVICE_DETACH_REJECTED - The VM rejected the attempt to detach the device.
//  DEVICE_ALREADY_DETACHED - The device is not currently attached
func (_class VBDClass) Unplug(sessionID SessionRef, self VBDRef) (_err error) {
	if IsMock {
		return _class.UnplugMock(sessionID, self)
	}	
	_method := "VBD.unplug"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVBDRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg)
	return
}


func VBDClassPlugMockDefault(sessionID SessionRef, self VBDRef) (_err error) {
	log.Println("VBD.Plug not mocked")
	_err = errors.New("VBD.Plug not mocked")
	return
}

var VBDClassPlugMockedCallback = VBDClassPlugMockDefault

func (_class VBDClass) PlugMock(sessionID SessionRef, self VBDRef) (_err error) {
	return VBDClassPlugMockedCallback(sessionID, self)
}
// Hotplug the specified VBD, dynamically attaching it to the running VM
func (_class VBDClass) Plug(sessionID SessionRef, self VBDRef) (_err error) {
	if IsMock {
		return _class.PlugMock(sessionID, self)
	}	
	_method := "VBD.plug"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVBDRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg)
	return
}


func VBDClassInsertMockDefault(sessionID SessionRef, vbd VBDRef, vdi VDIRef) (_err error) {
	log.Println("VBD.Insert not mocked")
	_err = errors.New("VBD.Insert not mocked")
	return
}

var VBDClassInsertMockedCallback = VBDClassInsertMockDefault

func (_class VBDClass) InsertMock(sessionID SessionRef, vbd VBDRef, vdi VDIRef) (_err error) {
	return VBDClassInsertMockedCallback(sessionID, vbd, vdi)
}
// Insert new media into the device
//
// Errors:
//  VBD_NOT_REMOVABLE_MEDIA - Media could not be ejected because it is not removable
//  VBD_NOT_EMPTY - Operation could not be performed because the drive is not empty
func (_class VBDClass) Insert(sessionID SessionRef, vbd VBDRef, vdi VDIRef) (_err error) {
	if IsMock {
		return _class.InsertMock(sessionID, vbd, vdi)
	}	
	_method := "VBD.insert"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_vbdArg, _err := convertVBDRefToXen(fmt.Sprintf("%s(%s)", _method, "vbd"), vbd)
	if _err != nil {
		return
	}
	_vdiArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "vdi"), vdi)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _vbdArg, _vdiArg)
	return
}


func VBDClassEjectMockDefault(sessionID SessionRef, vbd VBDRef) (_err error) {
	log.Println("VBD.Eject not mocked")
	_err = errors.New("VBD.Eject not mocked")
	return
}

var VBDClassEjectMockedCallback = VBDClassEjectMockDefault

func (_class VBDClass) EjectMock(sessionID SessionRef, vbd VBDRef) (_err error) {
	return VBDClassEjectMockedCallback(sessionID, vbd)
}
// Remove the media from the device and leave it empty
//
// Errors:
//  VBD_NOT_REMOVABLE_MEDIA - Media could not be ejected because it is not removable
//  VBD_IS_EMPTY - Operation could not be performed because the drive is empty
func (_class VBDClass) Eject(sessionID SessionRef, vbd VBDRef) (_err error) {
	if IsMock {
		return _class.EjectMock(sessionID, vbd)
	}	
	_method := "VBD.eject"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_vbdArg, _err := convertVBDRefToXen(fmt.Sprintf("%s(%s)", _method, "vbd"), vbd)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _vbdArg)
	return
}


func VBDClassRemoveFromQosAlgorithmParamsMockDefault(sessionID SessionRef, self VBDRef, key string) (_err error) {
	log.Println("VBD.RemoveFromQosAlgorithmParams not mocked")
	_err = errors.New("VBD.RemoveFromQosAlgorithmParams not mocked")
	return
}

var VBDClassRemoveFromQosAlgorithmParamsMockedCallback = VBDClassRemoveFromQosAlgorithmParamsMockDefault

func (_class VBDClass) RemoveFromQosAlgorithmParamsMock(sessionID SessionRef, self VBDRef, key string) (_err error) {
	return VBDClassRemoveFromQosAlgorithmParamsMockedCallback(sessionID, self, key)
}
// Remove the given key and its corresponding value from the qos/algorithm_params field of the given VBD.  If the key is not in that Map, then do nothing.
func (_class VBDClass) RemoveFromQosAlgorithmParams(sessionID SessionRef, self VBDRef, key string) (_err error) {
	if IsMock {
		return _class.RemoveFromQosAlgorithmParamsMock(sessionID, self, key)
	}	
	_method := "VBD.remove_from_qos_algorithm_params"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVBDRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func VBDClassAddToQosAlgorithmParamsMockDefault(sessionID SessionRef, self VBDRef, key string, value string) (_err error) {
	log.Println("VBD.AddToQosAlgorithmParams not mocked")
	_err = errors.New("VBD.AddToQosAlgorithmParams not mocked")
	return
}

var VBDClassAddToQosAlgorithmParamsMockedCallback = VBDClassAddToQosAlgorithmParamsMockDefault

func (_class VBDClass) AddToQosAlgorithmParamsMock(sessionID SessionRef, self VBDRef, key string, value string) (_err error) {
	return VBDClassAddToQosAlgorithmParamsMockedCallback(sessionID, self, key, value)
}
// Add the given key-value pair to the qos/algorithm_params field of the given VBD.
func (_class VBDClass) AddToQosAlgorithmParams(sessionID SessionRef, self VBDRef, key string, value string) (_err error) {
	if IsMock {
		return _class.AddToQosAlgorithmParamsMock(sessionID, self, key, value)
	}	
	_method := "VBD.add_to_qos_algorithm_params"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVBDRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func VBDClassSetQosAlgorithmParamsMockDefault(sessionID SessionRef, self VBDRef, value map[string]string) (_err error) {
	log.Println("VBD.SetQosAlgorithmParams not mocked")
	_err = errors.New("VBD.SetQosAlgorithmParams not mocked")
	return
}

var VBDClassSetQosAlgorithmParamsMockedCallback = VBDClassSetQosAlgorithmParamsMockDefault

func (_class VBDClass) SetQosAlgorithmParamsMock(sessionID SessionRef, self VBDRef, value map[string]string) (_err error) {
	return VBDClassSetQosAlgorithmParamsMockedCallback(sessionID, self, value)
}
// Set the qos/algorithm_params field of the given VBD.
func (_class VBDClass) SetQosAlgorithmParams(sessionID SessionRef, self VBDRef, value map[string]string) (_err error) {
	if IsMock {
		return _class.SetQosAlgorithmParamsMock(sessionID, self, value)
	}	
	_method := "VBD.set_qos_algorithm_params"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVBDRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func VBDClassSetQosAlgorithmTypeMockDefault(sessionID SessionRef, self VBDRef, value string) (_err error) {
	log.Println("VBD.SetQosAlgorithmType not mocked")
	_err = errors.New("VBD.SetQosAlgorithmType not mocked")
	return
}

var VBDClassSetQosAlgorithmTypeMockedCallback = VBDClassSetQosAlgorithmTypeMockDefault

func (_class VBDClass) SetQosAlgorithmTypeMock(sessionID SessionRef, self VBDRef, value string) (_err error) {
	return VBDClassSetQosAlgorithmTypeMockedCallback(sessionID, self, value)
}
// Set the qos/algorithm_type field of the given VBD.
func (_class VBDClass) SetQosAlgorithmType(sessionID SessionRef, self VBDRef, value string) (_err error) {
	if IsMock {
		return _class.SetQosAlgorithmTypeMock(sessionID, self, value)
	}	
	_method := "VBD.set_qos_algorithm_type"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVBDRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func VBDClassRemoveFromOtherConfigMockDefault(sessionID SessionRef, self VBDRef, key string) (_err error) {
	log.Println("VBD.RemoveFromOtherConfig not mocked")
	_err = errors.New("VBD.RemoveFromOtherConfig not mocked")
	return
}

var VBDClassRemoveFromOtherConfigMockedCallback = VBDClassRemoveFromOtherConfigMockDefault

func (_class VBDClass) RemoveFromOtherConfigMock(sessionID SessionRef, self VBDRef, key string) (_err error) {
	return VBDClassRemoveFromOtherConfigMockedCallback(sessionID, self, key)
}
// Remove the given key and its corresponding value from the other_config field of the given VBD.  If the key is not in that Map, then do nothing.
func (_class VBDClass) RemoveFromOtherConfig(sessionID SessionRef, self VBDRef, key string) (_err error) {
	if IsMock {
		return _class.RemoveFromOtherConfigMock(sessionID, self, key)
	}	
	_method := "VBD.remove_from_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVBDRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func VBDClassAddToOtherConfigMockDefault(sessionID SessionRef, self VBDRef, key string, value string) (_err error) {
	log.Println("VBD.AddToOtherConfig not mocked")
	_err = errors.New("VBD.AddToOtherConfig not mocked")
	return
}

var VBDClassAddToOtherConfigMockedCallback = VBDClassAddToOtherConfigMockDefault

func (_class VBDClass) AddToOtherConfigMock(sessionID SessionRef, self VBDRef, key string, value string) (_err error) {
	return VBDClassAddToOtherConfigMockedCallback(sessionID, self, key, value)
}
// Add the given key-value pair to the other_config field of the given VBD.
func (_class VBDClass) AddToOtherConfig(sessionID SessionRef, self VBDRef, key string, value string) (_err error) {
	if IsMock {
		return _class.AddToOtherConfigMock(sessionID, self, key, value)
	}	
	_method := "VBD.add_to_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVBDRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func VBDClassSetOtherConfigMockDefault(sessionID SessionRef, self VBDRef, value map[string]string) (_err error) {
	log.Println("VBD.SetOtherConfig not mocked")
	_err = errors.New("VBD.SetOtherConfig not mocked")
	return
}

var VBDClassSetOtherConfigMockedCallback = VBDClassSetOtherConfigMockDefault

func (_class VBDClass) SetOtherConfigMock(sessionID SessionRef, self VBDRef, value map[string]string) (_err error) {
	return VBDClassSetOtherConfigMockedCallback(sessionID, self, value)
}
// Set the other_config field of the given VBD.
func (_class VBDClass) SetOtherConfig(sessionID SessionRef, self VBDRef, value map[string]string) (_err error) {
	if IsMock {
		return _class.SetOtherConfigMock(sessionID, self, value)
	}	
	_method := "VBD.set_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVBDRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func VBDClassSetUnpluggableMockDefault(sessionID SessionRef, self VBDRef, value bool) (_err error) {
	log.Println("VBD.SetUnpluggable not mocked")
	_err = errors.New("VBD.SetUnpluggable not mocked")
	return
}

var VBDClassSetUnpluggableMockedCallback = VBDClassSetUnpluggableMockDefault

func (_class VBDClass) SetUnpluggableMock(sessionID SessionRef, self VBDRef, value bool) (_err error) {
	return VBDClassSetUnpluggableMockedCallback(sessionID, self, value)
}
// Set the unpluggable field of the given VBD.
func (_class VBDClass) SetUnpluggable(sessionID SessionRef, self VBDRef, value bool) (_err error) {
	if IsMock {
		return _class.SetUnpluggableMock(sessionID, self, value)
	}	
	_method := "VBD.set_unpluggable"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVBDRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertBoolToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}


func VBDClassSetTypeMockDefault(sessionID SessionRef, self VBDRef, value VbdType) (_err error) {
	log.Println("VBD.SetType not mocked")
	_err = errors.New("VBD.SetType not mocked")
	return
}

var VBDClassSetTypeMockedCallback = VBDClassSetTypeMockDefault

func (_class VBDClass) SetTypeMock(sessionID SessionRef, self VBDRef, value VbdType) (_err error) {
	return VBDClassSetTypeMockedCallback(sessionID, self, value)
}
// Set the type field of the given VBD.
func (_class VBDClass) SetType(sessionID SessionRef, self VBDRef, value VbdType) (_err error) {
	if IsMock {
		return _class.SetTypeMock(sessionID, self, value)
	}	
	_method := "VBD.set_type"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVBDRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertEnumVbdTypeToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}


func VBDClassSetBootableMockDefault(sessionID SessionRef, self VBDRef, value bool) (_err error) {
	log.Println("VBD.SetBootable not mocked")
	_err = errors.New("VBD.SetBootable not mocked")
	return
}

var VBDClassSetBootableMockedCallback = VBDClassSetBootableMockDefault

func (_class VBDClass) SetBootableMock(sessionID SessionRef, self VBDRef, value bool) (_err error) {
	return VBDClassSetBootableMockedCallback(sessionID, self, value)
}
// Set the bootable field of the given VBD.
func (_class VBDClass) SetBootable(sessionID SessionRef, self VBDRef, value bool) (_err error) {
	if IsMock {
		return _class.SetBootableMock(sessionID, self, value)
	}	
	_method := "VBD.set_bootable"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVBDRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertBoolToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}


func VBDClassSetUserdeviceMockDefault(sessionID SessionRef, self VBDRef, value string) (_err error) {
	log.Println("VBD.SetUserdevice not mocked")
	_err = errors.New("VBD.SetUserdevice not mocked")
	return
}

var VBDClassSetUserdeviceMockedCallback = VBDClassSetUserdeviceMockDefault

func (_class VBDClass) SetUserdeviceMock(sessionID SessionRef, self VBDRef, value string) (_err error) {
	return VBDClassSetUserdeviceMockedCallback(sessionID, self, value)
}
// Set the userdevice field of the given VBD.
func (_class VBDClass) SetUserdevice(sessionID SessionRef, self VBDRef, value string) (_err error) {
	if IsMock {
		return _class.SetUserdeviceMock(sessionID, self, value)
	}	
	_method := "VBD.set_userdevice"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVBDRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func VBDClassGetMetricsMockDefault(sessionID SessionRef, self VBDRef) (_retval VBDMetricsRef, _err error) {
	log.Println("VBD.GetMetrics not mocked")
	_err = errors.New("VBD.GetMetrics not mocked")
	return
}

var VBDClassGetMetricsMockedCallback = VBDClassGetMetricsMockDefault

func (_class VBDClass) GetMetricsMock(sessionID SessionRef, self VBDRef) (_retval VBDMetricsRef, _err error) {
	return VBDClassGetMetricsMockedCallback(sessionID, self)
}
// Get the metrics field of the given VBD.
func (_class VBDClass) GetMetrics(sessionID SessionRef, self VBDRef) (_retval VBDMetricsRef, _err error) {
	if IsMock {
		return _class.GetMetricsMock(sessionID, self)
	}	
	_method := "VBD.get_metrics"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVBDRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVBDMetricsRefToGo(_method + " -> ", _result.Value)
	return
}


func VBDClassGetQosSupportedAlgorithmsMockDefault(sessionID SessionRef, self VBDRef) (_retval []string, _err error) {
	log.Println("VBD.GetQosSupportedAlgorithms not mocked")
	_err = errors.New("VBD.GetQosSupportedAlgorithms not mocked")
	return
}

var VBDClassGetQosSupportedAlgorithmsMockedCallback = VBDClassGetQosSupportedAlgorithmsMockDefault

func (_class VBDClass) GetQosSupportedAlgorithmsMock(sessionID SessionRef, self VBDRef) (_retval []string, _err error) {
	return VBDClassGetQosSupportedAlgorithmsMockedCallback(sessionID, self)
}
// Get the qos/supported_algorithms field of the given VBD.
func (_class VBDClass) GetQosSupportedAlgorithms(sessionID SessionRef, self VBDRef) (_retval []string, _err error) {
	if IsMock {
		return _class.GetQosSupportedAlgorithmsMock(sessionID, self)
	}	
	_method := "VBD.get_qos_supported_algorithms"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVBDRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func VBDClassGetQosAlgorithmParamsMockDefault(sessionID SessionRef, self VBDRef) (_retval map[string]string, _err error) {
	log.Println("VBD.GetQosAlgorithmParams not mocked")
	_err = errors.New("VBD.GetQosAlgorithmParams not mocked")
	return
}

var VBDClassGetQosAlgorithmParamsMockedCallback = VBDClassGetQosAlgorithmParamsMockDefault

func (_class VBDClass) GetQosAlgorithmParamsMock(sessionID SessionRef, self VBDRef) (_retval map[string]string, _err error) {
	return VBDClassGetQosAlgorithmParamsMockedCallback(sessionID, self)
}
// Get the qos/algorithm_params field of the given VBD.
func (_class VBDClass) GetQosAlgorithmParams(sessionID SessionRef, self VBDRef) (_retval map[string]string, _err error) {
	if IsMock {
		return _class.GetQosAlgorithmParamsMock(sessionID, self)
	}	
	_method := "VBD.get_qos_algorithm_params"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVBDRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func VBDClassGetQosAlgorithmTypeMockDefault(sessionID SessionRef, self VBDRef) (_retval string, _err error) {
	log.Println("VBD.GetQosAlgorithmType not mocked")
	_err = errors.New("VBD.GetQosAlgorithmType not mocked")
	return
}

var VBDClassGetQosAlgorithmTypeMockedCallback = VBDClassGetQosAlgorithmTypeMockDefault

func (_class VBDClass) GetQosAlgorithmTypeMock(sessionID SessionRef, self VBDRef) (_retval string, _err error) {
	return VBDClassGetQosAlgorithmTypeMockedCallback(sessionID, self)
}
// Get the qos/algorithm_type field of the given VBD.
func (_class VBDClass) GetQosAlgorithmType(sessionID SessionRef, self VBDRef) (_retval string, _err error) {
	if IsMock {
		return _class.GetQosAlgorithmTypeMock(sessionID, self)
	}	
	_method := "VBD.get_qos_algorithm_type"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVBDRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func VBDClassGetRuntimePropertiesMockDefault(sessionID SessionRef, self VBDRef) (_retval map[string]string, _err error) {
	log.Println("VBD.GetRuntimeProperties not mocked")
	_err = errors.New("VBD.GetRuntimeProperties not mocked")
	return
}

var VBDClassGetRuntimePropertiesMockedCallback = VBDClassGetRuntimePropertiesMockDefault

func (_class VBDClass) GetRuntimePropertiesMock(sessionID SessionRef, self VBDRef) (_retval map[string]string, _err error) {
	return VBDClassGetRuntimePropertiesMockedCallback(sessionID, self)
}
// Get the runtime_properties field of the given VBD.
func (_class VBDClass) GetRuntimeProperties(sessionID SessionRef, self VBDRef) (_retval map[string]string, _err error) {
	if IsMock {
		return _class.GetRuntimePropertiesMock(sessionID, self)
	}	
	_method := "VBD.get_runtime_properties"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVBDRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func VBDClassGetStatusDetailMockDefault(sessionID SessionRef, self VBDRef) (_retval string, _err error) {
	log.Println("VBD.GetStatusDetail not mocked")
	_err = errors.New("VBD.GetStatusDetail not mocked")
	return
}

var VBDClassGetStatusDetailMockedCallback = VBDClassGetStatusDetailMockDefault

func (_class VBDClass) GetStatusDetailMock(sessionID SessionRef, self VBDRef) (_retval string, _err error) {
	return VBDClassGetStatusDetailMockedCallback(sessionID, self)
}
// Get the status_detail field of the given VBD.
func (_class VBDClass) GetStatusDetail(sessionID SessionRef, self VBDRef) (_retval string, _err error) {
	if IsMock {
		return _class.GetStatusDetailMock(sessionID, self)
	}	
	_method := "VBD.get_status_detail"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVBDRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func VBDClassGetStatusCodeMockDefault(sessionID SessionRef, self VBDRef) (_retval int, _err error) {
	log.Println("VBD.GetStatusCode not mocked")
	_err = errors.New("VBD.GetStatusCode not mocked")
	return
}

var VBDClassGetStatusCodeMockedCallback = VBDClassGetStatusCodeMockDefault

func (_class VBDClass) GetStatusCodeMock(sessionID SessionRef, self VBDRef) (_retval int, _err error) {
	return VBDClassGetStatusCodeMockedCallback(sessionID, self)
}
// Get the status_code field of the given VBD.
func (_class VBDClass) GetStatusCode(sessionID SessionRef, self VBDRef) (_retval int, _err error) {
	if IsMock {
		return _class.GetStatusCodeMock(sessionID, self)
	}	
	_method := "VBD.get_status_code"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVBDRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func VBDClassGetCurrentlyAttachedMockDefault(sessionID SessionRef, self VBDRef) (_retval bool, _err error) {
	log.Println("VBD.GetCurrentlyAttached not mocked")
	_err = errors.New("VBD.GetCurrentlyAttached not mocked")
	return
}

var VBDClassGetCurrentlyAttachedMockedCallback = VBDClassGetCurrentlyAttachedMockDefault

func (_class VBDClass) GetCurrentlyAttachedMock(sessionID SessionRef, self VBDRef) (_retval bool, _err error) {
	return VBDClassGetCurrentlyAttachedMockedCallback(sessionID, self)
}
// Get the currently_attached field of the given VBD.
func (_class VBDClass) GetCurrentlyAttached(sessionID SessionRef, self VBDRef) (_retval bool, _err error) {
	if IsMock {
		return _class.GetCurrentlyAttachedMock(sessionID, self)
	}	
	_method := "VBD.get_currently_attached"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVBDRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func VBDClassGetOtherConfigMockDefault(sessionID SessionRef, self VBDRef) (_retval map[string]string, _err error) {
	log.Println("VBD.GetOtherConfig not mocked")
	_err = errors.New("VBD.GetOtherConfig not mocked")
	return
}

var VBDClassGetOtherConfigMockedCallback = VBDClassGetOtherConfigMockDefault

func (_class VBDClass) GetOtherConfigMock(sessionID SessionRef, self VBDRef) (_retval map[string]string, _err error) {
	return VBDClassGetOtherConfigMockedCallback(sessionID, self)
}
// Get the other_config field of the given VBD.
func (_class VBDClass) GetOtherConfig(sessionID SessionRef, self VBDRef) (_retval map[string]string, _err error) {
	if IsMock {
		return _class.GetOtherConfigMock(sessionID, self)
	}	
	_method := "VBD.get_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVBDRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func VBDClassGetEmptyMockDefault(sessionID SessionRef, self VBDRef) (_retval bool, _err error) {
	log.Println("VBD.GetEmpty not mocked")
	_err = errors.New("VBD.GetEmpty not mocked")
	return
}

var VBDClassGetEmptyMockedCallback = VBDClassGetEmptyMockDefault

func (_class VBDClass) GetEmptyMock(sessionID SessionRef, self VBDRef) (_retval bool, _err error) {
	return VBDClassGetEmptyMockedCallback(sessionID, self)
}
// Get the empty field of the given VBD.
func (_class VBDClass) GetEmpty(sessionID SessionRef, self VBDRef) (_retval bool, _err error) {
	if IsMock {
		return _class.GetEmptyMock(sessionID, self)
	}	
	_method := "VBD.get_empty"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVBDRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func VBDClassGetStorageLockMockDefault(sessionID SessionRef, self VBDRef) (_retval bool, _err error) {
	log.Println("VBD.GetStorageLock not mocked")
	_err = errors.New("VBD.GetStorageLock not mocked")
	return
}

var VBDClassGetStorageLockMockedCallback = VBDClassGetStorageLockMockDefault

func (_class VBDClass) GetStorageLockMock(sessionID SessionRef, self VBDRef) (_retval bool, _err error) {
	return VBDClassGetStorageLockMockedCallback(sessionID, self)
}
// Get the storage_lock field of the given VBD.
func (_class VBDClass) GetStorageLock(sessionID SessionRef, self VBDRef) (_retval bool, _err error) {
	if IsMock {
		return _class.GetStorageLockMock(sessionID, self)
	}	
	_method := "VBD.get_storage_lock"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVBDRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func VBDClassGetUnpluggableMockDefault(sessionID SessionRef, self VBDRef) (_retval bool, _err error) {
	log.Println("VBD.GetUnpluggable not mocked")
	_err = errors.New("VBD.GetUnpluggable not mocked")
	return
}

var VBDClassGetUnpluggableMockedCallback = VBDClassGetUnpluggableMockDefault

func (_class VBDClass) GetUnpluggableMock(sessionID SessionRef, self VBDRef) (_retval bool, _err error) {
	return VBDClassGetUnpluggableMockedCallback(sessionID, self)
}
// Get the unpluggable field of the given VBD.
func (_class VBDClass) GetUnpluggable(sessionID SessionRef, self VBDRef) (_retval bool, _err error) {
	if IsMock {
		return _class.GetUnpluggableMock(sessionID, self)
	}	
	_method := "VBD.get_unpluggable"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVBDRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func VBDClassGetTypeMockDefault(sessionID SessionRef, self VBDRef) (_retval VbdType, _err error) {
	log.Println("VBD.GetType not mocked")
	_err = errors.New("VBD.GetType not mocked")
	return
}

var VBDClassGetTypeMockedCallback = VBDClassGetTypeMockDefault

func (_class VBDClass) GetTypeMock(sessionID SessionRef, self VBDRef) (_retval VbdType, _err error) {
	return VBDClassGetTypeMockedCallback(sessionID, self)
}
// Get the type field of the given VBD.
func (_class VBDClass) GetType(sessionID SessionRef, self VBDRef) (_retval VbdType, _err error) {
	if IsMock {
		return _class.GetTypeMock(sessionID, self)
	}	
	_method := "VBD.get_type"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVBDRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertEnumVbdTypeToGo(_method + " -> ", _result.Value)
	return
}


func VBDClassGetModeMockDefault(sessionID SessionRef, self VBDRef) (_retval VbdMode, _err error) {
	log.Println("VBD.GetMode not mocked")
	_err = errors.New("VBD.GetMode not mocked")
	return
}

var VBDClassGetModeMockedCallback = VBDClassGetModeMockDefault

func (_class VBDClass) GetModeMock(sessionID SessionRef, self VBDRef) (_retval VbdMode, _err error) {
	return VBDClassGetModeMockedCallback(sessionID, self)
}
// Get the mode field of the given VBD.
func (_class VBDClass) GetMode(sessionID SessionRef, self VBDRef) (_retval VbdMode, _err error) {
	if IsMock {
		return _class.GetModeMock(sessionID, self)
	}	
	_method := "VBD.get_mode"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVBDRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertEnumVbdModeToGo(_method + " -> ", _result.Value)
	return
}


func VBDClassGetBootableMockDefault(sessionID SessionRef, self VBDRef) (_retval bool, _err error) {
	log.Println("VBD.GetBootable not mocked")
	_err = errors.New("VBD.GetBootable not mocked")
	return
}

var VBDClassGetBootableMockedCallback = VBDClassGetBootableMockDefault

func (_class VBDClass) GetBootableMock(sessionID SessionRef, self VBDRef) (_retval bool, _err error) {
	return VBDClassGetBootableMockedCallback(sessionID, self)
}
// Get the bootable field of the given VBD.
func (_class VBDClass) GetBootable(sessionID SessionRef, self VBDRef) (_retval bool, _err error) {
	if IsMock {
		return _class.GetBootableMock(sessionID, self)
	}	
	_method := "VBD.get_bootable"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVBDRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func VBDClassGetUserdeviceMockDefault(sessionID SessionRef, self VBDRef) (_retval string, _err error) {
	log.Println("VBD.GetUserdevice not mocked")
	_err = errors.New("VBD.GetUserdevice not mocked")
	return
}

var VBDClassGetUserdeviceMockedCallback = VBDClassGetUserdeviceMockDefault

func (_class VBDClass) GetUserdeviceMock(sessionID SessionRef, self VBDRef) (_retval string, _err error) {
	return VBDClassGetUserdeviceMockedCallback(sessionID, self)
}
// Get the userdevice field of the given VBD.
func (_class VBDClass) GetUserdevice(sessionID SessionRef, self VBDRef) (_retval string, _err error) {
	if IsMock {
		return _class.GetUserdeviceMock(sessionID, self)
	}	
	_method := "VBD.get_userdevice"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVBDRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func VBDClassGetDeviceMockDefault(sessionID SessionRef, self VBDRef) (_retval string, _err error) {
	log.Println("VBD.GetDevice not mocked")
	_err = errors.New("VBD.GetDevice not mocked")
	return
}

var VBDClassGetDeviceMockedCallback = VBDClassGetDeviceMockDefault

func (_class VBDClass) GetDeviceMock(sessionID SessionRef, self VBDRef) (_retval string, _err error) {
	return VBDClassGetDeviceMockedCallback(sessionID, self)
}
// Get the device field of the given VBD.
func (_class VBDClass) GetDevice(sessionID SessionRef, self VBDRef) (_retval string, _err error) {
	if IsMock {
		return _class.GetDeviceMock(sessionID, self)
	}	
	_method := "VBD.get_device"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVBDRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func VBDClassGetVDIMockDefault(sessionID SessionRef, self VBDRef) (_retval VDIRef, _err error) {
	log.Println("VBD.GetVDI not mocked")
	_err = errors.New("VBD.GetVDI not mocked")
	return
}

var VBDClassGetVDIMockedCallback = VBDClassGetVDIMockDefault

func (_class VBDClass) GetVDIMock(sessionID SessionRef, self VBDRef) (_retval VDIRef, _err error) {
	return VBDClassGetVDIMockedCallback(sessionID, self)
}
// Get the VDI field of the given VBD.
func (_class VBDClass) GetVDI(sessionID SessionRef, self VBDRef) (_retval VDIRef, _err error) {
	if IsMock {
		return _class.GetVDIMock(sessionID, self)
	}	
	_method := "VBD.get_VDI"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVBDRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVDIRefToGo(_method + " -> ", _result.Value)
	return
}


func VBDClassGetVMMockDefault(sessionID SessionRef, self VBDRef) (_retval VMRef, _err error) {
	log.Println("VBD.GetVM not mocked")
	_err = errors.New("VBD.GetVM not mocked")
	return
}

var VBDClassGetVMMockedCallback = VBDClassGetVMMockDefault

func (_class VBDClass) GetVMMock(sessionID SessionRef, self VBDRef) (_retval VMRef, _err error) {
	return VBDClassGetVMMockedCallback(sessionID, self)
}
// Get the VM field of the given VBD.
func (_class VBDClass) GetVM(sessionID SessionRef, self VBDRef) (_retval VMRef, _err error) {
	if IsMock {
		return _class.GetVMMock(sessionID, self)
	}	
	_method := "VBD.get_VM"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVBDRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func VBDClassGetCurrentOperationsMockDefault(sessionID SessionRef, self VBDRef) (_retval map[string]VbdOperations, _err error) {
	log.Println("VBD.GetCurrentOperations not mocked")
	_err = errors.New("VBD.GetCurrentOperations not mocked")
	return
}

var VBDClassGetCurrentOperationsMockedCallback = VBDClassGetCurrentOperationsMockDefault

func (_class VBDClass) GetCurrentOperationsMock(sessionID SessionRef, self VBDRef) (_retval map[string]VbdOperations, _err error) {
	return VBDClassGetCurrentOperationsMockedCallback(sessionID, self)
}
// Get the current_operations field of the given VBD.
func (_class VBDClass) GetCurrentOperations(sessionID SessionRef, self VBDRef) (_retval map[string]VbdOperations, _err error) {
	if IsMock {
		return _class.GetCurrentOperationsMock(sessionID, self)
	}	
	_method := "VBD.get_current_operations"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVBDRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringToEnumVbdOperationsMapToGo(_method + " -> ", _result.Value)
	return
}


func VBDClassGetAllowedOperationsMockDefault(sessionID SessionRef, self VBDRef) (_retval []VbdOperations, _err error) {
	log.Println("VBD.GetAllowedOperations not mocked")
	_err = errors.New("VBD.GetAllowedOperations not mocked")
	return
}

var VBDClassGetAllowedOperationsMockedCallback = VBDClassGetAllowedOperationsMockDefault

func (_class VBDClass) GetAllowedOperationsMock(sessionID SessionRef, self VBDRef) (_retval []VbdOperations, _err error) {
	return VBDClassGetAllowedOperationsMockedCallback(sessionID, self)
}
// Get the allowed_operations field of the given VBD.
func (_class VBDClass) GetAllowedOperations(sessionID SessionRef, self VBDRef) (_retval []VbdOperations, _err error) {
	if IsMock {
		return _class.GetAllowedOperationsMock(sessionID, self)
	}	
	_method := "VBD.get_allowed_operations"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVBDRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertEnumVbdOperationsSetToGo(_method + " -> ", _result.Value)
	return
}


func VBDClassGetUUIDMockDefault(sessionID SessionRef, self VBDRef) (_retval string, _err error) {
	log.Println("VBD.GetUUID not mocked")
	_err = errors.New("VBD.GetUUID not mocked")
	return
}

var VBDClassGetUUIDMockedCallback = VBDClassGetUUIDMockDefault

func (_class VBDClass) GetUUIDMock(sessionID SessionRef, self VBDRef) (_retval string, _err error) {
	return VBDClassGetUUIDMockedCallback(sessionID, self)
}
// Get the uuid field of the given VBD.
func (_class VBDClass) GetUUID(sessionID SessionRef, self VBDRef) (_retval string, _err error) {
	if IsMock {
		return _class.GetUUIDMock(sessionID, self)
	}	
	_method := "VBD.get_uuid"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVBDRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func VBDClassDestroyMockDefault(sessionID SessionRef, self VBDRef) (_err error) {
	log.Println("VBD.Destroy not mocked")
	_err = errors.New("VBD.Destroy not mocked")
	return
}

var VBDClassDestroyMockedCallback = VBDClassDestroyMockDefault

func (_class VBDClass) DestroyMock(sessionID SessionRef, self VBDRef) (_err error) {
	return VBDClassDestroyMockedCallback(sessionID, self)
}
// Destroy the specified VBD instance.
func (_class VBDClass) Destroy(sessionID SessionRef, self VBDRef) (_err error) {
	if IsMock {
		return _class.DestroyMock(sessionID, self)
	}	
	_method := "VBD.destroy"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVBDRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg)
	return
}


func VBDClassCreateMockDefault(sessionID SessionRef, args VBDRecord) (_retval VBDRef, _err error) {
	log.Println("VBD.Create not mocked")
	_err = errors.New("VBD.Create not mocked")
	return
}

var VBDClassCreateMockedCallback = VBDClassCreateMockDefault

func (_class VBDClass) CreateMock(sessionID SessionRef, args VBDRecord) (_retval VBDRef, _err error) {
	return VBDClassCreateMockedCallback(sessionID, args)
}
// Create a new VBD instance, and return its handle.
// The constructor args are: VM*, VDI*, userdevice*, bootable*, mode*, type*, unpluggable, empty*, other_config*, qos_algorithm_type*, qos_algorithm_params* (* = non-optional).
func (_class VBDClass) Create(sessionID SessionRef, args VBDRecord) (_retval VBDRef, _err error) {
	if IsMock {
		return _class.CreateMock(sessionID, args)
	}	
	_method := "VBD.create"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_argsArg, _err := convertVBDRecordToXen(fmt.Sprintf("%s(%s)", _method, "args"), args)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _argsArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVBDRefToGo(_method + " -> ", _result.Value)
	return
}


func VBDClassGetByUUIDMockDefault(sessionID SessionRef, uuid string) (_retval VBDRef, _err error) {
	log.Println("VBD.GetByUUID not mocked")
	_err = errors.New("VBD.GetByUUID not mocked")
	return
}

var VBDClassGetByUUIDMockedCallback = VBDClassGetByUUIDMockDefault

func (_class VBDClass) GetByUUIDMock(sessionID SessionRef, uuid string) (_retval VBDRef, _err error) {
	return VBDClassGetByUUIDMockedCallback(sessionID, uuid)
}
// Get a reference to the VBD instance with the specified UUID.
func (_class VBDClass) GetByUUID(sessionID SessionRef, uuid string) (_retval VBDRef, _err error) {
	if IsMock {
		return _class.GetByUUIDMock(sessionID, uuid)
	}	
	_method := "VBD.get_by_uuid"
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
	_retval, _err = convertVBDRefToGo(_method + " -> ", _result.Value)
	return
}


func VBDClassGetRecordMockDefault(sessionID SessionRef, self VBDRef) (_retval VBDRecord, _err error) {
	log.Println("VBD.GetRecord not mocked")
	_err = errors.New("VBD.GetRecord not mocked")
	return
}

var VBDClassGetRecordMockedCallback = VBDClassGetRecordMockDefault

func (_class VBDClass) GetRecordMock(sessionID SessionRef, self VBDRef) (_retval VBDRecord, _err error) {
	return VBDClassGetRecordMockedCallback(sessionID, self)
}
// Get a record containing the current state of the given VBD.
func (_class VBDClass) GetRecord(sessionID SessionRef, self VBDRef) (_retval VBDRecord, _err error) {
	if IsMock {
		return _class.GetRecordMock(sessionID, self)
	}	
	_method := "VBD.get_record"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVBDRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVBDRecordToGo(_method + " -> ", _result.Value)
	return
}
