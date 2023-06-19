# RateLimitingUpdate200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**RateLimit**](RateLimit.md) |  | [optional] 
**Message** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewRateLimitingUpdate200Response

`func NewRateLimitingUpdate200Response() *RateLimitingUpdate200Response`

NewRateLimitingUpdate200Response instantiates a new RateLimitingUpdate200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRateLimitingUpdate200ResponseWithDefaults

`func NewRateLimitingUpdate200ResponseWithDefaults() *RateLimitingUpdate200Response`

NewRateLimitingUpdate200ResponseWithDefaults instantiates a new RateLimitingUpdate200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *RateLimitingUpdate200Response) GetData() RateLimit`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *RateLimitingUpdate200Response) GetDataOk() (*RateLimit, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *RateLimitingUpdate200Response) SetData(v RateLimit)`

SetData sets Data field to given value.

### HasData

`func (o *RateLimitingUpdate200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMessage

`func (o *RateLimitingUpdate200Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *RateLimitingUpdate200Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *RateLimitingUpdate200Response) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *RateLimitingUpdate200Response) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *RateLimitingUpdate200Response) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *RateLimitingUpdate200Response) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil

[[Back to Model list]](HOW-TO.md#documentation-for-models) [[Back to API list]](HOW-TO.md#documentation-for-api-endpoints) [[Back to README]](HOW-TO.md)


