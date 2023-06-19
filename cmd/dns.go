package main

import (
	"context"
	"log"
	"strconv"

	arvancloud "github.com/arvancloud/cdn-go/pkg"
	"github.com/gookit/gcli/v3/interact"
)

// List of functions for every DNS record type
var funcs = map[string]func() interface{}{
	"A":     createDNSRecord_A,
	"AAAA":  createDNSRecord_AAAA,
	"MX":    createDNSRecord_MX,
	"NS":    createDNSRecord_NS,
	"SRV":   createDNSRecord_SRV,
	"TXT":   createDNSRecord_TXT,
	"SPF":   createDNSRecord_SPF,
	"DKIM":  createDNSRecord_DKIM,
	"ANAME": createDNSRecord_ANAME,
	"CNAME": createDNSRecord_CNAME,
	"PTR":   createDNSRecord_PTR,
	"TLSA":  createDNSRecord_TLSA,
	"CAA":   createDNSRecord_CAA,
}

// GetDNSRecord will return a single DNS record
func GetDNSRecord(ctx context.Context, api *arvancloud.API, domain string, id string) {
	resource := arvancloud.Resource{
		Domain: domain,
	}

	u, err := api.GetDNSRecord(ctx, resource, id)
	if err != nil {
		log.Fatal(err.Error())
	}

	arvancloud.PrettyPrint(u)
}

// ListDNSRecords will list a part of DNS records
func ListDNSRecords(ctx context.Context, api *arvancloud.API, domain string, page int, perPage int) {
	param := arvancloud.ListDNSRecordsParams{
		Type:    "a",
		Page:    page,
		PerPage: perPage,
	}

	resource := arvancloud.Resource{
		Domain: domain,
	}

	u, err := api.ListDNSRecords(ctx, resource, param)
	if err != nil {
		log.Fatal(err.Error())
	}

	arvancloud.PrettyPrint(u)
}

// DeleteDNSRecord will delete a DNS record
func DeleteDNSRecord(ctx context.Context, api *arvancloud.API, domain string, id string) {
	resource := arvancloud.Resource{
		Domain: domain,
	}

	u, err := api.DeleteDNSRecord(ctx, resource, id)
	if err != nil {
		log.Fatal(err.Error())
	}

	arvancloud.PrettyPrint(u)
}

// CreateDNSRecord will create a DNS record
func CreateDNSRecord(ctx context.Context, api *arvancloud.API, domain string, record *arvancloud.CreateDNSRecordParams) {
	resource := arvancloud.Resource{
		Domain: domain,
	}

	u, err := api.CreateDNSRecord(ctx, resource, *record)
	if err != nil {
		log.Fatal(err.Error())
	}

	arvancloud.PrettyPrint(u)
}

// createDNSRecordParams will configure the parameters for creating a DNS record
func createDNSRecordParams(params *arvancloud.CreateDNSRecordParams) {
	recordName, _ := interact.ReadInput("Record name: ")
	params.Name = recordName

	recordType := interact.SelectOne(
		"Record type",
		[]string{
			"A",
			"AAAA",
			"MX",
			"NS",
			"SRV",
			"TXT",
			"SPF",
			"DKIM",
			"ANAME",
			"CNAME",
			"PTR",
			"TLSA",
			"CAA",
		},
		"",
		false,
	)
	params.Type = recordType

	ttl, _ := interact.ReadInput("Record TTL: ")
	recordTTL, err := strconv.Atoi(ttl)
	if err != nil {
		log.Fatal(err.Error())
	}
	params.TTL = recordTTL

	cloud := interact.SelectOne(
		"Record Cloud:",
		[]string{
			"True",
			"False",
		},
		"",
		false,
	)
	recordCloud, err := strconv.ParseBool(cloud)
	if err != nil {
		log.Fatal(err.Error())
	}
	params.Cloud = recordCloud

	recordUpstreamHTTPS := interact.SelectOne(
		"Record Upstream HTTPS:",
		[]string{
			"default",
			"auto",
			"http",
			"https",
		},
		"",
		false,
	)
	params.UpstreamHTTPS = recordUpstreamHTTPS

	// Call function based on record type
	if f, ok := funcs[recordType]; ok {
		params.Value = f()
	} else {
		log.Fatal("Unknown record type:", recordType)
	}

	params.IPFilterMode = createDNSRecord_IPFilterMode()
}

func createDNSRecord_A() interface{} {
	ip, _ := interact.ReadInput("Record IP: ")

	p, _ := interact.ReadInput("Record Port: ")
	port, err := strconv.Atoi(p)
	if err != nil {
		log.Fatal(err.Error())
	}

	w, _ := interact.ReadInput("Record Weight: ")
	weight, err := strconv.Atoi(w)
	if err != nil {
		log.Fatal(err.Error())
	}

	country, _ := interact.ReadInput("Record Country: ")

	return arvancloud.DNSRecord_Value_A{
		IP:      ip,
		Port:    port,
		Weight:  weight,
		Country: country,
	}
}

