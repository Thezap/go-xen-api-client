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

type VdiOperations string

const (
  // Cloning the VDI
	VdiOperationsClone VdiOperations = "clone"
  // Copying the VDI
	VdiOperationsCopy VdiOperations = "copy"
  // Resizing the VDI
	VdiOperationsResize VdiOperations = "resize"
  // Resizing the VDI which may or may not be online
	VdiOperationsResizeOnline VdiOperations = "resize_online"
  // Snapshotting the VDI
	VdiOperationsSnapshot VdiOperations = "snapshot"
  // Mirroring the VDI
	VdiOperationsMirror VdiOperations = "mirror"
  // Destroying the VDI
	VdiOperationsDestroy VdiOperations = "destroy"
  // Forget about the VDI
	VdiOperationsForget VdiOperations = "forget"
  // Refreshing the fields of the VDI
	VdiOperationsUpdate VdiOperations = "update"
  // Forcibly unlocking the VDI
	VdiOperationsForceUnlock VdiOperations = "force_unlock"
  // Generating static configuration
	VdiOperationsGenerateConfig VdiOperations = "generate_config"
  // Enabling changed block tracking for a VDI
	VdiOperationsEnableCbt VdiOperations = "enable_cbt"
  // Disabling changed block tracking for a VDI
	VdiOperationsDisableCbt VdiOperations = "disable_cbt"
  // Deleting the data of the VDI
	VdiOperationsDataDestroy VdiOperations = "data_destroy"
  // Exporting a bitmap that shows the changed blocks between two VDIs
	VdiOperationsListChangedBlocks VdiOperations = "list_changed_blocks"
  // Setting the on_boot field of the VDI
	VdiOperationsSetOnBoot VdiOperations = "set_on_boot"
  // Operations on this VDI are temporarily blocked
	VdiOperationsBlocked VdiOperations = "blocked"
)

type VdiType string

const (
  // a disk that may be replaced on upgrade
	VdiTypeSystem VdiType = "system"
  // a disk that is always preserved on upgrade
	VdiTypeUser VdiType = "user"
  // a disk that may be reformatted on upgrade
	VdiTypeEphemeral VdiType = "ephemeral"
  // a disk that stores a suspend image
	VdiTypeSuspend VdiType = "suspend"
  // a disk that stores VM crashdump information
	VdiTypeCrashdump VdiType = "crashdump"
  // a disk used for HA storage heartbeating
	VdiTypeHaStatefile VdiType = "ha_statefile"
  // a disk used for HA Pool metadata
	VdiTypeMetadata VdiType = "metadata"
  // a disk used for a general metadata redo-log
	VdiTypeRedoLog VdiType = "redo_log"
  // a disk that stores SR-level RRDs
	VdiTypeRrd VdiType = "rrd"
  // a disk that stores PVS cache data
	VdiTypePvsCache VdiType = "pvs_cache"
  // Metadata about a snapshot VDI that has been deleted: the set of blocks that changed between some previous version of the disk and the version tracked by the snapshot.
	VdiTypeCbtMetadata VdiType = "cbt_metadata"
)

type OnBoot string

const (
  // When a VM containing this VDI is started, the contents of the VDI are reset to the state they were in when this flag was last set.
	OnBootReset OnBoot = "reset"
  // Standard behaviour.
	OnBootPersist OnBoot = "persist"
)

type VDIRecord struct {
  // Unique identifier/object reference
	UUID string
  // a human-readable name
	NameLabel string
  // a notes field containing human-readable description
	NameDescription string
  // list of the operations allowed in this state. This list is advisory only and the server state may have changed by the time this field is read by a client.
	AllowedOperations []VdiOperations
  // links each of the running tasks using this object (by reference) to a current_operation enum which describes the nature of the task.
	CurrentOperations map[string]VdiOperations
  // storage repository in which the VDI resides
	SR SRRef
  // list of vbds that refer to this disk
	VBDs []VBDRef
  // list of crash dumps that refer to this disk
	CrashDumps []CrashdumpRef
  // size of disk as presented to the guest (in bytes). Note that, depending on storage backend type, requested size may not be respected exactly
	VirtualSize int
  // amount of physical space that the disk image is currently taking up on the storage repository (in bytes)
	PhysicalUtilisation int
  // type of the VDI
	Type VdiType
  // true if this disk may be shared
	Sharable bool
  // true if this disk may ONLY be mounted read-only
	ReadOnly bool
  // additional configuration
	OtherConfig map[string]string
  // true if this disk is locked at the storage level
	StorageLock bool
  // location information
	Location string
  // 
	Managed bool
  // true if SR scan operation reported this VDI as not present on disk
	Missing bool
  // This field is always null. Deprecated
	Parent VDIRef
  // data to be inserted into the xenstore tree (/local/domain/0/backend/vbd/<domid>/<device-id>/sm-data) after the VDI is attached. This is generally set by the SM backends on vdi_attach.
	XenstoreData map[string]string
  // SM dependent data
	SmConfig map[string]string
  // true if this is a snapshot.
	IsASnapshot bool
  // Ref pointing to the VDI this snapshot is of.
	SnapshotOf VDIRef
  // List pointing to all the VDIs snapshots.
	Snapshots []VDIRef
  // Date/time when this snapshot was created.
	SnapshotTime time.Time
  // user-specified tags for categorization purposes
	Tags []string
  // true if this VDI is to be cached in the local cache SR
	AllowCaching bool
  // The behaviour of this VDI on a VM boot
	OnBoot OnBoot
  // The pool whose metadata is contained in this VDI
	MetadataOfPool PoolRef
  // Whether this VDI contains the latest known accessible metadata for the pool
	MetadataLatest bool
  // Whether this VDI is a Tools ISO
	IsToolsIso bool
  // True if changed blocks are tracked for this VDI
	CbtEnabled bool
}

type VDIRef string

// A virtual disk image
type VDIClass struct {
	client *Client
}


func VDIClassGetAllRecordsMockDefault(sessionID SessionRef) (_retval map[VDIRef]VDIRecord, _err error) {
	log.Println("VDI.GetAllRecords not mocked")
	_err = errors.New("VDI.GetAllRecords not mocked")
	return
}

var VDIClassGetAllRecordsMockedCallback = VDIClassGetAllRecordsMockDefault

func (_class VDIClass) GetAllRecordsMock(sessionID SessionRef) (_retval map[VDIRef]VDIRecord, _err error) {
	return VDIClassGetAllRecordsMockedCallback(sessionID)
}
// Return a map of VDI references to VDI records for all VDIs known to the system.
func (_class VDIClass) GetAllRecords(sessionID SessionRef) (_retval map[VDIRef]VDIRecord, _err error) {
	if IsMock {
		return _class.GetAllRecordsMock(sessionID)
	}	
	_method := "VDI.get_all_records"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVDIRefToVDIRecordMapToGo(_method + " -> ", _result.Value)
	return
}


func VDIClassGetAllMockDefault(sessionID SessionRef) (_retval []VDIRef, _err error) {
	log.Println("VDI.GetAll not mocked")
	_err = errors.New("VDI.GetAll not mocked")
	return
}

var VDIClassGetAllMockedCallback = VDIClassGetAllMockDefault

func (_class VDIClass) GetAllMock(sessionID SessionRef) (_retval []VDIRef, _err error) {
	return VDIClassGetAllMockedCallback(sessionID)
}
// Return a list of all the VDIs known to the system.
func (_class VDIClass) GetAll(sessionID SessionRef) (_retval []VDIRef, _err error) {
	if IsMock {
		return _class.GetAllMock(sessionID)
	}	
	_method := "VDI.get_all"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVDIRefSetToGo(_method + " -> ", _result.Value)
	return
}


func VDIClassGetNbdInfoMockDefault(sessionID SessionRef, self VDIRef) (_retval []VdiNbdServerInfoRecord, _err error) {
	log.Println("VDI.GetNbdInfo not mocked")
	_err = errors.New("VDI.GetNbdInfo not mocked")
	return
}

var VDIClassGetNbdInfoMockedCallback = VDIClassGetNbdInfoMockDefault

func (_class VDIClass) GetNbdInfoMock(sessionID SessionRef, self VDIRef) (_retval []VdiNbdServerInfoRecord, _err error) {
	return VDIClassGetNbdInfoMockedCallback(sessionID, self)
}
// Get details specifying how to access this VDI via a Network Block Device server. For each of a set of NBD server addresses on which the VDI is available, the return value set contains a vdi_nbd_server_info object that contains an exportname to request once the NBD connection is established, and connection details for the address. An empty list is returned if there is no network that has a PIF on a host with access to the relevant SR, or if no such network has been assigned an NBD-related purpose in its purpose field. To access the given VDI, any of the vdi_nbd_server_info objects can be used to make a connection to a server, and then the VDI will be available by requesting the exportname.
//
// Errors:
//  VDI_INCOMPATIBLE_TYPE - This operation cannot be performed because the specified VDI is of an incompatible type (eg: an HA statefile cannot be attached to a guest)
func (_class VDIClass) GetNbdInfo(sessionID SessionRef, self VDIRef) (_retval []VdiNbdServerInfoRecord, _err error) {
	if IsMock {
		return _class.GetNbdInfoMock(sessionID, self)
	}	
	_method := "VDI.get_nbd_info"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVdiNbdServerInfoRecordSetToGo(_method + " -> ", _result.Value)
	return
}


func VDIClassListChangedBlocksMockDefault(sessionID SessionRef, vdiFrom VDIRef, vdiTo VDIRef) (_retval string, _err error) {
	log.Println("VDI.ListChangedBlocks not mocked")
	_err = errors.New("VDI.ListChangedBlocks not mocked")
	return
}

var VDIClassListChangedBlocksMockedCallback = VDIClassListChangedBlocksMockDefault

func (_class VDIClass) ListChangedBlocksMock(sessionID SessionRef, vdiFrom VDIRef, vdiTo VDIRef) (_retval string, _err error) {
	return VDIClassListChangedBlocksMockedCallback(sessionID, vdiFrom, vdiTo)
}
// Compare two VDIs in 64k block increments and report which blocks differ. This operation is not allowed when vdi_to is attached to a VM.
//
// Errors:
//  SR_OPERATION_NOT_SUPPORTED - The SR backend does not support the operation (check the SR's allowed operations)
//  VDI_MISSING - This operation cannot be performed because the specified VDI could not be found on the storage substrate
//  SR_NOT_ATTACHED - The SR is not attached.
//  SR_HAS_NO_PBDS - The SR has no attached PBDs
//  VDI_IN_USE - This operation cannot be performed because this VDI is in use by some other operation
func (_class VDIClass) ListChangedBlocks(sessionID SessionRef, vdiFrom VDIRef, vdiTo VDIRef) (_retval string, _err error) {
	if IsMock {
		return _class.ListChangedBlocksMock(sessionID, vdiFrom, vdiTo)
	}	
	_method := "VDI.list_changed_blocks"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_vdiFromArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "vdi_from"), vdiFrom)
	if _err != nil {
		return
	}
	_vdiToArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "vdi_to"), vdiTo)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _vdiFromArg, _vdiToArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringToGo(_method + " -> ", _result.Value)
	return
}


func VDIClassDataDestroyMockDefault(sessionID SessionRef, self VDIRef) (_err error) {
	log.Println("VDI.DataDestroy not mocked")
	_err = errors.New("VDI.DataDestroy not mocked")
	return
}

var VDIClassDataDestroyMockedCallback = VDIClassDataDestroyMockDefault

