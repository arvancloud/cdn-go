package arvancloud

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateDNSRecord(t *testing.T) {
	setup()
	defer teardown()

	input := CreateDNSRecordParams{
		Type: "a",
		Name: "@",
		Value: []interface{}{
			map[string]interface{}{
				"ip":      "1.2.3.4",
				"port":    1.0,
				"weight":  10.0,
				"country": "",
			},
			map[string]interface{}{
				"ip":      "5.6.7.8",
				"port":    2.0,
				"weight":  20.0,
				"country": "",
			},
		},
		TTL:           120,
		UpstreamHTTPS: "https",
		IPFilterMode: map[string]interface{}{
			"count":      "multi",
			"order":      "weighted",
			"geo_filter": "none",
		},
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method, "Expected method 'POST', got %s", r.Method)

		var p CreateDNSRecordParams
		err := json.NewDecoder(r.Body).Decode(&p)
		require.NoError(t, err)
		assert.Equal(t, input, p)

		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, `{
			"data": {
				"id": "714009ff-a43c-43c5-80e2-0b3ffc1344a4",
				"type": "a",
				"name": "@",
				"value": [
					{
						"ip": "1.2.3.4",
						"port": 1,
						"weight": 10,
						"country": ""
					},
					{
						"ip": "5.6.7.8",
						"port": 2,
						"weight": 20,
						"country": ""
					}
				],
				"ttl": 120,
				"cloud": false,
				"upstream_https": "https",
				"ip_filter_mode": {
					"count": "multi",
					"order": "weighted",
					"geo_filter": "none"
				},
				"is_protected": false,
				"created_at": "2023-03-31T08:57:51+00:00",
				"updated_at": "2023-03-31T08:57:51+00:00"
			},
			"message": "DNS record created successfully"
		}`)
	}

	mux.HandleFunc("/domains/"+testDomain+"/dns-records", handler)

	want := &CreateDNSRecord_Response{
		Data: map[string]interface{}{
			"id":             "714009ff-a43c-43c5-80e2-0b3ffc1344a4",
			"type":           input.Type,
			"name":           input.Name,
			"value":          input.Value,
			"ttl":            float64(input.TTL),
			"cloud":          false,
			"upstream_https": input.UpstreamHTTPS,
			"ip_filter_mode": input.IPFilterMode,
			"is_protected":   false,
			"created_at":     "2023-03-31T08:57:51+00:00",
			"updated_at":     "2023-03-31T08:57:51+00:00",
		},
		Message: "DNS record created successfully",
	}

	_, err := client.CreateDNSRecord(context.Background(), ResourceDomain(""), CreateDNSRecordParams{})
	assert.ErrorIs(t, err, ErrMissingDomain)

	actual, err := client.CreateDNSRecord(context.Background(), ResourceDomain(testDomain), input)
	require.NoError(t, err)

	assert.Equal(t, want, actual)
}

func TestGetDNSRecord(t *testing.T) {
	setup()
	defer teardown()

	recordID := "714009ff-a43c-43c5-80e2-0b3ffc1344a4"

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)

		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, `{
			"data": {
				"id": "714009ff-a43c-43c5-80e2-0b3ffc1344a4",
				"type": "a",
				"name": "@",
				"value": [
					{
						"ip": "1.2.3.4",
						"port": 1,
						"weight": 10,
						"country": ""
					},
					{
						"ip": "5.6.7.8",
						"port": 2,
						"weight": 20,
						"country": ""
					}
				],
				"ttl": 120,
				"cloud": false,
				"upstream_https": "https",
				"ip_filter_mode": {
					"count": "multi",
					"order": "weighted",
					"geo_filter": "none"
				},
				"created_at": "2023-03-31T08:57:51+00:00",
				"updated_at": "2023-03-31T08:57:51+00:00"
			}
		}`)
	}

	mux.HandleFunc("/domains/"+testDomain+"/dns-records/"+recordID, handler)

	want := &DNSRecord{
		ID:   recordID,
		Type: "a",
		Name: "@",
		Value: []interface{}{
			map[string]interface{}{
				"ip":      "1.2.3.4",
				"port":    1.0,
				"weight":  10.0,
				"country": "",
			},
			map[string]interface{}{
				"ip":      "5.6.7.8",
				"port":    2.0,
				"weight":  20.0,
				"country": "",
			},
		},
		TTL:           120,
		Cloud:         false,
		UpstreamHTTPS: "https",
		IPFilterMode: map[string]interface{}{
			"count":      "multi",
			"order":      "weighted",
			"geo_filter": "none",
		},
		CreatedAt: "2023-03-31T08:57:51+00:00",
		UpdatedAt: "2023-03-31T08:57:51+00:00",
	}

	_, err := client.GetDNSRecord(context.Background(), ResourceDomain(""), recordID)
	assert.ErrorIs(t, err, ErrMissingDomain)

	_, err = client.GetDNSRecord(context.Background(), ResourceDomain(testDomain), "")
	assert.ErrorIs(t, err, ErrMissingDNSRecordID)

	actual, err := client.GetDNSRecord(context.Background(), ResourceDomain(testDomain), recordID)
	require.NoError(t, err)

	assert.Equal(t, want, actual)
}

