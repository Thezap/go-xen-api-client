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

type BondMode string

const (
  // Source-level balancing
	BondModeBalanceSlb BondMode = "balance-slb"
  // Active/passive bonding: only one NIC is carrying traffic
	BondModeActiveBackup BondMode = "active-backup"
  // Link aggregation control protocol
	BondModeLacp BondMode = "lacp"
)

type BondRecord struct {
  // Unique identifier/object reference
	UUID string
  // The bonded interface
	Master PIFRef
  // The interfaces which are part of this bond
	Slaves []PIFRef
  // additional configuration
	OtherConfig map[string]string
  // The PIF of which the IP configuration and MAC were copied to the bond, and which will receive all configuration/VLANs/VIFs on the bond if the bond is destroyed
	PrimarySlave PIFRef
  // The algorithm used to distribute traffic among the bonded NICs
	Mode BondMode
  // Additional configuration properties specific to the bond mode.
	Properties map[string]string
  // Number of links up in this bond
	LinksUp int
}

type BondRef string

// 
type BondClass struct {
	client *Client
}

func (_class BondClass) GetAllRecords__mock(sessionID SessionRef) (_retval map[BondRef]BondRecord, _err error) {
	log.Println("Bond.GetAllRecords not mocked")
	_err = errors.New("Bond.GetAllRecords not mocked")
	return
}
// Return a map of Bond references to Bond records for all Bonds known to the system.
func (_class BondClass) GetAllRecords(sessionID SessionRef) (_retval map[BondRef]BondRecord, _err error) {
	if (IsMock) {
		return _class.GetAllRecords__mock(sessionID)
	}	
	_method := "Bond.get_all_records"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertBondRefToBondRecordMapToGo(_method + " -> ", _result.Value)
	return
}

func (_class BondClass) GetAll__mock(sessionID SessionRef) (_retval []BondRef, _err error) {
	log.Println("Bond.GetAll not mocked")
	_err = errors.New("Bond.GetAll not mocked")
	return
}
// Return a list of all the Bonds known to the system.
func (_class BondClass) GetAll(sessionID SessionRef) (_retval []BondRef, _err error) {
	if (IsMock) {
		return _class.GetAll__mock(sessionID)
	}	
	_method := "Bond.get_all"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertBondRefSetToGo(_method + " -> ", _result.Value)
	return
}

func (_class BondClass) SetProperty__mock(sessionID SessionRef, self BondRef, name string, value string) (_err error) {
	log.Println("Bond.SetProperty not mocked")
	_err = errors.New("Bond.SetProperty not mocked")
	return
}
// Set the value of a property of the bond
func (_class BondClass) SetProperty(sessionID SessionRef, self BondRef, name string, value string) (_err error) {
	if (IsMock) {
		return _class.SetProperty__mock(sessionID, self, name, value)
	}	
	_method := "Bond.set_property"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertBondRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_nameArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "name"), name)
	if _err != nil {
		return
	}
	_valueArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _nameArg, _valueArg)
	return
}

func (_class BondClass) SetMode__mock(sessionID SessionRef, self BondRef, value BondMode) (_err error) {
	log.Println("Bond.SetMode not mocked")
	_err = errors.New("Bond.SetMode not mocked")
	return
}
// Change the bond mode
func (_class BondClass) SetMode(sessionID SessionRef, self BondRef, value BondMode) (_err error) {
	if (IsMock) {
		return _class.SetMode__mock(sessionID, self, value)
	}	
	_method := "Bond.set_mode"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertBondRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertEnumBondModeToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}

func (_class BondClass) Destroy__mock(sessionID SessionRef, self BondRef) (_err error) {
	log.Println("Bond.Destroy not mocked")
	_err = errors.New("Bond.Destroy not mocked")
	return
}
// Destroy an interface bond
func (_class BondClass) Destroy(sessionID SessionRef, self BondRef) (_err error) {
	if (IsMock) {
		return _class.Destroy__mock(sessionID, self)
	}	
	_method := "Bond.destroy"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertBondRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg)
	return
}

