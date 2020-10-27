// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// RequestObjectRequestBuilder is request builder for RequestObject
type RequestObjectRequestBuilder struct{ BaseRequestBuilder }

// Request returns RequestObjectRequest
func (b *RequestObjectRequestBuilder) Request() *RequestObjectRequest {
	return &RequestObjectRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client,
			tenantID: b.tenantID, applicationID: b.applicationID, clientSecurityKey: b.clientSecurityKey},
	}
}

// RequestObjectRequest is request for RequestObject
type RequestObjectRequest struct{ BaseRequest }

// Get performs GET request for RequestObject
func (r *RequestObjectRequest) Get(ctx context.Context) (resObj *RequestObject, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for RequestObject
func (r *RequestObjectRequest) Update(ctx context.Context, reqObj *RequestObject) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for RequestObject
func (r *RequestObjectRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

//
type RequestObjectStopRequestBuilder struct{ BaseRequestBuilder }

// Stop action undocumented
func (b *RequestObjectRequestBuilder) Stop(reqObj *RequestObjectStopRequestParameter) *RequestObjectStopRequestBuilder {
	bb := &RequestObjectStopRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/stop"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type RequestObjectStopRequest struct{ BaseRequest }

//
func (b *RequestObjectStopRequestBuilder) Request() *RequestObjectStopRequest {
	return &RequestObjectStopRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject,
			tenantID: b.tenantID, applicationID: b.applicationID, clientSecurityKey: b.clientSecurityKey},
	}
}

//
func (r *RequestObjectStopRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

//
type RequestObjectRecordDecisionsRequestBuilder struct{ BaseRequestBuilder }

// RecordDecisions action undocumented
func (b *RequestObjectRequestBuilder) RecordDecisions(reqObj *RequestObjectRecordDecisionsRequestParameter) *RequestObjectRecordDecisionsRequestBuilder {
	bb := &RequestObjectRecordDecisionsRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/recordDecisions"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type RequestObjectRecordDecisionsRequest struct{ BaseRequest }

//
func (b *RequestObjectRecordDecisionsRequestBuilder) Request() *RequestObjectRecordDecisionsRequest {
	return &RequestObjectRecordDecisionsRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject,
			tenantID: b.tenantID, applicationID: b.applicationID, clientSecurityKey: b.clientSecurityKey},
	}
}

//
func (r *RequestObjectRecordDecisionsRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}
