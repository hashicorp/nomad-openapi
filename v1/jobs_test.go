package v1

import (
	"context"
	"fmt"
	"testing"
	"time"

	client "github.com/hashicorp/nomad-openapi/clients/go/v1"
	"github.com/hashicorp/nomad/command/agent"
	"github.com/hashicorp/nomad/helper/uuid"
	"github.com/hashicorp/nomad/nomad/mock"
	"github.com/hashicorp/nomad/nomad/structs"
	"github.com/stretchr/testify/require"
)

func postTestJob(s *agent.TestAgent, t *testing.T, job *client.Job) {
	testClient, err := NewTestClient(s)
	require.NoError(t, err)

	if job == nil {
		job = mockJob()
	}
	result, meta, err := testClient.Jobs().Post(writeOpts.Ctx(), job)

	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, meta)
}

func TestJobsGet(t *testing.T) {
	httpTest(t, nil, func(s *agent.TestAgent) {
		rpcRegister(t, s, mock.Job())

		testClient, err := NewTestClient(s)
		require.NoError(t, err)

		result, meta, err := testClient.Jobs().GetJobs(queryOpts.Ctx())
		require.NoError(t, err)
		require.NotNil(t, result)
		require.Len(t, *result, 1)
		require.NotNil(t, meta)
	})
}

func TestJobGet(t *testing.T) {
	httpTest(t, nil, func(s *agent.TestAgent) {
		job := mockJob()
		postTestJob(s, t, job)

		testClient, err := NewTestClient(s)
		require.NoError(t, err)

		result, meta, err := testClient.Jobs().GetJob(queryOpts.Ctx(), *job.ID)
		require.NoError(t, err)
		require.NotNil(t, result)
		require.NotNil(t, meta)

		if *job.ID != *result.ID {
			t.Fatalf("TestJobGet invalid job comparison: %s != %s", *job.ID, *result.ID)
		}
	})
}

func TestPostJob(t *testing.T) {
	httpTest(t, nil, func(s *agent.TestAgent) {
		job := mockJob()

		testClient, err := NewTestClient(s)
		require.NoError(t, err)

		result, meta, err := testClient.Jobs().Post(writeOpts.Ctx(), job)

		require.NoError(t, err)
		require.NotNil(t, result)
		require.NotNil(t, meta)
	})
}

func TestPlanJob(t *testing.T) {
	httpTest(t, nil, func(s *agent.TestAgent) {
		postTestJob(s, t, nil)

		testClient, err := NewTestClient(s)
		require.NoError(t, err)

		diffJob := mockJobWithDiff()

		result, meta, err := testClient.Jobs().Plan(writeOpts.Ctx(), diffJob, true)
		require.NoError(t, err)
		require.NotNil(t, result)
		require.NotNil(t, meta)
	})
}

func TestJobDelete(t *testing.T) {
	httpTest(t, nil, func(s *agent.TestAgent) {
		job := mock.Job()
		rpcRegister(t, s, job)

		// Make the HTTP request to do a soft delete
		testClient, err := NewTestClient(s)
		require.NoError(t, err)

		result, meta, err := testClient.Jobs().Delete(writeOpts.Ctx(), job.ID, false, false)
		require.NoError(t, err)
		require.NotNil(t, result)
		require.NotNil(t, meta)

		// Check the result
		require.NotNil(t, result.EvalID)
		require.NotEmpty(t, *result.EvalID)

		rpcJob := rpcGetJob(t, s, job.ID, job.Region, job.Namespace)
		require.NotNil(t, rpcJob)
		require.True(t, rpcJob.Stop)

		result, meta, err = testClient.Jobs().Delete(writeOpts.Ctx(), job.ID, true, false)
		require.NoError(t, err)
		require.NotNil(t, result)
		require.NotNil(t, meta)

		// Check the result
		require.NotNil(t, result.EvalID)
		require.NotEmpty(t, *result.EvalID)

		// Check the job is gone
		rpcJob = rpcGetJob(t, s, job.ID, job.Region, job.Namespace)
		require.Nil(t, rpcJob)
	})
}

