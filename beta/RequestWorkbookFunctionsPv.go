// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

//
type WorkbookFunctionsPvRequestBuilder struct{ BaseRequestBuilder }

// Pv action undocumented
func (b *WorkbookFunctionsRequestBuilder) Pv(reqObj *WorkbookFunctionsPvRequestParameter) *WorkbookFunctionsPvRequestBuilder {
	bb := &WorkbookFunctionsPvRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/pv"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type WorkbookFunctionsPvRequest struct{ BaseRequest }

//
func (b *WorkbookFunctionsPvRequestBuilder) Request() *WorkbookFunctionsPvRequest {
	return &WorkbookFunctionsPvRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject,
			tenantID: b.tenantID, applicationID: b.applicationID, clientSecurityKey: b.clientSecurityKey},
	}
}

//
func (r *WorkbookFunctionsPvRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}
