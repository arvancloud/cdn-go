/*
ArvanCloud CDN Services

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 4.99.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package r1cdn

import (
	"encoding/json"
	"time"
)

// checks if the TransportLayerProxy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransportLayerProxy{}

// TransportLayerProxy struct for TransportLayerProxy
type TransportLayerProxy struct {
	Id *string `json:"id,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	Status *string `json:"status,omitempty"`
	AppName *string `json:"app_name,omitempty"`
	Description *string `json:"description,omitempty"`
	Domain *string `json:"domain,omitempty"`
	Port *int32 `json:"port,omitempty"`
	Ip NullableString `json:"ip,omitempty"`
	ProxyProtocol *string `json:"proxy_protocol,omitempty"`
	BalanceAlgorithm *string `json:"balance_algorithm,omitempty"`
	Servers []TransportLayerProxyServersInner `json:"servers,omitempty"`
	FirewallDefaultAction *string `json:"firewall_default_action,omitempty"`
	Firewalls []TransportLayerProxyFirewallsInner `json:"firewalls,omitempty"`
}

// NewTransportLayerProxy instantiates a new TransportLayerProxy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransportLayerProxy() *TransportLayerProxy {
	this := TransportLayerProxy{}
	return &this
}

// NewTransportLayerProxyWithDefaults instantiates a new TransportLayerProxy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransportLayerProxyWithDefaults() *TransportLayerProxy {
	this := TransportLayerProxy{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TransportLayerProxy) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportLayerProxy) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TransportLayerProxy) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *TransportLayerProxy) SetId(v string) {
	o.Id = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *TransportLayerProxy) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportLayerProxy) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *TransportLayerProxy) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *TransportLayerProxy) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *TransportLayerProxy) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportLayerProxy) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *TransportLayerProxy) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *TransportLayerProxy) SetStatus(v string) {
	o.Status = &v
}

// GetAppName returns the AppName field value if set, zero value otherwise.
func (o *TransportLayerProxy) GetAppName() string {
	if o == nil || IsNil(o.AppName) {
		var ret string
		return ret
	}
	return *o.AppName
}

// GetAppNameOk returns a tuple with the AppName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportLayerProxy) GetAppNameOk() (*string, bool) {
	if o == nil || IsNil(o.AppName) {
		return nil, false
	}
	return o.AppName, true
}

// HasAppName returns a boolean if a field has been set.
func (o *TransportLayerProxy) HasAppName() bool {
	if o != nil && !IsNil(o.AppName) {
		return true
	}

	return false
}

// SetAppName gets a reference to the given string and assigns it to the AppName field.
func (o *TransportLayerProxy) SetAppName(v string) {
	o.AppName = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *TransportLayerProxy) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportLayerProxy) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *TransportLayerProxy) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *TransportLayerProxy) SetDescription(v string) {
	o.Description = &v
}

// GetDomain returns the Domain field value if set, zero value otherwise.
func (o *TransportLayerProxy) GetDomain() string {
	if o == nil || IsNil(o.Domain) {
		var ret string
		return ret
	}
	return *o.Domain
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportLayerProxy) GetDomainOk() (*string, bool) {
	if o == nil || IsNil(o.Domain) {
		return nil, false
	}
	return o.Domain, true
}

// HasDomain returns a boolean if a field has been set.
func (o *TransportLayerProxy) HasDomain() bool {
	if o != nil && !IsNil(o.Domain) {
		return true
	}

	return false
}

// SetDomain gets a reference to the given string and assigns it to the Domain field.
func (o *TransportLayerProxy) SetDomain(v string) {
	o.Domain = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *TransportLayerProxy) GetPort() int32 {
	if o == nil || IsNil(o.Port) {
		var ret int32
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportLayerProxy) GetPortOk() (*int32, bool) {
	if o == nil || IsNil(o.Port) {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *TransportLayerProxy) HasPort() bool {
	if o != nil && !IsNil(o.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given int32 and assigns it to the Port field.
func (o *TransportLayerProxy) SetPort(v int32) {
	o.Port = &v
}

// GetIp returns the Ip field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TransportLayerProxy) GetIp() string {
	if o == nil || IsNil(o.Ip.Get()) {
		var ret string
		return ret
	}
	return *o.Ip.Get()
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransportLayerProxy) GetIpOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Ip.Get(), o.Ip.IsSet()
}

// HasIp returns a boolean if a field has been set.
func (o *TransportLayerProxy) HasIp() bool {
	if o != nil && o.Ip.IsSet() {
		return true
	}

	return false
}

// SetIp gets a reference to the given NullableString and assigns it to the Ip field.
func (o *TransportLayerProxy) SetIp(v string) {
	o.Ip.Set(&v)
}
// SetIpNil sets the value for Ip to be an explicit nil
func (o *TransportLayerProxy) SetIpNil() {
	o.Ip.Set(nil)
}

// UnsetIp ensures that no value is present for Ip, not even an explicit nil
func (o *TransportLayerProxy) UnsetIp() {
	o.Ip.Unset()
}

// GetProxyProtocol returns the ProxyProtocol field value if set, zero value otherwise.
func (o *TransportLayerProxy) GetProxyProtocol() string {
	if o == nil || IsNil(o.ProxyProtocol) {
		var ret string
		return ret
	}
	return *o.ProxyProtocol
}

// GetProxyProtocolOk returns a tuple with the ProxyProtocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportLayerProxy) GetProxyProtocolOk() (*string, bool) {
	if o == nil || IsNil(o.ProxyProtocol) {
		return nil, false
	}
	return o.ProxyProtocol, true
}

// HasProxyProtocol returns a boolean if a field has been set.
func (o *TransportLayerProxy) HasProxyProtocol() bool {
	if o != nil && !IsNil(o.ProxyProtocol) {
		return true
	}

	return false
}

// SetProxyProtocol gets a reference to the given string and assigns it to the ProxyProtocol field.
func (o *TransportLayerProxy) SetProxyProtocol(v string) {
	o.ProxyProtocol = &v
}

// GetBalanceAlgorithm returns the BalanceAlgorithm field value if set, zero value otherwise.
func (o *TransportLayerProxy) GetBalanceAlgorithm() string {
	if o == nil || IsNil(o.BalanceAlgorithm) {
		var ret string
		return ret
	}
	return *o.BalanceAlgorithm
}

// GetBalanceAlgorithmOk returns a tuple with the BalanceAlgorithm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportLayerProxy) GetBalanceAlgorithmOk() (*string, bool) {
	if o == nil || IsNil(o.BalanceAlgorithm) {
		return nil, false
	}
	return o.BalanceAlgorithm, true
}

// HasBalanceAlgorithm returns a boolean if a field has been set.
func (o *TransportLayerProxy) HasBalanceAlgorithm() bool {
	if o != nil && !IsNil(o.BalanceAlgorithm) {
		return true
	}

	return false
}

// SetBalanceAlgorithm gets a reference to the given string and assigns it to the BalanceAlgorithm field.
func (o *TransportLayerProxy) SetBalanceAlgorithm(v string) {
	o.BalanceAlgorithm = &v
}

// GetServers returns the Servers field value if set, zero value otherwise.
func (o *TransportLayerProxy) GetServers() []TransportLayerProxyServersInner {
	if o == nil || IsNil(o.Servers) {
		var ret []TransportLayerProxyServersInner
		return ret
	}
	return o.Servers
}

// GetServersOk returns a tuple with the Servers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportLayerProxy) GetServersOk() ([]TransportLayerProxyServersInner, bool) {
	if o == nil || IsNil(o.Servers) {
		return nil, false
	}
	return o.Servers, true
}

// HasServers returns a boolean if a field has been set.
func (o *TransportLayerProxy) HasServers() bool {
	if o != nil && !IsNil(o.Servers) {
		return true
	}

	return false
}

// SetServers gets a reference to the given []TransportLayerProxyServersInner and assigns it to the Servers field.
func (o *TransportLayerProxy) SetServers(v []TransportLayerProxyServersInner) {
	o.Servers = v
}

// GetFirewallDefaultAction returns the FirewallDefaultAction field value if set, zero value otherwise.
func (o *TransportLayerProxy) GetFirewallDefaultAction() string {
	if o == nil || IsNil(o.FirewallDefaultAction) {
		var ret string
		return ret
	}
	return *o.FirewallDefaultAction
}

// GetFirewallDefaultActionOk returns a tuple with the FirewallDefaultAction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportLayerProxy) GetFirewallDefaultActionOk() (*string, bool) {
	if o == nil || IsNil(o.FirewallDefaultAction) {
		return nil, false
	}
	return o.FirewallDefaultAction, true
}

// HasFirewallDefaultAction returns a boolean if a field has been set.
func (o *TransportLayerProxy) HasFirewallDefaultAction() bool {
	if o != nil && !IsNil(o.FirewallDefaultAction) {
		return true
	}

	return false
}

// SetFirewallDefaultAction gets a reference to the given string and assigns it to the FirewallDefaultAction field.
func (o *TransportLayerProxy) SetFirewallDefaultAction(v string) {
	o.FirewallDefaultAction = &v
}

// GetFirewalls returns the Firewalls field value if set, zero value otherwise.
func (o *TransportLayerProxy) GetFirewalls() []TransportLayerProxyFirewallsInner {
	if o == nil || IsNil(o.Firewalls) {
		var ret []TransportLayerProxyFirewallsInner
		return ret
	}
	return o.Firewalls
}

// GetFirewallsOk returns a tuple with the Firewalls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportLayerProxy) GetFirewallsOk() ([]TransportLayerProxyFirewallsInner, bool) {
	if o == nil || IsNil(o.Firewalls) {
		return nil, false
	}
	return o.Firewalls, true
}

// HasFirewalls returns a boolean if a field has been set.
func (o *TransportLayerProxy) HasFirewalls() bool {
	if o != nil && !IsNil(o.Firewalls) {
		return true
	}

	return false
}

// SetFirewalls gets a reference to the given []TransportLayerProxyFirewallsInner and assigns it to the Firewalls field.
func (o *TransportLayerProxy) SetFirewalls(v []TransportLayerProxyFirewallsInner) {
	o.Firewalls = v
}

func (o TransportLayerProxy) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransportLayerProxy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: id is readOnly
	// skip: created_at is readOnly
	// skip: status is readOnly
	if !IsNil(o.AppName) {
		toSerialize["app_name"] = o.AppName
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Domain) {
		toSerialize["domain"] = o.Domain
	}
	if !IsNil(o.Port) {
		toSerialize["port"] = o.Port
	}
	if o.Ip.IsSet() {
		toSerialize["ip"] = o.Ip.Get()
	}
	if !IsNil(o.ProxyProtocol) {
		toSerialize["proxy_protocol"] = o.ProxyProtocol
	}
	if !IsNil(o.BalanceAlgorithm) {
		toSerialize["balance_algorithm"] = o.BalanceAlgorithm
	}
	if !IsNil(o.Servers) {
		toSerialize["servers"] = o.Servers
	}
	if !IsNil(o.FirewallDefaultAction) {
		toSerialize["firewall_default_action"] = o.FirewallDefaultAction
	}
	if !IsNil(o.Firewalls) {
		toSerialize["firewalls"] = o.Firewalls
	}
	return toSerialize, nil
}

type NullableTransportLayerProxy struct {
	value *TransportLayerProxy
	isSet bool
}

func (v NullableTransportLayerProxy) Get() *TransportLayerProxy {
	return v.value
}

func (v *NullableTransportLayerProxy) Set(val *TransportLayerProxy) {
	v.value = val
	v.isSet = true
}

func (v NullableTransportLayerProxy) IsSet() bool {
	return v.isSet
}

func (v *NullableTransportLayerProxy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransportLayerProxy(val *TransportLayerProxy) *NullableTransportLayerProxy {
	return &NullableTransportLayerProxy{value: val, isSet: true}
}

func (v NullableTransportLayerProxy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransportLayerProxy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

