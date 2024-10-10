# DailyResilienceModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Day** | **string** | Day when the resilience record was recorded. | 
**Contributors** | [**ResilienceContributors**](ResilienceContributors.md) | Contributors to the resilience score. | 
**Level** | [**LongTermResilienceLevel**](LongTermResilienceLevel.md) | Resilience level. | 

## Methods

### NewDailyResilienceModel

`func NewDailyResilienceModel(id string, day string, contributors ResilienceContributors, level LongTermResilienceLevel, ) *DailyResilienceModel`

NewDailyResilienceModel instantiates a new DailyResilienceModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDailyResilienceModelWithDefaults

`func NewDailyResilienceModelWithDefaults() *DailyResilienceModel`

NewDailyResilienceModelWithDefaults instantiates a new DailyResilienceModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DailyResilienceModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DailyResilienceModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DailyResilienceModel) SetId(v string)`

SetId sets Id field to given value.


### GetDay

`func (o *DailyResilienceModel) GetDay() string`

GetDay returns the Day field if non-nil, zero value otherwise.

### GetDayOk

`func (o *DailyResilienceModel) GetDayOk() (*string, bool)`

GetDayOk returns a tuple with the Day field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDay

`func (o *DailyResilienceModel) SetDay(v string)`

SetDay sets Day field to given value.


### GetContributors

`func (o *DailyResilienceModel) GetContributors() ResilienceContributors`

GetContributors returns the Contributors field if non-nil, zero value otherwise.

### GetContributorsOk

`func (o *DailyResilienceModel) GetContributorsOk() (*ResilienceContributors, bool)`

GetContributorsOk returns a tuple with the Contributors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContributors

`func (o *DailyResilienceModel) SetContributors(v ResilienceContributors)`

SetContributors sets Contributors field to given value.


### GetLevel

`func (o *DailyResilienceModel) GetLevel() LongTermResilienceLevel`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *DailyResilienceModel) GetLevelOk() (*LongTermResilienceLevel, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *DailyResilienceModel) SetLevel(v LongTermResilienceLevel)`

SetLevel sets Level field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


