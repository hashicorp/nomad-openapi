package api

import (
	"fmt"
	"time"

	"github.com/hashicorp/cronexpr"
	openapi "github.com/hashicorp/nomad-openapi/v1/client"
)

// Jobs encapsulates and extends the generated JobsApiService with convenience methods.
type Jobs struct {
	client *Client
	openapi.JobsApiService
}

func (c *Client) Jobs() *Jobs {
	return &Jobs{client: c}
}

// PlanOptions is used to pass through job planning parameters
type PlanOptions struct {
	Diff           bool
	PolicyOverride bool
}

func (j *Jobs) Plan(job *openapi.Job, diff bool) (*openapi.JobPlanResponse, *WriteMeta, error) {
	opts := PlanOptions{Diff: diff}
	return j.PlanOpts(job, &opts)
}

// PlanOpts returns a JobPlanResponse and the current Index or Error.
func (j *Jobs) PlanOpts(job *openapi.Job, opts *PlanOptions) (*openapi.JobPlanResponse, *WriteMeta, error) {
	if job == nil {
		return nil, nil, fmt.Errorf("must pass non-nil job")
	}

	writeMeta := &WriteMeta{}
	requestBody := *openapi.NewJobPlanRequest()
	requestBody.SetJob(*job)
	if opts != nil {
		requestBody.SetDiff(opts.Diff)
		requestBody.SetPolicyOverride(opts.PolicyOverride)
	}

	request := j.client.oapiClient.JobsApi.JobJobNamePlanPost(j.client.Ctx, *job.ID)
	request = request.JobPlanRequest(requestBody)

	result, response, err := request.Execute()
	err = parseWriteMeta(response, writeMeta)
	if err != nil {
		return nil, writeMeta, err
	}

	return &result, writeMeta, err
}

// RegisterOptions is used to pass through job registration parameters
type RegisterOptions struct {
	EnforceIndex   bool
	ModifyIndex    uint64
	PolicyOverride bool
	PreserveCounts bool
}

// EnforceRegister is used to register a job enforcing its job modify index.
func (j *Jobs) EnforceRegister(job *openapi.Job, modifyIndex uint64) (*openapi.JobRegisterResponse, *WriteMeta, error) {
	registerOpts := RegisterOptions{EnforceIndex: true, ModifyIndex: modifyIndex}
	return j.Register(job, &registerOpts)
}

func (j *Jobs) Register(job *openapi.Job, registerOpts *RegisterOptions) (*openapi.JobRegisterResponse, *WriteMeta, error) {
	if job == nil {
		return nil, nil, fmt.Errorf("must pass non-nil job")
	}

	// Format the request
	request := j.client.oapiClient.JobsApi.JobsPost(j.client.Ctx)
	registerRequest := openapi.NewJobRegisterRequest()
	// TODO: find a way to force this - too easy to overlook
	j.client.setWriteOptions(registerRequest.SetRegion, registerRequest.SetNamespace, registerRequest.SetSecretID, nil)
	registerRequest.SetJob(*job)

	if registerOpts != nil {
		if registerOpts.EnforceIndex {
			registerRequest.SetEnforceIndex(true)
			registerRequest.SetJobModifyIndex(int32(registerOpts.ModifyIndex))
		}
		registerRequest.SetPolicyOverride(registerOpts.PolicyOverride)
		registerRequest.SetPreserveCounts(registerOpts.PreserveCounts)
	}

	result, response, err := j.client.oapiClient.JobsApi.JobsPostExecute(request.JobRegisterRequest(*registerRequest))
	if err != nil {
		return nil, nil, err
	}

	writeMeta := &WriteMeta{}

	err = parseWriteMeta(response, writeMeta)
	if err != nil {
		return nil, writeMeta, err
	}

	return &result, writeMeta, nil
}

func (j *Jobs) IsMultiRegion(job *openapi.Job) bool {
	return job.Multiregion != nil && *job.Multiregion.Regions != nil && len(*job.Multiregion.Regions) > 0
}

func (j *Jobs) IsPeriodic(job *openapi.Job) bool {
	return job.Periodic != nil
}

func (j *Jobs) IsParameterized(job *openapi.Job) bool {
	dispatched := false
	if job.Dispatched != nil {
		dispatched = *job.Dispatched
	}
	return job.ParameterizedJob != nil && !dispatched
}

func (j *Jobs) GetLocation(job *openapi.Job) (*time.Location, error) {
	periodicConfig := *job.Periodic
	if job.Periodic == nil || periodicConfig.TimeZone == nil || *periodicConfig.TimeZone == "" {
		return time.UTC, nil
	}

	return time.LoadLocation(*periodicConfig.TimeZone)
}

// Next returns the closest time instant matching the spec that is after the
// passed time. If no matching instance exists, the zero value of time.Time is
// returned. The `time.Location` of the returned value matches that of the
// passed time.
func (j *Jobs) Next(p *openapi.PeriodicConfig, fromTime time.Time) (time.Time, error) {
	if *p.SpecType == PeriodicSpecCron {
		e, err := cronexpr.Parse(*p.Spec)
		if err != nil {
			return time.Time{}, fmt.Errorf("failed parsing cron expression %q: %v", *p.Spec, err)
		}
		return cronParseNext(e, fromTime, *p.Spec)
	}

	return time.Time{}, nil
}

// cronParseNext is a helper that parses the next time for the given expression
// but captures any panic that may occur in the underlying library.
// ---  THIS FUNCTION IS REPLICATED IN nomad/structs/structs.go
// and should be kept in sync.
func cronParseNext(e *cronexpr.Expression, fromTime time.Time, spec string) (t time.Time, err error) {
	defer func() {
		if recover() != nil {
			t = time.Time{}
			err = fmt.Errorf("failed parsing cron expression: %q", spec)
		}
	}()

	return e.Next(fromTime), nil
}
