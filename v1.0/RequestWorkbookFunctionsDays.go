// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

//
type WorkbookFunctionsDaysRequestBuilder struct{ BaseRequestBuilder }

// Days action undocumented
func (b *WorkbookFunctionsRequestBuilder) Days(reqObj *WorkbookFunctionsDaysRequestParameter) *WorkbookFunctionsDaysRequestBuilder {
	bb := &WorkbookFunctionsDaysRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/days"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type WorkbookFunctionsDaysRequest struct{ BaseRequest }

//
func (b *WorkbookFunctionsDaysRequestBuilder) Request() *WorkbookFunctionsDaysRequest {
	return &WorkbookFunctionsDaysRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject,
			tenantID: b.tenantID, applicationID: b.applicationID, clientSecurityKey: b.clientSecurityKey},
	}
}

//
func (r *WorkbookFunctionsDaysRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}
