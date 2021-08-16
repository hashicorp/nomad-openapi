package api

import (
	"strconv"
	"testing"

	"github.com/hashicorp/nomad-openapi/api/test"
	openapi "github.com/hashicorp/nomad-openapi/v1/client"
	"github.com/stretchr/testify/require"
)

func Test_JobsGet(t *testing.T) {
	t.Parallel()
	test.httpTest(t, nil, func(s *test.TestAgent) {

		client, ctx := openapi.NewClientAndContext(s.Config.BindAddr, strconv.Itoa(s.Config.Ports.HTTP))
		jobsRequest := client.JobsApi.JobsGet(ctx)
		jobs, apiResponse, err := client.JobsApi.JobsGetExecute(jobsRequest)

		require.NoError(t, err)
		require.Len(t, jobs, 3)

		test.ValidateMetaHeaders(t, apiResponse)
	})
}

func TestPostJob(t *testing.T) {
	t.Parallel()
	test.httpTest(t, nil, func(s *test.TestAgent) {
		job := getJob()

		client, err := NewWriteClient(&WriteOptions{})
		require.NoError(t, err)

		request := *openapi.NewJobRegisterRequest()
		request.SetJob(job)

		resp, r, err := client.oapiClient.JobsApi.JobsPost(client.Ctx).JobRegisterRequest(request).Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		require.NotNil(t, r)

		require.Equal(t, 200, r.StatusCode)
	})
}

func TestPlanJob(t *testing.T) {
	t.Parallel()
	test.httpTest(t, nil, func(s *test.TestAgent) {
		job := getJob()
		client, err := NewWriteClient(&WriteOptions{})
		require.NoError(t, err)

		request := *openapi.NewJobRegisterRequest()
		request.SetJob(job)

		resp, r, err := client.oapiClient.JobsApi.JobsPost(client.Ctx).JobRegisterRequest(request).Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		require.NotNil(t, r)

		require.Equal(t, 200, r.StatusCode)

		planRequest := *openapi.NewJobPlanRequest()
		planRequest.SetJob(getJobWithDiff())

		planResp, _, err := client.Jobs().Plan(planRequest.Job, true)

		require.NoError(t, err)
		require.NotNil(t, planResp)
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
