// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

//
type WorkbookFunctionsBitorRequestBuilder struct{ BaseRequestBuilder }

// Bitor action undocumented
func (b *WorkbookFunctionsRequestBuilder) Bitor(reqObj *WorkbookFunctionsBitorRequestParameter) *WorkbookFunctionsBitorRequestBuilder {
	bb := &WorkbookFunctionsBitorRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/bitor"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type WorkbookFunctionsBitorRequest struct{ BaseRequest }

//
func (b *WorkbookFunctionsBitorRequestBuilder) Request() *WorkbookFunctionsBitorRequest {
	return &WorkbookFunctionsBitorRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject,
			tenantID: b.tenantID, applicationID: b.applicationID, clientSecurityKey: b.clientSecurityKey},
	}
}

//
func (r *WorkbookFunctionsBitorRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}
