// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package asn1cgo

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "ProtocolIE-Container.h"
//#include "ProtocolIE-Field.h"
import "C"
import (
	"unsafe"

	"github.com/onosproject/onos-e2t/api/e2ap/v1beta1/e2appducontents"
)

func newE2SetupRequestIes(esv *e2appducontents.E2SetupRequestIes) (*C.ProtocolIE_Container_1544P11_t, error) {
	pIeC1544P11 := new(C.ProtocolIE_Container_1544P11_t)

	if esv.GetE2ApProtocolIes3() != nil {
		ie3C, err := newE2setupRequestIe3GlobalE2NodeID(esv.GetE2ApProtocolIes3())
		if err != nil {
			return nil, err
		}
		if _, err = C.asn_sequence_add(unsafe.Pointer(pIeC1544P11), unsafe.Pointer(ie3C)); err != nil {
			return nil, err
		}
	}

	if esv.GetE2ApProtocolIes10() != nil {
		ie10C, err := newE2setupRequestIe10RanFunctionList(esv.GetE2ApProtocolIes10())
		if err != nil {
			return nil, err
		}
		if _, err = C.asn_sequence_add(unsafe.Pointer(pIeC1544P11), unsafe.Pointer(ie10C)); err != nil {
			return nil, err
		}
	}

	return pIeC1544P11, nil
}

func decodeE2SetupRequestIes(protocolIEsC *C.ProtocolIE_Container_1544P11_t) (*e2appducontents.E2SetupRequestIes, error) {
	pIEs := new(e2appducontents.E2SetupRequestIes)

	ieCount := int(protocolIEsC.list.count)
	//fmt.Printf("1544P11 Type %T Count %v Size %v\n", *protocolIEsC.list.array, protocolIEsC.list.count, protocolIEsC.list.size)
	for i := 0; i < ieCount; i++ {
		offset := unsafe.Sizeof(unsafe.Pointer(*protocolIEsC.list.array)) * uintptr(i) // Forget the rest - this works - 7Nov20
		e2srIeC := *(**C.E2setupRequestIEs_t)(unsafe.Pointer(uintptr(unsafe.Pointer(protocolIEsC.list.array)) + offset))

		ie, err := decodeE2setupRequestIE(e2srIeC)
		if err != nil {
			return nil, err
		}
		if ie.E2ApProtocolIes3 != nil {
			pIEs.E2ApProtocolIes3 = ie.E2ApProtocolIes3
		}
		if ie.E2ApProtocolIes10 != nil {
			pIEs.E2ApProtocolIes10 = ie.E2ApProtocolIes10
		}
	}

	return pIEs, nil
}

func newE2SetupResponseIes(e2srIEs *e2appducontents.E2SetupResponseIes) (*C.ProtocolIE_Container_1544P12_t, error) {
	pIeC1544P12 := new(C.ProtocolIE_Container_1544P12_t)

	if e2srIEs.GetE2ApProtocolIes4() != nil {
		ie4C, err := newE2setupResponseIe4GlobalRicID(e2srIEs.GetE2ApProtocolIes4())
		if err != nil {
			return nil, err
		}
		if _, err = C.asn_sequence_add(unsafe.Pointer(pIeC1544P12), unsafe.Pointer(ie4C)); err != nil {
			return nil, err
		}
	}

	if e2srIEs.GetE2ApProtocolIes9() != nil {
		ie9C, err := newE2setupResponseIe9RanFunctionsAccepted(e2srIEs.GetE2ApProtocolIes9())
		if err != nil {
			return nil, err
		}
		if _, err = C.asn_sequence_add(unsafe.Pointer(pIeC1544P12), unsafe.Pointer(ie9C)); err != nil {
			return nil, err
		}
	}

	if e2srIEs.GetE2ApProtocolIes13() != nil {
		ie13C, err := newE2setupResponseIe13RanFunctionsRejected(e2srIEs.GetE2ApProtocolIes13())
		if err != nil {
			return nil, err
		}
		if _, err = C.asn_sequence_add(unsafe.Pointer(pIeC1544P12), unsafe.Pointer(ie13C)); err != nil {
			return nil, err
		}
	}
	return pIeC1544P12, nil
}

