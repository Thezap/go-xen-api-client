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

type PCIRecord struct {
  // Unique identifier/object reference
	UUID string
  // PCI class name
	ClassName string
  // Vendor name
	VendorName string
  // Device name
	DeviceName string
  // Physical machine that owns the PCI device
	Host HostRef
  // PCI ID of the physical device
	PciID string
  // List of dependent PCI devices
	Dependencies []PCIRef
  // Additional configuration
	OtherConfig map[string]string
  // Subsystem vendor name
	SubsystemVendorName string
  // Subsystem device name
	SubsystemDeviceName string
}

type PCIRef string

// A PCI device
type PCIClass struct {
	client *Client
}


func PCIClassGetAllRecordsMockDefault(sessionID SessionRef) (_retval map[PCIRef]PCIRecord, _err error) {
	log.Println("PCI.GetAllRecords not mocked")
	_err = errors.New("PCI.GetAllRecords not mocked")
	return
}

var PCIClassGetAllRecordsMockedCallback = PCIClassGetAllRecordsMockDefault

func (_class PCIClass) GetAllRecordsMock(sessionID SessionRef) (_retval map[PCIRef]PCIRecord, _err error) {
	return PCIClassGetAllRecordsMockedCallback(sessionID)
}
// Return a map of PCI references to PCI records for all PCIs known to the system.
func (_class PCIClass) GetAllRecords(sessionID SessionRef) (_retval map[PCIRef]PCIRecord, _err error) {
	if IsMock {
		return _class.GetAllRecordsMock(sessionID)
	}	
	_method := "PCI.get_all_records"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertPCIRefToPCIRecordMapToGo(_method + " -> ", _result.Value)
	return
}


func PCIClassGetAllMockDefault(sessionID SessionRef) (_retval []PCIRef, _err error) {
	log.Println("PCI.GetAll not mocked")
	_err = errors.New("PCI.GetAll not mocked")
	return
}

var PCIClassGetAllMockedCallback = PCIClassGetAllMockDefault

func (_class PCIClass) GetAllMock(sessionID SessionRef) (_retval []PCIRef, _err error) {
	return PCIClassGetAllMockedCallback(sessionID)
}
// Return a list of all the PCIs known to the system.
func (_class PCIClass) GetAll(sessionID SessionRef) (_retval []PCIRef, _err error) {
	if IsMock {
		return _class.GetAllMock(sessionID)
	}	
	_method := "PCI.get_all"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertPCIRefSetToGo(_method + " -> ", _result.Value)
	return
}


func PCIClassRemoveFromOtherConfigMockDefault(sessionID SessionRef, self PCIRef, key string) (_err error) {
	log.Println("PCI.RemoveFromOtherConfig not mocked")
	_err = errors.New("PCI.RemoveFromOtherConfig not mocked")
	return
}

var PCIClassRemoveFromOtherConfigMockedCallback = PCIClassRemoveFromOtherConfigMockDefault

func (_class PCIClass) RemoveFromOtherConfigMock(sessionID SessionRef, self PCIRef, key string) (_err error) {
	return PCIClassRemoveFromOtherConfigMockedCallback(sessionID, self, key)
}
// Remove the given key and its corresponding value from the other_config field of the given PCI.  If the key is not in that Map, then do nothing.
func (_class PCIClass) RemoveFromOtherConfig(sessionID SessionRef, self PCIRef, key string) (_err error) {
	if IsMock {
		return _class.RemoveFromOtherConfigMock(sessionID, self, key)
	}	
	_method := "PCI.remove_from_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPCIRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func PCIClassAddToOtherConfigMockDefault(sessionID SessionRef, self PCIRef, key string, value string) (_err error) {
	log.Println("PCI.AddToOtherConfig not mocked")
	_err = errors.New("PCI.AddToOtherConfig not mocked")
	return
}

var PCIClassAddToOtherConfigMockedCallback = PCIClassAddToOtherConfigMockDefault