func (_class BondClass) Create__mock(sessionID SessionRef, network NetworkRef, members []PIFRef, mac string, mode BondMode, properties map[string]string) (_retval BondRef, _err error) {
	log.Println("Bond.Create not mocked")
	_err = errors.New("Bond.Create not mocked")
	return
}
// Create an interface bond
func (_class BondClass) Create(sessionID SessionRef, network NetworkRef, members []PIFRef, mac string, mode BondMode, properties map[string]string) (_retval BondRef, _err error) {
	if (IsMock) {
		return _class.Create__mock(sessionID, network, members, mac, mode, properties)
	}	
	_method := "Bond.create"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_networkArg, _err := convertNetworkRefToXen(fmt.Sprintf("%s(%s)", _method, "network"), network)
	if _err != nil {
		return
	}
	_membersArg, _err := convertPIFRefSetToXen(fmt.Sprintf("%s(%s)", _method, "members"), members)
	if _err != nil {
		return
	}
	_macArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "MAC"), mac)
	if _err != nil {
		return
	}
	_modeArg, _err := convertEnumBondModeToXen(fmt.Sprintf("%s(%s)", _method, "mode"), mode)
	if _err != nil {
		return
	}
	_propertiesArg, _err := convertStringToStringMapToXen(fmt.Sprintf("%s(%s)", _method, "properties"), properties)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _networkArg, _membersArg, _macArg, _modeArg, _propertiesArg)
	if _err != nil {
		return
	}
	_retval, _err = convertBondRefToGo(_method + " -> ", _result.Value)
	return
}

func (_class BondClass) RemoveFromOtherConfig__mock(sessionID SessionRef, self BondRef, key string) (_err error) {
	log.Println("Bond.RemoveFromOtherConfig not mocked")
	_err = errors.New("Bond.RemoveFromOtherConfig not mocked")
	return
}
// Remove the given key and its corresponding value from the other_config field of the given Bond.  If the key is not in that Map, then do nothing.
func (_class BondClass) RemoveFromOtherConfig(sessionID SessionRef, self BondRef, key string) (_err error) {
	if (IsMock) {
		return _class.RemoveFromOtherConfig__mock(sessionID, self, key)
	}	
	_method := "Bond.remove_from_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertBondRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class BondClass) AddToOtherConfig__mock(sessionID SessionRef, self BondRef, key string, value string) (_err error) {
	log.Println("Bond.AddToOtherConfig not mocked")
	_err = errors.New("Bond.AddToOtherConfig not mocked")
	return
}
// Add the given key-value pair to the other_config field of the given Bond.
func (_class BondClass) AddToOtherConfig(sessionID SessionRef, self BondRef, key string, value string) (_err error) {
	if (IsMock) {
		return _class.AddToOtherConfig__mock(sessionID, self, key, value)
	}	
	_method := "Bond.add_to_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertBondRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class BondClass) SetOtherConfig__mock(sessionID SessionRef, self BondRef, value map[string]string) (_err error) {
	log.Println("Bond.SetOtherConfig not mocked")
	_err = errors.New("Bond.SetOtherConfig not mocked")
	return
}
// Set the other_config field of the given Bond.
func (_class BondClass) SetOtherConfig(sessionID SessionRef, self BondRef, value map[string]string) (_err error) {
	if (IsMock) {
		return _class.SetOtherConfig__mock(sessionID, self, value)
	}	
	_method := "Bond.set_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertBondRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class BondClass) GetLinksUp__mock(sessionID SessionRef, self BondRef) (_retval int, _err error) {
	log.Println("Bond.GetLinksUp not mocked")
	_err = errors.New("Bond.GetLinksUp not mocked")
	return
}
// Get the links_up field of the given Bond.
func (_class BondClass) GetLinksUp(sessionID SessionRef, self BondRef) (_retval int, _err error) {
	if (IsMock) {
		return _class.GetLinksUp__mock(sessionID, self)
	}	
	_method := "Bond.get_links_up"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertBondRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class BondClass) GetProperties__mock(sessionID SessionRef, self BondRef) (_retval map[string]string, _err error) {
	log.Println("Bond.GetProperties not mocked")
	_err = errors.New("Bond.GetProperties not mocked")
	return
}
// Get the properties field of the given Bond.
func (_class BondClass) GetProperties(sessionID SessionRef, self BondRef) (_retval map[string]string, _err error) {
	if (IsMock) {
		return _class.GetProperties__mock(sessionID, self)
	}	
	_method := "Bond.get_properties"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertBondRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class BondClass) GetMode__mock(sessionID SessionRef, self BondRef) (_retval BondMode, _err error) {
	log.Println("Bond.GetMode not mocked")
	_err = errors.New("Bond.GetMode not mocked")
	return
}
// Get the mode field of the given Bond.
func (_class BondClass) GetMode(sessionID SessionRef, self BondRef) (_retval BondMode, _err error) {
	if (IsMock) {
		return _class.GetMode__mock(sessionID, self)
	}	
	_method := "Bond.get_mode"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertBondRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertEnumBondModeToGo(_method + " -> ", _result.Value)
	return
}

