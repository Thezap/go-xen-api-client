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

type TunnelRecord struct {
  // Unique identifier/object reference
	UUID string
  // The interface through which the tunnel is accessed
	AccessPIF PIFRef
  // The interface used by the tunnel
	TransportPIF PIFRef
  // Status information about the tunnel
	Status map[string]string
  // Additional configuration
	OtherConfig map[string]string
}

type TunnelRef string

// A tunnel for network traffic
type TunnelClass struct {
	client *Client
}

func (_class TunnelClass) GetAllRecords__mock(sessionID SessionRef) (_retval map[TunnelRef]TunnelRecord, _err error) {
	log.Println("Tunnel.GetAllRecords not mocked")
	_err = errors.New("Tunnel.GetAllRecords not mocked")
	return
}
// Return a map of tunnel references to tunnel records for all tunnels known to the system.
func (_class TunnelClass) GetAllRecords(sessionID SessionRef) (_retval map[TunnelRef]TunnelRecord, _err error) {
	if (IsMock) {
		return _class.GetAllRecords__mock(sessionID)
	}	
	_method := "tunnel.get_all_records"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTunnelRefToTunnelRecordMapToGo(_method + " -> ", _result.Value)
	return
}

func (_class TunnelClass) GetAll__mock(sessionID SessionRef) (_retval []TunnelRef, _err error) {
	log.Println("Tunnel.GetAll not mocked")
	_err = errors.New("Tunnel.GetAll not mocked")
	return
}
// Return a list of all the tunnels known to the system.
func (_class TunnelClass) GetAll(sessionID SessionRef) (_retval []TunnelRef, _err error) {
	if (IsMock) {
		return _class.GetAll__mock(sessionID)
	}	
	_method := "tunnel.get_all"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTunnelRefSetToGo(_method + " -> ", _result.Value)
	return
}

func (_class TunnelClass) Destroy__mock(sessionID SessionRef, self TunnelRef) (_err error) {
	log.Println("Tunnel.Destroy not mocked")
	_err = errors.New("Tunnel.Destroy not mocked")
	return
}
// Destroy a tunnel
func (_class TunnelClass) Destroy(sessionID SessionRef, self TunnelRef) (_err error) {
	if (IsMock) {
		return _class.Destroy__mock(sessionID, self)
	}	
	_method := "tunnel.destroy"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertTunnelRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg)
	return
}

func (_class TunnelClass) Create__mock(sessionID SessionRef, transportPIF PIFRef, network NetworkRef) (_retval TunnelRef, _err error) {
	log.Println("Tunnel.Create not mocked")
	_err = errors.New("Tunnel.Create not mocked")
	return
}
// Create a tunnel
//
// Errors:
//  OPENVSWITCH_NOT_ACTIVE - This operation needs the OpenVSwitch networking backend to be enabled on all hosts in the pool.
//  TRANSPORT_PIF_NOT_CONFIGURED - The tunnel transport PIF has no IP configuration set.
//  IS_TUNNEL_ACCESS_PIF - You tried to create a VLAN or tunnel on top of a tunnel access PIF - use the underlying transport PIF instead.
func (_class TunnelClass) Create(sessionID SessionRef, transportPIF PIFRef, network NetworkRef) (_retval TunnelRef, _err error) {
	if (IsMock) {
		return _class.Create__mock(sessionID, transportPIF, network)
	}	
	_method := "tunnel.create"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_transportPIFArg, _err := convertPIFRefToXen(fmt.Sprintf("%s(%s)", _method, "transport_PIF"), transportPIF)
	if _err != nil {
		return
	}
	_networkArg, _err := convertNetworkRefToXen(fmt.Sprintf("%s(%s)", _method, "network"), network)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _transportPIFArg, _networkArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTunnelRefToGo(_method + " -> ", _result.Value)
	return
}

