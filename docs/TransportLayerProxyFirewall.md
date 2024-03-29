# TransportLayerProxyFirewall

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Access** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Match** | Pointer to [**TransportLayerProxyMatch**](TransportLayerProxyMatch.md) |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 

## Methods

### NewTransportLayerProxyFirewall

`func NewTransportLayerProxyFirewall() *TransportLayerProxyFirewall`

NewTransportLayerProxyFirewall instantiates a new TransportLayerProxyFirewall object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransportLayerProxyFirewallWithDefaults

`func NewTransportLayerProxyFirewallWithDefaults() *TransportLayerProxyFirewall`

NewTransportLayerProxyFirewallWithDefaults instantiates a new TransportLayerProxyFirewall object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccess

`func (o *TransportLayerProxyFirewall) GetAccess() string`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *TransportLayerProxyFirewall) GetAccessOk() (*string, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *TransportLayerProxyFirewall) SetAccess(v string)`

SetAccess sets Access field to given value.

### HasAccess

`func (o *TransportLayerProxyFirewall) HasAccess() bool`

HasAccess returns a boolean if a field has been set.

### GetName

`func (o *TransportLayerProxyFirewall) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TransportLayerProxyFirewall) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TransportLayerProxyFirewall) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TransportLayerProxyFirewall) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *TransportLayerProxyFirewall) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TransportLayerProxyFirewall) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TransportLayerProxyFirewall) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TransportLayerProxyFirewall) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetType

`func (o *TransportLayerProxyFirewall) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TransportLayerProxyFirewall) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TransportLayerProxyFirewall) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *TransportLayerProxyFirewall) HasType() bool`

HasType returns a boolean if a field has been set.

### GetMatch

`func (o *TransportLayerProxyFirewall) GetMatch() TransportLayerProxyMatch`

GetMatch returns the Match field if non-nil, zero value otherwise.

### GetMatchOk

`func (o *TransportLayerProxyFirewall) GetMatchOk() (*TransportLayerProxyMatch, bool)`

GetMatchOk returns a tuple with the Match field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatch

`func (o *TransportLayerProxyFirewall) SetMatch(v TransportLayerProxyMatch)`

SetMatch sets Match field to given value.

### HasMatch

`func (o *TransportLayerProxyFirewall) HasMatch() bool`

HasMatch returns a boolean if a field has been set.

### GetActive

`func (o *TransportLayerProxyFirewall) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *TransportLayerProxyFirewall) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *TransportLayerProxyFirewall) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *TransportLayerProxyFirewall) HasActive() bool`

HasActive returns a boolean if a field has been set.


[[Back to Model list]](HOW-TO.md#documentation-for-models) [[Back to API list]](HOW-TO.md#documentation-for-api-endpoints) [[Back to README]](HOW-TO.md)


