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

type BlobRecord struct {
  // Unique identifier/object reference
	UUID string
  // a human-readable name
	NameLabel string
  // a notes field containing human-readable description
	NameDescription string
  // Size of the binary data, in bytes
	Size int
  // True if the blob is publicly accessible
	Public bool
  // Time at which the data in the blob was last updated
	LastUpdated time.Time
  // The mime type associated with this object. Defaults to 'application/octet-stream' if the empty string is supplied
	MimeType string
}

type BlobRef string

// A placeholder for a binary blob
type BlobClass struct {
	client *Client
}


func BlobClassGetAllRecordsMockDefault(sessionID SessionRef) (_retval map[BlobRef]BlobRecord, _err error) {
	log.Println("Blob.GetAllRecords not mocked")
	_err = errors.New("Blob.GetAllRecords not mocked")
	return
}

var BlobClassGetAllRecordsMockedCallback = BlobClassGetAllRecordsMockDefault

func (_class BlobClass) GetAllRecordsMock(sessionID SessionRef) (_retval map[BlobRef]BlobRecord, _err error) {
	return BlobClassGetAllRecordsMockedCallback(sessionID)
}
// Return a map of blob references to blob records for all blobs known to the system.
func (_class BlobClass) GetAllRecords(sessionID SessionRef) (_retval map[BlobRef]BlobRecord, _err error) {
	if IsMock {
		return _class.GetAllRecordsMock(sessionID)
	}	
	_method := "blob.get_all_records"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertBlobRefToBlobRecordMapToGo(_method + " -> ", _result.Value)
	return
}


func BlobClassGetAllMockDefault(sessionID SessionRef) (_retval []BlobRef, _err error) {
	log.Println("Blob.GetAll not mocked")
	_err = errors.New("Blob.GetAll not mocked")
	return
}

var BlobClassGetAllMockedCallback = BlobClassGetAllMockDefault

func (_class BlobClass) GetAllMock(sessionID SessionRef) (_retval []BlobRef, _err error) {
	return BlobClassGetAllMockedCallback(sessionID)
}
// Return a list of all the blobs known to the system.
func (_class BlobClass) GetAll(sessionID SessionRef) (_retval []BlobRef, _err error) {
	if IsMock {
		return _class.GetAllMock(sessionID)
	}	
	_method := "blob.get_all"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertBlobRefSetToGo(_method + " -> ", _result.Value)
	return
}


func BlobClassDestroyMockDefault(sessionID SessionRef, self BlobRef) (_err error) {
	log.Println("Blob.Destroy not mocked")
	_err = errors.New("Blob.Destroy not mocked")
	return
}

var BlobClassDestroyMockedCallback = BlobClassDestroyMockDefault

func (_class BlobClass) DestroyMock(sessionID SessionRef, self BlobRef) (_err error) {
	return BlobClassDestroyMockedCallback(sessionID, self)
}
// 
func (_class BlobClass) Destroy(sessionID SessionRef, self BlobRef) (_err error) {
	if IsMock {
		return _class.DestroyMock(sessionID, self)
	}	
	_method := "blob.destroy"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertBlobRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg)
	return
}


func BlobClassCreateMockDefault(sessionID SessionRef, mimeType string, public bool) (_retval BlobRef, _err error) {
	log.Println("Blob.Create not mocked")
	_err = errors.New("Blob.Create not mocked")
	return
}

var BlobClassCreateMockedCallback = BlobClassCreateMockDefault

func (_class BlobClass) CreateMock(sessionID SessionRef, mimeType string, public bool) (_retval BlobRef, _err error) {
	return BlobClassCreateMockedCallback(sessionID, mimeType, public)
}
// Create a placeholder for a binary blob
func (_class BlobClass) Create(sessionID SessionRef, mimeType string, public bool) (_retval BlobRef, _err error) {
	if IsMock {
		return _class.CreateMock(sessionID, mimeType, public)
	}	
	_method := "blob.create"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
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
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _mimeTypeArg, _publicArg)
	if _err != nil {
		return
	}
	_retval, _err = convertBlobRefToGo(_method + " -> ", _result.Value)
	return
}


func BlobClassSetPublicMockDefault(sessionID SessionRef, self BlobRef, value bool) (_err error) {
	log.Println("Blob.SetPublic not mocked")
	_err = errors.New("Blob.SetPublic not mocked")
	return
}

var BlobClassSetPublicMockedCallback = BlobClassSetPublicMockDefault

