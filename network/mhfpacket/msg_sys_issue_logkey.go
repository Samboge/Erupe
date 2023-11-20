package mhfpacket

import (
	"errors"

	"erupe-ce/common/byteframe"
	"erupe-ce/network"
	"erupe-ce/network/clientctx"
)

// MsgSysIssueLogkey represents the MSG_SYS_ISSUE_LOGKEY
type MsgSysIssueLogkey struct {
	AckHandle uint32
	Unk0      uint16
}

// Opcode returns the ID associated with this packet type.
func (m *MsgSysIssueLogkey) Opcode() network.PacketID {
	return network.MSG_SYS_ISSUE_LOGKEY
}

// Parse parses the packet from binary
func (m *MsgSysIssueLogkey) Parse(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	m.AckHandle = bf.ReadUint32()
	m.Unk0 = bf.ReadUint16()
	bf.ReadUint16() // Zeroed
	return nil
}

// Build builds a binary packet from the current data.
func (m *MsgSysIssueLogkey) Build(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("NOT IMPLEMENTED")
}
