package v1

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"reflect"
	"strconv"
	"time"

	"github.com/hashicorp/cronexpr"
	client "github.com/hashicorp/nomad-openapi/clients/go/v1"
)

type Client struct {
	Ctx       context.Context
	apiClient *client.APIClient
}

func NewClient() (*Client, error) {
	nomadAddr := os.Getenv("NOMAD_ADDR")
	configuration := client.NewConfiguration()

	// Set to the dev agent default if empty
	if nomadAddr == "" {
		nomadAddr = "http://127.0.0.1:4646"
	}

	// Parse the URL so we can extract scheme and port
	nomadURL, err := url.Parse(nomadAddr)
	if err != nil {
		return nil, err
	}

	c := &Client{}

	c.Ctx = context.WithValue(context.Background(), client.ContextServerVariables, map[string]string{
		"scheme":  nomadURL.Scheme,
		"address": nomadURL.Hostname(),
		"port":    nomadURL.Port(),
	})

	err = configureAuth(configuration)
	if err != nil {
		return nil, err
	}

	c.apiClient = client.NewAPIClient(configuration)

	return c, nil
}

func configureAuth(configuration *client.Configuration) error {
	nomadToken := os.Getenv("NOMAD_TOKEN")
	// If environment has a NOMAD_TOKEN, set that in the header
	if nomadToken != "" {
		configuration.DefaultHeader["X-Nomad-Token"] = nomadToken
	}

	err := configureTLS(configuration)
	if err != nil {
		return err
	}
	return nil
}

// Configures TLS if the client environment contains TLS configuration settings.
func configureTLS(configuration *client.Configuration) error {
	tlsConfig, err := tlsConfigFromEnv()
	if err != nil {
		return err
	}

	// if environment is not configured for tls, return nil
	if tlsConfig == nil {
		return nil
	}

	// throw error if the environment is configured for TLS, but the HTTPClient
	// is already set.
	if configuration.HTTPClient != nil {
		return errors.New("client HTTPClient is already configured")
	}

	configuration.HTTPClient = &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: tlsConfig,
		},
	}

	return nil
}

func tlsConfigFromEnv() (*tls.Config, error) {
	serverCertPath := os.Getenv("NOMAD_CACERT")
	clientCertPath := os.Getenv("NOMAD_CLIENT_CERT")
	clientKeyPath := os.Getenv("NOMAD_CLIENT_KEY")

	// If environment is not configured for TLS, return
	if serverCertPath == "" || clientCertPath == "" || clientKeyPath == "" {
		return nil, nil
	}

	cfg := &tls.Config{}

	// Load the certificate authority certificate bytes
	serverCertBytes, err := os.ReadFile(serverCertPath)
	if err != nil {
		return nil, fmt.Errorf("error reading NOMAD_CACERT: %v", err)
	}

	// Load a client certificate from the client settings
	clientCert, err := tls.LoadX509KeyPair(clientCertPath, clientKeyPath)
	if err != nil {
		return nil, err
	}

	// Set the root CA from the server cert
	cfg.RootCAs = x509.NewCertPool()
	cfg.RootCAs.AppendCertsFromPEM(serverCertBytes)
	// Set the client certificate
	cfg.Certificates = []tls.Certificate{clientCert}

	cfg.ServerName = "server.global.nomad"
	region := os.Getenv("NOMAD_REGION")
	if region != "" {
		cfg.ServerName = fmt.Sprintf("server.%s.nomad", region)
	}

	return cfg, nil
}

// ExecQuery executes a request that returns query metadata.
func (c *Client) ExecQuery(ctx context.Context, request interface{}) (interface{}, *QueryMeta, error) {
	typeOf := reflect.TypeOf(request)

	_, ok := typeOf.MethodByName("Execute")
	if !ok {
		return nil, nil, errors.New("execQuery failed: no Execute method on interface")
	}

	request = c.setQueryOptions(ctx, request)

	valueOf := reflect.ValueOf(request)

	values := valueOf.MethodByName("Execute").Call([]reflect.Value{})
	if !values[2].IsNil() {
		return nil, nil, values[2].Interface().(error)
	}

	result := values[0].Interface()
	response := values[1].Interface().(*http.Response)

	meta, err := parseQueryMeta(response)
	if err != nil {
		return nil, nil, err
	}

	return result, meta, nil
}

