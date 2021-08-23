package api

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	openapi "github.com/hashicorp/nomad-openapi/v1/client"
	"github.com/hashicorp/nomad/api"
	"github.com/hashicorp/nomad/command/agent"
	"github.com/hashicorp/nomad/helper"
	"github.com/hashicorp/nomad/helper/uuid"
	"github.com/hashicorp/nomad/nomad/mock"
	"github.com/hashicorp/nomad/nomad/structs"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func postTestJob(s *agent.TestAgent, t *testing.T, job *openapi.Job) {
	client, err := NewTestWriteClient(s, writeOpts)
	require.NoError(t, err)

	if job == nil {
		job = mockJob()
	}
	resp, writeMeta, err := client.Jobs().Post(job)

	require.NoError(t, err)
	require.NotNil(t, resp)
	require.NotNil(t, writeMeta)
}

func TestJobsGet(t *testing.T) {
	t.Parallel()
	httpTest(t, nil, func(s *agent.TestAgent) {
		postTestJob(s, t, nil)

		client, err := NewTestQueryClient(s, queryOpts)
		require.NoError(t, err)

		jobs, queryMeta, err := client.Jobs().GetJobs()
		require.NoError(t, err)
		require.NotNil(t, jobs)
		require.Len(t, jobs, 1)
		require.NotNil(t, queryMeta)
	})
}

func TestJobGet(t *testing.T) {
	t.Parallel()
	httpTest(t, nil, func(s *agent.TestAgent) {
		job := mockJob()
		postTestJob(s, t, job)

		client, err := NewTestQueryClient(s, queryOpts)
		require.NoError(t, err)

		queryJob, queryMeta, err := client.Jobs().GetJob(*job.ID)
		require.NoError(t, err)
		require.NotNil(t, queryJob)
		require.NotNil(t, queryMeta)

		if *job.ID != *queryJob.ID {
			t.Fatalf("TestJobGet invalid job comparison: %s != %s", *job.ID, *queryJob.ID)
		}
	})
}

func TestPostJob(t *testing.T) {
	t.Parallel()
	httpTest(t, nil, func(s *agent.TestAgent) {
		job := mockJob()

		client, err := NewTestWriteClient(s, writeOpts)
		require.NoError(t, err)

		resp, meta, err := client.Jobs().Post(job)

		require.NoError(t, err)
		require.NotNil(t, resp)
		require.NotNil(t, meta)
	})
}

func TestPlanJob(t *testing.T) {
	t.Parallel()
	httpTest(t, nil, func(s *agent.TestAgent) {
		postTestJob(s, t, nil)

		client, err := NewTestWriteClient(s, writeOpts)
		require.NoError(t, err)

		diffJob := mockJobWithDiff()

		response, meta, err := client.Jobs().Plan(diffJob, true)
		require.NoError(t, err)
		require.NotNil(t, response)
		require.NotNil(t, meta)
	})
}

func TestJobDelete(t *testing.T) {
	t.Parallel()
	httpTest(t, nil, func(s *agent.TestAgent) {
		job := mockJob()
		postTestJob(s, t, job)

		// Make the HTTP request to do a soft delete
		client, err := NewTestWriteClient(s, writeOpts)
		require.NoError(t, err)

		response, writeMeta, err := client.Jobs().Delete(*job.ID, false, false)
		require.NoError(t, err)
		require.NotNil(t, response)
		require.NotNil(t, writeMeta)

		// Check the response
		require.NotNil(t, response.EvalID)
		require.NotEqual(t, "", *response.EvalID)

		// Check the job is still queryable
		client, err = NewTestQueryClient(s, queryOpts)
		require.NoError(t, err)

		rpcRequest := structs.JobSpecificRequest{
			JobID: *job.ID,
			QueryOptions: structs.QueryOptions{
				Region:    "global",
				Namespace: structs.DefaultNamespace,
			},
		}

		var rpcResponse structs.SingleJobResponse

		err = s.Agent.RPC("Job.GetJob", &rpcRequest, &rpcResponse)
		require.NoError(t, err)
		require.NotNil(t, rpcResponse.Job)
		require.True(t, rpcResponse.Job.Stop)

		// Make the HTTP request to do a purge delete
		client, err = NewTestWriteClient(s, writeOpts)
		require.NoError(t, err)

		response, writeMeta, err = client.Jobs().Delete(*job.ID, true, false)
		require.NoError(t, err)
		require.NotNil(t, response)
		require.NotNil(t, writeMeta)

		// Check the response
		require.NotNil(t, response.EvalID)
		require.NotEqual(t, "", *response.EvalID)

		// Check the job is gone
		err = s.Agent.RPC("Job.GetJob", rpcRequest, &rpcResponse)
		require.NoError(t, err)
		require.Nil(t, rpcResponse.Job)
	})
}

