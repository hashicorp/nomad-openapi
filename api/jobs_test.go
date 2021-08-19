package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	openapi "github.com/hashicorp/nomad-openapi/v1/client"
	"github.com/hashicorp/nomad/api"
	"github.com/hashicorp/nomad/command/agent"
	"github.com/hashicorp/nomad/nomad/mock"
	"github.com/hashicorp/nomad/nomad/structs"
	"github.com/stretchr/testify/require"
)

func TestJobsGet(t *testing.T) {
	t.Parallel()
	httpTest(t, nil, func(s *agent.TestAgent) {
		client, err := NewTestWriteClient(s, writeOpts)
		require.NoError(t, err)
		job := getJob()
		resp, writeMeta, err := client.Jobs().Post(&job)

		require.NoError(t, err)
		require.NotNil(t, resp)
		require.NotNil(t, writeMeta)

		client, err = NewTestQueryClient(s, queryOpts)
		require.NoError(t, err)

		jobs, queryMeta, err := client.Jobs().Get()
		require.NoError(t, err)
		require.NotNil(t, jobs)
		require.Len(t, jobs, 1)
		require.NotNil(t, queryMeta)
	})
}

func TestPostJob(t *testing.T) {
	t.Parallel()
	httpTest(t, nil, func(s *agent.TestAgent) {
		job := getJob()

		client, err := NewTestWriteClient(s, writeOpts)
		require.NoError(t, err)

		resp, meta, err := client.Jobs().Post(&job)

		require.NoError(t, err)
		require.NotNil(t, resp)
		require.NotNil(t, meta)
	})
}

func TestPlanJob(t *testing.T) {
	t.Parallel()
	httpTest(t, nil, func(s *agent.TestAgent) {
		job := getJob()
		client, err := NewTestWriteClient(s, writeOpts)
		require.NoError(t, err)

		request := *openapi.NewJobRegisterRequest()
		request.SetJob(job)

		resp, meta, err := client.Jobs().Post(&job)

		require.NoError(t, err)
		require.NotNil(t, resp)
		require.NotNil(t, meta)

		diffJob := getJobWithDiff()
		response, meta, err := client.Jobs().Plan(&diffJob, true)

		require.NoError(t, err)
		require.NotNil(t, response)
		require.NotNil(t, meta)
	})
}

func TestJobDelete(t *testing.T) {
	t.Parallel()
	httpTest(t, nil, func(s *agent.TestAgent) {
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

		// Make the HTTP request to do a soft delete
		req, err := http.NewRequest("DELETE", "/v1/job/"+job.ID, nil)
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
		dereg := obj.(structs.JobDeregisterResponse)
		if dereg.EvalID == "" {
			t.Fatalf("bad: %v", dereg)
		}

		// Check for the index
		if respW.HeaderMap.Get("X-Nomad-Index") == "" {
			t.Fatalf("missing index")
		}

		// Check the job is still queryable
		getReq1 := structs.JobSpecificRequest{
			JobID: job.ID,
			QueryOptions: structs.QueryOptions{
				Region:    "global",
				Namespace: structs.DefaultNamespace,
			},
		}
		var getResp1 structs.SingleJobResponse
		if err := s.Agent.RPC("Job.GetJob", &getReq1, &getResp1); err != nil {
			t.Fatalf("err: %v", err)
		}
		if getResp1.Job == nil {
			t.Fatalf("job doesn't exists")
		}
		if !getResp1.Job.Stop {
			t.Fatalf("job should be marked as stop")
		}

		// Make the HTTP request to do a purge delete
		req2, err := http.NewRequest("DELETE", "/v1/job/"+job.ID+"?purge=true", nil)
		if err != nil {
			t.Fatalf("err: %v", err)
		}
		respW.Flush()

		// Make the request
		obj, err = s.Server.JobSpecificRequest(respW, req2)
		if err != nil {
			t.Fatalf("err: %v", err)
		}

		// Check the response
		dereg = obj.(structs.JobDeregisterResponse)
		if dereg.EvalID == "" {
			t.Fatalf("bad: %v", dereg)
		}

		// Check for the index
		if respW.HeaderMap.Get("X-Nomad-Index") == "" {
			t.Fatalf("missing index")
		}

		// Check the job is gone
		getReq2 := structs.JobSpecificRequest{
			JobID: job.ID,
			QueryOptions: structs.QueryOptions{
				Region:    "global",
				Namespace: structs.DefaultNamespace,
			},
		}
		var getResp2 structs.SingleJobResponse
		if err := s.Agent.RPC("Job.GetJob", &getReq2, &getResp2); err != nil {
			t.Fatalf("err: %v", err)
		}
		if getResp2.Job != nil {
			t.Fatalf("job still exists")
		}
	})

	t.Parallel()
	httpTest(t, nil, func(s *agent.TestAgent) {
		job := getJob()
		client, err := NewTestWriteClient(s, writeOpts)
		require.NoError(t, err)

		request := *openapi.NewJobRegisterRequest()
		request.SetJob(job)

		resp, meta, err := client.Jobs().Post(&job)

		require.NoError(t, err)
		require.NotNil(t, resp)
		require.NotNil(t, meta)

		response, meta, err := client.Jobs().Delete(*job.ID, true, true)

		require.NoError(t, err)
		require.NotNil(t, response)
		require.NotNil(t, meta)
	})
}

