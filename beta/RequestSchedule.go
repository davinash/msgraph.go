// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// ScheduleRequestBuilder is request builder for Schedule
type ScheduleRequestBuilder struct{ BaseRequestBuilder }

// Request returns ScheduleRequest
func (b *ScheduleRequestBuilder) Request() *ScheduleRequest {
	return &ScheduleRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client,
			tenantID: b.tenantID, applicationID: b.applicationID, clientSecurityKey: b.clientSecurityKey},
	}
}

// ScheduleRequest is request for Schedule
type ScheduleRequest struct{ BaseRequest }

// Get performs GET request for Schedule
func (r *ScheduleRequest) Get(ctx context.Context) (resObj *Schedule, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for Schedule
func (r *ScheduleRequest) Update(ctx context.Context, reqObj *Schedule) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for Schedule
func (r *ScheduleRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// ScheduleChangeRequestObjectRequestBuilder is request builder for ScheduleChangeRequestObject
type ScheduleChangeRequestObjectRequestBuilder struct{ BaseRequestBuilder }

// Request returns ScheduleChangeRequestObjectRequest
func (b *ScheduleChangeRequestObjectRequestBuilder) Request() *ScheduleChangeRequestObjectRequest {
	return &ScheduleChangeRequestObjectRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client,
			tenantID: b.tenantID, applicationID: b.applicationID, clientSecurityKey: b.clientSecurityKey},
	}
}

// ScheduleChangeRequestObjectRequest is request for ScheduleChangeRequestObject
type ScheduleChangeRequestObjectRequest struct{ BaseRequest }

// Get performs GET request for ScheduleChangeRequestObject
func (r *ScheduleChangeRequestObjectRequest) Get(ctx context.Context) (resObj *ScheduleChangeRequestObject, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ScheduleChangeRequestObject
func (r *ScheduleChangeRequestObjectRequest) Update(ctx context.Context, reqObj *ScheduleChangeRequestObject) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ScheduleChangeRequestObject
func (r *ScheduleChangeRequestObjectRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

//
type ScheduleShareRequestBuilder struct{ BaseRequestBuilder }

// Share action undocumented
func (b *ScheduleRequestBuilder) Share(reqObj *ScheduleShareRequestParameter) *ScheduleShareRequestBuilder {
	bb := &ScheduleShareRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/share"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type ScheduleShareRequest struct{ BaseRequest }

//
func (b *ScheduleShareRequestBuilder) Request() *ScheduleShareRequest {
	return &ScheduleShareRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject,
			tenantID: b.tenantID, applicationID: b.applicationID, clientSecurityKey: b.clientSecurityKey},
	}
}

//
func (r *ScheduleShareRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

//
type ScheduleChangeRequestObjectApproveRequestBuilder struct{ BaseRequestBuilder }

// Approve action undocumented
func (b *ScheduleChangeRequestObjectRequestBuilder) Approve(reqObj *ScheduleChangeRequestObjectApproveRequestParameter) *ScheduleChangeRequestObjectApproveRequestBuilder {
	bb := &ScheduleChangeRequestObjectApproveRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/approve"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type ScheduleChangeRequestObjectApproveRequest struct{ BaseRequest }

//
func (b *ScheduleChangeRequestObjectApproveRequestBuilder) Request() *ScheduleChangeRequestObjectApproveRequest {
	return &ScheduleChangeRequestObjectApproveRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject,
			tenantID: b.tenantID, applicationID: b.applicationID, clientSecurityKey: b.clientSecurityKey},
	}
}

//
func (r *ScheduleChangeRequestObjectApproveRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

//
type ScheduleChangeRequestObjectDeclineRequestBuilder struct{ BaseRequestBuilder }

// Decline action undocumented
func (b *ScheduleChangeRequestObjectRequestBuilder) Decline(reqObj *ScheduleChangeRequestObjectDeclineRequestParameter) *ScheduleChangeRequestObjectDeclineRequestBuilder {
	bb := &ScheduleChangeRequestObjectDeclineRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/decline"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type ScheduleChangeRequestObjectDeclineRequest struct{ BaseRequest }

//
func (b *ScheduleChangeRequestObjectDeclineRequestBuilder) Request() *ScheduleChangeRequestObjectDeclineRequest {
	return &ScheduleChangeRequestObjectDeclineRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject,
			tenantID: b.tenantID, applicationID: b.applicationID, clientSecurityKey: b.clientSecurityKey},
	}
}

//
func (r *ScheduleChangeRequestObjectDeclineRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}
