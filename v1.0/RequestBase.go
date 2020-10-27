// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// BaseItemRequestBuilder is request builder for BaseItem
type BaseItemRequestBuilder struct{ BaseRequestBuilder }

// Request returns BaseItemRequest
func (b *BaseItemRequestBuilder) Request() *BaseItemRequest {
	return &BaseItemRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client,
			tenantID: b.tenantID, applicationID: b.applicationID, clientSecurityKey: b.clientSecurityKey},
	}
}

// BaseItemRequest is request for BaseItem
type BaseItemRequest struct{ BaseRequest }

// Get performs GET request for BaseItem
func (r *BaseItemRequest) Get(ctx context.Context) (resObj *BaseItem, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for BaseItem
func (r *BaseItemRequest) Update(ctx context.Context, reqObj *BaseItem) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for BaseItem
func (r *BaseItemRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
