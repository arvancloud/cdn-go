# DnsSecData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**DnsSec**](DnsSec.md) |  | [optional] 

## Methods

### NewDnsSecData

`func NewDnsSecData() *DnsSecData`

NewDnsSecData instantiates a new DnsSecData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDnsSecDataWithDefaults

`func NewDnsSecDataWithDefaults() *DnsSecData`

NewDnsSecDataWithDefaults instantiates a new DnsSecData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *DnsSecData) GetData() DnsSec`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *DnsSecData) GetDataOk() (*DnsSec, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *DnsSecData) SetData(v DnsSec)`

SetData sets Data field to given value.

### HasData

`func (o *DnsSecData) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](HOW-TO.md#documentation-for-models) [[Back to API list]](HOW-TO.md#documentation-for-api-endpoints) [[Back to README]](HOW-TO.md)


