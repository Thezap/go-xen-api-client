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

type VmppBackupType string

const (
  // The backup is a snapshot
	VmppBackupTypeSnapshot VmppBackupType = "snapshot"
  // The backup is a checkpoint
	VmppBackupTypeCheckpoint VmppBackupType = "checkpoint"
)

type VmppBackupFrequency string

const (
  // Hourly backups
	VmppBackupFrequencyHourly VmppBackupFrequency = "hourly"
  // Daily backups
	VmppBackupFrequencyDaily VmppBackupFrequency = "daily"
  // Weekly backups
	VmppBackupFrequencyWeekly VmppBackupFrequency = "weekly"
)

type VmppArchiveFrequency string

const (
  // Never archive
	VmppArchiveFrequencyNever VmppArchiveFrequency = "never"
  // Archive after backup
	VmppArchiveFrequencyAlwaysAfterBackup VmppArchiveFrequency = "always_after_backup"
  // Daily archives
	VmppArchiveFrequencyDaily VmppArchiveFrequency = "daily"
  // Weekly backups
	VmppArchiveFrequencyWeekly VmppArchiveFrequency = "weekly"
)

type VmppArchiveTargetType string

const (
  // No target config
	VmppArchiveTargetTypeNone VmppArchiveTargetType = "none"
  // CIFS target config
	VmppArchiveTargetTypeCifs VmppArchiveTargetType = "cifs"
  // NFS target config
	VmppArchiveTargetTypeNfs VmppArchiveTargetType = "nfs"
)

type VMPPRecord struct {
  // Unique identifier/object reference
	UUID string
  // a human-readable name
	NameLabel string
  // a notes field containing human-readable description
	NameDescription string
  // enable or disable this policy
	IsPolicyEnabled bool
  // type of the backup sub-policy
	BackupType VmppBackupType
  // maximum number of backups that should be stored at any time
	BackupRetentionValue int
  // frequency of the backup schedule
	BackupFrequency VmppBackupFrequency
  // schedule of the backup containing 'hour', 'min', 'days'. Date/time-related information is in Local Timezone
	BackupSchedule map[string]string
  // true if this protection policy's backup is running
	IsBackupRunning bool
  // time of the last backup
	BackupLastRunTime time.Time
  // type of the archive target config
	ArchiveTargetType VmppArchiveTargetType
  // configuration for the archive, including its 'location', 'username', 'password'
	ArchiveTargetConfig map[string]string
  // frequency of the archive schedule
	ArchiveFrequency VmppArchiveFrequency
  // schedule of the archive containing 'hour', 'min', 'days'. Date/time-related information is in Local Timezone
	ArchiveSchedule map[string]string
  // true if this protection policy's archive is running
	IsArchiveRunning bool
  // time of the last archive
	ArchiveLastRunTime time.Time
  // all VMs attached to this protection policy
	VMs []VMRef
  // true if alarm is enabled for this policy
	IsAlarmEnabled bool
  // configuration for the alarm
	AlarmConfig map[string]string
  // recent alerts
	RecentAlerts []string
}

type VMPPRef string

// VM Protection Policy
type VMPPClass struct {
	client *Client
}


var VMPPClass_GetAllRecordsMockedCallback = func (sessionID SessionRef) (_retval map[VMPPRef]VMPPRecord, _err error) {
	log.Println("VMPP.GetAllRecords not mocked")
	_err = errors.New("VMPP.GetAllRecords not mocked")
	return
}

func (_class VMPPClass) GetAllRecordsMock(sessionID SessionRef) (_retval map[VMPPRef]VMPPRecord, _err error) {
	return VMPPClass_GetAllRecordsMockedCallback(sessionID)
}
// Return a map of VMPP references to VMPP records for all VMPPs known to the system.
func (_class VMPPClass) GetAllRecords(sessionID SessionRef) (_retval map[VMPPRef]VMPPRecord, _err error) {
	if (IsMock) {
		return _class.GetAllRecordsMock(sessionID)
	}	
	_method := "VMPP.get_all_records"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVMPPRefToVMPPRecordMapToGo(_method + " -> ", _result.Value)
	return
}


var VMPPClass_GetAllMockedCallback = func (sessionID SessionRef) (_retval []VMPPRef, _err error) {
	log.Println("VMPP.GetAll not mocked")
	_err = errors.New("VMPP.GetAll not mocked")
	return
}

func (_class VMPPClass) GetAllMock(sessionID SessionRef) (_retval []VMPPRef, _err error) {
	return VMPPClass_GetAllMockedCallback(sessionID)
}
// Return a list of all the VMPPs known to the system.
func (_class VMPPClass) GetAll(sessionID SessionRef) (_retval []VMPPRef, _err error) {
	if (IsMock) {
		return _class.GetAllMock(sessionID)
	}	
	_method := "VMPP.get_all"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVMPPRefSetToGo(_method + " -> ", _result.Value)
	return
}


var VMPPClass_SetArchiveLastRunTimeMockedCallback = func (sessionID SessionRef, self VMPPRef, value time.Time) (_err error) {
	log.Println("VMPP.SetArchiveLastRunTime not mocked")
	_err = errors.New("VMPP.SetArchiveLastRunTime not mocked")
	return
}

