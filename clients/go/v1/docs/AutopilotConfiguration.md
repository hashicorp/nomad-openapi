# AutopilotConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CleanupDeadServers** | Pointer to **bool** |  | [optional] 
**CreateIndex** | Pointer to **int32** |  | [optional] 
**DisableUpgradeMigration** | Pointer to **bool** |  | [optional] 
**EnableCustomUpgrades** | Pointer to **bool** |  | [optional] 
**EnableRedundancyZones** | Pointer to **bool** |  | [optional] 
**LastContactThreshold** | Pointer to **string** |  | [optional] 
**MaxTrailingLogs** | Pointer to **int32** |  | [optional] 
**MinQuorum** | Pointer to **int32** |  | [optional] 
**ModifyIndex** | Pointer to **int32** |  | [optional] 
**ServerStabilizationTime** | Pointer to **string** |  | [optional] 

## Methods

### NewAutopilotConfiguration

`func NewAutopilotConfiguration() *AutopilotConfiguration`

NewAutopilotConfiguration instantiates a new AutopilotConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutopilotConfigurationWithDefaults

`func NewAutopilotConfigurationWithDefaults() *AutopilotConfiguration`

NewAutopilotConfigurationWithDefaults instantiates a new AutopilotConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCleanupDeadServers

`func (o *AutopilotConfiguration) GetCleanupDeadServers() bool`

GetCleanupDeadServers returns the CleanupDeadServers field if non-nil, zero value otherwise.

### GetCleanupDeadServersOk

`func (o *AutopilotConfiguration) GetCleanupDeadServersOk() (*bool, bool)`

GetCleanupDeadServersOk returns a tuple with the CleanupDeadServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCleanupDeadServers

`func (o *AutopilotConfiguration) SetCleanupDeadServers(v bool)`

SetCleanupDeadServers sets CleanupDeadServers field to given value.

### HasCleanupDeadServers

`func (o *AutopilotConfiguration) HasCleanupDeadServers() bool`

HasCleanupDeadServers returns a boolean if a field has been set.

### GetCreateIndex

`func (o *AutopilotConfiguration) GetCreateIndex() int32`

GetCreateIndex returns the CreateIndex field if non-nil, zero value otherwise.

### GetCreateIndexOk

`func (o *AutopilotConfiguration) GetCreateIndexOk() (*int32, bool)`

GetCreateIndexOk returns a tuple with the CreateIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateIndex

`func (o *AutopilotConfiguration) SetCreateIndex(v int32)`

SetCreateIndex sets CreateIndex field to given value.

### HasCreateIndex

`func (o *AutopilotConfiguration) HasCreateIndex() bool`

HasCreateIndex returns a boolean if a field has been set.

### GetDisableUpgradeMigration

`func (o *AutopilotConfiguration) GetDisableUpgradeMigration() bool`

GetDisableUpgradeMigration returns the DisableUpgradeMigration field if non-nil, zero value otherwise.

### GetDisableUpgradeMigrationOk

`func (o *AutopilotConfiguration) GetDisableUpgradeMigrationOk() (*bool, bool)`

GetDisableUpgradeMigrationOk returns a tuple with the DisableUpgradeMigration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableUpgradeMigration

`func (o *AutopilotConfiguration) SetDisableUpgradeMigration(v bool)`

SetDisableUpgradeMigration sets DisableUpgradeMigration field to given value.

### HasDisableUpgradeMigration

`func (o *AutopilotConfiguration) HasDisableUpgradeMigration() bool`

HasDisableUpgradeMigration returns a boolean if a field has been set.

### GetEnableCustomUpgrades

`func (o *AutopilotConfiguration) GetEnableCustomUpgrades() bool`

GetEnableCustomUpgrades returns the EnableCustomUpgrades field if non-nil, zero value otherwise.

### GetEnableCustomUpgradesOk

`func (o *AutopilotConfiguration) GetEnableCustomUpgradesOk() (*bool, bool)`

GetEnableCustomUpgradesOk returns a tuple with the EnableCustomUpgrades field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableCustomUpgrades

`func (o *AutopilotConfiguration) SetEnableCustomUpgrades(v bool)`

SetEnableCustomUpgrades sets EnableCustomUpgrades field to given value.

### HasEnableCustomUpgrades

`func (o *AutopilotConfiguration) HasEnableCustomUpgrades() bool`

HasEnableCustomUpgrades returns a boolean if a field has been set.

