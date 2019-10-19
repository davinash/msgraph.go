// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// WindowsUpdateForBusinessConfigurationExtendFeatureUpdatesPauseRequestParameter undocumented
type WindowsUpdateForBusinessConfigurationExtendFeatureUpdatesPauseRequestParameter struct {
}

// WindowsUpdateForBusinessConfigurationExtendQualityUpdatesPauseRequestParameter undocumented
type WindowsUpdateForBusinessConfigurationExtendQualityUpdatesPauseRequestParameter struct {
}

//
type WindowsUpdateForBusinessConfigurationExtendFeatureUpdatesPauseRequestBuilder struct{ BaseRequestBuilder }

// ExtendFeatureUpdatesPause action undocumented
func (b *WindowsUpdateForBusinessConfigurationRequestBuilder) ExtendFeatureUpdatesPause(reqObj *WindowsUpdateForBusinessConfigurationExtendFeatureUpdatesPauseRequestParameter) *WindowsUpdateForBusinessConfigurationExtendFeatureUpdatesPauseRequestBuilder {
	bb := &WindowsUpdateForBusinessConfigurationExtendFeatureUpdatesPauseRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/extendFeatureUpdatesPause"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type WindowsUpdateForBusinessConfigurationExtendFeatureUpdatesPauseRequest struct{ BaseRequest }

//
func (b *WindowsUpdateForBusinessConfigurationExtendFeatureUpdatesPauseRequestBuilder) Request() *WindowsUpdateForBusinessConfigurationExtendFeatureUpdatesPauseRequest {
	return &WindowsUpdateForBusinessConfigurationExtendFeatureUpdatesPauseRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *WindowsUpdateForBusinessConfigurationExtendFeatureUpdatesPauseRequest) Do(method, path string, reqObj interface{}) error {
	return r.JSONRequestWithPath(method, path, reqObj, nil)
}

//
func (r *WindowsUpdateForBusinessConfigurationExtendFeatureUpdatesPauseRequest) Post() error {
	return r.Do("POST", "", r.requestObject)
}

//
type WindowsUpdateForBusinessConfigurationExtendQualityUpdatesPauseRequestBuilder struct{ BaseRequestBuilder }

// ExtendQualityUpdatesPause action undocumented
func (b *WindowsUpdateForBusinessConfigurationRequestBuilder) ExtendQualityUpdatesPause(reqObj *WindowsUpdateForBusinessConfigurationExtendQualityUpdatesPauseRequestParameter) *WindowsUpdateForBusinessConfigurationExtendQualityUpdatesPauseRequestBuilder {
	bb := &WindowsUpdateForBusinessConfigurationExtendQualityUpdatesPauseRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/extendQualityUpdatesPause"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type WindowsUpdateForBusinessConfigurationExtendQualityUpdatesPauseRequest struct{ BaseRequest }

//
func (b *WindowsUpdateForBusinessConfigurationExtendQualityUpdatesPauseRequestBuilder) Request() *WindowsUpdateForBusinessConfigurationExtendQualityUpdatesPauseRequest {
	return &WindowsUpdateForBusinessConfigurationExtendQualityUpdatesPauseRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *WindowsUpdateForBusinessConfigurationExtendQualityUpdatesPauseRequest) Do(method, path string, reqObj interface{}) error {
	return r.JSONRequestWithPath(method, path, reqObj, nil)
}

//
func (r *WindowsUpdateForBusinessConfigurationExtendQualityUpdatesPauseRequest) Post() error {
	return r.Do("POST", "", r.requestObject)
}