func (_class VDIClass) DataDestroyMock(sessionID SessionRef, self VDIRef) (_err error) {
	return VDIClassDataDestroyMockedCallback(sessionID, self)
}
// Delete the data of the snapshot VDI, but keep its changed block tracking metadata. When successful, this call changes the type of the VDI to cbt_metadata. This operation is idempotent: calling it on a VDI of type cbt_metadata results in a no-op, and no error will be thrown.
//
// Errors:
//  SR_OPERATION_NOT_SUPPORTED - The SR backend does not support the operation (check the SR's allowed operations)
//  VDI_MISSING - This operation cannot be performed because the specified VDI could not be found on the storage substrate
//  SR_NOT_ATTACHED - The SR is not attached.
//  SR_HAS_NO_PBDS - The SR has no attached PBDs
//  OPERATION_NOT_ALLOWED - You attempted an operation that was not allowed.
//  VDI_INCOMPATIBLE_TYPE - This operation cannot be performed because the specified VDI is of an incompatible type (eg: an HA statefile cannot be attached to a guest)
//  VDI_NO_CBT_METADATA - The requested operation is not allowed because the specified VDI does not have changed block tracking metadata.
//  VDI_IN_USE - This operation cannot be performed because this VDI is in use by some other operation
//  VDI_IS_A_PHYSICAL_DEVICE - The operation cannot be performed on physical device
func (_class VDIClass) DataDestroy(sessionID SessionRef, self VDIRef) (_err error) {
	if IsMock {
		return _class.DataDestroyMock(sessionID, self)
	}	
	_method := "VDI.data_destroy"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg)
	return
}


func VDIClassDisableCbtMockDefault(sessionID SessionRef, self VDIRef) (_err error) {
	log.Println("VDI.DisableCbt not mocked")
	_err = errors.New("VDI.DisableCbt not mocked")
	return
}

var VDIClassDisableCbtMockedCallback = VDIClassDisableCbtMockDefault

func (_class VDIClass) DisableCbtMock(sessionID SessionRef, self VDIRef) (_err error) {
	return VDIClassDisableCbtMockedCallback(sessionID, self)
}
// Disable changed block tracking for the VDI. This call is only allowed on VDIs that support enabling CBT. It is an idempotent operation - disabling CBT for a VDI for which CBT is not enabled results in a no-op, and no error will be thrown.
//
// Errors:
//  SR_OPERATION_NOT_SUPPORTED - The SR backend does not support the operation (check the SR's allowed operations)
//  VDI_MISSING - This operation cannot be performed because the specified VDI could not be found on the storage substrate
//  SR_NOT_ATTACHED - The SR is not attached.
//  SR_HAS_NO_PBDS - The SR has no attached PBDs
//  OPERATION_NOT_ALLOWED - You attempted an operation that was not allowed.
//  VDI_INCOMPATIBLE_TYPE - This operation cannot be performed because the specified VDI is of an incompatible type (eg: an HA statefile cannot be attached to a guest)
//  VDI_ON_BOOT_MODE_INCOMPATIBLE_WITH_OPERATION - This operation is not permitted on VDIs in the 'on-boot=reset' mode, or on VMs having such VDIs.
func (_class VDIClass) DisableCbt(sessionID SessionRef, self VDIRef) (_err error) {
	if IsMock {
		return _class.DisableCbtMock(sessionID, self)
	}	
	_method := "VDI.disable_cbt"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg)
	return
}


func VDIClassEnableCbtMockDefault(sessionID SessionRef, self VDIRef) (_err error) {
	log.Println("VDI.EnableCbt not mocked")
	_err = errors.New("VDI.EnableCbt not mocked")
	return
}

var VDIClassEnableCbtMockedCallback = VDIClassEnableCbtMockDefault

func (_class VDIClass) EnableCbtMock(sessionID SessionRef, self VDIRef) (_err error) {
	return VDIClassEnableCbtMockedCallback(sessionID, self)
}
// Enable changed block tracking for the VDI. This call is idempotent - enabling CBT for a VDI for which CBT is already enabled results in a no-op, and no error will be thrown.
//
// Errors:
//  SR_OPERATION_NOT_SUPPORTED - The SR backend does not support the operation (check the SR's allowed operations)
//  VDI_MISSING - This operation cannot be performed because the specified VDI could not be found on the storage substrate
//  SR_NOT_ATTACHED - The SR is not attached.
//  SR_HAS_NO_PBDS - The SR has no attached PBDs
//  OPERATION_NOT_ALLOWED - You attempted an operation that was not allowed.
//  VDI_INCOMPATIBLE_TYPE - This operation cannot be performed because the specified VDI is of an incompatible type (eg: an HA statefile cannot be attached to a guest)
//  VDI_ON_BOOT_MODE_INCOMPATIBLE_WITH_OPERATION - This operation is not permitted on VDIs in the 'on-boot=reset' mode, or on VMs having such VDIs.
func (_class VDIClass) EnableCbt(sessionID SessionRef, self VDIRef) (_err error) {
	if IsMock {
		return _class.EnableCbtMock(sessionID, self)
	}	
	_method := "VDI.enable_cbt"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg)
	return
}


func VDIClassPoolMigrateMockDefault(sessionID SessionRef, vdi VDIRef, sr SRRef, options map[string]string) (_retval VDIRef, _err error) {
	log.Println("VDI.PoolMigrate not mocked")
	_err = errors.New("VDI.PoolMigrate not mocked")
	return
}

var VDIClassPoolMigrateMockedCallback = VDIClassPoolMigrateMockDefault

func (_class VDIClass) PoolMigrateMock(sessionID SessionRef, vdi VDIRef, sr SRRef, options map[string]string) (_retval VDIRef, _err error) {
	return VDIClassPoolMigrateMockedCallback(sessionID, vdi, sr, options)
}
// Migrate a VDI, which may be attached to a running guest, to a different SR. The destination SR must be visible to the guest.
func (_class VDIClass) PoolMigrate(sessionID SessionRef, vdi VDIRef, sr SRRef, options map[string]string) (_retval VDIRef, _err error) {
	if IsMock {
		return _class.PoolMigrateMock(sessionID, vdi, sr, options)
	}	
	_method := "VDI.pool_migrate"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_vdiArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "vdi"), vdi)
	if _err != nil {
		return
	}
	_srArg, _err := convertSRRefToXen(fmt.Sprintf("%s(%s)", _method, "sr"), sr)
	if _err != nil {
		return
	}
	_optionsArg, _err := convertStringToStringMapToXen(fmt.Sprintf("%s(%s)", _method, "options"), options)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _vdiArg, _srArg, _optionsArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVDIRefToGo(_method + " -> ", _result.Value)
	return
}


func VDIClassReadDatabasePoolUUIDMockDefault(sessionID SessionRef, self VDIRef) (_retval string, _err error) {
	log.Println("VDI.ReadDatabasePoolUUID not mocked")
	_err = errors.New("VDI.ReadDatabasePoolUUID not mocked")
	return
}

var VDIClassReadDatabasePoolUUIDMockedCallback = VDIClassReadDatabasePoolUUIDMockDefault

