# SleepModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**AverageBreath** | **NullableFloat32** |  | 
**AverageHeartRate** | **NullableFloat32** |  | 
**AverageHrv** | **NullableInt32** |  | 
**AwakeTime** | **NullableInt32** |  | 
**BedtimeEnd** | **string** | Bedtime end of the sleep. | 
**BedtimeStart** | **string** | Bedtime start of the sleep. | 
**Day** | **string** | Day that the sleep belongs to. | 
**DeepSleepDuration** | **NullableInt32** |  | 
**Efficiency** | **NullableInt32** |  | 
**HeartRate** | [**NullableSampleModel**](SampleModel.md) |  | 
**Hrv** | [**NullableSampleModel**](SampleModel.md) |  | 
**Latency** | **NullableInt32** |  | 
**LightSleepDuration** | **NullableInt32** |  | 
**LowBatteryAlert** | **bool** | Flag indicating if a low battery alert occurred. | 
**LowestHeartRate** | **NullableInt32** |  | 
**Movement30Sec** | **NullableString** |  | 
**Period** | **int32** | ECore sleep period identifier. | 
**Readiness** | [**NullableReadinessSummary**](ReadinessSummary.md) |  | 
**ReadinessScoreDelta** | **NullableInt32** |  | 
**RemSleepDuration** | **NullableInt32** |  | 
**RestlessPeriods** | **NullableInt32** |  | 
**SleepPhase5Min** | **NullableString** |  | 
**SleepScoreDelta** | **NullableInt32** |  | 
**SleepAlgorithmVersion** | [**NullableSleepAlgorithmVersion**](SleepAlgorithmVersion.md) |  | 
**TimeInBed** | **int32** | Duration spent in bed in seconds. | 
**TotalSleepDuration** | **NullableInt32** |  | 
**Type** | [**SleepType**](SleepType.md) |  | 

## Methods

### NewSleepModel

`func NewSleepModel(id string, averageBreath NullableFloat32, averageHeartRate NullableFloat32, averageHrv NullableInt32, awakeTime NullableInt32, bedtimeEnd string, bedtimeStart string, day string, deepSleepDuration NullableInt32, efficiency NullableInt32, heartRate NullableSampleModel, hrv NullableSampleModel, latency NullableInt32, lightSleepDuration NullableInt32, lowBatteryAlert bool, lowestHeartRate NullableInt32, movement30Sec NullableString, period int32, readiness NullableReadinessSummary, readinessScoreDelta NullableInt32, remSleepDuration NullableInt32, restlessPeriods NullableInt32, sleepPhase5Min NullableString, sleepScoreDelta NullableInt32, sleepAlgorithmVersion NullableSleepAlgorithmVersion, timeInBed int32, totalSleepDuration NullableInt32, type_ SleepType, ) *SleepModel`

NewSleepModel instantiates a new SleepModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSleepModelWithDefaults

`func NewSleepModelWithDefaults() *SleepModel`

