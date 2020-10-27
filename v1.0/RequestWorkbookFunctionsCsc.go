// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

//
type WorkbookFunctionsCscRequestBuilder struct{ BaseRequestBuilder }

// Csc action undocumented
func (b *WorkbookFunctionsRequestBuilder) Csc(reqObj *WorkbookFunctionsCscRequestParameter) *WorkbookFunctionsCscRequestBuilder {
	bb := &WorkbookFunctionsCscRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/csc"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type WorkbookFunctionsCscRequest struct{ BaseRequest }

//
func (b *WorkbookFunctionsCscRequestBuilder) Request() *WorkbookFunctionsCscRequest {
	return &WorkbookFunctionsCscRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject,
			tenantID: b.tenantID, applicationID: b.applicationID, clientSecurityKey: b.clientSecurityKey},
	}
}

//
func (r *WorkbookFunctionsCscRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}