func TestJobParse(t *testing.T) {
	t.Parallel()
	httpTest(t, nil, func(s *agent.TestAgent) {
		client, err := NewTestQueryClient(s, queryOpts)
		require.NoError(t, err)

		job, err := client.Jobs().Parse(jobHCL, false, false)
		require.NoError(t, err)
		require.NotNil(t, job)

		expected := mockJob()
		require.NotNil(t, job.Name)
		require.Equal(t, *job.Name, *expected.Name)

		if job.Datacenters == nil {
			expectedDatacenters := *expected.Datacenters
			jobDatacenters := *job.Datacenters
			require.NotEqual(t, jobDatacenters[0], expectedDatacenters[0])
		}
	})
}

func TestJobEvaluate(t *testing.T) {
	t.Parallel()
	httpTest(t, nil, func(s *agent.TestAgent) {
		job := mockJob()
		postTestJob(s, t, job)

		client, err := NewTestWriteClient(s, writeOpts)
		require.NoError(t, err)

		// Make the HTTP request
		model, meta, err := client.Jobs().Evaluate(*job.ID, false)
		require.NoError(t, err)
		require.NotNil(t, meta)
		require.NotNil(t, model)

		// Check the response
		require.NotNil(t, model.EvalID)
		require.NotEqual(t, "", *model.EvalID)
	})
}

func TestJobPeriodicForce(t *testing.T) {
	t.Parallel()
	httpTest(t, nil, func(s *agent.TestAgent) {
		// Create and register a periodic job.
		job := mockPeriodicJob()
		postTestJob(s, t, job)

		client, err := NewTestWriteClient(s, writeOpts)
		require.NoError(t, err)

		// Make the HTTP request
		model, meta, err := client.Jobs().PeriodicForce(*job.ID)
		require.NoError(t, err)
		require.NotNil(t, meta)
		// Check the response
		require.NotNil(t, model.EvalID)
		require.NotEqual(t, "", *model.EvalID)
	})
}

func TestJobSummary(t *testing.T) {
	t.Parallel()
	httpTest(t, nil, func(s *TestAgent) {
		TODO: Find current test
	})
}

func TestJobDispatch(t *testing.T) {
	t.Parallel()
	httpTest(t, nil, func(s *TestAgent) {
		// Create the parameterized job
		job := mock.BatchJob()
		job.ParameterizedJob = &structs.ParameterizedJobConfig{}

		args := structs.JobRegisterRequest{
			Job: job,
			WriteRequest: structs.WriteRequest{
				Region:    "global",
				Namespace: structs.DefaultNamespace,
			},
		}
		var resp structs.JobRegisterResponse
		if err := s.Agent.RPC("Job.Register", &args, &resp); err != nil {
			t.Fatalf("err: %v", err)
		}

		// Make the request
		respW := httptest.NewRecorder()
		args2 := structs.JobDispatchRequest{
			WriteRequest: structs.WriteRequest{
				Region:    "global",
				Namespace: structs.DefaultNamespace,
			},
		}
		buf := encodeReq(args2)

		// Make the HTTP request
		req2, err := http.NewRequest("PUT", "/v1/job/"+job.ID+"/dispatch", buf)
		if err != nil {
			t.Fatalf("err: %v", err)
		}
		respW.Flush()

		// Make the request
		obj, err := s.Server.JobSpecificRequest(respW, req2)
		if err != nil {
			t.Fatalf("err: %v", err)
		}

		// Check the response
		dispatch := obj.(structs.JobDispatchResponse)
		if dispatch.EvalID == "" {
			t.Fatalf("bad: %v", dispatch)
		}

		if dispatch.DispatchedJobID == "" {
			t.Fatalf("bad: %v", dispatch)
		}
	})
}