func TestJobParse(t *testing.T) {
	httpTest(t, nil, func(s *agent.TestAgent) {
		testClient, err := NewTestClient(s)
		require.NoError(t, err)

		result, err := testClient.Jobs().Parse(context.Background(), jobHCL, false, false)
		require.NoError(t, err)
		require.NotNil(t, result)

		expected := mockJob()
		require.NotNil(t, result.Name)
		require.Equal(t, *result.Name, *expected.Name)

		if result.Datacenters == nil {
			expectedDatacenters := *expected.Datacenters
			jobDatacenters := *result.Datacenters
			require.NotEqual(t, jobDatacenters[0], expectedDatacenters[0])
		}
	})
}

func TestJobEvaluate(t *testing.T) {
	httpTest(t, nil, func(s *agent.TestAgent) {
		job := mockJob()
		postTestJob(s, t, job)

		testClient, err := NewTestClient(s)
		require.NoError(t, err)

		// Make the HTTP request
		result, meta, err := testClient.Jobs().Evaluate(writeOpts.Ctx(), *job.ID, false)
		require.NoError(t, err)
		require.NotNil(t, meta)
		require.NotNil(t, result)

		// Check the response
		require.NotNil(t, result.EvalID)
		require.NotEmpty(t, *result.EvalID)
	})
}

func TestJobPeriodicForce(t *testing.T) {
	httpTest(t, nil, func(s *agent.TestAgent) {
		// Create and register a periodic job.
		job := mockPeriodicJob()
		postTestJob(s, t, job)

		testClient, err := NewTestClient(s)
		require.NoError(t, err)

		// Make the HTTP request
		result, meta, err := testClient.Jobs().PeriodicForce(writeOpts.Ctx(), *job.ID)
		require.NoError(t, err)
		require.NotNil(t, meta)

		// Check the response
		require.NotNil(t, result.EvalID)
		require.NotEmpty(t, *result.EvalID)
	})
}

func TestJobSummary(t *testing.T) {
	httpTest(t, nil, func(s *agent.TestAgent) {
		job := mockJob()
		postTestJob(s, t, job)

		testClient, err := NewTestClient(s)
		require.NoError(t, err)

		result, meta, err := testClient.Jobs().Summary(queryOpts.Ctx(), *job.ID)
		require.NoError(t, err)
		require.NotNil(t, meta)
		require.NotNil(t, result)
		require.Equal(t, *job.ID, *result.JobID)
	})
}

func TestJobDispatch(t *testing.T) {
	httpTest(t, nil, func(s *agent.TestAgent) {
		job := mock.BatchJob()
		job.ParameterizedJob = &structs.ParameterizedJobConfig{}

		rpcRegister(t, s, job)

		testClient, err := NewTestClient(s)
		require.NoError(t, err)

		result, meta, err := testClient.Jobs().Dispatch(writeOpts.Ctx(), job.ID, "", nil)
		require.NoError(t, err)
		require.NotNil(t, meta)
		require.NotEmpty(t, *result.EvalID)
		require.NotEmpty(t, *result.DispatchedJobID)
	})
}

