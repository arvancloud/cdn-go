# DdosData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**Ddos**](Ddos.md) |  | [optional] 
**Message** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewDdosData

`func NewDdosData() *DdosData`

NewDdosData instantiates a new DdosData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDdosDataWithDefaults

`func NewDdosDataWithDefaults() *DdosData`

NewDdosDataWithDefaults instantiates a new DdosData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *DdosData) GetData() Ddos`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *DdosData) GetDataOk() (*Ddos, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *DdosData) SetData(v Ddos)`

SetData sets Data field to given value.

### HasData

`func (o *DdosData) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMessage

`func (o *DdosData) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *DdosData) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *DdosData) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *DdosData) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *DdosData) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *DdosData) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil

[[Back to Model list]](HOW-TO.md#documentation-for-models) [[Back to API list]](HOW-TO.md#documentation-for-api-endpoints) [[Back to README]](HOW-TO.md)