NewSleepModelWithDefaults instantiates a new SleepModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SleepModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SleepModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SleepModel) SetId(v string)`

SetId sets Id field to given value.


### GetAverageBreath

`func (o *SleepModel) GetAverageBreath() float32`

GetAverageBreath returns the AverageBreath field if non-nil, zero value otherwise.

### GetAverageBreathOk

`func (o *SleepModel) GetAverageBreathOk() (*float32, bool)`

GetAverageBreathOk returns a tuple with the AverageBreath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageBreath

`func (o *SleepModel) SetAverageBreath(v float32)`

SetAverageBreath sets AverageBreath field to given value.


### SetAverageBreathNil

`func (o *SleepModel) SetAverageBreathNil(b bool)`

 SetAverageBreathNil sets the value for AverageBreath to be an explicit nil

### UnsetAverageBreath
`func (o *SleepModel) UnsetAverageBreath()`

UnsetAverageBreath ensures that no value is present for AverageBreath, not even an explicit nil
### GetAverageHeartRate

`func (o *SleepModel) GetAverageHeartRate() float32`

GetAverageHeartRate returns the AverageHeartRate field if non-nil, zero value otherwise.

### GetAverageHeartRateOk

`func (o *SleepModel) GetAverageHeartRateOk() (*float32, bool)`

GetAverageHeartRateOk returns a tuple with the AverageHeartRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageHeartRate

`func (o *SleepModel) SetAverageHeartRate(v float32)`

SetAverageHeartRate sets AverageHeartRate field to given value.


### SetAverageHeartRateNil

`func (o *SleepModel) SetAverageHeartRateNil(b bool)`

 SetAverageHeartRateNil sets the value for AverageHeartRate to be an explicit nil

### UnsetAverageHeartRate
`func (o *SleepModel) UnsetAverageHeartRate()`

UnsetAverageHeartRate ensures that no value is present for AverageHeartRate, not even an explicit nil
### GetAverageHrv

`func (o *SleepModel) GetAverageHrv() int32`

GetAverageHrv returns the AverageHrv field if non-nil, zero value otherwise.

### GetAverageHrvOk

`func (o *SleepModel) GetAverageHrvOk() (*int32, bool)`

GetAverageHrvOk returns a tuple with the AverageHrv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageHrv

`func (o *SleepModel) SetAverageHrv(v int32)`

SetAverageHrv sets AverageHrv field to given value.


### SetAverageHrvNil

`func (o *SleepModel) SetAverageHrvNil(b bool)`

 SetAverageHrvNil sets the value for AverageHrv to be an explicit nil

### UnsetAverageHrv
`func (o *SleepModel) UnsetAverageHrv()`

UnsetAverageHrv ensures that no value is present for AverageHrv, not even an explicit nil
### GetAwakeTime

`func (o *SleepModel) GetAwakeTime() int32`

GetAwakeTime returns the AwakeTime field if non-nil, zero value otherwise.

### GetAwakeTimeOk

`func (o *SleepModel) GetAwakeTimeOk() (*int32, bool)`

GetAwakeTimeOk returns a tuple with the AwakeTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwakeTime

`func (o *SleepModel) SetAwakeTime(v int32)`

SetAwakeTime sets AwakeTime field to given value.


### SetAwakeTimeNil

`func (o *SleepModel) SetAwakeTimeNil(b bool)`

 SetAwakeTimeNil sets the value for AwakeTime to be an explicit nil

### UnsetAwakeTime
`func (o *SleepModel) UnsetAwakeTime()`

UnsetAwakeTime ensures that no value is present for AwakeTime, not even an explicit nil
### GetBedtimeEnd

`func (o *SleepModel) GetBedtimeEnd() string`

GetBedtimeEnd returns the BedtimeEnd field if non-nil, zero value otherwise.

### GetBedtimeEndOk

`func (o *SleepModel) GetBedtimeEndOk() (*string, bool)`

GetBedtimeEndOk returns a tuple with the BedtimeEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBedtimeEnd

`func (o *SleepModel) SetBedtimeEnd(v string)`

SetBedtimeEnd sets BedtimeEnd field to given value.


### GetBedtimeStart

`func (o *SleepModel) GetBedtimeStart() string`

GetBedtimeStart returns the BedtimeStart field if non-nil, zero value otherwise.

### GetBedtimeStartOk

`func (o *SleepModel) GetBedtimeStartOk() (*string, bool)`

GetBedtimeStartOk returns a tuple with the BedtimeStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBedtimeStart

`func (o *SleepModel) SetBedtimeStart(v string)`

SetBedtimeStart sets BedtimeStart field to given value.


### GetDay

`func (o *SleepModel) GetDay() string`

GetDay returns the Day field if non-nil, zero value otherwise.

### GetDayOk

`func (o *SleepModel) GetDayOk() (*string, bool)`

GetDayOk returns a tuple with the Day field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDay

`func (o *SleepModel) SetDay(v string)`

SetDay sets Day field to given value.


### GetDeepSleepDuration

`func (o *SleepModel) GetDeepSleepDuration() int32`

GetDeepSleepDuration returns the DeepSleepDuration field if non-nil, zero value otherwise.

### GetDeepSleepDurationOk

`func (o *SleepModel) GetDeepSleepDurationOk() (*int32, bool)`

GetDeepSleepDurationOk returns a tuple with the DeepSleepDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeepSleepDuration

`func (o *SleepModel) SetDeepSleepDuration(v int32)`

SetDeepSleepDuration sets DeepSleepDuration field to given value.


### SetDeepSleepDurationNil

`func (o *SleepModel) SetDeepSleepDurationNil(b bool)`

 SetDeepSleepDurationNil sets the value for DeepSleepDuration to be an explicit nil

### UnsetDeepSleepDuration
`func (o *SleepModel) UnsetDeepSleepDuration()`

UnsetDeepSleepDuration ensures that no value is present for DeepSleepDuration, not even an explicit nil
### GetEfficiency

`func (o *SleepModel) GetEfficiency() int32`

GetEfficiency returns the Efficiency field if non-nil, zero value otherwise.

### GetEfficiencyOk

`func (o *SleepModel) GetEfficiencyOk() (*int32, bool)`

GetEfficiencyOk returns a tuple with the Efficiency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEfficiency

`func (o *SleepModel) SetEfficiency(v int32)`

SetEfficiency sets Efficiency field to given value.


### SetEfficiencyNil

`func (o *SleepModel) SetEfficiencyNil(b bool)`

 SetEfficiencyNil sets the value for Efficiency to be an explicit nil

### UnsetEfficiency
`func (o *SleepModel) UnsetEfficiency()`

UnsetEfficiency ensures that no value is present for Efficiency, not even an explicit nil
### GetHeartRate

`func (o *SleepModel) GetHeartRate() SampleModel`

GetHeartRate returns the HeartRate field if non-nil, zero value otherwise.

### GetHeartRateOk

`func (o *SleepModel) GetHeartRateOk() (*SampleModel, bool)`

GetHeartRateOk returns a tuple with the HeartRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeartRate

`func (o *SleepModel) SetHeartRate(v SampleModel)`

SetHeartRate sets HeartRate field to given value.


### SetHeartRateNil

`func (o *SleepModel) SetHeartRateNil(b bool)`

 SetHeartRateNil sets the value for HeartRate to be an explicit nil

### UnsetHeartRate
`func (o *SleepModel) UnsetHeartRate()`

UnsetHeartRate ensures that no value is present for HeartRate, not even an explicit nil
### GetHrv

`func (o *SleepModel) GetHrv() SampleModel`

GetHrv returns the Hrv field if non-nil, zero value otherwise.

### GetHrvOk

`func (o *SleepModel) GetHrvOk() (*SampleModel, bool)`

GetHrvOk returns a tuple with the Hrv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHrv

`func (o *SleepModel) SetHrv(v SampleModel)`

SetHrv sets Hrv field to given value.


### SetHrvNil

`func (o *SleepModel) SetHrvNil(b bool)`

 SetHrvNil sets the value for Hrv to be an explicit nil

### UnsetHrv
`func (o *SleepModel) UnsetHrv()`

UnsetHrv ensures that no value is present for Hrv, not even an explicit nil
### GetLatency

`func (o *SleepModel) GetLatency() int32`

GetLatency returns the Latency field if non-nil, zero value otherwise.

### GetLatencyOk

`func (o *SleepModel) GetLatencyOk() (*int32, bool)`

GetLatencyOk returns a tuple with the Latency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatency

`func (o *SleepModel) SetLatency(v int32)`

SetLatency sets Latency field to given value.


### SetLatencyNil

`func (o *SleepModel) SetLatencyNil(b bool)`

 SetLatencyNil sets the value for Latency to be an explicit nil

### UnsetLatency
`func (o *SleepModel) UnsetLatency()`

UnsetLatency ensures that no value is present for Latency, not even an explicit nil
### GetLightSleepDuration

`func (o *SleepModel) GetLightSleepDuration() int32`

GetLightSleepDuration returns the LightSleepDuration field if non-nil, zero value otherwise.

### GetLightSleepDurationOk

`func (o *SleepModel) GetLightSleepDurationOk() (*int32, bool)`

GetLightSleepDurationOk returns a tuple with the LightSleepDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLightSleepDuration

`func (o *SleepModel) SetLightSleepDuration(v int32)`

SetLightSleepDuration sets LightSleepDuration field to given value.


### SetLightSleepDurationNil

`func (o *SleepModel) SetLightSleepDurationNil(b bool)`

 SetLightSleepDurationNil sets the value for LightSleepDuration to be an explicit nil

### UnsetLightSleepDuration
`func (o *SleepModel) UnsetLightSleepDuration()`

UnsetLightSleepDuration ensures that no value is present for LightSleepDuration, not even an explicit nil
### GetLowBatteryAlert

`func (o *SleepModel) GetLowBatteryAlert() bool`

GetLowBatteryAlert returns the LowBatteryAlert field if non-nil, zero value otherwise.

### GetLowBatteryAlertOk

`func (o *SleepModel) GetLowBatteryAlertOk() (*bool, bool)`

GetLowBatteryAlertOk returns a tuple with the LowBatteryAlert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowBatteryAlert

`func (o *SleepModel) SetLowBatteryAlert(v bool)`

SetLowBatteryAlert sets LowBatteryAlert field to given value.


### GetLowestHeartRate

`func (o *SleepModel) GetLowestHeartRate() int32`

GetLowestHeartRate returns the LowestHeartRate field if non-nil, zero value otherwise.

### GetLowestHeartRateOk

`func (o *SleepModel) GetLowestHeartRateOk() (*int32, bool)`

GetLowestHeartRateOk returns a tuple with the LowestHeartRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowestHeartRate

`func (o *SleepModel) SetLowestHeartRate(v int32)`

SetLowestHeartRate sets LowestHeartRate field to given value.


### SetLowestHeartRateNil

`func (o *SleepModel) SetLowestHeartRateNil(b bool)`

 SetLowestHeartRateNil sets the value for LowestHeartRate to be an explicit nil

### UnsetLowestHeartRate
`func (o *SleepModel) UnsetLowestHeartRate()`

UnsetLowestHeartRate ensures that no value is present for LowestHeartRate, not even an explicit nil
### GetMovement30Sec

`func (o *SleepModel) GetMovement30Sec() string`

GetMovement30Sec returns the Movement30Sec field if non-nil, zero value otherwise.

### GetMovement30SecOk

`func (o *SleepModel) GetMovement30SecOk() (*string, bool)`

GetMovement30SecOk returns a tuple with the Movement30Sec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMovement30Sec

`func (o *SleepModel) SetMovement30Sec(v string)`

SetMovement30Sec sets Movement30Sec field to given value.


### SetMovement30SecNil

`func (o *SleepModel) SetMovement30SecNil(b bool)`

 SetMovement30SecNil sets the value for Movement30Sec to be an explicit nil

### UnsetMovement30Sec
`func (o *SleepModel) UnsetMovement30Sec()`

UnsetMovement30Sec ensures that no value is present for Movement30Sec, not even an explicit nil
### GetPeriod

`func (o *SleepModel) GetPeriod() int32`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *SleepModel) GetPeriodOk() (*int32, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *SleepModel) SetPeriod(v int32)`

