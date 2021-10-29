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
func (a *ACL) Policies(ctx context.Context) (*[]client.ACLPolicyListStub, *QueryMeta, error) {
	request := a.ACLApi().GetACLPolicies(a.client.Ctx)
	result, meta, err := a.client.ExecQuery(ctx, request)
	if err != nil {
		return nil, nil, err
	}

	final := result.([]client.ACLPolicyListStub)
	return &final, meta, nil
}

// returns objectSchema
// pass in policy name, no payload
func (a *ACL) GetPolicy(ctx context.Context, policyName string) (*client.ACLPolicy, *QueryMeta, error) {
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
// pass in policy name WITH payload
// can't use "Save", because there's an ACL "Save" for tokens as well
func (a *ACL) SavePolicy(ctx context.Context, policy *client.ACLPolicy) (*WriteMeta, error) {
	if *policy.Name == "" {
		return nil, errors.New("policy name is required")
	}

	// appending payload?  something that's taken care of by appendParams so I don't have to worry about it?
	request := a.ACLApi().PostACLPolicy(a.client.Ctx, *policy.Name)
	meta, err := a.client.ExecNoResponseWrite(ctx, request)
	if err != nil {
		return nil, err
	}

	return meta, nil
}

// pass in name, no payload
// can't use "Delete", because there's an ACL "Delete" for tokens as well
func (a *ACL) DeletePolicy(ctx context.Context, policyName string) (*WriteMeta, error) {
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

func (a *ACL) OnetimeToken(ctx context.Context) (*client.OneTimeToken, *WriteMeta, error) {
	request := a.ACLApi().PostACLTokenOnetime(a.client.Ctx)
	result, meta, err := a.client.ExecWrite(ctx, request)
	if err != nil {
		return nil, nil, err
	}

	final := result.(client.OneTimeToken)
	return &final, meta, nil
}

func (a *ACL) Exchange(ctx context.Context) (*client.ACLToken, *WriteMeta, error) {
	request := a.ACLApi().PostACLTokenOnetimeExchange(a.client.Ctx)
	result, meta, err := a.client.ExecWrite(ctx, request)
	if err != nil {
		return nil, nil, err
	}

	final := result.(client.ACLToken)
	return &final, meta, nil
}

func (a *ACL) Bootstrap(ctx context.Context) (*client.ACLToken, *WriteMeta, error) {
	request := a.ACLApi().PostACLBootstrap(a.client.Ctx)
	result, meta, err := a.client.ExecWrite(ctx, request)
	if err != nil {
		return nil, nil, err
	}

	final := result.(client.ACLToken)
	return &final, meta, nil
}

func (a *ACL) Tokens(ctx context.Context) (*[]client.ACLTokenListStub, *QueryMeta, error) {
	request := a.ACLApi().GetACLTokens(a.client.Ctx)
	result, meta, err := a.client.ExecQuery(ctx, request)
	if err != nil {
		return nil, nil, err
	}

	final := result.([]client.ACLTokenListStub)
	return &final, meta, nil
}

func (a *ACL) Self(ctx context.Context) (*client.ACLToken, *QueryMeta, error) {
	request := a.ACLApi().GetACLTokenSelf(a.client.Ctx)
	result, meta, err := a.client.ExecQuery(ctx, request)
	if err != nil {
		return nil, nil, err
	}

	final := result.(client.ACLToken)
	return &final, meta, nil
}

func (a *ACL) GetToken(ctx context.Context, tokenAccessor string) (*client.ACLToken, *QueryMeta, error) {
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

func (a *ACL) SaveToken(ctx context.Context, token *client.ACLToken) (*client.ACLToken, *WriteMeta, error) {
	if *token.AccessorID == "" {
		return nil, nil, errors.New("token accessor id is required")
	}

	request := a.ACLApi().PostACLToken(a.client.Ctx, *token.AccessorID)
	result, meta, err := a.client.ExecWrite(ctx, request)
	if err != nil {
		return nil, nil, err
	}

	final := result.(client.ACLToken)
	return &final, meta, nil
}

func (a *ACL) DeleteToken(ctx context.Context, tokenAccessor string) (*WriteMeta, error) {
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
