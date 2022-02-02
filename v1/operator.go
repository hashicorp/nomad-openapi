package v1

import (
	"context"

	client "github.com/hashicorp/nomad-openapi/clients/go/v1"
)

type Operator struct {
	client *Client
}

func (c *Client) Operator() *Operator {
	return &Operator{client: c}
}

func (o *Operator) OperatorApi() *client.OperatorApiService {
	return o.client.apiClient.OperatorApi
}

func (o *Operator) Raft(ctx context.Context) (*client.RaftConfiguration, error) {
	request := o.OperatorApi().GetOperatorRaftConfiguration(o.client.Ctx)

	result, err := o.client.ExecRequest(ctx, request)
	if err != nil {
		return nil, err
	}

	final := result.(client.RaftConfiguration)
	return &final, nil
}

func (o *Operator) Peer(ctx context.Context) error {
	request := o.OperatorApi().DeleteOperatorRaftPeer(o.client.Ctx)

	err := o.client.ExecNoResponseRequest(ctx, request)
	if err != nil {
		return err
	}

	return nil
}

func (o *Operator) Autopilot(ctx context.Context) (*client.AutopilotConfiguration, error) {
	request := o.OperatorApi().GetOperatorAutopilotConfiguration(o.client.Ctx)

	result, err := o.client.ExecRequest(ctx, request)
	if err != nil {
		return nil, err
	}

	final := result.(client.AutopilotConfiguration)
	return &final, nil
}

func (o *Operator) UpdateAutopilot(ctx context.Context, config *client.AutopilotConfiguration) (*bool, error) {
	request := o.OperatorApi().PutOperatorAutopilotConfiguration(o.client.Ctx)

	updateRequest := client.NewAutopilotConfiguration()
	updateRequest.SetCleanupDeadServers(*config.CleanupDeadServers)
	updateRequest.SetLastContactThreshold(*config.LastContactThreshold)
	updateRequest.SetMaxTrailingLogs(*config.MaxTrailingLogs)
	updateRequest.SetServerStabilizationTime(*config.ServerStabilizationTime)
	updateRequest.SetEnableRedundancyZones(*config.EnableRedundancyZones)
	updateRequest.SetDisableUpgradeMigration(*config.DisableUpgradeMigration)
	updateRequest.SetEnableCustomUpgrades(*config.EnableCustomUpgrades)

	request = request.AutopilotConfiguration(*updateRequest)

	result, err := o.client.ExecRequest(ctx, request)
	if err != nil {
		return nil, err
	}

	final := result.(bool)
	return &final, nil
}

func (o *Operator) AutopilotHealth(ctx context.Context) (*client.OperatorHealthReply, error) {
	request := o.OperatorApi().GetOperatorAutopilotHealth(o.client.Ctx)

	result, err := o.client.ExecRequest(ctx, request)
	if err != nil {
		return nil, err
	}

	final := result.(client.OperatorHealthReply)
	return &final, nil
}

func (o *Operator) Scheduler(ctx context.Context) (*client.SchedulerConfigurationResponse, *QueryMeta, error) {
	request := o.OperatorApi().GetOperatorSchedulerConfiguration(o.client.Ctx)

	result, meta, err := o.client.ExecQuery(ctx, request)
	if err != nil {
		return nil, nil, err
	}

	final := result.(client.SchedulerConfigurationResponse)
	return &final, meta, nil
}

// TODO: update Nomad version to pick up new fields (i.e. RejectJobRegistration)
func (o *Operator) UpdateScheduler(ctx context.Context, config *client.SchedulerConfiguration) (*client.SchedulerSetConfigurationResponse, *WriteMeta, error) {
	request := o.OperatorApi().PostOperatorSchedulerConfiguration(o.client.Ctx)

	updateRequest := client.NewSchedulerConfiguration()
	updateRequest.SetSchedulerAlgorithm(*config.SchedulerAlgorithm)
	updateRequest.SetMemoryOversubscriptionEnabled(*config.MemoryOversubscriptionEnabled)
	updateRequest.SetPreemptionConfig(*config.PreemptionConfig)

	request = request.SchedulerConfiguration(*updateRequest)

	result, meta, err := o.client.ExecWrite(ctx, request)
	if err != nil {
		return nil, nil, err
	}

	final := result.(client.SchedulerSetConfigurationResponse)
	return &final, meta, nil
}
