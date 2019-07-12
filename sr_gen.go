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

type StorageOperations string

const (
  // Scanning backends for new or deleted VDIs
	StorageOperationsScan StorageOperations = "scan"
  // Destroying the SR
	StorageOperationsDestroy StorageOperations = "destroy"
  // Forgetting about SR
	StorageOperationsForget StorageOperations = "forget"
  // Plugging a PBD into this SR
	StorageOperationsPlug StorageOperations = "plug"
  // Unplugging a PBD from this SR
	StorageOperationsUnplug StorageOperations = "unplug"
  // Refresh the fields on the SR
	StorageOperationsUpdate StorageOperations = "update"
  // Creating a new VDI
	StorageOperationsVdiCreate StorageOperations = "vdi_create"
  // Introducing a new VDI
	StorageOperationsVdiIntroduce StorageOperations = "vdi_introduce"
  // Destroying a VDI
	StorageOperationsVdiDestroy StorageOperations = "vdi_destroy"
  // Resizing a VDI
	StorageOperationsVdiResize StorageOperations = "vdi_resize"
  // Cloneing a VDI
	StorageOperationsVdiClone StorageOperations = "vdi_clone"
  // Snapshotting a VDI
	StorageOperationsVdiSnapshot StorageOperations = "vdi_snapshot"
  // Mirroring a VDI
	StorageOperationsVdiMirror StorageOperations = "vdi_mirror"
  // Enabling changed block tracking for a VDI
	StorageOperationsVdiEnableCbt StorageOperations = "vdi_enable_cbt"
  // Disabling changed block tracking for a VDI
	StorageOperationsVdiDisableCbt StorageOperations = "vdi_disable_cbt"
  // Deleting the data of the VDI
	StorageOperationsVdiDataDestroy StorageOperations = "vdi_data_destroy"
  // Exporting a bitmap that shows the changed blocks between two VDIs
	StorageOperationsVdiListChangedBlocks StorageOperations = "vdi_list_changed_blocks"
  // Setting the on_boot field of the VDI
	StorageOperationsVdiSetOnBoot StorageOperations = "vdi_set_on_boot"
  // Creating a PBD for this SR
	StorageOperationsPbdCreate StorageOperations = "pbd_create"
  // Destroying one of this SR's PBDs
	StorageOperationsPbdDestroy StorageOperations = "pbd_destroy"
)

type SRRecord struct {
  // Unique identifier/object reference
	UUID string
  // a human-readable name
	NameLabel string
  // a notes field containing human-readable description
	NameDescription string
  // list of the operations allowed in this state. This list is advisory only and the server state may have changed by the time this field is read by a client.
	AllowedOperations []StorageOperations
  // links each of the running tasks using this object (by reference) to a current_operation enum which describes the nature of the task.
	CurrentOperations map[string]StorageOperations
  // all virtual disks known to this storage repository
	VDIs []VDIRef
  // describes how particular hosts can see this storage repository
	PBDs []PBDRef
  // sum of virtual_sizes of all VDIs in this storage repository (in bytes)
	VirtualAllocation int
  // physical space currently utilised on this storage repository (in bytes). Note that for sparse disk formats, physical_utilisation may be less than virtual_allocation
	PhysicalUtilisation int
  // total physical size of the repository (in bytes)
	PhysicalSize int
  // type of the storage repository
	Type string
  // the type of the SR's content, if required (e.g. ISOs)
	ContentType string
  // true if this SR is (capable of being) shared between multiple hosts
	Shared bool
  // additional configuration
	OtherConfig map[string]string
  // user-specified tags for categorization purposes
	Tags []string
  // SM dependent data
	SmConfig map[string]string
  // Binary blobs associated with this SR
	Blobs map[string]BlobRef
  // True if this SR is assigned to be the local cache for its host
	LocalCacheEnabled bool
  // The disaster recovery task which introduced this SR
	IntroducedBy DRTaskRef
  // True if the SR is using aggregated local storage
	Clustered bool
  // True if this is the SR that contains the Tools ISO VDIs
	IsToolsSr bool
}

type SRRef string

// A storage repository
type SRClass struct {
	client *Client
}

func (_class SRClass) GetAllRecords__mock(sessionID SessionRef) (_retval map[SRRef]SRRecord, _err error) {
	log.Println("SR.GetAllRecords not mocked")
	_err = errors.New("SR.GetAllRecords not mocked")
	return
}
// Return a map of SR references to SR records for all SRs known to the system.
func (_class SRClass) GetAllRecords(sessionID SessionRef) (_retval map[SRRef]SRRecord, _err error) {
	if (IsMock) {
		return _class.GetAllRecords__mock(sessionID)
	}	
	_method := "SR.get_all_records"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertSRRefToSRRecordMapToGo(_method + " -> ", _result.Value)
	return
}

func (_class SRClass) GetAll__mock(sessionID SessionRef) (_retval []SRRef, _err error) {
	log.Println("SR.GetAll not mocked")
	_err = errors.New("SR.GetAll not mocked")
	return
}
// Return a list of all the SRs known to the system.
func (_class SRClass) GetAll(sessionID SessionRef) (_retval []SRRef, _err error) {
	if (IsMock) {
		return _class.GetAll__mock(sessionID)
	}	
	_method := "SR.get_all"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertSRRefSetToGo(_method + " -> ", _result.Value)
	return
}

func (_class SRClass) ForgetDataSourceArchives__mock(sessionID SessionRef, sr SRRef, dataSource string) (_err error) {
	log.Println("SR.ForgetDataSourceArchives not mocked")
	_err = errors.New("SR.ForgetDataSourceArchives not mocked")
	return
}
// Forget the recorded statistics related to the specified data source
func (_class SRClass) ForgetDataSourceArchives(sessionID SessionRef, sr SRRef, dataSource string) (_err error) {
	if (IsMock) {
		return _class.ForgetDataSourceArchives__mock(sessionID, sr, dataSource)
	}	
	_method := "SR.forget_data_source_archives"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_srArg, _err := convertSRRefToXen(fmt.Sprintf("%s(%s)", _method, "sr"), sr)
	if _err != nil {
		return
	}
	_dataSourceArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "data_source"), dataSource)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _srArg, _dataSourceArg)
	return
}

