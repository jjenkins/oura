/*
Oura API Documentation

Testing WebhookSubscriptionRoutesAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package oura

import (
	"context"
	"testing"

	openapiclient "github.com/jjenkins/oura"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_oura_WebhookSubscriptionRoutesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test WebhookSubscriptionRoutesAPIService CreateWebhookSubscriptionV2WebhookSubscriptionPost", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.WebhookSubscriptionRoutesAPI.CreateWebhookSubscriptionV2WebhookSubscriptionPost(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WebhookSubscriptionRoutesAPIService DeleteWebhookSubscriptionV2WebhookSubscriptionIdDelete", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id string

		httpRes, err := apiClient.WebhookSubscriptionRoutesAPI.DeleteWebhookSubscriptionV2WebhookSubscriptionIdDelete(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WebhookSubscriptionRoutesAPIService GetWebhookSubscriptionV2WebhookSubscriptionIdGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id string

		resp, httpRes, err := apiClient.WebhookSubscriptionRoutesAPI.GetWebhookSubscriptionV2WebhookSubscriptionIdGet(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WebhookSubscriptionRoutesAPIService ListWebhookSubscriptionsV2WebhookSubscriptionGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.WebhookSubscriptionRoutesAPI.ListWebhookSubscriptionsV2WebhookSubscriptionGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WebhookSubscriptionRoutesAPIService RenewWebhookSubscriptionV2WebhookSubscriptionRenewIdPut", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id string

		resp, httpRes, err := apiClient.WebhookSubscriptionRoutesAPI.RenewWebhookSubscriptionV2WebhookSubscriptionRenewIdPut(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WebhookSubscriptionRoutesAPIService UpdateWebhookSubscriptionV2WebhookSubscriptionIdPut", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id string

		resp, httpRes, err := apiClient.WebhookSubscriptionRoutesAPI.UpdateWebhookSubscriptionV2WebhookSubscriptionIdPut(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
