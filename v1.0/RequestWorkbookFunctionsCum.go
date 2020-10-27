// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

//
type WorkbookFunctionsCumIPmtRequestBuilder struct{ BaseRequestBuilder }

// CumIPmt action undocumented
func (b *WorkbookFunctionsRequestBuilder) CumIPmt(reqObj *WorkbookFunctionsCumIPmtRequestParameter) *WorkbookFunctionsCumIPmtRequestBuilder {
	bb := &WorkbookFunctionsCumIPmtRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/cumIPmt"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type WorkbookFunctionsCumIPmtRequest struct{ BaseRequest }

//
func (b *WorkbookFunctionsCumIPmtRequestBuilder) Request() *WorkbookFunctionsCumIPmtRequest {
	return &WorkbookFunctionsCumIPmtRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject,
			tenantID: b.tenantID, applicationID: b.applicationID, clientSecurityKey: b.clientSecurityKey},
	}
}

//
func (r *WorkbookFunctionsCumIPmtRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

//
type WorkbookFunctionsCumPrincRequestBuilder struct{ BaseRequestBuilder }

// CumPrinc action undocumented
func (b *WorkbookFunctionsRequestBuilder) CumPrinc(reqObj *WorkbookFunctionsCumPrincRequestParameter) *WorkbookFunctionsCumPrincRequestBuilder {
	bb := &WorkbookFunctionsCumPrincRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/cumPrinc"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type WorkbookFunctionsCumPrincRequest struct{ BaseRequest }

//
func (b *WorkbookFunctionsCumPrincRequestBuilder) Request() *WorkbookFunctionsCumPrincRequest {
	return &WorkbookFunctionsCumPrincRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject,
			tenantID: b.tenantID, applicationID: b.applicationID, clientSecurityKey: b.clientSecurityKey},
	}
}

//
func (r *WorkbookFunctionsCumPrincRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}
