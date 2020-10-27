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

// JournalPostRequestParameter undocumented
type JournalPostRequestParameter struct {
}

// Account is navigation property
func (b *JournalRequestBuilder) Account() *AccountRequestBuilder {
	bb := &AccountRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/account"
	return bb
}

// JournalLines returns request builder for JournalLine collection
func (b *JournalRequestBuilder) JournalLines() *JournalJournalLinesCollectionRequestBuilder {
	bb := &JournalJournalLinesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/journalLines"
	return bb
}

// JournalJournalLinesCollectionRequestBuilder is request builder for JournalLine collection
type JournalJournalLinesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for JournalLine collection
func (b *JournalJournalLinesCollectionRequestBuilder) Request() *JournalJournalLinesCollectionRequest {
	return &JournalJournalLinesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client,
			tenantID: b.tenantID, applicationID: b.applicationID, clientSecurityKey: b.clientSecurityKey},
	}
}

// ID returns request builder for JournalLine item
func (b *JournalJournalLinesCollectionRequestBuilder) ID(id string) *JournalLineRequestBuilder {
	bb := &JournalLineRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// JournalJournalLinesCollectionRequest is request for JournalLine collection
type JournalJournalLinesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for JournalLine collection
func (r *JournalJournalLinesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]JournalLine, error) {
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
	var values []JournalLine
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
			value  []JournalLine
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
func (r *JournalJournalLinesCollectionRequest) NewRequest(method, url string, body io.Reader) (*http.Request, error) {
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

// GetN performs GET request for JournalLine collection, max N pages
func (r *JournalJournalLinesCollectionRequest) GetN(ctx context.Context, n int) ([]JournalLine, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for JournalLine collection
func (r *JournalJournalLinesCollectionRequest) Get(ctx context.Context) ([]JournalLine, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for JournalLine collection
func (r *JournalJournalLinesCollectionRequest) Add(ctx context.Context, reqObj *JournalLine) (resObj *JournalLine, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Account is navigation property
func (b *JournalLineRequestBuilder) Account() *AccountRequestBuilder {
	bb := &AccountRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/account"
	return bb
}
