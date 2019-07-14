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

type TaskAllowedOperations string

const (
  // refers to the operation "cancel"
	TaskAllowedOperationsCancel TaskAllowedOperations = "cancel"
  // refers to the operation "destroy"
	TaskAllowedOperationsDestroy TaskAllowedOperations = "destroy"
)

type TaskStatusType string

const (
  // task is in progress
	TaskStatusTypePending TaskStatusType = "pending"
  // task was completed successfully
	TaskStatusTypeSuccess TaskStatusType = "success"
  // task has failed
	TaskStatusTypeFailure TaskStatusType = "failure"
  // task is being cancelled
	TaskStatusTypeCancelling TaskStatusType = "cancelling"
  // task has been cancelled
	TaskStatusTypeCancelled TaskStatusType = "cancelled"
)

type TaskRecord struct {
  // Unique identifier/object reference
	UUID string
  // a human-readable name
	NameLabel string
  // a notes field containing human-readable description
	NameDescription string
  // list of the operations allowed in this state. This list is advisory only and the server state may have changed by the time this field is read by a client.
	AllowedOperations []TaskAllowedOperations
  // links each of the running tasks using this object (by reference) to a current_operation enum which describes the nature of the task.
	CurrentOperations map[string]TaskAllowedOperations
  // Time task was created
	Created time.Time
  // Time task finished (i.e. succeeded or failed). If task-status is pending, then the value of this field has no meaning
	Finished time.Time
  // current status of the task
	Status TaskStatusType
  // the host on which the task is running
	ResidentOn HostRef
  // This field contains the estimated fraction of the task which is complete. This field should not be used to determine whether the task is complete - for this the status field of the task should be used.
	Progress float64
  // if the task has completed successfully, this field contains the type of the encoded result (i.e. name of the class whose reference is in the result field). Undefined otherwise.
	Type string
  // if the task has completed successfully, this field contains the result value (either Void or an object reference). Undefined otherwise.
	Result string
  // if the task has failed, this field contains the set of associated error strings. Undefined otherwise.
	ErrorInfo []string
  // additional configuration
	OtherConfig map[string]string
  // Ref pointing to the task this is a substask of.
	SubtaskOf TaskRef
  // List pointing to all the substasks.
	Subtasks []TaskRef
  // Function call trace for debugging.
	Backtrace string
}

type TaskRef string

// A long-running asynchronous task
type TaskClass struct {
	client *Client
}


var TaskClass_GetAllRecordsMockedCallback = func (sessionID SessionRef) (_retval map[TaskRef]TaskRecord, _err error) {
	log.Println("Task.GetAllRecords not mocked")
	_err = errors.New("Task.GetAllRecords not mocked")
	return
}

func (_class TaskClass) GetAllRecordsMock(sessionID SessionRef) (_retval map[TaskRef]TaskRecord, _err error) {
	return TaskClass_GetAllRecordsMockedCallback(sessionID)
}
// Return a map of task references to task records for all tasks known to the system.
func (_class TaskClass) GetAllRecords(sessionID SessionRef) (_retval map[TaskRef]TaskRecord, _err error) {
	if (IsMock) {
		return _class.GetAllRecordsMock(sessionID)
	}	
	_method := "task.get_all_records"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToTaskRecordMapToGo(_method + " -> ", _result.Value)
	return
}


var TaskClass_GetAllMockedCallback = func (sessionID SessionRef) (_retval []TaskRef, _err error) {
	log.Println("Task.GetAll not mocked")
	_err = errors.New("Task.GetAll not mocked")
	return
}

func (_class TaskClass) GetAllMock(sessionID SessionRef) (_retval []TaskRef, _err error) {
	return TaskClass_GetAllMockedCallback(sessionID)
}
// Return a list of all the tasks known to the system.
func (_class TaskClass) GetAll(sessionID SessionRef) (_retval []TaskRef, _err error) {
	if (IsMock) {
		return _class.GetAllMock(sessionID)
	}	
	_method := "task.get_all"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefSetToGo(_method + " -> ", _result.Value)
	return
}


var TaskClass_SetStatusMockedCallback = func (sessionID SessionRef, self TaskRef, value TaskStatusType) (_err error) {
	log.Println("Task.SetStatus not mocked")
	_err = errors.New("Task.SetStatus not mocked")
	return
}