### GetEnableRedundancyZones

`func (o *AutopilotConfiguration) GetEnableRedundancyZones() bool`

GetEnableRedundancyZones returns the EnableRedundancyZones field if non-nil, zero value otherwise.

### GetEnableRedundancyZonesOk

`func (o *AutopilotConfiguration) GetEnableRedundancyZonesOk() (*bool, bool)`

GetEnableRedundancyZonesOk returns a tuple with the EnableRedundancyZones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableRedundancyZones

`func (o *AutopilotConfiguration) SetEnableRedundancyZones(v bool)`

SetEnableRedundancyZones sets EnableRedundancyZones field to given value.

### HasEnableRedundancyZones

`func (o *AutopilotConfiguration) HasEnableRedundancyZones() bool`

HasEnableRedundancyZones returns a boolean if a field has been set.

### GetLastContactThreshold

`func (o *AutopilotConfiguration) GetLastContactThreshold() string`

GetLastContactThreshold returns the LastContactThreshold field if non-nil, zero value otherwise.

### GetLastContactThresholdOk

`func (o *AutopilotConfiguration) GetLastContactThresholdOk() (*string, bool)`

GetLastContactThresholdOk returns a tuple with the LastContactThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastContactThreshold

`func (o *AutopilotConfiguration) SetLastContactThreshold(v string)`

SetLastContactThreshold sets LastContactThreshold field to given value.

### HasLastContactThreshold

`func (o *AutopilotConfiguration) HasLastContactThreshold() bool`

HasLastContactThreshold returns a boolean if a field has been set.

### GetMaxTrailingLogs

`func (o *AutopilotConfiguration) GetMaxTrailingLogs() int32`

GetMaxTrailingLogs returns the MaxTrailingLogs field if non-nil, zero value otherwise.

### GetMaxTrailingLogsOk

`func (o *AutopilotConfiguration) GetMaxTrailingLogsOk() (*int32, bool)`

GetMaxTrailingLogsOk returns a tuple with the MaxTrailingLogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTrailingLogs

`func (o *AutopilotConfiguration) SetMaxTrailingLogs(v int32)`

SetMaxTrailingLogs sets MaxTrailingLogs field to given value.

### HasMaxTrailingLogs

`func (o *AutopilotConfiguration) HasMaxTrailingLogs() bool`

HasMaxTrailingLogs returns a boolean if a field has been set.

### GetMinQuorum

`func (o *AutopilotConfiguration) GetMinQuorum() int32`

GetMinQuorum returns the MinQuorum field if non-nil, zero value otherwise.

### GetMinQuorumOk

`func (o *AutopilotConfiguration) GetMinQuorumOk() (*int32, bool)`

GetMinQuorumOk returns a tuple with the MinQuorum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinQuorum

`func (o *AutopilotConfiguration) SetMinQuorum(v int32)`

SetMinQuorum sets MinQuorum field to given value.

### HasMinQuorum

`func (o *AutopilotConfiguration) HasMinQuorum() bool`

HasMinQuorum returns a boolean if a field has been set.

### GetModifyIndex

`func (o *AutopilotConfiguration) GetModifyIndex() int32`

GetModifyIndex returns the ModifyIndex field if non-nil, zero value otherwise.

### GetModifyIndexOk

`func (o *AutopilotConfiguration) GetModifyIndexOk() (*int32, bool)`

GetModifyIndexOk returns a tuple with the ModifyIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifyIndex

`func (o *AutopilotConfiguration) SetModifyIndex(v int32)`

SetModifyIndex sets ModifyIndex field to given value.

### HasModifyIndex

`func (o *AutopilotConfiguration) HasModifyIndex() bool`

HasModifyIndex returns a boolean if a field has been set.

### GetServerStabilizationTime

`func (o *AutopilotConfiguration) GetServerStabilizationTime() string`

GetServerStabilizationTime returns the ServerStabilizationTime field if non-nil, zero value otherwise.

### GetServerStabilizationTimeOk

`func (o *AutopilotConfiguration) GetServerStabilizationTimeOk() (*string, bool)`

GetServerStabilizationTimeOk returns a tuple with the ServerStabilizationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerStabilizationTime

`func (o *AutopilotConfiguration) SetServerStabilizationTime(v string)`

SetServerStabilizationTime sets ServerStabilizationTime field to given value.

### HasServerStabilizationTime

`func (o *AutopilotConfiguration) HasServerStabilizationTime() bool`

HasServerStabilizationTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


