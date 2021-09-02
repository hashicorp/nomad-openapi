package v1

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/hashicorp/cronexpr"
	"github.com/hashicorp/nomad-openapi/v1/client"
)

// Jobs encapsulates and extends the generated JobsApiService with convenience methods.
type Jobs struct {
	client *Client
}

func (c *Client) Jobs() *Jobs {
	return &Jobs{client: c}
}

func (j *Jobs) JobsApi() *client.JobsApiService {
	return j.client.apiClient.JobsApi
}

func (j *Jobs) Delete(ctx context.Context, jobName string, purge, global bool) (*client.JobDeregisterResponse, *WriteMeta, error) {
	if jobName == "" {
		return nil, nil, jobNameErr
	}

	request := j.JobsApi().DeleteJob(j.client.Ctx, jobName)
	request = j.client.setWriteOptions(ctx, request).(client.ApiDeleteJobRequest)
	request = request.Purge(purge).Global(global)

	result, response, err := request.Execute()
	if err != nil {
		return nil, nil, err
	}

	meta, err := parseWriteMeta(response)
	if err != nil {
		return nil, nil, err
	}

	return &result, meta, nil
}

func (j *Jobs) Deployment(ctx context.Context, jobName string) (*client.Deployment, *QueryMeta, error) {
	if jobName == "" {
		return nil, nil, jobNameErr
	}

	request := j.JobsApi().GetJobDeployment(j.client.Ctx, jobName)
	request = j.client.setQueryOptions(ctx, request).(client.ApiGetJobDeploymentRequest)

	result, response, err := request.Execute()
	if err != nil {
		return nil, nil, err
	}

	meta, err := parseQueryMeta(response)
	if err != nil {
		return nil, nil, err
	}

	return &result, meta, nil
}

func (j *Jobs) Deployments(ctx context.Context, jobName string) (*[]client.Deployment, *QueryMeta, error) {
	if jobName == "" {
		return nil, nil, jobNameErr
	}

	request := j.JobsApi().GetJobDeployments(j.client.Ctx, jobName)
	request = j.client.setQueryOptions(ctx, request).(client.ApiGetJobDeploymentsRequest)

	result, response, err := request.Execute()
	if err != nil {
		return nil, nil, err
	}

	meta, err := parseQueryMeta(response)
	if err != nil {
		return nil, nil, err
	}

	return &result, meta, nil
}

func (j *Jobs) Dispatch(ctx context.Context, jobName string, payload string, meta map[string]string) (*client.JobDispatchResponse, *WriteMeta, error) {
	if jobName == "" {
		return nil, nil, jobNameErr
	}

	request := j.JobsApi().PostJobDispatch(j.client.Ctx, jobName)
	request = j.client.setWriteOptions(ctx, request).(client.ApiPostJobDispatchRequest)

	dispatchRequest := client.NewJobDispatchRequest()
	dispatchRequest.SetJobID(jobName)
	dispatchRequest.SetPayload(payload)
	dispatchRequest.SetMeta(meta)

	request = request.JobDispatchRequest(*dispatchRequest)

	result, response, err := request.Execute()
	if err != nil {
		return nil, nil, err
	}

	writeMeta, err := parseWriteMeta(response)
	if err != nil {
		return nil, nil, err
	}

	return &result, writeMeta, nil
}

func (j *Jobs) EnforceRegister(ctx context.Context, job *client.Job, modifyIndex uint64) (*client.JobRegisterResponse, *WriteMeta, error) {
	return j.Register(ctx, job, &RegisterOpts{
		EnforceIndex: true,
		ModifyIndex:  modifyIndex,
	})
}

