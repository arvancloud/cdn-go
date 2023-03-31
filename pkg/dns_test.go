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

	handler := func(writer http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method, "Expected method 'POST', got %s", r.Method)

		var v CreateDNSRecordParams
		err := json.NewDecoder(r.Body).Decode(&v)
		require.NoError(t, err)
		assert.Equal(t, input, v)

		writer.Header().Set("content-type", "application/json")
		fmt.Fprint(writer, `{
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

		var v UpdateDNSRecordParams
		err := json.NewDecoder(r.Body).Decode(&v)
		require.NoError(t, err)
		assert.Equal(t, input, v)

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
