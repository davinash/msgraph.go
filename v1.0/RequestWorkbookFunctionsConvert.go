// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

//
type WorkbookFunctionsConvertRequestBuilder struct{ BaseRequestBuilder }

// Convert action undocumented
func (b *WorkbookFunctionsRequestBuilder) Convert(reqObj *WorkbookFunctionsConvertRequestParameter) *WorkbookFunctionsConvertRequestBuilder {
	bb := &WorkbookFunctionsConvertRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/convert"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type WorkbookFunctionsConvertRequest struct{ BaseRequest }

//
func (b *WorkbookFunctionsConvertRequestBuilder) Request() *WorkbookFunctionsConvertRequest {
	return &WorkbookFunctionsConvertRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject,
			tenantID: b.tenantID, applicationID: b.applicationID, clientSecurityKey: b.clientSecurityKey},
	}
}

//
func (r *WorkbookFunctionsConvertRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}
