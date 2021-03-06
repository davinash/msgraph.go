// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// ResourceOperationRequestBuilder is request builder for ResourceOperation
type ResourceOperationRequestBuilder struct{ BaseRequestBuilder }

// Request returns ResourceOperationRequest
func (b *ResourceOperationRequestBuilder) Request() *ResourceOperationRequest {
	return &ResourceOperationRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client,
			tenantID: b.tenantID, applicationID: b.applicationID, clientSecurityKey: b.clientSecurityKey},
	}
}

// ResourceOperationRequest is request for ResourceOperation
type ResourceOperationRequest struct{ BaseRequest }

// Get performs GET request for ResourceOperation
func (r *ResourceOperationRequest) Get(ctx context.Context) (resObj *ResourceOperation, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ResourceOperation
func (r *ResourceOperationRequest) Update(ctx context.Context, reqObj *ResourceOperation) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ResourceOperation
func (r *ResourceOperationRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
