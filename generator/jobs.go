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
					defaultQueryOpts,
					newResponseConfig(200, arraySchema, api.JobListStub{}, defaultQueryMeta, "GetJobsResponse"),
				),
				// TODO: Dedup
				newOperation(http.MethodPost, "jobUpdate", tags, "RegisterJob",
					newRequestBody(objectSchema, api.JobRegisterRequest{}),
					defaultWriteOpts,
					newResponseConfig(200, objectSchema, api.JobRegisterResponse{}, defaultQueryMeta, "RegisterJobResponse"),
				),
			},
		},
		// "/job/{jobName}"
		{
			Template: "/job/{jobName}",
			Operations: []*operation{
				newOperation(http.MethodGet, "jobQuery", tags, "GetJob",
					nil,
					appendParams(defaultQueryOpts, &jobNameParam),
					newResponseConfig(200, objectSchema, api.Job{}, defaultQueryMeta, "GetJobResponse"),
				),
				newOperation(http.MethodPost, "jobUpdate", tags, "PostJob",
					newRequestBody(objectSchema, api.JobRegisterRequest{}),
					appendParams(defaultWriteOpts, &jobNameParam),
					newResponseConfig(200, objectSchema, api.JobRegisterResponse{}, defaultQueryMeta, "PostJobResponse"),
				),
				newOperation(http.MethodDelete, "jobDelete", tags, "DeleteJob",
					nil,
					appendParams(defaultWriteOpts, &jobNameParam, &jobPurgeParam, &jobGlobalParam),
					newResponseConfig(200, objectSchema, api.JobDeregisterResponse{}, defaultWriteMeta, "DeleteJobResponse"),
				),
			},
		},
		// "/job/{jobName}/plan"
		{
			Template: "/job/{jobName}/plan",
			Operations: []*operation{
				newOperation(http.MethodPost, "jobPlan", tags, "PostJobPlan",
					newRequestBody(objectSchema, api.JobPlanRequest{}),
					appendParams(defaultWriteOpts, &jobNameParam),
					newResponseConfig(200, objectSchema, api.JobPlanResponse{}, defaultWriteMeta, "PostJobPlanResponse"),
				),
			},
		},
		// "/v1/jobs/parse",
		{
			Template: "/jobs/parse",
			Operations: []*operation{
				newOperation(http.MethodPost,
					httpServer.JobsParseRequest,
					tags,
					"PostJobParse",
					newRequestBody(objectSchema, api.JobsParseRequest{}),
					nil,
					newResponseConfig(200,
						objectSchema,
						api.Job{},
						nil,
						"PostJobParseResponse",
					),
				),
			},
		},
		// "/job/{jobName}/evaluate")
		{
			Template: "/job/{jobName}/evaluate",
			Operations: []*operation{
				newOperation(http.MethodPost,
					"jobForceEvaluate",
					tags,
					"PostJobEvaluate",
					newRequestBody(objectSchema, api.JobEvaluateRequest{}),
					appendParams(defaultWriteOpts, &jobNameParam),
					newResponseConfig(200,
						objectSchema,
						api.JobRegisterResponse{},
						defaultWriteMeta,
						"PostJobEvaluateResponse",
					),
				),
			},
		},
		//	"/job/{jobName}/allocations")
		{
			Template: "/job/{jobName}/allocations",
			Operations: []*operation{
				newOperation(http.MethodGet, "jobAllocations", tags, "GetJobAllocations",
					nil,
					appendParams(defaultQueryOpts, &jobNameParam, &jobAllocationsAllParam),
					newResponseConfig(200, arraySchema, api.AllocationListStub{}, defaultQueryMeta, "GetJobAllocationsResponse"),
				),
			},
		},
		//	"/job/{jobName}/evaluations")
		{
			Template: "/job/{jobName}/evaluations",
			Operations: []*operation{
				newOperation(http.MethodGet, "jobEvaluations", tags, "GetJobEvaluations",
					nil,
					appendParams(defaultQueryOpts, &jobNameParam),
					newResponseConfig(200, arraySchema, api.Evaluation{}, defaultQueryMeta, "GetJobEvaluationsResponse"),
				),
			},
		},
		//	"/job/{jobName}/periodic/force")
		{
			Template: "/job/{jobName}/periodic/force",
			Operations: []*operation{
				newOperation(http.MethodPost, "periodicForceRequest", tags, "PostJobPeriodicForce",
					nil,
					appendParams(defaultWriteOpts, &jobNameParam),
					newResponseConfig(200, objectSchema, structs.PeriodicForceResponse{}, defaultWriteMeta, "PostJobPeriodicForceResponse"),
				),
			},
		},
		//	"/job/{jobName}/summary")
		{
			Template: "/job/{jobName}/summary",
			Operations: []*operation{
				newOperation(http.MethodGet, "jobSummaryRequest", tags, "GetJobSummary",
					nil,
					appendParams(defaultQueryOpts, &jobNameParam),
					newResponseConfig(200, objectSchema, api.JobSummary{}, defaultQueryMeta, "GetJobSummaryResponse"),
					&responseConfig{Code: 404, Headers: defaultQueryMeta},
				),
			},
		},
		//	"/job/{jobName}/dispatch")
		{
			Template: "/job/{jobName}/dispatch",
			Operations: []*operation{
				newOperation(http.MethodPost, "jobDispatchRequest", tags, "PostJobDispatch",
					newRequestBody(objectSchema, api.JobDispatchRequest{}),
					appendParams(defaultWriteOpts, &jobNameParam),
					newResponseConfig(200, objectSchema, api.JobDispatchResponse{}, defaultWriteMeta, "PostJobDispatchResponse"),
				),
			},
		},
		//	"/job/{jobName}/versions")
		{
			Template: "/job/{jobName}/versions",
			Operations: []*operation{
				newOperation(http.MethodGet, "jobVersions", tags, "GetJobVersions",
					nil,
					appendParams(defaultQueryOpts, &jobNameParam, &jobDiffsParam),
					newResponseConfig(200, objectSchema, api.JobVersionsResponse{}, defaultQueryMeta, "GetJobVersionsResponse"),
					&responseConfig{Code: 404, Headers: defaultQueryMeta},
				),
			},
		},
		//  "/job/{jobName}/revert")
		{
			Template: "/job/{jobName}/revert",
			Operations: []*operation{
				newOperation(http.MethodPost, "jobRevert", tags, "PostJobRevert",
					newRequestBody(objectSchema, api.JobRevertRequest{}),
					appendParams(defaultWriteOpts, &jobNameParam),
					newResponseConfig(200, objectSchema, api.JobRegisterResponse{}, defaultWriteMeta, "PostJobRevertResponse"),
				),
			},
		},
		//	"/job/{jobName}/deployments")
		{
			Template: "/job/{jobName}/deployments",
			Operations: []*operation{
				newOperation(http.MethodGet, "jobDeployments", tags, "GetJobDeployments",
					nil,
					appendParams(defaultQueryOpts, &jobNameParam, &allParam),
					newResponseConfig(200, arraySchema, api.Deployment{}, defaultQueryMeta, "GetJobDeploymentsResponse"),
				),
			},
		},
		//	"/job/{jobName}/deployment")
		{
			Template: "/job/{jobName}/deployment",
			Operations: []*operation{
				newOperation(http.MethodGet, "jobLatestDeployment", tags, "GetJobDeployment",
					nil,
					appendParams(defaultQueryOpts, &jobNameParam),
					newResponseConfig(200, objectSchema, api.Deployment{}, defaultQueryMeta, "GetJobDeploymentResponse"),
				),
			},
		},
		//	"/job/{jobName}/stable")
		{
			Template: "/job/{jobName}/stable",
			Operations: []*operation{
				newOperation(http.MethodPost, "jobStable", tags, "PostJobStability",
					newRequestBody(objectSchema, api.JobStabilityRequest{}),
					appendParams(defaultWriteOpts, &jobNameParam),
					newResponseConfig(200, objectSchema, structs.JobStabilityResponse{}, defaultWriteMeta, "PostJobStabilityResponse"),
				),
			},
		},
		//	"/job/{jobName}/scale")
		{
			Template: "/job/{jobName}/scale",
			Operations: []*operation{
				newOperation(http.MethodGet, "jobScaleStatus", tags, "GetJobScaleStatus",
					nil,
					appendParams(defaultQueryOpts, &jobNameParam),
					newResponseConfig(200, objectSchema, api.JobScaleStatusResponse{}, defaultQueryMeta, "GetJobScaleStatusResponse"),
				),
				newOperation(http.MethodPost, "jobScaleAction", tags, "PostJobScalingRequest",
					newRequestBody(objectSchema, api.ScalingRequest{}),
					appendParams(defaultWriteOpts, &jobNameParam),
					newResponseConfig(200, objectSchema, structs.JobRegisterResponse{}, defaultWriteMeta, "PostJobScalingResponse"),
				),
			},
		},
		// "/v1/validate/job"
		{
			Template: "/validate/job",
			Operations: []*operation{
				newOperation(http.MethodPost,
					httpServer.ValidateJobRequest,
					tags, "PostJobValidateRequest",
					newRequestBody(
						objectSchema,
						api.JobValidateRequest{},
					),
					defaultWriteOpts,
					newResponseConfig(200,
						objectSchema,
						structs.JobValidateResponse{},
						defaultWriteMeta,
						"PostJobValidateResponse",
					),
				),
			},
		},
	}
}
