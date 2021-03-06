// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package e2

import (
	"context"
	"testing"
	"time"

	"github.com/onosproject/onos-ric-sdk-go/pkg/app"

	"github.com/onosproject/onos-api/go/onos/e2sub/subscription"
	subapi "github.com/onosproject/onos-api/go/onos/e2sub/subscription"
	e2client "github.com/onosproject/onos-ric-sdk-go/pkg/e2"
	"github.com/onosproject/onos-ric-sdk-go/pkg/e2/indication"
	sdksub "github.com/onosproject/onos-ric-sdk-go/pkg/e2/subscription"
	"github.com/stretchr/testify/assert"

	"github.com/onosproject/onos-e2t/test/utils"
)

const (
	defaultIndicationTimeout = 10 * time.Second
)

// getE2Client gets an E2 client
func getE2Client(t *testing.T, appID string) e2client.Client {
	clientConfig := e2client.Config{
		AppID: app.ID(appID),
		E2TService: e2client.ServiceConfig{
			Host: utils.E2TServiceHost,
			Port: utils.E2TServicePort,
		},
		SubscriptionService: e2client.ServiceConfig{
			Host: utils.SubscriptionServiceHost,
			Port: utils.SubscriptionServicePort,
		},
	}
	client, err := e2client.NewClient(clientConfig)
	assert.NoError(t, err)
	assert.NotNil(t, client)
	return client

}

// getSubClient returns an SDK subscription client
func getSubClient(t *testing.T) sdksub.Client {
	conn, err := utils.ConnectSubscriptionServiceHost()
	assert.NoError(t, err)
	assert.NotNil(t, conn)

	return sdksub.NewClient(conn)
}

// getSubscriptionList get  list of subscriptions
func getSubscriptionList(t *testing.T) []subscription.Subscription {
	subClient := getSubClient(t)
	subList, err := subClient.List(context.Background())
	assert.NoError(t, err)
	return subList
}

// checkSubscriptionGet make sure that the given subscription ID can be fetched via the subscription API
func checkSubscriptionGet(t *testing.T, expectedID subapi.ID) {
	subClient := getSubClient(t)

	fetched, err := subClient.Get(context.Background(), expectedID)
	assert.NoError(t, err)
	assert.NotNil(t, fetched)

	assert.Equal(t, expectedID, fetched.ID)
}

// checkSubscriptionIDInList makes sure that the give subscription ID appears once and only once in the subscription list
func checkSubscriptionIDInList(t *testing.T, expectedID subapi.ID, subList []subscription.Subscription) {
	found := 0
	for _, sub := range subList {
		if sub.ID == expectedID {
			found++
		}
	}
	assert.Equal(t, 1, found, "Subscription %s not found in subscription list", expectedID)
}

func checkIndicationMessage(t *testing.T, timeout time.Duration, ch chan indication.Indication) indication.Indication {
	select {
	case indicationMsg := <-ch:
		t.Log(indicationMsg)
		return indicationMsg
	case <-time.After(timeout):
		t.Fatal("failed to receive indication message")

	}
	return indication.Indication{}
}
