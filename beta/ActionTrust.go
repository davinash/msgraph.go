// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/davinash/msgraph.go/jsonx"
)

// TrustFrameworkKeySetGenerateKeyRequestParameter undocumented
type TrustFrameworkKeySetGenerateKeyRequestParameter struct {
	// Use undocumented
	Use *string `json:"use,omitempty"`
	// Kty undocumented
	Kty *string `json:"kty,omitempty"`
	// Nbf undocumented
	Nbf *int `json:"nbf,omitempty"`
	// Exp undocumented
	Exp *int `json:"exp,omitempty"`
}

// TrustFrameworkKeySetUploadSecretRequestParameter undocumented
type TrustFrameworkKeySetUploadSecretRequestParameter struct {
	// Use undocumented
	Use *string `json:"use,omitempty"`
	// K undocumented
	K *string `json:"k,omitempty"`
	// Nbf undocumented
	Nbf *int `json:"nbf,omitempty"`
	// Exp undocumented
	Exp *int `json:"exp,omitempty"`
}

// TrustFrameworkKeySetUploadCertificateRequestParameter undocumented
type TrustFrameworkKeySetUploadCertificateRequestParameter struct {
	// Key undocumented
	Key *string `json:"key,omitempty"`
}

// TrustFrameworkKeySetUploadPkcs12RequestParameter undocumented
type TrustFrameworkKeySetUploadPkcs12RequestParameter struct {
	// Key undocumented
	Key *string `json:"key,omitempty"`
	// Password undocumented
	Password *string `json:"password,omitempty"`
}

// KeySets returns request builder for TrustFrameworkKeySet collection
func (b *TrustFrameworkRequestBuilder) KeySets() *TrustFrameworkKeySetsCollectionRequestBuilder {
	bb := &TrustFrameworkKeySetsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/keySets"
	return bb
}

// TrustFrameworkKeySetsCollectionRequestBuilder is request builder for TrustFrameworkKeySet collection
type TrustFrameworkKeySetsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for TrustFrameworkKeySet collection
func (b *TrustFrameworkKeySetsCollectionRequestBuilder) Request() *TrustFrameworkKeySetsCollectionRequest {
	return &TrustFrameworkKeySetsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for TrustFrameworkKeySet item
func (b *TrustFrameworkKeySetsCollectionRequestBuilder) ID(id string) *TrustFrameworkKeySetRequestBuilder {
	bb := &TrustFrameworkKeySetRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// TrustFrameworkKeySetsCollectionRequest is request for TrustFrameworkKeySet collection
type TrustFrameworkKeySetsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for TrustFrameworkKeySet collection
func (r *TrustFrameworkKeySetsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]TrustFrameworkKeySet, error) {
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
	var values []TrustFrameworkKeySet
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
			value  []TrustFrameworkKeySet
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

// GetN performs GET request for TrustFrameworkKeySet collection, max N pages
func (r *TrustFrameworkKeySetsCollectionRequest) GetN(ctx context.Context, n int) ([]TrustFrameworkKeySet, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for TrustFrameworkKeySet collection
func (r *TrustFrameworkKeySetsCollectionRequest) Get(ctx context.Context) ([]TrustFrameworkKeySet, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for TrustFrameworkKeySet collection
func (r *TrustFrameworkKeySetsCollectionRequest) Add(ctx context.Context, reqObj *TrustFrameworkKeySet) (resObj *TrustFrameworkKeySet, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Policies returns request builder for TrustFrameworkPolicy collection
func (b *TrustFrameworkRequestBuilder) Policies() *TrustFrameworkPoliciesCollectionRequestBuilder {
	bb := &TrustFrameworkPoliciesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/policies"
	return bb
}

// TrustFrameworkPoliciesCollectionRequestBuilder is request builder for TrustFrameworkPolicy collection
type TrustFrameworkPoliciesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for TrustFrameworkPolicy collection
func (b *TrustFrameworkPoliciesCollectionRequestBuilder) Request() *TrustFrameworkPoliciesCollectionRequest {
	return &TrustFrameworkPoliciesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for TrustFrameworkPolicy item
func (b *TrustFrameworkPoliciesCollectionRequestBuilder) ID(id string) *TrustFrameworkPolicyRequestBuilder {
	bb := &TrustFrameworkPolicyRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// TrustFrameworkPoliciesCollectionRequest is request for TrustFrameworkPolicy collection
type TrustFrameworkPoliciesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for TrustFrameworkPolicy collection
func (r *TrustFrameworkPoliciesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]TrustFrameworkPolicy, error) {
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
	var values []TrustFrameworkPolicy
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
			value  []TrustFrameworkPolicy
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

// GetN performs GET request for TrustFrameworkPolicy collection, max N pages
func (r *TrustFrameworkPoliciesCollectionRequest) GetN(ctx context.Context, n int) ([]TrustFrameworkPolicy, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for TrustFrameworkPolicy collection
func (r *TrustFrameworkPoliciesCollectionRequest) Get(ctx context.Context) ([]TrustFrameworkPolicy, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for TrustFrameworkPolicy collection
func (r *TrustFrameworkPoliciesCollectionRequest) Add(ctx context.Context, reqObj *TrustFrameworkPolicy) (resObj *TrustFrameworkPolicy, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}
