// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// ChannelRequestBuilder is request builder for Channel
type ChannelRequestBuilder struct{ BaseRequestBuilder }

// Request returns ChannelRequest
func (b *ChannelRequestBuilder) Request() *ChannelRequest {
	return &ChannelRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client,
			tenantID: b.tenantID, applicationID: b.applicationID, clientSecurityKey: b.clientSecurityKey},
	}
}

// ChannelRequest is request for Channel
type ChannelRequest struct{ BaseRequest }

// Get performs GET request for Channel
func (r *ChannelRequest) Get(ctx context.Context) (resObj *Channel, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for Channel
func (r *ChannelRequest) Update(ctx context.Context, reqObj *Channel) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for Channel
func (r *ChannelRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