func (_class BondClass) GetPrimarySlave__mock(sessionID SessionRef, self BondRef) (_retval PIFRef, _err error) {
	log.Println("Bond.GetPrimarySlave not mocked")
	_err = errors.New("Bond.GetPrimarySlave not mocked")
	return
}
// Get the primary_slave field of the given Bond.
func (_class BondClass) GetPrimarySlave(sessionID SessionRef, self BondRef) (_retval PIFRef, _err error) {
	if (IsMock) {
		return _class.GetPrimarySlave__mock(sessionID, self)
	}	
	_method := "Bond.get_primary_slave"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertBondRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class BondClass) GetOtherConfig__mock(sessionID SessionRef, self BondRef) (_retval map[string]string, _err error) {
	log.Println("Bond.GetOtherConfig not mocked")
	_err = errors.New("Bond.GetOtherConfig not mocked")
	return
}
// Get the other_config field of the given Bond.
func (_class BondClass) GetOtherConfig(sessionID SessionRef, self BondRef) (_retval map[string]string, _err error) {
	if (IsMock) {
		return _class.GetOtherConfig__mock(sessionID, self)
	}	
	_method := "Bond.get_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertBondRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class BondClass) GetSlaves__mock(sessionID SessionRef, self BondRef) (_retval []PIFRef, _err error) {
	log.Println("Bond.GetSlaves not mocked")
	_err = errors.New("Bond.GetSlaves not mocked")
	return
}
// Get the slaves field of the given Bond.
func (_class BondClass) GetSlaves(sessionID SessionRef, self BondRef) (_retval []PIFRef, _err error) {
	if (IsMock) {
		return _class.GetSlaves__mock(sessionID, self)
	}	
	_method := "Bond.get_slaves"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertBondRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class BondClass) GetMaster__mock(sessionID SessionRef, self BondRef) (_retval PIFRef, _err error) {
	log.Println("Bond.GetMaster not mocked")
	_err = errors.New("Bond.GetMaster not mocked")
	return
}
// Get the master field of the given Bond.
func (_class BondClass) GetMaster(sessionID SessionRef, self BondRef) (_retval PIFRef, _err error) {
	if (IsMock) {
		return _class.GetMaster__mock(sessionID, self)
	}	
	_method := "Bond.get_master"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertBondRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class BondClass) GetUUID__mock(sessionID SessionRef, self BondRef) (_retval string, _err error) {
	log.Println("Bond.GetUUID not mocked")
	_err = errors.New("Bond.GetUUID not mocked")
	return
}
// Get the uuid field of the given Bond.
func (_class BondClass) GetUUID(sessionID SessionRef, self BondRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetUUID__mock(sessionID, self)
	}	
	_method := "Bond.get_uuid"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertBondRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class BondClass) GetByUUID__mock(sessionID SessionRef, uuid string) (_retval BondRef, _err error) {
	log.Println("Bond.GetByUUID not mocked")
	_err = errors.New("Bond.GetByUUID not mocked")
	return
}
// Get a reference to the Bond instance with the specified UUID.
func (_class BondClass) GetByUUID(sessionID SessionRef, uuid string) (_retval BondRef, _err error) {
	if (IsMock) {
		return _class.GetByUUID__mock(sessionID, uuid)
	}	
	_method := "Bond.get_by_uuid"
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
	_retval, _err = convertBondRefToGo(_method + " -> ", _result.Value)
	return
}

func (_class BondClass) GetRecord__mock(sessionID SessionRef, self BondRef) (_retval BondRecord, _err error) {
	log.Println("Bond.GetRecord not mocked")
	_err = errors.New("Bond.GetRecord not mocked")
	return
}
// Get a record containing the current state of the given Bond.
func (_class BondClass) GetRecord(sessionID SessionRef, self BondRef) (_retval BondRecord, _err error) {
	if (IsMock) {
		return _class.GetRecord__mock(sessionID, self)
	}	
	_method := "Bond.get_record"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertBondRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertBondRecordToGo(_method + " -> ", _result.Value)
	return
}
