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

type VmssFrequency string

const (
  // Hourly snapshots
	VmssFrequencyHourly VmssFrequency = "hourly"
  // Daily snapshots
	VmssFrequencyDaily VmssFrequency = "daily"
  // Weekly snapshots
	VmssFrequencyWeekly VmssFrequency = "weekly"
)

type VmssType string

const (
  // The snapshot is a disk snapshot
	VmssTypeSnapshot VmssType = "snapshot"
  // The snapshot is a checkpoint
	VmssTypeCheckpoint VmssType = "checkpoint"
  // The snapshot is a VSS
	VmssTypeSnapshotWithQuiesce VmssType = "snapshot_with_quiesce"
)

type VMSSRecord struct {
  // Unique identifier/object reference
	UUID string
  // a human-readable name
	NameLabel string
  // a notes field containing human-readable description
	NameDescription string
  // enable or disable this snapshot schedule
	Enabled bool
  // type of the snapshot schedule
	Type VmssType
  // maximum number of snapshots that should be stored at any time
	RetainedSnapshots int
  // frequency of taking snapshot from snapshot schedule
	Frequency VmssFrequency
  // schedule of the snapshot containing 'hour', 'min', 'days'. Date/time-related information is in Local Timezone
	Schedule map[string]string
  // time of the last snapshot
	LastRunTime time.Time
  // all VMs attached to this snapshot schedule
	VMs []VMRef
}

type VMSSRef string

// VM Snapshot Schedule
type VMSSClass struct {
	client *Client
}

func (_class VMSSClass) GetAllRecords__mock(sessionID SessionRef) (_retval map[VMSSRef]VMSSRecord, _err error) {
	log.Println("VMSS.GetAllRecords not mocked")
	_err = errors.New("VMSS.GetAllRecords not mocked")
	return
}
// Return a map of VMSS references to VMSS records for all VMSSs known to the system.
func (_class VMSSClass) GetAllRecords(sessionID SessionRef) (_retval map[VMSSRef]VMSSRecord, _err error) {
	if (IsMock) {
		return _class.GetAllRecords__mock(sessionID)
	}	
	_method := "VMSS.get_all_records"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVMSSRefToVMSSRecordMapToGo(_method + " -> ", _result.Value)
	return
}

func (_class VMSSClass) GetAll__mock(sessionID SessionRef) (_retval []VMSSRef, _err error) {
	log.Println("VMSS.GetAll not mocked")
	_err = errors.New("VMSS.GetAll not mocked")
	return
}
// Return a list of all the VMSSs known to the system.
func (_class VMSSClass) GetAll(sessionID SessionRef) (_retval []VMSSRef, _err error) {
	if (IsMock) {
		return _class.GetAll__mock(sessionID)
	}	
	_method := "VMSS.get_all"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVMSSRefSetToGo(_method + " -> ", _result.Value)
	return
}

func (_class VMSSClass) SetType__mock(sessionID SessionRef, self VMSSRef, value VmssType) (_err error) {
	log.Println("VMSS.SetType not mocked")
	_err = errors.New("VMSS.SetType not mocked")
	return
}
// 
func (_class VMSSClass) SetType(sessionID SessionRef, self VMSSRef, value VmssType) (_err error) {
	if (IsMock) {
		return _class.SetType__mock(sessionID, self, value)
	}	
	_method := "VMSS.set_type"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMSSRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertEnumVmssTypeToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}

func (_class VMSSClass) SetLastRunTime__mock(sessionID SessionRef, self VMSSRef, value time.Time) (_err error) {
	log.Println("VMSS.SetLastRunTime not mocked")
	_err = errors.New("VMSS.SetLastRunTime not mocked")
	return
}
// 
func (_class VMSSClass) SetLastRunTime(sessionID SessionRef, self VMSSRef, value time.Time) (_err error) {
	if (IsMock) {
		return _class.SetLastRunTime__mock(sessionID, self, value)
	}	
	_method := "VMSS.set_last_run_time"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMSSRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertTimeToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}