func TestListDNSRecords(t *testing.T) {
	setup()
	defer teardown()

	input := ListDNSRecordsParams{
		Search: "b",
		Type:   "A",
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)
		assert.Equal(t, input.Search, r.URL.Query().Get("search"))
		assert.Equal(t, input.Type, r.URL.Query().Get("type"))

		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, `{
			"data": [
				{
					"id": "714009ff-a43c-43c5-80e2-0b3ffc1344a4",
					"type": "a",
					"name": "@",
					"value": [
						{
							"ip": "1.2.3.4",
							"port": 1,
							"weight": 10,
							"country": ""
						},
						{
							"ip": "5.6.7.8",
							"port": 2,
							"weight": 20,
							"country": ""
						}
					],
					"ttl": 180,
					"cloud": false,
					"upstream_https": "https",
					"ip_filter_mode": {
						"count": "multi",
						"order": "weighted",
						"geo_filter": "none"
					},
					"created_at": "2023-03-31T08:57:51+00:00",
					"updated_at": "2023-03-31T15:41:26+00:00"
				},
				{
					"id": "ee29f172-0b0d-4f2d-ba12-9587291936e3",
					"type": "a",
					"name": "b",
					"value": [
						{
							"ip": "3.1.8.1",
							"port": 80,
							"weight": 100,
							"country": ""
						},
						{
							"ip": "4.8.1.2",
							"port": 80,
							"weight": 100,
							"country": ""
						}
					],
					"ttl": 120,
					"cloud": false,
					"upstream_https": "https",
					"ip_filter_mode": {
						"count": "multi",
						"order": "weighted",
						"geo_filter": "none"
					},
					"created_at": "2023-03-31T11:40:15+00:00",
					"updated_at": "2023-03-31T11:40:15+00:00"
				}
			],
			"links": {
				"first": "https://napi.arvancloud.ir/4.0/domains/wkmag.ir/dns-records?page=1",
				"last": "https://napi.arvancloud.ir/4.0/domains/wkmag.ir/dns-records?page=1",
				"prev": null,
				"next": null
			},
			"meta": {
				"current_page": 1,
				"from": 1,
				"last_page": 1,
				"links": [
					{
						"url": null,
						"label": "Previous",
						"active": false
					},
					{
						"url": "https://napi.arvancloud.ir/4.0/domains/wkmag.ir/dns-records?page=1",
						"label": "1",
						"active": true
					},
					{
						"url": null,
						"label": "Next",
						"active": false
					}
				],
				"path": "https://napi.arvancloud.ir/4.0/domains/wkmag.ir/dns-records",
				"per_page": 300,
				"to": 2,
				"total": 2
			}
		}`)
	}

	mux.HandleFunc("/domains/"+testDomain+"/dns-records", handler)

	want := []DNSRecord{{
		ID:   "714009ff-a43c-43c5-80e2-0b3ffc1344a4",
		Type: "a",
		Name: "@",
		Value: []interface{}{
			map[string]interface{}{
				"ip":      "1.2.3.4",
				"port":    1.0,
				"weight":  10.0,
				"country": "",
			},
			map[string]interface{}{
				"ip":      "5.6.7.8",
				"port":    2.0,
				"weight":  20.0,
				"country": "",
			},
		},
		TTL:           180,
		Cloud:         false,
		UpstreamHTTPS: "https",
		IPFilterMode: map[string]interface{}{
			"count":      "multi",
			"order":      "weighted",
			"geo_filter": "none",
		},
		CreatedAt: "2023-03-31T08:57:51+00:00",
		UpdatedAt: "2023-03-31T15:41:26+00:00",
	},
		{
			ID:   "ee29f172-0b0d-4f2d-ba12-9587291936e3",
			Type: "a",
			Name: "b",
			Value: []interface{}{
				map[string]interface{}{
					"ip":      "3.1.8.1",
					"port":    80.0,
					"weight":  100.0,
					"country": "",
				},
				map[string]interface{}{
					"ip":      "4.8.1.2",
					"port":    80.0,
					"weight":  100.0,
					"country": "",
				},
			},
			TTL:           120,
			Cloud:         false,
			UpstreamHTTPS: "https",
			IPFilterMode: map[string]interface{}{
				"count":      "multi",
				"order":      "weighted",
				"geo_filter": "none",
			},
			CreatedAt: "2023-03-31T11:40:15+00:00",
			UpdatedAt: "2023-03-31T11:40:15+00:00",
		},
	}

	_, err := client.ListDNSRecords(context.Background(), ResourceDomain(""), ListDNSRecordsParams{})
	assert.ErrorIs(t, err, ErrMissingDomain)

	actual, err := client.ListDNSRecords(context.Background(), ResourceDomain(testDomain), input)
	require.NoError(t, err)

	assert.Equal(t, want, actual)
}