func (_class VDIClass) ReadDatabasePoolUUIDMock(sessionID SessionRef, self VDIRef) (_retval string, _err error) {
	return VDIClassReadDatabasePoolUUIDMockedCallback(sessionID, self)
}
// Check the VDI cache for the pool UUID of the database on this VDI.
func (_class VDIClass) ReadDatabasePoolUUID(sessionID SessionRef, self VDIRef) (_retval string, _err error) {
	if IsMock {
		return _class.ReadDatabasePoolUUIDMock(sessionID, self)
	}	
	_method := "VDI.read_database_pool_uuid"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func VDIClassOpenDatabaseMockDefault(sessionID SessionRef, self VDIRef) (_retval SessionRef, _err error) {
	log.Println("VDI.OpenDatabase not mocked")
	_err = errors.New("VDI.OpenDatabase not mocked")
	return
}

var VDIClassOpenDatabaseMockedCallback = VDIClassOpenDatabaseMockDefault

func (_class VDIClass) OpenDatabaseMock(sessionID SessionRef, self VDIRef) (_retval SessionRef, _err error) {
	return VDIClassOpenDatabaseMockedCallback(sessionID, self)
}
// Load the metadata found on the supplied VDI and return a session reference which can be used in XenAPI calls to query its contents.
func (_class VDIClass) OpenDatabase(sessionID SessionRef, self VDIRef) (_retval SessionRef, _err error) {
	if IsMock {
		return _class.OpenDatabaseMock(sessionID, self)
	}	
	_method := "VDI.open_database"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertSessionRefToGo(_method + " -> ", _result.Value)
	return
}


func VDIClassSetAllowCachingMockDefault(sessionID SessionRef, self VDIRef, value bool) (_err error) {
	log.Println("VDI.SetAllowCaching not mocked")
	_err = errors.New("VDI.SetAllowCaching not mocked")
	return
}

var VDIClassSetAllowCachingMockedCallback = VDIClassSetAllowCachingMockDefault

func (_class VDIClass) SetAllowCachingMock(sessionID SessionRef, self VDIRef, value bool) (_err error) {
	return VDIClassSetAllowCachingMockedCallback(sessionID, self, value)
}
// Set the value of the allow_caching parameter. This value can only be changed when the VDI is not attached to a running VM. The caching behaviour is only affected by this flag for VHD-based VDIs that have one parent and no child VHDs. Moreover, caching only takes place when the host running the VM containing this VDI has a nominated SR for local caching.
func (_class VDIClass) SetAllowCaching(sessionID SessionRef, self VDIRef, value bool) (_err error) {
	if IsMock {
		return _class.SetAllowCachingMock(sessionID, self, value)
	}	
	_method := "VDI.set_allow_caching"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func VDIClassSetOnBootMockDefault(sessionID SessionRef, self VDIRef, value OnBoot) (_err error) {
	log.Println("VDI.SetOnBoot not mocked")
	_err = errors.New("VDI.SetOnBoot not mocked")
	return
}

var VDIClassSetOnBootMockedCallback = VDIClassSetOnBootMockDefault

func (_class VDIClass) SetOnBootMock(sessionID SessionRef, self VDIRef, value OnBoot) (_err error) {
	return VDIClassSetOnBootMockedCallback(sessionID, self, value)
}
// Set the value of the on_boot parameter. This value can only be changed when the VDI is not attached to a running VM.
func (_class VDIClass) SetOnBoot(sessionID SessionRef, self VDIRef, value OnBoot) (_err error) {
	if IsMock {
		return _class.SetOnBootMock(sessionID, self, value)
	}	
	_method := "VDI.set_on_boot"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertEnumOnBootToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}


func VDIClassSetNameDescriptionMockDefault(sessionID SessionRef, self VDIRef, value string) (_err error) {
	log.Println("VDI.SetNameDescription not mocked")
	_err = errors.New("VDI.SetNameDescription not mocked")
	return
}

var VDIClassSetNameDescriptionMockedCallback = VDIClassSetNameDescriptionMockDefault

func (_class VDIClass) SetNameDescriptionMock(sessionID SessionRef, self VDIRef, value string) (_err error) {
	return VDIClassSetNameDescriptionMockedCallback(sessionID, self, value)
}
// Set the name description of the VDI. This can only happen when its SR is currently attached.
func (_class VDIClass) SetNameDescription(sessionID SessionRef, self VDIRef, value string) (_err error) {
	if IsMock {
		return _class.SetNameDescriptionMock(sessionID, self, value)
	}	
	_method := "VDI.set_name_description"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func VDIClassSetNameLabelMockDefault(sessionID SessionRef, self VDIRef, value string) (_err error) {
	log.Println("VDI.SetNameLabel not mocked")
	_err = errors.New("VDI.SetNameLabel not mocked")
	return
}

var VDIClassSetNameLabelMockedCallback = VDIClassSetNameLabelMockDefault

func (_class VDIClass) SetNameLabelMock(sessionID SessionRef, self VDIRef, value string) (_err error) {
	return VDIClassSetNameLabelMockedCallback(sessionID, self, value)
}
// Set the name label of the VDI. This can only happen when then its SR is currently attached.
func (_class VDIClass) SetNameLabel(sessionID SessionRef, self VDIRef, value string) (_err error) {
	if IsMock {
		return _class.SetNameLabelMock(sessionID, self, value)
	}	
	_method := "VDI.set_name_label"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func VDIClassSetMetadataOfPoolMockDefault(sessionID SessionRef, self VDIRef, value PoolRef) (_err error) {
	log.Println("VDI.SetMetadataOfPool not mocked")
	_err = errors.New("VDI.SetMetadataOfPool not mocked")
	return
}

var VDIClassSetMetadataOfPoolMockedCallback = VDIClassSetMetadataOfPoolMockDefault

func (_class VDIClass) SetMetadataOfPoolMock(sessionID SessionRef, self VDIRef, value PoolRef) (_err error) {
	return VDIClassSetMetadataOfPoolMockedCallback(sessionID, self, value)
}
// Records the pool whose metadata is contained by this VDI.
func (_class VDIClass) SetMetadataOfPool(sessionID SessionRef, self VDIRef, value PoolRef) (_err error) {
	if IsMock {
		return _class.SetMetadataOfPoolMock(sessionID, self, value)
	}	
	_method := "VDI.set_metadata_of_pool"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}


func VDIClassSetSnapshotTimeMockDefault(sessionID SessionRef, self VDIRef, value time.Time) (_err error) {
	log.Println("VDI.SetSnapshotTime not mocked")
	_err = errors.New("VDI.SetSnapshotTime not mocked")
	return
}

var VDIClassSetSnapshotTimeMockedCallback = VDIClassSetSnapshotTimeMockDefault

func (_class VDIClass) SetSnapshotTimeMock(sessionID SessionRef, self VDIRef, value time.Time) (_err error) {
	return VDIClassSetSnapshotTimeMockedCallback(sessionID, self, value)
}
// Sets the snapshot time of this VDI.
func (_class VDIClass) SetSnapshotTime(sessionID SessionRef, self VDIRef, value time.Time) (_err error) {
	if IsMock {
		return _class.SetSnapshotTimeMock(sessionID, self, value)
	}	
	_method := "VDI.set_snapshot_time"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func VDIClassSetSnapshotOfMockDefault(sessionID SessionRef, self VDIRef, value VDIRef) (_err error) {
	log.Println("VDI.SetSnapshotOf not mocked")
	_err = errors.New("VDI.SetSnapshotOf not mocked")
	return
}

var VDIClassSetSnapshotOfMockedCallback = VDIClassSetSnapshotOfMockDefault

func (_class VDIClass) SetSnapshotOfMock(sessionID SessionRef, self VDIRef, value VDIRef) (_err error) {
	return VDIClassSetSnapshotOfMockedCallback(sessionID, self, value)
}
// Sets the VDI of which this VDI is a snapshot
func (_class VDIClass) SetSnapshotOf(sessionID SessionRef, self VDIRef, value VDIRef) (_err error) {
	if IsMock {
		return _class.SetSnapshotOfMock(sessionID, self, value)
	}	
	_method := "VDI.set_snapshot_of"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}


func VDIClassSetIsASnapshotMockDefault(sessionID SessionRef, self VDIRef, value bool) (_err error) {
	log.Println("VDI.SetIsASnapshot not mocked")
	_err = errors.New("VDI.SetIsASnapshot not mocked")
	return
}

var VDIClassSetIsASnapshotMockedCallback = VDIClassSetIsASnapshotMockDefault

func (_class VDIClass) SetIsASnapshotMock(sessionID SessionRef, self VDIRef, value bool) (_err error) {
	return VDIClassSetIsASnapshotMockedCallback(sessionID, self, value)
}
// Sets whether this VDI is a snapshot
func (_class VDIClass) SetIsASnapshot(sessionID SessionRef, self VDIRef, value bool) (_err error) {
	if IsMock {
		return _class.SetIsASnapshotMock(sessionID, self, value)
	}	
	_method := "VDI.set_is_a_snapshot"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func VDIClassSetPhysicalUtilisationMockDefault(sessionID SessionRef, self VDIRef, value int) (_err error) {
	log.Println("VDI.SetPhysicalUtilisation not mocked")
	_err = errors.New("VDI.SetPhysicalUtilisation not mocked")
	return
}

var VDIClassSetPhysicalUtilisationMockedCallback = VDIClassSetPhysicalUtilisationMockDefault

func (_class VDIClass) SetPhysicalUtilisationMock(sessionID SessionRef, self VDIRef, value int) (_err error) {
	return VDIClassSetPhysicalUtilisationMockedCallback(sessionID, self, value)
}
// Sets the VDI's physical_utilisation field
func (_class VDIClass) SetPhysicalUtilisation(sessionID SessionRef, self VDIRef, value int) (_err error) {
	if IsMock {
		return _class.SetPhysicalUtilisationMock(sessionID, self, value)
	}	
	_method := "VDI.set_physical_utilisation"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func VDIClassSetVirtualSizeMockDefault(sessionID SessionRef, self VDIRef, value int) (_err error) {
	log.Println("VDI.SetVirtualSize not mocked")
	_err = errors.New("VDI.SetVirtualSize not mocked")
	return
}

var VDIClassSetVirtualSizeMockedCallback = VDIClassSetVirtualSizeMockDefault

func (_class VDIClass) SetVirtualSizeMock(sessionID SessionRef, self VDIRef, value int) (_err error) {
	return VDIClassSetVirtualSizeMockedCallback(sessionID, self, value)
}
// Sets the VDI's virtual_size field
func (_class VDIClass) SetVirtualSize(sessionID SessionRef, self VDIRef, value int) (_err error) {
	if IsMock {
		return _class.SetVirtualSizeMock(sessionID, self, value)
	}	
	_method := "VDI.set_virtual_size"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func VDIClassSetMissingMockDefault(sessionID SessionRef, self VDIRef, value bool) (_err error) {
	log.Println("VDI.SetMissing not mocked")
	_err = errors.New("VDI.SetMissing not mocked")
	return
}

var VDIClassSetMissingMockedCallback = VDIClassSetMissingMockDefault

func (_class VDIClass) SetMissingMock(sessionID SessionRef, self VDIRef, value bool) (_err error) {
	return VDIClassSetMissingMockedCallback(sessionID, self, value)
}
// Sets the VDI's missing field
func (_class VDIClass) SetMissing(sessionID SessionRef, self VDIRef, value bool) (_err error) {
	if IsMock {
		return _class.SetMissingMock(sessionID, self, value)
	}	
	_method := "VDI.set_missing"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func VDIClassSetReadOnlyMockDefault(sessionID SessionRef, self VDIRef, value bool) (_err error) {
	log.Println("VDI.SetReadOnly not mocked")
	_err = errors.New("VDI.SetReadOnly not mocked")
	return
}

var VDIClassSetReadOnlyMockedCallback = VDIClassSetReadOnlyMockDefault

func (_class VDIClass) SetReadOnlyMock(sessionID SessionRef, self VDIRef, value bool) (_err error) {
	return VDIClassSetReadOnlyMockedCallback(sessionID, self, value)
}
// Sets the VDI's read_only field
func (_class VDIClass) SetReadOnly(sessionID SessionRef, self VDIRef, value bool) (_err error) {
	if IsMock {
		return _class.SetReadOnlyMock(sessionID, self, value)
	}	
	_method := "VDI.set_read_only"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func VDIClassSetSharableMockDefault(sessionID SessionRef, self VDIRef, value bool) (_err error) {
	log.Println("VDI.SetSharable not mocked")
	_err = errors.New("VDI.SetSharable not mocked")
	return
}

var VDIClassSetSharableMockedCallback = VDIClassSetSharableMockDefault

func (_class VDIClass) SetSharableMock(sessionID SessionRef, self VDIRef, value bool) (_err error) {
	return VDIClassSetSharableMockedCallback(sessionID, self, value)
}
// Sets the VDI's sharable field
func (_class VDIClass) SetSharable(sessionID SessionRef, self VDIRef, value bool) (_err error) {
	if IsMock {
		return _class.SetSharableMock(sessionID, self, value)
	}	
	_method := "VDI.set_sharable"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func VDIClassForgetMockDefault(sessionID SessionRef, vdi VDIRef) (_err error) {
	log.Println("VDI.Forget not mocked")
	_err = errors.New("VDI.Forget not mocked")
	return
}

var VDIClassForgetMockedCallback = VDIClassForgetMockDefault

func (_class VDIClass) ForgetMock(sessionID SessionRef, vdi VDIRef) (_err error) {
	return VDIClassForgetMockedCallback(sessionID, vdi)
}
// Removes a VDI record from the database
func (_class VDIClass) Forget(sessionID SessionRef, vdi VDIRef) (_err error) {
	if IsMock {
		return _class.ForgetMock(sessionID, vdi)
	}	
	_method := "VDI.forget"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_vdiArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "vdi"), vdi)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _vdiArg)
	return
}


func VDIClassSetManagedMockDefault(sessionID SessionRef, self VDIRef, value bool) (_err error) {
	log.Println("VDI.SetManaged not mocked")
	_err = errors.New("VDI.SetManaged not mocked")
	return
}

var VDIClassSetManagedMockedCallback = VDIClassSetManagedMockDefault

func (_class VDIClass) SetManagedMock(sessionID SessionRef, self VDIRef, value bool) (_err error) {
	return VDIClassSetManagedMockedCallback(sessionID, self, value)
}
// Sets the VDI's managed field
func (_class VDIClass) SetManaged(sessionID SessionRef, self VDIRef, value bool) (_err error) {
	if IsMock {
		return _class.SetManagedMock(sessionID, self, value)
	}	
	_method := "VDI.set_managed"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func VDIClassCopyMockDefault(sessionID SessionRef, vdi VDIRef, sr SRRef, baseVdi VDIRef, intoVdi VDIRef) (_retval VDIRef, _err error) {
	log.Println("VDI.Copy not mocked")
	_err = errors.New("VDI.Copy not mocked")
	return
}

var VDIClassCopyMockedCallback = VDIClassCopyMockDefault

func (_class VDIClass) CopyMock(sessionID SessionRef, vdi VDIRef, sr SRRef, baseVdi VDIRef, intoVdi VDIRef) (_retval VDIRef, _err error) {
	return VDIClassCopyMockedCallback(sessionID, vdi, sr, baseVdi, intoVdi)
}
// Copy either a full VDI or the block differences between two VDIs into either a fresh VDI or an existing VDI.
//
// Errors:
//  VDI_READONLY - The operation required write access but this VDI is read-only
//  VDI_TOO_SMALL - The VDI is too small. Please resize it to at least the minimum size.
//  VDI_NOT_SPARSE - The VDI is not stored using a sparse format. It is not possible to query and manipulate only the changed blocks (or 'block differences' or 'disk deltas') between two VDIs. Please select a VDI which uses a sparse-aware technology such as VHD.
func (_class VDIClass) Copy(sessionID SessionRef, vdi VDIRef, sr SRRef, baseVdi VDIRef, intoVdi VDIRef) (_retval VDIRef, _err error) {
	if IsMock {
		return _class.CopyMock(sessionID, vdi, sr, baseVdi, intoVdi)
	}	
	_method := "VDI.copy"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_vdiArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "vdi"), vdi)
	if _err != nil {
		return
	}
	_srArg, _err := convertSRRefToXen(fmt.Sprintf("%s(%s)", _method, "sr"), sr)
	if _err != nil {
		return
	}
	_baseVdiArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "base_vdi"), baseVdi)
	if _err != nil {
		return
	}
	_intoVdiArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "into_vdi"), intoVdi)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _vdiArg, _srArg, _baseVdiArg, _intoVdiArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVDIRefToGo(_method + " -> ", _result.Value)
	return
}


func VDIClassUpdateMockDefault(sessionID SessionRef, vdi VDIRef) (_err error) {
	log.Println("VDI.Update not mocked")
	_err = errors.New("VDI.Update not mocked")
	return
}

var VDIClassUpdateMockedCallback = VDIClassUpdateMockDefault

func (_class VDIClass) UpdateMock(sessionID SessionRef, vdi VDIRef) (_err error) {
	return VDIClassUpdateMockedCallback(sessionID, vdi)
}
// Ask the storage backend to refresh the fields in the VDI object
//
// Errors:
//  SR_OPERATION_NOT_SUPPORTED - The SR backend does not support the operation (check the SR's allowed operations)
func (_class VDIClass) Update(sessionID SessionRef, vdi VDIRef) (_err error) {
	if IsMock {
		return _class.UpdateMock(sessionID, vdi)
	}	
	_method := "VDI.update"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_vdiArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "vdi"), vdi)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _vdiArg)
	return
}


