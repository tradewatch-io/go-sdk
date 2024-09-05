# ErrorResponseBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | [**ErrorDetails**](ErrorDetails.md) |  | 

## Methods

### NewErrorResponseBody

`func NewErrorResponseBody(error_ ErrorDetails, ) *ErrorResponseBody`

NewErrorResponseBody instantiates a new ErrorResponseBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorResponseBodyWithDefaults

`func NewErrorResponseBodyWithDefaults() *ErrorResponseBody`

NewErrorResponseBodyWithDefaults instantiates a new ErrorResponseBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *ErrorResponseBody) GetError() ErrorDetails`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ErrorResponseBody) GetErrorOk() (*ErrorDetails, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ErrorResponseBody) SetError(v ErrorDetails)`

SetError sets Error field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