// ExecWrite executes a request that returns write metadata.
func (c *Client) ExecWrite(ctx context.Context, request interface{}) (interface{}, *WriteMeta, error) {
	typeOf := reflect.TypeOf(request)

	_, ok := typeOf.MethodByName("Execute")
	if !ok {
		return nil, nil, errors.New("ExecWrite failed: no Execute method on interface")
	}

	request = c.setWriteOptions(ctx, request)

	valueOf := reflect.ValueOf(request)

	values := valueOf.MethodByName("Execute").Call([]reflect.Value{})
	if !values[2].IsNil() {
		return nil, nil, values[2].Interface().(error)
	}

	result := values[0].Interface()
	response := values[1].Interface().(*http.Response)

	meta, err := parseWriteMeta(response)
	if err != nil {
		return nil, nil, err
	}

	return result, meta, nil
}

// ExecNoResponseWrite executes a request that returns write metadata, but no model.
func (c *Client) ExecNoResponseWrite(ctx context.Context, request interface{}) (*WriteMeta, error) {
	typeOf := reflect.TypeOf(request)

	_, ok := typeOf.MethodByName("Execute")
	if !ok {
		return nil, errors.New("ExecNoResponseWrite failed: no Execute method on interface")
	}

	request = c.setWriteOptions(ctx, request)

	valueOf := reflect.ValueOf(request)

	values := valueOf.MethodByName("Execute").Call([]reflect.Value{})
	if !values[1].IsNil() {
		return nil, values[1].Interface().(error)
	}

	response := values[0].Interface().(*http.Response)

	meta, err := parseWriteMeta(response)
	if err != nil {
		return nil, err
	}

	return meta, nil
}

// ExecRequest executes a client operation that does not return query or write metadata.
func (c *Client) ExecRequest(_ context.Context, request interface{}) (interface{}, error) {
	typeOf := reflect.TypeOf(request)
	valueOf := reflect.ValueOf(request)

	_, ok := typeOf.MethodByName("Execute")
	if !ok {
		return nil, errors.New("ExecRequest failed: no Execute method on interface")
	}

	values := valueOf.MethodByName("Execute").Call([]reflect.Value{})
	if !values[2].IsNil() {
		return nil, values[2].Interface().(error)
	}

	result := values[0].Interface()

	return result, nil
}

// ExecNoResponseRequest executes a client operation that does not return a model, query or write metadata.
func (c *Client) ExecNoResponseRequest(_ context.Context, request interface{}) error {
	typeOf := reflect.TypeOf(request)
	valueOf := reflect.ValueOf(request)

	_, ok := typeOf.MethodByName("Execute")
	if !ok {
		return errors.New("ExecNoResponseRequest failed: no Execute method on interface")
	}

	values := valueOf.MethodByName("Execute").Call([]reflect.Value{})
	if !values[1].IsNil() {
		return values[1].Interface().(error)
	}

	return nil
}

