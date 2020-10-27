// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

//
type WorkbookFunctionsAccrIntRequestBuilder struct{ BaseRequestBuilder }

// AccrInt action undocumented
func (b *WorkbookFunctionsRequestBuilder) AccrInt(reqObj *WorkbookFunctionsAccrIntRequestParameter) *WorkbookFunctionsAccrIntRequestBuilder {
	bb := &WorkbookFunctionsAccrIntRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/accrInt"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type WorkbookFunctionsAccrIntRequest struct{ BaseRequest }

//
func (b *WorkbookFunctionsAccrIntRequestBuilder) Request() *WorkbookFunctionsAccrIntRequest {
	return &WorkbookFunctionsAccrIntRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject,
			tenantID: b.tenantID, applicationID: b.applicationID, clientSecurityKey: b.clientSecurityKey},
	}
}

//
func (r *WorkbookFunctionsAccrIntRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

//
type WorkbookFunctionsAccrIntMRequestBuilder struct{ BaseRequestBuilder }

// AccrIntM action undocumented
func (b *WorkbookFunctionsRequestBuilder) AccrIntM(reqObj *WorkbookFunctionsAccrIntMRequestParameter) *WorkbookFunctionsAccrIntMRequestBuilder {
	bb := &WorkbookFunctionsAccrIntMRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/accrIntM"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type WorkbookFunctionsAccrIntMRequest struct{ BaseRequest }

//
func (b *WorkbookFunctionsAccrIntMRequestBuilder) Request() *WorkbookFunctionsAccrIntMRequest {
	return &WorkbookFunctionsAccrIntMRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject,
			tenantID: b.tenantID, applicationID: b.applicationID, clientSecurityKey: b.clientSecurityKey},
	}
}

//
func (r *WorkbookFunctionsAccrIntMRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}
