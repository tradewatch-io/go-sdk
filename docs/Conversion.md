# Conversion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Info** | [**ConversionInfo**](ConversionInfo.md) |  | 
**Query** | [**ConversionQuery**](ConversionQuery.md) |  | 
**Result** | **float32** |  | 

## Methods

### NewConversion

`func NewConversion(info ConversionInfo, query ConversionQuery, result float32, ) *Conversion`

NewConversion instantiates a new Conversion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConversionWithDefaults

`func NewConversionWithDefaults() *Conversion`

NewConversionWithDefaults instantiates a new Conversion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInfo

`func (o *Conversion) GetInfo() ConversionInfo`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *Conversion) GetInfoOk() (*ConversionInfo, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *Conversion) SetInfo(v ConversionInfo)`

SetInfo sets Info field to given value.


### GetQuery

`func (o *Conversion) GetQuery() ConversionQuery`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *Conversion) GetQueryOk() (*ConversionQuery, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *Conversion) SetQuery(v ConversionQuery)`

SetQuery sets Query field to given value.


### GetResult

`func (o *Conversion) GetResult() float32`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *Conversion) GetResultOk() (*float32, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *Conversion) SetResult(v float32)`

SetResult sets Result field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