func (_class TaskClass) SetStatusMock(sessionID SessionRef, self TaskRef, value TaskStatusType) (_err error) {
	return TaskClass_SetStatusMockedCallback(sessionID, self, value)
}
// Set the task status
func (_class TaskClass) SetStatus(sessionID SessionRef, self TaskRef, value TaskStatusType) (_err error) {
	if (IsMock) {
		return _class.SetStatusMock(sessionID, self, value)
	}	
	_method := "task.set_status"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertTaskRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertEnumTaskStatusTypeToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}


var TaskClass_CancelMockedCallback = func (sessionID SessionRef, task TaskRef) (_err error) {
	log.Println("Task.Cancel not mocked")
	_err = errors.New("Task.Cancel not mocked")
	return
}

func (_class TaskClass) CancelMock(sessionID SessionRef, task TaskRef) (_err error) {
	return TaskClass_CancelMockedCallback(sessionID, task)
}
// Request that a task be cancelled. Note that a task may fail to be cancelled and may complete or fail normally and note that, even when a task does cancel, it might take an arbitrary amount of time.
//
// Errors:
//  OPERATION_NOT_ALLOWED - You attempted an operation that was not allowed.
func (_class TaskClass) Cancel(sessionID SessionRef, task TaskRef) (_err error) {
	if (IsMock) {
		return _class.CancelMock(sessionID, task)
	}	
	_method := "task.cancel"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_taskArg, _err := convertTaskRefToXen(fmt.Sprintf("%s(%s)", _method, "task"), task)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _taskArg)
	return
}


var TaskClass_DestroyMockedCallback = func (sessionID SessionRef, self TaskRef) (_err error) {
	log.Println("Task.Destroy not mocked")
	_err = errors.New("Task.Destroy not mocked")
	return
}

func (_class TaskClass) DestroyMock(sessionID SessionRef, self TaskRef) (_err error) {
	return TaskClass_DestroyMockedCallback(sessionID, self)
}
// Destroy the task object
func (_class TaskClass) Destroy(sessionID SessionRef, self TaskRef) (_err error) {
	if (IsMock) {
		return _class.DestroyMock(sessionID, self)
	}	
	_method := "task.destroy"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertTaskRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg)
	return
}


var TaskClass_CreateMockedCallback = func (sessionID SessionRef, label string, description string) (_retval TaskRef, _err error) {
	log.Println("Task.Create not mocked")
	_err = errors.New("Task.Create not mocked")
	return
}

func (_class TaskClass) CreateMock(sessionID SessionRef, label string, description string) (_retval TaskRef, _err error) {
	return TaskClass_CreateMockedCallback(sessionID, label, description)
}
// Create a new task object which must be manually destroyed.
func (_class TaskClass) Create(sessionID SessionRef, label string, description string) (_retval TaskRef, _err error) {
	if (IsMock) {
		return _class.CreateMock(sessionID, label, description)
	}	
	_method := "task.create"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_labelArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "label"), label)
	if _err != nil {
		return
	}
	_descriptionArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "description"), description)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _labelArg, _descriptionArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}


var TaskClass_RemoveFromOtherConfigMockedCallback = func (sessionID SessionRef, self TaskRef, key string) (_err error) {
	log.Println("Task.RemoveFromOtherConfig not mocked")
	_err = errors.New("Task.RemoveFromOtherConfig not mocked")
	return
}

func (_class TaskClass) RemoveFromOtherConfigMock(sessionID SessionRef, self TaskRef, key string) (_err error) {
	return TaskClass_RemoveFromOtherConfigMockedCallback(sessionID, self, key)
}
// Remove the given key and its corresponding value from the other_config field of the given task.  If the key is not in that Map, then do nothing.
func (_class TaskClass) RemoveFromOtherConfig(sessionID SessionRef, self TaskRef, key string) (_err error) {
	if (IsMock) {
		return _class.RemoveFromOtherConfigMock(sessionID, self, key)
	}	
	_method := "task.remove_from_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertTaskRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var TaskClass_AddToOtherConfigMockedCallback = func (sessionID SessionRef, self TaskRef, key string, value string) (_err error) {
	log.Println("Task.AddToOtherConfig not mocked")
	_err = errors.New("Task.AddToOtherConfig not mocked")
	return
}

