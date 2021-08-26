package v1

import (
	"context"
	"fmt"
	"testing"
	"time"

	openapi "github.com/hashicorp/nomad-openapi/v1/client"
	"github.com/hashicorp/nomad/command/agent"
	"github.com/hashicorp/nomad/helper/uuid"
	"github.com/hashicorp/nomad/nomad/mock"
	"github.com/hashicorp/nomad/nomad/structs"
	"github.com/stretchr/testify/require"
)

func postTestJob(s *agent.TestAgent, t *testing.T, job *openapi.Job) {
	client, err := NewTestClient(s)
	require.NoError(t, err)

	if job == nil {
		job = mockJob()
	}
	resp, writeMeta, err := client.Jobs().Post(writeOpts.Ctx(), job)

	require.NoError(t, err)
	require.NotNil(t, resp)
	require.NotNil(t, writeMeta)
}

func TestJobsGet(t *testing.T) {
	t.Parallel()
	httpTest(t, nil, func(s *agent.TestAgent) {
		rpcRegister(t, s, mock.Job())

		client, err := NewTestClient(s)
		require.NoError(t, err)

		jobs, queryMeta, err := client.Jobs().GetJobs(queryOpts.Ctx())
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

		client, err := NewTestClient(s)
		require.NoError(t, err)

		queryJob, queryMeta, err := client.Jobs().GetJob(queryOpts.Ctx(), *job.ID)
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

		client, err := NewTestClient(s)
		require.NoError(t, err)

		resp, meta, err := client.Jobs().Post(writeOpts.Ctx(), job)

		require.NoError(t, err)
		require.NotNil(t, resp)
		require.NotNil(t, meta)
	})
}

func TestPlanJob(t *testing.T) {
	t.Parallel()
	httpTest(t, nil, func(s *agent.TestAgent) {
		postTestJob(s, t, nil)

		client, err := NewTestClient(s)
		require.NoError(t, err)

		diffJob := mockJobWithDiff()

		response, meta, err := client.Jobs().Plan(writeOpts.Ctx(), diffJob, true)
		require.NoError(t, err)
		require.NotNil(t, response)
		require.NotNil(t, meta)
	})
}

func TestJobDelete(t *testing.T) {
	t.Parallel()
	httpTest(t, nil, func(s *agent.TestAgent) {
		job := mock.Job()
		rpcRegister(t, s, job)

		// Make the HTTP request to do a soft delete
		client, err := NewTestClient(s)
		require.NoError(t, err)

		response, writeMeta, err := client.Jobs().Delete(writeOpts.Ctx(), job.ID, false, false)
		require.NoError(t, err)
		require.NotNil(t, response)
		require.NotNil(t, writeMeta)

		// Check the response
		require.NotNil(t, response.EvalID)
		require.NotEmpty(t, *response.EvalID)

		rpcJob := rpcGetJob(t, s, job.ID, job.Region, job.Namespace)
		require.NotNil(t, rpcJob)
		require.True(t, rpcJob.Stop)

		response, writeMeta, err = client.Jobs().Delete(writeOpts.Ctx(), job.ID, true, false)
		require.NoError(t, err)
		require.NotNil(t, response)
		require.NotNil(t, writeMeta)

		// Check the response
		require.NotNil(t, response.EvalID)
		require.NotEmpty(t, *response.EvalID)

		// Check the job is gone
		rpcJob = rpcGetJob(t, s, job.ID, job.Region, job.Namespace)
		require.Nil(t, rpcJob)
	})
}

func TestJobParse(t *testing.T) {
	t.Parallel()
	httpTest(t, nil, func(s *agent.TestAgent) {
		client, err := NewTestClient(s)
		require.NoError(t, err)

		job, err := client.Jobs().Parse(context.Background(), jobHCL, false, false)
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

		client, err := NewTestClient(s)
		require.NoError(t, err)

		// Make the HTTP request
		result, meta, err := client.Jobs().Evaluate(writeOpts.Ctx(), *job.ID, false)
		require.NoError(t, err)
		require.NotNil(t, meta)
		require.NotNil(t, result)

		// Check the response
		require.NotNil(t, result.EvalID)
		require.NotEmpty(t, *result.EvalID)
	})
}

