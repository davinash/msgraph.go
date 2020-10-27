// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/davinash/msgraph.go/jsonx"
)

// AgentGroups returns request builder for OnPremisesAgentGroup collection
func (b *OnPremisesAgentRequestBuilder) AgentGroups() *OnPremisesAgentAgentGroupsCollectionRequestBuilder {
	bb := &OnPremisesAgentAgentGroupsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/agentGroups"
	return bb
}

// OnPremisesAgentAgentGroupsCollectionRequestBuilder is request builder for OnPremisesAgentGroup collection
type OnPremisesAgentAgentGroupsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for OnPremisesAgentGroup collection
func (b *OnPremisesAgentAgentGroupsCollectionRequestBuilder) Request() *OnPremisesAgentAgentGroupsCollectionRequest {
	return &OnPremisesAgentAgentGroupsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for OnPremisesAgentGroup item
func (b *OnPremisesAgentAgentGroupsCollectionRequestBuilder) ID(id string) *OnPremisesAgentGroupRequestBuilder {
	bb := &OnPremisesAgentGroupRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// OnPremisesAgentAgentGroupsCollectionRequest is request for OnPremisesAgentGroup collection
type OnPremisesAgentAgentGroupsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for OnPremisesAgentGroup collection
func (r *OnPremisesAgentAgentGroupsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]OnPremisesAgentGroup, error) {
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
	var values []OnPremisesAgentGroup
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
			value  []OnPremisesAgentGroup
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

// GetN performs GET request for OnPremisesAgentGroup collection, max N pages
func (r *OnPremisesAgentAgentGroupsCollectionRequest) GetN(ctx context.Context, n int) ([]OnPremisesAgentGroup, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for OnPremisesAgentGroup collection
func (r *OnPremisesAgentAgentGroupsCollectionRequest) Get(ctx context.Context) ([]OnPremisesAgentGroup, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for OnPremisesAgentGroup collection
func (r *OnPremisesAgentAgentGroupsCollectionRequest) Add(ctx context.Context, reqObj *OnPremisesAgentGroup) (resObj *OnPremisesAgentGroup, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Agents returns request builder for OnPremisesAgent collection
func (b *OnPremisesAgentGroupRequestBuilder) Agents() *OnPremisesAgentGroupAgentsCollectionRequestBuilder {
	bb := &OnPremisesAgentGroupAgentsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/agents"
	return bb
}

// OnPremisesAgentGroupAgentsCollectionRequestBuilder is request builder for OnPremisesAgent collection
type OnPremisesAgentGroupAgentsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for OnPremisesAgent collection
func (b *OnPremisesAgentGroupAgentsCollectionRequestBuilder) Request() *OnPremisesAgentGroupAgentsCollectionRequest {
	return &OnPremisesAgentGroupAgentsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for OnPremisesAgent item
func (b *OnPremisesAgentGroupAgentsCollectionRequestBuilder) ID(id string) *OnPremisesAgentRequestBuilder {
	bb := &OnPremisesAgentRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// OnPremisesAgentGroupAgentsCollectionRequest is request for OnPremisesAgent collection
type OnPremisesAgentGroupAgentsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for OnPremisesAgent collection
func (r *OnPremisesAgentGroupAgentsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]OnPremisesAgent, error) {
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
	var values []OnPremisesAgent
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
			value  []OnPremisesAgent
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

// GetN performs GET request for OnPremisesAgent collection, max N pages
func (r *OnPremisesAgentGroupAgentsCollectionRequest) GetN(ctx context.Context, n int) ([]OnPremisesAgent, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for OnPremisesAgent collection
func (r *OnPremisesAgentGroupAgentsCollectionRequest) Get(ctx context.Context) ([]OnPremisesAgent, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for OnPremisesAgent collection
func (r *OnPremisesAgentGroupAgentsCollectionRequest) Add(ctx context.Context, reqObj *OnPremisesAgent) (resObj *OnPremisesAgent, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// PublishedResources returns request builder for PublishedResource collection
func (b *OnPremisesAgentGroupRequestBuilder) PublishedResources() *OnPremisesAgentGroupPublishedResourcesCollectionRequestBuilder {
	bb := &OnPremisesAgentGroupPublishedResourcesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/publishedResources"
	return bb
}

// OnPremisesAgentGroupPublishedResourcesCollectionRequestBuilder is request builder for PublishedResource collection
type OnPremisesAgentGroupPublishedResourcesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for PublishedResource collection
func (b *OnPremisesAgentGroupPublishedResourcesCollectionRequestBuilder) Request() *OnPremisesAgentGroupPublishedResourcesCollectionRequest {
	return &OnPremisesAgentGroupPublishedResourcesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for PublishedResource item
func (b *OnPremisesAgentGroupPublishedResourcesCollectionRequestBuilder) ID(id string) *PublishedResourceRequestBuilder {
	bb := &PublishedResourceRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// OnPremisesAgentGroupPublishedResourcesCollectionRequest is request for PublishedResource collection
type OnPremisesAgentGroupPublishedResourcesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for PublishedResource collection
func (r *OnPremisesAgentGroupPublishedResourcesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]PublishedResource, error) {
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
	var values []PublishedResource
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
			value  []PublishedResource
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

// GetN performs GET request for PublishedResource collection, max N pages
func (r *OnPremisesAgentGroupPublishedResourcesCollectionRequest) GetN(ctx context.Context, n int) ([]PublishedResource, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for PublishedResource collection
func (r *OnPremisesAgentGroupPublishedResourcesCollectionRequest) Get(ctx context.Context) ([]PublishedResource, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for PublishedResource collection
func (r *OnPremisesAgentGroupPublishedResourcesCollectionRequest) Add(ctx context.Context, reqObj *PublishedResource) (resObj *PublishedResource, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// AgentGroups returns request builder for OnPremisesAgentGroup collection
func (b *OnPremisesPublishingProfileRequestBuilder) AgentGroups() *OnPremisesPublishingProfileAgentGroupsCollectionRequestBuilder {
	bb := &OnPremisesPublishingProfileAgentGroupsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/agentGroups"
	return bb
}

// OnPremisesPublishingProfileAgentGroupsCollectionRequestBuilder is request builder for OnPremisesAgentGroup collection
type OnPremisesPublishingProfileAgentGroupsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for OnPremisesAgentGroup collection
func (b *OnPremisesPublishingProfileAgentGroupsCollectionRequestBuilder) Request() *OnPremisesPublishingProfileAgentGroupsCollectionRequest {
	return &OnPremisesPublishingProfileAgentGroupsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for OnPremisesAgentGroup item
func (b *OnPremisesPublishingProfileAgentGroupsCollectionRequestBuilder) ID(id string) *OnPremisesAgentGroupRequestBuilder {
	bb := &OnPremisesAgentGroupRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// OnPremisesPublishingProfileAgentGroupsCollectionRequest is request for OnPremisesAgentGroup collection
type OnPremisesPublishingProfileAgentGroupsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for OnPremisesAgentGroup collection
func (r *OnPremisesPublishingProfileAgentGroupsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]OnPremisesAgentGroup, error) {
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
	var values []OnPremisesAgentGroup
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
			value  []OnPremisesAgentGroup
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

// GetN performs GET request for OnPremisesAgentGroup collection, max N pages
func (r *OnPremisesPublishingProfileAgentGroupsCollectionRequest) GetN(ctx context.Context, n int) ([]OnPremisesAgentGroup, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for OnPremisesAgentGroup collection
func (r *OnPremisesPublishingProfileAgentGroupsCollectionRequest) Get(ctx context.Context) ([]OnPremisesAgentGroup, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for OnPremisesAgentGroup collection
func (r *OnPremisesPublishingProfileAgentGroupsCollectionRequest) Add(ctx context.Context, reqObj *OnPremisesAgentGroup) (resObj *OnPremisesAgentGroup, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Agents returns request builder for OnPremisesAgent collection
func (b *OnPremisesPublishingProfileRequestBuilder) Agents() *OnPremisesPublishingProfileAgentsCollectionRequestBuilder {
	bb := &OnPremisesPublishingProfileAgentsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/agents"
	return bb
}

// OnPremisesPublishingProfileAgentsCollectionRequestBuilder is request builder for OnPremisesAgent collection
type OnPremisesPublishingProfileAgentsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for OnPremisesAgent collection
func (b *OnPremisesPublishingProfileAgentsCollectionRequestBuilder) Request() *OnPremisesPublishingProfileAgentsCollectionRequest {
	return &OnPremisesPublishingProfileAgentsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for OnPremisesAgent item
func (b *OnPremisesPublishingProfileAgentsCollectionRequestBuilder) ID(id string) *OnPremisesAgentRequestBuilder {
	bb := &OnPremisesAgentRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// OnPremisesPublishingProfileAgentsCollectionRequest is request for OnPremisesAgent collection
type OnPremisesPublishingProfileAgentsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for OnPremisesAgent collection
func (r *OnPremisesPublishingProfileAgentsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]OnPremisesAgent, error) {
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
	var values []OnPremisesAgent
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
			value  []OnPremisesAgent
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

// GetN performs GET request for OnPremisesAgent collection, max N pages
func (r *OnPremisesPublishingProfileAgentsCollectionRequest) GetN(ctx context.Context, n int) ([]OnPremisesAgent, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for OnPremisesAgent collection
func (r *OnPremisesPublishingProfileAgentsCollectionRequest) Get(ctx context.Context) ([]OnPremisesAgent, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for OnPremisesAgent collection
func (r *OnPremisesPublishingProfileAgentsCollectionRequest) Add(ctx context.Context, reqObj *OnPremisesAgent) (resObj *OnPremisesAgent, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// PublishedResources returns request builder for PublishedResource collection
func (b *OnPremisesPublishingProfileRequestBuilder) PublishedResources() *OnPremisesPublishingProfilePublishedResourcesCollectionRequestBuilder {
	bb := &OnPremisesPublishingProfilePublishedResourcesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/publishedResources"
	return bb
}

// OnPremisesPublishingProfilePublishedResourcesCollectionRequestBuilder is request builder for PublishedResource collection
type OnPremisesPublishingProfilePublishedResourcesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for PublishedResource collection
func (b *OnPremisesPublishingProfilePublishedResourcesCollectionRequestBuilder) Request() *OnPremisesPublishingProfilePublishedResourcesCollectionRequest {
	return &OnPremisesPublishingProfilePublishedResourcesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for PublishedResource item
func (b *OnPremisesPublishingProfilePublishedResourcesCollectionRequestBuilder) ID(id string) *PublishedResourceRequestBuilder {
	bb := &PublishedResourceRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// OnPremisesPublishingProfilePublishedResourcesCollectionRequest is request for PublishedResource collection
type OnPremisesPublishingProfilePublishedResourcesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for PublishedResource collection
func (r *OnPremisesPublishingProfilePublishedResourcesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]PublishedResource, error) {
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
	var values []PublishedResource
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
			value  []PublishedResource
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

// GetN performs GET request for PublishedResource collection, max N pages
func (r *OnPremisesPublishingProfilePublishedResourcesCollectionRequest) GetN(ctx context.Context, n int) ([]PublishedResource, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for PublishedResource collection
func (r *OnPremisesPublishingProfilePublishedResourcesCollectionRequest) Get(ctx context.Context) ([]PublishedResource, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for PublishedResource collection
func (r *OnPremisesPublishingProfilePublishedResourcesCollectionRequest) Add(ctx context.Context, reqObj *PublishedResource) (resObj *PublishedResource, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}
