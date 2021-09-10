package v1

import (
	"context"
	"errors"

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

func (n *Namespaces) DeleteNamespace(ctx context.Context, name string) (*WriteMeta, error) {
	if name == "" {
		return nil, errors.New("namespace name is required")
	}
	request := n.NamespacesApi().DeleteNamespace(n.client.Ctx, name)
	request = n.client.setWriteOptions(ctx, request).(client.ApiDeleteNamespaceRequest)

	response, err := request.Execute()
	if err != nil {
		return nil, err
	}

	meta, err := parseWriteMeta(response)
	if err != nil {
		return nil, err
	}

	return meta, nil
}

func (n *Namespaces) GetNamespace(ctx context.Context, name string) (*client.Namespace, *QueryMeta, error) {
	if name == "" {
		return nil, nil, errors.New("namespace name is required")
	}

	request := n.NamespacesApi().GetNamespace(n.client.Ctx, name)
	request = n.client.setQueryOptions(ctx, request).(client.ApiGetNamespaceRequest)

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

func (n *Namespaces) PostNamespace(ctx context.Context, namespace *client.Namespace) (*WriteMeta, error) {
	if namespace == nil {
		return nil, errors.New("namespace is required")
	}
	request := n.NamespacesApi().PostNamespace(n.client.Ctx, *namespace.Name)
	request = n.client.setWriteOptions(ctx, request).(client.ApiPostNamespaceRequest)

	request = request.Namespace2(*namespace)

	response, err := request.Execute()
	if err != nil {
		return nil, err
	}

	meta, err := parseWriteMeta(response)
	if err != nil {
		return nil, err
	}

	return meta, nil
}