func TestJobVersions(t *testing.T) {
	httpTest(t, nil, func(s *agent.TestAgent) {
		// Create the job
		job := mock.Job()
		rpcRegister(t, s, job)

		job2 := mock.Job()
		job2.ID = job.ID
		job2.Priority = 100

		rpcRegister(t, s, job2)

		testClient, err := NewTestClient(s)
		require.NoError(t, err)

		// Make the HTTP request
		result, meta, err := testClient.Jobs().Versions(queryOpts.Ctx(), job.ID, true)
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
	httpTest(t, nil, func(s *agent.TestAgent) {
		rpcJob := mock.Job()
		rpcRegister(t, s, rpcJob)

		testClient, err := NewTestClient(s)
		require.NoError(t, err)

		job, _, err := testClient.Jobs().GetJob(queryOpts.Ctx(), rpcJob.ID)
		require.NoError(t, err)

		dcs := *job.Datacenters
		dcs = append(dcs, "foo")
		job.Datacenters = &dcs
		_, _, err = testClient.Jobs().Post(writeOpts.Ctx(), job)
		require.NoError(t, err)

		result, meta, err := testClient.Jobs().Revert(writeOpts.Ctx(), *job.ID, 0, 0, "", "")
		require.NoError(t, err)
		require.NotNil(t, meta)
		require.NotEmpty(t, *result.EvalID)
	})
}

func TestJobDeployment(t *testing.T) {
	httpTest(t, nil, func(s *agent.TestAgent) {
		job := mock.Job()
		rpcRegister(t, s, job)

		// Directly manipulate the state
		state := s.Agent.Server().State()
		deployment := mock.Deployment()
		deployment.JobID = job.ID
		deployment.JobCreateIndex = job.JobModifyIndex
		require.NoError(t, state.UpsertDeployment(1000, deployment))

		testClient, err := NewTestClient(s)
		require.NoError(t, err)
		// Make the HTTP request
		result, meta, err := testClient.Jobs().Deployment(queryOpts.Ctx(), job.ID)
		require.NoError(t, err)
		require.NotNil(t, meta)

		// Check the response
		require.NotNil(t, result.ID)
		require.Equal(t, deployment.ID, *result.ID)
	})
}

func TestJobDeployments(t *testing.T) {
	httpTest(t, nil, func(s *agent.TestAgent) {
		job := mock.Job()
		rpcRegister(t, s, job)

		testClient, err := NewTestClient(s)
		require.NoError(t, err)

		// Directly manipulate the state
		state := s.Agent.Server().State()
		deployment := mock.Deployment()
		deployment.JobID = job.ID
		deployment.JobCreateIndex = job.JobModifyIndex

		require.Nil(t, state.UpsertDeployment(1000, deployment), "UpsertDeployment")

		// Make the HTTP request
		result, meta, err := testClient.Jobs().Deployments(queryOpts.Ctx(), job.ID)
		require.NoError(t, err)
		require.NotNil(t, meta)

		// Check the response
		deployments := *result
		require.Len(t, deployments, 1)
		require.Equal(t, deployment.ID, *deployments[0].ID)
	})
}

func TestJobStable(t *testing.T) {
	httpTest(t, nil, func(s *agent.TestAgent) {
		// Create the job and register it twice
		job := mock.Job()
		rpcRegister(t, s, job)
		rpcRegister(t, s, job)

		testClient, err := NewTestClient(s)
		require.NoError(t, err)

		// Make the HTTP request
		result, meta, err := testClient.Jobs().Stability(queryOpts.Ctx(), job.ID, 0, true)
		require.NoError(t, err)
		require.NotNil(t, meta)

		// Check the response
		require.NotNil(t, result.Index)
		require.NotEqual(t, 0, *result.Index)
	})
}

func TestJobScaleStatus(t *testing.T) {
	httpTest(t, nil, func(s *agent.TestAgent) {
		job := mock.Job()
		rpcRegister(t, s, job)

		testClient, err := NewTestClient(s)
		require.NoError(t, err)

		// Make the HTTP request to scale the job group
		result, meta, err := testClient.Jobs().ScaleStatus(queryOpts.Ctx(), job.ID)
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
	httpTest(t, nil, func(s *agent.TestAgent) {
		job := mock.Job()
		rpcRegister(t, s, job)

		testClient, err := NewTestClient(s)
		require.NoError(t, err)

		tg := job.TaskGroups
		newCount := int64(tg[0].Count + 1)

		result, meta, err := testClient.Jobs().Scale(writeOpts.Ctx(), job.ID, newCount, "testing", map[string]string{
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

// TODO: Figure out how to force allocations with the TestAgent
//func TestGetJobAllocations(t *testing.T) {
//
//	agentConfFunc := func(c *agent.Config) {
//		c.Client.Enabled = true
//	}
//
//	httpTest(t, agentConfFunc, func(s *agent.TestAgent) {
//		job := mock.Job()
//		rpcRegister(t, s, job)
//
//		testClient, err := NewTestClient(s)
//		require.NoError(t, err)
//
//		testutil.WaitForResult(func() (bool, error) {
//			result, meta, err := testClient.Jobs().Allocations(queryOpts.Ctx(), job.Name, true)
//			require.NoError(t, err)
//			require.NotNil(t, result)
//			require.NotNil(t, meta)
//
//			if len(*result) < 1 {
//				return false, fmt.Errorf("no allocations yet")
//			}
//
//			allocs := *result
//			alloc := allocs[0]
//
//			if *alloc.ClientStatus != "running" {
//				return false, fmt.Errorf("alloc is not running yet: %v", alloc.ClientStatus)
//			}
//
//			return true, nil
//		}, func(err error) {
//			require.NoError(t, err)
//		})
//	})
//}

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
	defaultMigrateStrategy              = &client.MigrateStrategy{
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
	defaultLogConfig       = &client.LogConfig{
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

func mockPeriodicJob() *client.Job {
	job := mockJob()
	batch := structs.JobTypeBatch
	job.Type = &batch

	enabled := true
	specType := structs.PeriodicSpecCron
	spec := "*/30 * * * *"
	running := structs.JobStatusRunning

	job.Periodic = &client.PeriodicConfig{
		Enabled:  &enabled,
		SpecType: &specType,
		Spec:     &spec,
	}

	job.Status = &running
	tg := *job.TaskGroups
	tg[0].Migrate = nil

	return job
}

func mockJobWithDiff() *client.Job {
	job := mockJob()

	tgs := *job.TaskGroups
	tgs[0].Tasks = &[]client.Task{
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

func mockJob() *client.Job {
	jobID := fmt.Sprintf("%s-%s", id, uuid.Generate())
	return &client.Job{
		Region:      &globalRegion,
		ID:          &jobID,
		Name:        &jobName,
		Namespace:   &defaultNamespace,
		Type:        &jobTypeService,
		Priority:    &priority,
		AllAtOnce:   &allAtOnce,
		Datacenters: &[]string{"dc1"},
		Constraints: &[]client.Constraint{
			{
				LTarget: &lTarget,
				RTarget: &rTarget,
				Operand: &operand,
			},
		},
		TaskGroups: &[]client.TaskGroup{
			{
				Name:  &web,
				Count: &count,
				EphemeralDisk: &client.EphemeralDisk{
					SizeMB: &sizeMB,
				},
				RestartPolicy: &client.RestartPolicy{
					Attempts: &restartPolicyAttempts,
					Interval: &restartPolicyInterval,
					Delay:    &restartPolicyDelay,
					Mode:     &restartPolicyMode,
				},
				ReschedulePolicy: &client.ReschedulePolicy{
					Attempts:      &reschedulePolicyAttempts,
					Interval:      &reschedulePolicyInterval,
					Delay:         &reschedulePolicyDelay,
					DelayFunction: &reschedulePolicyDelayFunction,
					Unlimited:     &notUnlimited,
				},
				Migrate: defaultMigrateStrategy,
				Networks: &[]client.NetworkResource{
					{
						Mode: &hostMode,
						DynamicPorts: &[]client.Port{
							{Label: &httpLabel},
							{Label: &adminLabel},
						},
					},
				},
				Tasks: &[]client.Task{
					{
						Name:   &web,
						Driver: &execDriver,
						Config: &map[string]interface{}{
							"command": "/bin/date",
						},
						Env: &map[string]string{
							"FOO": "bar",
						},
						Services: &[]client.Service{
							{
								Name:      &frontEndTaskName,
								PortLabel: &httpLabel,
								Tags:      &[]string{"pci:${meta.pci-dss}", "datacenter:${node.datacenter}"},
								Checks: &[]client.ServiceCheck{
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
						Resources: &client.Resources{
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
