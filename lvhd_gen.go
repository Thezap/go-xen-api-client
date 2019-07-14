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

type LVHDRecord struct {
  // Unique identifier/object reference
	UUID string
}

type LVHDRef string

// LVHD SR specific operations
type LVHDClass struct {
	client *Client
}


var LVHDClass_EnableThinProvisioningMockedCallback = func (sessionID SessionRef, host HostRef, sr SRRef, initialAllocation int, allocationQuantum int) (_retval string, _err error) {
	log.Println("LVHD.EnableThinProvisioning not mocked")
	_err = errors.New("LVHD.EnableThinProvisioning not mocked")
	return
}

func (_class LVHDClass) EnableThinProvisioningMock(sessionID SessionRef, host HostRef, sr SRRef, initialAllocation int, allocationQuantum int) (_retval string, _err error) {
	return LVHDClass_EnableThinProvisioningMockedCallback(sessionID, host, sr, initialAllocation, allocationQuantum)
}
// Upgrades an LVHD SR to enable thin-provisioning. Future VDIs created in this SR will be thinly-provisioned, although existing VDIs will be left alone. Note that the SR must be attached to the SRmaster for upgrade to work.
func (_class LVHDClass) EnableThinProvisioning(sessionID SessionRef, host HostRef, sr SRRef, initialAllocation int, allocationQuantum int) (_retval string, _err error) {
	if (IsMock) {
		return _class.EnableThinProvisioningMock(sessionID, host, sr, initialAllocation, allocationQuantum)
	}	
	_method := "LVHD.enable_thin_provisioning"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_hostArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "host"), host)
	if _err != nil {
		return
	}
	_srArg, _err := convertSRRefToXen(fmt.Sprintf("%s(%s)", _method, "SR"), sr)
	if _err != nil {
		return
	}
	_initialAllocationArg, _err := convertIntToXen(fmt.Sprintf("%s(%s)", _method, "initial_allocation"), initialAllocation)
	if _err != nil {
		return
	}
	_allocationQuantumArg, _err := convertIntToXen(fmt.Sprintf("%s(%s)", _method, "allocation_quantum"), allocationQuantum)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _hostArg, _srArg, _initialAllocationArg, _allocationQuantumArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringToGo(_method + " -> ", _result.Value)
	return
}


var LVHDClass_GetUUIDMockedCallback = func (sessionID SessionRef, self LVHDRef) (_retval string, _err error) {
	log.Println("LVHD.GetUUID not mocked")
	_err = errors.New("LVHD.GetUUID not mocked")
	return
}

func (_class LVHDClass) GetUUIDMock(sessionID SessionRef, self LVHDRef) (_retval string, _err error) {
	return LVHDClass_GetUUIDMockedCallback(sessionID, self)
}
// Get the uuid field of the given LVHD.
func (_class LVHDClass) GetUUID(sessionID SessionRef, self LVHDRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetUUIDMock(sessionID, self)
	}	
	_method := "LVHD.get_uuid"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertLVHDRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var LVHDClass_GetByUUIDMockedCallback = func (sessionID SessionRef, uuid string) (_retval LVHDRef, _err error) {
	log.Println("LVHD.GetByUUID not mocked")
	_err = errors.New("LVHD.GetByUUID not mocked")
	return
}

func (_class LVHDClass) GetByUUIDMock(sessionID SessionRef, uuid string) (_retval LVHDRef, _err error) {
	return LVHDClass_GetByUUIDMockedCallback(sessionID, uuid)
}
// Get a reference to the LVHD instance with the specified UUID.
func (_class LVHDClass) GetByUUID(sessionID SessionRef, uuid string) (_retval LVHDRef, _err error) {
	if (IsMock) {
		return _class.GetByUUIDMock(sessionID, uuid)
	}	
	_method := "LVHD.get_by_uuid"
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
	_retval, _err = convertLVHDRefToGo(_method + " -> ", _result.Value)
	return
}


var LVHDClass_GetRecordMockedCallback = func (sessionID SessionRef, self LVHDRef) (_retval LVHDRecord, _err error) {
	log.Println("LVHD.GetRecord not mocked")
	_err = errors.New("LVHD.GetRecord not mocked")
	return
}

func (_class LVHDClass) GetRecordMock(sessionID SessionRef, self LVHDRef) (_retval LVHDRecord, _err error) {
	return LVHDClass_GetRecordMockedCallback(sessionID, self)
}
// Get a record containing the current state of the given LVHD.
func (_class LVHDClass) GetRecord(sessionID SessionRef, self LVHDRef) (_retval LVHDRecord, _err error) {
	if (IsMock) {
		return _class.GetRecordMock(sessionID, self)
	}	
	_method := "LVHD.get_record"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertLVHDRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertLVHDRecordToGo(_method + " -> ", _result.Value)
	return
}
