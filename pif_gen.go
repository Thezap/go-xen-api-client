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

type PifIgmpStatus string

const (
  // IGMP Snooping is enabled in the corresponding backend bridge.'
	PifIgmpStatusEnabled PifIgmpStatus = "enabled"
  // IGMP Snooping is disabled in the corresponding backend bridge.'
	PifIgmpStatusDisabled PifIgmpStatus = "disabled"
  // IGMP snooping status is unknown. If this is a VLAN master, then please consult the underlying VLAN slave PIF.
	PifIgmpStatusUnknown PifIgmpStatus = "unknown"
)

type IPConfigurationMode string

const (
  // Do not acquire an IP address
	IPConfigurationModeNone IPConfigurationMode = "None"
  // Acquire an IP address by DHCP
	IPConfigurationModeDHCP IPConfigurationMode = "DHCP"
  // Static IP address configuration
	IPConfigurationModeStatic IPConfigurationMode = "Static"
)

type Ipv6ConfigurationMode string

const (
  // Do not acquire an IPv6 address
	Ipv6ConfigurationModeNone Ipv6ConfigurationMode = "None"
  // Acquire an IPv6 address by DHCP
	Ipv6ConfigurationModeDHCP Ipv6ConfigurationMode = "DHCP"
  // Static IPv6 address configuration
	Ipv6ConfigurationModeStatic Ipv6ConfigurationMode = "Static"
  // Router assigned prefix delegation IPv6 allocation
	Ipv6ConfigurationModeAutoconf Ipv6ConfigurationMode = "Autoconf"
)

type PrimaryAddressType string

const (
  // Primary address is the IPv4 address
	PrimaryAddressTypeIPv4 PrimaryAddressType = "IPv4"
  // Primary address is the IPv6 address
	PrimaryAddressTypeIPv6 PrimaryAddressType = "IPv6"
)

type PIFRecord struct {
  // Unique identifier/object reference
	UUID string
  // machine-readable name of the interface (e.g. eth0)
	Device string
  // virtual network to which this pif is connected
	Network NetworkRef
  // physical machine to which this pif is connected
	Host HostRef
  // ethernet MAC address of physical interface
	MAC string
  // MTU in octets
	MTU int
  // VLAN tag for all traffic passing through this interface
	VLAN int
  // metrics associated with this PIF
	Metrics PIFMetricsRef
  // true if this represents a physical network interface
	Physical bool
  // true if this interface is online
	CurrentlyAttached bool
  // Sets if and how this interface gets an IP address
	IPConfigurationMode IPConfigurationMode
  // IP address
	IP string
  // IP netmask
	Netmask string
  // IP gateway
	Gateway string
  // IP address of DNS servers to use
	DNS string
  // Indicates which bond this interface is part of
	BondSlaveOf BondRef
  // Indicates this PIF represents the results of a bond
	BondMasterOf []BondRef
  // Indicates wich VLAN this interface receives untagged traffic from
	VLANMasterOf VLANRef
  // Indicates which VLANs this interface transmits tagged traffic to
	VLANSlaveOf []VLANRef
  // Indicates whether the control software is listening for connections on this interface
	Management bool
  // Additional configuration
	OtherConfig map[string]string
  // Prevent this PIF from being unplugged; set this to notify the management tool-stack that the PIF has a special use and should not be unplugged under any circumstances (e.g. because you're running storage traffic over it)
	DisallowUnplug bool
  // Indicates to which tunnel this PIF gives access
	TunnelAccessPIFOf []TunnelRef
  // Indicates to which tunnel this PIF provides transport
	TunnelTransportPIFOf []TunnelRef
  // Sets if and how this interface gets an IPv6 address
	Ipv6ConfigurationMode Ipv6ConfigurationMode
  // IPv6 address
	IPv6 []string
  // IPv6 gateway
	Ipv6Gateway string
  // Which protocol should define the primary address of this interface
	PrimaryAddressType PrimaryAddressType
  // Indicates whether the interface is managed by xapi. If it is not, then xapi will not configure the interface, the commands PIF.plug/unplug/reconfigure_ip(v6) can not be used, nor can the interface be bonded or have VLANs based on top through xapi.
	Managed bool
  // Additional configuration properties for the interface.
	Properties map[string]string
  // Additional capabilities on the interface.
	Capabilities []string
  // The IGMP snooping status of the corresponding network bridge
	IgmpSnoopingStatus PifIgmpStatus
}

type PIFRef string

// A physical network interface (note separate VLANs are represented as several PIFs)
type PIFClass struct {
	client *Client
}


var PIFClass_GetAllRecordsMockedCallback = func (sessionID SessionRef) (_retval map[PIFRef]PIFRecord, _err error) {
	log.Println("PIF.GetAllRecords not mocked")
	_err = errors.New("PIF.GetAllRecords not mocked")
	return
}

func (_class PIFClass) GetAllRecordsMock(sessionID SessionRef) (_retval map[PIFRef]PIFRecord, _err error) {
	return PIFClass_GetAllRecordsMockedCallback(sessionID)
}
// Return a map of PIF references to PIF records for all PIFs known to the system.
func (_class PIFClass) GetAllRecords(sessionID SessionRef) (_retval map[PIFRef]PIFRecord, _err error) {
	if (IsMock) {
		return _class.GetAllRecordsMock(sessionID)
	}	
	_method := "PIF.get_all_records"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertPIFRefToPIFRecordMapToGo(_method + " -> ", _result.Value)
	return
}


var PIFClass_GetAllMockedCallback = func (sessionID SessionRef) (_retval []PIFRef, _err error) {
	log.Println("PIF.GetAll not mocked")
	_err = errors.New("PIF.GetAll not mocked")
	return
}

func (_class PIFClass) GetAllMock(sessionID SessionRef) (_retval []PIFRef, _err error) {
	return PIFClass_GetAllMockedCallback(sessionID)
}
// Return a list of all the PIFs known to the system.
func (_class PIFClass) GetAll(sessionID SessionRef) (_retval []PIFRef, _err error) {
	if (IsMock) {
		return _class.GetAllMock(sessionID)
	}	
	_method := "PIF.get_all"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertPIFRefSetToGo(_method + " -> ", _result.Value)
	return
}


var PIFClass_SetPropertyMockedCallback = func (sessionID SessionRef, self PIFRef, name string, value string) (_err error) {
	log.Println("PIF.SetProperty not mocked")
	_err = errors.New("PIF.SetProperty not mocked")
	return
}

func (_class PIFClass) SetPropertyMock(sessionID SessionRef, self PIFRef, name string, value string) (_err error) {
	return PIFClass_SetPropertyMockedCallback(sessionID, self, name, value)
}
// Set the value of a property of the PIF
func (_class PIFClass) SetProperty(sessionID SessionRef, self PIFRef, name string, value string) (_err error) {
	if (IsMock) {
		return _class.SetPropertyMock(sessionID, self, name, value)
	}	
	_method := "PIF.set_property"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPIFRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_nameArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "name"), name)
	if _err != nil {
		return
	}
	_valueArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _nameArg, _valueArg)
	return
}


var PIFClass_DbForgetMockedCallback = func (sessionID SessionRef, self PIFRef) (_err error) {
	log.Println("PIF.DbForget not mocked")
	_err = errors.New("PIF.DbForget not mocked")
	return
}