func decodeE2SetupResponseIes(protocolIEsC *C.ProtocolIE_Container_1544P12_t) (*e2appducontents.E2SetupResponseIes, error) {
	pIEs := new(e2appducontents.E2SetupResponseIes)

	ieCount := int(protocolIEsC.list.count)
	//fmt.Printf("1544P11 Type %T Count %v Size %v\n", *protocolIEsC.list.array, protocolIEsC.list.count, protocolIEsC.list.size)
	for i := 0; i < ieCount; i++ {
		offset := unsafe.Sizeof(unsafe.Pointer(*protocolIEsC.list.array)) * uintptr(i)
		e2srIeC := *(**C.E2setupResponseIEs_t)(unsafe.Pointer(uintptr(unsafe.Pointer(protocolIEsC.list.array)) + offset))

		ie, err := decodeE2setupResponseIE(e2srIeC)
		if err != nil {
			return nil, err
		}
		if ie.E2ApProtocolIes4 != nil {
			pIEs.E2ApProtocolIes4 = ie.E2ApProtocolIes4
		}
		if ie.E2ApProtocolIes9 != nil {
			pIEs.E2ApProtocolIes9 = ie.E2ApProtocolIes9
		}
		if ie.E2ApProtocolIes13 != nil {
			pIEs.E2ApProtocolIes13 = ie.E2ApProtocolIes13
		}
	}

	return pIEs, nil
}

func newRicSubscriptionResponseIe(rsrIEs *e2appducontents.RicsubscriptionResponseIes) (*C.ProtocolIE_Container_1544P1_t, error) {
	pIeC1544P1 := new(C.ProtocolIE_Container_1544P1_t)

	if rsrIEs.GetE2ApProtocolIes5() != nil {
		ie5C, err := newRicSubscriptionResponseIe5RanFunctionID(rsrIEs.GetE2ApProtocolIes5())
		if err != nil {
			return nil, err
		}
		if _, err = C.asn_sequence_add(unsafe.Pointer(pIeC1544P1), unsafe.Pointer(ie5C)); err != nil {
			return nil, err
		}
	}
	if rsrIEs.GetE2ApProtocolIes17() != nil {
		ie17C, err := newRicSubscriptionResponseIe17RactionAdmittedList(rsrIEs.GetE2ApProtocolIes17())
		if err != nil {
			return nil, err
		}
		if _, err = C.asn_sequence_add(unsafe.Pointer(pIeC1544P1), unsafe.Pointer(ie17C)); err != nil {
			return nil, err
		}
	}
	// TODO: Comment back in when RICactionRejected is handled
	//if rsrIEs.GetE2ApProtocolIes18() != nil {
	//	ie18C, err := newE2setupResponseIe4GlobalRicID(rsrIEs.GetE2ApProtocolIes18())
	//	if err != nil {
	//		return nil, err
	//	}
	//	if _, err = C.asn_sequence_add(unsafe.Pointer(pIeC1544P1), unsafe.Pointer(ie18C)); err != nil {
	//		return nil, err
	//	}
	//}
	if rsrIEs.GetE2ApProtocolIes29() != nil {
		ie29C, err := newRicSubscriptionResponseIe29RicRequestID(rsrIEs.GetE2ApProtocolIes29())
		if err != nil {
			return nil, err
		}
		if _, err = C.asn_sequence_add(unsafe.Pointer(pIeC1544P1), unsafe.Pointer(ie29C)); err != nil {
			return nil, err
		}
	}
	return pIeC1544P1, nil
}

func decodeRicSubscriptionResponseIes(protocolIEsC *C.ProtocolIE_Container_1544P1_t) (*e2appducontents.RicsubscriptionResponseIes, error) {
	pIEs := new(e2appducontents.RicsubscriptionResponseIes)

	ieCount := int(protocolIEsC.list.count)
	//fmt.Printf("1544P1 Type %T Count %v Size %v\n", *protocolIEsC.list.array, protocolIEsC.list.count, protocolIEsC.list.size)
	for i := 0; i < ieCount; i++ {
		offset := unsafe.Sizeof(unsafe.Pointer(*protocolIEsC.list.array)) * uintptr(i)
		rsrIeC := *(**C.RICsubscriptionResponse_IEs_t)(unsafe.Pointer(uintptr(unsafe.Pointer(protocolIEsC.list.array)) + offset))

		ie, err := decodeRicSubscriptionResponseIE(rsrIeC)
		if err != nil {
			return nil, err
		}
		if ie.E2ApProtocolIes5 != nil {
			pIEs.E2ApProtocolIes5 = ie.E2ApProtocolIes5
		}
		if ie.E2ApProtocolIes17 != nil {
			pIEs.E2ApProtocolIes17 = ie.E2ApProtocolIes17
		}
		if ie.E2ApProtocolIes18 != nil {
			pIEs.E2ApProtocolIes18 = ie.E2ApProtocolIes18
		}
		if ie.E2ApProtocolIes29 != nil {
			pIEs.E2ApProtocolIes29 = ie.E2ApProtocolIes29
		}
	}

	return pIEs, nil
}

