// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package e2

import (
	"context"
	"testing"
	"time"

	subapi "github.com/onosproject/onos-api/go/onos/e2sub/subscription"
	"github.com/onosproject/onos-e2t/test/utils"
	e2client "github.com/onosproject/onos-ric-sdk-go/pkg/e2"
	"github.com/onosproject/onos-ric-sdk-go/pkg/e2/indication"
	"github.com/stretchr/testify/assert"
)

// TestE2NodeDownSubscription checks that a subscription channel read times out if
// the e2 node is down.
func (s *TestSuite) TestE2NodeDownSubscription(t *testing.T) {

	// Create a simulator
	sim := utils.CreateRanSimulatorWithName(t, "ran-simulator")
	assert.NotNil(t, sim)

	// Create an e2client
	clientConfig := e2client.Config{
		AppID: "subscription-e2node-down-test",
		SubscriptionService: e2client.ServiceConfig{
			Host: utils.SubscriptionServiceHost,
			Port: utils.SubscriptionServicePort,
		},
	}
	client, err := e2client.NewClient(clientConfig)
	assert.NoError(t, err)

	ch := make(chan indication.Indication)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	nodeIDs, err := utils.GetNodeIDs()
	assert.NoError(t, err)

	eventTriggerBytes, err := utils.CreateKpmEventTrigger(12)
	assert.NoError(t, err)

	subRequest := utils.Subscription{
		NodeID:               nodeIDs[0],
		EncodingType:         subapi.Encoding_ENCODING_PROTO,
		ActionType:           subapi.ActionType_ACTION_TYPE_REPORT,
		EventTrigger:         eventTriggerBytes,
		ServiceModelID:       utils.KpmServiceModelID,
		ActionID:             100,
		SubSequentActionType: subapi.SubsequentActionType_SUBSEQUENT_ACTION_TYPE_CONTINUE,
		TimeToWait:           subapi.TimeToWait_TIME_TO_WAIT_ZERO,
	}

	// Create a subscription request to indication messages from the client
	subReq, err := subRequest.Create()
	assert.NoError(t, err)

	// Cause the simulator to crash
	err = sim.Uninstall()
	assert.NoError(t, err)

	//  Create the subscription
	_, err = client.Subscribe(ctx, subReq, ch)
	assert.NoError(t, err)

	// Make sure that reads on the subscription channel time out. There should be no
	// indication messages available
	var gotIndication bool
	select {
	case indicationMsg := <-ch:
		// We got an indication. This is an error, as there is no E2 node to send one
		gotIndication = true
		t.Log(indicationMsg)

	case <-time.After(10 * time.Second):
		// The read timed out. This is the expected behavior.
		gotIndication = false
	}

	assert.False(t, gotIndication, "Indication message was delivered for a node that is down")
}