func (_class BlobClass) SetPublicMock(sessionID SessionRef, self BlobRef, value bool) (_err error) {
	return BlobClassSetPublicMockedCallback(sessionID, self, value)
}
// Set the public field of the given blob.
func (_class BlobClass) SetPublic(sessionID SessionRef, self BlobRef, value bool) (_err error) {
	if IsMock {
		return _class.SetPublicMock(sessionID, self, value)
	}	
	_method := "blob.set_public"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertBlobRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func BlobClassSetNameDescriptionMockDefault(sessionID SessionRef, self BlobRef, value string) (_err error) {
	log.Println("Blob.SetNameDescription not mocked")
	_err = errors.New("Blob.SetNameDescription not mocked")
	return
}

var BlobClassSetNameDescriptionMockedCallback = BlobClassSetNameDescriptionMockDefault

func (_class BlobClass) SetNameDescriptionMock(sessionID SessionRef, self BlobRef, value string) (_err error) {
	return BlobClassSetNameDescriptionMockedCallback(sessionID, self, value)
}
// Set the name/description field of the given blob.
func (_class BlobClass) SetNameDescription(sessionID SessionRef, self BlobRef, value string) (_err error) {
	if IsMock {
		return _class.SetNameDescriptionMock(sessionID, self, value)
	}	
	_method := "blob.set_name_description"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertBlobRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func BlobClassSetNameLabelMockDefault(sessionID SessionRef, self BlobRef, value string) (_err error) {
	log.Println("Blob.SetNameLabel not mocked")
	_err = errors.New("Blob.SetNameLabel not mocked")
	return
}

var BlobClassSetNameLabelMockedCallback = BlobClassSetNameLabelMockDefault

func (_class BlobClass) SetNameLabelMock(sessionID SessionRef, self BlobRef, value string) (_err error) {
	return BlobClassSetNameLabelMockedCallback(sessionID, self, value)
}
// Set the name/label field of the given blob.
func (_class BlobClass) SetNameLabel(sessionID SessionRef, self BlobRef, value string) (_err error) {
	if IsMock {
		return _class.SetNameLabelMock(sessionID, self, value)
	}	
	_method := "blob.set_name_label"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertBlobRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func BlobClassGetMimeTypeMockDefault(sessionID SessionRef, self BlobRef) (_retval string, _err error) {
	log.Println("Blob.GetMimeType not mocked")
	_err = errors.New("Blob.GetMimeType not mocked")
	return
}

var BlobClassGetMimeTypeMockedCallback = BlobClassGetMimeTypeMockDefault

func (_class BlobClass) GetMimeTypeMock(sessionID SessionRef, self BlobRef) (_retval string, _err error) {
	return BlobClassGetMimeTypeMockedCallback(sessionID, self)
}
// Get the mime_type field of the given blob.
func (_class BlobClass) GetMimeType(sessionID SessionRef, self BlobRef) (_retval string, _err error) {
	if IsMock {
		return _class.GetMimeTypeMock(sessionID, self)
	}	
	_method := "blob.get_mime_type"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertBlobRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func BlobClassGetLastUpdatedMockDefault(sessionID SessionRef, self BlobRef) (_retval time.Time, _err error) {
	log.Println("Blob.GetLastUpdated not mocked")
	_err = errors.New("Blob.GetLastUpdated not mocked")
	return
}

var BlobClassGetLastUpdatedMockedCallback = BlobClassGetLastUpdatedMockDefault

func (_class BlobClass) GetLastUpdatedMock(sessionID SessionRef, self BlobRef) (_retval time.Time, _err error) {
	return BlobClassGetLastUpdatedMockedCallback(sessionID, self)
}
// Get the last_updated field of the given blob.
func (_class BlobClass) GetLastUpdated(sessionID SessionRef, self BlobRef) (_retval time.Time, _err error) {
	if IsMock {
		return _class.GetLastUpdatedMock(sessionID, self)
	}	
	_method := "blob.get_last_updated"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertBlobRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func BlobClassGetPublicMockDefault(sessionID SessionRef, self BlobRef) (_retval bool, _err error) {
	log.Println("Blob.GetPublic not mocked")
	_err = errors.New("Blob.GetPublic not mocked")
	return
}

var BlobClassGetPublicMockedCallback = BlobClassGetPublicMockDefault

func (_class BlobClass) GetPublicMock(sessionID SessionRef, self BlobRef) (_retval bool, _err error) {
	return BlobClassGetPublicMockedCallback(sessionID, self)
}
// Get the public field of the given blob.
func (_class BlobClass) GetPublic(sessionID SessionRef, self BlobRef) (_retval bool, _err error) {
	if IsMock {
		return _class.GetPublicMock(sessionID, self)
	}	
	_method := "blob.get_public"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertBlobRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func BlobClassGetSizeMockDefault(sessionID SessionRef, self BlobRef) (_retval int, _err error) {
	log.Println("Blob.GetSize not mocked")
	_err = errors.New("Blob.GetSize not mocked")
	return
}

var BlobClassGetSizeMockedCallback = BlobClassGetSizeMockDefault

func (_class BlobClass) GetSizeMock(sessionID SessionRef, self BlobRef) (_retval int, _err error) {
	return BlobClassGetSizeMockedCallback(sessionID, self)
}
// Get the size field of the given blob.
func (_class BlobClass) GetSize(sessionID SessionRef, self BlobRef) (_retval int, _err error) {
	if IsMock {
		return _class.GetSizeMock(sessionID, self)
	}	
	_method := "blob.get_size"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertBlobRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func BlobClassGetNameDescriptionMockDefault(sessionID SessionRef, self BlobRef) (_retval string, _err error) {
	log.Println("Blob.GetNameDescription not mocked")
	_err = errors.New("Blob.GetNameDescription not mocked")
	return
}

var BlobClassGetNameDescriptionMockedCallback = BlobClassGetNameDescriptionMockDefault

func (_class BlobClass) GetNameDescriptionMock(sessionID SessionRef, self BlobRef) (_retval string, _err error) {
	return BlobClassGetNameDescriptionMockedCallback(sessionID, self)
}
// Get the name/description field of the given blob.
func (_class BlobClass) GetNameDescription(sessionID SessionRef, self BlobRef) (_retval string, _err error) {
	if IsMock {
		return _class.GetNameDescriptionMock(sessionID, self)
	}	
	_method := "blob.get_name_description"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertBlobRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func BlobClassGetNameLabelMockDefault(sessionID SessionRef, self BlobRef) (_retval string, _err error) {
	log.Println("Blob.GetNameLabel not mocked")
	_err = errors.New("Blob.GetNameLabel not mocked")
	return
}

var BlobClassGetNameLabelMockedCallback = BlobClassGetNameLabelMockDefault

func (_class BlobClass) GetNameLabelMock(sessionID SessionRef, self BlobRef) (_retval string, _err error) {
	return BlobClassGetNameLabelMockedCallback(sessionID, self)
}
// Get the name/label field of the given blob.
func (_class BlobClass) GetNameLabel(sessionID SessionRef, self BlobRef) (_retval string, _err error) {
	if IsMock {
		return _class.GetNameLabelMock(sessionID, self)
	}	
	_method := "blob.get_name_label"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertBlobRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func BlobClassGetUUIDMockDefault(sessionID SessionRef, self BlobRef) (_retval string, _err error) {
	log.Println("Blob.GetUUID not mocked")
	_err = errors.New("Blob.GetUUID not mocked")
	return
}

var BlobClassGetUUIDMockedCallback = BlobClassGetUUIDMockDefault

func (_class BlobClass) GetUUIDMock(sessionID SessionRef, self BlobRef) (_retval string, _err error) {
	return BlobClassGetUUIDMockedCallback(sessionID, self)
}
// Get the uuid field of the given blob.
func (_class BlobClass) GetUUID(sessionID SessionRef, self BlobRef) (_retval string, _err error) {
	if IsMock {
		return _class.GetUUIDMock(sessionID, self)
	}	
	_method := "blob.get_uuid"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertBlobRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


func BlobClassGetByNameLabelMockDefault(sessionID SessionRef, label string) (_retval []BlobRef, _err error) {
	log.Println("Blob.GetByNameLabel not mocked")
	_err = errors.New("Blob.GetByNameLabel not mocked")
	return
}

var BlobClassGetByNameLabelMockedCallback = BlobClassGetByNameLabelMockDefault

func (_class BlobClass) GetByNameLabelMock(sessionID SessionRef, label string) (_retval []BlobRef, _err error) {
	return BlobClassGetByNameLabelMockedCallback(sessionID, label)
}
// Get all the blob instances with the given label.
func (_class BlobClass) GetByNameLabel(sessionID SessionRef, label string) (_retval []BlobRef, _err error) {
	if IsMock {
		return _class.GetByNameLabelMock(sessionID, label)
	}	
	_method := "blob.get_by_name_label"
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
	_retval, _err = convertBlobRefSetToGo(_method + " -> ", _result.Value)
	return
}


func BlobClassGetByUUIDMockDefault(sessionID SessionRef, uuid string) (_retval BlobRef, _err error) {
	log.Println("Blob.GetByUUID not mocked")
	_err = errors.New("Blob.GetByUUID not mocked")
	return
}

var BlobClassGetByUUIDMockedCallback = BlobClassGetByUUIDMockDefault

func (_class BlobClass) GetByUUIDMock(sessionID SessionRef, uuid string) (_retval BlobRef, _err error) {
	return BlobClassGetByUUIDMockedCallback(sessionID, uuid)
}
// Get a reference to the blob instance with the specified UUID.
func (_class BlobClass) GetByUUID(sessionID SessionRef, uuid string) (_retval BlobRef, _err error) {
	if IsMock {
		return _class.GetByUUIDMock(sessionID, uuid)
	}	
	_method := "blob.get_by_uuid"
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
	_retval, _err = convertBlobRefToGo(_method + " -> ", _result.Value)
	return
}


func BlobClassGetRecordMockDefault(sessionID SessionRef, self BlobRef) (_retval BlobRecord, _err error) {
	log.Println("Blob.GetRecord not mocked")
	_err = errors.New("Blob.GetRecord not mocked")
	return
}

var BlobClassGetRecordMockedCallback = BlobClassGetRecordMockDefault

func (_class BlobClass) GetRecordMock(sessionID SessionRef, self BlobRef) (_retval BlobRecord, _err error) {
	return BlobClassGetRecordMockedCallback(sessionID, self)
}
// Get a record containing the current state of the given blob.
func (_class BlobClass) GetRecord(sessionID SessionRef, self BlobRef) (_retval BlobRecord, _err error) {
	if IsMock {
		return _class.GetRecordMock(sessionID, self)
	}	
	_method := "blob.get_record"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertBlobRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertBlobRecordToGo(_method + " -> ", _result.Value)
	return
}
