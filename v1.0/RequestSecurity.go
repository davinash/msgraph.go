// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// SecurityRequestBuilder is request builder for Security
type SecurityRequestBuilder struct{ BaseRequestBuilder }

// Request returns SecurityRequest
func (b *SecurityRequestBuilder) Request() *SecurityRequest {
	return &SecurityRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client,
			tenantID: b.tenantID, applicationID: b.applicationID, clientSecurityKey: b.clientSecurityKey},
	}
}

// SecurityRequest is request for Security
type SecurityRequest struct{ BaseRequest }

// Get performs GET request for Security
func (r *SecurityRequest) Get(ctx context.Context) (resObj *Security, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for Security
func (r *SecurityRequest) Update(ctx context.Context, reqObj *Security) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for Security
func (r *SecurityRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
