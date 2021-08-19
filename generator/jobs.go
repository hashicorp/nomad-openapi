package main

import (
	"net/http"

	"github.com/hashicorp/nomad/api"
	"github.com/hashicorp/nomad/nomad/structs"
)

func (v *v1api) getJobPaths() []*apiPath {
	tags := []string{"Jobs"}

	return []*apiPath{
		{
			Template: "/jobs",
			Operations: []*operation{
				newOperation(http.MethodGet, "jobListRequest", tags, "GetJobs",
					nil,
					queryOptions,
					newResponseConfig(200, arraySchema, api.JobListStub{}, queryMeta, "GetJobsResponse"),
				),
				newOperation(http.MethodPost, "jobUpdate", tags, "PostJob",
					newRequestBody(objectSchema, api.JobRegisterRequest{}),
					nil,
					newResponseConfig(200, objectSchema, api.JobRegisterResponse{}, queryMeta, "PostJobResponse"),
				),
			},
		},
		{
			Template: "/job/{jobName}",
			Operations: []*operation{
				newOperation(http.MethodGet, "jobQuery", tags, "GetJob",
					nil,
					appendParams(queryOptions, &jobNameParam),
					newResponseConfig(200, arraySchema, api.Job{}, queryMeta, "GetJobResponse"),
				),
				newOperation(http.MethodDelete, "jobDelete", tags, "DeleteJob",
					nil,
					appendParams(writeOptions, &jobNameParam, &jobPurgeParam, &jobGlobalParam),
					newResponseConfig(200, objectSchema, api.JobDeregisterResponse{}, writeMeta, "DeleteJobResponse"),
				),
			},
		},
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
				newOperation(http.MethodGet, "jobAllocations", tags, "GetAllocations",
					nil,
					appendParams(queryOptions, &jobNameParam),
					newResponseConfig(200, arraySchema, api.AllocationListStub{}, queryMeta, "GetAllocationsResponse"),
				),
			},
		},
		//	"/job/{jobName}/evaluations")
		{
			Template: "/job/{jobName}/evaluations",
			Operations: []*operation{
				newOperation(http.MethodGet, "jobEvaluations", tags, "GetEvaluations",
					nil,
					appendParams(queryOptions, &jobNameParam),
					newResponseConfig(200, arraySchema, api.Evaluation{}, queryMeta, "GetEvaluationsResponse"),
				),
			},
		},
		//	"/job/{jobName}/periodic/force")
		{
			Template: "/job/{jobName}/periodic/force",
			Operations: []*operation{
				newOperation(http.MethodPost, "periodicForceRequest", tags, "PostPeriodicForce",
					nil,
					appendParams(writeOptions, &jobNameParam),
					newResponseConfig(200, objectSchema, structs.PeriodicForceResponse{}, writeMeta, "PostPeriodicForceResponse"),
				),
			},
		},
		//	"/job/{jobName}/plan")
		//	"/job/{jobName}/summary")
		//	"/job/{jobName}/dispatch")
		//	"/job/{jobName}/versions")
		//  "/job/{jobName}/revert")
		//	"/job/{jobName}/deployments")
		//	"/job/{jobName}/deployment")
		//	"/job/{jobName}/stable")
		//	"/job/{jobName}/scale")
		//s.mux.HandleFunc("/v1/jobs/parse", s.wrap(s.JobsParseRequest))
		//s.mux.HandleFunc("/v1/job/", s.wrap(s.JobSpecificRequest))
		//s.mux.HandleFunc("/v1/validate/job", s.wrap(s.ValidateJobRequest))
	}
}