func (_class TaskClass) AddToOtherConfigMock(sessionID SessionRef, self TaskRef, key string, value string) (_err error) {
	return TaskClass_AddToOtherConfigMockedCallback(sessionID, self, key, value)
}
// Add the given key-value pair to the other_config field of the given task.
func (_class TaskClass) AddToOtherConfig(sessionID SessionRef, self TaskRef, key string, value string) (_err error) {
	if (IsMock) {
		return _class.AddToOtherConfigMock(sessionID, self, key, value)
	}	
	_method := "task.add_to_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertTaskRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var TaskClass_SetOtherConfigMockedCallback = func (sessionID SessionRef, self TaskRef, value map[string]string) (_err error) {
	log.Println("Task.SetOtherConfig not mocked")
	_err = errors.New("Task.SetOtherConfig not mocked")
	return
}

func (_class TaskClass) SetOtherConfigMock(sessionID SessionRef, self TaskRef, value map[string]string) (_err error) {
	return TaskClass_SetOtherConfigMockedCallback(sessionID, self, value)
}
// Set the other_config field of the given task.
func (_class TaskClass) SetOtherConfig(sessionID SessionRef, self TaskRef, value map[string]string) (_err error) {
	if (IsMock) {
		return _class.SetOtherConfigMock(sessionID, self, value)
	}	
	_method := "task.set_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertTaskRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var TaskClass_GetBacktraceMockedCallback = func (sessionID SessionRef, self TaskRef) (_retval string, _err error) {
	log.Println("Task.GetBacktrace not mocked")
	_err = errors.New("Task.GetBacktrace not mocked")
	return
}

func (_class TaskClass) GetBacktraceMock(sessionID SessionRef, self TaskRef) (_retval string, _err error) {
	return TaskClass_GetBacktraceMockedCallback(sessionID, self)
}
// Get the backtrace field of the given task.
func (_class TaskClass) GetBacktrace(sessionID SessionRef, self TaskRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetBacktraceMock(sessionID, self)
	}	
	_method := "task.get_backtrace"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertTaskRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var TaskClass_GetSubtasksMockedCallback = func (sessionID SessionRef, self TaskRef) (_retval []TaskRef, _err error) {
	log.Println("Task.GetSubtasks not mocked")
	_err = errors.New("Task.GetSubtasks not mocked")
	return
}

func (_class TaskClass) GetSubtasksMock(sessionID SessionRef, self TaskRef) (_retval []TaskRef, _err error) {
	return TaskClass_GetSubtasksMockedCallback(sessionID, self)
}
// Get the subtasks field of the given task.
func (_class TaskClass) GetSubtasks(sessionID SessionRef, self TaskRef) (_retval []TaskRef, _err error) {
	if (IsMock) {
		return _class.GetSubtasksMock(sessionID, self)
	}	
	_method := "task.get_subtasks"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertTaskRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefSetToGo(_method + " -> ", _result.Value)
	return
}


var TaskClass_GetSubtaskOfMockedCallback = func (sessionID SessionRef, self TaskRef) (_retval TaskRef, _err error) {
	log.Println("Task.GetSubtaskOf not mocked")
	_err = errors.New("Task.GetSubtaskOf not mocked")
	return
}

func (_class TaskClass) GetSubtaskOfMock(sessionID SessionRef, self TaskRef) (_retval TaskRef, _err error) {
	return TaskClass_GetSubtaskOfMockedCallback(sessionID, self)
}
// Get the subtask_of field of the given task.
func (_class TaskClass) GetSubtaskOf(sessionID SessionRef, self TaskRef) (_retval TaskRef, _err error) {
	if (IsMock) {
		return _class.GetSubtaskOfMock(sessionID, self)
	}	
	_method := "task.get_subtask_of"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertTaskRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}


var TaskClass_GetOtherConfigMockedCallback = func (sessionID SessionRef, self TaskRef) (_retval map[string]string, _err error) {
	log.Println("Task.GetOtherConfig not mocked")
	_err = errors.New("Task.GetOtherConfig not mocked")
	return
}