SetPeriod sets Period field to given value.


### GetReadiness

`func (o *SleepModel) GetReadiness() ReadinessSummary`

GetReadiness returns the Readiness field if non-nil, zero value otherwise.

### GetReadinessOk

`func (o *SleepModel) GetReadinessOk() (*ReadinessSummary, bool)`

GetReadinessOk returns a tuple with the Readiness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadiness

`func (o *SleepModel) SetReadiness(v ReadinessSummary)`

SetReadiness sets Readiness field to given value.


### SetReadinessNil

`func (o *SleepModel) SetReadinessNil(b bool)`

 SetReadinessNil sets the value for Readiness to be an explicit nil

### UnsetReadiness
`func (o *SleepModel) UnsetReadiness()`

UnsetReadiness ensures that no value is present for Readiness, not even an explicit nil
### GetReadinessScoreDelta

`func (o *SleepModel) GetReadinessScoreDelta() int32`

GetReadinessScoreDelta returns the ReadinessScoreDelta field if non-nil, zero value otherwise.

### GetReadinessScoreDeltaOk

`func (o *SleepModel) GetReadinessScoreDeltaOk() (*int32, bool)`

GetReadinessScoreDeltaOk returns a tuple with the ReadinessScoreDelta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadinessScoreDelta

`func (o *SleepModel) SetReadinessScoreDelta(v int32)`

