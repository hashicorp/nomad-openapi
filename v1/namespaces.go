// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package v1

import (
	"context"

	client "github.com/flytocolors/nomad-openapi/clients/go/v1"
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

func (n *Namespaces) DeleteNamespace(ctx context.Context, name string) (*WriteMeta, OpenAPIError) {
	if name == "" {
		return nil, &APIError{error: "namespace name is required"}
	}
	request := n.NamespacesApi().DeleteNamespace(n.client.Ctx, name)

	meta, err := n.client.ExecNoResponseWrite(ctx, request)
	if err != nil {
		return nil, err
	}

	return meta, nil
}

func (n *Namespaces) GetNamespace(ctx context.Context, name string) (*client.Namespace, *QueryMeta, OpenAPIError) {
	if name == "" {
		return nil, nil, &APIError{error: "namespace name is required"}
	}

	request := n.NamespacesApi().GetNamespace(n.client.Ctx, name)
	result, meta, err := n.client.ExecQuery(ctx, request)
	if err != nil {
		return nil, nil, err
	}

	final := result.(client.Namespace)
	return &final, meta, nil
}

func (n *Namespaces) GetNamespaces(ctx context.Context) (*[]client.Namespace, *QueryMeta, OpenAPIError) {
	request := n.NamespacesApi().GetNamespaces(n.client.Ctx)
	result, meta, err := n.client.ExecQuery(ctx, request)
	if err != nil {
		return nil, nil, err
	}

	final := result.([]client.Namespace)
	return &final, meta, nil
}

func (n *Namespaces) PostNamespace(ctx context.Context, namespace *client.Namespace) (*WriteMeta, OpenAPIError) {
	if namespace == nil {
		return nil, &APIError{error: "namespace is required"}
	}

	request := n.NamespacesApi().PostNamespace(n.client.Ctx, *namespace.Name)

	request = request.Namespace2(*namespace)

	meta, err := n.client.ExecNoResponseWrite(ctx, request)
	if err != nil {
		return nil, err
	}

	return meta, nil
}
