// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// EducationalActivityRequestBuilder is request builder for EducationalActivity
type EducationalActivityRequestBuilder struct{ BaseRequestBuilder }

// Request returns EducationalActivityRequest
func (b *EducationalActivityRequestBuilder) Request() *EducationalActivityRequest {
	return &EducationalActivityRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client,
			tenantID: b.tenantID, applicationID: b.applicationID, clientSecurityKey: b.clientSecurityKey},
	}
}

// EducationalActivityRequest is request for EducationalActivity
type EducationalActivityRequest struct{ BaseRequest }

// Get performs GET request for EducationalActivity
func (r *EducationalActivityRequest) Get(ctx context.Context) (resObj *EducationalActivity, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for EducationalActivity
func (r *EducationalActivityRequest) Update(ctx context.Context, reqObj *EducationalActivity) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for EducationalActivity
func (r *EducationalActivityRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
