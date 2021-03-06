// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package e2

import (
	"context"
	"testing"
	"time"

	subtaskapi "github.com/onosproject/onos-api/go/onos/e2sub/task"

	subapi "github.com/onosproject/onos-api/go/onos/e2sub/subscription"
	"github.com/onosproject/onos-e2t/test/utils"
	"github.com/onosproject/onos-ric-sdk-go/pkg/e2/indication"
	"github.com/stretchr/testify/assert"
)

type invalidSubscriptionTestCase struct {
	description         string
	encodingType        subapi.Encoding
	actionType          subapi.ActionType
	serviceModelName    subapi.ServiceModelName
	serviceModelVersion subapi.ServiceModelVersion
	eventTrigger        []byte
	actionID            int32
	expectedError       subtaskapi.Cause
	enabled             bool
}

func runTestCase(t *testing.T, testCase invalidSubscriptionTestCase) {
	if !testCase.enabled {
		t.Skip()
		return
	}
	e2Client := getE2Client(t, "invalid-subscriptions-id")

	ch := make(chan indication.Indication)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	nodeIDs, err := utils.GetNodeIDs()
	assert.NoError(t, err)

	subRequest := utils.Subscription{
		NodeID:               nodeIDs[0],
		EncodingType:         testCase.encodingType,
		ActionType:           testCase.actionType,
		EventTrigger:         testCase.eventTrigger,
		ServiceModelName:     testCase.serviceModelName,
		ServiceModelVersion:  testCase.serviceModelVersion,
		ActionID:             testCase.actionID,
		SubSequentActionType: subapi.SubsequentActionType_SUBSEQUENT_ACTION_TYPE_CONTINUE,
		TimeToWait:           subapi.TimeToWait_TIME_TO_WAIT_ZERO,
	}

	subReq, err := subRequest.Create()
	assert.NoError(t, err)

	sub, err := e2Client.Subscribe(ctx, subReq, ch)
	assert.NoError(t, err)

	select {
	case err = <-sub.Err():
		assert.Equal(t, testCase.expectedError.String(), err.Error())
	case <-time.After(10 * time.Second):
		t.Fatal("test is failed because of timeout")

	}
}

// TestInvalidSubscriptions tests invalid inputs into the SDK
func (s *TestSuite) TestInvalidSubscriptions(t *testing.T) {
	const actionID = 11
	eventTriggerBytes, err := utils.CreateKpmV1EventTrigger(12)
	assert.NoError(t, err)

	testCases := []invalidSubscriptionTestCase{
		{
			description:         "Non-existent Service Model ID",
			enabled:             true,
			encodingType:        subapi.Encoding_ENCODING_PROTO,
			actionType:          subapi.ActionType_ACTION_TYPE_REPORT,
			serviceModelName:    "no-such-service-model",
			serviceModelVersion: "v1",
			eventTrigger:        eventTriggerBytes,
			actionID:            actionID,
			expectedError:       subtaskapi.Cause_CAUSE_RIC_RAN_FUNCTION_ID_INVALID,
		},
		{
			description:         "Invalid action type",
			enabled:             true,
			encodingType:        subapi.Encoding_ENCODING_PROTO,
			actionType:          subapi.ActionType_ACTION_TYPE_INSERT,
			serviceModelName:    utils.KpmServiceModelName,
			serviceModelVersion: utils.KpmServiceModelVersion1,
			eventTrigger:        eventTriggerBytes,
			actionID:            actionID,
			expectedError:       subtaskapi.Cause_CAUSE_RIC_ACTION_NOT_SUPPORTED,
		},
		{
			description:         "Invalid encoding type",
			enabled:             true,
			encodingType:        77,
			actionType:          subapi.ActionType_ACTION_TYPE_REPORT,
			serviceModelName:    utils.KpmServiceModelName,
			serviceModelVersion: utils.KpmServiceModelVersion1,
			eventTrigger:        eventTriggerBytes,
			actionID:            actionID,
			expectedError:       subtaskapi.Cause_CAUSE_PROTOCOL_MESSAGE_NOT_COMPATIBLE_WITH_RECEIVER_STATE,
		},
		{
			description:         "Invalid action ID",
			enabled:             true,
			encodingType:        subapi.Encoding_ENCODING_PROTO,
			actionType:          subapi.ActionType_ACTION_TYPE_REPORT,
			serviceModelName:    utils.KpmServiceModelName,
			serviceModelVersion: utils.KpmServiceModelVersion1,
			eventTrigger:        eventTriggerBytes,
			actionID:            100000,
			expectedError:       subtaskapi.Cause_CAUSE_PROTOCOL_ABSTRACT_SYNTAX_ERROR_FALSELY_CONSTRUCTED_MESSAGE,
		},
		{
			description:         "Invalid event trigger",
			enabled:             true,
			encodingType:        subapi.Encoding_ENCODING_PROTO,
			actionType:          subapi.ActionType_ACTION_TYPE_REPORT,
			serviceModelName:    utils.KpmServiceModelName,
			serviceModelVersion: utils.KpmServiceModelVersion1,
			eventTrigger:        make([]byte, 50),
			actionID:            actionID,
			expectedError:       subtaskapi.Cause_CAUSE_PROTOCOL_ABSTRACT_SYNTAX_ERROR_FALSELY_CONSTRUCTED_MESSAGE,
		},
	}

	sim := utils.CreateRanSimulatorWithNameOrDie(t, "invalid-subscriptions")

	for _, testCase := range testCases {
		pinTestCase := testCase
		t.Run(pinTestCase.description, func(t *testing.T) {
			runTestCase(t, pinTestCase)
		})
	}
	err = sim.Uninstall()
	assert.NoError(t, err)
}
