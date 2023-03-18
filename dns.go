package arvancloud

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

func (api *API) CreateDNSRecord(ctx context.Context, rc ResourceContainer, record CreateDNSRecordParams) (*CreateDNSRecord_Response, error) {
	if rc.Domain == "" {
		return nil, ErrMissingDomain
	}

	uri := fmt.Sprintf("/domains/%s/dns-records", rc.Domain)
	response, err := api.makeRequestContext(ctx, http.MethodPost, uri, record)
	if err != nil {
		return nil, err
	}

	res := &CreateDNSRecord_Response{}
	err = json.Unmarshal(response, &res)
	if err != nil {
		return nil, fmt.Errorf(errUnmarshalError+": %w", err)
	}

	return res, nil
}

func (api *API) GetDNSRecord(ctx context.Context, rc ResourceContainer, recordID string) (*DNSRecord, error) {
	if rc.Domain == "" {
		return nil, ErrMissingDomain
	}

	if recordID == "" {
		return nil, ErrMissingDNSRecordID
	}

	uri := fmt.Sprintf("/domains/%s/dns-records/%s", rc.Domain, recordID)
	response, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}

	res := &DNSRecord_Response{}
	err = json.Unmarshal(response, &res)
	if err != nil {
		return nil, fmt.Errorf(errUnmarshalError+": %w", err)
	}

	return &res.Data, nil
}

func (api *API) ListDNSRecords(ctx context.Context, rc ResourceContainer, params ListDNSRecordsParams) ([]DNSRecord, error) {
	if rc.Domain == "" {
		return nil, ErrMissingDomain
	}

	if params.Page < 1 {
		params.Page = 1
	}

	uri := buildURI(fmt.Sprintf("/domains/%s/dns-records", rc.Domain), params)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
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

func (api *API) UpdateDNSRecord(ctx context.Context, rc ResourceContainer, recordID string, params UpdateDNSRecordParams) (*UpdateDNSRecord_Response, error) {
	if rc.Domain == "" {
		return nil, ErrMissingDomain
	}

	if recordID == "" {
		return nil, ErrMissingDNSRecordID
	}

	uri := fmt.Sprintf("/domains/%s/dns-records/%s", rc.Domain, recordID)
	response, err := api.makeRequestContext(ctx, http.MethodPut, uri, params)
	if err != nil {
		return nil, err
	}

	res := &UpdateDNSRecord_Response{}
	err = json.Unmarshal(response, &res)
	if err != nil {
		return nil, fmt.Errorf(errUnmarshalError+": %w", err)
	}

	return res, nil
}

func (api *API) DeleteDNSRecord(ctx context.Context, rc ResourceContainer, recordID string) (*DeleteDNSRecord_Response, error) {
	if rc.Domain == "" {
		return nil, ErrMissingDomain
	}

	if recordID == "" {
		return nil, ErrMissingDNSRecordID
	}

	uri := fmt.Sprintf("/domains/%s/dns-records/%s", rc.Domain, recordID)
	response, err := api.makeRequestContext(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return nil, err
	}

	res := &DeleteDNSRecord_Response{}
	err = json.Unmarshal(response, &res)
	if err != nil {
		return nil, fmt.Errorf(errUnmarshalError+": %w", err)
	}

	return res, nil
}
