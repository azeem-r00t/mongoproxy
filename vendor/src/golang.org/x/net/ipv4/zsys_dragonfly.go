// Created by cgo -godefs - DO NOT EDIT
// cgo -godefs defs_dragonfly.go

// +build dragonfly

package ipv4

const (
	sysIP_OPTIONS     = 0x1
	sysIP_HDRINCL     = 0x2
	sysIP_TOS         = 0x3
	sysIP_TTL         = 0x4
	sysIP_RECVOPTS    = 0x5
	sysIP_RECVRETOPTS = 0x6
	sysIP_RECVDSTADDR = 0x7
	sysIP_RETOPTS     = 0x8
	sysIP_RECVIF      = 0x14
	sysIP_RECVTTL     = 0x41

	sysIP_MULTICAST_IF    = 0x9
	sysIP_MULTICAST_TTL   = 0xa
	sysIP_MULTICAST_LOOP  = 0xb
	sysIP_MULTICAST_VIF   = 0xe
	sysIP_ADD_MEMBERSHIP  = 0xc
	sysIP_DROP_MEMBERSHIP = 0xd

	sysSizeofIPMreq = 0x8
)

type sysIPMreq struct {
	Multiaddr [4]byte /* in_addr */
	Interface [4]byte /* in_addr */
}
