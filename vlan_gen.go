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

type VLANRecord struct {
  // Unique identifier/object reference
	UUID string
  // interface on which traffic is tagged
	TaggedPIF PIFRef
  // interface on which traffic is untagged
	UntaggedPIF PIFRef
  // VLAN tag in use
	Tag int
  // additional configuration
	OtherConfig map[string]string
}

type VLANRef string

// A VLAN mux/demux
type VLANClass struct {
	client *Client
}


func VLANClassGetAllRecordsMockDefault(sessionID SessionRef) (_retval map[VLANRef]VLANRecord, _err error) {
	log.Println("VLAN.GetAllRecords not mocked")
	_err = errors.New("VLAN.GetAllRecords not mocked")
	return
}

var VLANClassGetAllRecordsMockedCallback = VLANClassGetAllRecordsMockDefault

func (_class VLANClass) GetAllRecordsMock(sessionID SessionRef) (_retval map[VLANRef]VLANRecord, _err error) {
	return VLANClassGetAllRecordsMockedCallback(sessionID)
}
// Return a map of VLAN references to VLAN records for all VLANs known to the system.
func (_class VLANClass) GetAllRecords(sessionID SessionRef) (_retval map[VLANRef]VLANRecord, _err error) {
	if IsMock {
		return _class.GetAllRecordsMock(sessionID)
	}	
	_method := "VLAN.get_all_records"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVLANRefToVLANRecordMapToGo(_method + " -> ", _result.Value)
	return
}


func VLANClassGetAllMockDefault(sessionID SessionRef) (_retval []VLANRef, _err error) {
	log.Println("VLAN.GetAll not mocked")
	_err = errors.New("VLAN.GetAll not mocked")
	return
}

var VLANClassGetAllMockedCallback = VLANClassGetAllMockDefault

func (_class VLANClass) GetAllMock(sessionID SessionRef) (_retval []VLANRef, _err error) {
	return VLANClassGetAllMockedCallback(sessionID)
}
// Return a list of all the VLANs known to the system.
func (_class VLANClass) GetAll(sessionID SessionRef) (_retval []VLANRef, _err error) {
	if IsMock {
		return _class.GetAllMock(sessionID)
	}	
	_method := "VLAN.get_all"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVLANRefSetToGo(_method + " -> ", _result.Value)
	return
}


func VLANClassDestroyMockDefault(sessionID SessionRef, self VLANRef) (_err error) {
	log.Println("VLAN.Destroy not mocked")
	_err = errors.New("VLAN.Destroy not mocked")
	return
}

var VLANClassDestroyMockedCallback = VLANClassDestroyMockDefault

func (_class VLANClass) DestroyMock(sessionID SessionRef, self VLANRef) (_err error) {
	return VLANClassDestroyMockedCallback(sessionID, self)
}
// Destroy a VLAN mux/demuxer
func (_class VLANClass) Destroy(sessionID SessionRef, self VLANRef) (_err error) {
	if IsMock {
		return _class.DestroyMock(sessionID, self)
	}	
	_method := "VLAN.destroy"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVLANRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg)
	return
}


func VLANClassCreateMockDefault(sessionID SessionRef, taggedPIF PIFRef, tag int, network NetworkRef) (_retval VLANRef, _err error) {
	log.Println("VLAN.Create not mocked")
	_err = errors.New("VLAN.Create not mocked")
	return
}

var VLANClassCreateMockedCallback = VLANClassCreateMockDefault

func (_class VLANClass) CreateMock(sessionID SessionRef, taggedPIF PIFRef, tag int, network NetworkRef) (_retval VLANRef, _err error) {
	return VLANClassCreateMockedCallback(sessionID, taggedPIF, tag, network)
}
// Create a VLAN mux/demuxer
func (_class VLANClass) Create(sessionID SessionRef, taggedPIF PIFRef, tag int, network NetworkRef) (_retval VLANRef, _err error) {
	if IsMock {
		return _class.CreateMock(sessionID, taggedPIF, tag, network)
	}	
	_method := "VLAN.create"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_taggedPIFArg, _err := convertPIFRefToXen(fmt.Sprintf("%s(%s)", _method, "tagged_PIF"), taggedPIF)
	if _err != nil {
		return
	}
	_tagArg, _err := convertIntToXen(fmt.Sprintf("%s(%s)", _method, "tag"), tag)
	if _err != nil {
		return
	}
	_networkArg, _err := convertNetworkRefToXen(fmt.Sprintf("%s(%s)", _method, "network"), network)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _taggedPIFArg, _tagArg, _networkArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVLANRefToGo(_method + " -> ", _result.Value)
	return
}


