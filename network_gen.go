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

type NetworkOperations string

const (
  // Indicates this network is attaching to a VIF or PIF
	NetworkOperationsAttaching NetworkOperations = "attaching"
)

type NetworkDefaultLockingMode string

const (
  // Treat all VIFs on this network with locking_mode = 'default' as if they have locking_mode = 'unlocked'
	NetworkDefaultLockingModeUnlocked NetworkDefaultLockingMode = "unlocked"
  // Treat all VIFs on this network with locking_mode = 'default' as if they have locking_mode = 'disabled'
	NetworkDefaultLockingModeDisabled NetworkDefaultLockingMode = "disabled"
)

type NetworkPurpose string

const (
  // Network Block Device service using TLS
	NetworkPurposeNbd NetworkPurpose = "nbd"
  // Network Block Device service without integrity or confidentiality: NOT RECOMMENDED
	NetworkPurposeInsecureNbd NetworkPurpose = "insecure_nbd"
)

type NetworkRecord struct {
  // Unique identifier/object reference
	UUID string
  // a human-readable name
	NameLabel string
  // a notes field containing human-readable description
	NameDescription string
  // list of the operations allowed in this state. This list is advisory only and the server state may have changed by the time this field is read by a client.
	AllowedOperations []NetworkOperations
  // links each of the running tasks using this object (by reference) to a current_operation enum which describes the nature of the task.
	CurrentOperations map[string]NetworkOperations
  // list of connected vifs
	VIFs []VIFRef
  // list of connected pifs
	PIFs []PIFRef
  // MTU in octets
	MTU int
  // additional configuration
	OtherConfig map[string]string
  // name of the bridge corresponding to this network on the local host
	Bridge string
  // true if the bridge is managed by xapi
	Managed bool
  // Binary blobs associated with this network
	Blobs map[string]BlobRef
  // user-specified tags for categorization purposes
	Tags []string
  // The network will use this value to determine the behaviour of all VIFs where locking_mode = default
	DefaultLockingMode NetworkDefaultLockingMode
  // The IP addresses assigned to VIFs on networks that have active xapi-managed DHCP
	AssignedIps map[VIFRef]string
  // Set of purposes for which the server will use this network
	Purpose []NetworkPurpose
}

type NetworkRef string

// A virtual network
type NetworkClass struct {
	client *Client
}


var NetworkClass_GetAllRecordsMockedCallback = func (sessionID SessionRef) (_retval map[NetworkRef]NetworkRecord, _err error) {
	log.Println("Network.GetAllRecords not mocked")
	_err = errors.New("Network.GetAllRecords not mocked")
	return
}

func (_class NetworkClass) GetAllRecordsMock(sessionID SessionRef) (_retval map[NetworkRef]NetworkRecord, _err error) {
	return NetworkClass_GetAllRecordsMockedCallback(sessionID)
}
// Return a map of network references to network records for all networks known to the system.
func (_class NetworkClass) GetAllRecords(sessionID SessionRef) (_retval map[NetworkRef]NetworkRecord, _err error) {
	if (IsMock) {
		return _class.GetAllRecordsMock(sessionID)
	}	
	_method := "network.get_all_records"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertNetworkRefToNetworkRecordMapToGo(_method + " -> ", _result.Value)
	return
}


var NetworkClass_GetAllMockedCallback = func (sessionID SessionRef) (_retval []NetworkRef, _err error) {
	log.Println("Network.GetAll not mocked")
	_err = errors.New("Network.GetAll not mocked")
	return
}

func (_class NetworkClass) GetAllMock(sessionID SessionRef) (_retval []NetworkRef, _err error) {
	return NetworkClass_GetAllMockedCallback(sessionID)
}
// Return a list of all the networks known to the system.
func (_class NetworkClass) GetAll(sessionID SessionRef) (_retval []NetworkRef, _err error) {
	if (IsMock) {
		return _class.GetAllMock(sessionID)
	}	
	_method := "network.get_all"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertNetworkRefSetToGo(_method + " -> ", _result.Value)
	return
}


var NetworkClass_RemovePurposeMockedCallback = func (sessionID SessionRef, self NetworkRef, value NetworkPurpose) (_err error) {
	log.Println("Network.RemovePurpose not mocked")
	_err = errors.New("Network.RemovePurpose not mocked")
	return
}