func (_class TaskClass) GetOtherConfigMock(sessionID SessionRef, self TaskRef) (_retval map[string]string, _err error) {
	return TaskClass_GetOtherConfigMockedCallback(sessionID, self)
}
// Get the other_config field of the given task.
func (_class TaskClass) GetOtherConfig(sessionID SessionRef, self TaskRef) (_retval map[string]string, _err error) {
	if (IsMock) {
		return _class.GetOtherConfigMock(sessionID, self)
	}	
	_method := "task.get_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertTaskRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var TaskClass_GetErrorInfoMockedCallback = func (sessionID SessionRef, self TaskRef) (_retval []string, _err error) {
	log.Println("Task.GetErrorInfo not mocked")
	_err = errors.New("Task.GetErrorInfo not mocked")
	return
}

func (_class TaskClass) GetErrorInfoMock(sessionID SessionRef, self TaskRef) (_retval []string, _err error) {
	return TaskClass_GetErrorInfoMockedCallback(sessionID, self)
}
// Get the error_info field of the given task.
func (_class TaskClass) GetErrorInfo(sessionID SessionRef, self TaskRef) (_retval []string, _err error) {
	if (IsMock) {
		return _class.GetErrorInfoMock(sessionID, self)
	}	
	_method := "task.get_error_info"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertTaskRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var TaskClass_GetResultMockedCallback = func (sessionID SessionRef, self TaskRef) (_retval string, _err error) {
	log.Println("Task.GetResult not mocked")
	_err = errors.New("Task.GetResult not mocked")
	return
}

func (_class TaskClass) GetResultMock(sessionID SessionRef, self TaskRef) (_retval string, _err error) {
	return TaskClass_GetResultMockedCallback(sessionID, self)
}
// Get the result field of the given task.
func (_class TaskClass) GetResult(sessionID SessionRef, self TaskRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetResultMock(sessionID, self)
	}	
	_method := "task.get_result"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertTaskRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var TaskClass_GetTypeMockedCallback = func (sessionID SessionRef, self TaskRef) (_retval string, _err error) {
	log.Println("Task.GetType not mocked")
	_err = errors.New("Task.GetType not mocked")
	return
}

func (_class TaskClass) GetTypeMock(sessionID SessionRef, self TaskRef) (_retval string, _err error) {
	return TaskClass_GetTypeMockedCallback(sessionID, self)
}
// Get the type field of the given task.
func (_class TaskClass) GetType(sessionID SessionRef, self TaskRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetTypeMock(sessionID, self)
	}	
	_method := "task.get_type"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertTaskRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var TaskClass_GetProgressMockedCallback = func (sessionID SessionRef, self TaskRef) (_retval float64, _err error) {
	log.Println("Task.GetProgress not mocked")
	_err = errors.New("Task.GetProgress not mocked")
	return
}

func (_class TaskClass) GetProgressMock(sessionID SessionRef, self TaskRef) (_retval float64, _err error) {
	return TaskClass_GetProgressMockedCallback(sessionID, self)
}
// Get the progress field of the given task.
func (_class TaskClass) GetProgress(sessionID SessionRef, self TaskRef) (_retval float64, _err error) {
	if (IsMock) {
		return _class.GetProgressMock(sessionID, self)
	}	
	_method := "task.get_progress"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertTaskRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertFloatToGo(_method + " -> ", _result.Value)
	return
}


var TaskClass_GetResidentOnMockedCallback = func (sessionID SessionRef, self TaskRef) (_retval HostRef, _err error) {
	log.Println("Task.GetResidentOn not mocked")
	_err = errors.New("Task.GetResidentOn not mocked")
	return
}

func (_class TaskClass) GetResidentOnMock(sessionID SessionRef, self TaskRef) (_retval HostRef, _err error) {
	return TaskClass_GetResidentOnMockedCallback(sessionID, self)
}
// Get the resident_on field of the given task.
func (_class TaskClass) GetResidentOn(sessionID SessionRef, self TaskRef) (_retval HostRef, _err error) {
	if (IsMock) {
		return _class.GetResidentOnMock(sessionID, self)
	}	
	_method := "task.get_resident_on"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertTaskRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var TaskClass_GetStatusMockedCallback = func (sessionID SessionRef, self TaskRef) (_retval TaskStatusType, _err error) {
	log.Println("Task.GetStatus not mocked")
	_err = errors.New("Task.GetStatus not mocked")
	return
}

func (_class TaskClass) GetStatusMock(sessionID SessionRef, self TaskRef) (_retval TaskStatusType, _err error) {
	return TaskClass_GetStatusMockedCallback(sessionID, self)
}
// Get the status field of the given task.
func (_class TaskClass) GetStatus(sessionID SessionRef, self TaskRef) (_retval TaskStatusType, _err error) {
	if (IsMock) {
		return _class.GetStatusMock(sessionID, self)
	}	
	_method := "task.get_status"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertTaskRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertEnumTaskStatusTypeToGo(_method + " -> ", _result.Value)
	return
}


var TaskClass_GetFinishedMockedCallback = func (sessionID SessionRef, self TaskRef) (_retval time.Time, _err error) {
	log.Println("Task.GetFinished not mocked")
	_err = errors.New("Task.GetFinished not mocked")
	return
}

func (_class TaskClass) GetFinishedMock(sessionID SessionRef, self TaskRef) (_retval time.Time, _err error) {
	return TaskClass_GetFinishedMockedCallback(sessionID, self)
}
// Get the finished field of the given task.
func (_class TaskClass) GetFinished(sessionID SessionRef, self TaskRef) (_retval time.Time, _err error) {
	if (IsMock) {
		return _class.GetFinishedMock(sessionID, self)
	}	
	_method := "task.get_finished"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertTaskRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var TaskClass_GetCreatedMockedCallback = func (sessionID SessionRef, self TaskRef) (_retval time.Time, _err error) {
	log.Println("Task.GetCreated not mocked")
	_err = errors.New("Task.GetCreated not mocked")
	return
}

func (_class TaskClass) GetCreatedMock(sessionID SessionRef, self TaskRef) (_retval time.Time, _err error) {
	return TaskClass_GetCreatedMockedCallback(sessionID, self)
}
// Get the created field of the given task.
func (_class TaskClass) GetCreated(sessionID SessionRef, self TaskRef) (_retval time.Time, _err error) {
	if (IsMock) {
		return _class.GetCreatedMock(sessionID, self)
	}	
	_method := "task.get_created"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertTaskRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var TaskClass_GetCurrentOperationsMockedCallback = func (sessionID SessionRef, self TaskRef) (_retval map[string]TaskAllowedOperations, _err error) {
	log.Println("Task.GetCurrentOperations not mocked")
	_err = errors.New("Task.GetCurrentOperations not mocked")
	return
}

func (_class TaskClass) GetCurrentOperationsMock(sessionID SessionRef, self TaskRef) (_retval map[string]TaskAllowedOperations, _err error) {
	return TaskClass_GetCurrentOperationsMockedCallback(sessionID, self)
}
// Get the current_operations field of the given task.
func (_class TaskClass) GetCurrentOperations(sessionID SessionRef, self TaskRef) (_retval map[string]TaskAllowedOperations, _err error) {
	if (IsMock) {
		return _class.GetCurrentOperationsMock(sessionID, self)
	}	
	_method := "task.get_current_operations"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertTaskRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringToEnumTaskAllowedOperationsMapToGo(_method + " -> ", _result.Value)
	return
}


var TaskClass_GetAllowedOperationsMockedCallback = func (sessionID SessionRef, self TaskRef) (_retval []TaskAllowedOperations, _err error) {
	log.Println("Task.GetAllowedOperations not mocked")
	_err = errors.New("Task.GetAllowedOperations not mocked")
	return
}

func (_class TaskClass) GetAllowedOperationsMock(sessionID SessionRef, self TaskRef) (_retval []TaskAllowedOperations, _err error) {
	return TaskClass_GetAllowedOperationsMockedCallback(sessionID, self)
}
// Get the allowed_operations field of the given task.
func (_class TaskClass) GetAllowedOperations(sessionID SessionRef, self TaskRef) (_retval []TaskAllowedOperations, _err error) {
	if (IsMock) {
		return _class.GetAllowedOperationsMock(sessionID, self)
	}	
	_method := "task.get_allowed_operations"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertTaskRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertEnumTaskAllowedOperationsSetToGo(_method + " -> ", _result.Value)
	return
}


var TaskClass_GetNameDescriptionMockedCallback = func (sessionID SessionRef, self TaskRef) (_retval string, _err error) {
	log.Println("Task.GetNameDescription not mocked")
	_err = errors.New("Task.GetNameDescription not mocked")
	return
}

func (_class TaskClass) GetNameDescriptionMock(sessionID SessionRef, self TaskRef) (_retval string, _err error) {
	return TaskClass_GetNameDescriptionMockedCallback(sessionID, self)
}
// Get the name/description field of the given task.
func (_class TaskClass) GetNameDescription(sessionID SessionRef, self TaskRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetNameDescriptionMock(sessionID, self)
	}	
	_method := "task.get_name_description"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertTaskRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var TaskClass_GetNameLabelMockedCallback = func (sessionID SessionRef, self TaskRef) (_retval string, _err error) {
	log.Println("Task.GetNameLabel not mocked")
	_err = errors.New("Task.GetNameLabel not mocked")
	return
}

func (_class TaskClass) GetNameLabelMock(sessionID SessionRef, self TaskRef) (_retval string, _err error) {
	return TaskClass_GetNameLabelMockedCallback(sessionID, self)
}
// Get the name/label field of the given task.
func (_class TaskClass) GetNameLabel(sessionID SessionRef, self TaskRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetNameLabelMock(sessionID, self)
	}	
	_method := "task.get_name_label"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertTaskRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var TaskClass_GetUUIDMockedCallback = func (sessionID SessionRef, self TaskRef) (_retval string, _err error) {
	log.Println("Task.GetUUID not mocked")
	_err = errors.New("Task.GetUUID not mocked")
	return
}

func (_class TaskClass) GetUUIDMock(sessionID SessionRef, self TaskRef) (_retval string, _err error) {
	return TaskClass_GetUUIDMockedCallback(sessionID, self)
}
// Get the uuid field of the given task.
func (_class TaskClass) GetUUID(sessionID SessionRef, self TaskRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetUUIDMock(sessionID, self)
	}	
	_method := "task.get_uuid"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertTaskRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var TaskClass_GetByNameLabelMockedCallback = func (sessionID SessionRef, label string) (_retval []TaskRef, _err error) {
	log.Println("Task.GetByNameLabel not mocked")
	_err = errors.New("Task.GetByNameLabel not mocked")
	return
}

func (_class TaskClass) GetByNameLabelMock(sessionID SessionRef, label string) (_retval []TaskRef, _err error) {
	return TaskClass_GetByNameLabelMockedCallback(sessionID, label)
}
// Get all the task instances with the given label.
func (_class TaskClass) GetByNameLabel(sessionID SessionRef, label string) (_retval []TaskRef, _err error) {
	if (IsMock) {
		return _class.GetByNameLabelMock(sessionID, label)
	}	
	_method := "task.get_by_name_label"
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
	_retval, _err = convertTaskRefSetToGo(_method + " -> ", _result.Value)
	return
}


var TaskClass_GetByUUIDMockedCallback = func (sessionID SessionRef, uuid string) (_retval TaskRef, _err error) {
	log.Println("Task.GetByUUID not mocked")
	_err = errors.New("Task.GetByUUID not mocked")
	return
}

func (_class TaskClass) GetByUUIDMock(sessionID SessionRef, uuid string) (_retval TaskRef, _err error) {
	return TaskClass_GetByUUIDMockedCallback(sessionID, uuid)
}
// Get a reference to the task instance with the specified UUID.
func (_class TaskClass) GetByUUID(sessionID SessionRef, uuid string) (_retval TaskRef, _err error) {
	if (IsMock) {
		return _class.GetByUUIDMock(sessionID, uuid)
	}	
	_method := "task.get_by_uuid"
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
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}


var TaskClass_GetRecordMockedCallback = func (sessionID SessionRef, self TaskRef) (_retval TaskRecord, _err error) {
	log.Println("Task.GetRecord not mocked")
	_err = errors.New("Task.GetRecord not mocked")
	return
}

func (_class TaskClass) GetRecordMock(sessionID SessionRef, self TaskRef) (_retval TaskRecord, _err error) {
	return TaskClass_GetRecordMockedCallback(sessionID, self)
}
// Get a record containing the current state of the given task.
func (_class TaskClass) GetRecord(sessionID SessionRef, self TaskRef) (_retval TaskRecord, _err error) {
	if (IsMock) {
		return _class.GetRecordMock(sessionID, self)
	}	
	_method := "task.get_record"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertTaskRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRecordToGo(_method + " -> ", _result.Value)
	return
}
