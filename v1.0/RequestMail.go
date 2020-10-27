// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// MailFolderRequestBuilder is request builder for MailFolder
type MailFolderRequestBuilder struct{ BaseRequestBuilder }

// Request returns MailFolderRequest
func (b *MailFolderRequestBuilder) Request() *MailFolderRequest {
	return &MailFolderRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client,
			tenantID: b.tenantID, applicationID: b.applicationID, clientSecurityKey: b.clientSecurityKey},
	}
}

// MailFolderRequest is request for MailFolder
type MailFolderRequest struct{ BaseRequest }

// Get performs GET request for MailFolder
func (r *MailFolderRequest) Get(ctx context.Context) (resObj *MailFolder, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for MailFolder
func (r *MailFolderRequest) Update(ctx context.Context, reqObj *MailFolder) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for MailFolder
func (r *MailFolderRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

//
type MailFolderCopyRequestBuilder struct{ BaseRequestBuilder }

// Copy action undocumented
func (b *MailFolderRequestBuilder) Copy(reqObj *MailFolderCopyRequestParameter) *MailFolderCopyRequestBuilder {
	bb := &MailFolderCopyRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/copy"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type MailFolderCopyRequest struct{ BaseRequest }

//
func (b *MailFolderCopyRequestBuilder) Request() *MailFolderCopyRequest {
	return &MailFolderCopyRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject,
			tenantID: b.tenantID, applicationID: b.applicationID, clientSecurityKey: b.clientSecurityKey},
	}
}

//
func (r *MailFolderCopyRequest) Post(ctx context.Context) (resObj *MailFolder, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

//
type MailFolderMoveRequestBuilder struct{ BaseRequestBuilder }

// Move action undocumented
func (b *MailFolderRequestBuilder) Move(reqObj *MailFolderMoveRequestParameter) *MailFolderMoveRequestBuilder {
	bb := &MailFolderMoveRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/move"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type MailFolderMoveRequest struct{ BaseRequest }

//
func (b *MailFolderMoveRequestBuilder) Request() *MailFolderMoveRequest {
	return &MailFolderMoveRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject,
			tenantID: b.tenantID, applicationID: b.applicationID, clientSecurityKey: b.clientSecurityKey},
	}
}

//
func (r *MailFolderMoveRequest) Post(ctx context.Context) (resObj *MailFolder, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}
