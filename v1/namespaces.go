package v1

import (
	"context"
	"errors"

	client "github.com/hashicorp/nomad-openapi/clients/go/v1"
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

	meta, err := n.client.ExecNoResponseWrite(ctx, request)
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
	result, meta, err := n.client.ExecQuery(ctx, request)
	if err != nil {
		return nil, nil, err
	}

	final := result.(client.Namespace)
	return &final, meta, nil
}
func (n *Namespaces) GetNamespaces(ctx context.Context) (*[]client.Namespace, *QueryMeta, error) {
	request := n.NamespacesApi().GetNamespaces(n.client.Ctx)
	result, meta, err := n.client.ExecQuery(ctx, request)
	if err != nil {
		return nil, nil, err
	}

	final := result.([]client.Namespace)
	return &final, meta, nil
}

func (n *Namespaces) PostNamespace(ctx context.Context, namespace *client.Namespace) (*WriteMeta, error) {
	if namespace == nil {
		return nil, errors.New("namespace is required")
	}
	request := n.NamespacesApi().PostNamespace(n.client.Ctx, *namespace.Name)

	request = request.Namespace2(*namespace)

	meta, err := n.client.ExecNoResponseWrite(ctx, request)
	if err != nil {
		return nil, err
	}

	return meta, nil
}