func (_class VMSSClass) RemoveFromSchedule__mock(sessionID SessionRef, self VMSSRef, key string) (_err error) {
	log.Println("VMSS.RemoveFromSchedule not mocked")
	_err = errors.New("VMSS.RemoveFromSchedule not mocked")
	return
}
// 
func (_class VMSSClass) RemoveFromSchedule(sessionID SessionRef, self VMSSRef, key string) (_err error) {
	if (IsMock) {
		return _class.RemoveFromSchedule__mock(sessionID, self, key)
	}	
	_method := "VMSS.remove_from_schedule"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMSSRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class VMSSClass) AddToSchedule__mock(sessionID SessionRef, self VMSSRef, key string, value string) (_err error) {
	log.Println("VMSS.AddToSchedule not mocked")
	_err = errors.New("VMSS.AddToSchedule not mocked")
	return
}
// 
func (_class VMSSClass) AddToSchedule(sessionID SessionRef, self VMSSRef, key string, value string) (_err error) {
	if (IsMock) {
		return _class.AddToSchedule__mock(sessionID, self, key, value)
	}	
	_method := "VMSS.add_to_schedule"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMSSRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class VMSSClass) SetSchedule__mock(sessionID SessionRef, self VMSSRef, value map[string]string) (_err error) {
	log.Println("VMSS.SetSchedule not mocked")
	_err = errors.New("VMSS.SetSchedule not mocked")
	return
}
// 
func (_class VMSSClass) SetSchedule(sessionID SessionRef, self VMSSRef, value map[string]string) (_err error) {
	if (IsMock) {
		return _class.SetSchedule__mock(sessionID, self, value)
	}	
	_method := "VMSS.set_schedule"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMSSRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class VMSSClass) SetFrequency__mock(sessionID SessionRef, self VMSSRef, value VmssFrequency) (_err error) {
	log.Println("VMSS.SetFrequency not mocked")
	_err = errors.New("VMSS.SetFrequency not mocked")
	return
}
// Set the value of the frequency field
func (_class VMSSClass) SetFrequency(sessionID SessionRef, self VMSSRef, value VmssFrequency) (_err error) {
	if (IsMock) {
		return _class.SetFrequency__mock(sessionID, self, value)
	}	
	_method := "VMSS.set_frequency"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMSSRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertEnumVmssFrequencyToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}

func (_class VMSSClass) SetRetainedSnapshots__mock(sessionID SessionRef, self VMSSRef, value int) (_err error) {
	log.Println("VMSS.SetRetainedSnapshots not mocked")
	_err = errors.New("VMSS.SetRetainedSnapshots not mocked")
	return
}
// 
func (_class VMSSClass) SetRetainedSnapshots(sessionID SessionRef, self VMSSRef, value int) (_err error) {
	if (IsMock) {
		return _class.SetRetainedSnapshots__mock(sessionID, self, value)
	}	
	_method := "VMSS.set_retained_snapshots"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMSSRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class VMSSClass) SnapshotNow__mock(sessionID SessionRef, vmss VMSSRef) (_retval string, _err error) {
	log.Println("VMSS.SnapshotNow not mocked")
	_err = errors.New("VMSS.SnapshotNow not mocked")
	return
}
// This call executes the snapshot schedule immediately
func (_class VMSSClass) SnapshotNow(sessionID SessionRef, vmss VMSSRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.SnapshotNow__mock(sessionID, vmss)
	}	
	_method := "VMSS.snapshot_now"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_vmssArg, _err := convertVMSSRefToXen(fmt.Sprintf("%s(%s)", _method, "vmss"), vmss)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _vmssArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringToGo(_method + " -> ", _result.Value)
	return
}

func (_class VMSSClass) SetEnabled__mock(sessionID SessionRef, self VMSSRef, value bool) (_err error) {
	log.Println("VMSS.SetEnabled not mocked")
	_err = errors.New("VMSS.SetEnabled not mocked")
	return
}
// Set the enabled field of the given VMSS.
func (_class VMSSClass) SetEnabled(sessionID SessionRef, self VMSSRef, value bool) (_err error) {
	if (IsMock) {
		return _class.SetEnabled__mock(sessionID, self, value)
	}	
	_method := "VMSS.set_enabled"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMSSRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class VMSSClass) SetNameDescription__mock(sessionID SessionRef, self VMSSRef, value string) (_err error) {
	log.Println("VMSS.SetNameDescription not mocked")
	_err = errors.New("VMSS.SetNameDescription not mocked")
	return
}
// Set the name/description field of the given VMSS.
func (_class VMSSClass) SetNameDescription(sessionID SessionRef, self VMSSRef, value string) (_err error) {
	if (IsMock) {
		return _class.SetNameDescription__mock(sessionID, self, value)
	}	
	_method := "VMSS.set_name_description"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMSSRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class VMSSClass) SetNameLabel__mock(sessionID SessionRef, self VMSSRef, value string) (_err error) {
	log.Println("VMSS.SetNameLabel not mocked")
	_err = errors.New("VMSS.SetNameLabel not mocked")
	return
}
// Set the name/label field of the given VMSS.
func (_class VMSSClass) SetNameLabel(sessionID SessionRef, self VMSSRef, value string) (_err error) {
	if (IsMock) {
		return _class.SetNameLabel__mock(sessionID, self, value)
	}	
	_method := "VMSS.set_name_label"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMSSRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class VMSSClass) GetVMs__mock(sessionID SessionRef, self VMSSRef) (_retval []VMRef, _err error) {
	log.Println("VMSS.GetVMs not mocked")
	_err = errors.New("VMSS.GetVMs not mocked")
	return
}
// Get the VMs field of the given VMSS.
func (_class VMSSClass) GetVMs(sessionID SessionRef, self VMSSRef) (_retval []VMRef, _err error) {
	if (IsMock) {
		return _class.GetVMs__mock(sessionID, self)
	}	
	_method := "VMSS.get_VMs"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMSSRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class VMSSClass) GetLastRunTime__mock(sessionID SessionRef, self VMSSRef) (_retval time.Time, _err error) {
	log.Println("VMSS.GetLastRunTime not mocked")
	_err = errors.New("VMSS.GetLastRunTime not mocked")
	return
}
// Get the last_run_time field of the given VMSS.
func (_class VMSSClass) GetLastRunTime(sessionID SessionRef, self VMSSRef) (_retval time.Time, _err error) {
	if (IsMock) {
		return _class.GetLastRunTime__mock(sessionID, self)
	}	
	_method := "VMSS.get_last_run_time"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMSSRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTimeToGo(_method + " -> ", _result.Value)
	return
}