func (j *Jobs) Evaluate(ctx context.Context, jobName string, forceReschedule bool) (*client.JobRegisterResponse, *WriteMeta, error) {
	if jobName == "" {
		return nil, nil, jobNameErr
	}

	request := j.JobsApi().PostJobEvaluate(j.client.Ctx, jobName)
	request = j.client.setWriteOptions(ctx, request).(client.ApiPostJobEvaluateRequest)

	evalRequest := client.NewJobEvaluateRequest()
	evalRequest.SetJobID(jobName)
	evalRequest.SetEvalOptions(client.EvalOptions{
		ForceReschedule: &forceReschedule,
	})

	request = request.JobEvaluateRequest(*evalRequest)

	result, response, err := request.Execute()
	if err != nil {
		return nil, nil, err
	}

	meta, err := parseWriteMeta(response)
	if err != nil {
		return nil, nil, err
	}

	return &result, meta, nil
}

func (j *Jobs) GetJob(ctx context.Context, jobName string) (*client.Job, *QueryMeta, error) {
	if jobName == "" {
		return nil, nil, jobNameErr
	}

	request := j.JobsApi().GetJob(j.client.Ctx, jobName)
	request = j.client.setQueryOptions(ctx, request).(client.ApiGetJobRequest)

	result, response, err := request.Execute()
	if err != nil {
		return nil, nil, err
	}

	meta, err := parseQueryMeta(response)
	if err != nil {
		return nil, nil, err
	}

	return &result, meta, nil
}

func (j *Jobs) GetJobs(ctx context.Context) ([]client.JobListStub, *QueryMeta, error) {
	request := j.JobsApi().GetJobs(j.client.Ctx)
	// TODO: Find a way to make this automatic
	request = j.client.setQueryOptions(ctx, request).(client.ApiGetJobsRequest)

	result, response, err := request.Execute()
	if err != nil {
		return nil, nil, err
	}

	meta, err := parseQueryMeta(response)
	if err != nil {
		return nil, nil, err
	}

	return result, meta, nil
}

func (j *Jobs) Parse(ctx context.Context, hcl string, canonicalize, hclV1 bool) (*client.Job, error) {
	if hcl == "" {
		return nil, errors.New("job hcl not set")
	}

	request := j.JobsApi().PostJobParse(j.client.Ctx)

	parseRequest := client.NewJobsParseRequest()
	parseRequest.SetJobHCL(hcl)
	parseRequest.SetCanonicalize(canonicalize)
	parseRequest.SetHclv1(hclV1)

	request = request.JobsParseRequest(*parseRequest)

	result, _, err := request.Execute()
	if err != nil {
		return nil, err
	}

	return &result, err
}

type PlanOpts struct {
	Diff           bool
	PolicyOverride bool
}

func (j *Jobs) PeriodicForce(ctx context.Context, jobName string) (*client.PeriodicForceResponse, *WriteMeta, error) {
	if jobName == "" {
		return nil, nil, jobNameErr
	}

	request := j.JobsApi().PostJobPeriodicForce(j.client.Ctx, jobName)
	request = j.client.setWriteOptions(ctx, request).(client.ApiPostJobPeriodicForceRequest)

	result, response, err := request.Execute()
	if err != nil {
		return nil, nil, err
	}

	meta, err := parseWriteMeta(response)
	if err != nil {
		return nil, nil, err
	}

	return &result, meta, nil
}

func (j *Jobs) Plan(ctx context.Context, job *client.Job, diff bool) (*client.JobPlanResponse, *WriteMeta, error) {
	return j.PlanOpts(ctx, job, &PlanOpts{Diff: diff})
}

func (j *Jobs) PlanOpts(ctx context.Context, job *client.Job, opts *PlanOpts) (*client.JobPlanResponse, *WriteMeta, error) {
	requestBody := *client.NewJobPlanRequest()
	requestBody.SetJob(*job)
	if opts != nil {
		requestBody.SetDiff(opts.Diff)
		requestBody.SetPolicyOverride(opts.PolicyOverride)
	}

	request := j.JobsApi().PostJobPlan(j.client.Ctx, *job.ID)
	request = request.JobPlanRequest(requestBody)
	request = j.client.setWriteOptions(ctx, request).(client.ApiPostJobPlanRequest)

	result, response, err := request.Execute()
	if err != nil {
		return nil, nil, err
	}

	meta, err := parseWriteMeta(response)
	if err != nil {
		return nil, nil, err
	}

	return &result, meta, err
}

