// +build templates

// This file is a part of msgraph.go/gen/templates.
// Anything until the first appearance of "// BEGIN" line will be ignored.

package msgraph

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"sync"
	"time"

	"github.com/davinash/msgraph.go/jsonx"
	"github.com/rickb777/date/period"
)

// BEGIN - everything below this line will be copied to the output

// Binary is type alias for Edm.Binary
type Binary []byte

// Stream is type alias for Edm.Stream
type Stream []byte

// UUID is type alias for Edm.Guid
type UUID string

// Date is type alias for Edm.Date
type Date string

// NewDate creates Date from time.Time
func NewDate(t time.Time) *Date {
	s := t.Format("2006-01-02")
	return (*Date)(&s)
}

// Time converts Date to time.Time
func (d *Date) Time() (time.Time, error) {
	return time.Parse("2006-01-02", string(*d))
}

// TimeOfDay is type alis for Edm.TimeOfDay
type TimeOfDay string

// NewTimeOfDay creates NewTimeOfDay from time.Time
func NewTimeOfDay(t time.Time) *TimeOfDay {
	s := t.Format("15:04:05")
	return (*TimeOfDay)(&s)
}

// Time converts TimeOfDay to time.Time
func (t *TimeOfDay) Time() (time.Time, error) {
	return time.Parse("15:04:05", string(*t))
}

// Duration is type alias for Edm.Duration
type Duration string

// NewDuration creates Duration from time.Duration
func NewDuration(d time.Duration) *Duration {
	p, _ := period.NewOf(d)
	s := p.String()
	return (*Duration)(&s)
}

// Time converts Duration to time.Duration
func (d *Duration) Time() (time.Duration, error) {
	p, err := period.Parse(string(*d))
	if err != nil {
		return 0, err
	}
	return p.DurationApprox(), nil
}

// Object is the common ancestor of all models
type Object struct {
	// AdditionalData contains all other fields not defined above
	AdditionalData map[string]interface{} `json:"-" jsonx:"true"`
}

// SetAdditionalData sets object's additional data
func (o *Object) SetAdditionalData(key string, val interface{}) {
	if o.AdditionalData == nil {
		o.AdditionalData = map[string]interface{}{key: val}
	} else {
		o.AdditionalData[key] = val
	}
}

// GetAdditionalData gets object's additional data
func (o *Object) GetAdditionalData(key string) (interface{}, bool) {
	if o.AdditionalData == nil {
		return nil, false
	} else {
		val, ok := o.AdditionalData[key]
		return val, ok
	}
}

// ErrorObject is common error object
type ErrorObject struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Object
}

// ErrorResponse is common error response
type ErrorResponse struct {
	ErrorObject ErrorObject    `json:"error"`
	Response    *http.Response `json:"-"`
	Object
}

// Error implements error interface
func (r *ErrorResponse) Error() string {
	b, _ := jsonx.Marshal(r)
	return fmt.Sprintf("%s: %s", r.Status(), string(b))
}

// Status returns status, "000 Unknown" if response is nil
func (r *ErrorResponse) Status() string {
	if r.Response == nil {
		return "000 Unknown"
	}
	return r.Response.Status
}

// StatusCode returns status code, 0 if response is nil
func (r *ErrorResponse) StatusCode() int {
	if r.Response == nil {
		return 0
	}
	return r.Response.StatusCode
}

// Paging is struct returned to paging requests
type Paging struct {
	NextLink string          `json:"@odata.nextLink"`
	Value    json.RawMessage `json:"value"`
}

// BaseRequestBuilder is base request builder
type BaseRequestBuilder struct {
	baseURL           string
	client            *http.Client
	requestObject     interface{}
	tenantID          string
	applicationID     string
	clientSecurityKey string
	apiCall           sync.Mutex
	token             Token
}

// URL returns URL
func (r *BaseRequestBuilder) URL() string {
	return r.baseURL
}

// SetURL sets the baseURL
func (r *BaseRequestBuilder) SetURL(baseURL string) {
	r.baseURL = baseURL
}

// BaseRequest is base request
type BaseRequest struct {
	baseURL           string
	client            *http.Client
	requestObject     interface{}
	header            http.Header
	query             url.Values
	tenantID          string
	applicationID     string
	clientSecurityKey string
	token             Token
	apiCall           sync.Mutex
}

// URL returns URL with queries
func (r *BaseRequest) URL() string {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.baseURL + query
}

// Client returns HTTP client
func (r *BaseRequest) Client() *http.Client {
	return r.client
}

// Header returns headers of the request
func (r *BaseRequest) Header() http.Header {
	if r.header == nil {
		r.header = http.Header{}
	}
	return r.header
}

// Query returns queries of the request
func (r *BaseRequest) Query() url.Values {
	if r.query == nil {
		r.query = url.Values{}
	}
	return r.query
}

// Expand adds $expand query
func (r *BaseRequest) Expand(value string) {
	if r.query == nil {
		r.query = url.Values{}
	}
	r.query.Add("$expand", value)
}

// Select adds $select query
func (r *BaseRequest) Select(value string) {
	if r.query == nil {
		r.query = url.Values{}
	}
	r.query.Add("$select", value)
}

// Top adds $top query
func (r *BaseRequest) Top(value int) {
	if r.query == nil {
		r.query = url.Values{}
	}
	r.query.Add("$top", strconv.Itoa(value))
}

// Filter adds $filter query
func (r *BaseRequest) Filter(value string) {
	if r.query == nil {
		r.query = url.Values{}
	}
	r.query.Add("$filter", value)
}

