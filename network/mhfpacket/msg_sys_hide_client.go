package mhfpacket

import (
	"errors"

	"erupe-ce/common/byteframe"
	"erupe-ce/network"
	"erupe-ce/network/clientctx"
)

// MsgSysHideClient represents the MSG_SYS_HIDE_CLIENT
type MsgSysHideClient struct {
	Hide bool
}

// Opcode returns the ID associated with this packet type.
func (m *MsgSysHideClient) Opcode() network.PacketID {
	return network.MSG_SYS_HIDE_CLIENT
}

// Parse parses the packet from binary
func (m *MsgSysHideClient) Parse(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	m.Hide = bf.ReadBool()
	bf.ReadUint8() // Zeroed
	bf.ReadUint8() // Zeroed
	bf.ReadUint8() // Zeroed
	return nil
}

// Build builds a binary packet from the current data.
func (m *MsgSysHideClient) Build(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("NOT IMPLEMENTED")
}
