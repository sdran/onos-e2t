// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package asn1cgo

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "RANfunctionRevision.h"
import "C"
import (
	"github.com/onosproject/onos-e2t/api/e2ap/v1beta1/e2apies"
)

func newRanFunctionRevision(revision *e2apies.RanfunctionRevision) C.long {
	return C.long(revision.Value)
}

func decodeRanFunctionRevision(ranFunctionRevisionC *C.RANfunctionRevision_t) *e2apies.RanfunctionRevision {
	result := e2apies.RanfunctionRevision{
		Value: int32(*ranFunctionRevisionC),
	}

	return &result
}
