// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

//
type WorkbookFunctionsFindRequestBuilder struct{ BaseRequestBuilder }

// Find action undocumented
func (b *WorkbookFunctionsRequestBuilder) Find(reqObj *WorkbookFunctionsFindRequestParameter) *WorkbookFunctionsFindRequestBuilder {
	bb := &WorkbookFunctionsFindRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/find"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type WorkbookFunctionsFindRequest struct{ BaseRequest }

//
func (b *WorkbookFunctionsFindRequestBuilder) Request() *WorkbookFunctionsFindRequest {
	return &WorkbookFunctionsFindRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject,
			tenantID: b.tenantID, applicationID: b.applicationID, clientSecurityKey: b.clientSecurityKey},
	}
}

//
func (r *WorkbookFunctionsFindRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

//
type WorkbookFunctionsFindBRequestBuilder struct{ BaseRequestBuilder }

// FindB action undocumented
func (b *WorkbookFunctionsRequestBuilder) FindB(reqObj *WorkbookFunctionsFindBRequestParameter) *WorkbookFunctionsFindBRequestBuilder {
	bb := &WorkbookFunctionsFindBRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/findB"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type WorkbookFunctionsFindBRequest struct{ BaseRequest }

//
func (b *WorkbookFunctionsFindBRequestBuilder) Request() *WorkbookFunctionsFindBRequest {
	return &WorkbookFunctionsFindBRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject,
			tenantID: b.tenantID, applicationID: b.applicationID, clientSecurityKey: b.clientSecurityKey},
	}
}

//
func (r *WorkbookFunctionsFindBRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}