func (j *Jobs) Post(ctx context.Context, job *client.Job) (*client.JobRegisterResponse, *WriteMeta, error) {
	return j.Register(ctx, job, nil)
}

type RegisterOpts struct {
	EnforceIndex   bool
	ModifyIndex    uint64
	PolicyOverride bool
	PreserveCounts bool
}

func (j *Jobs) Register(ctx context.Context, job *client.Job, registerOpts *RegisterOpts) (*client.JobRegisterResponse, *WriteMeta, error) {
	if job == nil {
		return nil, nil, fmt.Errorf("must pass non-nil job")
	}

	request := j.JobsApi().RegisterJob(j.client.Ctx)
	request = j.client.setWriteOptions(ctx, request).(client.ApiRegisterJobRequest)

	registerRequest := client.NewJobRegisterRequest()
	registerRequest.SetJob(*job)

	if registerOpts != nil {
		if registerOpts.EnforceIndex {
			registerRequest.SetEnforceIndex(true)
			registerRequest.SetJobModifyIndex(int32(registerOpts.ModifyIndex))
		}
		registerRequest.SetPolicyOverride(registerOpts.PolicyOverride)
		registerRequest.SetPreserveCounts(registerOpts.PreserveCounts)
	}

	request = request.JobRegisterRequest(*registerRequest)

	result, response, err := request.Execute()
	if err != nil {
		return nil, nil, err
	}

	meta, err := parseWriteMeta(response)
	if err != nil {
		return nil, nil, err
	}

	return &result, meta, nil
}

func (j *Jobs) Revert(ctx context.Context, jobName string, versionNumber, enforcePriorVersion int32, consulToken, vaultToken string) (*client.JobRegisterResponse, *WriteMeta, error) {
	if jobName == "" {
		return nil, nil, jobNameErr
	}

	request := j.JobsApi().PostJobRevert(j.client.Ctx, jobName)
	request = j.client.setWriteOptions(ctx, request).(client.ApiPostJobRevertRequest)

	revertRequest := client.NewJobRevertRequest()
	revertRequest.SetJobVersion(versionNumber)
	revertRequest.SetJobID(jobName)

	if enforcePriorVersion != 0 {
		revertRequest.SetEnforcePriorVersion(enforcePriorVersion)
	}

	if consulToken != "" {
		revertRequest.SetConsulToken(consulToken)
	}

	if vaultToken != "" {
		revertRequest.SetVaultToken(vaultToken)
	}

	request = request.JobRevertRequest(*revertRequest)

	result, response, err := request.Execute()
	if err != nil {
		return nil, nil, err
	}

	meta, err := parseWriteMeta(response)
	if err != nil {
		return nil, nil, err
	}

	return &result, meta, nil
}

func (j *Jobs) Scale(ctx context.Context, jobName string, count int64, msg string, target map[string]string) (*client.JobRegisterResponse, *WriteMeta, error) {
	if jobName == "" {
		return nil, nil, jobNameErr
	}

	request := j.JobsApi().PostJobScalingRequest(j.client.Ctx, jobName)

	scalingRequest := client.NewScalingRequest()
	scalingRequest.SetCount(count)
	scalingRequest.SetMessage(msg)
	scalingRequest.SetTarget(target)

	request = j.client.setWriteOptions(ctx, request).(client.ApiPostJobScalingRequestRequest)
	request = request.ScalingRequest(*scalingRequest)

	result, response, err := request.Execute()
	if err != nil {
		return nil, nil, err
	}

	meta, err := parseWriteMeta(response)
	if err != nil {
		return nil, nil, err
	}

	return &result, meta, nil
}

func (j *Jobs) ScaleStatus(ctx context.Context, jobName string) (*client.JobScaleStatusResponse, *QueryMeta, error) {
	if jobName == "" {
		return nil, nil, jobNameErr
	}

	request := j.JobsApi().GetJobScaleStatus(j.client.Ctx, jobName)
	request = j.client.setQueryOptions(ctx, request).(client.ApiGetJobScaleStatusRequest)

	result, response, err := request.Execute()
	if err != nil {
		return nil, nil, err
	}

	meta, err := parseQueryMeta(response)
	if err != nil {
		return nil, nil, err
	}

	return &result, meta, nil
}