func (_class SRClass) QueryDataSource__mock(sessionID SessionRef, sr SRRef, dataSource string) (_retval float64, _err error) {
	log.Println("SR.QueryDataSource not mocked")
	_err = errors.New("SR.QueryDataSource not mocked")
	return
}
// Query the latest value of the specified data source
func (_class SRClass) QueryDataSource(sessionID SessionRef, sr SRRef, dataSource string) (_retval float64, _err error) {
	if (IsMock) {
		return _class.QueryDataSource__mock(sessionID, sr, dataSource)
	}	
	_method := "SR.query_data_source"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_srArg, _err := convertSRRefToXen(fmt.Sprintf("%s(%s)", _method, "sr"), sr)
	if _err != nil {
		return
	}
	_dataSourceArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "data_source"), dataSource)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _srArg, _dataSourceArg)
	if _err != nil {
		return
	}
	_retval, _err = convertFloatToGo(_method + " -> ", _result.Value)
	return
}

func (_class SRClass) RecordDataSource__mock(sessionID SessionRef, sr SRRef, dataSource string) (_err error) {
	log.Println("SR.RecordDataSource not mocked")
	_err = errors.New("SR.RecordDataSource not mocked")
	return
}
// Start recording the specified data source
func (_class SRClass) RecordDataSource(sessionID SessionRef, sr SRRef, dataSource string) (_err error) {
	if (IsMock) {
		return _class.RecordDataSource__mock(sessionID, sr, dataSource)
	}	
	_method := "SR.record_data_source"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_srArg, _err := convertSRRefToXen(fmt.Sprintf("%s(%s)", _method, "sr"), sr)
	if _err != nil {
		return
	}
	_dataSourceArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "data_source"), dataSource)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _srArg, _dataSourceArg)
	return
}

func (_class SRClass) GetDataSources__mock(sessionID SessionRef, sr SRRef) (_retval []DataSourceRecord, _err error) {
	log.Println("SR.GetDataSources not mocked")
	_err = errors.New("SR.GetDataSources not mocked")
	return
}
// 
func (_class SRClass) GetDataSources(sessionID SessionRef, sr SRRef) (_retval []DataSourceRecord, _err error) {
	if (IsMock) {
		return _class.GetDataSources__mock(sessionID, sr)
	}	
	_method := "SR.get_data_sources"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_srArg, _err := convertSRRefToXen(fmt.Sprintf("%s(%s)", _method, "sr"), sr)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _srArg)
	if _err != nil {
		return
	}
	_retval, _err = convertDataSourceRecordSetToGo(_method + " -> ", _result.Value)
	return
}

func (_class SRClass) DisableDatabaseReplication__mock(sessionID SessionRef, sr SRRef) (_err error) {
	log.Println("SR.DisableDatabaseReplication not mocked")
	_err = errors.New("SR.DisableDatabaseReplication not mocked")
	return
}
// 
func (_class SRClass) DisableDatabaseReplication(sessionID SessionRef, sr SRRef) (_err error) {
	if (IsMock) {
		return _class.DisableDatabaseReplication__mock(sessionID, sr)
	}	
	_method := "SR.disable_database_replication"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_srArg, _err := convertSRRefToXen(fmt.Sprintf("%s(%s)", _method, "sr"), sr)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _srArg)
	return
}

func (_class SRClass) EnableDatabaseReplication__mock(sessionID SessionRef, sr SRRef) (_err error) {
	log.Println("SR.EnableDatabaseReplication not mocked")
	_err = errors.New("SR.EnableDatabaseReplication not mocked")
	return
}
// 
func (_class SRClass) EnableDatabaseReplication(sessionID SessionRef, sr SRRef) (_err error) {
	if (IsMock) {
		return _class.EnableDatabaseReplication__mock(sessionID, sr)
	}	
	_method := "SR.enable_database_replication"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_srArg, _err := convertSRRefToXen(fmt.Sprintf("%s(%s)", _method, "sr"), sr)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _srArg)
	return
}

func (_class SRClass) AssertSupportsDatabaseReplication__mock(sessionID SessionRef, sr SRRef) (_err error) {
	log.Println("SR.AssertSupportsDatabaseReplication not mocked")
	_err = errors.New("SR.AssertSupportsDatabaseReplication not mocked")
	return
}
// Returns successfully if the given SR supports database replication. Otherwise returns an error to explain why not.
func (_class SRClass) AssertSupportsDatabaseReplication(sessionID SessionRef, sr SRRef) (_err error) {
	if (IsMock) {
		return _class.AssertSupportsDatabaseReplication__mock(sessionID, sr)
	}	
	_method := "SR.assert_supports_database_replication"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_srArg, _err := convertSRRefToXen(fmt.Sprintf("%s(%s)", _method, "sr"), sr)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _srArg)
	return
}

func (_class SRClass) AssertCanHostHaStatefile__mock(sessionID SessionRef, sr SRRef) (_err error) {
	log.Println("SR.AssertCanHostHaStatefile not mocked")
	_err = errors.New("SR.AssertCanHostHaStatefile not mocked")
	return
}
// Returns successfully if the given SR can host an HA statefile. Otherwise returns an error to explain why not
func (_class SRClass) AssertCanHostHaStatefile(sessionID SessionRef, sr SRRef) (_err error) {
	if (IsMock) {
		return _class.AssertCanHostHaStatefile__mock(sessionID, sr)
	}	
	_method := "SR.assert_can_host_ha_statefile"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_srArg, _err := convertSRRefToXen(fmt.Sprintf("%s(%s)", _method, "sr"), sr)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _srArg)
	return
}

