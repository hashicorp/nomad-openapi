package v1

import (
	"context"

	client "github.com/hashicorp/nomad-openapi/clients/go/v1"
)

type Plugins struct {
	client *Client
}

func (c *Client) Plugins() *Plugins {
	return &Plugins{client: c}
}

func (p *Plugins) PluginsApi() *client.PluginsApiService {
	return p.client.apiClient.PluginsApi
}

func (p *Plugins) Get(ctx context.Context) (*[]client.CSIPluginListStub, OpenAPIError) {
	request := p.PluginsApi().GetPlugins(p.client.Ctx)
	result, err := p.client.ExecRequest(ctx, request)
	if err != nil {
		return nil, err
	}

	final := result.([]client.CSIPluginListStub)
	return &final, nil
}

func (p *Plugins) GetPlugin(ctx context.Context, pluginID string) (*[]client.CSIPlugin, *QueryMeta, OpenAPIError) {
	if pluginID == "" {
		return nil, nil, &APIError{error: "plugin id is required"}
	}

	request := p.PluginsApi().GetPluginCSI(p.client.Ctx, pluginID)
	result, meta, err := p.client.ExecQuery(ctx, request)
	if err != nil {
		return nil, nil, err
	}

	final := result.([]client.CSIPlugin)
	return &final, meta, nil
}