func VLANClassRemoveFromOtherConfigMockDefault(sessionID SessionRef, self VLANRef, key string) (_err error) {
	log.Println("VLAN.RemoveFromOtherConfig not mocked")
	_err = errors.New("VLAN.RemoveFromOtherConfig not mocked")
	return
}

var VLANClassRemoveFromOtherConfigMockedCallback = VLANClassRemoveFromOtherConfigMockDefault

func (_class VLANClass) RemoveFromOtherConfigMock(sessionID SessionRef, self VLANRef, key string) (_err error) {
	return VLANClassRemoveFromOtherConfigMockedCallback(sessionID, self, key)
}
// Remove the given key and its corresponding value from the other_config field of the given VLAN.  If the key is not in that Map, then do nothing.
func (_class VLANClass) RemoveFromOtherConfig(sessionID SessionRef, self VLANRef, key string) (_err error) {
	if IsMock {
		return _class.RemoveFromOtherConfigMock(sessionID, self, key)
	}	
	_method := "VLAN.remove_from_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVLANRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func VLANClassAddToOtherConfigMockDefault(sessionID SessionRef, self VLANRef, key string, value string) (_err error) {
	log.Println("VLAN.AddToOtherConfig not mocked")
	_err = errors.New("VLAN.AddToOtherConfig not mocked")
	return
}

var VLANClassAddToOtherConfigMockedCallback = VLANClassAddToOtherConfigMockDefault

func (_class VLANClass) AddToOtherConfigMock(sessionID SessionRef, self VLANRef, key string, value string) (_err error) {
	return VLANClassAddToOtherConfigMockedCallback(sessionID, self, key, value)
}
// Add the given key-value pair to the other_config field of the given VLAN.
func (_class VLANClass) AddToOtherConfig(sessionID SessionRef, self VLANRef, key string, value string) (_err error) {
	if IsMock {
		return _class.AddToOtherConfigMock(sessionID, self, key, value)
	}	
	_method := "VLAN.add_to_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVLANRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func VLANClassSetOtherConfigMockDefault(sessionID SessionRef, self VLANRef, value map[string]string) (_err error) {
	log.Println("VLAN.SetOtherConfig not mocked")
	_err = errors.New("VLAN.SetOtherConfig not mocked")
	return
}

var VLANClassSetOtherConfigMockedCallback = VLANClassSetOtherConfigMockDefault

func (_class VLANClass) SetOtherConfigMock(sessionID SessionRef, self VLANRef, value map[string]string) (_err error) {
	return VLANClassSetOtherConfigMockedCallback(sessionID, self, value)
}
// Set the other_config field of the given VLAN.
func (_class VLANClass) SetOtherConfig(sessionID SessionRef, self VLANRef, value map[string]string) (_err error) {
	if IsMock {
		return _class.SetOtherConfigMock(sessionID, self, value)
	}	
	_method := "VLAN.set_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVLANRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func VLANClassGetOtherConfigMockDefault(sessionID SessionRef, self VLANRef) (_retval map[string]string, _err error) {
	log.Println("VLAN.GetOtherConfig not mocked")
	_err = errors.New("VLAN.GetOtherConfig not mocked")
	return
}

var VLANClassGetOtherConfigMockedCallback = VLANClassGetOtherConfigMockDefault

func (_class VLANClass) GetOtherConfigMock(sessionID SessionRef, self VLANRef) (_retval map[string]string, _err error) {
	return VLANClassGetOtherConfigMockedCallback(sessionID, self)
}
// Get the other_config field of the given VLAN.
func (_class VLANClass) GetOtherConfig(sessionID SessionRef, self VLANRef) (_retval map[string]string, _err error) {
	if IsMock {
		return _class.GetOtherConfigMock(sessionID, self)
	}	
	_method := "VLAN.get_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVLANRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func VLANClassGetTagMockDefault(sessionID SessionRef, self VLANRef) (_retval int, _err error) {
	log.Println("VLAN.GetTag not mocked")
	_err = errors.New("VLAN.GetTag not mocked")
	return
}

var VLANClassGetTagMockedCallback = VLANClassGetTagMockDefault

