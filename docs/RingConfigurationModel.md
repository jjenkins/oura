# RingConfigurationModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Color** | Pointer to [**NullableRingColor**](RingColor.md) |  | [optional] 
**Design** | Pointer to [**NullableRingDesign**](RingDesign.md) |  | [optional] 
**FirmwareVersion** | Pointer to **NullableString** |  | [optional] 
**HardwareType** | Pointer to [**NullableRingHardwareType**](RingHardwareType.md) |  | [optional] 
**SetUpAt** | Pointer to **string** |  | [optional] 
**Size** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewRingConfigurationModel

`func NewRingConfigurationModel(id string, ) *RingConfigurationModel`

NewRingConfigurationModel instantiates a new RingConfigurationModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRingConfigurationModelWithDefaults

`func NewRingConfigurationModelWithDefaults() *RingConfigurationModel`

NewRingConfigurationModelWithDefaults instantiates a new RingConfigurationModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RingConfigurationModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RingConfigurationModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RingConfigurationModel) SetId(v string)`

SetId sets Id field to given value.


### GetColor

`func (o *RingConfigurationModel) GetColor() RingColor`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *RingConfigurationModel) GetColorOk() (*RingColor, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *RingConfigurationModel) SetColor(v RingColor)`

SetColor sets Color field to given value.

### HasColor

`func (o *RingConfigurationModel) HasColor() bool`

HasColor returns a boolean if a field has been set.

### SetColorNil

`func (o *RingConfigurationModel) SetColorNil(b bool)`

 SetColorNil sets the value for Color to be an explicit nil

### UnsetColor
`func (o *RingConfigurationModel) UnsetColor()`

UnsetColor ensures that no value is present for Color, not even an explicit nil
### GetDesign

`func (o *RingConfigurationModel) GetDesign() RingDesign`

GetDesign returns the Design field if non-nil, zero value otherwise.

### GetDesignOk

`func (o *RingConfigurationModel) GetDesignOk() (*RingDesign, bool)`

GetDesignOk returns a tuple with the Design field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesign

`func (o *RingConfigurationModel) SetDesign(v RingDesign)`

SetDesign sets Design field to given value.

### HasDesign

`func (o *RingConfigurationModel) HasDesign() bool`

HasDesign returns a boolean if a field has been set.

### SetDesignNil

`func (o *RingConfigurationModel) SetDesignNil(b bool)`

 SetDesignNil sets the value for Design to be an explicit nil

### UnsetDesign
`func (o *RingConfigurationModel) UnsetDesign()`

UnsetDesign ensures that no value is present for Design, not even an explicit nil
### GetFirmwareVersion

`func (o *RingConfigurationModel) GetFirmwareVersion() string`

GetFirmwareVersion returns the FirmwareVersion field if non-nil, zero value otherwise.

### GetFirmwareVersionOk

`func (o *RingConfigurationModel) GetFirmwareVersionOk() (*string, bool)`

GetFirmwareVersionOk returns a tuple with the FirmwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmwareVersion

`func (o *RingConfigurationModel) SetFirmwareVersion(v string)`

SetFirmwareVersion sets FirmwareVersion field to given value.

### HasFirmwareVersion

`func (o *RingConfigurationModel) HasFirmwareVersion() bool`

HasFirmwareVersion returns a boolean if a field has been set.

### SetFirmwareVersionNil

`func (o *RingConfigurationModel) SetFirmwareVersionNil(b bool)`

 SetFirmwareVersionNil sets the value for FirmwareVersion to be an explicit nil

### UnsetFirmwareVersion
`func (o *RingConfigurationModel) UnsetFirmwareVersion()`

UnsetFirmwareVersion ensures that no value is present for FirmwareVersion, not even an explicit nil
### GetHardwareType

`func (o *RingConfigurationModel) GetHardwareType() RingHardwareType`

GetHardwareType returns the HardwareType field if non-nil, zero value otherwise.

### GetHardwareTypeOk

`func (o *RingConfigurationModel) GetHardwareTypeOk() (*RingHardwareType, bool)`

GetHardwareTypeOk returns a tuple with the HardwareType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardwareType

`func (o *RingConfigurationModel) SetHardwareType(v RingHardwareType)`

SetHardwareType sets HardwareType field to given value.

### HasHardwareType

`func (o *RingConfigurationModel) HasHardwareType() bool`

HasHardwareType returns a boolean if a field has been set.

### SetHardwareTypeNil

`func (o *RingConfigurationModel) SetHardwareTypeNil(b bool)`

 SetHardwareTypeNil sets the value for HardwareType to be an explicit nil

### UnsetHardwareType
`func (o *RingConfigurationModel) UnsetHardwareType()`

UnsetHardwareType ensures that no value is present for HardwareType, not even an explicit nil
### GetSetUpAt

`func (o *RingConfigurationModel) GetSetUpAt() string`

GetSetUpAt returns the SetUpAt field if non-nil, zero value otherwise.

### GetSetUpAtOk

`func (o *RingConfigurationModel) GetSetUpAtOk() (*string, bool)`

GetSetUpAtOk returns a tuple with the SetUpAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetUpAt

`func (o *RingConfigurationModel) SetSetUpAt(v string)`

SetSetUpAt sets SetUpAt field to given value.

### HasSetUpAt

`func (o *RingConfigurationModel) HasSetUpAt() bool`

HasSetUpAt returns a boolean if a field has been set.

### GetSize

`func (o *RingConfigurationModel) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *RingConfigurationModel) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *RingConfigurationModel) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *RingConfigurationModel) HasSize() bool`

HasSize returns a boolean if a field has been set.

### SetSizeNil

`func (o *RingConfigurationModel) SetSizeNil(b bool)`

 SetSizeNil sets the value for Size to be an explicit nil

### UnsetSize
`func (o *RingConfigurationModel) UnsetSize()`

UnsetSize ensures that no value is present for Size, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


