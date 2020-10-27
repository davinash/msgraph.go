// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/davinash/msgraph.go/jsonx"
)

// ApplicationAddKeyRequestParameter undocumented
type ApplicationAddKeyRequestParameter struct {
	// KeyCredential undocumented
	KeyCredential *KeyCredential `json:"keyCredential,omitempty"`
	// PasswordCredential undocumented
	PasswordCredential *PasswordCredential `json:"passwordCredential,omitempty"`
	// Proof undocumented
	Proof *string `json:"proof,omitempty"`
}

// ApplicationAddPasswordRequestParameter undocumented
type ApplicationAddPasswordRequestParameter struct {
	// PasswordCredential undocumented
	PasswordCredential *PasswordCredential `json:"passwordCredential,omitempty"`
}

// ApplicationRemoveKeyRequestParameter undocumented
type ApplicationRemoveKeyRequestParameter struct {
	// KeyID undocumented
	KeyID *UUID `json:"keyId,omitempty"`
	// Proof undocumented
	Proof *string `json:"proof,omitempty"`
}

// ApplicationRemovePasswordRequestParameter undocumented
type ApplicationRemovePasswordRequestParameter struct {
	// KeyID undocumented
	KeyID *UUID `json:"keyId,omitempty"`
}

// ApplicationTemplateInstantiateRequestParameter undocumented
type ApplicationTemplateInstantiateRequestParameter struct {
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
}

// CreatedOnBehalfOf is navigation property
func (b *ApplicationRequestBuilder) CreatedOnBehalfOf() *DirectoryObjectRequestBuilder {
	bb := &DirectoryObjectRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/createdOnBehalfOf"
	return bb
}

// ExtensionProperties returns request builder for ExtensionProperty collection
func (b *ApplicationRequestBuilder) ExtensionProperties() *ApplicationExtensionPropertiesCollectionRequestBuilder {
	bb := &ApplicationExtensionPropertiesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/extensionProperties"
	return bb
}

// ApplicationExtensionPropertiesCollectionRequestBuilder is request builder for ExtensionProperty collection
type ApplicationExtensionPropertiesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for ExtensionProperty collection
func (b *ApplicationExtensionPropertiesCollectionRequestBuilder) Request() *ApplicationExtensionPropertiesCollectionRequest {
	return &ApplicationExtensionPropertiesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client,
			tenantID: b.tenantID, applicationID: b.applicationID, clientSecurityKey: b.clientSecurityKey},
	}
}

// ID returns request builder for ExtensionProperty item
func (b *ApplicationExtensionPropertiesCollectionRequestBuilder) ID(id string) *ExtensionPropertyRequestBuilder {
	bb := &ExtensionPropertyRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// ApplicationExtensionPropertiesCollectionRequest is request for ExtensionProperty collection
type ApplicationExtensionPropertiesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for ExtensionProperty collection
func (r *ApplicationExtensionPropertiesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]ExtensionProperty, error) {
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
	var values []ExtensionProperty
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
			value  []ExtensionProperty
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
		req, err = r.NewRequest("GET", paging.NextLink, nil)
		if ctx != nil {
			req = req.WithContext(ctx)
		}
		res, err = r.client.Do(req)
		if err != nil {
			return nil, err
		}
	}
}

// NewRequest Wrapper over the http.NewRequest with adding auth tokens
func (r *ApplicationExtensionPropertiesCollectionRequest) NewRequest(method, url string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequestWithContext(context.Background(), method, url, body)
	if err != nil {
		return nil, err
	}
	err = r.getAuthToken()
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", r.token.GetAccessToken())
	return req, err
}

