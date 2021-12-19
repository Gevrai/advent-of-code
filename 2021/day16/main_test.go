// Copyright (c) Kaloom Inc., 2021
//
// This unpublished material is property of Kaloom Inc.
// All rights reserved.
// Reproduction or distribution, in whole or in part, is
// forbidden except by express written permission of Kaloom Inc.

package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewPacket(t *testing.T) {
	tests := []struct {
		packetHex  string
		sumVersion int
	}{
		{"D2FE28", 6},
		{"38006F45291200", 9},
		{"EE00D40C823060", 14},
		{"8A004A801A8002F478", 16},
		{"620080001611562C8802118E34", 12},
		{"C0015000016115A2E0802F182340", 23},
		{"A0016C880162017C3686B18A3D4780", 31},
	}
	for _, tt := range tests {
		t.Run(tt.packetHex, func(t *testing.T) {
			require.NotPanics(t, func() {
				stream := NewStreamFromHexadecimal(tt.packetHex)
				packet := NewPacket(stream)
				require.Equal(t, tt.sumVersion, packet.SumPacketVersions())
			})
		})
	}
}
