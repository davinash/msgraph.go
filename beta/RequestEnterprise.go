// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// EnterpriseCodeSigningCertificateRequestBuilder is request builder for EnterpriseCodeSigningCertificate
type EnterpriseCodeSigningCertificateRequestBuilder struct{ BaseRequestBuilder }

// Request returns EnterpriseCodeSigningCertificateRequest
func (b *EnterpriseCodeSigningCertificateRequestBuilder) Request() *EnterpriseCodeSigningCertificateRequest {
	return &EnterpriseCodeSigningCertificateRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client,
			tenantID: b.tenantID, applicationID: b.applicationID, clientSecurityKey: b.clientSecurityKey},
	}
}

// EnterpriseCodeSigningCertificateRequest is request for EnterpriseCodeSigningCertificate
type EnterpriseCodeSigningCertificateRequest struct{ BaseRequest }

// Get performs GET request for EnterpriseCodeSigningCertificate
func (r *EnterpriseCodeSigningCertificateRequest) Get(ctx context.Context) (resObj *EnterpriseCodeSigningCertificate, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for EnterpriseCodeSigningCertificate
func (r *EnterpriseCodeSigningCertificateRequest) Update(ctx context.Context, reqObj *EnterpriseCodeSigningCertificate) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for EnterpriseCodeSigningCertificate
func (r *EnterpriseCodeSigningCertificateRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