func (j *Jobs) Stability(ctx context.Context, jobName string, versionNumber int32, stable bool) (*client.JobStabilityResponse, *WriteMeta, error) {
	if jobName == "" {
		return nil, nil, jobNameErr
	}

	request := j.JobsApi().PostJobStability(j.client.Ctx, jobName)

	stabilityRequest := client.NewJobStabilityRequest()
	stabilityRequest.SetJobID(jobName)
	stabilityRequest.SetJobVersion(versionNumber)
	stabilityRequest.SetStable(stable)

	request = request.JobStabilityRequest(*stabilityRequest)

	request = j.client.setWriteOptions(ctx, request).(client.ApiPostJobStabilityRequest)

	result, response, err := request.Execute()
	if err != nil {
		return nil, nil, err
	}

	meta, err := parseWriteMeta(response)
	if err != nil {
		return nil, nil, err
	}

	return &result, meta, nil
}

func (j *Jobs) Summary(ctx context.Context, jobName string) (*client.JobSummary, *QueryMeta, error) {
	if jobName == "" {
		return nil, nil, jobNameErr
	}

	request := j.JobsApi().GetJobSummary(j.client.Ctx, jobName)
	request = j.client.setQueryOptions(ctx, request).(client.ApiGetJobSummaryRequest)

	result, response, err := request.Execute()
	if err != nil {
		return nil, nil, err
	}

	meta, err := parseQueryMeta(response)
	if err != nil {
		return nil, nil, err
	}

	return &result, meta, nil
}

func (j *Jobs) Versions(ctx context.Context, jobName string, withDiffs bool) (*client.JobVersionsResponse, *QueryMeta, error) {
	if jobName == "" {
		return nil, nil, jobNameErr
	}

	request := j.JobsApi().GetJobVersions(j.client.Ctx, jobName).Diffs(withDiffs)
	request = j.client.setQueryOptions(ctx, request).(client.ApiGetJobVersionsRequest)

	result, response, err := request.Execute()
	if err != nil {
		return nil, nil, err
	}

	meta, err := parseQueryMeta(response)
	if err != nil {
		return nil, nil, err
	}

	return &result, meta, nil
}

func (j *Jobs) GetLocation(job *client.Job) (*time.Location, error) {
	periodicConfig := *job.Periodic
	if job.Periodic == nil || periodicConfig.TimeZone == nil || *periodicConfig.TimeZone == "" {
		return time.UTC, nil
	}

	return time.LoadLocation(*periodicConfig.TimeZone)
}

func (j *Jobs) IsMultiRegion(job *client.Job) bool {
	return job.Multiregion != nil && *job.Multiregion.Regions != nil && len(*job.Multiregion.Regions) > 0
}

func (j *Jobs) IsParameterized(job *client.Job) bool {
	dispatched := false
	if job.Dispatched != nil {
		dispatched = *job.Dispatched
	}
	return job.ParameterizedJob != nil && !dispatched
}

func (j *Jobs) IsPeriodic(job *client.Job) bool {
	return job.Periodic != nil
}

// Next returns the closest time instant matching the spec that is after the
// passed time. If no matching instance exists, the zero value of time.Time is
// returned. The `time.Location` of the returned value matches that of the
// passed time.
func (j *Jobs) Next(p *client.PeriodicConfig, fromTime time.Time) (time.Time, error) {
	if *p.SpecType == PeriodicSpecCron {
		e, err := cronexpr.Parse(*p.Spec)
		if err != nil {
			return time.Time{}, fmt.Errorf("failed parsing cron expression %q: %v", *p.Spec, err)
		}
		return cronParseNext(e, fromTime, *p.Spec)
	}

	return time.Time{}, nil
}

var jobNameErr = errors.New("jobName must be specified")
