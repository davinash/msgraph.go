// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/davinash/msgraph.go/jsonx"
)

// ManagedDevices returns request builder for ManagedDevice collection
func (b *DetectedAppRequestBuilder) ManagedDevices() *DetectedAppManagedDevicesCollectionRequestBuilder {
	bb := &DetectedAppManagedDevicesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/managedDevices"
	return bb
}

// DetectedAppManagedDevicesCollectionRequestBuilder is request builder for ManagedDevice collection
type DetectedAppManagedDevicesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for ManagedDevice collection
func (b *DetectedAppManagedDevicesCollectionRequestBuilder) Request() *DetectedAppManagedDevicesCollectionRequest {
	return &DetectedAppManagedDevicesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for ManagedDevice item
func (b *DetectedAppManagedDevicesCollectionRequestBuilder) ID(id string) *ManagedDeviceRequestBuilder {
	bb := &ManagedDeviceRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// DetectedAppManagedDevicesCollectionRequest is request for ManagedDevice collection
type DetectedAppManagedDevicesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for ManagedDevice collection
func (r *DetectedAppManagedDevicesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]ManagedDevice, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []ManagedDevice
	for {
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			res.Body.Close()
			errRes := &ErrorResponse{Response: res}
			err := jsonx.Unmarshal(b, errRes)
			if err != nil {
				return nil, fmt.Errorf("%s: %s", res.Status, string(b))
			}
			return nil, errRes
		}
		var (
			paging Paging
			value  []ManagedDevice
		)
		err := jsonx.NewDecoder(res.Body).Decode(&paging)
		res.Body.Close()
		if err != nil {
			return nil, err
		}
		err = jsonx.Unmarshal(paging.Value, &value)
		if err != nil {
			return nil, err
		}
		values = append(values, value...)
		if n >= 0 {
			n--
		}
		if n == 0 || len(paging.NextLink) == 0 {
			return values, nil
		}
		req, err = http.NewRequest("GET", paging.NextLink, nil)
		if ctx != nil {
			req = req.WithContext(ctx)
		}
		res, err = r.client.Do(req)
		if err != nil {
			return nil, err
		}
	}
}

// GetN performs GET request for ManagedDevice collection, max N pages
func (r *DetectedAppManagedDevicesCollectionRequest) GetN(ctx context.Context, n int) ([]ManagedDevice, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for ManagedDevice collection
func (r *DetectedAppManagedDevicesCollectionRequest) Get(ctx context.Context) ([]ManagedDevice, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for ManagedDevice collection
func (r *DetectedAppManagedDevicesCollectionRequest) Add(ctx context.Context, reqObj *ManagedDevice) (resObj *ManagedDevice, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}
