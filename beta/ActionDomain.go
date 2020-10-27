// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/davinash/msgraph.go/jsonx"
)

// DomainForceDeleteRequestParameter undocumented
type DomainForceDeleteRequestParameter struct {
	// DisableUserAccounts undocumented
	DisableUserAccounts *bool `json:"disableUserAccounts,omitempty"`
}

// DomainVerifyRequestParameter undocumented
type DomainVerifyRequestParameter struct {
}

// DomainNameReferences returns request builder for DirectoryObject collection
func (b *DomainRequestBuilder) DomainNameReferences() *DomainDomainNameReferencesCollectionRequestBuilder {
	bb := &DomainDomainNameReferencesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/domainNameReferences"
	return bb
}

// DomainDomainNameReferencesCollectionRequestBuilder is request builder for DirectoryObject collection
type DomainDomainNameReferencesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for DirectoryObject collection
func (b *DomainDomainNameReferencesCollectionRequestBuilder) Request() *DomainDomainNameReferencesCollectionRequest {
	return &DomainDomainNameReferencesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for DirectoryObject item
func (b *DomainDomainNameReferencesCollectionRequestBuilder) ID(id string) *DirectoryObjectRequestBuilder {
	bb := &DirectoryObjectRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// DomainDomainNameReferencesCollectionRequest is request for DirectoryObject collection
type DomainDomainNameReferencesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for DirectoryObject collection
func (r *DomainDomainNameReferencesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]DirectoryObject, error) {
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

// GetN performs GET request for DirectoryObject collection, max N pages
func (r *DomainDomainNameReferencesCollectionRequest) GetN(ctx context.Context, n int) ([]DirectoryObject, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for DirectoryObject collection
func (r *DomainDomainNameReferencesCollectionRequest) Get(ctx context.Context) ([]DirectoryObject, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for DirectoryObject collection
func (r *DomainDomainNameReferencesCollectionRequest) Add(ctx context.Context, reqObj *DirectoryObject) (resObj *DirectoryObject, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// ServiceConfigurationRecords returns request builder for DomainDNSRecord collection
func (b *DomainRequestBuilder) ServiceConfigurationRecords() *DomainServiceConfigurationRecordsCollectionRequestBuilder {
	bb := &DomainServiceConfigurationRecordsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/serviceConfigurationRecords"
	return bb
}

// DomainServiceConfigurationRecordsCollectionRequestBuilder is request builder for DomainDNSRecord collection
type DomainServiceConfigurationRecordsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for DomainDNSRecord collection
func (b *DomainServiceConfigurationRecordsCollectionRequestBuilder) Request() *DomainServiceConfigurationRecordsCollectionRequest {
	return &DomainServiceConfigurationRecordsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for DomainDNSRecord item
func (b *DomainServiceConfigurationRecordsCollectionRequestBuilder) ID(id string) *DomainDNSRecordRequestBuilder {
	bb := &DomainDNSRecordRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// DomainServiceConfigurationRecordsCollectionRequest is request for DomainDNSRecord collection
type DomainServiceConfigurationRecordsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for DomainDNSRecord collection
func (r *DomainServiceConfigurationRecordsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]DomainDNSRecord, error) {
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
	var values []DomainDNSRecord
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
			value  []DomainDNSRecord
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

// GetN performs GET request for DomainDNSRecord collection, max N pages
func (r *DomainServiceConfigurationRecordsCollectionRequest) GetN(ctx context.Context, n int) ([]DomainDNSRecord, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for DomainDNSRecord collection
func (r *DomainServiceConfigurationRecordsCollectionRequest) Get(ctx context.Context) ([]DomainDNSRecord, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for DomainDNSRecord collection
func (r *DomainServiceConfigurationRecordsCollectionRequest) Add(ctx context.Context, reqObj *DomainDNSRecord) (resObj *DomainDNSRecord, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// VerificationDNSRecords returns request builder for DomainDNSRecord collection
func (b *DomainRequestBuilder) VerificationDNSRecords() *DomainVerificationDNSRecordsCollectionRequestBuilder {
	bb := &DomainVerificationDNSRecordsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/verificationDnsRecords"
	return bb
}

// DomainVerificationDNSRecordsCollectionRequestBuilder is request builder for DomainDNSRecord collection
type DomainVerificationDNSRecordsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for DomainDNSRecord collection
func (b *DomainVerificationDNSRecordsCollectionRequestBuilder) Request() *DomainVerificationDNSRecordsCollectionRequest {
	return &DomainVerificationDNSRecordsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for DomainDNSRecord item
func (b *DomainVerificationDNSRecordsCollectionRequestBuilder) ID(id string) *DomainDNSRecordRequestBuilder {
	bb := &DomainDNSRecordRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// DomainVerificationDNSRecordsCollectionRequest is request for DomainDNSRecord collection
type DomainVerificationDNSRecordsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for DomainDNSRecord collection
func (r *DomainVerificationDNSRecordsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]DomainDNSRecord, error) {
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
	var values []DomainDNSRecord
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
			value  []DomainDNSRecord
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

// GetN performs GET request for DomainDNSRecord collection, max N pages
func (r *DomainVerificationDNSRecordsCollectionRequest) GetN(ctx context.Context, n int) ([]DomainDNSRecord, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for DomainDNSRecord collection
func (r *DomainVerificationDNSRecordsCollectionRequest) Get(ctx context.Context) ([]DomainDNSRecord, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for DomainDNSRecord collection
func (r *DomainVerificationDNSRecordsCollectionRequest) Add(ctx context.Context, reqObj *DomainDNSRecord) (resObj *DomainDNSRecord, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}
