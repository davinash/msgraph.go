// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// PersonRequestBuilder is request builder for Person
type PersonRequestBuilder struct{ BaseRequestBuilder }

// Request returns PersonRequest
func (b *PersonRequestBuilder) Request() *PersonRequest {
	return &PersonRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client,
			tenantID: b.tenantID, applicationID: b.applicationID, clientSecurityKey: b.clientSecurityKey},
	}
}

// PersonRequest is request for Person
type PersonRequest struct{ BaseRequest }

// Get performs GET request for Person
func (r *PersonRequest) Get(ctx context.Context) (resObj *Person, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for Person
func (r *PersonRequest) Update(ctx context.Context, reqObj *Person) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for Person
func (r *PersonRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// PersonAnniversaryRequestBuilder is request builder for PersonAnniversary
type PersonAnniversaryRequestBuilder struct{ BaseRequestBuilder }

// Request returns PersonAnniversaryRequest
func (b *PersonAnniversaryRequestBuilder) Request() *PersonAnniversaryRequest {
	return &PersonAnniversaryRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client,
			tenantID: b.tenantID, applicationID: b.applicationID, clientSecurityKey: b.clientSecurityKey},
	}
}

// PersonAnniversaryRequest is request for PersonAnniversary
type PersonAnniversaryRequest struct{ BaseRequest }

// Get performs GET request for PersonAnniversary
func (r *PersonAnniversaryRequest) Get(ctx context.Context) (resObj *PersonAnniversary, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for PersonAnniversary
func (r *PersonAnniversaryRequest) Update(ctx context.Context, reqObj *PersonAnniversary) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for PersonAnniversary
func (r *PersonAnniversaryRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// PersonInterestRequestBuilder is request builder for PersonInterest
type PersonInterestRequestBuilder struct{ BaseRequestBuilder }

// Request returns PersonInterestRequest
func (b *PersonInterestRequestBuilder) Request() *PersonInterestRequest {
	return &PersonInterestRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client,
			tenantID: b.tenantID, applicationID: b.applicationID, clientSecurityKey: b.clientSecurityKey},
	}
}

// PersonInterestRequest is request for PersonInterest
type PersonInterestRequest struct{ BaseRequest }

// Get performs GET request for PersonInterest
func (r *PersonInterestRequest) Get(ctx context.Context) (resObj *PersonInterest, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for PersonInterest
func (r *PersonInterestRequest) Update(ctx context.Context, reqObj *PersonInterest) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for PersonInterest
func (r *PersonInterestRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// PersonNameRequestBuilder is request builder for PersonName
type PersonNameRequestBuilder struct{ BaseRequestBuilder }

// Request returns PersonNameRequest
func (b *PersonNameRequestBuilder) Request() *PersonNameRequest {
	return &PersonNameRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client,
			tenantID: b.tenantID, applicationID: b.applicationID, clientSecurityKey: b.clientSecurityKey},
	}
}

// PersonNameRequest is request for PersonName
type PersonNameRequest struct{ BaseRequest }

// Get performs GET request for PersonName
func (r *PersonNameRequest) Get(ctx context.Context) (resObj *PersonName, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for PersonName
func (r *PersonNameRequest) Update(ctx context.Context, reqObj *PersonName) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for PersonName
func (r *PersonNameRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// PersonWebsiteRequestBuilder is request builder for PersonWebsite
type PersonWebsiteRequestBuilder struct{ BaseRequestBuilder }

// Request returns PersonWebsiteRequest
func (b *PersonWebsiteRequestBuilder) Request() *PersonWebsiteRequest {
	return &PersonWebsiteRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client,
			tenantID: b.tenantID, applicationID: b.applicationID, clientSecurityKey: b.clientSecurityKey},
	}
}

// PersonWebsiteRequest is request for PersonWebsite
type PersonWebsiteRequest struct{ BaseRequest }

// Get performs GET request for PersonWebsite
func (r *PersonWebsiteRequest) Get(ctx context.Context) (resObj *PersonWebsite, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for PersonWebsite
func (r *PersonWebsiteRequest) Update(ctx context.Context, reqObj *PersonWebsite) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for PersonWebsite
func (r *PersonWebsiteRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
