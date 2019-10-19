// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// DeviceManagementComplexSettingInstanceRequestBuilder is request builder for DeviceManagementComplexSettingInstance
type DeviceManagementComplexSettingInstanceRequestBuilder struct{ BaseRequestBuilder }

// Request returns DeviceManagementComplexSettingInstanceRequest
func (b *DeviceManagementComplexSettingInstanceRequestBuilder) Request() *DeviceManagementComplexSettingInstanceRequest {
	return &DeviceManagementComplexSettingInstanceRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// DeviceManagementComplexSettingInstanceRequest is request for DeviceManagementComplexSettingInstance
type DeviceManagementComplexSettingInstanceRequest struct{ BaseRequest }

// Do performs HTTP request for DeviceManagementComplexSettingInstance
func (r *DeviceManagementComplexSettingInstanceRequest) Do(method, path string, reqObj interface{}) (resObj *DeviceManagementComplexSettingInstance, err error) {
	err = r.JSONRequestWithPath(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for DeviceManagementComplexSettingInstance
func (r *DeviceManagementComplexSettingInstanceRequest) Get() (*DeviceManagementComplexSettingInstance, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for DeviceManagementComplexSettingInstance
func (r *DeviceManagementComplexSettingInstanceRequest) Update(reqObj *DeviceManagementComplexSettingInstance) (*DeviceManagementComplexSettingInstance, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for DeviceManagementComplexSettingInstance
func (r *DeviceManagementComplexSettingInstanceRequest) Delete() error {
	return r.JSONRequestWithPath("DELETE", "", nil, nil)
}

// Value returns request builder for DeviceManagementSettingInstance collection
func (b *DeviceManagementComplexSettingInstanceRequestBuilder) Value() *DeviceManagementComplexSettingInstanceValueCollectionRequestBuilder {
	bb := &DeviceManagementComplexSettingInstanceValueCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/value"
	return bb
}

// DeviceManagementComplexSettingInstanceValueCollectionRequestBuilder is request builder for DeviceManagementSettingInstance collection
type DeviceManagementComplexSettingInstanceValueCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for DeviceManagementSettingInstance collection
func (b *DeviceManagementComplexSettingInstanceValueCollectionRequestBuilder) Request() *DeviceManagementComplexSettingInstanceValueCollectionRequest {
	return &DeviceManagementComplexSettingInstanceValueCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for DeviceManagementSettingInstance item
func (b *DeviceManagementComplexSettingInstanceValueCollectionRequestBuilder) ID(id string) *DeviceManagementSettingInstanceRequestBuilder {
	bb := &DeviceManagementSettingInstanceRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// DeviceManagementComplexSettingInstanceValueCollectionRequest is request for DeviceManagementSettingInstance collection
type DeviceManagementComplexSettingInstanceValueCollectionRequest struct{ BaseRequest }

// Do performs HTTP request for DeviceManagementSettingInstance collection
func (r *DeviceManagementComplexSettingInstanceValueCollectionRequest) Do(method, path string, reqObj interface{}) (resObj *DeviceManagementSettingInstance, err error) {
	err = r.JSONRequestWithPath(method, path, reqObj, &resObj)
	return
}

// Paging perfoms paging operation for DeviceManagementSettingInstance collection
func (r *DeviceManagementComplexSettingInstanceValueCollectionRequest) Paging(method, path string, obj interface{}) ([]DeviceManagementSettingInstance, error) {
	req, err := r.NewJSONRequestWithPath(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []DeviceManagementSettingInstance
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  []DeviceManagementSettingInstance
		)
		err := json.NewDecoder(res.Body).Decode(&paging)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(paging.Value, &value)
		if err != nil {
			return nil, err
		}
		values = append(values, value...)
		if len(paging.NextLink) == 0 {
			return values, nil
		}
		res, err = r.client.Get(paging.NextLink)
		if err != nil {
			return nil, err
		}
	}
}

// Get performs GET request for DeviceManagementSettingInstance collection
func (r *DeviceManagementComplexSettingInstanceValueCollectionRequest) Get() ([]DeviceManagementSettingInstance, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for DeviceManagementSettingInstance collection
func (r *DeviceManagementComplexSettingInstanceValueCollectionRequest) Add(reqObj *DeviceManagementSettingInstance) (*DeviceManagementSettingInstance, error) {
	return r.Do("POST", "", reqObj)
}