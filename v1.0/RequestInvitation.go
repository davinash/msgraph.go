// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// InvitationRequestBuilder is request builder for Invitation
type InvitationRequestBuilder struct{ BaseRequestBuilder }

// Request returns InvitationRequest
func (b *InvitationRequestBuilder) Request() *InvitationRequest {
	return &InvitationRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client,
			tenantID: b.tenantID, applicationID: b.applicationID, clientSecurityKey: b.clientSecurityKey},
	}
}

// InvitationRequest is request for Invitation
type InvitationRequest struct{ BaseRequest }

// Get performs GET request for Invitation
func (r *InvitationRequest) Get(ctx context.Context) (resObj *Invitation, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for Invitation
func (r *InvitationRequest) Update(ctx context.Context, reqObj *Invitation) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for Invitation
func (r *InvitationRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