func newRicSubscriptionRequestIes(rsrIEs *e2appducontents.RicsubscriptionRequestIes) (*C.ProtocolIE_Container_1544P0_t, error) {
	pIeC1544P0 := new(C.ProtocolIE_Container_1544P0_t)

	if rsrIEs.GetE2ApProtocolIes5() != nil {
		ie5C, err := newRicSubscriptionRequestIe5RanFunctionID(rsrIEs.E2ApProtocolIes5)
		if err != nil {
			return nil, err
		}
		if _, err = C.asn_sequence_add(unsafe.Pointer(pIeC1544P0), unsafe.Pointer(ie5C)); err != nil {
			return nil, err
		}
	}

	if rsrIEs.GetE2ApProtocolIes29() != nil {
		ie29C, err := newRicSubscriptionRequestIe29RicRequestID(rsrIEs.E2ApProtocolIes29)
		if err != nil {
			return nil, err
		}
		if _, err = C.asn_sequence_add(unsafe.Pointer(pIeC1544P0), unsafe.Pointer(ie29C)); err != nil {
			return nil, err
		}
	}

	if rsrIEs.GetE2ApProtocolIes30() != nil {
		ie30C, err := newRicSubscriptionRequestIe30RicSubscriptionDetails(rsrIEs.E2ApProtocolIes30)
		if err != nil {
			return nil, err
		}
		if _, err = C.asn_sequence_add(unsafe.Pointer(pIeC1544P0), unsafe.Pointer(ie30C)); err != nil {
			return nil, err
		}
	}

	return pIeC1544P0, nil
}

func decodeRicSubscriptionRequestIes(protocolIEsC *C.ProtocolIE_Container_1544P0_t) (*e2appducontents.RicsubscriptionRequestIes, error) {
	pIEs := new(e2appducontents.RicsubscriptionRequestIes)

	ieCount := int(protocolIEsC.list.count)
	//	fmt.Printf("1544P0 Type %T Count %v Size %v\n", *protocolIEsC.list.array, protocolIEsC.list.count, protocolIEsC.list.size)
	for i := 0; i < ieCount; i++ {
		offset := unsafe.Sizeof(unsafe.Pointer(*protocolIEsC.list.array)) * uintptr(i)
		rsrIeC := *(**C.RICsubscriptionRequest_IEs_t)(unsafe.Pointer(uintptr(unsafe.Pointer(protocolIEsC.list.array)) + offset))

		ie, err := decodeRicSubscriptionRequestIE(rsrIeC)
		if err != nil {
			return nil, err
		}
		if ie.E2ApProtocolIes5 != nil {
			pIEs.E2ApProtocolIes5 = ie.E2ApProtocolIes5
		}
		if ie.E2ApProtocolIes29 != nil {
			pIEs.E2ApProtocolIes29 = ie.E2ApProtocolIes29
		}
		if ie.E2ApProtocolIes30 != nil {
			pIEs.E2ApProtocolIes30 = ie.E2ApProtocolIes30
		}
	}

	return pIEs, nil
}