func TestHTTP_JobVersions(t *testing.T) {
	t.Parallel()
	httpTest(t, nil, func(s *TestAgent) {
		// Create the job
		job := mock.Job()
		args := structs.JobRegisterRequest{
			Job: job,
			WriteRequest: structs.WriteRequest{
				Region:    "global",
				Namespace: structs.DefaultNamespace,
			},
		}
		var resp structs.JobRegisterResponse
		if err := s.Agent.RPC("Job.Register", &args, &resp); err != nil {
			t.Fatalf("err: %v", err)
		}

		job2 := mock.Job()
		job2.ID = job.ID
		job2.Priority = 100

		args2 := structs.JobRegisterRequest{
			Job: job2,
			WriteRequest: structs.WriteRequest{
				Region:    "global",
				Namespace: structs.DefaultNamespace,
			},
		}
		var resp2 structs.JobRegisterResponse
		if err := s.Agent.RPC("Job.Register", &args2, &resp2); err != nil {
			t.Fatalf("err: %v", err)
		}

		// Make the HTTP request
		req, err := http.NewRequest("GET", "/v1/job/"+job.ID+"/versions?diffs=true", nil)
		if err != nil {
			t.Fatalf("err: %v", err)
		}
		respW := httptest.NewRecorder()

		// Make the request
		obj, err := s.Server.JobSpecificRequest(respW, req)
		if err != nil {
			t.Fatalf("err: %v", err)
		}

		// Check the response
		vResp := obj.(structs.JobVersionsResponse)
		versions := vResp.Versions
		if len(versions) != 2 {
			t.Fatalf("got %d versions; want 2", len(versions))
		}

		if v := versions[0]; v.Version != 1 || v.Priority != 100 {
			t.Fatalf("bad %v", v)
		}

		if v := versions[1]; v.Version != 0 {
			t.Fatalf("bad %v", v)
		}

		if len(vResp.Diffs) != 1 {
			t.Fatalf("bad %v", vResp)
		}

		// Check for the index
		if respW.HeaderMap.Get("X-Nomad-Index") == "" {
			t.Fatalf("missing index")
		}
		if respW.HeaderMap.Get("X-Nomad-KnownLeader") != "true" {
			t.Fatalf("missing known leader")
		}
		if respW.HeaderMap.Get("X-Nomad-LastContact") == "" {
			t.Fatalf("missing last contact")
		}
	})
}

func TestHTTP_JobRevert(t *testing.T) {
	t.Parallel()
	httpTest(t, nil, func(s *TestAgent) {
		// Create the job and register it twice
		job := mock.Job()
		regReq := structs.JobRegisterRequest{
			Job: job,
			WriteRequest: structs.WriteRequest{
				Region:    "global",
				Namespace: structs.DefaultNamespace,
			},
		}
		var regResp structs.JobRegisterResponse
		if err := s.Agent.RPC("Job.Register", &regReq, &regResp); err != nil {
			t.Fatalf("err: %v", err)
		}

		// Change the job to get a new version
		job.Datacenters = append(job.Datacenters, "foo")
		if err := s.Agent.RPC("Job.Register", &regReq, &regResp); err != nil {
			t.Fatalf("err: %v", err)
		}

		args := structs.JobRevertRequest{
			JobID:      job.ID,
			JobVersion: 0,
			WriteRequest: structs.WriteRequest{
				Region:    "global",
				Namespace: structs.DefaultNamespace,
			},
		}
		buf := encodeReq(args)

		// Make the HTTP request
		req, err := http.NewRequest("PUT", "/v1/job/"+job.ID+"/revert", buf)
		if err != nil {
			t.Fatalf("err: %v", err)
		}
		respW := httptest.NewRecorder()

		// Make the request
		obj, err := s.Server.JobSpecificRequest(respW, req)
		if err != nil {
			t.Fatalf("err: %v", err)
		}

		// Check the response
		revertResp := obj.(structs.JobRegisterResponse)
		if revertResp.EvalID == "" {
			t.Fatalf("bad: %v", revertResp)
		}

		// Check for the index
		if respW.HeaderMap.Get("X-Nomad-Index") == "" {
			t.Fatalf("missing index")
		}
	})
}

