package api

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"

	oac "github.com/hashicorp/nomad-openapi/v1/client"
)

type Client struct {
	Ctx        context.Context
	config     *Config
	oapiClient *oac.APIClient
}

func newClient() (*Client, error) {
	nomadAddr := os.Getenv("NOMAD_ADDR")
	nomadURL, err := url.Parse(nomadAddr)
	if err != nil {
		return nil, err
	}

	client := &Client{}

	client.Ctx = context.WithValue(context.Background(), oac.ContextServerVariables, map[string]string{
		"scheme":  nomadURL.Scheme,
		"address": nomadURL.Hostname(),
		"port":    nomadURL.Port(),
	})

	configuration := oac.NewConfiguration()
	client.oapiClient = oac.NewAPIClient(configuration)

	return client, nil
}

func NewQueryClient(opts *QueryOptions) (*Client, error) {
	client, err := newClient()
	if err != nil {
		return nil, err
	}

	client.config = &Config{
		QueryOptions: opts,
	}
	client.SetRegion(opts.Region)
	client.SetNamespace(opts.Namespace)

	return client, nil
}

func NewWriteClient(opts *WriteOptions) (*Client, error) {
	client, err := newClient()
	if err != nil {
		return nil, err
	}

	client.config = &Config{
		WriteOptions: opts,
	}
	client.SetRegion(opts.Region)
	client.SetNamespace(opts.Namespace)

	return client, nil
}
func (c *Client) SetRegion(region string) {
	if region == "" {
		c.config.Region = GlobalRegion
		return
	}
	c.config.Region = region
}

func (c *Client) SetNamespace(namespace string) {
	if namespace == "" {
		c.config.Namespace = DefaultNamespace
		return
	}
	c.config.Namespace = namespace
}

// setQueryOptions is used to annotate the request with
// additional query options
func (c *Client) setQueryOptions(regionFn func(string), namespaceFn func(string), authTokenFn func(string),
	staleFn func(string), waitIndexFn func(uint64), waitTimeFn func(int32), prefixFn func(string)) {
	cfg := c.config
	opts := c.config.QueryOptions

	if cfg.Region != "" && regionFn != nil {
		regionFn(cfg.Region)
	}
	if cfg.Namespace != "" && namespaceFn != nil {
		namespaceFn(cfg.Namespace)
	}
	if opts.AuthToken != "" && authTokenFn != nil {
		authTokenFn(opts.AuthToken)
	}
	if opts.AllowStale {
		staleFn("")
	}
	if opts.WaitIndex != 0 {
		waitIndexFn(opts.WaitIndex)
	}
	if opts.WaitTime != 0 {
		w, err := strconv.ParseInt(durToMsec(opts.WaitTime), 10, 32)
		if err == nil {
			waitTimeFn(int32(w))
		}
	}
	if opts.Prefix != "" {
		prefixFn(opts.Prefix)
	}
	// TODO: Handle extra params
	//if c.config.Params != nil {
	//}
}

// setWriteOptions is used to annotate an openapi request with
// additional write options
func (c *Client) setWriteOptions(regionFn func(string), namespaceFn func(string), authTokenFn func(string), idempotencyTokenFn func(string)) {
	cfg := c.config

	if cfg.Region != "" && regionFn != nil {
		regionFn(cfg.Region)
	}
	if cfg.Namespace != "" && namespaceFn != nil {
		namespaceFn(cfg.Namespace)
	}
	if cfg.WriteOptions.AuthToken != "" && authTokenFn != nil {
		authTokenFn(cfg.WriteOptions.AuthToken)
	}
	if cfg.WriteOptions.IdempotencyToken != "" && idempotencyTokenFn != nil {
		idempotencyTokenFn(cfg.WriteOptions.IdempotencyToken)
	}
}

// QueryOptions are used to parametrize a query
type QueryOptions struct {
	// Providing a datacenter overwrites the region provided
	// by the Config
	Region string

	// Namespace is the target namespace for the query.
	Namespace string

	// AllowStale allows any Nomad server (non-leader) to service
	// a read. This allows for lower latency and higher throughput
	AllowStale bool

	// WaitIndex is used to enable a blocking query. Waits
	// until the timeout or the next index is reached
	WaitIndex uint64

	// WaitTime is used to bound the duration of a wait.
	// Defaults to that of the Config, but can be overridden.
	WaitTime time.Duration

	// If set, used as prefix for resource list searches
	Prefix string

	// Set HTTP parameters on the query.
	Params map[string]string

	// AuthToken is the secret ID of an ACL token
	AuthToken string

	// PerPage is the number of entries to be returned in queries that support
	// paginated lists.
	PerPage int32

	// NextToken is the token used indicate where to start paging for queries
	// that support paginated lists.
	NextToken string

	// ctx is an optional context pass through to the underlying HTTP
	// request layer. Use Context() and WithContext() to manage this.
	ctx context.Context
}