SetReadinessScoreDelta sets ReadinessScoreDelta field to given value.


### SetReadinessScoreDeltaNil

`func (o *SleepModel) SetReadinessScoreDeltaNil(b bool)`

 SetReadinessScoreDeltaNil sets the value for ReadinessScoreDelta to be an explicit nil

### UnsetReadinessScoreDelta
`func (o *SleepModel) UnsetReadinessScoreDelta()`

UnsetReadinessScoreDelta ensures that no value is present for ReadinessScoreDelta, not even an explicit nil
### GetRemSleepDuration

`func (o *SleepModel) GetRemSleepDuration() int32`

GetRemSleepDuration returns the RemSleepDuration field if non-nil, zero value otherwise.

### GetRemSleepDurationOk

`func (o *SleepModel) GetRemSleepDurationOk() (*int32, bool)`

GetRemSleepDurationOk returns a tuple with the RemSleepDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemSleepDuration

`func (o *SleepModel) SetRemSleepDuration(v int32)`

SetRemSleepDuration sets RemSleepDuration field to given value.


### SetRemSleepDurationNil

`func (o *SleepModel) SetRemSleepDurationNil(b bool)`

 SetRemSleepDurationNil sets the value for RemSleepDuration to be an explicit nil

### UnsetRemSleepDuration
`func (o *SleepModel) UnsetRemSleepDuration()`

