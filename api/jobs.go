package api

import (
	"errors"
	"fmt"
	"time"

	"github.com/hashicorp/cronexpr"
	openapi "github.com/hashicorp/nomad-openapi/v1/client"
)

// Jobs encapsulates and extends the generated JobsApiService with convenience methods.
type Jobs struct {
	client *Client
}

func (c *Client) Jobs() *Jobs {
	return &Jobs{client: c}
}

func (j *Jobs) JobsApi() *openapi.JobsApiService {
	return j.client.oapiClient.JobsApi
}

func (j *Jobs) GetJobs() ([]openapi.JobListStub, *QueryMeta, error) {
	jobsApi := j.JobsApi()
	request := jobsApi.GetJobs(j.client.Ctx)
	// TODO: Find a way to make this automatic
	request = j.client.setQueryOptions(request).(openapi.ApiGetJobsRequest)

	response, apiResponse, err := jobsApi.GetJobsExecute(request)
	if err != nil {
		return nil, nil, err
	}

	meta, err := parseQueryMeta(apiResponse)
	if err != nil {
		return nil, nil, err
	}

	return response, meta, nil
}

// PlanOpts is used to pass through job planning parameters
type PlanOpts struct {
	Diff           bool
	PolicyOverride bool
}

func (j *Jobs) Plan(job *openapi.Job, diff bool) (*openapi.JobPlanResponse, *WriteMeta, error) {
	opts := PlanOpts{Diff: diff}
	return j.PlanOpts(job, &opts)
}

// PlanOpts returns a JobPlanResponse and the current Index or Error.
func (j *Jobs) PlanOpts(job *openapi.Job, opts *PlanOpts) (*openapi.JobPlanResponse, *WriteMeta, error) {
	jobsApi := j.JobsApi()
	requestBody := *openapi.NewJobPlanRequest()
	requestBody.SetJob(*job)
	if opts != nil {
		requestBody.SetDiff(opts.Diff)
		requestBody.SetPolicyOverride(opts.PolicyOverride)
	}

	request := jobsApi.PostJobPlan(j.client.Ctx, *job.ID)
	request = request.JobPlanRequest(requestBody)
	// TODO: Figure out why this name is so wonky
	request = j.client.setWriteOptions(request).(openapi.ApiPostJobPlanRequest)

	result, response, err := request.Execute()
	meta, err := parseWriteMeta(response)
	if err != nil {
		return nil, nil, err
	}

	return &result, meta, err
}

// RegisterOpts is used to pass through job registration parameters
type RegisterOpts struct {
	EnforceIndex   bool
	ModifyIndex    uint64
	PolicyOverride bool
	PreserveCounts bool
}

// EnforceRegister is used to register a job enforcing its job modify index.
func (j *Jobs) EnforceRegister(job *openapi.Job, modifyIndex uint64) (*openapi.JobRegisterResponse, *WriteMeta, error) {
	registerOpts := RegisterOpts{EnforceIndex: true, ModifyIndex: modifyIndex}
	return j.Register(job, &registerOpts)
}

func (j *Jobs) Register(job *openapi.Job, registerOpts *RegisterOpts) (*openapi.JobRegisterResponse, *WriteMeta, error) {
	if job == nil {
		return nil, nil, fmt.Errorf("must pass non-nil job")
	}

	jobsApi := j.JobsApi()
	// Format the request j.client.Ctx
	request := jobsApi.PostJob(j.client.Ctx)
	registerRequest := openapi.NewJobRegisterRequest()

	registerRequest.SetJob(*job)

	if j.client.config.Region != "" {
		registerRequest.SetRegion(j.client.config.Region)
	}

	if j.client.config.Namespace != "" {
		registerRequest.SetNamespace(j.client.config.Namespace)
	}

	if j.client.config.WriteOpts.AuthToken != "" {
		registerRequest.SetSecretID(j.client.config.WriteOpts.AuthToken)
	}

	if registerOpts != nil {
		if registerOpts.EnforceIndex {
			registerRequest.SetEnforceIndex(true)
			registerRequest.SetJobModifyIndex(int32(registerOpts.ModifyIndex))
		}
		registerRequest.SetPolicyOverride(registerOpts.PolicyOverride)
		registerRequest.SetPreserveCounts(registerOpts.PreserveCounts)
	}

	result, response, err := j.client.oapiClient.JobsApi.PostJobExecute(request.JobRegisterRequest(*registerRequest))
	if err != nil {
		return nil, nil, err
	}

	meta, err := parseWriteMeta(response)
	if err != nil {
		return nil, nil, err
	}

	return &result, meta, nil
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

func (j *Jobs) Post(job *openapi.Job) (*openapi.JobRegisterResponse, *WriteMeta, error) {
	return j.Register(job, nil)
}

func (j *Jobs) Delete(jobName string, purge, global bool) (*openapi.JobDeregisterResponse, *WriteMeta, error) {
	if jobName == "" {
		return nil, nil, errors.New("jobName must be specified")
	}
	jobsApi := j.JobsApi()
	request := jobsApi.DeleteJob(j.client.Ctx, jobName)
	request = j.client.setWriteOptions(request).(openapi.ApiDeleteJobRequest).Purge(purge).Global(global)

	result, response, err := jobsApi.DeleteJobExecute(request)
	if err != nil {
		return nil, nil, err
	}

	meta, err := parseWriteMeta(response)
	if err != nil {
		return nil, nil, err
	}

	return &result, meta, nil
}
