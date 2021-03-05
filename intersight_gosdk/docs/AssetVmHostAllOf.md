# AssetVmHostAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "asset.VmHost"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "asset.VmHost"]
**AccountMoid** | Pointer to **string** | Reference to virtualization target account ID. | [optional] [readonly] 
**Connected** | Pointer to **int64** | The connection status of the host. 1 represents being connected and 0 represents being disconnected. | [optional] [readonly] 
**RegisteredDeviceMoid** | Pointer to **string** | Reference to virtualization target device ID. | [optional] [readonly] 

## Methods

### NewAssetVmHostAllOf

`func NewAssetVmHostAllOf(classId string, objectType string, ) *AssetVmHostAllOf`

NewAssetVmHostAllOf instantiates a new AssetVmHostAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetVmHostAllOfWithDefaults

`func NewAssetVmHostAllOfWithDefaults() *AssetVmHostAllOf`

NewAssetVmHostAllOfWithDefaults instantiates a new AssetVmHostAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *AssetVmHostAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *AssetVmHostAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *AssetVmHostAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *AssetVmHostAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *AssetVmHostAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *AssetVmHostAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAccountMoid

`func (o *AssetVmHostAllOf) GetAccountMoid() string`

GetAccountMoid returns the AccountMoid field if non-nil, zero value otherwise.

### GetAccountMoidOk

`func (o *AssetVmHostAllOf) GetAccountMoidOk() (*string, bool)`

GetAccountMoidOk returns a tuple with the AccountMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountMoid

`func (o *AssetVmHostAllOf) SetAccountMoid(v string)`

SetAccountMoid sets AccountMoid field to given value.

### HasAccountMoid

`func (o *AssetVmHostAllOf) HasAccountMoid() bool`

HasAccountMoid returns a boolean if a field has been set.

### GetConnected

`func (o *AssetVmHostAllOf) GetConnected() int64`

GetConnected returns the Connected field if non-nil, zero value otherwise.

### GetConnectedOk

`func (o *AssetVmHostAllOf) GetConnectedOk() (*int64, bool)`

GetConnectedOk returns a tuple with the Connected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnected

`func (o *AssetVmHostAllOf) SetConnected(v int64)`

SetConnected sets Connected field to given value.

### HasConnected

`func (o *AssetVmHostAllOf) HasConnected() bool`

HasConnected returns a boolean if a field has been set.

### GetRegisteredDeviceMoid

`func (o *AssetVmHostAllOf) GetRegisteredDeviceMoid() string`

GetRegisteredDeviceMoid returns the RegisteredDeviceMoid field if non-nil, zero value otherwise.

### GetRegisteredDeviceMoidOk

`func (o *AssetVmHostAllOf) GetRegisteredDeviceMoidOk() (*string, bool)`

GetRegisteredDeviceMoidOk returns a tuple with the RegisteredDeviceMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDeviceMoid

`func (o *AssetVmHostAllOf) SetRegisteredDeviceMoid(v string)`

SetRegisteredDeviceMoid sets RegisteredDeviceMoid field to given value.

### HasRegisteredDeviceMoid

`func (o *AssetVmHostAllOf) HasRegisteredDeviceMoid() bool`

HasRegisteredDeviceMoid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

