// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// WorkbookTableRequestBuilder is request builder for WorkbookTable
type WorkbookTableRequestBuilder struct{ BaseRequestBuilder }

// Request returns WorkbookTableRequest
func (b *WorkbookTableRequestBuilder) Request() *WorkbookTableRequest {
	return &WorkbookTableRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// WorkbookTableRequest is request for WorkbookTable
type WorkbookTableRequest struct{ BaseRequest }

// Do performs HTTP request for WorkbookTable
func (r *WorkbookTableRequest) Do(method, path string, reqObj interface{}) (resObj *WorkbookTable, err error) {
	err = r.JSONRequestWithPath(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for WorkbookTable
func (r *WorkbookTableRequest) Get() (*WorkbookTable, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for WorkbookTable
func (r *WorkbookTableRequest) Update(reqObj *WorkbookTable) (*WorkbookTable, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for WorkbookTable
func (r *WorkbookTableRequest) Delete() error {
	return r.JSONRequestWithPath("DELETE", "", nil, nil)
}

// Columns returns request builder for WorkbookTableColumn collection
func (b *WorkbookTableRequestBuilder) Columns() *WorkbookTableColumnsCollectionRequestBuilder {
	bb := &WorkbookTableColumnsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/columns"
	return bb
}

// WorkbookTableColumnsCollectionRequestBuilder is request builder for WorkbookTableColumn collection
type WorkbookTableColumnsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for WorkbookTableColumn collection
func (b *WorkbookTableColumnsCollectionRequestBuilder) Request() *WorkbookTableColumnsCollectionRequest {
	return &WorkbookTableColumnsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for WorkbookTableColumn item
func (b *WorkbookTableColumnsCollectionRequestBuilder) ID(id string) *WorkbookTableColumnRequestBuilder {
	bb := &WorkbookTableColumnRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// WorkbookTableColumnsCollectionRequest is request for WorkbookTableColumn collection
type WorkbookTableColumnsCollectionRequest struct{ BaseRequest }

// Do performs HTTP request for WorkbookTableColumn collection
func (r *WorkbookTableColumnsCollectionRequest) Do(method, path string, reqObj interface{}) (resObj *WorkbookTableColumn, err error) {
	err = r.JSONRequestWithPath(method, path, reqObj, &resObj)
	return
}

// Paging perfoms paging operation for WorkbookTableColumn collection
func (r *WorkbookTableColumnsCollectionRequest) Paging(method, path string, obj interface{}) ([]WorkbookTableColumn, error) {
	req, err := r.NewJSONRequestWithPath(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []WorkbookTableColumn
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  []WorkbookTableColumn
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

// Get performs GET request for WorkbookTableColumn collection
func (r *WorkbookTableColumnsCollectionRequest) Get() ([]WorkbookTableColumn, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for WorkbookTableColumn collection
func (r *WorkbookTableColumnsCollectionRequest) Add(reqObj *WorkbookTableColumn) (*WorkbookTableColumn, error) {
	return r.Do("POST", "", reqObj)
}

// Rows returns request builder for WorkbookTableRow collection
func (b *WorkbookTableRequestBuilder) Rows() *WorkbookTableRowsCollectionRequestBuilder {
	bb := &WorkbookTableRowsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/rows"
	return bb
}

// WorkbookTableRowsCollectionRequestBuilder is request builder for WorkbookTableRow collection
type WorkbookTableRowsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for WorkbookTableRow collection
func (b *WorkbookTableRowsCollectionRequestBuilder) Request() *WorkbookTableRowsCollectionRequest {
	return &WorkbookTableRowsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for WorkbookTableRow item
func (b *WorkbookTableRowsCollectionRequestBuilder) ID(id string) *WorkbookTableRowRequestBuilder {
	bb := &WorkbookTableRowRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// WorkbookTableRowsCollectionRequest is request for WorkbookTableRow collection
type WorkbookTableRowsCollectionRequest struct{ BaseRequest }

// Do performs HTTP request for WorkbookTableRow collection
func (r *WorkbookTableRowsCollectionRequest) Do(method, path string, reqObj interface{}) (resObj *WorkbookTableRow, err error) {
	err = r.JSONRequestWithPath(method, path, reqObj, &resObj)
	return
}

// Paging perfoms paging operation for WorkbookTableRow collection
func (r *WorkbookTableRowsCollectionRequest) Paging(method, path string, obj interface{}) ([]WorkbookTableRow, error) {
	req, err := r.NewJSONRequestWithPath(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []WorkbookTableRow
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  []WorkbookTableRow
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

// Get performs GET request for WorkbookTableRow collection
func (r *WorkbookTableRowsCollectionRequest) Get() ([]WorkbookTableRow, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for WorkbookTableRow collection
func (r *WorkbookTableRowsCollectionRequest) Add(reqObj *WorkbookTableRow) (*WorkbookTableRow, error) {
	return r.Do("POST", "", reqObj)
}

// Sort is navigation property
func (b *WorkbookTableRequestBuilder) Sort() *WorkbookTableSortRequestBuilder {
	bb := &WorkbookTableSortRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/sort"
	return bb
}

// Worksheet is navigation property
func (b *WorkbookTableRequestBuilder) Worksheet() *WorkbookWorksheetRequestBuilder {
	bb := &WorkbookWorksheetRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/worksheet"
	return bb
}