func (_class PCIClass) AddToOtherConfigMock(sessionID SessionRef, self PCIRef, key string, value string) (_err error) {
	return PCIClassAddToOtherConfigMockedCallback(sessionID, self, key, value)
}
// Add the given key-value pair to the other_config field of the given PCI.
func (_class PCIClass) AddToOtherConfig(sessionID SessionRef, self PCIRef, key string, value string) (_err error) {
	if IsMock {
		return _class.AddToOtherConfigMock(sessionID, self, key, value)
	}	
	_method := "PCI.add_to_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPCIRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func PCIClassSetOtherConfigMockDefault(sessionID SessionRef, self PCIRef, value map[string]string) (_err error) {
	log.Println("PCI.SetOtherConfig not mocked")
	_err = errors.New("PCI.SetOtherConfig not mocked")
	return
}

var PCIClassSetOtherConfigMockedCallback = PCIClassSetOtherConfigMockDefault

func (_class PCIClass) SetOtherConfigMock(sessionID SessionRef, self PCIRef, value map[string]string) (_err error) {
	return PCIClassSetOtherConfigMockedCallback(sessionID, self, value)
}
// Set the other_config field of the given PCI.
func (_class PCIClass) SetOtherConfig(sessionID SessionRef, self PCIRef, value map[string]string) (_err error) {
	if IsMock {
		return _class.SetOtherConfigMock(sessionID, self, value)
	}	
	_method := "PCI.set_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPCIRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func PCIClassGetSubsystemDeviceNameMockDefault(sessionID SessionRef, self PCIRef) (_retval string, _err error) {
	log.Println("PCI.GetSubsystemDeviceName not mocked")
	_err = errors.New("PCI.GetSubsystemDeviceName not mocked")
	return
}

var PCIClassGetSubsystemDeviceNameMockedCallback = PCIClassGetSubsystemDeviceNameMockDefault

