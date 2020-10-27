// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// ActivityHistoryItemRequestBuilder is request builder for ActivityHistoryItem
type ActivityHistoryItemRequestBuilder struct{ BaseRequestBuilder }

// Request returns ActivityHistoryItemRequest
func (b *ActivityHistoryItemRequestBuilder) Request() *ActivityHistoryItemRequest {
	return &ActivityHistoryItemRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client,
			tenantID: b.tenantID, applicationID: b.applicationID, clientSecurityKey: b.clientSecurityKey},
	}
}

// ActivityHistoryItemRequest is request for ActivityHistoryItem
type ActivityHistoryItemRequest struct{ BaseRequest }

// Get performs GET request for ActivityHistoryItem
func (r *ActivityHistoryItemRequest) Get(ctx context.Context) (resObj *ActivityHistoryItem, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ActivityHistoryItem
func (r *ActivityHistoryItemRequest) Update(ctx context.Context, reqObj *ActivityHistoryItem) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ActivityHistoryItem
func (r *ActivityHistoryItemRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// ActivityStatisticsRequestBuilder is request builder for ActivityStatistics
type ActivityStatisticsRequestBuilder struct{ BaseRequestBuilder }

// Request returns ActivityStatisticsRequest
func (b *ActivityStatisticsRequestBuilder) Request() *ActivityStatisticsRequest {
	return &ActivityStatisticsRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client,
			tenantID: b.tenantID, applicationID: b.applicationID, clientSecurityKey: b.clientSecurityKey},
	}
}

// ActivityStatisticsRequest is request for ActivityStatistics
type ActivityStatisticsRequest struct{ BaseRequest }

// Get performs GET request for ActivityStatistics
func (r *ActivityStatisticsRequest) Get(ctx context.Context) (resObj *ActivityStatistics, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ActivityStatistics
func (r *ActivityStatisticsRequest) Update(ctx context.Context, reqObj *ActivityStatistics) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ActivityStatistics
func (r *ActivityStatisticsRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
