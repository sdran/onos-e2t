// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0
package pdubuilder

import (
	"github.com/onosproject/onos-e2t/api/e2ap/v1beta1/e2apies"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/asn1cgo"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/types"
	"gotest.tools/assert"
	"testing"
)

func TestRicIndication(t *testing.T) {
	ricRequestID := types.RicRequest{
		RequestorID: 21,
		InstanceID:  22,
	}
	var ranFuncID types.RanFunctionID = 9
	var ricAction = e2apies.RicactionType_RICACTION_TYPE_POLICY
	var ricIndicationType = e2apies.RicindicationType_RICINDICATION_TYPE_INSERT
	var ricSn types.RicIndicationSn = 1
	var ricIndHd types.RicIndicationHeader = []byte("123")
	var ricIndMsg types.RicIndicationMessage = []byte("456")
	var ricCallPrID types.RicCallProcessID = []byte("789")
	newE2apPdu, err := RicIndicationE2apPdu(ricRequestID,
		ranFuncID, ricAction, ricSn, ricIndicationType, ricIndHd, ricIndMsg, ricCallPrID)
	assert.NilError(t, err)
	assert.Assert(t, newE2apPdu != nil)

	xer, err := asn1cgo.XerEncodeE2apPdu(newE2apPdu)
	assert.NilError(t, err)
	t.Logf("RIC Indication XER\n%s", string(xer))
}
