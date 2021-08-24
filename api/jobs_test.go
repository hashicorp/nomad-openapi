package api

import (
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
		result, meta, err := client.Jobs().Evaluate(*job.ID, false)
		require.NoError(t, err)
		require.NotNil(t, meta)
		require.NotNil(t, result)

		// Check the response
		require.NotNil(t, result.EvalID)
		require.NotEqual(t, "", *result.EvalID)
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
		result, meta, err := client.Jobs().PeriodicForce(*job.ID)
		require.NoError(t, err)
		require.NotNil(t, meta)

		// Check the response
		require.NotNil(t, result.EvalID)
		require.NotEqual(t, "", *result.EvalID)
	})
}

func TestJobSummary(t *testing.T) {
	t.Parallel()
	httpTest(t, nil, func(s *agent.TestAgent) {
		job := mockJob()
		postTestJob(s, t, job)

		client, err := NewTestQueryClient(s, queryOpts)
		require.NoError(t, err)

		result, meta, err := client.Jobs().Summary(*job.ID)
		require.NoError(t, err)
		require.NotNil(t, meta)
		require.NotNil(t, result)
		require.Equal(t, *job.ID, *result.JobID)
	})
}

func TestJobDispatch(t *testing.T) {
	t.Parallel()
	httpTest(t, nil, func(s *agent.TestAgent) {
		job := mockJob()
		postTestJob(s, t, job)

		client, err := NewTestWriteClient(s, writeOpts)
		require.NoError(t, err)

		result, meta, err := client.Jobs().Dispatch(*job.ID, "", nil)
		require.NoError(t, err)
		require.NotNil(t, meta)
		require.NotEqual(t, "", *result.EvalID)
		require.NotEqual(t, "", *result.DispatchedJobID)
	})
}

func TestJobVersions(t *testing.T) {
	t.Parallel()
	httpTest(t, nil, func(s *agent.TestAgent) {
		job := mockJob()
		p1 := int32(1)
		job.Priority = &p1
		postTestJob(s, t, job)

		client, err := NewTestWriteClient(s, writeOpts)
		require.NoError(t, err)

		p100 := int32(100)
		job.Priority = &p100
		_, _, err = client.Jobs().Register(job, &RegisterOpts{})
		require.NoError(t, err)

		// Make the HTTP request
		result, meta, err := client.Jobs().Versions(*job.ID, true)
		require.NoError(t, err)
		require.NotNil(t, meta)
		versions := *result.Versions
		require.Len(t, versions, 2)
		v1 := versions[0]
		v0 := versions[1]
		require.Equal(t, 1, v1.Version)
		require.Equal(t, 100, v1.Priority)
		require.Equal(t, 0, v0.Version)
		require.NotEqual(t, 1, *result.Diffs)
	})
}

func TestJobRevert(t *testing.T) {
	t.Parallel()
	httpTest(t, nil, func(s *agent.TestAgent) {
		job := mockJob()
		postTestJob(s, t, job)

		client, err := NewTestWriteClient(s, writeOpts)
		require.NoError(t, err)

		dcs := *job.Datacenters
		dcs = append(dcs, "foo")
		job.Datacenters = &dcs
		_, _, err = client.Jobs().Post(job)
		require.NoError(t, err)

		result, meta, err := client.Jobs().Revert(*job.ID, 0)
		require.NoError(t, err)
		require.NotNil(t, meta)
		require.NotEqual(t, "", *result.EvalID)
	})
}

func TestJobDeployment(t *testing.T) {
	t.Parallel()
	httpTest(t, nil, func(s *agent.TestAgent) {
		job := mockJob()
		postTestJob(s, t, job)

		client, err := NewTestWriteClient(s, writeOpts)
		require.NoError(t, err)

		// Make the HTTP request
		result, meta, err := client.Jobs().Deployment(*job.ID)
		require.NoError(t, err)
		require.NotNil(t, meta)

		// Check the response
		require.Equal(t, *job.ID, *result.JobID)

	})
}

func TestJobDeployments(t *testing.T) {
	t.Parallel()
	httpTest(t, nil, func(s *agent.TestAgent) {
		job := mockJob()
		postTestJob(s, t, job)

		client, err := NewTestQueryClient(s, queryOpts)
		require.NoError(t, err)

		// Directly manipulate the state
		state := s.Agent.Server().State()
		d := mock.Deployment()
		d.JobID = *job.ID
		d.JobCreateIndex = uint64(*job.JobModifyIndex)

		require.Nil(t, state.UpsertDeployment(1000, d), "UpsertDeployment")

		// Make the HTTP request
		result, meta, err := client.Jobs().Deployments(*job.ID)
		require.NoError(t, err)
		require.NotNil(t, meta)

		// Check the response
		deployments := *result
		require.Len(t, deployments, 1)
		require.Equal(t, d.ID, deployments[0].ID)

	})
}

func TestJobStable(t *testing.T) {
	t.Parallel()
	httpTest(t, nil, func(s *agent.TestAgent) {
		// Create the job and register it twice
		job := mockJob()
		postTestJob(s, t, job)
		postTestJob(s, t, job)

		client, err := NewTestWriteClient(s, writeOpts)
		require.NoError(t, err)

		// Make the HTTP request
		result, meta, err := client.Jobs().Stability(*job.ID, 0, true)
		require.NoError(t, err)
		require.NotNil(t, meta)

		// Check the response
		require.NotEqual(t, 0, *result.LastIndex)
	})
}

func TestJobScaleStatus(t *testing.T) {
	t.Parallel()
	httpTest(t, nil, func(s *agent.TestAgent) {
		job := mockJob()
		postTestJob(s, t, job)

		client, err := NewTestQueryClient(s, queryOpts)
		require.NoError(t, err)

		// Make the HTTP request to scale the job group
		result, meta, err := client.Jobs().ScaleStatus(*job.ID)
		require.NoError(t, err)
		require.NotNil(t, meta)

		// Check the response
		status := *result
		tg := *job.TaskGroups
		statusTG := *status.TaskGroups
		require.Equal(t, *tg[0].Count, statusTG[*tg[0].Name].Desired)
	})
}

func TestJobScaleTaskGroup(t *testing.T) {
	t.Parallel()
	httpTest(t, nil, func(s *agent.TestAgent) {
		job := mockJob()
		postTestJob(s, t, job)

		client, err := NewTestWriteClient(s, writeOpts)
		require.NoError(t, err)

		tg := *job.TaskGroups
		newCount := int64(*tg[0].Count + 1)

		result, meta, err := client.Jobs().Scale(*job.ID, newCount, "testing", map[string]string{
			"Job":   *job.ID,
			"Group": *tg[0].Name,
		})

		require.NoError(t, err)
		require.NotNil(t, meta)

		// Check the response
		require.NotEmpty(t, *result.EvalID)

		// Check that the group count was changed
		getReq := structs.JobSpecificRequest{
			JobID: *job.ID,
			QueryOptions: structs.QueryOptions{
				Region:    "global",
				Namespace: structs.DefaultNamespace,
			},
		}
		var getResp structs.SingleJobResponse
		err = s.Agent.RPC("Job.GetJob", &getReq, &getResp)
		require.NoError(t, err)
		require.NotNil(t, getResp.Job)
		require.Equal(t, newCount, getResp.Job.TaskGroups[0].Count)
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
