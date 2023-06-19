# DomainWaf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsEnabled** | Pointer to **bool** |  | [optional] [readonly] 
**Mode** | Pointer to **string** |  | [optional] 
**Packages** | Pointer to [**[]DomainWafPackage**](DomainWafPackage.md) | Pacakges and their configurations that are used to configure WAF. | [optional] [readonly] 
**Rules** | Pointer to [**[]WafRule**](WafRule.md) |  | [optional] [readonly] 
**FWafStatus** | Pointer to **string** |  | [optional] 
**FThreshold** | Pointer to **string** |  | [optional] 
**FProvider** | Pointer to **string** |  | [optional] 

## Methods

### NewDomainWaf

`func NewDomainWaf() *DomainWaf`

NewDomainWaf instantiates a new DomainWaf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainWafWithDefaults

`func NewDomainWafWithDefaults() *DomainWaf`

NewDomainWafWithDefaults instantiates a new DomainWaf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsEnabled

`func (o *DomainWaf) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *DomainWaf) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *DomainWaf) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *DomainWaf) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### GetMode

`func (o *DomainWaf) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *DomainWaf) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *DomainWaf) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *DomainWaf) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetPackages

`func (o *DomainWaf) GetPackages() []DomainWafPackage`

GetPackages returns the Packages field if non-nil, zero value otherwise.

### GetPackagesOk

`func (o *DomainWaf) GetPackagesOk() (*[]DomainWafPackage, bool)`

GetPackagesOk returns a tuple with the Packages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackages

`func (o *DomainWaf) SetPackages(v []DomainWafPackage)`

SetPackages sets Packages field to given value.

### HasPackages

`func (o *DomainWaf) HasPackages() bool`

HasPackages returns a boolean if a field has been set.

### GetRules

`func (o *DomainWaf) GetRules() []WafRule`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *DomainWaf) GetRulesOk() (*[]WafRule, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *DomainWaf) SetRules(v []WafRule)`

SetRules sets Rules field to given value.

### HasRules

`func (o *DomainWaf) HasRules() bool`

HasRules returns a boolean if a field has been set.

### GetFWafStatus

`func (o *DomainWaf) GetFWafStatus() string`

GetFWafStatus returns the FWafStatus field if non-nil, zero value otherwise.

### GetFWafStatusOk

`func (o *DomainWaf) GetFWafStatusOk() (*string, bool)`

GetFWafStatusOk returns a tuple with the FWafStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFWafStatus

`func (o *DomainWaf) SetFWafStatus(v string)`

SetFWafStatus sets FWafStatus field to given value.

### HasFWafStatus

`func (o *DomainWaf) HasFWafStatus() bool`

HasFWafStatus returns a boolean if a field has been set.

### GetFThreshold

`func (o *DomainWaf) GetFThreshold() string`

GetFThreshold returns the FThreshold field if non-nil, zero value otherwise.

### GetFThresholdOk

`func (o *DomainWaf) GetFThresholdOk() (*string, bool)`

GetFThresholdOk returns a tuple with the FThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFThreshold

`func (o *DomainWaf) SetFThreshold(v string)`

SetFThreshold sets FThreshold field to given value.

### HasFThreshold

`func (o *DomainWaf) HasFThreshold() bool`

HasFThreshold returns a boolean if a field has been set.

### GetFProvider

`func (o *DomainWaf) GetFProvider() string`

GetFProvider returns the FProvider field if non-nil, zero value otherwise.

### GetFProviderOk

`func (o *DomainWaf) GetFProviderOk() (*string, bool)`

GetFProviderOk returns a tuple with the FProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFProvider

`func (o *DomainWaf) SetFProvider(v string)`

SetFProvider sets FProvider field to given value.

### HasFProvider

`func (o *DomainWaf) HasFProvider() bool`

HasFProvider returns a boolean if a field has been set.


[[Back to Model list]](HOW-TO.md#documentation-for-models) [[Back to API list]](HOW-TO.md#documentation-for-api-endpoints) [[Back to README]](HOW-TO.md)