func (_class NetworkClass) RemovePurposeMock(sessionID SessionRef, self NetworkRef, value NetworkPurpose) (_err error) {
	return NetworkClass_RemovePurposeMockedCallback(sessionID, self, value)
}
// Remove a purpose from a network (if present)
func (_class NetworkClass) RemovePurpose(sessionID SessionRef, self NetworkRef, value NetworkPurpose) (_err error) {
	if (IsMock) {
		return _class.RemovePurposeMock(sessionID, self, value)
	}	
	_method := "network.remove_purpose"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertNetworkRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertEnumNetworkPurposeToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}


var NetworkClass_AddPurposeMockedCallback = func (sessionID SessionRef, self NetworkRef, value NetworkPurpose) (_err error) {
	log.Println("Network.AddPurpose not mocked")
	_err = errors.New("Network.AddPurpose not mocked")
	return
}

func (_class NetworkClass) AddPurposeMock(sessionID SessionRef, self NetworkRef, value NetworkPurpose) (_err error) {
	return NetworkClass_AddPurposeMockedCallback(sessionID, self, value)
}
// Give a network a new purpose (if not present already)
//
// Errors:
//  NETWORK_INCOMPATIBLE_PURPOSES - You tried to add a purpose to a network but the new purpose is not compatible with an existing purpose of the network or other networks.
func (_class NetworkClass) AddPurpose(sessionID SessionRef, self NetworkRef, value NetworkPurpose) (_err error) {
	if (IsMock) {
		return _class.AddPurposeMock(sessionID, self, value)
	}	
	_method := "network.add_purpose"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertNetworkRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertEnumNetworkPurposeToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}


var NetworkClass_SetDefaultLockingModeMockedCallback = func (sessionID SessionRef, network NetworkRef, value NetworkDefaultLockingMode) (_err error) {
	log.Println("Network.SetDefaultLockingMode not mocked")
	_err = errors.New("Network.SetDefaultLockingMode not mocked")
	return
}

func (_class NetworkClass) SetDefaultLockingModeMock(sessionID SessionRef, network NetworkRef, value NetworkDefaultLockingMode) (_err error) {
	return NetworkClass_SetDefaultLockingModeMockedCallback(sessionID, network, value)
}
// Set the default locking mode for VIFs attached to this network
func (_class NetworkClass) SetDefaultLockingMode(sessionID SessionRef, network NetworkRef, value NetworkDefaultLockingMode) (_err error) {
	if (IsMock) {
		return _class.SetDefaultLockingModeMock(sessionID, network, value)
	}	
	_method := "network.set_default_locking_mode"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_networkArg, _err := convertNetworkRefToXen(fmt.Sprintf("%s(%s)", _method, "network"), network)
	if _err != nil {
		return
	}
	_valueArg, _err := convertEnumNetworkDefaultLockingModeToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _networkArg, _valueArg)
	return
}


var NetworkClass_CreateNewBlobMockedCallback = func (sessionID SessionRef, network NetworkRef, name string, mimeType string, public bool) (_retval BlobRef, _err error) {
	log.Println("Network.CreateNewBlob not mocked")
	_err = errors.New("Network.CreateNewBlob not mocked")
	return
}

func (_class NetworkClass) CreateNewBlobMock(sessionID SessionRef, network NetworkRef, name string, mimeType string, public bool) (_retval BlobRef, _err error) {
	return NetworkClass_CreateNewBlobMockedCallback(sessionID, network, name, mimeType, public)
}
// Create a placeholder for a named binary blob of data that is associated with this pool
func (_class NetworkClass) CreateNewBlob(sessionID SessionRef, network NetworkRef, name string, mimeType string, public bool) (_retval BlobRef, _err error) {
	if (IsMock) {
		return _class.CreateNewBlobMock(sessionID, network, name, mimeType, public)
	}	
	_method := "network.create_new_blob"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_networkArg, _err := convertNetworkRefToXen(fmt.Sprintf("%s(%s)", _method, "network"), network)
	if _err != nil {
		return
	}
	_nameArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "name"), name)
	if _err != nil {
		return
	}
	_mimeTypeArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "mime_type"), mimeType)
	if _err != nil {
		return
	}
	_publicArg, _err := convertBoolToXen(fmt.Sprintf("%s(%s)", _method, "public"), public)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _networkArg, _nameArg, _mimeTypeArg, _publicArg)
	if _err != nil {
		return
	}
	_retval, _err = convertBlobRefToGo(_method + " -> ", _result.Value)
	return
}


