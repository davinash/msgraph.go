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

// SalesInvoiceCancelAndSendRequestParameter undocumented
type SalesInvoiceCancelAndSendRequestParameter struct {
}

// SalesInvoiceCancelRequestParameter undocumented
type SalesInvoiceCancelRequestParameter struct {
}

// SalesInvoicePostAndSendRequestParameter undocumented
type SalesInvoicePostAndSendRequestParameter struct {
}

// SalesInvoicePostRequestParameter undocumented
type SalesInvoicePostRequestParameter struct {
}

// SalesInvoiceSendRequestParameter undocumented
type SalesInvoiceSendRequestParameter struct {
}

// SalesQuoteMakeInvoiceRequestParameter undocumented
type SalesQuoteMakeInvoiceRequestParameter struct {
}

// SalesQuoteSendRequestParameter undocumented
type SalesQuoteSendRequestParameter struct {
}

// Currency is navigation property
func (b *SalesCreditMemoRequestBuilder) Currency() *CurrencyRequestBuilder {
	bb := &CurrencyRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/currency"
	return bb
}

// Customer is navigation property
func (b *SalesCreditMemoRequestBuilder) Customer() *CustomerRequestBuilder {
	bb := &CustomerRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/customer"
	return bb
}

// PaymentTerm is navigation property
func (b *SalesCreditMemoRequestBuilder) PaymentTerm() *PaymentTermRequestBuilder {
	bb := &PaymentTermRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/paymentTerm"
	return bb
}

// SalesCreditMemoLines returns request builder for SalesCreditMemoLine collection
func (b *SalesCreditMemoRequestBuilder) SalesCreditMemoLines() *SalesCreditMemoSalesCreditMemoLinesCollectionRequestBuilder {
	bb := &SalesCreditMemoSalesCreditMemoLinesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/salesCreditMemoLines"
	return bb
}

// SalesCreditMemoSalesCreditMemoLinesCollectionRequestBuilder is request builder for SalesCreditMemoLine collection
type SalesCreditMemoSalesCreditMemoLinesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for SalesCreditMemoLine collection
func (b *SalesCreditMemoSalesCreditMemoLinesCollectionRequestBuilder) Request() *SalesCreditMemoSalesCreditMemoLinesCollectionRequest {
	return &SalesCreditMemoSalesCreditMemoLinesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client,
			tenantID: b.tenantID, applicationID: b.applicationID, clientSecurityKey: b.clientSecurityKey},
	}
}

