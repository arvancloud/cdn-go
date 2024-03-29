/*
ArvanCloud CDN Services

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 4.99.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package r1cdn

import (
	"encoding/json"
)

// checks if the DomainWafData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DomainWafData{}

// DomainWafData struct for DomainWafData
type DomainWafData struct {
	Data *DomainWaf `json:"data,omitempty"`
}

// NewDomainWafData instantiates a new DomainWafData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDomainWafData() *DomainWafData {
	this := DomainWafData{}
	return &this
}

// NewDomainWafDataWithDefaults instantiates a new DomainWafData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDomainWafDataWithDefaults() *DomainWafData {
	this := DomainWafData{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *DomainWafData) GetData() DomainWaf {
	if o == nil || IsNil(o.Data) {
		var ret DomainWaf
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainWafData) GetDataOk() (*DomainWaf, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *DomainWafData) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given DomainWaf and assigns it to the Data field.
func (o *DomainWafData) SetData(v DomainWaf) {
	o.Data = &v
}

func (o DomainWafData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DomainWafData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableDomainWafData struct {
	value *DomainWafData
	isSet bool
}

func (v NullableDomainWafData) Get() *DomainWafData {
	return v.value
}

func (v *NullableDomainWafData) Set(val *DomainWafData) {
	v.value = val
	v.isSet = true
}

func (v NullableDomainWafData) IsSet() bool {
	return v.isSet
}

func (v *NullableDomainWafData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDomainWafData(val *DomainWafData) *NullableDomainWafData {
	return &NullableDomainWafData{value: val, isSet: true}
}

func (v NullableDomainWafData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDomainWafData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


