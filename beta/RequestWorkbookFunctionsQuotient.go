// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

//
type WorkbookFunctionsQuotientRequestBuilder struct{ BaseRequestBuilder }

// Quotient action undocumented
func (b *WorkbookFunctionsRequestBuilder) Quotient(reqObj *WorkbookFunctionsQuotientRequestParameter) *WorkbookFunctionsQuotientRequestBuilder {
	bb := &WorkbookFunctionsQuotientRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/quotient"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type WorkbookFunctionsQuotientRequest struct{ BaseRequest }

//
func (b *WorkbookFunctionsQuotientRequestBuilder) Request() *WorkbookFunctionsQuotientRequest {
	return &WorkbookFunctionsQuotientRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject,
			tenantID: b.tenantID, applicationID: b.applicationID, clientSecurityKey: b.clientSecurityKey},
	}
}

//
func (r *WorkbookFunctionsQuotientRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}
