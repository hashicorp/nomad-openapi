package v1

import (
	"context"

	"github.com/hashicorp/nomad-openapi/v1/client"
)

type Agent struct {
	client *Client
}

func (c *Client) Agent() *Agent {
	return &Agent{client: c}
}

func (a *Agent) AgentApi() *client.AgentApiService {
	return a.client.apiClient.AgentApi
}

func (a *Agent) ForceLeave(ctx context.Context) error {

	request := a.AgentApi().PostAgentForceLeave(a.client.Ctx)
	request = a.client.setWriteOptions(ctx, request).(client.ApiPostAgentForceLeaveRequest)

	_, err := request.Execute()
	if err != nil {
		return err
	}

	return nil
}

func (a *Agent) Health(ctx context.Context, agentTypes []string) (*client.AgentHealthResponse, error) {

	request := a.AgentApi().GetAgentHealth(a.client.Ctx)
	request = a.client.setQueryOptions(ctx, request).(client.ApiGetAgentHealthRequest)

	request = request.Type_(agentTypes)

	result, _, err := request.Execute()
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (a *Agent) Host(ctx context.Context) (*client.HostDataResponse, error) {

	request := a.AgentApi().GetAgentHostData(a.client.Ctx)
	request = a.client.setQueryOptions(ctx, request).(client.ApiGetAgentHostDataRequest)

	result, _, err := request.Execute()
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (a *Agent) Join(ctx context.Context, addresses []string) (*client.JoinResponse, error) {

	request := a.AgentApi().PostAgentJoin(a.client.Ctx)
	request = a.client.setWriteOptions(ctx, request).(client.ApiPostAgentJoinRequest)

	request = request.Address(addresses)

	result, _, err := request.Execute()
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (a *Agent) Keyring(ctx context.Context, action string, key string) (*client.KeyringResponse, error) {

	request := a.AgentApi().PostAgentKeyringAction(a.client.Ctx, action)
	request = a.client.setWriteOptions(ctx, request).(client.ApiPostAgentKeyringActionRequest)

	keyringRequest := client.NewKeyringRequest()
	keyringRequest.SetKey(key)

	request = request.KeyringRequest(*keyringRequest)

	result, _, err := request.Execute()
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (a *Agent) Members(ctx context.Context) (*client.ServerMembers, error) {

	request := a.AgentApi().GetAgentMembers(a.client.Ctx)
	request = a.client.setQueryOptions(ctx, request).(client.ApiGetAgentMembersRequest)

	result, _, err := request.Execute()
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (a *Agent) Self(ctx context.Context) (*client.AgentSelf, error) {
	request := a.AgentApi().GetAgentSelf(a.client.Ctx)
	request = a.client.setQueryOptions(ctx, request).(client.ApiGetAgentSelfRequest)

	result, _, err := request.Execute()
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (a *Agent) Servers(ctx context.Context) (*[]string, error) {
	request := a.AgentApi().GetAgentServers(a.client.Ctx)
	request = a.client.setQueryOptions(ctx, request).(client.ApiGetAgentServersRequest)

	result, _, err := request.Execute()
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (a *Agent) UpdateServers(ctx context.Context, addresses []string) error {
	request := a.AgentApi().PostAgentServers(a.client.Ctx)
	request = a.client.setWriteOptions(ctx, request).(client.ApiPostAgentServersRequest)

	request = request.Address(addresses)

	_, err := request.Execute()
	if err != nil {
		return err
	}

	return nil
}
