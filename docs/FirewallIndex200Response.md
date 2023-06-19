# FirewallIndex200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**Firewall**](Firewall.md) |  | [optional] 
**Message** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewFirewallIndex200Response

`func NewFirewallIndex200Response() *FirewallIndex200Response`

NewFirewallIndex200Response instantiates a new FirewallIndex200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirewallIndex200ResponseWithDefaults

`func NewFirewallIndex200ResponseWithDefaults() *FirewallIndex200Response`

NewFirewallIndex200ResponseWithDefaults instantiates a new FirewallIndex200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *FirewallIndex200Response) GetData() Firewall`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *FirewallIndex200Response) GetDataOk() (*Firewall, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *FirewallIndex200Response) SetData(v Firewall)`

SetData sets Data field to given value.

### HasData

`func (o *FirewallIndex200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMessage

`func (o *FirewallIndex200Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *FirewallIndex200Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *FirewallIndex200Response) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *FirewallIndex200Response) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *FirewallIndex200Response) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *FirewallIndex200Response) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil

[[Back to Model list]](HOW-TO.md#documentation-for-models) [[Back to API list]](HOW-TO.md#documentation-for-api-endpoints) [[Back to README]](HOW-TO.md)


