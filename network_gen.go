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


func NetworkClassGetAllRecordsMockDefault(sessionID SessionRef) (_retval map[NetworkRef]NetworkRecord, _err error) {
	log.Println("Network.GetAllRecords not mocked")
	_err = errors.New("Network.GetAllRecords not mocked")
	return
}

var NetworkClassGetAllRecordsMockedCallback = NetworkClassGetAllRecordsMockDefault

func (_class NetworkClass) GetAllRecordsMock(sessionID SessionRef) (_retval map[NetworkRef]NetworkRecord, _err error) {
	return NetworkClassGetAllRecordsMockedCallback(sessionID)
}
// Return a map of network references to network records for all networks known to the system.
func (_class NetworkClass) GetAllRecords(sessionID SessionRef) (_retval map[NetworkRef]NetworkRecord, _err error) {
	if IsMock {
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


func NetworkClassGetAllMockDefault(sessionID SessionRef) (_retval []NetworkRef, _err error) {
	log.Println("Network.GetAll not mocked")
	_err = errors.New("Network.GetAll not mocked")
	return
}

var NetworkClassGetAllMockedCallback = NetworkClassGetAllMockDefault

func (_class NetworkClass) GetAllMock(sessionID SessionRef) (_retval []NetworkRef, _err error) {
	return NetworkClassGetAllMockedCallback(sessionID)
}
// Return a list of all the networks known to the system.
func (_class NetworkClass) GetAll(sessionID SessionRef) (_retval []NetworkRef, _err error) {
	if IsMock {
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


func NetworkClassRemovePurposeMockDefault(sessionID SessionRef, self NetworkRef, value NetworkPurpose) (_err error) {
	log.Println("Network.RemovePurpose not mocked")
	_err = errors.New("Network.RemovePurpose not mocked")
	return
}

var NetworkClassRemovePurposeMockedCallback = NetworkClassRemovePurposeMockDefault

func (_class NetworkClass) RemovePurposeMock(sessionID SessionRef, self NetworkRef, value NetworkPurpose) (_err error) {
	return NetworkClassRemovePurposeMockedCallback(sessionID, self, value)
}
// Remove a purpose from a network (if present)
func (_class NetworkClass) RemovePurpose(sessionID SessionRef, self NetworkRef, value NetworkPurpose) (_err error) {
	if IsMock {
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


func NetworkClassAddPurposeMockDefault(sessionID SessionRef, self NetworkRef, value NetworkPurpose) (_err error) {
	log.Println("Network.AddPurpose not mocked")
	_err = errors.New("Network.AddPurpose not mocked")
	return
}

var NetworkClassAddPurposeMockedCallback = NetworkClassAddPurposeMockDefault

func (_class NetworkClass) AddPurposeMock(sessionID SessionRef, self NetworkRef, value NetworkPurpose) (_err error) {
	return NetworkClassAddPurposeMockedCallback(sessionID, self, value)
}
// Give a network a new purpose (if not present already)
//
// Errors:
//  NETWORK_INCOMPATIBLE_PURPOSES - You tried to add a purpose to a network but the new purpose is not compatible with an existing purpose of the network or other networks.
func (_class NetworkClass) AddPurpose(sessionID SessionRef, self NetworkRef, value NetworkPurpose) (_err error) {
	if IsMock {
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


func NetworkClassSetDefaultLockingModeMockDefault(sessionID SessionRef, network NetworkRef, value NetworkDefaultLockingMode) (_err error) {
	log.Println("Network.SetDefaultLockingMode not mocked")
	_err = errors.New("Network.SetDefaultLockingMode not mocked")
	return
}

var NetworkClassSetDefaultLockingModeMockedCallback = NetworkClassSetDefaultLockingModeMockDefault

func (_class NetworkClass) SetDefaultLockingModeMock(sessionID SessionRef, network NetworkRef, value NetworkDefaultLockingMode) (_err error) {
	return NetworkClassSetDefaultLockingModeMockedCallback(sessionID, network, value)
}
// Set the default locking mode for VIFs attached to this network
func (_class NetworkClass) SetDefaultLockingMode(sessionID SessionRef, network NetworkRef, value NetworkDefaultLockingMode) (_err error) {
	if IsMock {
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


func NetworkClassCreateNewBlobMockDefault(sessionID SessionRef, network NetworkRef, name string, mimeType string, public bool) (_retval BlobRef, _err error) {
	log.Println("Network.CreateNewBlob not mocked")
	_err = errors.New("Network.CreateNewBlob not mocked")
	return
}

var NetworkClassCreateNewBlobMockedCallback = NetworkClassCreateNewBlobMockDefault

func (_class NetworkClass) CreateNewBlobMock(sessionID SessionRef, network NetworkRef, name string, mimeType string, public bool) (_retval BlobRef, _err error) {
	return NetworkClassCreateNewBlobMockedCallback(sessionID, network, name, mimeType, public)
}
// Create a placeholder for a named binary blob of data that is associated with this pool
func (_class NetworkClass) CreateNewBlob(sessionID SessionRef, network NetworkRef, name string, mimeType string, public bool) (_retval BlobRef, _err error) {
	if IsMock {
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


func NetworkClassRemoveTagsMockDefault(sessionID SessionRef, self NetworkRef, value string) (_err error) {
	log.Println("Network.RemoveTags not mocked")
	_err = errors.New("Network.RemoveTags not mocked")
	return
}

var NetworkClassRemoveTagsMockedCallback = NetworkClassRemoveTagsMockDefault

func (_class NetworkClass) RemoveTagsMock(sessionID SessionRef, self NetworkRef, value string) (_err error) {
	return NetworkClassRemoveTagsMockedCallback(sessionID, self, value)
}
// Remove the given value from the tags field of the given network.  If the value is not in that Set, then do nothing.
func (_class NetworkClass) RemoveTags(sessionID SessionRef, self NetworkRef, value string) (_err error) {
	if IsMock {
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


func NetworkClassAddTagsMockDefault(sessionID SessionRef, self NetworkRef, value string) (_err error) {
	log.Println("Network.AddTags not mocked")
	_err = errors.New("Network.AddTags not mocked")
	return
}

var NetworkClassAddTagsMockedCallback = NetworkClassAddTagsMockDefault

func (_class NetworkClass) AddTagsMock(sessionID SessionRef, self NetworkRef, value string) (_err error) {
	return NetworkClassAddTagsMockedCallback(sessionID, self, value)
}
// Add the given value to the tags field of the given network.  If the value is already in that Set, then do nothing.
func (_class NetworkClass) AddTags(sessionID SessionRef, self NetworkRef, value string) (_err error) {
	if IsMock {
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


func NetworkClassSetTagsMockDefault(sessionID SessionRef, self NetworkRef, value []string) (_err error) {
	log.Println("Network.SetTags not mocked")
	_err = errors.New("Network.SetTags not mocked")
	return
}

var NetworkClassSetTagsMockedCallback = NetworkClassSetTagsMockDefault

func (_class NetworkClass) SetTagsMock(sessionID SessionRef, self NetworkRef, value []string) (_err error) {
	return NetworkClassSetTagsMockedCallback(sessionID, self, value)
}
// Set the tags field of the given network.
func (_class NetworkClass) SetTags(sessionID SessionRef, self NetworkRef, value []string) (_err error) {
	if IsMock {
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


func NetworkClassRemoveFromOtherConfigMockDefault(sessionID SessionRef, self NetworkRef, key string) (_err error) {
	log.Println("Network.RemoveFromOtherConfig not mocked")
	_err = errors.New("Network.RemoveFromOtherConfig not mocked")
	return
}

var NetworkClassRemoveFromOtherConfigMockedCallback = NetworkClassRemoveFromOtherConfigMockDefault

func (_class NetworkClass) RemoveFromOtherConfigMock(sessionID SessionRef, self NetworkRef, key string) (_err error) {
	return NetworkClassRemoveFromOtherConfigMockedCallback(sessionID, self, key)
}
// Remove the given key and its corresponding value from the other_config field of the given network.  If the key is not in that Map, then do nothing.
func (_class NetworkClass) RemoveFromOtherConfig(sessionID SessionRef, self NetworkRef, key string) (_err error) {
	if IsMock {
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


func NetworkClassAddToOtherConfigMockDefault(sessionID SessionRef, self NetworkRef, key string, value string) (_err error) {
	log.Println("Network.AddToOtherConfig not mocked")
	_err = errors.New("Network.AddToOtherConfig not mocked")
	return
}

var NetworkClassAddToOtherConfigMockedCallback = NetworkClassAddToOtherConfigMockDefault

func (_class NetworkClass) AddToOtherConfigMock(sessionID SessionRef, self NetworkRef, key string, value string) (_err error) {
	return NetworkClassAddToOtherConfigMockedCallback(sessionID, self, key, value)
}
// Add the given key-value pair to the other_config field of the given network.
func (_class NetworkClass) AddToOtherConfig(sessionID SessionRef, self NetworkRef, key string, value string) (_err error) {
	if IsMock {
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


func NetworkClassSetOtherConfigMockDefault(sessionID SessionRef, self NetworkRef, value map[string]string) (_err error) {
	log.Println("Network.SetOtherConfig not mocked")
	_err = errors.New("Network.SetOtherConfig not mocked")
	return
}

var NetworkClassSetOtherConfigMockedCallback = NetworkClassSetOtherConfigMockDefault

func (_class NetworkClass) SetOtherConfigMock(sessionID SessionRef, self NetworkRef, value map[string]string) (_err error) {
	return NetworkClassSetOtherConfigMockedCallback(sessionID, self, value)
}
// Set the other_config field of the given network.
func (_class NetworkClass) SetOtherConfig(sessionID SessionRef, self NetworkRef, value map[string]string) (_err error) {
	if IsMock {
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


func NetworkClassSetMTUMockDefault(sessionID SessionRef, self NetworkRef, value int) (_err error) {
	log.Println("Network.SetMTU not mocked")
	_err = errors.New("Network.SetMTU not mocked")
	return
}

var NetworkClassSetMTUMockedCallback = NetworkClassSetMTUMockDefault

func (_class NetworkClass) SetMTUMock(sessionID SessionRef, self NetworkRef, value int) (_err error) {
	return NetworkClassSetMTUMockedCallback(sessionID, self, value)
}
// Set the MTU field of the given network.
func (_class NetworkClass) SetMTU(sessionID SessionRef, self NetworkRef, value int) (_err error) {
	if IsMock {
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


func NetworkClassSetNameDescriptionMockDefault(sessionID SessionRef, self NetworkRef, value string) (_err error) {
	log.Println("Network.SetNameDescription not mocked")
	_err = errors.New("Network.SetNameDescription not mocked")
	return
}

var NetworkClassSetNameDescriptionMockedCallback = NetworkClassSetNameDescriptionMockDefault

func (_class NetworkClass) SetNameDescriptionMock(sessionID SessionRef, self NetworkRef, value string) (_err error) {
	return NetworkClassSetNameDescriptionMockedCallback(sessionID, self, value)
}
// Set the name/description field of the given network.
func (_class NetworkClass) SetNameDescription(sessionID SessionRef, self NetworkRef, value string) (_err error) {
	if IsMock {
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


func NetworkClassSetNameLabelMockDefault(sessionID SessionRef, self NetworkRef, value string) (_err error) {
	log.Println("Network.SetNameLabel not mocked")
	_err = errors.New("Network.SetNameLabel not mocked")
	return
}

var NetworkClassSetNameLabelMockedCallback = NetworkClassSetNameLabelMockDefault

func (_class NetworkClass) SetNameLabelMock(sessionID SessionRef, self NetworkRef, value string) (_err error) {
	return NetworkClassSetNameLabelMockedCallback(sessionID, self, value)
}
// Set the name/label field of the given network.
func (_class NetworkClass) SetNameLabel(sessionID SessionRef, self NetworkRef, value string) (_err error) {
	if IsMock {
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


func NetworkClassGetPurposeMockDefault(sessionID SessionRef, self NetworkRef) (_retval []NetworkPurpose, _err error) {
	log.Println("Network.GetPurpose not mocked")
	_err = errors.New("Network.GetPurpose not mocked")
	return
}

var NetworkClassGetPurposeMockedCallback = NetworkClassGetPurposeMockDefault

func (_class NetworkClass) GetPurposeMock(sessionID SessionRef, self NetworkRef) (_retval []NetworkPurpose, _err error) {
	return NetworkClassGetPurposeMockedCallback(sessionID, self)
}
// Get the purpose field of the given network.
func (_class NetworkClass) GetPurpose(sessionID SessionRef, self NetworkRef) (_retval []NetworkPurpose, _err error) {
	if IsMock {
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


func NetworkClassGetAssignedIpsMockDefault(sessionID SessionRef, self NetworkRef) (_retval map[VIFRef]string, _err error) {
	log.Println("Network.GetAssignedIps not mocked")
	_err = errors.New("Network.GetAssignedIps not mocked")
	return
}

var NetworkClassGetAssignedIpsMockedCallback = NetworkClassGetAssignedIpsMockDefault

func (_class NetworkClass) GetAssignedIpsMock(sessionID SessionRef, self NetworkRef) (_retval map[VIFRef]string, _err error) {
	return NetworkClassGetAssignedIpsMockedCallback(sessionID, self)
}
// Get the assigned_ips field of the given network.
func (_class NetworkClass) GetAssignedIps(sessionID SessionRef, self NetworkRef) (_retval map[VIFRef]string, _err error) {
	if IsMock {
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


func NetworkClassGetDefaultLockingModeMockDefault(sessionID SessionRef, self NetworkRef) (_retval NetworkDefaultLockingMode, _err error) {
	log.Println("Network.GetDefaultLockingMode not mocked")
	_err = errors.New("Network.GetDefaultLockingMode not mocked")
	return
}

var NetworkClassGetDefaultLockingModeMockedCallback = NetworkClassGetDefaultLockingModeMockDefault

func (_class NetworkClass) GetDefaultLockingModeMock(sessionID SessionRef, self NetworkRef) (_retval NetworkDefaultLockingMode, _err error) {
	return NetworkClassGetDefaultLockingModeMockedCallback(sessionID, self)
}
// Get the default_locking_mode field of the given network.
func (_class NetworkClass) GetDefaultLockingMode(sessionID SessionRef, self NetworkRef) (_retval NetworkDefaultLockingMode, _err error) {
	if IsMock {
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


func NetworkClassGetTagsMockDefault(sessionID SessionRef, self NetworkRef) (_retval []string, _err error) {
	log.Println("Network.GetTags not mocked")
	_err = errors.New("Network.GetTags not mocked")
	return
}

var NetworkClassGetTagsMockedCallback = NetworkClassGetTagsMockDefault

func (_class NetworkClass) GetTagsMock(sessionID SessionRef, self NetworkRef) (_retval []string, _err error) {
	return NetworkClassGetTagsMockedCallback(sessionID, self)
}
// Get the tags field of the given network.
func (_class NetworkClass) GetTags(sessionID SessionRef, self NetworkRef) (_retval []string, _err error) {
	if IsMock {
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


func NetworkClassGetBlobsMockDefault(sessionID SessionRef, self NetworkRef) (_retval map[string]BlobRef, _err error) {
	log.Println("Network.GetBlobs not mocked")
	_err = errors.New("Network.GetBlobs not mocked")
	return
}

var NetworkClassGetBlobsMockedCallback = NetworkClassGetBlobsMockDefault

func (_class NetworkClass) GetBlobsMock(sessionID SessionRef, self NetworkRef) (_retval map[string]BlobRef, _err error) {
	return NetworkClassGetBlobsMockedCallback(sessionID, self)
}
// Get the blobs field of the given network.
func (_class NetworkClass) GetBlobs(sessionID SessionRef, self NetworkRef) (_retval map[string]BlobRef, _err error) {
	if IsMock {
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


func NetworkClassGetManagedMockDefault(sessionID SessionRef, self NetworkRef) (_retval bool, _err error) {
	log.Println("Network.GetManaged not mocked")
	_err = errors.New("Network.GetManaged not mocked")
	return
}

var NetworkClassGetManagedMockedCallback = NetworkClassGetManagedMockDefault

func (_class NetworkClass) GetManagedMock(sessionID SessionRef, self NetworkRef) (_retval bool, _err error) {
	return NetworkClassGetManagedMockedCallback(sessionID, self)
}
// Get the managed field of the given network.
func (_class NetworkClass) GetManaged(sessionID SessionRef, self NetworkRef) (_retval bool, _err error) {
	if IsMock {
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


func NetworkClassGetBridgeMockDefault(sessionID SessionRef, self NetworkRef) (_retval string, _err error) {
	log.Println("Network.GetBridge not mocked")
	_err = errors.New("Network.GetBridge not mocked")
	return
}

var NetworkClassGetBridgeMockedCallback = NetworkClassGetBridgeMockDefault

func (_class NetworkClass) GetBridgeMock(sessionID SessionRef, self NetworkRef) (_retval string, _err error) {
	return NetworkClassGetBridgeMockedCallback(sessionID, self)
}
// Get the bridge field of the given network.
func (_class NetworkClass) GetBridge(sessionID SessionRef, self NetworkRef) (_retval string, _err error) {
	if IsMock {
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


func NetworkClassGetOtherConfigMockDefault(sessionID SessionRef, self NetworkRef) (_retval map[string]string, _err error) {
	log.Println("Network.GetOtherConfig not mocked")
	_err = errors.New("Network.GetOtherConfig not mocked")
	return
}

var NetworkClassGetOtherConfigMockedCallback = NetworkClassGetOtherConfigMockDefault

func (_class NetworkClass) GetOtherConfigMock(sessionID SessionRef, self NetworkRef) (_retval map[string]string, _err error) {
	return NetworkClassGetOtherConfigMockedCallback(sessionID, self)
}
// Get the other_config field of the given network.
func (_class NetworkClass) GetOtherConfig(sessionID SessionRef, self NetworkRef) (_retval map[string]string, _err error) {
	if IsMock {
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


func NetworkClassGetMTUMockDefault(sessionID SessionRef, self NetworkRef) (_retval int, _err error) {
	log.Println("Network.GetMTU not mocked")
	_err = errors.New("Network.GetMTU not mocked")
	return
}

var NetworkClassGetMTUMockedCallback = NetworkClassGetMTUMockDefault

func (_class NetworkClass) GetMTUMock(sessionID SessionRef, self NetworkRef) (_retval int, _err error) {
	return NetworkClassGetMTUMockedCallback(sessionID, self)
}
// Get the MTU field of the given network.
func (_class NetworkClass) GetMTU(sessionID SessionRef, self NetworkRef) (_retval int, _err error) {
	if IsMock {
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


func NetworkClassGetPIFsMockDefault(sessionID SessionRef, self NetworkRef) (_retval []PIFRef, _err error) {
	log.Println("Network.GetPIFs not mocked")
	_err = errors.New("Network.GetPIFs not mocked")
	return
}

var NetworkClassGetPIFsMockedCallback = NetworkClassGetPIFsMockDefault

func (_class NetworkClass) GetPIFsMock(sessionID SessionRef, self NetworkRef) (_retval []PIFRef, _err error) {
	return NetworkClassGetPIFsMockedCallback(sessionID, self)
}
// Get the PIFs field of the given network.
func (_class NetworkClass) GetPIFs(sessionID SessionRef, self NetworkRef) (_retval []PIFRef, _err error) {
	if IsMock {
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


func NetworkClassGetVIFsMockDefault(sessionID SessionRef, self NetworkRef) (_retval []VIFRef, _err error) {
	log.Println("Network.GetVIFs not mocked")
	_err = errors.New("Network.GetVIFs not mocked")
	return
}

var NetworkClassGetVIFsMockedCallback = NetworkClassGetVIFsMockDefault

func (_class NetworkClass) GetVIFsMock(sessionID SessionRef, self NetworkRef) (_retval []VIFRef, _err error) {
	return NetworkClassGetVIFsMockedCallback(sessionID, self)
}
// Get the VIFs field of the given network.
func (_class NetworkClass) GetVIFs(sessionID SessionRef, self NetworkRef) (_retval []VIFRef, _err error) {
	if IsMock {
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


func NetworkClassGetCurrentOperationsMockDefault(sessionID SessionRef, self NetworkRef) (_retval map[string]NetworkOperations, _err error) {
	log.Println("Network.GetCurrentOperations not mocked")
	_err = errors.New("Network.GetCurrentOperations not mocked")
	return
}

var NetworkClassGetCurrentOperationsMockedCallback = NetworkClassGetCurrentOperationsMockDefault

func (_class NetworkClass) GetCurrentOperationsMock(sessionID SessionRef, self NetworkRef) (_retval map[string]NetworkOperations, _err error) {
	return NetworkClassGetCurrentOperationsMockedCallback(sessionID, self)
}
// Get the current_operations field of the given network.
func (_class NetworkClass) GetCurrentOperations(sessionID SessionRef, self NetworkRef) (_retval map[string]NetworkOperations, _err error) {
	if IsMock {
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


func NetworkClassGetAllowedOperationsMockDefault(sessionID SessionRef, self NetworkRef) (_retval []NetworkOperations, _err error) {
	log.Println("Network.GetAllowedOperations not mocked")
	_err = errors.New("Network.GetAllowedOperations not mocked")
	return
}

var NetworkClassGetAllowedOperationsMockedCallback = NetworkClassGetAllowedOperationsMockDefault

func (_class NetworkClass) GetAllowedOperationsMock(sessionID SessionRef, self NetworkRef) (_retval []NetworkOperations, _err error) {
	return NetworkClassGetAllowedOperationsMockedCallback(sessionID, self)
}
// Get the allowed_operations field of the given network.
func (_class NetworkClass) GetAllowedOperations(sessionID SessionRef, self NetworkRef) (_retval []NetworkOperations, _err error) {
	if IsMock {
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


func NetworkClassGetNameDescriptionMockDefault(sessionID SessionRef, self NetworkRef) (_retval string, _err error) {
	log.Println("Network.GetNameDescription not mocked")
	_err = errors.New("Network.GetNameDescription not mocked")
	return
}

var NetworkClassGetNameDescriptionMockedCallback = NetworkClassGetNameDescriptionMockDefault

func (_class NetworkClass) GetNameDescriptionMock(sessionID SessionRef, self NetworkRef) (_retval string, _err error) {
	return NetworkClassGetNameDescriptionMockedCallback(sessionID, self)
}
// Get the name/description field of the given network.
func (_class NetworkClass) GetNameDescription(sessionID SessionRef, self NetworkRef) (_retval string, _err error) {
	if IsMock {
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


func NetworkClassGetNameLabelMockDefault(sessionID SessionRef, self NetworkRef) (_retval string, _err error) {
	log.Println("Network.GetNameLabel not mocked")
	_err = errors.New("Network.GetNameLabel not mocked")
	return
}

var NetworkClassGetNameLabelMockedCallback = NetworkClassGetNameLabelMockDefault

func (_class NetworkClass) GetNameLabelMock(sessionID SessionRef, self NetworkRef) (_retval string, _err error) {
	return NetworkClassGetNameLabelMockedCallback(sessionID, self)
}
// Get the name/label field of the given network.
func (_class NetworkClass) GetNameLabel(sessionID SessionRef, self NetworkRef) (_retval string, _err error) {
	if IsMock {
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


func NetworkClassGetUUIDMockDefault(sessionID SessionRef, self NetworkRef) (_retval string, _err error) {
	log.Println("Network.GetUUID not mocked")
	_err = errors.New("Network.GetUUID not mocked")
	return
}

var NetworkClassGetUUIDMockedCallback = NetworkClassGetUUIDMockDefault

func (_class NetworkClass) GetUUIDMock(sessionID SessionRef, self NetworkRef) (_retval string, _err error) {
	return NetworkClassGetUUIDMockedCallback(sessionID, self)
}
// Get the uuid field of the given network.
func (_class NetworkClass) GetUUID(sessionID SessionRef, self NetworkRef) (_retval string, _err error) {
	if IsMock {
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


func NetworkClassGetByNameLabelMockDefault(sessionID SessionRef, label string) (_retval []NetworkRef, _err error) {
	log.Println("Network.GetByNameLabel not mocked")
	_err = errors.New("Network.GetByNameLabel not mocked")
	return
}

var NetworkClassGetByNameLabelMockedCallback = NetworkClassGetByNameLabelMockDefault

func (_class NetworkClass) GetByNameLabelMock(sessionID SessionRef, label string) (_retval []NetworkRef, _err error) {
	return NetworkClassGetByNameLabelMockedCallback(sessionID, label)
}
// Get all the network instances with the given label.
func (_class NetworkClass) GetByNameLabel(sessionID SessionRef, label string) (_retval []NetworkRef, _err error) {
	if IsMock {
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


func NetworkClassDestroyMockDefault(sessionID SessionRef, self NetworkRef) (_err error) {
	log.Println("Network.Destroy not mocked")
	_err = errors.New("Network.Destroy not mocked")
	return
}

var NetworkClassDestroyMockedCallback = NetworkClassDestroyMockDefault

func (_class NetworkClass) DestroyMock(sessionID SessionRef, self NetworkRef) (_err error) {
	return NetworkClassDestroyMockedCallback(sessionID, self)
}
// Destroy the specified network instance.
func (_class NetworkClass) Destroy(sessionID SessionRef, self NetworkRef) (_err error) {
	if IsMock {
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


func NetworkClassCreateMockDefault(sessionID SessionRef, args NetworkRecord) (_retval NetworkRef, _err error) {
	log.Println("Network.Create not mocked")
	_err = errors.New("Network.Create not mocked")
	return
}

var NetworkClassCreateMockedCallback = NetworkClassCreateMockDefault

func (_class NetworkClass) CreateMock(sessionID SessionRef, args NetworkRecord) (_retval NetworkRef, _err error) {
	return NetworkClassCreateMockedCallback(sessionID, args)
}
// Create a new network instance, and return its handle.
// The constructor args are: name_label, name_description, MTU, other_config*, bridge, managed, tags (* = non-optional).
func (_class NetworkClass) Create(sessionID SessionRef, args NetworkRecord) (_retval NetworkRef, _err error) {
	if IsMock {
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


func NetworkClassGetByUUIDMockDefault(sessionID SessionRef, uuid string) (_retval NetworkRef, _err error) {
	log.Println("Network.GetByUUID not mocked")
	_err = errors.New("Network.GetByUUID not mocked")
	return
}

var NetworkClassGetByUUIDMockedCallback = NetworkClassGetByUUIDMockDefault

func (_class NetworkClass) GetByUUIDMock(sessionID SessionRef, uuid string) (_retval NetworkRef, _err error) {
	return NetworkClassGetByUUIDMockedCallback(sessionID, uuid)
}
// Get a reference to the network instance with the specified UUID.
func (_class NetworkClass) GetByUUID(sessionID SessionRef, uuid string) (_retval NetworkRef, _err error) {
	if IsMock {
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


func NetworkClassGetRecordMockDefault(sessionID SessionRef, self NetworkRef) (_retval NetworkRecord, _err error) {
	log.Println("Network.GetRecord not mocked")
	_err = errors.New("Network.GetRecord not mocked")
	return
}

var NetworkClassGetRecordMockedCallback = NetworkClassGetRecordMockDefault

func (_class NetworkClass) GetRecordMock(sessionID SessionRef, self NetworkRef) (_retval NetworkRecord, _err error) {
	return NetworkClassGetRecordMockedCallback(sessionID, self)
}
// Get a record containing the current state of the given network.
func (_class NetworkClass) GetRecord(sessionID SessionRef, self NetworkRef) (_retval NetworkRecord, _err error) {
	if IsMock {
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
