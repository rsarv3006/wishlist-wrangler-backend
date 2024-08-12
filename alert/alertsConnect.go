package alert

import (
	"wishlist-wrangler-api/config"

	"github.com/apialerts/apialerts-go"
)

func Connect() *apialerts.Client {

	apiKey := config.Config("API_ALERTS_KEY")
	client := apialerts.ApiAlertsClient()
	client.SetApiKey(apiKey)

	return client
}