// ID returns request builder for SalesCreditMemoLine item
func (b *SalesCreditMemoSalesCreditMemoLinesCollectionRequestBuilder) ID(id string) *SalesCreditMemoLineRequestBuilder {
	bb := &SalesCreditMemoLineRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// SalesCreditMemoSalesCreditMemoLinesCollectionRequest is request for SalesCreditMemoLine collection
type SalesCreditMemoSalesCreditMemoLinesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for SalesCreditMemoLine collection
func (r *SalesCreditMemoSalesCreditMemoLinesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]SalesCreditMemoLine, error) {
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
	var values []SalesCreditMemoLine
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
			value  []SalesCreditMemoLine
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
func (r *SalesCreditMemoSalesCreditMemoLinesCollectionRequest) NewRequest(method, url string, body io.Reader) (*http.Request, error) {
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

// GetN performs GET request for SalesCreditMemoLine collection, max N pages
func (r *SalesCreditMemoSalesCreditMemoLinesCollectionRequest) GetN(ctx context.Context, n int) ([]SalesCreditMemoLine, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for SalesCreditMemoLine collection
func (r *SalesCreditMemoSalesCreditMemoLinesCollectionRequest) Get(ctx context.Context) ([]SalesCreditMemoLine, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for SalesCreditMemoLine collection
func (r *SalesCreditMemoSalesCreditMemoLinesCollectionRequest) Add(ctx context.Context, reqObj *SalesCreditMemoLine) (resObj *SalesCreditMemoLine, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Account is navigation property
func (b *SalesCreditMemoLineRequestBuilder) Account() *AccountRequestBuilder {
	bb := &AccountRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/account"
	return bb
}

// Item is navigation property
func (b *SalesCreditMemoLineRequestBuilder) Item() *ItemRequestBuilder {
	bb := &ItemRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/item"
	return bb
}

// Currency is navigation property
func (b *SalesInvoiceRequestBuilder) Currency() *CurrencyRequestBuilder {
	bb := &CurrencyRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/currency"
	return bb
}

// Customer is navigation property
func (b *SalesInvoiceRequestBuilder) Customer() *CustomerRequestBuilder {
	bb := &CustomerRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/customer"
	return bb
}

// PaymentTerm is navigation property
func (b *SalesInvoiceRequestBuilder) PaymentTerm() *PaymentTermRequestBuilder {
	bb := &PaymentTermRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/paymentTerm"
	return bb
}

// SalesInvoiceLines returns request builder for SalesInvoiceLine collection
func (b *SalesInvoiceRequestBuilder) SalesInvoiceLines() *SalesInvoiceSalesInvoiceLinesCollectionRequestBuilder {
	bb := &SalesInvoiceSalesInvoiceLinesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/salesInvoiceLines"
	return bb
}

// SalesInvoiceSalesInvoiceLinesCollectionRequestBuilder is request builder for SalesInvoiceLine collection
type SalesInvoiceSalesInvoiceLinesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for SalesInvoiceLine collection
func (b *SalesInvoiceSalesInvoiceLinesCollectionRequestBuilder) Request() *SalesInvoiceSalesInvoiceLinesCollectionRequest {
	return &SalesInvoiceSalesInvoiceLinesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client,
			tenantID: b.tenantID, applicationID: b.applicationID, clientSecurityKey: b.clientSecurityKey},
	}
}

// ID returns request builder for SalesInvoiceLine item
func (b *SalesInvoiceSalesInvoiceLinesCollectionRequestBuilder) ID(id string) *SalesInvoiceLineRequestBuilder {
	bb := &SalesInvoiceLineRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// SalesInvoiceSalesInvoiceLinesCollectionRequest is request for SalesInvoiceLine collection
type SalesInvoiceSalesInvoiceLinesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for SalesInvoiceLine collection
func (r *SalesInvoiceSalesInvoiceLinesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]SalesInvoiceLine, error) {
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
	var values []SalesInvoiceLine
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
			value  []SalesInvoiceLine
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
func (r *SalesInvoiceSalesInvoiceLinesCollectionRequest) NewRequest(method, url string, body io.Reader) (*http.Request, error) {
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

// GetN performs GET request for SalesInvoiceLine collection, max N pages
func (r *SalesInvoiceSalesInvoiceLinesCollectionRequest) GetN(ctx context.Context, n int) ([]SalesInvoiceLine, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for SalesInvoiceLine collection
func (r *SalesInvoiceSalesInvoiceLinesCollectionRequest) Get(ctx context.Context) ([]SalesInvoiceLine, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for SalesInvoiceLine collection
func (r *SalesInvoiceSalesInvoiceLinesCollectionRequest) Add(ctx context.Context, reqObj *SalesInvoiceLine) (resObj *SalesInvoiceLine, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// ShipmentMethod is navigation property
func (b *SalesInvoiceRequestBuilder) ShipmentMethod() *ShipmentMethodRequestBuilder {
	bb := &ShipmentMethodRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/shipmentMethod"
	return bb
}

// Account is navigation property
func (b *SalesInvoiceLineRequestBuilder) Account() *AccountRequestBuilder {
	bb := &AccountRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/account"
	return bb
}

// Item is navigation property
func (b *SalesInvoiceLineRequestBuilder) Item() *ItemRequestBuilder {
	bb := &ItemRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/item"
	return bb
}

// Currency is navigation property
func (b *SalesOrderRequestBuilder) Currency() *CurrencyRequestBuilder {
	bb := &CurrencyRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/currency"
	return bb
}

// Customer is navigation property
func (b *SalesOrderRequestBuilder) Customer() *CustomerRequestBuilder {
	bb := &CustomerRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/customer"
	return bb
}

// PaymentTerm is navigation property
func (b *SalesOrderRequestBuilder) PaymentTerm() *PaymentTermRequestBuilder {
	bb := &PaymentTermRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/paymentTerm"
	return bb
}

// SalesOrderLines returns request builder for SalesOrderLine collection
func (b *SalesOrderRequestBuilder) SalesOrderLines() *SalesOrderSalesOrderLinesCollectionRequestBuilder {
	bb := &SalesOrderSalesOrderLinesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/salesOrderLines"
	return bb
}

// SalesOrderSalesOrderLinesCollectionRequestBuilder is request builder for SalesOrderLine collection
type SalesOrderSalesOrderLinesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for SalesOrderLine collection
func (b *SalesOrderSalesOrderLinesCollectionRequestBuilder) Request() *SalesOrderSalesOrderLinesCollectionRequest {
	return &SalesOrderSalesOrderLinesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client,
			tenantID: b.tenantID, applicationID: b.applicationID, clientSecurityKey: b.clientSecurityKey},
	}
}

// ID returns request builder for SalesOrderLine item
func (b *SalesOrderSalesOrderLinesCollectionRequestBuilder) ID(id string) *SalesOrderLineRequestBuilder {
	bb := &SalesOrderLineRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// SalesOrderSalesOrderLinesCollectionRequest is request for SalesOrderLine collection
type SalesOrderSalesOrderLinesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for SalesOrderLine collection
func (r *SalesOrderSalesOrderLinesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]SalesOrderLine, error) {
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
	var values []SalesOrderLine
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
			value  []SalesOrderLine
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
func (r *SalesOrderSalesOrderLinesCollectionRequest) NewRequest(method, url string, body io.Reader) (*http.Request, error) {
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

// GetN performs GET request for SalesOrderLine collection, max N pages
func (r *SalesOrderSalesOrderLinesCollectionRequest) GetN(ctx context.Context, n int) ([]SalesOrderLine, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for SalesOrderLine collection
func (r *SalesOrderSalesOrderLinesCollectionRequest) Get(ctx context.Context) ([]SalesOrderLine, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for SalesOrderLine collection
func (r *SalesOrderSalesOrderLinesCollectionRequest) Add(ctx context.Context, reqObj *SalesOrderLine) (resObj *SalesOrderLine, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Account is navigation property
func (b *SalesOrderLineRequestBuilder) Account() *AccountRequestBuilder {
	bb := &AccountRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/account"
	return bb
}

// Item is navigation property
func (b *SalesOrderLineRequestBuilder) Item() *ItemRequestBuilder {
	bb := &ItemRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/item"
	return bb
}

// Currency is navigation property
func (b *SalesQuoteRequestBuilder) Currency() *CurrencyRequestBuilder {
	bb := &CurrencyRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/currency"
	return bb
}

// Customer is navigation property
func (b *SalesQuoteRequestBuilder) Customer() *CustomerRequestBuilder {
	bb := &CustomerRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/customer"
	return bb
}

// PaymentTerm is navigation property
func (b *SalesQuoteRequestBuilder) PaymentTerm() *PaymentTermRequestBuilder {
	bb := &PaymentTermRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/paymentTerm"
	return bb
}

// SalesQuoteLines returns request builder for SalesQuoteLine collection
func (b *SalesQuoteRequestBuilder) SalesQuoteLines() *SalesQuoteSalesQuoteLinesCollectionRequestBuilder {
	bb := &SalesQuoteSalesQuoteLinesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/salesQuoteLines"
	return bb
}

// SalesQuoteSalesQuoteLinesCollectionRequestBuilder is request builder for SalesQuoteLine collection
type SalesQuoteSalesQuoteLinesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for SalesQuoteLine collection
func (b *SalesQuoteSalesQuoteLinesCollectionRequestBuilder) Request() *SalesQuoteSalesQuoteLinesCollectionRequest {
	return &SalesQuoteSalesQuoteLinesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client,
			tenantID: b.tenantID, applicationID: b.applicationID, clientSecurityKey: b.clientSecurityKey},
	}
}

// ID returns request builder for SalesQuoteLine item
func (b *SalesQuoteSalesQuoteLinesCollectionRequestBuilder) ID(id string) *SalesQuoteLineRequestBuilder {
	bb := &SalesQuoteLineRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// SalesQuoteSalesQuoteLinesCollectionRequest is request for SalesQuoteLine collection
type SalesQuoteSalesQuoteLinesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for SalesQuoteLine collection
func (r *SalesQuoteSalesQuoteLinesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]SalesQuoteLine, error) {
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
	var values []SalesQuoteLine
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
			value  []SalesQuoteLine
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
func (r *SalesQuoteSalesQuoteLinesCollectionRequest) NewRequest(method, url string, body io.Reader) (*http.Request, error) {
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

// GetN performs GET request for SalesQuoteLine collection, max N pages
func (r *SalesQuoteSalesQuoteLinesCollectionRequest) GetN(ctx context.Context, n int) ([]SalesQuoteLine, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for SalesQuoteLine collection
func (r *SalesQuoteSalesQuoteLinesCollectionRequest) Get(ctx context.Context) ([]SalesQuoteLine, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for SalesQuoteLine collection
func (r *SalesQuoteSalesQuoteLinesCollectionRequest) Add(ctx context.Context, reqObj *SalesQuoteLine) (resObj *SalesQuoteLine, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// ShipmentMethod is navigation property
func (b *SalesQuoteRequestBuilder) ShipmentMethod() *ShipmentMethodRequestBuilder {
	bb := &ShipmentMethodRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/shipmentMethod"
	return bb
}

// Account is navigation property
func (b *SalesQuoteLineRequestBuilder) Account() *AccountRequestBuilder {
	bb := &AccountRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/account"
	return bb
}

// Item is navigation property
func (b *SalesQuoteLineRequestBuilder) Item() *ItemRequestBuilder {
	bb := &ItemRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/item"
	return bb
}
