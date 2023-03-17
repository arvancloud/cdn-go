package arvancloud

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type DNSRecord struct {
	ID               string      `json:"id,omitempty"`
	Type             string      `json:"type,omitempty"`
	Name             string      `json:"name,omitempty"`
	Value            interface{} `json:"value,omitempty"`
	TTL              int         `json:"ttl,omitempty"`
	Cloud            bool        `json:"cloud,omitempty"`
	UpstreamHTTPS    string      `json:"upstream_https,omitempty"`
	IPFilterMode     interface{} `json:"ip_filter_mode,omitempty"`
	IsProtected      bool        `json:"is_protected,omitempty"`
	CreatedAt        string      `json:"created_at,omitempty"`
	UpdatedAt        string      `json:"updated_at,omitempty"`
	MonitoringStatus string      `json:"monitoring_status,omitempty"`
	HealthCheck      interface{} `json:"health_check,omitempty"`
}

type CreateDNSRecordParams struct {
	Type          string      `json:"type,omitempty"`
	Name          string      `json:"name,omitempty"`
	Value         interface{} `json:"value,omitempty"`
	TTL           int         `json:"ttl,omitempty"`
	Cloud         bool        `json:"cloud,omitempty"`
	UpstreamHTTPS string      `json:"upstream_https,omitempty"`
	IPFilterMode  interface{} `json:"ip_filter_mode,omitempty"`
}

type DNSRecord_Value_A struct {
	IP      string `json:"ip,omitempty"`
	Port    int    `json:"port,omitempty"`
	Weight  int    `json:"weight,omitempty"`
	Country string `json:"country,omitempty"`
}

type DNSRecord_Value_AAAA = DNSRecord_Value_A

type DNSRecord_Value_MX struct {
	Host     string `json:"host,omitempty"`
	Priority int    `json:"priority,omitempty"`
}

type DNSRecord_Value_NS struct {
	Host string `json:"host,omitempty"`
}

type DNSRecord_Value_SRV struct {
	Target   string `json:"target,omitempty"`
	Port     int    `json:"port,omitempty"`
	Weight   int    `json:"weight,omitempty"`
	Priority int    `json:"priority,omitempty"`
}

type DNSRecord_Value_TXT struct {
	Text string `json:"text,omitempty"`
}

type DNSRecord_Value_SPF = DNSRecord_Value_TXT

type DNSRecord_Value_DKIM = DNSRecord_Value_TXT

type DNSRecord_Value_ANAME struct {
	Location   string `json:"location,omitempty"`
	HostHeader string `json:"host_header,omitempty"`
	Port       int    `json:"port,omitempty"`
}

type DNSRecord_Value_CNAME struct {
	Host       string `json:"host,omitempty"`
	HostHeader string `json:"host_header,omitempty"`
	Port       int    `json:"port,omitempty"`
}

type DNSRecord_Value_PTR struct {
	Domain string `json:"domain,omitempty"`
}

type DNSRecord_Value_TLSA struct {
	Usage        string `json:"usage,omitempty"`
	Selector     string `json:"selector,omitempty"`
	MatchingType string `json:"matching_type,omitempty"`
	Certificate  string `json:"certificate,omitempty"`
}

type DNSRecord_Value_CAA struct {
	Value string `json:"value,omitempty"`
	Tag   string `json:"tag,omitempty"`
}

type DNSRecord_IPFilterMode struct {
	Count     string `json:"count,omitempty"`
	Order     string `json:"order,omitempty"`
	GeoFilter string `json:"geo_filter,omitempty"`
}

type CreateDNSRecord_Response struct {
	Message string              `json:"message,omitempty"`
	Status  bool                `json:"status,omitempty"`
	Errors  map[string][]string `json:"errors,omitempty"`
	Data    interface{}         `json:"data,omitempty"`
}

type DNSRecord_Response struct {
	Data DNSRecord `json:"data,omitempty"`
}

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

var ErrMissingDNSRecordID = errors.New("required DNS record ID missing")

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
