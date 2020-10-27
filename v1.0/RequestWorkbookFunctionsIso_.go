// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

//
type WorkbookFunctionsIso_CeilingRequestBuilder struct{ BaseRequestBuilder }

// Iso_Ceiling action undocumented
func (b *WorkbookFunctionsRequestBuilder) Iso_Ceiling(reqObj *WorkbookFunctionsIso_CeilingRequestParameter) *WorkbookFunctionsIso_CeilingRequestBuilder {
	bb := &WorkbookFunctionsIso_CeilingRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/iso_Ceiling"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type WorkbookFunctionsIso_CeilingRequest struct{ BaseRequest }

//
func (b *WorkbookFunctionsIso_CeilingRequestBuilder) Request() *WorkbookFunctionsIso_CeilingRequest {
	return &WorkbookFunctionsIso_CeilingRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject,
			tenantID: b.tenantID, applicationID: b.applicationID, clientSecurityKey: b.clientSecurityKey},
	}
}

//
func (r *WorkbookFunctionsIso_CeilingRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}
