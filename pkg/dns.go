package arvancloud

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// CreateDNSRecord will create a DNS record
// ? Documentation: https://www.arvancloud.ir/api/cdn/4.0#tag/DNS-Management/operation/dns_records.store
func (api *API) CreateDNSRecord(ctx context.Context, resource Resource, record CreateDNSRecordParams) (res *CreateDNSRecord_Response, err error) {
	if resource.Domain == "" {
		return nil, ErrMissingDomain
	}

	uri := fmt.Sprintf("/domains/%s/dns-records", resource.Domain)
	response, err := api.makeRequest(ctx, http.MethodPost, uri, record)
	if err != nil {
		return nil, err
	}

	res = &CreateDNSRecord_Response{}
	err = json.Unmarshal(response, &res)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return res, nil
}

// GetDNSRecord will return a single DNS record
// ? Documentation: https://www.arvancloud.ir/api/cdn/4.0#tag/DNS-Management/operation/dns_records.show
func (api *API) GetDNSRecord(ctx context.Context, resource Resource, recordID string) (*DNSRecord, error) {
	if resource.Domain == "" {
		return nil, ErrMissingDomain
	}

	if recordID == "" {
		return nil, ErrMissingDNSRecordID
	}

	uri := fmt.Sprintf("/domains/%s/dns-records/%s", resource.Domain, recordID)
	response, err := api.makeRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}

	res := &DNSRecord_Response{}
	err = json.Unmarshal(response, &res)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return &res.Data, nil
}

// ListDNSRecords will list a part of DNS records
// ? Documentation: https://www.arvancloud.ir/api/cdn/4.0#tag/DNS-Management/operation/dns_records.list
func (api *API) ListDNSRecords(ctx context.Context, resource Resource, params ListDNSRecordsParams) ([]DNSRecord, error) {
	if resource.Domain == "" {
		return nil, ErrMissingDomain
	}

	if params.Page < 1 {
		params.Page = 1
	}

	uri := buildURI(fmt.Sprintf("/domains/%s/dns-records", resource.Domain), params)
	res, err := api.makeRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}

	var listResponse ListDNSRecord_Response
	err = json.Unmarshal(res, &listResponse)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return listResponse.Data, nil
}

// UpdateDNSRecord will update a DNS record
// ? Documentation: https://www.arvancloud.ir/api/cdn/4.0#tag/DNS-Management/operation/dns_records.update
func (api *API) UpdateDNSRecord(ctx context.Context, resource Resource, recordID string, params UpdateDNSRecordParams) (*UpdateDNSRecord_Response, error) {
	if resource.Domain == "" {
		return nil, ErrMissingDomain
	}

	if recordID == "" {
		return nil, ErrMissingDNSRecordID
	}

	uri := fmt.Sprintf("/domains/%s/dns-records/%s", resource.Domain, recordID)
	response, err := api.makeRequest(ctx, http.MethodPut, uri, params)
	if err != nil {
		return nil, err
	}

	res := &UpdateDNSRecord_Response{}
	err = json.Unmarshal(response, &res)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return res, nil
}

// DeleteDNSRecord will delete a DNS record
// ? Documentation: https://www.arvancloud.ir/api/cdn/4.0#tag/DNS-Management/operation/dns_records.remove
func (api *API) DeleteDNSRecord(ctx context.Context, resource Resource, recordID string) (*DeleteDNSRecord_Response, error) {
	if resource.Domain == "" {
		return nil, ErrMissingDomain
	}

	if recordID == "" {
		return nil, ErrMissingDNSRecordID
	}

	uri := fmt.Sprintf("/domains/%s/dns-records/%s", resource.Domain, recordID)
	response, err := api.makeRequest(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return nil, err
	}

	res := &DeleteDNSRecord_Response{}
	err = json.Unmarshal(response, &res)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return res, nil
}
