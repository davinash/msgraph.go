// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

//
type WorkbookFunctionsProductRequestBuilder struct{ BaseRequestBuilder }

// Product action undocumented
func (b *WorkbookFunctionsRequestBuilder) Product(reqObj *WorkbookFunctionsProductRequestParameter) *WorkbookFunctionsProductRequestBuilder {
	bb := &WorkbookFunctionsProductRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/product"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type WorkbookFunctionsProductRequest struct{ BaseRequest }

//
func (b *WorkbookFunctionsProductRequestBuilder) Request() *WorkbookFunctionsProductRequest {
	return &WorkbookFunctionsProductRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject,
			tenantID: b.tenantID, applicationID: b.applicationID, clientSecurityKey: b.clientSecurityKey},
	}
}

//
func (r *WorkbookFunctionsProductRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}