func (_class PIFClass) DbForgetMock(sessionID SessionRef, self PIFRef) (_err error) {
	return PIFClass_DbForgetMockedCallback(sessionID, self)
}
// Destroy a PIF database record.
func (_class PIFClass) DbForget(sessionID SessionRef, self PIFRef) (_err error) {
	if (IsMock) {
		return _class.DbForgetMock(sessionID, self)
	}	
	_method := "PIF.db_forget"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPIFRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg)
	return
}


var PIFClass_DbIntroduceMockedCallback = func (sessionID SessionRef, device string, network NetworkRef, host HostRef, mac string, mtu int, vlan int, physical bool, ipConfigurationMode IPConfigurationMode, ip string, netmask string, gateway string, dns string, bondSlaveOf BondRef, vlanMasterOf VLANRef, management bool, otherConfig map[string]string, disallowUnplug bool, ipv6ConfigurationMode Ipv6ConfigurationMode, ipv6 []string, ipv6Gateway string, primaryAddressType PrimaryAddressType, managed bool, properties map[string]string) (_retval PIFRef, _err error) {
	log.Println("PIF.DbIntroduce not mocked")
	_err = errors.New("PIF.DbIntroduce not mocked")
	return
}

func (_class PIFClass) DbIntroduceMock(sessionID SessionRef, device string, network NetworkRef, host HostRef, mac string, mtu int, vlan int, physical bool, ipConfigurationMode IPConfigurationMode, ip string, netmask string, gateway string, dns string, bondSlaveOf BondRef, vlanMasterOf VLANRef, management bool, otherConfig map[string]string, disallowUnplug bool, ipv6ConfigurationMode Ipv6ConfigurationMode, ipv6 []string, ipv6Gateway string, primaryAddressType PrimaryAddressType, managed bool, properties map[string]string) (_retval PIFRef, _err error) {
	return PIFClass_DbIntroduceMockedCallback(sessionID, device, network, host, mac, mtu, vlan, physical, ipConfigurationMode, ip, netmask, gateway, dns, bondSlaveOf, vlanMasterOf, management, otherConfig, disallowUnplug, ipv6ConfigurationMode, ipv6, ipv6Gateway, primaryAddressType, managed, properties)
}
// Create a new PIF record in the database only
func (_class PIFClass) DbIntroduce(sessionID SessionRef, device string, network NetworkRef, host HostRef, mac string, mtu int, vlan int, physical bool, ipConfigurationMode IPConfigurationMode, ip string, netmask string, gateway string, dns string, bondSlaveOf BondRef, vlanMasterOf VLANRef, management bool, otherConfig map[string]string, disallowUnplug bool, ipv6ConfigurationMode Ipv6ConfigurationMode, ipv6 []string, ipv6Gateway string, primaryAddressType PrimaryAddressType, managed bool, properties map[string]string) (_retval PIFRef, _err error) {
	if (IsMock) {
		return _class.DbIntroduceMock(sessionID, device, network, host, mac, mtu, vlan, physical, ipConfigurationMode, ip, netmask, gateway, dns, bondSlaveOf, vlanMasterOf, management, otherConfig, disallowUnplug, ipv6ConfigurationMode, ipv6, ipv6Gateway, primaryAddressType, managed, properties)
	}	
	_method := "PIF.db_introduce"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_deviceArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "device"), device)
	if _err != nil {
		return
	}
	_networkArg, _err := convertNetworkRefToXen(fmt.Sprintf("%s(%s)", _method, "network"), network)
	if _err != nil {
		return
	}
	_hostArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "host"), host)
	if _err != nil {
		return
	}
	_macArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "MAC"), mac)
	if _err != nil {
		return
	}
	_mtuArg, _err := convertIntToXen(fmt.Sprintf("%s(%s)", _method, "MTU"), mtu)
	if _err != nil {
		return
	}
	_vlanArg, _err := convertIntToXen(fmt.Sprintf("%s(%s)", _method, "VLAN"), vlan)
	if _err != nil {
		return
	}
	_physicalArg, _err := convertBoolToXen(fmt.Sprintf("%s(%s)", _method, "physical"), physical)
	if _err != nil {
		return
	}
	_ipConfigurationModeArg, _err := convertEnumIPConfigurationModeToXen(fmt.Sprintf("%s(%s)", _method, "ip_configuration_mode"), ipConfigurationMode)
	if _err != nil {
		return
	}
	_ipArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "IP"), ip)
	if _err != nil {
		return
	}
	_netmaskArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "netmask"), netmask)
	if _err != nil {
		return
	}
	_gatewayArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "gateway"), gateway)
	if _err != nil {
		return
	}
	_dnsArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "DNS"), dns)
	if _err != nil {
		return
	}
	_bondSlaveOfArg, _err := convertBondRefToXen(fmt.Sprintf("%s(%s)", _method, "bond_slave_of"), bondSlaveOf)
	if _err != nil {
		return
	}
	_vlanMasterOfArg, _err := convertVLANRefToXen(fmt.Sprintf("%s(%s)", _method, "VLAN_master_of"), vlanMasterOf)
	if _err != nil {
		return
	}
	_managementArg, _err := convertBoolToXen(fmt.Sprintf("%s(%s)", _method, "management"), management)
	if _err != nil {
		return
	}
	_otherConfigArg, _err := convertStringToStringMapToXen(fmt.Sprintf("%s(%s)", _method, "other_config"), otherConfig)
	if _err != nil {
		return
	}
	_disallowUnplugArg, _err := convertBoolToXen(fmt.Sprintf("%s(%s)", _method, "disallow_unplug"), disallowUnplug)
	if _err != nil {
		return
	}
	_ipv6ConfigurationModeArg, _err := convertEnumIpv6ConfigurationModeToXen(fmt.Sprintf("%s(%s)", _method, "ipv6_configuration_mode"), ipv6ConfigurationMode)
	if _err != nil {
		return
	}
	_ipv6Arg, _err := convertStringSetToXen(fmt.Sprintf("%s(%s)", _method, "IPv6"), ipv6)
	if _err != nil {
		return
	}
	_ipv6GatewayArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "ipv6_gateway"), ipv6Gateway)
	if _err != nil {
		return
	}
	_primaryAddressTypeArg, _err := convertEnumPrimaryAddressTypeToXen(fmt.Sprintf("%s(%s)", _method, "primary_address_type"), primaryAddressType)
	if _err != nil {
		return
	}
	_managedArg, _err := convertBoolToXen(fmt.Sprintf("%s(%s)", _method, "managed"), managed)
	if _err != nil {
		return
	}
	_propertiesArg, _err := convertStringToStringMapToXen(fmt.Sprintf("%s(%s)", _method, "properties"), properties)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _deviceArg, _networkArg, _hostArg, _macArg, _mtuArg, _vlanArg, _physicalArg, _ipConfigurationModeArg, _ipArg, _netmaskArg, _gatewayArg, _dnsArg, _bondSlaveOfArg, _vlanMasterOfArg, _managementArg, _otherConfigArg, _disallowUnplugArg, _ipv6ConfigurationModeArg, _ipv6Arg, _ipv6GatewayArg, _primaryAddressTypeArg, _managedArg, _propertiesArg)
	if _err != nil {
		return
	}
	_retval, _err = convertPIFRefToGo(_method + " -> ", _result.Value)
	return
}