func TestUpdateDNSRecord(t *testing.T) {
	setup()
	defer teardown()

	recordID := "714009ff-a43c-43c5-80e2-0b3ffc1344a4"

	input := UpdateDNSRecordParams{
		Type: "a",
		Name: "@",
		Value: []interface{}{
			map[string]interface{}{
				"ip":      "1.2.3.5",
				"port":    2.0,
				"weight":  20.0,
				"country": "",
			},
			map[string]interface{}{
				"ip":      "5.6.7.9",
				"port":    3.0,
				"weight":  30.0,
				"country": "",
			},
		},
		TTL:           180,
		UpstreamHTTPS: "http",
		IPFilterMode: map[string]interface{}{
			"count":      "multi",
			"order":      "weighted",
			"geo_filter": "none",
		},
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPut, r.Method, "Expected method 'PUT', got %s", r.Method)

		var p UpdateDNSRecordParams
		err := json.NewDecoder(r.Body).Decode(&p)
		require.NoError(t, err)
		assert.Equal(t, input, p)

		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, `{
			"data": {
				"id": "714009ff-a43c-43c5-80e2-0b3ffc1344a4",
				"type": "a",
				"name": "@",
				"value": [
					{
						"ip": "1.2.3.5",
						"port": 2,
						"weight": 20,
						"country": ""
					},
					{
						"ip": "5.6.7.9",
						"port": 3,
						"weight": 30,
						"country": ""
					}
				],
				"ttl": 180,
				"cloud": false,
				"upstream_https": "http",
				"ip_filter_mode": {
					"count": "multi",
					"order": "weighted",
					"geo_filter": "none"
				},
				"is_protected": false,
				"created_at": "2023-03-31T08:57:51+00:00",
				"updated_at": "2023-03-31T15:41:26+00:00"
			},
			"message": "DNS record updated"
		}`)
	}

	mux.HandleFunc("/domains/"+testDomain+"/dns-records/"+recordID, handler)

	want := &UpdateDNSRecord_Response{
		Data: map[string]interface{}{
			"id":             recordID,
			"type":           input.Type,
			"name":           input.Name,
			"value":          input.Value,
			"ttl":            float64(input.TTL),
			"cloud":          false,
			"upstream_https": input.UpstreamHTTPS,
			"ip_filter_mode": input.IPFilterMode,
			"is_protected":   false,
			"created_at":     "2023-03-31T08:57:51+00:00",
			"updated_at":     "2023-03-31T15:41:26+00:00",
		},
		Message: "DNS record updated",
	}

	_, err := client.UpdateDNSRecord(context.Background(), ResourceDomain(""), recordID, UpdateDNSRecordParams{})
	assert.ErrorIs(t, err, ErrMissingDomain)

	_, err = client.UpdateDNSRecord(context.Background(), ResourceDomain(testDomain), "", UpdateDNSRecordParams{})
	assert.ErrorIs(t, err, ErrMissingDNSRecordID)

	actual, err := client.UpdateDNSRecord(context.Background(), ResourceDomain(testDomain), recordID, input)
	require.NoError(t, err)

	assert.Equal(t, want, actual)
}

func TestDeleteDNSRecord(t *testing.T) {
	setup()
	defer teardown()

	recordID := "714009ff-a43c-43c5-80e2-0b3ffc1344a4"

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodDelete, r.Method, "Expected method 'DELETE', got %s", r.Method)

		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, `{
			"message": "DNS record deleted"
		}`)
	}

	mux.HandleFunc("/domains/"+testDomain+"/dns-records/"+recordID, handler)

	want := &DeleteDNSRecord_Response{
		Message: "DNS record deleted",
	}

	_, err := client.DeleteDNSRecord(context.Background(), ResourceDomain(""), recordID)
	assert.ErrorIs(t, err, ErrMissingDomain)

	_, err = client.DeleteDNSRecord(context.Background(), ResourceDomain(testDomain), "")
	assert.ErrorIs(t, err, ErrMissingDNSRecordID)

	actual, err := client.DeleteDNSRecord(context.Background(), ResourceDomain(testDomain), recordID)
	require.NoError(t, err)

	assert.Equal(t, want, actual)
}