// WriteOptions are used to parametrize a write operation
type WriteOptions struct {
	// Providing a datacenter overwrites the region provided
	// by the Config
	Region string

	// Namespace is the target namespace for the operation.
	Namespace string

	// AuthToken is the secret ID of an ACL token
	AuthToken string

	// ctx is an optional context pass through to the underlying HTTP
	// request layer. Use Context() and WithContext() to manage this.
	ctx context.Context

	// IdempotencyToken can be used to ensure the operation is idempotent.
	IdempotencyToken string
}

// QueryMeta is used to return metadata about a query
type QueryMeta struct {
	// LastIndex can be used as a WaitIndex to perform
	// a blocking query
	LastIndex uint64

	// LastContact is the time of last contact from the leader for the
	// server servicing the request
	LastContact time.Duration

	// KnownLeader indicates if there is a known leader
	KnownLeader bool

	// RequestTime is how long did the request took
	RequestTime time.Duration
}

func parseQueryMeta(resp *http.Response, q *QueryMeta) error {
	header := resp.Header

	// Parse the X-Nomad-Index
	index, err := strconv.ParseUint(header.Get("X-Nomad-Index"), 10, 64)
	if err != nil {
		return fmt.Errorf("failed to parse X-Nomad-Index: %v", err)
	}
	q.LastIndex = index

	// Parse the X-Nomad-LastContact
	last, err := strconv.ParseUint(header.Get("X-Nomad-LastContact"), 10, 64)
	if err != nil {
		return fmt.Errorf("failed to parse X-Nomad-LastContact: %v", err)
	}
	q.LastContact = time.Duration(last) * time.Millisecond

	// Parse the X-Nomad-KnownLeader
	switch header.Get("X-Nomad-KnownLeader") {
	case "true":
		q.KnownLeader = true
	default:
		q.KnownLeader = false
	}
	return nil
}

// WriteMeta is used to return metadata about a write operation
type WriteMeta struct {
	// LastIndex can be used as a WaitIndex to perform
	// a blocking query
	LastIndex uint64

	// RequestTime is how long did the request took
	RequestTime time.Duration
}

// parseWriteMeta is used to help parse write meta-data
func parseWriteMeta(resp *http.Response, q *WriteMeta) error {
	header := resp.Header

	// Parse the X-Nomad-Index
	index, err := strconv.ParseUint(header.Get("X-Nomad-Index"), 10, 64)
	if err != nil {
		return fmt.Errorf("failed to parse X-Nomad-Index: %v", err)
	}
	q.LastIndex = index
	return nil
}

// durToMsec converts a duration to a millisecond specified string
func durToMsec(dur time.Duration) string {
	return fmt.Sprintf("%dms", dur/time.Millisecond)
}

func headerByKey(resp *http.Response, key string) (string, error) {
	h := resp.Header.Get(key)
	if len(h) == 0 {
		return "", errors.New("key not found")
	}
	return h, nil
}

type Config struct {
	// Region to use. If not provided, the default agent region is used.
	Region string

	// Namespace to use. If not provided the default namespace is used.
	Namespace string

	QueryOptions *QueryOptions

	WriteOptions *WriteOptions
}

const (
	// JobTypeService indicates a long-running processes
	JobTypeService = "service"

	// JobTypeBatch indicates a short-lived process
	JobTypeBatch = "batch"

	// JobTypeSystem indicates a system process that should run on all clients
	JobTypeSystem = "system"

	// PeriodicSpecCron is used for a cron spec.
	PeriodicSpecCron = "cron"

	// DefaultNamespace is the default namespace.
	DefaultNamespace = "default"

	// GlobalRegion is a sentinel region value dor Job configuration
	// that users may specify to indicate the job should be run on
	// the region of the node that the job was submitted to.
	// For Client configuration, if no region information is given,
	// the client node will default to be part of the GlobalRegion.
	GlobalRegion = "global"

	// RegisterEnforceIndexErrPrefix is the prefix to use in errors caused by
	// enforcing the job modify index during registers.
	RegisterEnforceIndexErrPrefix = "Enforcing job modify index"
)
