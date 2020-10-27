// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// NotificationMessageTemplateRequestBuilder is request builder for NotificationMessageTemplate
type NotificationMessageTemplateRequestBuilder struct{ BaseRequestBuilder }

// Request returns NotificationMessageTemplateRequest
func (b *NotificationMessageTemplateRequestBuilder) Request() *NotificationMessageTemplateRequest {
	return &NotificationMessageTemplateRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client,
			tenantID: b.tenantID, applicationID: b.applicationID, clientSecurityKey: b.clientSecurityKey},
	}
}

// NotificationMessageTemplateRequest is request for NotificationMessageTemplate
type NotificationMessageTemplateRequest struct{ BaseRequest }

// Get performs GET request for NotificationMessageTemplate
func (r *NotificationMessageTemplateRequest) Get(ctx context.Context) (resObj *NotificationMessageTemplate, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for NotificationMessageTemplate
func (r *NotificationMessageTemplateRequest) Update(ctx context.Context, reqObj *NotificationMessageTemplate) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for NotificationMessageTemplate
func (r *NotificationMessageTemplateRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

//
type NotificationMessageTemplateSendTestMessageRequestBuilder struct{ BaseRequestBuilder }

// SendTestMessage action undocumented
func (b *NotificationMessageTemplateRequestBuilder) SendTestMessage(reqObj *NotificationMessageTemplateSendTestMessageRequestParameter) *NotificationMessageTemplateSendTestMessageRequestBuilder {
	bb := &NotificationMessageTemplateSendTestMessageRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/sendTestMessage"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type NotificationMessageTemplateSendTestMessageRequest struct{ BaseRequest }

//
func (b *NotificationMessageTemplateSendTestMessageRequestBuilder) Request() *NotificationMessageTemplateSendTestMessageRequest {
	return &NotificationMessageTemplateSendTestMessageRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject,
			tenantID: b.tenantID, applicationID: b.applicationID, clientSecurityKey: b.clientSecurityKey},
	}
}

//
func (r *NotificationMessageTemplateSendTestMessageRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}
