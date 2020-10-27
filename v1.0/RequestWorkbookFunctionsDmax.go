// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

//
type WorkbookFunctionsDmaxRequestBuilder struct{ BaseRequestBuilder }

// Dmax action undocumented
func (b *WorkbookFunctionsRequestBuilder) Dmax(reqObj *WorkbookFunctionsDmaxRequestParameter) *WorkbookFunctionsDmaxRequestBuilder {
	bb := &WorkbookFunctionsDmaxRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/dmax"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type WorkbookFunctionsDmaxRequest struct{ BaseRequest }

//
func (b *WorkbookFunctionsDmaxRequestBuilder) Request() *WorkbookFunctionsDmaxRequest {
	return &WorkbookFunctionsDmaxRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject,
			tenantID: b.tenantID, applicationID: b.applicationID, clientSecurityKey: b.clientSecurityKey},
	}
}

//
func (r *WorkbookFunctionsDmaxRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}
