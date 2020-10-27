// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

//
type WorkbookFunctionsAmorDegrcRequestBuilder struct{ BaseRequestBuilder }

// AmorDegrc action undocumented
func (b *WorkbookFunctionsRequestBuilder) AmorDegrc(reqObj *WorkbookFunctionsAmorDegrcRequestParameter) *WorkbookFunctionsAmorDegrcRequestBuilder {
	bb := &WorkbookFunctionsAmorDegrcRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/amorDegrc"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type WorkbookFunctionsAmorDegrcRequest struct{ BaseRequest }

//
func (b *WorkbookFunctionsAmorDegrcRequestBuilder) Request() *WorkbookFunctionsAmorDegrcRequest {
	return &WorkbookFunctionsAmorDegrcRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject,
			tenantID: b.tenantID, applicationID: b.applicationID, clientSecurityKey: b.clientSecurityKey},
	}
}

//
func (r *WorkbookFunctionsAmorDegrcRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

//
type WorkbookFunctionsAmorLincRequestBuilder struct{ BaseRequestBuilder }

// AmorLinc action undocumented
func (b *WorkbookFunctionsRequestBuilder) AmorLinc(reqObj *WorkbookFunctionsAmorLincRequestParameter) *WorkbookFunctionsAmorLincRequestBuilder {
	bb := &WorkbookFunctionsAmorLincRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/amorLinc"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type WorkbookFunctionsAmorLincRequest struct{ BaseRequest }

//
func (b *WorkbookFunctionsAmorLincRequestBuilder) Request() *WorkbookFunctionsAmorLincRequest {
	return &WorkbookFunctionsAmorLincRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject,
			tenantID: b.tenantID, applicationID: b.applicationID, clientSecurityKey: b.clientSecurityKey},
	}
}

//
func (r *WorkbookFunctionsAmorLincRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}
