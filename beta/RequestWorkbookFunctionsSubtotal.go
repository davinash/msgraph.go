// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

//
type WorkbookFunctionsSubtotalRequestBuilder struct{ BaseRequestBuilder }

// Subtotal action undocumented
func (b *WorkbookFunctionsRequestBuilder) Subtotal(reqObj *WorkbookFunctionsSubtotalRequestParameter) *WorkbookFunctionsSubtotalRequestBuilder {
	bb := &WorkbookFunctionsSubtotalRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/subtotal"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type WorkbookFunctionsSubtotalRequest struct{ BaseRequest }

//
func (b *WorkbookFunctionsSubtotalRequestBuilder) Request() *WorkbookFunctionsSubtotalRequest {
	return &WorkbookFunctionsSubtotalRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject,
			tenantID: b.tenantID, applicationID: b.applicationID, clientSecurityKey: b.clientSecurityKey},
	}
}

//
func (r *WorkbookFunctionsSubtotalRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}