func VDIClassDbForgetMockDefault(sessionID SessionRef, vdi VDIRef) (_err error) {
	log.Println("VDI.DbForget not mocked")
	_err = errors.New("VDI.DbForget not mocked")
	return
}

var VDIClassDbForgetMockedCallback = VDIClassDbForgetMockDefault

func (_class VDIClass) DbForgetMock(sessionID SessionRef, vdi VDIRef) (_err error) {
	return VDIClassDbForgetMockedCallback(sessionID, vdi)
}
// Removes a VDI record from the database
func (_class VDIClass) DbForget(sessionID SessionRef, vdi VDIRef) (_err error) {
	if IsMock {
		return _class.DbForgetMock(sessionID, vdi)
	}	
	_method := "VDI.db_forget"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_vdiArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "vdi"), vdi)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _vdiArg)
	return
}


func VDIClassDbIntroduceMockDefault(sessionID SessionRef, uuid string, nameLabel string, nameDescription string, sr SRRef, atype VdiType, sharable bool, readOnly bool, otherConfig map[string]string, location string, xenstoreData map[string]string, smConfig map[string]string, managed bool, virtualSize int, physicalUtilisation int, metadataOfPool PoolRef, isASnapshot bool, snapshotTime time.Time, snapshotOf VDIRef, cbtEnabled bool) (_retval VDIRef, _err error) {
	log.Println("VDI.DbIntroduce not mocked")
	_err = errors.New("VDI.DbIntroduce not mocked")
	return
}

var VDIClassDbIntroduceMockedCallback = VDIClassDbIntroduceMockDefault

func (_class VDIClass) DbIntroduceMock(sessionID SessionRef, uuid string, nameLabel string, nameDescription string, sr SRRef, atype VdiType, sharable bool, readOnly bool, otherConfig map[string]string, location string, xenstoreData map[string]string, smConfig map[string]string, managed bool, virtualSize int, physicalUtilisation int, metadataOfPool PoolRef, isASnapshot bool, snapshotTime time.Time, snapshotOf VDIRef, cbtEnabled bool) (_retval VDIRef, _err error) {
	return VDIClassDbIntroduceMockedCallback(sessionID, uuid, nameLabel, nameDescription, sr, atype, sharable, readOnly, otherConfig, location, xenstoreData, smConfig, managed, virtualSize, physicalUtilisation, metadataOfPool, isASnapshot, snapshotTime, snapshotOf, cbtEnabled)
}
// Create a new VDI record in the database only
func (_class VDIClass) DbIntroduce(sessionID SessionRef, uuid string, nameLabel string, nameDescription string, sr SRRef, atype VdiType, sharable bool, readOnly bool, otherConfig map[string]string, location string, xenstoreData map[string]string, smConfig map[string]string, managed bool, virtualSize int, physicalUtilisation int, metadataOfPool PoolRef, isASnapshot bool, snapshotTime time.Time, snapshotOf VDIRef, cbtEnabled bool) (_retval VDIRef, _err error) {
	if IsMock {
		return _class.DbIntroduceMock(sessionID, uuid, nameLabel, nameDescription, sr, atype, sharable, readOnly, otherConfig, location, xenstoreData, smConfig, managed, virtualSize, physicalUtilisation, metadataOfPool, isASnapshot, snapshotTime, snapshotOf, cbtEnabled)
	}	
	_method := "VDI.db_introduce"
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
	_srArg, _err := convertSRRefToXen(fmt.Sprintf("%s(%s)", _method, "SR"), sr)
	if _err != nil {
		return
	}
	_atypeArg, _err := convertEnumVdiTypeToXen(fmt.Sprintf("%s(%s)", _method, "type"), atype)
	if _err != nil {
		return
	}
	_sharableArg, _err := convertBoolToXen(fmt.Sprintf("%s(%s)", _method, "sharable"), sharable)
	if _err != nil {
		return
	}
	_readOnlyArg, _err := convertBoolToXen(fmt.Sprintf("%s(%s)", _method, "read_only"), readOnly)
	if _err != nil {
		return
	}
	_otherConfigArg, _err := convertStringToStringMapToXen(fmt.Sprintf("%s(%s)", _method, "other_config"), otherConfig)
	if _err != nil {
		return
	}
	_locationArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "location"), location)
	if _err != nil {
		return
	}
	_xenstoreDataArg, _err := convertStringToStringMapToXen(fmt.Sprintf("%s(%s)", _method, "xenstore_data"), xenstoreData)
	if _err != nil {
		return
	}
	_smConfigArg, _err := convertStringToStringMapToXen(fmt.Sprintf("%s(%s)", _method, "sm_config"), smConfig)
	if _err != nil {
		return
	}
	_managedArg, _err := convertBoolToXen(fmt.Sprintf("%s(%s)", _method, "managed"), managed)
	if _err != nil {
		return
	}
	_virtualSizeArg, _err := convertIntToXen(fmt.Sprintf("%s(%s)", _method, "virtual_size"), virtualSize)
	if _err != nil {
		return
	}
	_physicalUtilisationArg, _err := convertIntToXen(fmt.Sprintf("%s(%s)", _method, "physical_utilisation"), physicalUtilisation)
	if _err != nil {
		return
	}
	_metadataOfPoolArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "metadata_of_pool"), metadataOfPool)
	if _err != nil {
		return
	}
	_isASnapshotArg, _err := convertBoolToXen(fmt.Sprintf("%s(%s)", _method, "is_a_snapshot"), isASnapshot)
	if _err != nil {
		return
	}
	_snapshotTimeArg, _err := convertTimeToXen(fmt.Sprintf("%s(%s)", _method, "snapshot_time"), snapshotTime)
	if _err != nil {
		return
	}
	_snapshotOfArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "snapshot_of"), snapshotOf)
	if _err != nil {
		return
	}
	_cbtEnabledArg, _err := convertBoolToXen(fmt.Sprintf("%s(%s)", _method, "cbt_enabled"), cbtEnabled)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _uuidArg, _nameLabelArg, _nameDescriptionArg, _srArg, _atypeArg, _sharableArg, _readOnlyArg, _otherConfigArg, _locationArg, _xenstoreDataArg, _smConfigArg, _managedArg, _virtualSizeArg, _physicalUtilisationArg, _metadataOfPoolArg, _isASnapshotArg, _snapshotTimeArg, _snapshotOfArg, _cbtEnabledArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVDIRefToGo(_method + " -> ", _result.Value)
	return
}


func VDIClassIntroduceMockDefault(sessionID SessionRef, uuid string, nameLabel string, nameDescription string, sr SRRef, atype VdiType, sharable bool, readOnly bool, otherConfig map[string]string, location string, xenstoreData map[string]string, smConfig map[string]string, managed bool, virtualSize int, physicalUtilisation int, metadataOfPool PoolRef, isASnapshot bool, snapshotTime time.Time, snapshotOf VDIRef) (_retval VDIRef, _err error) {
	log.Println("VDI.Introduce not mocked")
	_err = errors.New("VDI.Introduce not mocked")
	return
}

var VDIClassIntroduceMockedCallback = VDIClassIntroduceMockDefault

func (_class VDIClass) IntroduceMock(sessionID SessionRef, uuid string, nameLabel string, nameDescription string, sr SRRef, atype VdiType, sharable bool, readOnly bool, otherConfig map[string]string, location string, xenstoreData map[string]string, smConfig map[string]string, managed bool, virtualSize int, physicalUtilisation int, metadataOfPool PoolRef, isASnapshot bool, snapshotTime time.Time, snapshotOf VDIRef) (_retval VDIRef, _err error) {
	return VDIClassIntroduceMockedCallback(sessionID, uuid, nameLabel, nameDescription, sr, atype, sharable, readOnly, otherConfig, location, xenstoreData, smConfig, managed, virtualSize, physicalUtilisation, metadataOfPool, isASnapshot, snapshotTime, snapshotOf)
}
// Create a new VDI record in the database only
//
// Errors:
//  SR_OPERATION_NOT_SUPPORTED - The SR backend does not support the operation (check the SR's allowed operations)
func (_class VDIClass) Introduce(sessionID SessionRef, uuid string, nameLabel string, nameDescription string, sr SRRef, atype VdiType, sharable bool, readOnly bool, otherConfig map[string]string, location string, xenstoreData map[string]string, smConfig map[string]string, managed bool, virtualSize int, physicalUtilisation int, metadataOfPool PoolRef, isASnapshot bool, snapshotTime time.Time, snapshotOf VDIRef) (_retval VDIRef, _err error) {
	if IsMock {
		return _class.IntroduceMock(sessionID, uuid, nameLabel, nameDescription, sr, atype, sharable, readOnly, otherConfig, location, xenstoreData, smConfig, managed, virtualSize, physicalUtilisation, metadataOfPool, isASnapshot, snapshotTime, snapshotOf)
	}	
	_method := "VDI.introduce"
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
	_srArg, _err := convertSRRefToXen(fmt.Sprintf("%s(%s)", _method, "SR"), sr)
	if _err != nil {
		return
	}
	_atypeArg, _err := convertEnumVdiTypeToXen(fmt.Sprintf("%s(%s)", _method, "type"), atype)
	if _err != nil {
		return
	}
	_sharableArg, _err := convertBoolToXen(fmt.Sprintf("%s(%s)", _method, "sharable"), sharable)
	if _err != nil {
		return
	}
	_readOnlyArg, _err := convertBoolToXen(fmt.Sprintf("%s(%s)", _method, "read_only"), readOnly)
	if _err != nil {
		return
	}
	_otherConfigArg, _err := convertStringToStringMapToXen(fmt.Sprintf("%s(%s)", _method, "other_config"), otherConfig)
	if _err != nil {
		return
	}
	_locationArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "location"), location)
	if _err != nil {
		return
	}
	_xenstoreDataArg, _err := convertStringToStringMapToXen(fmt.Sprintf("%s(%s)", _method, "xenstore_data"), xenstoreData)
	if _err != nil {
		return
	}
	_smConfigArg, _err := convertStringToStringMapToXen(fmt.Sprintf("%s(%s)", _method, "sm_config"), smConfig)
	if _err != nil {
		return
	}
	_managedArg, _err := convertBoolToXen(fmt.Sprintf("%s(%s)", _method, "managed"), managed)
	if _err != nil {
		return
	}
	_virtualSizeArg, _err := convertIntToXen(fmt.Sprintf("%s(%s)", _method, "virtual_size"), virtualSize)
	if _err != nil {
		return
	}
	_physicalUtilisationArg, _err := convertIntToXen(fmt.Sprintf("%s(%s)", _method, "physical_utilisation"), physicalUtilisation)
	if _err != nil {
		return
	}
	_metadataOfPoolArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "metadata_of_pool"), metadataOfPool)
	if _err != nil {
		return
	}
	_isASnapshotArg, _err := convertBoolToXen(fmt.Sprintf("%s(%s)", _method, "is_a_snapshot"), isASnapshot)
	if _err != nil {
		return
	}
	_snapshotTimeArg, _err := convertTimeToXen(fmt.Sprintf("%s(%s)", _method, "snapshot_time"), snapshotTime)
	if _err != nil {
		return
	}
	_snapshotOfArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "snapshot_of"), snapshotOf)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _uuidArg, _nameLabelArg, _nameDescriptionArg, _srArg, _atypeArg, _sharableArg, _readOnlyArg, _otherConfigArg, _locationArg, _xenstoreDataArg, _smConfigArg, _managedArg, _virtualSizeArg, _physicalUtilisationArg, _metadataOfPoolArg, _isASnapshotArg, _snapshotTimeArg, _snapshotOfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVDIRefToGo(_method + " -> ", _result.Value)
	return
}


func VDIClassResizeOnlineMockDefault(sessionID SessionRef, vdi VDIRef, size int) (_err error) {
	log.Println("VDI.ResizeOnline not mocked")
	_err = errors.New("VDI.ResizeOnline not mocked")
	return
}

var VDIClassResizeOnlineMockedCallback = VDIClassResizeOnlineMockDefault

