# Waf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsEnabled** | Pointer to **bool** |  | [optional] [readonly] 
**Mode** | Pointer to **string** |  | [optional] 
**Packages** | Pointer to [**[]DomainWafPackage**](DomainWafPackage.md) | Pacakges and their configurations that are used to configure WAF. | [optional] [readonly] 
**Rules** | Pointer to [**[]WafRule**](WafRule.md) |  | [optional] [readonly] 

## Methods

### NewWaf

`func NewWaf() *Waf`

NewWaf instantiates a new Waf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWafWithDefaults

`func NewWafWithDefaults() *Waf`

NewWafWithDefaults instantiates a new Waf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsEnabled

`func (o *Waf) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *Waf) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *Waf) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *Waf) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### GetMode

`func (o *Waf) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *Waf) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *Waf) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *Waf) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetPackages

`func (o *Waf) GetPackages() []DomainWafPackage`

GetPackages returns the Packages field if non-nil, zero value otherwise.

### GetPackagesOk

`func (o *Waf) GetPackagesOk() (*[]DomainWafPackage, bool)`

GetPackagesOk returns a tuple with the Packages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackages

`func (o *Waf) SetPackages(v []DomainWafPackage)`

SetPackages sets Packages field to given value.

### HasPackages

`func (o *Waf) HasPackages() bool`

HasPackages returns a boolean if a field has been set.

### GetRules

`func (o *Waf) GetRules() []WafRule`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *Waf) GetRulesOk() (*[]WafRule, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *Waf) SetRules(v []WafRule)`

SetRules sets Rules field to given value.

### HasRules

`func (o *Waf) HasRules() bool`

HasRules returns a boolean if a field has been set.


[[Back to Model list]](HOW-TO.md#documentation-for-models) [[Back to API list]](HOW-TO.md#documentation-for-api-endpoints) [[Back to README]](HOW-TO.md)