UnsetRemSleepDuration ensures that no value is present for RemSleepDuration, not even an explicit nil
### GetRestlessPeriods

`func (o *SleepModel) GetRestlessPeriods() int32`

GetRestlessPeriods returns the RestlessPeriods field if non-nil, zero value otherwise.

### GetRestlessPeriodsOk

`func (o *SleepModel) GetRestlessPeriodsOk() (*int32, bool)`

GetRestlessPeriodsOk returns a tuple with the RestlessPeriods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestlessPeriods

`func (o *SleepModel) SetRestlessPeriods(v int32)`

SetRestlessPeriods sets RestlessPeriods field to given value.


### SetRestlessPeriodsNil

`func (o *SleepModel) SetRestlessPeriodsNil(b bool)`

 SetRestlessPeriodsNil sets the value for RestlessPeriods to be an explicit nil

### UnsetRestlessPeriods
`func (o *SleepModel) UnsetRestlessPeriods()`

UnsetRestlessPeriods ensures that no value is present for RestlessPeriods, not even an explicit nil
### GetSleepPhase5Min

`func (o *SleepModel) GetSleepPhase5Min() string`

GetSleepPhase5Min returns the SleepPhase5Min field if non-nil, zero value otherwise.

### GetSleepPhase5MinOk

`func (o *SleepModel) GetSleepPhase5MinOk() (*string, bool)`

GetSleepPhase5MinOk returns a tuple with the SleepPhase5Min field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSleepPhase5Min

`func (o *SleepModel) SetSleepPhase5Min(v string)`

SetSleepPhase5Min sets SleepPhase5Min field to given value.


### SetSleepPhase5MinNil

`func (o *SleepModel) SetSleepPhase5MinNil(b bool)`

 SetSleepPhase5MinNil sets the value for SleepPhase5Min to be an explicit nil