func (_class VMSSClass) GetSchedule__mock(sessionID SessionRef, self VMSSRef) (_retval map[string]string, _err error) {
	log.Println("VMSS.GetSchedule not mocked")
	_err = errors.New("VMSS.GetSchedule not mocked")
	return
}
// Get the schedule field of the given VMSS.
func (_class VMSSClass) GetSchedule(sessionID SessionRef, self VMSSRef) (_retval map[string]string, _err error) {
	if (IsMock) {
		return _class.GetSchedule__mock(sessionID, self)
	}	
	_method := "VMSS.get_schedule"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMSSRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class VMSSClass) GetFrequency__mock(sessionID SessionRef, self VMSSRef) (_retval VmssFrequency, _err error) {
	log.Println("VMSS.GetFrequency not mocked")
	_err = errors.New("VMSS.GetFrequency not mocked")
	return
}
// Get the frequency field of the given VMSS.
func (_class VMSSClass) GetFrequency(sessionID SessionRef, self VMSSRef) (_retval VmssFrequency, _err error) {
	if (IsMock) {
		return _class.GetFrequency__mock(sessionID, self)
	}	
	_method := "VMSS.get_frequency"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMSSRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertEnumVmssFrequencyToGo(_method + " -> ", _result.Value)
	return
}

func (_class VMSSClass) GetRetainedSnapshots__mock(sessionID SessionRef, self VMSSRef) (_retval int, _err error) {
	log.Println("VMSS.GetRetainedSnapshots not mocked")
	_err = errors.New("VMSS.GetRetainedSnapshots not mocked")
	return
}
// Get the retained_snapshots field of the given VMSS.
func (_class VMSSClass) GetRetainedSnapshots(sessionID SessionRef, self VMSSRef) (_retval int, _err error) {
	if (IsMock) {
		return _class.GetRetainedSnapshots__mock(sessionID, self)
	}	
	_method := "VMSS.get_retained_snapshots"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMSSRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class VMSSClass) GetType__mock(sessionID SessionRef, self VMSSRef) (_retval VmssType, _err error) {
	log.Println("VMSS.GetType not mocked")
	_err = errors.New("VMSS.GetType not mocked")
	return
}
// Get the type field of the given VMSS.
func (_class VMSSClass) GetType(sessionID SessionRef, self VMSSRef) (_retval VmssType, _err error) {
	if (IsMock) {
		return _class.GetType__mock(sessionID, self)
	}	
	_method := "VMSS.get_type"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMSSRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertEnumVmssTypeToGo(_method + " -> ", _result.Value)
	return
}

