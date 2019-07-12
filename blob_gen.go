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

func (_class BlobClass) GetAllRecords__mock(sessionID SessionRef) (_retval map[BlobRef]BlobRecord, _err error) {
	log.Println("Blob.GetAllRecords not mocked")
	_err = errors.New("Blob.GetAllRecords not mocked")
	return
}
// Return a map of blob references to blob records for all blobs known to the system.
func (_class BlobClass) GetAllRecords(sessionID SessionRef) (_retval map[BlobRef]BlobRecord, _err error) {
	if (IsMock) {
		return _class.GetAllRecords__mock(sessionID)
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

func (_class BlobClass) GetAll__mock(sessionID SessionRef) (_retval []BlobRef, _err error) {
	log.Println("Blob.GetAll not mocked")
	_err = errors.New("Blob.GetAll not mocked")
	return
}
// Return a list of all the blobs known to the system.
func (_class BlobClass) GetAll(sessionID SessionRef) (_retval []BlobRef, _err error) {
	if (IsMock) {
		return _class.GetAll__mock(sessionID)
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

func (_class BlobClass) Destroy__mock(sessionID SessionRef, self BlobRef) (_err error) {
	log.Println("Blob.Destroy not mocked")
	_err = errors.New("Blob.Destroy not mocked")
	return
}
// 
func (_class BlobClass) Destroy(sessionID SessionRef, self BlobRef) (_err error) {
	if (IsMock) {
		return _class.Destroy__mock(sessionID, self)
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

func (_class BlobClass) Create__mock(sessionID SessionRef, mimeType string, public bool) (_retval BlobRef, _err error) {
	log.Println("Blob.Create not mocked")
	_err = errors.New("Blob.Create not mocked")
	return
}
// Create a placeholder for a binary blob
func (_class BlobClass) Create(sessionID SessionRef, mimeType string, public bool) (_retval BlobRef, _err error) {
	if (IsMock) {
		return _class.Create__mock(sessionID, mimeType, public)
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

func (_class BlobClass) SetPublic__mock(sessionID SessionRef, self BlobRef, value bool) (_err error) {
	log.Println("Blob.SetPublic not mocked")
	_err = errors.New("Blob.SetPublic not mocked")
	return
}
// Set the public field of the given blob.
func (_class BlobClass) SetPublic(sessionID SessionRef, self BlobRef, value bool) (_err error) {
	if (IsMock) {
		return _class.SetPublic__mock(sessionID, self, value)
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

func (_class BlobClass) SetNameDescription__mock(sessionID SessionRef, self BlobRef, value string) (_err error) {
	log.Println("Blob.SetNameDescription not mocked")
	_err = errors.New("Blob.SetNameDescription not mocked")
	return
}
// Set the name/description field of the given blob.
func (_class BlobClass) SetNameDescription(sessionID SessionRef, self BlobRef, value string) (_err error) {
	if (IsMock) {
		return _class.SetNameDescription__mock(sessionID, self, value)
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

func (_class BlobClass) SetNameLabel__mock(sessionID SessionRef, self BlobRef, value string) (_err error) {
	log.Println("Blob.SetNameLabel not mocked")
	_err = errors.New("Blob.SetNameLabel not mocked")
	return
}
// Set the name/label field of the given blob.
func (_class BlobClass) SetNameLabel(sessionID SessionRef, self BlobRef, value string) (_err error) {
	if (IsMock) {
		return _class.SetNameLabel__mock(sessionID, self, value)
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

func (_class BlobClass) GetMimeType__mock(sessionID SessionRef, self BlobRef) (_retval string, _err error) {
	log.Println("Blob.GetMimeType not mocked")
	_err = errors.New("Blob.GetMimeType not mocked")
	return
}
// Get the mime_type field of the given blob.
func (_class BlobClass) GetMimeType(sessionID SessionRef, self BlobRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetMimeType__mock(sessionID, self)
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

func (_class BlobClass) GetLastUpdated__mock(sessionID SessionRef, self BlobRef) (_retval time.Time, _err error) {
	log.Println("Blob.GetLastUpdated not mocked")
	_err = errors.New("Blob.GetLastUpdated not mocked")
	return
}
// Get the last_updated field of the given blob.
func (_class BlobClass) GetLastUpdated(sessionID SessionRef, self BlobRef) (_retval time.Time, _err error) {
	if (IsMock) {
		return _class.GetLastUpdated__mock(sessionID, self)
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

func (_class BlobClass) GetPublic__mock(sessionID SessionRef, self BlobRef) (_retval bool, _err error) {
	log.Println("Blob.GetPublic not mocked")
	_err = errors.New("Blob.GetPublic not mocked")
	return
}
// Get the public field of the given blob.
func (_class BlobClass) GetPublic(sessionID SessionRef, self BlobRef) (_retval bool, _err error) {
	if (IsMock) {
		return _class.GetPublic__mock(sessionID, self)
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

func (_class BlobClass) GetSize__mock(sessionID SessionRef, self BlobRef) (_retval int, _err error) {
	log.Println("Blob.GetSize not mocked")
	_err = errors.New("Blob.GetSize not mocked")
	return
}
// Get the size field of the given blob.
func (_class BlobClass) GetSize(sessionID SessionRef, self BlobRef) (_retval int, _err error) {
	if (IsMock) {
		return _class.GetSize__mock(sessionID, self)
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

func (_class BlobClass) GetNameDescription__mock(sessionID SessionRef, self BlobRef) (_retval string, _err error) {
	log.Println("Blob.GetNameDescription not mocked")
	_err = errors.New("Blob.GetNameDescription not mocked")
	return
}
// Get the name/description field of the given blob.
func (_class BlobClass) GetNameDescription(sessionID SessionRef, self BlobRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetNameDescription__mock(sessionID, self)
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

func (_class BlobClass) GetNameLabel__mock(sessionID SessionRef, self BlobRef) (_retval string, _err error) {
	log.Println("Blob.GetNameLabel not mocked")
	_err = errors.New("Blob.GetNameLabel not mocked")
	return
}
// Get the name/label field of the given blob.
func (_class BlobClass) GetNameLabel(sessionID SessionRef, self BlobRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetNameLabel__mock(sessionID, self)
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

func (_class BlobClass) GetUUID__mock(sessionID SessionRef, self BlobRef) (_retval string, _err error) {
	log.Println("Blob.GetUUID not mocked")
	_err = errors.New("Blob.GetUUID not mocked")
	return
}
// Get the uuid field of the given blob.
func (_class BlobClass) GetUUID(sessionID SessionRef, self BlobRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetUUID__mock(sessionID, self)
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

func (_class BlobClass) GetByNameLabel__mock(sessionID SessionRef, label string) (_retval []BlobRef, _err error) {
	log.Println("Blob.GetByNameLabel not mocked")
	_err = errors.New("Blob.GetByNameLabel not mocked")
	return
}
// Get all the blob instances with the given label.
func (_class BlobClass) GetByNameLabel(sessionID SessionRef, label string) (_retval []BlobRef, _err error) {
	if (IsMock) {
		return _class.GetByNameLabel__mock(sessionID, label)
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

func (_class BlobClass) GetByUUID__mock(sessionID SessionRef, uuid string) (_retval BlobRef, _err error) {
	log.Println("Blob.GetByUUID not mocked")
	_err = errors.New("Blob.GetByUUID not mocked")
	return
}
// Get a reference to the blob instance with the specified UUID.
func (_class BlobClass) GetByUUID(sessionID SessionRef, uuid string) (_retval BlobRef, _err error) {
	if (IsMock) {
		return _class.GetByUUID__mock(sessionID, uuid)
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

func (_class BlobClass) GetRecord__mock(sessionID SessionRef, self BlobRef) (_retval BlobRecord, _err error) {
	log.Println("Blob.GetRecord not mocked")
	_err = errors.New("Blob.GetRecord not mocked")
	return
}
// Get a record containing the current state of the given blob.
func (_class BlobClass) GetRecord(sessionID SessionRef, self BlobRef) (_retval BlobRecord, _err error) {
	if (IsMock) {
		return _class.GetRecord__mock(sessionID, self)
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