var PIFClass_PlugMockedCallback = func (sessionID SessionRef, self PIFRef) (_err error) {
	log.Println("PIF.Plug not mocked")
	_err = errors.New("PIF.Plug not mocked")
	return
}

func (_class PIFClass) PlugMock(sessionID SessionRef, self PIFRef) (_err error) {
	return PIFClass_PlugMockedCallback(sessionID, self)
}
// Attempt to bring up a physical interface
//
// Errors:
//  TRANSPORT_PIF_NOT_CONFIGURED - The tunnel transport PIF has no IP configuration set.
func (_class PIFClass) Plug(sessionID SessionRef, self PIFRef) (_err error) {
	if (IsMock) {
		return _class.PlugMock(sessionID, self)
	}	
	_method := "PIF.plug"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPIFRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg)
	return
}


var PIFClass_UnplugMockedCallback = func (sessionID SessionRef, self PIFRef) (_err error) {
	log.Println("PIF.Unplug not mocked")
	_err = errors.New("PIF.Unplug not mocked")
	return
}

func (_class PIFClass) UnplugMock(sessionID SessionRef, self PIFRef) (_err error) {
	return PIFClass_UnplugMockedCallback(sessionID, self)
}
// Attempt to bring down a physical interface
//
// Errors:
//  HA_OPERATION_WOULD_BREAK_FAILOVER_PLAN - This operation cannot be performed because it would invalidate VM failover planning such that the system would be unable to guarantee to restart protected VMs after a Host failure.
//  VIF_IN_USE - Network has active VIFs
//  PIF_DOES_NOT_ALLOW_UNPLUG - The operation you requested cannot be performed because the specified PIF does not allow unplug.
//  PIF_HAS_FCOE_SR_IN_USE - The operation you requested cannot be performed because the specified PIF has FCoE SR in use.
func (_class PIFClass) Unplug(sessionID SessionRef, self PIFRef) (_err error) {
	if (IsMock) {
		return _class.UnplugMock(sessionID, self)
	}	
	_method := "PIF.unplug"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPIFRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg)
	return
}


var PIFClass_ForgetMockedCallback = func (sessionID SessionRef, self PIFRef) (_err error) {
	log.Println("PIF.Forget not mocked")
	_err = errors.New("PIF.Forget not mocked")
	return
}

func (_class PIFClass) ForgetMock(sessionID SessionRef, self PIFRef) (_err error) {
	return PIFClass_ForgetMockedCallback(sessionID, self)
}
// Destroy the PIF object matching a particular network interface
//
// Errors:
//  PIF_TUNNEL_STILL_EXISTS - Operation cannot proceed while a tunnel exists on this interface.
func (_class PIFClass) Forget(sessionID SessionRef, self PIFRef) (_err error) {
	if (IsMock) {
		return _class.ForgetMock(sessionID, self)
	}	
	_method := "PIF.forget"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPIFRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg)
	return
}


var PIFClass_IntroduceMockedCallback = func (sessionID SessionRef, host HostRef, mac string, device string, managed bool) (_retval PIFRef, _err error) {
	log.Println("PIF.Introduce not mocked")
	_err = errors.New("PIF.Introduce not mocked")
	return
}

func (_class PIFClass) IntroduceMock(sessionID SessionRef, host HostRef, mac string, device string, managed bool) (_retval PIFRef, _err error) {
	return PIFClass_IntroduceMockedCallback(sessionID, host, mac, device, managed)
}
// Create a PIF object matching a particular network interface
func (_class PIFClass) Introduce(sessionID SessionRef, host HostRef, mac string, device string, managed bool) (_retval PIFRef, _err error) {
	if (IsMock) {
		return _class.IntroduceMock(sessionID, host, mac, device, managed)
	}	
	_method := "PIF.introduce"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_hostArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "host"), host)
	if _err != nil {
		return
	}
	_macArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "MAC"), mac)
	if _err != nil {
		return
	}
	_deviceArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "device"), device)
	if _err != nil {
		return
	}
	_managedArg, _err := convertBoolToXen(fmt.Sprintf("%s(%s)", _method, "managed"), managed)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _hostArg, _macArg, _deviceArg, _managedArg)
	if _err != nil {
		return
	}
	_retval, _err = convertPIFRefToGo(_method + " -> ", _result.Value)
	return
}


var PIFClass_ScanMockedCallback = func (sessionID SessionRef, host HostRef) (_err error) {
	log.Println("PIF.Scan not mocked")
	_err = errors.New("PIF.Scan not mocked")
	return
}

func (_class PIFClass) ScanMock(sessionID SessionRef, host HostRef) (_err error) {
	return PIFClass_ScanMockedCallback(sessionID, host)
}
// Scan for physical interfaces on a host and create PIF objects to represent them
func (_class PIFClass) Scan(sessionID SessionRef, host HostRef) (_err error) {
	if (IsMock) {
		return _class.ScanMock(sessionID, host)
	}	
	_method := "PIF.scan"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_hostArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "host"), host)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _hostArg)
	return
}


var PIFClass_SetPrimaryAddressTypeMockedCallback = func (sessionID SessionRef, self PIFRef, primaryAddressType PrimaryAddressType) (_err error) {
	log.Println("PIF.SetPrimaryAddressType not mocked")
	_err = errors.New("PIF.SetPrimaryAddressType not mocked")
	return
}

func (_class PIFClass) SetPrimaryAddressTypeMock(sessionID SessionRef, self PIFRef, primaryAddressType PrimaryAddressType) (_err error) {
	return PIFClass_SetPrimaryAddressTypeMockedCallback(sessionID, self, primaryAddressType)
}
// Change the primary address type used by this PIF
func (_class PIFClass) SetPrimaryAddressType(sessionID SessionRef, self PIFRef, primaryAddressType PrimaryAddressType) (_err error) {
	if (IsMock) {
		return _class.SetPrimaryAddressTypeMock(sessionID, self, primaryAddressType)
	}	
	_method := "PIF.set_primary_address_type"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPIFRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_primaryAddressTypeArg, _err := convertEnumPrimaryAddressTypeToXen(fmt.Sprintf("%s(%s)", _method, "primary_address_type"), primaryAddressType)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _primaryAddressTypeArg)
	return
}


var PIFClass_ReconfigureIpv6MockedCallback = func (sessionID SessionRef, self PIFRef, mode Ipv6ConfigurationMode, ipv6 string, gateway string, dns string) (_err error) {
	log.Println("PIF.ReconfigureIpv6 not mocked")
	_err = errors.New("PIF.ReconfigureIpv6 not mocked")
	return
}

func (_class PIFClass) ReconfigureIpv6Mock(sessionID SessionRef, self PIFRef, mode Ipv6ConfigurationMode, ipv6 string, gateway string, dns string) (_err error) {
	return PIFClass_ReconfigureIpv6MockedCallback(sessionID, self, mode, ipv6, gateway, dns)
}
// Reconfigure the IPv6 address settings for this interface
func (_class PIFClass) ReconfigureIpv6(sessionID SessionRef, self PIFRef, mode Ipv6ConfigurationMode, ipv6 string, gateway string, dns string) (_err error) {
	if (IsMock) {
		return _class.ReconfigureIpv6Mock(sessionID, self, mode, ipv6, gateway, dns)
	}	
	_method := "PIF.reconfigure_ipv6"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPIFRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_modeArg, _err := convertEnumIpv6ConfigurationModeToXen(fmt.Sprintf("%s(%s)", _method, "mode"), mode)
	if _err != nil {
		return
	}
	_ipv6Arg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "IPv6"), ipv6)
	if _err != nil {
		return
	}
	_gatewayArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "gateway"), gateway)
	if _err != nil {
		return
	}
	_dnsArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "DNS"), dns)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _modeArg, _ipv6Arg, _gatewayArg, _dnsArg)
	return
}