func TestJobPeriodicForce(t *testing.T) {
	t.Parallel()
	httpTest(t, nil, func(s *agent.TestAgent) {
		// Create and register a periodic job.
		job := mockPeriodicJob()
		postTestJob(s, t, job)

		client, err := NewTestClient(s)
		require.NoError(t, err)

		// Make the HTTP request
		result, meta, err := client.Jobs().PeriodicForce(writeOpts.Ctx(), *job.ID)
		require.NoError(t, err)
		require.NotNil(t, meta)

		// Check the response
		require.NotNil(t, result.EvalID)
		require.NotEmpty(t, *result.EvalID)
	})
}

func TestJobSummary(t *testing.T) {
	t.Parallel()
	httpTest(t, nil, func(s *agent.TestAgent) {
		job := mockJob()
		postTestJob(s, t, job)

		client, err := NewTestClient(s)
		require.NoError(t, err)

		result, meta, err := client.Jobs().Summary(queryOpts.Ctx(), *job.ID)
		require.NoError(t, err)
		require.NotNil(t, meta)
		require.NotNil(t, result)
		require.Equal(t, *job.ID, *result.JobID)
	})
}

func TestJobDispatch(t *testing.T) {
	t.Parallel()
	httpTest(t, nil, func(s *agent.TestAgent) {
		job := mock.BatchJob()
		job.ParameterizedJob = &structs.ParameterizedJobConfig{}

		rpcRegister(t, s, job)

		client, err := NewTestClient(s)
		require.NoError(t, err)

		result, meta, err := client.Jobs().Dispatch(writeOpts.Ctx(), job.ID, "", nil)
		require.NoError(t, err)
		require.NotNil(t, meta)
		require.NotEmpty(t, *result.EvalID)
		require.NotEmpty(t, *result.DispatchedJobID)
	})
}

func TestJobVersions(t *testing.T) {
	t.Parallel()
	httpTest(t, nil, func(s *agent.TestAgent) {
		// Create the job
		job := mock.Job()
		rpcRegister(t, s, job)

		job2 := mock.Job()
		job2.ID = job.ID
		job2.Priority = 100

		rpcRegister(t, s, job2)

		client, err := NewTestClient(s)
		require.NoError(t, err)

		// Make the HTTP request
		result, meta, err := client.Jobs().Versions(queryOpts.Ctx(), job.ID, true)
		require.NoError(t, err)
		require.NotNil(t, meta)

		versions := *result.Versions
		require.Len(t, versions, 2)

		v1 := versions[0]
		v0 := versions[1]
		require.Equal(t, int32(1), *v1.Version)
		require.Equal(t, int32(100), *v1.Priority)
		require.Equal(t, int32(0), *v0.Version)
		require.Len(t, *result.Diffs, 1)
	})
}

func TestJobRevert(t *testing.T) {
	t.Parallel()
	httpTest(t, nil, func(s *agent.TestAgent) {
		job := mockJob()
		postTestJob(s, t, job)

		client, err := NewTestClient(s)
		require.NoError(t, err)

		dcs := *job.Datacenters
		dcs = append(dcs, "foo")
		job.Datacenters = &dcs
		_, _, err = client.Jobs().Post(writeOpts.Ctx(), job)
		require.NoError(t, err)

		result, meta, err := client.Jobs().Revert(writeOpts.Ctx(), *job.ID, 0, 0, "", "")
		require.NoError(t, err)
		require.NotNil(t, meta)
		require.NotEmpty(t, *result.EvalID)
	})
}

func TestJobDeployment(t *testing.T) {
	t.Parallel()
	httpTest(t, nil, func(s *agent.TestAgent) {
		job := mock.Job()
		rpcRegister(t, s, job)

		// Directly manipulate the state
		state := s.Agent.Server().State()
		deployment := mock.Deployment()
		deployment.JobID = job.ID
		deployment.JobCreateIndex = job.JobModifyIndex
		require.NoError(t, state.UpsertDeployment(1000, deployment))

		client, err := NewTestClient(s)
		require.NoError(t, err)
		// Make the HTTP request
		result, meta, err := client.Jobs().Deployment(queryOpts.Ctx(), job.ID)
		require.NoError(t, err)
		require.NotNil(t, meta)

		// Check the response
		require.NotNil(t, result.ID)
		require.Equal(t, deployment.ID, *result.ID)
	})
}

