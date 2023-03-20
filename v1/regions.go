// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package v1

import (
	"context"
	"sort"

	client "github.com/flytocolors/nomad-openapi/clients/go/v1"
)

type Regions struct {
	client *Client
}

func (c *Client) Regions() *Regions {
	return &Regions{client: c}
}

func (r *Regions) RegionsApi() *client.RegionsApiService {
	return r.client.apiClient.RegionsApi
}

func (r *Regions) GetRegions(ctx context.Context) (*[]string, OpenAPIError) {
	request := r.RegionsApi().GetRegions(r.client.Ctx)

	result, err := r.client.ExecRequest(ctx, request)
	if err != nil {
		return nil, err
	}

	final := result.([]string)
	// Sort the results, so the output is always consistent.
	sort.Strings(final)
	return &final, nil
}