func (_class SRClass) SetPhysicalUtilisation__mock(sessionID SessionRef, self SRRef, value int) (_err error) {
	log.Println("SR.SetPhysicalUtilisation not mocked")
	_err = errors.New("SR.SetPhysicalUtilisation not mocked")
	return
}
// Sets the SR's physical_utilisation field
func (_class SRClass) SetPhysicalUtilisation(sessionID SessionRef, self SRRef, value int) (_err error) {
	if (IsMock) {
		return _class.SetPhysicalUtilisation__mock(sessionID, self, value)
	}	
	_method := "SR.set_physical_utilisation"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertSRRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class SRClass) SetVirtualAllocation__mock(sessionID SessionRef, self SRRef, value int) (_err error) {
	log.Println("SR.SetVirtualAllocation not mocked")
	_err = errors.New("SR.SetVirtualAllocation not mocked")
	return
}
// Sets the SR's virtual_allocation field
func (_class SRClass) SetVirtualAllocation(sessionID SessionRef, self SRRef, value int) (_err error) {
	if (IsMock) {
		return _class.SetVirtualAllocation__mock(sessionID, self, value)
	}	
	_method := "SR.set_virtual_allocation"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertSRRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class SRClass) SetPhysicalSize__mock(sessionID SessionRef, self SRRef, value int) (_err error) {
	log.Println("SR.SetPhysicalSize not mocked")
	_err = errors.New("SR.SetPhysicalSize not mocked")
	return
}
// Sets the SR's physical_size field
func (_class SRClass) SetPhysicalSize(sessionID SessionRef, self SRRef, value int) (_err error) {
	if (IsMock) {
		return _class.SetPhysicalSize__mock(sessionID, self, value)
	}	
	_method := "SR.set_physical_size"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertSRRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class SRClass) CreateNewBlob__mock(sessionID SessionRef, sr SRRef, name string, mimeType string, public bool) (_retval BlobRef, _err error) {
	log.Println("SR.CreateNewBlob not mocked")
	_err = errors.New("SR.CreateNewBlob not mocked")
	return
}
// Create a placeholder for a named binary blob of data that is associated with this SR
func (_class SRClass) CreateNewBlob(sessionID SessionRef, sr SRRef, name string, mimeType string, public bool) (_retval BlobRef, _err error) {
	if (IsMock) {
		return _class.CreateNewBlob__mock(sessionID, sr, name, mimeType, public)
	}	
	_method := "SR.create_new_blob"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_srArg, _err := convertSRRefToXen(fmt.Sprintf("%s(%s)", _method, "sr"), sr)
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
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _srArg, _nameArg, _mimeTypeArg, _publicArg)
	if _err != nil {
		return
	}
	_retval, _err = convertBlobRefToGo(_method + " -> ", _result.Value)
	return
}

func (_class SRClass) SetNameDescription__mock(sessionID SessionRef, sr SRRef, value string) (_err error) {
	log.Println("SR.SetNameDescription not mocked")
	_err = errors.New("SR.SetNameDescription not mocked")
	return
}
// Set the name description of the SR
func (_class SRClass) SetNameDescription(sessionID SessionRef, sr SRRef, value string) (_err error) {
	if (IsMock) {
		return _class.SetNameDescription__mock(sessionID, sr, value)
	}	
	_method := "SR.set_name_description"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_srArg, _err := convertSRRefToXen(fmt.Sprintf("%s(%s)", _method, "sr"), sr)
	if _err != nil {
		return
	}
	_valueArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _srArg, _valueArg)
	return
}

func (_class SRClass) SetNameLabel__mock(sessionID SessionRef, sr SRRef, value string) (_err error) {
	log.Println("SR.SetNameLabel not mocked")
	_err = errors.New("SR.SetNameLabel not mocked")
	return
}
// Set the name label of the SR
func (_class SRClass) SetNameLabel(sessionID SessionRef, sr SRRef, value string) (_err error) {
	if (IsMock) {
		return _class.SetNameLabel__mock(sessionID, sr, value)
	}	
	_method := "SR.set_name_label"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_srArg, _err := convertSRRefToXen(fmt.Sprintf("%s(%s)", _method, "sr"), sr)
	if _err != nil {
		return
	}
	_valueArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _srArg, _valueArg)
	return
}

func (_class SRClass) SetShared__mock(sessionID SessionRef, sr SRRef, value bool) (_err error) {
	log.Println("SR.SetShared not mocked")
	_err = errors.New("SR.SetShared not mocked")
	return
}
// Sets the shared flag on the SR
func (_class SRClass) SetShared(sessionID SessionRef, sr SRRef, value bool) (_err error) {
	if (IsMock) {
		return _class.SetShared__mock(sessionID, sr, value)
	}	
	_method := "SR.set_shared"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_srArg, _err := convertSRRefToXen(fmt.Sprintf("%s(%s)", _method, "sr"), sr)
	if _err != nil {
		return
	}
	_valueArg, _err := convertBoolToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _srArg, _valueArg)
	return
}

func (_class SRClass) Probe__mock(sessionID SessionRef, host HostRef, deviceConfig map[string]string, atype string, smConfig map[string]string) (_retval string, _err error) {
	log.Println("SR.Probe not mocked")
	_err = errors.New("SR.Probe not mocked")
	return
}
// Perform a backend-specific scan, using the given device_config.  If the device_config is complete, then this will return a list of the SRs present of this type on the device, if any.  If the device_config is partial, then a backend-specific scan will be performed, returning results that will guide the user in improving the device_config.
func (_class SRClass) Probe(sessionID SessionRef, host HostRef, deviceConfig map[string]string, atype string, smConfig map[string]string) (_retval string, _err error) {
	if (IsMock) {
		return _class.Probe__mock(sessionID, host, deviceConfig, atype, smConfig)
	}	
	_method := "SR.probe"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_hostArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "host"), host)
	if _err != nil {
		return
	}
	_deviceConfigArg, _err := convertStringToStringMapToXen(fmt.Sprintf("%s(%s)", _method, "device_config"), deviceConfig)
	if _err != nil {
		return
	}
	_atypeArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "type"), atype)
	if _err != nil {
		return
	}
	_smConfigArg, _err := convertStringToStringMapToXen(fmt.Sprintf("%s(%s)", _method, "sm_config"), smConfig)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _hostArg, _deviceConfigArg, _atypeArg, _smConfigArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringToGo(_method + " -> ", _result.Value)
	return
}

