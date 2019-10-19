// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// UnifiedRoleDefinitionRequestBuilder is request builder for UnifiedRoleDefinition
type UnifiedRoleDefinitionRequestBuilder struct{ BaseRequestBuilder }

// Request returns UnifiedRoleDefinitionRequest
func (b *UnifiedRoleDefinitionRequestBuilder) Request() *UnifiedRoleDefinitionRequest {
	return &UnifiedRoleDefinitionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// UnifiedRoleDefinitionRequest is request for UnifiedRoleDefinition
type UnifiedRoleDefinitionRequest struct{ BaseRequest }

// Do performs HTTP request for UnifiedRoleDefinition
func (r *UnifiedRoleDefinitionRequest) Do(method, path string, reqObj interface{}) (resObj *UnifiedRoleDefinition, err error) {
	err = r.JSONRequestWithPath(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for UnifiedRoleDefinition
func (r *UnifiedRoleDefinitionRequest) Get() (*UnifiedRoleDefinition, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for UnifiedRoleDefinition
func (r *UnifiedRoleDefinitionRequest) Update(reqObj *UnifiedRoleDefinition) (*UnifiedRoleDefinition, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for UnifiedRoleDefinition
func (r *UnifiedRoleDefinitionRequest) Delete() error {
	return r.JSONRequestWithPath("DELETE", "", nil, nil)
}