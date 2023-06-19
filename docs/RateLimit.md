# RateLimit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DdosDetection** | Pointer to **bool** |  | [optional] 
**ExcludeSources** | Pointer to **[]string** |  | [optional] 
**Rules** | Pointer to [**[]RateLimitRuleView**](RateLimitRuleView.md) |  | [optional] [readonly] 

## Methods

### NewRateLimit

`func NewRateLimit() *RateLimit`

NewRateLimit instantiates a new RateLimit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRateLimitWithDefaults

`func NewRateLimitWithDefaults() *RateLimit`

NewRateLimitWithDefaults instantiates a new RateLimit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDdosDetection

`func (o *RateLimit) GetDdosDetection() bool`

GetDdosDetection returns the DdosDetection field if non-nil, zero value otherwise.

### GetDdosDetectionOk

`func (o *RateLimit) GetDdosDetectionOk() (*bool, bool)`

GetDdosDetectionOk returns a tuple with the DdosDetection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdosDetection

`func (o *RateLimit) SetDdosDetection(v bool)`

SetDdosDetection sets DdosDetection field to given value.

### HasDdosDetection

`func (o *RateLimit) HasDdosDetection() bool`

HasDdosDetection returns a boolean if a field has been set.

### GetExcludeSources

`func (o *RateLimit) GetExcludeSources() []string`

GetExcludeSources returns the ExcludeSources field if non-nil, zero value otherwise.

### GetExcludeSourcesOk

`func (o *RateLimit) GetExcludeSourcesOk() (*[]string, bool)`

GetExcludeSourcesOk returns a tuple with the ExcludeSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeSources

`func (o *RateLimit) SetExcludeSources(v []string)`

SetExcludeSources sets ExcludeSources field to given value.

### HasExcludeSources

`func (o *RateLimit) HasExcludeSources() bool`

HasExcludeSources returns a boolean if a field has been set.

### GetRules

`func (o *RateLimit) GetRules() []RateLimitRuleView`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *RateLimit) GetRulesOk() (*[]RateLimitRuleView, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *RateLimit) SetRules(v []RateLimitRuleView)`

SetRules sets Rules field to given value.

### HasRules

`func (o *RateLimit) HasRules() bool`

HasRules returns a boolean if a field has been set.


[[Back to Model list]](HOW-TO.md#documentation-for-models) [[Back to API list]](HOW-TO.md#documentation-for-api-endpoints) [[Back to README]](HOW-TO.md)


