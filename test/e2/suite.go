// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package e2

import (
	"github.com/onosproject/helmit/pkg/test"
	"github.com/onosproject/onos-e2t/test/utils"
)

// TestSuite is the primary onos-e2t test suite
type TestSuite struct {
	test.Suite
}

// SetupTestSuite sets up the onos-e2t test suite
func (s *TestSuite) SetupTestSuite() error {
	sdran, err := utils.CreateSdranRelease()
	if err != nil {
		return err
	}
	return sdran.Install(true)
}