var PIFClass_ReconfigureIPMockedCallback = func (sessionID SessionRef, self PIFRef, mode IPConfigurationMode, ip string, netmask string, gateway string, dns string) (_err error) {
	log.Println("PIF.ReconfigureIP not mocked")
	_err = errors.New("PIF.ReconfigureIP not mocked")
	return
}

func (_class PIFClass) ReconfigureIPMock(sessionID SessionRef, self PIFRef, mode IPConfigurationMode, ip string, netmask string, gateway string, dns string) (_err error) {
	return PIFClass_ReconfigureIPMockedCallback(sessionID, self, mode, ip, netmask, gateway, dns)
}
// Reconfigure the IP address settings for this interface
func (_class PIFClass) ReconfigureIP(sessionID SessionRef, self PIFRef, mode IPConfigurationMode, ip string, netmask string, gateway string, dns string) (_err error) {
	if (IsMock) {
		return _class.ReconfigureIPMock(sessionID, self, mode, ip, netmask, gateway, dns)
	}	
	_method := "PIF.reconfigure_ip"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPIFRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_modeArg, _err := convertEnumIPConfigurationModeToXen(fmt.Sprintf("%s(%s)", _method, "mode"), mode)
	if _err != nil {
		return
	}
	_ipArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "IP"), ip)
	if _err != nil {
		return
	}
	_netmaskArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "netmask"), netmask)
	if _err != nil {
		return
	}
	_gatewayArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "gateway"), gateway)
	if _err != nil {
		return
	}
	_dnsArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "DNS"), dns)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _modeArg, _ipArg, _netmaskArg, _gatewayArg, _dnsArg)
	return
}


var PIFClass_DestroyMockedCallback = func (sessionID SessionRef, self PIFRef) (_err error) {
	log.Println("PIF.Destroy not mocked")
	_err = errors.New("PIF.Destroy not mocked")
	return
}

func (_class PIFClass) DestroyMock(sessionID SessionRef, self PIFRef) (_err error) {
	return PIFClass_DestroyMockedCallback(sessionID, self)
}
// Destroy the PIF object (provided it is a VLAN interface). This call is deprecated: use VLAN.destroy or Bond.destroy instead
//
// Errors:
//  PIF_IS_PHYSICAL - You tried to destroy a PIF, but it represents an aspect of the physical host configuration, and so cannot be destroyed.  The parameter echoes the PIF handle you gave.
func (_class PIFClass) Destroy(sessionID SessionRef, self PIFRef) (_err error) {
	if (IsMock) {
		return _class.DestroyMock(sessionID, self)
	}	
	_method := "PIF.destroy"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPIFRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg)
	return
}


var PIFClass_CreateVLANMockedCallback = func (sessionID SessionRef, device string, network NetworkRef, host HostRef, vlan int) (_retval PIFRef, _err error) {
	log.Println("PIF.CreateVLAN not mocked")
	_err = errors.New("PIF.CreateVLAN not mocked")
	return
}

func (_class PIFClass) CreateVLANMock(sessionID SessionRef, device string, network NetworkRef, host HostRef, vlan int) (_retval PIFRef, _err error) {
	return PIFClass_CreateVLANMockedCallback(sessionID, device, network, host, vlan)
}
// Create a VLAN interface from an existing physical interface. This call is deprecated: use VLAN.create instead
//
// Errors:
//  VLAN_TAG_INVALID - You tried to create a VLAN, but the tag you gave was invalid -- it must be between 0 and 4094.  The parameter echoes the VLAN tag you gave.
func (_class PIFClass) CreateVLAN(sessionID SessionRef, device string, network NetworkRef, host HostRef, vlan int) (_retval PIFRef, _err error) {
	if (IsMock) {
		return _class.CreateVLANMock(sessionID, device, network, host, vlan)
	}	
	_method := "PIF.create_VLAN"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_deviceArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "device"), device)
	if _err != nil {
		return
	}
	_networkArg, _err := convertNetworkRefToXen(fmt.Sprintf("%s(%s)", _method, "network"), network)
	if _err != nil {
		return
	}
	_hostArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "host"), host)
	if _err != nil {
		return
	}
	_vlanArg, _err := convertIntToXen(fmt.Sprintf("%s(%s)", _method, "VLAN"), vlan)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _deviceArg, _networkArg, _hostArg, _vlanArg)
	if _err != nil {
		return
	}
	_retval, _err = convertPIFRefToGo(_method + " -> ", _result.Value)
	return
}


var PIFClass_SetDisallowUnplugMockedCallback = func (sessionID SessionRef, self PIFRef, value bool) (_err error) {
	log.Println("PIF.SetDisallowUnplug not mocked")
	_err = errors.New("PIF.SetDisallowUnplug not mocked")
	return
}

