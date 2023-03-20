package main

import (
	"context"
	"log"

	arvancloud "github.com/arvancloud/cdn-go/pkg"
)

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
