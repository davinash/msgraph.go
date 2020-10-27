// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// EnrollmentConfigurationAssignmentRequestBuilder is request builder for EnrollmentConfigurationAssignment
type EnrollmentConfigurationAssignmentRequestBuilder struct{ BaseRequestBuilder }

// Request returns EnrollmentConfigurationAssignmentRequest
func (b *EnrollmentConfigurationAssignmentRequestBuilder) Request() *EnrollmentConfigurationAssignmentRequest {
	return &EnrollmentConfigurationAssignmentRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client,
			tenantID: b.tenantID, applicationID: b.applicationID, clientSecurityKey: b.clientSecurityKey},
	}
}

// EnrollmentConfigurationAssignmentRequest is request for EnrollmentConfigurationAssignment
type EnrollmentConfigurationAssignmentRequest struct{ BaseRequest }

// Get performs GET request for EnrollmentConfigurationAssignment
func (r *EnrollmentConfigurationAssignmentRequest) Get(ctx context.Context) (resObj *EnrollmentConfigurationAssignment, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for EnrollmentConfigurationAssignment
func (r *EnrollmentConfigurationAssignmentRequest) Update(ctx context.Context, reqObj *EnrollmentConfigurationAssignment) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for EnrollmentConfigurationAssignment
func (r *EnrollmentConfigurationAssignmentRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// EnrollmentProfileRequestBuilder is request builder for EnrollmentProfile
type EnrollmentProfileRequestBuilder struct{ BaseRequestBuilder }

// Request returns EnrollmentProfileRequest
func (b *EnrollmentProfileRequestBuilder) Request() *EnrollmentProfileRequest {
	return &EnrollmentProfileRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client,
			tenantID: b.tenantID, applicationID: b.applicationID, clientSecurityKey: b.clientSecurityKey},
	}
}

// EnrollmentProfileRequest is request for EnrollmentProfile
type EnrollmentProfileRequest struct{ BaseRequest }

// Get performs GET request for EnrollmentProfile
func (r *EnrollmentProfileRequest) Get(ctx context.Context) (resObj *EnrollmentProfile, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for EnrollmentProfile
func (r *EnrollmentProfileRequest) Update(ctx context.Context, reqObj *EnrollmentProfile) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for EnrollmentProfile
func (r *EnrollmentProfileRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

//
type EnrollmentProfileSetDefaultProfileRequestBuilder struct{ BaseRequestBuilder }

// SetDefaultProfile action undocumented
func (b *EnrollmentProfileRequestBuilder) SetDefaultProfile(reqObj *EnrollmentProfileSetDefaultProfileRequestParameter) *EnrollmentProfileSetDefaultProfileRequestBuilder {
	bb := &EnrollmentProfileSetDefaultProfileRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/setDefaultProfile"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type EnrollmentProfileSetDefaultProfileRequest struct{ BaseRequest }

//
func (b *EnrollmentProfileSetDefaultProfileRequestBuilder) Request() *EnrollmentProfileSetDefaultProfileRequest {
	return &EnrollmentProfileSetDefaultProfileRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject,
			tenantID: b.tenantID, applicationID: b.applicationID, clientSecurityKey: b.clientSecurityKey},
	}
}

//
func (r *EnrollmentProfileSetDefaultProfileRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

//
type EnrollmentProfileUpdateDeviceProfileAssignmentRequestBuilder struct{ BaseRequestBuilder }

// UpdateDeviceProfileAssignment action undocumented
func (b *EnrollmentProfileRequestBuilder) UpdateDeviceProfileAssignment(reqObj *EnrollmentProfileUpdateDeviceProfileAssignmentRequestParameter) *EnrollmentProfileUpdateDeviceProfileAssignmentRequestBuilder {
	bb := &EnrollmentProfileUpdateDeviceProfileAssignmentRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/updateDeviceProfileAssignment"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type EnrollmentProfileUpdateDeviceProfileAssignmentRequest struct{ BaseRequest }

//
func (b *EnrollmentProfileUpdateDeviceProfileAssignmentRequestBuilder) Request() *EnrollmentProfileUpdateDeviceProfileAssignmentRequest {
	return &EnrollmentProfileUpdateDeviceProfileAssignmentRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject,
			tenantID: b.tenantID, applicationID: b.applicationID, clientSecurityKey: b.clientSecurityKey},
	}
}

//
func (r *EnrollmentProfileUpdateDeviceProfileAssignmentRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}
