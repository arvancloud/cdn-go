# FeatureSet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Label** | Pointer to **string** |  | [optional] 
**Features** | Pointer to [**[]FeatureDefinition**](FeatureDefinition.md) |  | [optional] 

## Methods

### NewFeatureSet

`func NewFeatureSet() *FeatureSet`

NewFeatureSet instantiates a new FeatureSet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeatureSetWithDefaults

`func NewFeatureSetWithDefaults() *FeatureSet`

NewFeatureSetWithDefaults instantiates a new FeatureSet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FeatureSet) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FeatureSet) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FeatureSet) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FeatureSet) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLabel

`func (o *FeatureSet) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *FeatureSet) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *FeatureSet) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *FeatureSet) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetFeatures

`func (o *FeatureSet) GetFeatures() []FeatureDefinition`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *FeatureSet) GetFeaturesOk() (*[]FeatureDefinition, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *FeatureSet) SetFeatures(v []FeatureDefinition)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *FeatureSet) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.


[[Back to Model list]](HOW-TO.md#documentation-for-models) [[Back to API list]](HOW-TO.md#documentation-for-api-endpoints) [[Back to README]](HOW-TO.md)


