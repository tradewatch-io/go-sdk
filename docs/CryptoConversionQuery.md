# CryptoConversionQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**From** | **string** |  | 
**To** | **string** |  | 
**Amount** | Pointer to **float32** |  | [optional] [default to 10]
**Precision** | Pointer to **int32** |  | [optional] [default to 8]

## Methods

### NewCryptoConversionQuery

`func NewCryptoConversionQuery(from string, to string, ) *CryptoConversionQuery`

NewCryptoConversionQuery instantiates a new CryptoConversionQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCryptoConversionQueryWithDefaults

`func NewCryptoConversionQueryWithDefaults() *CryptoConversionQuery`

NewCryptoConversionQueryWithDefaults instantiates a new CryptoConversionQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFrom

`func (o *CryptoConversionQuery) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *CryptoConversionQuery) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *CryptoConversionQuery) SetFrom(v string)`

SetFrom sets From field to given value.


### GetTo

`func (o *CryptoConversionQuery) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *CryptoConversionQuery) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *CryptoConversionQuery) SetTo(v string)`

SetTo sets To field to given value.


### GetAmount

`func (o *CryptoConversionQuery) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *CryptoConversionQuery) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *CryptoConversionQuery) SetAmount(v float32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *CryptoConversionQuery) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetPrecision

`func (o *CryptoConversionQuery) GetPrecision() int32`

GetPrecision returns the Precision field if non-nil, zero value otherwise.

### GetPrecisionOk

`func (o *CryptoConversionQuery) GetPrecisionOk() (*int32, bool)`

GetPrecisionOk returns a tuple with the Precision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrecision

`func (o *CryptoConversionQuery) SetPrecision(v int32)`

SetPrecision sets Precision field to given value.

### HasPrecision

`func (o *CryptoConversionQuery) HasPrecision() bool`

HasPrecision returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


