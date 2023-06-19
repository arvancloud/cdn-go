package arvancloud

// DSNRecord is a DSN record structure for a domain
type DNSRecord struct {
	ID            string      `json:"id,omitempty"`
	Type          string      `json:"type,omitempty"`
	Name          string      `json:"name,omitempty"`
	Value         interface{} `json:"value,omitempty"`
	TTL           int         `json:"ttl,omitempty"`
	Cloud         bool        `json:"cloud,omitempty"`
	UpstreamHTTPS string      `json:"upstream_https,omitempty"`
	IPFilterMode  interface{} `json:"ip_filter_mode,omitempty"`
	IsProtected   bool        `json:"is_protected,omitempty"`
	CreatedAt     string      `json:"created_at,omitempty"`
	UpdatedAt     string      `json:"updated_at,omitempty"`
}

// DNSRecord_Response is response structure contains
// DSN record's data
type DNSRecord_Response struct {
	Data DNSRecord `json:"data,omitempty"`
}

// ----------------------------------------------------------- DSN Records

// DNSRecord_Value_A is a structure for A record
type DNSRecord_Value_A struct {
	IP      string `json:"ip,omitempty"`
	Port    int    `json:"port,omitempty"`
	Weight  int    `json:"weight,omitempty"`
	Country string `json:"country,omitempty"`
}

// DNSRecord_Value_AAAA is a structure for AAAA record
type DNSRecord_Value_AAAA = DNSRecord_Value_A

// DNSRecord_Value_MX is a structure for MX record
type DNSRecord_Value_MX struct {
	Host     string `json:"host,omitempty"`
	Priority int    `json:"priority,omitempty"`
}

// DNSRecord_Value_NS is a structure for NS record
type DNSRecord_Value_NS struct {
	Host string `json:"host,omitempty"`
}

// DNSRecord_Value_SRV is a structure for SRV record
type DNSRecord_Value_SRV struct {
	Target   string `json:"target,omitempty"`
	Port     int    `json:"port,omitempty"`
	Weight   int    `json:"weight,omitempty"`
	Priority int    `json:"priority,omitempty"`
}

// DNSRecord_Value_TXT is a structure for TXT record
type DNSRecord_Value_TXT struct {
	Text string `json:"text,omitempty"`
}

// DNSRecord_Value_SPF is a structure for SPF record
type DNSRecord_Value_SPF = DNSRecord_Value_TXT

// DNSRecord_Value_DKIM is a structure for DKIM record
type DNSRecord_Value_DKIM = DNSRecord_Value_TXT

// DNSRecord_Value_ANAME is a structure for ANAME record
type DNSRecord_Value_ANAME struct {
	Location   string `json:"location,omitempty"`
	HostHeader string `json:"host_header,omitempty"`
	Port       int    `json:"port,omitempty"`
}

// DNSRecord_Value_CNAME is a structure for CNAME record
type DNSRecord_Value_CNAME struct {
	Host       string `json:"host,omitempty"`
	HostHeader string `json:"host_header,omitempty"`
	Port       int    `json:"port,omitempty"`
}

// DNSRecord_Value_PTR is a structure for PTR record
type DNSRecord_Value_PTR struct {
	Domain string `json:"domain,omitempty"`
}

// DNSRecord_Value_TLSA is a structure for TLSA record
type DNSRecord_Value_TLSA struct {
	Usage        string `json:"usage,omitempty"`
	Selector     string `json:"selector,omitempty"`
	MatchingType string `json:"matching_type,omitempty"`
	Certificate  string `json:"certificate,omitempty"`
}

// DNSRecord_Value_CAA is a structure for CAA record
type DNSRecord_Value_CAA struct {
	Value string `json:"value,omitempty"`
	Tag   string `json:"tag,omitempty"`
}

// ----------------------------------------------------------- DSN Operations - Create

// CreateDNSRecordParams is a structure for all needed parameters
// to create DNS record
type CreateDNSRecordParams struct {
	Type          string      `json:"type,omitempty"`
	Name          string      `json:"name,omitempty"`
	Value         interface{} `json:"value,omitempty"`
	TTL           int         `json:"ttl,omitempty"`
	Cloud         bool        `json:"cloud,omitempty"`
	UpstreamHTTPS string      `json:"upstream_https,omitempty"`
	IPFilterMode  interface{} `json:"ip_filter_mode,omitempty"`
}

// DNSRecord_IPFilterMode is a structure for IP Filter Mode when
// creating a DNS record
type DNSRecord_IPFilterMode struct {
	Count     string `json:"count,omitempty"`
	Order     string `json:"order,omitempty"`
	GeoFilter string `json:"geo_filter,omitempty"`
}

// CreateDNSRecord_Response is a response structure when creating
// a DNS record of a domain
type CreateDNSRecord_Response struct {
	Message string              `json:"message,omitempty"`
	Status  bool                `json:"status,omitempty"`
	Errors  map[string][]string `json:"errors,omitempty"`
	Data    interface{}         `json:"data,omitempty"`
}

// ----------------------------------------------------------- DSN Operations - List

// ListDNSRecordsParams is a structure for all needed parameters
// to list all DNS records of a domain
type ListDNSRecordsParams struct {
	Search  string `url:"search,omitempty"`
	Type    string `url:"type,omitempty"`
	Page    int    `url:"page,omitempty"`
	PerPage int    `url:"per_page,omitempty"`
}

// ListDNSRecord_Response is a response structure when listing
// DNS records of a domain
type ListDNSRecord_Response struct {
	Data  []DNSRecord `json:"data"`
	Meta  interface{} `json:"meta,omitempty"`
	Links interface{} `json:"links,omitempty"`
}

// ----------------------------------------------------------- DSN Operations - Update

type UpdateDNSRecordParams = CreateDNSRecordParams

// UpdateDNSRecord_Response is a response structure when updating
// a DNS record of a domain
type UpdateDNSRecord_Response = CreateDNSRecord_Response

// ----------------------------------------------------------- DSN Operations - Delete

// DeleteDNSRecord_Response is a response structure when deleting
// a DNS record of a domain
type DeleteDNSRecord_Response = CreateDNSRecord_Response
