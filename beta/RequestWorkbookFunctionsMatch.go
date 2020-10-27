// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

//
type WorkbookFunctionsMatchRequestBuilder struct{ BaseRequestBuilder }

// Match action undocumented
func (b *WorkbookFunctionsRequestBuilder) Match(reqObj *WorkbookFunctionsMatchRequestParameter) *WorkbookFunctionsMatchRequestBuilder {
	bb := &WorkbookFunctionsMatchRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/match"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type WorkbookFunctionsMatchRequest struct{ BaseRequest }

//
func (b *WorkbookFunctionsMatchRequestBuilder) Request() *WorkbookFunctionsMatchRequest {
	return &WorkbookFunctionsMatchRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject,
			tenantID: b.tenantID, applicationID: b.applicationID, clientSecurityKey: b.clientSecurityKey},
	}
}

//
func (r *WorkbookFunctionsMatchRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}
