package v1

import (
	"context"

	"github.com/hashicorp/nomad-openapi/v1/client"
)

type Namespaces struct {
	client *Client
}

func (c *Client) Namespaces() *Namespaces {
	return &Namespaces{client: c}
}

func (n *Namespaces) NamespacesApi() *client.NamespacesApiService {
	return n.client.apiClient.NamespacesApi
}

func (n *Namespaces) GetNamespaces(ctx context.Context) (*[]client.Namespace, *QueryMeta, error) {
	request := n.NamespacesApi().GetNamespaces(n.client.Ctx)
	request = n.client.setQueryOptions(ctx, request).(client.ApiGetNamespacesRequest)

	result, response, err := request.Execute()
	if err != nil {
		return nil, nil, err
	}

	meta, err := parseQueryMeta(response)
	if err != nil {
		return nil, nil, err
	}

	return &result, meta, nil
}
