// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0
package pdubuilder

import (
	"fmt"
	"github.com/onosproject/onos-e2t/api/e2ap/v1beta2"
	e2ap_commondatatypes "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-commondatatypes"
	e2ap_ies "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-ies"
	e2appducontents "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-pdu-contents"
	e2appdudescriptions "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-pdu-descriptions"
)

func CreateE2connectionUpdateAcknowledgeE2apPdu() (*e2appdudescriptions.E2ApPdu, error) {

	connectionSetup := e2appducontents.E2ConnectionUpdateAckIes_E2ConnectionUpdateAckIes39{
		Id:          int32(v1beta2.ProtocolIeIDE2connectionSetup),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		ConnectionSetup: &e2appducontents.E2ConnectionUpdateList{
			Value: make([]*e2appducontents.E2ConnectionUpdateItemIes, 0),
		},
		Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_OPTIONAL),
	}

	csi := &e2appducontents.E2ConnectionUpdateItemIes{
		Id:          int32(v1beta2.ProtocolIeIDE2connectionUpdateItem),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value: &e2appducontents.E2ConnectionUpdateItem{
			TnlInformation: &e2ap_ies.Tnlinformation{
				TnlPort: &e2ap_commondatatypes.BitString{ //ToDo - pass as a parameter
					Value: 0x89bcd,
					Len:   16,
				},
				TnlAddress: &e2ap_commondatatypes.BitString{ //ToDo - pass as a parameter
					Value: 0x89bcd,
					Len:   64,
				},
			},
		},
		Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_OPTIONAL),
	}
	connectionSetup.ConnectionSetup.Value = append(connectionSetup.ConnectionSetup.Value, csi)

	connectionSetupFailed := e2appducontents.E2ConnectionUpdateAckIes_E2ConnectionUpdateAckIes40{
		Id:          int32(v1beta2.ProtocolIeIDE2connectionSetupFailed),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		ConnectionSetupFailed: &e2appducontents.E2ConnectionSetupFailedList{
			Value: make([]*e2appducontents.E2ConnectionSetupFailedItemIes, 0),
		},
		Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_OPTIONAL),
	}

	csfi := &e2appducontents.E2ConnectionSetupFailedItemIes{
		Id:          int32(v1beta2.ProtocolIeIDE2connectionSetupFailedItem),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE),
		Value: &e2appducontents.E2ConnectionSetupFailedItem{
			TnlInformation: &e2ap_ies.Tnlinformation{
				TnlPort: &e2ap_commondatatypes.BitString{ //ToDo - pass as a parameter
					Value: 0x89bcd,
					Len:   16,
				},
				TnlAddress: &e2ap_commondatatypes.BitString{ //ToDo - pass as a parameter
					Value: 0x89bcd,
					Len:   64,
				},
			},
			Cause: &e2ap_ies.Cause{
				Cause: &e2ap_ies.Cause_Protocol{
					Protocol: e2ap_ies.CauseProtocol_CAUSE_PROTOCOL_TRANSFER_SYNTAX_ERROR,
				},
			},
		},
		Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
	}
	connectionSetupFailed.ConnectionSetupFailed.Value = append(connectionSetupFailed.ConnectionSetupFailed.Value, csfi)

	e2apPdu := e2appdudescriptions.E2ApPdu{
		E2ApPdu: &e2appdudescriptions.E2ApPdu_SuccessfulOutcome{
			SuccessfulOutcome: &e2appdudescriptions.SuccessfulOutcome{
				ProcedureCode: &e2appdudescriptions.E2ApElementaryProcedures{
					E2ConnectionUpdate: &e2appdudescriptions.E2ConnectionUpdateEp{
						SuccessfulOutcome: &e2appducontents.E2ConnectionUpdateAcknowledge{
							ProtocolIes: &e2appducontents.E2ConnectionUpdateAckIes{
								E2ApProtocolIes39: &connectionSetup,       //E2 Connection Setup List
								E2ApProtocolIes40: &connectionSetupFailed, //E2 Connection Setup Failed List
							},
						},
					},
				},
			},
		},
	}
	if err := e2apPdu.Validate(); err != nil {
		return nil, fmt.Errorf("error validating E2ApPDU %s", err.Error())
	}
	return &e2apPdu, nil
}