func (_class PIFClass) SetDisallowUnplugMock(sessionID SessionRef, self PIFRef, value bool) (_err error) {
	return PIFClass_SetDisallowUnplugMockedCallback(sessionID, self, value)
}
// Set the disallow_unplug field of the given PIF.
func (_class PIFClass) SetDisallowUnplug(sessionID SessionRef, self PIFRef, value bool) (_err error) {
	if (IsMock) {
		return _class.SetDisallowUnplugMock(sessionID, self, value)
	}	
	_method := "PIF.set_disallow_unplug"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPIFRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var PIFClass_RemoveFromOtherConfigMockedCallback = func (sessionID SessionRef, self PIFRef, key string) (_err error) {
	log.Println("PIF.RemoveFromOtherConfig not mocked")
	_err = errors.New("PIF.RemoveFromOtherConfig not mocked")
	return
}

func (_class PIFClass) RemoveFromOtherConfigMock(sessionID SessionRef, self PIFRef, key string) (_err error) {
	return PIFClass_RemoveFromOtherConfigMockedCallback(sessionID, self, key)
}
// Remove the given key and its corresponding value from the other_config field of the given PIF.  If the key is not in that Map, then do nothing.
func (_class PIFClass) RemoveFromOtherConfig(sessionID SessionRef, self PIFRef, key string) (_err error) {
	if (IsMock) {
		return _class.RemoveFromOtherConfigMock(sessionID, self, key)
	}	
	_method := "PIF.remove_from_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPIFRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var PIFClass_AddToOtherConfigMockedCallback = func (sessionID SessionRef, self PIFRef, key string, value string) (_err error) {
	log.Println("PIF.AddToOtherConfig not mocked")
	_err = errors.New("PIF.AddToOtherConfig not mocked")
	return
}

func (_class PIFClass) AddToOtherConfigMock(sessionID SessionRef, self PIFRef, key string, value string) (_err error) {
	return PIFClass_AddToOtherConfigMockedCallback(sessionID, self, key, value)
}
// Add the given key-value pair to the other_config field of the given PIF.
func (_class PIFClass) AddToOtherConfig(sessionID SessionRef, self PIFRef, key string, value string) (_err error) {
	if (IsMock) {
		return _class.AddToOtherConfigMock(sessionID, self, key, value)
	}	
	_method := "PIF.add_to_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPIFRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var PIFClass_SetOtherConfigMockedCallback = func (sessionID SessionRef, self PIFRef, value map[string]string) (_err error) {
	log.Println("PIF.SetOtherConfig not mocked")
	_err = errors.New("PIF.SetOtherConfig not mocked")
	return
}

func (_class PIFClass) SetOtherConfigMock(sessionID SessionRef, self PIFRef, value map[string]string) (_err error) {
	return PIFClass_SetOtherConfigMockedCallback(sessionID, self, value)
}
// Set the other_config field of the given PIF.
func (_class PIFClass) SetOtherConfig(sessionID SessionRef, self PIFRef, value map[string]string) (_err error) {
	if (IsMock) {
		return _class.SetOtherConfigMock(sessionID, self, value)
	}	
	_method := "PIF.set_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPIFRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var PIFClass_GetIgmpSnoopingStatusMockedCallback = func (sessionID SessionRef, self PIFRef) (_retval PifIgmpStatus, _err error) {
	log.Println("PIF.GetIgmpSnoopingStatus not mocked")
	_err = errors.New("PIF.GetIgmpSnoopingStatus not mocked")
	return
}

func (_class PIFClass) GetIgmpSnoopingStatusMock(sessionID SessionRef, self PIFRef) (_retval PifIgmpStatus, _err error) {
	return PIFClass_GetIgmpSnoopingStatusMockedCallback(sessionID, self)
}
// Get the igmp_snooping_status field of the given PIF.
func (_class PIFClass) GetIgmpSnoopingStatus(sessionID SessionRef, self PIFRef) (_retval PifIgmpStatus, _err error) {
	if (IsMock) {
		return _class.GetIgmpSnoopingStatusMock(sessionID, self)
	}	
	_method := "PIF.get_igmp_snooping_status"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPIFRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertEnumPifIgmpStatusToGo(_method + " -> ", _result.Value)
	return
}


var PIFClass_GetCapabilitiesMockedCallback = func (sessionID SessionRef, self PIFRef) (_retval []string, _err error) {
	log.Println("PIF.GetCapabilities not mocked")
	_err = errors.New("PIF.GetCapabilities not mocked")
	return
}

func (_class PIFClass) GetCapabilitiesMock(sessionID SessionRef, self PIFRef) (_retval []string, _err error) {
	return PIFClass_GetCapabilitiesMockedCallback(sessionID, self)
}
// Get the capabilities field of the given PIF.
func (_class PIFClass) GetCapabilities(sessionID SessionRef, self PIFRef) (_retval []string, _err error) {
	if (IsMock) {
		return _class.GetCapabilitiesMock(sessionID, self)
	}	
	_method := "PIF.get_capabilities"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPIFRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var PIFClass_GetPropertiesMockedCallback = func (sessionID SessionRef, self PIFRef) (_retval map[string]string, _err error) {
	log.Println("PIF.GetProperties not mocked")
	_err = errors.New("PIF.GetProperties not mocked")
	return
}

func (_class PIFClass) GetPropertiesMock(sessionID SessionRef, self PIFRef) (_retval map[string]string, _err error) {
	return PIFClass_GetPropertiesMockedCallback(sessionID, self)
}
// Get the properties field of the given PIF.
func (_class PIFClass) GetProperties(sessionID SessionRef, self PIFRef) (_retval map[string]string, _err error) {
	if (IsMock) {
		return _class.GetPropertiesMock(sessionID, self)
	}	
	_method := "PIF.get_properties"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPIFRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var PIFClass_GetManagedMockedCallback = func (sessionID SessionRef, self PIFRef) (_retval bool, _err error) {
	log.Println("PIF.GetManaged not mocked")
	_err = errors.New("PIF.GetManaged not mocked")
	return
}

func (_class PIFClass) GetManagedMock(sessionID SessionRef, self PIFRef) (_retval bool, _err error) {
	return PIFClass_GetManagedMockedCallback(sessionID, self)
}
// Get the managed field of the given PIF.
func (_class PIFClass) GetManaged(sessionID SessionRef, self PIFRef) (_retval bool, _err error) {
	if (IsMock) {
		return _class.GetManagedMock(sessionID, self)
	}	
	_method := "PIF.get_managed"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPIFRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var PIFClass_GetPrimaryAddressTypeMockedCallback = func (sessionID SessionRef, self PIFRef) (_retval PrimaryAddressType, _err error) {
	log.Println("PIF.GetPrimaryAddressType not mocked")
	_err = errors.New("PIF.GetPrimaryAddressType not mocked")
	return
}

func (_class PIFClass) GetPrimaryAddressTypeMock(sessionID SessionRef, self PIFRef) (_retval PrimaryAddressType, _err error) {
	return PIFClass_GetPrimaryAddressTypeMockedCallback(sessionID, self)
}
// Get the primary_address_type field of the given PIF.
func (_class PIFClass) GetPrimaryAddressType(sessionID SessionRef, self PIFRef) (_retval PrimaryAddressType, _err error) {
	if (IsMock) {
		return _class.GetPrimaryAddressTypeMock(sessionID, self)
	}	
	_method := "PIF.get_primary_address_type"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPIFRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertEnumPrimaryAddressTypeToGo(_method + " -> ", _result.Value)
	return
}


var PIFClass_GetIpv6GatewayMockedCallback = func (sessionID SessionRef, self PIFRef) (_retval string, _err error) {
	log.Println("PIF.GetIpv6Gateway not mocked")
	_err = errors.New("PIF.GetIpv6Gateway not mocked")
	return
}

func (_class PIFClass) GetIpv6GatewayMock(sessionID SessionRef, self PIFRef) (_retval string, _err error) {
	return PIFClass_GetIpv6GatewayMockedCallback(sessionID, self)
}
// Get the ipv6_gateway field of the given PIF.
func (_class PIFClass) GetIpv6Gateway(sessionID SessionRef, self PIFRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetIpv6GatewayMock(sessionID, self)
	}	
	_method := "PIF.get_ipv6_gateway"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPIFRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var PIFClass_GetIPv6MockedCallback = func (sessionID SessionRef, self PIFRef) (_retval []string, _err error) {
	log.Println("PIF.GetIPv6 not mocked")
	_err = errors.New("PIF.GetIPv6 not mocked")
	return
}

func (_class PIFClass) GetIPv6Mock(sessionID SessionRef, self PIFRef) (_retval []string, _err error) {
	return PIFClass_GetIPv6MockedCallback(sessionID, self)
}
// Get the IPv6 field of the given PIF.
func (_class PIFClass) GetIPv6(sessionID SessionRef, self PIFRef) (_retval []string, _err error) {
	if (IsMock) {
		return _class.GetIPv6Mock(sessionID, self)
	}	
	_method := "PIF.get_IPv6"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPIFRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var PIFClass_GetIpv6ConfigurationModeMockedCallback = func (sessionID SessionRef, self PIFRef) (_retval Ipv6ConfigurationMode, _err error) {
	log.Println("PIF.GetIpv6ConfigurationMode not mocked")
	_err = errors.New("PIF.GetIpv6ConfigurationMode not mocked")
	return
}

func (_class PIFClass) GetIpv6ConfigurationModeMock(sessionID SessionRef, self PIFRef) (_retval Ipv6ConfigurationMode, _err error) {
	return PIFClass_GetIpv6ConfigurationModeMockedCallback(sessionID, self)
}
// Get the ipv6_configuration_mode field of the given PIF.
func (_class PIFClass) GetIpv6ConfigurationMode(sessionID SessionRef, self PIFRef) (_retval Ipv6ConfigurationMode, _err error) {
	if (IsMock) {
		return _class.GetIpv6ConfigurationModeMock(sessionID, self)
	}	
	_method := "PIF.get_ipv6_configuration_mode"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPIFRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertEnumIpv6ConfigurationModeToGo(_method + " -> ", _result.Value)
	return
}


var PIFClass_GetTunnelTransportPIFOfMockedCallback = func (sessionID SessionRef, self PIFRef) (_retval []TunnelRef, _err error) {
	log.Println("PIF.GetTunnelTransportPIFOf not mocked")
	_err = errors.New("PIF.GetTunnelTransportPIFOf not mocked")
	return
}

func (_class PIFClass) GetTunnelTransportPIFOfMock(sessionID SessionRef, self PIFRef) (_retval []TunnelRef, _err error) {
	return PIFClass_GetTunnelTransportPIFOfMockedCallback(sessionID, self)
}
// Get the tunnel_transport_PIF_of field of the given PIF.
func (_class PIFClass) GetTunnelTransportPIFOf(sessionID SessionRef, self PIFRef) (_retval []TunnelRef, _err error) {
	if (IsMock) {
		return _class.GetTunnelTransportPIFOfMock(sessionID, self)
	}	
	_method := "PIF.get_tunnel_transport_PIF_of"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPIFRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTunnelRefSetToGo(_method + " -> ", _result.Value)
	return
}


var PIFClass_GetTunnelAccessPIFOfMockedCallback = func (sessionID SessionRef, self PIFRef) (_retval []TunnelRef, _err error) {
	log.Println("PIF.GetTunnelAccessPIFOf not mocked")
	_err = errors.New("PIF.GetTunnelAccessPIFOf not mocked")
	return
}

func (_class PIFClass) GetTunnelAccessPIFOfMock(sessionID SessionRef, self PIFRef) (_retval []TunnelRef, _err error) {
	return PIFClass_GetTunnelAccessPIFOfMockedCallback(sessionID, self)
}
// Get the tunnel_access_PIF_of field of the given PIF.
func (_class PIFClass) GetTunnelAccessPIFOf(sessionID SessionRef, self PIFRef) (_retval []TunnelRef, _err error) {
	if (IsMock) {
		return _class.GetTunnelAccessPIFOfMock(sessionID, self)
	}	
	_method := "PIF.get_tunnel_access_PIF_of"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPIFRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTunnelRefSetToGo(_method + " -> ", _result.Value)
	return
}


var PIFClass_GetDisallowUnplugMockedCallback = func (sessionID SessionRef, self PIFRef) (_retval bool, _err error) {
	log.Println("PIF.GetDisallowUnplug not mocked")
	_err = errors.New("PIF.GetDisallowUnplug not mocked")
	return
}

func (_class PIFClass) GetDisallowUnplugMock(sessionID SessionRef, self PIFRef) (_retval bool, _err error) {
	return PIFClass_GetDisallowUnplugMockedCallback(sessionID, self)
}
// Get the disallow_unplug field of the given PIF.
func (_class PIFClass) GetDisallowUnplug(sessionID SessionRef, self PIFRef) (_retval bool, _err error) {
	if (IsMock) {
		return _class.GetDisallowUnplugMock(sessionID, self)
	}	
	_method := "PIF.get_disallow_unplug"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPIFRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var PIFClass_GetOtherConfigMockedCallback = func (sessionID SessionRef, self PIFRef) (_retval map[string]string, _err error) {
	log.Println("PIF.GetOtherConfig not mocked")
	_err = errors.New("PIF.GetOtherConfig not mocked")
	return
}

func (_class PIFClass) GetOtherConfigMock(sessionID SessionRef, self PIFRef) (_retval map[string]string, _err error) {
	return PIFClass_GetOtherConfigMockedCallback(sessionID, self)
}
// Get the other_config field of the given PIF.
func (_class PIFClass) GetOtherConfig(sessionID SessionRef, self PIFRef) (_retval map[string]string, _err error) {
	if (IsMock) {
		return _class.GetOtherConfigMock(sessionID, self)
	}	
	_method := "PIF.get_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPIFRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var PIFClass_GetManagementMockedCallback = func (sessionID SessionRef, self PIFRef) (_retval bool, _err error) {
	log.Println("PIF.GetManagement not mocked")
	_err = errors.New("PIF.GetManagement not mocked")
	return
}

func (_class PIFClass) GetManagementMock(sessionID SessionRef, self PIFRef) (_retval bool, _err error) {
	return PIFClass_GetManagementMockedCallback(sessionID, self)
}
// Get the management field of the given PIF.
func (_class PIFClass) GetManagement(sessionID SessionRef, self PIFRef) (_retval bool, _err error) {
	if (IsMock) {
		return _class.GetManagementMock(sessionID, self)
	}	
	_method := "PIF.get_management"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPIFRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var PIFClass_GetVLANSlaveOfMockedCallback = func (sessionID SessionRef, self PIFRef) (_retval []VLANRef, _err error) {
	log.Println("PIF.GetVLANSlaveOf not mocked")
	_err = errors.New("PIF.GetVLANSlaveOf not mocked")
	return
}

func (_class PIFClass) GetVLANSlaveOfMock(sessionID SessionRef, self PIFRef) (_retval []VLANRef, _err error) {
	return PIFClass_GetVLANSlaveOfMockedCallback(sessionID, self)
}
// Get the VLAN_slave_of field of the given PIF.
func (_class PIFClass) GetVLANSlaveOf(sessionID SessionRef, self PIFRef) (_retval []VLANRef, _err error) {
	if (IsMock) {
		return _class.GetVLANSlaveOfMock(sessionID, self)
	}	
	_method := "PIF.get_VLAN_slave_of"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPIFRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVLANRefSetToGo(_method + " -> ", _result.Value)
	return
}


var PIFClass_GetVLANMasterOfMockedCallback = func (sessionID SessionRef, self PIFRef) (_retval VLANRef, _err error) {
	log.Println("PIF.GetVLANMasterOf not mocked")
	_err = errors.New("PIF.GetVLANMasterOf not mocked")
	return
}

func (_class PIFClass) GetVLANMasterOfMock(sessionID SessionRef, self PIFRef) (_retval VLANRef, _err error) {
	return PIFClass_GetVLANMasterOfMockedCallback(sessionID, self)
}
// Get the VLAN_master_of field of the given PIF.
func (_class PIFClass) GetVLANMasterOf(sessionID SessionRef, self PIFRef) (_retval VLANRef, _err error) {
	if (IsMock) {
		return _class.GetVLANMasterOfMock(sessionID, self)
	}	
	_method := "PIF.get_VLAN_master_of"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPIFRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVLANRefToGo(_method + " -> ", _result.Value)
	return
}


var PIFClass_GetBondMasterOfMockedCallback = func (sessionID SessionRef, self PIFRef) (_retval []BondRef, _err error) {
	log.Println("PIF.GetBondMasterOf not mocked")
	_err = errors.New("PIF.GetBondMasterOf not mocked")
	return
}

func (_class PIFClass) GetBondMasterOfMock(sessionID SessionRef, self PIFRef) (_retval []BondRef, _err error) {
	return PIFClass_GetBondMasterOfMockedCallback(sessionID, self)
}
// Get the bond_master_of field of the given PIF.
func (_class PIFClass) GetBondMasterOf(sessionID SessionRef, self PIFRef) (_retval []BondRef, _err error) {
	if (IsMock) {
		return _class.GetBondMasterOfMock(sessionID, self)
	}	
	_method := "PIF.get_bond_master_of"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPIFRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertBondRefSetToGo(_method + " -> ", _result.Value)
	return
}


var PIFClass_GetBondSlaveOfMockedCallback = func (sessionID SessionRef, self PIFRef) (_retval BondRef, _err error) {
	log.Println("PIF.GetBondSlaveOf not mocked")
	_err = errors.New("PIF.GetBondSlaveOf not mocked")
	return
}

func (_class PIFClass) GetBondSlaveOfMock(sessionID SessionRef, self PIFRef) (_retval BondRef, _err error) {
	return PIFClass_GetBondSlaveOfMockedCallback(sessionID, self)
}
// Get the bond_slave_of field of the given PIF.
func (_class PIFClass) GetBondSlaveOf(sessionID SessionRef, self PIFRef) (_retval BondRef, _err error) {
	if (IsMock) {
		return _class.GetBondSlaveOfMock(sessionID, self)
	}	
	_method := "PIF.get_bond_slave_of"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPIFRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertBondRefToGo(_method + " -> ", _result.Value)
	return
}


var PIFClass_GetDNSMockedCallback = func (sessionID SessionRef, self PIFRef) (_retval string, _err error) {
	log.Println("PIF.GetDNS not mocked")
	_err = errors.New("PIF.GetDNS not mocked")
	return
}

func (_class PIFClass) GetDNSMock(sessionID SessionRef, self PIFRef) (_retval string, _err error) {
	return PIFClass_GetDNSMockedCallback(sessionID, self)
}
// Get the DNS field of the given PIF.
func (_class PIFClass) GetDNS(sessionID SessionRef, self PIFRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetDNSMock(sessionID, self)
	}	
	_method := "PIF.get_DNS"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPIFRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var PIFClass_GetGatewayMockedCallback = func (sessionID SessionRef, self PIFRef) (_retval string, _err error) {
	log.Println("PIF.GetGateway not mocked")
	_err = errors.New("PIF.GetGateway not mocked")
	return
}

func (_class PIFClass) GetGatewayMock(sessionID SessionRef, self PIFRef) (_retval string, _err error) {
	return PIFClass_GetGatewayMockedCallback(sessionID, self)
}
// Get the gateway field of the given PIF.
func (_class PIFClass) GetGateway(sessionID SessionRef, self PIFRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetGatewayMock(sessionID, self)
	}	
	_method := "PIF.get_gateway"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPIFRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var PIFClass_GetNetmaskMockedCallback = func (sessionID SessionRef, self PIFRef) (_retval string, _err error) {
	log.Println("PIF.GetNetmask not mocked")
	_err = errors.New("PIF.GetNetmask not mocked")
	return
}

func (_class PIFClass) GetNetmaskMock(sessionID SessionRef, self PIFRef) (_retval string, _err error) {
	return PIFClass_GetNetmaskMockedCallback(sessionID, self)
}
// Get the netmask field of the given PIF.
func (_class PIFClass) GetNetmask(sessionID SessionRef, self PIFRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetNetmaskMock(sessionID, self)
	}	
	_method := "PIF.get_netmask"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPIFRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var PIFClass_GetIPMockedCallback = func (sessionID SessionRef, self PIFRef) (_retval string, _err error) {
	log.Println("PIF.GetIP not mocked")
	_err = errors.New("PIF.GetIP not mocked")
	return
}

func (_class PIFClass) GetIPMock(sessionID SessionRef, self PIFRef) (_retval string, _err error) {
	return PIFClass_GetIPMockedCallback(sessionID, self)
}
// Get the IP field of the given PIF.
func (_class PIFClass) GetIP(sessionID SessionRef, self PIFRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetIPMock(sessionID, self)
	}	
	_method := "PIF.get_IP"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPIFRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var PIFClass_GetIPConfigurationModeMockedCallback = func (sessionID SessionRef, self PIFRef) (_retval IPConfigurationMode, _err error) {
	log.Println("PIF.GetIPConfigurationMode not mocked")
	_err = errors.New("PIF.GetIPConfigurationMode not mocked")
	return
}

func (_class PIFClass) GetIPConfigurationModeMock(sessionID SessionRef, self PIFRef) (_retval IPConfigurationMode, _err error) {
	return PIFClass_GetIPConfigurationModeMockedCallback(sessionID, self)
}
// Get the ip_configuration_mode field of the given PIF.
func (_class PIFClass) GetIPConfigurationMode(sessionID SessionRef, self PIFRef) (_retval IPConfigurationMode, _err error) {
	if (IsMock) {
		return _class.GetIPConfigurationModeMock(sessionID, self)
	}	
	_method := "PIF.get_ip_configuration_mode"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPIFRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertEnumIPConfigurationModeToGo(_method + " -> ", _result.Value)
	return
}


var PIFClass_GetCurrentlyAttachedMockedCallback = func (sessionID SessionRef, self PIFRef) (_retval bool, _err error) {
	log.Println("PIF.GetCurrentlyAttached not mocked")
	_err = errors.New("PIF.GetCurrentlyAttached not mocked")
	return
}

func (_class PIFClass) GetCurrentlyAttachedMock(sessionID SessionRef, self PIFRef) (_retval bool, _err error) {
	return PIFClass_GetCurrentlyAttachedMockedCallback(sessionID, self)
}
// Get the currently_attached field of the given PIF.
func (_class PIFClass) GetCurrentlyAttached(sessionID SessionRef, self PIFRef) (_retval bool, _err error) {
	if (IsMock) {
		return _class.GetCurrentlyAttachedMock(sessionID, self)
	}	
	_method := "PIF.get_currently_attached"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPIFRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var PIFClass_GetPhysicalMockedCallback = func (sessionID SessionRef, self PIFRef) (_retval bool, _err error) {
	log.Println("PIF.GetPhysical not mocked")
	_err = errors.New("PIF.GetPhysical not mocked")
	return
}

func (_class PIFClass) GetPhysicalMock(sessionID SessionRef, self PIFRef) (_retval bool, _err error) {
	return PIFClass_GetPhysicalMockedCallback(sessionID, self)
}
// Get the physical field of the given PIF.
func (_class PIFClass) GetPhysical(sessionID SessionRef, self PIFRef) (_retval bool, _err error) {
	if (IsMock) {
		return _class.GetPhysicalMock(sessionID, self)
	}	
	_method := "PIF.get_physical"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPIFRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var PIFClass_GetMetricsMockedCallback = func (sessionID SessionRef, self PIFRef) (_retval PIFMetricsRef, _err error) {
	log.Println("PIF.GetMetrics not mocked")
	_err = errors.New("PIF.GetMetrics not mocked")
	return
}

func (_class PIFClass) GetMetricsMock(sessionID SessionRef, self PIFRef) (_retval PIFMetricsRef, _err error) {
	return PIFClass_GetMetricsMockedCallback(sessionID, self)
}
// Get the metrics field of the given PIF.
func (_class PIFClass) GetMetrics(sessionID SessionRef, self PIFRef) (_retval PIFMetricsRef, _err error) {
	if (IsMock) {
		return _class.GetMetricsMock(sessionID, self)
	}	
	_method := "PIF.get_metrics"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPIFRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertPIFMetricsRefToGo(_method + " -> ", _result.Value)
	return
}


var PIFClass_GetVLANMockedCallback = func (sessionID SessionRef, self PIFRef) (_retval int, _err error) {
	log.Println("PIF.GetVLAN not mocked")
	_err = errors.New("PIF.GetVLAN not mocked")
	return
}

func (_class PIFClass) GetVLANMock(sessionID SessionRef, self PIFRef) (_retval int, _err error) {
	return PIFClass_GetVLANMockedCallback(sessionID, self)
}
// Get the VLAN field of the given PIF.
func (_class PIFClass) GetVLAN(sessionID SessionRef, self PIFRef) (_retval int, _err error) {
	if (IsMock) {
		return _class.GetVLANMock(sessionID, self)
	}	
	_method := "PIF.get_VLAN"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPIFRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var PIFClass_GetMTUMockedCallback = func (sessionID SessionRef, self PIFRef) (_retval int, _err error) {
	log.Println("PIF.GetMTU not mocked")
	_err = errors.New("PIF.GetMTU not mocked")
	return
}

func (_class PIFClass) GetMTUMock(sessionID SessionRef, self PIFRef) (_retval int, _err error) {
	return PIFClass_GetMTUMockedCallback(sessionID, self)
}
// Get the MTU field of the given PIF.
func (_class PIFClass) GetMTU(sessionID SessionRef, self PIFRef) (_retval int, _err error) {
	if (IsMock) {
		return _class.GetMTUMock(sessionID, self)
	}	
	_method := "PIF.get_MTU"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPIFRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var PIFClass_GetMACMockedCallback = func (sessionID SessionRef, self PIFRef) (_retval string, _err error) {
	log.Println("PIF.GetMAC not mocked")
	_err = errors.New("PIF.GetMAC not mocked")
	return
}

func (_class PIFClass) GetMACMock(sessionID SessionRef, self PIFRef) (_retval string, _err error) {
	return PIFClass_GetMACMockedCallback(sessionID, self)
}
// Get the MAC field of the given PIF.
func (_class PIFClass) GetMAC(sessionID SessionRef, self PIFRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetMACMock(sessionID, self)
	}	
	_method := "PIF.get_MAC"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPIFRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var PIFClass_GetHostMockedCallback = func (sessionID SessionRef, self PIFRef) (_retval HostRef, _err error) {
	log.Println("PIF.GetHost not mocked")
	_err = errors.New("PIF.GetHost not mocked")
	return
}

func (_class PIFClass) GetHostMock(sessionID SessionRef, self PIFRef) (_retval HostRef, _err error) {
	return PIFClass_GetHostMockedCallback(sessionID, self)
}
// Get the host field of the given PIF.
func (_class PIFClass) GetHost(sessionID SessionRef, self PIFRef) (_retval HostRef, _err error) {
	if (IsMock) {
		return _class.GetHostMock(sessionID, self)
	}	
	_method := "PIF.get_host"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPIFRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var PIFClass_GetNetworkMockedCallback = func (sessionID SessionRef, self PIFRef) (_retval NetworkRef, _err error) {
	log.Println("PIF.GetNetwork not mocked")
	_err = errors.New("PIF.GetNetwork not mocked")
	return
}

func (_class PIFClass) GetNetworkMock(sessionID SessionRef, self PIFRef) (_retval NetworkRef, _err error) {
	return PIFClass_GetNetworkMockedCallback(sessionID, self)
}
// Get the network field of the given PIF.
func (_class PIFClass) GetNetwork(sessionID SessionRef, self PIFRef) (_retval NetworkRef, _err error) {
	if (IsMock) {
		return _class.GetNetworkMock(sessionID, self)
	}	
	_method := "PIF.get_network"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPIFRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertNetworkRefToGo(_method + " -> ", _result.Value)
	return
}


var PIFClass_GetDeviceMockedCallback = func (sessionID SessionRef, self PIFRef) (_retval string, _err error) {
	log.Println("PIF.GetDevice not mocked")
	_err = errors.New("PIF.GetDevice not mocked")
	return
}

func (_class PIFClass) GetDeviceMock(sessionID SessionRef, self PIFRef) (_retval string, _err error) {
	return PIFClass_GetDeviceMockedCallback(sessionID, self)
}
// Get the device field of the given PIF.
func (_class PIFClass) GetDevice(sessionID SessionRef, self PIFRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetDeviceMock(sessionID, self)
	}	
	_method := "PIF.get_device"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPIFRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var PIFClass_GetUUIDMockedCallback = func (sessionID SessionRef, self PIFRef) (_retval string, _err error) {
	log.Println("PIF.GetUUID not mocked")
	_err = errors.New("PIF.GetUUID not mocked")
	return
}

func (_class PIFClass) GetUUIDMock(sessionID SessionRef, self PIFRef) (_retval string, _err error) {
	return PIFClass_GetUUIDMockedCallback(sessionID, self)
}
// Get the uuid field of the given PIF.
func (_class PIFClass) GetUUID(sessionID SessionRef, self PIFRef) (_retval string, _err error) {
	if (IsMock) {
		return _class.GetUUIDMock(sessionID, self)
	}	
	_method := "PIF.get_uuid"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPIFRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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


var PIFClass_GetByUUIDMockedCallback = func (sessionID SessionRef, uuid string) (_retval PIFRef, _err error) {
	log.Println("PIF.GetByUUID not mocked")
	_err = errors.New("PIF.GetByUUID not mocked")
	return
}

func (_class PIFClass) GetByUUIDMock(sessionID SessionRef, uuid string) (_retval PIFRef, _err error) {
	return PIFClass_GetByUUIDMockedCallback(sessionID, uuid)
}
// Get a reference to the PIF instance with the specified UUID.
func (_class PIFClass) GetByUUID(sessionID SessionRef, uuid string) (_retval PIFRef, _err error) {
	if (IsMock) {
		return _class.GetByUUIDMock(sessionID, uuid)
	}	
	_method := "PIF.get_by_uuid"
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
	_retval, _err = convertPIFRefToGo(_method + " -> ", _result.Value)
	return
}


var PIFClass_GetRecordMockedCallback = func (sessionID SessionRef, self PIFRef) (_retval PIFRecord, _err error) {
	log.Println("PIF.GetRecord not mocked")
	_err = errors.New("PIF.GetRecord not mocked")
	return
}

func (_class PIFClass) GetRecordMock(sessionID SessionRef, self PIFRef) (_retval PIFRecord, _err error) {
	return PIFClass_GetRecordMockedCallback(sessionID, self)
}
// Get a record containing the current state of the given PIF.
func (_class PIFClass) GetRecord(sessionID SessionRef, self PIFRef) (_retval PIFRecord, _err error) {
	if (IsMock) {
		return _class.GetRecordMock(sessionID, self)
	}	
	_method := "PIF.get_record"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPIFRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertPIFRecordToGo(_method + " -> ", _result.Value)
	return
}
