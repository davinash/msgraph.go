// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// CallRequestBuilder is request builder for Call
type CallRequestBuilder struct{ BaseRequestBuilder }

// Request returns CallRequest
func (b *CallRequestBuilder) Request() *CallRequest {
	return &CallRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client,
			tenantID: b.tenantID, applicationID: b.applicationID, clientSecurityKey: b.clientSecurityKey},
	}
}

// CallRequest is request for Call
type CallRequest struct{ BaseRequest }

// Get performs GET request for Call
func (r *CallRequest) Get(ctx context.Context) (resObj *Call, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for Call
func (r *CallRequest) Update(ctx context.Context, reqObj *Call) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for Call
func (r *CallRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

//
type CallAnswerRequestBuilder struct{ BaseRequestBuilder }

// Answer action undocumented
func (b *CallRequestBuilder) Answer(reqObj *CallAnswerRequestParameter) *CallAnswerRequestBuilder {
	bb := &CallAnswerRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/answer"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type CallAnswerRequest struct{ BaseRequest }

//
func (b *CallAnswerRequestBuilder) Request() *CallAnswerRequest {
	return &CallAnswerRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject,
			tenantID: b.tenantID, applicationID: b.applicationID, clientSecurityKey: b.clientSecurityKey},
	}
}

//
func (r *CallAnswerRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

//
type CallCancelMediaProcessingRequestBuilder struct{ BaseRequestBuilder }

// CancelMediaProcessing action undocumented
func (b *CallRequestBuilder) CancelMediaProcessing(reqObj *CallCancelMediaProcessingRequestParameter) *CallCancelMediaProcessingRequestBuilder {
	bb := &CallCancelMediaProcessingRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/cancelMediaProcessing"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type CallCancelMediaProcessingRequest struct{ BaseRequest }

//
func (b *CallCancelMediaProcessingRequestBuilder) Request() *CallCancelMediaProcessingRequest {
	return &CallCancelMediaProcessingRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject,
			tenantID: b.tenantID, applicationID: b.applicationID, clientSecurityKey: b.clientSecurityKey},
	}
}

//
func (r *CallCancelMediaProcessingRequest) Post(ctx context.Context) (resObj *CancelMediaProcessingOperation, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

//
type CallChangeScreenSharingRoleRequestBuilder struct{ BaseRequestBuilder }

// ChangeScreenSharingRole action undocumented
func (b *CallRequestBuilder) ChangeScreenSharingRole(reqObj *CallChangeScreenSharingRoleRequestParameter) *CallChangeScreenSharingRoleRequestBuilder {
	bb := &CallChangeScreenSharingRoleRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/changeScreenSharingRole"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type CallChangeScreenSharingRoleRequest struct{ BaseRequest }

//
func (b *CallChangeScreenSharingRoleRequestBuilder) Request() *CallChangeScreenSharingRoleRequest {
	return &CallChangeScreenSharingRoleRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject,
			tenantID: b.tenantID, applicationID: b.applicationID, clientSecurityKey: b.clientSecurityKey},
	}
}

//
func (r *CallChangeScreenSharingRoleRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

//
type CallKeepAliveRequestBuilder struct{ BaseRequestBuilder }

// KeepAlive action undocumented
func (b *CallRequestBuilder) KeepAlive(reqObj *CallKeepAliveRequestParameter) *CallKeepAliveRequestBuilder {
	bb := &CallKeepAliveRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/keepAlive"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type CallKeepAliveRequest struct{ BaseRequest }

//
func (b *CallKeepAliveRequestBuilder) Request() *CallKeepAliveRequest {
	return &CallKeepAliveRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject,
			tenantID: b.tenantID, applicationID: b.applicationID, clientSecurityKey: b.clientSecurityKey},
	}
}

//
func (r *CallKeepAliveRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

//
type CallMuteRequestBuilder struct{ BaseRequestBuilder }

// Mute action undocumented
func (b *CallRequestBuilder) Mute(reqObj *CallMuteRequestParameter) *CallMuteRequestBuilder {
	bb := &CallMuteRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/mute"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type CallMuteRequest struct{ BaseRequest }

//
func (b *CallMuteRequestBuilder) Request() *CallMuteRequest {
	return &CallMuteRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject,
			tenantID: b.tenantID, applicationID: b.applicationID, clientSecurityKey: b.clientSecurityKey},
	}
}

//
func (r *CallMuteRequest) Post(ctx context.Context) (resObj *MuteParticipantOperation, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

//
type CallPlayPromptRequestBuilder struct{ BaseRequestBuilder }

// PlayPrompt action undocumented
func (b *CallRequestBuilder) PlayPrompt(reqObj *CallPlayPromptRequestParameter) *CallPlayPromptRequestBuilder {
	bb := &CallPlayPromptRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/playPrompt"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type CallPlayPromptRequest struct{ BaseRequest }

//
func (b *CallPlayPromptRequestBuilder) Request() *CallPlayPromptRequest {
	return &CallPlayPromptRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject,
			tenantID: b.tenantID, applicationID: b.applicationID, clientSecurityKey: b.clientSecurityKey},
	}
}

//
func (r *CallPlayPromptRequest) Post(ctx context.Context) (resObj *PlayPromptOperation, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

//
type CallRecordRequestBuilder struct{ BaseRequestBuilder }

// Record action undocumented
func (b *CallRequestBuilder) Record(reqObj *CallRecordRequestParameter) *CallRecordRequestBuilder {
	bb := &CallRecordRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/record"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type CallRecordRequest struct{ BaseRequest }

//
func (b *CallRecordRequestBuilder) Request() *CallRecordRequest {
	return &CallRecordRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject,
			tenantID: b.tenantID, applicationID: b.applicationID, clientSecurityKey: b.clientSecurityKey},
	}
}

//
func (r *CallRecordRequest) Post(ctx context.Context) (resObj *RecordOperation, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

//
type CallRecordResponseRequestBuilder struct{ BaseRequestBuilder }

// RecordResponse action undocumented
func (b *CallRequestBuilder) RecordResponse(reqObj *CallRecordResponseRequestParameter) *CallRecordResponseRequestBuilder {
	bb := &CallRecordResponseRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/recordResponse"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type CallRecordResponseRequest struct{ BaseRequest }

//
func (b *CallRecordResponseRequestBuilder) Request() *CallRecordResponseRequest {
	return &CallRecordResponseRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject,
			tenantID: b.tenantID, applicationID: b.applicationID, clientSecurityKey: b.clientSecurityKey},
	}
}

//
func (r *CallRecordResponseRequest) Post(ctx context.Context) (resObj *RecordOperation, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

//
type CallRedirectRequestBuilder struct{ BaseRequestBuilder }

// Redirect action undocumented
func (b *CallRequestBuilder) Redirect(reqObj *CallRedirectRequestParameter) *CallRedirectRequestBuilder {
	bb := &CallRedirectRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/redirect"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type CallRedirectRequest struct{ BaseRequest }

//
func (b *CallRedirectRequestBuilder) Request() *CallRedirectRequest {
	return &CallRedirectRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject,
			tenantID: b.tenantID, applicationID: b.applicationID, clientSecurityKey: b.clientSecurityKey},
	}
}

//
func (r *CallRedirectRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

//
type CallRejectRequestBuilder struct{ BaseRequestBuilder }

// Reject action undocumented
func (b *CallRequestBuilder) Reject(reqObj *CallRejectRequestParameter) *CallRejectRequestBuilder {
	bb := &CallRejectRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/reject"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type CallRejectRequest struct{ BaseRequest }

//
func (b *CallRejectRequestBuilder) Request() *CallRejectRequest {
	return &CallRejectRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject,
			tenantID: b.tenantID, applicationID: b.applicationID, clientSecurityKey: b.clientSecurityKey},
	}
}

//
func (r *CallRejectRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

//
type CallSubscribeToToneRequestBuilder struct{ BaseRequestBuilder }

// SubscribeToTone action undocumented
func (b *CallRequestBuilder) SubscribeToTone(reqObj *CallSubscribeToToneRequestParameter) *CallSubscribeToToneRequestBuilder {
	bb := &CallSubscribeToToneRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/subscribeToTone"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type CallSubscribeToToneRequest struct{ BaseRequest }

//
func (b *CallSubscribeToToneRequestBuilder) Request() *CallSubscribeToToneRequest {
	return &CallSubscribeToToneRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject,
			tenantID: b.tenantID, applicationID: b.applicationID, clientSecurityKey: b.clientSecurityKey},
	}
}

//
func (r *CallSubscribeToToneRequest) Post(ctx context.Context) (resObj *SubscribeToToneOperation, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

//
type CallTransferRequestBuilder struct{ BaseRequestBuilder }

// Transfer action undocumented
func (b *CallRequestBuilder) Transfer(reqObj *CallTransferRequestParameter) *CallTransferRequestBuilder {
	bb := &CallTransferRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/transfer"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type CallTransferRequest struct{ BaseRequest }

//
func (b *CallTransferRequestBuilder) Request() *CallTransferRequest {
	return &CallTransferRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject,
			tenantID: b.tenantID, applicationID: b.applicationID, clientSecurityKey: b.clientSecurityKey},
	}
}

//
func (r *CallTransferRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

//
type CallUnmuteRequestBuilder struct{ BaseRequestBuilder }

// Unmute action undocumented
func (b *CallRequestBuilder) Unmute(reqObj *CallUnmuteRequestParameter) *CallUnmuteRequestBuilder {
	bb := &CallUnmuteRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/unmute"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type CallUnmuteRequest struct{ BaseRequest }

//
func (b *CallUnmuteRequestBuilder) Request() *CallUnmuteRequest {
	return &CallUnmuteRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject,
			tenantID: b.tenantID, applicationID: b.applicationID, clientSecurityKey: b.clientSecurityKey},
	}
}

//
func (r *CallUnmuteRequest) Post(ctx context.Context) (resObj *UnmuteParticipantOperation, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

//
type CallUpdateRecordingStatusRequestBuilder struct{ BaseRequestBuilder }

// UpdateRecordingStatus action undocumented
func (b *CallRequestBuilder) UpdateRecordingStatus(reqObj *CallUpdateRecordingStatusRequestParameter) *CallUpdateRecordingStatusRequestBuilder {
	bb := &CallUpdateRecordingStatusRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/updateRecordingStatus"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type CallUpdateRecordingStatusRequest struct{ BaseRequest }

//
func (b *CallUpdateRecordingStatusRequestBuilder) Request() *CallUpdateRecordingStatusRequest {
	return &CallUpdateRecordingStatusRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject,
			tenantID: b.tenantID, applicationID: b.applicationID, clientSecurityKey: b.clientSecurityKey},
	}
}

//
func (r *CallUpdateRecordingStatusRequest) Post(ctx context.Context) (resObj *UpdateRecordingStatusOperation, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}
