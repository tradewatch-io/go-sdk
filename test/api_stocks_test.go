/*
tradewatch.io

Testing StocksAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package tradewatch

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/tradewatch-io/go-sdk"
)

func Test_tradewatch_StocksAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test StocksAPIService getCountries", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.StocksAPI.getCountries(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test StocksAPIService getQuote", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var symbol string

		resp, httpRes, err := apiClient.StocksAPI.getQuote(context.Background(), symbol).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test StocksAPIService getSymbols", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.StocksAPI.getSymbols(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}