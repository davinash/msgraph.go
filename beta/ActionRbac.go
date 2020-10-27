// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/davinash/msgraph.go/jsonx"
)

// RoleAssignments returns request builder for UnifiedRoleAssignment collection
func (b *RbacApplicationRequestBuilder) RoleAssignments() *RbacApplicationRoleAssignmentsCollectionRequestBuilder {
	bb := &RbacApplicationRoleAssignmentsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/roleAssignments"
	return bb
}

// RbacApplicationRoleAssignmentsCollectionRequestBuilder is request builder for UnifiedRoleAssignment collection
type RbacApplicationRoleAssignmentsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for UnifiedRoleAssignment collection
func (b *RbacApplicationRoleAssignmentsCollectionRequestBuilder) Request() *RbacApplicationRoleAssignmentsCollectionRequest {
	return &RbacApplicationRoleAssignmentsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client,
			tenantID: b.tenantID, applicationID: b.applicationID, clientSecurityKey: b.clientSecurityKey},
	}
}

// ID returns request builder for UnifiedRoleAssignment item
func (b *RbacApplicationRoleAssignmentsCollectionRequestBuilder) ID(id string) *UnifiedRoleAssignmentRequestBuilder {
	bb := &UnifiedRoleAssignmentRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// RbacApplicationRoleAssignmentsCollectionRequest is request for UnifiedRoleAssignment collection
type RbacApplicationRoleAssignmentsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for UnifiedRoleAssignment collection
func (r *RbacApplicationRoleAssignmentsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]UnifiedRoleAssignment, error) {
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
	var values []UnifiedRoleAssignment
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
			value  []UnifiedRoleAssignment
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
		req, err = r.NewRequest("GET", paging.NextLink, nil)
		if ctx != nil {
			req = req.WithContext(ctx)
		}
		res, err = r.client.Do(req)
		if err != nil {
			return nil, err
		}
	}
}

// NewRequest Wrapper over the http.NewRequest with adding auth tokens
func (r *RbacApplicationRoleAssignmentsCollectionRequest) NewRequest(method, url string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequestWithContext(context.Background(), method, url, body)
	if err != nil {
		return nil, err
	}
	err = r.getAuthToken()
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", r.token.GetAccessToken())
	return req, err
}

// GetN performs GET request for UnifiedRoleAssignment collection, max N pages
func (r *RbacApplicationRoleAssignmentsCollectionRequest) GetN(ctx context.Context, n int) ([]UnifiedRoleAssignment, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for UnifiedRoleAssignment collection
func (r *RbacApplicationRoleAssignmentsCollectionRequest) Get(ctx context.Context) ([]UnifiedRoleAssignment, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for UnifiedRoleAssignment collection
func (r *RbacApplicationRoleAssignmentsCollectionRequest) Add(ctx context.Context, reqObj *UnifiedRoleAssignment) (resObj *UnifiedRoleAssignment, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// RoleDefinitions returns request builder for UnifiedRoleDefinition collection
func (b *RbacApplicationRequestBuilder) RoleDefinitions() *RbacApplicationRoleDefinitionsCollectionRequestBuilder {
	bb := &RbacApplicationRoleDefinitionsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/roleDefinitions"
	return bb
}

// RbacApplicationRoleDefinitionsCollectionRequestBuilder is request builder for UnifiedRoleDefinition collection
type RbacApplicationRoleDefinitionsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for UnifiedRoleDefinition collection
func (b *RbacApplicationRoleDefinitionsCollectionRequestBuilder) Request() *RbacApplicationRoleDefinitionsCollectionRequest {
	return &RbacApplicationRoleDefinitionsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client,
			tenantID: b.tenantID, applicationID: b.applicationID, clientSecurityKey: b.clientSecurityKey},
	}
}

// ID returns request builder for UnifiedRoleDefinition item
func (b *RbacApplicationRoleDefinitionsCollectionRequestBuilder) ID(id string) *UnifiedRoleDefinitionRequestBuilder {
	bb := &UnifiedRoleDefinitionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// RbacApplicationRoleDefinitionsCollectionRequest is request for UnifiedRoleDefinition collection
type RbacApplicationRoleDefinitionsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for UnifiedRoleDefinition collection
func (r *RbacApplicationRoleDefinitionsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]UnifiedRoleDefinition, error) {
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
	var values []UnifiedRoleDefinition
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
			value  []UnifiedRoleDefinition
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
		req, err = r.NewRequest("GET", paging.NextLink, nil)
		if ctx != nil {
			req = req.WithContext(ctx)
		}
		res, err = r.client.Do(req)
		if err != nil {
			return nil, err
		}
	}
}

// NewRequest Wrapper over the http.NewRequest with adding auth tokens
func (r *RbacApplicationRoleDefinitionsCollectionRequest) NewRequest(method, url string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequestWithContext(context.Background(), method, url, body)
	if err != nil {
		return nil, err
	}
	err = r.getAuthToken()
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", r.token.GetAccessToken())
	return req, err
}

// GetN performs GET request for UnifiedRoleDefinition collection, max N pages
func (r *RbacApplicationRoleDefinitionsCollectionRequest) GetN(ctx context.Context, n int) ([]UnifiedRoleDefinition, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for UnifiedRoleDefinition collection
func (r *RbacApplicationRoleDefinitionsCollectionRequest) Get(ctx context.Context) ([]UnifiedRoleDefinition, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for UnifiedRoleDefinition collection
func (r *RbacApplicationRoleDefinitionsCollectionRequest) Add(ctx context.Context, reqObj *UnifiedRoleDefinition) (resObj *UnifiedRoleDefinition, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}