// setQueryOptions is used to annotate the request with
// additional query options
func (c *Client) setQueryOptions(ctx context.Context, iface interface{}) interface{} {
	opts, ok := ctx.Value("QueryOpts").(*QueryOpts)
	if !ok || opts == nil {
		return iface
	}

	typeOf := reflect.TypeOf(iface)
	valueOf := reflect.ValueOf(iface)

	_, ok = typeOf.MethodByName(RegionKey)
	if ok && opts.Region != "" {
		iface = valueOf.MethodByName(RegionKey).Call([]reflect.Value{reflect.ValueOf(opts.Region)})[0].Interface()
		valueOf = reflect.ValueOf(iface)
	}

	_, ok = typeOf.MethodByName(NamespaceKey)
	if ok && opts.Namespace != "" {
		iface = valueOf.MethodByName(NamespaceKey).Call([]reflect.Value{reflect.ValueOf(opts.Namespace)})[0].Interface()
		valueOf = reflect.ValueOf(iface)
	}

	_, ok = typeOf.MethodByName("XNomadToken")
	if ok && opts.AuthToken != "" {
		iface = valueOf.MethodByName("XNomadToken").Call([]reflect.Value{reflect.ValueOf(opts.AuthToken)})[0].Interface()
		valueOf = reflect.ValueOf(iface)
	}

	_, ok = typeOf.MethodByName("Stale")
	if ok && opts.AllowStale {
		iface = valueOf.MethodByName("Stale").Call([]reflect.Value{reflect.ValueOf("")})[0].Interface()
		valueOf = reflect.ValueOf(iface)
	}

	_, ok = typeOf.MethodByName("Index")
	if ok && opts.WaitIndex != 0 {
		iface = valueOf.MethodByName("Index").Call([]reflect.Value{reflect.ValueOf(opts.WaitIndex)})[0].Interface()
		valueOf = reflect.ValueOf(iface)
	}

	_, ok = typeOf.MethodByName("Wait")
	if ok && opts.WaitTime != 0 {
		iface = valueOf.MethodByName("Wait").Call([]reflect.Value{reflect.ValueOf(fmt.Sprintf("%dms", opts.WaitTime))})[0].Interface()
		valueOf = reflect.ValueOf(iface)
	}

	_, ok = typeOf.MethodByName(PrefixKey)
	if ok && opts.Prefix != "" {
		iface = valueOf.MethodByName(PrefixKey).Call([]reflect.Value{reflect.ValueOf(opts.Prefix)})[0].Interface()
		valueOf = reflect.ValueOf(iface)
	}

	_, ok = typeOf.MethodByName(PerPageKey)
	if ok && opts.PerPage != 0 {
		iface = valueOf.MethodByName(PerPageKey).Call([]reflect.Value{reflect.ValueOf(opts.PerPage)})[0].Interface()
		valueOf = reflect.ValueOf(iface)
	}

	_, ok = typeOf.MethodByName(NextTokenKey)
	if ok && opts.NextToken != "" {
		iface = valueOf.MethodByName(NextTokenKey).Call([]reflect.Value{reflect.ValueOf(opts.NextToken)})[0].Interface()
		valueOf = reflect.ValueOf(iface)
	}

	// TODO: Handle extra params
	//if c.config.Params != nil {
	//}

	return iface
}