func TestJobParse(t *testing.T) {
	t.Parallel()
	httpTest(t, nil, func(s *agent.TestAgent) {
		buf := encodeReq(api.JobsParseRequest{JobHCL: mock.HCL()})
		req, err := http.NewRequest("POST", "/v1/jobs/parse", buf)
		if err != nil {
			t.Fatalf("err: %v", err)
		}

		respW := httptest.NewRecorder()

		obj, err := s.Server.JobsParseRequest(respW, req)
		if err != nil {
			t.Fatalf("err: %v", err)
		}
		if obj == nil {
			t.Fatal("response should not be nil")
		}

		job := obj.(*api.Job)
		expected := mock.Job()
		if job.Name == nil || *job.Name != expected.Name {
			t.Fatalf("job name is '%s', expected '%s'", *job.Name, expected.Name)
		}

		if job.Datacenters == nil ||
			job.Datacenters[0] != expected.Datacenters[0] {
			t.Fatalf("job datacenters is '%s', expected '%s'",
				job.Datacenters[0], expected.Datacenters[0])
		}
	})
}

func TestHTTP_JobForceEvaluate(t *testing.T) {
	t.Parallel()
	httpTest(t, nil, func(s *agent.TestAgent) {
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

		// Make the HTTP request
		req, err := http.NewRequest("POST", "/v1/job/"+job.ID+"/evaluate", nil)
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
		reg := obj.(structs.JobRegisterResponse)
		if reg.EvalID == "" {
			t.Fatalf("bad: %v", reg)
		}

		// Check for the index
		if respW.HeaderMap.Get("X-Nomad-Index") == "" {
			t.Fatalf("missing index")
		}
	})
}

func TestJobPeriodicForce(t *testing.T) {
	t.Parallel()
	httpTest(t, nil, func(s *agent.TestAgent) {
		// Create and register a periodic job.
		job := mock.PeriodicJob()
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

		// Make the HTTP request
		req, err := http.NewRequest("POST", "/v1/job/"+job.ID+"/periodic/force", nil)
		if err != nil {
			t.Fatalf("err: %v", err)
		}
		respW := httptest.NewRecorder()

		// Make the request
		obj, err := s.Server.JobSpecificRequest(respW, req)
		if err != nil {
			t.Fatalf("err: %v", err)
		}

		// Check for the index
		if respW.HeaderMap.Get("X-Nomad-Index") == "" {
			t.Fatalf("missing index")
		}

		// Check the response
		r := obj.(structs.PeriodicForceResponse)
		if r.EvalID == "" {
			t.Fatalf("bad: %#v", r)
		}
	})
}

func TestJobGet(t *testing.T) {
	t.Parallel()
	httpTest(t, nil, func(s *agent.TestAgent) {
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

		// Make the HTTP request
		req, err := http.NewRequest("GET", "/v1/job/"+job.ID, nil)
		if err != nil {
			t.Fatalf("err: %v", err)
		}
		respW := httptest.NewRecorder()

		// Make the request
		obj, err := s.Server.JobSpecificRequest(respW, req)
		if err != nil {
			t.Fatalf("err: %v", err)
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

		// Check the job
		j := obj.(*structs.Job)
		if j.ID != job.ID {
			t.Fatalf("bad: %#v", j)
		}
	})
}

var (
	id               = "cache"
	dbLabel          = "db"
	toPort     int32 = 6379
	jobName          = "cache"
	redisCache       = "redis-cache"
	docker           = "docker"
	redis            = "redis"
)

func getJobWithDiff() openapi.Job {
	job := getJob()

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

func getJob() openapi.Job {
	return openapi.Job{
		Datacenters: &[]string{"dc1"},
		ID:          &id,
		TaskGroups: &[]openapi.TaskGroup{
			{
				Name: &jobName,
				Networks: &[]openapi.NetworkResource{
					{
						DynamicPorts: &[]openapi.Port{
							{
								Label: &dbLabel,
								To:    &toPort,
							},
						},
					},
				},
				Services: &[]openapi.Service{
					{
						Name:      &redisCache,
						PortLabel: &dbLabel,
					},
				},
				Tasks: &[]openapi.Task{
					{
						Config: &map[string]interface{}{
							"image": "redis:3.2",
							"ports": []string{dbLabel},
						},
						Driver: &docker,
						Name:   &redis,
					},
				},
			},
		},
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