func TestHTTP_JobDeployments(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()
	httpTest(t, nil, func(s *TestAgent) {
		// Create the job
		j := mock.Job()
		args := structs.JobRegisterRequest{
			Job: j,
			WriteRequest: structs.WriteRequest{
				Region:    "global",
				Namespace: structs.DefaultNamespace,
			},
		}
		var resp structs.JobRegisterResponse
		assert.Nil(s.Agent.RPC("Job.Register", &args, &resp), "JobRegister")

		// Directly manipulate the state
		state := s.Agent.server.State()
		d := mock.Deployment()
		d.JobID = j.ID
		d.JobCreateIndex = resp.JobModifyIndex

		assert.Nil(state.UpsertDeployment(1000, d), "UpsertDeployment")

		// Make the HTTP request
		req, err := http.NewRequest("GET", "/v1/job/"+j.ID+"/deployments", nil)
		assert.Nil(err, "HTTP")
		respW := httptest.NewRecorder()

		// Make the request
		obj, err := s.Server.JobSpecificRequest(respW, req)
		assert.Nil(err, "JobSpecificRequest")

		// Check the response
		deploys := obj.([]*structs.Deployment)
		assert.Len(deploys, 1, "deployments")
		assert.Equal(d.ID, deploys[0].ID, "deployment id")

		assert.NotZero(respW.HeaderMap.Get("X-Nomad-Index"), "missing index")
		assert.Equal("true", respW.HeaderMap.Get("X-Nomad-KnownLeader"), "missing known leader")
		assert.NotZero(respW.HeaderMap.Get("X-Nomad-LastContact"), "missing last contact")
	})
}

func TestHTTP_JobStable(t *testing.T) {
	t.Parallel()
	httpTest(t, nil, func(s *TestAgent) {
		// Create the job and register it twice
		job := mock.Job()
		regReq := structs.JobRegisterRequest{
			Job: job,
			WriteRequest: structs.WriteRequest{
				Region:    "global",
				Namespace: structs.DefaultNamespace,
			},
		}
		var regResp structs.JobRegisterResponse
		if err := s.Agent.RPC("Job.Register", &regReq, &regResp); err != nil {
			t.Fatalf("err: %v", err)
		}

		if err := s.Agent.RPC("Job.Register", &regReq, &regResp); err != nil {
			t.Fatalf("err: %v", err)
		}

		args := structs.JobStabilityRequest{
			JobID:      job.ID,
			JobVersion: 0,
			Stable:     true,
			WriteRequest: structs.WriteRequest{
				Region:    "global",
				Namespace: structs.DefaultNamespace,
			},
		}
		buf := encodeReq(args)

		// Make the HTTP request
		req, err := http.NewRequest("PUT", "/v1/job/"+job.ID+"/stable", buf)
		if err != nil {
			t.Fatalf("err: %v", err)
		}
		respW := httptest.NewRecorder()

		// Make the request
		obj, err := s.Server.JobSpecificRequest(respW, req)
		if err != nil {
			t.Fatalf("err: %v", err)
		}

		// Check the response
		stableResp := obj.(structs.JobStabilityResponse)
		if stableResp.Index == 0 {
			t.Fatalf("bad: %v", stableResp)
		}

		// Check for the index
		if respW.HeaderMap.Get("X-Nomad-Index") == "" {
			t.Fatalf("missing index")
		}
	})
}

func TestJobScaleStatus(t *testing.T) {
t.Parallel()

require := require.New(t)

httpTest(t, nil, func(s *TestAgent) {
// Create the job
job := mock.Job()
args := structs.JobRegisterRequest{
Job: job,
WriteRequest: structs.WriteRequest{
Region:    "global",
Namespace: structs.DefaultNamespace,
},
}
var resp structs.JobRegisterResponse
if err := s.Agent.RPC("Job.Register", &args, &resp); err != nil {
t.Fatalf("err: %v", err)
}

// Make the HTTP request to scale the job group
req, err := http.NewRequest("GET", "/v1/job/"+job.ID+"/scale", nil)
require.NoError(err)
respW := httptest.NewRecorder()

// Make the request
obj, err := s.Server.JobSpecificRequest(respW, req)
require.NoError(err)

// Check the response
status := obj.(*structs.JobScaleStatus)
require.NotEmpty(resp.EvalID)
require.Equal(job.TaskGroups[0].Count, status.TaskGroups[job.TaskGroups[0].Name].Desired)

// Check for the index
require.NotEmpty(respW.Header().Get("X-Nomad-Index"))
})
}

