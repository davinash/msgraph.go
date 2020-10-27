// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// AccountRequestBuilder is request builder for Account
type AccountRequestBuilder struct{ BaseRequestBuilder }

// Request returns AccountRequest
func (b *AccountRequestBuilder) Request() *AccountRequest {
	return &AccountRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client,
			tenantID: b.tenantID, applicationID: b.applicationID, clientSecurityKey: b.clientSecurityKey},
	}
}

// AccountRequest is request for Account
type AccountRequest struct{ BaseRequest }

// Get performs GET request for Account
func (r *AccountRequest) Get(ctx context.Context) (resObj *Account, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for Account
func (r *AccountRequest) Update(ctx context.Context, reqObj *Account) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for Account
func (r *AccountRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
