// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// AdvancedThreatProtectionOnboardingDeviceSettingStateRequestBuilder is request builder for AdvancedThreatProtectionOnboardingDeviceSettingState
type AdvancedThreatProtectionOnboardingDeviceSettingStateRequestBuilder struct{ BaseRequestBuilder }

// Request returns AdvancedThreatProtectionOnboardingDeviceSettingStateRequest
func (b *AdvancedThreatProtectionOnboardingDeviceSettingStateRequestBuilder) Request() *AdvancedThreatProtectionOnboardingDeviceSettingStateRequest {
	return &AdvancedThreatProtectionOnboardingDeviceSettingStateRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client,
			tenantID: b.tenantID, applicationID: b.applicationID, clientSecurityKey: b.clientSecurityKey},
	}
}

// AdvancedThreatProtectionOnboardingDeviceSettingStateRequest is request for AdvancedThreatProtectionOnboardingDeviceSettingState
type AdvancedThreatProtectionOnboardingDeviceSettingStateRequest struct{ BaseRequest }

// Get performs GET request for AdvancedThreatProtectionOnboardingDeviceSettingState
func (r *AdvancedThreatProtectionOnboardingDeviceSettingStateRequest) Get(ctx context.Context) (resObj *AdvancedThreatProtectionOnboardingDeviceSettingState, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for AdvancedThreatProtectionOnboardingDeviceSettingState
func (r *AdvancedThreatProtectionOnboardingDeviceSettingStateRequest) Update(ctx context.Context, reqObj *AdvancedThreatProtectionOnboardingDeviceSettingState) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for AdvancedThreatProtectionOnboardingDeviceSettingState
func (r *AdvancedThreatProtectionOnboardingDeviceSettingStateRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// AdvancedThreatProtectionOnboardingStateSummaryRequestBuilder is request builder for AdvancedThreatProtectionOnboardingStateSummary
type AdvancedThreatProtectionOnboardingStateSummaryRequestBuilder struct{ BaseRequestBuilder }

// Request returns AdvancedThreatProtectionOnboardingStateSummaryRequest
func (b *AdvancedThreatProtectionOnboardingStateSummaryRequestBuilder) Request() *AdvancedThreatProtectionOnboardingStateSummaryRequest {
	return &AdvancedThreatProtectionOnboardingStateSummaryRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client,
			tenantID: b.tenantID, applicationID: b.applicationID, clientSecurityKey: b.clientSecurityKey},
	}
}

// AdvancedThreatProtectionOnboardingStateSummaryRequest is request for AdvancedThreatProtectionOnboardingStateSummary
type AdvancedThreatProtectionOnboardingStateSummaryRequest struct{ BaseRequest }

// Get performs GET request for AdvancedThreatProtectionOnboardingStateSummary
func (r *AdvancedThreatProtectionOnboardingStateSummaryRequest) Get(ctx context.Context) (resObj *AdvancedThreatProtectionOnboardingStateSummary, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for AdvancedThreatProtectionOnboardingStateSummary
func (r *AdvancedThreatProtectionOnboardingStateSummaryRequest) Update(ctx context.Context, reqObj *AdvancedThreatProtectionOnboardingStateSummary) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for AdvancedThreatProtectionOnboardingStateSummary
func (r *AdvancedThreatProtectionOnboardingStateSummaryRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