func (_class VMSSClass) GetEnabled__mock(sessionID SessionRef, self VMSSRef) (_retval bool, _err error) {
	log.Println("VMSS.GetEnabled not mocked")
	_err = errors.New("VMSS.GetEnabled not mocked")
	return
}
// Get the enabled field of the given VMSS.
func (_class VMSSClass) GetEnabled(sessionID SessionRef, self VMSSRef) (_retval bool, _err error) {
	if (IsMock) {
		return _class.GetEnabled__mock(sessionID, self)
	}	
	_method := "VMSS.get_enabled"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMSSRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class VMSSClass) GetNameDescription__mock(sessionID SessionRef, self VMSSRef) (_retval string, _err error) {
	log.Println("VMSS.GetNameDescription not mocked")
	_err = errors.New("VMSS.GetNameDescription not mocked")
	return
}
// Get the name/description field of the given VMSS.
func (_class VMSSClass) GetNameDescription(sessionID SessionRef, self VMSSRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetNameDescription__mock(sessionID, self)
	}	
	_method := "VMSS.get_name_description"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMSSRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class VMSSClass) GetNameLabel__mock(sessionID SessionRef, self VMSSRef) (_retval string, _err error) {
	log.Println("VMSS.GetNameLabel not mocked")
	_err = errors.New("VMSS.GetNameLabel not mocked")
	return
}
// Get the name/label field of the given VMSS.
func (_class VMSSClass) GetNameLabel(sessionID SessionRef, self VMSSRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetNameLabel__mock(sessionID, self)
	}	
	_method := "VMSS.get_name_label"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMSSRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class VMSSClass) GetUUID__mock(sessionID SessionRef, self VMSSRef) (_retval string, _err error) {
	log.Println("VMSS.GetUUID not mocked")
	_err = errors.New("VMSS.GetUUID not mocked")
	return
}
// Get the uuid field of the given VMSS.
func (_class VMSSClass) GetUUID(sessionID SessionRef, self VMSSRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetUUID__mock(sessionID, self)
	}	
	_method := "VMSS.get_uuid"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMSSRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class VMSSClass) GetByNameLabel__mock(sessionID SessionRef, label string) (_retval []VMSSRef, _err error) {
	log.Println("VMSS.GetByNameLabel not mocked")
	_err = errors.New("VMSS.GetByNameLabel not mocked")
	return
}
// Get all the VMSS instances with the given label.
func (_class VMSSClass) GetByNameLabel(sessionID SessionRef, label string) (_retval []VMSSRef, _err error) {
	if (IsMock) {
		return _class.GetByNameLabel__mock(sessionID, label)
	}	
	_method := "VMSS.get_by_name_label"
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
	_retval, _err = convertVMSSRefSetToGo(_method + " -> ", _result.Value)
	return
}

func (_class VMSSClass) Destroy__mock(sessionID SessionRef, self VMSSRef) (_err error) {
	log.Println("VMSS.Destroy not mocked")
	_err = errors.New("VMSS.Destroy not mocked")
	return
}
// Destroy the specified VMSS instance.
func (_class VMSSClass) Destroy(sessionID SessionRef, self VMSSRef) (_err error) {
	if (IsMock) {
		return _class.Destroy__mock(sessionID, self)
	}	
	_method := "VMSS.destroy"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMSSRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg)
	return
}

func (_class VMSSClass) Create__mock(sessionID SessionRef, args VMSSRecord) (_retval VMSSRef, _err error) {
	log.Println("VMSS.Create not mocked")
	_err = errors.New("VMSS.Create not mocked")
	return
}
// Create a new VMSS instance, and return its handle.
// The constructor args are: name_label, name_description, enabled, type*, retained_snapshots, frequency*, schedule (* = non-optional).
func (_class VMSSClass) Create(sessionID SessionRef, args VMSSRecord) (_retval VMSSRef, _err error) {
	if (IsMock) {
		return _class.Create__mock(sessionID, args)
	}	
	_method := "VMSS.create"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_argsArg, _err := convertVMSSRecordToXen(fmt.Sprintf("%s(%s)", _method, "args"), args)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _argsArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVMSSRefToGo(_method + " -> ", _result.Value)
	return
}

func (_class VMSSClass) GetByUUID__mock(sessionID SessionRef, uuid string) (_retval VMSSRef, _err error) {
	log.Println("VMSS.GetByUUID not mocked")
	_err = errors.New("VMSS.GetByUUID not mocked")
	return
}
// Get a reference to the VMSS instance with the specified UUID.
func (_class VMSSClass) GetByUUID(sessionID SessionRef, uuid string) (_retval VMSSRef, _err error) {
	if (IsMock) {
		return _class.GetByUUID__mock(sessionID, uuid)
	}	
	_method := "VMSS.get_by_uuid"
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
	_retval, _err = convertVMSSRefToGo(_method + " -> ", _result.Value)
	return
}

func (_class VMSSClass) GetRecord__mock(sessionID SessionRef, self VMSSRef) (_retval VMSSRecord, _err error) {
	log.Println("VMSS.GetRecord not mocked")
	_err = errors.New("VMSS.GetRecord not mocked")
	return
}
// Get a record containing the current state of the given VMSS.
func (_class VMSSClass) GetRecord(sessionID SessionRef, self VMSSRef) (_retval VMSSRecord, _err error) {
	if (IsMock) {
		return _class.GetRecord__mock(sessionID, self)
	}	
	_method := "VMSS.get_record"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMSSRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVMSSRecordToGo(_method + " -> ", _result.Value)
	return
}