func (_class SRClass) Scan__mock(sessionID SessionRef, sr SRRef) (_err error) {
	log.Println("SR.Scan not mocked")
	_err = errors.New("SR.Scan not mocked")
	return
}
// Refreshes the list of VDIs associated with an SR
func (_class SRClass) Scan(sessionID SessionRef, sr SRRef) (_err error) {
	if (IsMock) {
		return _class.Scan__mock(sessionID, sr)
	}	
	_method := "SR.scan"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_srArg, _err := convertSRRefToXen(fmt.Sprintf("%s(%s)", _method, "sr"), sr)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _srArg)
	return
}

func (_class SRClass) GetSupportedTypes__mock(sessionID SessionRef) (_retval []string, _err error) {
	log.Println("SR.GetSupportedTypes not mocked")
	_err = errors.New("SR.GetSupportedTypes not mocked")
	return
}
// Return a set of all the SR types supported by the system
func (_class SRClass) GetSupportedTypes(sessionID SessionRef) (_retval []string, _err error) {
	if (IsMock) {
		return _class.GetSupportedTypes__mock(sessionID)
	}	
	_method := "SR.get_supported_types"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringSetToGo(_method + " -> ", _result.Value)
	return
}

func (_class SRClass) Update__mock(sessionID SessionRef, sr SRRef) (_err error) {
	log.Println("SR.Update not mocked")
	_err = errors.New("SR.Update not mocked")
	return
}
// Refresh the fields on the SR object
func (_class SRClass) Update(sessionID SessionRef, sr SRRef) (_err error) {
	if (IsMock) {
		return _class.Update__mock(sessionID, sr)
	}	
	_method := "SR.update"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_srArg, _err := convertSRRefToXen(fmt.Sprintf("%s(%s)", _method, "sr"), sr)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _srArg)
	return
}

func (_class SRClass) Forget__mock(sessionID SessionRef, sr SRRef) (_err error) {
	log.Println("SR.Forget not mocked")
	_err = errors.New("SR.Forget not mocked")
	return
}
// Removing specified SR-record from database, without attempting to remove SR from disk
//
// Errors:
//  SR_HAS_PBD - The SR is still connected to a host via a PBD. It cannot be destroyed or forgotten.
func (_class SRClass) Forget(sessionID SessionRef, sr SRRef) (_err error) {
	if (IsMock) {
		return _class.Forget__mock(sessionID, sr)
	}	
	_method := "SR.forget"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_srArg, _err := convertSRRefToXen(fmt.Sprintf("%s(%s)", _method, "sr"), sr)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _srArg)
	return
}

func (_class SRClass) Destroy__mock(sessionID SessionRef, sr SRRef) (_err error) {
	log.Println("SR.Destroy not mocked")
	_err = errors.New("SR.Destroy not mocked")
	return
}
// Destroy specified SR, removing SR-record from database and remove SR from disk. (In order to affect this operation the appropriate device_config is read from the specified SR's PBD on current host)
//
// Errors:
//  SR_HAS_PBD - The SR is still connected to a host via a PBD. It cannot be destroyed or forgotten.
func (_class SRClass) Destroy(sessionID SessionRef, sr SRRef) (_err error) {
	if (IsMock) {
		return _class.Destroy__mock(sessionID, sr)
	}	
	_method := "SR.destroy"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_srArg, _err := convertSRRefToXen(fmt.Sprintf("%s(%s)", _method, "sr"), sr)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _srArg)
	return
}

func (_class SRClass) Make__mock(sessionID SessionRef, host HostRef, deviceConfig map[string]string, physicalSize int, nameLabel string, nameDescription string, atype string, contentType string, smConfig map[string]string) (_retval string, _err error) {
	log.Println("SR.Make not mocked")
	_err = errors.New("SR.Make not mocked")
	return
}
// Create a new Storage Repository on disk. This call is deprecated: use SR.create instead.
func (_class SRClass) Make(sessionID SessionRef, host HostRef, deviceConfig map[string]string, physicalSize int, nameLabel string, nameDescription string, atype string, contentType string, smConfig map[string]string) (_retval string, _err error) {
	if (IsMock) {
		return _class.Make__mock(sessionID, host, deviceConfig, physicalSize, nameLabel, nameDescription, atype, contentType, smConfig)
	}	
	_method := "SR.make"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_hostArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "host"), host)
	if _err != nil {
		return
	}
	_deviceConfigArg, _err := convertStringToStringMapToXen(fmt.Sprintf("%s(%s)", _method, "device_config"), deviceConfig)
	if _err != nil {
		return
	}
	_physicalSizeArg, _err := convertIntToXen(fmt.Sprintf("%s(%s)", _method, "physical_size"), physicalSize)
	if _err != nil {
		return
	}
	_nameLabelArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "name_label"), nameLabel)
	if _err != nil {
		return
	}
	_nameDescriptionArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "name_description"), nameDescription)
	if _err != nil {
		return
	}
	_atypeArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "type"), atype)
	if _err != nil {
		return
	}
	_contentTypeArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "content_type"), contentType)
	if _err != nil {
		return
	}
	_smConfigArg, _err := convertStringToStringMapToXen(fmt.Sprintf("%s(%s)", _method, "sm_config"), smConfig)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _hostArg, _deviceConfigArg, _physicalSizeArg, _nameLabelArg, _nameDescriptionArg, _atypeArg, _contentTypeArg, _smConfigArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringToGo(_method + " -> ", _result.Value)
	return
}

