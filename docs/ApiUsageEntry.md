# ApiUsageEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Requests** | **int32** |  | 
**Data** | [**ApiUsageDataTransfer**](ApiUsageDataTransfer.md) |  | 

## Methods

### NewApiUsageEntry

`func NewApiUsageEntry(requests int32, data ApiUsageDataTransfer, ) *ApiUsageEntry`

NewApiUsageEntry instantiates a new ApiUsageEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiUsageEntryWithDefaults

`func NewApiUsageEntryWithDefaults() *ApiUsageEntry`

NewApiUsageEntryWithDefaults instantiates a new ApiUsageEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequests

`func (o *ApiUsageEntry) GetRequests() int32`

GetRequests returns the Requests field if non-nil, zero value otherwise.

### GetRequestsOk

`func (o *ApiUsageEntry) GetRequestsOk() (*int32, bool)`

GetRequestsOk returns a tuple with the Requests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequests

`func (o *ApiUsageEntry) SetRequests(v int32)`

SetRequests sets Requests field to given value.


### GetData

`func (o *ApiUsageEntry) GetData() ApiUsageDataTransfer`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ApiUsageEntry) GetDataOk() (*ApiUsageDataTransfer, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ApiUsageEntry) SetData(v ApiUsageDataTransfer)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