// Skip adds $skip query
func (r *BaseRequest) Skip(value int) {
	if r.query == nil {
		r.query = url.Values{}
	}
	r.query.Add("$skip", strconv.Itoa(value))
}

// OrderBy adds $orderby query
func (r *BaseRequest) OrderBy(value string) {
	if r.query == nil {
		r.query = url.Values{}
	}
	r.query.Add("$orderby", value)
}

// NewRequest returns new HTTP request
func (r *BaseRequest) NewRequest(method, path string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, r.baseURL+path, body)
	if err != nil {
		return nil, err
	}
	if r.header != nil {
		for key, values := range r.header {
			for _, value := range values {
				req.Header.Add(key, value)
			}
		}
	}
	err = r.getAuthToken()
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", r.token.GetAccessToken())
	return req, nil
}

func (r *BaseRequest) getAuthToken() error {
	r.apiCall.Lock()
	defer r.apiCall.Unlock() // unlock when the func returns
	// Check token
	if r.token.WantsToBeRefreshed() { // Token not valid anymore?
		err := r.refreshToken()
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *BaseRequest) refreshToken() error {
	if r.tenantID == "" {
		return fmt.Errorf("tenant ID is empty")
	}
	resource := fmt.Sprintf("/%v/oauth2/token", r.tenantID)
	data := url.Values{}
	data.Add("grant_type", "client_credentials")
	data.Add("client_id", r.applicationID)
	data.Add("client_secret", r.clientSecurityKey)
	data.Add("resource", "https://graph.microsoft.com")

	u, err := url.ParseRequestURI("https://login.microsoftonline.com")
	if err != nil {
		return fmt.Errorf("unable to parse URI: %v", err)
	}

	u.Path = resource
	req, err := http.NewRequest("POST", u.String(), bytes.NewBufferString(data.Encode()))

	if err != nil {
		return fmt.Errorf("HTTP Request Error: %v", err)
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	var newToken Token
	_, err = r.performRequest(req, &newToken) // perform the prepared request
	if err != nil {
		return fmt.Errorf("error on getting msgraph Token: %v", err)
	}
	r.token = newToken
	return err
}

func (r *BaseRequest) performRequest(req *http.Request, v interface{}) (int, error) {
	httpClient := &http.Client{
		Transport: &http.Transport{
			TLSHandshakeTimeout:   100 * time.Second,
			ResponseHeaderTimeout: 100 * time.Second,
			ExpectContinueTimeout: 10 * time.Second,
			IdleConnTimeout:       1000 * time.Second,
		},
		Timeout: time.Second * 500,
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		if resp == nil {
			return 400, fmt.Errorf("HTTP response error: %v ", err)
		}
		return resp.StatusCode, fmt.Errorf("HTTP response error: %v ", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return resp.StatusCode, fmt.Errorf("StatusCode is not OK: %v. Body: %v ", resp.StatusCode, string(body))
	} else if resp.StatusCode == 204 {
		return resp.StatusCode, err
	}
	if err != nil {
		return resp.StatusCode, fmt.Errorf("HTTP response read error: %v of http.Request: %v", err, req.URL)
	}
	return resp.StatusCode, json.Unmarshal(body, &v) // return the error of the json unmarshal
}

// NewJSONRequest returns new HTTP request with JSON payload
func (r *BaseRequest) NewJSONRequest(method, path string, obj interface{}) (*http.Request, error) {
	buf := &bytes.Buffer{}
	if obj != nil {
		err := jsonx.NewEncoder(buf).Encode(obj)
		if err != nil {
			return nil, err
		}
	}
	req, err := r.NewRequest(method, path, buf)
	if err != nil {
		return nil, err
	}
	if obj != nil {
		req.Header.Add("Content-Type", "application/json")
	}
	return req, nil
}

// DecodeJSONResponse decodes HTTP response with JSON payload
func (r *BaseRequest) DecodeJSONResponse(res *http.Response, obj interface{}) error {
	switch res.StatusCode {
	case http.StatusOK, http.StatusCreated:
		if obj != nil {
			err := jsonx.NewDecoder(res.Body).Decode(obj)
			if err != nil {
				return err
			}
		}
		return nil
	case http.StatusNoContent:
		return nil
	default:
		b, _ := ioutil.ReadAll(res.Body)
		errRes := &ErrorResponse{Response: res}
		err := jsonx.Unmarshal(b, errRes)
		if err != nil {
			return fmt.Errorf("%s: %s", res.Status, string(b))
		}
		return errRes
	}
}

// JSONRequest issues HTTP request with JSON payload
func (r *BaseRequest) JSONRequest(ctx context.Context, method, path string, reqObj, resObj interface{}) error {
	req, err := r.NewJSONRequest(method, path, reqObj)
	if err != nil {
		return err
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}
	res, err := r.client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	return r.DecodeJSONResponse(res, resObj)
}

// GraphServiceRequestBuilder is GraphService reuqest builder
type GraphServiceRequestBuilder struct {
	BaseRequestBuilder
}

// NewClient returns GraphService request builder with default base URL
func NewClient(cli *http.Client, tenantId string, applicationId string, clientSecurityKey string) *GraphServiceRequestBuilder {
	return &GraphServiceRequestBuilder{
		BaseRequestBuilder: BaseRequestBuilder{baseURL: defaultBaseURL, client: cli,
			tenantID: tenantId, applicationID: applicationId, clientSecurityKey: clientSecurityKey},
	}
}
