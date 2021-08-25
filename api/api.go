package api

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"reflect"
	"strconv"
	"time"

	"github.com/hashicorp/cronexpr"
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

func NewQueryClient(opts *QueryOpts) (*Client, error) {
	client, err := newClient()
	if err != nil {
		return nil, err
	}

	client.config = &Config{
		QueryOpts: opts,
	}
	client.SetRegion(opts.Region)
	client.SetNamespace(opts.Namespace)

	return client, nil
}

func NewWriteClient(opts *WriteOpts) (*Client, error) {
	client, err := newClient()
	if err != nil {
		return nil, err
	}

	client.config = &Config{
		WriteOpts: opts,
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
func (c *Client) setQueryOptions(iface interface{}) interface{} {
	cfg := c.config
	opts := c.config.QueryOpts
	typeOf := reflect.TypeOf(iface)
	valueOf := reflect.ValueOf(iface)

	_, ok := typeOf.MethodByName("Region")
	if ok && cfg.Region != "" {
		iface = valueOf.MethodByName("Region").Call([]reflect.Value{reflect.ValueOf(cfg.Region)})[0].Interface()
	}
	_, ok = typeOf.MethodByName("Namespace")
	if ok && cfg.Namespace != "" {
		iface = valueOf.MethodByName("Namespace").Call([]reflect.Value{reflect.ValueOf(cfg.Namespace)})[0].Interface()
	}
	_, ok = typeOf.MethodByName("XNomadToken")
	if ok && opts.AuthToken != "" {
		iface = valueOf.MethodByName("XNomadToken").Call([]reflect.Value{reflect.ValueOf(opts.AuthToken)})[0].Interface()
	}
	_, ok = typeOf.MethodByName("Stale")
	if ok && opts.AllowStale {
		iface = valueOf.MethodByName("Stale").Call([]reflect.Value{reflect.ValueOf("")})[0].Interface()
	}
	_, ok = typeOf.MethodByName("Index")
	if ok && opts.WaitIndex != 0 {
		idx := int32(opts.WaitIndex)
		iface = valueOf.MethodByName("Index").Call([]reflect.Value{reflect.ValueOf(idx)})[0].Interface()
	}
	_, ok = typeOf.MethodByName("Wait")
	if ok && opts.WaitTime != 0 {
		w, err := strconv.ParseInt(durToMsec(opts.WaitTime), 10, 32)
		if err == nil {
			iface = valueOf.MethodByName("Wait").Call([]reflect.Value{reflect.ValueOf(int32(w))})[0].Interface()
		}
	}
	_, ok = typeOf.MethodByName("Prefix")
	if ok && opts.Prefix != "" {
		iface = valueOf.MethodByName("Prefix").Call([]reflect.Value{reflect.ValueOf(opts.Prefix)})[0].Interface()
	}
	_, ok = typeOf.MethodByName("PerPage")
	if ok && opts.PerPage != 0 {
		iface = valueOf.MethodByName("PerPage").Call([]reflect.Value{reflect.ValueOf(opts.PerPage)})[0].Interface()
	}
	_, ok = typeOf.MethodByName("NextToken")
	if ok && opts.NextToken != "" {
		iface = valueOf.MethodByName("NextToken").Call([]reflect.Value{reflect.ValueOf(opts.NextToken)})[0].Interface()
	}
	// TODO: Handle extra params
	//if c.config.Params != nil {
	//}

	return iface
}

// setWriteOptions is used to annotate an openapi request with additional write options.
func (c *Client) setWriteOptions(iface interface{}) interface{} {
	cfg := c.config
	opts := c.config.WriteOpts
	typeOf := reflect.TypeOf(iface)
	valueOf := reflect.ValueOf(iface)

	_, ok := typeOf.MethodByName("Region")
	if ok && cfg.Region != "" {
		iface = valueOf.MethodByName("Region").Call([]reflect.Value{reflect.ValueOf(cfg.Region)})[0].Interface()
	}

	_, ok = typeOf.MethodByName("Namespace")
	if ok && cfg.Namespace != "" {
		iface = valueOf.MethodByName("Namespace").Call([]reflect.Value{reflect.ValueOf(cfg.Namespace)})[0].Interface()
	}

	_, ok = typeOf.MethodByName("XNomadToken")
	if ok && opts.AuthToken != "" {
		iface = valueOf.MethodByName("XNomadToken").Call([]reflect.Value{reflect.ValueOf(opts.AuthToken)})[0].Interface()
	}

	_, ok = typeOf.MethodByName("IdempotencyToken")
	if ok && opts.IdempotencyToken != "" {
		iface = valueOf.MethodByName("IdempotencyToken").Call([]reflect.Value{reflect.ValueOf(opts.IdempotencyToken)})[0].Interface()
	}

	return iface
}

// QueryOpts are used to parametrize a query
type QueryOpts struct {
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

// WriteOpts are used to parametrize a write operation
type WriteOpts struct {
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
	// LastIndex can be used as a Index to perform
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

func parseQueryMeta(resp *http.Response) (*QueryMeta, error) {
	q := &QueryMeta{}
	header := resp.Header

	// Parse the X-Nomad-Index
	indexHeader := header.Get("X-Nomad-Index")
	if indexHeader == "" {
		return nil, fmt.Errorf("X-Nomad-Index header not found")
	}
	index, err := strconv.ParseUint(indexHeader, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("failed to parse X-Nomad-Index: %v", err)
	}
	q.LastIndex = index

	// Parse the X-Nomad-LastContact
	lastHeader := header.Get("X-Nomad-Lastcontact")
	if lastHeader == "" {
		return nil, fmt.Errorf("X-Nomad-Lastcontact header not found")
	}
	last, err := strconv.ParseUint(lastHeader, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("failed to parse X-Nomad-Lastcontact: %v", err)
	}
	q.LastContact = time.Duration(last) * time.Millisecond

	leaderHeader := header.Get("X-Nomad-Knownleader")
	if leaderHeader == "" {
		return nil, fmt.Errorf("X-Nomad-Knownleader header not found")
	}
	// Parse the X-Nomad-KnownLeader
	switch leaderHeader {
	case "true":
		q.KnownLeader = true
	default:
		q.KnownLeader = false
	}
	return q, nil
}

// WriteMeta is used to return metadata about a write operation
type WriteMeta struct {
	// LastIndex can be used as a Index to perform
	// a blocking query
	LastIndex uint64

	// RequestTime is how long did the request took
	RequestTime time.Duration
}

// parseWriteMeta is used to help parse write meta-data
func parseWriteMeta(resp *http.Response) (*WriteMeta, error) {
	header := resp.Header.Get("X-Nomad-Index")
	if header == "" {
		return nil, errors.New("X-Nomad-Index header not found")
	}

	// Parse the X-Nomad-Index
	index, err := strconv.ParseUint(header, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("failed to parse X-Nomad-Index: %v", err)
	}

	return &WriteMeta{LastIndex: index}, nil
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

	QueryOpts *QueryOpts

	WriteOpts *WriteOpts
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

// cronParseNext is a helper that parses the next time for the given expression
// but captures any panic that may occur in the underlying library.
// ---  THIS FUNCTION IS REPLICATED IN nomad/structs/structs.go
// and should be kept in sync.
func cronParseNext(e *cronexpr.Expression, fromTime time.Time, spec string) (t time.Time, err error) {
	defer func() {
		if recover() != nil {
			t = time.Time{}
			err = fmt.Errorf("failed parsing cron expression: %q", spec)
		}
	}()

	return e.Next(fromTime), nil
}
