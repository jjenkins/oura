# ResilienceContributors

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SleepRecovery** | **float32** | Sleep recovery contributor to the resilience score. Range: [0, 100] | 
**DaytimeRecovery** | **float32** | Daytime recovery contributor to the resilience score. Range: [0, 100] | 
**Stress** | **float32** | Stress contributor to the resilience score. Range: [0, 100] | 

## Methods

### NewResilienceContributors

`func NewResilienceContributors(sleepRecovery float32, daytimeRecovery float32, stress float32, ) *ResilienceContributors`

NewResilienceContributors instantiates a new ResilienceContributors object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResilienceContributorsWithDefaults

`func NewResilienceContributorsWithDefaults() *ResilienceContributors`

NewResilienceContributorsWithDefaults instantiates a new ResilienceContributors object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSleepRecovery

`func (o *ResilienceContributors) GetSleepRecovery() float32`

GetSleepRecovery returns the SleepRecovery field if non-nil, zero value otherwise.

### GetSleepRecoveryOk

`func (o *ResilienceContributors) GetSleepRecoveryOk() (*float32, bool)`

GetSleepRecoveryOk returns a tuple with the SleepRecovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSleepRecovery

`func (o *ResilienceContributors) SetSleepRecovery(v float32)`

SetSleepRecovery sets SleepRecovery field to given value.


### GetDaytimeRecovery

`func (o *ResilienceContributors) GetDaytimeRecovery() float32`

GetDaytimeRecovery returns the DaytimeRecovery field if non-nil, zero value otherwise.

### GetDaytimeRecoveryOk

`func (o *ResilienceContributors) GetDaytimeRecoveryOk() (*float32, bool)`

GetDaytimeRecoveryOk returns a tuple with the DaytimeRecovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaytimeRecovery

`func (o *ResilienceContributors) SetDaytimeRecovery(v float32)`

SetDaytimeRecovery sets DaytimeRecovery field to given value.


### GetStress

`func (o *ResilienceContributors) GetStress() float32`

GetStress returns the Stress field if non-nil, zero value otherwise.

### GetStressOk

`func (o *ResilienceContributors) GetStressOk() (*float32, bool)`

GetStressOk returns a tuple with the Stress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStress

`func (o *ResilienceContributors) SetStress(v float32)`

SetStress sets Stress field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


