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


func SRClassGetAllRecordsMockDefault(sessionID SessionRef) (_retval map[SRRef]SRRecord, _err error) {
	log.Println("SR.GetAllRecords not mocked")
	_err = errors.New("SR.GetAllRecords not mocked")
	return
}

var SRClassGetAllRecordsMockedCallback = SRClassGetAllRecordsMockDefault

func (_class SRClass) GetAllRecordsMock(sessionID SessionRef) (_retval map[SRRef]SRRecord, _err error) {
	return SRClassGetAllRecordsMockedCallback(sessionID)
}
// Return a map of SR references to SR records for all SRs known to the system.
func (_class SRClass) GetAllRecords(sessionID SessionRef) (_retval map[SRRef]SRRecord, _err error) {
	if IsMock {
		return _class.GetAllRecordsMock(sessionID)
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


func SRClassGetAllMockDefault(sessionID SessionRef) (_retval []SRRef, _err error) {
	log.Println("SR.GetAll not mocked")
	_err = errors.New("SR.GetAll not mocked")
	return
}

var SRClassGetAllMockedCallback = SRClassGetAllMockDefault

func (_class SRClass) GetAllMock(sessionID SessionRef) (_retval []SRRef, _err error) {
	return SRClassGetAllMockedCallback(sessionID)
}
// Return a list of all the SRs known to the system.
func (_class SRClass) GetAll(sessionID SessionRef) (_retval []SRRef, _err error) {
	if IsMock {
		return _class.GetAllMock(sessionID)
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


func SRClassForgetDataSourceArchivesMockDefault(sessionID SessionRef, sr SRRef, dataSource string) (_err error) {
	log.Println("SR.ForgetDataSourceArchives not mocked")
	_err = errors.New("SR.ForgetDataSourceArchives not mocked")
	return
}

var SRClassForgetDataSourceArchivesMockedCallback = SRClassForgetDataSourceArchivesMockDefault

func (_class SRClass) ForgetDataSourceArchivesMock(sessionID SessionRef, sr SRRef, dataSource string) (_err error) {
	return SRClassForgetDataSourceArchivesMockedCallback(sessionID, sr, dataSource)
}
// Forget the recorded statistics related to the specified data source
func (_class SRClass) ForgetDataSourceArchives(sessionID SessionRef, sr SRRef, dataSource string) (_err error) {
	if IsMock {
		return _class.ForgetDataSourceArchivesMock(sessionID, sr, dataSource)
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


func SRClassQueryDataSourceMockDefault(sessionID SessionRef, sr SRRef, dataSource string) (_retval float64, _err error) {
	log.Println("SR.QueryDataSource not mocked")
	_err = errors.New("SR.QueryDataSource not mocked")
	return
}

var SRClassQueryDataSourceMockedCallback = SRClassQueryDataSourceMockDefault

func (_class SRClass) QueryDataSourceMock(sessionID SessionRef, sr SRRef, dataSource string) (_retval float64, _err error) {
	return SRClassQueryDataSourceMockedCallback(sessionID, sr, dataSource)
}
// Query the latest value of the specified data source
func (_class SRClass) QueryDataSource(sessionID SessionRef, sr SRRef, dataSource string) (_retval float64, _err error) {
	if IsMock {
		return _class.QueryDataSourceMock(sessionID, sr, dataSource)
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


func SRClassRecordDataSourceMockDefault(sessionID SessionRef, sr SRRef, dataSource string) (_err error) {
	log.Println("SR.RecordDataSource not mocked")
	_err = errors.New("SR.RecordDataSource not mocked")
	return
}

var SRClassRecordDataSourceMockedCallback = SRClassRecordDataSourceMockDefault

func (_class SRClass) RecordDataSourceMock(sessionID SessionRef, sr SRRef, dataSource string) (_err error) {
	return SRClassRecordDataSourceMockedCallback(sessionID, sr, dataSource)
}
// Start recording the specified data source
func (_class SRClass) RecordDataSource(sessionID SessionRef, sr SRRef, dataSource string) (_err error) {
	if IsMock {
		return _class.RecordDataSourceMock(sessionID, sr, dataSource)
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


func SRClassGetDataSourcesMockDefault(sessionID SessionRef, sr SRRef) (_retval []DataSourceRecord, _err error) {
	log.Println("SR.GetDataSources not mocked")
	_err = errors.New("SR.GetDataSources not mocked")
	return
}

var SRClassGetDataSourcesMockedCallback = SRClassGetDataSourcesMockDefault

func (_class SRClass) GetDataSourcesMock(sessionID SessionRef, sr SRRef) (_retval []DataSourceRecord, _err error) {
	return SRClassGetDataSourcesMockedCallback(sessionID, sr)
}
// 
func (_class SRClass) GetDataSources(sessionID SessionRef, sr SRRef) (_retval []DataSourceRecord, _err error) {
	if IsMock {
		return _class.GetDataSourcesMock(sessionID, sr)
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


func SRClassDisableDatabaseReplicationMockDefault(sessionID SessionRef, sr SRRef) (_err error) {
	log.Println("SR.DisableDatabaseReplication not mocked")
	_err = errors.New("SR.DisableDatabaseReplication not mocked")
	return
}

var SRClassDisableDatabaseReplicationMockedCallback = SRClassDisableDatabaseReplicationMockDefault

func (_class SRClass) DisableDatabaseReplicationMock(sessionID SessionRef, sr SRRef) (_err error) {
	return SRClassDisableDatabaseReplicationMockedCallback(sessionID, sr)
}
// 
func (_class SRClass) DisableDatabaseReplication(sessionID SessionRef, sr SRRef) (_err error) {
	if IsMock {
		return _class.DisableDatabaseReplicationMock(sessionID, sr)
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


func SRClassEnableDatabaseReplicationMockDefault(sessionID SessionRef, sr SRRef) (_err error) {
	log.Println("SR.EnableDatabaseReplication not mocked")
	_err = errors.New("SR.EnableDatabaseReplication not mocked")
	return
}

var SRClassEnableDatabaseReplicationMockedCallback = SRClassEnableDatabaseReplicationMockDefault

func (_class SRClass) EnableDatabaseReplicationMock(sessionID SessionRef, sr SRRef) (_err error) {
	return SRClassEnableDatabaseReplicationMockedCallback(sessionID, sr)
}
// 
func (_class SRClass) EnableDatabaseReplication(sessionID SessionRef, sr SRRef) (_err error) {
	if IsMock {
		return _class.EnableDatabaseReplicationMock(sessionID, sr)
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


func SRClassAssertSupportsDatabaseReplicationMockDefault(sessionID SessionRef, sr SRRef) (_err error) {
	log.Println("SR.AssertSupportsDatabaseReplication not mocked")
	_err = errors.New("SR.AssertSupportsDatabaseReplication not mocked")
	return
}

var SRClassAssertSupportsDatabaseReplicationMockedCallback = SRClassAssertSupportsDatabaseReplicationMockDefault

func (_class SRClass) AssertSupportsDatabaseReplicationMock(sessionID SessionRef, sr SRRef) (_err error) {
	return SRClassAssertSupportsDatabaseReplicationMockedCallback(sessionID, sr)
}
// Returns successfully if the given SR supports database replication. Otherwise returns an error to explain why not.
func (_class SRClass) AssertSupportsDatabaseReplication(sessionID SessionRef, sr SRRef) (_err error) {
	if IsMock {
		return _class.AssertSupportsDatabaseReplicationMock(sessionID, sr)
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


func SRClassAssertCanHostHaStatefileMockDefault(sessionID SessionRef, sr SRRef) (_err error) {
	log.Println("SR.AssertCanHostHaStatefile not mocked")
	_err = errors.New("SR.AssertCanHostHaStatefile not mocked")
	return
}

var SRClassAssertCanHostHaStatefileMockedCallback = SRClassAssertCanHostHaStatefileMockDefault

func (_class SRClass) AssertCanHostHaStatefileMock(sessionID SessionRef, sr SRRef) (_err error) {
	return SRClassAssertCanHostHaStatefileMockedCallback(sessionID, sr)
}
// Returns successfully if the given SR can host an HA statefile. Otherwise returns an error to explain why not
func (_class SRClass) AssertCanHostHaStatefile(sessionID SessionRef, sr SRRef) (_err error) {
	if IsMock {
		return _class.AssertCanHostHaStatefileMock(sessionID, sr)
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


func SRClassSetPhysicalUtilisationMockDefault(sessionID SessionRef, self SRRef, value int) (_err error) {
	log.Println("SR.SetPhysicalUtilisation not mocked")
	_err = errors.New("SR.SetPhysicalUtilisation not mocked")
	return
}

var SRClassSetPhysicalUtilisationMockedCallback = SRClassSetPhysicalUtilisationMockDefault

func (_class SRClass) SetPhysicalUtilisationMock(sessionID SessionRef, self SRRef, value int) (_err error) {
	return SRClassSetPhysicalUtilisationMockedCallback(sessionID, self, value)
}
// Sets the SR's physical_utilisation field
func (_class SRClass) SetPhysicalUtilisation(sessionID SessionRef, self SRRef, value int) (_err error) {
	if IsMock {
		return _class.SetPhysicalUtilisationMock(sessionID, self, value)
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


func SRClassSetVirtualAllocationMockDefault(sessionID SessionRef, self SRRef, value int) (_err error) {
	log.Println("SR.SetVirtualAllocation not mocked")
	_err = errors.New("SR.SetVirtualAllocation not mocked")
	return
}

var SRClassSetVirtualAllocationMockedCallback = SRClassSetVirtualAllocationMockDefault

func (_class SRClass) SetVirtualAllocationMock(sessionID SessionRef, self SRRef, value int) (_err error) {
	return SRClassSetVirtualAllocationMockedCallback(sessionID, self, value)
}
// Sets the SR's virtual_allocation field
func (_class SRClass) SetVirtualAllocation(sessionID SessionRef, self SRRef, value int) (_err error) {
	if IsMock {
		return _class.SetVirtualAllocationMock(sessionID, self, value)
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


func SRClassSetPhysicalSizeMockDefault(sessionID SessionRef, self SRRef, value int) (_err error) {
	log.Println("SR.SetPhysicalSize not mocked")
	_err = errors.New("SR.SetPhysicalSize not mocked")
	return
}

var SRClassSetPhysicalSizeMockedCallback = SRClassSetPhysicalSizeMockDefault

func (_class SRClass) SetPhysicalSizeMock(sessionID SessionRef, self SRRef, value int) (_err error) {
	return SRClassSetPhysicalSizeMockedCallback(sessionID, self, value)
}
// Sets the SR's physical_size field
func (_class SRClass) SetPhysicalSize(sessionID SessionRef, self SRRef, value int) (_err error) {
	if IsMock {
		return _class.SetPhysicalSizeMock(sessionID, self, value)
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


func SRClassCreateNewBlobMockDefault(sessionID SessionRef, sr SRRef, name string, mimeType string, public bool) (_retval BlobRef, _err error) {
	log.Println("SR.CreateNewBlob not mocked")
	_err = errors.New("SR.CreateNewBlob not mocked")
	return
}

var SRClassCreateNewBlobMockedCallback = SRClassCreateNewBlobMockDefault

func (_class SRClass) CreateNewBlobMock(sessionID SessionRef, sr SRRef, name string, mimeType string, public bool) (_retval BlobRef, _err error) {
	return SRClassCreateNewBlobMockedCallback(sessionID, sr, name, mimeType, public)
}
// Create a placeholder for a named binary blob of data that is associated with this SR
func (_class SRClass) CreateNewBlob(sessionID SessionRef, sr SRRef, name string, mimeType string, public bool) (_retval BlobRef, _err error) {
	if IsMock {
		return _class.CreateNewBlobMock(sessionID, sr, name, mimeType, public)
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


func SRClassSetNameDescriptionMockDefault(sessionID SessionRef, sr SRRef, value string) (_err error) {
	log.Println("SR.SetNameDescription not mocked")
	_err = errors.New("SR.SetNameDescription not mocked")
	return
}

var SRClassSetNameDescriptionMockedCallback = SRClassSetNameDescriptionMockDefault

func (_class SRClass) SetNameDescriptionMock(sessionID SessionRef, sr SRRef, value string) (_err error) {
	return SRClassSetNameDescriptionMockedCallback(sessionID, sr, value)
}
// Set the name description of the SR
func (_class SRClass) SetNameDescription(sessionID SessionRef, sr SRRef, value string) (_err error) {
	if IsMock {
		return _class.SetNameDescriptionMock(sessionID, sr, value)
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


func SRClassSetNameLabelMockDefault(sessionID SessionRef, sr SRRef, value string) (_err error) {
	log.Println("SR.SetNameLabel not mocked")
	_err = errors.New("SR.SetNameLabel not mocked")
	return
}

var SRClassSetNameLabelMockedCallback = SRClassSetNameLabelMockDefault

func (_class SRClass) SetNameLabelMock(sessionID SessionRef, sr SRRef, value string) (_err error) {
	return SRClassSetNameLabelMockedCallback(sessionID, sr, value)
}
// Set the name label of the SR
func (_class SRClass) SetNameLabel(sessionID SessionRef, sr SRRef, value string) (_err error) {
	if IsMock {
		return _class.SetNameLabelMock(sessionID, sr, value)
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


func SRClassSetSharedMockDefault(sessionID SessionRef, sr SRRef, value bool) (_err error) {
	log.Println("SR.SetShared not mocked")
	_err = errors.New("SR.SetShared not mocked")
	return
}

var SRClassSetSharedMockedCallback = SRClassSetSharedMockDefault

func (_class SRClass) SetSharedMock(sessionID SessionRef, sr SRRef, value bool) (_err error) {
	return SRClassSetSharedMockedCallback(sessionID, sr, value)
}
// Sets the shared flag on the SR
func (_class SRClass) SetShared(sessionID SessionRef, sr SRRef, value bool) (_err error) {
	if IsMock {
		return _class.SetSharedMock(sessionID, sr, value)
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


func SRClassProbeMockDefault(sessionID SessionRef, host HostRef, deviceConfig map[string]string, atype string, smConfig map[string]string) (_retval string, _err error) {
	log.Println("SR.Probe not mocked")
	_err = errors.New("SR.Probe not mocked")
	return
}

var SRClassProbeMockedCallback = SRClassProbeMockDefault

func (_class SRClass) ProbeMock(sessionID SessionRef, host HostRef, deviceConfig map[string]string, atype string, smConfig map[string]string) (_retval string, _err error) {
	return SRClassProbeMockedCallback(sessionID, host, deviceConfig, atype, smConfig)
}
// Perform a backend-specific scan, using the given device_config.  If the device_config is complete, then this will return a list of the SRs present of this type on the device, if any.  If the device_config is partial, then a backend-specific scan will be performed, returning results that will guide the user in improving the device_config.
func (_class SRClass) Probe(sessionID SessionRef, host HostRef, deviceConfig map[string]string, atype string, smConfig map[string]string) (_retval string, _err error) {
	if IsMock {
		return _class.ProbeMock(sessionID, host, deviceConfig, atype, smConfig)
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


func SRClassScanMockDefault(sessionID SessionRef, sr SRRef) (_err error) {
	log.Println("SR.Scan not mocked")
	_err = errors.New("SR.Scan not mocked")
	return
}

var SRClassScanMockedCallback = SRClassScanMockDefault

func (_class SRClass) ScanMock(sessionID SessionRef, sr SRRef) (_err error) {
	return SRClassScanMockedCallback(sessionID, sr)
}
// Refreshes the list of VDIs associated with an SR
func (_class SRClass) Scan(sessionID SessionRef, sr SRRef) (_err error) {
	if IsMock {
		return _class.ScanMock(sessionID, sr)
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


func SRClassGetSupportedTypesMockDefault(sessionID SessionRef) (_retval []string, _err error) {
	log.Println("SR.GetSupportedTypes not mocked")
	_err = errors.New("SR.GetSupportedTypes not mocked")
	return
}

var SRClassGetSupportedTypesMockedCallback = SRClassGetSupportedTypesMockDefault

func (_class SRClass) GetSupportedTypesMock(sessionID SessionRef) (_retval []string, _err error) {
	return SRClassGetSupportedTypesMockedCallback(sessionID)
}
// Return a set of all the SR types supported by the system
func (_class SRClass) GetSupportedTypes(sessionID SessionRef) (_retval []string, _err error) {
	if IsMock {
		return _class.GetSupportedTypesMock(sessionID)
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


func SRClassUpdateMockDefault(sessionID SessionRef, sr SRRef) (_err error) {
	log.Println("SR.Update not mocked")
	_err = errors.New("SR.Update not mocked")
	return
}

var SRClassUpdateMockedCallback = SRClassUpdateMockDefault

func (_class SRClass) UpdateMock(sessionID SessionRef, sr SRRef) (_err error) {
	return SRClassUpdateMockedCallback(sessionID, sr)
}
// Refresh the fields on the SR object
func (_class SRClass) Update(sessionID SessionRef, sr SRRef) (_err error) {
	if IsMock {
		return _class.UpdateMock(sessionID, sr)
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


func SRClassForgetMockDefault(sessionID SessionRef, sr SRRef) (_err error) {
	log.Println("SR.Forget not mocked")
	_err = errors.New("SR.Forget not mocked")
	return
}

var SRClassForgetMockedCallback = SRClassForgetMockDefault

func (_class SRClass) ForgetMock(sessionID SessionRef, sr SRRef) (_err error) {
	return SRClassForgetMockedCallback(sessionID, sr)
}
// Removing specified SR-record from database, without attempting to remove SR from disk
//
// Errors:
//  SR_HAS_PBD - The SR is still connected to a host via a PBD. It cannot be destroyed or forgotten.
func (_class SRClass) Forget(sessionID SessionRef, sr SRRef) (_err error) {
	if IsMock {
		return _class.ForgetMock(sessionID, sr)
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


func SRClassDestroyMockDefault(sessionID SessionRef, sr SRRef) (_err error) {
	log.Println("SR.Destroy not mocked")
	_err = errors.New("SR.Destroy not mocked")
	return
}

var SRClassDestroyMockedCallback = SRClassDestroyMockDefault

func (_class SRClass) DestroyMock(sessionID SessionRef, sr SRRef) (_err error) {
	return SRClassDestroyMockedCallback(sessionID, sr)
}
// Destroy specified SR, removing SR-record from database and remove SR from disk. (In order to affect this operation the appropriate device_config is read from the specified SR's PBD on current host)
//
// Errors:
//  SR_HAS_PBD - The SR is still connected to a host via a PBD. It cannot be destroyed or forgotten.
func (_class SRClass) Destroy(sessionID SessionRef, sr SRRef) (_err error) {
	if IsMock {
		return _class.DestroyMock(sessionID, sr)
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


func SRClassMakeMockDefault(sessionID SessionRef, host HostRef, deviceConfig map[string]string, physicalSize int, nameLabel string, nameDescription string, atype string, contentType string, smConfig map[string]string) (_retval string, _err error) {
	log.Println("SR.Make not mocked")
	_err = errors.New("SR.Make not mocked")
	return
}

var SRClassMakeMockedCallback = SRClassMakeMockDefault

func (_class SRClass) MakeMock(sessionID SessionRef, host HostRef, deviceConfig map[string]string, physicalSize int, nameLabel string, nameDescription string, atype string, contentType string, smConfig map[string]string) (_retval string, _err error) {
	return SRClassMakeMockedCallback(sessionID, host, deviceConfig, physicalSize, nameLabel, nameDescription, atype, contentType, smConfig)
}
// Create a new Storage Repository on disk. This call is deprecated: use SR.create instead.
func (_class SRClass) Make(sessionID SessionRef, host HostRef, deviceConfig map[string]string, physicalSize int, nameLabel string, nameDescription string, atype string, contentType string, smConfig map[string]string) (_retval string, _err error) {
	if IsMock {
		return _class.MakeMock(sessionID, host, deviceConfig, physicalSize, nameLabel, nameDescription, atype, contentType, smConfig)
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


func SRClassIntroduceMockDefault(sessionID SessionRef, uuid string, nameLabel string, nameDescription string, atype string, contentType string, shared bool, smConfig map[string]string) (_retval SRRef, _err error) {
	log.Println("SR.Introduce not mocked")
	_err = errors.New("SR.Introduce not mocked")
	return
}

var SRClassIntroduceMockedCallback = SRClassIntroduceMockDefault

func (_class SRClass) IntroduceMock(sessionID SessionRef, uuid string, nameLabel string, nameDescription string, atype string, contentType string, shared bool, smConfig map[string]string) (_retval SRRef, _err error) {
	return SRClassIntroduceMockedCallback(sessionID, uuid, nameLabel, nameDescription, atype, contentType, shared, smConfig)
}
// Introduce a new Storage Repository into the managed system
func (_class SRClass) Introduce(sessionID SessionRef, uuid string, nameLabel string, nameDescription string, atype string, contentType string, shared bool, smConfig map[string]string) (_retval SRRef, _err error) {
	if IsMock {
		return _class.IntroduceMock(sessionID, uuid, nameLabel, nameDescription, atype, contentType, shared, smConfig)
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


func SRClassCreateMockDefault(sessionID SessionRef, host HostRef, deviceConfig map[string]string, physicalSize int, nameLabel string, nameDescription string, atype string, contentType string, shared bool, smConfig map[string]string) (_retval SRRef, _err error) {
	log.Println("SR.Create not mocked")
	_err = errors.New("SR.Create not mocked")
	return
}

var SRClassCreateMockedCallback = SRClassCreateMockDefault

func (_class SRClass) CreateMock(sessionID SessionRef, host HostRef, deviceConfig map[string]string, physicalSize int, nameLabel string, nameDescription string, atype string, contentType string, shared bool, smConfig map[string]string) (_retval SRRef, _err error) {
	return SRClassCreateMockedCallback(sessionID, host, deviceConfig, physicalSize, nameLabel, nameDescription, atype, contentType, shared, smConfig)
}
// Create a new Storage Repository and introduce it into the managed system, creating both SR record and PBD record to attach it to current host (with specified device_config parameters)
//
// Errors:
//  SR_UNKNOWN_DRIVER - The SR could not be connected because the driver was not recognised.
func (_class SRClass) Create(sessionID SessionRef, host HostRef, deviceConfig map[string]string, physicalSize int, nameLabel string, nameDescription string, atype string, contentType string, shared bool, smConfig map[string]string) (_retval SRRef, _err error) {
	if IsMock {
		return _class.CreateMock(sessionID, host, deviceConfig, physicalSize, nameLabel, nameDescription, atype, contentType, shared, smConfig)
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


func SRClassRemoveFromSmConfigMockDefault(sessionID SessionRef, self SRRef, key string) (_err error) {
	log.Println("SR.RemoveFromSmConfig not mocked")
	_err = errors.New("SR.RemoveFromSmConfig not mocked")
	return
}

var SRClassRemoveFromSmConfigMockedCallback = SRClassRemoveFromSmConfigMockDefault

func (_class SRClass) RemoveFromSmConfigMock(sessionID SessionRef, self SRRef, key string) (_err error) {
	return SRClassRemoveFromSmConfigMockedCallback(sessionID, self, key)
}
// Remove the given key and its corresponding value from the sm_config field of the given SR.  If the key is not in that Map, then do nothing.
func (_class SRClass) RemoveFromSmConfig(sessionID SessionRef, self SRRef, key string) (_err error) {
	if IsMock {
		return _class.RemoveFromSmConfigMock(sessionID, self, key)
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


func SRClassAddToSmConfigMockDefault(sessionID SessionRef, self SRRef, key string, value string) (_err error) {
	log.Println("SR.AddToSmConfig not mocked")
	_err = errors.New("SR.AddToSmConfig not mocked")
	return
}

var SRClassAddToSmConfigMockedCallback = SRClassAddToSmConfigMockDefault

func (_class SRClass) AddToSmConfigMock(sessionID SessionRef, self SRRef, key string, value string) (_err error) {
	return SRClassAddToSmConfigMockedCallback(sessionID, self, key, value)
}
// Add the given key-value pair to the sm_config field of the given SR.
func (_class SRClass) AddToSmConfig(sessionID SessionRef, self SRRef, key string, value string) (_err error) {
	if IsMock {
		return _class.AddToSmConfigMock(sessionID, self, key, value)
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


func SRClassSetSmConfigMockDefault(sessionID SessionRef, self SRRef, value map[string]string) (_err error) {
	log.Println("SR.SetSmConfig not mocked")
	_err = errors.New("SR.SetSmConfig not mocked")
	return
}

var SRClassSetSmConfigMockedCallback = SRClassSetSmConfigMockDefault

func (_class SRClass) SetSmConfigMock(sessionID SessionRef, self SRRef, value map[string]string) (_err error) {
	return SRClassSetSmConfigMockedCallback(sessionID, self, value)
}
// Set the sm_config field of the given SR.
func (_class SRClass) SetSmConfig(sessionID SessionRef, self SRRef, value map[string]string) (_err error) {
	if IsMock {
		return _class.SetSmConfigMock(sessionID, self, value)
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


func SRClassRemoveTagsMockDefault(sessionID SessionRef, self SRRef, value string) (_err error) {
	log.Println("SR.RemoveTags not mocked")
	_err = errors.New("SR.RemoveTags not mocked")
	return
}

var SRClassRemoveTagsMockedCallback = SRClassRemoveTagsMockDefault

func (_class SRClass) RemoveTagsMock(sessionID SessionRef, self SRRef, value string) (_err error) {
	return SRClassRemoveTagsMockedCallback(sessionID, self, value)
}
// Remove the given value from the tags field of the given SR.  If the value is not in that Set, then do nothing.
func (_class SRClass) RemoveTags(sessionID SessionRef, self SRRef, value string) (_err error) {
	if IsMock {
		return _class.RemoveTagsMock(sessionID, self, value)
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


func SRClassAddTagsMockDefault(sessionID SessionRef, self SRRef, value string) (_err error) {
	log.Println("SR.AddTags not mocked")
	_err = errors.New("SR.AddTags not mocked")
	return
}

var SRClassAddTagsMockedCallback = SRClassAddTagsMockDefault

func (_class SRClass) AddTagsMock(sessionID SessionRef, self SRRef, value string) (_err error) {
	return SRClassAddTagsMockedCallback(sessionID, self, value)
}
// Add the given value to the tags field of the given SR.  If the value is already in that Set, then do nothing.
func (_class SRClass) AddTags(sessionID SessionRef, self SRRef, value string) (_err error) {
	if IsMock {
		return _class.AddTagsMock(sessionID, self, value)
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


func SRClassSetTagsMockDefault(sessionID SessionRef, self SRRef, value []string) (_err error) {
	log.Println("SR.SetTags not mocked")
	_err = errors.New("SR.SetTags not mocked")
	return
}

var SRClassSetTagsMockedCallback = SRClassSetTagsMockDefault

func (_class SRClass) SetTagsMock(sessionID SessionRef, self SRRef, value []string) (_err error) {
	return SRClassSetTagsMockedCallback(sessionID, self, value)
}
// Set the tags field of the given SR.
func (_class SRClass) SetTags(sessionID SessionRef, self SRRef, value []string) (_err error) {
	if IsMock {
		return _class.SetTagsMock(sessionID, self, value)
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


func SRClassRemoveFromOtherConfigMockDefault(sessionID SessionRef, self SRRef, key string) (_err error) {
	log.Println("SR.RemoveFromOtherConfig not mocked")
	_err = errors.New("SR.RemoveFromOtherConfig not mocked")
	return
}

var SRClassRemoveFromOtherConfigMockedCallback = SRClassRemoveFromOtherConfigMockDefault

func (_class SRClass) RemoveFromOtherConfigMock(sessionID SessionRef, self SRRef, key string) (_err error) {
	return SRClassRemoveFromOtherConfigMockedCallback(sessionID, self, key)
}
// Remove the given key and its corresponding value from the other_config field of the given SR.  If the key is not in that Map, then do nothing.
func (_class SRClass) RemoveFromOtherConfig(sessionID SessionRef, self SRRef, key string) (_err error) {
	if IsMock {
		return _class.RemoveFromOtherConfigMock(sessionID, self, key)
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


func SRClassAddToOtherConfigMockDefault(sessionID SessionRef, self SRRef, key string, value string) (_err error) {
	log.Println("SR.AddToOtherConfig not mocked")
	_err = errors.New("SR.AddToOtherConfig not mocked")
	return
}

var SRClassAddToOtherConfigMockedCallback = SRClassAddToOtherConfigMockDefault

func (_class SRClass) AddToOtherConfigMock(sessionID SessionRef, self SRRef, key string, value string) (_err error) {
	return SRClassAddToOtherConfigMockedCallback(sessionID, self, key, value)
}
// Add the given key-value pair to the other_config field of the given SR.
func (_class SRClass) AddToOtherConfig(sessionID SessionRef, self SRRef, key string, value string) (_err error) {
	if IsMock {
		return _class.AddToOtherConfigMock(sessionID, self, key, value)
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


func SRClassSetOtherConfigMockDefault(sessionID SessionRef, self SRRef, value map[string]string) (_err error) {
	log.Println("SR.SetOtherConfig not mocked")
	_err = errors.New("SR.SetOtherConfig not mocked")
	return
}

var SRClassSetOtherConfigMockedCallback = SRClassSetOtherConfigMockDefault

func (_class SRClass) SetOtherConfigMock(sessionID SessionRef, self SRRef, value map[string]string) (_err error) {
	return SRClassSetOtherConfigMockedCallback(sessionID, self, value)
}
// Set the other_config field of the given SR.
func (_class SRClass) SetOtherConfig(sessionID SessionRef, self SRRef, value map[string]string) (_err error) {
	if IsMock {
		return _class.SetOtherConfigMock(sessionID, self, value)
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


func SRClassGetIsToolsSrMockDefault(sessionID SessionRef, self SRRef) (_retval bool, _err error) {
	log.Println("SR.GetIsToolsSr not mocked")
	_err = errors.New("SR.GetIsToolsSr not mocked")
	return
}

var SRClassGetIsToolsSrMockedCallback = SRClassGetIsToolsSrMockDefault

func (_class SRClass) GetIsToolsSrMock(sessionID SessionRef, self SRRef) (_retval bool, _err error) {
	return SRClassGetIsToolsSrMockedCallback(sessionID, self)
}
// Get the is_tools_sr field of the given SR.
func (_class SRClass) GetIsToolsSr(sessionID SessionRef, self SRRef) (_retval bool, _err error) {
	if IsMock {
		return _class.GetIsToolsSrMock(sessionID, self)
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


func SRClassGetClusteredMockDefault(sessionID SessionRef, self SRRef) (_retval bool, _err error) {
	log.Println("SR.GetClustered not mocked")
	_err = errors.New("SR.GetClustered not mocked")
	return
}

var SRClassGetClusteredMockedCallback = SRClassGetClusteredMockDefault

func (_class SRClass) GetClusteredMock(sessionID SessionRef, self SRRef) (_retval bool, _err error) {
	return SRClassGetClusteredMockedCallback(sessionID, self)
}
// Get the clustered field of the given SR.
func (_class SRClass) GetClustered(sessionID SessionRef, self SRRef) (_retval bool, _err error) {
	if IsMock {
		return _class.GetClusteredMock(sessionID, self)
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


func SRClassGetIntroducedByMockDefault(sessionID SessionRef, self SRRef) (_retval DRTaskRef, _err error) {
	log.Println("SR.GetIntroducedBy not mocked")
	_err = errors.New("SR.GetIntroducedBy not mocked")
	return
}

var SRClassGetIntroducedByMockedCallback = SRClassGetIntroducedByMockDefault

func (_class SRClass) GetIntroducedByMock(sessionID SessionRef, self SRRef) (_retval DRTaskRef, _err error) {
	return SRClassGetIntroducedByMockedCallback(sessionID, self)
}
// Get the introduced_by field of the given SR.
func (_class SRClass) GetIntroducedBy(sessionID SessionRef, self SRRef) (_retval DRTaskRef, _err error) {
	if IsMock {
		return _class.GetIntroducedByMock(sessionID, self)
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


func SRClassGetLocalCacheEnabledMockDefault(sessionID SessionRef, self SRRef) (_retval bool, _err error) {
	log.Println("SR.GetLocalCacheEnabled not mocked")
	_err = errors.New("SR.GetLocalCacheEnabled not mocked")
	return
}

var SRClassGetLocalCacheEnabledMockedCallback = SRClassGetLocalCacheEnabledMockDefault

func (_class SRClass) GetLocalCacheEnabledMock(sessionID SessionRef, self SRRef) (_retval bool, _err error) {
	return SRClassGetLocalCacheEnabledMockedCallback(sessionID, self)
}
// Get the local_cache_enabled field of the given SR.
func (_class SRClass) GetLocalCacheEnabled(sessionID SessionRef, self SRRef) (_retval bool, _err error) {
	if IsMock {
		return _class.GetLocalCacheEnabledMock(sessionID, self)
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


func SRClassGetBlobsMockDefault(sessionID SessionRef, self SRRef) (_retval map[string]BlobRef, _err error) {
	log.Println("SR.GetBlobs not mocked")
	_err = errors.New("SR.GetBlobs not mocked")
	return
}

var SRClassGetBlobsMockedCallback = SRClassGetBlobsMockDefault

func (_class SRClass) GetBlobsMock(sessionID SessionRef, self SRRef) (_retval map[string]BlobRef, _err error) {
	return SRClassGetBlobsMockedCallback(sessionID, self)
}
// Get the blobs field of the given SR.
func (_class SRClass) GetBlobs(sessionID SessionRef, self SRRef) (_retval map[string]BlobRef, _err error) {
	if IsMock {
		return _class.GetBlobsMock(sessionID, self)
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


func SRClassGetSmConfigMockDefault(sessionID SessionRef, self SRRef) (_retval map[string]string, _err error) {
	log.Println("SR.GetSmConfig not mocked")
	_err = errors.New("SR.GetSmConfig not mocked")
	return
}

var SRClassGetSmConfigMockedCallback = SRClassGetSmConfigMockDefault

func (_class SRClass) GetSmConfigMock(sessionID SessionRef, self SRRef) (_retval map[string]string, _err error) {
	return SRClassGetSmConfigMockedCallback(sessionID, self)
}
// Get the sm_config field of the given SR.
func (_class SRClass) GetSmConfig(sessionID SessionRef, self SRRef) (_retval map[string]string, _err error) {
	if IsMock {
		return _class.GetSmConfigMock(sessionID, self)
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


func SRClassGetTagsMockDefault(sessionID SessionRef, self SRRef) (_retval []string, _err error) {
	log.Println("SR.GetTags not mocked")
	_err = errors.New("SR.GetTags not mocked")
	return
}

var SRClassGetTagsMockedCallback = SRClassGetTagsMockDefault

func (_class SRClass) GetTagsMock(sessionID SessionRef, self SRRef) (_retval []string, _err error) {
	return SRClassGetTagsMockedCallback(sessionID, self)
}
// Get the tags field of the given SR.
func (_class SRClass) GetTags(sessionID SessionRef, self SRRef) (_retval []string, _err error) {
	if IsMock {
		return _class.GetTagsMock(sessionID, self)
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


func SRClassGetOtherConfigMockDefault(sessionID SessionRef, self SRRef) (_retval map[string]string, _err error) {
	log.Println("SR.GetOtherConfig not mocked")
	_err = errors.New("SR.GetOtherConfig not mocked")
	return
}

var SRClassGetOtherConfigMockedCallback = SRClassGetOtherConfigMockDefault

func (_class SRClass) GetOtherConfigMock(sessionID SessionRef, self SRRef) (_retval map[string]string, _err error) {
	return SRClassGetOtherConfigMockedCallback(sessionID, self)
}
// Get the other_config field of the given SR.
func (_class SRClass) GetOtherConfig(sessionID SessionRef, self SRRef) (_retval map[string]string, _err error) {
	if IsMock {
		return _class.GetOtherConfigMock(sessionID, self)
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


func SRClassGetSharedMockDefault(sessionID SessionRef, self SRRef) (_retval bool, _err error) {
	log.Println("SR.GetShared not mocked")
	_err = errors.New("SR.GetShared not mocked")
	return
}

var SRClassGetSharedMockedCallback = SRClassGetSharedMockDefault

func (_class SRClass) GetSharedMock(sessionID SessionRef, self SRRef) (_retval bool, _err error) {
	return SRClassGetSharedMockedCallback(sessionID, self)
}
// Get the shared field of the given SR.
func (_class SRClass) GetShared(sessionID SessionRef, self SRRef) (_retval bool, _err error) {
	if IsMock {
		return _class.GetSharedMock(sessionID, self)
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


func SRClassGetContentTypeMockDefault(sessionID SessionRef, self SRRef) (_retval string, _err error) {
	log.Println("SR.GetContentType not mocked")
	_err = errors.New("SR.GetContentType not mocked")
	return
}

var SRClassGetContentTypeMockedCallback = SRClassGetContentTypeMockDefault

func (_class SRClass) GetContentTypeMock(sessionID SessionRef, self SRRef) (_retval string, _err error) {
	return SRClassGetContentTypeMockedCallback(sessionID, self)
}
// Get the content_type field of the given SR.
func (_class SRClass) GetContentType(sessionID SessionRef, self SRRef) (_retval string, _err error) {
	if IsMock {
		return _class.GetContentTypeMock(sessionID, self)
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


func SRClassGetTypeMockDefault(sessionID SessionRef, self SRRef) (_retval string, _err error) {
	log.Println("SR.GetType not mocked")
	_err = errors.New("SR.GetType not mocked")
	return
}

var SRClassGetTypeMockedCallback = SRClassGetTypeMockDefault

func (_class SRClass) GetTypeMock(sessionID SessionRef, self SRRef) (_retval string, _err error) {
	return SRClassGetTypeMockedCallback(sessionID, self)
}
// Get the type field of the given SR.
func (_class SRClass) GetType(sessionID SessionRef, self SRRef) (_retval string, _err error) {
	if IsMock {
		return _class.GetTypeMock(sessionID, self)
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


func SRClassGetPhysicalSizeMockDefault(sessionID SessionRef, self SRRef) (_retval int, _err error) {
	log.Println("SR.GetPhysicalSize not mocked")
	_err = errors.New("SR.GetPhysicalSize not mocked")
	return
}

var SRClassGetPhysicalSizeMockedCallback = SRClassGetPhysicalSizeMockDefault

func (_class SRClass) GetPhysicalSizeMock(sessionID SessionRef, self SRRef) (_retval int, _err error) {
	return SRClassGetPhysicalSizeMockedCallback(sessionID, self)
}
// Get the physical_size field of the given SR.
func (_class SRClass) GetPhysicalSize(sessionID SessionRef, self SRRef) (_retval int, _err error) {
	if IsMock {
		return _class.GetPhysicalSizeMock(sessionID, self)
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


func SRClassGetPhysicalUtilisationMockDefault(sessionID SessionRef, self SRRef) (_retval int, _err error) {
	log.Println("SR.GetPhysicalUtilisation not mocked")
	_err = errors.New("SR.GetPhysicalUtilisation not mocked")
	return
}

var SRClassGetPhysicalUtilisationMockedCallback = SRClassGetPhysicalUtilisationMockDefault

func (_class SRClass) GetPhysicalUtilisationMock(sessionID SessionRef, self SRRef) (_retval int, _err error) {
	return SRClassGetPhysicalUtilisationMockedCallback(sessionID, self)
}
// Get the physical_utilisation field of the given SR.
func (_class SRClass) GetPhysicalUtilisation(sessionID SessionRef, self SRRef) (_retval int, _err error) {
	if IsMock {
		return _class.GetPhysicalUtilisationMock(sessionID, self)
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


func SRClassGetVirtualAllocationMockDefault(sessionID SessionRef, self SRRef) (_retval int, _err error) {
	log.Println("SR.GetVirtualAllocation not mocked")
	_err = errors.New("SR.GetVirtualAllocation not mocked")
	return
}

var SRClassGetVirtualAllocationMockedCallback = SRClassGetVirtualAllocationMockDefault

func (_class SRClass) GetVirtualAllocationMock(sessionID SessionRef, self SRRef) (_retval int, _err error) {
	return SRClassGetVirtualAllocationMockedCallback(sessionID, self)
}
// Get the virtual_allocation field of the given SR.
func (_class SRClass) GetVirtualAllocation(sessionID SessionRef, self SRRef) (_retval int, _err error) {
	if IsMock {
		return _class.GetVirtualAllocationMock(sessionID, self)
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


func SRClassGetPBDsMockDefault(sessionID SessionRef, self SRRef) (_retval []PBDRef, _err error) {
	log.Println("SR.GetPBDs not mocked")
	_err = errors.New("SR.GetPBDs not mocked")
	return
}

var SRClassGetPBDsMockedCallback = SRClassGetPBDsMockDefault

func (_class SRClass) GetPBDsMock(sessionID SessionRef, self SRRef) (_retval []PBDRef, _err error) {
	return SRClassGetPBDsMockedCallback(sessionID, self)
}
// Get the PBDs field of the given SR.
func (_class SRClass) GetPBDs(sessionID SessionRef, self SRRef) (_retval []PBDRef, _err error) {
	if IsMock {
		return _class.GetPBDsMock(sessionID, self)
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


func SRClassGetVDIsMockDefault(sessionID SessionRef, self SRRef) (_retval []VDIRef, _err error) {
	log.Println("SR.GetVDIs not mocked")
	_err = errors.New("SR.GetVDIs not mocked")
	return
}

var SRClassGetVDIsMockedCallback = SRClassGetVDIsMockDefault

func (_class SRClass) GetVDIsMock(sessionID SessionRef, self SRRef) (_retval []VDIRef, _err error) {
	return SRClassGetVDIsMockedCallback(sessionID, self)
}
// Get the VDIs field of the given SR.
func (_class SRClass) GetVDIs(sessionID SessionRef, self SRRef) (_retval []VDIRef, _err error) {
	if IsMock {
		return _class.GetVDIsMock(sessionID, self)
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


func SRClassGetCurrentOperationsMockDefault(sessionID SessionRef, self SRRef) (_retval map[string]StorageOperations, _err error) {
	log.Println("SR.GetCurrentOperations not mocked")
	_err = errors.New("SR.GetCurrentOperations not mocked")
	return
}

var SRClassGetCurrentOperationsMockedCallback = SRClassGetCurrentOperationsMockDefault

func (_class SRClass) GetCurrentOperationsMock(sessionID SessionRef, self SRRef) (_retval map[string]StorageOperations, _err error) {
	return SRClassGetCurrentOperationsMockedCallback(sessionID, self)
}
// Get the current_operations field of the given SR.
func (_class SRClass) GetCurrentOperations(sessionID SessionRef, self SRRef) (_retval map[string]StorageOperations, _err error) {
	if IsMock {
		return _class.GetCurrentOperationsMock(sessionID, self)
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


func SRClassGetAllowedOperationsMockDefault(sessionID SessionRef, self SRRef) (_retval []StorageOperations, _err error) {
	log.Println("SR.GetAllowedOperations not mocked")
	_err = errors.New("SR.GetAllowedOperations not mocked")
	return
}

var SRClassGetAllowedOperationsMockedCallback = SRClassGetAllowedOperationsMockDefault

func (_class SRClass) GetAllowedOperationsMock(sessionID SessionRef, self SRRef) (_retval []StorageOperations, _err error) {
	return SRClassGetAllowedOperationsMockedCallback(sessionID, self)
}
// Get the allowed_operations field of the given SR.
func (_class SRClass) GetAllowedOperations(sessionID SessionRef, self SRRef) (_retval []StorageOperations, _err error) {
	if IsMock {
		return _class.GetAllowedOperationsMock(sessionID, self)
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


func SRClassGetNameDescriptionMockDefault(sessionID SessionRef, self SRRef) (_retval string, _err error) {
	log.Println("SR.GetNameDescription not mocked")
	_err = errors.New("SR.GetNameDescription not mocked")
	return
}

var SRClassGetNameDescriptionMockedCallback = SRClassGetNameDescriptionMockDefault

func (_class SRClass) GetNameDescriptionMock(sessionID SessionRef, self SRRef) (_retval string, _err error) {
	return SRClassGetNameDescriptionMockedCallback(sessionID, self)
}
// Get the name/description field of the given SR.
func (_class SRClass) GetNameDescription(sessionID SessionRef, self SRRef) (_retval string, _err error) {
	if IsMock {
		return _class.GetNameDescriptionMock(sessionID, self)
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


func SRClassGetNameLabelMockDefault(sessionID SessionRef, self SRRef) (_retval string, _err error) {
	log.Println("SR.GetNameLabel not mocked")
	_err = errors.New("SR.GetNameLabel not mocked")
	return
}

var SRClassGetNameLabelMockedCallback = SRClassGetNameLabelMockDefault

func (_class SRClass) GetNameLabelMock(sessionID SessionRef, self SRRef) (_retval string, _err error) {
	return SRClassGetNameLabelMockedCallback(sessionID, self)
}
// Get the name/label field of the given SR.
func (_class SRClass) GetNameLabel(sessionID SessionRef, self SRRef) (_retval string, _err error) {
	if IsMock {
		return _class.GetNameLabelMock(sessionID, self)
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


func SRClassGetUUIDMockDefault(sessionID SessionRef, self SRRef) (_retval string, _err error) {
	log.Println("SR.GetUUID not mocked")
	_err = errors.New("SR.GetUUID not mocked")
	return
}

var SRClassGetUUIDMockedCallback = SRClassGetUUIDMockDefault

func (_class SRClass) GetUUIDMock(sessionID SessionRef, self SRRef) (_retval string, _err error) {
	return SRClassGetUUIDMockedCallback(sessionID, self)
}
// Get the uuid field of the given SR.
func (_class SRClass) GetUUID(sessionID SessionRef, self SRRef) (_retval string, _err error) {
	if IsMock {
		return _class.GetUUIDMock(sessionID, self)
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


func SRClassGetByNameLabelMockDefault(sessionID SessionRef, label string) (_retval []SRRef, _err error) {
	log.Println("SR.GetByNameLabel not mocked")
	_err = errors.New("SR.GetByNameLabel not mocked")
	return
}

var SRClassGetByNameLabelMockedCallback = SRClassGetByNameLabelMockDefault

func (_class SRClass) GetByNameLabelMock(sessionID SessionRef, label string) (_retval []SRRef, _err error) {
	return SRClassGetByNameLabelMockedCallback(sessionID, label)
}
// Get all the SR instances with the given label.
func (_class SRClass) GetByNameLabel(sessionID SessionRef, label string) (_retval []SRRef, _err error) {
	if IsMock {
		return _class.GetByNameLabelMock(sessionID, label)
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


func SRClassGetByUUIDMockDefault(sessionID SessionRef, uuid string) (_retval SRRef, _err error) {
	log.Println("SR.GetByUUID not mocked")
	_err = errors.New("SR.GetByUUID not mocked")
	return
}

var SRClassGetByUUIDMockedCallback = SRClassGetByUUIDMockDefault

func (_class SRClass) GetByUUIDMock(sessionID SessionRef, uuid string) (_retval SRRef, _err error) {
	return SRClassGetByUUIDMockedCallback(sessionID, uuid)
}
// Get a reference to the SR instance with the specified UUID.
func (_class SRClass) GetByUUID(sessionID SessionRef, uuid string) (_retval SRRef, _err error) {
	if IsMock {
		return _class.GetByUUIDMock(sessionID, uuid)
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


func SRClassGetRecordMockDefault(sessionID SessionRef, self SRRef) (_retval SRRecord, _err error) {
	log.Println("SR.GetRecord not mocked")
	_err = errors.New("SR.GetRecord not mocked")
	return
}

var SRClassGetRecordMockedCallback = SRClassGetRecordMockDefault

func (_class SRClass) GetRecordMock(sessionID SessionRef, self SRRef) (_retval SRRecord, _err error) {
	return SRClassGetRecordMockedCallback(sessionID, self)
}
// Get a record containing the current state of the given SR.
func (_class SRClass) GetRecord(sessionID SessionRef, self SRRef) (_retval SRRecord, _err error) {
	if IsMock {
		return _class.GetRecordMock(sessionID, self)
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
