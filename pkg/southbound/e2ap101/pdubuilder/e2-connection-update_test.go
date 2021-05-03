// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0
package pdubuilder

import (
	"gotest.tools/assert"
	"testing"
)

func TestE2connectionUpdate(t *testing.T) {
	newE2apPdu, err := CreateE2connectionUpdateE2apPdu()
	assert.NilError(t, err)
	assert.Assert(t, newE2apPdu != nil)

	//xer, err := asn1cgo.XerEncodeE2apPdu(newE2apPdu)
	//assert.NilError(t, err)
	//t.Logf("RicServiceQuery E2AP PDU XER\n%s", string(xer))
	//
	//per, err := asn1cgo.PerEncodeE2apPdu(newE2apPdu)
	//assert.NilError(t, err)
	//t.Logf("RicServiceQuery E2AP PDU PER\n%v", hex.Dump(per))
}
