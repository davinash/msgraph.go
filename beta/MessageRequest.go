// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// MessageRequestBuilder is request builder for Message
type MessageRequestBuilder struct{ BaseRequestBuilder }

// Request returns MessageRequest
func (b *MessageRequestBuilder) Request() *MessageRequest {
	return &MessageRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// MessageRequest is request for Message
type MessageRequest struct{ BaseRequest }

// Do performs HTTP request for Message
func (r *MessageRequest) Do(method, path string, reqObj interface{}) (resObj *Message, err error) {
	err = r.JSONRequestWithPath(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for Message
func (r *MessageRequest) Get() (*Message, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for Message
func (r *MessageRequest) Update(reqObj *Message) (*Message, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for Message
func (r *MessageRequest) Delete() error {
	return r.JSONRequestWithPath("DELETE", "", nil, nil)
}

// Attachments returns request builder for Attachment collection
func (b *MessageRequestBuilder) Attachments() *MessageAttachmentsCollectionRequestBuilder {
	bb := &MessageAttachmentsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/attachments"
	return bb
}

// MessageAttachmentsCollectionRequestBuilder is request builder for Attachment collection
type MessageAttachmentsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for Attachment collection
func (b *MessageAttachmentsCollectionRequestBuilder) Request() *MessageAttachmentsCollectionRequest {
	return &MessageAttachmentsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for Attachment item
func (b *MessageAttachmentsCollectionRequestBuilder) ID(id string) *AttachmentRequestBuilder {
	bb := &AttachmentRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// MessageAttachmentsCollectionRequest is request for Attachment collection
type MessageAttachmentsCollectionRequest struct{ BaseRequest }

// Do performs HTTP request for Attachment collection
func (r *MessageAttachmentsCollectionRequest) Do(method, path string, reqObj interface{}) (resObj *Attachment, err error) {
	err = r.JSONRequestWithPath(method, path, reqObj, &resObj)
	return
}

// Paging perfoms paging operation for Attachment collection
func (r *MessageAttachmentsCollectionRequest) Paging(method, path string, obj interface{}) ([]Attachment, error) {
	req, err := r.NewJSONRequestWithPath(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []Attachment
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  []Attachment
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

// Get performs GET request for Attachment collection
func (r *MessageAttachmentsCollectionRequest) Get() ([]Attachment, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for Attachment collection
func (r *MessageAttachmentsCollectionRequest) Add(reqObj *Attachment) (*Attachment, error) {
	return r.Do("POST", "", reqObj)
}

// Extensions returns request builder for Extension collection
func (b *MessageRequestBuilder) Extensions() *MessageExtensionsCollectionRequestBuilder {
	bb := &MessageExtensionsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/extensions"
	return bb
}

// MessageExtensionsCollectionRequestBuilder is request builder for Extension collection
type MessageExtensionsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for Extension collection
func (b *MessageExtensionsCollectionRequestBuilder) Request() *MessageExtensionsCollectionRequest {
	return &MessageExtensionsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for Extension item
func (b *MessageExtensionsCollectionRequestBuilder) ID(id string) *ExtensionRequestBuilder {
	bb := &ExtensionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// MessageExtensionsCollectionRequest is request for Extension collection
type MessageExtensionsCollectionRequest struct{ BaseRequest }

// Do performs HTTP request for Extension collection
func (r *MessageExtensionsCollectionRequest) Do(method, path string, reqObj interface{}) (resObj *Extension, err error) {
	err = r.JSONRequestWithPath(method, path, reqObj, &resObj)
	return
}

// Paging perfoms paging operation for Extension collection
func (r *MessageExtensionsCollectionRequest) Paging(method, path string, obj interface{}) ([]Extension, error) {
	req, err := r.NewJSONRequestWithPath(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []Extension
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  []Extension
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

// Get performs GET request for Extension collection
func (r *MessageExtensionsCollectionRequest) Get() ([]Extension, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for Extension collection
func (r *MessageExtensionsCollectionRequest) Add(reqObj *Extension) (*Extension, error) {
	return r.Do("POST", "", reqObj)
}

// Mentions returns request builder for Mention collection
func (b *MessageRequestBuilder) Mentions() *MessageMentionsCollectionRequestBuilder {
	bb := &MessageMentionsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/mentions"
	return bb
}

// MessageMentionsCollectionRequestBuilder is request builder for Mention collection
type MessageMentionsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for Mention collection
func (b *MessageMentionsCollectionRequestBuilder) Request() *MessageMentionsCollectionRequest {
	return &MessageMentionsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for Mention item
func (b *MessageMentionsCollectionRequestBuilder) ID(id string) *MentionRequestBuilder {
	bb := &MentionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// MessageMentionsCollectionRequest is request for Mention collection
type MessageMentionsCollectionRequest struct{ BaseRequest }

// Do performs HTTP request for Mention collection
func (r *MessageMentionsCollectionRequest) Do(method, path string, reqObj interface{}) (resObj *Mention, err error) {
	err = r.JSONRequestWithPath(method, path, reqObj, &resObj)
	return
}

// Paging perfoms paging operation for Mention collection
func (r *MessageMentionsCollectionRequest) Paging(method, path string, obj interface{}) ([]Mention, error) {
	req, err := r.NewJSONRequestWithPath(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []Mention
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  []Mention
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

// Get performs GET request for Mention collection
func (r *MessageMentionsCollectionRequest) Get() ([]Mention, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for Mention collection
func (r *MessageMentionsCollectionRequest) Add(reqObj *Mention) (*Mention, error) {
	return r.Do("POST", "", reqObj)
}

// MultiValueExtendedProperties returns request builder for MultiValueLegacyExtendedProperty collection
func (b *MessageRequestBuilder) MultiValueExtendedProperties() *MessageMultiValueExtendedPropertiesCollectionRequestBuilder {
	bb := &MessageMultiValueExtendedPropertiesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/multiValueExtendedProperties"
	return bb
}

// MessageMultiValueExtendedPropertiesCollectionRequestBuilder is request builder for MultiValueLegacyExtendedProperty collection
type MessageMultiValueExtendedPropertiesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for MultiValueLegacyExtendedProperty collection
func (b *MessageMultiValueExtendedPropertiesCollectionRequestBuilder) Request() *MessageMultiValueExtendedPropertiesCollectionRequest {
	return &MessageMultiValueExtendedPropertiesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for MultiValueLegacyExtendedProperty item
func (b *MessageMultiValueExtendedPropertiesCollectionRequestBuilder) ID(id string) *MultiValueLegacyExtendedPropertyRequestBuilder {
	bb := &MultiValueLegacyExtendedPropertyRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// MessageMultiValueExtendedPropertiesCollectionRequest is request for MultiValueLegacyExtendedProperty collection
type MessageMultiValueExtendedPropertiesCollectionRequest struct{ BaseRequest }

// Do performs HTTP request for MultiValueLegacyExtendedProperty collection
func (r *MessageMultiValueExtendedPropertiesCollectionRequest) Do(method, path string, reqObj interface{}) (resObj *MultiValueLegacyExtendedProperty, err error) {
	err = r.JSONRequestWithPath(method, path, reqObj, &resObj)
	return
}

// Paging perfoms paging operation for MultiValueLegacyExtendedProperty collection
func (r *MessageMultiValueExtendedPropertiesCollectionRequest) Paging(method, path string, obj interface{}) ([]MultiValueLegacyExtendedProperty, error) {
	req, err := r.NewJSONRequestWithPath(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []MultiValueLegacyExtendedProperty
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  []MultiValueLegacyExtendedProperty
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

// Get performs GET request for MultiValueLegacyExtendedProperty collection
func (r *MessageMultiValueExtendedPropertiesCollectionRequest) Get() ([]MultiValueLegacyExtendedProperty, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for MultiValueLegacyExtendedProperty collection
func (r *MessageMultiValueExtendedPropertiesCollectionRequest) Add(reqObj *MultiValueLegacyExtendedProperty) (*MultiValueLegacyExtendedProperty, error) {
	return r.Do("POST", "", reqObj)
}

// SingleValueExtendedProperties returns request builder for SingleValueLegacyExtendedProperty collection
func (b *MessageRequestBuilder) SingleValueExtendedProperties() *MessageSingleValueExtendedPropertiesCollectionRequestBuilder {
	bb := &MessageSingleValueExtendedPropertiesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/singleValueExtendedProperties"
	return bb
}

// MessageSingleValueExtendedPropertiesCollectionRequestBuilder is request builder for SingleValueLegacyExtendedProperty collection
type MessageSingleValueExtendedPropertiesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for SingleValueLegacyExtendedProperty collection
func (b *MessageSingleValueExtendedPropertiesCollectionRequestBuilder) Request() *MessageSingleValueExtendedPropertiesCollectionRequest {
	return &MessageSingleValueExtendedPropertiesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for SingleValueLegacyExtendedProperty item
func (b *MessageSingleValueExtendedPropertiesCollectionRequestBuilder) ID(id string) *SingleValueLegacyExtendedPropertyRequestBuilder {
	bb := &SingleValueLegacyExtendedPropertyRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// MessageSingleValueExtendedPropertiesCollectionRequest is request for SingleValueLegacyExtendedProperty collection
type MessageSingleValueExtendedPropertiesCollectionRequest struct{ BaseRequest }

// Do performs HTTP request for SingleValueLegacyExtendedProperty collection
func (r *MessageSingleValueExtendedPropertiesCollectionRequest) Do(method, path string, reqObj interface{}) (resObj *SingleValueLegacyExtendedProperty, err error) {
	err = r.JSONRequestWithPath(method, path, reqObj, &resObj)
	return
}

// Paging perfoms paging operation for SingleValueLegacyExtendedProperty collection
func (r *MessageSingleValueExtendedPropertiesCollectionRequest) Paging(method, path string, obj interface{}) ([]SingleValueLegacyExtendedProperty, error) {
	req, err := r.NewJSONRequestWithPath(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []SingleValueLegacyExtendedProperty
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  []SingleValueLegacyExtendedProperty
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

// Get performs GET request for SingleValueLegacyExtendedProperty collection
func (r *MessageSingleValueExtendedPropertiesCollectionRequest) Get() ([]SingleValueLegacyExtendedProperty, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for SingleValueLegacyExtendedProperty collection
func (r *MessageSingleValueExtendedPropertiesCollectionRequest) Add(reqObj *SingleValueLegacyExtendedProperty) (*SingleValueLegacyExtendedProperty, error) {
	return r.Do("POST", "", reqObj)
}