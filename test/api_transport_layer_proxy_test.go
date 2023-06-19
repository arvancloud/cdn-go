/*
ArvanCloud CDN Services

Testing TransportLayerProxyApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package r1cdn

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/arvancloud/cdn-go"
)

func Test_r1cdn_TransportLayerProxyApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test TransportLayerProxyApiService TransportLayerProxiesDestroy", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var domain string
		var transportLayerProxyId string

		resp, httpRes, err := apiClient.TransportLayerProxyApi.TransportLayerProxiesDestroy(context.Background(), domain, transportLayerProxyId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TransportLayerProxyApiService TransportLayerProxiesIndex", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var domain string

		resp, httpRes, err := apiClient.TransportLayerProxyApi.TransportLayerProxiesIndex(context.Background(), domain).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TransportLayerProxyApiService TransportLayerProxiesShow", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var domain string
		var transportLayerProxyId string

		resp, httpRes, err := apiClient.TransportLayerProxyApi.TransportLayerProxiesShow(context.Background(), domain, transportLayerProxyId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TransportLayerProxyApiService TransportLayerProxiesStore", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var domain string

		resp, httpRes, err := apiClient.TransportLayerProxyApi.TransportLayerProxiesStore(context.Background(), domain).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TransportLayerProxyApiService TransportLayerProxiesUpdate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var domain string
		var transportLayerProxyId string

		resp, httpRes, err := apiClient.TransportLayerProxyApi.TransportLayerProxiesUpdate(context.Background(), domain, transportLayerProxyId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}