func (_class SRClass) Introduce__mock(sessionID SessionRef, uuid string, nameLabel string, nameDescription string, atype string, contentType string, shared bool, smConfig map[string]string) (_retval SRRef, _err error) {
	log.Println("SR.Introduce not mocked")
	_err = errors.New("SR.Introduce not mocked")
	return
}
// Introduce a new Storage Repository into the managed system
func (_class SRClass) Introduce(sessionID SessionRef, uuid string, nameLabel string, nameDescription string, atype string, contentType string, shared bool, smConfig map[string]string) (_retval SRRef, _err error) {
	if (IsMock) {
		return _class.Introduce__mock(sessionID, uuid, nameLabel, nameDescription, atype, contentType, shared, smConfig)
	}	
	_method := "SR.introduce"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_uuidArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "uuid"), uuid)
	if _err != nil {
		return
	}
	_nameLabelArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "name_label"), nameLabel)
	if _err != nil {
		return
	}
	_nameDescriptionArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "name_description"), nameDescription)
	if _err != nil {
		return
	}
	_atypeArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "type"), atype)
	if _err != nil {
		return
	}
	_contentTypeArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "content_type"), contentType)
	if _err != nil {
		return
	}
	_sharedArg, _err := convertBoolToXen(fmt.Sprintf("%s(%s)", _method, "shared"), shared)
	if _err != nil {
		return
	}
	_smConfigArg, _err := convertStringToStringMapToXen(fmt.Sprintf("%s(%s)", _method, "sm_config"), smConfig)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _uuidArg, _nameLabelArg, _nameDescriptionArg, _atypeArg, _contentTypeArg, _sharedArg, _smConfigArg)
	if _err != nil {
		return
	}
	_retval, _err = convertSRRefToGo(_method + " -> ", _result.Value)
	return
}

func (_class SRClass) Create__mock(sessionID SessionRef, host HostRef, deviceConfig map[string]string, physicalSize int, nameLabel string, nameDescription string, atype string, contentType string, shared bool, smConfig map[string]string) (_retval SRRef, _err error) {
	log.Println("SR.Create not mocked")
	_err = errors.New("SR.Create not mocked")
	return
}
// Create a new Storage Repository and introduce it into the managed system, creating both SR record and PBD record to attach it to current host (with specified device_config parameters)
//
// Errors:
//  SR_UNKNOWN_DRIVER - The SR could not be connected because the driver was not recognised.
func (_class SRClass) Create(sessionID SessionRef, host HostRef, deviceConfig map[string]string, physicalSize int, nameLabel string, nameDescription string, atype string, contentType string, shared bool, smConfig map[string]string) (_retval SRRef, _err error) {
	if (IsMock) {
		return _class.Create__mock(sessionID, host, deviceConfig, physicalSize, nameLabel, nameDescription, atype, contentType, shared, smConfig)
	}	
	_method := "SR.create"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_hostArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "host"), host)
	if _err != nil {
		return
	}
	_deviceConfigArg, _err := convertStringToStringMapToXen(fmt.Sprintf("%s(%s)", _method, "device_config"), deviceConfig)
	if _err != nil {
		return
	}
	_physicalSizeArg, _err := convertIntToXen(fmt.Sprintf("%s(%s)", _method, "physical_size"), physicalSize)
	if _err != nil {
		return
	}
	_nameLabelArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "name_label"), nameLabel)
	if _err != nil {
		return
	}
	_nameDescriptionArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "name_description"), nameDescription)
	if _err != nil {
		return
	}
	_atypeArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "type"), atype)
	if _err != nil {
		return
	}
	_contentTypeArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "content_type"), contentType)
	if _err != nil {
		return
	}
	_sharedArg, _err := convertBoolToXen(fmt.Sprintf("%s(%s)", _method, "shared"), shared)
	if _err != nil {
		return
	}
	_smConfigArg, _err := convertStringToStringMapToXen(fmt.Sprintf("%s(%s)", _method, "sm_config"), smConfig)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _hostArg, _deviceConfigArg, _physicalSizeArg, _nameLabelArg, _nameDescriptionArg, _atypeArg, _contentTypeArg, _sharedArg, _smConfigArg)
	if _err != nil {
		return
	}
	_retval, _err = convertSRRefToGo(_method + " -> ", _result.Value)
	return
}

