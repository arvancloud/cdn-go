# ReportsAttacksIndex200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]AttackReportItem**](AttackReportItem.md) |  | [optional] 
**Links** | Pointer to [**PaginatedResponseLinks**](PaginatedResponseLinks.md) |  | [optional] 
**Meta** | Pointer to [**PaginatedResponseMeta**](PaginatedResponseMeta.md) |  | [optional] 

## Methods

### NewReportsAttacksIndex200Response

`func NewReportsAttacksIndex200Response() *ReportsAttacksIndex200Response`

NewReportsAttacksIndex200Response instantiates a new ReportsAttacksIndex200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReportsAttacksIndex200ResponseWithDefaults

`func NewReportsAttacksIndex200ResponseWithDefaults() *ReportsAttacksIndex200Response`

NewReportsAttacksIndex200ResponseWithDefaults instantiates a new ReportsAttacksIndex200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ReportsAttacksIndex200Response) GetData() []AttackReportItem`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ReportsAttacksIndex200Response) GetDataOk() (*[]AttackReportItem, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ReportsAttacksIndex200Response) SetData(v []AttackReportItem)`

SetData sets Data field to given value.

### HasData

`func (o *ReportsAttacksIndex200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetLinks

`func (o *ReportsAttacksIndex200Response) GetLinks() PaginatedResponseLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ReportsAttacksIndex200Response) GetLinksOk() (*PaginatedResponseLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ReportsAttacksIndex200Response) SetLinks(v PaginatedResponseLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ReportsAttacksIndex200Response) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMeta

`func (o *ReportsAttacksIndex200Response) GetMeta() PaginatedResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ReportsAttacksIndex200Response) GetMetaOk() (*PaginatedResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ReportsAttacksIndex200Response) SetMeta(v PaginatedResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ReportsAttacksIndex200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](HOW-TO.md#documentation-for-models) [[Back to API list]](HOW-TO.md#documentation-for-api-endpoints) [[Back to README]](HOW-TO.md)