// GetN performs GET request for ExtensionProperty collection, max N pages
func (r *ApplicationExtensionPropertiesCollectionRequest) GetN(ctx context.Context, n int) ([]ExtensionProperty, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for ExtensionProperty collection
func (r *ApplicationExtensionPropertiesCollectionRequest) Get(ctx context.Context) ([]ExtensionProperty, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for ExtensionProperty collection
func (r *ApplicationExtensionPropertiesCollectionRequest) Add(ctx context.Context, reqObj *ExtensionProperty) (resObj *ExtensionProperty, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Owners returns request builder for DirectoryObject collection
func (b *ApplicationRequestBuilder) Owners() *ApplicationOwnersCollectionRequestBuilder {
	bb := &ApplicationOwnersCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/owners"
	return bb
}

// ApplicationOwnersCollectionRequestBuilder is request builder for DirectoryObject collection
type ApplicationOwnersCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for DirectoryObject collection
func (b *ApplicationOwnersCollectionRequestBuilder) Request() *ApplicationOwnersCollectionRequest {
	return &ApplicationOwnersCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client,
			tenantID: b.tenantID, applicationID: b.applicationID, clientSecurityKey: b.clientSecurityKey},
	}
}

// ID returns request builder for DirectoryObject item
func (b *ApplicationOwnersCollectionRequestBuilder) ID(id string) *DirectoryObjectRequestBuilder {
	bb := &DirectoryObjectRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// ApplicationOwnersCollectionRequest is request for DirectoryObject collection
type ApplicationOwnersCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for DirectoryObject collection
func (r *ApplicationOwnersCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]DirectoryObject, error) {
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
	var values []DirectoryObject
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
			value  []DirectoryObject
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
		req, err = r.NewRequest("GET", paging.NextLink, nil)
		if ctx != nil {
			req = req.WithContext(ctx)
		}
		res, err = r.client.Do(req)
		if err != nil {
			return nil, err
		}
	}
}

// NewRequest Wrapper over the http.NewRequest with adding auth tokens
func (r *ApplicationOwnersCollectionRequest) NewRequest(method, url string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequestWithContext(context.Background(), method, url, body)
	if err != nil {
		return nil, err
	}
	err = r.getAuthToken()
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", r.token.GetAccessToken())
	return req, err
}

// GetN performs GET request for DirectoryObject collection, max N pages
func (r *ApplicationOwnersCollectionRequest) GetN(ctx context.Context, n int) ([]DirectoryObject, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for DirectoryObject collection
func (r *ApplicationOwnersCollectionRequest) Get(ctx context.Context) ([]DirectoryObject, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for DirectoryObject collection
func (r *ApplicationOwnersCollectionRequest) Add(ctx context.Context, reqObj *DirectoryObject) (resObj *DirectoryObject, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Policies returns request builder for DirectoryObject collection
func (b *ApplicationRequestBuilder) Policies() *ApplicationPoliciesCollectionRequestBuilder {
	bb := &ApplicationPoliciesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/policies"
	return bb
}

// ApplicationPoliciesCollectionRequestBuilder is request builder for DirectoryObject collection
type ApplicationPoliciesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for DirectoryObject collection
func (b *ApplicationPoliciesCollectionRequestBuilder) Request() *ApplicationPoliciesCollectionRequest {
	return &ApplicationPoliciesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client,
			tenantID: b.tenantID, applicationID: b.applicationID, clientSecurityKey: b.clientSecurityKey},
	}
}

// ID returns request builder for DirectoryObject item
func (b *ApplicationPoliciesCollectionRequestBuilder) ID(id string) *DirectoryObjectRequestBuilder {
	bb := &DirectoryObjectRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// ApplicationPoliciesCollectionRequest is request for DirectoryObject collection
type ApplicationPoliciesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for DirectoryObject collection
func (r *ApplicationPoliciesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]DirectoryObject, error) {
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
	var values []DirectoryObject
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
			value  []DirectoryObject
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
		req, err = r.NewRequest("GET", paging.NextLink, nil)
		if ctx != nil {
			req = req.WithContext(ctx)
		}
		res, err = r.client.Do(req)
		if err != nil {
			return nil, err
		}
	}
}

// NewRequest Wrapper over the http.NewRequest with adding auth tokens
func (r *ApplicationPoliciesCollectionRequest) NewRequest(method, url string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequestWithContext(context.Background(), method, url, body)
	if err != nil {
		return nil, err
	}
	err = r.getAuthToken()
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", r.token.GetAccessToken())
	return req, err
}

// GetN performs GET request for DirectoryObject collection, max N pages
func (r *ApplicationPoliciesCollectionRequest) GetN(ctx context.Context, n int) ([]DirectoryObject, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for DirectoryObject collection
func (r *ApplicationPoliciesCollectionRequest) Get(ctx context.Context) ([]DirectoryObject, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for DirectoryObject collection
func (r *ApplicationPoliciesCollectionRequest) Add(ctx context.Context, reqObj *DirectoryObject) (resObj *DirectoryObject, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Synchronization is navigation property
func (b *ApplicationRequestBuilder) Synchronization() *SynchronizationRequestBuilder {
	bb := &SynchronizationRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/synchronization"
	return bb
}