func TestJobDeployments(t *testing.T) {
	t.Parallel()
	httpTest(t, nil, func(s *agent.TestAgent) {
		job := mock.Job()
		rpcRegister(t, s, job)

		client, err := NewTestClient(s)
		require.NoError(t, err)

		// Directly manipulate the state
		state := s.Agent.Server().State()
		deployment := mock.Deployment()
		deployment.JobID = job.ID
		deployment.JobCreateIndex = job.JobModifyIndex

		require.Nil(t, state.UpsertDeployment(1000, deployment), "UpsertDeployment")

		// Make the HTTP request
		result, meta, err := client.Jobs().Deployments(queryOpts.Ctx(), job.ID)
		require.NoError(t, err)
		require.NotNil(t, meta)

		// Check the response
		deployments := *result
		require.Len(t, deployments, 1)
		require.Equal(t, deployment.ID, *deployments[0].ID)
	})
}

func TestJobStable(t *testing.T) {
	t.Parallel()
	httpTest(t, nil, func(s *agent.TestAgent) {
		// Create the job and register it twice
		job := mock.Job()
		rpcRegister(t, s, job)
		rpcRegister(t, s, job)

		client, err := NewTestClient(s)
		require.NoError(t, err)

		// Make the HTTP request
		result, meta, err := client.Jobs().Stability(queryOpts.Ctx(), job.ID, 0, true)
		require.NoError(t, err)
		require.NotNil(t, meta)

		// Check the response
		require.NotNil(t, result.Index)
		require.NotEqual(t, 0, *result.Index)
	})
}

func TestJobScaleStatus(t *testing.T) {
	t.Parallel()
	httpTest(t, nil, func(s *agent.TestAgent) {
		job := mock.Job()
		rpcRegister(t, s, job)

		client, err := NewTestClient(s)
		require.NoError(t, err)

		// Make the HTTP request to scale the job group
		result, meta, err := client.Jobs().ScaleStatus(queryOpts.Ctx(), job.ID)
		require.NoError(t, err)
		require.NotNil(t, meta)

		// Check the response
		status := *result
		tg := job.TaskGroups
		statusTG := *status.TaskGroups
		require.Equal(t, int32(tg[0].Count), *statusTG[tg[0].Name].Desired)
	})
}

func TestJobScaleTaskGroup(t *testing.T) {
	t.Parallel()
	httpTest(t, nil, func(s *agent.TestAgent) {
		job := mock.Job()
		rpcRegister(t, s, job)

		client, err := NewTestClient(s)
		require.NoError(t, err)

		tg := job.TaskGroups
		newCount := int64(tg[0].Count + 1)

		result, meta, err := client.Jobs().Scale(writeOpts.Ctx(), job.ID, newCount, "testing", map[string]string{
			"Job":   job.ID,
			"Group": tg[0].Name,
		})

		require.NoError(t, err)
		require.NotNil(t, meta)

		// Check the response
		require.NotEmpty(t, *result.EvalID)

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
		require.NoError(t, err)
		require.NotNil(t, getResp.Job)
		require.Equal(t, newCount, int64(getResp.Job.TaskGroups[0].Count))
	})
}

var (
	id                                  = "mock-service"
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
	jobID := fmt.Sprintf("%s-%s", id, uuid.Generate())
	return &openapi.Job{
		Region:      &globalRegion,
		ID:          &jobID,
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

func rpcGetJob(t *testing.T, s *agent.TestAgent, jobID, region, namespace string) *structs.Job {
	rpcRequest := structs.JobSpecificRequest{
		JobID: jobID,
		QueryOptions: structs.QueryOptions{
			Region:    region,
			Namespace: namespace,
		},
	}

	var rpcResponse structs.SingleJobResponse

	err := s.Agent.RPC("Job.GetJob", &rpcRequest, &rpcResponse)
	require.NoError(t, err)

	return rpcResponse.Job
}

// rpcRegister is used to register jobs directly with the RPC layer. This is
// useful for scenarios where the test needs to ensure the job is fully registered
// to prevent race conditions.
func rpcRegister(t *testing.T, s *agent.TestAgent, job *structs.Job) {
	request := structs.JobRegisterRequest{
		Job: job,
		WriteRequest: structs.WriteRequest{
			Region:    "global",
			Namespace: structs.DefaultNamespace,
		},
	}
	var response structs.JobRegisterResponse
	require.NoError(t, s.Agent.RPC("Job.Register", &request, &response))
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
