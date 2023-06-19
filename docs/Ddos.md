# Ddos

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsEnabled** | Pointer to **bool** |  | [optional] [readonly] 
**ProtectionMode** | Pointer to **string** |  | [optional] 
**CaptchaService** | Pointer to **string** |  | [optional] 
**Ttl** | Pointer to **int32** | Time in seconds for cookie max-age | [optional] 
**HttpsOnly** | Pointer to **bool** | Adds \&quot;SameSite&#x3D;None; Secure\&quot; to set-cookie header | [optional] 
**Preflight** | Pointer to [**DdosPreflight**](DdosPreflight.md) |  | [optional] 
**Rules** | Pointer to [**[]DdosRule**](DdosRule.md) |  | [optional] [readonly] 

## Methods

### NewDdos

`func NewDdos() *Ddos`

NewDdos instantiates a new Ddos object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDdosWithDefaults

`func NewDdosWithDefaults() *Ddos`

NewDdosWithDefaults instantiates a new Ddos object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsEnabled

`func (o *Ddos) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *Ddos) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *Ddos) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *Ddos) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### GetProtectionMode

`func (o *Ddos) GetProtectionMode() string`

GetProtectionMode returns the ProtectionMode field if non-nil, zero value otherwise.

### GetProtectionModeOk

`func (o *Ddos) GetProtectionModeOk() (*string, bool)`

GetProtectionModeOk returns a tuple with the ProtectionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionMode

`func (o *Ddos) SetProtectionMode(v string)`

SetProtectionMode sets ProtectionMode field to given value.

### HasProtectionMode

`func (o *Ddos) HasProtectionMode() bool`

HasProtectionMode returns a boolean if a field has been set.

### GetCaptchaService

`func (o *Ddos) GetCaptchaService() string`

GetCaptchaService returns the CaptchaService field if non-nil, zero value otherwise.

### GetCaptchaServiceOk

`func (o *Ddos) GetCaptchaServiceOk() (*string, bool)`

GetCaptchaServiceOk returns a tuple with the CaptchaService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptchaService

`func (o *Ddos) SetCaptchaService(v string)`

SetCaptchaService sets CaptchaService field to given value.

### HasCaptchaService

`func (o *Ddos) HasCaptchaService() bool`

HasCaptchaService returns a boolean if a field has been set.

### GetTtl

`func (o *Ddos) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *Ddos) GetTtlOk() (*int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *Ddos) SetTtl(v int32)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *Ddos) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetHttpsOnly

`func (o *Ddos) GetHttpsOnly() bool`

GetHttpsOnly returns the HttpsOnly field if non-nil, zero value otherwise.

### GetHttpsOnlyOk

`func (o *Ddos) GetHttpsOnlyOk() (*bool, bool)`

GetHttpsOnlyOk returns a tuple with the HttpsOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpsOnly

`func (o *Ddos) SetHttpsOnly(v bool)`

SetHttpsOnly sets HttpsOnly field to given value.

### HasHttpsOnly

`func (o *Ddos) HasHttpsOnly() bool`

HasHttpsOnly returns a boolean if a field has been set.

### GetPreflight

`func (o *Ddos) GetPreflight() DdosPreflight`

GetPreflight returns the Preflight field if non-nil, zero value otherwise.

### GetPreflightOk

`func (o *Ddos) GetPreflightOk() (*DdosPreflight, bool)`

GetPreflightOk returns a tuple with the Preflight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreflight

`func (o *Ddos) SetPreflight(v DdosPreflight)`

SetPreflight sets Preflight field to given value.

### HasPreflight

`func (o *Ddos) HasPreflight() bool`

HasPreflight returns a boolean if a field has been set.

### GetRules

`func (o *Ddos) GetRules() []DdosRule`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *Ddos) GetRulesOk() (*[]DdosRule, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *Ddos) SetRules(v []DdosRule)`

SetRules sets Rules field to given value.

### HasRules

`func (o *Ddos) HasRules() bool`

HasRules returns a boolean if a field has been set.


[[Back to Model list]](HOW-TO.md#documentation-for-models) [[Back to API list]](HOW-TO.md#documentation-for-api-endpoints) [[Back to README]](HOW-TO.md)


