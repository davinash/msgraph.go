// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// ActiveDirectoryWindowsAutopilotDeploymentProfileRequestBuilder is request builder for ActiveDirectoryWindowsAutopilotDeploymentProfile
type ActiveDirectoryWindowsAutopilotDeploymentProfileRequestBuilder struct{ BaseRequestBuilder }

// Request returns ActiveDirectoryWindowsAutopilotDeploymentProfileRequest
func (b *ActiveDirectoryWindowsAutopilotDeploymentProfileRequestBuilder) Request() *ActiveDirectoryWindowsAutopilotDeploymentProfileRequest {
	return &ActiveDirectoryWindowsAutopilotDeploymentProfileRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client,
			tenantID: b.tenantID, applicationID: b.applicationID, clientSecurityKey: b.clientSecurityKey},
	}
}

// ActiveDirectoryWindowsAutopilotDeploymentProfileRequest is request for ActiveDirectoryWindowsAutopilotDeploymentProfile
type ActiveDirectoryWindowsAutopilotDeploymentProfileRequest struct{ BaseRequest }

// Get performs GET request for ActiveDirectoryWindowsAutopilotDeploymentProfile
func (r *ActiveDirectoryWindowsAutopilotDeploymentProfileRequest) Get(ctx context.Context) (resObj *ActiveDirectoryWindowsAutopilotDeploymentProfile, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ActiveDirectoryWindowsAutopilotDeploymentProfile
func (r *ActiveDirectoryWindowsAutopilotDeploymentProfileRequest) Update(ctx context.Context, reqObj *ActiveDirectoryWindowsAutopilotDeploymentProfile) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ActiveDirectoryWindowsAutopilotDeploymentProfile
func (r *ActiveDirectoryWindowsAutopilotDeploymentProfileRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