func (_class SRClass) RemoveFromSmConfig__mock(sessionID SessionRef, self SRRef, key string) (_err error) {
	log.Println("SR.RemoveFromSmConfig not mocked")
	_err = errors.New("SR.RemoveFromSmConfig not mocked")
	return
}
// Remove the given key and its corresponding value from the sm_config field of the given SR.  If the key is not in that Map, then do nothing.
func (_class SRClass) RemoveFromSmConfig(sessionID SessionRef, self SRRef, key string) (_err error) {
	if (IsMock) {
		return _class.RemoveFromSmConfig__mock(sessionID, self, key)
	}	
	_method := "SR.remove_from_sm_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertSRRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class SRClass) AddToSmConfig__mock(sessionID SessionRef, self SRRef, key string, value string) (_err error) {
	log.Println("SR.AddToSmConfig not mocked")
	_err = errors.New("SR.AddToSmConfig not mocked")
	return
}
// Add the given key-value pair to the sm_config field of the given SR.
func (_class SRClass) AddToSmConfig(sessionID SessionRef, self SRRef, key string, value string) (_err error) {
	if (IsMock) {
		return _class.AddToSmConfig__mock(sessionID, self, key, value)
	}	
	_method := "SR.add_to_sm_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertSRRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class SRClass) SetSmConfig__mock(sessionID SessionRef, self SRRef, value map[string]string) (_err error) {
	log.Println("SR.SetSmConfig not mocked")
	_err = errors.New("SR.SetSmConfig not mocked")
	return
}
// Set the sm_config field of the given SR.
func (_class SRClass) SetSmConfig(sessionID SessionRef, self SRRef, value map[string]string) (_err error) {
	if (IsMock) {
		return _class.SetSmConfig__mock(sessionID, self, value)
	}	
	_method := "SR.set_sm_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertSRRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class SRClass) RemoveTags__mock(sessionID SessionRef, self SRRef, value string) (_err error) {
	log.Println("SR.RemoveTags not mocked")
	_err = errors.New("SR.RemoveTags not mocked")
	return
}
// Remove the given value from the tags field of the given SR.  If the value is not in that Set, then do nothing.
func (_class SRClass) RemoveTags(sessionID SessionRef, self SRRef, value string) (_err error) {
	if (IsMock) {
		return _class.RemoveTags__mock(sessionID, self, value)
	}	
	_method := "SR.remove_tags"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertSRRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class SRClass) AddTags__mock(sessionID SessionRef, self SRRef, value string) (_err error) {
	log.Println("SR.AddTags not mocked")
	_err = errors.New("SR.AddTags not mocked")
	return
}
// Add the given value to the tags field of the given SR.  If the value is already in that Set, then do nothing.
func (_class SRClass) AddTags(sessionID SessionRef, self SRRef, value string) (_err error) {
	if (IsMock) {
		return _class.AddTags__mock(sessionID, self, value)
	}	
	_method := "SR.add_tags"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertSRRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class SRClass) SetTags__mock(sessionID SessionRef, self SRRef, value []string) (_err error) {
	log.Println("SR.SetTags not mocked")
	_err = errors.New("SR.SetTags not mocked")
	return
}
// Set the tags field of the given SR.
func (_class SRClass) SetTags(sessionID SessionRef, self SRRef, value []string) (_err error) {
	if (IsMock) {
		return _class.SetTags__mock(sessionID, self, value)
	}	
	_method := "SR.set_tags"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertSRRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class SRClass) RemoveFromOtherConfig__mock(sessionID SessionRef, self SRRef, key string) (_err error) {
	log.Println("SR.RemoveFromOtherConfig not mocked")
	_err = errors.New("SR.RemoveFromOtherConfig not mocked")
	return
}
// Remove the given key and its corresponding value from the other_config field of the given SR.  If the key is not in that Map, then do nothing.
func (_class SRClass) RemoveFromOtherConfig(sessionID SessionRef, self SRRef, key string) (_err error) {
	if (IsMock) {
		return _class.RemoveFromOtherConfig__mock(sessionID, self, key)
	}	
	_method := "SR.remove_from_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertSRRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class SRClass) AddToOtherConfig__mock(sessionID SessionRef, self SRRef, key string, value string) (_err error) {
	log.Println("SR.AddToOtherConfig not mocked")
	_err = errors.New("SR.AddToOtherConfig not mocked")
	return
}
// Add the given key-value pair to the other_config field of the given SR.
func (_class SRClass) AddToOtherConfig(sessionID SessionRef, self SRRef, key string, value string) (_err error) {
	if (IsMock) {
		return _class.AddToOtherConfig__mock(sessionID, self, key, value)
	}	
	_method := "SR.add_to_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertSRRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class SRClass) SetOtherConfig__mock(sessionID SessionRef, self SRRef, value map[string]string) (_err error) {
	log.Println("SR.SetOtherConfig not mocked")
	_err = errors.New("SR.SetOtherConfig not mocked")
	return
}
// Set the other_config field of the given SR.
func (_class SRClass) SetOtherConfig(sessionID SessionRef, self SRRef, value map[string]string) (_err error) {
	if (IsMock) {
		return _class.SetOtherConfig__mock(sessionID, self, value)
	}	
	_method := "SR.set_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertSRRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class SRClass) GetIsToolsSr__mock(sessionID SessionRef, self SRRef) (_retval bool, _err error) {
	log.Println("SR.GetIsToolsSr not mocked")
	_err = errors.New("SR.GetIsToolsSr not mocked")
	return
}
// Get the is_tools_sr field of the given SR.
func (_class SRClass) GetIsToolsSr(sessionID SessionRef, self SRRef) (_retval bool, _err error) {
	if (IsMock) {
		return _class.GetIsToolsSr__mock(sessionID, self)
	}	
	_method := "SR.get_is_tools_sr"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertSRRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class SRClass) GetClustered__mock(sessionID SessionRef, self SRRef) (_retval bool, _err error) {
	log.Println("SR.GetClustered not mocked")
	_err = errors.New("SR.GetClustered not mocked")
	return
}
// Get the clustered field of the given SR.
func (_class SRClass) GetClustered(sessionID SessionRef, self SRRef) (_retval bool, _err error) {
	if (IsMock) {
		return _class.GetClustered__mock(sessionID, self)
	}	
	_method := "SR.get_clustered"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertSRRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class SRClass) GetIntroducedBy__mock(sessionID SessionRef, self SRRef) (_retval DRTaskRef, _err error) {
	log.Println("SR.GetIntroducedBy not mocked")
	_err = errors.New("SR.GetIntroducedBy not mocked")
	return
}
// Get the introduced_by field of the given SR.
func (_class SRClass) GetIntroducedBy(sessionID SessionRef, self SRRef) (_retval DRTaskRef, _err error) {
	if (IsMock) {
		return _class.GetIntroducedBy__mock(sessionID, self)
	}	
	_method := "SR.get_introduced_by"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertSRRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertDRTaskRefToGo(_method + " -> ", _result.Value)
	return
}