func (_class VDIClass) ResizeOnlineMock(sessionID SessionRef, vdi VDIRef, size int) (_err error) {
	return VDIClassResizeOnlineMockedCallback(sessionID, vdi, size)
}
// Resize the VDI which may or may not be attached to running guests.
func (_class VDIClass) ResizeOnline(sessionID SessionRef, vdi VDIRef, size int) (_err error) {
	if IsMock {
		return _class.ResizeOnlineMock(sessionID, vdi, size)
	}	
	_method := "VDI.resize_online"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_vdiArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "vdi"), vdi)
	if _err != nil {
		return
	}
	_sizeArg, _err := convertIntToXen(fmt.Sprintf("%s(%s)", _method, "size"), size)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _vdiArg, _sizeArg)
	return
}


func VDIClassResizeMockDefault(sessionID SessionRef, vdi VDIRef, size int) (_err error) {
	log.Println("VDI.Resize not mocked")
	_err = errors.New("VDI.Resize not mocked")
	return
}

var VDIClassResizeMockedCallback = VDIClassResizeMockDefault

func (_class VDIClass) ResizeMock(sessionID SessionRef, vdi VDIRef, size int) (_err error) {
	return VDIClassResizeMockedCallback(sessionID, vdi, size)
}
// Resize the VDI.
func (_class VDIClass) Resize(sessionID SessionRef, vdi VDIRef, size int) (_err error) {
	if IsMock {
		return _class.ResizeMock(sessionID, vdi, size)
	}	
	_method := "VDI.resize"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_vdiArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "vdi"), vdi)
	if _err != nil {
		return
	}
	_sizeArg, _err := convertIntToXen(fmt.Sprintf("%s(%s)", _method, "size"), size)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _vdiArg, _sizeArg)
	return
}


func VDIClassCloneMockDefault(sessionID SessionRef, vdi VDIRef, driverParams map[string]string) (_retval VDIRef, _err error) {
	log.Println("VDI.Clone not mocked")
	_err = errors.New("VDI.Clone not mocked")
	return
}

var VDIClassCloneMockedCallback = VDIClassCloneMockDefault

func (_class VDIClass) CloneMock(sessionID SessionRef, vdi VDIRef, driverParams map[string]string) (_retval VDIRef, _err error) {
	return VDIClassCloneMockedCallback(sessionID, vdi, driverParams)
}
// Take an exact copy of the VDI and return a reference to the new disk. If any driver_params are specified then these are passed through to the storage-specific substrate driver that implements the clone operation. NB the clone lives in the same Storage Repository as its parent.
func (_class VDIClass) Clone(sessionID SessionRef, vdi VDIRef, driverParams map[string]string) (_retval VDIRef, _err error) {
	if IsMock {
		return _class.CloneMock(sessionID, vdi, driverParams)
	}	
	_method := "VDI.clone"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_vdiArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "vdi"), vdi)
	if _err != nil {
		return
	}
	_driverParamsArg, _err := convertStringToStringMapToXen(fmt.Sprintf("%s(%s)", _method, "driver_params"), driverParams)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _vdiArg, _driverParamsArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVDIRefToGo(_method + " -> ", _result.Value)
	return
}


func VDIClassSnapshotMockDefault(sessionID SessionRef, vdi VDIRef, driverParams map[string]string) (_retval VDIRef, _err error) {
	log.Println("VDI.Snapshot not mocked")
	_err = errors.New("VDI.Snapshot not mocked")
	return
}

var VDIClassSnapshotMockedCallback = VDIClassSnapshotMockDefault

func (_class VDIClass) SnapshotMock(sessionID SessionRef, vdi VDIRef, driverParams map[string]string) (_retval VDIRef, _err error) {
	return VDIClassSnapshotMockedCallback(sessionID, vdi, driverParams)
}
// Take a read-only snapshot of the VDI, returning a reference to the snapshot. If any driver_params are specified then these are passed through to the storage-specific substrate driver that takes the snapshot. NB the snapshot lives in the same Storage Repository as its parent.
func (_class VDIClass) Snapshot(sessionID SessionRef, vdi VDIRef, driverParams map[string]string) (_retval VDIRef, _err error) {
	if IsMock {
		return _class.SnapshotMock(sessionID, vdi, driverParams)
	}	
	_method := "VDI.snapshot"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_vdiArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "vdi"), vdi)
	if _err != nil {
		return
	}
	_driverParamsArg, _err := convertStringToStringMapToXen(fmt.Sprintf("%s(%s)", _method, "driver_params"), driverParams)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _vdiArg, _driverParamsArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVDIRefToGo(_method + " -> ", _result.Value)
	return
}


func VDIClassRemoveTagsMockDefault(sessionID SessionRef, self VDIRef, value string) (_err error) {
	log.Println("VDI.RemoveTags not mocked")
	_err = errors.New("VDI.RemoveTags not mocked")
	return
}

var VDIClassRemoveTagsMockedCallback = VDIClassRemoveTagsMockDefault

