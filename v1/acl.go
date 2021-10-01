package v1

import (
	"context"
	"errors"

	client "github.com/hashicorp/nomad-openapi/clients/go/v1"
)

type ACL struct {
	client *Client
}

func (c *Client) ACL() *ACL {
	return &ACL{client: c}
}

func (a *ACL) ACLApi() *client.ACLApiService {
	return a.client.apiClient.ACLApi
}

// returns arraySchema
func (a *ACL) GetACLPolicies(ctx context.Context) (*[]client.ACLPolicyListStub, *QueryMeta, error) {
	request := a.ACLApi().GetACLPolicies(a.client.Ctx)
	result, meta, err := a.client.ExecQuery(ctx, request)
	if err != nil {
		return nil, nil, err
	}

	final := result.([]client.ACLPolicyListStub)
	return &final, meta, nil
}

// returns objectSchema
func (a *ACL) GetACLPolicy(ctx context.Context, policyName string) (*client.ACLPolicy, *QueryMeta, error) {
	if policyName == "" {
		return nil, nil, errors.New("policy name is required")
	}

	request := a.ACLApi().GetACLPolicy(a.client.Ctx, policyName)
	result, meta, err := a.client.ExecQuery(ctx, request)
	if err != nil {
		return nil, nil, err
	}

	final := result.(client.ACLPolicy)
	return &final, meta, nil
}

// returns nilSchema
func (a *ACL) PostACLPolicy(ctx context.Context, policyName string) (*WriteMeta, error) {
	if policyName == "" {
		return nil, errors.New("policy name is required")
	}

	// appending payload?  something that's taken care of by appendParams so I don't have to worry about it?
	request := a.ACLApi().PostACLPolicy(a.client.Ctx, policyName)
	meta, err := a.client.ExecNoResponseWrite(ctx, request)
	if err != nil {
		return nil, err
	}

	return meta, nil
}

func (a *ACL) DeleteACLPolicy(ctx context.Context, policyName string) (*WriteMeta, error) {
	if policyName == "" {
		return nil, errors.New("policy name is required")
	}

	request := a.ACLApi().DeleteACLPolicy(a.client.Ctx, policyName)
	meta, err := a.client.ExecNoResponseWrite(ctx, request)
	if err != nil {
		return nil, err
	}

	return meta, nil
}

func (a *ACL) PostACLTokenOnetime(ctx context.Context) (*client.OneTimeToken, *WriteMeta, error) {
	request := a.ACLApi().PostACLTokenOnetime(a.client.Ctx)
	result, meta, err := a.client.ExecWrite(ctx, request)
	if err != nil {
		return nil, nil, err
	}

	final := result.(client.OneTimeToken)
	return &final, meta, nil
}

func (a *ACL) PostACLTokenOnetimeExchange(ctx context.Context) (*client.ACLToken, *WriteMeta, error) {
	request := a.ACLApi().PostACLTokenOnetimeExchange(a.client.Ctx)
	result, meta, err := a.client.ExecWrite(ctx, request)
	if err != nil {
		return nil, nil, err
	}

	final := result.(client.ACLToken)
	return &final, meta, nil
}

func (a *ACL) PostACLBootstrap(ctx context.Context) (*client.ACLToken, *WriteMeta, error) {
	request := a.ACLApi().PostACLBootstrap(a.client.Ctx)
	result, meta, err := a.client.ExecWrite(ctx, request)
	if err != nil {
		return nil, nil, err
	}

	final := result.(client.ACLToken)
	return &final, meta, nil
}

func (a *ACL) GetACLTokens(ctx context.Context) (*[]client.ACLTokenListStub, *QueryMeta, error) {
	request := a.ACLApi().GetACLTokens(a.client.Ctx)
	result, meta, err := a.client.ExecQuery(ctx, request)
	if err != nil {
		return nil, nil, err
	}

	final := result.([]client.ACLTokenListStub)
	return &final, meta, nil
}

func (a *ACL) GetACLTokenSelf(ctx context.Context) (*client.ACLToken, *QueryMeta, error) {
	request := a.ACLApi().GetACLTokenSelf(a.client.Ctx)
	result, meta, err := a.client.ExecQuery(ctx, request)
	if err != nil {
		return nil, nil, err
	}

	final := result.(client.ACLToken)
	return &final, meta, nil
}

func (a *ACL) GetACLToken(ctx context.Context, tokenAccessor string) (*client.ACLToken, *QueryMeta, error) {
	if tokenAccessor == "" {
		return nil, nil, errors.New("token accessor id is required")
	}

	request := a.ACLApi().GetACLToken(a.client.Ctx, tokenAccessor)
	result, meta, err := a.client.ExecQuery(ctx, request)
	if err != nil {
		return nil, nil, err
	}

	final := result.(client.ACLToken)
	return &final, meta, nil
}

func (a *ACL) PostACLToken(ctx context.Context, tokenAccessor string) (*client.ACLToken, *WriteMeta, error) {
	if tokenAccessor == "" {
		return nil, nil, errors.New("token accessor id is required")
	}

	request := a.ACLApi().PostACLToken(a.client.Ctx, tokenAccessor)
	result, meta, err := a.client.ExecWrite(ctx, request)
	if err != nil {
		return nil, nil, err
	}

	final := result.(client.ACLToken)
	return &final, meta, nil
}

func (a *ACL) DeleteACLToken(ctx context.Context, tokenAccessor string) (*WriteMeta, error) {
	if tokenAccessor == "" {
		return nil, errors.New("token accessor id is required")
	}

	request := a.ACLApi().DeleteACLToken(a.client.Ctx, tokenAccessor)
	meta, err := a.client.ExecNoResponseWrite(ctx, request)
	if err != nil {
		return nil, err
	}

	return meta, nil
}