func (_class SRClass) GetLocalCacheEnabled__mock(sessionID SessionRef, self SRRef) (_retval bool, _err error) {
	log.Println("SR.GetLocalCacheEnabled not mocked")
	_err = errors.New("SR.GetLocalCacheEnabled not mocked")
	return
}
// Get the local_cache_enabled field of the given SR.
func (_class SRClass) GetLocalCacheEnabled(sessionID SessionRef, self SRRef) (_retval bool, _err error) {
	if (IsMock) {
		return _class.GetLocalCacheEnabled__mock(sessionID, self)
	}	
	_method := "SR.get_local_cache_enabled"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertSRRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class SRClass) GetBlobs__mock(sessionID SessionRef, self SRRef) (_retval map[string]BlobRef, _err error) {
	log.Println("SR.GetBlobs not mocked")
	_err = errors.New("SR.GetBlobs not mocked")
	return
}
// Get the blobs field of the given SR.
func (_class SRClass) GetBlobs(sessionID SessionRef, self SRRef) (_retval map[string]BlobRef, _err error) {
	if (IsMock) {
		return _class.GetBlobs__mock(sessionID, self)
	}	
	_method := "SR.get_blobs"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertSRRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class SRClass) GetSmConfig__mock(sessionID SessionRef, self SRRef) (_retval map[string]string, _err error) {
	log.Println("SR.GetSmConfig not mocked")
	_err = errors.New("SR.GetSmConfig not mocked")
	return
}
// Get the sm_config field of the given SR.
func (_class SRClass) GetSmConfig(sessionID SessionRef, self SRRef) (_retval map[string]string, _err error) {
	if (IsMock) {
		return _class.GetSmConfig__mock(sessionID, self)
	}	
	_method := "SR.get_sm_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertSRRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class SRClass) GetTags__mock(sessionID SessionRef, self SRRef) (_retval []string, _err error) {
	log.Println("SR.GetTags not mocked")
	_err = errors.New("SR.GetTags not mocked")
	return
}
// Get the tags field of the given SR.
func (_class SRClass) GetTags(sessionID SessionRef, self SRRef) (_retval []string, _err error) {
	if (IsMock) {
		return _class.GetTags__mock(sessionID, self)
	}	
	_method := "SR.get_tags"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertSRRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class SRClass) GetOtherConfig__mock(sessionID SessionRef, self SRRef) (_retval map[string]string, _err error) {
	log.Println("SR.GetOtherConfig not mocked")
	_err = errors.New("SR.GetOtherConfig not mocked")
	return
}
// Get the other_config field of the given SR.
func (_class SRClass) GetOtherConfig(sessionID SessionRef, self SRRef) (_retval map[string]string, _err error) {
	if (IsMock) {
		return _class.GetOtherConfig__mock(sessionID, self)
	}	
	_method := "SR.get_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertSRRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class SRClass) GetShared__mock(sessionID SessionRef, self SRRef) (_retval bool, _err error) {
	log.Println("SR.GetShared not mocked")
	_err = errors.New("SR.GetShared not mocked")
	return
}
// Get the shared field of the given SR.
func (_class SRClass) GetShared(sessionID SessionRef, self SRRef) (_retval bool, _err error) {
	if (IsMock) {
		return _class.GetShared__mock(sessionID, self)
	}	
	_method := "SR.get_shared"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertSRRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class SRClass) GetContentType__mock(sessionID SessionRef, self SRRef) (_retval string, _err error) {
	log.Println("SR.GetContentType not mocked")
	_err = errors.New("SR.GetContentType not mocked")
	return
}
// Get the content_type field of the given SR.
func (_class SRClass) GetContentType(sessionID SessionRef, self SRRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetContentType__mock(sessionID, self)
	}	
	_method := "SR.get_content_type"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertSRRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class SRClass) GetType__mock(sessionID SessionRef, self SRRef) (_retval string, _err error) {
	log.Println("SR.GetType not mocked")
	_err = errors.New("SR.GetType not mocked")
	return
}
// Get the type field of the given SR.
func (_class SRClass) GetType(sessionID SessionRef, self SRRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetType__mock(sessionID, self)
	}	
	_method := "SR.get_type"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertSRRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class SRClass) GetPhysicalSize__mock(sessionID SessionRef, self SRRef) (_retval int, _err error) {
	log.Println("SR.GetPhysicalSize not mocked")
	_err = errors.New("SR.GetPhysicalSize not mocked")
	return
}
// Get the physical_size field of the given SR.
func (_class SRClass) GetPhysicalSize(sessionID SessionRef, self SRRef) (_retval int, _err error) {
	if (IsMock) {
		return _class.GetPhysicalSize__mock(sessionID, self)
	}	
	_method := "SR.get_physical_size"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertSRRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class SRClass) GetPhysicalUtilisation__mock(sessionID SessionRef, self SRRef) (_retval int, _err error) {
	log.Println("SR.GetPhysicalUtilisation not mocked")
	_err = errors.New("SR.GetPhysicalUtilisation not mocked")
	return
}
// Get the physical_utilisation field of the given SR.
func (_class SRClass) GetPhysicalUtilisation(sessionID SessionRef, self SRRef) (_retval int, _err error) {
	if (IsMock) {
		return _class.GetPhysicalUtilisation__mock(sessionID, self)
	}	
	_method := "SR.get_physical_utilisation"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertSRRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class SRClass) GetVirtualAllocation__mock(sessionID SessionRef, self SRRef) (_retval int, _err error) {
	log.Println("SR.GetVirtualAllocation not mocked")
	_err = errors.New("SR.GetVirtualAllocation not mocked")
	return
}
// Get the virtual_allocation field of the given SR.
func (_class SRClass) GetVirtualAllocation(sessionID SessionRef, self SRRef) (_retval int, _err error) {
	if (IsMock) {
		return _class.GetVirtualAllocation__mock(sessionID, self)
	}	
	_method := "SR.get_virtual_allocation"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertSRRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class SRClass) GetPBDs__mock(sessionID SessionRef, self SRRef) (_retval []PBDRef, _err error) {
	log.Println("SR.GetPBDs not mocked")
	_err = errors.New("SR.GetPBDs not mocked")
	return
}
// Get the PBDs field of the given SR.
func (_class SRClass) GetPBDs(sessionID SessionRef, self SRRef) (_retval []PBDRef, _err error) {
	if (IsMock) {
		return _class.GetPBDs__mock(sessionID, self)
	}	
	_method := "SR.get_PBDs"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertSRRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertPBDRefSetToGo(_method + " -> ", _result.Value)
	return
}

