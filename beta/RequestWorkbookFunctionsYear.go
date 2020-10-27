// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

//
type WorkbookFunctionsYearRequestBuilder struct{ BaseRequestBuilder }

// Year action undocumented
func (b *WorkbookFunctionsRequestBuilder) Year(reqObj *WorkbookFunctionsYearRequestParameter) *WorkbookFunctionsYearRequestBuilder {
	bb := &WorkbookFunctionsYearRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/year"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type WorkbookFunctionsYearRequest struct{ BaseRequest }

//
func (b *WorkbookFunctionsYearRequestBuilder) Request() *WorkbookFunctionsYearRequest {
	return &WorkbookFunctionsYearRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject,
			tenantID: b.tenantID, applicationID: b.applicationID, clientSecurityKey: b.clientSecurityKey},
	}
}

//
func (r *WorkbookFunctionsYearRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

//
type WorkbookFunctionsYearFracRequestBuilder struct{ BaseRequestBuilder }

// YearFrac action undocumented
func (b *WorkbookFunctionsRequestBuilder) YearFrac(reqObj *WorkbookFunctionsYearFracRequestParameter) *WorkbookFunctionsYearFracRequestBuilder {
	bb := &WorkbookFunctionsYearFracRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/yearFrac"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type WorkbookFunctionsYearFracRequest struct{ BaseRequest }

//
func (b *WorkbookFunctionsYearFracRequestBuilder) Request() *WorkbookFunctionsYearFracRequest {
	return &WorkbookFunctionsYearFracRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject,
			tenantID: b.tenantID, applicationID: b.applicationID, clientSecurityKey: b.clientSecurityKey},
	}
}

//
func (r *WorkbookFunctionsYearFracRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}
