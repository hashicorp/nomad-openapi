package api

import (
	"testing"

	openapi "github.com/hashicorp/nomad-openapi/v1/client"
	"github.com/hashicorp/nomad/command/agent"
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
