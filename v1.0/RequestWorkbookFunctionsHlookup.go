// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

//
type WorkbookFunctionsHlookupRequestBuilder struct{ BaseRequestBuilder }

// Hlookup action undocumented
func (b *WorkbookFunctionsRequestBuilder) Hlookup(reqObj *WorkbookFunctionsHlookupRequestParameter) *WorkbookFunctionsHlookupRequestBuilder {
	bb := &WorkbookFunctionsHlookupRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/hlookup"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type WorkbookFunctionsHlookupRequest struct{ BaseRequest }

//
func (b *WorkbookFunctionsHlookupRequestBuilder) Request() *WorkbookFunctionsHlookupRequest {
	return &WorkbookFunctionsHlookupRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject,
			tenantID: b.tenantID, applicationID: b.applicationID, clientSecurityKey: b.clientSecurityKey},
	}
}

//
func (r *WorkbookFunctionsHlookupRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}
