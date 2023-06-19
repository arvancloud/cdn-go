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

// checks if the DomainPurgeTags type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DomainPurgeTags{}

// DomainPurgeTags struct for DomainPurgeTags
type DomainPurgeTags struct {
	DomainId *string `json:"domain_id,omitempty"`
	Tags []string `json:"tags,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

// NewDomainPurgeTags instantiates a new DomainPurgeTags object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDomainPurgeTags() *DomainPurgeTags {
	this := DomainPurgeTags{}
	return &this
}

// NewDomainPurgeTagsWithDefaults instantiates a new DomainPurgeTags object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDomainPurgeTagsWithDefaults() *DomainPurgeTags {
	this := DomainPurgeTags{}
	return &this
}

// GetDomainId returns the DomainId field value if set, zero value otherwise.
func (o *DomainPurgeTags) GetDomainId() string {
	if o == nil || IsNil(o.DomainId) {
		var ret string
		return ret
	}
	return *o.DomainId
}

// GetDomainIdOk returns a tuple with the DomainId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainPurgeTags) GetDomainIdOk() (*string, bool) {
	if o == nil || IsNil(o.DomainId) {
		return nil, false
	}
	return o.DomainId, true
}

// HasDomainId returns a boolean if a field has been set.
func (o *DomainPurgeTags) HasDomainId() bool {
	if o != nil && !IsNil(o.DomainId) {
		return true
	}

	return false
}

// SetDomainId gets a reference to the given string and assigns it to the DomainId field.
func (o *DomainPurgeTags) SetDomainId(v string) {
	o.DomainId = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *DomainPurgeTags) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainPurgeTags) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *DomainPurgeTags) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *DomainPurgeTags) SetTags(v []string) {
	o.Tags = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *DomainPurgeTags) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainPurgeTags) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *DomainPurgeTags) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *DomainPurgeTags) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *DomainPurgeTags) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainPurgeTags) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *DomainPurgeTags) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *DomainPurgeTags) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o DomainPurgeTags) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DomainPurgeTags) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DomainId) {
		toSerialize["domain_id"] = o.DomainId
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	return toSerialize, nil
}

type NullableDomainPurgeTags struct {
	value *DomainPurgeTags
	isSet bool
}

func (v NullableDomainPurgeTags) Get() *DomainPurgeTags {
	return v.value
}

func (v *NullableDomainPurgeTags) Set(val *DomainPurgeTags) {
	v.value = val
	v.isSet = true
}

func (v NullableDomainPurgeTags) IsSet() bool {
	return v.isSet
}

func (v *NullableDomainPurgeTags) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDomainPurgeTags(val *DomainPurgeTags) *NullableDomainPurgeTags {
	return &NullableDomainPurgeTags{value: val, isSet: true}
}

func (v NullableDomainPurgeTags) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDomainPurgeTags) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

