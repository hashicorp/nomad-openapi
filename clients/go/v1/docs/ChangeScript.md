# ChangeScript

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Args** | Pointer to **[]string** |  | [optional] 
**Command** | Pointer to **string** |  | [optional] 
**FailOnError** | Pointer to **bool** |  | [optional] 
**Timeout** | Pointer to **int64** |  | [optional] 

## Methods

### NewChangeScript

`func NewChangeScript() *ChangeScript`

NewChangeScript instantiates a new ChangeScript object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChangeScriptWithDefaults

`func NewChangeScriptWithDefaults() *ChangeScript`

NewChangeScriptWithDefaults instantiates a new ChangeScript object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArgs

`func (o *ChangeScript) GetArgs() []string`

GetArgs returns the Args field if non-nil, zero value otherwise.

### GetArgsOk

`func (o *ChangeScript) GetArgsOk() (*[]string, bool)`

GetArgsOk returns a tuple with the Args field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArgs

`func (o *ChangeScript) SetArgs(v []string)`

SetArgs sets Args field to given value.

### HasArgs

`func (o *ChangeScript) HasArgs() bool`

HasArgs returns a boolean if a field has been set.

### GetCommand

`func (o *ChangeScript) GetCommand() string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *ChangeScript) GetCommandOk() (*string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *ChangeScript) SetCommand(v string)`

SetCommand sets Command field to given value.

### HasCommand

`func (o *ChangeScript) HasCommand() bool`

HasCommand returns a boolean if a field has been set.

### GetFailOnError

`func (o *ChangeScript) GetFailOnError() bool`

GetFailOnError returns the FailOnError field if non-nil, zero value otherwise.

### GetFailOnErrorOk

`func (o *ChangeScript) GetFailOnErrorOk() (*bool, bool)`

GetFailOnErrorOk returns a tuple with the FailOnError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailOnError

`func (o *ChangeScript) SetFailOnError(v bool)`

SetFailOnError sets FailOnError field to given value.

### HasFailOnError

`func (o *ChangeScript) HasFailOnError() bool`

HasFailOnError returns a boolean if a field has been set.

### GetTimeout

`func (o *ChangeScript) GetTimeout() int64`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *ChangeScript) GetTimeoutOk() (*int64, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *ChangeScript) SetTimeout(v int64)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *ChangeScript) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


