// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/davinash/msgraph.go/jsonx"
)

// TargetedManagedAppConfigurationRequestBuilder is request builder for TargetedManagedAppConfiguration
type TargetedManagedAppConfigurationRequestBuilder struct{ BaseRequestBuilder }

// Request returns TargetedManagedAppConfigurationRequest
func (b *TargetedManagedAppConfigurationRequestBuilder) Request() *TargetedManagedAppConfigurationRequest {
	return &TargetedManagedAppConfigurationRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// TargetedManagedAppConfigurationRequest is request for TargetedManagedAppConfiguration
type TargetedManagedAppConfigurationRequest struct{ BaseRequest }

// Get performs GET request for TargetedManagedAppConfiguration
func (r *TargetedManagedAppConfigurationRequest) Get(ctx context.Context) (resObj *TargetedManagedAppConfiguration, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for TargetedManagedAppConfiguration
func (r *TargetedManagedAppConfigurationRequest) Update(ctx context.Context, reqObj *TargetedManagedAppConfiguration) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for TargetedManagedAppConfiguration
func (r *TargetedManagedAppConfigurationRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// TargetedManagedAppPolicyAssignmentRequestBuilder is request builder for TargetedManagedAppPolicyAssignment
type TargetedManagedAppPolicyAssignmentRequestBuilder struct{ BaseRequestBuilder }

// Request returns TargetedManagedAppPolicyAssignmentRequest
func (b *TargetedManagedAppPolicyAssignmentRequestBuilder) Request() *TargetedManagedAppPolicyAssignmentRequest {
	return &TargetedManagedAppPolicyAssignmentRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// TargetedManagedAppPolicyAssignmentRequest is request for TargetedManagedAppPolicyAssignment
type TargetedManagedAppPolicyAssignmentRequest struct{ BaseRequest }

// Get performs GET request for TargetedManagedAppPolicyAssignment
func (r *TargetedManagedAppPolicyAssignmentRequest) Get(ctx context.Context) (resObj *TargetedManagedAppPolicyAssignment, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for TargetedManagedAppPolicyAssignment
func (r *TargetedManagedAppPolicyAssignmentRequest) Update(ctx context.Context, reqObj *TargetedManagedAppPolicyAssignment) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for TargetedManagedAppPolicyAssignment
func (r *TargetedManagedAppPolicyAssignmentRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// TargetedManagedAppProtectionRequestBuilder is request builder for TargetedManagedAppProtection
type TargetedManagedAppProtectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns TargetedManagedAppProtectionRequest
func (b *TargetedManagedAppProtectionRequestBuilder) Request() *TargetedManagedAppProtectionRequest {
	return &TargetedManagedAppProtectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// TargetedManagedAppProtectionRequest is request for TargetedManagedAppProtection
type TargetedManagedAppProtectionRequest struct{ BaseRequest }

// Get performs GET request for TargetedManagedAppProtection
func (r *TargetedManagedAppProtectionRequest) Get(ctx context.Context) (resObj *TargetedManagedAppProtection, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for TargetedManagedAppProtection
func (r *TargetedManagedAppProtectionRequest) Update(ctx context.Context, reqObj *TargetedManagedAppProtection) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for TargetedManagedAppProtection
func (r *TargetedManagedAppProtectionRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

//
type TargetedManagedAppConfigurationCollectionHasPayloadLinksRequestBuilder struct{ BaseRequestBuilder }

// HasPayloadLinks action undocumented
func (b *DeviceAppManagementTargetedManagedAppConfigurationsCollectionRequestBuilder) HasPayloadLinks(reqObj *TargetedManagedAppConfigurationCollectionHasPayloadLinksRequestParameter) *TargetedManagedAppConfigurationCollectionHasPayloadLinksRequestBuilder {
	bb := &TargetedManagedAppConfigurationCollectionHasPayloadLinksRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/hasPayloadLinks"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type TargetedManagedAppConfigurationCollectionHasPayloadLinksRequest struct{ BaseRequest }

//
func (b *TargetedManagedAppConfigurationCollectionHasPayloadLinksRequestBuilder) Request() *TargetedManagedAppConfigurationCollectionHasPayloadLinksRequest {
	return &TargetedManagedAppConfigurationCollectionHasPayloadLinksRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *TargetedManagedAppConfigurationCollectionHasPayloadLinksRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]HasPayloadLinkResultItem, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []HasPayloadLinkResultItem
	for {
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			res.Body.Close()
			errRes := &ErrorResponse{Response: res}
			err := jsonx.Unmarshal(b, errRes)
			if err != nil {
				return nil, fmt.Errorf("%s: %s", res.Status, string(b))
			}
			return nil, errRes
		}
		var (
			paging Paging
			value  []HasPayloadLinkResultItem
		)
		err := jsonx.NewDecoder(res.Body).Decode(&paging)
		res.Body.Close()
		if err != nil {
			return nil, err
		}
		err = jsonx.Unmarshal(paging.Value, &value)
		if err != nil {
			return nil, err
		}
		values = append(values, value...)
		if n >= 0 {
			n--
		}
		if n == 0 || len(paging.NextLink) == 0 {
			return values, nil
		}
		req, err = http.NewRequest("GET", paging.NextLink, nil)
		if ctx != nil {
			req = req.WithContext(ctx)
		}
		res, err = r.client.Do(req)
		if err != nil {
			return nil, err
		}
	}
}

//
func (r *TargetedManagedAppConfigurationCollectionHasPayloadLinksRequest) PostN(ctx context.Context, n int) ([]HasPayloadLinkResultItem, error) {
	return r.Paging(ctx, "POST", "", r.requestObject, n)
}

//
func (r *TargetedManagedAppConfigurationCollectionHasPayloadLinksRequest) Post(ctx context.Context) ([]HasPayloadLinkResultItem, error) {
	return r.Paging(ctx, "POST", "", r.requestObject, 0)
}

//
type TargetedManagedAppConfigurationAssignRequestBuilder struct{ BaseRequestBuilder }

// Assign action undocumented
func (b *TargetedManagedAppConfigurationRequestBuilder) Assign(reqObj *TargetedManagedAppConfigurationAssignRequestParameter) *TargetedManagedAppConfigurationAssignRequestBuilder {
	bb := &TargetedManagedAppConfigurationAssignRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/assign"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type TargetedManagedAppConfigurationAssignRequest struct{ BaseRequest }

//
func (b *TargetedManagedAppConfigurationAssignRequestBuilder) Request() *TargetedManagedAppConfigurationAssignRequest {
	return &TargetedManagedAppConfigurationAssignRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *TargetedManagedAppConfigurationAssignRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

//
type TargetedManagedAppConfigurationTargetAppsRequestBuilder struct{ BaseRequestBuilder }

// TargetApps action undocumented
func (b *TargetedManagedAppConfigurationRequestBuilder) TargetApps(reqObj *TargetedManagedAppConfigurationTargetAppsRequestParameter) *TargetedManagedAppConfigurationTargetAppsRequestBuilder {
	bb := &TargetedManagedAppConfigurationTargetAppsRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/targetApps"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type TargetedManagedAppConfigurationTargetAppsRequest struct{ BaseRequest }

//
func (b *TargetedManagedAppConfigurationTargetAppsRequestBuilder) Request() *TargetedManagedAppConfigurationTargetAppsRequest {
	return &TargetedManagedAppConfigurationTargetAppsRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *TargetedManagedAppConfigurationTargetAppsRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

//
type TargetedManagedAppProtectionAssignRequestBuilder struct{ BaseRequestBuilder }

// Assign action undocumented
func (b *TargetedManagedAppProtectionRequestBuilder) Assign(reqObj *TargetedManagedAppProtectionAssignRequestParameter) *TargetedManagedAppProtectionAssignRequestBuilder {
	bb := &TargetedManagedAppProtectionAssignRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/assign"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type TargetedManagedAppProtectionAssignRequest struct{ BaseRequest }

//
func (b *TargetedManagedAppProtectionAssignRequestBuilder) Request() *TargetedManagedAppProtectionAssignRequest {
	return &TargetedManagedAppProtectionAssignRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *TargetedManagedAppProtectionAssignRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}