func newRicIndicationIEs(riIes *e2appducontents.RicindicationIes) (*C.ProtocolIE_Container_1544P6_t, error) {
	pIeC1544P6 := new(C.ProtocolIE_Container_1544P6_t)

	if riIes.GetE2ApProtocolIes5() != nil {
		ie5c, err := newRicIndicationIe5RanFunctionID(riIes.GetE2ApProtocolIes5())
		if err != nil {
			return nil, err
		}
		if _, err = C.asn_sequence_add(unsafe.Pointer(pIeC1544P6), unsafe.Pointer(ie5c)); err != nil {
			return nil, err
		}
	}

	if riIes.GetE2ApProtocolIes15() != nil {
		ie15c, err := newRicIndicationIe15RicActionID(riIes.GetE2ApProtocolIes15())
		if err != nil {
			return nil, err
		}
		if _, err = C.asn_sequence_add(unsafe.Pointer(pIeC1544P6), unsafe.Pointer(ie15c)); err != nil {
			return nil, err
		}
	}

	if riIes.GetE2ApProtocolIes20() != nil {
		ie20c, err := newRicIndicationIe20RiccallProcessID(riIes.GetE2ApProtocolIes20())
		if err != nil {
			return nil, err
		}
		if _, err = C.asn_sequence_add(unsafe.Pointer(pIeC1544P6), unsafe.Pointer(ie20c)); err != nil {
			return nil, err
		}
	}

	if riIes.GetE2ApProtocolIes25() != nil {
		ie25c, err := newRicIndicationIe25RicIndicationHeader(riIes.GetE2ApProtocolIes25())
		if err != nil {
			return nil, err
		}
		if _, err = C.asn_sequence_add(unsafe.Pointer(pIeC1544P6), unsafe.Pointer(ie25c)); err != nil {
			return nil, err
		}
	}

	if riIes.GetE2ApProtocolIes26() != nil {
		ie26c, err := newRicIndicationIe26RicIndicationMessage(riIes.GetE2ApProtocolIes26())
		if err != nil {
			return nil, err
		}
		if _, err = C.asn_sequence_add(unsafe.Pointer(pIeC1544P6), unsafe.Pointer(ie26c)); err != nil {
			return nil, err
		}
	}

	if riIes.GetE2ApProtocolIes27() != nil {
		ie27c, err := newRicIndicationIe27RicIndicationSn(riIes.GetE2ApProtocolIes27())
		if err != nil {
			return nil, err
		}
		if _, err = C.asn_sequence_add(unsafe.Pointer(pIeC1544P6), unsafe.Pointer(ie27c)); err != nil {
			return nil, err
		}
	}

	if riIes.GetE2ApProtocolIes28() != nil {
		ie28c, err := newRicIndicationIe28RicIndicationType(riIes.GetE2ApProtocolIes28())
		if err != nil {
			return nil, err
		}
		if _, err = C.asn_sequence_add(unsafe.Pointer(pIeC1544P6), unsafe.Pointer(ie28c)); err != nil {
			return nil, err
		}
	}

	if riIes.GetE2ApProtocolIes29() != nil {
		ie29c, err := newRicIndicationIe29RicRequestID(riIes.GetE2ApProtocolIes29())
		if err != nil {
			return nil, err
		}
		if _, err = C.asn_sequence_add(unsafe.Pointer(pIeC1544P6), unsafe.Pointer(ie29c)); err != nil {
			return nil, err
		}
	}

	return pIeC1544P6, nil
}

func decodeRicIndicationIes(protocolIEsC *C.ProtocolIE_Container_1544P6_t) (*e2appducontents.RicindicationIes, error) {
	pIEs := new(e2appducontents.RicindicationIes)

	ieCount := int(protocolIEsC.list.count)
	//fmt.Printf("1544P6 Type %T Count %v Size %v\n", *protocolIEsC.list.array, protocolIEsC.list.count, protocolIEsC.list.size)
	for i := 0; i < ieCount; i++ {
		offset := unsafe.Sizeof(unsafe.Pointer(*protocolIEsC.list.array)) * uintptr(i) // Forget the rest - this works - 7Nov20
		riIeC := *(**C.RICindication_IEs_t)(unsafe.Pointer(uintptr(unsafe.Pointer(protocolIEsC.list.array)) + offset))

		ie, err := decodeRicIndicationIE(riIeC)
		if err != nil {
			return nil, err
		}
		if ie.E2ApProtocolIes5 != nil {
			pIEs.E2ApProtocolIes5 = ie.E2ApProtocolIes5
		}
		if ie.E2ApProtocolIes15 != nil {
			pIEs.E2ApProtocolIes15 = ie.E2ApProtocolIes15
		}
		if ie.E2ApProtocolIes20 != nil {
			pIEs.E2ApProtocolIes20 = ie.E2ApProtocolIes20
		}
		if ie.E2ApProtocolIes25 != nil {
			pIEs.E2ApProtocolIes25 = ie.E2ApProtocolIes25
		}
		if ie.E2ApProtocolIes26 != nil {
			pIEs.E2ApProtocolIes26 = ie.E2ApProtocolIes26
		}
		if ie.E2ApProtocolIes27 != nil {
			pIEs.E2ApProtocolIes27 = ie.E2ApProtocolIes27
		}
		if ie.E2ApProtocolIes28 != nil {
			pIEs.E2ApProtocolIes28 = ie.E2ApProtocolIes28
		}
		if ie.E2ApProtocolIes29 != nil {
			pIEs.E2ApProtocolIes29 = ie.E2ApProtocolIes29
		}
	}

	return pIEs, nil
}