var NetworkClass_RemoveTagsMockedCallback = func (sessionID SessionRef, self NetworkRef, value string) (_err error) {
	log.Println("Network.RemoveTags not mocked")
	_err = errors.New("Network.RemoveTags not mocked")
	return
}

func (_class NetworkClass) RemoveTagsMock(sessionID SessionRef, self NetworkRef, value string) (_err error) {
	return NetworkClass_RemoveTagsMockedCallback(sessionID, self, value)
}
// Remove the given value from the tags field of the given network.  If the value is not in that Set, then do nothing.
func (_class NetworkClass) RemoveTags(sessionID SessionRef, self NetworkRef, value string) (_err error) {
	if (IsMock) {
		return _class.RemoveTagsMock(sessionID, self, value)
	}	
	_method := "network.remove_tags"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertNetworkRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var NetworkClass_AddTagsMockedCallback = func (sessionID SessionRef, self NetworkRef, value string) (_err error) {
	log.Println("Network.AddTags not mocked")
	_err = errors.New("Network.AddTags not mocked")
	return
}

func (_class NetworkClass) AddTagsMock(sessionID SessionRef, self NetworkRef, value string) (_err error) {
	return NetworkClass_AddTagsMockedCallback(sessionID, self, value)
}
// Add the given value to the tags field of the given network.  If the value is already in that Set, then do nothing.
func (_class NetworkClass) AddTags(sessionID SessionRef, self NetworkRef, value string) (_err error) {
	if (IsMock) {
		return _class.AddTagsMock(sessionID, self, value)
	}	
	_method := "network.add_tags"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertNetworkRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var NetworkClass_SetTagsMockedCallback = func (sessionID SessionRef, self NetworkRef, value []string) (_err error) {
	log.Println("Network.SetTags not mocked")
	_err = errors.New("Network.SetTags not mocked")
	return
}

func (_class NetworkClass) SetTagsMock(sessionID SessionRef, self NetworkRef, value []string) (_err error) {
	return NetworkClass_SetTagsMockedCallback(sessionID, self, value)
}
// Set the tags field of the given network.
func (_class NetworkClass) SetTags(sessionID SessionRef, self NetworkRef, value []string) (_err error) {
	if (IsMock) {
		return _class.SetTagsMock(sessionID, self, value)
	}	
	_method := "network.set_tags"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertNetworkRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertStringSetToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}


var NetworkClass_RemoveFromOtherConfigMockedCallback = func (sessionID SessionRef, self NetworkRef, key string) (_err error) {
	log.Println("Network.RemoveFromOtherConfig not mocked")
	_err = errors.New("Network.RemoveFromOtherConfig not mocked")
	return
}

func (_class NetworkClass) RemoveFromOtherConfigMock(sessionID SessionRef, self NetworkRef, key string) (_err error) {
	return NetworkClass_RemoveFromOtherConfigMockedCallback(sessionID, self, key)
}
// Remove the given key and its corresponding value from the other_config field of the given network.  If the key is not in that Map, then do nothing.
func (_class NetworkClass) RemoveFromOtherConfig(sessionID SessionRef, self NetworkRef, key string) (_err error) {
	if (IsMock) {
		return _class.RemoveFromOtherConfigMock(sessionID, self, key)
	}	
	_method := "network.remove_from_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertNetworkRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var NetworkClass_AddToOtherConfigMockedCallback = func (sessionID SessionRef, self NetworkRef, key string, value string) (_err error) {
	log.Println("Network.AddToOtherConfig not mocked")
	_err = errors.New("Network.AddToOtherConfig not mocked")
	return
}

func (_class NetworkClass) AddToOtherConfigMock(sessionID SessionRef, self NetworkRef, key string, value string) (_err error) {
	return NetworkClass_AddToOtherConfigMockedCallback(sessionID, self, key, value)
}
// Add the given key-value pair to the other_config field of the given network.
func (_class NetworkClass) AddToOtherConfig(sessionID SessionRef, self NetworkRef, key string, value string) (_err error) {
	if (IsMock) {
		return _class.AddToOtherConfigMock(sessionID, self, key, value)
	}	
	_method := "network.add_to_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertNetworkRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var NetworkClass_SetOtherConfigMockedCallback = func (sessionID SessionRef, self NetworkRef, value map[string]string) (_err error) {
	log.Println("Network.SetOtherConfig not mocked")
	_err = errors.New("Network.SetOtherConfig not mocked")
	return
}

func (_class NetworkClass) SetOtherConfigMock(sessionID SessionRef, self NetworkRef, value map[string]string) (_err error) {
	return NetworkClass_SetOtherConfigMockedCallback(sessionID, self, value)
}
// Set the other_config field of the given network.
func (_class NetworkClass) SetOtherConfig(sessionID SessionRef, self NetworkRef, value map[string]string) (_err error) {
	if (IsMock) {
		return _class.SetOtherConfigMock(sessionID, self, value)
	}	
	_method := "network.set_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertNetworkRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var NetworkClass_SetMTUMockedCallback = func (sessionID SessionRef, self NetworkRef, value int) (_err error) {
	log.Println("Network.SetMTU not mocked")
	_err = errors.New("Network.SetMTU not mocked")
	return
}

func (_class NetworkClass) SetMTUMock(sessionID SessionRef, self NetworkRef, value int) (_err error) {
	return NetworkClass_SetMTUMockedCallback(sessionID, self, value)
}
// Set the MTU field of the given network.
func (_class NetworkClass) SetMTU(sessionID SessionRef, self NetworkRef, value int) (_err error) {
	if (IsMock) {
		return _class.SetMTUMock(sessionID, self, value)
	}	
	_method := "network.set_MTU"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertNetworkRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertIntToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}


var NetworkClass_SetNameDescriptionMockedCallback = func (sessionID SessionRef, self NetworkRef, value string) (_err error) {
	log.Println("Network.SetNameDescription not mocked")
	_err = errors.New("Network.SetNameDescription not mocked")
	return
}

func (_class NetworkClass) SetNameDescriptionMock(sessionID SessionRef, self NetworkRef, value string) (_err error) {
	return NetworkClass_SetNameDescriptionMockedCallback(sessionID, self, value)
}
// Set the name/description field of the given network.
func (_class NetworkClass) SetNameDescription(sessionID SessionRef, self NetworkRef, value string) (_err error) {
	if (IsMock) {
		return _class.SetNameDescriptionMock(sessionID, self, value)
	}	
	_method := "network.set_name_description"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertNetworkRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var NetworkClass_SetNameLabelMockedCallback = func (sessionID SessionRef, self NetworkRef, value string) (_err error) {
	log.Println("Network.SetNameLabel not mocked")
	_err = errors.New("Network.SetNameLabel not mocked")
	return
}

func (_class NetworkClass) SetNameLabelMock(sessionID SessionRef, self NetworkRef, value string) (_err error) {
	return NetworkClass_SetNameLabelMockedCallback(sessionID, self, value)
}
// Set the name/label field of the given network.
func (_class NetworkClass) SetNameLabel(sessionID SessionRef, self NetworkRef, value string) (_err error) {
	if (IsMock) {
		return _class.SetNameLabelMock(sessionID, self, value)
	}	
	_method := "network.set_name_label"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertNetworkRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var NetworkClass_GetPurposeMockedCallback = func (sessionID SessionRef, self NetworkRef) (_retval []NetworkPurpose, _err error) {
	log.Println("Network.GetPurpose not mocked")
	_err = errors.New("Network.GetPurpose not mocked")
	return
}

func (_class NetworkClass) GetPurposeMock(sessionID SessionRef, self NetworkRef) (_retval []NetworkPurpose, _err error) {
	return NetworkClass_GetPurposeMockedCallback(sessionID, self)
}
// Get the purpose field of the given network.
func (_class NetworkClass) GetPurpose(sessionID SessionRef, self NetworkRef) (_retval []NetworkPurpose, _err error) {
	if (IsMock) {
		return _class.GetPurposeMock(sessionID, self)
	}	
	_method := "network.get_purpose"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertNetworkRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertEnumNetworkPurposeSetToGo(_method + " -> ", _result.Value)
	return
}


var NetworkClass_GetAssignedIpsMockedCallback = func (sessionID SessionRef, self NetworkRef) (_retval map[VIFRef]string, _err error) {
	log.Println("Network.GetAssignedIps not mocked")
	_err = errors.New("Network.GetAssignedIps not mocked")
	return
}

func (_class NetworkClass) GetAssignedIpsMock(sessionID SessionRef, self NetworkRef) (_retval map[VIFRef]string, _err error) {
	return NetworkClass_GetAssignedIpsMockedCallback(sessionID, self)
}
// Get the assigned_ips field of the given network.
func (_class NetworkClass) GetAssignedIps(sessionID SessionRef, self NetworkRef) (_retval map[VIFRef]string, _err error) {
	if (IsMock) {
		return _class.GetAssignedIpsMock(sessionID, self)
	}	
	_method := "network.get_assigned_ips"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertNetworkRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVIFRefToStringMapToGo(_method + " -> ", _result.Value)
	return
}


var NetworkClass_GetDefaultLockingModeMockedCallback = func (sessionID SessionRef, self NetworkRef) (_retval NetworkDefaultLockingMode, _err error) {
	log.Println("Network.GetDefaultLockingMode not mocked")
	_err = errors.New("Network.GetDefaultLockingMode not mocked")
	return
}

func (_class NetworkClass) GetDefaultLockingModeMock(sessionID SessionRef, self NetworkRef) (_retval NetworkDefaultLockingMode, _err error) {
	return NetworkClass_GetDefaultLockingModeMockedCallback(sessionID, self)
}
// Get the default_locking_mode field of the given network.
func (_class NetworkClass) GetDefaultLockingMode(sessionID SessionRef, self NetworkRef) (_retval NetworkDefaultLockingMode, _err error) {
	if (IsMock) {
		return _class.GetDefaultLockingModeMock(sessionID, self)
	}	
	_method := "network.get_default_locking_mode"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertNetworkRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertEnumNetworkDefaultLockingModeToGo(_method + " -> ", _result.Value)
	return
}


var NetworkClass_GetTagsMockedCallback = func (sessionID SessionRef, self NetworkRef) (_retval []string, _err error) {
	log.Println("Network.GetTags not mocked")
	_err = errors.New("Network.GetTags not mocked")
	return
}

func (_class NetworkClass) GetTagsMock(sessionID SessionRef, self NetworkRef) (_retval []string, _err error) {
	return NetworkClass_GetTagsMockedCallback(sessionID, self)
}
// Get the tags field of the given network.
func (_class NetworkClass) GetTags(sessionID SessionRef, self NetworkRef) (_retval []string, _err error) {
	if (IsMock) {
		return _class.GetTagsMock(sessionID, self)
	}	
	_method := "network.get_tags"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertNetworkRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var NetworkClass_GetBlobsMockedCallback = func (sessionID SessionRef, self NetworkRef) (_retval map[string]BlobRef, _err error) {
	log.Println("Network.GetBlobs not mocked")
	_err = errors.New("Network.GetBlobs not mocked")
	return
}

func (_class NetworkClass) GetBlobsMock(sessionID SessionRef, self NetworkRef) (_retval map[string]BlobRef, _err error) {
	return NetworkClass_GetBlobsMockedCallback(sessionID, self)
}
// Get the blobs field of the given network.
func (_class NetworkClass) GetBlobs(sessionID SessionRef, self NetworkRef) (_retval map[string]BlobRef, _err error) {
	if (IsMock) {
		return _class.GetBlobsMock(sessionID, self)
	}	
	_method := "network.get_blobs"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertNetworkRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringToBlobRefMapToGo(_method + " -> ", _result.Value)
	return
}


var NetworkClass_GetManagedMockedCallback = func (sessionID SessionRef, self NetworkRef) (_retval bool, _err error) {
	log.Println("Network.GetManaged not mocked")
	_err = errors.New("Network.GetManaged not mocked")
	return
}

func (_class NetworkClass) GetManagedMock(sessionID SessionRef, self NetworkRef) (_retval bool, _err error) {
	return NetworkClass_GetManagedMockedCallback(sessionID, self)
}
// Get the managed field of the given network.
func (_class NetworkClass) GetManaged(sessionID SessionRef, self NetworkRef) (_retval bool, _err error) {
	if (IsMock) {
		return _class.GetManagedMock(sessionID, self)
	}	
	_method := "network.get_managed"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertNetworkRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var NetworkClass_GetBridgeMockedCallback = func (sessionID SessionRef, self NetworkRef) (_retval string, _err error) {
	log.Println("Network.GetBridge not mocked")
	_err = errors.New("Network.GetBridge not mocked")
	return
}

func (_class NetworkClass) GetBridgeMock(sessionID SessionRef, self NetworkRef) (_retval string, _err error) {
	return NetworkClass_GetBridgeMockedCallback(sessionID, self)
}
// Get the bridge field of the given network.
func (_class NetworkClass) GetBridge(sessionID SessionRef, self NetworkRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetBridgeMock(sessionID, self)
	}	
	_method := "network.get_bridge"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertNetworkRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var NetworkClass_GetOtherConfigMockedCallback = func (sessionID SessionRef, self NetworkRef) (_retval map[string]string, _err error) {
	log.Println("Network.GetOtherConfig not mocked")
	_err = errors.New("Network.GetOtherConfig not mocked")
	return
}

func (_class NetworkClass) GetOtherConfigMock(sessionID SessionRef, self NetworkRef) (_retval map[string]string, _err error) {
	return NetworkClass_GetOtherConfigMockedCallback(sessionID, self)
}
// Get the other_config field of the given network.
func (_class NetworkClass) GetOtherConfig(sessionID SessionRef, self NetworkRef) (_retval map[string]string, _err error) {
	if (IsMock) {
		return _class.GetOtherConfigMock(sessionID, self)
	}	
	_method := "network.get_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertNetworkRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var NetworkClass_GetMTUMockedCallback = func (sessionID SessionRef, self NetworkRef) (_retval int, _err error) {
	log.Println("Network.GetMTU not mocked")
	_err = errors.New("Network.GetMTU not mocked")
	return
}

func (_class NetworkClass) GetMTUMock(sessionID SessionRef, self NetworkRef) (_retval int, _err error) {
	return NetworkClass_GetMTUMockedCallback(sessionID, self)
}
// Get the MTU field of the given network.
func (_class NetworkClass) GetMTU(sessionID SessionRef, self NetworkRef) (_retval int, _err error) {
	if (IsMock) {
		return _class.GetMTUMock(sessionID, self)
	}	
	_method := "network.get_MTU"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertNetworkRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var NetworkClass_GetPIFsMockedCallback = func (sessionID SessionRef, self NetworkRef) (_retval []PIFRef, _err error) {
	log.Println("Network.GetPIFs not mocked")
	_err = errors.New("Network.GetPIFs not mocked")
	return
}

func (_class NetworkClass) GetPIFsMock(sessionID SessionRef, self NetworkRef) (_retval []PIFRef, _err error) {
	return NetworkClass_GetPIFsMockedCallback(sessionID, self)
}
// Get the PIFs field of the given network.
func (_class NetworkClass) GetPIFs(sessionID SessionRef, self NetworkRef) (_retval []PIFRef, _err error) {
	if (IsMock) {
		return _class.GetPIFsMock(sessionID, self)
	}	
	_method := "network.get_PIFs"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertNetworkRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertPIFRefSetToGo(_method + " -> ", _result.Value)
	return
}


var NetworkClass_GetVIFsMockedCallback = func (sessionID SessionRef, self NetworkRef) (_retval []VIFRef, _err error) {
	log.Println("Network.GetVIFs not mocked")
	_err = errors.New("Network.GetVIFs not mocked")
	return
}

func (_class NetworkClass) GetVIFsMock(sessionID SessionRef, self NetworkRef) (_retval []VIFRef, _err error) {
	return NetworkClass_GetVIFsMockedCallback(sessionID, self)
}
// Get the VIFs field of the given network.
func (_class NetworkClass) GetVIFs(sessionID SessionRef, self NetworkRef) (_retval []VIFRef, _err error) {
	if (IsMock) {
		return _class.GetVIFsMock(sessionID, self)
	}	
	_method := "network.get_VIFs"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertNetworkRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVIFRefSetToGo(_method + " -> ", _result.Value)
	return
}


var NetworkClass_GetCurrentOperationsMockedCallback = func (sessionID SessionRef, self NetworkRef) (_retval map[string]NetworkOperations, _err error) {
	log.Println("Network.GetCurrentOperations not mocked")
	_err = errors.New("Network.GetCurrentOperations not mocked")
	return
}

func (_class NetworkClass) GetCurrentOperationsMock(sessionID SessionRef, self NetworkRef) (_retval map[string]NetworkOperations, _err error) {
	return NetworkClass_GetCurrentOperationsMockedCallback(sessionID, self)
}
// Get the current_operations field of the given network.
func (_class NetworkClass) GetCurrentOperations(sessionID SessionRef, self NetworkRef) (_retval map[string]NetworkOperations, _err error) {
	if (IsMock) {
		return _class.GetCurrentOperationsMock(sessionID, self)
	}	
	_method := "network.get_current_operations"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertNetworkRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringToEnumNetworkOperationsMapToGo(_method + " -> ", _result.Value)
	return
}


var NetworkClass_GetAllowedOperationsMockedCallback = func (sessionID SessionRef, self NetworkRef) (_retval []NetworkOperations, _err error) {
	log.Println("Network.GetAllowedOperations not mocked")
	_err = errors.New("Network.GetAllowedOperations not mocked")
	return
}

func (_class NetworkClass) GetAllowedOperationsMock(sessionID SessionRef, self NetworkRef) (_retval []NetworkOperations, _err error) {
	return NetworkClass_GetAllowedOperationsMockedCallback(sessionID, self)
}
// Get the allowed_operations field of the given network.
func (_class NetworkClass) GetAllowedOperations(sessionID SessionRef, self NetworkRef) (_retval []NetworkOperations, _err error) {
	if (IsMock) {
		return _class.GetAllowedOperationsMock(sessionID, self)
	}	
	_method := "network.get_allowed_operations"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertNetworkRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertEnumNetworkOperationsSetToGo(_method + " -> ", _result.Value)
	return
}


var NetworkClass_GetNameDescriptionMockedCallback = func (sessionID SessionRef, self NetworkRef) (_retval string, _err error) {
	log.Println("Network.GetNameDescription not mocked")
	_err = errors.New("Network.GetNameDescription not mocked")
	return
}

func (_class NetworkClass) GetNameDescriptionMock(sessionID SessionRef, self NetworkRef) (_retval string, _err error) {
	return NetworkClass_GetNameDescriptionMockedCallback(sessionID, self)
}
// Get the name/description field of the given network.
func (_class NetworkClass) GetNameDescription(sessionID SessionRef, self NetworkRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetNameDescriptionMock(sessionID, self)
	}	
	_method := "network.get_name_description"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertNetworkRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var NetworkClass_GetNameLabelMockedCallback = func (sessionID SessionRef, self NetworkRef) (_retval string, _err error) {
	log.Println("Network.GetNameLabel not mocked")
	_err = errors.New("Network.GetNameLabel not mocked")
	return
}

func (_class NetworkClass) GetNameLabelMock(sessionID SessionRef, self NetworkRef) (_retval string, _err error) {
	return NetworkClass_GetNameLabelMockedCallback(sessionID, self)
}
// Get the name/label field of the given network.
func (_class NetworkClass) GetNameLabel(sessionID SessionRef, self NetworkRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetNameLabelMock(sessionID, self)
	}	
	_method := "network.get_name_label"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertNetworkRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var NetworkClass_GetUUIDMockedCallback = func (sessionID SessionRef, self NetworkRef) (_retval string, _err error) {
	log.Println("Network.GetUUID not mocked")
	_err = errors.New("Network.GetUUID not mocked")
	return
}

func (_class NetworkClass) GetUUIDMock(sessionID SessionRef, self NetworkRef) (_retval string, _err error) {
	return NetworkClass_GetUUIDMockedCallback(sessionID, self)
}
// Get the uuid field of the given network.
func (_class NetworkClass) GetUUID(sessionID SessionRef, self NetworkRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetUUIDMock(sessionID, self)
	}	
	_method := "network.get_uuid"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertNetworkRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var NetworkClass_GetByNameLabelMockedCallback = func (sessionID SessionRef, label string) (_retval []NetworkRef, _err error) {
	log.Println("Network.GetByNameLabel not mocked")
	_err = errors.New("Network.GetByNameLabel not mocked")
	return
}

func (_class NetworkClass) GetByNameLabelMock(sessionID SessionRef, label string) (_retval []NetworkRef, _err error) {
	return NetworkClass_GetByNameLabelMockedCallback(sessionID, label)
}
// Get all the network instances with the given label.
func (_class NetworkClass) GetByNameLabel(sessionID SessionRef, label string) (_retval []NetworkRef, _err error) {
	if (IsMock) {
		return _class.GetByNameLabelMock(sessionID, label)
	}	
	_method := "network.get_by_name_label"
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
	_retval, _err = convertNetworkRefSetToGo(_method + " -> ", _result.Value)
	return
}


var NetworkClass_DestroyMockedCallback = func (sessionID SessionRef, self NetworkRef) (_err error) {
	log.Println("Network.Destroy not mocked")
	_err = errors.New("Network.Destroy not mocked")
	return
}

func (_class NetworkClass) DestroyMock(sessionID SessionRef, self NetworkRef) (_err error) {
	return NetworkClass_DestroyMockedCallback(sessionID, self)
}
// Destroy the specified network instance.
func (_class NetworkClass) Destroy(sessionID SessionRef, self NetworkRef) (_err error) {
	if (IsMock) {
		return _class.DestroyMock(sessionID, self)
	}	
	_method := "network.destroy"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertNetworkRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg)
	return
}


var NetworkClass_CreateMockedCallback = func (sessionID SessionRef, args NetworkRecord) (_retval NetworkRef, _err error) {
	log.Println("Network.Create not mocked")
	_err = errors.New("Network.Create not mocked")
	return
}

func (_class NetworkClass) CreateMock(sessionID SessionRef, args NetworkRecord) (_retval NetworkRef, _err error) {
	return NetworkClass_CreateMockedCallback(sessionID, args)
}
// Create a new network instance, and return its handle.
// The constructor args are: name_label, name_description, MTU, other_config*, bridge, managed, tags (* = non-optional).
func (_class NetworkClass) Create(sessionID SessionRef, args NetworkRecord) (_retval NetworkRef, _err error) {
	if (IsMock) {
		return _class.CreateMock(sessionID, args)
	}	
	_method := "network.create"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_argsArg, _err := convertNetworkRecordToXen(fmt.Sprintf("%s(%s)", _method, "args"), args)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _argsArg)
	if _err != nil {
		return
	}
	_retval, _err = convertNetworkRefToGo(_method + " -> ", _result.Value)
	return
}


var NetworkClass_GetByUUIDMockedCallback = func (sessionID SessionRef, uuid string) (_retval NetworkRef, _err error) {
	log.Println("Network.GetByUUID not mocked")
	_err = errors.New("Network.GetByUUID not mocked")
	return
}

func (_class NetworkClass) GetByUUIDMock(sessionID SessionRef, uuid string) (_retval NetworkRef, _err error) {
	return NetworkClass_GetByUUIDMockedCallback(sessionID, uuid)
}
// Get a reference to the network instance with the specified UUID.
func (_class NetworkClass) GetByUUID(sessionID SessionRef, uuid string) (_retval NetworkRef, _err error) {
	if (IsMock) {
		return _class.GetByUUIDMock(sessionID, uuid)
	}	
	_method := "network.get_by_uuid"
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
	_retval, _err = convertNetworkRefToGo(_method + " -> ", _result.Value)
	return
}


var NetworkClass_GetRecordMockedCallback = func (sessionID SessionRef, self NetworkRef) (_retval NetworkRecord, _err error) {
	log.Println("Network.GetRecord not mocked")
	_err = errors.New("Network.GetRecord not mocked")
	return
}

func (_class NetworkClass) GetRecordMock(sessionID SessionRef, self NetworkRef) (_retval NetworkRecord, _err error) {
	return NetworkClass_GetRecordMockedCallback(sessionID, self)
}
// Get a record containing the current state of the given network.
func (_class NetworkClass) GetRecord(sessionID SessionRef, self NetworkRef) (_retval NetworkRecord, _err error) {
	if (IsMock) {
		return _class.GetRecordMock(sessionID, self)
	}	
	_method := "network.get_record"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertNetworkRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertNetworkRecordToGo(_method + " -> ", _result.Value)
	return
}
