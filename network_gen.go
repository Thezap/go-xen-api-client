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

func (_class NetworkClass) GetAllRecords__mock(sessionID SessionRef) (_retval map[NetworkRef]NetworkRecord, _err error) {
	log.Println("Network.GetAllRecords not mocked")
	_err = errors.New("Network.GetAllRecords not mocked")
	return
}
// Return a map of network references to network records for all networks known to the system.
func (_class NetworkClass) GetAllRecords(sessionID SessionRef) (_retval map[NetworkRef]NetworkRecord, _err error) {
	if (IsMock) {
		return _class.GetAllRecords__mock(sessionID)
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

func (_class NetworkClass) GetAll__mock(sessionID SessionRef) (_retval []NetworkRef, _err error) {
	log.Println("Network.GetAll not mocked")
	_err = errors.New("Network.GetAll not mocked")
	return
}
// Return a list of all the networks known to the system.
func (_class NetworkClass) GetAll(sessionID SessionRef) (_retval []NetworkRef, _err error) {
	if (IsMock) {
		return _class.GetAll__mock(sessionID)
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

func (_class NetworkClass) RemovePurpose__mock(sessionID SessionRef, self NetworkRef, value NetworkPurpose) (_err error) {
	log.Println("Network.RemovePurpose not mocked")
	_err = errors.New("Network.RemovePurpose not mocked")
	return
}
// Remove a purpose from a network (if present)
func (_class NetworkClass) RemovePurpose(sessionID SessionRef, self NetworkRef, value NetworkPurpose) (_err error) {
	if (IsMock) {
		return _class.RemovePurpose__mock(sessionID, self, value)
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

func (_class NetworkClass) AddPurpose__mock(sessionID SessionRef, self NetworkRef, value NetworkPurpose) (_err error) {
	log.Println("Network.AddPurpose not mocked")
	_err = errors.New("Network.AddPurpose not mocked")
	return
}
// Give a network a new purpose (if not present already)
//
// Errors:
//  NETWORK_INCOMPATIBLE_PURPOSES - You tried to add a purpose to a network but the new purpose is not compatible with an existing purpose of the network or other networks.
func (_class NetworkClass) AddPurpose(sessionID SessionRef, self NetworkRef, value NetworkPurpose) (_err error) {
	if (IsMock) {
		return _class.AddPurpose__mock(sessionID, self, value)
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

func (_class NetworkClass) SetDefaultLockingMode__mock(sessionID SessionRef, network NetworkRef, value NetworkDefaultLockingMode) (_err error) {
	log.Println("Network.SetDefaultLockingMode not mocked")
	_err = errors.New("Network.SetDefaultLockingMode not mocked")
	return
}
// Set the default locking mode for VIFs attached to this network
func (_class NetworkClass) SetDefaultLockingMode(sessionID SessionRef, network NetworkRef, value NetworkDefaultLockingMode) (_err error) {
	if (IsMock) {
		return _class.SetDefaultLockingMode__mock(sessionID, network, value)
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

func (_class NetworkClass) CreateNewBlob__mock(sessionID SessionRef, network NetworkRef, name string, mimeType string, public bool) (_retval BlobRef, _err error) {
	log.Println("Network.CreateNewBlob not mocked")
	_err = errors.New("Network.CreateNewBlob not mocked")
	return
}
// Create a placeholder for a named binary blob of data that is associated with this pool
func (_class NetworkClass) CreateNewBlob(sessionID SessionRef, network NetworkRef, name string, mimeType string, public bool) (_retval BlobRef, _err error) {
	if (IsMock) {
		return _class.CreateNewBlob__mock(sessionID, network, name, mimeType, public)
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

func (_class NetworkClass) RemoveTags__mock(sessionID SessionRef, self NetworkRef, value string) (_err error) {
	log.Println("Network.RemoveTags not mocked")
	_err = errors.New("Network.RemoveTags not mocked")
	return
}
// Remove the given value from the tags field of the given network.  If the value is not in that Set, then do nothing.
func (_class NetworkClass) RemoveTags(sessionID SessionRef, self NetworkRef, value string) (_err error) {
	if (IsMock) {
		return _class.RemoveTags__mock(sessionID, self, value)
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

func (_class NetworkClass) AddTags__mock(sessionID SessionRef, self NetworkRef, value string) (_err error) {
	log.Println("Network.AddTags not mocked")
	_err = errors.New("Network.AddTags not mocked")
	return
}
// Add the given value to the tags field of the given network.  If the value is already in that Set, then do nothing.
func (_class NetworkClass) AddTags(sessionID SessionRef, self NetworkRef, value string) (_err error) {
	if (IsMock) {
		return _class.AddTags__mock(sessionID, self, value)
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

func (_class NetworkClass) SetTags__mock(sessionID SessionRef, self NetworkRef, value []string) (_err error) {
	log.Println("Network.SetTags not mocked")
	_err = errors.New("Network.SetTags not mocked")
	return
}
// Set the tags field of the given network.
func (_class NetworkClass) SetTags(sessionID SessionRef, self NetworkRef, value []string) (_err error) {
	if (IsMock) {
		return _class.SetTags__mock(sessionID, self, value)
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

func (_class NetworkClass) RemoveFromOtherConfig__mock(sessionID SessionRef, self NetworkRef, key string) (_err error) {
	log.Println("Network.RemoveFromOtherConfig not mocked")
	_err = errors.New("Network.RemoveFromOtherConfig not mocked")
	return
}
// Remove the given key and its corresponding value from the other_config field of the given network.  If the key is not in that Map, then do nothing.
func (_class NetworkClass) RemoveFromOtherConfig(sessionID SessionRef, self NetworkRef, key string) (_err error) {
	if (IsMock) {
		return _class.RemoveFromOtherConfig__mock(sessionID, self, key)
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

func (_class NetworkClass) AddToOtherConfig__mock(sessionID SessionRef, self NetworkRef, key string, value string) (_err error) {
	log.Println("Network.AddToOtherConfig not mocked")
	_err = errors.New("Network.AddToOtherConfig not mocked")
	return
}
// Add the given key-value pair to the other_config field of the given network.
func (_class NetworkClass) AddToOtherConfig(sessionID SessionRef, self NetworkRef, key string, value string) (_err error) {
	if (IsMock) {
		return _class.AddToOtherConfig__mock(sessionID, self, key, value)
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

func (_class NetworkClass) SetOtherConfig__mock(sessionID SessionRef, self NetworkRef, value map[string]string) (_err error) {
	log.Println("Network.SetOtherConfig not mocked")
	_err = errors.New("Network.SetOtherConfig not mocked")
	return
}
// Set the other_config field of the given network.
func (_class NetworkClass) SetOtherConfig(sessionID SessionRef, self NetworkRef, value map[string]string) (_err error) {
	if (IsMock) {
		return _class.SetOtherConfig__mock(sessionID, self, value)
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

func (_class NetworkClass) SetMTU__mock(sessionID SessionRef, self NetworkRef, value int) (_err error) {
	log.Println("Network.SetMTU not mocked")
	_err = errors.New("Network.SetMTU not mocked")
	return
}
// Set the MTU field of the given network.
func (_class NetworkClass) SetMTU(sessionID SessionRef, self NetworkRef, value int) (_err error) {
	if (IsMock) {
		return _class.SetMTU__mock(sessionID, self, value)
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

func (_class NetworkClass) SetNameDescription__mock(sessionID SessionRef, self NetworkRef, value string) (_err error) {
	log.Println("Network.SetNameDescription not mocked")
	_err = errors.New("Network.SetNameDescription not mocked")
	return
}
// Set the name/description field of the given network.
func (_class NetworkClass) SetNameDescription(sessionID SessionRef, self NetworkRef, value string) (_err error) {
	if (IsMock) {
		return _class.SetNameDescription__mock(sessionID, self, value)
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

func (_class NetworkClass) SetNameLabel__mock(sessionID SessionRef, self NetworkRef, value string) (_err error) {
	log.Println("Network.SetNameLabel not mocked")
	_err = errors.New("Network.SetNameLabel not mocked")
	return
}
// Set the name/label field of the given network.
func (_class NetworkClass) SetNameLabel(sessionID SessionRef, self NetworkRef, value string) (_err error) {
	if (IsMock) {
		return _class.SetNameLabel__mock(sessionID, self, value)
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

func (_class NetworkClass) GetPurpose__mock(sessionID SessionRef, self NetworkRef) (_retval []NetworkPurpose, _err error) {
	log.Println("Network.GetPurpose not mocked")
	_err = errors.New("Network.GetPurpose not mocked")
	return
}
// Get the purpose field of the given network.
func (_class NetworkClass) GetPurpose(sessionID SessionRef, self NetworkRef) (_retval []NetworkPurpose, _err error) {
	if (IsMock) {
		return _class.GetPurpose__mock(sessionID, self)
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

func (_class NetworkClass) GetAssignedIps__mock(sessionID SessionRef, self NetworkRef) (_retval map[VIFRef]string, _err error) {
	log.Println("Network.GetAssignedIps not mocked")
	_err = errors.New("Network.GetAssignedIps not mocked")
	return
}
// Get the assigned_ips field of the given network.
func (_class NetworkClass) GetAssignedIps(sessionID SessionRef, self NetworkRef) (_retval map[VIFRef]string, _err error) {
	if (IsMock) {
		return _class.GetAssignedIps__mock(sessionID, self)
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

func (_class NetworkClass) GetDefaultLockingMode__mock(sessionID SessionRef, self NetworkRef) (_retval NetworkDefaultLockingMode, _err error) {
	log.Println("Network.GetDefaultLockingMode not mocked")
	_err = errors.New("Network.GetDefaultLockingMode not mocked")
	return
}
// Get the default_locking_mode field of the given network.
func (_class NetworkClass) GetDefaultLockingMode(sessionID SessionRef, self NetworkRef) (_retval NetworkDefaultLockingMode, _err error) {
	if (IsMock) {
		return _class.GetDefaultLockingMode__mock(sessionID, self)
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

func (_class NetworkClass) GetTags__mock(sessionID SessionRef, self NetworkRef) (_retval []string, _err error) {
	log.Println("Network.GetTags not mocked")
	_err = errors.New("Network.GetTags not mocked")
	return
}
// Get the tags field of the given network.
func (_class NetworkClass) GetTags(sessionID SessionRef, self NetworkRef) (_retval []string, _err error) {
	if (IsMock) {
		return _class.GetTags__mock(sessionID, self)
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

func (_class NetworkClass) GetBlobs__mock(sessionID SessionRef, self NetworkRef) (_retval map[string]BlobRef, _err error) {
	log.Println("Network.GetBlobs not mocked")
	_err = errors.New("Network.GetBlobs not mocked")
	return
}
// Get the blobs field of the given network.
func (_class NetworkClass) GetBlobs(sessionID SessionRef, self NetworkRef) (_retval map[string]BlobRef, _err error) {
	if (IsMock) {
		return _class.GetBlobs__mock(sessionID, self)
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

func (_class NetworkClass) GetManaged__mock(sessionID SessionRef, self NetworkRef) (_retval bool, _err error) {
	log.Println("Network.GetManaged not mocked")
	_err = errors.New("Network.GetManaged not mocked")
	return
}
// Get the managed field of the given network.
func (_class NetworkClass) GetManaged(sessionID SessionRef, self NetworkRef) (_retval bool, _err error) {
	if (IsMock) {
		return _class.GetManaged__mock(sessionID, self)
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

func (_class NetworkClass) GetBridge__mock(sessionID SessionRef, self NetworkRef) (_retval string, _err error) {
	log.Println("Network.GetBridge not mocked")
	_err = errors.New("Network.GetBridge not mocked")
	return
}
// Get the bridge field of the given network.
func (_class NetworkClass) GetBridge(sessionID SessionRef, self NetworkRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetBridge__mock(sessionID, self)
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

func (_class NetworkClass) GetOtherConfig__mock(sessionID SessionRef, self NetworkRef) (_retval map[string]string, _err error) {
	log.Println("Network.GetOtherConfig not mocked")
	_err = errors.New("Network.GetOtherConfig not mocked")
	return
}
// Get the other_config field of the given network.
func (_class NetworkClass) GetOtherConfig(sessionID SessionRef, self NetworkRef) (_retval map[string]string, _err error) {
	if (IsMock) {
		return _class.GetOtherConfig__mock(sessionID, self)
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

func (_class NetworkClass) GetMTU__mock(sessionID SessionRef, self NetworkRef) (_retval int, _err error) {
	log.Println("Network.GetMTU not mocked")
	_err = errors.New("Network.GetMTU not mocked")
	return
}
// Get the MTU field of the given network.
func (_class NetworkClass) GetMTU(sessionID SessionRef, self NetworkRef) (_retval int, _err error) {
	if (IsMock) {
		return _class.GetMTU__mock(sessionID, self)
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

func (_class NetworkClass) GetPIFs__mock(sessionID SessionRef, self NetworkRef) (_retval []PIFRef, _err error) {
	log.Println("Network.GetPIFs not mocked")
	_err = errors.New("Network.GetPIFs not mocked")
	return
}
// Get the PIFs field of the given network.
func (_class NetworkClass) GetPIFs(sessionID SessionRef, self NetworkRef) (_retval []PIFRef, _err error) {
	if (IsMock) {
		return _class.GetPIFs__mock(sessionID, self)
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

func (_class NetworkClass) GetVIFs__mock(sessionID SessionRef, self NetworkRef) (_retval []VIFRef, _err error) {
	log.Println("Network.GetVIFs not mocked")
	_err = errors.New("Network.GetVIFs not mocked")
	return
}
// Get the VIFs field of the given network.
func (_class NetworkClass) GetVIFs(sessionID SessionRef, self NetworkRef) (_retval []VIFRef, _err error) {
	if (IsMock) {
		return _class.GetVIFs__mock(sessionID, self)
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

func (_class NetworkClass) GetCurrentOperations__mock(sessionID SessionRef, self NetworkRef) (_retval map[string]NetworkOperations, _err error) {
	log.Println("Network.GetCurrentOperations not mocked")
	_err = errors.New("Network.GetCurrentOperations not mocked")
	return
}
// Get the current_operations field of the given network.
func (_class NetworkClass) GetCurrentOperations(sessionID SessionRef, self NetworkRef) (_retval map[string]NetworkOperations, _err error) {
	if (IsMock) {
		return _class.GetCurrentOperations__mock(sessionID, self)
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

func (_class NetworkClass) GetAllowedOperations__mock(sessionID SessionRef, self NetworkRef) (_retval []NetworkOperations, _err error) {
	log.Println("Network.GetAllowedOperations not mocked")
	_err = errors.New("Network.GetAllowedOperations not mocked")
	return
}
// Get the allowed_operations field of the given network.
func (_class NetworkClass) GetAllowedOperations(sessionID SessionRef, self NetworkRef) (_retval []NetworkOperations, _err error) {
	if (IsMock) {
		return _class.GetAllowedOperations__mock(sessionID, self)
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

func (_class NetworkClass) GetNameDescription__mock(sessionID SessionRef, self NetworkRef) (_retval string, _err error) {
	log.Println("Network.GetNameDescription not mocked")
	_err = errors.New("Network.GetNameDescription not mocked")
	return
}
// Get the name/description field of the given network.
func (_class NetworkClass) GetNameDescription(sessionID SessionRef, self NetworkRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetNameDescription__mock(sessionID, self)
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

func (_class NetworkClass) GetNameLabel__mock(sessionID SessionRef, self NetworkRef) (_retval string, _err error) {
	log.Println("Network.GetNameLabel not mocked")
	_err = errors.New("Network.GetNameLabel not mocked")
	return
}
// Get the name/label field of the given network.
func (_class NetworkClass) GetNameLabel(sessionID SessionRef, self NetworkRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetNameLabel__mock(sessionID, self)
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

func (_class NetworkClass) GetUUID__mock(sessionID SessionRef, self NetworkRef) (_retval string, _err error) {
	log.Println("Network.GetUUID not mocked")
	_err = errors.New("Network.GetUUID not mocked")
	return
}
// Get the uuid field of the given network.
func (_class NetworkClass) GetUUID(sessionID SessionRef, self NetworkRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetUUID__mock(sessionID, self)
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

func (_class NetworkClass) GetByNameLabel__mock(sessionID SessionRef, label string) (_retval []NetworkRef, _err error) {
	log.Println("Network.GetByNameLabel not mocked")
	_err = errors.New("Network.GetByNameLabel not mocked")
	return
}
// Get all the network instances with the given label.
func (_class NetworkClass) GetByNameLabel(sessionID SessionRef, label string) (_retval []NetworkRef, _err error) {
	if (IsMock) {
		return _class.GetByNameLabel__mock(sessionID, label)
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

func (_class NetworkClass) Destroy__mock(sessionID SessionRef, self NetworkRef) (_err error) {
	log.Println("Network.Destroy not mocked")
	_err = errors.New("Network.Destroy not mocked")
	return
}
// Destroy the specified network instance.
func (_class NetworkClass) Destroy(sessionID SessionRef, self NetworkRef) (_err error) {
	if (IsMock) {
		return _class.Destroy__mock(sessionID, self)
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

func (_class NetworkClass) Create__mock(sessionID SessionRef, args NetworkRecord) (_retval NetworkRef, _err error) {
	log.Println("Network.Create not mocked")
	_err = errors.New("Network.Create not mocked")
	return
}
// Create a new network instance, and return its handle.
// The constructor args are: name_label, name_description, MTU, other_config*, bridge, managed, tags (* = non-optional).
func (_class NetworkClass) Create(sessionID SessionRef, args NetworkRecord) (_retval NetworkRef, _err error) {
	if (IsMock) {
		return _class.Create__mock(sessionID, args)
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

func (_class NetworkClass) GetByUUID__mock(sessionID SessionRef, uuid string) (_retval NetworkRef, _err error) {
	log.Println("Network.GetByUUID not mocked")
	_err = errors.New("Network.GetByUUID not mocked")
	return
}
// Get a reference to the network instance with the specified UUID.
func (_class NetworkClass) GetByUUID(sessionID SessionRef, uuid string) (_retval NetworkRef, _err error) {
	if (IsMock) {
		return _class.GetByUUID__mock(sessionID, uuid)
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

func (_class NetworkClass) GetRecord__mock(sessionID SessionRef, self NetworkRef) (_retval NetworkRecord, _err error) {
	log.Println("Network.GetRecord not mocked")
	_err = errors.New("Network.GetRecord not mocked")
	return
}
// Get a record containing the current state of the given network.
func (_class NetworkClass) GetRecord(sessionID SessionRef, self NetworkRef) (_retval NetworkRecord, _err error) {
	if (IsMock) {
		return _class.GetRecord__mock(sessionID, self)
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
