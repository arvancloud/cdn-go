# Firewall

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsEnabled** | Pointer to **bool** |  | [optional] [readonly] 
**DefaultAction** | Pointer to **string** |  | [optional] 
**DefaultActionDetails** | Pointer to **map[string]interface{}** |  | [optional] 
**VerifySni** | Pointer to **bool** | True to verify that SNI and hostname are equal | [optional] [default to false]
**Rules** | Pointer to [**[]FirewallRuleView**](FirewallRuleView.md) |  | [optional] [readonly] 

## Methods

### NewFirewall

`func NewFirewall() *Firewall`

NewFirewall instantiates a new Firewall object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirewallWithDefaults

`func NewFirewallWithDefaults() *Firewall`

NewFirewallWithDefaults instantiates a new Firewall object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsEnabled

`func (o *Firewall) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *Firewall) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *Firewall) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *Firewall) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### GetDefaultAction

`func (o *Firewall) GetDefaultAction() string`

GetDefaultAction returns the DefaultAction field if non-nil, zero value otherwise.

### GetDefaultActionOk

`func (o *Firewall) GetDefaultActionOk() (*string, bool)`

GetDefaultActionOk returns a tuple with the DefaultAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultAction

`func (o *Firewall) SetDefaultAction(v string)`

SetDefaultAction sets DefaultAction field to given value.

### HasDefaultAction

`func (o *Firewall) HasDefaultAction() bool`

HasDefaultAction returns a boolean if a field has been set.

### GetDefaultActionDetails

`func (o *Firewall) GetDefaultActionDetails() map[string]interface{}`

GetDefaultActionDetails returns the DefaultActionDetails field if non-nil, zero value otherwise.

### GetDefaultActionDetailsOk

`func (o *Firewall) GetDefaultActionDetailsOk() (*map[string]interface{}, bool)`

GetDefaultActionDetailsOk returns a tuple with the DefaultActionDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultActionDetails

`func (o *Firewall) SetDefaultActionDetails(v map[string]interface{})`

SetDefaultActionDetails sets DefaultActionDetails field to given value.

### HasDefaultActionDetails

`func (o *Firewall) HasDefaultActionDetails() bool`

HasDefaultActionDetails returns a boolean if a field has been set.

### SetDefaultActionDetailsNil

`func (o *Firewall) SetDefaultActionDetailsNil(b bool)`

 SetDefaultActionDetailsNil sets the value for DefaultActionDetails to be an explicit nil

### UnsetDefaultActionDetails
`func (o *Firewall) UnsetDefaultActionDetails()`

UnsetDefaultActionDetails ensures that no value is present for DefaultActionDetails, not even an explicit nil
### GetVerifySni

`func (o *Firewall) GetVerifySni() bool`

GetVerifySni returns the VerifySni field if non-nil, zero value otherwise.

### GetVerifySniOk

`func (o *Firewall) GetVerifySniOk() (*bool, bool)`

GetVerifySniOk returns a tuple with the VerifySni field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifySni

`func (o *Firewall) SetVerifySni(v bool)`

SetVerifySni sets VerifySni field to given value.

### HasVerifySni

`func (o *Firewall) HasVerifySni() bool`

HasVerifySni returns a boolean if a field has been set.

### GetRules

`func (o *Firewall) GetRules() []FirewallRuleView`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *Firewall) GetRulesOk() (*[]FirewallRuleView, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *Firewall) SetRules(v []FirewallRuleView)`

SetRules sets Rules field to given value.

### HasRules

`func (o *Firewall) HasRules() bool`

HasRules returns a boolean if a field has been set.


[[Back to Model list]](HOW-TO.md#documentation-for-models) [[Back to API list]](HOW-TO.md#documentation-for-api-endpoints) [[Back to README]](HOW-TO.md)