// setWriteOptions is used to annotate an openapi request with additional write options.
func (c *Client) setWriteOptions(ctx context.Context, iface interface{}) interface{} {
	opts, ok := ctx.Value("WriteOpts").(*WriteOpts)
	if !ok || opts == nil {
		return iface
	}
	typeOf := reflect.TypeOf(iface)
	valueOf := reflect.ValueOf(iface)

	_, ok = typeOf.MethodByName(RegionKey)
	if ok && opts.Region != "" {
		iface = valueOf.MethodByName(RegionKey).Call([]reflect.Value{reflect.ValueOf(opts.Region)})[0].Interface()
		valueOf = reflect.ValueOf(iface)
	}

	_, ok = typeOf.MethodByName(NamespaceKey)
	if ok && opts.Namespace != "" {
		iface = valueOf.MethodByName(NamespaceKey).Call([]reflect.Value{reflect.ValueOf(opts.Namespace)})[0].Interface()
		valueOf = reflect.ValueOf(iface)
	}

	_, ok = typeOf.MethodByName("XNomadToken")
	if ok && opts.AuthToken != "" {
		iface = valueOf.MethodByName("XNomadToken").Call([]reflect.Value{reflect.ValueOf(opts.AuthToken)})[0].Interface()
		valueOf = reflect.ValueOf(iface)
	}

	_, ok = typeOf.MethodByName(IdempotencyTokenKey)
	if ok && opts.IdempotencyToken != "" {
		iface = valueOf.MethodByName(IdempotencyTokenKey).Call([]reflect.Value{reflect.ValueOf(opts.IdempotencyToken)})[0].Interface()
		valueOf = reflect.ValueOf(iface)
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
	WaitIndex int32

	// WaitTime is used to bound the duration of a call and is set in milliseconds.
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
}

// DefaultQueryOpts returns a QueryOpts for the default namespace and global region.
func DefaultQueryOpts() *QueryOpts {
	opts := &QueryOpts{}
	return opts.
		WithNamespace(DefaultNamespace).
		WithRegion(GlobalRegion)
}

func (q *QueryOpts) WithRegion(region string) *QueryOpts {
	q.Region = region
	return q
}

func (q *QueryOpts) WithNamespace(namespace string) *QueryOpts {
	q.Namespace = namespace
	return q
}

func (q *QueryOpts) WithAuthToken(authToken string) *QueryOpts {
	q.AuthToken = authToken
	return q
}

func (q *QueryOpts) WithPrefix(prefix string) *QueryOpts {
	q.Prefix = prefix
	return q
}

func (q *QueryOpts) WithPerPage(perPage int32) *QueryOpts {
	q.PerPage = perPage
	return q
}

func (q *QueryOpts) WithWaitIndex(waitIndex int32) *QueryOpts {
	q.WaitIndex = waitIndex
	return q
}

func (q *QueryOpts) WithWaitTime(waitTime time.Duration) *QueryOpts {
	q.WaitTime = waitTime
	return q
}

func (q *QueryOpts) WithNextToken(nextToken string) *QueryOpts {
	q.NextToken = nextToken
	return q
}

func (q *QueryOpts) WithAllowStale(allowStale bool) *QueryOpts {
	q.AllowStale = allowStale
	return q
}

func (q *QueryOpts) Ctx() context.Context {
	return context.WithValue(context.Background(), "QueryOpts", q)
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

	// IdempotencyToken can be used to ensure the operation is idempotent.
	IdempotencyToken string
}

// DefaultWriteOpts returns a WriteOps for the default namespace and the global region
func DefaultWriteOpts() *WriteOpts {
	opts := &WriteOpts{}
	return opts.
		WithNamespace(DefaultNamespace).
		WithRegion(GlobalRegion)
}

func (w *WriteOpts) WithRegion(region string) *WriteOpts {
	w.Region = region
	return w
}

func (w *WriteOpts) WithNamespace(namespace string) *WriteOpts {
	w.Namespace = namespace
	return w
}

func (w *WriteOpts) WithAuthToken(authToken string) *WriteOpts {
	w.AuthToken = authToken
	return w
}

func (w *WriteOpts) WithIdempotencyToken(idempotencyToken string) *WriteOpts {
	w.IdempotencyToken = idempotencyToken
	return w
}

func (w *WriteOpts) Ctx() context.Context {
	return context.WithValue(context.Background(), "WriteOpts", w)
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

	// RegionKey can be used to prevent hard coded string key accessor errors
	RegionKey = "Region"

	// NamespaceKey can be used to prevent hard coded string key accessor errors
	NamespaceKey = "Namespace"

	// ParamsKey can be used to prevent hard coded string key accessor errors
	ParamsKey = "Params"

	// PerPageKey can be used to prevent hard coded string key accessor errors
	PerPageKey = "PerPage"

	// NextTokenKey can be used to prevent hard coded string key accessor errors
	NextTokenKey = "NextToken"

	// PrefixKey can be used to prevent hard coded string key accessor errors
	PrefixKey = "Prefix"

	// AuthTokenKey can be used to prevent hard coded string key accessor errors
	AuthTokenKey = "AuthToken"

	// AllowStaleKey can be used to prevent hard coded string key accessor errors
	AllowStaleKey = "AllowStale"

	// WaitTimeKey can be used to prevent hard coded string key accessor errors
	WaitTimeKey = "WaitTime"

	// WriteIndexKey can be used to prevent hard coded string key accessor errors
	WaitIndexKey = "WaitIndex"

	// IdempotencyTokenKey can be used to prevent hard coded string key accessor errors
	IdempotencyTokenKey = "IdempotencyToken"
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