func TestHTTPJobScaleTaskGroup(t *testing.T) {
	t.Parallel()

	require := require.New(t)

	httpTest(t, nil, func(s *TestAgent) {
		// Create the job
		job := mock.Job()
		args := structs.JobRegisterRequest{
			Job: job,
			WriteRequest: structs.WriteRequest{
				Region:    "global",
				Namespace: structs.DefaultNamespace,
			},
		}
		var resp structs.JobRegisterResponse
		require.NoError(s.Agent.RPC("Job.Register", &args, &resp))

		newCount := job.TaskGroups[0].Count + 1
		scaleReq := &api.ScalingRequest{
			Count:   helper.Int64ToPtr(int64(newCount)),
			Message: "testing",
			Target: map[string]string{
				"Job":   job.ID,
				"Group": job.TaskGroups[0].Name,
			},
		}
		buf := encodeReq(scaleReq)

		// Make the HTTP request to scale the job group
		req, err := http.NewRequest("POST", "/v1/job/"+job.ID+"/scale", buf)
		require.NoError(err)
		respW := httptest.NewRecorder()

		// Make the request
		obj, err := s.Server.JobSpecificRequest(respW, req)
		require.NoError(err)

		// Check the response
		resp = obj.(structs.JobRegisterResponse)
		require.NotEmpty(resp.EvalID)

		// Check for the index
		require.NotEmpty(respW.Header().Get("X-Nomad-Index"))

		// Check that the group count was changed
		getReq := structs.JobSpecificRequest{
			JobID: job.ID,
			QueryOptions: structs.QueryOptions{
				Region:    "global",
				Namespace: structs.DefaultNamespace,
			},
		}
		var getResp structs.SingleJobResponse
		err = s.Agent.RPC("Job.GetJob", &getReq, &getResp)
		require.NoError(err)
		require.NotNil(getResp.Job)
		require.Equal(newCount, getResp.Job.TaskGroups[0].Count)
	})
}


var (
	id                                  = fmt.Sprintf("mock-service-%s", uuid.Generate())
	dbLabel                             = "db"
	jobName                             = "my-job"
	jobTypeService                      = "service"
	docker                              = "docker"
	redis                               = "redis"
	pendingStatus                       = "pending"
	globalRegion                        = "global"
	defaultNamespace                    = "default"
	allAtOnce                           = false
	lTarget                             = "${attr.kernel.name}"
	rTarget                             = "linux"
	operand                             = "="
	web                                 = "web"
	priority                      int32 = 50
	count                         int32 = 10
	sizeMB                        int32 = 150
	restartPolicyAttempts         int32 = 3
	restartPolicyInterval               = int64(10 * time.Minute)
	restartPolicyDelay                  = int64(1 * time.Minute)
	restartPolicyMode                   = structs.RestartPolicyModeDelay
	reschedulePolicyAttempts      int32 = 2
	reschedulePolicyInterval            = int64(10 * time.Minute)
	reschedulePolicyDelay               = int64(5 * time.Second)
	reschedulePolicyDelayFunction       = "constant"
	maxParallel                   int32 = 1
	checks                              = "checks"
	minHealthyTime                      = int64(10 * time.Second)
	healthyDeadline                     = int64(5 * time.Minute)
	defaultMigrateStrategy              = &openapi.MigrateStrategy{
		MaxParallel:     &maxParallel,
		HealthCheck:     &checks,
		MinHealthyTime:  &minHealthyTime,
		HealthyDeadline: &healthyDeadline,
	}
	hostMode               = "host"
	httpLabel              = "http"
	adminLabel             = "admin"
	execDriver             = "exec"
	frontEndTaskName       = "${TASK}-frontend"
	serviceCheckName       = "check-table"
	serviceCheckType       = structs.ServiceCheckScript
	serviceCheckCommand    = "/usr/local/check-table-${meta.database}"
	serviceCheckInterval   = int64(30 * time.Second)
	serviceCheckTimeout    = int64(5 * time.Second)
	adminTaskName          = "${TASK}-admin"
	logConfigMaxFile       = int32(10)
	logConfigMaxFileSizeMB = int32(10)
	defaultLogConfig       = &openapi.LogConfig{
		MaxFiles:      &logConfigMaxFile,
		MaxFileSizeMB: &logConfigMaxFileSizeMB,
	}
	resourcesCPU      = int32(500)
	resourcesMemoryMB = int32(256)
	version           = int32(0)
	createIndex       = int32(42)
	modifyIndex       = int32(99)
	jobModifyIndex    = int32(99)
	notUnlimited      = false
)