### UnsetSleepPhase5Min
`func (o *SleepModel) UnsetSleepPhase5Min()`

UnsetSleepPhase5Min ensures that no value is present for SleepPhase5Min, not even an explicit nil
### GetSleepScoreDelta

`func (o *SleepModel) GetSleepScoreDelta() int32`

GetSleepScoreDelta returns the SleepScoreDelta field if non-nil, zero value otherwise.

### GetSleepScoreDeltaOk

`func (o *SleepModel) GetSleepScoreDeltaOk() (*int32, bool)`

GetSleepScoreDeltaOk returns a tuple with the SleepScoreDelta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSleepScoreDelta

`func (o *SleepModel) SetSleepScoreDelta(v int32)`

SetSleepScoreDelta sets SleepScoreDelta field to given value.


### SetSleepScoreDeltaNil

`func (o *SleepModel) SetSleepScoreDeltaNil(b bool)`

 SetSleepScoreDeltaNil sets the value for SleepScoreDelta to be an explicit nil

### UnsetSleepScoreDelta
`func (o *SleepModel) UnsetSleepScoreDelta()`

UnsetSleepScoreDelta ensures that no value is present for SleepScoreDelta, not even an explicit nil
### GetSleepAlgorithmVersion

`func (o *SleepModel) GetSleepAlgorithmVersion() SleepAlgorithmVersion`

GetSleepAlgorithmVersion returns the SleepAlgorithmVersion field if non-nil, zero value otherwise.

### GetSleepAlgorithmVersionOk

`func (o *SleepModel) GetSleepAlgorithmVersionOk() (*SleepAlgorithmVersion, bool)`

GetSleepAlgorithmVersionOk returns a tuple with the SleepAlgorithmVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSleepAlgorithmVersion

`func (o *SleepModel) SetSleepAlgorithmVersion(v SleepAlgorithmVersion)`

SetSleepAlgorithmVersion sets SleepAlgorithmVersion field to given value.


### SetSleepAlgorithmVersionNil

`func (o *SleepModel) SetSleepAlgorithmVersionNil(b bool)`

 SetSleepAlgorithmVersionNil sets the value for SleepAlgorithmVersion to be an explicit nil

### UnsetSleepAlgorithmVersion
`func (o *SleepModel) UnsetSleepAlgorithmVersion()`

UnsetSleepAlgorithmVersion ensures that no value is present for SleepAlgorithmVersion, not even an explicit nil
### GetTimeInBed

`func (o *SleepModel) GetTimeInBed() int32`

GetTimeInBed returns the TimeInBed field if non-nil, zero value otherwise.

### GetTimeInBedOk

`func (o *SleepModel) GetTimeInBedOk() (*int32, bool)`

GetTimeInBedOk returns a tuple with the TimeInBed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeInBed

`func (o *SleepModel) SetTimeInBed(v int32)`

SetTimeInBed sets TimeInBed field to given value.


### GetTotalSleepDuration

`func (o *SleepModel) GetTotalSleepDuration() int32`

GetTotalSleepDuration returns the TotalSleepDuration field if non-nil, zero value otherwise.

### GetTotalSleepDurationOk

`func (o *SleepModel) GetTotalSleepDurationOk() (*int32, bool)`

GetTotalSleepDurationOk returns a tuple with the TotalSleepDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSleepDuration

`func (o *SleepModel) SetTotalSleepDuration(v int32)`

SetTotalSleepDuration sets TotalSleepDuration field to given value.


### SetTotalSleepDurationNil

`func (o *SleepModel) SetTotalSleepDurationNil(b bool)`

 SetTotalSleepDurationNil sets the value for TotalSleepDuration to be an explicit nil

### UnsetTotalSleepDuration
`func (o *SleepModel) UnsetTotalSleepDuration()`

UnsetTotalSleepDuration ensures that no value is present for TotalSleepDuration, not even an explicit nil
### GetType

`func (o *SleepModel) GetType() SleepType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SleepModel) GetTypeOk() (*SleepType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SleepModel) SetType(v SleepType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


