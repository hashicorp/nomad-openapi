package v1

import (
	"context"
	"errors"

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

func (p *Plugins) List(ctx context.Context) (*[]client.CSIPluginListStub, *QueryMeta, error) {
	request := p.PluginsApi().GetPlugins(p.client.Ctx)
	result, meta, err := p.client.ExecQuery(ctx, request)
	if err != nil {
		return nil, nil, err
	}

	final := result.([]client.CSIPluginListStub)
	return &final, meta, nil
}

func (p *Plugins) CSI(ctx context.Context, pluginId string) (*[]client.CSIPlugin, *QueryMeta, error) {
	if pluginId == "" {
		return nil, nil, errors.New("plugin id is required")
	}

	request := p.PluginsApi().GetPluginCSI(p.client.Ctx, pluginId)
	result, meta, err := p.client.ExecQuery(ctx, request)
	if err != nil {
		return nil, nil, err
	}

	final := result.([]client.CSIPlugin)
	return &final, meta, nil
}
