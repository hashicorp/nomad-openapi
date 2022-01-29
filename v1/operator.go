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

func (o *Operator) Raft(ctx context.Context) (*[]client.RaftConfigurationResponse, error) {
	request := o.OperatorApi().GetOperatorRaftConfiguration(o.client.Ctx)

	result, err := o.client.ExecRequest(ctx, request)
	if err != nil {
		return nil, err
	}

	final := result.([]client.RaftConfigurationResponse)
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

func (o *Operator) UpdateAutopilot(
	ctx context.Context,
	cleanupdeadservers bool,
	lastcontactthreshold string,
	maxtrailinglogs int32,
	serverstabilizationtime string,
	enableredundancyzones bool,
	disableupgrademigration bool,
	enablecustomupgrades bool,
) (*bool, error) {
	request := o.OperatorApi().PutOperatorAutopilotConfiguration(o.client.Ctx)

	// QUESTION: is there a better way to handle this?
	// https://www.nomadproject.io/api-docs/operator/autopilot#update-autopilot-configuration
	updateRequest := client.NewAutopilotConfiguration()
	updateRequest.SetCleanupDeadServers(cleanupdeadservers)
	updateRequest.SetLastContactThreshold(lastcontactthreshold)
	updateRequest.SetMaxTrailingLogs(maxtrailinglogs)
	updateRequest.SetServerStabilizationTime(serverstabilizationtime)
	updateRequest.SetEnableRedundancyZones(enableredundancyzones)
	updateRequest.SetDisableUpgradeMigration(disableupgrademigration)
	updateRequest.SetEnableCustomUpgrades(enablecustomupgrades)

	request = request.AutopilotConfiguration(*updateRequest)

	//request.Execute() // returns *http.Response, error, but where's the bool!?!?
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

// NOTE: there's no RejectJobRegistration?
func (o *Operator) UpdateScheduler(
	ctx context.Context,
	scheduleralgorithm string,
	memoryoversubscriptionenabled bool,
	preemptionconfig *client.PreemptionConfig,
) (*client.SchedulerSetConfigurationResponse, *WriteMeta, error) {
	request := o.OperatorApi().PostOperatorSchedulerConfiguration(o.client.Ctx)

	// QUESTION: how do I pass in the premptionconfig?
	updateRequest := client.NewSchedulerConfiguration()
	updateRequest.SetSchedulerAlgorithm(scheduleralgorithm)
	updateRequest.SetMemoryOversubscriptionEnabled(memoryoversubscriptionenabled)
	updateRequest.SetPreemptionConfig(*preemptionconfig)

	request = request.SchedulerConfiguration(*updateRequest)

	result, meta, err := o.client.ExecWrite(ctx, request)
	if err != nil {
		return nil, nil, err
	}

	final := result.(client.SchedulerSetConfigurationResponse)
	return &final, meta, nil
}
