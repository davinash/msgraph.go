// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

//
type WorkbookFunctionsOct2BinRequestBuilder struct{ BaseRequestBuilder }

// Oct2Bin action undocumented
func (b *WorkbookFunctionsRequestBuilder) Oct2Bin(reqObj *WorkbookFunctionsOct2BinRequestParameter) *WorkbookFunctionsOct2BinRequestBuilder {
	bb := &WorkbookFunctionsOct2BinRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/oct2Bin"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type WorkbookFunctionsOct2BinRequest struct{ BaseRequest }

//
func (b *WorkbookFunctionsOct2BinRequestBuilder) Request() *WorkbookFunctionsOct2BinRequest {
	return &WorkbookFunctionsOct2BinRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject,
			tenantID: b.tenantID, applicationID: b.applicationID, clientSecurityKey: b.clientSecurityKey},
	}
}

//
func (r *WorkbookFunctionsOct2BinRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

//
type WorkbookFunctionsOct2DecRequestBuilder struct{ BaseRequestBuilder }

// Oct2Dec action undocumented
func (b *WorkbookFunctionsRequestBuilder) Oct2Dec(reqObj *WorkbookFunctionsOct2DecRequestParameter) *WorkbookFunctionsOct2DecRequestBuilder {
	bb := &WorkbookFunctionsOct2DecRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/oct2Dec"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type WorkbookFunctionsOct2DecRequest struct{ BaseRequest }

//
func (b *WorkbookFunctionsOct2DecRequestBuilder) Request() *WorkbookFunctionsOct2DecRequest {
	return &WorkbookFunctionsOct2DecRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject,
			tenantID: b.tenantID, applicationID: b.applicationID, clientSecurityKey: b.clientSecurityKey},
	}
}

//
func (r *WorkbookFunctionsOct2DecRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

//
type WorkbookFunctionsOct2HexRequestBuilder struct{ BaseRequestBuilder }

// Oct2Hex action undocumented
func (b *WorkbookFunctionsRequestBuilder) Oct2Hex(reqObj *WorkbookFunctionsOct2HexRequestParameter) *WorkbookFunctionsOct2HexRequestBuilder {
	bb := &WorkbookFunctionsOct2HexRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/oct2Hex"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type WorkbookFunctionsOct2HexRequest struct{ BaseRequest }

//
func (b *WorkbookFunctionsOct2HexRequestBuilder) Request() *WorkbookFunctionsOct2HexRequest {
	return &WorkbookFunctionsOct2HexRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject,
			tenantID: b.tenantID, applicationID: b.applicationID, clientSecurityKey: b.clientSecurityKey},
	}
}

//
func (r *WorkbookFunctionsOct2HexRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}
