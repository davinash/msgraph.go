// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// DeviceCategoryRequestBuilder is request builder for DeviceCategory
type DeviceCategoryRequestBuilder struct{ BaseRequestBuilder }

// Request returns DeviceCategoryRequest
func (b *DeviceCategoryRequestBuilder) Request() *DeviceCategoryRequest {
	return &DeviceCategoryRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// DeviceCategoryRequest is request for DeviceCategory
type DeviceCategoryRequest struct{ BaseRequest }

// Do performs HTTP request for DeviceCategory
func (r *DeviceCategoryRequest) Do(method, path string, reqObj interface{}) (resObj *DeviceCategory, err error) {
	err = r.JSONRequestWithPath(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for DeviceCategory
func (r *DeviceCategoryRequest) Get() (*DeviceCategory, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for DeviceCategory
func (r *DeviceCategoryRequest) Update(reqObj *DeviceCategory) (*DeviceCategory, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for DeviceCategory
func (r *DeviceCategoryRequest) Delete() error {
	return r.JSONRequestWithPath("DELETE", "", nil, nil)
}