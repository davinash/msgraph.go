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

// CalendarGetScheduleRequestParameter undocumented
type CalendarGetScheduleRequestParameter struct {
	// Schedules undocumented
	Schedules []string `json:"Schedules,omitempty"`
	// EndTime undocumented
	EndTime *DateTimeTimeZone `json:"EndTime,omitempty"`
	// StartTime undocumented
	StartTime *DateTimeTimeZone `json:"StartTime,omitempty"`
	// AvailabilityViewInterval undocumented
	AvailabilityViewInterval *int `json:"AvailabilityViewInterval,omitempty"`
}

// CalendarView returns request builder for Event collection
func (b *CalendarRequestBuilder) CalendarView() *CalendarCalendarViewCollectionRequestBuilder {
	bb := &CalendarCalendarViewCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/calendarView"
	return bb
}

// CalendarCalendarViewCollectionRequestBuilder is request builder for Event collection
type CalendarCalendarViewCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for Event collection
func (b *CalendarCalendarViewCollectionRequestBuilder) Request() *CalendarCalendarViewCollectionRequest {
	return &CalendarCalendarViewCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client,
			tenantID: b.tenantID, applicationID: b.applicationID, clientSecurityKey: b.clientSecurityKey},
	}
}

// ID returns request builder for Event item
func (b *CalendarCalendarViewCollectionRequestBuilder) ID(id string) *EventRequestBuilder {
	bb := &EventRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// CalendarCalendarViewCollectionRequest is request for Event collection
type CalendarCalendarViewCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for Event collection
func (r *CalendarCalendarViewCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]Event, error) {
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
	var values []Event
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
			value  []Event
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
func (r *CalendarCalendarViewCollectionRequest) NewRequest(method, url string, body io.Reader) (*http.Request, error) {
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

// GetN performs GET request for Event collection, max N pages
func (r *CalendarCalendarViewCollectionRequest) GetN(ctx context.Context, n int) ([]Event, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for Event collection
func (r *CalendarCalendarViewCollectionRequest) Get(ctx context.Context) ([]Event, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for Event collection
func (r *CalendarCalendarViewCollectionRequest) Add(ctx context.Context, reqObj *Event) (resObj *Event, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Events returns request builder for Event collection
func (b *CalendarRequestBuilder) Events() *CalendarEventsCollectionRequestBuilder {
	bb := &CalendarEventsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/events"
	return bb
}

// CalendarEventsCollectionRequestBuilder is request builder for Event collection
type CalendarEventsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for Event collection
func (b *CalendarEventsCollectionRequestBuilder) Request() *CalendarEventsCollectionRequest {
	return &CalendarEventsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client,
			tenantID: b.tenantID, applicationID: b.applicationID, clientSecurityKey: b.clientSecurityKey},
	}
}

// ID returns request builder for Event item
func (b *CalendarEventsCollectionRequestBuilder) ID(id string) *EventRequestBuilder {
	bb := &EventRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// CalendarEventsCollectionRequest is request for Event collection
type CalendarEventsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for Event collection
func (r *CalendarEventsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]Event, error) {
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
	var values []Event
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
			value  []Event
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
func (r *CalendarEventsCollectionRequest) NewRequest(method, url string, body io.Reader) (*http.Request, error) {
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

// GetN performs GET request for Event collection, max N pages
func (r *CalendarEventsCollectionRequest) GetN(ctx context.Context, n int) ([]Event, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for Event collection
func (r *CalendarEventsCollectionRequest) Get(ctx context.Context) ([]Event, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for Event collection
func (r *CalendarEventsCollectionRequest) Add(ctx context.Context, reqObj *Event) (resObj *Event, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// MultiValueExtendedProperties returns request builder for MultiValueLegacyExtendedProperty collection
func (b *CalendarRequestBuilder) MultiValueExtendedProperties() *CalendarMultiValueExtendedPropertiesCollectionRequestBuilder {
	bb := &CalendarMultiValueExtendedPropertiesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/multiValueExtendedProperties"
	return bb
}

// CalendarMultiValueExtendedPropertiesCollectionRequestBuilder is request builder for MultiValueLegacyExtendedProperty collection
type CalendarMultiValueExtendedPropertiesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for MultiValueLegacyExtendedProperty collection
func (b *CalendarMultiValueExtendedPropertiesCollectionRequestBuilder) Request() *CalendarMultiValueExtendedPropertiesCollectionRequest {
	return &CalendarMultiValueExtendedPropertiesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client,
			tenantID: b.tenantID, applicationID: b.applicationID, clientSecurityKey: b.clientSecurityKey},
	}
}

// ID returns request builder for MultiValueLegacyExtendedProperty item
func (b *CalendarMultiValueExtendedPropertiesCollectionRequestBuilder) ID(id string) *MultiValueLegacyExtendedPropertyRequestBuilder {
	bb := &MultiValueLegacyExtendedPropertyRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// CalendarMultiValueExtendedPropertiesCollectionRequest is request for MultiValueLegacyExtendedProperty collection
type CalendarMultiValueExtendedPropertiesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for MultiValueLegacyExtendedProperty collection
func (r *CalendarMultiValueExtendedPropertiesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]MultiValueLegacyExtendedProperty, error) {
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
	var values []MultiValueLegacyExtendedProperty
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
			value  []MultiValueLegacyExtendedProperty
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
func (r *CalendarMultiValueExtendedPropertiesCollectionRequest) NewRequest(method, url string, body io.Reader) (*http.Request, error) {
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

// GetN performs GET request for MultiValueLegacyExtendedProperty collection, max N pages
func (r *CalendarMultiValueExtendedPropertiesCollectionRequest) GetN(ctx context.Context, n int) ([]MultiValueLegacyExtendedProperty, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for MultiValueLegacyExtendedProperty collection
func (r *CalendarMultiValueExtendedPropertiesCollectionRequest) Get(ctx context.Context) ([]MultiValueLegacyExtendedProperty, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for MultiValueLegacyExtendedProperty collection
func (r *CalendarMultiValueExtendedPropertiesCollectionRequest) Add(ctx context.Context, reqObj *MultiValueLegacyExtendedProperty) (resObj *MultiValueLegacyExtendedProperty, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// SingleValueExtendedProperties returns request builder for SingleValueLegacyExtendedProperty collection
func (b *CalendarRequestBuilder) SingleValueExtendedProperties() *CalendarSingleValueExtendedPropertiesCollectionRequestBuilder {
	bb := &CalendarSingleValueExtendedPropertiesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/singleValueExtendedProperties"
	return bb
}

// CalendarSingleValueExtendedPropertiesCollectionRequestBuilder is request builder for SingleValueLegacyExtendedProperty collection
type CalendarSingleValueExtendedPropertiesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for SingleValueLegacyExtendedProperty collection
func (b *CalendarSingleValueExtendedPropertiesCollectionRequestBuilder) Request() *CalendarSingleValueExtendedPropertiesCollectionRequest {
	return &CalendarSingleValueExtendedPropertiesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client,
			tenantID: b.tenantID, applicationID: b.applicationID, clientSecurityKey: b.clientSecurityKey},
	}
}

// ID returns request builder for SingleValueLegacyExtendedProperty item
func (b *CalendarSingleValueExtendedPropertiesCollectionRequestBuilder) ID(id string) *SingleValueLegacyExtendedPropertyRequestBuilder {
	bb := &SingleValueLegacyExtendedPropertyRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// CalendarSingleValueExtendedPropertiesCollectionRequest is request for SingleValueLegacyExtendedProperty collection
type CalendarSingleValueExtendedPropertiesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for SingleValueLegacyExtendedProperty collection
func (r *CalendarSingleValueExtendedPropertiesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]SingleValueLegacyExtendedProperty, error) {
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
	var values []SingleValueLegacyExtendedProperty
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
			value  []SingleValueLegacyExtendedProperty
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
func (r *CalendarSingleValueExtendedPropertiesCollectionRequest) NewRequest(method, url string, body io.Reader) (*http.Request, error) {
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

// GetN performs GET request for SingleValueLegacyExtendedProperty collection, max N pages
func (r *CalendarSingleValueExtendedPropertiesCollectionRequest) GetN(ctx context.Context, n int) ([]SingleValueLegacyExtendedProperty, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for SingleValueLegacyExtendedProperty collection
func (r *CalendarSingleValueExtendedPropertiesCollectionRequest) Get(ctx context.Context) ([]SingleValueLegacyExtendedProperty, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for SingleValueLegacyExtendedProperty collection
func (r *CalendarSingleValueExtendedPropertiesCollectionRequest) Add(ctx context.Context, reqObj *SingleValueLegacyExtendedProperty) (resObj *SingleValueLegacyExtendedProperty, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Calendars returns request builder for Calendar collection
func (b *CalendarGroupRequestBuilder) Calendars() *CalendarGroupCalendarsCollectionRequestBuilder {
	bb := &CalendarGroupCalendarsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/calendars"
	return bb
}

// CalendarGroupCalendarsCollectionRequestBuilder is request builder for Calendar collection
type CalendarGroupCalendarsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for Calendar collection
func (b *CalendarGroupCalendarsCollectionRequestBuilder) Request() *CalendarGroupCalendarsCollectionRequest {
	return &CalendarGroupCalendarsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client,
			tenantID: b.tenantID, applicationID: b.applicationID, clientSecurityKey: b.clientSecurityKey},
	}
}

// ID returns request builder for Calendar item
func (b *CalendarGroupCalendarsCollectionRequestBuilder) ID(id string) *CalendarRequestBuilder {
	bb := &CalendarRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// CalendarGroupCalendarsCollectionRequest is request for Calendar collection
type CalendarGroupCalendarsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for Calendar collection
func (r *CalendarGroupCalendarsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]Calendar, error) {
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
	var values []Calendar
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
			value  []Calendar
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
func (r *CalendarGroupCalendarsCollectionRequest) NewRequest(method, url string, body io.Reader) (*http.Request, error) {
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

// GetN performs GET request for Calendar collection, max N pages
func (r *CalendarGroupCalendarsCollectionRequest) GetN(ctx context.Context, n int) ([]Calendar, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for Calendar collection
func (r *CalendarGroupCalendarsCollectionRequest) Get(ctx context.Context) ([]Calendar, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for Calendar collection
func (r *CalendarGroupCalendarsCollectionRequest) Add(ctx context.Context, reqObj *Calendar) (resObj *Calendar, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}