func (_class VLANClass) GetTagMock(sessionID SessionRef, self VLANRef) (_retval int, _err error) {
	return VLANClassGetTagMockedCallback(sessionID, self)
}
// Get the tag field of the given VLAN.
func (_class VLANClass) GetTag(sessionID SessionRef, self VLANRef) (_retval int, _err error) {
	if IsMock {
		return _class.GetTagMock(sessionID, self)
	}	
	_method := "VLAN.get_tag"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVLANRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func VLANClassGetUntaggedPIFMockDefault(sessionID SessionRef, self VLANRef) (_retval PIFRef, _err error) {
	log.Println("VLAN.GetUntaggedPIF not mocked")
	_err = errors.New("VLAN.GetUntaggedPIF not mocked")
	return
}

var VLANClassGetUntaggedPIFMockedCallback = VLANClassGetUntaggedPIFMockDefault

func (_class VLANClass) GetUntaggedPIFMock(sessionID SessionRef, self VLANRef) (_retval PIFRef, _err error) {
	return VLANClassGetUntaggedPIFMockedCallback(sessionID, self)
}
// Get the untagged_PIF field of the given VLAN.
func (_class VLANClass) GetUntaggedPIF(sessionID SessionRef, self VLANRef) (_retval PIFRef, _err error) {
	if IsMock {
		return _class.GetUntaggedPIFMock(sessionID, self)
	}	
	_method := "VLAN.get_untagged_PIF"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVLANRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertPIFRefToGo(_method + " -> ", _result.Value)
	return
}


func VLANClassGetTaggedPIFMockDefault(sessionID SessionRef, self VLANRef) (_retval PIFRef, _err error) {
	log.Println("VLAN.GetTaggedPIF not mocked")
	_err = errors.New("VLAN.GetTaggedPIF not mocked")
	return
}

var VLANClassGetTaggedPIFMockedCallback = VLANClassGetTaggedPIFMockDefault

func (_class VLANClass) GetTaggedPIFMock(sessionID SessionRef, self VLANRef) (_retval PIFRef, _err error) {
	return VLANClassGetTaggedPIFMockedCallback(sessionID, self)
}
// Get the tagged_PIF field of the given VLAN.
func (_class VLANClass) GetTaggedPIF(sessionID SessionRef, self VLANRef) (_retval PIFRef, _err error) {
	if IsMock {
		return _class.GetTaggedPIFMock(sessionID, self)
	}	
	_method := "VLAN.get_tagged_PIF"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVLANRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertPIFRefToGo(_method + " -> ", _result.Value)
	return
}


func VLANClassGetUUIDMockDefault(sessionID SessionRef, self VLANRef) (_retval string, _err error) {
	log.Println("VLAN.GetUUID not mocked")
	_err = errors.New("VLAN.GetUUID not mocked")
	return
}

var VLANClassGetUUIDMockedCallback = VLANClassGetUUIDMockDefault

func (_class VLANClass) GetUUIDMock(sessionID SessionRef, self VLANRef) (_retval string, _err error) {
	return VLANClassGetUUIDMockedCallback(sessionID, self)
}
// Get the uuid field of the given VLAN.
func (_class VLANClass) GetUUID(sessionID SessionRef, self VLANRef) (_retval string, _err error) {
	if IsMock {
		return _class.GetUUIDMock(sessionID, self)
	}	
	_method := "VLAN.get_uuid"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVLANRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func VLANClassGetByUUIDMockDefault(sessionID SessionRef, uuid string) (_retval VLANRef, _err error) {
	log.Println("VLAN.GetByUUID not mocked")
	_err = errors.New("VLAN.GetByUUID not mocked")
	return
}

var VLANClassGetByUUIDMockedCallback = VLANClassGetByUUIDMockDefault

func (_class VLANClass) GetByUUIDMock(sessionID SessionRef, uuid string) (_retval VLANRef, _err error) {
	return VLANClassGetByUUIDMockedCallback(sessionID, uuid)
}
// Get a reference to the VLAN instance with the specified UUID.
func (_class VLANClass) GetByUUID(sessionID SessionRef, uuid string) (_retval VLANRef, _err error) {
	if IsMock {
		return _class.GetByUUIDMock(sessionID, uuid)
	}	
	_method := "VLAN.get_by_uuid"
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
	_retval, _err = convertVLANRefToGo(_method + " -> ", _result.Value)
	return
}


func VLANClassGetRecordMockDefault(sessionID SessionRef, self VLANRef) (_retval VLANRecord, _err error) {
	log.Println("VLAN.GetRecord not mocked")
	_err = errors.New("VLAN.GetRecord not mocked")
	return
}

var VLANClassGetRecordMockedCallback = VLANClassGetRecordMockDefault

func (_class VLANClass) GetRecordMock(sessionID SessionRef, self VLANRef) (_retval VLANRecord, _err error) {
	return VLANClassGetRecordMockedCallback(sessionID, self)
}
// Get a record containing the current state of the given VLAN.
func (_class VLANClass) GetRecord(sessionID SessionRef, self VLANRef) (_retval VLANRecord, _err error) {
	if IsMock {
		return _class.GetRecordMock(sessionID, self)
	}	
	_method := "VLAN.get_record"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVLANRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVLANRecordToGo(_method + " -> ", _result.Value)
	return
}
