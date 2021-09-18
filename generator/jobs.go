package main

import (
	"net/http"

	"github.com/hashicorp/nomad/api"
	"github.com/hashicorp/nomad/nomad/structs"
)

func (v *v1api) getJobPaths() []*apiPath {
	tags := []string{"Jobs"}

	return []*apiPath{
		// "/jobs"
		{
			Template: "/jobs",
			Operations: []*operation{
				newOperation(http.MethodGet, "jobListRequest", tags, "GetJobs",
					nil,
					queryOptions,
					newResponseConfig(200, arraySchema, api.JobListStub{}, queryMeta, "GetJobsResponse"),
				),
				// TODO: Dedup
				newOperation(http.MethodPost, "jobUpdate", tags, "RegisterJob",
					newRequestBody(objectSchema, api.JobRegisterRequest{}),
					writeOptions,
					newResponseConfig(200, objectSchema, api.JobRegisterResponse{}, queryMeta, "RegisterJobResponse"),
				),
			},
		},
		// "/job/{jobName}"
		{
			Template: "/job/{jobName}",
			Operations: []*operation{
				newOperation(http.MethodGet, "jobQuery", tags, "GetJob",
					nil,
					appendParams(queryOptions, &jobNameParam),
					newResponseConfig(200, objectSchema, api.Job{}, queryMeta, "GetJobResponse"),
				),
				newOperation(http.MethodPost, "jobUpdate", tags, "PostJob",
					newRequestBody(objectSchema, api.JobRegisterRequest{}),
					appendParams(writeOptions, &jobNameParam),
					newResponseConfig(200, objectSchema, api.JobRegisterResponse{}, queryMeta, "PostJobResponse"),
				),
				newOperation(http.MethodDelete, "jobDelete", tags, "DeleteJob",
					nil,
					appendParams(writeOptions, &jobNameParam, &jobPurgeParam, &jobGlobalParam),
					newResponseConfig(200, objectSchema, api.JobDeregisterResponse{}, writeMeta, "DeleteJobResponse"),
				),
			},
		},
		// "/job/{jobName}/plan"
		{
			Template: "/job/{jobName}/plan",
			Operations: []*operation{
				newOperation(http.MethodPost, "jobPlan", tags, "PostJobPlan",
					newRequestBody(objectSchema, api.JobPlanRequest{}),
					appendParams(writeOptions, &jobNameParam),
					newResponseConfig(200, objectSchema, api.JobPlanResponse{}, writeMeta, "PostJobPlanResponse"),
				),
			},
		},
		// "/v1/jobs/parse",
		{
			Template: "/jobs/parse",
			Operations: []*operation{
				newOperation(http.MethodPost, "JobsParseRequest", tags, "PostJobParse",
					newRequestBody(objectSchema, api.JobsParseRequest{}),
					nil,
					newResponseConfig(200, objectSchema, api.Job{}, nil, "PostJobParseResponse"),
				),
			},
		},
		// "/job/{jobName}/evaluate")
		{
			Template: "/job/{jobName}/evaluate",
			Operations: []*operation{
				newOperation(http.MethodPost, "jobForceEvaluate", tags, "PostJobEvaluate",
					newRequestBody(objectSchema, api.JobEvaluateRequest{}),
					appendParams(writeOptions, &jobNameParam),
					newResponseConfig(200, objectSchema, api.JobRegisterResponse{}, writeMeta, "PostJobEvaluateResponse"),
				),
			},
		},
		//	"/job/{jobName}/allocations")
		{
			Template: "/job/{jobName}/allocations",
			Operations: []*operation{
				newOperation(http.MethodGet, "jobAllocations", tags, "GetJobAllocations",
					nil,
					appendParams(queryOptions, &jobNameParam),
					newResponseConfig(200, arraySchema, api.AllocationListStub{}, queryMeta, "GetJobAllocationsResponse"),
				),
			},
		},
		//	"/job/{jobName}/evaluations")
		{
			Template: "/job/{jobName}/evaluations",
			Operations: []*operation{
				newOperation(http.MethodGet, "jobEvaluations", tags, "GetJobEvaluations",
					nil,
					appendParams(queryOptions, &jobNameParam),
					newResponseConfig(200, arraySchema, api.Evaluation{}, queryMeta, "GetJobEvaluationsResponse"),
				),
			},
		},
		//	"/job/{jobName}/periodic/force")
		{
			Template: "/job/{jobName}/periodic/force",
			Operations: []*operation{
				newOperation(http.MethodPost, "periodicForceRequest", tags, "PostJobPeriodicForce",
					nil,
					appendParams(writeOptions, &jobNameParam),
					newResponseConfig(200, objectSchema, structs.PeriodicForceResponse{}, writeMeta, "PostJobPeriodicForceResponse"),
				),
			},
		},
		//	"/job/{jobName}/summary")
		{
			Template: "/job/{jobName}/summary",
			Operations: []*operation{
				newOperation(http.MethodGet, "jobSummaryRequest", tags, "GetJobSummary",
					nil,
					appendParams(queryOptions, &jobNameParam),
					newResponseConfig(200, objectSchema, api.JobSummary{}, queryMeta, "GetJobSummaryResponse"),
					&responseConfig{Code: 404, Headers: queryMeta},
				),
			},
		},
		//	"/job/{jobName}/dispatch")
		{
			Template: "/job/{jobName}/dispatch",
			Operations: []*operation{
				newOperation(http.MethodPost, "jobDispatchRequest", tags, "PostJobDispatch",
					newRequestBody(objectSchema, api.JobDispatchRequest{}),
					appendParams(writeOptions, &jobNameParam),
					newResponseConfig(200, objectSchema, api.JobDispatchResponse{}, writeMeta, "PostJobDispatchResponse"),
				),
			},
		},
		//	"/job/{jobName}/versions")
		{
			Template: "/job/{jobName}/versions",
			Operations: []*operation{
				newOperation(http.MethodGet, "jobVersions", tags, "GetJobVersions",
					nil,
					appendParams(queryOptions, &jobNameParam, &jobDiffsParam),
					newResponseConfig(200, objectSchema, api.JobVersionsResponse{}, queryMeta, "GetJobVersionsResponse"),
					&responseConfig{Code: 404, Headers: queryMeta},
				),
			},
		},
		//  "/job/{jobName}/revert")
		{
			Template: "/job/{jobName}/revert",
			Operations: []*operation{
				newOperation(http.MethodPost, "jobRevert", tags, "PostJobRevert",
					newRequestBody(objectSchema, api.JobRevertRequest{}),
					appendParams(writeOptions, &jobNameParam),
					newResponseConfig(200, objectSchema, api.JobRegisterResponse{}, writeMeta, "PostJobRevertResponse"),
				),
			},
		},
		//	"/job/{jobName}/deployments")
		{
			Template: "/job/{jobName}/deployments",
			Operations: []*operation{
				newOperation(http.MethodGet, "jobDeployments", tags, "GetJobDeployments",
					nil,
					appendParams(queryOptions, &jobNameParam, &allParam),
					newResponseConfig(200, arraySchema, api.Deployment{}, queryMeta, "GetJobDeploymentsResponse"),
				),
			},
		},
		//	"/job/{jobName}/deployment")
		{
			Template: "/job/{jobName}/deployment",
			Operations: []*operation{
				newOperation(http.MethodGet, "jobLatestDeployment", tags, "GetJobDeployment",
					nil,
					appendParams(queryOptions, &jobNameParam),
					newResponseConfig(200, objectSchema, api.Deployment{}, queryMeta, "GetJobDeploymentResponse"),
				),
			},
		},
		//	"/job/{jobName}/stable")
		{
			Template: "/job/{jobName}/stable",
			Operations: []*operation{
				newOperation(http.MethodPost, "jobStable", tags, "PostJobStability",
					newRequestBody(objectSchema, api.JobStabilityRequest{}),
					appendParams(writeOptions, &jobNameParam),
					newResponseConfig(200, objectSchema, structs.JobStabilityResponse{}, writeMeta, "PostJobStabilityResponse"),
				),
			},
		},
		//	"/job/{jobName}/scale")
		{
			Template: "/job/{jobName}/scale",
			Operations: []*operation{
				newOperation(http.MethodGet, "jobScaleStatus", tags, "GetJobScaleStatus",
					nil,
					appendParams(queryOptions, &jobNameParam),
					newResponseConfig(200, objectSchema, api.JobScaleStatusResponse{}, queryMeta, "GetJobScaleStatusResponse"),
				),
				newOperation(http.MethodPost, "jobScaleAction", tags, "PostJobScalingRequest",
					newRequestBody(objectSchema, api.ScalingRequest{}),
					appendParams(writeOptions, &jobNameParam),
					newResponseConfig(200, objectSchema, structs.JobRegisterResponse{}, writeMeta, "PostJobScalingResponse"),
				),
			},
		},
		// "/v1/validate/job"
		{
			Template: "/validate/job",
			Operations: []*operation{
				newOperation(http.MethodPost, "ValidateJobRequest", tags, "PostJobValidateRequest",
					newRequestBody(objectSchema, api.JobValidateRequest{}),
					writeOptions,
					newResponseConfig(200, objectSchema, structs.JobValidateResponse{}, writeMeta, "PostJobValidateResponse"),
				),
			},
		},
	}
}