func (_class VMPPClass) SetArchiveLastRunTimeMock(sessionID SessionRef, self VMPPRef, value time.Time) (_err error) {
	return VMPPClass_SetArchiveLastRunTimeMockedCallback(sessionID, self, value)
}
// 
func (_class VMPPClass) SetArchiveLastRunTime(sessionID SessionRef, self VMPPRef, value time.Time) (_err error) {
	if (IsMock) {
		return _class.SetArchiveLastRunTimeMock(sessionID, self, value)
	}	
	_method := "VMPP.set_archive_last_run_time"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMPPRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var VMPPClass_SetBackupLastRunTimeMockedCallback = func (sessionID SessionRef, self VMPPRef, value time.Time) (_err error) {
	log.Println("VMPP.SetBackupLastRunTime not mocked")
	_err = errors.New("VMPP.SetBackupLastRunTime not mocked")
	return
}

func (_class VMPPClass) SetBackupLastRunTimeMock(sessionID SessionRef, self VMPPRef, value time.Time) (_err error) {
	return VMPPClass_SetBackupLastRunTimeMockedCallback(sessionID, self, value)
}
// 
func (_class VMPPClass) SetBackupLastRunTime(sessionID SessionRef, self VMPPRef, value time.Time) (_err error) {
	if (IsMock) {
		return _class.SetBackupLastRunTimeMock(sessionID, self, value)
	}	
	_method := "VMPP.set_backup_last_run_time"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMPPRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var VMPPClass_RemoveFromAlarmConfigMockedCallback = func (sessionID SessionRef, self VMPPRef, key string) (_err error) {
	log.Println("VMPP.RemoveFromAlarmConfig not mocked")
	_err = errors.New("VMPP.RemoveFromAlarmConfig not mocked")
	return
}

func (_class VMPPClass) RemoveFromAlarmConfigMock(sessionID SessionRef, self VMPPRef, key string) (_err error) {
	return VMPPClass_RemoveFromAlarmConfigMockedCallback(sessionID, self, key)
}
// 
func (_class VMPPClass) RemoveFromAlarmConfig(sessionID SessionRef, self VMPPRef, key string) (_err error) {
	if (IsMock) {
		return _class.RemoveFromAlarmConfigMock(sessionID, self, key)
	}	
	_method := "VMPP.remove_from_alarm_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMPPRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var VMPPClass_RemoveFromArchiveScheduleMockedCallback = func (sessionID SessionRef, self VMPPRef, key string) (_err error) {
	log.Println("VMPP.RemoveFromArchiveSchedule not mocked")
	_err = errors.New("VMPP.RemoveFromArchiveSchedule not mocked")
	return
}

func (_class VMPPClass) RemoveFromArchiveScheduleMock(sessionID SessionRef, self VMPPRef, key string) (_err error) {
	return VMPPClass_RemoveFromArchiveScheduleMockedCallback(sessionID, self, key)
}
// 
func (_class VMPPClass) RemoveFromArchiveSchedule(sessionID SessionRef, self VMPPRef, key string) (_err error) {
	if (IsMock) {
		return _class.RemoveFromArchiveScheduleMock(sessionID, self, key)
	}	
	_method := "VMPP.remove_from_archive_schedule"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMPPRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var VMPPClass_RemoveFromArchiveTargetConfigMockedCallback = func (sessionID SessionRef, self VMPPRef, key string) (_err error) {
	log.Println("VMPP.RemoveFromArchiveTargetConfig not mocked")
	_err = errors.New("VMPP.RemoveFromArchiveTargetConfig not mocked")
	return
}

func (_class VMPPClass) RemoveFromArchiveTargetConfigMock(sessionID SessionRef, self VMPPRef, key string) (_err error) {
	return VMPPClass_RemoveFromArchiveTargetConfigMockedCallback(sessionID, self, key)
}
// 
func (_class VMPPClass) RemoveFromArchiveTargetConfig(sessionID SessionRef, self VMPPRef, key string) (_err error) {
	if (IsMock) {
		return _class.RemoveFromArchiveTargetConfigMock(sessionID, self, key)
	}	
	_method := "VMPP.remove_from_archive_target_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMPPRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var VMPPClass_RemoveFromBackupScheduleMockedCallback = func (sessionID SessionRef, self VMPPRef, key string) (_err error) {
	log.Println("VMPP.RemoveFromBackupSchedule not mocked")
	_err = errors.New("VMPP.RemoveFromBackupSchedule not mocked")
	return
}

func (_class VMPPClass) RemoveFromBackupScheduleMock(sessionID SessionRef, self VMPPRef, key string) (_err error) {
	return VMPPClass_RemoveFromBackupScheduleMockedCallback(sessionID, self, key)
}
// 
func (_class VMPPClass) RemoveFromBackupSchedule(sessionID SessionRef, self VMPPRef, key string) (_err error) {
	if (IsMock) {
		return _class.RemoveFromBackupScheduleMock(sessionID, self, key)
	}	
	_method := "VMPP.remove_from_backup_schedule"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMPPRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var VMPPClass_AddToAlarmConfigMockedCallback = func (sessionID SessionRef, self VMPPRef, key string, value string) (_err error) {
	log.Println("VMPP.AddToAlarmConfig not mocked")
	_err = errors.New("VMPP.AddToAlarmConfig not mocked")
	return
}

func (_class VMPPClass) AddToAlarmConfigMock(sessionID SessionRef, self VMPPRef, key string, value string) (_err error) {
	return VMPPClass_AddToAlarmConfigMockedCallback(sessionID, self, key, value)
}
// 
func (_class VMPPClass) AddToAlarmConfig(sessionID SessionRef, self VMPPRef, key string, value string) (_err error) {
	if (IsMock) {
		return _class.AddToAlarmConfigMock(sessionID, self, key, value)
	}	
	_method := "VMPP.add_to_alarm_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMPPRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var VMPPClass_AddToArchiveScheduleMockedCallback = func (sessionID SessionRef, self VMPPRef, key string, value string) (_err error) {
	log.Println("VMPP.AddToArchiveSchedule not mocked")
	_err = errors.New("VMPP.AddToArchiveSchedule not mocked")
	return
}

func (_class VMPPClass) AddToArchiveScheduleMock(sessionID SessionRef, self VMPPRef, key string, value string) (_err error) {
	return VMPPClass_AddToArchiveScheduleMockedCallback(sessionID, self, key, value)
}
// 
func (_class VMPPClass) AddToArchiveSchedule(sessionID SessionRef, self VMPPRef, key string, value string) (_err error) {
	if (IsMock) {
		return _class.AddToArchiveScheduleMock(sessionID, self, key, value)
	}	
	_method := "VMPP.add_to_archive_schedule"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMPPRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var VMPPClass_AddToArchiveTargetConfigMockedCallback = func (sessionID SessionRef, self VMPPRef, key string, value string) (_err error) {
	log.Println("VMPP.AddToArchiveTargetConfig not mocked")
	_err = errors.New("VMPP.AddToArchiveTargetConfig not mocked")
	return
}

func (_class VMPPClass) AddToArchiveTargetConfigMock(sessionID SessionRef, self VMPPRef, key string, value string) (_err error) {
	return VMPPClass_AddToArchiveTargetConfigMockedCallback(sessionID, self, key, value)
}
// 
func (_class VMPPClass) AddToArchiveTargetConfig(sessionID SessionRef, self VMPPRef, key string, value string) (_err error) {
	if (IsMock) {
		return _class.AddToArchiveTargetConfigMock(sessionID, self, key, value)
	}	
	_method := "VMPP.add_to_archive_target_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMPPRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var VMPPClass_AddToBackupScheduleMockedCallback = func (sessionID SessionRef, self VMPPRef, key string, value string) (_err error) {
	log.Println("VMPP.AddToBackupSchedule not mocked")
	_err = errors.New("VMPP.AddToBackupSchedule not mocked")
	return
}

func (_class VMPPClass) AddToBackupScheduleMock(sessionID SessionRef, self VMPPRef, key string, value string) (_err error) {
	return VMPPClass_AddToBackupScheduleMockedCallback(sessionID, self, key, value)
}
// 
func (_class VMPPClass) AddToBackupSchedule(sessionID SessionRef, self VMPPRef, key string, value string) (_err error) {
	if (IsMock) {
		return _class.AddToBackupScheduleMock(sessionID, self, key, value)
	}	
	_method := "VMPP.add_to_backup_schedule"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMPPRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var VMPPClass_SetAlarmConfigMockedCallback = func (sessionID SessionRef, self VMPPRef, value map[string]string) (_err error) {
	log.Println("VMPP.SetAlarmConfig not mocked")
	_err = errors.New("VMPP.SetAlarmConfig not mocked")
	return
}

func (_class VMPPClass) SetAlarmConfigMock(sessionID SessionRef, self VMPPRef, value map[string]string) (_err error) {
	return VMPPClass_SetAlarmConfigMockedCallback(sessionID, self, value)
}
// 
func (_class VMPPClass) SetAlarmConfig(sessionID SessionRef, self VMPPRef, value map[string]string) (_err error) {
	if (IsMock) {
		return _class.SetAlarmConfigMock(sessionID, self, value)
	}	
	_method := "VMPP.set_alarm_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMPPRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var VMPPClass_SetIsAlarmEnabledMockedCallback = func (sessionID SessionRef, self VMPPRef, value bool) (_err error) {
	log.Println("VMPP.SetIsAlarmEnabled not mocked")
	_err = errors.New("VMPP.SetIsAlarmEnabled not mocked")
	return
}

func (_class VMPPClass) SetIsAlarmEnabledMock(sessionID SessionRef, self VMPPRef, value bool) (_err error) {
	return VMPPClass_SetIsAlarmEnabledMockedCallback(sessionID, self, value)
}
// Set the value of the is_alarm_enabled field
func (_class VMPPClass) SetIsAlarmEnabled(sessionID SessionRef, self VMPPRef, value bool) (_err error) {
	if (IsMock) {
		return _class.SetIsAlarmEnabledMock(sessionID, self, value)
	}	
	_method := "VMPP.set_is_alarm_enabled"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMPPRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var VMPPClass_SetArchiveTargetConfigMockedCallback = func (sessionID SessionRef, self VMPPRef, value map[string]string) (_err error) {
	log.Println("VMPP.SetArchiveTargetConfig not mocked")
	_err = errors.New("VMPP.SetArchiveTargetConfig not mocked")
	return
}

func (_class VMPPClass) SetArchiveTargetConfigMock(sessionID SessionRef, self VMPPRef, value map[string]string) (_err error) {
	return VMPPClass_SetArchiveTargetConfigMockedCallback(sessionID, self, value)
}
// 
func (_class VMPPClass) SetArchiveTargetConfig(sessionID SessionRef, self VMPPRef, value map[string]string) (_err error) {
	if (IsMock) {
		return _class.SetArchiveTargetConfigMock(sessionID, self, value)
	}	
	_method := "VMPP.set_archive_target_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMPPRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var VMPPClass_SetArchiveTargetTypeMockedCallback = func (sessionID SessionRef, self VMPPRef, value VmppArchiveTargetType) (_err error) {
	log.Println("VMPP.SetArchiveTargetType not mocked")
	_err = errors.New("VMPP.SetArchiveTargetType not mocked")
	return
}

func (_class VMPPClass) SetArchiveTargetTypeMock(sessionID SessionRef, self VMPPRef, value VmppArchiveTargetType) (_err error) {
	return VMPPClass_SetArchiveTargetTypeMockedCallback(sessionID, self, value)
}
// Set the value of the archive_target_config_type field
func (_class VMPPClass) SetArchiveTargetType(sessionID SessionRef, self VMPPRef, value VmppArchiveTargetType) (_err error) {
	if (IsMock) {
		return _class.SetArchiveTargetTypeMock(sessionID, self, value)
	}	
	_method := "VMPP.set_archive_target_type"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMPPRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertEnumVmppArchiveTargetTypeToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}


var VMPPClass_SetArchiveScheduleMockedCallback = func (sessionID SessionRef, self VMPPRef, value map[string]string) (_err error) {
	log.Println("VMPP.SetArchiveSchedule not mocked")
	_err = errors.New("VMPP.SetArchiveSchedule not mocked")
	return
}

func (_class VMPPClass) SetArchiveScheduleMock(sessionID SessionRef, self VMPPRef, value map[string]string) (_err error) {
	return VMPPClass_SetArchiveScheduleMockedCallback(sessionID, self, value)
}
// 
func (_class VMPPClass) SetArchiveSchedule(sessionID SessionRef, self VMPPRef, value map[string]string) (_err error) {
	if (IsMock) {
		return _class.SetArchiveScheduleMock(sessionID, self, value)
	}	
	_method := "VMPP.set_archive_schedule"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMPPRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var VMPPClass_SetArchiveFrequencyMockedCallback = func (sessionID SessionRef, self VMPPRef, value VmppArchiveFrequency) (_err error) {
	log.Println("VMPP.SetArchiveFrequency not mocked")
	_err = errors.New("VMPP.SetArchiveFrequency not mocked")
	return
}

func (_class VMPPClass) SetArchiveFrequencyMock(sessionID SessionRef, self VMPPRef, value VmppArchiveFrequency) (_err error) {
	return VMPPClass_SetArchiveFrequencyMockedCallback(sessionID, self, value)
}
// Set the value of the archive_frequency field
func (_class VMPPClass) SetArchiveFrequency(sessionID SessionRef, self VMPPRef, value VmppArchiveFrequency) (_err error) {
	if (IsMock) {
		return _class.SetArchiveFrequencyMock(sessionID, self, value)
	}	
	_method := "VMPP.set_archive_frequency"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMPPRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertEnumVmppArchiveFrequencyToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}


var VMPPClass_SetBackupScheduleMockedCallback = func (sessionID SessionRef, self VMPPRef, value map[string]string) (_err error) {
	log.Println("VMPP.SetBackupSchedule not mocked")
	_err = errors.New("VMPP.SetBackupSchedule not mocked")
	return
}

func (_class VMPPClass) SetBackupScheduleMock(sessionID SessionRef, self VMPPRef, value map[string]string) (_err error) {
	return VMPPClass_SetBackupScheduleMockedCallback(sessionID, self, value)
}
// 
func (_class VMPPClass) SetBackupSchedule(sessionID SessionRef, self VMPPRef, value map[string]string) (_err error) {
	if (IsMock) {
		return _class.SetBackupScheduleMock(sessionID, self, value)
	}	
	_method := "VMPP.set_backup_schedule"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMPPRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var VMPPClass_SetBackupFrequencyMockedCallback = func (sessionID SessionRef, self VMPPRef, value VmppBackupFrequency) (_err error) {
	log.Println("VMPP.SetBackupFrequency not mocked")
	_err = errors.New("VMPP.SetBackupFrequency not mocked")
	return
}

func (_class VMPPClass) SetBackupFrequencyMock(sessionID SessionRef, self VMPPRef, value VmppBackupFrequency) (_err error) {
	return VMPPClass_SetBackupFrequencyMockedCallback(sessionID, self, value)
}
// Set the value of the backup_frequency field
func (_class VMPPClass) SetBackupFrequency(sessionID SessionRef, self VMPPRef, value VmppBackupFrequency) (_err error) {
	if (IsMock) {
		return _class.SetBackupFrequencyMock(sessionID, self, value)
	}	
	_method := "VMPP.set_backup_frequency"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMPPRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertEnumVmppBackupFrequencyToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}


var VMPPClass_SetBackupRetentionValueMockedCallback = func (sessionID SessionRef, self VMPPRef, value int) (_err error) {
	log.Println("VMPP.SetBackupRetentionValue not mocked")
	_err = errors.New("VMPP.SetBackupRetentionValue not mocked")
	return
}

func (_class VMPPClass) SetBackupRetentionValueMock(sessionID SessionRef, self VMPPRef, value int) (_err error) {
	return VMPPClass_SetBackupRetentionValueMockedCallback(sessionID, self, value)
}
// 
func (_class VMPPClass) SetBackupRetentionValue(sessionID SessionRef, self VMPPRef, value int) (_err error) {
	if (IsMock) {
		return _class.SetBackupRetentionValueMock(sessionID, self, value)
	}	
	_method := "VMPP.set_backup_retention_value"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMPPRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var VMPPClass_GetAlertsMockedCallback = func (sessionID SessionRef, vmpp VMPPRef, hoursFromNow int) (_retval []string, _err error) {
	log.Println("VMPP.GetAlerts not mocked")
	_err = errors.New("VMPP.GetAlerts not mocked")
	return
}

func (_class VMPPClass) GetAlertsMock(sessionID SessionRef, vmpp VMPPRef, hoursFromNow int) (_retval []string, _err error) {
	return VMPPClass_GetAlertsMockedCallback(sessionID, vmpp, hoursFromNow)
}
// This call fetches a history of alerts for a given protection policy
func (_class VMPPClass) GetAlerts(sessionID SessionRef, vmpp VMPPRef, hoursFromNow int) (_retval []string, _err error) {
	if (IsMock) {
		return _class.GetAlertsMock(sessionID, vmpp, hoursFromNow)
	}	
	_method := "VMPP.get_alerts"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_vmppArg, _err := convertVMPPRefToXen(fmt.Sprintf("%s(%s)", _method, "vmpp"), vmpp)
	if _err != nil {
		return
	}
	_hoursFromNowArg, _err := convertIntToXen(fmt.Sprintf("%s(%s)", _method, "hours_from_now"), hoursFromNow)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _vmppArg, _hoursFromNowArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringSetToGo(_method + " -> ", _result.Value)
	return
}


var VMPPClass_ArchiveNowMockedCallback = func (sessionID SessionRef, snapshot VMRef) (_retval string, _err error) {
	log.Println("VMPP.ArchiveNow not mocked")
	_err = errors.New("VMPP.ArchiveNow not mocked")
	return
}

func (_class VMPPClass) ArchiveNowMock(sessionID SessionRef, snapshot VMRef) (_retval string, _err error) {
	return VMPPClass_ArchiveNowMockedCallback(sessionID, snapshot)
}
// This call archives the snapshot provided as a parameter
func (_class VMPPClass) ArchiveNow(sessionID SessionRef, snapshot VMRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.ArchiveNowMock(sessionID, snapshot)
	}	
	_method := "VMPP.archive_now"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_snapshotArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "snapshot"), snapshot)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _snapshotArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringToGo(_method + " -> ", _result.Value)
	return
}


var VMPPClass_ProtectNowMockedCallback = func (sessionID SessionRef, vmpp VMPPRef) (_retval string, _err error) {
	log.Println("VMPP.ProtectNow not mocked")
	_err = errors.New("VMPP.ProtectNow not mocked")
	return
}

func (_class VMPPClass) ProtectNowMock(sessionID SessionRef, vmpp VMPPRef) (_retval string, _err error) {
	return VMPPClass_ProtectNowMockedCallback(sessionID, vmpp)
}
// This call executes the protection policy immediately
func (_class VMPPClass) ProtectNow(sessionID SessionRef, vmpp VMPPRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.ProtectNowMock(sessionID, vmpp)
	}	
	_method := "VMPP.protect_now"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_vmppArg, _err := convertVMPPRefToXen(fmt.Sprintf("%s(%s)", _method, "vmpp"), vmpp)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _vmppArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringToGo(_method + " -> ", _result.Value)
	return
}


var VMPPClass_SetBackupTypeMockedCallback = func (sessionID SessionRef, self VMPPRef, value VmppBackupType) (_err error) {
	log.Println("VMPP.SetBackupType not mocked")
	_err = errors.New("VMPP.SetBackupType not mocked")
	return
}

func (_class VMPPClass) SetBackupTypeMock(sessionID SessionRef, self VMPPRef, value VmppBackupType) (_err error) {
	return VMPPClass_SetBackupTypeMockedCallback(sessionID, self, value)
}
// Set the backup_type field of the given VMPP.
func (_class VMPPClass) SetBackupType(sessionID SessionRef, self VMPPRef, value VmppBackupType) (_err error) {
	if (IsMock) {
		return _class.SetBackupTypeMock(sessionID, self, value)
	}	
	_method := "VMPP.set_backup_type"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMPPRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertEnumVmppBackupTypeToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}


var VMPPClass_SetIsPolicyEnabledMockedCallback = func (sessionID SessionRef, self VMPPRef, value bool) (_err error) {
	log.Println("VMPP.SetIsPolicyEnabled not mocked")
	_err = errors.New("VMPP.SetIsPolicyEnabled not mocked")
	return
}

func (_class VMPPClass) SetIsPolicyEnabledMock(sessionID SessionRef, self VMPPRef, value bool) (_err error) {
	return VMPPClass_SetIsPolicyEnabledMockedCallback(sessionID, self, value)
}
// Set the is_policy_enabled field of the given VMPP.
func (_class VMPPClass) SetIsPolicyEnabled(sessionID SessionRef, self VMPPRef, value bool) (_err error) {
	if (IsMock) {
		return _class.SetIsPolicyEnabledMock(sessionID, self, value)
	}	
	_method := "VMPP.set_is_policy_enabled"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMPPRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var VMPPClass_SetNameDescriptionMockedCallback = func (sessionID SessionRef, self VMPPRef, value string) (_err error) {
	log.Println("VMPP.SetNameDescription not mocked")
	_err = errors.New("VMPP.SetNameDescription not mocked")
	return
}

func (_class VMPPClass) SetNameDescriptionMock(sessionID SessionRef, self VMPPRef, value string) (_err error) {
	return VMPPClass_SetNameDescriptionMockedCallback(sessionID, self, value)
}
// Set the name/description field of the given VMPP.
func (_class VMPPClass) SetNameDescription(sessionID SessionRef, self VMPPRef, value string) (_err error) {
	if (IsMock) {
		return _class.SetNameDescriptionMock(sessionID, self, value)
	}	
	_method := "VMPP.set_name_description"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMPPRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var VMPPClass_SetNameLabelMockedCallback = func (sessionID SessionRef, self VMPPRef, value string) (_err error) {
	log.Println("VMPP.SetNameLabel not mocked")
	_err = errors.New("VMPP.SetNameLabel not mocked")
	return
}

func (_class VMPPClass) SetNameLabelMock(sessionID SessionRef, self VMPPRef, value string) (_err error) {
	return VMPPClass_SetNameLabelMockedCallback(sessionID, self, value)
}
// Set the name/label field of the given VMPP.
func (_class VMPPClass) SetNameLabel(sessionID SessionRef, self VMPPRef, value string) (_err error) {
	if (IsMock) {
		return _class.SetNameLabelMock(sessionID, self, value)
	}	
	_method := "VMPP.set_name_label"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMPPRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var VMPPClass_GetRecentAlertsMockedCallback = func (sessionID SessionRef, self VMPPRef) (_retval []string, _err error) {
	log.Println("VMPP.GetRecentAlerts not mocked")
	_err = errors.New("VMPP.GetRecentAlerts not mocked")
	return
}

func (_class VMPPClass) GetRecentAlertsMock(sessionID SessionRef, self VMPPRef) (_retval []string, _err error) {
	return VMPPClass_GetRecentAlertsMockedCallback(sessionID, self)
}
// Get the recent_alerts field of the given VMPP.
func (_class VMPPClass) GetRecentAlerts(sessionID SessionRef, self VMPPRef) (_retval []string, _err error) {
	if (IsMock) {
		return _class.GetRecentAlertsMock(sessionID, self)
	}	
	_method := "VMPP.get_recent_alerts"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMPPRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var VMPPClass_GetAlarmConfigMockedCallback = func (sessionID SessionRef, self VMPPRef) (_retval map[string]string, _err error) {
	log.Println("VMPP.GetAlarmConfig not mocked")
	_err = errors.New("VMPP.GetAlarmConfig not mocked")
	return
}

func (_class VMPPClass) GetAlarmConfigMock(sessionID SessionRef, self VMPPRef) (_retval map[string]string, _err error) {
	return VMPPClass_GetAlarmConfigMockedCallback(sessionID, self)
}
// Get the alarm_config field of the given VMPP.
func (_class VMPPClass) GetAlarmConfig(sessionID SessionRef, self VMPPRef) (_retval map[string]string, _err error) {
	if (IsMock) {
		return _class.GetAlarmConfigMock(sessionID, self)
	}	
	_method := "VMPP.get_alarm_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMPPRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var VMPPClass_GetIsAlarmEnabledMockedCallback = func (sessionID SessionRef, self VMPPRef) (_retval bool, _err error) {
	log.Println("VMPP.GetIsAlarmEnabled not mocked")
	_err = errors.New("VMPP.GetIsAlarmEnabled not mocked")
	return
}

func (_class VMPPClass) GetIsAlarmEnabledMock(sessionID SessionRef, self VMPPRef) (_retval bool, _err error) {
	return VMPPClass_GetIsAlarmEnabledMockedCallback(sessionID, self)
}
// Get the is_alarm_enabled field of the given VMPP.
func (_class VMPPClass) GetIsAlarmEnabled(sessionID SessionRef, self VMPPRef) (_retval bool, _err error) {
	if (IsMock) {
		return _class.GetIsAlarmEnabledMock(sessionID, self)
	}	
	_method := "VMPP.get_is_alarm_enabled"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMPPRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var VMPPClass_GetVMsMockedCallback = func (sessionID SessionRef, self VMPPRef) (_retval []VMRef, _err error) {
	log.Println("VMPP.GetVMs not mocked")
	_err = errors.New("VMPP.GetVMs not mocked")
	return
}

func (_class VMPPClass) GetVMsMock(sessionID SessionRef, self VMPPRef) (_retval []VMRef, _err error) {
	return VMPPClass_GetVMsMockedCallback(sessionID, self)
}
// Get the VMs field of the given VMPP.
func (_class VMPPClass) GetVMs(sessionID SessionRef, self VMPPRef) (_retval []VMRef, _err error) {
	if (IsMock) {
		return _class.GetVMsMock(sessionID, self)
	}	
	_method := "VMPP.get_VMs"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMPPRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var VMPPClass_GetArchiveLastRunTimeMockedCallback = func (sessionID SessionRef, self VMPPRef) (_retval time.Time, _err error) {
	log.Println("VMPP.GetArchiveLastRunTime not mocked")
	_err = errors.New("VMPP.GetArchiveLastRunTime not mocked")
	return
}

func (_class VMPPClass) GetArchiveLastRunTimeMock(sessionID SessionRef, self VMPPRef) (_retval time.Time, _err error) {
	return VMPPClass_GetArchiveLastRunTimeMockedCallback(sessionID, self)
}
// Get the archive_last_run_time field of the given VMPP.
func (_class VMPPClass) GetArchiveLastRunTime(sessionID SessionRef, self VMPPRef) (_retval time.Time, _err error) {
	if (IsMock) {
		return _class.GetArchiveLastRunTimeMock(sessionID, self)
	}	
	_method := "VMPP.get_archive_last_run_time"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMPPRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var VMPPClass_GetIsArchiveRunningMockedCallback = func (sessionID SessionRef, self VMPPRef) (_retval bool, _err error) {
	log.Println("VMPP.GetIsArchiveRunning not mocked")
	_err = errors.New("VMPP.GetIsArchiveRunning not mocked")
	return
}

func (_class VMPPClass) GetIsArchiveRunningMock(sessionID SessionRef, self VMPPRef) (_retval bool, _err error) {
	return VMPPClass_GetIsArchiveRunningMockedCallback(sessionID, self)
}
// Get the is_archive_running field of the given VMPP.
func (_class VMPPClass) GetIsArchiveRunning(sessionID SessionRef, self VMPPRef) (_retval bool, _err error) {
	if (IsMock) {
		return _class.GetIsArchiveRunningMock(sessionID, self)
	}	
	_method := "VMPP.get_is_archive_running"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMPPRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var VMPPClass_GetArchiveScheduleMockedCallback = func (sessionID SessionRef, self VMPPRef) (_retval map[string]string, _err error) {
	log.Println("VMPP.GetArchiveSchedule not mocked")
	_err = errors.New("VMPP.GetArchiveSchedule not mocked")
	return
}

func (_class VMPPClass) GetArchiveScheduleMock(sessionID SessionRef, self VMPPRef) (_retval map[string]string, _err error) {
	return VMPPClass_GetArchiveScheduleMockedCallback(sessionID, self)
}
// Get the archive_schedule field of the given VMPP.
func (_class VMPPClass) GetArchiveSchedule(sessionID SessionRef, self VMPPRef) (_retval map[string]string, _err error) {
	if (IsMock) {
		return _class.GetArchiveScheduleMock(sessionID, self)
	}	
	_method := "VMPP.get_archive_schedule"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMPPRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var VMPPClass_GetArchiveFrequencyMockedCallback = func (sessionID SessionRef, self VMPPRef) (_retval VmppArchiveFrequency, _err error) {
	log.Println("VMPP.GetArchiveFrequency not mocked")
	_err = errors.New("VMPP.GetArchiveFrequency not mocked")
	return
}

func (_class VMPPClass) GetArchiveFrequencyMock(sessionID SessionRef, self VMPPRef) (_retval VmppArchiveFrequency, _err error) {
	return VMPPClass_GetArchiveFrequencyMockedCallback(sessionID, self)
}
// Get the archive_frequency field of the given VMPP.
func (_class VMPPClass) GetArchiveFrequency(sessionID SessionRef, self VMPPRef) (_retval VmppArchiveFrequency, _err error) {
	if (IsMock) {
		return _class.GetArchiveFrequencyMock(sessionID, self)
	}	
	_method := "VMPP.get_archive_frequency"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMPPRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertEnumVmppArchiveFrequencyToGo(_method + " -> ", _result.Value)
	return
}


var VMPPClass_GetArchiveTargetConfigMockedCallback = func (sessionID SessionRef, self VMPPRef) (_retval map[string]string, _err error) {
	log.Println("VMPP.GetArchiveTargetConfig not mocked")
	_err = errors.New("VMPP.GetArchiveTargetConfig not mocked")
	return
}

func (_class VMPPClass) GetArchiveTargetConfigMock(sessionID SessionRef, self VMPPRef) (_retval map[string]string, _err error) {
	return VMPPClass_GetArchiveTargetConfigMockedCallback(sessionID, self)
}
// Get the archive_target_config field of the given VMPP.
func (_class VMPPClass) GetArchiveTargetConfig(sessionID SessionRef, self VMPPRef) (_retval map[string]string, _err error) {
	if (IsMock) {
		return _class.GetArchiveTargetConfigMock(sessionID, self)
	}	
	_method := "VMPP.get_archive_target_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMPPRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var VMPPClass_GetArchiveTargetTypeMockedCallback = func (sessionID SessionRef, self VMPPRef) (_retval VmppArchiveTargetType, _err error) {
	log.Println("VMPP.GetArchiveTargetType not mocked")
	_err = errors.New("VMPP.GetArchiveTargetType not mocked")
	return
}

func (_class VMPPClass) GetArchiveTargetTypeMock(sessionID SessionRef, self VMPPRef) (_retval VmppArchiveTargetType, _err error) {
	return VMPPClass_GetArchiveTargetTypeMockedCallback(sessionID, self)
}
// Get the archive_target_type field of the given VMPP.
func (_class VMPPClass) GetArchiveTargetType(sessionID SessionRef, self VMPPRef) (_retval VmppArchiveTargetType, _err error) {
	if (IsMock) {
		return _class.GetArchiveTargetTypeMock(sessionID, self)
	}	
	_method := "VMPP.get_archive_target_type"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMPPRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertEnumVmppArchiveTargetTypeToGo(_method + " -> ", _result.Value)
	return
}


var VMPPClass_GetBackupLastRunTimeMockedCallback = func (sessionID SessionRef, self VMPPRef) (_retval time.Time, _err error) {
	log.Println("VMPP.GetBackupLastRunTime not mocked")
	_err = errors.New("VMPP.GetBackupLastRunTime not mocked")
	return
}

func (_class VMPPClass) GetBackupLastRunTimeMock(sessionID SessionRef, self VMPPRef) (_retval time.Time, _err error) {
	return VMPPClass_GetBackupLastRunTimeMockedCallback(sessionID, self)
}
// Get the backup_last_run_time field of the given VMPP.
func (_class VMPPClass) GetBackupLastRunTime(sessionID SessionRef, self VMPPRef) (_retval time.Time, _err error) {
	if (IsMock) {
		return _class.GetBackupLastRunTimeMock(sessionID, self)
	}	
	_method := "VMPP.get_backup_last_run_time"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMPPRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var VMPPClass_GetIsBackupRunningMockedCallback = func (sessionID SessionRef, self VMPPRef) (_retval bool, _err error) {
	log.Println("VMPP.GetIsBackupRunning not mocked")
	_err = errors.New("VMPP.GetIsBackupRunning not mocked")
	return
}

func (_class VMPPClass) GetIsBackupRunningMock(sessionID SessionRef, self VMPPRef) (_retval bool, _err error) {
	return VMPPClass_GetIsBackupRunningMockedCallback(sessionID, self)
}
// Get the is_backup_running field of the given VMPP.
func (_class VMPPClass) GetIsBackupRunning(sessionID SessionRef, self VMPPRef) (_retval bool, _err error) {
	if (IsMock) {
		return _class.GetIsBackupRunningMock(sessionID, self)
	}	
	_method := "VMPP.get_is_backup_running"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMPPRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var VMPPClass_GetBackupScheduleMockedCallback = func (sessionID SessionRef, self VMPPRef) (_retval map[string]string, _err error) {
	log.Println("VMPP.GetBackupSchedule not mocked")
	_err = errors.New("VMPP.GetBackupSchedule not mocked")
	return
}

func (_class VMPPClass) GetBackupScheduleMock(sessionID SessionRef, self VMPPRef) (_retval map[string]string, _err error) {
	return VMPPClass_GetBackupScheduleMockedCallback(sessionID, self)
}
// Get the backup_schedule field of the given VMPP.
func (_class VMPPClass) GetBackupSchedule(sessionID SessionRef, self VMPPRef) (_retval map[string]string, _err error) {
	if (IsMock) {
		return _class.GetBackupScheduleMock(sessionID, self)
	}	
	_method := "VMPP.get_backup_schedule"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMPPRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var VMPPClass_GetBackupFrequencyMockedCallback = func (sessionID SessionRef, self VMPPRef) (_retval VmppBackupFrequency, _err error) {
	log.Println("VMPP.GetBackupFrequency not mocked")
	_err = errors.New("VMPP.GetBackupFrequency not mocked")
	return
}

func (_class VMPPClass) GetBackupFrequencyMock(sessionID SessionRef, self VMPPRef) (_retval VmppBackupFrequency, _err error) {
	return VMPPClass_GetBackupFrequencyMockedCallback(sessionID, self)
}
// Get the backup_frequency field of the given VMPP.
func (_class VMPPClass) GetBackupFrequency(sessionID SessionRef, self VMPPRef) (_retval VmppBackupFrequency, _err error) {
	if (IsMock) {
		return _class.GetBackupFrequencyMock(sessionID, self)
	}	
	_method := "VMPP.get_backup_frequency"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMPPRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertEnumVmppBackupFrequencyToGo(_method + " -> ", _result.Value)
	return
}


var VMPPClass_GetBackupRetentionValueMockedCallback = func (sessionID SessionRef, self VMPPRef) (_retval int, _err error) {
	log.Println("VMPP.GetBackupRetentionValue not mocked")
	_err = errors.New("VMPP.GetBackupRetentionValue not mocked")
	return
}

func (_class VMPPClass) GetBackupRetentionValueMock(sessionID SessionRef, self VMPPRef) (_retval int, _err error) {
	return VMPPClass_GetBackupRetentionValueMockedCallback(sessionID, self)
}
// Get the backup_retention_value field of the given VMPP.
func (_class VMPPClass) GetBackupRetentionValue(sessionID SessionRef, self VMPPRef) (_retval int, _err error) {
	if (IsMock) {
		return _class.GetBackupRetentionValueMock(sessionID, self)
	}	
	_method := "VMPP.get_backup_retention_value"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMPPRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var VMPPClass_GetBackupTypeMockedCallback = func (sessionID SessionRef, self VMPPRef) (_retval VmppBackupType, _err error) {
	log.Println("VMPP.GetBackupType not mocked")
	_err = errors.New("VMPP.GetBackupType not mocked")
	return
}

func (_class VMPPClass) GetBackupTypeMock(sessionID SessionRef, self VMPPRef) (_retval VmppBackupType, _err error) {
	return VMPPClass_GetBackupTypeMockedCallback(sessionID, self)
}
// Get the backup_type field of the given VMPP.
func (_class VMPPClass) GetBackupType(sessionID SessionRef, self VMPPRef) (_retval VmppBackupType, _err error) {
	if (IsMock) {
		return _class.GetBackupTypeMock(sessionID, self)
	}	
	_method := "VMPP.get_backup_type"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMPPRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertEnumVmppBackupTypeToGo(_method + " -> ", _result.Value)
	return
}


var VMPPClass_GetIsPolicyEnabledMockedCallback = func (sessionID SessionRef, self VMPPRef) (_retval bool, _err error) {
	log.Println("VMPP.GetIsPolicyEnabled not mocked")
	_err = errors.New("VMPP.GetIsPolicyEnabled not mocked")
	return
}

func (_class VMPPClass) GetIsPolicyEnabledMock(sessionID SessionRef, self VMPPRef) (_retval bool, _err error) {
	return VMPPClass_GetIsPolicyEnabledMockedCallback(sessionID, self)
}
// Get the is_policy_enabled field of the given VMPP.
func (_class VMPPClass) GetIsPolicyEnabled(sessionID SessionRef, self VMPPRef) (_retval bool, _err error) {
	if (IsMock) {
		return _class.GetIsPolicyEnabledMock(sessionID, self)
	}	
	_method := "VMPP.get_is_policy_enabled"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMPPRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var VMPPClass_GetNameDescriptionMockedCallback = func (sessionID SessionRef, self VMPPRef) (_retval string, _err error) {
	log.Println("VMPP.GetNameDescription not mocked")
	_err = errors.New("VMPP.GetNameDescription not mocked")
	return
}

func (_class VMPPClass) GetNameDescriptionMock(sessionID SessionRef, self VMPPRef) (_retval string, _err error) {
	return VMPPClass_GetNameDescriptionMockedCallback(sessionID, self)
}
// Get the name/description field of the given VMPP.
func (_class VMPPClass) GetNameDescription(sessionID SessionRef, self VMPPRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetNameDescriptionMock(sessionID, self)
	}	
	_method := "VMPP.get_name_description"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMPPRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var VMPPClass_GetNameLabelMockedCallback = func (sessionID SessionRef, self VMPPRef) (_retval string, _err error) {
	log.Println("VMPP.GetNameLabel not mocked")
	_err = errors.New("VMPP.GetNameLabel not mocked")
	return
}

func (_class VMPPClass) GetNameLabelMock(sessionID SessionRef, self VMPPRef) (_retval string, _err error) {
	return VMPPClass_GetNameLabelMockedCallback(sessionID, self)
}
// Get the name/label field of the given VMPP.
func (_class VMPPClass) GetNameLabel(sessionID SessionRef, self VMPPRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetNameLabelMock(sessionID, self)
	}	
	_method := "VMPP.get_name_label"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMPPRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var VMPPClass_GetUUIDMockedCallback = func (sessionID SessionRef, self VMPPRef) (_retval string, _err error) {
	log.Println("VMPP.GetUUID not mocked")
	_err = errors.New("VMPP.GetUUID not mocked")
	return
}

func (_class VMPPClass) GetUUIDMock(sessionID SessionRef, self VMPPRef) (_retval string, _err error) {
	return VMPPClass_GetUUIDMockedCallback(sessionID, self)
}
// Get the uuid field of the given VMPP.
func (_class VMPPClass) GetUUID(sessionID SessionRef, self VMPPRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetUUIDMock(sessionID, self)
	}	
	_method := "VMPP.get_uuid"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMPPRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var VMPPClass_GetByNameLabelMockedCallback = func (sessionID SessionRef, label string) (_retval []VMPPRef, _err error) {
	log.Println("VMPP.GetByNameLabel not mocked")
	_err = errors.New("VMPP.GetByNameLabel not mocked")
	return
}

func (_class VMPPClass) GetByNameLabelMock(sessionID SessionRef, label string) (_retval []VMPPRef, _err error) {
	return VMPPClass_GetByNameLabelMockedCallback(sessionID, label)
}
// Get all the VMPP instances with the given label.
func (_class VMPPClass) GetByNameLabel(sessionID SessionRef, label string) (_retval []VMPPRef, _err error) {
	if (IsMock) {
		return _class.GetByNameLabelMock(sessionID, label)
	}	
	_method := "VMPP.get_by_name_label"
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
	_retval, _err = convertVMPPRefSetToGo(_method + " -> ", _result.Value)
	return
}


var VMPPClass_DestroyMockedCallback = func (sessionID SessionRef, self VMPPRef) (_err error) {
	log.Println("VMPP.Destroy not mocked")
	_err = errors.New("VMPP.Destroy not mocked")
	return
}

func (_class VMPPClass) DestroyMock(sessionID SessionRef, self VMPPRef) (_err error) {
	return VMPPClass_DestroyMockedCallback(sessionID, self)
}
// Destroy the specified VMPP instance.
func (_class VMPPClass) Destroy(sessionID SessionRef, self VMPPRef) (_err error) {
	if (IsMock) {
		return _class.DestroyMock(sessionID, self)
	}	
	_method := "VMPP.destroy"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMPPRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg)
	return
}


var VMPPClass_CreateMockedCallback = func (sessionID SessionRef, args VMPPRecord) (_retval VMPPRef, _err error) {
	log.Println("VMPP.Create not mocked")
	_err = errors.New("VMPP.Create not mocked")
	return
}

func (_class VMPPClass) CreateMock(sessionID SessionRef, args VMPPRecord) (_retval VMPPRef, _err error) {
	return VMPPClass_CreateMockedCallback(sessionID, args)
}
// Create a new VMPP instance, and return its handle.
// The constructor args are: name_label, name_description, is_policy_enabled, backup_type, backup_retention_value, backup_frequency, backup_schedule, archive_target_type, archive_target_config, archive_frequency, archive_schedule, is_alarm_enabled, alarm_config (* = non-optional).
func (_class VMPPClass) Create(sessionID SessionRef, args VMPPRecord) (_retval VMPPRef, _err error) {
	if (IsMock) {
		return _class.CreateMock(sessionID, args)
	}	
	_method := "VMPP.create"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_argsArg, _err := convertVMPPRecordToXen(fmt.Sprintf("%s(%s)", _method, "args"), args)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _argsArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVMPPRefToGo(_method + " -> ", _result.Value)
	return
}


var VMPPClass_GetByUUIDMockedCallback = func (sessionID SessionRef, uuid string) (_retval VMPPRef, _err error) {
	log.Println("VMPP.GetByUUID not mocked")
	_err = errors.New("VMPP.GetByUUID not mocked")
	return
}

func (_class VMPPClass) GetByUUIDMock(sessionID SessionRef, uuid string) (_retval VMPPRef, _err error) {
	return VMPPClass_GetByUUIDMockedCallback(sessionID, uuid)
}
// Get a reference to the VMPP instance with the specified UUID.
func (_class VMPPClass) GetByUUID(sessionID SessionRef, uuid string) (_retval VMPPRef, _err error) {
	if (IsMock) {
		return _class.GetByUUIDMock(sessionID, uuid)
	}	
	_method := "VMPP.get_by_uuid"
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
	_retval, _err = convertVMPPRefToGo(_method + " -> ", _result.Value)
	return
}


var VMPPClass_GetRecordMockedCallback = func (sessionID SessionRef, self VMPPRef) (_retval VMPPRecord, _err error) {
	log.Println("VMPP.GetRecord not mocked")
	_err = errors.New("VMPP.GetRecord not mocked")
	return
}

func (_class VMPPClass) GetRecordMock(sessionID SessionRef, self VMPPRef) (_retval VMPPRecord, _err error) {
	return VMPPClass_GetRecordMockedCallback(sessionID, self)
}
// Get a record containing the current state of the given VMPP.
func (_class VMPPClass) GetRecord(sessionID SessionRef, self VMPPRef) (_retval VMPPRecord, _err error) {
	if (IsMock) {
		return _class.GetRecordMock(sessionID, self)
	}	
	_method := "VMPP.get_record"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMPPRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVMPPRecordToGo(_method + " -> ", _result.Value)
	return
}