func (_class TunnelClass) RemoveFromOtherConfig__mock(sessionID SessionRef, self TunnelRef, key string) (_err error) {
	log.Println("Tunnel.RemoveFromOtherConfig not mocked")
	_err = errors.New("Tunnel.RemoveFromOtherConfig not mocked")
	return
}
// Remove the given key and its corresponding value from the other_config field of the given tunnel.  If the key is not in that Map, then do nothing.
func (_class TunnelClass) RemoveFromOtherConfig(sessionID SessionRef, self TunnelRef, key string) (_err error) {
	if (IsMock) {
		return _class.RemoveFromOtherConfig__mock(sessionID, self, key)
	}	
	_method := "tunnel.remove_from_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertTunnelRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class TunnelClass) AddToOtherConfig__mock(sessionID SessionRef, self TunnelRef, key string, value string) (_err error) {
	log.Println("Tunnel.AddToOtherConfig not mocked")
	_err = errors.New("Tunnel.AddToOtherConfig not mocked")
	return
}
// Add the given key-value pair to the other_config field of the given tunnel.
func (_class TunnelClass) AddToOtherConfig(sessionID SessionRef, self TunnelRef, key string, value string) (_err error) {
	if (IsMock) {
		return _class.AddToOtherConfig__mock(sessionID, self, key, value)
	}	
	_method := "tunnel.add_to_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertTunnelRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class TunnelClass) SetOtherConfig__mock(sessionID SessionRef, self TunnelRef, value map[string]string) (_err error) {
	log.Println("Tunnel.SetOtherConfig not mocked")
	_err = errors.New("Tunnel.SetOtherConfig not mocked")
	return
}
// Set the other_config field of the given tunnel.
func (_class TunnelClass) SetOtherConfig(sessionID SessionRef, self TunnelRef, value map[string]string) (_err error) {
	if (IsMock) {
		return _class.SetOtherConfig__mock(sessionID, self, value)
	}	
	_method := "tunnel.set_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertTunnelRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class TunnelClass) RemoveFromStatus__mock(sessionID SessionRef, self TunnelRef, key string) (_err error) {
	log.Println("Tunnel.RemoveFromStatus not mocked")
	_err = errors.New("Tunnel.RemoveFromStatus not mocked")
	return
}
// Remove the given key and its corresponding value from the status field of the given tunnel.  If the key is not in that Map, then do nothing.
func (_class TunnelClass) RemoveFromStatus(sessionID SessionRef, self TunnelRef, key string) (_err error) {
	if (IsMock) {
		return _class.RemoveFromStatus__mock(sessionID, self, key)
	}	
	_method := "tunnel.remove_from_status"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertTunnelRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class TunnelClass) AddToStatus__mock(sessionID SessionRef, self TunnelRef, key string, value string) (_err error) {
	log.Println("Tunnel.AddToStatus not mocked")
	_err = errors.New("Tunnel.AddToStatus not mocked")
	return
}
// Add the given key-value pair to the status field of the given tunnel.
func (_class TunnelClass) AddToStatus(sessionID SessionRef, self TunnelRef, key string, value string) (_err error) {
	if (IsMock) {
		return _class.AddToStatus__mock(sessionID, self, key, value)
	}	
	_method := "tunnel.add_to_status"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertTunnelRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class TunnelClass) SetStatus__mock(sessionID SessionRef, self TunnelRef, value map[string]string) (_err error) {
	log.Println("Tunnel.SetStatus not mocked")
	_err = errors.New("Tunnel.SetStatus not mocked")
	return
}
// Set the status field of the given tunnel.
func (_class TunnelClass) SetStatus(sessionID SessionRef, self TunnelRef, value map[string]string) (_err error) {
	if (IsMock) {
		return _class.SetStatus__mock(sessionID, self, value)
	}	
	_method := "tunnel.set_status"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertTunnelRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class TunnelClass) GetOtherConfig__mock(sessionID SessionRef, self TunnelRef) (_retval map[string]string, _err error) {
	log.Println("Tunnel.GetOtherConfig not mocked")
	_err = errors.New("Tunnel.GetOtherConfig not mocked")
	return
}
// Get the other_config field of the given tunnel.
func (_class TunnelClass) GetOtherConfig(sessionID SessionRef, self TunnelRef) (_retval map[string]string, _err error) {
	if (IsMock) {
		return _class.GetOtherConfig__mock(sessionID, self)
	}	
	_method := "tunnel.get_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertTunnelRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class TunnelClass) GetStatus__mock(sessionID SessionRef, self TunnelRef) (_retval map[string]string, _err error) {
	log.Println("Tunnel.GetStatus not mocked")
	_err = errors.New("Tunnel.GetStatus not mocked")
	return
}
// Get the status field of the given tunnel.
func (_class TunnelClass) GetStatus(sessionID SessionRef, self TunnelRef) (_retval map[string]string, _err error) {
	if (IsMock) {
		return _class.GetStatus__mock(sessionID, self)
	}	
	_method := "tunnel.get_status"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertTunnelRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class TunnelClass) GetTransportPIF__mock(sessionID SessionRef, self TunnelRef) (_retval PIFRef, _err error) {
	log.Println("Tunnel.GetTransportPIF not mocked")
	_err = errors.New("Tunnel.GetTransportPIF not mocked")
	return
}
// Get the transport_PIF field of the given tunnel.
func (_class TunnelClass) GetTransportPIF(sessionID SessionRef, self TunnelRef) (_retval PIFRef, _err error) {
	if (IsMock) {
		return _class.GetTransportPIF__mock(sessionID, self)
	}	
	_method := "tunnel.get_transport_PIF"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertTunnelRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class TunnelClass) GetAccessPIF__mock(sessionID SessionRef, self TunnelRef) (_retval PIFRef, _err error) {
	log.Println("Tunnel.GetAccessPIF not mocked")
	_err = errors.New("Tunnel.GetAccessPIF not mocked")
	return
}
// Get the access_PIF field of the given tunnel.
func (_class TunnelClass) GetAccessPIF(sessionID SessionRef, self TunnelRef) (_retval PIFRef, _err error) {
	if (IsMock) {
		return _class.GetAccessPIF__mock(sessionID, self)
	}	
	_method := "tunnel.get_access_PIF"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertTunnelRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class TunnelClass) GetUUID__mock(sessionID SessionRef, self TunnelRef) (_retval string, _err error) {
	log.Println("Tunnel.GetUUID not mocked")
	_err = errors.New("Tunnel.GetUUID not mocked")
	return
}
// Get the uuid field of the given tunnel.
func (_class TunnelClass) GetUUID(sessionID SessionRef, self TunnelRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetUUID__mock(sessionID, self)
	}	
	_method := "tunnel.get_uuid"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertTunnelRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class TunnelClass) GetByUUID__mock(sessionID SessionRef, uuid string) (_retval TunnelRef, _err error) {
	log.Println("Tunnel.GetByUUID not mocked")
	_err = errors.New("Tunnel.GetByUUID not mocked")
	return
}
// Get a reference to the tunnel instance with the specified UUID.
func (_class TunnelClass) GetByUUID(sessionID SessionRef, uuid string) (_retval TunnelRef, _err error) {
	if (IsMock) {
		return _class.GetByUUID__mock(sessionID, uuid)
	}	
	_method := "tunnel.get_by_uuid"
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
	_retval, _err = convertTunnelRefToGo(_method + " -> ", _result.Value)
	return
}

func (_class TunnelClass) GetRecord__mock(sessionID SessionRef, self TunnelRef) (_retval TunnelRecord, _err error) {
	log.Println("Tunnel.GetRecord not mocked")
	_err = errors.New("Tunnel.GetRecord not mocked")
	return
}
// Get a record containing the current state of the given tunnel.
func (_class TunnelClass) GetRecord(sessionID SessionRef, self TunnelRef) (_retval TunnelRecord, _err error) {
	if (IsMock) {
		return _class.GetRecord__mock(sessionID, self)
	}	
	_method := "tunnel.get_record"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertTunnelRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTunnelRecordToGo(_method + " -> ", _result.Value)
	return
}