func (_class SRClass) GetVDIs__mock(sessionID SessionRef, self SRRef) (_retval []VDIRef, _err error) {
	log.Println("SR.GetVDIs not mocked")
	_err = errors.New("SR.GetVDIs not mocked")
	return
}
// Get the VDIs field of the given SR.
func (_class SRClass) GetVDIs(sessionID SessionRef, self SRRef) (_retval []VDIRef, _err error) {
	if (IsMock) {
		return _class.GetVDIs__mock(sessionID, self)
	}	
	_method := "SR.get_VDIs"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertSRRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVDIRefSetToGo(_method + " -> ", _result.Value)
	return
}

func (_class SRClass) GetCurrentOperations__mock(sessionID SessionRef, self SRRef) (_retval map[string]StorageOperations, _err error) {
	log.Println("SR.GetCurrentOperations not mocked")
	_err = errors.New("SR.GetCurrentOperations not mocked")
	return
}
// Get the current_operations field of the given SR.
func (_class SRClass) GetCurrentOperations(sessionID SessionRef, self SRRef) (_retval map[string]StorageOperations, _err error) {
	if (IsMock) {
		return _class.GetCurrentOperations__mock(sessionID, self)
	}	
	_method := "SR.get_current_operations"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertSRRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringToEnumStorageOperationsMapToGo(_method + " -> ", _result.Value)
	return
}

func (_class SRClass) GetAllowedOperations__mock(sessionID SessionRef, self SRRef) (_retval []StorageOperations, _err error) {
	log.Println("SR.GetAllowedOperations not mocked")
	_err = errors.New("SR.GetAllowedOperations not mocked")
	return
}
// Get the allowed_operations field of the given SR.
func (_class SRClass) GetAllowedOperations(sessionID SessionRef, self SRRef) (_retval []StorageOperations, _err error) {
	if (IsMock) {
		return _class.GetAllowedOperations__mock(sessionID, self)
	}	
	_method := "SR.get_allowed_operations"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertSRRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertEnumStorageOperationsSetToGo(_method + " -> ", _result.Value)
	return
}

func (_class SRClass) GetNameDescription__mock(sessionID SessionRef, self SRRef) (_retval string, _err error) {
	log.Println("SR.GetNameDescription not mocked")
	_err = errors.New("SR.GetNameDescription not mocked")
	return
}
// Get the name/description field of the given SR.
func (_class SRClass) GetNameDescription(sessionID SessionRef, self SRRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetNameDescription__mock(sessionID, self)
	}	
	_method := "SR.get_name_description"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertSRRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class SRClass) GetNameLabel__mock(sessionID SessionRef, self SRRef) (_retval string, _err error) {
	log.Println("SR.GetNameLabel not mocked")
	_err = errors.New("SR.GetNameLabel not mocked")
	return
}
// Get the name/label field of the given SR.
func (_class SRClass) GetNameLabel(sessionID SessionRef, self SRRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetNameLabel__mock(sessionID, self)
	}	
	_method := "SR.get_name_label"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertSRRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class SRClass) GetUUID__mock(sessionID SessionRef, self SRRef) (_retval string, _err error) {
	log.Println("SR.GetUUID not mocked")
	_err = errors.New("SR.GetUUID not mocked")
	return
}
// Get the uuid field of the given SR.
func (_class SRClass) GetUUID(sessionID SessionRef, self SRRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetUUID__mock(sessionID, self)
	}	
	_method := "SR.get_uuid"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertSRRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class SRClass) GetByNameLabel__mock(sessionID SessionRef, label string) (_retval []SRRef, _err error) {
	log.Println("SR.GetByNameLabel not mocked")
	_err = errors.New("SR.GetByNameLabel not mocked")
	return
}
// Get all the SR instances with the given label.
func (_class SRClass) GetByNameLabel(sessionID SessionRef, label string) (_retval []SRRef, _err error) {
	if (IsMock) {
		return _class.GetByNameLabel__mock(sessionID, label)
	}	
	_method := "SR.get_by_name_label"
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
	_retval, _err = convertSRRefSetToGo(_method + " -> ", _result.Value)
	return
}

func (_class SRClass) GetByUUID__mock(sessionID SessionRef, uuid string) (_retval SRRef, _err error) {
	log.Println("SR.GetByUUID not mocked")
	_err = errors.New("SR.GetByUUID not mocked")
	return
}
// Get a reference to the SR instance with the specified UUID.
func (_class SRClass) GetByUUID(sessionID SessionRef, uuid string) (_retval SRRef, _err error) {
	if (IsMock) {
		return _class.GetByUUID__mock(sessionID, uuid)
	}	
	_method := "SR.get_by_uuid"
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
	_retval, _err = convertSRRefToGo(_method + " -> ", _result.Value)
	return
}

func (_class SRClass) GetRecord__mock(sessionID SessionRef, self SRRef) (_retval SRRecord, _err error) {
	log.Println("SR.GetRecord not mocked")
	_err = errors.New("SR.GetRecord not mocked")
	return
}
// Get a record containing the current state of the given SR.
func (_class SRClass) GetRecord(sessionID SessionRef, self SRRef) (_retval SRRecord, _err error) {
	if (IsMock) {
		return _class.GetRecord__mock(sessionID, self)
	}	
	_method := "SR.get_record"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertSRRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertSRRecordToGo(_method + " -> ", _result.Value)
	return
}