func createDNSRecord_AAAA() interface{} {
	ip, _ := interact.ReadInput("Record IP: ")

	p, _ := interact.ReadInput("Record Port: ")
	port, err := strconv.Atoi(p)
	if err != nil {
		log.Fatal(err.Error())
	}

	w, _ := interact.ReadInput("Record Weight: ")
	weight, err := strconv.Atoi(w)
	if err != nil {
		log.Fatal(err.Error())
	}

	country, _ := interact.ReadInput("Record Country: ")

	return arvancloud.DNSRecord_Value_AAAA{
		IP:      ip,
		Port:    port,
		Weight:  weight,
		Country: country,
	}
}

func createDNSRecord_MX() interface{} {
	host, _ := interact.ReadInput("Record Host: ")

	p, _ := interact.ReadInput("Record Priority: ")
	priority, err := strconv.Atoi(p)
	if err != nil {
		log.Fatal(err.Error())
	}

	return arvancloud.DNSRecord_Value_MX{
		Host:     host,
		Priority: priority,
	}
}

func createDNSRecord_NS() interface{} {
	host, _ := interact.ReadInput("Record Host: ")

	return arvancloud.DNSRecord_Value_NS{
		Host: host,
	}
}

func createDNSRecord_SRV() interface{} {
	target, _ := interact.ReadInput("Record Target: ")

	p, _ := interact.ReadInput("Record Port: ")
	port, err := strconv.Atoi(p)
	if err != nil {
		log.Fatal(err.Error())
	}

	w, _ := interact.ReadInput("Record Weight: ")
	weight, err := strconv.Atoi(w)
	if err != nil {
		log.Fatal(err.Error())
	}

	p, _ = interact.ReadInput("Record Priority: ")
	priority, err := strconv.Atoi(p)
	if err != nil {
		log.Fatal(err.Error())
	}

	return arvancloud.DNSRecord_Value_SRV{
		Target:   target,
		Port:     port,
		Weight:   weight,
		Priority: priority,
	}
}

func createDNSRecord_TXT() interface{} {
	text, _ := interact.ReadInput("Record Text: ")

	return arvancloud.DNSRecord_Value_TXT{
		Text: text,
	}
}

func createDNSRecord_SPF() interface{} {
	text, _ := interact.ReadInput("Record Text: ")

	return arvancloud.DNSRecord_Value_SPF{
		Text: text,
	}
}

func createDNSRecord_DKIM() interface{} {
	text, _ := interact.ReadInput("Record Text: ")

	return arvancloud.DNSRecord_Value_DKIM{
		Text: text,
	}
}

func createDNSRecord_ANAME() interface{} {
	location, _ := interact.ReadInput("Record Location: ")

	host_header := interact.SelectOne(
		"Record Host header:",
		[]string{
			"source",
			"dest",
		},
		"",
		false,
	)

	p, _ := interact.ReadInput("Record Port: ")
	port, err := strconv.Atoi(p)
	if err != nil {
		log.Fatal(err.Error())
	}

	return arvancloud.DNSRecord_Value_ANAME{
		Location:   location,
		HostHeader: host_header,
		Port:       port,
	}
}

func createDNSRecord_CNAME() interface{} {
	host, _ := interact.ReadInput("Record Host: ")

	host_header := interact.SelectOne(
		"Record Host header:",
		[]string{
			"source",
			"dest",
		},
		"",
		false,
	)

	p, _ := interact.ReadInput("Record Port: ")
	port, err := strconv.Atoi(p)
	if err != nil {
		log.Fatal(err.Error())
	}

	return arvancloud.DNSRecord_Value_CNAME{
		Host:       host,
		HostHeader: host_header,
		Port:       port,
	}
}

func createDNSRecord_PTR() interface{} {
	domain, _ := interact.ReadInput("Record Domain: ")

	return arvancloud.DNSRecord_Value_PTR{
		Domain: domain,
	}
}

func createDNSRecord_TLSA() interface{} {
	usage, _ := interact.ReadInput("Record Usage: ")
	selector, _ := interact.ReadInput("Record Selector: ")
	matching_type, _ := interact.ReadInput("Record Matching type: ")
	certificate, _ := interact.ReadInput("Record Certificate: ")

	return arvancloud.DNSRecord_Value_TLSA{
		Usage:        usage,
		Selector:     selector,
		MatchingType: matching_type,
		Certificate:  certificate,
	}
}

func createDNSRecord_CAA() interface{} {
	value, _ := interact.ReadInput("Record Value: ")

	tag := interact.SelectOne(
		"Record Tag:",
		[]string{
			"issuewild",
			"issue",
			"iodef",
		},
		"",
		false,
	)

	return arvancloud.DNSRecord_Value_CAA{
		Value: value,
		Tag:   tag,
	}
}

func createDNSRecord_IPFilterMode() interface{} {
	count := interact.SelectOne(
		"Record Count:",
		[]string{
			"single",
			"multi",
		},
		"",
		false,
	)

	order := interact.SelectOne(
		"Record Order:",
		[]string{
			"none",
			"weighted",
			"rr",
		},
		"",
		false,
	)

	geo_filter := interact.SelectOne(
		"Record Geo filter:",
		[]string{
			"none",
			"location",
			"country",
		},
		"",
		false,
	)

	return arvancloud.DNSRecord_IPFilterMode{
		Count:     count,
		Order:     order,
		GeoFilter: geo_filter,
	}
}