func (_class VDIClass) RemoveTagsMock(sessionID SessionRef, self VDIRef, value string) (_err error) {
	return VDIClassRemoveTagsMockedCallback(sessionID, self, value)
}
// Remove the given value from the tags field of the given VDI.  If the value is not in that Set, then do nothing.
func (_class VDIClass) RemoveTags(sessionID SessionRef, self VDIRef, value string) (_err error) {
	if IsMock {
		return _class.RemoveTagsMock(sessionID, self, value)
	}	
	_method := "VDI.remove_tags"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func VDIClassAddTagsMockDefault(sessionID SessionRef, self VDIRef, value string) (_err error) {
	log.Println("VDI.AddTags not mocked")
	_err = errors.New("VDI.AddTags not mocked")
	return
}

var VDIClassAddTagsMockedCallback = VDIClassAddTagsMockDefault

func (_class VDIClass) AddTagsMock(sessionID SessionRef, self VDIRef, value string) (_err error) {
	return VDIClassAddTagsMockedCallback(sessionID, self, value)
}
// Add the given value to the tags field of the given VDI.  If the value is already in that Set, then do nothing.
func (_class VDIClass) AddTags(sessionID SessionRef, self VDIRef, value string) (_err error) {
	if IsMock {
		return _class.AddTagsMock(sessionID, self, value)
	}	
	_method := "VDI.add_tags"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func VDIClassSetTagsMockDefault(sessionID SessionRef, self VDIRef, value []string) (_err error) {
	log.Println("VDI.SetTags not mocked")
	_err = errors.New("VDI.SetTags not mocked")
	return
}

var VDIClassSetTagsMockedCallback = VDIClassSetTagsMockDefault

func (_class VDIClass) SetTagsMock(sessionID SessionRef, self VDIRef, value []string) (_err error) {
	return VDIClassSetTagsMockedCallback(sessionID, self, value)
}
// Set the tags field of the given VDI.
func (_class VDIClass) SetTags(sessionID SessionRef, self VDIRef, value []string) (_err error) {
	if IsMock {
		return _class.SetTagsMock(sessionID, self, value)
	}	
	_method := "VDI.set_tags"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func VDIClassRemoveFromSmConfigMockDefault(sessionID SessionRef, self VDIRef, key string) (_err error) {
	log.Println("VDI.RemoveFromSmConfig not mocked")
	_err = errors.New("VDI.RemoveFromSmConfig not mocked")
	return
}

var VDIClassRemoveFromSmConfigMockedCallback = VDIClassRemoveFromSmConfigMockDefault

func (_class VDIClass) RemoveFromSmConfigMock(sessionID SessionRef, self VDIRef, key string) (_err error) {
	return VDIClassRemoveFromSmConfigMockedCallback(sessionID, self, key)
}
// Remove the given key and its corresponding value from the sm_config field of the given VDI.  If the key is not in that Map, then do nothing.
func (_class VDIClass) RemoveFromSmConfig(sessionID SessionRef, self VDIRef, key string) (_err error) {
	if IsMock {
		return _class.RemoveFromSmConfigMock(sessionID, self, key)
	}	
	_method := "VDI.remove_from_sm_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func VDIClassAddToSmConfigMockDefault(sessionID SessionRef, self VDIRef, key string, value string) (_err error) {
	log.Println("VDI.AddToSmConfig not mocked")
	_err = errors.New("VDI.AddToSmConfig not mocked")
	return
}

var VDIClassAddToSmConfigMockedCallback = VDIClassAddToSmConfigMockDefault

func (_class VDIClass) AddToSmConfigMock(sessionID SessionRef, self VDIRef, key string, value string) (_err error) {
	return VDIClassAddToSmConfigMockedCallback(sessionID, self, key, value)
}
// Add the given key-value pair to the sm_config field of the given VDI.
func (_class VDIClass) AddToSmConfig(sessionID SessionRef, self VDIRef, key string, value string) (_err error) {
	if IsMock {
		return _class.AddToSmConfigMock(sessionID, self, key, value)
	}	
	_method := "VDI.add_to_sm_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func VDIClassSetSmConfigMockDefault(sessionID SessionRef, self VDIRef, value map[string]string) (_err error) {
	log.Println("VDI.SetSmConfig not mocked")
	_err = errors.New("VDI.SetSmConfig not mocked")
	return
}

var VDIClassSetSmConfigMockedCallback = VDIClassSetSmConfigMockDefault

func (_class VDIClass) SetSmConfigMock(sessionID SessionRef, self VDIRef, value map[string]string) (_err error) {
	return VDIClassSetSmConfigMockedCallback(sessionID, self, value)
}
// Set the sm_config field of the given VDI.
func (_class VDIClass) SetSmConfig(sessionID SessionRef, self VDIRef, value map[string]string) (_err error) {
	if IsMock {
		return _class.SetSmConfigMock(sessionID, self, value)
	}	
	_method := "VDI.set_sm_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func VDIClassRemoveFromXenstoreDataMockDefault(sessionID SessionRef, self VDIRef, key string) (_err error) {
	log.Println("VDI.RemoveFromXenstoreData not mocked")
	_err = errors.New("VDI.RemoveFromXenstoreData not mocked")
	return
}

var VDIClassRemoveFromXenstoreDataMockedCallback = VDIClassRemoveFromXenstoreDataMockDefault

func (_class VDIClass) RemoveFromXenstoreDataMock(sessionID SessionRef, self VDIRef, key string) (_err error) {
	return VDIClassRemoveFromXenstoreDataMockedCallback(sessionID, self, key)
}
// Remove the given key and its corresponding value from the xenstore_data field of the given VDI.  If the key is not in that Map, then do nothing.
func (_class VDIClass) RemoveFromXenstoreData(sessionID SessionRef, self VDIRef, key string) (_err error) {
	if IsMock {
		return _class.RemoveFromXenstoreDataMock(sessionID, self, key)
	}	
	_method := "VDI.remove_from_xenstore_data"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func VDIClassAddToXenstoreDataMockDefault(sessionID SessionRef, self VDIRef, key string, value string) (_err error) {
	log.Println("VDI.AddToXenstoreData not mocked")
	_err = errors.New("VDI.AddToXenstoreData not mocked")
	return
}

var VDIClassAddToXenstoreDataMockedCallback = VDIClassAddToXenstoreDataMockDefault

func (_class VDIClass) AddToXenstoreDataMock(sessionID SessionRef, self VDIRef, key string, value string) (_err error) {
	return VDIClassAddToXenstoreDataMockedCallback(sessionID, self, key, value)
}
// Add the given key-value pair to the xenstore_data field of the given VDI.
func (_class VDIClass) AddToXenstoreData(sessionID SessionRef, self VDIRef, key string, value string) (_err error) {
	if IsMock {
		return _class.AddToXenstoreDataMock(sessionID, self, key, value)
	}	
	_method := "VDI.add_to_xenstore_data"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func VDIClassSetXenstoreDataMockDefault(sessionID SessionRef, self VDIRef, value map[string]string) (_err error) {
	log.Println("VDI.SetXenstoreData not mocked")
	_err = errors.New("VDI.SetXenstoreData not mocked")
	return
}

var VDIClassSetXenstoreDataMockedCallback = VDIClassSetXenstoreDataMockDefault

func (_class VDIClass) SetXenstoreDataMock(sessionID SessionRef, self VDIRef, value map[string]string) (_err error) {
	return VDIClassSetXenstoreDataMockedCallback(sessionID, self, value)
}
// Set the xenstore_data field of the given VDI.
func (_class VDIClass) SetXenstoreData(sessionID SessionRef, self VDIRef, value map[string]string) (_err error) {
	if IsMock {
		return _class.SetXenstoreDataMock(sessionID, self, value)
	}	
	_method := "VDI.set_xenstore_data"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func VDIClassRemoveFromOtherConfigMockDefault(sessionID SessionRef, self VDIRef, key string) (_err error) {
	log.Println("VDI.RemoveFromOtherConfig not mocked")
	_err = errors.New("VDI.RemoveFromOtherConfig not mocked")
	return
}

var VDIClassRemoveFromOtherConfigMockedCallback = VDIClassRemoveFromOtherConfigMockDefault

func (_class VDIClass) RemoveFromOtherConfigMock(sessionID SessionRef, self VDIRef, key string) (_err error) {
	return VDIClassRemoveFromOtherConfigMockedCallback(sessionID, self, key)
}
// Remove the given key and its corresponding value from the other_config field of the given VDI.  If the key is not in that Map, then do nothing.
func (_class VDIClass) RemoveFromOtherConfig(sessionID SessionRef, self VDIRef, key string) (_err error) {
	if IsMock {
		return _class.RemoveFromOtherConfigMock(sessionID, self, key)
	}	
	_method := "VDI.remove_from_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func VDIClassAddToOtherConfigMockDefault(sessionID SessionRef, self VDIRef, key string, value string) (_err error) {
	log.Println("VDI.AddToOtherConfig not mocked")
	_err = errors.New("VDI.AddToOtherConfig not mocked")
	return
}

var VDIClassAddToOtherConfigMockedCallback = VDIClassAddToOtherConfigMockDefault

func (_class VDIClass) AddToOtherConfigMock(sessionID SessionRef, self VDIRef, key string, value string) (_err error) {
	return VDIClassAddToOtherConfigMockedCallback(sessionID, self, key, value)
}
// Add the given key-value pair to the other_config field of the given VDI.
func (_class VDIClass) AddToOtherConfig(sessionID SessionRef, self VDIRef, key string, value string) (_err error) {
	if IsMock {
		return _class.AddToOtherConfigMock(sessionID, self, key, value)
	}	
	_method := "VDI.add_to_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func VDIClassSetOtherConfigMockDefault(sessionID SessionRef, self VDIRef, value map[string]string) (_err error) {
	log.Println("VDI.SetOtherConfig not mocked")
	_err = errors.New("VDI.SetOtherConfig not mocked")
	return
}

var VDIClassSetOtherConfigMockedCallback = VDIClassSetOtherConfigMockDefault

func (_class VDIClass) SetOtherConfigMock(sessionID SessionRef, self VDIRef, value map[string]string) (_err error) {
	return VDIClassSetOtherConfigMockedCallback(sessionID, self, value)
}
// Set the other_config field of the given VDI.
func (_class VDIClass) SetOtherConfig(sessionID SessionRef, self VDIRef, value map[string]string) (_err error) {
	if IsMock {
		return _class.SetOtherConfigMock(sessionID, self, value)
	}	
	_method := "VDI.set_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func VDIClassGetCbtEnabledMockDefault(sessionID SessionRef, self VDIRef) (_retval bool, _err error) {
	log.Println("VDI.GetCbtEnabled not mocked")
	_err = errors.New("VDI.GetCbtEnabled not mocked")
	return
}

var VDIClassGetCbtEnabledMockedCallback = VDIClassGetCbtEnabledMockDefault

func (_class VDIClass) GetCbtEnabledMock(sessionID SessionRef, self VDIRef) (_retval bool, _err error) {
	return VDIClassGetCbtEnabledMockedCallback(sessionID, self)
}
// Get the cbt_enabled field of the given VDI.
func (_class VDIClass) GetCbtEnabled(sessionID SessionRef, self VDIRef) (_retval bool, _err error) {
	if IsMock {
		return _class.GetCbtEnabledMock(sessionID, self)
	}	
	_method := "VDI.get_cbt_enabled"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func VDIClassGetIsToolsIsoMockDefault(sessionID SessionRef, self VDIRef) (_retval bool, _err error) {
	log.Println("VDI.GetIsToolsIso not mocked")
	_err = errors.New("VDI.GetIsToolsIso not mocked")
	return
}

var VDIClassGetIsToolsIsoMockedCallback = VDIClassGetIsToolsIsoMockDefault

func (_class VDIClass) GetIsToolsIsoMock(sessionID SessionRef, self VDIRef) (_retval bool, _err error) {
	return VDIClassGetIsToolsIsoMockedCallback(sessionID, self)
}
// Get the is_tools_iso field of the given VDI.
func (_class VDIClass) GetIsToolsIso(sessionID SessionRef, self VDIRef) (_retval bool, _err error) {
	if IsMock {
		return _class.GetIsToolsIsoMock(sessionID, self)
	}	
	_method := "VDI.get_is_tools_iso"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func VDIClassGetMetadataLatestMockDefault(sessionID SessionRef, self VDIRef) (_retval bool, _err error) {
	log.Println("VDI.GetMetadataLatest not mocked")
	_err = errors.New("VDI.GetMetadataLatest not mocked")
	return
}

var VDIClassGetMetadataLatestMockedCallback = VDIClassGetMetadataLatestMockDefault

func (_class VDIClass) GetMetadataLatestMock(sessionID SessionRef, self VDIRef) (_retval bool, _err error) {
	return VDIClassGetMetadataLatestMockedCallback(sessionID, self)
}
// Get the metadata_latest field of the given VDI.
func (_class VDIClass) GetMetadataLatest(sessionID SessionRef, self VDIRef) (_retval bool, _err error) {
	if IsMock {
		return _class.GetMetadataLatestMock(sessionID, self)
	}	
	_method := "VDI.get_metadata_latest"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func VDIClassGetMetadataOfPoolMockDefault(sessionID SessionRef, self VDIRef) (_retval PoolRef, _err error) {
	log.Println("VDI.GetMetadataOfPool not mocked")
	_err = errors.New("VDI.GetMetadataOfPool not mocked")
	return
}

var VDIClassGetMetadataOfPoolMockedCallback = VDIClassGetMetadataOfPoolMockDefault

func (_class VDIClass) GetMetadataOfPoolMock(sessionID SessionRef, self VDIRef) (_retval PoolRef, _err error) {
	return VDIClassGetMetadataOfPoolMockedCallback(sessionID, self)
}
// Get the metadata_of_pool field of the given VDI.
func (_class VDIClass) GetMetadataOfPool(sessionID SessionRef, self VDIRef) (_retval PoolRef, _err error) {
	if IsMock {
		return _class.GetMetadataOfPoolMock(sessionID, self)
	}	
	_method := "VDI.get_metadata_of_pool"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertPoolRefToGo(_method + " -> ", _result.Value)
	return
}


func VDIClassGetOnBootMockDefault(sessionID SessionRef, self VDIRef) (_retval OnBoot, _err error) {
	log.Println("VDI.GetOnBoot not mocked")
	_err = errors.New("VDI.GetOnBoot not mocked")
	return
}

var VDIClassGetOnBootMockedCallback = VDIClassGetOnBootMockDefault

func (_class VDIClass) GetOnBootMock(sessionID SessionRef, self VDIRef) (_retval OnBoot, _err error) {
	return VDIClassGetOnBootMockedCallback(sessionID, self)
}
// Get the on_boot field of the given VDI.
func (_class VDIClass) GetOnBoot(sessionID SessionRef, self VDIRef) (_retval OnBoot, _err error) {
	if IsMock {
		return _class.GetOnBootMock(sessionID, self)
	}	
	_method := "VDI.get_on_boot"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertEnumOnBootToGo(_method + " -> ", _result.Value)
	return
}


func VDIClassGetAllowCachingMockDefault(sessionID SessionRef, self VDIRef) (_retval bool, _err error) {
	log.Println("VDI.GetAllowCaching not mocked")
	_err = errors.New("VDI.GetAllowCaching not mocked")
	return
}

var VDIClassGetAllowCachingMockedCallback = VDIClassGetAllowCachingMockDefault

func (_class VDIClass) GetAllowCachingMock(sessionID SessionRef, self VDIRef) (_retval bool, _err error) {
	return VDIClassGetAllowCachingMockedCallback(sessionID, self)
}
// Get the allow_caching field of the given VDI.
func (_class VDIClass) GetAllowCaching(sessionID SessionRef, self VDIRef) (_retval bool, _err error) {
	if IsMock {
		return _class.GetAllowCachingMock(sessionID, self)
	}	
	_method := "VDI.get_allow_caching"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func VDIClassGetTagsMockDefault(sessionID SessionRef, self VDIRef) (_retval []string, _err error) {
	log.Println("VDI.GetTags not mocked")
	_err = errors.New("VDI.GetTags not mocked")
	return
}

var VDIClassGetTagsMockedCallback = VDIClassGetTagsMockDefault

func (_class VDIClass) GetTagsMock(sessionID SessionRef, self VDIRef) (_retval []string, _err error) {
	return VDIClassGetTagsMockedCallback(sessionID, self)
}
// Get the tags field of the given VDI.
func (_class VDIClass) GetTags(sessionID SessionRef, self VDIRef) (_retval []string, _err error) {
	if IsMock {
		return _class.GetTagsMock(sessionID, self)
	}	
	_method := "VDI.get_tags"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func VDIClassGetSnapshotTimeMockDefault(sessionID SessionRef, self VDIRef) (_retval time.Time, _err error) {
	log.Println("VDI.GetSnapshotTime not mocked")
	_err = errors.New("VDI.GetSnapshotTime not mocked")
	return
}

var VDIClassGetSnapshotTimeMockedCallback = VDIClassGetSnapshotTimeMockDefault

func (_class VDIClass) GetSnapshotTimeMock(sessionID SessionRef, self VDIRef) (_retval time.Time, _err error) {
	return VDIClassGetSnapshotTimeMockedCallback(sessionID, self)
}
// Get the snapshot_time field of the given VDI.
func (_class VDIClass) GetSnapshotTime(sessionID SessionRef, self VDIRef) (_retval time.Time, _err error) {
	if IsMock {
		return _class.GetSnapshotTimeMock(sessionID, self)
	}	
	_method := "VDI.get_snapshot_time"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func VDIClassGetSnapshotsMockDefault(sessionID SessionRef, self VDIRef) (_retval []VDIRef, _err error) {
	log.Println("VDI.GetSnapshots not mocked")
	_err = errors.New("VDI.GetSnapshots not mocked")
	return
}

var VDIClassGetSnapshotsMockedCallback = VDIClassGetSnapshotsMockDefault

func (_class VDIClass) GetSnapshotsMock(sessionID SessionRef, self VDIRef) (_retval []VDIRef, _err error) {
	return VDIClassGetSnapshotsMockedCallback(sessionID, self)
}
// Get the snapshots field of the given VDI.
func (_class VDIClass) GetSnapshots(sessionID SessionRef, self VDIRef) (_retval []VDIRef, _err error) {
	if IsMock {
		return _class.GetSnapshotsMock(sessionID, self)
	}	
	_method := "VDI.get_snapshots"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func VDIClassGetSnapshotOfMockDefault(sessionID SessionRef, self VDIRef) (_retval VDIRef, _err error) {
	log.Println("VDI.GetSnapshotOf not mocked")
	_err = errors.New("VDI.GetSnapshotOf not mocked")
	return
}

var VDIClassGetSnapshotOfMockedCallback = VDIClassGetSnapshotOfMockDefault

func (_class VDIClass) GetSnapshotOfMock(sessionID SessionRef, self VDIRef) (_retval VDIRef, _err error) {
	return VDIClassGetSnapshotOfMockedCallback(sessionID, self)
}
// Get the snapshot_of field of the given VDI.
func (_class VDIClass) GetSnapshotOf(sessionID SessionRef, self VDIRef) (_retval VDIRef, _err error) {
	if IsMock {
		return _class.GetSnapshotOfMock(sessionID, self)
	}	
	_method := "VDI.get_snapshot_of"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVDIRefToGo(_method + " -> ", _result.Value)
	return
}


func VDIClassGetIsASnapshotMockDefault(sessionID SessionRef, self VDIRef) (_retval bool, _err error) {
	log.Println("VDI.GetIsASnapshot not mocked")
	_err = errors.New("VDI.GetIsASnapshot not mocked")
	return
}

var VDIClassGetIsASnapshotMockedCallback = VDIClassGetIsASnapshotMockDefault

func (_class VDIClass) GetIsASnapshotMock(sessionID SessionRef, self VDIRef) (_retval bool, _err error) {
	return VDIClassGetIsASnapshotMockedCallback(sessionID, self)
}
// Get the is_a_snapshot field of the given VDI.
func (_class VDIClass) GetIsASnapshot(sessionID SessionRef, self VDIRef) (_retval bool, _err error) {
	if IsMock {
		return _class.GetIsASnapshotMock(sessionID, self)
	}	
	_method := "VDI.get_is_a_snapshot"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func VDIClassGetSmConfigMockDefault(sessionID SessionRef, self VDIRef) (_retval map[string]string, _err error) {
	log.Println("VDI.GetSmConfig not mocked")
	_err = errors.New("VDI.GetSmConfig not mocked")
	return
}

var VDIClassGetSmConfigMockedCallback = VDIClassGetSmConfigMockDefault

func (_class VDIClass) GetSmConfigMock(sessionID SessionRef, self VDIRef) (_retval map[string]string, _err error) {
	return VDIClassGetSmConfigMockedCallback(sessionID, self)
}
// Get the sm_config field of the given VDI.
func (_class VDIClass) GetSmConfig(sessionID SessionRef, self VDIRef) (_retval map[string]string, _err error) {
	if IsMock {
		return _class.GetSmConfigMock(sessionID, self)
	}	
	_method := "VDI.get_sm_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func VDIClassGetXenstoreDataMockDefault(sessionID SessionRef, self VDIRef) (_retval map[string]string, _err error) {
	log.Println("VDI.GetXenstoreData not mocked")
	_err = errors.New("VDI.GetXenstoreData not mocked")
	return
}

var VDIClassGetXenstoreDataMockedCallback = VDIClassGetXenstoreDataMockDefault

func (_class VDIClass) GetXenstoreDataMock(sessionID SessionRef, self VDIRef) (_retval map[string]string, _err error) {
	return VDIClassGetXenstoreDataMockedCallback(sessionID, self)
}
// Get the xenstore_data field of the given VDI.
func (_class VDIClass) GetXenstoreData(sessionID SessionRef, self VDIRef) (_retval map[string]string, _err error) {
	if IsMock {
		return _class.GetXenstoreDataMock(sessionID, self)
	}	
	_method := "VDI.get_xenstore_data"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func VDIClassGetParentMockDefault(sessionID SessionRef, self VDIRef) (_retval VDIRef, _err error) {
	log.Println("VDI.GetParent not mocked")
	_err = errors.New("VDI.GetParent not mocked")
	return
}

var VDIClassGetParentMockedCallback = VDIClassGetParentMockDefault

func (_class VDIClass) GetParentMock(sessionID SessionRef, self VDIRef) (_retval VDIRef, _err error) {
	return VDIClassGetParentMockedCallback(sessionID, self)
}
// Get the parent field of the given VDI.
func (_class VDIClass) GetParent(sessionID SessionRef, self VDIRef) (_retval VDIRef, _err error) {
	if IsMock {
		return _class.GetParentMock(sessionID, self)
	}	
	_method := "VDI.get_parent"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVDIRefToGo(_method + " -> ", _result.Value)
	return
}


func VDIClassGetMissingMockDefault(sessionID SessionRef, self VDIRef) (_retval bool, _err error) {
	log.Println("VDI.GetMissing not mocked")
	_err = errors.New("VDI.GetMissing not mocked")
	return
}

var VDIClassGetMissingMockedCallback = VDIClassGetMissingMockDefault

func (_class VDIClass) GetMissingMock(sessionID SessionRef, self VDIRef) (_retval bool, _err error) {
	return VDIClassGetMissingMockedCallback(sessionID, self)
}
// Get the missing field of the given VDI.
func (_class VDIClass) GetMissing(sessionID SessionRef, self VDIRef) (_retval bool, _err error) {
	if IsMock {
		return _class.GetMissingMock(sessionID, self)
	}	
	_method := "VDI.get_missing"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func VDIClassGetManagedMockDefault(sessionID SessionRef, self VDIRef) (_retval bool, _err error) {
	log.Println("VDI.GetManaged not mocked")
	_err = errors.New("VDI.GetManaged not mocked")
	return
}

var VDIClassGetManagedMockedCallback = VDIClassGetManagedMockDefault

func (_class VDIClass) GetManagedMock(sessionID SessionRef, self VDIRef) (_retval bool, _err error) {
	return VDIClassGetManagedMockedCallback(sessionID, self)
}
// Get the managed field of the given VDI.
func (_class VDIClass) GetManaged(sessionID SessionRef, self VDIRef) (_retval bool, _err error) {
	if IsMock {
		return _class.GetManagedMock(sessionID, self)
	}	
	_method := "VDI.get_managed"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func VDIClassGetLocationMockDefault(sessionID SessionRef, self VDIRef) (_retval string, _err error) {
	log.Println("VDI.GetLocation not mocked")
	_err = errors.New("VDI.GetLocation not mocked")
	return
}

var VDIClassGetLocationMockedCallback = VDIClassGetLocationMockDefault

func (_class VDIClass) GetLocationMock(sessionID SessionRef, self VDIRef) (_retval string, _err error) {
	return VDIClassGetLocationMockedCallback(sessionID, self)
}
// Get the location field of the given VDI.
func (_class VDIClass) GetLocation(sessionID SessionRef, self VDIRef) (_retval string, _err error) {
	if IsMock {
		return _class.GetLocationMock(sessionID, self)
	}	
	_method := "VDI.get_location"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func VDIClassGetStorageLockMockDefault(sessionID SessionRef, self VDIRef) (_retval bool, _err error) {
	log.Println("VDI.GetStorageLock not mocked")
	_err = errors.New("VDI.GetStorageLock not mocked")
	return
}

var VDIClassGetStorageLockMockedCallback = VDIClassGetStorageLockMockDefault

func (_class VDIClass) GetStorageLockMock(sessionID SessionRef, self VDIRef) (_retval bool, _err error) {
	return VDIClassGetStorageLockMockedCallback(sessionID, self)
}
// Get the storage_lock field of the given VDI.
func (_class VDIClass) GetStorageLock(sessionID SessionRef, self VDIRef) (_retval bool, _err error) {
	if IsMock {
		return _class.GetStorageLockMock(sessionID, self)
	}	
	_method := "VDI.get_storage_lock"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func VDIClassGetOtherConfigMockDefault(sessionID SessionRef, self VDIRef) (_retval map[string]string, _err error) {
	log.Println("VDI.GetOtherConfig not mocked")
	_err = errors.New("VDI.GetOtherConfig not mocked")
	return
}

var VDIClassGetOtherConfigMockedCallback = VDIClassGetOtherConfigMockDefault

func (_class VDIClass) GetOtherConfigMock(sessionID SessionRef, self VDIRef) (_retval map[string]string, _err error) {
	return VDIClassGetOtherConfigMockedCallback(sessionID, self)
}
// Get the other_config field of the given VDI.
func (_class VDIClass) GetOtherConfig(sessionID SessionRef, self VDIRef) (_retval map[string]string, _err error) {
	if IsMock {
		return _class.GetOtherConfigMock(sessionID, self)
	}	
	_method := "VDI.get_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func VDIClassGetReadOnlyMockDefault(sessionID SessionRef, self VDIRef) (_retval bool, _err error) {
	log.Println("VDI.GetReadOnly not mocked")
	_err = errors.New("VDI.GetReadOnly not mocked")
	return
}

var VDIClassGetReadOnlyMockedCallback = VDIClassGetReadOnlyMockDefault

func (_class VDIClass) GetReadOnlyMock(sessionID SessionRef, self VDIRef) (_retval bool, _err error) {
	return VDIClassGetReadOnlyMockedCallback(sessionID, self)
}
// Get the read_only field of the given VDI.
func (_class VDIClass) GetReadOnly(sessionID SessionRef, self VDIRef) (_retval bool, _err error) {
	if IsMock {
		return _class.GetReadOnlyMock(sessionID, self)
	}	
	_method := "VDI.get_read_only"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func VDIClassGetSharableMockDefault(sessionID SessionRef, self VDIRef) (_retval bool, _err error) {
	log.Println("VDI.GetSharable not mocked")
	_err = errors.New("VDI.GetSharable not mocked")
	return
}

var VDIClassGetSharableMockedCallback = VDIClassGetSharableMockDefault

func (_class VDIClass) GetSharableMock(sessionID SessionRef, self VDIRef) (_retval bool, _err error) {
	return VDIClassGetSharableMockedCallback(sessionID, self)
}
// Get the sharable field of the given VDI.
func (_class VDIClass) GetSharable(sessionID SessionRef, self VDIRef) (_retval bool, _err error) {
	if IsMock {
		return _class.GetSharableMock(sessionID, self)
	}	
	_method := "VDI.get_sharable"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func VDIClassGetTypeMockDefault(sessionID SessionRef, self VDIRef) (_retval VdiType, _err error) {
	log.Println("VDI.GetType not mocked")
	_err = errors.New("VDI.GetType not mocked")
	return
}

var VDIClassGetTypeMockedCallback = VDIClassGetTypeMockDefault

func (_class VDIClass) GetTypeMock(sessionID SessionRef, self VDIRef) (_retval VdiType, _err error) {
	return VDIClassGetTypeMockedCallback(sessionID, self)
}
// Get the type field of the given VDI.
func (_class VDIClass) GetType(sessionID SessionRef, self VDIRef) (_retval VdiType, _err error) {
	if IsMock {
		return _class.GetTypeMock(sessionID, self)
	}	
	_method := "VDI.get_type"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertEnumVdiTypeToGo(_method + " -> ", _result.Value)
	return
}


func VDIClassGetPhysicalUtilisationMockDefault(sessionID SessionRef, self VDIRef) (_retval int, _err error) {
	log.Println("VDI.GetPhysicalUtilisation not mocked")
	_err = errors.New("VDI.GetPhysicalUtilisation not mocked")
	return
}

var VDIClassGetPhysicalUtilisationMockedCallback = VDIClassGetPhysicalUtilisationMockDefault

func (_class VDIClass) GetPhysicalUtilisationMock(sessionID SessionRef, self VDIRef) (_retval int, _err error) {
	return VDIClassGetPhysicalUtilisationMockedCallback(sessionID, self)
}
// Get the physical_utilisation field of the given VDI.
func (_class VDIClass) GetPhysicalUtilisation(sessionID SessionRef, self VDIRef) (_retval int, _err error) {
	if IsMock {
		return _class.GetPhysicalUtilisationMock(sessionID, self)
	}	
	_method := "VDI.get_physical_utilisation"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func VDIClassGetVirtualSizeMockDefault(sessionID SessionRef, self VDIRef) (_retval int, _err error) {
	log.Println("VDI.GetVirtualSize not mocked")
	_err = errors.New("VDI.GetVirtualSize not mocked")
	return
}

var VDIClassGetVirtualSizeMockedCallback = VDIClassGetVirtualSizeMockDefault

func (_class VDIClass) GetVirtualSizeMock(sessionID SessionRef, self VDIRef) (_retval int, _err error) {
	return VDIClassGetVirtualSizeMockedCallback(sessionID, self)
}
// Get the virtual_size field of the given VDI.
func (_class VDIClass) GetVirtualSize(sessionID SessionRef, self VDIRef) (_retval int, _err error) {
	if IsMock {
		return _class.GetVirtualSizeMock(sessionID, self)
	}	
	_method := "VDI.get_virtual_size"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func VDIClassGetCrashDumpsMockDefault(sessionID SessionRef, self VDIRef) (_retval []CrashdumpRef, _err error) {
	log.Println("VDI.GetCrashDumps not mocked")
	_err = errors.New("VDI.GetCrashDumps not mocked")
	return
}

var VDIClassGetCrashDumpsMockedCallback = VDIClassGetCrashDumpsMockDefault

func (_class VDIClass) GetCrashDumpsMock(sessionID SessionRef, self VDIRef) (_retval []CrashdumpRef, _err error) {
	return VDIClassGetCrashDumpsMockedCallback(sessionID, self)
}
// Get the crash_dumps field of the given VDI.
func (_class VDIClass) GetCrashDumps(sessionID SessionRef, self VDIRef) (_retval []CrashdumpRef, _err error) {
	if IsMock {
		return _class.GetCrashDumpsMock(sessionID, self)
	}	
	_method := "VDI.get_crash_dumps"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertCrashdumpRefSetToGo(_method + " -> ", _result.Value)
	return
}


func VDIClassGetVBDsMockDefault(sessionID SessionRef, self VDIRef) (_retval []VBDRef, _err error) {
	log.Println("VDI.GetVBDs not mocked")
	_err = errors.New("VDI.GetVBDs not mocked")
	return
}

var VDIClassGetVBDsMockedCallback = VDIClassGetVBDsMockDefault

func (_class VDIClass) GetVBDsMock(sessionID SessionRef, self VDIRef) (_retval []VBDRef, _err error) {
	return VDIClassGetVBDsMockedCallback(sessionID, self)
}
// Get the VBDs field of the given VDI.
func (_class VDIClass) GetVBDs(sessionID SessionRef, self VDIRef) (_retval []VBDRef, _err error) {
	if IsMock {
		return _class.GetVBDsMock(sessionID, self)
	}	
	_method := "VDI.get_VBDs"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVBDRefSetToGo(_method + " -> ", _result.Value)
	return
}


func VDIClassGetSRMockDefault(sessionID SessionRef, self VDIRef) (_retval SRRef, _err error) {
	log.Println("VDI.GetSR not mocked")
	_err = errors.New("VDI.GetSR not mocked")
	return
}

var VDIClassGetSRMockedCallback = VDIClassGetSRMockDefault

func (_class VDIClass) GetSRMock(sessionID SessionRef, self VDIRef) (_retval SRRef, _err error) {
	return VDIClassGetSRMockedCallback(sessionID, self)
}
// Get the SR field of the given VDI.
func (_class VDIClass) GetSR(sessionID SessionRef, self VDIRef) (_retval SRRef, _err error) {
	if IsMock {
		return _class.GetSRMock(sessionID, self)
	}	
	_method := "VDI.get_SR"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertSRRefToGo(_method + " -> ", _result.Value)
	return
}


func VDIClassGetCurrentOperationsMockDefault(sessionID SessionRef, self VDIRef) (_retval map[string]VdiOperations, _err error) {
	log.Println("VDI.GetCurrentOperations not mocked")
	_err = errors.New("VDI.GetCurrentOperations not mocked")
	return
}

var VDIClassGetCurrentOperationsMockedCallback = VDIClassGetCurrentOperationsMockDefault

func (_class VDIClass) GetCurrentOperationsMock(sessionID SessionRef, self VDIRef) (_retval map[string]VdiOperations, _err error) {
	return VDIClassGetCurrentOperationsMockedCallback(sessionID, self)
}
// Get the current_operations field of the given VDI.
func (_class VDIClass) GetCurrentOperations(sessionID SessionRef, self VDIRef) (_retval map[string]VdiOperations, _err error) {
	if IsMock {
		return _class.GetCurrentOperationsMock(sessionID, self)
	}	
	_method := "VDI.get_current_operations"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringToEnumVdiOperationsMapToGo(_method + " -> ", _result.Value)
	return
}


func VDIClassGetAllowedOperationsMockDefault(sessionID SessionRef, self VDIRef) (_retval []VdiOperations, _err error) {
	log.Println("VDI.GetAllowedOperations not mocked")
	_err = errors.New("VDI.GetAllowedOperations not mocked")
	return
}

var VDIClassGetAllowedOperationsMockedCallback = VDIClassGetAllowedOperationsMockDefault

func (_class VDIClass) GetAllowedOperationsMock(sessionID SessionRef, self VDIRef) (_retval []VdiOperations, _err error) {
	return VDIClassGetAllowedOperationsMockedCallback(sessionID, self)
}
// Get the allowed_operations field of the given VDI.
func (_class VDIClass) GetAllowedOperations(sessionID SessionRef, self VDIRef) (_retval []VdiOperations, _err error) {
	if IsMock {
		return _class.GetAllowedOperationsMock(sessionID, self)
	}	
	_method := "VDI.get_allowed_operations"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertEnumVdiOperationsSetToGo(_method + " -> ", _result.Value)
	return
}


func VDIClassGetNameDescriptionMockDefault(sessionID SessionRef, self VDIRef) (_retval string, _err error) {
	log.Println("VDI.GetNameDescription not mocked")
	_err = errors.New("VDI.GetNameDescription not mocked")
	return
}

var VDIClassGetNameDescriptionMockedCallback = VDIClassGetNameDescriptionMockDefault

func (_class VDIClass) GetNameDescriptionMock(sessionID SessionRef, self VDIRef) (_retval string, _err error) {
	return VDIClassGetNameDescriptionMockedCallback(sessionID, self)
}
// Get the name/description field of the given VDI.
func (_class VDIClass) GetNameDescription(sessionID SessionRef, self VDIRef) (_retval string, _err error) {
	if IsMock {
		return _class.GetNameDescriptionMock(sessionID, self)
	}	
	_method := "VDI.get_name_description"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func VDIClassGetNameLabelMockDefault(sessionID SessionRef, self VDIRef) (_retval string, _err error) {
	log.Println("VDI.GetNameLabel not mocked")
	_err = errors.New("VDI.GetNameLabel not mocked")
	return
}

var VDIClassGetNameLabelMockedCallback = VDIClassGetNameLabelMockDefault

func (_class VDIClass) GetNameLabelMock(sessionID SessionRef, self VDIRef) (_retval string, _err error) {
	return VDIClassGetNameLabelMockedCallback(sessionID, self)
}
// Get the name/label field of the given VDI.
func (_class VDIClass) GetNameLabel(sessionID SessionRef, self VDIRef) (_retval string, _err error) {
	if IsMock {
		return _class.GetNameLabelMock(sessionID, self)
	}	
	_method := "VDI.get_name_label"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func VDIClassGetUUIDMockDefault(sessionID SessionRef, self VDIRef) (_retval string, _err error) {
	log.Println("VDI.GetUUID not mocked")
	_err = errors.New("VDI.GetUUID not mocked")
	return
}

var VDIClassGetUUIDMockedCallback = VDIClassGetUUIDMockDefault

func (_class VDIClass) GetUUIDMock(sessionID SessionRef, self VDIRef) (_retval string, _err error) {
	return VDIClassGetUUIDMockedCallback(sessionID, self)
}
// Get the uuid field of the given VDI.
func (_class VDIClass) GetUUID(sessionID SessionRef, self VDIRef) (_retval string, _err error) {
	if IsMock {
		return _class.GetUUIDMock(sessionID, self)
	}	
	_method := "VDI.get_uuid"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func VDIClassGetByNameLabelMockDefault(sessionID SessionRef, label string) (_retval []VDIRef, _err error) {
	log.Println("VDI.GetByNameLabel not mocked")
	_err = errors.New("VDI.GetByNameLabel not mocked")
	return
}

var VDIClassGetByNameLabelMockedCallback = VDIClassGetByNameLabelMockDefault

func (_class VDIClass) GetByNameLabelMock(sessionID SessionRef, label string) (_retval []VDIRef, _err error) {
	return VDIClassGetByNameLabelMockedCallback(sessionID, label)
}
// Get all the VDI instances with the given label.
func (_class VDIClass) GetByNameLabel(sessionID SessionRef, label string) (_retval []VDIRef, _err error) {
	if IsMock {
		return _class.GetByNameLabelMock(sessionID, label)
	}	
	_method := "VDI.get_by_name_label"
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
	_retval, _err = convertVDIRefSetToGo(_method + " -> ", _result.Value)
	return
}


func VDIClassDestroyMockDefault(sessionID SessionRef, self VDIRef) (_err error) {
	log.Println("VDI.Destroy not mocked")
	_err = errors.New("VDI.Destroy not mocked")
	return
}

var VDIClassDestroyMockedCallback = VDIClassDestroyMockDefault

func (_class VDIClass) DestroyMock(sessionID SessionRef, self VDIRef) (_err error) {
	return VDIClassDestroyMockedCallback(sessionID, self)
}
// Destroy the specified VDI instance.
func (_class VDIClass) Destroy(sessionID SessionRef, self VDIRef) (_err error) {
	if IsMock {
		return _class.DestroyMock(sessionID, self)
	}	
	_method := "VDI.destroy"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg)
	return
}


func VDIClassCreateMockDefault(sessionID SessionRef, args VDIRecord) (_retval VDIRef, _err error) {
	log.Println("VDI.Create not mocked")
	_err = errors.New("VDI.Create not mocked")
	return
}

var VDIClassCreateMockedCallback = VDIClassCreateMockDefault

func (_class VDIClass) CreateMock(sessionID SessionRef, args VDIRecord) (_retval VDIRef, _err error) {
	return VDIClassCreateMockedCallback(sessionID, args)
}
// Create a new VDI instance, and return its handle.
// The constructor args are: name_label, name_description, SR*, virtual_size*, type*, sharable*, read_only*, other_config*, xenstore_data, sm_config, tags (* = non-optional).
func (_class VDIClass) Create(sessionID SessionRef, args VDIRecord) (_retval VDIRef, _err error) {
	if IsMock {
		return _class.CreateMock(sessionID, args)
	}	
	_method := "VDI.create"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_argsArg, _err := convertVDIRecordToXen(fmt.Sprintf("%s(%s)", _method, "args"), args)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _argsArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVDIRefToGo(_method + " -> ", _result.Value)
	return
}


func VDIClassGetByUUIDMockDefault(sessionID SessionRef, uuid string) (_retval VDIRef, _err error) {
	log.Println("VDI.GetByUUID not mocked")
	_err = errors.New("VDI.GetByUUID not mocked")
	return
}

var VDIClassGetByUUIDMockedCallback = VDIClassGetByUUIDMockDefault

func (_class VDIClass) GetByUUIDMock(sessionID SessionRef, uuid string) (_retval VDIRef, _err error) {
	return VDIClassGetByUUIDMockedCallback(sessionID, uuid)
}
// Get a reference to the VDI instance with the specified UUID.
func (_class VDIClass) GetByUUID(sessionID SessionRef, uuid string) (_retval VDIRef, _err error) {
	if IsMock {
		return _class.GetByUUIDMock(sessionID, uuid)
	}	
	_method := "VDI.get_by_uuid"
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
	_retval, _err = convertVDIRefToGo(_method + " -> ", _result.Value)
	return
}


func VDIClassGetRecordMockDefault(sessionID SessionRef, self VDIRef) (_retval VDIRecord, _err error) {
	log.Println("VDI.GetRecord not mocked")
	_err = errors.New("VDI.GetRecord not mocked")
	return
}

var VDIClassGetRecordMockedCallback = VDIClassGetRecordMockDefault

func (_class VDIClass) GetRecordMock(sessionID SessionRef, self VDIRef) (_retval VDIRecord, _err error) {
	return VDIClassGetRecordMockedCallback(sessionID, self)
}
// Get a record containing the current state of the given VDI.
func (_class VDIClass) GetRecord(sessionID SessionRef, self VDIRef) (_retval VDIRecord, _err error) {
	if IsMock {
		return _class.GetRecordMock(sessionID, self)
	}	
	_method := "VDI.get_record"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVDIRecordToGo(_method + " -> ", _result.Value)
	return
}