func (_class PCIClass) GetSubsystemDeviceNameMock(sessionID SessionRef, self PCIRef) (_retval string, _err error) {
	return PCIClassGetSubsystemDeviceNameMockedCallback(sessionID, self)
}
// Get the subsystem_device_name field of the given PCI.
func (_class PCIClass) GetSubsystemDeviceName(sessionID SessionRef, self PCIRef) (_retval string, _err error) {
	if IsMock {
		return _class.GetSubsystemDeviceNameMock(sessionID, self)
	}	
	_method := "PCI.get_subsystem_device_name"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPCIRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func PCIClassGetSubsystemVendorNameMockDefault(sessionID SessionRef, self PCIRef) (_retval string, _err error) {
	log.Println("PCI.GetSubsystemVendorName not mocked")
	_err = errors.New("PCI.GetSubsystemVendorName not mocked")
	return
}

var PCIClassGetSubsystemVendorNameMockedCallback = PCIClassGetSubsystemVendorNameMockDefault

func (_class PCIClass) GetSubsystemVendorNameMock(sessionID SessionRef, self PCIRef) (_retval string, _err error) {
	return PCIClassGetSubsystemVendorNameMockedCallback(sessionID, self)
}
// Get the subsystem_vendor_name field of the given PCI.
func (_class PCIClass) GetSubsystemVendorName(sessionID SessionRef, self PCIRef) (_retval string, _err error) {
	if IsMock {
		return _class.GetSubsystemVendorNameMock(sessionID, self)
	}	
	_method := "PCI.get_subsystem_vendor_name"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPCIRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func PCIClassGetOtherConfigMockDefault(sessionID SessionRef, self PCIRef) (_retval map[string]string, _err error) {
	log.Println("PCI.GetOtherConfig not mocked")
	_err = errors.New("PCI.GetOtherConfig not mocked")
	return
}

var PCIClassGetOtherConfigMockedCallback = PCIClassGetOtherConfigMockDefault

func (_class PCIClass) GetOtherConfigMock(sessionID SessionRef, self PCIRef) (_retval map[string]string, _err error) {
	return PCIClassGetOtherConfigMockedCallback(sessionID, self)
}
// Get the other_config field of the given PCI.
func (_class PCIClass) GetOtherConfig(sessionID SessionRef, self PCIRef) (_retval map[string]string, _err error) {
	if IsMock {
		return _class.GetOtherConfigMock(sessionID, self)
	}	
	_method := "PCI.get_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPCIRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func PCIClassGetDependenciesMockDefault(sessionID SessionRef, self PCIRef) (_retval []PCIRef, _err error) {
	log.Println("PCI.GetDependencies not mocked")
	_err = errors.New("PCI.GetDependencies not mocked")
	return
}

var PCIClassGetDependenciesMockedCallback = PCIClassGetDependenciesMockDefault

func (_class PCIClass) GetDependenciesMock(sessionID SessionRef, self PCIRef) (_retval []PCIRef, _err error) {
	return PCIClassGetDependenciesMockedCallback(sessionID, self)
}
// Get the dependencies field of the given PCI.
func (_class PCIClass) GetDependencies(sessionID SessionRef, self PCIRef) (_retval []PCIRef, _err error) {
	if IsMock {
		return _class.GetDependenciesMock(sessionID, self)
	}	
	_method := "PCI.get_dependencies"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPCIRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertPCIRefSetToGo(_method + " -> ", _result.Value)
	return
}


func PCIClassGetPciIDMockDefault(sessionID SessionRef, self PCIRef) (_retval string, _err error) {
	log.Println("PCI.GetPciID not mocked")
	_err = errors.New("PCI.GetPciID not mocked")
	return
}

var PCIClassGetPciIDMockedCallback = PCIClassGetPciIDMockDefault

func (_class PCIClass) GetPciIDMock(sessionID SessionRef, self PCIRef) (_retval string, _err error) {
	return PCIClassGetPciIDMockedCallback(sessionID, self)
}
// Get the pci_id field of the given PCI.
func (_class PCIClass) GetPciID(sessionID SessionRef, self PCIRef) (_retval string, _err error) {
	if IsMock {
		return _class.GetPciIDMock(sessionID, self)
	}	
	_method := "PCI.get_pci_id"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPCIRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func PCIClassGetHostMockDefault(sessionID SessionRef, self PCIRef) (_retval HostRef, _err error) {
	log.Println("PCI.GetHost not mocked")
	_err = errors.New("PCI.GetHost not mocked")
	return
}

var PCIClassGetHostMockedCallback = PCIClassGetHostMockDefault

func (_class PCIClass) GetHostMock(sessionID SessionRef, self PCIRef) (_retval HostRef, _err error) {
	return PCIClassGetHostMockedCallback(sessionID, self)
}
// Get the host field of the given PCI.
func (_class PCIClass) GetHost(sessionID SessionRef, self PCIRef) (_retval HostRef, _err error) {
	if IsMock {
		return _class.GetHostMock(sessionID, self)
	}	
	_method := "PCI.get_host"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPCIRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func PCIClassGetDeviceNameMockDefault(sessionID SessionRef, self PCIRef) (_retval string, _err error) {
	log.Println("PCI.GetDeviceName not mocked")
	_err = errors.New("PCI.GetDeviceName not mocked")
	return
}

var PCIClassGetDeviceNameMockedCallback = PCIClassGetDeviceNameMockDefault

func (_class PCIClass) GetDeviceNameMock(sessionID SessionRef, self PCIRef) (_retval string, _err error) {
	return PCIClassGetDeviceNameMockedCallback(sessionID, self)
}
// Get the device_name field of the given PCI.
func (_class PCIClass) GetDeviceName(sessionID SessionRef, self PCIRef) (_retval string, _err error) {
	if IsMock {
		return _class.GetDeviceNameMock(sessionID, self)
	}	
	_method := "PCI.get_device_name"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPCIRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func PCIClassGetVendorNameMockDefault(sessionID SessionRef, self PCIRef) (_retval string, _err error) {
	log.Println("PCI.GetVendorName not mocked")
	_err = errors.New("PCI.GetVendorName not mocked")
	return
}

var PCIClassGetVendorNameMockedCallback = PCIClassGetVendorNameMockDefault

func (_class PCIClass) GetVendorNameMock(sessionID SessionRef, self PCIRef) (_retval string, _err error) {
	return PCIClassGetVendorNameMockedCallback(sessionID, self)
}
// Get the vendor_name field of the given PCI.
func (_class PCIClass) GetVendorName(sessionID SessionRef, self PCIRef) (_retval string, _err error) {
	if IsMock {
		return _class.GetVendorNameMock(sessionID, self)
	}	
	_method := "PCI.get_vendor_name"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPCIRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func PCIClassGetClassNameMockDefault(sessionID SessionRef, self PCIRef) (_retval string, _err error) {
	log.Println("PCI.GetClassName not mocked")
	_err = errors.New("PCI.GetClassName not mocked")
	return
}

var PCIClassGetClassNameMockedCallback = PCIClassGetClassNameMockDefault

func (_class PCIClass) GetClassNameMock(sessionID SessionRef, self PCIRef) (_retval string, _err error) {
	return PCIClassGetClassNameMockedCallback(sessionID, self)
}
// Get the class_name field of the given PCI.
func (_class PCIClass) GetClassName(sessionID SessionRef, self PCIRef) (_retval string, _err error) {
	if IsMock {
		return _class.GetClassNameMock(sessionID, self)
	}	
	_method := "PCI.get_class_name"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPCIRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func PCIClassGetUUIDMockDefault(sessionID SessionRef, self PCIRef) (_retval string, _err error) {
	log.Println("PCI.GetUUID not mocked")
	_err = errors.New("PCI.GetUUID not mocked")
	return
}

var PCIClassGetUUIDMockedCallback = PCIClassGetUUIDMockDefault

func (_class PCIClass) GetUUIDMock(sessionID SessionRef, self PCIRef) (_retval string, _err error) {
	return PCIClassGetUUIDMockedCallback(sessionID, self)
}
// Get the uuid field of the given PCI.
func (_class PCIClass) GetUUID(sessionID SessionRef, self PCIRef) (_retval string, _err error) {
	if IsMock {
		return _class.GetUUIDMock(sessionID, self)
	}	
	_method := "PCI.get_uuid"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPCIRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func PCIClassGetByUUIDMockDefault(sessionID SessionRef, uuid string) (_retval PCIRef, _err error) {
	log.Println("PCI.GetByUUID not mocked")
	_err = errors.New("PCI.GetByUUID not mocked")
	return
}

var PCIClassGetByUUIDMockedCallback = PCIClassGetByUUIDMockDefault

func (_class PCIClass) GetByUUIDMock(sessionID SessionRef, uuid string) (_retval PCIRef, _err error) {
	return PCIClassGetByUUIDMockedCallback(sessionID, uuid)
}
// Get a reference to the PCI instance with the specified UUID.
func (_class PCIClass) GetByUUID(sessionID SessionRef, uuid string) (_retval PCIRef, _err error) {
	if IsMock {
		return _class.GetByUUIDMock(sessionID, uuid)
	}	
	_method := "PCI.get_by_uuid"
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
	_retval, _err = convertPCIRefToGo(_method + " -> ", _result.Value)
	return
}


func PCIClassGetRecordMockDefault(sessionID SessionRef, self PCIRef) (_retval PCIRecord, _err error) {
	log.Println("PCI.GetRecord not mocked")
	_err = errors.New("PCI.GetRecord not mocked")
	return
}

var PCIClassGetRecordMockedCallback = PCIClassGetRecordMockDefault

func (_class PCIClass) GetRecordMock(sessionID SessionRef, self PCIRef) (_retval PCIRecord, _err error) {
	return PCIClassGetRecordMockedCallback(sessionID, self)
}
// Get a record containing the current state of the given PCI.
func (_class PCIClass) GetRecord(sessionID SessionRef, self PCIRef) (_retval PCIRecord, _err error) {
	if IsMock {
		return _class.GetRecordMock(sessionID, self)
	}	
	_method := "PCI.get_record"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPCIRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertPCIRecordToGo(_method + " -> ", _result.Value)
	return
}