func mockPeriodicJob() *openapi.Job {
	job := mockJob()
	batch := structs.JobTypeBatch
	job.Type = &batch

	enabled := true
	specType := structs.PeriodicSpecCron
	spec := "*/30 * * * *"
	running := structs.JobStatusRunning

	job.Periodic = &openapi.PeriodicConfig{
		Enabled:  &enabled,
		SpecType: &specType,
		Spec:     &spec,
	}

	job.Status = &running
	tg := *job.TaskGroups
	tg[0].Migrate = nil

	return job
}

func mockJobWithDiff() *openapi.Job {
	job := mockJob()

	tgs := *job.TaskGroups
	tgs[0].Tasks = &[]openapi.Task{
		{
			Config: &map[string]interface{}{
				"image": "redis:3.4",
				"ports": []string{dbLabel},
			},
			Driver: &docker,
			Name:   &redis,
		},
	}

	return job
}

func mockJob() *openapi.Job {
	return &openapi.Job{
		Region:      &globalRegion,
		ID:          &id,
		Name:        &jobName,
		Namespace:   &defaultNamespace,
		Type:        &jobTypeService,
		Priority:    &priority,
		AllAtOnce:   &allAtOnce,
		Datacenters: &[]string{"dc1"},
		Constraints: &[]openapi.Constraint{
			{
				LTarget: &lTarget,
				RTarget: &rTarget,
				Operand: &operand,
			},
		},
		TaskGroups: &[]openapi.TaskGroup{
			{
				Name:  &web,
				Count: &count,
				EphemeralDisk: &openapi.EphemeralDisk{
					SizeMB: &sizeMB,
				},
				RestartPolicy: &openapi.RestartPolicy{
					Attempts: &restartPolicyAttempts,
					Interval: &restartPolicyInterval,
					Delay:    &restartPolicyDelay,
					Mode:     &restartPolicyMode,
				},
				ReschedulePolicy: &openapi.ReschedulePolicy{
					Attempts:      &reschedulePolicyAttempts,
					Interval:      &reschedulePolicyInterval,
					Delay:         &reschedulePolicyDelay,
					DelayFunction: &reschedulePolicyDelayFunction,
					Unlimited:     &notUnlimited,
				},
				Migrate: defaultMigrateStrategy,
				Networks: &[]openapi.NetworkResource{
					{
						Mode: &hostMode,
						DynamicPorts: &[]openapi.Port{
							{Label: &httpLabel},
							{Label: &adminLabel},
						},
					},
				},
				Tasks: &[]openapi.Task{
					{
						Name:   &web,
						Driver: &execDriver,
						Config: &map[string]interface{}{
							"command": "/bin/date",
						},
						Env: &map[string]string{
							"FOO": "bar",
						},
						Services: &[]openapi.Service{
							{
								Name:      &frontEndTaskName,
								PortLabel: &httpLabel,
								Tags:      &[]string{"pci:${meta.pci-dss}", "datacenter:${node.datacenter}"},
								Checks: &[]openapi.ServiceCheck{
									{
										Name:     &serviceCheckName,
										Type:     &serviceCheckType,
										Command:  &serviceCheckCommand,
										Args:     &[]string{"${meta.version}"},
										Interval: &serviceCheckInterval,
										Timeout:  &serviceCheckTimeout,
									},
								},
							},
							{
								Name:      &adminTaskName,
								PortLabel: &adminLabel,
							},
						},
						LogConfig: defaultLogConfig,
						Resources: &openapi.Resources{
							CPU:      &resourcesCPU,
							MemoryMB: &resourcesMemoryMB,
						},
						Meta: &map[string]string{
							"foo": "bar",
						},
					},
				},
				Meta: &map[string]string{
					"elb_check_type":     "http",
					"elb_check_interval": "30s",
					"elb_check_min":      "3",
				},
			},
		},
		Meta: &map[string]string{
			"owner": "armon",
		},
		Status:         &pendingStatus,
		Version:        &version,
		CreateIndex:    &createIndex,
		ModifyIndex:    &modifyIndex,
		JobModifyIndex: &jobModifyIndex,
	}
}

var jobHCL = `job "my-job" {
	datacenters = ["dc1"]
	type = "service"
	constraint {
		attribute = "${attr.kernel.name}"
		value = "linux"
	}

	group "web" {
		count = 10
		restart {
			attempts = 3
			interval = "10m"
			delay = "1m"
			mode = "delay"
		}
		task "web" {
			driver = "exec"
			config {
				command = "/bin/date"
			}
			resources {
				cpu = 500
				memory = 256
			}
		}
	}
}
`
