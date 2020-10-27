// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

//
type WorkbookFunctionsBinom_DistRequestBuilder struct{ BaseRequestBuilder }

// Binom_Dist action undocumented
func (b *WorkbookFunctionsRequestBuilder) Binom_Dist(reqObj *WorkbookFunctionsBinom_DistRequestParameter) *WorkbookFunctionsBinom_DistRequestBuilder {
	bb := &WorkbookFunctionsBinom_DistRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/binom_Dist"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type WorkbookFunctionsBinom_DistRequest struct{ BaseRequest }

//
func (b *WorkbookFunctionsBinom_DistRequestBuilder) Request() *WorkbookFunctionsBinom_DistRequest {
	return &WorkbookFunctionsBinom_DistRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject,
			tenantID: b.tenantID, applicationID: b.applicationID, clientSecurityKey: b.clientSecurityKey},
	}
}

//
func (r *WorkbookFunctionsBinom_DistRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

//
type WorkbookFunctionsBinom_Dist_RangeRequestBuilder struct{ BaseRequestBuilder }

// Binom_Dist_Range action undocumented
func (b *WorkbookFunctionsRequestBuilder) Binom_Dist_Range(reqObj *WorkbookFunctionsBinom_Dist_RangeRequestParameter) *WorkbookFunctionsBinom_Dist_RangeRequestBuilder {
	bb := &WorkbookFunctionsBinom_Dist_RangeRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/binom_Dist_Range"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type WorkbookFunctionsBinom_Dist_RangeRequest struct{ BaseRequest }

//
func (b *WorkbookFunctionsBinom_Dist_RangeRequestBuilder) Request() *WorkbookFunctionsBinom_Dist_RangeRequest {
	return &WorkbookFunctionsBinom_Dist_RangeRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject,
			tenantID: b.tenantID, applicationID: b.applicationID, clientSecurityKey: b.clientSecurityKey},
	}
}

//
func (r *WorkbookFunctionsBinom_Dist_RangeRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

//
type WorkbookFunctionsBinom_InvRequestBuilder struct{ BaseRequestBuilder }

// Binom_Inv action undocumented
func (b *WorkbookFunctionsRequestBuilder) Binom_Inv(reqObj *WorkbookFunctionsBinom_InvRequestParameter) *WorkbookFunctionsBinom_InvRequestBuilder {
	bb := &WorkbookFunctionsBinom_InvRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/binom_Inv"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type WorkbookFunctionsBinom_InvRequest struct{ BaseRequest }

//
func (b *WorkbookFunctionsBinom_InvRequestBuilder) Request() *WorkbookFunctionsBinom_InvRequest {
	return &WorkbookFunctionsBinom_InvRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject,
			tenantID: b.tenantID, applicationID: b.applicationID, clientSecurityKey: b.clientSecurityKey},
	}
}

//
func (r *WorkbookFunctionsBinom_InvRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}
