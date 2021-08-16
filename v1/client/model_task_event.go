/*
 * Nomad
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.1.3
 * Contact: support@hashicorp.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// TaskEvent struct for TaskEvent
type TaskEvent struct {
	Details *map[string]string `json:"Details,omitempty"`
	DiskLimit *int64 `json:"DiskLimit,omitempty"`
	DiskSize *int64 `json:"DiskSize,omitempty"`
	DisplayMessage *string `json:"DisplayMessage,omitempty"`
	DownloadError *string `json:"DownloadError,omitempty"`
	DriverError *string `json:"DriverError,omitempty"`
	DriverMessage *string `json:"DriverMessage,omitempty"`
	ExitCode *int32 `json:"ExitCode,omitempty"`
	FailedSibling *string `json:"FailedSibling,omitempty"`
	FailsTask *bool `json:"FailsTask,omitempty"`
	GenericSource *string `json:"GenericSource,omitempty"`
	KillError *string `json:"KillError,omitempty"`
	KillReason *string `json:"KillReason,omitempty"`
	KillTimeout *int64 `json:"KillTimeout,omitempty"`
	Message *string `json:"Message,omitempty"`
	RestartReason *string `json:"RestartReason,omitempty"`
	SetupError *string `json:"SetupError,omitempty"`
	Signal *int32 `json:"Signal,omitempty"`
	StartDelay *int64 `json:"StartDelay,omitempty"`
	TaskSignal *string `json:"TaskSignal,omitempty"`
	TaskSignalReason *string `json:"TaskSignalReason,omitempty"`
	Time *int64 `json:"Time,omitempty"`
	Type *string `json:"Type,omitempty"`
	ValidationError *string `json:"ValidationError,omitempty"`
	VaultError *string `json:"VaultError,omitempty"`
}

// NewTaskEvent instantiates a new TaskEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaskEvent() *TaskEvent {
	this := TaskEvent{}
	return &this
}

// NewTaskEventWithDefaults instantiates a new TaskEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaskEventWithDefaults() *TaskEvent {
	this := TaskEvent{}
	return &this
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *TaskEvent) GetDetails() map[string]string {
	if o == nil || o.Details == nil {
		var ret map[string]string
		return ret
	}
	return *o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskEvent) GetDetailsOk() (*map[string]string, bool) {
	if o == nil || o.Details == nil {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *TaskEvent) HasDetails() bool {
	if o != nil && o.Details != nil {
		return true
	}

	return false
}

// SetDetails gets a reference to the given map[string]string and assigns it to the Details field.
func (o *TaskEvent) SetDetails(v map[string]string) {
	o.Details = &v
}

// GetDiskLimit returns the DiskLimit field value if set, zero value otherwise.
func (o *TaskEvent) GetDiskLimit() int64 {
	if o == nil || o.DiskLimit == nil {
		var ret int64
		return ret
	}
	return *o.DiskLimit
}

// GetDiskLimitOk returns a tuple with the DiskLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskEvent) GetDiskLimitOk() (*int64, bool) {
	if o == nil || o.DiskLimit == nil {
		return nil, false
	}
	return o.DiskLimit, true
}

// HasDiskLimit returns a boolean if a field has been set.
func (o *TaskEvent) HasDiskLimit() bool {
	if o != nil && o.DiskLimit != nil {
		return true
	}

	return false
}

// SetDiskLimit gets a reference to the given int64 and assigns it to the DiskLimit field.
func (o *TaskEvent) SetDiskLimit(v int64) {
	o.DiskLimit = &v
}

// GetDiskSize returns the DiskSize field value if set, zero value otherwise.
func (o *TaskEvent) GetDiskSize() int64 {
	if o == nil || o.DiskSize == nil {
		var ret int64
		return ret
	}
	return *o.DiskSize
}

// GetDiskSizeOk returns a tuple with the DiskSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskEvent) GetDiskSizeOk() (*int64, bool) {
	if o == nil || o.DiskSize == nil {
		return nil, false
	}
	return o.DiskSize, true
}

// HasDiskSize returns a boolean if a field has been set.
func (o *TaskEvent) HasDiskSize() bool {
	if o != nil && o.DiskSize != nil {
		return true
	}

	return false
}

// SetDiskSize gets a reference to the given int64 and assigns it to the DiskSize field.
func (o *TaskEvent) SetDiskSize(v int64) {
	o.DiskSize = &v
}

// GetDisplayMessage returns the DisplayMessage field value if set, zero value otherwise.
func (o *TaskEvent) GetDisplayMessage() string {
	if o == nil || o.DisplayMessage == nil {
		var ret string
		return ret
	}
	return *o.DisplayMessage
}

// GetDisplayMessageOk returns a tuple with the DisplayMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskEvent) GetDisplayMessageOk() (*string, bool) {
	if o == nil || o.DisplayMessage == nil {
		return nil, false
	}
	return o.DisplayMessage, true
}

// HasDisplayMessage returns a boolean if a field has been set.
func (o *TaskEvent) HasDisplayMessage() bool {
	if o != nil && o.DisplayMessage != nil {
		return true
	}

	return false
}

// SetDisplayMessage gets a reference to the given string and assigns it to the DisplayMessage field.
func (o *TaskEvent) SetDisplayMessage(v string) {
	o.DisplayMessage = &v
}

// GetDownloadError returns the DownloadError field value if set, zero value otherwise.
func (o *TaskEvent) GetDownloadError() string {
	if o == nil || o.DownloadError == nil {
		var ret string
		return ret
	}
	return *o.DownloadError
}

// GetDownloadErrorOk returns a tuple with the DownloadError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskEvent) GetDownloadErrorOk() (*string, bool) {
	if o == nil || o.DownloadError == nil {
		return nil, false
	}
	return o.DownloadError, true
}

// HasDownloadError returns a boolean if a field has been set.
func (o *TaskEvent) HasDownloadError() bool {
	if o != nil && o.DownloadError != nil {
		return true
	}

	return false
}

// SetDownloadError gets a reference to the given string and assigns it to the DownloadError field.
func (o *TaskEvent) SetDownloadError(v string) {
	o.DownloadError = &v
}

// GetDriverError returns the DriverError field value if set, zero value otherwise.
func (o *TaskEvent) GetDriverError() string {
	if o == nil || o.DriverError == nil {
		var ret string
		return ret
	}
	return *o.DriverError
}

// GetDriverErrorOk returns a tuple with the DriverError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskEvent) GetDriverErrorOk() (*string, bool) {
	if o == nil || o.DriverError == nil {
		return nil, false
	}
	return o.DriverError, true
}

// HasDriverError returns a boolean if a field has been set.
func (o *TaskEvent) HasDriverError() bool {
	if o != nil && o.DriverError != nil {
		return true
	}

	return false
}

// SetDriverError gets a reference to the given string and assigns it to the DriverError field.
func (o *TaskEvent) SetDriverError(v string) {
	o.DriverError = &v
}

// GetDriverMessage returns the DriverMessage field value if set, zero value otherwise.
func (o *TaskEvent) GetDriverMessage() string {
	if o == nil || o.DriverMessage == nil {
		var ret string
		return ret
	}
	return *o.DriverMessage
}

// GetDriverMessageOk returns a tuple with the DriverMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskEvent) GetDriverMessageOk() (*string, bool) {
	if o == nil || o.DriverMessage == nil {
		return nil, false
	}
	return o.DriverMessage, true
}

// HasDriverMessage returns a boolean if a field has been set.
func (o *TaskEvent) HasDriverMessage() bool {
	if o != nil && o.DriverMessage != nil {
		return true
	}

	return false
}

// SetDriverMessage gets a reference to the given string and assigns it to the DriverMessage field.
func (o *TaskEvent) SetDriverMessage(v string) {
	o.DriverMessage = &v
}

// GetExitCode returns the ExitCode field value if set, zero value otherwise.
func (o *TaskEvent) GetExitCode() int32 {
	if o == nil || o.ExitCode == nil {
		var ret int32
		return ret
	}
	return *o.ExitCode
}

// GetExitCodeOk returns a tuple with the ExitCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskEvent) GetExitCodeOk() (*int32, bool) {
	if o == nil || o.ExitCode == nil {
		return nil, false
	}
	return o.ExitCode, true
}

// HasExitCode returns a boolean if a field has been set.
func (o *TaskEvent) HasExitCode() bool {
	if o != nil && o.ExitCode != nil {
		return true
	}

	return false
}

// SetExitCode gets a reference to the given int32 and assigns it to the ExitCode field.
func (o *TaskEvent) SetExitCode(v int32) {
	o.ExitCode = &v
}

// GetFailedSibling returns the FailedSibling field value if set, zero value otherwise.
func (o *TaskEvent) GetFailedSibling() string {
	if o == nil || o.FailedSibling == nil {
		var ret string
		return ret
	}
	return *o.FailedSibling
}

// GetFailedSiblingOk returns a tuple with the FailedSibling field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskEvent) GetFailedSiblingOk() (*string, bool) {
	if o == nil || o.FailedSibling == nil {
		return nil, false
	}
	return o.FailedSibling, true
}

// HasFailedSibling returns a boolean if a field has been set.
func (o *TaskEvent) HasFailedSibling() bool {
	if o != nil && o.FailedSibling != nil {
		return true
	}

	return false
}

// SetFailedSibling gets a reference to the given string and assigns it to the FailedSibling field.
func (o *TaskEvent) SetFailedSibling(v string) {
	o.FailedSibling = &v
}

// GetFailsTask returns the FailsTask field value if set, zero value otherwise.
func (o *TaskEvent) GetFailsTask() bool {
	if o == nil || o.FailsTask == nil {
		var ret bool
		return ret
	}
	return *o.FailsTask
}

// GetFailsTaskOk returns a tuple with the FailsTask field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskEvent) GetFailsTaskOk() (*bool, bool) {
	if o == nil || o.FailsTask == nil {
		return nil, false
	}
	return o.FailsTask, true
}

// HasFailsTask returns a boolean if a field has been set.
func (o *TaskEvent) HasFailsTask() bool {
	if o != nil && o.FailsTask != nil {
		return true
	}

	return false
}

// SetFailsTask gets a reference to the given bool and assigns it to the FailsTask field.
func (o *TaskEvent) SetFailsTask(v bool) {
	o.FailsTask = &v
}

// GetGenericSource returns the GenericSource field value if set, zero value otherwise.
func (o *TaskEvent) GetGenericSource() string {
	if o == nil || o.GenericSource == nil {
		var ret string
		return ret
	}
	return *o.GenericSource
}

// GetGenericSourceOk returns a tuple with the GenericSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskEvent) GetGenericSourceOk() (*string, bool) {
	if o == nil || o.GenericSource == nil {
		return nil, false
	}
	return o.GenericSource, true
}

// HasGenericSource returns a boolean if a field has been set.
func (o *TaskEvent) HasGenericSource() bool {
	if o != nil && o.GenericSource != nil {
		return true
	}

	return false
}

// SetGenericSource gets a reference to the given string and assigns it to the GenericSource field.
func (o *TaskEvent) SetGenericSource(v string) {
	o.GenericSource = &v
}

// GetKillError returns the KillError field value if set, zero value otherwise.
func (o *TaskEvent) GetKillError() string {
	if o == nil || o.KillError == nil {
		var ret string
		return ret
	}
	return *o.KillError
}

// GetKillErrorOk returns a tuple with the KillError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskEvent) GetKillErrorOk() (*string, bool) {
	if o == nil || o.KillError == nil {
		return nil, false
	}
	return o.KillError, true
}

// HasKillError returns a boolean if a field has been set.
func (o *TaskEvent) HasKillError() bool {
	if o != nil && o.KillError != nil {
		return true
	}

	return false
}

// SetKillError gets a reference to the given string and assigns it to the KillError field.
func (o *TaskEvent) SetKillError(v string) {
	o.KillError = &v
}

// GetKillReason returns the KillReason field value if set, zero value otherwise.
func (o *TaskEvent) GetKillReason() string {
	if o == nil || o.KillReason == nil {
		var ret string
		return ret
	}
	return *o.KillReason
}

// GetKillReasonOk returns a tuple with the KillReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskEvent) GetKillReasonOk() (*string, bool) {
	if o == nil || o.KillReason == nil {
		return nil, false
	}
	return o.KillReason, true
}

// HasKillReason returns a boolean if a field has been set.
func (o *TaskEvent) HasKillReason() bool {
	if o != nil && o.KillReason != nil {
		return true
	}

	return false
}

// SetKillReason gets a reference to the given string and assigns it to the KillReason field.
func (o *TaskEvent) SetKillReason(v string) {
	o.KillReason = &v
}

// GetKillTimeout returns the KillTimeout field value if set, zero value otherwise.
func (o *TaskEvent) GetKillTimeout() int64 {
	if o == nil || o.KillTimeout == nil {
		var ret int64
		return ret
	}
	return *o.KillTimeout
}

// GetKillTimeoutOk returns a tuple with the KillTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskEvent) GetKillTimeoutOk() (*int64, bool) {
	if o == nil || o.KillTimeout == nil {
		return nil, false
	}
	return o.KillTimeout, true
}

// HasKillTimeout returns a boolean if a field has been set.
func (o *TaskEvent) HasKillTimeout() bool {
	if o != nil && o.KillTimeout != nil {
		return true
	}

	return false
}

// SetKillTimeout gets a reference to the given int64 and assigns it to the KillTimeout field.
func (o *TaskEvent) SetKillTimeout(v int64) {
	o.KillTimeout = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *TaskEvent) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskEvent) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *TaskEvent) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *TaskEvent) SetMessage(v string) {
	o.Message = &v
}

// GetRestartReason returns the RestartReason field value if set, zero value otherwise.
func (o *TaskEvent) GetRestartReason() string {
	if o == nil || o.RestartReason == nil {
		var ret string
		return ret
	}
	return *o.RestartReason
}

// GetRestartReasonOk returns a tuple with the RestartReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskEvent) GetRestartReasonOk() (*string, bool) {
	if o == nil || o.RestartReason == nil {
		return nil, false
	}
	return o.RestartReason, true
}

// HasRestartReason returns a boolean if a field has been set.
func (o *TaskEvent) HasRestartReason() bool {
	if o != nil && o.RestartReason != nil {
		return true
	}

	return false
}

// SetRestartReason gets a reference to the given string and assigns it to the RestartReason field.
func (o *TaskEvent) SetRestartReason(v string) {
	o.RestartReason = &v
}

// GetSetupError returns the SetupError field value if set, zero value otherwise.
func (o *TaskEvent) GetSetupError() string {
	if o == nil || o.SetupError == nil {
		var ret string
		return ret
	}
	return *o.SetupError
}

// GetSetupErrorOk returns a tuple with the SetupError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskEvent) GetSetupErrorOk() (*string, bool) {
	if o == nil || o.SetupError == nil {
		return nil, false
	}
	return o.SetupError, true
}

// HasSetupError returns a boolean if a field has been set.
func (o *TaskEvent) HasSetupError() bool {
	if o != nil && o.SetupError != nil {
		return true
	}

	return false
}

// SetSetupError gets a reference to the given string and assigns it to the SetupError field.
func (o *TaskEvent) SetSetupError(v string) {
	o.SetupError = &v
}

// GetSignal returns the Signal field value if set, zero value otherwise.
func (o *TaskEvent) GetSignal() int32 {
	if o == nil || o.Signal == nil {
		var ret int32
		return ret
	}
	return *o.Signal
}

// GetSignalOk returns a tuple with the Signal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskEvent) GetSignalOk() (*int32, bool) {
	if o == nil || o.Signal == nil {
		return nil, false
	}
	return o.Signal, true
}

// HasSignal returns a boolean if a field has been set.
func (o *TaskEvent) HasSignal() bool {
	if o != nil && o.Signal != nil {
		return true
	}

	return false
}

// SetSignal gets a reference to the given int32 and assigns it to the Signal field.
func (o *TaskEvent) SetSignal(v int32) {
	o.Signal = &v
}

// GetStartDelay returns the StartDelay field value if set, zero value otherwise.
func (o *TaskEvent) GetStartDelay() int64 {
	if o == nil || o.StartDelay == nil {
		var ret int64
		return ret
	}
	return *o.StartDelay
}

// GetStartDelayOk returns a tuple with the StartDelay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskEvent) GetStartDelayOk() (*int64, bool) {
	if o == nil || o.StartDelay == nil {
		return nil, false
	}
	return o.StartDelay, true
}

// HasStartDelay returns a boolean if a field has been set.
func (o *TaskEvent) HasStartDelay() bool {
	if o != nil && o.StartDelay != nil {
		return true
	}

	return false
}

// SetStartDelay gets a reference to the given int64 and assigns it to the StartDelay field.
func (o *TaskEvent) SetStartDelay(v int64) {
	o.StartDelay = &v
}

// GetTaskSignal returns the TaskSignal field value if set, zero value otherwise.
func (o *TaskEvent) GetTaskSignal() string {
	if o == nil || o.TaskSignal == nil {
		var ret string
		return ret
	}
	return *o.TaskSignal
}

// GetTaskSignalOk returns a tuple with the TaskSignal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskEvent) GetTaskSignalOk() (*string, bool) {
	if o == nil || o.TaskSignal == nil {
		return nil, false
	}
	return o.TaskSignal, true
}

// HasTaskSignal returns a boolean if a field has been set.
func (o *TaskEvent) HasTaskSignal() bool {
	if o != nil && o.TaskSignal != nil {
		return true
	}

	return false
}

// SetTaskSignal gets a reference to the given string and assigns it to the TaskSignal field.
func (o *TaskEvent) SetTaskSignal(v string) {
	o.TaskSignal = &v
}

// GetTaskSignalReason returns the TaskSignalReason field value if set, zero value otherwise.
func (o *TaskEvent) GetTaskSignalReason() string {
	if o == nil || o.TaskSignalReason == nil {
		var ret string
		return ret
	}
	return *o.TaskSignalReason
}

// GetTaskSignalReasonOk returns a tuple with the TaskSignalReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskEvent) GetTaskSignalReasonOk() (*string, bool) {
	if o == nil || o.TaskSignalReason == nil {
		return nil, false
	}
	return o.TaskSignalReason, true
}

// HasTaskSignalReason returns a boolean if a field has been set.
func (o *TaskEvent) HasTaskSignalReason() bool {
	if o != nil && o.TaskSignalReason != nil {
		return true
	}

	return false
}

// SetTaskSignalReason gets a reference to the given string and assigns it to the TaskSignalReason field.
func (o *TaskEvent) SetTaskSignalReason(v string) {
	o.TaskSignalReason = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *TaskEvent) GetTime() int64 {
	if o == nil || o.Time == nil {
		var ret int64
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskEvent) GetTimeOk() (*int64, bool) {
	if o == nil || o.Time == nil {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *TaskEvent) HasTime() bool {
	if o != nil && o.Time != nil {
		return true
	}

	return false
}

// SetTime gets a reference to the given int64 and assigns it to the Time field.
func (o *TaskEvent) SetTime(v int64) {
	o.Time = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *TaskEvent) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskEvent) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *TaskEvent) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *TaskEvent) SetType(v string) {
	o.Type = &v
}

// GetValidationError returns the ValidationError field value if set, zero value otherwise.
func (o *TaskEvent) GetValidationError() string {
	if o == nil || o.ValidationError == nil {
		var ret string
		return ret
	}
	return *o.ValidationError
}

// GetValidationErrorOk returns a tuple with the ValidationError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskEvent) GetValidationErrorOk() (*string, bool) {
	if o == nil || o.ValidationError == nil {
		return nil, false
	}
	return o.ValidationError, true
}

// HasValidationError returns a boolean if a field has been set.
func (o *TaskEvent) HasValidationError() bool {
	if o != nil && o.ValidationError != nil {
		return true
	}

	return false
}

// SetValidationError gets a reference to the given string and assigns it to the ValidationError field.
func (o *TaskEvent) SetValidationError(v string) {
	o.ValidationError = &v
}

// GetVaultError returns the VaultError field value if set, zero value otherwise.
func (o *TaskEvent) GetVaultError() string {
	if o == nil || o.VaultError == nil {
		var ret string
		return ret
	}
	return *o.VaultError
}

// GetVaultErrorOk returns a tuple with the VaultError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskEvent) GetVaultErrorOk() (*string, bool) {
	if o == nil || o.VaultError == nil {
		return nil, false
	}
	return o.VaultError, true
}

// HasVaultError returns a boolean if a field has been set.
func (o *TaskEvent) HasVaultError() bool {
	if o != nil && o.VaultError != nil {
		return true
	}

	return false
}

// SetVaultError gets a reference to the given string and assigns it to the VaultError field.
func (o *TaskEvent) SetVaultError(v string) {
	o.VaultError = &v
}

func (o TaskEvent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Details != nil {
		toSerialize["Details"] = o.Details
	}
	if o.DiskLimit != nil {
		toSerialize["DiskLimit"] = o.DiskLimit
	}
	if o.DiskSize != nil {
		toSerialize["DiskSize"] = o.DiskSize
	}
	if o.DisplayMessage != nil {
		toSerialize["DisplayMessage"] = o.DisplayMessage
	}
	if o.DownloadError != nil {
		toSerialize["DownloadError"] = o.DownloadError
	}
	if o.DriverError != nil {
		toSerialize["DriverError"] = o.DriverError
	}
	if o.DriverMessage != nil {
		toSerialize["DriverMessage"] = o.DriverMessage
	}
	if o.ExitCode != nil {
		toSerialize["ExitCode"] = o.ExitCode
	}
	if o.FailedSibling != nil {
		toSerialize["FailedSibling"] = o.FailedSibling
	}
	if o.FailsTask != nil {
		toSerialize["FailsTask"] = o.FailsTask
	}
	if o.GenericSource != nil {
		toSerialize["GenericSource"] = o.GenericSource
	}
	if o.KillError != nil {
		toSerialize["KillError"] = o.KillError
	}
	if o.KillReason != nil {
		toSerialize["KillReason"] = o.KillReason
	}
	if o.KillTimeout != nil {
		toSerialize["KillTimeout"] = o.KillTimeout
	}
	if o.Message != nil {
		toSerialize["Message"] = o.Message
	}
	if o.RestartReason != nil {
		toSerialize["RestartReason"] = o.RestartReason
	}
	if o.SetupError != nil {
		toSerialize["SetupError"] = o.SetupError
	}
	if o.Signal != nil {
		toSerialize["Signal"] = o.Signal
	}
	if o.StartDelay != nil {
		toSerialize["StartDelay"] = o.StartDelay
	}
	if o.TaskSignal != nil {
		toSerialize["TaskSignal"] = o.TaskSignal
	}
	if o.TaskSignalReason != nil {
		toSerialize["TaskSignalReason"] = o.TaskSignalReason
	}
	if o.Time != nil {
		toSerialize["Time"] = o.Time
	}
	if o.Type != nil {
		toSerialize["Type"] = o.Type
	}
	if o.ValidationError != nil {
		toSerialize["ValidationError"] = o.ValidationError
	}
	if o.VaultError != nil {
		toSerialize["VaultError"] = o.VaultError
	}
	return json.Marshal(toSerialize)
}

type NullableTaskEvent struct {
	value *TaskEvent
	isSet bool
}

func (v NullableTaskEvent) Get() *TaskEvent {
	return v.value
}

func (v *NullableTaskEvent) Set(val *TaskEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableTaskEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableTaskEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaskEvent(val *TaskEvent) *NullableTaskEvent {
	return &NullableTaskEvent{value: val, isSet: true}
}

func (v NullableTaskEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaskEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


